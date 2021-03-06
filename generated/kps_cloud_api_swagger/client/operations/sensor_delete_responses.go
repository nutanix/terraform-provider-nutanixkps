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

// SensorDeleteReader is a Reader for the SensorDelete structure.
type SensorDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SensorDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSensorDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSensorDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSensorDeleteOK creates a SensorDeleteOK with default headers values
func NewSensorDeleteOK() *SensorDeleteOK {
	return &SensorDeleteOK{}
}

/*SensorDeleteOK handles this case with default header values.

Ok
*/
type SensorDeleteOK struct {
	Payload *models.DeleteDocumentResponse
}

func (o *SensorDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/sensors/{id}][%d] sensorDeleteOK  %+v", 200, o.Payload)
}

func (o *SensorDeleteOK) GetPayload() *models.DeleteDocumentResponse {
	return o.Payload
}

func (o *SensorDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeleteDocumentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSensorDeleteDefault creates a SensorDeleteDefault with default headers values
func NewSensorDeleteDefault(code int) *SensorDeleteDefault {
	return &SensorDeleteDefault{
		_statusCode: code,
	}
}

/*SensorDeleteDefault handles this case with default header values.

generic API error response
*/
type SensorDeleteDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the sensor delete default response
func (o *SensorDeleteDefault) Code() int {
	return o._statusCode
}

func (o *SensorDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /v1/sensors/{id}][%d] SensorDelete default  %+v", o._statusCode, o.Payload)
}

func (o *SensorDeleteDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *SensorDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
