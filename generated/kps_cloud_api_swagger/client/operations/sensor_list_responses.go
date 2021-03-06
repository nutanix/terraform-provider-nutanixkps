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

// SensorListReader is a Reader for the SensorList structure.
type SensorListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SensorListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSensorListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSensorListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSensorListOK creates a SensorListOK with default headers values
func NewSensorListOK() *SensorListOK {
	return &SensorListOK{}
}

/*SensorListOK handles this case with default header values.

Ok
*/
type SensorListOK struct {
	Payload []*models.Sensor
}

func (o *SensorListOK) Error() string {
	return fmt.Sprintf("[GET /v1/sensors][%d] sensorListOK  %+v", 200, o.Payload)
}

func (o *SensorListOK) GetPayload() []*models.Sensor {
	return o.Payload
}

func (o *SensorListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSensorListDefault creates a SensorListDefault with default headers values
func NewSensorListDefault(code int) *SensorListDefault {
	return &SensorListDefault{
		_statusCode: code,
	}
}

/*SensorListDefault handles this case with default header values.

generic API error response
*/
type SensorListDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the sensor list default response
func (o *SensorListDefault) Code() int {
	return o._statusCode
}

func (o *SensorListDefault) Error() string {
	return fmt.Sprintf("[GET /v1/sensors][%d] SensorList default  %+v", o._statusCode, o.Payload)
}

func (o *SensorListDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *SensorListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
