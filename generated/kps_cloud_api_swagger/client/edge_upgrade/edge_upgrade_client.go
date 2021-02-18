// Code generated by go-swagger; DO NOT EDIT.

package edge_upgrade

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new edge upgrade API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for edge upgrade API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
EdgeGetUpgradesV2 lists compatible edge software upgrades by edge ID ntnx ignore

Retrieves compatible software upgrades for the given edge ID.
*/
func (a *Client) EdgeGetUpgradesV2(params *EdgeGetUpgradesV2Params, authInfo runtime.ClientAuthInfoWriter) (*EdgeGetUpgradesV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEdgeGetUpgradesV2Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "EdgeGetUpgradesV2",
		Method:             "GET",
		PathPattern:        "/v1.0/edges/{edgeId}/upgradecompatible",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &EdgeGetUpgradesV2Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*EdgeGetUpgradesV2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*EdgeGetUpgradesV2Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
EdgeUpgradeListV2 lists available edge software upgrades ntnx ignore

Lists all possible software upgrades that are available for all detected edges.
*/
func (a *Client) EdgeUpgradeListV2(params *EdgeUpgradeListV2Params, authInfo runtime.ClientAuthInfoWriter) (*EdgeUpgradeListV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEdgeUpgradeListV2Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "EdgeUpgradeListV2",
		Method:             "GET",
		PathPattern:        "/v1.0/edgescompatibleupgrades",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &EdgeUpgradeListV2Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*EdgeUpgradeListV2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*EdgeUpgradeListV2Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ExecuteEdgeUpgradeV2 upgrades the edge software ntnx ignore

Upgrades the edge software.
*/
func (a *Client) ExecuteEdgeUpgradeV2(params *ExecuteEdgeUpgradeV2Params, authInfo runtime.ClientAuthInfoWriter) (*ExecuteEdgeUpgradeV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewExecuteEdgeUpgradeV2Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ExecuteEdgeUpgradeV2",
		Method:             "POST",
		PathPattern:        "/v1.0/edges/upgrade",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ExecuteEdgeUpgradeV2Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ExecuteEdgeUpgradeV2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ExecuteEdgeUpgradeV2Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
