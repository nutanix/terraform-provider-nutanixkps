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

// ApplicationLogEntriesGetV2Reader is a Reader for the ApplicationLogEntriesGetV2 structure.
type ApplicationLogEntriesGetV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ApplicationLogEntriesGetV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewApplicationLogEntriesGetV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewApplicationLogEntriesGetV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewApplicationLogEntriesGetV2OK creates a ApplicationLogEntriesGetV2OK with default headers values
func NewApplicationLogEntriesGetV2OK() *ApplicationLogEntriesGetV2OK {
	return &ApplicationLogEntriesGetV2OK{}
}

/*ApplicationLogEntriesGetV2OK handles this case with default header values.

Ok
*/
type ApplicationLogEntriesGetV2OK struct {
	Payload *models.LogEntriesListPayload
}

func (o *ApplicationLogEntriesGetV2OK) Error() string {
	return fmt.Sprintf("[GET /v1.0/logs/applications/{id}][%d] applicationLogEntriesGetV2OK  %+v", 200, o.Payload)
}

func (o *ApplicationLogEntriesGetV2OK) GetPayload() *models.LogEntriesListPayload {
	return o.Payload
}

func (o *ApplicationLogEntriesGetV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LogEntriesListPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewApplicationLogEntriesGetV2Default creates a ApplicationLogEntriesGetV2Default with default headers values
func NewApplicationLogEntriesGetV2Default(code int) *ApplicationLogEntriesGetV2Default {
	return &ApplicationLogEntriesGetV2Default{
		_statusCode: code,
	}
}

/*ApplicationLogEntriesGetV2Default handles this case with default header values.

generic API error response
*/
type ApplicationLogEntriesGetV2Default struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the application log entries get v2 default response
func (o *ApplicationLogEntriesGetV2Default) Code() int {
	return o._statusCode
}

func (o *ApplicationLogEntriesGetV2Default) Error() string {
	return fmt.Sprintf("[GET /v1.0/logs/applications/{id}][%d] ApplicationLogEntriesGetV2 default  %+v", o._statusCode, o.Payload)
}

func (o *ApplicationLogEntriesGetV2Default) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *ApplicationLogEntriesGetV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
