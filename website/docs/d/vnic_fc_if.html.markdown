
---
layout: "intersight"
page_title: "Intersight: intersight_vnic_fc_if"
sidebar_current: "docs-intersight-data-source-vnic-fc-if"
description: |-
Virtual Fibre Channel Interface.
---

# Data Source: intersight_vnic._fc_if
Virtual Fibre Channel Interface.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the virtual fibre channel interface. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `order`:(int) The order in which the virtual interface is brought up. The order assigned to an interface should be unique for all the Ethernet and Fibre-Channel interfaces on each PCI link on a VIC adapter. The maximum value of PCI order is limited by the number of virtual interfaces (Ethernet and Fibre-Channel) on each PCI link on a VIC adapter. All VIC adapters have a single PCI link except VIC 1385 which has two. 
* `persistent_bindings`:(bool) Enables retention of LUN ID associations in memory until they are manually cleared. 
* `type`:(string) VHBA Type configuration for SAN Connectivity Policy. This configuration is supported only on Cisco VIC 14XX series and higher series of adapters. 
* `vif_id`:(int) This should be the same as the channel number of the vfc created on switch in order to set up the data path. The property is applicable only for FI attached servers where a vfc is created on the switch for every vHBA. 
* `wwpn`:(string) The WWPN address that is assigned to the vhba based on the wwn pool that has been assigned to the SAN Connectivity Policy. 
