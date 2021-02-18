// Code generated by go-swagger; DO NOT EDIT.

package service_domain

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"
)

// ServiceDomainListReader is a Reader for the ServiceDomainList structure.
type ServiceDomainListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServiceDomainListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServiceDomainListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewServiceDomainListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewServiceDomainListOK creates a ServiceDomainListOK with default headers values
func NewServiceDomainListOK() *ServiceDomainListOK {
	return &ServiceDomainListOK{}
}

/*ServiceDomainListOK handles this case with default header values.

Ok
*/
type ServiceDomainListOK struct {
	Payload *models.ServiceDomainListPayload
}

func (o *ServiceDomainListOK) Error() string {
	return fmt.Sprintf("[GET /v1.0/servicedomains][%d] serviceDomainListOK  %+v", 200, o.Payload)
}

func (o *ServiceDomainListOK) GetPayload() *models.ServiceDomainListPayload {
	return o.Payload
}

func (o *ServiceDomainListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceDomainListPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceDomainListDefault creates a ServiceDomainListDefault with default headers values
func NewServiceDomainListDefault(code int) *ServiceDomainListDefault {
	return &ServiceDomainListDefault{
		_statusCode: code,
	}
}

/*ServiceDomainListDefault handles this case with default header values.

generic API error response
*/
type ServiceDomainListDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the service domain list default response
func (o *ServiceDomainListDefault) Code() int {
	return o._statusCode
}

func (o *ServiceDomainListDefault) Error() string {
	return fmt.Sprintf("[GET /v1.0/servicedomains][%d] ServiceDomainList default  %+v", o._statusCode, o.Payload)
}

func (o *ServiceDomainListDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *ServiceDomainListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
