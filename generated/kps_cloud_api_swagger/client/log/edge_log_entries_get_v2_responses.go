// Code generated by go-swagger; DO NOT EDIT.

package log

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"
)

// EdgeLogEntriesGetV2Reader is a Reader for the EdgeLogEntriesGetV2 structure.
type EdgeLogEntriesGetV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeLogEntriesGetV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeLogEntriesGetV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewEdgeLogEntriesGetV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEdgeLogEntriesGetV2OK creates a EdgeLogEntriesGetV2OK with default headers values
func NewEdgeLogEntriesGetV2OK() *EdgeLogEntriesGetV2OK {
	return &EdgeLogEntriesGetV2OK{}
}

/*EdgeLogEntriesGetV2OK handles this case with default header values.

Ok
*/
type EdgeLogEntriesGetV2OK struct {
	Payload *models.LogEntriesListPayload
}

func (o *EdgeLogEntriesGetV2OK) Error() string {
	return fmt.Sprintf("[GET /v1.0/logs/edges/{id}][%d] edgeLogEntriesGetV2OK  %+v", 200, o.Payload)
}

func (o *EdgeLogEntriesGetV2OK) GetPayload() *models.LogEntriesListPayload {
	return o.Payload
}

func (o *EdgeLogEntriesGetV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LogEntriesListPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeLogEntriesGetV2Default creates a EdgeLogEntriesGetV2Default with default headers values
func NewEdgeLogEntriesGetV2Default(code int) *EdgeLogEntriesGetV2Default {
	return &EdgeLogEntriesGetV2Default{
		_statusCode: code,
	}
}

/*EdgeLogEntriesGetV2Default handles this case with default header values.

generic API error response
*/
type EdgeLogEntriesGetV2Default struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the edge log entries get v2 default response
func (o *EdgeLogEntriesGetV2Default) Code() int {
	return o._statusCode
}

func (o *EdgeLogEntriesGetV2Default) Error() string {
	return fmt.Sprintf("[GET /v1.0/logs/edges/{id}][%d] EdgeLogEntriesGetV2 default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeLogEntriesGetV2Default) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *EdgeLogEntriesGetV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
