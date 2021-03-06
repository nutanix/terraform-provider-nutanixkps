// Code generated by go-swagger; DO NOT EDIT.

package k8s_dashboard

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new k8s dashboard API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for k8s dashboard API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
K8sDashboardAddViewonlyUsers adds kubernetes dashboard viewonly users to a service domain karbon cluster ntnx ignore

Add kubernetes dashboard viewonly users to a Service Domain by its ID {svcDomainId}.
*/
func (a *Client) K8sDashboardAddViewonlyUsers(params *K8sDashboardAddViewonlyUsersParams, authInfo runtime.ClientAuthInfoWriter) (*K8sDashboardAddViewonlyUsersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewK8sDashboardAddViewonlyUsersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "K8sDashboardAddViewonlyUsers",
		Method:             "POST",
		PathPattern:        "/v1.0/k8sdashboard/{svcDomainId}/viewonlyUsersAdd",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &K8sDashboardAddViewonlyUsersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*K8sDashboardAddViewonlyUsersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*K8sDashboardAddViewonlyUsersDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
K8sDashboardGetAdminKubeConfig gets admin kube config for kubectl ntnx ignore

Get Admin KubeConfig for kubectl. Caller must be infra admin.
svcDomainId is ID of service domain or kubernetes cluster
*/
func (a *Client) K8sDashboardGetAdminKubeConfig(params *K8sDashboardGetAdminKubeConfigParams, authInfo runtime.ClientAuthInfoWriter) (*K8sDashboardGetAdminKubeConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewK8sDashboardGetAdminKubeConfigParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "K8sDashboardGetAdminKubeConfig",
		Method:             "GET",
		PathPattern:        "/v1.0/k8sdashboard/{svcDomainId}/adminKubeConfig",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &K8sDashboardGetAdminKubeConfigReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*K8sDashboardGetAdminKubeConfigOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*K8sDashboardGetAdminKubeConfigDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
K8sDashboardGetAdminToken gets admin token for kubernetes dashboard ntnx ignore

Get Admin Token for Kubernetes Dashboard. Caller must be infra admin.
svcDomainId is ID of service domain or kubernetes cluster
*/
func (a *Client) K8sDashboardGetAdminToken(params *K8sDashboardGetAdminTokenParams, authInfo runtime.ClientAuthInfoWriter) (*K8sDashboardGetAdminTokenOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewK8sDashboardGetAdminTokenParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "K8sDashboardGetAdminToken",
		Method:             "GET",
		PathPattern:        "/v1.0/k8sdashboard/{svcDomainId}/adminToken",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &K8sDashboardGetAdminTokenReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*K8sDashboardGetAdminTokenOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*K8sDashboardGetAdminTokenDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
K8sDashboardGetUserKubeConfig gets user kube config for kubectl ntnx ignore

Get User KubeConfig for kubectl.
svcDomainId is ID of service domain or kubernetes cluster
*/
func (a *Client) K8sDashboardGetUserKubeConfig(params *K8sDashboardGetUserKubeConfigParams, authInfo runtime.ClientAuthInfoWriter) (*K8sDashboardGetUserKubeConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewK8sDashboardGetUserKubeConfigParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "K8sDashboardGetUserKubeConfig",
		Method:             "GET",
		PathPattern:        "/v1.0/k8sdashboard/{svcDomainId}/userKubeConfig",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &K8sDashboardGetUserKubeConfigReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*K8sDashboardGetUserKubeConfigOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*K8sDashboardGetUserKubeConfigDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
K8sDashboardGetUserToken gets user token for kubernetes dashboard ntnx ignore

Get User Token for Kubernetes Dashboard.
svcDomainId is ID of service domain or kubernetes cluster
*/
func (a *Client) K8sDashboardGetUserToken(params *K8sDashboardGetUserTokenParams, authInfo runtime.ClientAuthInfoWriter) (*K8sDashboardGetUserTokenOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewK8sDashboardGetUserTokenParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "K8sDashboardGetUserToken",
		Method:             "GET",
		PathPattern:        "/v1.0/k8sdashboard/{svcDomainId}/userToken",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &K8sDashboardGetUserTokenReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*K8sDashboardGetUserTokenOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*K8sDashboardGetUserTokenDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
K8sDashboardGetViewonlyKubeConfig gets viewonly kube config for kubectl ntnx ignore

Get Viewonly KubeConfig for kubectl. Caller must have viewonly access.
svcDomainId is ID of service domain or kubernetes cluster
*/
func (a *Client) K8sDashboardGetViewonlyKubeConfig(params *K8sDashboardGetViewonlyKubeConfigParams, authInfo runtime.ClientAuthInfoWriter) (*K8sDashboardGetViewonlyKubeConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewK8sDashboardGetViewonlyKubeConfigParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "K8sDashboardGetViewonlyKubeConfig",
		Method:             "GET",
		PathPattern:        "/v1.0/k8sdashboard/{svcDomainId}/viewonlyKubeConfig",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &K8sDashboardGetViewonlyKubeConfigReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*K8sDashboardGetViewonlyKubeConfigOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*K8sDashboardGetViewonlyKubeConfigDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
K8sDashboardGetViewonlyToken gets viewonly token for kubernetes dashboard ntnx ignore

Get Viewonly Token for Kubernetes Dashboard. Caller must have viewonly access.
svcDomainId is ID of service domain or kubernetes cluster
*/
func (a *Client) K8sDashboardGetViewonlyToken(params *K8sDashboardGetViewonlyTokenParams, authInfo runtime.ClientAuthInfoWriter) (*K8sDashboardGetViewonlyTokenOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewK8sDashboardGetViewonlyTokenParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "K8sDashboardGetViewonlyToken",
		Method:             "GET",
		PathPattern:        "/v1.0/k8sdashboard/{svcDomainId}/viewonlyToken",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &K8sDashboardGetViewonlyTokenReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*K8sDashboardGetViewonlyTokenOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*K8sDashboardGetViewonlyTokenDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
K8sDashboardGetViewonlyUsers gets all kubernetes dashboard viewonly users associated with a service domain karbon cluster ntnx ignore

Retrieves a list of all kubernetes dashboard viewonly users associated with a Service Domain by its ID {svcDomainId}.
*/
func (a *Client) K8sDashboardGetViewonlyUsers(params *K8sDashboardGetViewonlyUsersParams, authInfo runtime.ClientAuthInfoWriter) (*K8sDashboardGetViewonlyUsersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewK8sDashboardGetViewonlyUsersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "K8sDashboardGetViewonlyUsers",
		Method:             "GET",
		PathPattern:        "/v1.0/k8sdashboard/{svcDomainId}/viewonlyUsers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &K8sDashboardGetViewonlyUsersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*K8sDashboardGetViewonlyUsersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*K8sDashboardGetViewonlyUsersDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
K8sDashboardRemoveViewonlyUsers removes kubernetes dashboard viewonly users from a service domain karbon cluster ntnx ignore

Remove kubernetes dashboard viewonly users from a Service Domain by its ID {svcDomainId}.
*/
func (a *Client) K8sDashboardRemoveViewonlyUsers(params *K8sDashboardRemoveViewonlyUsersParams, authInfo runtime.ClientAuthInfoWriter) (*K8sDashboardRemoveViewonlyUsersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewK8sDashboardRemoveViewonlyUsersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "K8sDashboardRemoveViewonlyUsers",
		Method:             "POST",
		PathPattern:        "/v1.0/k8sdashboard/{svcDomainId}/viewonlyUsersRemove",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &K8sDashboardRemoveViewonlyUsersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*K8sDashboardRemoveViewonlyUsersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*K8sDashboardRemoveViewonlyUsersDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
