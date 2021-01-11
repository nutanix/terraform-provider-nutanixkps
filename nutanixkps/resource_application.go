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

const (
	HELMPACKAINGTYPE string = "helm"
)

func resourceApplication() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceApplicationCreate,
		ReadContext:   resourceApplicationRead,
		UpdateContext: resourceApplicationUpdate,
		DeleteContext: resourceApplicationDelete,
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
			"edge_ids": {
				Required: true,
				Type:     schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"helm_chart_filename": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"helm_values_filename": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tenant_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceApplicationCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Print("[Debug] Entering resourceApplicationCreate")
	client := m.(*nutanixkpsclient.Client)
	applicationName := d.Get("name").(string)
	applicationDescription := d.Get("description").(string)
	applicationProjectID := d.Get("project_id").(string)
	applicationEdgeIDs := expandEdgeIDs(d.Get("edge_ids").([]interface{}))
	applicationHelmChartFileName := d.Get("helm_chart_filename").(string)
	applicationHelmValuesFileName := ""
	if applicationHelmValuesFileNameInput, ok := d.Get("helm_values_filename").(string); ok {
		applicationHelmValuesFileName = applicationHelmValuesFileNameInput
	}
	application := &models.ApplicationV2{}
	application.Description = applicationDescription
	application.Name = &applicationName
	application.EdgeIDs = applicationEdgeIDs
	application.PackagingType = HELMPACKAINGTYPE
	application.ProjectID = applicationProjectID
	utils.PrintToJSON(application, "application: ")
	applicationString, err := applicationToString(application)
	if err != nil {
		return diag.FromErr(err)
	}
	log.Printf("applicationString: %s", applicationString)
	log.Print("[resourceApplicationCreate] Pre create")
	if !utils.FileExists(applicationHelmChartFileName) {
		return diag.Errorf("file %s does not exist!", applicationHelmChartFileName)
	}
	if applicationHelmValuesFileName != "" && !utils.FileExists(applicationHelmChartFileName) {
		return diag.Errorf("file %s does not exist!", applicationHelmValuesFileName)
	}
	applicationID, createErr := client.HelmApplicationCreate(applicationString, applicationHelmChartFileName, applicationHelmValuesFileName)
	if createErr != nil {
		return diag.Errorf("error occured creating application %s: %s", applicationName, nutanixkpsclient.APIErrorToError(createErr))
	}

	d.SetId(*applicationID)
	return resourceApplicationRead(ctx, d, m)

}

func resourceApplicationUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//todo

	return resourceApplicationRead(ctx, d, m)
}

func resourceApplicationRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Print("[Debug] Entering resourceApplicationRead")
	client := m.(*nutanixkpsclient.Client)
	var diags diag.Diagnostics
	applicationID := d.Id()
	application, errGet := client.ApplicationV2Get(applicationID)
	if errGet != nil {
		d.SetId("")
		return nil
	}

	utils.PrintToJSON(application, "read application: ")
	if err := d.Set("name", application.Name); err != nil {
		return diag.Errorf("failed to set attribute name for application %s", d.Id())
	}
	if err := d.Set("description", application.Description); err != nil {
		return diag.Errorf("failed to set attribute description for application %s", d.Id())
	}
	if err := d.Set("tenant_id", application.TenantID); err != nil {
		return diag.Errorf("failed to set attribute tenant_id for application %s", d.Id())
	}

	// if err := d.Set("gateway", application.Gateway); err != nil {
	// 	return diag.Errorf("failed to set attribute gateway for application %s", d.Id())
	// }
	// if err := d.Set("ip_address", application.IPAddress); err != nil {
	// 	return diag.Errorf("failed to set attribute ip_address for application %s", d.Id())
	// }
	// if err := d.Set("is_bootstrap_master", application.IsBootstrapMaster); err != nil {
	// 	return diag.Errorf("failed to set attribute is_bootstrap_master for application %s", d.Id())
	// }
	// if err := d.Set("serial_number", application.SerialNumber); err != nil {
	// 	return diag.Errorf("failed to set attribute serial_number for application %s", d.Id())
	// }
	// if err := d.Set("subnet", application.Subnet); err != nil {
	// 	return diag.Errorf("failed to set attribute subnet for application %s", d.Id())
	// }
	// if err := d.Set("service_domain_id", application.SvcDomainID); err != nil {
	// 	return diag.Errorf("failed to set attribute service_domain_id for application %s", d.Id())
	// }
	// if err := d.Set("role", flatNodeRole); err != nil {
	// 	return diag.Errorf("failed to set attribute role for application %s", d.Id())
	// }
	return diags
}

func resourceApplicationDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*nutanixkpsclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	applicationID := d.Id()

	_, err := client.ApplicationV2Delete(applicationID)
	if err != nil {
		return diag.FromErr(nutanixkpsclient.APIErrorToError(err))
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}

func expandEdgeIDs(edgeIDInterface []interface{}) []string {
	sList := make([]string, len(edgeIDInterface))
	for i, in := range edgeIDInterface {
		sList[i] = in.(string)
	}
	return sList
}

func applicationToString(application *models.ApplicationV2) (string, error) {
	b, err := application.MarshalBinary()
	if err != nil {
		return "", err
	}
	return string(b), nil

}
