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

// EdgeGetSensorsV2Reader is a Reader for the EdgeGetSensorsV2 structure.
type EdgeGetSensorsV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeGetSensorsV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeGetSensorsV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewEdgeGetSensorsV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEdgeGetSensorsV2OK creates a EdgeGetSensorsV2OK with default headers values
func NewEdgeGetSensorsV2OK() *EdgeGetSensorsV2OK {
	return &EdgeGetSensorsV2OK{}
}

/*EdgeGetSensorsV2OK handles this case with default header values.

Ok
*/
type EdgeGetSensorsV2OK struct {
	Payload *models.SensorListPayload
}

func (o *EdgeGetSensorsV2OK) Error() string {
	return fmt.Sprintf("[GET /v1.0/edges/{edgeId}/sensors][%d] edgeGetSensorsV2OK  %+v", 200, o.Payload)
}

func (o *EdgeGetSensorsV2OK) GetPayload() *models.SensorListPayload {
	return o.Payload
}

func (o *EdgeGetSensorsV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SensorListPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeGetSensorsV2Default creates a EdgeGetSensorsV2Default with default headers values
func NewEdgeGetSensorsV2Default(code int) *EdgeGetSensorsV2Default {
	return &EdgeGetSensorsV2Default{
		_statusCode: code,
	}
}

/*EdgeGetSensorsV2Default handles this case with default header values.

generic API error response
*/
type EdgeGetSensorsV2Default struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the edge get sensors v2 default response
func (o *EdgeGetSensorsV2Default) Code() int {
	return o._statusCode
}

func (o *EdgeGetSensorsV2Default) Error() string {
	return fmt.Sprintf("[GET /v1.0/edges/{edgeId}/sensors][%d] EdgeGetSensorsV2 default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeGetSensorsV2Default) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *EdgeGetSensorsV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
