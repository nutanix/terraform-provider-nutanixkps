// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/client/application"
	"sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/client/application_status"
	"sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/client/auditlog"
	"sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/client/auth"
	"sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/client/category"
	"sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/client/certificate"
	"sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/client/cloud_profile"
	"sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/client/container_registry"
	"sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/client/data_driver_class"
	"sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/client/data_driver_config"
	"sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/client/data_driver_instance"
	"sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/client/data_pipeline"
	"sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/client/data_source"
	"sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/client/edge"
	"sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/client/edge_info"
	"sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/client/edge_inventory_delta"
	"sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/client/edge_upgrade"
	"sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/client/event"
	"sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/client/function"
	"sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/client/helm"
	"sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/client/http_service_proxy"
	"sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/client/infra_config"
	"sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/client/k8s_dashboard"
	"sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/client/kiali"
	"sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/client/kubernetes_cluster"
	"sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/client/log"
	"sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/client/log_collector"
	"sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/client/m_l_model"
	"sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/client/m_l_model_status"
	"sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/client/node"
	"sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/client/node_info"
	"sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/client/operations"
	"sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/client/project"
	"sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/client/project_service"
	"sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/client/proxy"
	"sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/client/runtime_environment"
	"sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/client/sensor"
	"sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/client/service"
	"sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/client/service_binding"
	"sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/client/service_class"
	"sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/client/service_domain"
	"sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/client/service_domain_info"
	"sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/client/service_instance"
	"sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/client/software_update"
	"sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/client/ssh"
	"sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/client/storage_profile"
	"sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/client/tenant"
	"sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/client/tenant_props"
	"sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/client/user"
	"sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/client/user_api_token"
	"sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/client/user_props"
	"sherlock-terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/client/user_public_key"
)

// Default sherlock HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http", "https"}

// NewHTTPClient creates a new sherlock HTTP client.
func NewHTTPClient(formats strfmt.Registry) *Sherlock {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new sherlock HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *Sherlock {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new sherlock client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Sherlock {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(Sherlock)
	cli.Transport = transport

	cli.Application = application.New(transport, formats)

	cli.ApplicationStatus = application_status.New(transport, formats)

	cli.Auditlog = auditlog.New(transport, formats)

	cli.Auth = auth.New(transport, formats)

	cli.Category = category.New(transport, formats)

	cli.Certificate = certificate.New(transport, formats)

	cli.CloudProfile = cloud_profile.New(transport, formats)

	cli.ContainerRegistry = container_registry.New(transport, formats)

	cli.DataDriverClass = data_driver_class.New(transport, formats)

	cli.DataDriverConfig = data_driver_config.New(transport, formats)

	cli.DataDriverInstance = data_driver_instance.New(transport, formats)

	cli.DataPipeline = data_pipeline.New(transport, formats)

	cli.DataSource = data_source.New(transport, formats)

	cli.Edge = edge.New(transport, formats)

	cli.EdgeInfo = edge_info.New(transport, formats)

	cli.EdgeInventoryDelta = edge_inventory_delta.New(transport, formats)

	cli.EdgeUpgrade = edge_upgrade.New(transport, formats)

	cli.Event = event.New(transport, formats)

	cli.Function = function.New(transport, formats)

	cli.Helm = helm.New(transport, formats)

	cli.HTTPServiceProxy = http_service_proxy.New(transport, formats)

	cli.InfraConfig = infra_config.New(transport, formats)

	cli.K8sDashboard = k8s_dashboard.New(transport, formats)

	cli.Kiali = kiali.New(transport, formats)

	cli.KubernetesCluster = kubernetes_cluster.New(transport, formats)

	cli.Log = log.New(transport, formats)

	cli.LogCollector = log_collector.New(transport, formats)

	cli.MlModel = m_l_model.New(transport, formats)

	cli.MlModelStatus = m_l_model_status.New(transport, formats)

	cli.Node = node.New(transport, formats)

	cli.NodeInfo = node_info.New(transport, formats)

	cli.Operations = operations.New(transport, formats)

	cli.Project = project.New(transport, formats)

	cli.ProjectService = project_service.New(transport, formats)

	cli.Proxy = proxy.New(transport, formats)

	cli.RuntimeEnvironment = runtime_environment.New(transport, formats)

	cli.Sensor = sensor.New(transport, formats)

	cli.Service = service.New(transport, formats)

	cli.ServiceBinding = service_binding.New(transport, formats)

	cli.ServiceClass = service_class.New(transport, formats)

	cli.ServiceDomain = service_domain.New(transport, formats)

	cli.ServiceDomainInfo = service_domain_info.New(transport, formats)

	cli.ServiceInstance = service_instance.New(transport, formats)

	cli.SoftwareUpdate = software_update.New(transport, formats)

	cli.SSH = ssh.New(transport, formats)

	cli.StorageProfile = storage_profile.New(transport, formats)

	cli.Tenant = tenant.New(transport, formats)

	cli.TenantProps = tenant_props.New(transport, formats)

	cli.User = user.New(transport, formats)

	cli.UserAPIToken = user_api_token.New(transport, formats)

	cli.UserProps = user_props.New(transport, formats)

	cli.UserPublicKey = user_public_key.New(transport, formats)

	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// Sherlock is a client for sherlock
type Sherlock struct {
	Application *application.Client

	ApplicationStatus *application_status.Client

	Auditlog *auditlog.Client

	Auth *auth.Client

	Category *category.Client

	Certificate *certificate.Client

	CloudProfile *cloud_profile.Client

	ContainerRegistry *container_registry.Client

	DataDriverClass *data_driver_class.Client

	DataDriverConfig *data_driver_config.Client

	DataDriverInstance *data_driver_instance.Client

	DataPipeline *data_pipeline.Client

	DataSource *data_source.Client

	Edge *edge.Client

	EdgeInfo *edge_info.Client

	EdgeInventoryDelta *edge_inventory_delta.Client

	EdgeUpgrade *edge_upgrade.Client

	Event *event.Client

	Function *function.Client

	Helm *helm.Client

	HTTPServiceProxy *http_service_proxy.Client

	InfraConfig *infra_config.Client

	K8sDashboard *k8s_dashboard.Client

	Kiali *kiali.Client

	KubernetesCluster *kubernetes_cluster.Client

	Log *log.Client

	LogCollector *log_collector.Client

	MlModel *m_l_model.Client

	MlModelStatus *m_l_model_status.Client

	Node *node.Client

	NodeInfo *node_info.Client

	Operations *operations.Client

	Project *project.Client

	ProjectService *project_service.Client

	Proxy *proxy.Client

	RuntimeEnvironment *runtime_environment.Client

	Sensor *sensor.Client

	Service *service.Client

	ServiceBinding *service_binding.Client

	ServiceClass *service_class.Client

	ServiceDomain *service_domain.Client

	ServiceDomainInfo *service_domain_info.Client

	ServiceInstance *service_instance.Client

	SoftwareUpdate *software_update.Client

	SSH *ssh.Client

	StorageProfile *storage_profile.Client

	Tenant *tenant.Client

	TenantProps *tenant_props.Client

	User *user.Client

	UserAPIToken *user_api_token.Client

	UserProps *user_props.Client

	UserPublicKey *user_public_key.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *Sherlock) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.Application.SetTransport(transport)

	c.ApplicationStatus.SetTransport(transport)

	c.Auditlog.SetTransport(transport)

	c.Auth.SetTransport(transport)

	c.Category.SetTransport(transport)

	c.Certificate.SetTransport(transport)

	c.CloudProfile.SetTransport(transport)

	c.ContainerRegistry.SetTransport(transport)

	c.DataDriverClass.SetTransport(transport)

	c.DataDriverConfig.SetTransport(transport)

	c.DataDriverInstance.SetTransport(transport)

	c.DataPipeline.SetTransport(transport)

	c.DataSource.SetTransport(transport)

	c.Edge.SetTransport(transport)

	c.EdgeInfo.SetTransport(transport)

	c.EdgeInventoryDelta.SetTransport(transport)

	c.EdgeUpgrade.SetTransport(transport)

	c.Event.SetTransport(transport)

	c.Function.SetTransport(transport)

	c.Helm.SetTransport(transport)

	c.HTTPServiceProxy.SetTransport(transport)

	c.InfraConfig.SetTransport(transport)

	c.K8sDashboard.SetTransport(transport)

	c.Kiali.SetTransport(transport)

	c.KubernetesCluster.SetTransport(transport)

	c.Log.SetTransport(transport)

	c.LogCollector.SetTransport(transport)

	c.MlModel.SetTransport(transport)

	c.MlModelStatus.SetTransport(transport)

	c.Node.SetTransport(transport)

	c.NodeInfo.SetTransport(transport)

	c.Operations.SetTransport(transport)

	c.Project.SetTransport(transport)

	c.ProjectService.SetTransport(transport)

	c.Proxy.SetTransport(transport)

	c.RuntimeEnvironment.SetTransport(transport)

	c.Sensor.SetTransport(transport)

	c.Service.SetTransport(transport)

	c.ServiceBinding.SetTransport(transport)

	c.ServiceClass.SetTransport(transport)

	c.ServiceDomain.SetTransport(transport)

	c.ServiceDomainInfo.SetTransport(transport)

	c.ServiceInstance.SetTransport(transport)

	c.SoftwareUpdate.SetTransport(transport)

	c.SSH.SetTransport(transport)

	c.StorageProfile.SetTransport(transport)

	c.Tenant.SetTransport(transport)

	c.TenantProps.SetTransport(transport)

	c.User.SetTransport(transport)

	c.UserAPIToken.SetTransport(transport)

	c.UserProps.SetTransport(transport)

	c.UserPublicKey.SetTransport(transport)

}
