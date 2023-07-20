// Copyright Envoy Gateway Authors
// SPDX-License-Identifier: Apache-2.0
// The full text of the Apache license is available in the LICENSE file at
// the root of the repo.

package cmd

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/http/pprof"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/envoyproxy/gateway/internal/envoygateway/config"
	extensionregistry "github.com/envoyproxy/gateway/internal/extension/registry"
	gatewayapirunner "github.com/envoyproxy/gateway/internal/gatewayapi/runner"
	ratelimitrunner "github.com/envoyproxy/gateway/internal/globalratelimit/runner"
	infrarunner "github.com/envoyproxy/gateway/internal/infrastructure/runner"
	"github.com/envoyproxy/gateway/internal/logging"
	"github.com/envoyproxy/gateway/internal/message"
	providerrunner "github.com/envoyproxy/gateway/internal/provider/runner"
	"github.com/envoyproxy/gateway/internal/runner"
	xdsserverrunner "github.com/envoyproxy/gateway/internal/xds/server/runner"
	xdstranslatorrunner "github.com/envoyproxy/gateway/internal/xds/translator/runner"
)

var (
	// cfgPath is the path to the EnvoyGateway configuration file.
	cfgPath string
)

// getServerCommand returns the server cobra command to be executed.
func getServerCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "server",
		Aliases: []string{"serve"},
		Short:   "Serve Envoy Gateway",
		RunE: func(cmd *cobra.Command, args []string) error {
			return server()
		},
	}
	cmd.PersistentFlags().StringVarP(&cfgPath, "config-path", "c", "",
		"The path to the configuration file.")

	return cmd
}

// server serves Envoy Gateway.
func server() error {
	ctx := ctrl.SetupSignalHandler()
	cfg, err := getConfig()
	if err != nil {
		return err
	}

	if err := setupAdminServer(cfg); err != nil {
		return err
	}
	if err := setupRunners(ctx, cfg); err != nil {
		return err
	}

	return nil
}

// setupRunners setups all the runners required for the Envoy Gateway to
// fulfill its tasks.
func setupRunners(ctx context.Context, cfg *config.Server) error {
	if err := initRunnerManager(cfg); err != nil {
		return err
	}

	if err := runner.Manager().Run(ctx); err != nil {
		return err
	}

	return nil
}

// setupAdminServer setups the admin server of Envoy Gateway
func setupAdminServer(cfg *config.Server) error {
	logger := cfg.Logger.WithName("admin-server")

	if cfg.EnvoyGateway.Admin.Debug {
		spewConfig := spew.NewDefaultConfig()
		spewConfig.DisableMethods = true
		spewConfig.Dump(cfg)
	}

	adminHandlers := http.NewServeMux()
	address := cfg.EnvoyGateway.GetEnvoyGatewayAdmin().Address

	if cfg.EnvoyGateway.GetEnvoyGatewayAdmin().Debug {
		// Serve pprof endpoints to aid in live debugging.
		adminHandlers.HandleFunc("/debug/pprof/", pprof.Index)
		adminHandlers.HandleFunc("/debug/pprof/profile", pprof.Profile)
		adminHandlers.HandleFunc("/debug/pprof/trace", pprof.Trace)
		adminHandlers.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
		adminHandlers.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	}

	adminServer := &http.Server{
		Handler:           adminHandlers,
		Addr:              net.JoinHostPort(address.Host, fmt.Sprint(address.Port)),
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       15 * time.Second,
	}

	logger.Info("starting admin server")
	// Listen And serve admin server.
	go func() {
		if err := adminServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error(err, "failed to start admin server")
		}
	}()

	return nil
}

// initRunnerManager inits eg-manager and registers all
// the root-runnners of eg.
func initRunnerManager(cfg *config.Server) error {
	runner.Manager().Init(*cfg)
	// Setup the Extension Manager
	extMgr, err := extensionregistry.NewManager(cfg)
	if err != nil {
		return err
	}

	pResources := new(message.ProviderResources)
	// Start the Provider Service
	// It fetches the resources from the configured provider type
	// and publishes it
	providerrunner.Register(providerrunner.Resources{
		ProviderResources: pResources,
	}, *cfg)

	xdsIR := new(message.XdsIR)
	infraIR := new(message.InfraIR)
	// Start the GatewayAPI Translator Runner
	// It subscribes to the provider resources, translates it to xDS IR
	// and infra IR resources and publishes them.
	gatewayapirunner.Register(gatewayapirunner.Resources{
		ProviderResources: pResources,
		XdsIR:             xdsIR,
		InfraIR:           infraIR,
		ExtensionManager:  extMgr,
	}, *cfg)

	xds := new(message.Xds)
	// Start the Xds Translator Service
	// It subscribes to the xdsIR, translates it into xds Resources and publishes it.
	xdstranslatorrunner.Register(xdstranslatorrunner.Resources{
		XdsIR:            xdsIR,
		Xds:              xds,
		ExtensionManager: extMgr,
	}, *cfg)

	// Start the Infra Manager Runner
	// It subscribes to the infraIR, translates it into Envoy Proxy infrastructure
	// resources such as K8s deployment and services.
	infrarunner.Register(infrarunner.Resources{
		InfraIR: infraIR,
	}, *cfg)

	// Start the xDS Server
	// It subscribes to the xds Resources and configures the remote Envoy Proxy
	// via the xDS Protocol.
	xdsserverrunner.Register(xdsserverrunner.Resources{
		Xds: xds,
	}, *cfg)

	// Start the global rateLimit if it has been enabled through the config
	if cfg.EnvoyGateway.RateLimit != nil {
		// Start the Global RateLimit xDS Server
		// It subscribes to the xds Resources and translates it to Envoy Ratelimit configuration.
		ratelimitrunner.Register(ratelimitrunner.Resources{
			XdsIR: xdsIR,
		}, *cfg)
	}

	return nil
}

// getConfig gets the Server configuration
func getConfig() (*config.Server, error) {
	return getConfigByPath(cfgPath)
}

// make `cfgPath` an argument to test it without polluting the global var
func getConfigByPath(cfgPath string) (*config.Server, error) {
	// Initialize with default config parameters.
	cfg, err := config.New()
	if err != nil {
		return nil, err
	}

	logger := cfg.Logger

	// Read the config file.
	if cfgPath == "" {
		// Use default config parameters
		logger.Info("No config file provided, using default parameters")
	} else {
		// Load the config file.
		eg, err := config.Decode(cfgPath)
		if err != nil {
			logger.Error(err, "failed to decode config file", "name", cfgPath)
			return nil, err
		}
		// Set defaults for unset fields
		eg.SetEnvoyGatewayDefaults()
		cfg.EnvoyGateway = eg
		// update cfg logger
		eg.Logging.SetEnvoyGatewayLoggingDefaults()
		cfg.Logger = logging.NewLogger(eg.Logging)
	}

	if err := cfg.Validate(); err != nil {
		return nil, err
	}
	return cfg, nil
}
