// Code generated by go-swagger; DO NOT EDIT.

package auditlog

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"
)

// QueryAuditLogsV2Reader is a Reader for the QueryAuditLogsV2 structure.
type QueryAuditLogsV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryAuditLogsV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryAuditLogsV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewQueryAuditLogsV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQueryAuditLogsV2OK creates a QueryAuditLogsV2OK with default headers values
func NewQueryAuditLogsV2OK() *QueryAuditLogsV2OK {
	return &QueryAuditLogsV2OK{}
}

/*QueryAuditLogsV2OK handles this case with default header values.

Ok
*/
type QueryAuditLogsV2OK struct {
	Payload []*models.AuditLogV2
}

func (o *QueryAuditLogsV2OK) Error() string {
	return fmt.Sprintf("[POST /v1.0/auditlogsV2][%d] queryAuditLogsV2OK  %+v", 200, o.Payload)
}

func (o *QueryAuditLogsV2OK) GetPayload() []*models.AuditLogV2 {
	return o.Payload
}

func (o *QueryAuditLogsV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryAuditLogsV2Default creates a QueryAuditLogsV2Default with default headers values
func NewQueryAuditLogsV2Default(code int) *QueryAuditLogsV2Default {
	return &QueryAuditLogsV2Default{
		_statusCode: code,
	}
}

/*QueryAuditLogsV2Default handles this case with default header values.

generic API error response
*/
type QueryAuditLogsV2Default struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the query audit logs v2 default response
func (o *QueryAuditLogsV2Default) Code() int {
	return o._statusCode
}

func (o *QueryAuditLogsV2Default) Error() string {
	return fmt.Sprintf("[POST /v1.0/auditlogsV2][%d] QueryAuditLogsV2 default  %+v", o._statusCode, o.Payload)
}

func (o *QueryAuditLogsV2Default) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *QueryAuditLogsV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
