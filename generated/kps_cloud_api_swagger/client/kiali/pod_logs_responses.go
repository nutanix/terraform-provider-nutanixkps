// Code generated by go-swagger; DO NOT EDIT.

package kiali

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"
)

// PodLogsReader is a Reader for the PodLogs structure.
type PodLogsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PodLogsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPodLogsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPodLogsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPodLogsOK creates a PodLogsOK with default headers values
func NewPodLogsOK() *PodLogsOK {
	return &PodLogsOK{}
}

/*PodLogsOK handles this case with default header values.

Listing all the information related to a workload
*/
type PodLogsOK struct {
	Payload *models.Workload
}

func (o *PodLogsOK) Error() string {
	return fmt.Sprintf("[GET /v1.0/kiali/namespaces/{namespace}/pods/{pod}/logs][%d] podLogsOK  %+v", 200, o.Payload)
}

func (o *PodLogsOK) GetPayload() *models.Workload {
	return o.Payload
}

func (o *PodLogsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Workload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPodLogsDefault creates a PodLogsDefault with default headers values
func NewPodLogsDefault(code int) *PodLogsDefault {
	return &PodLogsDefault{
		_statusCode: code,
	}
}

/*PodLogsDefault handles this case with default header values.

generic API error response
*/
type PodLogsDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the pod logs default response
func (o *PodLogsDefault) Code() int {
	return o._statusCode
}

func (o *PodLogsDefault) Error() string {
	return fmt.Sprintf("[GET /v1.0/kiali/namespaces/{namespace}/pods/{pod}/logs][%d] podLogs default  %+v", o._statusCode, o.Payload)
}

func (o *PodLogsDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *PodLogsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
