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

// PostIter8ExperimentsReader is a Reader for the PostIter8Experiments structure.
type PostIter8ExperimentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostIter8ExperimentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostIter8ExperimentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostIter8ExperimentsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostIter8ExperimentsOK creates a PostIter8ExperimentsOK with default headers values
func NewPostIter8ExperimentsOK() *PostIter8ExperimentsOK {
	return &PostIter8ExperimentsOK{}
}

/*PostIter8ExperimentsOK handles this case with default header values.

Return a Iter8 Experiment detail
*/
type PostIter8ExperimentsOK struct {
	Payload *models.Iter8ExperimentDetail
}

func (o *PostIter8ExperimentsOK) Error() string {
	return fmt.Sprintf("[POST /v1.0/kiali/iter8/namespaces/{namespace}/experiments][%d] postIter8ExperimentsOK  %+v", 200, o.Payload)
}

func (o *PostIter8ExperimentsOK) GetPayload() *models.Iter8ExperimentDetail {
	return o.Payload
}

func (o *PostIter8ExperimentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Iter8ExperimentDetail)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostIter8ExperimentsDefault creates a PostIter8ExperimentsDefault with default headers values
func NewPostIter8ExperimentsDefault(code int) *PostIter8ExperimentsDefault {
	return &PostIter8ExperimentsDefault{
		_statusCode: code,
	}
}

/*PostIter8ExperimentsDefault handles this case with default header values.

generic API error response
*/
type PostIter8ExperimentsDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the post iter8 experiments default response
func (o *PostIter8ExperimentsDefault) Code() int {
	return o._statusCode
}

func (o *PostIter8ExperimentsDefault) Error() string {
	return fmt.Sprintf("[POST /v1.0/kiali/iter8/namespaces/{namespace}/experiments][%d] postIter8Experiments default  %+v", o._statusCode, o.Payload)
}

func (o *PostIter8ExperimentsDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *PostIter8ExperimentsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
