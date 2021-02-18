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

// GetThreeScaleServiceReader is a Reader for the GetThreeScaleService structure.
type GetThreeScaleServiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetThreeScaleServiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetThreeScaleServiceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetThreeScaleServiceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetThreeScaleServiceOK creates a GetThreeScaleServiceOK with default headers values
func NewGetThreeScaleServiceOK() *GetThreeScaleServiceOK {
	return &GetThreeScaleServiceOK{}
}

/*GetThreeScaleServiceOK handles this case with default header values.

Return Threescale rule definition for a given service
*/
type GetThreeScaleServiceOK struct {
	Payload *models.ThreeScaleServiceRule
}

func (o *GetThreeScaleServiceOK) Error() string {
	return fmt.Sprintf("[GET /v1.0/kiali/threescale/namespaces/{namespace}/services/{service}][%d] getThreeScaleServiceOK  %+v", 200, o.Payload)
}

func (o *GetThreeScaleServiceOK) GetPayload() *models.ThreeScaleServiceRule {
	return o.Payload
}

func (o *GetThreeScaleServiceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ThreeScaleServiceRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetThreeScaleServiceDefault creates a GetThreeScaleServiceDefault with default headers values
func NewGetThreeScaleServiceDefault(code int) *GetThreeScaleServiceDefault {
	return &GetThreeScaleServiceDefault{
		_statusCode: code,
	}
}

/*GetThreeScaleServiceDefault handles this case with default header values.

generic API error response
*/
type GetThreeScaleServiceDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the get three scale service default response
func (o *GetThreeScaleServiceDefault) Code() int {
	return o._statusCode
}

func (o *GetThreeScaleServiceDefault) Error() string {
	return fmt.Sprintf("[GET /v1.0/kiali/threescale/namespaces/{namespace}/services/{service}][%d] getThreeScaleService default  %+v", o._statusCode, o.Payload)
}

func (o *GetThreeScaleServiceDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *GetThreeScaleServiceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
