// Code generated by go-swagger; DO NOT EDIT.

package application

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new application API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for application API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
ApplicationCreateV2 creates an application

Create an application.
*/
func (a *Client) ApplicationCreateV2(params *ApplicationCreateV2Params, authInfo runtime.ClientAuthInfoWriter) (*ApplicationCreateV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewApplicationCreateV2Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ApplicationCreateV2",
		Method:             "POST",
		PathPattern:        "/v1.0/applications",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ApplicationCreateV2Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ApplicationCreateV2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ApplicationCreateV2Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ApplicationDeleteV2 deletes application specified by the application ID

Deletes the application with the given ID {id}.
*/
func (a *Client) ApplicationDeleteV2(params *ApplicationDeleteV2Params, authInfo runtime.ClientAuthInfoWriter) (*ApplicationDeleteV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewApplicationDeleteV2Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ApplicationDeleteV2",
		Method:             "DELETE",
		PathPattern:        "/v1.0/applications/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ApplicationDeleteV2Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ApplicationDeleteV2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ApplicationDeleteV2Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ApplicationGetV2 gets application by application ID

Retrieves the application with the given ID {id}.
*/
func (a *Client) ApplicationGetV2(params *ApplicationGetV2Params, authInfo runtime.ClientAuthInfoWriter) (*ApplicationGetV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewApplicationGetV2Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ApplicationGetV2",
		Method:             "GET",
		PathPattern:        "/v1.0/applications/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ApplicationGetV2Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ApplicationGetV2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ApplicationGetV2Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ApplicationListV2 gets all applications

Retrieves a list of all applications.
*/
func (a *Client) ApplicationListV2(params *ApplicationListV2Params, authInfo runtime.ClientAuthInfoWriter) (*ApplicationListV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewApplicationListV2Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ApplicationListV2",
		Method:             "GET",
		PathPattern:        "/v1.0/applications",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ApplicationListV2Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ApplicationListV2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ApplicationListV2Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ApplicationUpdateV3 updates a specific application with ID id

Update a specific application with ID {id}.
You cannot change the project associated with the application or the application ID.
You can change all other attributes.
*/
func (a *Client) ApplicationUpdateV3(params *ApplicationUpdateV3Params, authInfo runtime.ClientAuthInfoWriter) (*ApplicationUpdateV3OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewApplicationUpdateV3Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ApplicationUpdateV3",
		Method:             "PUT",
		PathPattern:        "/v1.0/applications/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ApplicationUpdateV3Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ApplicationUpdateV3OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ApplicationUpdateV3Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetApplicationContainers gets containers of an application specified by application ID running on a specific edge

Gets the containers of an application with the given ID {id} running on edge with id {edgeId}.
*/
func (a *Client) GetApplicationContainers(params *GetApplicationContainersParams, authInfo runtime.ClientAuthInfoWriter) (*GetApplicationContainersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetApplicationContainersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetApplicationContainers",
		Method:             "GET",
		PathPattern:        "/v1.0/applications/{id}/containers/{edgeId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetApplicationContainersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetApplicationContainersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetApplicationContainersDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
HelmAppCreate creates a new helm chart based app ntnx ignore

Create a new helm chart based app and return the charts uuid.
*/
func (a *Client) HelmAppCreate(params *HelmAppCreateParams, authInfo runtime.ClientAuthInfoWriter) (*HelmAppCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHelmAppCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "HelmAppCreate",
		Method:             "POST",
		PathPattern:        "/v1.0/helmapp",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &HelmAppCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*HelmAppCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*HelmAppCreateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
HelmAppGetYaml gets application by application ID ntnx ignore

Retrieves the application with the given ID {id}.
*/
func (a *Client) HelmAppGetYaml(params *HelmAppGetYamlParams, authInfo runtime.ClientAuthInfoWriter) (*HelmAppGetYamlOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHelmAppGetYamlParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "HelmAppGetYaml",
		Method:             "GET",
		PathPattern:        "/v1.0/helmapp/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &HelmAppGetYamlReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*HelmAppGetYamlOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*HelmAppGetYamlDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
HelmValuesCreate adds a values file to the helm chart identified by id ntnx ignore

Adds a values file to the helm chart identified by id.
*/
func (a *Client) HelmValuesCreate(params *HelmValuesCreateParams, authInfo runtime.ClientAuthInfoWriter) (*HelmValuesCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewHelmValuesCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "HelmValuesCreate",
		Method:             "POST",
		PathPattern:        "/v1.0/helmapp/{id}/values",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &HelmValuesCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*HelmValuesCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*HelmValuesCreateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ProjectGetApplicationsV2 gets all applications in a project according to project ID

Retrieves a list of all applications in a project with ID {projectId}.
*/
func (a *Client) ProjectGetApplicationsV2(params *ProjectGetApplicationsV2Params, authInfo runtime.ClientAuthInfoWriter) (*ProjectGetApplicationsV2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProjectGetApplicationsV2Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ProjectGetApplicationsV2",
		Method:             "GET",
		PathPattern:        "/v1.0/projects/{projectId}/applications",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ProjectGetApplicationsV2Reader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ProjectGetApplicationsV2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ProjectGetApplicationsV2Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
RenderApplication renders application ID running on a specific edge ntnx ignore

Render application template with the given ID {id} running on edge with id {edgeId}.
*/
func (a *Client) RenderApplication(params *RenderApplicationParams, authInfo runtime.ClientAuthInfoWriter) (*RenderApplicationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRenderApplicationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RenderApplication",
		Method:             "POST",
		PathPattern:        "/v1/applications/{id}/render/{edgeId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RenderApplicationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RenderApplicationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*RenderApplicationDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
