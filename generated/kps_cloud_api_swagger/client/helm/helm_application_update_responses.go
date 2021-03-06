// Code generated by go-swagger; DO NOT EDIT.

package helm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"
)

// HelmApplicationUpdateReader is a Reader for the HelmApplicationUpdate structure.
type HelmApplicationUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HelmApplicationUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHelmApplicationUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewHelmApplicationUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewHelmApplicationUpdateOK creates a HelmApplicationUpdateOK with default headers values
func NewHelmApplicationUpdateOK() *HelmApplicationUpdateOK {
	return &HelmApplicationUpdateOK{}
}

/*HelmApplicationUpdateOK handles this case with default header values.

Ok
*/
type HelmApplicationUpdateOK struct {
	Payload *models.UpdateDocumentResponseV2
}

func (o *HelmApplicationUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /v1.0/helm/apps/{id}][%d] helmApplicationUpdateOK  %+v", 200, o.Payload)
}

func (o *HelmApplicationUpdateOK) GetPayload() *models.UpdateDocumentResponseV2 {
	return o.Payload
}

func (o *HelmApplicationUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UpdateDocumentResponseV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHelmApplicationUpdateDefault creates a HelmApplicationUpdateDefault with default headers values
func NewHelmApplicationUpdateDefault(code int) *HelmApplicationUpdateDefault {
	return &HelmApplicationUpdateDefault{
		_statusCode: code,
	}
}

/*HelmApplicationUpdateDefault handles this case with default header values.

generic API error response
*/
type HelmApplicationUpdateDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the helm application update default response
func (o *HelmApplicationUpdateDefault) Code() int {
	return o._statusCode
}

func (o *HelmApplicationUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /v1.0/helm/apps/{id}][%d] HelmApplicationUpdate default  %+v", o._statusCode, o.Payload)
}

func (o *HelmApplicationUpdateDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *HelmApplicationUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
