// Code generated by go-swagger; DO NOT EDIT.

package runtime_environment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"
)

// ProjectGetRuntimeEnvironmentsReader is a Reader for the ProjectGetRuntimeEnvironments structure.
type ProjectGetRuntimeEnvironmentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectGetRuntimeEnvironmentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectGetRuntimeEnvironmentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewProjectGetRuntimeEnvironmentsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewProjectGetRuntimeEnvironmentsOK creates a ProjectGetRuntimeEnvironmentsOK with default headers values
func NewProjectGetRuntimeEnvironmentsOK() *ProjectGetRuntimeEnvironmentsOK {
	return &ProjectGetRuntimeEnvironmentsOK{}
}

/*ProjectGetRuntimeEnvironmentsOK handles this case with default header values.

Ok
*/
type ProjectGetRuntimeEnvironmentsOK struct {
	Payload *models.RuntimeEnvironmentListPayload
}

func (o *ProjectGetRuntimeEnvironmentsOK) Error() string {
	return fmt.Sprintf("[GET /v1.0/projects/{projectId}/runtimeenvironments][%d] projectGetRuntimeEnvironmentsOK  %+v", 200, o.Payload)
}

func (o *ProjectGetRuntimeEnvironmentsOK) GetPayload() *models.RuntimeEnvironmentListPayload {
	return o.Payload
}

func (o *ProjectGetRuntimeEnvironmentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeEnvironmentListPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectGetRuntimeEnvironmentsDefault creates a ProjectGetRuntimeEnvironmentsDefault with default headers values
func NewProjectGetRuntimeEnvironmentsDefault(code int) *ProjectGetRuntimeEnvironmentsDefault {
	return &ProjectGetRuntimeEnvironmentsDefault{
		_statusCode: code,
	}
}

/*ProjectGetRuntimeEnvironmentsDefault handles this case with default header values.

generic API error response
*/
type ProjectGetRuntimeEnvironmentsDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the project get runtime environments default response
func (o *ProjectGetRuntimeEnvironmentsDefault) Code() int {
	return o._statusCode
}

func (o *ProjectGetRuntimeEnvironmentsDefault) Error() string {
	return fmt.Sprintf("[GET /v1.0/projects/{projectId}/runtimeenvironments][%d] ProjectGetRuntimeEnvironments default  %+v", o._statusCode, o.Payload)
}

func (o *ProjectGetRuntimeEnvironmentsDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *ProjectGetRuntimeEnvironmentsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
