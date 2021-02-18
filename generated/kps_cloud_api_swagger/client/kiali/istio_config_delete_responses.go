// Code generated by go-swagger; DO NOT EDIT.

package kiali

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"
)

// IstioConfigDeleteReader is a Reader for the IstioConfigDelete structure.
type IstioConfigDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IstioConfigDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 400:
		result := NewIstioConfigDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIstioConfigDeleteBadRequest creates a IstioConfigDeleteBadRequest with default headers values
func NewIstioConfigDeleteBadRequest() *IstioConfigDeleteBadRequest {
	return &IstioConfigDeleteBadRequest{}
}

/*IstioConfigDeleteBadRequest handles this case with default header values.

generic API error response
*/
type IstioConfigDeleteBadRequest struct {
	Payload *models.APIErrorPayload
}

func (o *IstioConfigDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /v1.0/kiali/namespaces/{namespace}/istio/{object_type}/{object}][%d] istioConfigDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *IstioConfigDeleteBadRequest) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *IstioConfigDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
