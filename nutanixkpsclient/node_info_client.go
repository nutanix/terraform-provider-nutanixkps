package nutanixkpsclient

import (
	"terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/client/node_info"
	"terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"

	runtime "github.com/go-openapi/runtime/client"
)

// NodeInfoList - Returns all NodeInfos
func (c *Client) NodeInfoList() ([]*models.NodeInfo, *models.APIErrorPayload) {
	nodeInfoListParams := node_info.NewNodeInfoListParams()
	ok, err := c.CloudmgmtAPIClient.NodeInfo.NodeInfoList(nodeInfoListParams, runtime.BearerToken(c.Token))
	if err != nil {
		if apiErr, ok := err.(*node_info.NodeInfoListDefault); ok {
			return nil, apiErr.Payload
		}
		return nil, errorToAPIError(err)
	}
	return ok.Payload.NodeInfoList, nil
}

// NodeInfoGet - Get a NodeInfo
func (c *Client) NodeInfoGet(nodeID string) (nodeModel *models.NodeInfo, apiErr *models.APIErrorPayload) {
	nodeInfoGetParams := node_info.NewNodeInfoGetParams()
	nodeInfoGetParams.NodeID = nodeID
	ok, err := c.CloudmgmtAPIClient.NodeInfo.NodeInfoGet(nodeInfoGetParams, runtime.BearerToken(c.Token))
	if err != nil {
		if apiErr, ok := err.(*node_info.NodeInfoGetDefault); ok {
			return nil, apiErr.Payload
		}
		return nil, errorToAPIError(err)
	}
	return ok.Payload, nil
}
