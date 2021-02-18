package nutanixkpsclient

import (
	"sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"

	"sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/client/auth"
	"sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/client/service_domain"

	runtime "github.com/go-openapi/runtime/client"
)

func (c *Client) Login(username, password *string) (string, *models.APIErrorPayload) {
	loginParams := auth.NewLoginCallV2Params()
	loginParams.Request = &models.Credential{Email: username, Password: password}
	ok, err := c.CloudmgmtAPIClient.Auth.LoginCallV2(loginParams)
	if err != nil {
		apiErr, castOK := err.(*auth.LoginCallV2Default)
		if castOK {
			return "", apiErr.Payload
		}
		return "", errorToAPIError(err)
	}
	return *ok.Payload.Token, nil
}

// ServiceDomainList - Returns all ServiceDomains
func (c *Client) ServiceDomainList() ([]*models.ServiceDomain, *models.APIErrorPayload) {
	serviceDomainListParams := service_domain.NewServiceDomainListParams()
	ok, err := c.CloudmgmtAPIClient.ServiceDomain.ServiceDomainList(serviceDomainListParams, runtime.BearerToken(c.Token))
	if err != nil {
		if apiErr, ok := err.(*service_domain.ServiceDomainListDefault); ok {
			return nil, apiErr.Payload
		}
		return nil, errorToAPIError(err)
	}
	return ok.Payload.SvcDomainList, nil
}

// ServiceDomainGet - Get a ServiceDomain
func (c *Client) ServiceDomainGet(svcDomainID string) (serviceDomain *models.ServiceDomain, apiErr *models.APIErrorPayload) {
	serviceDomainGetParams := service_domain.NewServiceDomainGetParams()
	serviceDomainGetParams.SvcDomainID = svcDomainID
	ok, err := c.CloudmgmtAPIClient.ServiceDomain.ServiceDomainGet(serviceDomainGetParams, runtime.BearerToken(c.Token))
	if err != nil {
		if apiErr, ok := err.(*service_domain.ServiceDomainGetDefault); ok {
			return nil, apiErr.Payload
		}
		return nil, errorToAPIError(err)
	}
	return ok.Payload, nil
}

// ServiceDomainCreate - Create a ServiceDomain
func (c *Client) ServiceDomainCreate(serviceDomain *models.ServiceDomain) (serviceDomainID *string, apiErr *models.APIErrorPayload) {
	serviceDomainCreateParams := service_domain.NewServiceDomainCreateParams()
	serviceDomainCreateParams.Body = serviceDomain
	ok, err := c.CloudmgmtAPIClient.ServiceDomain.ServiceDomainCreate(serviceDomainCreateParams, runtime.BearerToken(c.Token))
	if err != nil {
		if apiErr, ok := err.(*service_domain.ServiceDomainCreateDefault); ok {
			return nil, apiErr.Payload
		}
		return nil, errorToAPIError(err)
	}
	serviceDomainID = ok.Payload.ID
	return serviceDomainID, nil
}

// ServiceDomainUpdate - Update a ServiceDomain
func (c *Client) ServiceDomainUpdate(svcDomainID string, serviceDomain *models.ServiceDomain) (serviceDomainID *string, apiErr *models.APIErrorPayload) {
	serviceDomainUpdateParams := service_domain.NewServiceDomainUpdateParams()
	serviceDomainUpdateParams.SvcDomainID = svcDomainID
	serviceDomainUpdateParams.Body = serviceDomain
	ok, err := c.CloudmgmtAPIClient.ServiceDomain.ServiceDomainUpdate(serviceDomainUpdateParams, runtime.BearerToken(c.Token))
	if err != nil {
		if apiErr, ok := err.(*service_domain.ServiceDomainUpdateDefault); ok {
			return nil, apiErr.Payload
		}
		return nil, errorToAPIError(err)
	}
	serviceDomainID = ok.Payload.ID
	return serviceDomainID, nil
}

// ServiceDomainDelete - Delete a ServiceDomain
func (c *Client) ServiceDomainDelete(svcDomainID string) (serviceDomainID *string, apiErr *models.APIErrorPayload) {
	serviceDomainDeleteParams := service_domain.NewServiceDomainDeleteParams()
	serviceDomainDeleteParams.SvcDomainID = svcDomainID
	ok, err := c.CloudmgmtAPIClient.ServiceDomain.ServiceDomainDelete(serviceDomainDeleteParams, runtime.BearerToken(c.Token))
	if err != nil {
		if apiErr, ok := err.(*service_domain.ServiceDomainDeleteDefault); ok {
			return nil, apiErr.Payload
		}
		return nil, errorToAPIError(err)
	}
	serviceDomainID = ok.Payload.ID
	return serviceDomainID, nil

}
