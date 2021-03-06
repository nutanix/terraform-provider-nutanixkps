// Code generated by go-swagger; DO NOT EDIT.

package service_binding

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"
)

// ServiceBindingGetReader is a Reader for the ServiceBindingGet structure.
type ServiceBindingGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServiceBindingGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServiceBindingGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewServiceBindingGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewServiceBindingGetOK creates a ServiceBindingGetOK with default headers values
func NewServiceBindingGetOK() *ServiceBindingGetOK {
	return &ServiceBindingGetOK{}
}

/*ServiceBindingGetOK handles this case with default header values.

Ok
*/
type ServiceBindingGetOK struct {
	Payload *models.ServiceBinding
}

func (o *ServiceBindingGetOK) Error() string {
	return fmt.Sprintf("[GET /v1.0/servicebindings/{svcBindingId}][%d] serviceBindingGetOK  %+v", 200, o.Payload)
}

func (o *ServiceBindingGetOK) GetPayload() *models.ServiceBinding {
	return o.Payload
}

func (o *ServiceBindingGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceBinding)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBindingGetDefault creates a ServiceBindingGetDefault with default headers values
func NewServiceBindingGetDefault(code int) *ServiceBindingGetDefault {
	return &ServiceBindingGetDefault{
		_statusCode: code,
	}
}

/*ServiceBindingGetDefault handles this case with default header values.

generic API error response
*/
type ServiceBindingGetDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the service binding get default response
func (o *ServiceBindingGetDefault) Code() int {
	return o._statusCode
}

func (o *ServiceBindingGetDefault) Error() string {
	return fmt.Sprintf("[GET /v1.0/servicebindings/{svcBindingId}][%d] ServiceBindingGet default  %+v", o._statusCode, o.Payload)
}

func (o *ServiceBindingGetDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *ServiceBindingGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
