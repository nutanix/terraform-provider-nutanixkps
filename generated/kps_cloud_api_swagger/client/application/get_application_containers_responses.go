// Code generated by go-swagger; DO NOT EDIT.

package application

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"
)

// GetApplicationContainersReader is a Reader for the GetApplicationContainers structure.
type GetApplicationContainersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetApplicationContainersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApplicationContainersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetApplicationContainersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetApplicationContainersOK creates a GetApplicationContainersOK with default headers values
func NewGetApplicationContainersOK() *GetApplicationContainersOK {
	return &GetApplicationContainersOK{}
}

/*GetApplicationContainersOK handles this case with default header values.

GetApplicationContainersResponse is the API response that
returns a list of container names for a given app on a given edge.
*/
type GetApplicationContainersOK struct {
	Payload *models.ApplicationContainers
}

func (o *GetApplicationContainersOK) Error() string {
	return fmt.Sprintf("[GET /v1.0/applications/{id}/containers/{edgeId}][%d] getApplicationContainersOK  %+v", 200, o.Payload)
}

func (o *GetApplicationContainersOK) GetPayload() *models.ApplicationContainers {
	return o.Payload
}

func (o *GetApplicationContainersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApplicationContainers)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApplicationContainersDefault creates a GetApplicationContainersDefault with default headers values
func NewGetApplicationContainersDefault(code int) *GetApplicationContainersDefault {
	return &GetApplicationContainersDefault{
		_statusCode: code,
	}
}

/*GetApplicationContainersDefault handles this case with default header values.

generic API error response
*/
type GetApplicationContainersDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the get application containers default response
func (o *GetApplicationContainersDefault) Code() int {
	return o._statusCode
}

func (o *GetApplicationContainersDefault) Error() string {
	return fmt.Sprintf("[GET /v1.0/applications/{id}/containers/{edgeId}][%d] GetApplicationContainers default  %+v", o._statusCode, o.Payload)
}

func (o *GetApplicationContainersDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *GetApplicationContainersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
