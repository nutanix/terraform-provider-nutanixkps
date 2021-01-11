package nutanixkpsclient

import (
	"sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"
)

type KPSClient interface {
	Login(username, password *string) (token string, apiErr *models.APIErrorPayload)
	ServiceDomainList() (serviceDomains []*models.ServiceDomain, apiErr *models.APIErrorPayload)
	ServiceDomainCreate(serviceDomain *models.ServiceDomain) (serviceDomainID *string, apiErr *models.APIErrorPayload)
	ServiceDomainUpdate(svcDomainID string, serviceDomain *models.ServiceDomain) (serviceDomainID *string, apiErr *models.APIErrorPayload)
	ServiceDomainGet(svcDomainID string) (serviceDomain *models.ServiceDomain, apiErr *models.APIErrorPayload)
	ServiceDomainDelete(svcDomainID string) (serviceDomainID *string, apiErr *models.APIErrorPayload)
}
