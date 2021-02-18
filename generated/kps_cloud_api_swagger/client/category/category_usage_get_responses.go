// Code generated by go-swagger; DO NOT EDIT.

package category

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"
)

// CategoryUsageGetReader is a Reader for the CategoryUsageGet structure.
type CategoryUsageGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CategoryUsageGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCategoryUsageGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCategoryUsageGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCategoryUsageGetOK creates a CategoryUsageGetOK with default headers values
func NewCategoryUsageGetOK() *CategoryUsageGetOK {
	return &CategoryUsageGetOK{}
}

/*CategoryUsageGetOK handles this case with default header values.

Ok
*/
type CategoryUsageGetOK struct {
	Payload *models.CategoryDetailUsageInfo
}

func (o *CategoryUsageGetOK) Error() string {
	return fmt.Sprintf("[GET /v1.0/categoriesusage/{id}][%d] categoryUsageGetOK  %+v", 200, o.Payload)
}

func (o *CategoryUsageGetOK) GetPayload() *models.CategoryDetailUsageInfo {
	return o.Payload
}

func (o *CategoryUsageGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CategoryDetailUsageInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCategoryUsageGetDefault creates a CategoryUsageGetDefault with default headers values
func NewCategoryUsageGetDefault(code int) *CategoryUsageGetDefault {
	return &CategoryUsageGetDefault{
		_statusCode: code,
	}
}

/*CategoryUsageGetDefault handles this case with default header values.

generic API error response
*/
type CategoryUsageGetDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the category usage get default response
func (o *CategoryUsageGetDefault) Code() int {
	return o._statusCode
}

func (o *CategoryUsageGetDefault) Error() string {
	return fmt.Sprintf("[GET /v1.0/categoriesusage/{id}][%d] CategoryUsageGet default  %+v", o._statusCode, o.Payload)
}

func (o *CategoryUsageGetDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *CategoryUsageGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
