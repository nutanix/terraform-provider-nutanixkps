// Code generated by go-swagger; DO NOT EDIT.

package user_props

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"
)

// UserPropsGetV2Reader is a Reader for the UserPropsGetV2 structure.
type UserPropsGetV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserPropsGetV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUserPropsGetV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUserPropsGetV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUserPropsGetV2OK creates a UserPropsGetV2OK with default headers values
func NewUserPropsGetV2OK() *UserPropsGetV2OK {
	return &UserPropsGetV2OK{}
}

/*UserPropsGetV2OK handles this case with default header values.

Ok
*/
type UserPropsGetV2OK struct {
	Payload *models.UserProps
}

func (o *UserPropsGetV2OK) Error() string {
	return fmt.Sprintf("[GET /v1.0/userprops/{id}][%d] userPropsGetV2OK  %+v", 200, o.Payload)
}

func (o *UserPropsGetV2OK) GetPayload() *models.UserProps {
	return o.Payload
}

func (o *UserPropsGetV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserProps)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserPropsGetV2Default creates a UserPropsGetV2Default with default headers values
func NewUserPropsGetV2Default(code int) *UserPropsGetV2Default {
	return &UserPropsGetV2Default{
		_statusCode: code,
	}
}

/*UserPropsGetV2Default handles this case with default header values.

generic API error response
*/
type UserPropsGetV2Default struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the user props get v2 default response
func (o *UserPropsGetV2Default) Code() int {
	return o._statusCode
}

func (o *UserPropsGetV2Default) Error() string {
	return fmt.Sprintf("[GET /v1.0/userprops/{id}][%d] UserPropsGetV2 default  %+v", o._statusCode, o.Payload)
}

func (o *UserPropsGetV2Default) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *UserPropsGetV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
