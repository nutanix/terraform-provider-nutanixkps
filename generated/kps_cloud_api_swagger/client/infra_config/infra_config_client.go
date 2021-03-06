// Code generated by go-swagger; DO NOT EDIT.

package infra_config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new infra config API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for infra config API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
InfraConfigGet gets a edge cluster config by its ID ntnx ignore

Retrieves a edgeCluster config with the given ID {id}.
*/
func (a *Client) InfraConfigGet(params *InfraConfigGetParams, authInfo runtime.ClientAuthInfoWriter) (*InfraConfigGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInfraConfigGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "InfraConfigGet",
		Method:             "GET",
		PathPattern:        "/v1.0/edgecluster/{id}/infraconfig",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &InfraConfigGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*InfraConfigGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*InfraConfigGetDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
