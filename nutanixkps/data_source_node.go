package nutanixkps

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"terraform-provider-nutanixkps/nutanixkpsclient"

	"terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"
)

func dataSourceNode() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceNodeRead,
		Schema:      NodeDataSourceMap(),
		Description: "Describes a Karbon Platform Services Service Domain Node. A Service Domain Node is the baremetal/virtual machine that is being managed by Karbon Platform Services.",
	}
}

func dataSourceNodeRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*nutanixkpsclient.Client)
	nodeID, iok := d.GetOk("node_id")
	nodeNameInput, nok := d.GetOk("node_name")
	if !iok && !nok {
		return diag.Errorf("please provide one of node_id or node_name attributes")
	}
	var node *models.Node
	// Search by ID
	if iok {
		var errGet *models.APIErrorPayload
		node, errGet = client.NodeGet(nodeID.(string))
		if errGet != nil {
			d.SetId("")
			return diag.FromErr(nutanixkpsclient.APIErrorToError(errGet))
		}
	} else {
		var errSearch error
		node, errSearch = searchNodeByName(client, nodeNameInput.(string))
		if errSearch != nil {
			d.SetId("")
			return diag.FromErr(errSearch)
		}
	}

	flatNodeRole, err := flattenNodeRole(node.Role)
	if err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("name", node.Name); err != nil {
		return diag.Errorf("failed to set attribute name for node %s", d.Id())
	}
	if err := d.Set("description", node.Description); err != nil {
		return diag.Errorf("failed to set attribute description for node %s", d.Id())
	}
	if err := d.Set("gateway", node.Gateway); err != nil {
		return diag.Errorf("failed to set attribute gateway for node %s", d.Id())
	}
	if err := d.Set("ip_address", node.IPAddress); err != nil {
		return diag.Errorf("failed to set attribute ip_address for node %s", d.Id())
	}
	if err := d.Set("is_bootstrap_master", node.IsBootstrapMaster); err != nil {
		return diag.Errorf("failed to set attribute is_bootstrap_master for node %s", d.Id())
	}
	if err := d.Set("serial_number", node.SerialNumber); err != nil {
		return diag.Errorf("failed to set attribute serial_number for node %s", d.Id())
	}
	if err := d.Set("subnet", node.Subnet); err != nil {
		return diag.Errorf("failed to set attribute subnet for node %s", d.Id())
	}
	if err := d.Set("service_domain_id", node.SvcDomainID); err != nil {
		return diag.Errorf("failed to set attribute service_domain_id for node %s", d.Id())
	}
	if err := d.Set("role", flatNodeRole); err != nil {
		return diag.Errorf("failed to set attribute role for node %s", d.Id())
	}

	d.SetId(node.ID)

	return nil
}

func searchNodeByName(client *nutanixkpsclient.Client, nodeName string) (*models.Node, error) {
	nodeList := make([]*models.Node, 0)
	nodes, errGet := client.NodeList()
	if errGet != nil {
		return nil, nutanixkpsclient.APIErrorToError(errGet)
	}
	for _, n := range nodes {
		if *n.Name == nodeName {
			nodeList = append(nodeList, n)

		}
	}
	switch a := len(nodeList); {
	case a < 1:
		return nil, fmt.Errorf("No node found with name %s", nodeName)
	case a > 1:
		return nil, fmt.Errorf("Multiple nodes found with name %s", nodeName)
	default:
		return nodeList[0], nil
	}
}

func NodeDataSourceMap() map[string]*schema.Schema {
	ndsm := NodeElementDataSourceMap()
	ndsm["node_id"] = &schema.Schema{
		Type:          schema.TypeString,
		Optional:      true,
		ConflictsWith: []string{"node_name"},
	}
	ndsm["node_name"] = &schema.Schema{
		Type:          schema.TypeString,
		Optional:      true,
		ConflictsWith: []string{"node_id"},
	}
	return ndsm
}

func NodeElementDataSourceMap() map[string]*schema.Schema {
	return map[string]*schema.Schema{
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
		"serial_number": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"service_domain_id": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"is_bootstrap_master": {
			Type:     schema.TypeBool,
			Optional: true,
			Computed: true,
		},
		"role": {
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"master": {
						Type:     schema.TypeBool,
						Optional: true,
						Computed: true,
					},
					"worker": {
						Type:     schema.TypeBool,
						Optional: true,
						Computed: true,
					},
				},
			},
		},
		"gateway": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validation.IsIPAddress,
		},
		"ip_address": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validation.IsIPAddress,
		},
		"subnet": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validation.IsIPAddress,
		},
	}
}
