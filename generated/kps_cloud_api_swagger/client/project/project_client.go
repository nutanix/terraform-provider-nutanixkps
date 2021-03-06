// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new project API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for project API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
ProjectCreateV2 creates a project

Creates a project.
*/
func (a *Client) ProjectCreateV2(params *ProjectCreateV2Params, authInfo runtime.ClientAuthInfoWriter) (*ProjectCreateV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProjectCreateV2Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ProjectCreateV2",
		Method:             "POST",
		PathPattern:        "/v1.0/projects",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ProjectCreateV2Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ProjectCreateV2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ProjectCreateV2Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ProjectDeleteV2 deletes a project by ID

Deletes a project with the given ID {id}.
*/
func (a *Client) ProjectDeleteV2(params *ProjectDeleteV2Params, authInfo runtime.ClientAuthInfoWriter) (*ProjectDeleteV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProjectDeleteV2Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ProjectDeleteV2",
		Method:             "DELETE",
		PathPattern:        "/v1.0/projects/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ProjectDeleteV2Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ProjectDeleteV2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ProjectDeleteV2Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ProjectGetV2 gets project by its ID

Retrieves the project by its given ID {projectId}.
*/
func (a *Client) ProjectGetV2(params *ProjectGetV2Params, authInfo runtime.ClientAuthInfoWriter) (*ProjectGetV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProjectGetV2Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ProjectGetV2",
		Method:             "GET",
		PathPattern:        "/v1.0/projects/{projectId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ProjectGetV2Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ProjectGetV2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ProjectGetV2Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ProjectListV2 gets projects

Retrieves all projects.
*/
func (a *Client) ProjectListV2(params *ProjectListV2Params, authInfo runtime.ClientAuthInfoWriter) (*ProjectListV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProjectListV2Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ProjectListV2",
		Method:             "GET",
		PathPattern:        "/v1.0/projects",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ProjectListV2Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ProjectListV2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ProjectListV2Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ProjectUpdateV3 updates a project by its ID

Updates a project by its given ID {id}.
*/
func (a *Client) ProjectUpdateV3(params *ProjectUpdateV3Params, authInfo runtime.ClientAuthInfoWriter) (*ProjectUpdateV3OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProjectUpdateV3Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ProjectUpdateV3",
		Method:             "PUT",
		PathPattern:        "/v1.0/projects/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ProjectUpdateV3Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ProjectUpdateV3OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ProjectUpdateV3Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
