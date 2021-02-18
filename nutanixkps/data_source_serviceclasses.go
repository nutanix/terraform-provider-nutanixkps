package nutanixkps

import (
	"context"
	"terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"
	"terraform-provider-nutanixkps/nutanixkpsclient"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

//todo check error for reading empty serviceClasses..
func dataSourceServiceClasses() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceServiceClassesRead,
		Schema: map[string]*schema.Schema{
			"service_classes": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: ServiceClassElementDataSourceMap(),
				},
			},
		},
	}
}

func dataSourceServiceClassesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*nutanixkpsclient.Client)

	serviceClasses, errGet := client.ServiceClassList()
	if errGet != nil {
		return diag.FromErr(nutanixkpsclient.APIErrorToError(errGet))
	}
	flatServiceClasses, err := flattenServiceClasses(serviceClasses)
	if err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("service_classes", flatServiceClasses); err != nil {
		return diag.Errorf("failed to set service_classes output: %s", err)
	}
	d.SetId(resource.UniqueId())

	return nil
}

func flattenServiceClasses(serviceClasses []*models.ServiceClass) ([]map[string]interface{}, error) {
	flatServiceClasses := make([]map[string]interface{}, 0)
	for _, sc := range serviceClasses {
		flatServiceClassTags, err := flattenServiceClassTags(sc.Tags)
		if err != nil {
			return nil, err
		}
		n := map[string]interface{}{
			"name":                       sc.Name,
			"description":                sc.Description,
			"type":                       sc.Type,
			"service_version":            sc.SvcVersion,
			"min_service_domain_version": sc.MinSvcDomainVersion,
			"state":                      sc.State,
			"tags":                       flatServiceClassTags,
		}
		flatServiceClasses = append(flatServiceClasses, n)
	}
	return flatServiceClasses, nil
}
