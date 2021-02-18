// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"
)

// OAuthTokenCallV2Reader is a Reader for the OAuthTokenCallV2 structure.
type OAuthTokenCallV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OAuthTokenCallV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOAuthTokenCallV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewOAuthTokenCallV2Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewOAuthTokenCallV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewOAuthTokenCallV2OK creates a OAuthTokenCallV2OK with default headers values
func NewOAuthTokenCallV2OK() *OAuthTokenCallV2OK {
	return &OAuthTokenCallV2OK{}
}

/*OAuthTokenCallV2OK handles this case with default header values.

Ok
*/
type OAuthTokenCallV2OK struct {
	Payload *models.LoginResponse
}

func (o *OAuthTokenCallV2OK) Error() string {
	return fmt.Sprintf("[POST /v1.0/oauth2/token][%d] oAuthTokenCallV2OK  %+v", 200, o.Payload)
}

func (o *OAuthTokenCallV2OK) GetPayload() *models.LoginResponse {
	return o.Payload
}

func (o *OAuthTokenCallV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LoginResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOAuthTokenCallV2Unauthorized creates a OAuthTokenCallV2Unauthorized with default headers values
func NewOAuthTokenCallV2Unauthorized() *OAuthTokenCallV2Unauthorized {
	return &OAuthTokenCallV2Unauthorized{}
}

/*OAuthTokenCallV2Unauthorized handles this case with default header values.

Login failed
*/
type OAuthTokenCallV2Unauthorized struct {
	Payload *models.APIErrorPayload
}

func (o *OAuthTokenCallV2Unauthorized) Error() string {
	return fmt.Sprintf("[POST /v1.0/oauth2/token][%d] oAuthTokenCallV2Unauthorized  %+v", 401, o.Payload)
}

func (o *OAuthTokenCallV2Unauthorized) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *OAuthTokenCallV2Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOAuthTokenCallV2Default creates a OAuthTokenCallV2Default with default headers values
func NewOAuthTokenCallV2Default(code int) *OAuthTokenCallV2Default {
	return &OAuthTokenCallV2Default{
		_statusCode: code,
	}
}

/*OAuthTokenCallV2Default handles this case with default header values.

generic API error response
*/
type OAuthTokenCallV2Default struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the o auth token call v2 default response
func (o *OAuthTokenCallV2Default) Code() int {
	return o._statusCode
}

func (o *OAuthTokenCallV2Default) Error() string {
	return fmt.Sprintf("[POST /v1.0/oauth2/token][%d] OAuthTokenCallV2 default  %+v", o._statusCode, o.Payload)
}

func (o *OAuthTokenCallV2Default) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *OAuthTokenCallV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
