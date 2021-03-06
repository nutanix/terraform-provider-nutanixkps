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

// OAuthAuthorizeCallV2Reader is a Reader for the OAuthAuthorizeCallV2 structure.
type OAuthAuthorizeCallV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OAuthAuthorizeCallV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 401:
		result := NewOAuthAuthorizeCallV2Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewOAuthAuthorizeCallV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewOAuthAuthorizeCallV2Unauthorized creates a OAuthAuthorizeCallV2Unauthorized with default headers values
func NewOAuthAuthorizeCallV2Unauthorized() *OAuthAuthorizeCallV2Unauthorized {
	return &OAuthAuthorizeCallV2Unauthorized{}
}

/*OAuthAuthorizeCallV2Unauthorized handles this case with default header values.

Login failed
*/
type OAuthAuthorizeCallV2Unauthorized struct {
	Payload *models.APIErrorPayload
}

func (o *OAuthAuthorizeCallV2Unauthorized) Error() string {
	return fmt.Sprintf("[GET /v1.0/oauth2/authorize][%d] oAuthAuthorizeCallV2Unauthorized  %+v", 401, o.Payload)
}

func (o *OAuthAuthorizeCallV2Unauthorized) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *OAuthAuthorizeCallV2Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOAuthAuthorizeCallV2Default creates a OAuthAuthorizeCallV2Default with default headers values
func NewOAuthAuthorizeCallV2Default(code int) *OAuthAuthorizeCallV2Default {
	return &OAuthAuthorizeCallV2Default{
		_statusCode: code,
	}
}

/*OAuthAuthorizeCallV2Default handles this case with default header values.

generic API error response
*/
type OAuthAuthorizeCallV2Default struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the o auth authorize call v2 default response
func (o *OAuthAuthorizeCallV2Default) Code() int {
	return o._statusCode
}

func (o *OAuthAuthorizeCallV2Default) Error() string {
	return fmt.Sprintf("[GET /v1.0/oauth2/authorize][%d] OAuthAuthorizeCallV2 default  %+v", o._statusCode, o.Payload)
}

func (o *OAuthAuthorizeCallV2Default) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *OAuthAuthorizeCallV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
