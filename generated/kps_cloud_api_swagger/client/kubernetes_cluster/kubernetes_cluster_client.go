// Code generated by go-swagger; DO NOT EDIT.

package kubernetes_cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new kubernetes cluster API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for kubernetes cluster API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
KubernetesClusterInstaller gets the kubernetes clusters helm installer


Gets the kubernetes cluster helm installer.
*/
func (a *Client) KubernetesClusterInstaller(params *KubernetesClusterInstallerParams, authInfo runtime.ClientAuthInfoWriter) (*KubernetesClusterInstallerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewKubernetesClusterInstallerParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "KubernetesClusterInstaller",
		Method:             "GET",
		PathPattern:        "/v1.0/kubernetescluster-installer",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &KubernetesClusterInstallerReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*KubernetesClusterInstallerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*KubernetesClusterInstallerDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
KubernetesClustersCreate creates a kubernetes cluster

Create a kubernetes cluster.
*/
func (a *Client) KubernetesClustersCreate(params *KubernetesClustersCreateParams, authInfo runtime.ClientAuthInfoWriter) (*KubernetesClustersCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewKubernetesClustersCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "KubernetesClustersCreate",
		Method:             "POST",
		PathPattern:        "/v1.0/kubernetesclusters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &KubernetesClustersCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*KubernetesClustersCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*KubernetesClustersCreateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
KubernetesClustersDelete deletes a kubernetes cluster as specified by its ID

Deletes a kubernetes cluster by its ID {id}.
*/
func (a *Client) KubernetesClustersDelete(params *KubernetesClustersDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*KubernetesClustersDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewKubernetesClustersDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "KubernetesClustersDelete",
		Method:             "DELETE",
		PathPattern:        "/v1.0/kubernetesclusters/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &KubernetesClustersDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*KubernetesClustersDeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*KubernetesClustersDeleteDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
KubernetesClustersGet gets single kubernetes cluster
*/
func (a *Client) KubernetesClustersGet(params *KubernetesClustersGetParams, authInfo runtime.ClientAuthInfoWriter) (*KubernetesClustersGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewKubernetesClustersGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "KubernetesClustersGet",
		Method:             "GET",
		PathPattern:        "/v1.0/kubernetesclusters/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &KubernetesClustersGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*KubernetesClustersGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*KubernetesClustersGetDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
KubernetesClustersHandle retrieves the certificate and private key for the kubernetes cluster by its given ID ntnx ignore

Retrieves the certificate and private key for the kubernetes cluster by its given ID.
*/
func (a *Client) KubernetesClustersHandle(params *KubernetesClustersHandleParams) (*KubernetesClustersHandleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewKubernetesClustersHandleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "KubernetesClustersHandle",
		Method:             "POST",
		PathPattern:        "/v1.0/kubernetesclusters/{id}/handle",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &KubernetesClustersHandleReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*KubernetesClustersHandleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*KubernetesClustersHandleDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
KubernetesClustersList gets all kubernetes clusters

Retrieves a list of all kubernetes clusters.
*/
func (a *Client) KubernetesClustersList(params *KubernetesClustersListParams, authInfo runtime.ClientAuthInfoWriter) (*KubernetesClustersListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewKubernetesClustersListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "KubernetesClustersList",
		Method:             "GET",
		PathPattern:        "/v1.0/kubernetesclusters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &KubernetesClustersListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*KubernetesClustersListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*KubernetesClustersListDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
KubernetesClustersUpdate updates a kubernetes cluster by its ID

Updates a kubernetes cluster by its ID {id}.
*/
func (a *Client) KubernetesClustersUpdate(params *KubernetesClustersUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*KubernetesClustersUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewKubernetesClustersUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "KubernetesClustersUpdate",
		Method:             "PUT",
		PathPattern:        "/v1.0/kubernetesclusters/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &KubernetesClustersUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*KubernetesClustersUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*KubernetesClustersUpdateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
