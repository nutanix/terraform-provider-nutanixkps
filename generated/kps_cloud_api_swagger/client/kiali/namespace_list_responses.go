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

// NamespaceListReader is a Reader for the NamespaceList structure.
type NamespaceListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NamespaceListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNamespaceListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNamespaceListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNamespaceListOK creates a NamespaceListOK with default headers values
func NewNamespaceListOK() *NamespaceListOK {
	return &NamespaceListOK{}
}

/*NamespaceListOK handles this case with default header values.

List of Namespaces
*/
type NamespaceListOK struct {
	Payload []*models.Namespace
}

func (o *NamespaceListOK) Error() string {
	return fmt.Sprintf("[GET /v1.0/kiali/namespaces][%d] namespaceListOK  %+v", 200, o.Payload)
}

func (o *NamespaceListOK) GetPayload() []*models.Namespace {
	return o.Payload
}

func (o *NamespaceListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNamespaceListDefault creates a NamespaceListDefault with default headers values
func NewNamespaceListDefault(code int) *NamespaceListDefault {
	return &NamespaceListDefault{
		_statusCode: code,
	}
}

/*NamespaceListDefault handles this case with default header values.

generic API error response
*/
type NamespaceListDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the namespace list default response
func (o *NamespaceListDefault) Code() int {
	return o._statusCode
}

func (o *NamespaceListDefault) Error() string {
	return fmt.Sprintf("[GET /v1.0/kiali/namespaces][%d] namespaceList default  %+v", o._statusCode, o.Payload)
}

func (o *NamespaceListDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *NamespaceListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
