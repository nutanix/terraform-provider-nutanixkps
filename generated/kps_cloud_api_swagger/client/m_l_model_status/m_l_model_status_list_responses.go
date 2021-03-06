// Code generated by go-swagger; DO NOT EDIT.

package m_l_model_status

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"
)

// MLModelStatusListReader is a Reader for the MLModelStatusList structure.
type MLModelStatusListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MLModelStatusListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMLModelStatusListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewMLModelStatusListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewMLModelStatusListOK creates a MLModelStatusListOK with default headers values
func NewMLModelStatusListOK() *MLModelStatusListOK {
	return &MLModelStatusListOK{}
}

/*MLModelStatusListOK handles this case with default header values.

Ok
*/
type MLModelStatusListOK struct {
	Payload *models.MLModelStatusListPayload
}

func (o *MLModelStatusListOK) Error() string {
	return fmt.Sprintf("[GET /v1.0/mlmodelstatuses][%d] mLModelStatusListOK  %+v", 200, o.Payload)
}

func (o *MLModelStatusListOK) GetPayload() *models.MLModelStatusListPayload {
	return o.Payload
}

func (o *MLModelStatusListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MLModelStatusListPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMLModelStatusListDefault creates a MLModelStatusListDefault with default headers values
func NewMLModelStatusListDefault(code int) *MLModelStatusListDefault {
	return &MLModelStatusListDefault{
		_statusCode: code,
	}
}

/*MLModelStatusListDefault handles this case with default header values.

generic API error response
*/
type MLModelStatusListDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the m l model status list default response
func (o *MLModelStatusListDefault) Code() int {
	return o._statusCode
}

func (o *MLModelStatusListDefault) Error() string {
	return fmt.Sprintf("[GET /v1.0/mlmodelstatuses][%d] MLModelStatusList default  %+v", o._statusCode, o.Payload)
}

func (o *MLModelStatusListDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *MLModelStatusListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
