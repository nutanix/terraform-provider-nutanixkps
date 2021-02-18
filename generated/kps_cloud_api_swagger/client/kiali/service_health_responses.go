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

// ServiceHealthReader is a Reader for the ServiceHealth structure.
type ServiceHealthReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServiceHealthReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServiceHealthOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewServiceHealthDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewServiceHealthOK creates a ServiceHealthOK with default headers values
func NewServiceHealthOK() *ServiceHealthOK {
	return &ServiceHealthOK{}
}

/*ServiceHealthOK handles this case with default header values.

serviceHealthResponse contains aggregated health from various sources, for a given service
*/
type ServiceHealthOK struct {
	Payload *models.ServiceHealth
}

func (o *ServiceHealthOK) Error() string {
	return fmt.Sprintf("[GET /v1.0/kiali/namespaces/{namespace}/services/{service}/health][%d] serviceHealthOK  %+v", 200, o.Payload)
}

func (o *ServiceHealthOK) GetPayload() *models.ServiceHealth {
	return o.Payload
}

func (o *ServiceHealthOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceHealth)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceHealthDefault creates a ServiceHealthDefault with default headers values
func NewServiceHealthDefault(code int) *ServiceHealthDefault {
	return &ServiceHealthDefault{
		_statusCode: code,
	}
}

/*ServiceHealthDefault handles this case with default header values.

generic API error response
*/
type ServiceHealthDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the service health default response
func (o *ServiceHealthDefault) Code() int {
	return o._statusCode
}

func (o *ServiceHealthDefault) Error() string {
	return fmt.Sprintf("[GET /v1.0/kiali/namespaces/{namespace}/services/{service}/health][%d] serviceHealth default  %+v", o._statusCode, o.Payload)
}

func (o *ServiceHealthDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *ServiceHealthDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
