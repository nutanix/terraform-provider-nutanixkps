package nutanixkpsclient

import (
	"terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/client/service_instance"
	"terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"

	cr "github.com/go-openapi/runtime/client"
)

// ServiceInstanceList - Returns all ServiceInstances
// func (c *Client) ServiceInstanceList() ([]*models.ServiceInstance, *models.APIErrorPayload) {
// 	serviceInstanceListParams := service_instance.NewServiceInstanceListParams()
// 	ok, err := c.CloudmgmtAPIClient.ServiceInstance.ServiceInstanceList(serviceInstanceListParams, cr.BearerToken(c.Token))
// 	if err != nil {
// 		if apiErr, ok := err.(*service_instance.ServiceInstanceListDefault); ok {
// 			return nil, apiErr.Payload
// 		}
// 		return nil, errorToAPIError(err)
// 	}
// 	return ok.Payload.ServiceInstanceList, nil
// }

// ServiceInstanceGet - Get a ServiceInstance
func (c *Client) ServiceInstanceGet(serviceInstanceID string) (serviceInstanceModel *models.ServiceInstance, apiErr *models.APIErrorPayload) {
	serviceInstanceGetParams := service_instance.NewServiceInstanceGetParams()
	serviceInstanceGetParams.SvcInstanceID = serviceInstanceID
	ok, err := c.CloudmgmtAPIClient.ServiceInstance.ServiceInstanceGet(serviceInstanceGetParams, cr.BearerToken(c.Token))
	if err != nil {
		if apiErr, ok := err.(*service_instance.ServiceInstanceGetDefault); ok {
			return nil, apiErr.Payload
		}
		return nil, errorToAPIError(err)
	}
	return ok.Payload, nil
}

// // ServiceInstanceCreate - Create a ServiceInstance
func (c *Client) ServiceInstanceCreate(serviceInstanceModel *models.ServiceInstanceParam) (serviceInstanceID *string, apiErr *models.APIErrorPayload) {
	serviceInstanceCreateParams := service_instance.NewServiceInstanceCreateParams()
	serviceInstanceCreateParams.Body = serviceInstanceModel
	ok, err := c.CloudmgmtAPIClient.ServiceInstance.ServiceInstanceCreate(serviceInstanceCreateParams, cr.BearerToken(c.Token))
	if err != nil {
		if apiErr, ok := err.(*service_instance.ServiceInstanceCreateDefault); ok {
			return nil, apiErr.Payload
		}
		return nil, errorToAPIError(err)
	}
	serviceInstanceID = ok.Payload.ID
	return serviceInstanceID, nil
}

//todo
// // ServiceDomainUpdate - Update a ServiceDomain
// func (c *Client) ServiceDomainUpdate(svcDomainID string, serviceDomain *models.ServiceDomain) (deletedServiceInstanceID *string, apiErr *models.APIErrorPayload) {
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
// 	deletedServiceInstanceID = ok.Payload.ID
// 	return deletedServiceInstanceID, nil
// }

// ServiceInstanceDelete - Delete a ServiceInstance
func (c *Client) ServiceInstanceDelete(serviceInstanceID string) (deletedServiceInstanceID *string, apiErr *models.APIErrorPayload) {
	serviceInstanceDeleteParams := service_instance.NewServiceInstanceDeleteParams()
	serviceInstanceDeleteParams.SvcInstanceID = serviceInstanceID
	ok, err := c.CloudmgmtAPIClient.ServiceInstance.ServiceInstanceDelete(serviceInstanceDeleteParams, cr.BearerToken(c.Token))
	if err != nil {
		if apiErr, ok := err.(*service_instance.ServiceInstanceDeleteDefault); ok {
			return nil, apiErr.Payload
		}
		return nil, errorToAPIError(err)
	}
	deletedServiceInstanceID = ok.Payload.ID
	return deletedServiceInstanceID, nil
}
