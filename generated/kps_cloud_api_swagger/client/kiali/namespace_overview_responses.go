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

// NamespaceOverviewReader is a Reader for the NamespaceOverview structure.
type NamespaceOverviewReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NamespaceOverviewReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNamespaceOverviewOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNamespaceOverviewDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNamespaceOverviewOK creates a NamespaceOverviewOK with default headers values
func NewNamespaceOverviewOK() *NamespaceOverviewOK {
	return &NamespaceOverviewOK{}
}

/*NamespaceOverviewOK handles this case with default header values.

Return Namespace combined Info
*/
type NamespaceOverviewOK struct {
	Payload *models.NamespaceOverview
}

func (o *NamespaceOverviewOK) Error() string {
	return fmt.Sprintf("[GET /v1.0/kiali/namespaces/{namespace}/overview][%d] namespaceOverviewOK  %+v", 200, o.Payload)
}

func (o *NamespaceOverviewOK) GetPayload() *models.NamespaceOverview {
	return o.Payload
}

func (o *NamespaceOverviewOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NamespaceOverview)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNamespaceOverviewDefault creates a NamespaceOverviewDefault with default headers values
func NewNamespaceOverviewDefault(code int) *NamespaceOverviewDefault {
	return &NamespaceOverviewDefault{
		_statusCode: code,
	}
}

/*NamespaceOverviewDefault handles this case with default header values.

generic API error response
*/
type NamespaceOverviewDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the namespace overview default response
func (o *NamespaceOverviewDefault) Code() int {
	return o._statusCode
}

func (o *NamespaceOverviewDefault) Error() string {
	return fmt.Sprintf("[GET /v1.0/kiali/namespaces/{namespace}/overview][%d] namespaceOverview default  %+v", o._statusCode, o.Payload)
}

func (o *NamespaceOverviewDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *NamespaceOverviewDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
