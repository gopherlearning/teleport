/*
 * Teleport
 * Copyright (C) 2024  Gravitational, Inc.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package agent

import (
	"context"
	"errors"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"

	"github.com/google/renameio/v2"
	"github.com/gravitational/trace"
)

const (
	updateServiceTemplate = `# teleport-update
# DO NOT EDIT THIS FILE
[Unit]
Description=Teleport auto-update service

[Service]
Type=oneshot
ExecStart={{.UpdaterCommand}}
`
	updateTimerTemplate = `# teleport-update
# DO NOT EDIT THIS FILE
[Unit]
Description=Teleport auto-update timer unit

[Timer]
OnActiveSec=1m
OnUnitActiveSec=5m
RandomizedDelaySec=1m

[Install]
WantedBy={{.TeleportService}}
`
)

// Namespace represents a namespace within various system paths for a isolated installation of Teleport.
type Namespace string

var alphanum = regexp.MustCompile("^[a-zA-Z0-9-]*$")

// NewNamespace validates and returns a Namespace.
// Namespaces must be alphanumeric + `-`.
func NewNamespace(name string) (Namespace, error) {
	if !alphanum.MatchString(name) {
		return "", trace.Errorf("invalid namespace name %q, must be alphanumeric", name)
	}
	return Namespace(name), nil
}

// Setup installs service and timer files for the teleport-update binary.
// Afterwords, Setup reloads systemd and enables the timer with --now.
func (ns Namespace) Setup(ctx context.Context, log *slog.Logger, linkDir, dataDir string) error {
	err := ns.writeConfigFiles(linkDir)
	if err != nil {
		return trace.Errorf("failed to write teleport-update systemd config files: %w", err)
	}
	svc := &SystemdService{
		ServiceName: ns.UpdaterTimer(),
		Log:         log,
	}
	if err := svc.Sync(ctx); err != nil {
		return trace.Errorf("failed to sync systemd config: %w", err)
	}
	if err := svc.Enable(ctx, true); err != nil {
		return trace.Errorf("failed to enable teleport-update systemd timer: %w", err)
	}
	return nil
}

// Teardown removes all traces of the auto-updater, including its configuration.
func (ns Namespace) Teardown(ctx context.Context, log *slog.Logger, linkDir, dataDir string) error {
	svc := &SystemdService{
		ServiceName: ns.UpdaterTimer(),
		Log:         log,
	}
	if err := svc.Disable(ctx); err != nil {
		return trace.Errorf("failed to disable teleport-update systemd timer: %w", err)
	}
	servicePath := filepath.Join(ns.LinkDir(linkDir), serviceDir, ns.UpdaterService())
	if err := os.Remove(servicePath); err != nil && !errors.Is(err, fs.ErrNotExist) {
		return trace.Errorf("failed to remove teleport-update systemd service: %w", err)
	}
	timerPath := filepath.Join(ns.LinkDir(linkDir), serviceDir, ns.UpdaterTimer())
	if err := os.Remove(timerPath); err != nil && !errors.Is(err, fs.ErrNotExist) {
		return trace.Errorf("failed to remove teleport-update systemd timer: %w", err)
	}
	if err := svc.Sync(ctx); err != nil {
		return trace.Errorf("failed to sync systemd config: %w", err)
	}
	if err := os.RemoveAll(filepath.Join(ns.DataDir(dataDir), VersionsDirName)); err != nil {
		return trace.Errorf("failed to remove versions directory: %w", err)
	}
	return nil
}

func (ns Namespace) writeConfigFiles(linkDir string) error {
	var args string
	if ns != "" {
		args = " --namespace=" + string(ns)
	}
	err := writeTemplate(
		ns.UpdaterServicePath(linkDir), updateServiceTemplate,
		struct{ UpdaterCommand string }{
			ns.UpdaterBinaryPath(linkDir) + args + " update",
		},
	)
	if err != nil {
		return trace.Wrap(err)
	}
	err = writeTemplate(
		ns.UpdaterTimerPath(linkDir), updateTimerTemplate,
		struct{ TeleportService string }{ns.TeleportService()},
	)
	if err != nil {
		return trace.Wrap(err)
	}
	return nil
}

func writeTemplate(path, t string, values any) error {
	dir, file := filepath.Split(path)
	if err := os.MkdirAll(dir, systemDirMode); err != nil {
		return trace.Wrap(err)
	}
	opts := []renameio.Option{
		renameio.WithPermissions(configFileMode),
		renameio.WithExistingPermissions(),
	}
	f, err := renameio.NewPendingFile(path, opts...)
	if err != nil {
		return trace.Wrap(err)
	}
	defer f.Cleanup()

	tmpl, err := template.New(file).Parse(t)
	if err != nil {
		return trace.Wrap(err)
	}
	err = tmpl.Execute(f, values)
	if err != nil {
		return trace.Wrap(err)
	}
	return trace.Wrap(f.CloseAtomicallyReplace())
}

// TeleportName is the namespaced name for Teleport (e.g., teleport_mycluster).
func (ns Namespace) TeleportName() string {
	return "teleport" + ns.suffix()
}

// TeleportService is the namespaced name for the Teleport systemd service (e.g., teleport_mycluster.service).
func (ns Namespace) TeleportService() string {
	return ns.TeleportName() + ".service"
}

// TeleportConfigPath is the namespaced path for the Teleport config file (e.g., /etc/teleport_mycluster.yaml).
func (ns Namespace) TeleportConfigPath() string {
	return filepath.Join("/etc", ns.TeleportName()+".yaml")
}

// TeleportPIDPath is the namespaced path for the Teleport PID file (e.g., /run/teleport_mycluster.pid).
func (ns Namespace) TeleportPIDPath() string {
	return filepath.Join("/run", ns.TeleportName()+".pid")
}

// UpdaterName is the namespaced name for the updater binary.
func (ns Namespace) UpdaterName() string {
	return BinaryName + ns.suffix()
}

// UpdaterService is the namespaced name for the updater systemd service (e.g., updater_mycluster.service).
func (ns Namespace) UpdaterService() string {
	return ns.UpdaterName() + ".service"
}

// UpdaterTimer is the namespaced name for the updater systemd timer (e.g., updater_mycluster.timer).
func (ns Namespace) UpdaterTimer() string {
	return ns.UpdaterName() + ".timer"
}

// DataDir returns the namespaced data directory path.
func (ns Namespace) DataDir(globalDataDir string) string {
	if ns == "" {
		return globalDataDir
	}
	return filepath.Join(globalDataDir, "ns", string(ns))
}

// LinkDir returns the namespaced link directory path.
func (ns Namespace) LinkDir(globalLinkDir string) string {
	if ns == "" {
		return globalLinkDir
	}
	return filepath.Join(globalLinkDir, "teleport", string(ns))
}

// UpdaterBinaryPath returns the namespaced path for the updater binary.
func (ns Namespace) UpdaterBinaryPath(globalLinkDir string) string {
	return filepath.Join(
		ns.LinkDir(globalLinkDir), "bin",
		BinaryName,
	)
}

// LinkServicePath returns the namespaced path for the linked service.
// Note: /usr/local/lib/systemd/system/teleport_mycluster.service or
// /lib/systemd/system/teleport.service when not namespaced.
func (ns Namespace) LinkServicePath(globalLinkDir string) string {
	if ns == "" {
		globalLinkDir = "/"
	}
	return filepath.Join(
		globalLinkDir, serviceDir,
		ns.TeleportService(),
	)
}

// UpdaterServicePath returns the namespaced path for the updater service.
// Note: /usr/local/lib/systemd/system/teleport-update_mycluster.service or
// /usr/local/lib/systemd/system/teleport-update.service when not namespaced.
func (ns Namespace) UpdaterServicePath(globalLinkDir string) string {
	return filepath.Join(
		globalLinkDir, serviceDir,
		ns.UpdaterService(),
	)
}

// UpdaterTimerPath returns the namespaced path for the updater timer.
// Note: /usr/local/lib/systemd/system/teleport-update_mycluster.timer or
// /usr/local/lib/systemd/system/teleport-update.timer when not namespaced.
func (ns Namespace) UpdaterTimerPath(globalLinkDir string) string {
	return filepath.Join(
		globalLinkDir, serviceDir,
		ns.UpdaterTimer(),
	)
}

func (ns Namespace) suffix() string {
	if ns == "" {
		return ""
	}
	return "_" + string(ns)
}

// ReplaceTeleportService replaces the default paths in the Teleport service config with namespaced paths.
func (ns Namespace) ReplaceTeleportService(orig []byte, linkDir string) []byte {
	cfg := string(orig)
	cfg = strings.ReplaceAll(cfg, "/usr/local/", ns.LinkDir(linkDir)+"/")
	cfg = strings.ReplaceAll(cfg, "/etc/teleport.yaml", ns.TeleportConfigPath())
	cfg = strings.ReplaceAll(cfg, "/run/teleport.pid", ns.TeleportPIDPath())
	cfg = strings.ReplaceAll(cfg, "/etc/default/teleport", "/etc/default/"+ns.TeleportName())
	return []byte(cfg)
}
