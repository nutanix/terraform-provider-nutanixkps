package nutanixkpsclient

import (
	"fmt"
	"log"
	"sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/client/storage_profile"
	"sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"
	"sherlock-terraform-provider-nutanixkps/utils"

	runtime "github.com/go-openapi/runtime/client"
) // 	"sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"

// 	"sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/client/auth"

// 	runtime "github.com/go-openapi/runtime/client"

// StorageProfileList - Returns all StorageProfiles
func (c *Client) StorageProfileList(SvcDomainID string) ([]*models.StorageProfile, *models.APIErrorPayload) {
	storageProfileListParams := storage_profile.NewSvcDomainGetStorageProfilesParams()
	storageProfileListParams.SvcDomainID = SvcDomainID
	ok, err := c.CloudmgmtAPIClient.StorageProfile.SvcDomainGetStorageProfiles(storageProfileListParams, runtime.BearerToken(c.Token))
	if err != nil {
		if apiErr, ok := err.(*storage_profile.SvcDomainGetStorageProfilesDefault); ok {
			return nil, apiErr.Payload
		}
		return nil, errorToAPIError(err)
	}
	return ok.Payload.StorageProfileList, nil
}

// StorageProfileGet - Get a StorageProfile
func (c *Client) StorageProfileGet(svcDomainID, storageProfileID string) (storageProfile *models.StorageProfile, apiErr *models.APIErrorPayload) {
	storageProfiles, err := c.StorageProfileList(svcDomainID)
	if err != nil {
		return nil, err
	}
	for _, sp := range storageProfiles {
		if sp.ID == storageProfileID {
			return sp, nil
		}
	}
	return nil, errorToAPIError(fmt.Errorf("no storage profile found with id: %s in service domain", storageProfileID))
}

// StorageProfileCreate - Create a StorageProfile
func (c *Client) StorageProfileCreate(SvcDomainID string, storageProfile *models.StorageProfile) (storageProfileID *string, apiErr *models.APIErrorPayload) {
	storageProfileCreateParams := storage_profile.NewStorageProfileCreateParams()
	storageProfileCreateParams.SvcDomainID = SvcDomainID
	storageProfileCreateParams.Body = storageProfile
	ok, err := c.CloudmgmtAPIClient.StorageProfile.StorageProfileCreate(storageProfileCreateParams, runtime.BearerToken(c.Token))
	if err != nil {
		if apiErr, ok := err.(*storage_profile.StorageProfileCreateDefault); ok {
			return nil, apiErr.Payload
		}
		return nil, errorToAPIError(err)
	}
	log.Printf("[DEBUG] StorageProfileCreate error was nil")
	utils.PrintToJSON(ok, "ok: ")
	utils.PrintToJSON(ok.Payload, "ok.Payload: ")
	storageProfileID = ok.Payload.ID
	return storageProfileID, nil
}

// // StorageProfileUpdate - Update a StorageProfile
// func (c *Client) StorageProfileUpdate(storageProfileID string, storageProfile *models.StorageProfile) (storageProfileID *string, apiErr *models.APIErrorPayload) {
// 	storageProfileUpdateParams := storage_profile.NewStorageProfileUpdateParams()
// 	storageProfileUpdateParams.SvcDomainID = storageProfileID
// 	storageProfileUpdateParams.Body = storageProfile
// 	ok, err := c.CloudmgmtAPIClient.StorageProfile.StorageProfileUpdate(storageProfileUpdateParams, runtime.BearerToken(c.Token))
// 	if err != nil {
// 		if apiErr, ok := err.(*storage_profile.StorageProfileUpdateDefault); ok {
// 			return nil, apiErr.Payload
// 		}
// 		return nil, errorToAPIError(err)
// 	}
// 	storageProfileID = ok.Payload.ID
// 	return storageProfileID, nil
// }

// StorageProfileDelete - Delete a StorageProfile
// func (c *Client) StorageProfileDelete(storageProfileID string) (storageProfileID *string, apiErr *models.APIErrorPayload) {
// 	// storageProfileDeleteParams := storage_profile.NewStorageProfileDeleteParams()
// 	storageProfileDeleteParams := storage_profile.
// 	storageProfileDeleteParams.SvcDomainID = storageProfileID
// 	ok, err := c.CloudmgmtAPIClient.StorageProfile.StorageProfileDelete(storageProfileDeleteParams, runtime.BearerToken(c.Token))
// 	if err != nil {
// 		if apiErr, ok := err.(*storage_profile.StorageProfileDeleteDefault); ok {
// 			return nil, apiErr.Payload
// 		}
// 		return nil, errorToAPIError(err)
// 	}
// 	storageProfileID = ok.Payload.ID
// 	return storageProfileID, nil

// }
