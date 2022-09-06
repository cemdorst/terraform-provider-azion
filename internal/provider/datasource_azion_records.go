package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/cemdorst/apiclient"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceRecords() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceRecordsRead,

		Schema: map[string]*schema.Schema{
			"records": {
				Computed: true,
				Optional: true,
				Type:     schema.TypeList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"answers_list": {
							Computed: true,
							Type:     schema.TypeList,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"description": {
							Computed: true,
							Type:     schema.TypeString,
						},
						"policy": {
							Computed: true,
							Type:     schema.TypeString,
						},
						"record_type": {
							Computed: true,
							Type:     schema.TypeString,
						},
						"ttl": {
							Computed: true,
							Type:     schema.TypeInt,
						},
						"record_id": {
							Computed: true,
							Type:     schema.TypeInt,
						},
						"entry": {
							Computed: true,
							Type:     schema.TypeString,
						},
					},
				},
			},
			"zone_id": {
				Required: true,
				Type:     schema.TypeInt,
			},
		},
	}
}

func dataSourceRecordsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	id := d.Get("zone_id")
	url := fmt.Sprintf("/intelligent_dns/%v/records", id)
	method := "GET"

	client := meta.(*apiclient.Client)
	response_bytes, err := client.DoRequest(method, url)
	if err != nil {
		return diag.FromErr(err)
	}

	reader_a := bytes.NewReader(response_bytes)
	decode_reader_a := make(map[string]interface{}, 0)
	if err = json.NewDecoder(reader_a).Decode(&decode_reader_a); err != nil {
		return diag.FromErr(err)
	}

	decode_reader_b := make(map[string]interface{}, 0)
	bytes_of_reader_a, err := json.Marshal(decode_reader_a["results"])
	if err != nil {
		log.Fatal(err)
	}

	reader_b := bytes.NewReader(bytes_of_reader_a)
	if err = json.NewDecoder(reader_b).Decode(&decode_reader_b); err != nil {
		return diag.FromErr(err)
	}

	if err = d.Set("records", decode_reader_b["records"]); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
