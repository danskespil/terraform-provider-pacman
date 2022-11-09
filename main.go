package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"

  "github.com/danskespil/terraform-provider-pacman/pacman"
)

func main() {
  plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return pacman.Provider()
		},
	})
}
