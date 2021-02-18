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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewWorkloadDashboardParams creates a new WorkloadDashboardParams object
// with the default values initialized.
func NewWorkloadDashboardParams() *WorkloadDashboardParams {
	var (
		avgDefault             = bool(true)
		directionDefault       = string("outbound")
		durationDefault        = int64(1800)
		rateFuncDefault        = string("rate")
		rateIntervalDefault    = string("1m")
		reporterDefault        = string("source")
		requestProtocolDefault = string("all protocols")
		stepDefault            = int64(15)
	)
	return &WorkloadDashboardParams{
		Avg:             &avgDefault,
		Direction:       &directionDefault,
		Duration:        &durationDefault,
		RateFunc:        &rateFuncDefault,
		RateInterval:    &rateIntervalDefault,
		Reporter:        &reporterDefault,
		RequestProtocol: &requestProtocolDefault,
		Step:            &stepDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewWorkloadDashboardParamsWithTimeout creates a new WorkloadDashboardParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWorkloadDashboardParamsWithTimeout(timeout time.Duration) *WorkloadDashboardParams {
	var (
		avgDefault             = bool(true)
		directionDefault       = string("outbound")
		durationDefault        = int64(1800)
		rateFuncDefault        = string("rate")
		rateIntervalDefault    = string("1m")
		reporterDefault        = string("source")
		requestProtocolDefault = string("all protocols")
		stepDefault            = int64(15)
	)
	return &WorkloadDashboardParams{
		Avg:             &avgDefault,
		Direction:       &directionDefault,
		Duration:        &durationDefault,
		RateFunc:        &rateFuncDefault,
		RateInterval:    &rateIntervalDefault,
		Reporter:        &reporterDefault,
		RequestProtocol: &requestProtocolDefault,
		Step:            &stepDefault,

		timeout: timeout,
	}
}

// NewWorkloadDashboardParamsWithContext creates a new WorkloadDashboardParams object
// with the default values initialized, and the ability to set a context for a request
func NewWorkloadDashboardParamsWithContext(ctx context.Context) *WorkloadDashboardParams {
	var (
		avgDefault             = bool(true)
		directionDefault       = string("outbound")
		durationDefault        = int64(1800)
		rateFuncDefault        = string("rate")
		rateIntervalDefault    = string("1m")
		reporterDefault        = string("source")
		requestProtocolDefault = string("all protocols")
		stepDefault            = int64(15)
	)
	return &WorkloadDashboardParams{
		Avg:             &avgDefault,
		Direction:       &directionDefault,
		Duration:        &durationDefault,
		RateFunc:        &rateFuncDefault,
		RateInterval:    &rateIntervalDefault,
		Reporter:        &reporterDefault,
		RequestProtocol: &requestProtocolDefault,
		Step:            &stepDefault,

		Context: ctx,
	}
}

// NewWorkloadDashboardParamsWithHTTPClient creates a new WorkloadDashboardParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWorkloadDashboardParamsWithHTTPClient(client *http.Client) *WorkloadDashboardParams {
	var (
		avgDefault             = bool(true)
		directionDefault       = string("outbound")
		durationDefault        = int64(1800)
		rateFuncDefault        = string("rate")
		rateIntervalDefault    = string("1m")
		reporterDefault        = string("source")
		requestProtocolDefault = string("all protocols")
		stepDefault            = int64(15)
	)
	return &WorkloadDashboardParams{
		Avg:             &avgDefault,
		Direction:       &directionDefault,
		Duration:        &durationDefault,
		RateFunc:        &rateFuncDefault,
		RateInterval:    &rateIntervalDefault,
		Reporter:        &reporterDefault,
		RequestProtocol: &requestProtocolDefault,
		Step:            &stepDefault,
		HTTPClient:      client,
	}
}

/*WorkloadDashboardParams contains all the parameters to send to the API endpoint
for the workload dashboard operation typically these are written to a http.Request
*/
type WorkloadDashboardParams struct {

	/*Authorization
	  Format: Bearer &lt;token>, with &lt;token> from login API response.

	*/
	Authorization string
	/*Avg
	  Flag for fetching histogram average. Default is true.

	*/
	Avg *bool
	/*ByLabels
	  List of labels to use for grouping metrics (via Prometheus 'by' clause).

	*/
	ByLabels []string
	/*Direction
	  Traffic direction: 'inbound' or 'outbound'.

	*/
	Direction *string
	/*Duration
	  Duration of the query period, in seconds.

	*/
	Duration *int64
	/*Namespace
	  The namespace name.

	*/
	Namespace string
	/*Quantiles
	  List of quantiles to fetch. Fetch no quantiles when empty. Ex: [0.5, 0.95, 0.99].

	*/
	Quantiles []string
	/*RateFunc
	  Prometheus function used to calculate rate: 'rate' or 'irate'.

	*/
	RateFunc *string
	/*RateInterval
	  Interval used for rate and histogram calculation.

	*/
	RateInterval *string
	/*Reporter
	  Istio telemetry reporter: 'source' or 'destination'.

	*/
	Reporter *string
	/*RequestProtocol
	  Desired request protocol for the telemetry: For example, 'http' or 'grpc'.

	*/
	RequestProtocol *string
	/*ServiceDomain
	  ID of ServiceDomain to access.

	*/
	ServiceDomain string
	/*Step
	  Step between [graph] datapoints, in seconds.

	*/
	Step *int64
	/*Workload
	  The workload name.

	*/
	Workload string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the workload dashboard params
func (o *WorkloadDashboardParams) WithTimeout(timeout time.Duration) *WorkloadDashboardParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the workload dashboard params
func (o *WorkloadDashboardParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the workload dashboard params
func (o *WorkloadDashboardParams) WithContext(ctx context.Context) *WorkloadDashboardParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the workload dashboard params
func (o *WorkloadDashboardParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the workload dashboard params
func (o *WorkloadDashboardParams) WithHTTPClient(client *http.Client) *WorkloadDashboardParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the workload dashboard params
func (o *WorkloadDashboardParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the workload dashboard params
func (o *WorkloadDashboardParams) WithAuthorization(authorization string) *WorkloadDashboardParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the workload dashboard params
func (o *WorkloadDashboardParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithAvg adds the avg to the workload dashboard params
func (o *WorkloadDashboardParams) WithAvg(avg *bool) *WorkloadDashboardParams {
	o.SetAvg(avg)
	return o
}

// SetAvg adds the avg to the workload dashboard params
func (o *WorkloadDashboardParams) SetAvg(avg *bool) {
	o.Avg = avg
}

// WithByLabels adds the byLabels to the workload dashboard params
func (o *WorkloadDashboardParams) WithByLabels(byLabels []string) *WorkloadDashboardParams {
	o.SetByLabels(byLabels)
	return o
}

// SetByLabels adds the byLabels to the workload dashboard params
func (o *WorkloadDashboardParams) SetByLabels(byLabels []string) {
	o.ByLabels = byLabels
}

// WithDirection adds the direction to the workload dashboard params
func (o *WorkloadDashboardParams) WithDirection(direction *string) *WorkloadDashboardParams {
	o.SetDirection(direction)
	return o
}

// SetDirection adds the direction to the workload dashboard params
func (o *WorkloadDashboardParams) SetDirection(direction *string) {
	o.Direction = direction
}

// WithDuration adds the duration to the workload dashboard params
func (o *WorkloadDashboardParams) WithDuration(duration *int64) *WorkloadDashboardParams {
	o.SetDuration(duration)
	return o
}

// SetDuration adds the duration to the workload dashboard params
func (o *WorkloadDashboardParams) SetDuration(duration *int64) {
	o.Duration = duration
}

// WithNamespace adds the namespace to the workload dashboard params
func (o *WorkloadDashboardParams) WithNamespace(namespace string) *WorkloadDashboardParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the workload dashboard params
func (o *WorkloadDashboardParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithQuantiles adds the quantiles to the workload dashboard params
func (o *WorkloadDashboardParams) WithQuantiles(quantiles []string) *WorkloadDashboardParams {
	o.SetQuantiles(quantiles)
	return o
}

// SetQuantiles adds the quantiles to the workload dashboard params
func (o *WorkloadDashboardParams) SetQuantiles(quantiles []string) {
	o.Quantiles = quantiles
}

// WithRateFunc adds the rateFunc to the workload dashboard params
func (o *WorkloadDashboardParams) WithRateFunc(rateFunc *string) *WorkloadDashboardParams {
	o.SetRateFunc(rateFunc)
	return o
}

// SetRateFunc adds the rateFunc to the workload dashboard params
func (o *WorkloadDashboardParams) SetRateFunc(rateFunc *string) {
	o.RateFunc = rateFunc
}

// WithRateInterval adds the rateInterval to the workload dashboard params
func (o *WorkloadDashboardParams) WithRateInterval(rateInterval *string) *WorkloadDashboardParams {
	o.SetRateInterval(rateInterval)
	return o
}

// SetRateInterval adds the rateInterval to the workload dashboard params
func (o *WorkloadDashboardParams) SetRateInterval(rateInterval *string) {
	o.RateInterval = rateInterval
}

// WithReporter adds the reporter to the workload dashboard params
func (o *WorkloadDashboardParams) WithReporter(reporter *string) *WorkloadDashboardParams {
	o.SetReporter(reporter)
	return o
}

// SetReporter adds the reporter to the workload dashboard params
func (o *WorkloadDashboardParams) SetReporter(reporter *string) {
	o.Reporter = reporter
}

// WithRequestProtocol adds the requestProtocol to the workload dashboard params
func (o *WorkloadDashboardParams) WithRequestProtocol(requestProtocol *string) *WorkloadDashboardParams {
	o.SetRequestProtocol(requestProtocol)
	return o
}

// SetRequestProtocol adds the requestProtocol to the workload dashboard params
func (o *WorkloadDashboardParams) SetRequestProtocol(requestProtocol *string) {
	o.RequestProtocol = requestProtocol
}

// WithServiceDomain adds the serviceDomain to the workload dashboard params
func (o *WorkloadDashboardParams) WithServiceDomain(serviceDomain string) *WorkloadDashboardParams {
	o.SetServiceDomain(serviceDomain)
	return o
}

// SetServiceDomain adds the serviceDomain to the workload dashboard params
func (o *WorkloadDashboardParams) SetServiceDomain(serviceDomain string) {
	o.ServiceDomain = serviceDomain
}

// WithStep adds the step to the workload dashboard params
func (o *WorkloadDashboardParams) WithStep(step *int64) *WorkloadDashboardParams {
	o.SetStep(step)
	return o
}

// SetStep adds the step to the workload dashboard params
func (o *WorkloadDashboardParams) SetStep(step *int64) {
	o.Step = step
}

// WithWorkload adds the workload to the workload dashboard params
func (o *WorkloadDashboardParams) WithWorkload(workload string) *WorkloadDashboardParams {
	o.SetWorkload(workload)
	return o
}

// SetWorkload adds the workload to the workload dashboard params
func (o *WorkloadDashboardParams) SetWorkload(workload string) {
	o.Workload = workload
}

// WriteToRequest writes these params to a swagger request
func (o *WorkloadDashboardParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	if o.Avg != nil {

		// query param avg
		var qrAvg bool
		if o.Avg != nil {
			qrAvg = *o.Avg
		}
		qAvg := swag.FormatBool(qrAvg)
		if qAvg != "" {
			if err := r.SetQueryParam("avg", qAvg); err != nil {
				return err
			}
		}

	}

	valuesByLabels := o.ByLabels

	joinedByLabels := swag.JoinByFormat(valuesByLabels, "")
	// query array param byLabels[]
	if err := r.SetQueryParam("byLabels[]", joinedByLabels...); err != nil {
		return err
	}

	if o.Direction != nil {

		// query param direction
		var qrDirection string
		if o.Direction != nil {
			qrDirection = *o.Direction
		}
		qDirection := qrDirection
		if qDirection != "" {
			if err := r.SetQueryParam("direction", qDirection); err != nil {
				return err
			}
		}

	}

	if o.Duration != nil {

		// query param duration
		var qrDuration int64
		if o.Duration != nil {
			qrDuration = *o.Duration
		}
		qDuration := swag.FormatInt64(qrDuration)
		if qDuration != "" {
			if err := r.SetQueryParam("duration", qDuration); err != nil {
				return err
			}
		}

	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	valuesQuantiles := o.Quantiles

	joinedQuantiles := swag.JoinByFormat(valuesQuantiles, "")
	// query array param quantiles[]
	if err := r.SetQueryParam("quantiles[]", joinedQuantiles...); err != nil {
		return err
	}

	if o.RateFunc != nil {

		// query param rateFunc
		var qrRateFunc string
		if o.RateFunc != nil {
			qrRateFunc = *o.RateFunc
		}
		qRateFunc := qrRateFunc
		if qRateFunc != "" {
			if err := r.SetQueryParam("rateFunc", qRateFunc); err != nil {
				return err
			}
		}

	}

	if o.RateInterval != nil {

		// query param rateInterval
		var qrRateInterval string
		if o.RateInterval != nil {
			qrRateInterval = *o.RateInterval
		}
		qRateInterval := qrRateInterval
		if qRateInterval != "" {
			if err := r.SetQueryParam("rateInterval", qRateInterval); err != nil {
				return err
			}
		}

	}

	if o.Reporter != nil {

		// query param reporter
		var qrReporter string
		if o.Reporter != nil {
			qrReporter = *o.Reporter
		}
		qReporter := qrReporter
		if qReporter != "" {
			if err := r.SetQueryParam("reporter", qReporter); err != nil {
				return err
			}
		}

	}

	if o.RequestProtocol != nil {

		// query param requestProtocol
		var qrRequestProtocol string
		if o.RequestProtocol != nil {
			qrRequestProtocol = *o.RequestProtocol
		}
		qRequestProtocol := qrRequestProtocol
		if qRequestProtocol != "" {
			if err := r.SetQueryParam("requestProtocol", qRequestProtocol); err != nil {
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

	if o.Step != nil {

		// query param step
		var qrStep int64
		if o.Step != nil {
			qrStep = *o.Step
		}
		qStep := swag.FormatInt64(qrStep)
		if qStep != "" {
			if err := r.SetQueryParam("step", qStep); err != nil {
				return err
			}
		}

	}

	// path param workload
	if err := r.SetPathParam("workload", o.Workload); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
