package nutanixkpsclient

import (
	"terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/client/user"
	"terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"

	cr "github.com/go-openapi/runtime/client"
)

// UserList - Returns all Users
func (c *Client) UserList() ([]*models.User, *models.APIErrorPayload) {
	userListParams := user.NewUserListV2Params()
	ok, err := c.CloudmgmtAPIClient.User.UserListV2(userListParams, cr.BearerToken(c.Token))
	if err != nil {
		if apiErr, ok := err.(*user.UserListV2Default); ok {
			return nil, apiErr.Payload
		}
		return nil, errorToAPIError(err)
	}
	return ok.Payload.UserList, nil
}

// UserGet - Get a User
func (c *Client) UserGet(userID string) (userModel *models.User, apiErr *models.APIErrorPayload) {
	userGetParams := user.NewUserGetV2Params()
	userGetParams.ID = userID
	ok, err := c.CloudmgmtAPIClient.User.UserGetV2(userGetParams, cr.BearerToken(c.Token))
	if err != nil {
		if apiErr, ok := err.(*user.UserGetV2Default); ok {
			return nil, apiErr.Payload
		}
		return nil, errorToAPIError(err)
	}
	return ok.Payload, nil
}
