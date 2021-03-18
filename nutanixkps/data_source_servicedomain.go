package nutanixkps

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"terraform-provider-nutanixkps/nutanixkpsclient"

	"terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"
)

func dataSourceServiceDomains() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceServiceDomainsRead,
		Schema: map[string]*schema.Schema{
			"servicedomains": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
		Description: "Describes a Nutanix KPS Service Domain.",
	}
}

func flattenServiceDomainData(serviceDomains *[]*models.ServiceDomain) []interface{} {
	if serviceDomains != nil {
		sds := make([]interface{}, len(*serviceDomains), len(*serviceDomains))

		for i, serviceDomain := range *serviceDomains {
			sd := make(map[string]interface{})

			sd["id"] = serviceDomain.ID
			sd["name"] = serviceDomain.Name
			sd["description"] = serviceDomain.Description

			sds[i] = sd
		}

		return sds
	}

	return make([]interface{}, 0)
}

func dataSourceServiceDomainsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*nutanixkpsclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	servicedomains, err := client.ServiceDomainList()
	if err != nil {
		return diag.FromErr(errors.New(*err.Message))
	}

	sds := flattenServiceDomainData(&servicedomains)
	if err := d.Set("servicedomains", sds); err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
