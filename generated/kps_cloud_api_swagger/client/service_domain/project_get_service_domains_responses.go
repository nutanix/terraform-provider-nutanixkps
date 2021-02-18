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

// ProjectGetServiceDomainsReader is a Reader for the ProjectGetServiceDomains structure.
type ProjectGetServiceDomainsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectGetServiceDomainsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectGetServiceDomainsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewProjectGetServiceDomainsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewProjectGetServiceDomainsOK creates a ProjectGetServiceDomainsOK with default headers values
func NewProjectGetServiceDomainsOK() *ProjectGetServiceDomainsOK {
	return &ProjectGetServiceDomainsOK{}
}

/*ProjectGetServiceDomainsOK handles this case with default header values.

Ok
*/
type ProjectGetServiceDomainsOK struct {
	Payload *models.ServiceDomainListPayload
}

func (o *ProjectGetServiceDomainsOK) Error() string {
	return fmt.Sprintf("[GET /v1.0/projects/{projectId}/servicedomains][%d] projectGetServiceDomainsOK  %+v", 200, o.Payload)
}

func (o *ProjectGetServiceDomainsOK) GetPayload() *models.ServiceDomainListPayload {
	return o.Payload
}

func (o *ProjectGetServiceDomainsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceDomainListPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectGetServiceDomainsDefault creates a ProjectGetServiceDomainsDefault with default headers values
func NewProjectGetServiceDomainsDefault(code int) *ProjectGetServiceDomainsDefault {
	return &ProjectGetServiceDomainsDefault{
		_statusCode: code,
	}
}

/*ProjectGetServiceDomainsDefault handles this case with default header values.

generic API error response
*/
type ProjectGetServiceDomainsDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the project get service domains default response
func (o *ProjectGetServiceDomainsDefault) Code() int {
	return o._statusCode
}

func (o *ProjectGetServiceDomainsDefault) Error() string {
	return fmt.Sprintf("[GET /v1.0/projects/{projectId}/servicedomains][%d] ProjectGetServiceDomains default  %+v", o._statusCode, o.Payload)
}

func (o *ProjectGetServiceDomainsDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *ProjectGetServiceDomainsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
