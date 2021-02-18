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

// ServiceDomainCreateReader is a Reader for the ServiceDomainCreate structure.
type ServiceDomainCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServiceDomainCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServiceDomainCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewServiceDomainCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewServiceDomainCreateOK creates a ServiceDomainCreateOK with default headers values
func NewServiceDomainCreateOK() *ServiceDomainCreateOK {
	return &ServiceDomainCreateOK{}
}

/*ServiceDomainCreateOK handles this case with default header values.

Ok
*/
type ServiceDomainCreateOK struct {
	Payload *models.CreateDocumentResponseV2
}

func (o *ServiceDomainCreateOK) Error() string {
	return fmt.Sprintf("[POST /v1.0/servicedomains][%d] serviceDomainCreateOK  %+v", 200, o.Payload)
}

func (o *ServiceDomainCreateOK) GetPayload() *models.CreateDocumentResponseV2 {
	return o.Payload
}

func (o *ServiceDomainCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateDocumentResponseV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceDomainCreateDefault creates a ServiceDomainCreateDefault with default headers values
func NewServiceDomainCreateDefault(code int) *ServiceDomainCreateDefault {
	return &ServiceDomainCreateDefault{
		_statusCode: code,
	}
}

/*ServiceDomainCreateDefault handles this case with default header values.

generic API error response
*/
type ServiceDomainCreateDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the service domain create default response
func (o *ServiceDomainCreateDefault) Code() int {
	return o._statusCode
}

func (o *ServiceDomainCreateDefault) Error() string {
	return fmt.Sprintf("[POST /v1.0/servicedomains][%d] ServiceDomainCreate default  %+v", o._statusCode, o.Payload)
}

func (o *ServiceDomainCreateDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *ServiceDomainCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
