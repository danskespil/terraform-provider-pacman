package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

type Client struct {
  BaseUri string
  Client  *http.Client
  Header  *http.Header
}

func New (baseUri string, username string, password string) (*Client) {
  u := strings.TrimSuffix(baseUri, "/")

  c := &http.Client{
    Timeout: 300 * time.Second,
  }

  r := &http.Request{Header: map[string][]string{
    "Content-Type":  {"application/json"},
  }}
  r.SetBasicAuth(username, password)
  h := r.Header.Clone()

  return &Client{u, c, &h}
}

func validateRequestType (requestType string) bool {
  validTypes := []string{"GET", "POST", "PUT", "DELETE"}

  for _, vt := range validTypes {
    if vt == requestType {
      return true
    }
  }
  return false
}

func (c *Client) newRequest (requestType string, path string, params map[string]string) (*http.Response, diag.Diagnostics) {
  var diags diag.Diagnostics

  if vr := validateRequestType(requestType); vr != true {
    diags = append(diags, diag.Diagnostic{
      Severity:      diag.Error,
      Summary:       fmt.Sprintf("Could not generate HTTP request for path: %s", path),
      Detail:        fmt.Sprintf("%s is not a valid HTTP request type", requestType),
      AttributePath: cty.Path{cty.GetAttrStep{Name: "client_new_request"}},
    })
    return nil, diags
  }

  jb, _ := json.Marshal(params)
  body := bytes.NewBuffer(jb)

  req, err := http.NewRequest(requestType, fmt.Sprintf("%s/%s", c.BaseUri, path), body)
  if err != nil {
    diags = append(diags, diag.Diagnostic{
      Severity:      diag.Error,
      Summary:       fmt.Sprintf("Could not generate %s request for path: %s", requestType, path),
      Detail:        err.Error(),
      AttributePath: cty.Path{cty.GetAttrStep{Name: "client_new_request"}},
    })
    return nil, diags
  }

  req.Header = *c.Header

  res, err := c.Client.Do(req)
  if err != nil {
    diags = append(diags, diag.Diagnostic{
      Severity:      diag.Error,
      Summary:       fmt.Sprintf("Could not generate %s request for path: %s", requestType, path),
      Detail:        err.Error(),
      AttributePath: cty.Path{cty.GetAttrStep{Name: "client_new_request"}},
    })
    return nil, diags
  }

  return res, diags
}
