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

func dataSourceVirtualMachine() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceVirtualMachineRead,
		Schema:      VirtualMachineDataSourceMap(),
	}
}

func dataSourceVirtualMachineRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	publicIP, iok := d.GetOk("public_ip_address")
	if !iok {
		return diag.Errorf("Please provide public IP address of the virtual machine")
	}
	cloudFqdn, cok := d.GetOk("cloud_fqdn")
	if !cok {
		return diag.Errorf("Please provide cloud FQDN where this VM is / will be onboarded")
	}
	serialNumber, _ := getSerialNumber(publicIP.(string))
	resp, _ := configureCloudFqdn(publicIP.(string), cloudFqdn.(string))
	if resp != 0 {
		return diag.Errorf("Failed to configure cloud FQDN on the virtual machine")
	}
	if err := d.Set("serial_number", serialNumber); err != nil {
		return diag.Errorf("Failed to set attribute serial_number for the virtual machine")
	}
	d.SetId(publicIP.(string))
	return nil
}

func VirtualMachineDataSourceMap() map[string]*schema.Schema {
	vmdsm := VirtualMachineElementDataSourceMap()
	vmdsm["public_ip_address"] = &schema.Schema{
		Type:          schema.TypeString,
		Required:      true,
		ValidateFunc: validation.IsIPAddress,
	}
	return vmdsm
}

func VirtualMachineElementDataSourceMap() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"public_ip_address": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
			ValidateFunc: validation.IsIPAddress,
		},
		"serial_number": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"cloud_fqdn": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
	}
}

// getSerialNumber - Loop until node serial number is available
func getSerialNumber(publicIP string) (serialNumber string, err error) {
	log.Printf("Querying node serial number")
	snUrl := "http://" + publicIP + ":8080/v1/sn"
	var done bool = false
	for !done {
		res, err := http.Get(snUrl)
		if err != nil {
			log.Printf("Error while getting node serial number: %s", err)
		}
		if res != nil {
			serialNumber, parseErr := ioutil.ReadAll(res.Body)
			res.Body.Close()
			if parseErr != nil {
				log.Printf("Error while parsing node serial number: %s", parseErr)
				return "", parseErr
			}
			log.Printf("VM serial number: %s", string(serialNumber))
			return string(serialNumber), nil
		}
		time.Sleep(1000 * time.Millisecond)
	}
	return "", nil
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
