<a href="https://terraform.io">
    <img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" alt="Terraform logo" title="Terraform" align="right" height="50" />
</a>

# Terraform Provider for Nutanix KPS

The Terraform Karbon Platform Services provider is a plugin to run the lifecycle management of your KPS resources.

## Quick Starts

- [Homepage](https://registry.terraform.io/providers/nutanix/nutanixkps/latest)
- [Usage examples](https://registry.terraform.io/providers/nutanix/nutanixkps/latest/docs)

## Building on Mac? 
Note: You need to install make on Mac using command line developer tools.
<pre>
make installkpsclient
make build
make install
</pre>

## How to validate provider?
<pre>
cd examples
terraform init
terraform plan
terraform apply -auto-approve
</pre>
