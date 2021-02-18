// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"
)

// DataSourceDeleteReader is a Reader for the DataSourceDelete structure.
type DataSourceDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DataSourceDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDataSourceDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDataSourceDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDataSourceDeleteOK creates a DataSourceDeleteOK with default headers values
func NewDataSourceDeleteOK() *DataSourceDeleteOK {
	return &DataSourceDeleteOK{}
}

/*DataSourceDeleteOK handles this case with default header values.

Ok
*/
type DataSourceDeleteOK struct {
	Payload *models.DeleteDocumentResponse
}

func (o *DataSourceDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/datasources/{id}][%d] dataSourceDeleteOK  %+v", 200, o.Payload)
}

func (o *DataSourceDeleteOK) GetPayload() *models.DeleteDocumentResponse {
	return o.Payload
}

func (o *DataSourceDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeleteDocumentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDataSourceDeleteDefault creates a DataSourceDeleteDefault with default headers values
func NewDataSourceDeleteDefault(code int) *DataSourceDeleteDefault {
	return &DataSourceDeleteDefault{
		_statusCode: code,
	}
}

/*DataSourceDeleteDefault handles this case with default header values.

generic API error response
*/
type DataSourceDeleteDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the data source delete default response
func (o *DataSourceDeleteDefault) Code() int {
	return o._statusCode
}

func (o *DataSourceDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /v1/datasources/{id}][%d] DataSourceDelete default  %+v", o._statusCode, o.Payload)
}

func (o *DataSourceDeleteDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *DataSourceDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
