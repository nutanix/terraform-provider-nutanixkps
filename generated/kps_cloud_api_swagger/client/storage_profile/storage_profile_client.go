// Code generated by go-swagger; DO NOT EDIT.

package storage_profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new storage profile API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for storage profile API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
StorageProfileCreate creates a storage profile ntnx ignore

Create a storage profile on the given service domain ID {svcDomainId}.
*/
func (a *Client) StorageProfileCreate(params *StorageProfileCreateParams, authInfo runtime.ClientAuthInfoWriter) (*StorageProfileCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStorageProfileCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "StorageProfileCreate",
		Method:             "POST",
		PathPattern:        "/v1.0/servicedomains/{svcDomainId}/storageprofiles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &StorageProfileCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*StorageProfileCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*StorageProfileCreateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
StorageProfileUpdate updates storage profile ntnx ignore

Update the storage profile with {id} on the given service domain ID {svcDomainId}.
*/
func (a *Client) StorageProfileUpdate(params *StorageProfileUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*StorageProfileUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStorageProfileUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "StorageProfileUpdate",
		Method:             "PUT",
		PathPattern:        "/v1.0/servicedomains/{svcDomainId}/storageprofiles/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &StorageProfileUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*StorageProfileUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*StorageProfileUpdateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
SvcDomainGetStorageProfiles gets storage profiles according to service domain ID ntnx ignore

Retrieves all storage profiles for a service domain with a given ID {svcDomainId}
*/
func (a *Client) SvcDomainGetStorageProfiles(params *SvcDomainGetStorageProfilesParams, authInfo runtime.ClientAuthInfoWriter) (*SvcDomainGetStorageProfilesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSvcDomainGetStorageProfilesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SvcDomainGetStorageProfiles",
		Method:             "GET",
		PathPattern:        "/v1.0/servicedomains/{svcDomainId}/storageprofiles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SvcDomainGetStorageProfilesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SvcDomainGetStorageProfilesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SvcDomainGetStorageProfilesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
