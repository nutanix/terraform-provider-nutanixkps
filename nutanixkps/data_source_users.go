package nutanixkps

import (
	"context"
	"terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"
	"terraform-provider-nutanixkps/nutanixkpsclient"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

//todo check error for reading empty users..
func dataSourceUsers() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceUsersRead,
		Schema: map[string]*schema.Schema{
			"users": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: UserElementDataSourceMap(),
				},
			},
		},
	}
}

func dataSourceUsersRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*nutanixkpsclient.Client)

	users, errGet := client.UserList()
	if errGet != nil {
		return diag.FromErr(nutanixkpsclient.APIErrorToError(errGet))
	}
	flatUsers, err := flattenUsers(users)
	if err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("users", flatUsers); err != nil {
		return diag.Errorf("failed to set users output: %s", err)
	}
	d.SetId(resource.UniqueId())

	return nil
}

func flattenUsers(users []*models.User) ([]map[string]interface{}, error) {
	flatUsers := make([]map[string]interface{}, 0)
	for _, u := range users {
		n := map[string]interface{}{
			"name":      u.Name,
			"role":      u.Role,
			"email":     u.Email,
			"tenant_id": u.TenantID,
		}
		flatUsers = append(flatUsers, n)
	}
	return flatUsers, nil
}
