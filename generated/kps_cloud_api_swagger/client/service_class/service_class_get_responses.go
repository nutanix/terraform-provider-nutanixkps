// Code generated by go-swagger; DO NOT EDIT.

package service_class

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"
)

// ServiceClassGetReader is a Reader for the ServiceClassGet structure.
type ServiceClassGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServiceClassGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServiceClassGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewServiceClassGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewServiceClassGetOK creates a ServiceClassGetOK with default headers values
func NewServiceClassGetOK() *ServiceClassGetOK {
	return &ServiceClassGetOK{}
}

/*ServiceClassGetOK handles this case with default header values.

Ok
*/
type ServiceClassGetOK struct {
	Payload *models.ServiceClass
}

func (o *ServiceClassGetOK) Error() string {
	return fmt.Sprintf("[GET /v1.0/serviceclasses/{svcClassId}][%d] serviceClassGetOK  %+v", 200, o.Payload)
}

func (o *ServiceClassGetOK) GetPayload() *models.ServiceClass {
	return o.Payload
}

func (o *ServiceClassGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceClass)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceClassGetDefault creates a ServiceClassGetDefault with default headers values
func NewServiceClassGetDefault(code int) *ServiceClassGetDefault {
	return &ServiceClassGetDefault{
		_statusCode: code,
	}
}

/*ServiceClassGetDefault handles this case with default header values.

generic API error response
*/
type ServiceClassGetDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the service class get default response
func (o *ServiceClassGetDefault) Code() int {
	return o._statusCode
}

func (o *ServiceClassGetDefault) Error() string {
	return fmt.Sprintf("[GET /v1.0/serviceclasses/{svcClassId}][%d] ServiceClassGet default  %+v", o._statusCode, o.Payload)
}

func (o *ServiceClassGetDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *ServiceClassGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
