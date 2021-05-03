package nutanixkps

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	// "github.com/hashicorp/terraform/helper/validation"

	"sherlock-terraform-provider-nutanixkps/nutanixkpsclient"
	"sherlock-terraform-provider-nutanixkps/utils"

	"sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"
)

const (
	DEFAULTWAITTIMEOUT = 60
	MINIMUMWAITTIMEOUT = 10
	WAITMINTIMEOUT     = 30
	WAITDELAY          = 60
	NODEINFOERROR      = "node info error"
)

func resourceNode() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceNodeCreate,
		ReadContext:   resourceNodeRead,
		UpdateContext: resourceNodeUpdate,
		DeleteContext: resourceNodeDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Description: `Name of the node: 
				Name must include lowercase alphanumeric characters and must start and end with an lowercase alphanumeric character.
				Dash (-) and dot (.) characters are allowed as delimiters. Maximum length of 60 characters.`,
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Description: "Describe the node. For example, the main purpose or use case of the node.",
				Type:     schema.TypeString,
				Required: true,
			},
			"serial_number": &schema.Schema{
				Description: "Node serial number. Any alpha characters must be in all capitals.",
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"service_domain_id": &schema.Schema{
				Description: "Id of the service domain to which this node belongs",
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"is_bootstrap_master": &schema.Schema{
				Description: "Default setting is true. Set to false indicates this node is not a bootstrap master.",
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"role": &schema.Schema{
				Description: `Set the role as master or worker. Default setting is true to enable the role as master as well as worker. 
				Set to false to disable a role.`,
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"master": {
							Type:     schema.TypeBool,
							Required: true,
						},
						"worker": {
							Type:     schema.TypeBool,
							Required: true,
						},
					},
				},
			},
			"gateway": &schema.Schema{
				Description: "Gateway IPv4 address for this node",
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.IsIPAddress,
			},
			"ip_address": &schema.Schema{
				Description: "IPv4 address of this node",
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.IsIPAddress,
			},
			"subnet": &schema.Schema{
				Description: "Subnet mask for this node",
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.IsIPAddress,
			},
			"wait_for_onboarding": &schema.Schema{
				Description: `Default setting is false and the terraform provider does not wait for the node to be onboarded. 
				Set to true indicates that the terraform provider waits for the node to be onboarded.`,
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"wait_timeout_minutes": &schema.Schema{
				Description: "Wait timeout in minutes",
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      DEFAULTWAITTIMEOUT,
				ValidateFunc: validation.IntAtLeast(MINIMUMWAITTIMEOUT),
			},
		},
		Description: "Describes a Karbon Platform Services Service Domain Node. A Service Domain Node is the baremetal/virtual machine that is being managed by Karbon Platform Services.",
	}
}

func resourceNodeCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Print("[Debug] Entering resourceNodeCreate")
	client := m.(*nutanixkpsclient.Client)
	// var diags diag.Diagnostics
	nodeName := d.Get("name").(string)
	nodeDescription := d.Get("description").(string)
	nodeGateway := d.Get("gateway").(string)
	nodeIPAddress := d.Get("ip_address").(string)
	nodeIsBootstrapMaster := d.Get("is_bootstrap_master").(bool)
	nodeSerialNumber := d.Get("serial_number").(string)
	nodeSubnet := d.Get("subnet").(string)
	nodeServiceDomainID := d.Get("service_domain_id").(string)
	nodeRole := d.Get("role").([]interface{})
	waitForOnboarding := d.Get("wait_for_onboarding").(bool)
	waitTimeoutMinutes := d.Get("wait_timeout_minutes").(int)
	///READ WAITTIMEOUT VARS
	expandedNodeRole, err := expandNodeRole(nodeRole)
	if err != nil {
		return diag.Errorf("error occured while expanding role for node %s: %s", nodeName, err)
	}

	// Warning or errors can be collected in a slice type
	node := &models.Node{}
	node.Description = nodeDescription
	node.Gateway = &nodeGateway
	node.IPAddress = &nodeIPAddress
	node.IsBootstrapMaster = nodeIsBootstrapMaster
	node.Name = &nodeName
	node.SerialNumber = &nodeSerialNumber
	node.Subnet = &nodeSubnet
	node.SvcDomainID = &nodeServiceDomainID
	node.Role = expandedNodeRole
	nodeID, createErr := client.NodeCreate(node)
	if createErr != nil {
		return diag.Errorf("error occured while creating node %s: %s", nodeName, nutanixkpsclient.APIErrorToError(createErr))
	}
	if waitForOnboarding {
		err := WaitForNodeOnboarding(client, int64(waitTimeoutMinutes), *nodeID)
		if err != nil {
			return diag.Errorf("An error occured while waiting for onboarding of node %s.: %s.", nodeName, err)
		}
	}
	d.SetId(*nodeID)
	return resourceNodeRead(ctx, d, m)

}

func resourceNodeUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Print("[Debug] Entering resourceNodeUpdate")
	client := m.(*nutanixkpsclient.Client)
	nodeId := d.Id()
	nodeName := d.Get("name").(string)
	nodeDescription := d.Get("description").(string)
	nodeGateway := d.Get("gateway").(string)
	nodeIPAddress := d.Get("ip_address").(string)
	nodeIsBootstrapMaster := d.Get("is_bootstrap_master").(bool)
	nodeSerialNumber := d.Get("serial_number").(string)
	nodeSubnet := d.Get("subnet").(string)
	nodeServiceDomainID := d.Get("service_domain_id").(string)
	nodeRole := d.Get("role").([]interface{})
	expandedNodeRole, err := expandNodeRole(nodeRole)
	if err != nil {
		return diag.Errorf("error occured while expanding role for node %s: %s", nodeName, err)
	}

	node := &models.Node{}
	node.ID = nodeId
	node.Description = nodeDescription
	node.Gateway = &nodeGateway
	node.IPAddress = &nodeIPAddress
	node.IsBootstrapMaster = nodeIsBootstrapMaster
	node.Name = &nodeName
	node.SerialNumber = &nodeSerialNumber
	node.Subnet = &nodeSubnet
	node.SvcDomainID = &nodeServiceDomainID
	node.Role = expandedNodeRole
	nodeIdAfterUpdate, updateErr := client.NodeUpdate(nodeId, node)
	if updateErr != nil {
		return diag.Errorf("An error occured while updating node %s: %s", nodeName, nutanixkpsclient.APIErrorToError(updateErr))
	}
	d.SetId(*nodeIdAfterUpdate)
	return resourceNodeRead(ctx, d, m)
}

func resourceNodeRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Print("[Debug] Entering resourceNodeRead")
	client := m.(*nutanixkpsclient.Client)
	var diags diag.Diagnostics
	nodeID := d.Id()
	node, errGet := client.NodeGet(nodeID)
	if errGet != nil {
		d.SetId("")
		return nil
	}

	utils.PrintToJSON(node, "read node: ")
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
	return diags
}

func resourceNodeDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*nutanixkpsclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	nodeID := d.Id()

	_, err := client.NodeDelete(nodeID)
	if err != nil {
		return diag.FromErr(nutanixkpsclient.APIErrorToError(err))
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}

func expandNodeRole(roles []interface{}) (*models.NodeRole, error) {

	if len(roles) != 1 {
		return nil, fmt.Errorf("exactly one role must be passed")
	}
	nodeRole := &models.NodeRole{}
	role := roles[0].(map[string]interface{})
	if roleMaster, okMaster := role["master"]; okMaster {
		nodeRole.Master = roleMaster.(bool)
	} else {
		return nil, fmt.Errorf("key master was not part of role definition")
	}
	if roleWorker, okWorker := role["worker"]; okWorker {
		nodeRole.Worker = roleWorker.(bool)
	} else {
		return nil, fmt.Errorf("key worker was not part of role definition")
	}
	return nodeRole, nil
}

func flattenNodeRole(nodeRole *models.NodeRole) ([]map[string]interface{}, error) {
	if nodeRole == nil {
		return nil, fmt.Errorf("cannot flatten nil node role")
	}
	flatNodeRole := make([]map[string]interface{}, 0)
	return append(flatNodeRole, map[string]interface{}{
		"worker": nodeRole.Worker,
		"master": nodeRole.Master,
	}), nil
}

func WaitForNodeOnboarding(client *nutanixkpsclient.Client, waitTimeoutMinutes int64, nodeUUID string) error {
	log.Printf("Starting wait for node**************: %s", nodeUUID)
	stateConf := &resource.StateChangeConf{
		Pending:    []string{"ONBOARDING"},
		Target:     []string{"ONBOARDED", "HEALTHY"},
		Refresh:    onboardingStateRefreshFunc(client, nodeUUID),
		Timeout:    time.Duration(waitTimeoutMinutes) * time.Minute,
		Delay:      WAITDELAY * time.Second,
	}

	if _, errWaitTask := stateConf.WaitForState(); errWaitTask != nil {
		return fmt.Errorf("Error waiting for nodes to be onboarded: %s", errWaitTask)
	}
	log.Printf("Ending wait for node: %s", nodeUUID)
	return nil
}

func onboardingStateRefreshFunc(client *nutanixkpsclient.Client, nodeUUID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		state := "UNKNOWN"
		ni, err := client.NodeInfoGet(nodeUUID)

		if err != nil {
			convErr := nutanixkpsclient.APIErrorToError(err)
			if strings.Contains(fmt.Sprint(convErr), "is not found") {
				return nil, "", convErr
			}
		}
		if ni.Onboarded {
			state = "ONBOARDED"
		} else {
			state = "ONBOARDING"
		}
		if ni.Healthy {
			state = "HEALTHY"
		}
		log.Printf("[onboardingStateRefreshFunc] state: %s", state)
		return ni, state, nil
	}
}

