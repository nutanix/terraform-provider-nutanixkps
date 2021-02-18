package nutanixkpsclient

import (
	"terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/client/node"
	"terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"

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

//todo
// // ServiceDomainUpdate - Update a ServiceDomain
// func (c *Client) ServiceDomainUpdate(svcDomainID string, serviceDomain *models.ServiceDomain) (deletedNodeID *string, apiErr *models.APIErrorPayload) {
// 	serviceDomainUpdateParams := service_domain.NewServiceDomainUpdateParams()
// 	serviceDomainUpdateParams.SvcDomainID = svcDomainID
// 	serviceDomainUpdateParams.Body = serviceDomain
// 	ok, err := c.CloudmgmtAPIClient.ServiceDomain.ServiceDomainUpdate(serviceDomainUpdateParams, runtime.BearerToken(c.Token))
// 	if err != nil {
// 		if apiErr, ok := err.(*service_domain.ServiceDomainUpdateDefault); ok {
// 			return nil, apiErr.Payload
// 		}
// 		return nil, errorToAPIError(err)
// 	}
// 	deletedNodeID = ok.Payload.ID
// 	return deletedNodeID, nil
// }

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
