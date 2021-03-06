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

// EdgeCreateReader is a Reader for the EdgeCreate structure.
type EdgeCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewEdgeCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEdgeCreateOK creates a EdgeCreateOK with default headers values
func NewEdgeCreateOK() *EdgeCreateOK {
	return &EdgeCreateOK{}
}

/*EdgeCreateOK handles this case with default header values.

Ok
*/
type EdgeCreateOK struct {
	Payload *models.CreateDocumentResponse
}

func (o *EdgeCreateOK) Error() string {
	return fmt.Sprintf("[POST /v1/edges][%d] edgeCreateOK  %+v", 200, o.Payload)
}

func (o *EdgeCreateOK) GetPayload() *models.CreateDocumentResponse {
	return o.Payload
}

func (o *EdgeCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateDocumentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeCreateDefault creates a EdgeCreateDefault with default headers values
func NewEdgeCreateDefault(code int) *EdgeCreateDefault {
	return &EdgeCreateDefault{
		_statusCode: code,
	}
}

/*EdgeCreateDefault handles this case with default header values.

generic API error response
*/
type EdgeCreateDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the edge create default response
func (o *EdgeCreateDefault) Code() int {
	return o._statusCode
}

func (o *EdgeCreateDefault) Error() string {
	return fmt.Sprintf("[POST /v1/edges][%d] EdgeCreate default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeCreateDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *EdgeCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
