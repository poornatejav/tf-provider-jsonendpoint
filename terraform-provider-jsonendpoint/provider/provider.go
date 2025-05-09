package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// New returns a terraform resource provider
func New() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"base_url": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Base URL of the JSON endpoint (e.g., http://localhost:9000)",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"jsonendpoint_item": resourceJSONEndpoint(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"jsonendpoint_fetch": DataSourceJSONEndpointFetch(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

// providerConfigure parses provider-level settings and returns a Config
func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics
	base := d.Get("base_url").(string)
	if base == "" {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Missing base_url",
			Detail:   "The provider base_url must be set.",
		})
		return nil, diags
	}
	return &Config{BaseURL: base}, diags
}
