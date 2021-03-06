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

// UserGetReader is a Reader for the UserGet structure.
type UserGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUserGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUserGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUserGetOK creates a UserGetOK with default headers values
func NewUserGetOK() *UserGetOK {
	return &UserGetOK{}
}

/*UserGetOK handles this case with default header values.

Ok
*/
type UserGetOK struct {
	Payload *models.User
}

func (o *UserGetOK) Error() string {
	return fmt.Sprintf("[GET /v1/users/{id}][%d] userGetOK  %+v", 200, o.Payload)
}

func (o *UserGetOK) GetPayload() *models.User {
	return o.Payload
}

func (o *UserGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserGetDefault creates a UserGetDefault with default headers values
func NewUserGetDefault(code int) *UserGetDefault {
	return &UserGetDefault{
		_statusCode: code,
	}
}

/*UserGetDefault handles this case with default header values.

generic API error response
*/
type UserGetDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the user get default response
func (o *UserGetDefault) Code() int {
	return o._statusCode
}

func (o *UserGetDefault) Error() string {
	return fmt.Sprintf("[GET /v1/users/{id}][%d] UserGet default  %+v", o._statusCode, o.Payload)
}

func (o *UserGetDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *UserGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
