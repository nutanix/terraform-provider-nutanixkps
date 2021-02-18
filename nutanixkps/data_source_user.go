package nutanixkps

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"sherlock-terraform-provider-nutanixkps/nutanixkpsclient"

	"sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"
)

func dataSourceUser() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceUserRead,
		Schema:      UserDataSourceMap(),
	}
}

func dataSourceUserRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*nutanixkpsclient.Client)
	userID, iok := d.GetOk("user_id")
	userNameInput, nok := d.GetOk("user_name")
	if !iok && !nok {
		return diag.Errorf("please provide one of user_id or user_name attributes")
	}
	var user *models.User
	// Search by ID
	if iok {
		var errGet *models.APIErrorPayload
		user, errGet = client.UserGet(userID.(string))
		if errGet != nil {
			d.SetId("")
			return diag.FromErr(nutanixkpsclient.APIErrorToError(errGet))
		}
	} else {
		var errSearch error
		user, errSearch = searchUserByName(client, userNameInput.(string))
		if errSearch != nil {
			d.SetId("")
			return diag.FromErr(errSearch)
		}
	}

	if err := d.Set("name", user.Name); err != nil {
		return diag.Errorf("failed to set attribute name for user %s", d.Id())
	}
	if err := d.Set("role", user.Role); err != nil {
		return diag.Errorf("failed to set attribute role for user %s", d.Id())
	}
	if err := d.Set("email", user.Email); err != nil {
		return diag.Errorf("failed to set attribute email for user %s", d.Id())
	}
	if err := d.Set("tenant_id", user.TenantID); err != nil {
		return diag.Errorf("failed to set attribute tenant_id for user %s", d.Id())
	}
	d.SetId(user.ID)

	return nil
}

func searchUserByName(client *nutanixkpsclient.Client, userName string) (*models.User, error) {
	userList := make([]*models.User, 0)
	useres, errGet := client.UserList()
	if errGet != nil {
		return nil, nutanixkpsclient.APIErrorToError(errGet)
	}
	for _, n := range useres {
		if *n.Name == userName {
			userList = append(userList, n)

		}
	}
	switch a := len(userList); {
	case a < 1:
		return nil, fmt.Errorf("No user found with name %s", userName)
	case a > 1:
		return nil, fmt.Errorf("Multiple useres found with name %s", userName)
	default:
		return userList[0], nil
	}
}

func UserDataSourceMap() map[string]*schema.Schema {
	ndsm := UserElementDataSourceMap()
	ndsm["user_id"] = &schema.Schema{
		Type:          schema.TypeString,
		Optional:      true,
		ConflictsWith: []string{"user_name"},
	}
	ndsm["user_name"] = &schema.Schema{
		Type:          schema.TypeString,
		Optional:      true,
		ConflictsWith: []string{"user_id"},
	}
	return ndsm
}

func UserElementDataSourceMap() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"role": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"email": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"tenant_id": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}
