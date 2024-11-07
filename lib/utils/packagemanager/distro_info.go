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
	"context"
	"net/http"
	"slices"

	"github.com/gravitational/trace"

	"github.com/gravitational/teleport/lib/linux"
)

// PackageManagerForSystem returns the PackageManager for the current detected linux distro.
func PackageManagerForSystem(osRelease *linux.OSRelease, fsRootPrefix string, binariesLocation BinariesLocation, httpGet func(string) (*http.Response, error)) (PackageManager, error) {
	aptWellKnownIDs := []string{"debian", "ubuntu"}
	legacyAPT := []string{"xenial", "trusty"}

	yumWellKnownIDs := []string{"amzn", "rhel", "centos", "rocky", "almalinux"}

	zypperWellKnownIDs := []string{"sles", "opensuse-tumbleweed", "opensuse-leap"}

	switch {
	case slices.Contains(aptWellKnownIDs, osRelease.ID):
		if slices.Contains(legacyAPT, osRelease.VersionCodename) {
			pm, err := NewAPTLegacy(&APTConfig{
				fsRootPrefix: fsRootPrefix,
				bins:         binariesLocation,
				httpGet:      httpGet,
			})
			if err != nil {
				return nil, trace.Wrap(err)
			}
			return pm, nil
		}

		pm, err := NewAPT(&APTConfig{
			fsRootPrefix: fsRootPrefix,
			bins:         binariesLocation,
			httpGet:      httpGet,
		})
		if err != nil {
			return nil, trace.Wrap(err)
		}
		return pm, nil

	case slices.Contains(yumWellKnownIDs, osRelease.ID):
		pm, err := NewYUM(&YUMConfig{
			fsRootPrefix: fsRootPrefix,
			bins:         binariesLocation,
		})
		if err != nil {
			return nil, trace.Wrap(err)
		}
		return pm, nil

	case slices.Contains(zypperWellKnownIDs, osRelease.ID):
		pm, err := NewZypper(&ZypperConfig{
			fsRootPrefix: fsRootPrefix,
			bins:         binariesLocation,
		})
		if err != nil {
			return nil, trace.Wrap(err)
		}
		return pm, nil

	default:
		return nil, trace.NotFound("package manager for etc/os-release:ID=%s not found", osRelease.ID)
	}
}

// PackageManager describes the required methods to implement a package manager.
type PackageManager interface {
	// AddTeleportRepository adds the Teleport repository using a specific channel.
	// developmentRepo indicates whether to use the production or development artifacts.
	// Example for APT:
	// - production repo is at https://apt.releases.teleport.dev/
	// - development repo is at https://apt.releases.development.teleport.dev/ and it contains dev and pre-release builds (eg alpha versions).
	AddTeleportRepository(ctx context.Context, ldi *linux.OSRelease, repoChannel string, developmentRepo bool) error
	// InstallPackages installs a list of packages.
	// If a PackageVersion does not contain the version, then it will install the latest available.
	InstallPackages(context.Context, []PackageVersion) error
}
