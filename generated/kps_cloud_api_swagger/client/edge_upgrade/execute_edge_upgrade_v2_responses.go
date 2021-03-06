// Code generated by go-swagger; DO NOT EDIT.

package edge_upgrade

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"
)

// ExecuteEdgeUpgradeV2Reader is a Reader for the ExecuteEdgeUpgradeV2 structure.
type ExecuteEdgeUpgradeV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExecuteEdgeUpgradeV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExecuteEdgeUpgradeV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExecuteEdgeUpgradeV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExecuteEdgeUpgradeV2OK creates a ExecuteEdgeUpgradeV2OK with default headers values
func NewExecuteEdgeUpgradeV2OK() *ExecuteEdgeUpgradeV2OK {
	return &ExecuteEdgeUpgradeV2OK{}
}

/*ExecuteEdgeUpgradeV2OK handles this case with default header values.

Ok
*/
type ExecuteEdgeUpgradeV2OK struct {
	Payload *models.CreateDocumentResponseV2
}

func (o *ExecuteEdgeUpgradeV2OK) Error() string {
	return fmt.Sprintf("[POST /v1.0/edges/upgrade][%d] executeEdgeUpgradeV2OK  %+v", 200, o.Payload)
}

func (o *ExecuteEdgeUpgradeV2OK) GetPayload() *models.CreateDocumentResponseV2 {
	return o.Payload
}

func (o *ExecuteEdgeUpgradeV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateDocumentResponseV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExecuteEdgeUpgradeV2Default creates a ExecuteEdgeUpgradeV2Default with default headers values
func NewExecuteEdgeUpgradeV2Default(code int) *ExecuteEdgeUpgradeV2Default {
	return &ExecuteEdgeUpgradeV2Default{
		_statusCode: code,
	}
}

/*ExecuteEdgeUpgradeV2Default handles this case with default header values.

generic API error response
*/
type ExecuteEdgeUpgradeV2Default struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the execute edge upgrade v2 default response
func (o *ExecuteEdgeUpgradeV2Default) Code() int {
	return o._statusCode
}

func (o *ExecuteEdgeUpgradeV2Default) Error() string {
	return fmt.Sprintf("[POST /v1.0/edges/upgrade][%d] ExecuteEdgeUpgradeV2 default  %+v", o._statusCode, o.Payload)
}

func (o *ExecuteEdgeUpgradeV2Default) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *ExecuteEdgeUpgradeV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
