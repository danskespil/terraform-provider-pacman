package pacman

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

  "github.com/danskespil/terraform-provider-pacman/client"
)

func resourceAsset() *schema.Resource {
  return &schema.Resource{
		CreateContext: resourceAssetCreate,
		ReadContext:   resourceAssetRead,
		UpdateContext: resourceAssetUpdate,
		DeleteContext: resourceAssetDelete,
		Importer: &schema.ResourceImporter{
			State: resourceAssetImport,
		},
    Schema: map[string]*schema.Schema{
      "workgroup_id": {
			  Type:		  schema.TypeInt,
	      Computed: true,
      },
      "asset_id": {
			  Type:		  schema.TypeInt,
	      Computed: true,
      },
      "asset_name": {
        Type:		  schema.TypeString,
	      Required: true,
        ForceNew: true,
      },
      "dns_name": {
        Type:		  schema.TypeString,
	      Computed: true,
      },
      "domain_name": {
        Type:		  schema.TypeString,
	      Required: true,
      },
      "ip_address": {
        Type:		  schema.TypeString,
	      Required: true,
      },
      "mac_address": {
        Type:		  schema.TypeString,
	      Computed: true,
      },
      "asset_type": {
        Type:		  schema.TypeString,
	      Required: true,
      },
      "operating_system": {
        Type:		  schema.TypeString,
	      Computed: true,
      },
      "create_date": {
        Type:		  schema.TypeString,
	      Computed: true,
      },
      "last_update_date": {
        Type:		  schema.TypeString,
	      Computed: true,
      },
    },
  }
}

func resourceAssetCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
  var diags diag.Diagnostics

  name := d.Get("asset_name").(string)
  domain_name := d.Get("domain_name").(string)
  asset_type := d.Get("asset_type").(string)
  ip_address := d.Get("ip_address").(string)

  c := m.(*client.Client)

  a, diag := c.CreateAsset(name, domain_name, asset_type, ip_address)
  if diag.HasError() == true {
    return diag
  }

  d.Set("workgroup_id", a.WorkgroupID)
  d.Set("asset_id", a.AssetID)
  d.Set("dns_name", a.DNSName)
  d.Set("mac_address", a.MacAddress)
  d.Set("operating_system", a.OperatingSystem)
  d.Set("create_date", strconv.FormatInt(a.CreateDate.Unix(), 10))
  d.Set("last_update_date", strconv.FormatInt(a.LastUpdateDate.Unix(), 10))

  d.SetId(strconv.Itoa(a.AssetID))

  return diags
}

func resourceAssetRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
  var diags diag.Diagnostics

  c := m.(*client.Client)

  a, diag := c.ReadAsset(d.Id())
  if diag.HasError() == true {
    return diag
  }

  if a.AssetID == 0 {
    d.SetId("")
  }
  d.Set("workgroup_id", a.WorkgroupID)
  d.Set("asset_id", a.AssetID)
  d.Set("asset_name", a.AssetName)
  d.Set("dns_name", a.DNSName)
  d.Set("domain_name", a.DomainName)
  d.Set("ip_address", a.IPAddress)
  d.Set("mac_address", a.MacAddress)
  d.Set("asset_type", a.AssetType)
  d.Set("operating_system", a.OperatingSystem)
  d.Set("create_date", strconv.FormatInt(a.CreateDate.Unix(), 10))
  d.Set("last_update_date", strconv.FormatInt(a.LastUpdateDate.Unix(), 10))

  return diags
}

func resourceAssetUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
  name := d.Get("asset_name").(string)
  domain_name := d.Get("domain_name").(string)
  asset_type := d.Get("asset_type").(string)
  ip_address := d.Get("ip_address").(string)

  c := m.(*client.Client)

  _, diag := c.UpdateAsset(name, domain_name, asset_type, ip_address)
  if diag.HasError() == true {
    return diag
  }

  return resourceAssetRead(ctx, d, m)
}

func resourceAssetDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
  var diags diag.Diagnostics

  c := m.(*client.Client)

  diag := c.DeleteAsset(d.Id())
  if diag.HasError() == true {
    return diag
  }

  d.SetId("")

  return diags
}

func resourceAssetImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	return []*schema.ResourceData{d}, nil
}
