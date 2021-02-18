package nutanixkps

import (
	"context"
	"fmt"
	"log"
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	// "github.com/hashicorp/terraform/helper/validation"

	"terraform-provider-nutanixkps/nutanixkpsclient"
	"terraform-provider-nutanixkps/utils"

	"terraform-provider-nutanixkps/generated/kps_cloud_api_swagger/models"
)

const (
	NUTANIXVOLUMESTYPE = "NutanixVolumes"
)

func resourceStorageProfile() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceStorageProfileCreate,
		ReadContext:   resourceStorageProfileRead,
		UpdateContext: resourceStorageProfileUpdate,
		DeleteContext: resourceStorageProfileDelete,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Description: "Name of the service instance: Maximum length of 200 characters.",
				Type:     schema.TypeString,
				Required: true,
			},
			"service_domain_id": {
				Description: "ID of the Service Domain to which this Storage Profile belongs.",
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"type": {
				Description: "Type of the Storage Profile, auto-computed. No input required for this field.",
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"is_default": {
				Description: `Default setting is true.
				Set to false indicates this storage profile is not the default profile for the service domain.`,
				Type:     schema.TypeBool,
				Required: true,
			},
			"nutanix_volumes_config": {
				Description: "Configuration for the Nutanix AOS cluster storage which uses Nutanix Volumes as the backend storage for this profile.",
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				// ConflictsWith: []string{"active_passive_config", "single_master_config"},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"data_services_ip": {
							Description: `IPv4 data services address used by Nutanix Volumes.
							Nutanix Volumes is a load-balanced iSCSI target storage feature of AOS.
							You can obtain this field value by logging in to the Nutanix Prism console.
							In the Prism web console, see Cluster Details to get the cluster data services IP address.`,
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.IsIPAddress,
						},
						"data_services_port": {
							Description: "Data services default port 3260.",
							Type:         schema.TypeInt,
							Optional:     true,
							Default:      3260,
							ValidateFunc: validation.IntAtLeast(1),
						},
						"flash_mode": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"prism_element_cluster_port": {
							Description: "Prism Element cluster default port 9440.",
							Type:         schema.TypeInt,
							Optional:     true,
							Default:      9440,
							ValidateFunc: validation.IntAtLeast(1),
						},
						"prism_element_cluster_vip": {
							Description: `This field sets a logical IP address for the cluster.
							You can obtain this field value by logging in to the Nutanix Prism console. 
							In the Prism web console, see Cluster Details to get the cluster virtual IP address.`,
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.IsIPAddress,
						},
						"prism_element_password": {
							Description: `Password for the Prism element cluster.
							You can obtain the credentials by contacting your Prism Element cluster administrator if you don't have them already.`,
							Type:        schema.TypeString,
							Required:    true,
							Sensitive:   true,
							DefaultFunc: schema.EnvDefaultFunc("NUTANIX_PE_PASSWORD", nil),
						},
						"prism_element_username": {
							Description: `User name for the Prism element cluster.
							You can obtain the credentials by contacting your Prism Element cluster administrator if you don't have them already.`,
							Type:        schema.TypeString,
							Required:    true,
							DefaultFunc: schema.EnvDefaultFunc("NUTANIX_PE_USERNAME", nil),
						},
						"prism_element_storage_container_name": {
							Description: `Name of the storage container: The maximum length is 75 characters.
							Allowed characters are uppercase and lowercase standard Latin letters (A-Z and a-z), Simplified Chinese, decimal digits (0-9), dots (.), hyphens (-), and underscores (_).`,
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
		},
	}
}

func resourceStorageProfileCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Print("[Debug] Entering resourceStorageProfileCreate")
	client := m.(*nutanixkpsclient.Client)
	// var diags diag.Diagnostics
	storageProfileName := d.Get("name").(string)
	storageProfileServiceDomainID := d.Get("service_domain_id").(string)
	storageProfileIsDefault := d.Get("is_default").(bool)

	// Warning or errors can be collected in a slice type
	storageProfile := &models.StorageProfile{}
	storageProfile.Name = &storageProfileName
	storageProfile.IsDefault = storageProfileIsDefault

	storageProfileNutanixVolumesConfig, okNutanixVol := d.GetOk("nutanix_volumes_config")
	if okNutanixVol {
		expandedStorageProfileNutanixVolumesConfig, err := expandNutanixVolumeStorageProfile(storageProfileNutanixVolumesConfig.([]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		t := NUTANIXVOLUMESTYPE
		storageProfile.Type = &t
		storageProfile.NutanixVolumesConfig = expandedStorageProfileNutanixVolumesConfig
	}

	utils.PrintToJSON(storageProfile, "to create storageProfile: ")
	storageProfileID, createErr := client.StorageProfileCreate(storageProfileServiceDomainID, storageProfile)
	if createErr != nil {
		return diag.FromErr(nutanixkpsclient.APIErrorToError(createErr))
	}

	d.SetId(*storageProfileID)
	return resourceStorageProfileRead(ctx, d, m)

}

func resourceStorageProfileUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//todo

	return resourceStorageProfileRead(ctx, d, m)
}

func resourceStorageProfileRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Print("[Debug] Entering resourceStorageProfileRead")
	client := m.(*nutanixkpsclient.Client)
	var diags diag.Diagnostics
	storageProfileID := d.Id()
	svcDomainID := d.Get("service_domain_id").(string)
	storageProfile, errGet := client.StorageProfileGet(svcDomainID, storageProfileID)
	if errGet != nil {
		d.SetId("")
		return nil
	}

	utils.PrintToJSON(storageProfile, "read storageProfile: ")

	if err := d.Set("name", storageProfile.Name); err != nil {
		return diag.Errorf("failed to set attribute name for storageProfile %s", d.Id())
	}

	if err := d.Set("type", storageProfile.Type); err != nil {
		return diag.Errorf("failed to set attribute type for storageProfile %s", d.Id())
	}

	d.SetId(storageProfile.ID)
	return diags
}

func resourceStorageProfileDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// client := m.(*nutanixkpsclient.Client)

	// Warning or errors can be collected in a slice type
	// var diags diag.Diagnostics

	// storageProfileID := d.Id()

	// _, err := client.StorageProfileDelete(storageProfileID)
	// if err != nil {
	// 	return diag.FromErr(nutanixkpsclient.APIErrorToError(err))
	// }

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")
	return nil
	// return diags
}

func expandNutanixVolumeStorageProfile(nutanixVolumeStorageProfiles []interface{}) (*models.NutanixVolumesStorageProfileConfig, error) {
	if len(nutanixVolumeStorageProfiles) != 1 {
		return nil, fmt.Errorf("exactly one nutanix volume storage profile must be passed")
	}
	var err error
	nutanixVolumesStorageProfileConfig := &models.NutanixVolumesStorageProfileConfig{}
	nvsp := nutanixVolumeStorageProfiles[0].(map[string]interface{})
	// if dataServicesIP, okDataServicesIP := nvsp["data_services_ip"]; okDataServicesIP {
	// 	dataServicesIPCasted := dataServicesIP.(string)
	// 	nutanixVolumesStorageProfileConfig.DataServicesIP = &dataServicesIPCasted
	// } else {
	// 	return nil, fmt.Errorf("data_services_ip was not set in nutanix volume storage profile")
	// }
	err = expandNutanixVolumeField(nutanixVolumesStorageProfileConfig, nvsp, "DataServicesIP", "data_services_ip")
	if err != nil {
		return nil, err
	}
	// if dataServicesPort, okDataServicesPort := nvsp["data_services_port"]; okDataServicesPort {
	// 	dataServicesPortCasted := int64(dataServicesPort.(int))
	// 	nutanixVolumesStorageProfileConfig.DataServicesPort = &dataServicesPortCasted
	// } else {
	// 	return nil, fmt.Errorf("data_services_port was not set in nutanix volume storage profile")
	// }
	err = expandNutanixVolumeField(nutanixVolumesStorageProfileConfig, nvsp, "DataServicesPort", "data_services_port")
	if err != nil {
		return nil, err
	}
	// if prismElementClusterPort, okPrismElementClusterPort := nvsp["prism_element_cluster_port"]; okPrismElementClusterPort {
	// 	prismElementClusterPortCasted := int64(prismElementClusterPort.(int))
	// 	nutanixVolumesStorageProfileConfig.PrismElementClusterPort = &prismElementClusterPortCasted
	// } else {
	// 	return nil, fmt.Errorf("prism_element_cluster_port was not set in nutanix volume storage profile")
	// }
	err = expandNutanixVolumeField(nutanixVolumesStorageProfileConfig, nvsp, "PrismElementClusterPort", "prism_element_cluster_port")
	if err != nil {
		return nil, err
	}
	// if flashMode, okFlashMode := nvsp["flash_mode"]; okFlashMode {
	// 	nutanixVolumesStorageProfileConfig.FlashMode = flashMode.(bool)
	// } else {
	// 	return nil, fmt.Errorf("flash_mode was not set in nutanix volume storage profile")
	// }
	err = expandNutanixVolumeField(nutanixVolumesStorageProfileConfig, nvsp, "FlashMode", "flash_mode")
	if err != nil {
		return nil, err
	}
	// if prismElementClusterVIP, okPrismElementClusterVIP := nvsp["prism_element_cluster_vip"]; okPrismElementClusterVIP {
	// 	prismElementClusterVIPCasted := prismElementClusterVIP.(string)
	// 	nutanixVolumesStorageProfileConfig.PrismElementClusterVIP = &prismElementClusterVIPCasted
	// } else {
	// 	return nil, fmt.Errorf("prism_element_cluster_vip was not set in nutanix volume storage profile")
	// }
	err = expandNutanixVolumeField(nutanixVolumesStorageProfileConfig, nvsp, "PrismElementClusterVIP", "prism_element_cluster_vip")
	if err != nil {
		return nil, err
	}
	// if prismElementPassword, okPrismElementPassword := nvsp["prism_element_password"]; okPrismElementPassword {
	// 	prismElementPasswordCasted := prismElementPassword.(string)
	// 	nutanixVolumesStorageProfileConfig.PrismElementPassword = &prismElementPasswordCasted
	// } else {
	// 	return nil, fmt.Errorf("prism_element_password was not set in nutanix volume storage profile")
	// }
	err = expandNutanixVolumeField(nutanixVolumesStorageProfileConfig, nvsp, "PrismElementPassword", "prism_element_password")
	if err != nil {
		return nil, err
	}
	err = expandNutanixVolumeField(nutanixVolumesStorageProfileConfig, nvsp, "PrismElementUserName", "prism_element_username")
	if err != nil {
		return nil, err
	}
	err = expandNutanixVolumeField(nutanixVolumesStorageProfileConfig, nvsp, "StorageContainerName", "prism_element_storage_container_name")
	if err != nil {
		return nil, err
	}
	utils.PrintToJSON(nutanixVolumesStorageProfileConfig, "nutanixVolumesStorageProfileConfig: ")
	return nutanixVolumesStorageProfileConfig, nil
}

func expandNutanixVolumeField(nvspc interface{}, toExpandMap map[string]interface{}, sfield, mfield string) error {
	utils.PrintToJSON(nvspc, "PRE nvspc: ")
	log.Printf("sfield: %s", sfield)
	mapVal, okMapVal := toExpandMap[mfield]
	if !okMapVal {
		return fmt.Errorf("%s was not set in nutanix volume storage profile", mfield)
	}
	ps := reflect.ValueOf(nvspc)
	// struct
	s := ps.Elem()
	if s.Kind() == reflect.Struct {
		// exported field
		f := s.FieldByName(sfield)
		if f.IsValid() {
			if f.CanSet() {
				// log.Print("f.Kind(): ")
				// log.Print(f.Kind())
				// log.Print("f.Type(): ")
				// log.Print(f.Type())
				log.Print("f.Type().String()  ")
				log.Print(f.Type().String() == "*string")
				switch v := f.Type().String(); {
				case v == "*int64":
					log.Printf("entering == *int")
					mapValInt := int64(mapVal.(int))
					f.Set(reflect.ValueOf(&mapValInt))
				case v == "*string":
					log.Printf("entering == string")
					pmapVal := mapVal.(string)
					// f.SetPointer(&pmapVal)
					f.Set(reflect.ValueOf(&pmapVal))
				case v == "bool":
					log.Printf("entering == bool")
					f.SetBool(mapVal.(bool))
				default:
					return fmt.Errorf("unsupported type for field %s", sfield)
				}
			} else {
				return fmt.Errorf("cannot set value of field %s", sfield)
			}
		} else {
			return fmt.Errorf("invalid field %s", sfield)
		}
	} else {
		return fmt.Errorf("struct kind was not reflect.Struct")
	}
	utils.PrintToJSON(nvspc, "POST nvspc: ")
	return nil
}
