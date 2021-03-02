---
page_title: "Onboarding an AHV Service Domain"
---

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

Finally leverage Nutanix KPS service domain module which can be found on our [Github](https://github.com/nutanix-xi/sherlock-developer/tree/master/automation/infrastructure/terraform/modules/service_domain) to onboard the recently created VM as service domain. This modules replaces all the configurations you would usually make in the UI as well as provides more detailed configurations such as node info, storage profile info, and others.

```hcl
module "service_domain" {
  source = "../modules/service_domain"
  depends_on = [
    nutanix_virtual_machine.kps_servicedomain_instance
  ]
  instance_info = var.instance_info
  cloud_info = var.cloud_info
  node_info = var.node_info
  service_domain_info = var.service_domain_info
  storage_profile_info = var.storage_profile_info
  kps_servicedomain_instance_details = nutanix_virtual_machine.kps_servicedomain_instance
  storage_config = var.nutanix_volumes_config
  create_ebs_storage_profile = 0
  create_nutanixvolumes_storage_profile = var.create_storage_profile
  private_instance_ips = nutanix_virtual_machine.kps_servicedomain_instance[*].nic_list[0].ip_endpoint_list[0].ip
  public_instance_ips = nutanix_virtual_machine.kps_servicedomain_instance[*].nic_list[0].ip_endpoint_list[0].ip
  wait_for_onboarding = var.wait_for_onboarding
}
```

For more information, take a look at our [Github Workshop](https://github.com/nutanix-xi/sherlock-developer/blob/master/automation/infrastructure/terraform/nutanix/main.tf) regarding onboarding an AHV Service Domain.