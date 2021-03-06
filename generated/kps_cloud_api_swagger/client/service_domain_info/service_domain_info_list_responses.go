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

// ServiceDomainInfoListReader is a Reader for the ServiceDomainInfoList structure.
type ServiceDomainInfoListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServiceDomainInfoListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServiceDomainInfoListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewServiceDomainInfoListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewServiceDomainInfoListOK creates a ServiceDomainInfoListOK with default headers values
func NewServiceDomainInfoListOK() *ServiceDomainInfoListOK {
	return &ServiceDomainInfoListOK{}
}

/*ServiceDomainInfoListOK handles this case with default header values.

Ok
*/
type ServiceDomainInfoListOK struct {
	Payload *models.ServiceDomainInfoListPayload
}

func (o *ServiceDomainInfoListOK) Error() string {
	return fmt.Sprintf("[GET /v1.0/servicedomainsinfo][%d] serviceDomainInfoListOK  %+v", 200, o.Payload)
}

func (o *ServiceDomainInfoListOK) GetPayload() *models.ServiceDomainInfoListPayload {
	return o.Payload
}

func (o *ServiceDomainInfoListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceDomainInfoListPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceDomainInfoListDefault creates a ServiceDomainInfoListDefault with default headers values
func NewServiceDomainInfoListDefault(code int) *ServiceDomainInfoListDefault {
	return &ServiceDomainInfoListDefault{
		_statusCode: code,
	}
}

/*ServiceDomainInfoListDefault handles this case with default header values.

generic API error response
*/
type ServiceDomainInfoListDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the service domain info list default response
func (o *ServiceDomainInfoListDefault) Code() int {
	return o._statusCode
}

func (o *ServiceDomainInfoListDefault) Error() string {
	return fmt.Sprintf("[GET /v1.0/servicedomainsinfo][%d] ServiceDomainInfoList default  %+v", o._statusCode, o.Payload)
}

func (o *ServiceDomainInfoListDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *ServiceDomainInfoListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
