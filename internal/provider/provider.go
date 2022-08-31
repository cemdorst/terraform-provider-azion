package provider

import (
	"context"

	"github.com/cemdorst/apiclient"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const HostURL string = "https://api.azionapi.net:443"

func init() {
	// Set descriptions to support markdown syntax, this will be used in document generation
	// and the language server.
	schema.DescriptionKind = schema.StringMarkdown

	// Customize the content of descriptions when output. For example you can add defaults on
	// to the exported descriptions if present.
	// schema.SchemaDescriptionBuilder = func(s *schema.Schema) string {
	// 	desc := s.Description
	// 	if s.Default != nil {
	// 		desc += fmt.Sprintf(" Defaults to `%v`.", s.Default)
	// 	}
	// 	return strings.TrimSpace(desc)
	// }
}

func New(version string) func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
			Schema: map[string]*schema.Schema{
				"apikey": &schema.Schema{
					Type:        schema.TypeString,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc("AZION_APIKEY", nil),
				},
			},
			DataSourcesMap: map[string]*schema.Resource{
				"azion_idns_zones": DataSourceIDNS(),
				"azion_idns_zone":  DataSourceIDNSID(),
			},
			ResourcesMap: map[string]*schema.Resource{},
		}
		p.ConfigureContextFunc = configure(version, p)
		return p
	}

}

func configure(version string, p *schema.Provider) func(ctx context.Context, d *schema.ResourceData) (any, diag.Diagnostics) {
	return func(ctx context.Context, d *schema.ResourceData) (any, diag.Diagnostics) {
		var diags diag.Diagnostics

		var c apiclient.Client

		c.New(d.Get("apikey").(string), HostURL)

		return &c, diags
	}
}
