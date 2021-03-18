package nutanixkps

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"sherlock-terraform-provider-nutanixkps/nutanixkpsclient"
)

// Provider -
func Provider() *schema.Provider {

	descriptions := map[string]string{
		"host": "Karbon Platform Services endpoint for hosting the Service Domain. ",
		"username": "Username for the Karbon Platform Services local user account user.",
		"password": "Password for the Karbon Platform Services local user account user.",
		"api_token": "API authentication token used during API requests. You can generate API keys in the KPS cloud management console by using the Manage API Keys option.",
	}
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"host": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("NUTANIX_KPS_HOST", nil),
				Description: descriptions["host"],
			},
			"username": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("NUTANIX_KPS_USERNAME", nil),
				Description: descriptions["username"],
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("NUTANIX_KPS_PASSWORD", nil),
				Description: descriptions["password"],
			},
			"api_token": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("NUTANIX_KPS_API_TOKEN", nil),
				Description: descriptions["api_token"],
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"nutanixkps_servicedomain":  resourceServiceDomain(),
			"nutanixkps_node":           resourceNode(),
			"nutanixkps_storageprofile": resourceStorageProfile(),
			"nutanixkps_application":     resourceApplication(),
			"nutanixkps_project":         resourceProject(),
			"nutanixkps_serviceinstance": resourceServiceInstance(),
			"nutanixkps_servicebinding":  resourceServiceBinding(),
			"nutanixkps_vm_cloud_config": resourceVMCloudConfig(),
      
		},
		DataSourcesMap: map[string]*schema.Resource{
			"nutanixkps_servicedomains": dataSourceServiceDomains(),
			"nutanixkps_node":           dataSourceNode(),
			"nutanixkps_nodes":          dataSourceNodes(),
			"nutanixkps_serviceclass":   dataSourceServiceClass(),
			"nutanixkps_serviceclasses": dataSourceServiceClasses(),
			"nutanixkps_user":           dataSourceUser(),
			"nutanixkps_users":          dataSourceUsers(),
			"nutanixkps_vm_config": dataSourceVMConfig(), 
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	username := d.Get("username").(string)
	password := d.Get("password").(string)
	apiToken := d.Get("api_token").(string)

	var host *string

	hVal, ok := d.GetOk("host")
	if ok {
		tempHost := hVal.(string)
		host = &tempHost
	}

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	// Check if both the API token and username/password is passed
	if ((username != "") || (password != "")) && (apiToken != "") {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Cannot use api_token together with username and password",
		})

		return nil, diags
	}

	if ((username != "") && (password != "")) || (apiToken != "") {
		c, err := nutanixkpsclient.NewClient(host, &apiToken, &username, &password)
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to create KPS nutanixkpsclient",
				Detail:   err.Error(),
			})

			return nil, diags
		}

		return c, diags
	}

	c, err := nutanixkpsclient.NewClient(host, nil, nil, nil)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create KPS nutanixkpsclient",
			Detail:   err.Error(),
		})
		return nil, diags
	}

	return c, diags
}
