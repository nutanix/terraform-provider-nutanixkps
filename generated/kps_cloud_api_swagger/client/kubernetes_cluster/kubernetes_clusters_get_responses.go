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

// KubernetesClustersGetReader is a Reader for the KubernetesClustersGet structure.
type KubernetesClustersGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KubernetesClustersGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKubernetesClustersGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewKubernetesClustersGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewKubernetesClustersGetOK creates a KubernetesClustersGetOK with default headers values
func NewKubernetesClustersGetOK() *KubernetesClustersGetOK {
	return &KubernetesClustersGetOK{}
}

/*KubernetesClustersGetOK handles this case with default header values.

Ok
*/
type KubernetesClustersGetOK struct {
	Payload *models.KubernetesCluster
}

func (o *KubernetesClustersGetOK) Error() string {
	return fmt.Sprintf("[GET /v1.0/kubernetesclusters/{id}][%d] kubernetesClustersGetOK  %+v", 200, o.Payload)
}

func (o *KubernetesClustersGetOK) GetPayload() *models.KubernetesCluster {
	return o.Payload
}

func (o *KubernetesClustersGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KubernetesCluster)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewKubernetesClustersGetDefault creates a KubernetesClustersGetDefault with default headers values
func NewKubernetesClustersGetDefault(code int) *KubernetesClustersGetDefault {
	return &KubernetesClustersGetDefault{
		_statusCode: code,
	}
}

/*KubernetesClustersGetDefault handles this case with default header values.

generic API error response
*/
type KubernetesClustersGetDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the kubernetes clusters get default response
func (o *KubernetesClustersGetDefault) Code() int {
	return o._statusCode
}

func (o *KubernetesClustersGetDefault) Error() string {
	return fmt.Sprintf("[GET /v1.0/kubernetesclusters/{id}][%d] KubernetesClustersGet default  %+v", o._statusCode, o.Payload)
}

func (o *KubernetesClustersGetDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *KubernetesClustersGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
