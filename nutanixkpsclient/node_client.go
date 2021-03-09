package nutanixkpsclient

import (
	"sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/client/node"
	"sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"

	runtime "github.com/go-openapi/runtime/client"
)

// NodeList - Returns all Nodes
func (c *Client) NodeList() ([]*models.Node, *models.APIErrorPayload) {
	nodeListParams := node.NewNodeListParams()
	ok, err := c.CloudmgmtAPIClient.Node.NodeList(nodeListParams, runtime.BearerToken(c.Token))
	if err != nil {
		if apiErr, ok := err.(*node.NodeListDefault); ok {
			return nil, apiErr.Payload
		}
		return nil, errorToAPIError(err)
	}
	return ok.Payload.NodeList, nil
}

// NodeGet - Get a Node
func (c *Client) NodeGet(nodeID string) (nodeModel *models.Node, apiErr *models.APIErrorPayload) {
	nodeGetParams := node.NewNodeGetParams()
	nodeGetParams.NodeID = nodeID
	ok, err := c.CloudmgmtAPIClient.Node.NodeGet(nodeGetParams, runtime.BearerToken(c.Token))
	if err != nil {
		if apiErr, ok := err.(*node.NodeGetDefault); ok {
			return nil, apiErr.Payload
		}
		return nil, errorToAPIError(err)
	}
	return ok.Payload, nil
}

// // NodeCreate - Create a Node
func (c *Client) NodeCreate(nodeModel *models.Node) (nodeID *string, apiErr *models.APIErrorPayload) {
	nodeCreateParams := node.NewNodeCreateParams()
	nodeCreateParams.Body = nodeModel
	ok, err := c.CloudmgmtAPIClient.Node.NodeCreate(nodeCreateParams, runtime.BearerToken(c.Token))
	if err != nil {
		if apiErr, ok := err.(*node.NodeCreateDefault); ok {
			return nil, apiErr.Payload
		}
		return nil, errorToAPIError(err)
	}
	nodeID = ok.Payload.ID
	return nodeID, nil
}

// // NodeUpdate - Update a Node
func (c *Client) NodeUpdate(nodeID string, nodeModel *models.Node) (nodeIDAfterUpdate *string, apiErr *models.APIErrorPayload) {
	nodeUpdateParams := node.NewNodeUpdateParams()
	nodeUpdateParams.NodeID = nodeID
	nodeUpdateParams.Body = nodeModel
	ok, err := c.CloudmgmtAPIClient.Node.NodeUpdate(nodeUpdateParams, runtime.BearerToken(c.Token))
	if err != nil {
		if apiErr, ok := err.(*node.NodeUpdateDefault); ok {
			return nil, apiErr.Payload
		}
		return nil, errorToAPIError(err)
	}
	nodeIDAfterUpdate = ok.Payload.ID
	return nodeIDAfterUpdate, nil
}

// NodeDelete - Delete a Node
func (c *Client) NodeDelete(nodeID string) (deletedNodeID *string, apiErr *models.APIErrorPayload) {
	nodeDeleteParams := node.NewNodeDeleteParams()
	nodeDeleteParams.NodeID = nodeID
	ok, err := c.CloudmgmtAPIClient.Node.NodeDelete(nodeDeleteParams, runtime.BearerToken(c.Token))
	if err != nil {
		if apiErr, ok := err.(*node.NodeDeleteDefault); ok {
			return nil, apiErr.Payload
		}
		return nil, errorToAPIError(err)
	}
	deletedNodeID = ok.Payload.ID
	return deletedNodeID, nil
}
