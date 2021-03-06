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

// UpsertEventsReader is a Reader for the UpsertEvents structure.
type UpsertEventsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpsertEventsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpsertEventsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpsertEventsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpsertEventsOK creates a UpsertEventsOK with default headers values
func NewUpsertEventsOK() *UpsertEventsOK {
	return &UpsertEventsOK{}
}

/*UpsertEventsOK handles this case with default header values.

Ok
*/
type UpsertEventsOK struct {
	Payload []*models.Event
}

func (o *UpsertEventsOK) Error() string {
	return fmt.Sprintf("[PUT /v1/events][%d] upsertEventsOK  %+v", 200, o.Payload)
}

func (o *UpsertEventsOK) GetPayload() []*models.Event {
	return o.Payload
}

func (o *UpsertEventsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpsertEventsDefault creates a UpsertEventsDefault with default headers values
func NewUpsertEventsDefault(code int) *UpsertEventsDefault {
	return &UpsertEventsDefault{
		_statusCode: code,
	}
}

/*UpsertEventsDefault handles this case with default header values.

generic API error response
*/
type UpsertEventsDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the upsert events default response
func (o *UpsertEventsDefault) Code() int {
	return o._statusCode
}

func (o *UpsertEventsDefault) Error() string {
	return fmt.Sprintf("[PUT /v1/events][%d] UpsertEvents default  %+v", o._statusCode, o.Payload)
}

func (o *UpsertEventsDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *UpsertEventsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
