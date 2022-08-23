package provider

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cemdorst/apiclient"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/mitchellh/hashstructure/v2"
)

func DataSourceIDNS() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIDNSRead,
		Schema: map[string]*schema.Schema{
			"hosted_zones": {
				Optional: true,
				Computed: true,
				Type:     schema.TypeList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"domain": {
							Required: true,
							Type:     schema.TypeString,
						},
						"is_active": {
							Required: true,
							Type:     schema.TypeBool,
						},
						"name": {
							Required: true,
							Type:     schema.TypeString,
						},
						"id": {
							Required: true,
							Type:     schema.TypeInt,
						},
					},
				},
			},
		},
	}

}

type idnsResponse struct {
	Count int `json:"count"`
	Links struct {
		Previous interface{} `json:"previous"`
		Next     interface{} `json:"next"`
	} `json:"links"`
	TotalPages int `json:"total_pages"`
	Results    []struct {
		Domain   string `json:"domain"`
		IsActive bool   `json:"is_active"`
		Name     string `json:"name"`
		ID       int    `json:"id"`
	} `json:"results"`
	SchemaVersion int `json:"schema_version"`
}

func flattenResults(r idnsResponse) []interface{} {

	var flat []interface{}

	for i := 0; i < len(r.Results); i++ {
		flat = append(flat, map[string]interface{}{
			"domain":    r.Results[i].Domain,
			"is_active": r.Results[i].IsActive,
			"name":      r.Results[i].Name,
			"id":        r.Results[i].ID,
		})
	}
	return flat
}

func dataSourceIDNSRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	url := "/intelligent_dns"
	method := "GET"

	var diags diag.Diagnostics
	var result idnsResponse

	client := meta.(*apiclient.Client)
	r, err := client.DoRequest(method, url)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := json.Unmarshal(r, &result); err != nil {
		return diag.FromErr(err)
	}

	hash, err := hashstructure.Hash(result, hashstructure.FormatV2, nil)
	if err != nil {
		diag.Errorf("unable to create hash from attribute: %s", err)
	}

	d.Set("hosted_zones", flattenResults(result))
	d.SetId((fmt.Sprintf("%d", hash)))

	return diags
}
