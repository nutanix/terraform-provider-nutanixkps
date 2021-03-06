// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"
)

// ProjectGetApplicationsReader is a Reader for the ProjectGetApplications structure.
type ProjectGetApplicationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectGetApplicationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectGetApplicationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewProjectGetApplicationsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewProjectGetApplicationsOK creates a ProjectGetApplicationsOK with default headers values
func NewProjectGetApplicationsOK() *ProjectGetApplicationsOK {
	return &ProjectGetApplicationsOK{}
}

/*ProjectGetApplicationsOK handles this case with default header values.

Ok
*/
type ProjectGetApplicationsOK struct {
	Payload []*models.Application
}

func (o *ProjectGetApplicationsOK) Error() string {
	return fmt.Sprintf("[GET /v1/projects/{projectId}/applications][%d] projectGetApplicationsOK  %+v", 200, o.Payload)
}

func (o *ProjectGetApplicationsOK) GetPayload() []*models.Application {
	return o.Payload
}

func (o *ProjectGetApplicationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectGetApplicationsDefault creates a ProjectGetApplicationsDefault with default headers values
func NewProjectGetApplicationsDefault(code int) *ProjectGetApplicationsDefault {
	return &ProjectGetApplicationsDefault{
		_statusCode: code,
	}
}

/*ProjectGetApplicationsDefault handles this case with default header values.

generic API error response
*/
type ProjectGetApplicationsDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the project get applications default response
func (o *ProjectGetApplicationsDefault) Code() int {
	return o._statusCode
}

func (o *ProjectGetApplicationsDefault) Error() string {
	return fmt.Sprintf("[GET /v1/projects/{projectId}/applications][%d] ProjectGetApplications default  %+v", o._statusCode, o.Payload)
}

func (o *ProjectGetApplicationsDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *ProjectGetApplicationsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
