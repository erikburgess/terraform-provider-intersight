
---
layout: "intersight"
page_title: "Intersight: intersight_equipment_chassis"
sidebar_current: "docs-intersight-data-source-equipment-chassis"
description: |-
A physical holder housing blade servers.
---

# Data Source: intersight_equipment._chassis
A physical holder housing blade servers.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `chassis_id`:(int) The assigned identifier for a chassis. 
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `connection_path`:(string) This field identifies the connectivity path for the chassis enclosure. 
* `connection_status`:(string) This field identifies the connectivity status for the chassis enclosure. 
* `description`:(string) This field is to provide description for chassis model. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `fault_summary`:(int) This field summarizes the faults on the chassis enclosure. 
* `management_mode`:(string) The management mode of the blade server chassis. 
* `model`:(string) This field identifies the model of the given component. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) This field identifies the name for the chassis enclosure. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `oper_state`:(string) This field identifies the Chassis Operational State. 
* `part_number`:(string) Part Number identifier for the chassis enclosure. 
* `pid`:(string) This field identifies the Product ID for the chassis enclosure. 
* `platform_type`:(string) The platform type that the chassis is a part of. 
* `product_name`:(string) This field identifies the Product Name for the chassis enclosure. 
* `revision`:(string) This field identifies the revision of the given component. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field identifies the serial of the given component. 
* `sku`:(string) This field identifies the Stock Keeping Unit for the chassis enclosure. 
* `vendor`:(string) This field identifies the vendor of the given component. 
* `vid`:(string) This field identifies the Vendor ID for the chassis enclosure. 
