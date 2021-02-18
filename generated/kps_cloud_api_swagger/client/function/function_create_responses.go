// Code generated by go-swagger; DO NOT EDIT.

package function

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"
)

// FunctionCreateReader is a Reader for the FunctionCreate structure.
type FunctionCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FunctionCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFunctionCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFunctionCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFunctionCreateOK creates a FunctionCreateOK with default headers values
func NewFunctionCreateOK() *FunctionCreateOK {
	return &FunctionCreateOK{}
}

/*FunctionCreateOK handles this case with default header values.

Ok
*/
type FunctionCreateOK struct {
	Payload *models.CreateDocumentResponseV2
}

func (o *FunctionCreateOK) Error() string {
	return fmt.Sprintf("[POST /v1.0/functions][%d] functionCreateOK  %+v", 200, o.Payload)
}

func (o *FunctionCreateOK) GetPayload() *models.CreateDocumentResponseV2 {
	return o.Payload
}

func (o *FunctionCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateDocumentResponseV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFunctionCreateDefault creates a FunctionCreateDefault with default headers values
func NewFunctionCreateDefault(code int) *FunctionCreateDefault {
	return &FunctionCreateDefault{
		_statusCode: code,
	}
}

/*FunctionCreateDefault handles this case with default header values.

generic API error response
*/
type FunctionCreateDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the function create default response
func (o *FunctionCreateDefault) Code() int {
	return o._statusCode
}

func (o *FunctionCreateDefault) Error() string {
	return fmt.Sprintf("[POST /v1.0/functions][%d] FunctionCreate default  %+v", o._statusCode, o.Payload)
}

func (o *FunctionCreateDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *FunctionCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
