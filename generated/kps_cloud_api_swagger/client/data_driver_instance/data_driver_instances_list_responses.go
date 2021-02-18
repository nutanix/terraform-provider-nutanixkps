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

// DataDriverInstancesListReader is a Reader for the DataDriverInstancesList structure.
type DataDriverInstancesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DataDriverInstancesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDataDriverInstancesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDataDriverInstancesListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDataDriverInstancesListOK creates a DataDriverInstancesListOK with default headers values
func NewDataDriverInstancesListOK() *DataDriverInstancesListOK {
	return &DataDriverInstancesListOK{}
}

/*DataDriverInstancesListOK handles this case with default header values.

DataDriverInstanceListResponse is a a data driver instance listing response
*/
type DataDriverInstancesListOK struct {
	Payload *models.DataDriverInstanceListResponsePayload
}

func (o *DataDriverInstancesListOK) Error() string {
	return fmt.Sprintf("[GET /v1.0/datadriverinstances][%d] dataDriverInstancesListOK  %+v", 200, o.Payload)
}

func (o *DataDriverInstancesListOK) GetPayload() *models.DataDriverInstanceListResponsePayload {
	return o.Payload
}

func (o *DataDriverInstancesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DataDriverInstanceListResponsePayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDataDriverInstancesListDefault creates a DataDriverInstancesListDefault with default headers values
func NewDataDriverInstancesListDefault(code int) *DataDriverInstancesListDefault {
	return &DataDriverInstancesListDefault{
		_statusCode: code,
	}
}

/*DataDriverInstancesListDefault handles this case with default header values.

generic API error response
*/
type DataDriverInstancesListDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the data driver instances list default response
func (o *DataDriverInstancesListDefault) Code() int {
	return o._statusCode
}

func (o *DataDriverInstancesListDefault) Error() string {
	return fmt.Sprintf("[GET /v1.0/datadriverinstances][%d] DataDriverInstancesList default  %+v", o._statusCode, o.Payload)
}

func (o *DataDriverInstancesListDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *DataDriverInstancesListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
