// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"
)

// WsMessagingOnCreateCloudCredsReader is a Reader for the WsMessagingOnCreateCloudCreds structure.
type WsMessagingOnCreateCloudCredsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WsMessagingOnCreateCloudCredsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWsMessagingOnCreateCloudCredsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewWsMessagingOnCreateCloudCredsOK creates a WsMessagingOnCreateCloudCredsOK with default headers values
func NewWsMessagingOnCreateCloudCredsOK() *WsMessagingOnCreateCloudCredsOK {
	return &WsMessagingOnCreateCloudCredsOK{}
}

/*WsMessagingOnCreateCloudCredsOK handles this case with default header values.

Ok
*/
type WsMessagingOnCreateCloudCredsOK struct {
	Payload *models.ResponseBase
}

func (o *WsMessagingOnCreateCloudCredsOK) Error() string {
	return fmt.Sprintf("[POST /v1/wsdocs/onCreateCloudCreds][%d] wsMessagingOnCreateCloudCredsOK  %+v", 200, o.Payload)
}

func (o *WsMessagingOnCreateCloudCredsOK) GetPayload() *models.ResponseBase {
	return o.Payload
}

func (o *WsMessagingOnCreateCloudCredsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseBase)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
