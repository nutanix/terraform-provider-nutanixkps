// Code generated by go-swagger; DO NOT EDIT.

package edge_info

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new edge info API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for edge info API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
EdgeInfoGetV2 gets all edge service domain information by edge ID ntnx ignore

Once installed, the Karbon Platform Services Service Domain software provides the service domain infrastructure.

Retrieves all resource, build, and version details for a given service domain by ID.
The ID is the service domain serial number used when you added the service domain.

This request also requires an Authorization header which specifies your API key.
This example Python request is for a service domain with serial number 5f31b963-acac-4368-a157-0c7dc0d32000,
where bearer_api_key is your actual API key.

import http.client

conn = http.client.HTTPSConnection("karbon.nutanix.com")
headers = { 'authorization': "bearer_api_key" }
conn.request("GET", "//v1.0/edges/5f31b963-acac-4368-a157-0c7dc0d32000/info", headers=headers)
res = conn.getresponse()
data = res.read()
print(data.decode("utf-8"))
*/
func (a *Client) EdgeInfoGetV2(params *EdgeInfoGetV2Params, authInfo runtime.ClientAuthInfoWriter) (*EdgeInfoGetV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEdgeInfoGetV2Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "EdgeInfoGetV2",
		Method:             "GET",
		PathPattern:        "/v1.0/edges/{edgeId}/info",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &EdgeInfoGetV2Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*EdgeInfoGetV2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*EdgeInfoGetV2Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
EdgeInfoListV2 gets edge resource build and version details ntnx ignore

Retrieves all edge resource, build, and version details.
*/
func (a *Client) EdgeInfoListV2(params *EdgeInfoListV2Params, authInfo runtime.ClientAuthInfoWriter) (*EdgeInfoListV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEdgeInfoListV2Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "EdgeInfoListV2",
		Method:             "GET",
		PathPattern:        "/v1.0/edgesinfo",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &EdgeInfoListV2Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*EdgeInfoListV2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*EdgeInfoListV2Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
EdgeInfoUpdateV2 updates edge resource build and version details for a given edge ID ntnx ignore

Once installed, the Xi IoT edge software provides the service domain infrastructure.

Update resource, build, and version details for a given service domain VM by ID.
The ID is the service domain serial number used when you added the service domain.

This request also requires an Authorization header which specifies your API key.
It returns the updated service domain information in JSON format.
*/
func (a *Client) EdgeInfoUpdateV2(params *EdgeInfoUpdateV2Params, authInfo runtime.ClientAuthInfoWriter) (*EdgeInfoUpdateV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEdgeInfoUpdateV2Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "EdgeInfoUpdateV2",
		Method:             "PUT",
		PathPattern:        "/v1.0/edges/{id}/info",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &EdgeInfoUpdateV2Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*EdgeInfoUpdateV2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*EdgeInfoUpdateV2Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ProjectGetEdgesInfoV2 gets all edge information for a project by ID ntnx ignore

Retrieves all edge resource, build, and version details by project ID.
*/
func (a *Client) ProjectGetEdgesInfoV2(params *ProjectGetEdgesInfoV2Params, authInfo runtime.ClientAuthInfoWriter) (*ProjectGetEdgesInfoV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProjectGetEdgesInfoV2Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ProjectGetEdgesInfoV2",
		Method:             "GET",
		PathPattern:        "/v1.0/projects/{projectId}/edgesinfo",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ProjectGetEdgesInfoV2Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ProjectGetEdgesInfoV2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ProjectGetEdgesInfoV2Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
