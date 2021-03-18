package nutanixkps

import (
	"context"
	"net/http"
	"time"
	"io/ioutil"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceVMConfig() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceVMConfigRead,
		Schema:      VMConfigDataSourceMap(),
		Description: "Describes a Karbon Platform Services Virtual Machine.",
	}
}

func dataSourceVMConfigRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	publicIP, iok := d.GetOk("public_ip_address")
	if !iok {
		return diag.Errorf("Please provide public IP address of the virtual machine")
	}
	serialNumber, _ := getSerialNumber(publicIP.(string))
	if err := d.Set("serial_number", serialNumber); err != nil {
		return diag.Errorf("Failed to set attribute serial_number for the virtual machine")
	}
	d.SetId(publicIP.(string))
	return nil
}

func VMConfigDataSourceMap() map[string]*schema.Schema {
	vmdsm := VMConfigElementDataSourceMap()
	vmdsm["public_ip_address"] = &schema.Schema{
		Type:          schema.TypeString,
		Required:      true,
		ValidateFunc: validation.IsIPAddress,
	}
	return vmdsm
}

func VMConfigElementDataSourceMap() map[string]*schema.Schema {
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
