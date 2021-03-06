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

// WsMessagingOnUpdateEdgeReader is a Reader for the WsMessagingOnUpdateEdge structure.
type WsMessagingOnUpdateEdgeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WsMessagingOnUpdateEdgeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWsMessagingOnUpdateEdgeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewWsMessagingOnUpdateEdgeOK creates a WsMessagingOnUpdateEdgeOK with default headers values
func NewWsMessagingOnUpdateEdgeOK() *WsMessagingOnUpdateEdgeOK {
	return &WsMessagingOnUpdateEdgeOK{}
}

/*WsMessagingOnUpdateEdgeOK handles this case with default header values.

Ok
*/
type WsMessagingOnUpdateEdgeOK struct {
	Payload *models.ResponseBase
}

func (o *WsMessagingOnUpdateEdgeOK) Error() string {
	return fmt.Sprintf("[POST /v1/wsdocs/onUpdateEdge][%d] wsMessagingOnUpdateEdgeOK  %+v", 200, o.Payload)
}

func (o *WsMessagingOnUpdateEdgeOK) GetPayload() *models.ResponseBase {
	return o.Payload
}

func (o *WsMessagingOnUpdateEdgeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseBase)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
