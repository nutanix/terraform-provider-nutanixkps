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

// OAuthAuthorizeCallReader is a Reader for the OAuthAuthorizeCall structure.
type OAuthAuthorizeCallReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OAuthAuthorizeCallReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 401:
		result := NewOAuthAuthorizeCallUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewOAuthAuthorizeCallDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewOAuthAuthorizeCallUnauthorized creates a OAuthAuthorizeCallUnauthorized with default headers values
func NewOAuthAuthorizeCallUnauthorized() *OAuthAuthorizeCallUnauthorized {
	return &OAuthAuthorizeCallUnauthorized{}
}

/*OAuthAuthorizeCallUnauthorized handles this case with default header values.

Login failed
*/
type OAuthAuthorizeCallUnauthorized struct {
	Payload *models.APIErrorPayload
}

func (o *OAuthAuthorizeCallUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/oauth2/authorize][%d] oAuthAuthorizeCallUnauthorized  %+v", 401, o.Payload)
}

func (o *OAuthAuthorizeCallUnauthorized) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *OAuthAuthorizeCallUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOAuthAuthorizeCallDefault creates a OAuthAuthorizeCallDefault with default headers values
func NewOAuthAuthorizeCallDefault(code int) *OAuthAuthorizeCallDefault {
	return &OAuthAuthorizeCallDefault{
		_statusCode: code,
	}
}

/*OAuthAuthorizeCallDefault handles this case with default header values.

generic API error response
*/
type OAuthAuthorizeCallDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the o auth authorize call default response
func (o *OAuthAuthorizeCallDefault) Code() int {
	return o._statusCode
}

func (o *OAuthAuthorizeCallDefault) Error() string {
	return fmt.Sprintf("[GET /v1/oauth2/authorize][%d] OAuthAuthorizeCall default  %+v", o._statusCode, o.Payload)
}

func (o *OAuthAuthorizeCallDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *OAuthAuthorizeCallDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
