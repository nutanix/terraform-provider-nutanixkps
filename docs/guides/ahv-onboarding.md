---
page_title: "Onboarding an AHV Service Domain"
---

## Onboarding an AHV Service Domain

This guides provides samples on how to onboard an AHV Service Domain to Karbon Platform Services.

### Provider Setup

Begin by instantiating your `nutanixkps` and `nutanix` providers in terraform. This helps the provider autheticate interactions for each user in regards to our API. Simply enter you MyNutanix credentials to get started. 

```hcl
provider "nutanixkps" {
  host = var.cloud_info["cloud_fqdn"]
  username = var.cloud_info["cloud_user_name"]
  password = var.cloud_info["cloud_user_pwd"]
}

provider "nutanix" {
  username = var.provider_info["username"]
  password = var.provider_info["password"]
  endpoint = var.provider_info["endpoint"]
  insecure = var.provider_info["insecure"]
  port     = var.provider_info["port"]
}
```

### Nutanix QCOW2/Service Domain image

Next create a resource to manage your nutanix image for AHV which you can download from our [Support Portal](https://portal.nutanix.com/page/downloads?product=karbonplatformservices). You can name as well as describe the resource being created. There are two options on how to pass in the downloaded qcow2 image into the resource. The user can either provide a local path or you can use a URI path in which case `source_path` will need to be changed to `source_uri`. 

```hcl
data "nutanix_clusters" "clusters" {}

resource "nutanix_image" "kps_servicedomain_image" {
  name        = var.image_config["name"]
  description = var.image_config["description"]
  source_path  = var.image_config["source_path"]
  depends_on = [
    data.nutanix_clusters.clusters
  ]
}
```

### Nutanix Virtual Machine

The image resource can be used to create a virtual machine which will later be onboarded as a KPS Service Domain. This resouce will allow you to configure attributes of the VM such as number of sockets, vCPUs per socket, and memory size.

```hcl
data "nutanix_subnet" "sherlock_net" {
  subnet_name = "sherlock_net"
}

# Create Nutanix KPS SD VM Instance

resource "nutanix_virtual_machine" "kps_servicedomain_instance" {
  provider = nutanix
  cluster_uuid= data.nutanix_clusters.clusters.entities.0.metadata.uuid

  count = var.instance_info["instance_count"]
  name = join("-", [var.instance_info["instance_name_prefix"], count.index])
  description = var.nutanix_vm_config["description"]

  num_vcpus_per_socket = var.nutanix_vm_config["num_vcpus_per_socket"]
  num_sockets          = var.nutanix_vm_config["num_sockets"]
  memory_size_mib      = var.nutanix_vm_config["memory_size_mib"]

  nic_list {
    subnet_name = data.nutanix_subnet.sherlock_net.name
    subnet_uuid = data.nutanix_subnet.sherlock_net.metadata.uuid
  }

  disk_list {
    # data_source_reference in the Nutanix API refers to where the source for
    # the disk device will come from. Could be a clone of a different VM or a
    # image like we're doing here
    data_source_reference = {
      kind = "image"
      uuid = data.nutanix_image.kps_servicedomain_image.image_id
    }
  }

  depends_on = [
    data.nutanix_image.kps_servicedomain_image
  ]
}
```

### Nutanix KPS Service Domain

Finally leverage Nutanix KPS service domain resource to onboard the recently created VM as service domain. This modules replaces all the configurations you would usually make in the UI as well as provides more detailed configurations such as node info, storage profile info, and others.

```hcl
resource "nutanixkps_servicedomain" "kps_servicedomain" {
  name = var.service_domain_info["sd_name"]
  description = var.service_domain_info["sd_description"]
  virtual_ip = var.service_domain_info["sd_virtual_ip"]
}

output "servicedomains" {
  value = nutanixkps_servicedomain.kps_servicedomain
}

data "nutanixkps_virtual_machine" "kps_virtual_machine" {
  count = var.instance_info["instance_count"]
  public_ip_address =  nutanix_virtual_machine.kps_servicedomain_instance[count.index].nic_list[0].ip_endpoint_list[0].ip
  cloud_fqdn = var.cloud_info["cloud_fqdn"]
  depends_on = [
    nutanix_virtual_machine.kps_servicedomain_instance
  ]
}

resource "nutanixkps_node" "nodes" {
  count = var.instance_info["instance_count"]
  name = "${var.instance_info["instance_name_prefix"]}-${count.index}"
  description = "Node added to Service Domain through Terraform"
  service_domain_id = nutanixkps_servicedomain.kps_servicedomain.id
  serial_number = data.nutanixkps_virtual_machine.kps_virtual_machine[count.index].serial_number
  ip_address = nutanix_virtual_machine.kps_servicedomain_instance[count.index].nic_list[0].ip_endpoint_list[0].ip
  gateway = var.node_info["node_gateway"]
  subnet = var.node_info["node_subnet"]
  role {
    master = true
    worker = true
  }
  wait_for_onboarding = var.wait_for_onboarding
  cloud_fqdn = var.cloud_info["cloud_fqdn"]
  depends_on = [
    nutanixkps_servicedomain.kps_servicedomain,
    data.nutanixkps_virtual_machine.kps_virtual_machine
  ]
}

resource "nutanixkps_storageprofile" "nutanixvolumes_storage_profile" {
  // count is used as a boolean here, only effective when creating Nutanix VM based service domain
  count = var.create_storage_profile
  name = var.storage_profile_info["name"]
  service_domain_id = nutanixkps_servicedomain.kps_servicedomain.id
  is_default = var.storage_profile_info["isDefault"]
  nutanix_volumes_config {
    data_services_ip = var.nutanix_volumes_config["dataServicesIP"]
    data_services_port = var.nutanix_volumes_config["dataServicesPort"]
    flash_mode = var.nutanix_volumes_config["flashMode"]
    prism_element_cluster_vip = var.nutanix_volumes_config["prismElementClusterVIP"]
    prism_element_cluster_port = var.nutanix_volumes_config["prismElementClusterPort"]
    prism_element_password = var.nutanix_volumes_config["prismElementPassword"]
    prism_element_username = var.nutanix_volumes_config["prismElementUserName"]
    prism_element_storage_container_name = var.nutanix_volumes_config["storageContainerName"]
  }

  depends_on = [
    nutanixkps_servicedomain.kps_servicedomain
  ]
}
```

For more information, take a look at our [Github Workshop](https://github.com/nutanix/karbon-platform-services/tree/master/automation/infrastructure/terraform) regarding onboarding an AHV Service Domain.