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

// DataStreamDeleteReader is a Reader for the DataStreamDelete structure.
type DataStreamDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DataStreamDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDataStreamDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDataStreamDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDataStreamDeleteOK creates a DataStreamDeleteOK with default headers values
func NewDataStreamDeleteOK() *DataStreamDeleteOK {
	return &DataStreamDeleteOK{}
}

/*DataStreamDeleteOK handles this case with default header values.

Ok
*/
type DataStreamDeleteOK struct {
	Payload *models.DeleteDocumentResponse
}

func (o *DataStreamDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/datastreams/{id}][%d] dataStreamDeleteOK  %+v", 200, o.Payload)
}

func (o *DataStreamDeleteOK) GetPayload() *models.DeleteDocumentResponse {
	return o.Payload
}

func (o *DataStreamDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeleteDocumentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDataStreamDeleteDefault creates a DataStreamDeleteDefault with default headers values
func NewDataStreamDeleteDefault(code int) *DataStreamDeleteDefault {
	return &DataStreamDeleteDefault{
		_statusCode: code,
	}
}

/*DataStreamDeleteDefault handles this case with default header values.

generic API error response
*/
type DataStreamDeleteDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the data stream delete default response
func (o *DataStreamDeleteDefault) Code() int {
	return o._statusCode
}

func (o *DataStreamDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /v1/datastreams/{id}][%d] DataStreamDelete default  %+v", o._statusCode, o.Payload)
}

func (o *DataStreamDeleteDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *DataStreamDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
