package nutanixkps

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	// "github.com/hashicorp/terraform/helper/validation"

	"terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"
	"terraform-provider-nutanixkps/nutanixkpsclient"
	"terraform-provider-nutanixkps/utils"
)

func resourceServiceBinding() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceServiceBindingCreate,
		ReadContext:   resourceServiceBindingRead,
		UpdateContext: resourceServiceBindingUpdate,
		DeleteContext: resourceServiceBindingDelete,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Description: "Name of the Service Binding. Maximum length of 200 characters.",
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"bind_resource": {
				Description: "The resource, service domain or project, to which the service class will be bound",
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Description: "ID of the resource, either a service domain or a project",
							Type:     schema.TypeString,
							Required: true,
						},
						"type": &schema.Schema{
							Description: "Valid values for this field are: 'SERVICEDOMAIN' or 'PROJECT'",
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringInSlice(getSupportedServiceClassScopeTypes(), false),
						},
					},
				},
			},
			"service_class_id": &schema.Schema{
				Description: `Service class ID of the service you want to bind. 
				For example, ID of the AI Inferencing or the Istio service. 
				You can obtain service class IDs by querying the ServiceClass resource.`,
				Type:     schema.TypeString,
				Required: true,
			},
		},
		Description: "Describes a Karbon Platform Services Service Binding.",
	}
}

func resourceServiceBindingCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Print("[Debug] Entering resourceServiceBindingCreate")
	client := m.(*nutanixkpsclient.Client)
	serviceBindingName := d.Get("name").(string)
	serviceBindingServiceClassID := d.Get("service_class_id").(string)
	serviceBindingBindResource := d.Get("bind_resource").([]interface{})
	expandedBindResource, err := expandBindResource(serviceBindingBindResource)
	if err != nil {
		return diag.Errorf("error occured while creating service binding %s: %s", serviceBindingName, err)
	}
	serviceBindingParam := &models.ServiceBindingParam{}
	serviceBindingParam.Name = &serviceBindingName
	serviceBindingParam.SvcClassID = &serviceBindingServiceClassID
	serviceBindingParam.BindResource = expandedBindResource
	utils.PrintToJSON(serviceBindingParam, "serviceBindingParam: ")

	log.Print("[resourceServiceBindingCreate] Pre create")

	serviceBindingID, createErr := client.ServiceBindingCreate(serviceBindingParam)
	if createErr != nil {
		return diag.Errorf("error occured creating serviceBinding %s: %s", serviceBindingName, nutanixkpsclient.APIErrorToError(createErr))
	}

	d.SetId(*serviceBindingID)
	return resourceServiceBindingRead(ctx, d, m)

}

func resourceServiceBindingUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//todo

	return resourceServiceBindingRead(ctx, d, m)
}

func resourceServiceBindingRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Print("[Debug] Entering resourceServiceBindingRead")
	client := m.(*nutanixkpsclient.Client)
	var diags diag.Diagnostics
	serviceBindingID := d.Id()
	serviceBinding, errGet := client.ServiceBindingGet(serviceBindingID)
	if errGet != nil {
		d.SetId("")
		return nil
	}
	//flatten bindresource
	flatBindResource, err := flattenBindResource(serviceBinding.BindResource)
	if err != nil {
		diag.Errorf("error occured flattening bind resource for service binding %s: %s", d.Id(), err)
	}
	utils.PrintToJSON(serviceBinding, "read serviceBinding: ")
	if err := d.Set("name", serviceBinding.Name); err != nil {
		return diag.Errorf("failed to set attribute name for serviceBinding %s", d.Id())
	}
	if err := d.Set("service_class_id", serviceBinding.SvcClassID); err != nil {
		return diag.Errorf("failed to set attribute service_class_id for serviceBinding %s", d.Id())
	}
	if err := d.Set("bind_resource", flatBindResource); err != nil {
		return diag.Errorf("failed to set attribute bind_resource for serviceBinding %s", d.Id())
	}
	return diags
}

func resourceServiceBindingDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*nutanixkpsclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	serviceBindingID := d.Id()

	_, err := client.ServiceBindingDelete(serviceBindingID)
	if err != nil {
		return diag.FromErr(nutanixkpsclient.APIErrorToError(err))
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}

func expandBindResource(bindResources []interface{}) (*models.ServiceBindingResource, error) {
	if len(bindResources) != 1 {
		return nil, fmt.Errorf("exactly one service bind resource must be passed")
	}
	svcBindResource := &models.ServiceBindingResource{}
	bindResource := bindResources[0].(map[string]interface{})
	if bindResourceID, okID := bindResource["id"]; okID {
		svcBindResource.ID = bindResourceID.(string)
	} else {
		return nil, fmt.Errorf("key id was not part of bind resource definition")
	}
	if bindResourceType, okType := bindResource["type"]; okType {
		switch t := bindResourceType.(string); t {
		case "SERVICEDOMAIN":
			svcBindResource.Type = models.ServiceBindingResourceTypeSERVICEDOMAIN
		case "PROJECT":
			svcBindResource.Type = models.ServiceBindingResourceTypePROJECT
		default:
			return nil, fmt.Errorf("invalid bind resource type: %s", t)
		}
	} else {
		return nil, fmt.Errorf("key type was not part of bind resource definition")
	}
	return svcBindResource, nil
}

func getSupportedServiceClassScopeTypes() []string {
	return []string{
		"SERVICEDOMAIN",
		"PROJECT",
	}
}

func flattenBindResource(bindResource *models.ServiceBindingResource) ([]map[string]interface{}, error) {
	if bindResource == nil {
		return nil, fmt.Errorf("cannot flatten nil node role")
	}
	utils.PrintToJSON(bindResource, "[flattenBindResource] bindResource: ")
	flatBindResource := make([]map[string]interface{}, 0)
	return append(flatBindResource, map[string]interface{}{
		"id":   bindResource.ID,
		"type": bindResource.Type,
	}), nil
}
