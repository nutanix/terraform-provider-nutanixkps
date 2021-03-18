package nutanixkps

import (
	"context"
	"fmt"
	"sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"
	"sherlock-terraform-provider-nutanixkps/nutanixkpsclient"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

//todo check error for reading empty nodes..
func dataSourceNodes() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceNodesRead,
		Schema: map[string]*schema.Schema{
			"nodes": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: NodeElementDataSourceMap(),
				},
			},
		},
	}
}

func dataSourceNodesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*nutanixkpsclient.Client)

	nodes, errGet := client.NodeList()
	if errGet != nil {
		return diag.FromErr(nutanixkpsclient.APIErrorToError(errGet))
	}
	flatNodes, err := flattenNodes(nodes)
	if err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("nodes", flatNodes); err != nil {
		return diag.Errorf("failed to set nodes output: %s", err)
	}
	d.SetId(resource.UniqueId())

	return nil
}

func flattenNodes(nodes []*models.Node) ([]map[string]interface{}, error) {
	if nodes == nil {
		return nil, fmt.Errorf("cannot flatten nil nodes")
	}
	flatNodes := make([]map[string]interface{}, 0)
	for _, node := range nodes {
		flatNodeRole, err := flattenNodeRole(node.Role)
		if err != nil {
			return nil, err
		}
		n := map[string]interface{}{
			"name":                node.Name,
			"description":         node.Description,
			"gateway":             node.Gateway,
			"ip_address":          node.IPAddress,
			"is_bootstrap_master": node.IsBootstrapMaster,
			"serial_number":       node.SerialNumber,
			"subnet":              node.Subnet,
			"service_domain_id":   node.SvcDomainID,
			"role":                flatNodeRole,
		}
		flatNodes = append(flatNodes, n)
	}
	return flatNodes, nil
}
