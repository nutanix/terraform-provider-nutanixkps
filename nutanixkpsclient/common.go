package nutanixkpsclient

import (
	"errors"
	"sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"
)

func errorToAPIError(err error) *models.APIErrorPayload {
	errString := err.Error()
	return &models.APIErrorPayload{Message: &errString}
}

func APIErrorToError(apiErrorPayload *models.APIErrorPayload) (err error) {
	return errors.New(*apiErrorPayload.Message)
}
