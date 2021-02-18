// Code generated by go-swagger; DO NOT EDIT.

package kiali

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"
)

// ServiceDetailsReader is a Reader for the ServiceDetails structure.
type ServiceDetailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServiceDetailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServiceDetailsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewServiceDetailsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewServiceDetailsOK creates a ServiceDetailsOK with default headers values
func NewServiceDetailsOK() *ServiceDetailsOK {
	return &ServiceDetailsOK{}
}

/*ServiceDetailsOK handles this case with default header values.

Listing all the information related to a workload
*/
type ServiceDetailsOK struct {
	Payload *models.ServiceDetails
}

func (o *ServiceDetailsOK) Error() string {
	return fmt.Sprintf("[GET /v1.0/kiali/namespaces/{namespace}/services/{service}][%d] serviceDetailsOK  %+v", 200, o.Payload)
}

func (o *ServiceDetailsOK) GetPayload() *models.ServiceDetails {
	return o.Payload
}

func (o *ServiceDetailsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceDetailsDefault creates a ServiceDetailsDefault with default headers values
func NewServiceDetailsDefault(code int) *ServiceDetailsDefault {
	return &ServiceDetailsDefault{
		_statusCode: code,
	}
}

/*ServiceDetailsDefault handles this case with default header values.

generic API error response
*/
type ServiceDetailsDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the service details default response
func (o *ServiceDetailsDefault) Code() int {
	return o._statusCode
}

func (o *ServiceDetailsDefault) Error() string {
	return fmt.Sprintf("[GET /v1.0/kiali/namespaces/{namespace}/services/{service}][%d] serviceDetails default  %+v", o._statusCode, o.Payload)
}

func (o *ServiceDetailsDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *ServiceDetailsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
