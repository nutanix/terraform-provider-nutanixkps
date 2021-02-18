// Code generated by go-swagger; DO NOT EDIT.

package sensor

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"
)

// SensorCreateV2Reader is a Reader for the SensorCreateV2 structure.
type SensorCreateV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SensorCreateV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSensorCreateV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSensorCreateV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSensorCreateV2OK creates a SensorCreateV2OK with default headers values
func NewSensorCreateV2OK() *SensorCreateV2OK {
	return &SensorCreateV2OK{}
}

/*SensorCreateV2OK handles this case with default header values.

Ok
*/
type SensorCreateV2OK struct {
	Payload *models.CreateDocumentResponseV2
}

func (o *SensorCreateV2OK) Error() string {
	return fmt.Sprintf("[POST /v1.0/sensors][%d] sensorCreateV2OK  %+v", 200, o.Payload)
}

func (o *SensorCreateV2OK) GetPayload() *models.CreateDocumentResponseV2 {
	return o.Payload
}

func (o *SensorCreateV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateDocumentResponseV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSensorCreateV2Default creates a SensorCreateV2Default with default headers values
func NewSensorCreateV2Default(code int) *SensorCreateV2Default {
	return &SensorCreateV2Default{
		_statusCode: code,
	}
}

/*SensorCreateV2Default handles this case with default header values.

generic API error response
*/
type SensorCreateV2Default struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the sensor create v2 default response
func (o *SensorCreateV2Default) Code() int {
	return o._statusCode
}

func (o *SensorCreateV2Default) Error() string {
	return fmt.Sprintf("[POST /v1.0/sensors][%d] SensorCreateV2 default  %+v", o._statusCode, o.Payload)
}

func (o *SensorCreateV2Default) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *SensorCreateV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
