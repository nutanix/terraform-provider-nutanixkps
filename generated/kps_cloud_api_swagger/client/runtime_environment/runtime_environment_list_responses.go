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

// RuntimeEnvironmentListReader is a Reader for the RuntimeEnvironmentList structure.
type RuntimeEnvironmentListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RuntimeEnvironmentListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRuntimeEnvironmentListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewRuntimeEnvironmentListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRuntimeEnvironmentListOK creates a RuntimeEnvironmentListOK with default headers values
func NewRuntimeEnvironmentListOK() *RuntimeEnvironmentListOK {
	return &RuntimeEnvironmentListOK{}
}

/*RuntimeEnvironmentListOK handles this case with default header values.

Ok
*/
type RuntimeEnvironmentListOK struct {
	Payload *models.RuntimeEnvironmentListPayload
}

func (o *RuntimeEnvironmentListOK) Error() string {
	return fmt.Sprintf("[GET /v1.0/runtimeenvironments][%d] runtimeEnvironmentListOK  %+v", 200, o.Payload)
}

func (o *RuntimeEnvironmentListOK) GetPayload() *models.RuntimeEnvironmentListPayload {
	return o.Payload
}

func (o *RuntimeEnvironmentListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RuntimeEnvironmentListPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRuntimeEnvironmentListDefault creates a RuntimeEnvironmentListDefault with default headers values
func NewRuntimeEnvironmentListDefault(code int) *RuntimeEnvironmentListDefault {
	return &RuntimeEnvironmentListDefault{
		_statusCode: code,
	}
}

/*RuntimeEnvironmentListDefault handles this case with default header values.

generic API error response
*/
type RuntimeEnvironmentListDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the runtime environment list default response
func (o *RuntimeEnvironmentListDefault) Code() int {
	return o._statusCode
}

func (o *RuntimeEnvironmentListDefault) Error() string {
	return fmt.Sprintf("[GET /v1.0/runtimeenvironments][%d] RuntimeEnvironmentList default  %+v", o._statusCode, o.Payload)
}

func (o *RuntimeEnvironmentListDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *RuntimeEnvironmentListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
