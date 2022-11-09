package pacman

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

  "github.com/danskespil/terraform-provider-pacman/client"
)

func Provider() *schema.Provider {
	return &schema.Provider{
    Schema: map[string]*schema.Schema{
      "uri": {
        Type:     schema.TypeString,
        Required: true,
      },
      "username": {
        Type:     schema.TypeString,
        Required: true,
      },
      "password": {
        Type:      schema.TypeString,
        Required:  true,
        Sensitive: true,
      },
    },
    ResourcesMap: map[string]*schema.Resource{
      "pacman_asset": resourceAsset(),
    },
	  DataSourcesMap: map[string]*schema.Resource{
    },
    ConfigureContextFunc: providerConfigure,
  }
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
  uri := d.Get("uri").(string)
  username := d.Get("username").(string)
  password := d.Get("password").(string)

  c := client.New(uri, username, password)
  return c, nil
}
