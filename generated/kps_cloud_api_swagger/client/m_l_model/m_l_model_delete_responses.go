// Code generated by go-swagger; DO NOT EDIT.

package m_l_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"
)

// MLModelDeleteReader is a Reader for the MLModelDelete structure.
type MLModelDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MLModelDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMLModelDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewMLModelDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewMLModelDeleteOK creates a MLModelDeleteOK with default headers values
func NewMLModelDeleteOK() *MLModelDeleteOK {
	return &MLModelDeleteOK{}
}

/*MLModelDeleteOK handles this case with default header values.

Ok
*/
type MLModelDeleteOK struct {
	Payload *models.DeleteDocumentResponseV2
}

func (o *MLModelDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /v1.0/mlmodels/{id}][%d] mLModelDeleteOK  %+v", 200, o.Payload)
}

func (o *MLModelDeleteOK) GetPayload() *models.DeleteDocumentResponseV2 {
	return o.Payload
}

func (o *MLModelDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeleteDocumentResponseV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMLModelDeleteDefault creates a MLModelDeleteDefault with default headers values
func NewMLModelDeleteDefault(code int) *MLModelDeleteDefault {
	return &MLModelDeleteDefault{
		_statusCode: code,
	}
}

/*MLModelDeleteDefault handles this case with default header values.

generic API error response
*/
type MLModelDeleteDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the m l model delete default response
func (o *MLModelDeleteDefault) Code() int {
	return o._statusCode
}

func (o *MLModelDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /v1.0/mlmodels/{id}][%d] MLModelDelete default  %+v", o._statusCode, o.Payload)
}

func (o *MLModelDeleteDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *MLModelDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
