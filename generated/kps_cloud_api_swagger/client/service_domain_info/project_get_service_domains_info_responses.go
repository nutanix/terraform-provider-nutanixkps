// Code generated by go-swagger; DO NOT EDIT.

package service_domain_info

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"
)

// ProjectGetServiceDomainsInfoReader is a Reader for the ProjectGetServiceDomainsInfo structure.
type ProjectGetServiceDomainsInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectGetServiceDomainsInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectGetServiceDomainsInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewProjectGetServiceDomainsInfoDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewProjectGetServiceDomainsInfoOK creates a ProjectGetServiceDomainsInfoOK with default headers values
func NewProjectGetServiceDomainsInfoOK() *ProjectGetServiceDomainsInfoOK {
	return &ProjectGetServiceDomainsInfoOK{}
}

/*ProjectGetServiceDomainsInfoOK handles this case with default header values.

Ok
*/
type ProjectGetServiceDomainsInfoOK struct {
	Payload *models.ServiceDomainInfoListPayload
}

func (o *ProjectGetServiceDomainsInfoOK) Error() string {
	return fmt.Sprintf("[GET /v1.0/projects/{projectId}/servicedomainsinfo][%d] projectGetServiceDomainsInfoOK  %+v", 200, o.Payload)
}

func (o *ProjectGetServiceDomainsInfoOK) GetPayload() *models.ServiceDomainInfoListPayload {
	return o.Payload
}

func (o *ProjectGetServiceDomainsInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceDomainInfoListPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectGetServiceDomainsInfoDefault creates a ProjectGetServiceDomainsInfoDefault with default headers values
func NewProjectGetServiceDomainsInfoDefault(code int) *ProjectGetServiceDomainsInfoDefault {
	return &ProjectGetServiceDomainsInfoDefault{
		_statusCode: code,
	}
}

/*ProjectGetServiceDomainsInfoDefault handles this case with default header values.

generic API error response
*/
type ProjectGetServiceDomainsInfoDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the project get service domains info default response
func (o *ProjectGetServiceDomainsInfoDefault) Code() int {
	return o._statusCode
}

func (o *ProjectGetServiceDomainsInfoDefault) Error() string {
	return fmt.Sprintf("[GET /v1.0/projects/{projectId}/servicedomainsinfo][%d] ProjectGetServiceDomainsInfo default  %+v", o._statusCode, o.Payload)
}

func (o *ProjectGetServiceDomainsInfoDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *ProjectGetServiceDomainsInfoDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
