package nutanixkpsclient

import (
	"sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/client/service_class"
	"sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"

	cr "github.com/go-openapi/runtime/client"
)

// ServiceClassList - Returns all ServiceClasses
func (c *Client) ServiceClassList() ([]*models.ServiceClass, *models.APIErrorPayload) {
	serviceClassListParams := service_class.NewServiceClassListParams()
	ok, err := c.CloudmgmtAPIClient.ServiceClass.ServiceClassList(serviceClassListParams, cr.BearerToken(c.Token))
	if err != nil {
		if apiErr, ok := err.(*service_class.ServiceClassListDefault); ok {
			return nil, apiErr.Payload
		}
		return nil, errorToAPIError(err)
	}
	return ok.Payload.SvcClassList, nil
}

// ServiceClassGet - Get a ServiceClass
func (c *Client) ServiceClassGet(serviceClassID string) (serviceClassModel *models.ServiceClass, apiErr *models.APIErrorPayload) {
	serviceClassGetParams := service_class.NewServiceClassGetParams()
	serviceClassGetParams.SvcClassID = serviceClassID
	ok, err := c.CloudmgmtAPIClient.ServiceClass.ServiceClassGet(serviceClassGetParams, cr.BearerToken(c.Token))
	if err != nil {
		if apiErr, ok := err.(*service_class.ServiceClassGetDefault); ok {
			return nil, apiErr.Payload
		}
		return nil, errorToAPIError(err)
	}
	return ok.Payload, nil
}

// // ServiceClassCreate - Create a ServiceClass
// func (c *Client) ServiceClassCreate(serviceClassModel *models.Project) (serviceClassID *string, apiErr *models.APIErrorPayload) {
// 	serviceClassCreateParams :=service_class.NewProjectCreateV2Params()
// 	serviceClassCreateParams.Doc = serviceClassModel
// 	ok, err := c.CloudmgmtAPIClient.Project.ProjectCreateV2(serviceClassCreateParams, cr.BearerToken(c.Token))
// 	if err != nil {
// 		if apiErr, ok := err.(*serviceClass.ProjectCreateV2Default); ok {
// 			return nil, apiErr.Payload
// 		}
// 		return nil, errorToAPIError(err)
// 	}
// 	serviceClassID = ok.Payload.ID
// 	return serviceClassID, nil
// }

//todo
// // ServiceDomainUpdate - Update a ServiceDomain
// func (c *Client) ServiceDomainUpdate(svcDomainID string, serviceDomain *models.ServiceDomain) (deletedServiceClassID *string, apiErr *models.APIErrorPayload) {
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
// 	deletedServiceClassID = ok.Payload.ID
// 	return deletedServiceClassID, nil
// }

// ServiceClassDelete - Delete a ServiceClass
// func (c *Client) ServiceClassDelete(serviceClassID string) (deletedServiceClassID *string, apiErr *models.APIErrorPayload) {
// 	serviceClassDeleteParams :=service_class.NewProjectDeleteV2Params()
// 	serviceClassDeleteParams.ID = serviceClassID
// 	ok, err := c.CloudmgmtAPIClient.Project.ProjectDeleteV2(serviceClassDeleteParams, cr.BearerToken(c.Token))
// 	if err != nil {
// 		if apiErr, ok := err.(*serviceClass.ProjectDeleteV2Default); ok {
// 			return nil, apiErr.Payload
// 		}
// 		return nil, errorToAPIError(err)
// 	}
// 	deletedServiceClassID = ok.Payload.ID
// 	return deletedServiceClassID, nil
// }
