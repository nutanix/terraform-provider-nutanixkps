// Code generated by go-swagger; DO NOT EDIT.

package application

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"
)

// ApplicationDeleteV2Reader is a Reader for the ApplicationDeleteV2 structure.
type ApplicationDeleteV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ApplicationDeleteV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewApplicationDeleteV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewApplicationDeleteV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewApplicationDeleteV2OK creates a ApplicationDeleteV2OK with default headers values
func NewApplicationDeleteV2OK() *ApplicationDeleteV2OK {
	return &ApplicationDeleteV2OK{}
}

/*ApplicationDeleteV2OK handles this case with default header values.

Ok
*/
type ApplicationDeleteV2OK struct {
	Payload *models.DeleteDocumentResponseV2
}

func (o *ApplicationDeleteV2OK) Error() string {
	return fmt.Sprintf("[DELETE /v1.0/applications/{id}][%d] applicationDeleteV2OK  %+v", 200, o.Payload)
}

func (o *ApplicationDeleteV2OK) GetPayload() *models.DeleteDocumentResponseV2 {
	return o.Payload
}

func (o *ApplicationDeleteV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeleteDocumentResponseV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewApplicationDeleteV2Default creates a ApplicationDeleteV2Default with default headers values
func NewApplicationDeleteV2Default(code int) *ApplicationDeleteV2Default {
	return &ApplicationDeleteV2Default{
		_statusCode: code,
	}
}

/*ApplicationDeleteV2Default handles this case with default header values.

generic API error response
*/
type ApplicationDeleteV2Default struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the application delete v2 default response
func (o *ApplicationDeleteV2Default) Code() int {
	return o._statusCode
}

func (o *ApplicationDeleteV2Default) Error() string {
	return fmt.Sprintf("[DELETE /v1.0/applications/{id}][%d] ApplicationDeleteV2 default  %+v", o._statusCode, o.Payload)
}

func (o *ApplicationDeleteV2Default) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *ApplicationDeleteV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
