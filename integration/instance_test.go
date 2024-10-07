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

package integration

import (
	"context"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/gravitational/trace"
	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"

	"github.com/gravitational/teleport"
	"github.com/gravitational/teleport/api/breaker"
	"github.com/gravitational/teleport/api/types"
	"github.com/gravitational/teleport/integration/helpers"
	"github.com/gravitational/teleport/lib"
	"github.com/gravitational/teleport/lib/backend"
	"github.com/gravitational/teleport/lib/cloud/imds"
	"github.com/gravitational/teleport/lib/defaults"
	"github.com/gravitational/teleport/lib/service"
	"github.com/gravitational/teleport/lib/service/servicecfg"
	"github.com/gravitational/teleport/lib/services"
	"github.com/gravitational/teleport/lib/utils"
)

// basicDirCopy performs a very simplistic recursive copy from one directory to another. this helper was
// written specifically for setting up teleport data directories for testing purposes and may not be
// suitable for other applications.
func basicDirCopy(src string, dst string) error {
	entries, err := os.ReadDir(src)
	if err != nil {
		return trace.Wrap(err)
	}

	if err := os.MkdirAll(dst, teleport.PrivateDirMode); err != nil {
		return trace.Wrap(err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			if err := basicDirCopy(filepath.Join(src, entry.Name()), filepath.Join(dst, entry.Name())); err != nil {
				return trace.Wrap(err)
			}
			continue
		}

		data, err := os.ReadFile(filepath.Join(src, entry.Name()))
		if err != nil {
			return trace.Wrap(err)
		}

		if err := os.WriteFile(filepath.Join(dst, entry.Name()), data, teleport.PrivateDirMode); err != nil {
			return trace.Wrap(err)
		}
	}

	return nil
}

// TestInstanceCertReissue tests the reissuance of an instance certificate when
// the instance has malformed system roles using pre-constructed data directories
// generated by an older teleport version that permitted token mix-and-match.
func TestInstanceCertReissue(t *testing.T) {
	t.Setenv("_insecuredevmode_no_parallel", "1") // panic if the test is or will become parallel
	lib.SetInsecureDevMode(true)
	defer lib.SetInsecureDevMode(false)

	var eg errgroup.Group
	defer func() { require.NoError(t, eg.Wait()) }()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Create temporary directories for the auth and agent data directories.
	authDir, agentDir := t.TempDir(), t.TempDir()

	// Write the instance assets to the temporary directories to set up pre-existing
	// state for our teleport instances to use.
	require.NoError(t, basicDirCopy("testdata/auth", authDir))
	require.NoError(t, basicDirCopy("testdata/agent", agentDir))

	authCfg := servicecfg.MakeDefaultConfig()
	authCfg.Version = defaults.TeleportConfigVersionV3
	authCfg.DataDir = authDir
	authCfg.Auth.Enabled = true
	authCfg.Auth.ListenAddr.Addr = helpers.NewListener(t, service.ListenerAuth, &authCfg.FileDescriptors)
	authCfg.SetAuthServerAddress(authCfg.Auth.ListenAddr)
	// ensure auth server is using the pre-constructed sqlite db
	authCfg.Auth.StorageConfig.Params = backend.Params{defaults.BackendPath: filepath.Join(authDir, defaults.BackendDir)}
	var err error
	authCfg.Auth.ClusterName, err = services.NewClusterNameWithRandomID(types.ClusterNameSpecV2{
		ClusterName: "auth-server",
	})
	require.NoError(t, err)
	authCfg.Auth.NetworkingConfig.SetProxyListenerMode(types.ProxyListenerMode_Multiplex)

	authCfg.Proxy.Enabled = true
	authCfg.Proxy.DisableWebInterface = true
	proxyAddr := helpers.NewListener(t, service.ListenerProxyWeb, &authCfg.FileDescriptors)
	authCfg.Proxy.WebAddr.Addr = proxyAddr

	authCfg.SSH.Enabled = true
	authCfg.SSH.Addr.Addr = "localhost:0"
	authCfg.CircuitBreakerConfig = breaker.NoopBreakerConfig()
	authCfg.Log = utils.NewLoggerForTests()
	authCfg.InstanceMetadataClient = imds.NewDisabledIMDSClient()

	authProc, err := service.NewTeleport(authCfg)
	require.NoError(t, err)
	require.NoError(t, authProc.Start())
	eg.Go(func() error { return authProc.WaitForSignals(ctx, nil) })

	authIdentity, err := authProc.GetIdentityForTesting(t, types.RoleInstance)
	require.NoError(t, err)
	require.ElementsMatch(t, []string{string(types.RoleAuth), string(types.RoleProxy)}, authIdentity.SystemRoles)

	authEventCh := make(chan service.Event)
	authProc.ListenForEvents(ctx, service.TeleportCredentialsUpdatedEvent, authEventCh)

	timeoutCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	_, err = authProc.WaitForEvent(timeoutCtx, service.TeleportCredentialsUpdatedEvent)
	require.NoError(t, err, "timed out waiting for second auth identity")

	authIdentity, err = authProc.GetIdentityForTesting(t, types.RoleInstance)
	require.NoError(t, err)
	require.ElementsMatch(t, []string{string(types.RoleAuth), string(types.RoleProxy), string(types.RoleNode)}, authIdentity.SystemRoles)

	agentCfg := servicecfg.MakeDefaultConfig()
	agentCfg.Version = defaults.TeleportConfigVersionV3
	agentCfg.DataDir = agentDir
	agentCfg.ProxyServer = utils.NetAddr{
		AddrNetwork: "tcp",
		Addr:        proxyAddr,
	}

	agentCfg.Auth.Enabled = false
	agentCfg.Proxy.Enabled = false
	agentCfg.SSH.Enabled = true

	agentCfg.WindowsDesktop.Enabled = true
	agentCfg.CircuitBreakerConfig = breaker.NoopBreakerConfig()
	agentCfg.Log = utils.NewLoggerForTests()
	agentCfg.MaxRetryPeriod = time.Second
	agentCfg.InstanceMetadataClient = imds.NewDisabledIMDSClient()

	agentProc, err := service.NewTeleport(agentCfg)
	require.NoError(t, err)
	require.NoError(t, agentProc.Start())
	eg.Go(func() error { return agentProc.WaitForSignals(ctx, nil) })

	agentIdentity, err := agentProc.GetIdentityForTesting(t, types.RoleInstance)
	require.NoError(t, err)
	require.ElementsMatch(t, []string{string(types.RoleNode)}, agentIdentity.SystemRoles)

	_, err = agentProc.WaitForEvent(timeoutCtx, service.TeleportCredentialsUpdatedEvent)
	require.NoError(t, err, "timed out waiting for second agent identity")

	agentIdentity, err = agentProc.GetIdentityForTesting(t, types.RoleInstance)
	require.NoError(t, err)
	require.ElementsMatch(t, []string{string(types.RoleNode), string(types.RoleWindowsDesktop)}, agentIdentity.SystemRoles)
}
