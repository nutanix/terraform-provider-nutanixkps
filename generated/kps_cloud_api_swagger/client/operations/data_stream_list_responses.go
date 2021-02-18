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

// DataStreamListReader is a Reader for the DataStreamList structure.
type DataStreamListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DataStreamListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDataStreamListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDataStreamListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDataStreamListOK creates a DataStreamListOK with default headers values
func NewDataStreamListOK() *DataStreamListOK {
	return &DataStreamListOK{}
}

/*DataStreamListOK handles this case with default header values.

Ok
*/
type DataStreamListOK struct {
	Payload []*models.DataStream
}

func (o *DataStreamListOK) Error() string {
	return fmt.Sprintf("[GET /v1/datastreams][%d] dataStreamListOK  %+v", 200, o.Payload)
}

func (o *DataStreamListOK) GetPayload() []*models.DataStream {
	return o.Payload
}

func (o *DataStreamListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDataStreamListDefault creates a DataStreamListDefault with default headers values
func NewDataStreamListDefault(code int) *DataStreamListDefault {
	return &DataStreamListDefault{
		_statusCode: code,
	}
}

/*DataStreamListDefault handles this case with default header values.

generic API error response
*/
type DataStreamListDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the data stream list default response
func (o *DataStreamListDefault) Code() int {
	return o._statusCode
}

func (o *DataStreamListDefault) Error() string {
	return fmt.Sprintf("[GET /v1/datastreams][%d] DataStreamList default  %+v", o._statusCode, o.Payload)
}

func (o *DataStreamListDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *DataStreamListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
