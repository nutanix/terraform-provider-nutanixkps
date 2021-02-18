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

// CategoryUpdateV2Reader is a Reader for the CategoryUpdateV2 structure.
type CategoryUpdateV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CategoryUpdateV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCategoryUpdateV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCategoryUpdateV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCategoryUpdateV2OK creates a CategoryUpdateV2OK with default headers values
func NewCategoryUpdateV2OK() *CategoryUpdateV2OK {
	return &CategoryUpdateV2OK{}
}

/*CategoryUpdateV2OK handles this case with default header values.

Ok
*/
type CategoryUpdateV2OK struct {
	Payload *models.UpdateDocumentResponse
}

func (o *CategoryUpdateV2OK) Error() string {
	return fmt.Sprintf("[PUT /v1/categories/{id}][%d] categoryUpdateV2OK  %+v", 200, o.Payload)
}

func (o *CategoryUpdateV2OK) GetPayload() *models.UpdateDocumentResponse {
	return o.Payload
}

func (o *CategoryUpdateV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UpdateDocumentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCategoryUpdateV2Default creates a CategoryUpdateV2Default with default headers values
func NewCategoryUpdateV2Default(code int) *CategoryUpdateV2Default {
	return &CategoryUpdateV2Default{
		_statusCode: code,
	}
}

/*CategoryUpdateV2Default handles this case with default header values.

generic API error response
*/
type CategoryUpdateV2Default struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the category update v2 default response
func (o *CategoryUpdateV2Default) Code() int {
	return o._statusCode
}

func (o *CategoryUpdateV2Default) Error() string {
	return fmt.Sprintf("[PUT /v1/categories/{id}][%d] CategoryUpdateV2 default  %+v", o._statusCode, o.Payload)
}

func (o *CategoryUpdateV2Default) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *CategoryUpdateV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
