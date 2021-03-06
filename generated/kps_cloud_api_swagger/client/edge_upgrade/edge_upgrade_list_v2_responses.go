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

// EdgeUpgradeListV2Reader is a Reader for the EdgeUpgradeListV2 structure.
type EdgeUpgradeListV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeUpgradeListV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeUpgradeListV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewEdgeUpgradeListV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEdgeUpgradeListV2OK creates a EdgeUpgradeListV2OK with default headers values
func NewEdgeUpgradeListV2OK() *EdgeUpgradeListV2OK {
	return &EdgeUpgradeListV2OK{}
}

/*EdgeUpgradeListV2OK handles this case with default header values.

Ok
*/
type EdgeUpgradeListV2OK struct {
	Payload *models.EdgeUpgradeListPayload
}

func (o *EdgeUpgradeListV2OK) Error() string {
	return fmt.Sprintf("[GET /v1.0/edgescompatibleupgrades][%d] edgeUpgradeListV2OK  %+v", 200, o.Payload)
}

func (o *EdgeUpgradeListV2OK) GetPayload() *models.EdgeUpgradeListPayload {
	return o.Payload
}

func (o *EdgeUpgradeListV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EdgeUpgradeListPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeUpgradeListV2Default creates a EdgeUpgradeListV2Default with default headers values
func NewEdgeUpgradeListV2Default(code int) *EdgeUpgradeListV2Default {
	return &EdgeUpgradeListV2Default{
		_statusCode: code,
	}
}

/*EdgeUpgradeListV2Default handles this case with default header values.

generic API error response
*/
type EdgeUpgradeListV2Default struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the edge upgrade list v2 default response
func (o *EdgeUpgradeListV2Default) Code() int {
	return o._statusCode
}

func (o *EdgeUpgradeListV2Default) Error() string {
	return fmt.Sprintf("[GET /v1.0/edgescompatibleupgrades][%d] EdgeUpgradeListV2 default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeUpgradeListV2Default) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *EdgeUpgradeListV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
