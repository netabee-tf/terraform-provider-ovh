---
layout: "ovh"
page_title: "OVH: cloud_project_network_private"
sidebar_current: "docs-ovh-resource-cloud-project-network-private-x"
description: |-
  Creates a private network in a public cloud project.
---

# ovh_cloud_project_network_private

Creates a private network in a public cloud project.

## Example Usage

```hcl
resource "ovh_cloud_project_network_private" "net" {
   project_id = "67890"
   name       = "admin_network"
   regions    = ["GRA1", "BHS1"]
}
```

## Argument Reference

The following arguments are supported:


* `project_id` - (Optional) Deprecated. The id of the public cloud project. If omitted,
    the `OVH_PROJECT_ID` environment variable is used.
    One of `service_name` or `project_id` is required. Conflits with `service_name`.

* `service_name` - (Optional) The id of the public cloud project. If omitted,
    the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used. 
    One of `service_name` or `project_id` is required. Conflits with `project_id`.

* `name` - (Required) The name of the network.

* `vlan_id` - a vlan id to associate with the network.
   Changing this value recreates the resource. Defaults to 0.

* `regions` - an array of valid OVH public cloud region ID in which the network
   will be available. Ex.: "GRA1". Defaults to all public cloud regions.

## Attributes Reference

The following attributes are exported:

* `project_id` - See Argument Reference above.
* `service_name` - See Argument Reference above.
* `name` - See Argument Reference above.
* `vlan_id` - See Argument Reference above.
* `regions` - See Argument Reference above.
* `regions_status` - A map representing the status of the network per region.
* `regions_status/region` - The id of the region.
* `regions_status/status` - The status of the network in the region.
* `status` - the status of the network. should be normally set to 'ACTIVE'.
* `type` - the type of the network. Either 'private' or 'public'. 