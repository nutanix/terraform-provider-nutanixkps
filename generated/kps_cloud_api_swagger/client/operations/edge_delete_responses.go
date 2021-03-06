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

// EdgeDeleteReader is a Reader for the EdgeDelete structure.
type EdgeDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewEdgeDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEdgeDeleteOK creates a EdgeDeleteOK with default headers values
func NewEdgeDeleteOK() *EdgeDeleteOK {
	return &EdgeDeleteOK{}
}

/*EdgeDeleteOK handles this case with default header values.

Ok
*/
type EdgeDeleteOK struct {
	Payload *models.DeleteDocumentResponse
}

func (o *EdgeDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/edges/{edgeId}][%d] edgeDeleteOK  %+v", 200, o.Payload)
}

func (o *EdgeDeleteOK) GetPayload() *models.DeleteDocumentResponse {
	return o.Payload
}

func (o *EdgeDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeleteDocumentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeDeleteDefault creates a EdgeDeleteDefault with default headers values
func NewEdgeDeleteDefault(code int) *EdgeDeleteDefault {
	return &EdgeDeleteDefault{
		_statusCode: code,
	}
}

/*EdgeDeleteDefault handles this case with default header values.

generic API error response
*/
type EdgeDeleteDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the edge delete default response
func (o *EdgeDeleteDefault) Code() int {
	return o._statusCode
}

func (o *EdgeDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /v1/edges/{edgeId}][%d] EdgeDelete default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeDeleteDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *EdgeDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
