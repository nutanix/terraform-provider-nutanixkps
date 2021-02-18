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

func resourceProject() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceProjectCreate,
		ReadContext:   resourceProjectRead,
		UpdateContext: resourceProjectUpdate,
		DeleteContext: resourceProjectDelete,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Description: "Name of the project: Maximum length of 200 characters.",
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": &schema.Schema{
				Description: `Describe the project. For example, the main purpose or use case of the project.
				Maximum length of 200 characters.`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"edge_ids": {
				Description: `List of service domain and / or node IDs which are part of this project. 
				Even if you want to add a single ID, use square brackets. For example, [ '<svc domain id>' ]`,
				Required: true,
				Type:     schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"privileged": &schema.Schema{
				Description: `Default setting is false. 
				Set to true indicates that the project can run with elevated privileges`,
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"user": {
				Description: "List of users who can access this project",
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"role": {
							Description: "Valid values for Role are: PROJECT_ADMIN, PROJECT_USER",
							Type:     schema.TypeString,
							Required: true,
						},
						"user_id": {
							Description: "ID of the user to be added to the project",
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"edge_selector_type": &schema.Schema{
				Description: `Type for selecting nodes / service domains belonging to this project. 
				Valid values for this field are: 'Category' or 'Explicit'.
				Currently the terraform provider supports only Explicit mode.`,
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "Explicit",
				ValidateFunc: validation.StringInSlice(getSupportedEdgeSelectorTypes(), false),
			},
			"tenant_id": &schema.Schema{
				Description: "Id of the tenant to which this application belongs",
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
		},
	}
}

func resourceProjectCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Print("[Debug] Entering resourceProjectCreate")
	client := m.(*nutanixkpsclient.Client)
	projectName := d.Get("name").(string)
	projectDescription := d.Get("description").(string)
	projectPrivileged := d.Get("privileged").(bool)
	projectEdgeSelectorType := d.Get("edge_selector_type").(string)

	projectEdgeIDs := expandEdgeIDs(d.Get("edge_ids").([]interface{}))
	projectUsers, err := expandProjectUsers(d.Get("user").([]interface{}))

	if err != nil {
		return diag.Errorf("error expanding project user for project %s: %s", projectName, err)
	}
	project := &models.Project{}
	project.Description = &projectDescription
	project.Name = &projectName
	project.EdgeIDs = projectEdgeIDs
	project.Users = projectUsers
	project.Privileged = projectPrivileged
	project.EdgeSelectorType = &projectEdgeSelectorType
	utils.PrintToJSON(project, "project: ")

	log.Print("[resourceProjectCreate] Pre create")

	projectID, createErr := client.ProjectV2Create(project)
	if createErr != nil {
		return diag.Errorf("error creating project %s: %s", projectName, nutanixkpsclient.APIErrorToError(createErr))

	}

	d.SetId(*projectID)
	return resourceProjectRead(ctx, d, m)

}

func resourceProjectUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//todo

	return resourceProjectRead(ctx, d, m)
}

func resourceProjectRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Print("[Debug] Entering resourceProjectRead")
	client := m.(*nutanixkpsclient.Client)
	var diags diag.Diagnostics
	projectID := d.Id()
	project, errGet := client.ProjectV2Get(projectID)
	if errGet != nil {
		d.SetId("")
		return nil
	}

	flatUsers, err := flattenProjectUsers(project.Users)
	if err != nil {
		return diag.Errorf("error flattening user for project %s: %s", d.Id(), err)
	}

	utils.PrintToJSON(project, "read project: ")
	if err := d.Set("name", project.Name); err != nil {
		return diag.Errorf("failed to set attribute name for project %s", d.Id())
	}
	if err := d.Set("description", project.Description); err != nil {
		return diag.Errorf("failed to set attribute description for project %s", d.Id())
	}
	if err := d.Set("user", flatUsers); err != nil {
		return diag.Errorf("failed to set attribute user for project %s", d.Id())
	}
	if err := d.Set("tenant_id", project.TenantID); err != nil {
		return diag.Errorf("failed to set attribute tenant_id for project %s", d.Id())
	}
	if err := d.Set("privileged", project.Privileged); err != nil {
		return diag.Errorf("failed to set attribute privileged for project %s", d.Id())
	}

	return diags
}

func resourceProjectDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*nutanixkpsclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	projectID := d.Id()

	_, err := client.ProjectV2Delete(projectID)
	if err != nil {
		return diag.FromErr(nutanixkpsclient.APIErrorToError(err))
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}

func expandProjectUsers(projectUsers []interface{}) ([]*models.ProjectUserInfo, error) {
	expandedProjectUsers := make([]*models.ProjectUserInfo, 0)
	for _, pu := range projectUsers {
		puCast := pu.(map[string]interface{})
		puObj := models.ProjectUserInfo{}
		if puRole, okRole := puCast["role"]; okRole {
			pur := puRole.(string)
			puObj.Role = &pur
		} else {
			return nil, fmt.Errorf("failed to expand project user, key role not found")
		}
		if puUserID, okUserID := puCast["user_id"]; okUserID {
			puID := puUserID.(string)
			puObj.UserID = &puID
		} else {
			return nil, fmt.Errorf("failed to expand project user, key user_id not found")
		}
		expandedProjectUsers = append(expandedProjectUsers, &puObj)
	}
	return expandedProjectUsers, nil
}

func flattenProjectUsers(projectUsers []*models.ProjectUserInfo) ([]map[string]interface{}, error) {
	flatProjectusers := make([]map[string]interface{}, 0)
	for _, pu := range projectUsers {
		projectUser := map[string]interface{}{}
		if pu.Role != nil {
			projectUser["role"] = *pu.Role
		} else {
			return flatProjectusers, fmt.Errorf("role was nil for project user")
		}
		if pu.UserID != nil {
			projectUser["user_id"] = *pu.UserID
		} else {
			return flatProjectusers, fmt.Errorf("user_id was nil for project user")
		}
		flatProjectusers = append(flatProjectusers, projectUser)
	}
	return flatProjectusers, nil
}

func getSupportedEdgeSelectorTypes() []string {
	return []string{
		"Explicit",
		"Category",
	}
}
