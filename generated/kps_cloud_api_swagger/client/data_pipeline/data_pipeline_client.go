// Code generated by go-swagger; DO NOT EDIT.

package data_pipeline

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new data pipeline API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for data pipeline API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DataPipelineCreate creates a data pipeline

Create a data pipeline.
*/
func (a *Client) DataPipelineCreate(params *DataPipelineCreateParams, authInfo runtime.ClientAuthInfoWriter) (*DataPipelineCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDataPipelineCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DataPipelineCreate",
		Method:             "POST",
		PathPattern:        "/v1.0/datapipelines",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DataPipelineCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DataPipelineCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DataPipelineCreateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DataPipelineDelete deletes data pipeline by its ID

Delete the data pipeline with the given id {id}.
*/
func (a *Client) DataPipelineDelete(params *DataPipelineDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*DataPipelineDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDataPipelineDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DataPipelineDelete",
		Method:             "DELETE",
		PathPattern:        "/v1.0/datapipelines/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DataPipelineDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DataPipelineDeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DataPipelineDeleteDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DataPipelineGet lists data pipeline by its ID

Retrieves a data pipelines with a given ID {id}.
*/
func (a *Client) DataPipelineGet(params *DataPipelineGetParams, authInfo runtime.ClientAuthInfoWriter) (*DataPipelineGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDataPipelineGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DataPipelineGet",
		Method:             "GET",
		PathPattern:        "/v1.0/datapipelines/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DataPipelineGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DataPipelineGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DataPipelineGetDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DataPipelineList gets data pipelines

Retrieves all data pipelines for a tenant.
*/
func (a *Client) DataPipelineList(params *DataPipelineListParams, authInfo runtime.ClientAuthInfoWriter) (*DataPipelineListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDataPipelineListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DataPipelineList",
		Method:             "GET",
		PathPattern:        "/v1.0/datapipelines",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DataPipelineListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DataPipelineListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DataPipelineListDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DataPipelineUpdate updates a data pipeline by its ID

Update a data pipeline with a given ID {id}.
*/
func (a *Client) DataPipelineUpdate(params *DataPipelineUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*DataPipelineUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDataPipelineUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DataPipelineUpdate",
		Method:             "PUT",
		PathPattern:        "/v1.0/datapipelines/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DataPipelineUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DataPipelineUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DataPipelineUpdateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetDataPipelineContainers gets containers of a data pipeline specified by datapipeline ID running on a specific edge

Gets the containers of a data pipeline with the given ID {id} running on edge with id {edgeId}.
*/
func (a *Client) GetDataPipelineContainers(params *GetDataPipelineContainersParams, authInfo runtime.ClientAuthInfoWriter) (*GetDataPipelineContainersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDataPipelineContainersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetDataPipelineContainers",
		Method:             "GET",
		PathPattern:        "/v1.0/datapipelines/{id}/containers/{edgeId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetDataPipelineContainersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDataPipelineContainersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetDataPipelineContainersDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ProjectGetDataPipelines gets data pipelines for a project

Retrieves all data pipelines for a project with a given ID {projectId}.
*/
func (a *Client) ProjectGetDataPipelines(params *ProjectGetDataPipelinesParams, authInfo runtime.ClientAuthInfoWriter) (*ProjectGetDataPipelinesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProjectGetDataPipelinesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ProjectGetDataPipelines",
		Method:             "GET",
		PathPattern:        "/v1.0/projects/{projectId}/datapipelines",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ProjectGetDataPipelinesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ProjectGetDataPipelinesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ProjectGetDataPipelinesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
