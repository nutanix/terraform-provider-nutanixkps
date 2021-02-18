package nutanixkpsclient

import (
	"os"
	"sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/client/application"
	"sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/client/helm"
	"sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"
	"sherlock-terraform-provider-nutanixkps/utils"

	cr "github.com/go-openapi/runtime/client"
)

// ApplicationV2List - Returns all ApplicationV2s
// func (c *Client) ApplicationV2List() ([]*models.ApplicationV2V2, *models.APIErrorPayload) {
// 	applicationListParams := application.NewApplicationV2ListParams()
// 	ok, err := c.CloudmgmtAPIClient.ApplicationV2.ApplicationV2List(applicationListParams, cr.BearerToken(c.Token))
// 	if err != nil {
// 		if apiErr, ok := err.(*application.ApplicationV2ListDefault); ok {
// 			return nil, apiErr.Payload
// 		}
// 		return nil, errorToAPIError(err)
// 	}
// 	return ok.Payload.ApplicationV2List, nil
// }

// ApplicationV2Get - Get a ApplicationV2
func (c *Client) ApplicationV2Get(applicationID string) (applicationModel *models.ApplicationV2, apiErr *models.APIErrorPayload) {
	applicationGetParams := application.NewApplicationGetV2Params()
	applicationGetParams.ID = applicationID
	ok, err := c.CloudmgmtAPIClient.Application.ApplicationGetV2(applicationGetParams, cr.BearerToken(c.Token))
	if err != nil {
		if apiErr, ok := err.(*application.ApplicationGetV2Default); ok {
			return nil, apiErr.Payload
		}
		return nil, errorToAPIError(err)
	}
	return ok.Payload, nil
}

// // HelmChartCreate - Create a ApplicationV2
func (c *Client) HelmApplicationCreate(application, helmChartFilename, helmChartValues string) (applicationID *string, apiErr *models.APIErrorPayload) {
	helmApplicationCreateParams := helm.NewHelmApplicationCreateParams()
	helmApplicationCreateParams.Application = application
	helmApplicationCreateParams.Chart = mustGetFile(helmChartFilename)
	if helmChartValues != "" {
		helmApplicationCreateParams.Values = mustGetFile(helmChartValues)
	}
	utils.PrintToJSON(helmApplicationCreateParams, "helmApplicationCreateParams: ")
	ok, err := c.CloudmgmtAPIClient.Helm.HelmApplicationCreate(helmApplicationCreateParams, cr.BearerToken(c.Token))
	if err != nil {
		if apiErr, ok := err.(*helm.HelmApplicationCreateDefault); ok {
			return nil, apiErr.Payload
		}
		return nil, errorToAPIError(err)
	}
	applicationID = ok.Payload.ID
	return applicationID, nil
}

func mustGetFile(path string) *os.File {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	return f
}

// // // ApplicationV2Create - Create a ApplicationV2
// func (c *Client) ApplicationV2Create(applicationModel *models.ApplicationV2) (applicationID *string, apiErr *models.APIErrorPayload) {
// 	applicationCreateParams := applicationV2.NewApplicationV2CreateParams()
// 	applicationCreateParams.Body = applicationModel
// 	ok, err := c.CloudmgmtAPIClient.ApplicationV2.ApplicationV2Create(applicationCreateParams, cr.BearerToken(c.Token))
// 	if err != nil {
// 		if apiErr, ok := err.(*application.ApplicationV2CreateDefault); ok {
// 			return nil, apiErr.Payload
// 		}
// 		return nil, errorToAPIError(err)
// 	}
// 	applicationID = ok.Payload.ID
// 	return applicationID, nil
// }

//todo
// // ServiceDomainUpdate - Update a ServiceDomain
// func (c *Client) ServiceDomainUpdate(svcDomainID string, serviceDomain *models.ServiceDomain) (deletedApplicationV2ID *string, apiErr *models.APIErrorPayload) {
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
// 	deletedApplicationV2ID = ok.Payload.ID
// 	return deletedApplicationV2ID, nil
// }

// ApplicationV2Delete - Delete a ApplicationV2
func (c *Client) ApplicationV2Delete(applicationID string) (deletedApplicationV2ID *string, apiErr *models.APIErrorPayload) {
	applicationDeleteParams := application.NewApplicationDeleteV2Params()
	applicationDeleteParams.ID = applicationID
	ok, err := c.CloudmgmtAPIClient.Application.ApplicationDeleteV2(applicationDeleteParams, cr.BearerToken(c.Token))
	if err != nil {
		if apiErr, ok := err.(*application.ApplicationDeleteV2Default); ok {
			return nil, apiErr.Payload
		}
		return nil, errorToAPIError(err)
	}
	deletedApplicationV2ID = ok.Payload.ID
	return deletedApplicationV2ID, nil
}
