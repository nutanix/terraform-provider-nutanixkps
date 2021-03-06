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

// IstioConfigCreateReader is a Reader for the IstioConfigCreate structure.
type IstioConfigCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IstioConfigCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIstioConfigCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewIstioConfigCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIstioConfigCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIstioConfigCreateOK creates a IstioConfigCreateOK with default headers values
func NewIstioConfigCreateOK() *IstioConfigCreateOK {
	return &IstioConfigCreateOK{}
}

/*IstioConfigCreateOK handles this case with default header values.

IstioConfig details of an specific Istio Object
*/
type IstioConfigCreateOK struct {
	Payload *models.IstioConfigDetails
}

func (o *IstioConfigCreateOK) Error() string {
	return fmt.Sprintf("[POST /v1.0/kiali/namespaces/{namespace}/istio/{object_type}][%d] istioConfigCreateOK  %+v", 200, o.Payload)
}

func (o *IstioConfigCreateOK) GetPayload() *models.IstioConfigDetails {
	return o.Payload
}

func (o *IstioConfigCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IstioConfigDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIstioConfigCreateCreated creates a IstioConfigCreateCreated with default headers values
func NewIstioConfigCreateCreated() *IstioConfigCreateCreated {
	return &IstioConfigCreateCreated{}
}

/*IstioConfigCreateCreated handles this case with default header values.

IstioConfig details of an specific Istio Object
*/
type IstioConfigCreateCreated struct {
	Payload *models.IstioConfigDetails
}

func (o *IstioConfigCreateCreated) Error() string {
	return fmt.Sprintf("[POST /v1.0/kiali/namespaces/{namespace}/istio/{object_type}][%d] istioConfigCreateCreated  %+v", 201, o.Payload)
}

func (o *IstioConfigCreateCreated) GetPayload() *models.IstioConfigDetails {
	return o.Payload
}

func (o *IstioConfigCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IstioConfigDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIstioConfigCreateDefault creates a IstioConfigCreateDefault with default headers values
func NewIstioConfigCreateDefault(code int) *IstioConfigCreateDefault {
	return &IstioConfigCreateDefault{
		_statusCode: code,
	}
}

/*IstioConfigCreateDefault handles this case with default header values.

generic API error response
*/
type IstioConfigCreateDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the istio config create default response
func (o *IstioConfigCreateDefault) Code() int {
	return o._statusCode
}

func (o *IstioConfigCreateDefault) Error() string {
	return fmt.Sprintf("[POST /v1.0/kiali/namespaces/{namespace}/istio/{object_type}][%d] istioConfigCreate default  %+v", o._statusCode, o.Payload)
}

func (o *IstioConfigCreateDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *IstioConfigCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
