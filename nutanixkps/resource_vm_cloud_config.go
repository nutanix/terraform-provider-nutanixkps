package nutanixkps

import (
	"context"
	"net/http"
	"time"
	"io/ioutil"
	"strconv"
	"strings"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)
func resourceVMCloudConfig() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceVMCloudConfigCreate,
		ReadContext:   resourceVMCloudConfigRead,
		UpdateContext: resourceVMCloudConfigUpdate,
		DeleteContext: resourceVMCloudConfigDelete,
		Schema: map[string]*schema.Schema{
			"public_ip_address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.IsIPAddress,
			},
			"cloud_fqdn": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceVMCloudConfigCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Print("[Debug] Entering resourceVMCloudConfigCreate")
	publicIP, iok := d.GetOk("public_ip_address")
	if !iok {
		return diag.Errorf("Please provide public IP address of the virtual machine")
	}
	cloudFqdn, cok := d.GetOk("cloud_fqdn")
	if !cok {
		return diag.Errorf("Please provide cloud FQDN where this VM is / will be onboarded")
	}
	resp, _ := configureCloudFqdn(publicIP.(string), cloudFqdn.(string))
	if resp != 0 {
		return diag.Errorf("Failed to configure cloud FQDN on the virtual machine")
	}
	d.SetId(publicIP.(string))
	return resourceVMCloudConfigRead(ctx, d, m)
	
}

func resourceVMCloudConfigUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//todo

	return resourceVMCloudConfigRead(ctx, d, m)
}

func resourceVMCloudConfigDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// todo implement delete
	d.SetId("")

	return nil
}

func resourceVMCloudConfigRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	_, iok := d.GetOk("public_ip_address")
	if !iok {
		return diag.Errorf("Please provide public IP address of the virtual machine")
	}
	// dummy read - since this resource is going to be used for configuring cloud fqdn, never to query the fqdn
	if err := d.Set("cloud_fqdn", "test.ntnxsherlock.com"); err != nil {
		return diag.Errorf("failed to set attribute cloud_fqdn for vm_cloud_config %s", d.Id())
	}
	return diags
}

// ConfigureCloudFqdn - Configure cloud FQDN on the given node
func configureCloudFqdn(publicIP string, cloudFqdn string) (respCode int, err error) {
	configUrl := "http://" + publicIP + ":8080/v1/configure"
	contentType := "application/json"
	data := "{ \"CloudMgmtFQDN\": \"" + cloudFqdn + "\" }"
	stringReader := strings.NewReader(data)
	var done bool = false
	for !done {
		res, err := http.Post(configUrl, contentType, stringReader)
		if err != nil {
			log.Printf("Configure Cloud FQDN Error: %s", err)
		} else {
			byteRes, parseErr := ioutil.ReadAll(res.Body)
			res.Body.Close()
			if parseErr != nil {
				log.Printf("Configure Cloud FQDN Parse Error: %s", parseErr)
				return -1, parseErr
			} else {
				respCode, _ = strconv.Atoi(string(byteRes))
				return respCode, nil
			}
		}
		time.Sleep(1000 * time.Millisecond)
	}
	return -1, nil
}

