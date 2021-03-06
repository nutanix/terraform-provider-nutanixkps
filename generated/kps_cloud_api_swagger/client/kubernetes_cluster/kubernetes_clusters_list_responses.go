// Code generated by go-swagger; DO NOT EDIT.

package kubernetes_cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"
)

// KubernetesClustersListReader is a Reader for the KubernetesClustersList structure.
type KubernetesClustersListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KubernetesClustersListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKubernetesClustersListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewKubernetesClustersListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewKubernetesClustersListOK creates a KubernetesClustersListOK with default headers values
func NewKubernetesClustersListOK() *KubernetesClustersListOK {
	return &KubernetesClustersListOK{}
}

/*KubernetesClustersListOK handles this case with default header values.

Ok
*/
type KubernetesClustersListOK struct {
	Payload *models.KubernetesClustersListResponsePayload
}

func (o *KubernetesClustersListOK) Error() string {
	return fmt.Sprintf("[GET /v1.0/kubernetesclusters][%d] kubernetesClustersListOK  %+v", 200, o.Payload)
}

func (o *KubernetesClustersListOK) GetPayload() *models.KubernetesClustersListResponsePayload {
	return o.Payload
}

func (o *KubernetesClustersListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KubernetesClustersListResponsePayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesClustersListDefault creates a KubernetesClustersListDefault with default headers values
func NewKubernetesClustersListDefault(code int) *KubernetesClustersListDefault {
	return &KubernetesClustersListDefault{
		_statusCode: code,
	}
}

/*KubernetesClustersListDefault handles this case with default header values.

generic API error response
*/
type KubernetesClustersListDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the kubernetes clusters list default response
func (o *KubernetesClustersListDefault) Code() int {
	return o._statusCode
}

func (o *KubernetesClustersListDefault) Error() string {
	return fmt.Sprintf("[GET /v1.0/kubernetesclusters][%d] KubernetesClustersList default  %+v", o._statusCode, o.Payload)
}

func (o *KubernetesClustersListDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *KubernetesClustersListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
