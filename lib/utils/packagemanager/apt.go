// Teleport
// Copyright (C) 2024 Gravitational, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package packagemanager

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/gravitational/trace"

	"github.com/gravitational/teleport"
	"github.com/gravitational/teleport/lib/linux"
	"github.com/gravitational/teleport/lib/utils"
)

const (
	aptFilePermsRepository = 0o644
)

type aptRepoMetadata struct {
	endpoint       string
	pubKeyEndpoint string
	sourceListFile string
	pubKeyFile     string
}

// APT is a wrapper for apt package manager.
// This package manager is used in Debian/Ubuntu and distros based on this distribution.
type APT struct {
	*APTConfig

	// legacy indicates that the old method of adding repos must be used.
	// This is used in Xenial (16.04) and Trusty (14.04) Ubuntu releases.
	legacy bool
}

// APTConfig contains the configurable fields for setting up the APT package manager.
type APTConfig struct {
	logger       *slog.Logger
	fsRootPrefix string
	bins         BinariesLocation
	httpGet      func(url string) (resp *http.Response, err error)
}

// CheckAndSetDefaults checks and sets default config values.
func (p *APTConfig) CheckAndSetDefaults() error {
	if p == nil {
		return trace.BadParameter("config is required")
	}

	p.bins.CheckAndSetDefaults()

	if p.fsRootPrefix == "" {
		p.fsRootPrefix = "/"
	}

	if p.logger == nil {
		p.logger = slog.Default()
	}

	if p.httpGet == nil {
		return trace.BadParameter("httpGet is required")
	}

	return nil
}

// NewAPT creates a new APT package manager.
func NewAPT(cfg *APTConfig) (*APT, error) {
	if err := cfg.CheckAndSetDefaults(); err != nil {
		return nil, trace.Wrap(err)
	}
	return &APT{APTConfig: cfg}, nil
}

// NewAPTLegacy creates a new APT for legacy ubuntu versions (Xenial and Trusty).
func NewAPTLegacy(cfg *APTConfig) (*APT, error) {
	pm, err := NewAPT(cfg)
	if err != nil {
		return nil, trace.Wrap(err)
	}
	pm.legacy = true
	pm.logger = pm.logger.With("legacy", "true")
	return pm, nil
}

// AddTeleportRepository adds the Teleport repository to the current system.
func (pm *APT) AddTeleportRepository(ctx context.Context, linuxInfo *linux.OSRelease, repoChannel string, developmentRepo bool) error {
	repoMetadata := teleportAPTRepo(developmentRepo)

	aptRepoEndpoint := repoMetadata.endpoint
	pm.logger.InfoContext(ctx, "Fetching Teleport repository key", "endpoint", repoMetadata.pubKeyEndpoint)

	resp, err := pm.httpGet(repoMetadata.pubKeyEndpoint)
	if err != nil {
		return trace.Wrap(err)
	}
	defer resp.Body.Close()
	publicKey, err := utils.ReadAtMost(resp.Body, teleport.MaxHTTPResponseSize)
	if err != nil {
		return trace.Wrap(err)
	}

	aptTeleportSourceListFile := filepath.Join(pm.fsRootPrefix, repoMetadata.sourceListFile)
	aptTeleportPublicKeyFile := filepath.Join(pm.fsRootPrefix, repoMetadata.pubKeyFile)
	// Format for teleport repo entry should look like this:
	// deb [signed-by=/usr/share/keyrings/teleport-archive-keyring.asc]  https://apt.releases.teleport.dev/${ID?} ${VERSION_CODENAME?} $RepoChannel"
	// Or (for non-production releases like alpha.X)
	// deb [signed-by=/usr/share/keyrings/teleport-development-archive-keyring.asc]  https://apt.releases.development.teleport.dev/${ID?} ${VERSION_CODENAME?} $RepoChannel"
	teleportRepoMetadata := fmt.Sprintf("deb [signed-by=%s] %s%s %s %s", aptTeleportPublicKeyFile, aptRepoEndpoint, linuxInfo.ID, linuxInfo.VersionCodename, repoChannel)

	switch {
	case pm.legacy:
		pm.logger.InfoContext(ctx, "Trusting Teleport repository key", "command", "apt-key add -")
		aptKeyAddCMD := exec.CommandContext(ctx, pm.bins.AptKey, "add", "-")
		aptKeyAddCMD.Stdin = bytes.NewReader(publicKey)
		aptKeyAddCMDOutput, err := aptKeyAddCMD.CombinedOutput()
		if err != nil {
			return trace.Wrap(err, string(aptKeyAddCMDOutput))
		}
		teleportRepoMetadata = fmt.Sprintf("deb %s %s %s", aptRepoEndpoint, linuxInfo.VersionCodename, repoChannel)

	default:
		pm.logger.InfoContext(ctx, "Writing Teleport repository key", "destination", aptTeleportPublicKeyFile)
		if err := os.WriteFile(aptTeleportPublicKeyFile, publicKey, aptFilePermsRepository); err != nil {
			return trace.Wrap(err)
		}
	}

	pm.logger.InfoContext(ctx, "Adding repository metadata", "apt_source_file", aptTeleportSourceListFile, "metadata", teleportRepoMetadata)
	if err := os.WriteFile(aptTeleportSourceListFile, []byte(teleportRepoMetadata), aptFilePermsRepository); err != nil {
		return trace.Wrap(err)
	}

	pm.logger.InfoContext(ctx, "Updating apt sources", "command", "apt-get update")
	updateReposCMD := exec.CommandContext(ctx, pm.bins.AptGet, "update")
	updateReposCMDOutput, err := updateReposCMD.CombinedOutput()
	if err != nil {
		return trace.Wrap(err, string(updateReposCMDOutput))
	}
	return nil
}

// InstallPackages installs one or multiple packages into the current system.
func (pm *APT) InstallPackages(ctx context.Context, packageList []PackageVersion) error {
	if len(packageList) == 0 {
		return nil
	}

	installArgs := make([]string, 0, len(packageList)+2)
	installArgs = append(installArgs, "install", "-y")

	for _, pv := range packageList {
		if pv.Version != "" {
			installArgs = append(installArgs, pv.Name+"="+pv.Version)
			continue
		}
		installArgs = append(installArgs, pv.Name)
	}

	pm.logger.InfoContext(ctx, "Installing", "command", "apt-get "+strings.Join(installArgs, " "))

	installPackagesCMD := exec.CommandContext(ctx, pm.bins.AptGet, installArgs...)
	installPackagesCMDOutput, err := installPackagesCMD.CombinedOutput()
	if err != nil {
		return trace.Wrap(err, string(installPackagesCMDOutput))
	}
	return nil
}
