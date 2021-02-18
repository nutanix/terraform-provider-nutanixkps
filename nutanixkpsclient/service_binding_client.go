package nutanixkpsclient

import (
	"sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/client/service_binding"
	"sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"

	cr "github.com/go-openapi/runtime/client"
)

// ServiceBindingList - Returns all ServiceBindings
// func (c *Client) ServiceBindingList() ([]*models.ServiceBinding, *models.APIErrorPayload) {
// 	serviceBindingListParams := service_binding.NewServiceBindingListParams()
// 	ok, err := c.CloudmgmtAPIClient.ServiceBinding.ServiceBindingList(serviceBindingListParams, cr.BearerToken(c.Token))
// 	if err != nil {
// 		if apiErr, ok := err.(*service_binding.ServiceBindingListDefault); ok {
// 			return nil, apiErr.Payload
// 		}
// 		return nil, errorToAPIError(err)
// 	}
// 	return ok.Payload.ServiceBindingList, nil
// }

// ServiceBindingGet - Get a ServiceBinding
func (c *Client) ServiceBindingGet(serviceBindingID string) (serviceBindingModel *models.ServiceBinding, apiErr *models.APIErrorPayload) {
	serviceBindingGetParams := service_binding.NewServiceBindingGetParams()
	serviceBindingGetParams.SvcBindingID = serviceBindingID
	ok, err := c.CloudmgmtAPIClient.ServiceBinding.ServiceBindingGet(serviceBindingGetParams, cr.BearerToken(c.Token))
	if err != nil {
		if apiErr, ok := err.(*service_binding.ServiceBindingGetDefault); ok {
			return nil, apiErr.Payload
		}
		return nil, errorToAPIError(err)
	}
	return ok.Payload, nil
}

// // ServiceBindingCreate - Create a ServiceBinding
func (c *Client) ServiceBindingCreate(serviceBindingModel *models.ServiceBindingParam) (serviceBindingID *string, apiErr *models.APIErrorPayload) {
	serviceBindingCreateParams := service_binding.NewServiceBindingCreateParams()
	serviceBindingCreateParams.Body = serviceBindingModel
	ok, err := c.CloudmgmtAPIClient.ServiceBinding.ServiceBindingCreate(serviceBindingCreateParams, cr.BearerToken(c.Token))
	if err != nil {
		if apiErr, ok := err.(*service_binding.ServiceBindingCreateDefault); ok {
			return nil, apiErr.Payload
		}
		return nil, errorToAPIError(err)
	}
	serviceBindingID = ok.Payload.ID
	return serviceBindingID, nil
}

//todo
// // ServiceDomainUpdate - Update a ServiceDomain
// func (c *Client) ServiceDomainUpdate(svcDomainID string, serviceDomain *models.ServiceDomain) (deletedServiceBindingID *string, apiErr *models.APIErrorPayload) {
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
// 	deletedServiceBindingID = ok.Payload.ID
// 	return deletedServiceBindingID, nil
// }

// ServiceBindingDelete - Delete a ServiceBinding
func (c *Client) ServiceBindingDelete(serviceBindingID string) (deletedServiceBindingID *string, apiErr *models.APIErrorPayload) {
	serviceBindingDeleteParams := service_binding.NewServiceBindingDeleteParams()
	serviceBindingDeleteParams.SvcBindingID = serviceBindingID
	ok, err := c.CloudmgmtAPIClient.ServiceBinding.ServiceBindingDelete(serviceBindingDeleteParams, cr.BearerToken(c.Token))
	if err != nil {
		if apiErr, ok := err.(*service_binding.ServiceBindingDeleteDefault); ok {
			return nil, apiErr.Payload
		}
		return nil, errorToAPIError(err)
	}
	deletedServiceBindingID = ok.Payload.ID
	return deletedServiceBindingID, nil
}
