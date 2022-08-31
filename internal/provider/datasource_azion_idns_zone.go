package provider

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cemdorst/apiclient"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceIDNSID() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIDNSReadID,
		Schema: map[string]*schema.Schema{
			"domain": {
				Computed: true,
				Type:     schema.TypeString,
			},
			"is_active": {
				Computed: true,
				Type:     schema.TypeBool,
			},
			"name": {
				Computed: true,
				Type:     schema.TypeString,
			},
			"id": {
				Required: true,
				Type:     schema.TypeInt,
			},
			"nxttl": {
				Computed: true,
				Type:     schema.TypeInt,
			},
			"retry": {
				Computed: true,
				Type:     schema.TypeInt,
			},
			"soattl": {
				Computed: true,
				Type:     schema.TypeInt,
			},
			"refresh": {
				Computed: true,
				Type:     schema.TypeInt,
			},
			"expiry": {
				Computed: true,
				Type:     schema.TypeInt,
			},
			"nameservers": {
				Computed: true,
				Type:     schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

type idnsResponseID struct {
	Results struct {
		NxTTL       int      `json:"nx_ttl"`
		Domain      string   `json:"domain"`
		Retry       int      `json:"retry"`
		Name        string   `json:"name"`
		Nameservers []string `json:"nameservers"`
		SoaTTL      int      `json:"soa_ttl"`
		IsActive    bool     `json:"is_active"`
		Refresh     int      `json:"refresh"`
		Expiry      int      `json:"expiry"`
		ID          int      `json:"id"`
	} `json:"results"`
	SchemaVersion int `json:"schema_version"`
}

func flattenIDResults(r idnsResponseID) interface{} {
	var flat interface{}

	flat = map[string]interface{}{
		"nxttl":       r.Results.NxTTL,
		"domain":      r.Results.Domain,
		"retry":       r.Results.Retry,
		"name":        r.Results.Name,
		"nameservers": r.Results.Nameservers,
		"soattl":      r.Results.SoaTTL,
		"is_active":   r.Results.IsActive,
		"refresh":     r.Results.Refresh,
		"expire":      r.Results.Expiry,
		"id":          r.Results.ID,
	}
	return flat
}

func dataSourceIDNSReadID(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	id := d.Get("id")
	url := fmt.Sprintf("/intelligent_dns/%v", id)
	method := "GET"

	var diags diag.Diagnostics
	var result idnsResponseID

	client := meta.(*apiclient.Client)
	r, err := client.DoRequest(method, url)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := json.Unmarshal(r, &result); err != nil {
		return diag.FromErr(err)
	}

	if err != nil {
		diag.Errorf("unable to create hash from attribute: %s", err)
	}

	d.SetId((fmt.Sprintf("%d", result.Results.ID)))
	d.Set("domain", result.Results.Domain)
	d.Set("expiry", result.Results.Expiry)
	d.Set("is_active", result.Results.IsActive)
	d.Set("name", result.Results.Name)
	d.Set("nameservers", result.Results.Nameservers)
	d.Set("nxttl", result.Results.NxTTL)
	d.Set("refresh", result.Results.Refresh)
	d.Set("retry", result.Results.Retry)
	d.Set("soattl", result.Results.SoaTTL)

	return diags
}
