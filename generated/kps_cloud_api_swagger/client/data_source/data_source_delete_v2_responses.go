// Code generated by go-swagger; DO NOT EDIT.

package data_source

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"
)

// DataSourceDeleteV2Reader is a Reader for the DataSourceDeleteV2 structure.
type DataSourceDeleteV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DataSourceDeleteV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDataSourceDeleteV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDataSourceDeleteV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDataSourceDeleteV2OK creates a DataSourceDeleteV2OK with default headers values
func NewDataSourceDeleteV2OK() *DataSourceDeleteV2OK {
	return &DataSourceDeleteV2OK{}
}

/*DataSourceDeleteV2OK handles this case with default header values.

Ok
*/
type DataSourceDeleteV2OK struct {
	Payload *models.DeleteDocumentResponseV2
}

func (o *DataSourceDeleteV2OK) Error() string {
	return fmt.Sprintf("[DELETE /v1.0/datasources/{id}][%d] dataSourceDeleteV2OK  %+v", 200, o.Payload)
}

func (o *DataSourceDeleteV2OK) GetPayload() *models.DeleteDocumentResponseV2 {
	return o.Payload
}

func (o *DataSourceDeleteV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeleteDocumentResponseV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDataSourceDeleteV2Default creates a DataSourceDeleteV2Default with default headers values
func NewDataSourceDeleteV2Default(code int) *DataSourceDeleteV2Default {
	return &DataSourceDeleteV2Default{
		_statusCode: code,
	}
}

/*DataSourceDeleteV2Default handles this case with default header values.

generic API error response
*/
type DataSourceDeleteV2Default struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the data source delete v2 default response
func (o *DataSourceDeleteV2Default) Code() int {
	return o._statusCode
}

func (o *DataSourceDeleteV2Default) Error() string {
	return fmt.Sprintf("[DELETE /v1.0/datasources/{id}][%d] DataSourceDeleteV2 default  %+v", o._statusCode, o.Payload)
}

func (o *DataSourceDeleteV2Default) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *DataSourceDeleteV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
