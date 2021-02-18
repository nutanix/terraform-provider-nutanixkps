package nutanixkps

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"sherlock-terraform-provider-nutanixkps/nutanixkpsclient"

	"sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"
)

func dataSourceServiceClass() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceServiceClassRead,
		Schema:      ServiceClassDataSourceMap(),
	}
}

func dataSourceServiceClassRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*nutanixkpsclient.Client)
	serviceClassID, iok := d.GetOk("service_class_id")
	serviceClassNameInput, nok := d.GetOk("service_class_name")
	if !iok && !nok {
		return diag.Errorf("please provide one of service_class_id or service_class_name attributes")
	}
	var serviceClass *models.ServiceClass
	// Search by ID
	if iok {
		var errGet *models.APIErrorPayload
		serviceClass, errGet = client.ServiceClassGet(serviceClassID.(string))
		if errGet != nil {
			d.SetId("")
			return diag.FromErr(nutanixkpsclient.APIErrorToError(errGet))
		}
	} else {
		var errSearch error
		serviceClass, errSearch = searchServiceClassByName(client, serviceClassNameInput.(string))
		if errSearch != nil {
			d.SetId("")
			return diag.FromErr(errSearch)
		}
	}

	flatServiceClassTags, err := flattenServiceClassTags(serviceClass.Tags)
	if err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("type", serviceClass.Type); err != nil {
		return diag.Errorf("failed to set attribute type for serviceClass %s", d.Id())
	}
	if err := d.Set("service_version", serviceClass.SvcVersion); err != nil {
		return diag.Errorf("failed to set attribute service_version for serviceClass %s", d.Id())
	}
	if err := d.Set("min_service_domain_version", serviceClass.MinSvcDomainVersion); err != nil {
		return diag.Errorf("failed to set attribute min_service_domain_version for serviceClass %s", d.Id())
	}
	if err := d.Set("name", serviceClass.Name); err != nil {
		return diag.Errorf("failed to set attribute name for serviceClass %s", d.Id())
	}
	if err := d.Set("description", serviceClass.Description); err != nil {
		return diag.Errorf("failed to set attribute description for serviceClass %s", d.Id())
	}

	if err := d.Set("state", serviceClass.State); err != nil {
		return diag.Errorf("failed to set attribute state for serviceClass %s", d.Id())
	}
	if err := d.Set("tags", flatServiceClassTags); err != nil {
		return diag.Errorf("failed to set attribute role for serviceClass %s", d.Id())
	}

	d.SetId(serviceClass.ID)

	return nil
}

func searchServiceClassByName(client *nutanixkpsclient.Client, serviceClassName string) (*models.ServiceClass, error) {
	serviceClassList := make([]*models.ServiceClass, 0)
	serviceClasses, errGet := client.ServiceClassList()
	if errGet != nil {
		return nil, nutanixkpsclient.APIErrorToError(errGet)
	}
	for _, n := range serviceClasses {
		if *n.Name == serviceClassName {
			serviceClassList = append(serviceClassList, n)

		}
	}
	switch a := len(serviceClassList); {
	case a < 1:
		return nil, fmt.Errorf("No serviceClass found with name %s", serviceClassName)
	case a > 1:
		return nil, fmt.Errorf("Multiple serviceClasses found with name %s", serviceClassName)
	default:
		return serviceClassList[0], nil
	}
}

func ServiceClassDataSourceMap() map[string]*schema.Schema {
	ndsm := ServiceClassElementDataSourceMap()
	ndsm["service_class_id"] = &schema.Schema{
		Type:          schema.TypeString,
		Optional:      true,
		ConflictsWith: []string{"service_class_name"},
	}
	ndsm["service_class_name"] = &schema.Schema{
		Type:          schema.TypeString,
		Optional:      true,
		ConflictsWith: []string{"service_class_id"},
	}
	return ndsm
}

func ServiceClassElementDataSourceMap() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"type": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"service_version": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"min_service_domain_version": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"description": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"state": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"tags": {
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"name": {
						Type:     schema.TypeString,
						Optional: true,
						Computed: true,
					},
					"value": {
						Type:     schema.TypeString,
						Optional: true,
						Computed: true,
					},
				},
			},
		},
	}
}

func flattenServiceClassTags(serviceClassTags []*models.ServiceClassTag) ([]map[string]interface{}, error) {
	flatServiceClassTags := make([]map[string]interface{}, 0)
	for _, sct := range serviceClassTags {
		if sct.Name == "" {
			return flatServiceClassTags, fmt.Errorf("could not flatten service class tags because name was empty")
		}
		flatServiceClassTags = append(flatServiceClassTags, map[string]interface{}{
			"name":  sct.Name,
			"value": sct.Value,
		})
	}
	return flatServiceClassTags, nil
}
