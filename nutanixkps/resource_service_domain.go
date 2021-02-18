package nutanixkps

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"terraform-provider-nutanixkps/nutanixkpsclient"

	"terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"
)

func resourceServiceDomain() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceServiceDomainCreate,
		ReadContext:   resourceServiceDomainRead,
		UpdateContext: resourceServiceDomainUpdate,
		DeleteContext: resourceServiceDomainDelete,
		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Description: "Last updated timestamp, auto-computed. No input required for this field.",
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Description: `Name of the service domain: 
				Name must include lowercase alphanumeric characters and must start and end with an lowercase alphanumeric character.
				Dash (-) and dot (.) characters are allowed as delimiters. Maximum length of 60 characters.`,
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Description: "Describe the service domain. For example, the main purpose or use case of the service domain.",
				Type:     schema.TypeString,
				Required: true,
			},
			"virtual_ip": &schema.Schema{
				Description: "Virtual IPv4 address for this Service Domain",
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.IsIPAddress,
				ForceNew:     true,
			},
		},
	}
}

func resourceServiceDomainCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*nutanixkpsclient.Client)

	serviceDomainName := d.Get("name").(string)
	serviceDomainDescription := d.Get("description").(string)
	serviceDomainVirtualIP := d.Get("virtual_ip").(string)
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	var serviceDomain models.ServiceDomain
	serviceDomain.Name = &serviceDomainName
	serviceDomain.Description = serviceDomainDescription
	serviceDomain.VirtualIP = serviceDomainVirtualIP

	serviceDomainID, err := client.ServiceDomainCreate(&serviceDomain)
	if err != nil {
		return diag.FromErr(nutanixkpsclient.APIErrorToError(err))
	}

	d.SetId(*serviceDomainID)
	resourceServiceDomainRead(ctx, d, m)

	return diags
}

func resourceServiceDomainUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*nutanixkpsclient.Client)

	serviceDomainID := d.Id()

	if d.HasChange("name") || d.HasChange("description") {
		serviceDomainName := d.Get("name").(string)
		serviceDomainDescription := d.Get("description").(string)

		var serviceDomain models.ServiceDomain
		serviceDomain.Name = &serviceDomainName
		serviceDomain.Description = serviceDomainDescription
		_, err := client.ServiceDomainUpdate(serviceDomainID, &serviceDomain)
		if err != nil {
			return diag.FromErr(nutanixkpsclient.APIErrorToError(err))
		}

		d.Set("last_updated", time.Now().Format(time.RFC850))
	}

	return resourceServiceDomainRead(ctx, d, m)
}

func resourceServiceDomainRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*nutanixkpsclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	ServiceDomainID := d.Id()

	serviceDomain, err := client.ServiceDomainGet(ServiceDomainID)
	if err != nil {
		return diag.FromErr(nutanixkpsclient.APIErrorToError(err))
	}

	if err := d.Set("name", serviceDomain.Name); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("description", serviceDomain.Description); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("virtual_ip", serviceDomain.VirtualIP); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceServiceDomainDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*nutanixkpsclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	ServiceDomainID := d.Id()

	_, err := client.ServiceDomainDelete(ServiceDomainID)
	if err != nil {
		return diag.FromErr(nutanixkpsclient.APIErrorToError(err))
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
