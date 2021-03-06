// Code generated by go-swagger; DO NOT EDIT.

package container_registry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"
)

// ContainerRegistryListV2Reader is a Reader for the ContainerRegistryListV2 structure.
type ContainerRegistryListV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContainerRegistryListV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewContainerRegistryListV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewContainerRegistryListV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewContainerRegistryListV2OK creates a ContainerRegistryListV2OK with default headers values
func NewContainerRegistryListV2OK() *ContainerRegistryListV2OK {
	return &ContainerRegistryListV2OK{}
}

/*ContainerRegistryListV2OK handles this case with default header values.

Ok
*/
type ContainerRegistryListV2OK struct {
	Payload *models.ContainerRegistryListPayload
}

func (o *ContainerRegistryListV2OK) Error() string {
	return fmt.Sprintf("[GET /v1.0/containerregistries][%d] containerRegistryListV2OK  %+v", 200, o.Payload)
}

func (o *ContainerRegistryListV2OK) GetPayload() *models.ContainerRegistryListPayload {
	return o.Payload
}

func (o *ContainerRegistryListV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ContainerRegistryListPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerRegistryListV2Default creates a ContainerRegistryListV2Default with default headers values
func NewContainerRegistryListV2Default(code int) *ContainerRegistryListV2Default {
	return &ContainerRegistryListV2Default{
		_statusCode: code,
	}
}

/*ContainerRegistryListV2Default handles this case with default header values.

generic API error response
*/
type ContainerRegistryListV2Default struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the container registry list v2 default response
func (o *ContainerRegistryListV2Default) Code() int {
	return o._statusCode
}

func (o *ContainerRegistryListV2Default) Error() string {
	return fmt.Sprintf("[GET /v1.0/containerregistries][%d] ContainerRegistryListV2 default  %+v", o._statusCode, o.Payload)
}

func (o *ContainerRegistryListV2Default) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *ContainerRegistryListV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
