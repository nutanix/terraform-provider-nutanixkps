// Code generated by go-swagger; DO NOT EDIT.

package project_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"
)

// ProjectServiceListReader is a Reader for the ProjectServiceList structure.
type ProjectServiceListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectServiceListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectServiceListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewProjectServiceListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewProjectServiceListOK creates a ProjectServiceListOK with default headers values
func NewProjectServiceListOK() *ProjectServiceListOK {
	return &ProjectServiceListOK{}
}

/*ProjectServiceListOK handles this case with default header values.

Ok
*/
type ProjectServiceListOK struct {
	Payload models.ProjectServiceListPayload
}

func (o *ProjectServiceListOK) Error() string {
	return fmt.Sprintf("[GET /v1.0/projectservices][%d] projectServiceListOK  %+v", 200, o.Payload)
}

func (o *ProjectServiceListOK) GetPayload() models.ProjectServiceListPayload {
	return o.Payload
}

func (o *ProjectServiceListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectServiceListDefault creates a ProjectServiceListDefault with default headers values
func NewProjectServiceListDefault(code int) *ProjectServiceListDefault {
	return &ProjectServiceListDefault{
		_statusCode: code,
	}
}

/*ProjectServiceListDefault handles this case with default header values.

generic API error response
*/
type ProjectServiceListDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the project service list default response
func (o *ProjectServiceListDefault) Code() int {
	return o._statusCode
}

func (o *ProjectServiceListDefault) Error() string {
	return fmt.Sprintf("[GET /v1.0/projectservices][%d] ProjectServiceList default  %+v", o._statusCode, o.Payload)
}

func (o *ProjectServiceListDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *ProjectServiceListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
