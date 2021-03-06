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

// LoginCallReader is a Reader for the LoginCall structure.
type LoginCallReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LoginCallReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLoginCallOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewLoginCallUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewLoginCallDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLoginCallOK creates a LoginCallOK with default headers values
func NewLoginCallOK() *LoginCallOK {
	return &LoginCallOK{}
}

/*LoginCallOK handles this case with default header values.

Ok
*/
type LoginCallOK struct {
	Payload *models.LoginResponse
}

func (o *LoginCallOK) Error() string {
	return fmt.Sprintf("[POST /v1/login][%d] loginCallOK  %+v", 200, o.Payload)
}

func (o *LoginCallOK) GetPayload() *models.LoginResponse {
	return o.Payload
}

func (o *LoginCallOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LoginResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLoginCallUnauthorized creates a LoginCallUnauthorized with default headers values
func NewLoginCallUnauthorized() *LoginCallUnauthorized {
	return &LoginCallUnauthorized{}
}

/*LoginCallUnauthorized handles this case with default header values.

Login failed
*/
type LoginCallUnauthorized struct {
	Payload *models.APIErrorPayload
}

func (o *LoginCallUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/login][%d] loginCallUnauthorized  %+v", 401, o.Payload)
}

func (o *LoginCallUnauthorized) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *LoginCallUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLoginCallDefault creates a LoginCallDefault with default headers values
func NewLoginCallDefault(code int) *LoginCallDefault {
	return &LoginCallDefault{
		_statusCode: code,
	}
}

/*LoginCallDefault handles this case with default header values.

generic API error response
*/
type LoginCallDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the login call default response
func (o *LoginCallDefault) Code() int {
	return o._statusCode
}

func (o *LoginCallDefault) Error() string {
	return fmt.Sprintf("[POST /v1/login][%d] LoginCall default  %+v", o._statusCode, o.Payload)
}

func (o *LoginCallDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *LoginCallDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
