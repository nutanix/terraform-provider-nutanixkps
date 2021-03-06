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

// UserUpdateReader is a Reader for the UserUpdate structure.
type UserUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUserUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUserUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUserUpdateOK creates a UserUpdateOK with default headers values
func NewUserUpdateOK() *UserUpdateOK {
	return &UserUpdateOK{}
}

/*UserUpdateOK handles this case with default header values.

Ok
*/
type UserUpdateOK struct {
	Payload *models.UpdateDocumentResponse
}

func (o *UserUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /v1/users][%d] userUpdateOK  %+v", 200, o.Payload)
}

func (o *UserUpdateOK) GetPayload() *models.UpdateDocumentResponse {
	return o.Payload
}

func (o *UserUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UpdateDocumentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserUpdateDefault creates a UserUpdateDefault with default headers values
func NewUserUpdateDefault(code int) *UserUpdateDefault {
	return &UserUpdateDefault{
		_statusCode: code,
	}
}

/*UserUpdateDefault handles this case with default header values.

generic API error response
*/
type UserUpdateDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the user update default response
func (o *UserUpdateDefault) Code() int {
	return o._statusCode
}

func (o *UserUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /v1/users][%d] UserUpdate default  %+v", o._statusCode, o.Payload)
}

func (o *UserUpdateDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *UserUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
