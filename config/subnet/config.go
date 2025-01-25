package subnet

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("hcloud_network_subnet", func(r *config.Resource) {
		r.ShortGroup = "subnet"

		r.References["network_id"] = config.Reference{
			TerraformName: "hcloud_network",
		}
	})
}
