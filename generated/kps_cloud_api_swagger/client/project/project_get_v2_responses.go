// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"
)

// ProjectGetV2Reader is a Reader for the ProjectGetV2 structure.
type ProjectGetV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectGetV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectGetV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewProjectGetV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewProjectGetV2OK creates a ProjectGetV2OK with default headers values
func NewProjectGetV2OK() *ProjectGetV2OK {
	return &ProjectGetV2OK{}
}

/*ProjectGetV2OK handles this case with default header values.

Ok
*/
type ProjectGetV2OK struct {
	Payload *models.Project
}

func (o *ProjectGetV2OK) Error() string {
	return fmt.Sprintf("[GET /v1.0/projects/{projectId}][%d] projectGetV2OK  %+v", 200, o.Payload)
}

func (o *ProjectGetV2OK) GetPayload() *models.Project {
	return o.Payload
}

func (o *ProjectGetV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Project)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectGetV2Default creates a ProjectGetV2Default with default headers values
func NewProjectGetV2Default(code int) *ProjectGetV2Default {
	return &ProjectGetV2Default{
		_statusCode: code,
	}
}

/*ProjectGetV2Default handles this case with default header values.

generic API error response
*/
type ProjectGetV2Default struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the project get v2 default response
func (o *ProjectGetV2Default) Code() int {
	return o._statusCode
}

func (o *ProjectGetV2Default) Error() string {
	return fmt.Sprintf("[GET /v1.0/projects/{projectId}][%d] ProjectGetV2 default  %+v", o._statusCode, o.Payload)
}

func (o *ProjectGetV2Default) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *ProjectGetV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
