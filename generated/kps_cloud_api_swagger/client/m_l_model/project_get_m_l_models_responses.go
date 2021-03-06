// Code generated by go-swagger; DO NOT EDIT.

package m_l_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"
)

// ProjectGetMLModelsReader is a Reader for the ProjectGetMLModels structure.
type ProjectGetMLModelsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectGetMLModelsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectGetMLModelsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewProjectGetMLModelsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewProjectGetMLModelsOK creates a ProjectGetMLModelsOK with default headers values
func NewProjectGetMLModelsOK() *ProjectGetMLModelsOK {
	return &ProjectGetMLModelsOK{}
}

/*ProjectGetMLModelsOK handles this case with default header values.

Ok
*/
type ProjectGetMLModelsOK struct {
	Payload *models.MLModelListResponsePayload
}

func (o *ProjectGetMLModelsOK) Error() string {
	return fmt.Sprintf("[GET /v1.0/projects/{projectId}/mlmodels][%d] projectGetMLModelsOK  %+v", 200, o.Payload)
}

func (o *ProjectGetMLModelsOK) GetPayload() *models.MLModelListResponsePayload {
	return o.Payload
}

func (o *ProjectGetMLModelsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MLModelListResponsePayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectGetMLModelsDefault creates a ProjectGetMLModelsDefault with default headers values
func NewProjectGetMLModelsDefault(code int) *ProjectGetMLModelsDefault {
	return &ProjectGetMLModelsDefault{
		_statusCode: code,
	}
}

/*ProjectGetMLModelsDefault handles this case with default header values.

generic API error response
*/
type ProjectGetMLModelsDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the project get m l models default response
func (o *ProjectGetMLModelsDefault) Code() int {
	return o._statusCode
}

func (o *ProjectGetMLModelsDefault) Error() string {
	return fmt.Sprintf("[GET /v1.0/projects/{projectId}/mlmodels][%d] ProjectGetMLModels default  %+v", o._statusCode, o.Payload)
}

func (o *ProjectGetMLModelsDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *ProjectGetMLModelsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
