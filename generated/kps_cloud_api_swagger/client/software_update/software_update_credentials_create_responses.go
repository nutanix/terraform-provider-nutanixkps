// Code generated by go-swagger; DO NOT EDIT.

package software_update

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"
)

// SoftwareUpdateCredentialsCreateReader is a Reader for the SoftwareUpdateCredentialsCreate structure.
type SoftwareUpdateCredentialsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SoftwareUpdateCredentialsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSoftwareUpdateCredentialsCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSoftwareUpdateCredentialsCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSoftwareUpdateCredentialsCreateOK creates a SoftwareUpdateCredentialsCreateOK with default headers values
func NewSoftwareUpdateCredentialsCreateOK() *SoftwareUpdateCredentialsCreateOK {
	return &SoftwareUpdateCredentialsCreateOK{}
}

/*SoftwareUpdateCredentialsCreateOK handles this case with default header values.

Ok
*/
type SoftwareUpdateCredentialsCreateOK struct {
	Payload *models.SoftwareUpdateCredentialsCreatePayload
}

func (o *SoftwareUpdateCredentialsCreateOK) Error() string {
	return fmt.Sprintf("[POST /v1.0/softwareupdates/credentials][%d] softwareUpdateCredentialsCreateOK  %+v", 200, o.Payload)
}

func (o *SoftwareUpdateCredentialsCreateOK) GetPayload() *models.SoftwareUpdateCredentialsCreatePayload {
	return o.Payload
}

func (o *SoftwareUpdateCredentialsCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SoftwareUpdateCredentialsCreatePayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSoftwareUpdateCredentialsCreateDefault creates a SoftwareUpdateCredentialsCreateDefault with default headers values
func NewSoftwareUpdateCredentialsCreateDefault(code int) *SoftwareUpdateCredentialsCreateDefault {
	return &SoftwareUpdateCredentialsCreateDefault{
		_statusCode: code,
	}
}

/*SoftwareUpdateCredentialsCreateDefault handles this case with default header values.

generic API error response
*/
type SoftwareUpdateCredentialsCreateDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the software update credentials create default response
func (o *SoftwareUpdateCredentialsCreateDefault) Code() int {
	return o._statusCode
}

func (o *SoftwareUpdateCredentialsCreateDefault) Error() string {
	return fmt.Sprintf("[POST /v1.0/softwareupdates/credentials][%d] SoftwareUpdateCredentialsCreate default  %+v", o._statusCode, o.Payload)
}

func (o *SoftwareUpdateCredentialsCreateDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *SoftwareUpdateCredentialsCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
