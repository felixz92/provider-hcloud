/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	"github.com/felixz92/provider-hcloud/config/network"
	"github.com/felixz92/provider-hcloud/config/server"
	"github.com/felixz92/provider-hcloud/config/subnet"

	ujconfig "github.com/crossplane/upjet/pkg/config"
)

const (
	resourcePrefix = "hcloud"
	modulePath     = "github.com/felixz92/provider-hcloud"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("fzx.github.com"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		server.Configure,
		network.Configure,
		subnet.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
