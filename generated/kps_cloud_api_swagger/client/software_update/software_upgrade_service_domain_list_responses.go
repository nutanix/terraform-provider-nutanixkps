// Code generated by go-swagger; DO NOT EDIT.

package software_update

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"
)

// SoftwareUpgradeServiceDomainListReader is a Reader for the SoftwareUpgradeServiceDomainList structure.
type SoftwareUpgradeServiceDomainListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SoftwareUpgradeServiceDomainListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSoftwareUpgradeServiceDomainListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSoftwareUpgradeServiceDomainListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSoftwareUpgradeServiceDomainListOK creates a SoftwareUpgradeServiceDomainListOK with default headers values
func NewSoftwareUpgradeServiceDomainListOK() *SoftwareUpgradeServiceDomainListOK {
	return &SoftwareUpgradeServiceDomainListOK{}
}

/*SoftwareUpgradeServiceDomainListOK handles this case with default header values.

Ok
*/
type SoftwareUpgradeServiceDomainListOK struct {
	Payload *models.SoftwareUpdateServiceDomainListPayload
}

func (o *SoftwareUpgradeServiceDomainListOK) Error() string {
	return fmt.Sprintf("[GET /v1.0/softwareupdates/upgrades/{batchId}/servicedomains][%d] softwareUpgradeServiceDomainListOK  %+v", 200, o.Payload)
}

func (o *SoftwareUpgradeServiceDomainListOK) GetPayload() *models.SoftwareUpdateServiceDomainListPayload {
	return o.Payload
}

func (o *SoftwareUpgradeServiceDomainListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SoftwareUpdateServiceDomainListPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSoftwareUpgradeServiceDomainListDefault creates a SoftwareUpgradeServiceDomainListDefault with default headers values
func NewSoftwareUpgradeServiceDomainListDefault(code int) *SoftwareUpgradeServiceDomainListDefault {
	return &SoftwareUpgradeServiceDomainListDefault{
		_statusCode: code,
	}
}

/*SoftwareUpgradeServiceDomainListDefault handles this case with default header values.

generic API error response
*/
type SoftwareUpgradeServiceDomainListDefault struct {
	_statusCode int

	Payload *models.APIErrorPayload
}

// Code gets the status code for the software upgrade service domain list default response
func (o *SoftwareUpgradeServiceDomainListDefault) Code() int {
	return o._statusCode
}

func (o *SoftwareUpgradeServiceDomainListDefault) Error() string {
	return fmt.Sprintf("[GET /v1.0/softwareupdates/upgrades/{batchId}/servicedomains][%d] SoftwareUpgradeServiceDomainList default  %+v", o._statusCode, o.Payload)
}

func (o *SoftwareUpgradeServiceDomainListDefault) GetPayload() *models.APIErrorPayload {
	return o.Payload
}

func (o *SoftwareUpgradeServiceDomainListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorPayload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
