package snyk

import (
	"context"

	"github.com/lendi-au/terraform-provider-snyk/snyk/api"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceOrganization() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceOrganizationRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"slug": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"url": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceOrganizationRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	so := m.(api.SnykOptions)
	id := d.Get("id").(string)

	org, err := api.GetOrganization(so, id)

	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("id", org.Id)
	d.Set("created", org.Created.String())
	d.Set("name", org.Name)
	d.Set("slug", org.Slug)
	d.Set("url", org.Url)

	d.SetId(org.Id)

	return diags
}
