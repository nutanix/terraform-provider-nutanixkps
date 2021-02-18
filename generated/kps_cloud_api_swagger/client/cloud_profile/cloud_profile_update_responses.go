// Code generated by go-swagger; DO NOT EDIT.

package cloud_profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"
)

// CloudProfileUpdateReader is a Reader for the CloudProfileUpdate structure.
type CloudProfileUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CloudProfileUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCloudProfileUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCloudProfileUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCloudProfileUpdateOK creates a CloudProfileUpdateOK with default headers values
func NewCloudProfileUpdateOK() *CloudProfileUpdateOK {
	return &CloudProfileUpdateOK{}
}

/*CloudProfileUpdateOK handles this case with default header values.

Ok
*/
type CloudProfileUpdateOK struct {
	Payload *models.UpdateDocumentResponseV2
}

func (o *CloudProfileUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /v1.0/cloudprofiles/{id}][%d] cloudProfileUpdateOK  %+v", 200, o.Payload)
}

func (o *CloudProfileUpdateOK) GetPayload() *models.UpdateDocumentResponseV2 {
	return o.Payload
}

func (o *CloudProfileUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UpdateDocumentResponseV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloudProfileUpdateDefault creates a CloudProfileUpdateDefault with default headers values
func NewCloudProfileUpdateDefault(code int) *CloudProfileUpdateDefault {
	return &CloudProfileUpdateDefault{
		_statusCode: code,
	}
}

/*CloudProfileUpdateDefault handles this case with default header values.

generic API error response
*/
type CloudProfileUpdateDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the cloud profile update default response
func (o *CloudProfileUpdateDefault) Code() int {
	return o._statusCode
}

func (o *CloudProfileUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /v1.0/cloudprofiles/{id}][%d] CloudProfileUpdate default  %+v", o._statusCode, o.Payload)
}

func (o *CloudProfileUpdateDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *CloudProfileUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
