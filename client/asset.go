package client

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

const assetPath = "asset"

type Asset struct {
  WorkgroupID     int       `json:"WorkgroupID"`
  AssetID         int       `json:"AssetID"`
  AssetName       string    `json:"AssetName"`
  DNSName         string    `json:"DNSName"`
  DomainName      string    `json:"DomainName"`
  IPAddress       string    `json:"IPAddress"`
  MacAddress      string    `json:"MacAddress"`
  AssetType       string    `json:"AssetType"`
  OperatingSystem string    `json:"OperatingSystem"`
  CreateDate      time.Time `json:"CreateDate"`
  LastUpdateDate  time.Time `json:"LastUpdateDate"`
}

func (c *Client) CreateAsset(name string, domain string, assetType string, ip string) (*Asset, diag.Diagnostics) {
  ctyPath := cty.Path{cty.GetAttrStep{Name: "client_asset_create"}}

  params := map[string]string{
    "AssetName":  name,
    "DomainName": domain,
    "AssetType":  assetType,
    "IPAddress":  ip,
  }

  res, diags := c.newRequest("POST", assetPath, params)
  if diags.HasError() == true {
    return nil, diags
  }
  defer res.Body.Close()

  a := &Asset{}
  if res.StatusCode != 200 {
    diags = append(diags, diag.Diagnostic{
      Severity:      diag.Error,
      Summary:       fmt.Sprintf("Could not create asset: %s", name),
      Detail:        fmt.Sprintf("Status returned by request: %s", res.Status),
      AttributePath: ctyPath,
    })
  } else {
    err := json.NewDecoder(res.Body).Decode(&a)
    if err != nil {
      diags = append(diags, diag.Diagnostic{
        Severity:      diag.Warning,
        Summary:       fmt.Sprintf("Asset possibly created: %s", name),
        Detail:        "Request was successful but server response could not be parsed. See errors for more information.",
        AttributePath: ctyPath,
      })

      diags = append(diags, diag.Diagnostic{
        Severity:      diag.Error,
        Summary:       fmt.Sprintf("Could not parse response for asset: %s", name),
        Detail:        err.Error(),
        AttributePath: ctyPath,
      })
    }
  }

  return a, diags
}

func (c *Client) ReadAsset(id string) (*Asset, diag.Diagnostics) {
  ctyPath := cty.Path{cty.GetAttrStep{Name: "client_asset_read"}}

  res, diags := c.newRequest("GET", fmt.Sprintf("%s/%s", assetPath, id), nil)
  if diags.HasError() == true {
    return nil, diags
  }
  defer res.Body.Close()

  a := &Asset{}
  if res.StatusCode == 200 {
    err := json.NewDecoder(res.Body).Decode(&a)
    if err != nil {
      diags = append(diags, diag.Diagnostic{
        Severity:      diag.Error,
        Summary:       fmt.Sprintf("Could not read asset: %s", id),
        Detail:        err.Error(),
        AttributePath: ctyPath,
      })
    }
  } else if res.StatusCode != 404 {
    diags = append(diags, diag.Diagnostic{
      Severity:      diag.Error,
      Summary:       fmt.Sprintf("Could not read asset: %s", id),
      Detail:        fmt.Sprintf("Request failed with status %s", res.Status),
      AttributePath: ctyPath,
    })
  }

  return a, diags
}

func (c *Client) UpdateAsset(name string, domain string, assetType string, ip string) (*Asset, diag.Diagnostics) {
  ctyPath := cty.Path{cty.GetAttrStep{Name: "client_asset_update"}}

  params := map[string]string{
    "AssetName":  name,
    "DomainName": domain,
    "AssetType":  assetType,
    "IPAddress":  ip,
  }

  res, diags := c.newRequest("PUT", assetPath, params)
  if diags.HasError() == true {
    return nil, diags
  }
  defer res.Body.Close()

  a := &Asset{}
  if res.StatusCode != 200 {
    diags = append(diags, diag.Diagnostic{
      Severity:      diag.Error,
      Summary:       fmt.Sprintf("Could not update asset: %s", name),
      Detail:        fmt.Sprintf("Status returned by request: %s", res.Status),
      AttributePath: ctyPath,
    })
  } else {
    err := json.NewDecoder(res.Body).Decode(&a)
    if err != nil {
      diags = append(diags, diag.Diagnostic{
        Severity:      diag.Warning,
        Summary:       fmt.Sprintf("Asset possibly updated: %s", name),
        Detail:        "Request was successful but server response could not be parsed. See errors for more information.",
        AttributePath: ctyPath,
      })

      diags = append(diags, diag.Diagnostic{
        Severity:      diag.Error,
        Summary:       fmt.Sprintf("Could not parse response for asset: %s", name),
        Detail:        err.Error(),
        AttributePath: ctyPath,
      })
    }
  }

  return a, diags
}

func (c *Client) DeleteAsset(id string) diag.Diagnostics {
  ctyPath := cty.Path{cty.GetAttrStep{Name: "client_asset_update"}}

  res, diags := c.newRequest("DELETE", fmt.Sprintf("%s/%s", assetPath, id), map[string]string{})
  if diags.HasError() == true {
    return diags
  }
  defer res.Body.Close()

  if res.StatusCode != 200 {
    diags = append(diags, diag.Diagnostic{
      Severity:      diag.Error,
      Summary:       fmt.Sprintf("Could not delete asset: %s", id),
      Detail:        fmt.Sprintf("Status returned by request: %s", res.Status),
      AttributePath: ctyPath,
    })
  }

  return diags
}
