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

// ProjectServiceDeleteReader is a Reader for the ProjectServiceDelete structure.
type ProjectServiceDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectServiceDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectServiceDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewProjectServiceDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewProjectServiceDeleteOK creates a ProjectServiceDeleteOK with default headers values
func NewProjectServiceDeleteOK() *ProjectServiceDeleteOK {
	return &ProjectServiceDeleteOK{}
}

/*ProjectServiceDeleteOK handles this case with default header values.

Ok
*/
type ProjectServiceDeleteOK struct {
	Payload *models.DeleteDocumentResponseV2
}

func (o *ProjectServiceDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /v1.0/projectservices/{id}][%d] projectServiceDeleteOK  %+v", 200, o.Payload)
}

func (o *ProjectServiceDeleteOK) GetPayload() *models.DeleteDocumentResponseV2 {
	return o.Payload
}

func (o *ProjectServiceDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeleteDocumentResponseV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectServiceDeleteDefault creates a ProjectServiceDeleteDefault with default headers values
func NewProjectServiceDeleteDefault(code int) *ProjectServiceDeleteDefault {
	return &ProjectServiceDeleteDefault{
		_statusCode: code,
	}
}

/*ProjectServiceDeleteDefault handles this case with default header values.

generic API error response
*/
type ProjectServiceDeleteDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the project service delete default response
func (o *ProjectServiceDeleteDefault) Code() int {
	return o._statusCode
}

func (o *ProjectServiceDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /v1.0/projectservices/{id}][%d] ProjectServiceDelete default  %+v", o._statusCode, o.Payload)
}

func (o *ProjectServiceDeleteDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *ProjectServiceDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
