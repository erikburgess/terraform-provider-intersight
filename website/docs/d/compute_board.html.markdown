
---
layout: "intersight"
page_title: "Intersight: intersight_compute_board"
sidebar_current: "docs-intersight-data-source-compute-board"
description: |-
Mother board of a server.
---

# Data Source: intersight_compute._board
Mother board of a server.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `board_id`:(int) The identity of the motherboard. 
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `cpu_type_controller`:(string) The type of central processing unit on the mother board. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `model`:(string) This field identifies the model of the given component. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `oper_power_state`:(string) Current power state of the mother board of the server. 
* `presence`:(string) Identifies the presence of the mother board of the server. 
* `revision`:(string) This field identifies the revision of the given component. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field identifies the serial of the given component. 
* `vendor`:(string) This field identifies the vendor of the given component. 
