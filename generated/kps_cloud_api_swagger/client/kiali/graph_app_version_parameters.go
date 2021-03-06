// Code generated by go-swagger; DO NOT EDIT.

package kiali

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGraphAppVersionParams creates a new GraphAppVersionParams object
// with the default values initialized.
func NewGraphAppVersionParams() *GraphAppVersionParams {
	var (
		appendersDefault          = string("run all appenders")
		durationDefault           = string("10m")
		graphTypeDefault          = string("workload")
		groupByDefault            = string("none")
		injectServiceNodesDefault = string("false")
		queryTimeDefault          = string("now")
	)
	return &GraphAppVersionParams{
		Appenders:          &appendersDefault,
		Duration:           &durationDefault,
		GraphType:          &graphTypeDefault,
		GroupBy:            &groupByDefault,
		InjectServiceNodes: &injectServiceNodesDefault,
		QueryTime:          &queryTimeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGraphAppVersionParamsWithTimeout creates a new GraphAppVersionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGraphAppVersionParamsWithTimeout(timeout time.Duration) *GraphAppVersionParams {
	var (
		appendersDefault          = string("run all appenders")
		durationDefault           = string("10m")
		graphTypeDefault          = string("workload")
		groupByDefault            = string("none")
		injectServiceNodesDefault = string("false")
		queryTimeDefault          = string("now")
	)
	return &GraphAppVersionParams{
		Appenders:          &appendersDefault,
		Duration:           &durationDefault,
		GraphType:          &graphTypeDefault,
		GroupBy:            &groupByDefault,
		InjectServiceNodes: &injectServiceNodesDefault,
		QueryTime:          &queryTimeDefault,

		timeout: timeout,
	}
}

// NewGraphAppVersionParamsWithContext creates a new GraphAppVersionParams object
// with the default values initialized, and the ability to set a context for a request
func NewGraphAppVersionParamsWithContext(ctx context.Context) *GraphAppVersionParams {
	var (
		appendersDefault          = string("run all appenders")
		durationDefault           = string("10m")
		graphTypeDefault          = string("workload")
		groupByDefault            = string("none")
		injectServiceNodesDefault = string("false")
		queryTimeDefault          = string("now")
	)
	return &GraphAppVersionParams{
		Appenders:          &appendersDefault,
		Duration:           &durationDefault,
		GraphType:          &graphTypeDefault,
		GroupBy:            &groupByDefault,
		InjectServiceNodes: &injectServiceNodesDefault,
		QueryTime:          &queryTimeDefault,

		Context: ctx,
	}
}

// NewGraphAppVersionParamsWithHTTPClient creates a new GraphAppVersionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGraphAppVersionParamsWithHTTPClient(client *http.Client) *GraphAppVersionParams {
	var (
		appendersDefault          = string("run all appenders")
		durationDefault           = string("10m")
		graphTypeDefault          = string("workload")
		groupByDefault            = string("none")
		injectServiceNodesDefault = string("false")
		queryTimeDefault          = string("now")
	)
	return &GraphAppVersionParams{
		Appenders:          &appendersDefault,
		Duration:           &durationDefault,
		GraphType:          &graphTypeDefault,
		GroupBy:            &groupByDefault,
		InjectServiceNodes: &injectServiceNodesDefault,
		QueryTime:          &queryTimeDefault,
		HTTPClient:         client,
	}
}

/*GraphAppVersionParams contains all the parameters to send to the API endpoint
for the graph app version operation typically these are written to a http.Request
*/
type GraphAppVersionParams struct {

	/*Authorization
	  Format: Bearer &lt;token>, with &lt;token> from login API response.

	*/
	Authorization string
	/*App
	  The app name (label value).

	*/
	App string
	/*Appenders
	  Comma-separated list of Appenders to run. Available appenders: [deadNode, istio, responseTime, securityPolicy, serviceEntry, sidecarsCheck, unusedNode].

	*/
	Appenders *string
	/*Duration
	  Query time-range duration (Golang string duration).

	*/
	Duration *string
	/*GraphType
	  Graph type. Available graph types: [app, service, versionedApp, workload].

	*/
	GraphType *string
	/*GroupBy
	  App box grouping characteristic. Available groupings: [app, none, version].

	*/
	GroupBy *string
	/*InjectServiceNodes
	  Flag for injecting the requested service node between source and destination nodes.

	*/
	InjectServiceNodes *string
	/*Namespace
	  The namespace name.

	*/
	Namespace string
	/*QueryTime
	  Unix time (seconds) for query such that time range is [queryTime-duration..queryTime]. Default is now.

	*/
	QueryTime *string
	/*ServiceDomain
	  ID of ServiceDomain to access.

	*/
	ServiceDomain string
	/*Version
	  The app version (label value).

	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the graph app version params
func (o *GraphAppVersionParams) WithTimeout(timeout time.Duration) *GraphAppVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the graph app version params
func (o *GraphAppVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the graph app version params
func (o *GraphAppVersionParams) WithContext(ctx context.Context) *GraphAppVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the graph app version params
func (o *GraphAppVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the graph app version params
func (o *GraphAppVersionParams) WithHTTPClient(client *http.Client) *GraphAppVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the graph app version params
func (o *GraphAppVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the graph app version params
func (o *GraphAppVersionParams) WithAuthorization(authorization string) *GraphAppVersionParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the graph app version params
func (o *GraphAppVersionParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithApp adds the app to the graph app version params
func (o *GraphAppVersionParams) WithApp(app string) *GraphAppVersionParams {
	o.SetApp(app)
	return o
}

// SetApp adds the app to the graph app version params
func (o *GraphAppVersionParams) SetApp(app string) {
	o.App = app
}

// WithAppenders adds the appenders to the graph app version params
func (o *GraphAppVersionParams) WithAppenders(appenders *string) *GraphAppVersionParams {
	o.SetAppenders(appenders)
	return o
}

// SetAppenders adds the appenders to the graph app version params
func (o *GraphAppVersionParams) SetAppenders(appenders *string) {
	o.Appenders = appenders
}

// WithDuration adds the duration to the graph app version params
func (o *GraphAppVersionParams) WithDuration(duration *string) *GraphAppVersionParams {
	o.SetDuration(duration)
	return o
}

// SetDuration adds the duration to the graph app version params
func (o *GraphAppVersionParams) SetDuration(duration *string) {
	o.Duration = duration
}

// WithGraphType adds the graphType to the graph app version params
func (o *GraphAppVersionParams) WithGraphType(graphType *string) *GraphAppVersionParams {
	o.SetGraphType(graphType)
	return o
}

// SetGraphType adds the graphType to the graph app version params
func (o *GraphAppVersionParams) SetGraphType(graphType *string) {
	o.GraphType = graphType
}

// WithGroupBy adds the groupBy to the graph app version params
func (o *GraphAppVersionParams) WithGroupBy(groupBy *string) *GraphAppVersionParams {
	o.SetGroupBy(groupBy)
	return o
}

// SetGroupBy adds the groupBy to the graph app version params
func (o *GraphAppVersionParams) SetGroupBy(groupBy *string) {
	o.GroupBy = groupBy
}

// WithInjectServiceNodes adds the injectServiceNodes to the graph app version params
func (o *GraphAppVersionParams) WithInjectServiceNodes(injectServiceNodes *string) *GraphAppVersionParams {
	o.SetInjectServiceNodes(injectServiceNodes)
	return o
}

// SetInjectServiceNodes adds the injectServiceNodes to the graph app version params
func (o *GraphAppVersionParams) SetInjectServiceNodes(injectServiceNodes *string) {
	o.InjectServiceNodes = injectServiceNodes
}

// WithNamespace adds the namespace to the graph app version params
func (o *GraphAppVersionParams) WithNamespace(namespace string) *GraphAppVersionParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the graph app version params
func (o *GraphAppVersionParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithQueryTime adds the queryTime to the graph app version params
func (o *GraphAppVersionParams) WithQueryTime(queryTime *string) *GraphAppVersionParams {
	o.SetQueryTime(queryTime)
	return o
}

// SetQueryTime adds the queryTime to the graph app version params
func (o *GraphAppVersionParams) SetQueryTime(queryTime *string) {
	o.QueryTime = queryTime
}

// WithServiceDomain adds the serviceDomain to the graph app version params
func (o *GraphAppVersionParams) WithServiceDomain(serviceDomain string) *GraphAppVersionParams {
	o.SetServiceDomain(serviceDomain)
	return o
}

// SetServiceDomain adds the serviceDomain to the graph app version params
func (o *GraphAppVersionParams) SetServiceDomain(serviceDomain string) {
	o.ServiceDomain = serviceDomain
}

// WithVersion adds the version to the graph app version params
func (o *GraphAppVersionParams) WithVersion(version string) *GraphAppVersionParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the graph app version params
func (o *GraphAppVersionParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *GraphAppVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// path param app
	if err := r.SetPathParam("app", o.App); err != nil {
		return err
	}

	if o.Appenders != nil {

		// query param appenders
		var qrAppenders string
		if o.Appenders != nil {
			qrAppenders = *o.Appenders
		}
		qAppenders := qrAppenders
		if qAppenders != "" {
			if err := r.SetQueryParam("appenders", qAppenders); err != nil {
				return err
			}
		}

	}

	if o.Duration != nil {

		// query param duration
		var qrDuration string
		if o.Duration != nil {
			qrDuration = *o.Duration
		}
		qDuration := qrDuration
		if qDuration != "" {
			if err := r.SetQueryParam("duration", qDuration); err != nil {
				return err
			}
		}

	}

	if o.GraphType != nil {

		// query param graphType
		var qrGraphType string
		if o.GraphType != nil {
			qrGraphType = *o.GraphType
		}
		qGraphType := qrGraphType
		if qGraphType != "" {
			if err := r.SetQueryParam("graphType", qGraphType); err != nil {
				return err
			}
		}

	}

	if o.GroupBy != nil {

		// query param groupBy
		var qrGroupBy string
		if o.GroupBy != nil {
			qrGroupBy = *o.GroupBy
		}
		qGroupBy := qrGroupBy
		if qGroupBy != "" {
			if err := r.SetQueryParam("groupBy", qGroupBy); err != nil {
				return err
			}
		}

	}

	if o.InjectServiceNodes != nil {

		// query param injectServiceNodes
		var qrInjectServiceNodes string
		if o.InjectServiceNodes != nil {
			qrInjectServiceNodes = *o.InjectServiceNodes
		}
		qInjectServiceNodes := qrInjectServiceNodes
		if qInjectServiceNodes != "" {
			if err := r.SetQueryParam("injectServiceNodes", qInjectServiceNodes); err != nil {
				return err
			}
		}

	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	if o.QueryTime != nil {

		// query param queryTime
		var qrQueryTime string
		if o.QueryTime != nil {
			qrQueryTime = *o.QueryTime
		}
		qQueryTime := qrQueryTime
		if qQueryTime != "" {
			if err := r.SetQueryParam("queryTime", qQueryTime); err != nil {
				return err
			}
		}

	}

	// query param serviceDomain
	qrServiceDomain := o.ServiceDomain
	qServiceDomain := qrServiceDomain
	if qServiceDomain != "" {
		if err := r.SetQueryParam("serviceDomain", qServiceDomain); err != nil {
			return err
		}
	}

	// path param version
	if err := r.SetPathParam("version", o.Version); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
