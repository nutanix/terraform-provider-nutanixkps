package nutanixkpsclient

import (
	"terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/client/project"
	"terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"

	cr "github.com/go-openapi/runtime/client"
)

// ProjectV2List - Returns all ProjectV2s
// func (c *Client) ProjectV2List() ([]*models.ProjectV2V2, *models.APIErrorPayload) {
// 	projectListParams := project.NewProjectV2ListParams()
// 	ok, err := c.CloudmgmtAPIClient.ProjectV2.ProjectV2List(projectListParams, cr.BearerToken(c.Token))
// 	if err != nil {
// 		if apiErr, ok := err.(*project.ProjectV2ListDefault); ok {
// 			return nil, apiErr.Payload
// 		}
// 		return nil, errorToAPIError(err)
// 	}
// 	return ok.Payload.ProjectV2List, nil
// }

// ProjectV2Get - Get a ProjectV2
func (c *Client) ProjectV2Get(projectID string) (projectModel *models.Project, apiErr *models.APIErrorPayload) {
	projectGetParams := project.NewProjectGetV2Params()
	projectGetParams.ProjectID = projectID
	ok, err := c.CloudmgmtAPIClient.Project.ProjectGetV2(projectGetParams, cr.BearerToken(c.Token))
	if err != nil {
		if apiErr, ok := err.(*project.ProjectGetV2Default); ok {
			return nil, apiErr.Payload
		}
		return nil, errorToAPIError(err)
	}
	return ok.Payload, nil
}

// // ProjectV2Create - Create a ProjectV2
func (c *Client) ProjectV2Create(projectModel *models.Project) (projectID *string, apiErr *models.APIErrorPayload) {
	projectCreateParams := project.NewProjectCreateV2Params()
	projectCreateParams.Doc = projectModel
	ok, err := c.CloudmgmtAPIClient.Project.ProjectCreateV2(projectCreateParams, cr.BearerToken(c.Token))
	if err != nil {
		if apiErr, ok := err.(*project.ProjectCreateV2Default); ok {
			return nil, apiErr.Payload
		}
		return nil, errorToAPIError(err)
	}
	projectID = ok.Payload.ID
	return projectID, nil
}

//todo
// // ServiceDomainUpdate - Update a ServiceDomain
// func (c *Client) ServiceDomainUpdate(svcDomainID string, serviceDomain *models.ServiceDomain) (deletedProjectV2ID *string, apiErr *models.APIErrorPayload) {
// 	serviceDomainUpdateParams := service_domain.NewServiceDomainUpdateParams()
// 	serviceDomainUpdateParams.SvcDomainID = svcDomainID
// 	serviceDomainUpdateParams.Body = serviceDomain
// 	ok, err := c.CloudmgmtAPIClient.ServiceDomain.ServiceDomainUpdate(serviceDomainUpdateParams, cr.BearerToken(c.Token))
// 	if err != nil {
// 		if apiErr, ok := err.(*service_domain.ServiceDomainUpdateDefault); ok {
// 			return nil, apiErr.Payload
// 		}
// 		return nil, errorToAPIError(err)
// 	}
// 	deletedProjectV2ID = ok.Payload.ID
// 	return deletedProjectV2ID, nil
// }

// ProjectV2Delete - Delete a ProjectV2
func (c *Client) ProjectV2Delete(projectID string) (deletedProjectV2ID *string, apiErr *models.APIErrorPayload) {
	projectDeleteParams := project.NewProjectDeleteV2Params()
	projectDeleteParams.ID = projectID
	ok, err := c.CloudmgmtAPIClient.Project.ProjectDeleteV2(projectDeleteParams, cr.BearerToken(c.Token))
	if err != nil {
		if apiErr, ok := err.(*project.ProjectDeleteV2Default); ok {
			return nil, apiErr.Payload
		}
		return nil, errorToAPIError(err)
	}
	deletedProjectV2ID = ok.Payload.ID
	return deletedProjectV2ID, nil
}
