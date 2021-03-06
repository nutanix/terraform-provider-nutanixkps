// Code generated by go-swagger; DO NOT EDIT.

package data_driver_instance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"
)

// DataDriverInstanceCreateReader is a Reader for the DataDriverInstanceCreate structure.
type DataDriverInstanceCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DataDriverInstanceCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDataDriverInstanceCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDataDriverInstanceCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDataDriverInstanceCreateOK creates a DataDriverInstanceCreateOK with default headers values
func NewDataDriverInstanceCreateOK() *DataDriverInstanceCreateOK {
	return &DataDriverInstanceCreateOK{}
}

/*DataDriverInstanceCreateOK handles this case with default header values.

Ok
*/
type DataDriverInstanceCreateOK struct {
	Payload *models.CreateDocumentResponseV2
}

func (o *DataDriverInstanceCreateOK) Error() string {
	return fmt.Sprintf("[POST /v1.0/datadriverinstances][%d] dataDriverInstanceCreateOK  %+v", 200, o.Payload)
}

func (o *DataDriverInstanceCreateOK) GetPayload() *models.CreateDocumentResponseV2 {
	return o.Payload
}

func (o *DataDriverInstanceCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateDocumentResponseV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDataDriverInstanceCreateDefault creates a DataDriverInstanceCreateDefault with default headers values
func NewDataDriverInstanceCreateDefault(code int) *DataDriverInstanceCreateDefault {
	return &DataDriverInstanceCreateDefault{
		_statusCode: code,
	}
}

/*DataDriverInstanceCreateDefault handles this case with default header values.

generic API error response
*/
type DataDriverInstanceCreateDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the data driver instance create default response
func (o *DataDriverInstanceCreateDefault) Code() int {
	return o._statusCode
}

func (o *DataDriverInstanceCreateDefault) Error() string {
	return fmt.Sprintf("[POST /v1.0/datadriverinstances][%d] DataDriverInstanceCreate default  %+v", o._statusCode, o.Payload)
}

func (o *DataDriverInstanceCreateDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *DataDriverInstanceCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
