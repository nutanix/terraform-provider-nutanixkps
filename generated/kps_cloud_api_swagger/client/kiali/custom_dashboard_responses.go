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

// CustomDashboardReader is a Reader for the CustomDashboard structure.
type CustomDashboardReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomDashboardReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomDashboardOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomDashboardDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomDashboardOK creates a CustomDashboardOK with default headers values
func NewCustomDashboardOK() *CustomDashboardOK {
	return &CustomDashboardOK{}
}

/*CustomDashboardOK handles this case with default header values.

Dashboard response model
*/
type CustomDashboardOK struct {
	Payload *models.MonitoringDashboard
}

func (o *CustomDashboardOK) Error() string {
	return fmt.Sprintf("[GET /v1.0/kiali/namespaces/{namespace}/customdashboard/{dashboard}][%d] customDashboardOK  %+v", 200, o.Payload)
}

func (o *CustomDashboardOK) GetPayload() *models.MonitoringDashboard {
	return o.Payload
}

func (o *CustomDashboardOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MonitoringDashboard)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomDashboardDefault creates a CustomDashboardDefault with default headers values
func NewCustomDashboardDefault(code int) *CustomDashboardDefault {
	return &CustomDashboardDefault{
		_statusCode: code,
	}
}

/*CustomDashboardDefault handles this case with default header values.

generic API error response
*/
type CustomDashboardDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the custom dashboard default response
func (o *CustomDashboardDefault) Code() int {
	return o._statusCode
}

func (o *CustomDashboardDefault) Error() string {
	return fmt.Sprintf("[GET /v1.0/kiali/namespaces/{namespace}/customdashboard/{dashboard}][%d] customDashboard default  %+v", o._statusCode, o.Payload)
}

func (o *CustomDashboardDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *CustomDashboardDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
