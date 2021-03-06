---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "nutanixkps_servicedomain Resource - terraform-provider-nutanixkps"
subcategory: ""
description: |-
  Describes a Karbon Platform Services Service Domain.
---

# Resource `nutanixkps_servicedomain`

Describes a Karbon Platform Services Service Domain.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **description** (String) Describe the service domain. For example, the main purpose or use case of the service domain.
- **name** (String) Name of the service domain: 
				Name must include lowercase alphanumeric characters and must start and end with an lowercase alphanumeric character.
				Dash (-) and dot (.) characters are allowed as delimiters. Maximum length of 60 characters.

### Optional

- **id** (String) The ID of this resource.
- **last_updated** (String) Last updated timestamp, auto-computed. No input required for this field.
- **virtual_ip** (String) Virtual IPv4 address for this Service Domain


