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

// EdgeGetSensorsReader is a Reader for the EdgeGetSensors structure.
type EdgeGetSensorsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeGetSensorsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeGetSensorsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewEdgeGetSensorsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEdgeGetSensorsOK creates a EdgeGetSensorsOK with default headers values
func NewEdgeGetSensorsOK() *EdgeGetSensorsOK {
	return &EdgeGetSensorsOK{}
}

/*EdgeGetSensorsOK handles this case with default header values.

Ok
*/
type EdgeGetSensorsOK struct {
	Payload []*models.Sensor
}

func (o *EdgeGetSensorsOK) Error() string {
	return fmt.Sprintf("[GET /v1/edges/{edgeId}/sensors][%d] edgeGetSensorsOK  %+v", 200, o.Payload)
}

func (o *EdgeGetSensorsOK) GetPayload() []*models.Sensor {
	return o.Payload
}

func (o *EdgeGetSensorsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeGetSensorsDefault creates a EdgeGetSensorsDefault with default headers values
func NewEdgeGetSensorsDefault(code int) *EdgeGetSensorsDefault {
	return &EdgeGetSensorsDefault{
		_statusCode: code,
	}
}

/*EdgeGetSensorsDefault handles this case with default header values.

generic API error response
*/
type EdgeGetSensorsDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the edge get sensors default response
func (o *EdgeGetSensorsDefault) Code() int {
	return o._statusCode
}

func (o *EdgeGetSensorsDefault) Error() string {
	return fmt.Sprintf("[GET /v1/edges/{edgeId}/sensors][%d] EdgeGetSensors default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeGetSensorsDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *EdgeGetSensorsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
