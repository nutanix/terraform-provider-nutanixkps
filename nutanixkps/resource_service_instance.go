package nutanixkps

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	// "github.com/hashicorp/terraform/helper/validation"

	"sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"
	"sherlock-terraform-provider-nutanixkps/nutanixkpsclient"
	"sherlock-terraform-provider-nutanixkps/utils"
)

func resourceServiceInstance() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceServiceInstanceCreate,
		ReadContext:   resourceServiceInstanceRead,
		UpdateContext: resourceServiceInstanceUpdate,
		DeleteContext: resourceServiceInstanceDelete,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"project_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"service_class_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceServiceInstanceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Print("[Debug] Entering resourceServiceInstanceCreate")
	client := m.(*nutanixkpsclient.Client)
	serviceInstanceName := d.Get("name").(string)
	serviceInstanceDescription := d.Get("description").(string)
	serviceInstanceProjectID := d.Get("project_id").(string)
	serviceInstanceServiceClassID := d.Get("service_class_id").(string)

	serviceInstanceParam := &models.ServiceInstanceParam{}
	serviceInstanceParam.Description = serviceInstanceDescription
	serviceInstanceParam.Name = serviceInstanceName
	serviceInstanceParam.ScopeID = &serviceInstanceProjectID
	serviceInstanceParam.SvcClassID = &serviceInstanceServiceClassID
	utils.PrintToJSON(serviceInstanceParam, "serviceInstanceParam: ")

	log.Print("[resourceServiceInstanceCreate] Pre create")

	serviceInstanceID, createErr := client.ServiceInstanceCreate(serviceInstanceParam)
	if createErr != nil {
		return diag.Errorf("error occured creating serviceInstance %s: %s", serviceInstanceName, nutanixkpsclient.APIErrorToError(createErr))
	}

	d.SetId(*serviceInstanceID)
	return resourceServiceInstanceRead(ctx, d, m)

}

func resourceServiceInstanceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//todo

	return resourceServiceInstanceRead(ctx, d, m)
}

func resourceServiceInstanceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Print("[Debug] Entering resourceServiceInstanceRead")
	client := m.(*nutanixkpsclient.Client)
	var diags diag.Diagnostics
	serviceInstanceID := d.Id()
	serviceInstance, errGet := client.ServiceInstanceGet(serviceInstanceID)
	if errGet != nil {
		d.SetId("")
		return nil
	}

	utils.PrintToJSON(serviceInstance, "read serviceInstance: ")
	if err := d.Set("name", serviceInstance.Name); err != nil {
		return diag.Errorf("failed to set attribute name for serviceInstance %s", d.Id())
	}
	if err := d.Set("description", serviceInstance.Description); err != nil {
		return diag.Errorf("failed to set attribute description for serviceInstance %s", d.Id())
	}
	if err := d.Set("project_id", serviceInstance.ScopeID); err != nil {
		return diag.Errorf("failed to set attribute project_id for serviceInstance %s", d.Id())
	}
	if err := d.Set("service_class_id", serviceInstance.SvcClassID); err != nil {
		return diag.Errorf("failed to set attribute service_class_id for serviceInstance %s", d.Id())
	}
	return diags
}

func resourceServiceInstanceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*nutanixkpsclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	serviceInstanceID := d.Id()

	_, err := client.ServiceInstanceDelete(serviceInstanceID)
	if err != nil {
		return diag.FromErr(nutanixkpsclient.APIErrorToError(err))
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
