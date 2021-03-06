// Code generated by go-swagger; DO NOT EDIT.

package log

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"
)

// LogRequestDownloadV2Reader is a Reader for the LogRequestDownloadV2 structure.
type LogRequestDownloadV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LogRequestDownloadV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLogRequestDownloadV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLogRequestDownloadV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLogRequestDownloadV2OK creates a LogRequestDownloadV2OK with default headers values
func NewLogRequestDownloadV2OK() *LogRequestDownloadV2OK {
	return &LogRequestDownloadV2OK{}
}

/*LogRequestDownloadV2OK handles this case with default header values.

Ok
*/
type LogRequestDownloadV2OK struct {
	Payload *models.LogDownloadPayload
}

func (o *LogRequestDownloadV2OK) Error() string {
	return fmt.Sprintf("[POST /v1.0/logs/requestdownload][%d] logRequestDownloadV2OK  %+v", 200, o.Payload)
}

func (o *LogRequestDownloadV2OK) GetPayload() *models.LogDownloadPayload {
	return o.Payload
}

func (o *LogRequestDownloadV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LogDownloadPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLogRequestDownloadV2Default creates a LogRequestDownloadV2Default with default headers values
func NewLogRequestDownloadV2Default(code int) *LogRequestDownloadV2Default {
	return &LogRequestDownloadV2Default{
		_statusCode: code,
	}
}

/*LogRequestDownloadV2Default handles this case with default header values.

generic API error response
*/
type LogRequestDownloadV2Default struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the log request download v2 default response
func (o *LogRequestDownloadV2Default) Code() int {
	return o._statusCode
}

func (o *LogRequestDownloadV2Default) Error() string {
	return fmt.Sprintf("[POST /v1.0/logs/requestdownload][%d] LogRequestDownloadV2 default  %+v", o._statusCode, o.Payload)
}

func (o *LogRequestDownloadV2Default) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *LogRequestDownloadV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
