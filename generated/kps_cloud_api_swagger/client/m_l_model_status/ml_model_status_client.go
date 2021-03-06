// Code generated by go-swagger; DO NOT EDIT.

package m_l_model_status

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new m l model status API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for m l model status API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
MLModelStatusGet gets m l model status by model ID

Retrieve status for an ML model with the given ID {id}.
*/
func (a *Client) MLModelStatusGet(params *MLModelStatusGetParams, authInfo runtime.ClientAuthInfoWriter) (*MLModelStatusGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMLModelStatusGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "MLModelStatusGet",
		Method:             "GET",
		PathPattern:        "/v1.0/mlmodelstatuses/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &MLModelStatusGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*MLModelStatusGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*MLModelStatusGetDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
MLModelStatusList gets status for all m l models

Retrieves status for all ML models.
*/
func (a *Client) MLModelStatusList(params *MLModelStatusListParams, authInfo runtime.ClientAuthInfoWriter) (*MLModelStatusListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMLModelStatusListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "MLModelStatusList",
		Method:             "GET",
		PathPattern:        "/v1.0/mlmodelstatuses",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &MLModelStatusListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*MLModelStatusListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*MLModelStatusListDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
