
---
layout: "intersight"
page_title: "Intersight: intersight_niatelemetry_nia_feature_usage"
sidebar_current: "docs-intersight-data-source-niatelemetry-nia-feature-usage"
description: |-
Object available at Device connector scope for feature and fabric information. This applies to APIC environment currently.
---

# Data Source: intersight_niatelemetry._nia_feature_usage
Object available at Device connector scope for feature and fabric information. This applies to APIC environment currently.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `apic_count`:(int) Number of APIC controllers. This determines the value of controllers for the fabric. 
* `app_center_count`:(int) ACI APPs feature usage. This determines the total number of apps installed on the fabric. 
* `ave`:(string) AVE feature usage. This determines if ACI virtual edge feature is enabled or disabled. 
* `bd_count`:(int) Number of BDs. This determines the total number of Broadcast Domains across the fabric. 
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `consistency_checker_app`:(string) Consistency checker application usage. This determines if the fabric has Consistency checker application installed. 
* `contract_count`:(int) Number of contracts. This determines the total number of Contracts configured across the fabric. 
* `dns_count`:(int) DNS feature usage. This determines the total number of DNS configurations across the fabric. 
* `eigrp_count`:(int) Eigrp feature usage. This determines the total number of EIGRP sessions across the fabric. 
* `epg_count`:(int) Number of EPGs. This determines the total number of End Point Groups across the fabric. 
* `hsrp_count`:(int) Hsrp feature usage. This determines the total number of HSRP sessions across the fabric. 
* `ibgp_count`:(int) Ibgp feature usage. This determines the total number of BGP sessions across the fabric. 
* `igmp_access_list_count`:(int) IGMP Access List feature usage. This determines the total number of IGMP access lists configured across the fabric. 
* `igmp_snoop`:(string) IGMP Snooping feature usage. This determines if this feature is enabled or disabled. 
* `ip_epg_count`:(int) Number of IP based EPGs. This determines the total number of IP End Point Groups across the fabric. 
* `isis_count`:(int) Isis feature usage. TThis determines the total number of ISIS sessions across the fabric. 
* `l2_multicast`:(string) L2Multicast feature usage. This determines if this Layer 2 Multicast feature is being enabled / disabled on the fabric. 
* `leaf_count`:(int) Number of Leafs. This determines the total number of Leaf switches in the fabric. 
* `maintenance_mode_count`:(int) Maintenance Mode feature usage. This determines the number of switches that are currently in maintenance mode. 
* `management_over_v6_count`:(int) Management over IPv6 feature usage. This determines the total number of IPv6 configurtaions in the fabric. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `nir`:(string) NIR application usage. This determines if the fabric has NIR application installed. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `opflex_kubernetes_count`:(int) Opflex for Kubernetes feature usage. This determines the total number of VMM sessions of type kubernetes. 
* `ospf_count`:(int) Ospf feature usage. This determines the total number of OSPF sessions across the fabric. 
* `poe_count`:(int) POE feature usage. This determines the total number of POE configurations across the fabric. 
* `qin_vni_tunnel_count`:(int) QinVniTunnel feature usage. This determines if the qinVniTunnel feature is being used on the fabric and the scale of it. 
* `remote_leaf_count`:(int) Number of remote Leafs. This determines if this feature is being enabled or disabled. 
* `scvmm_count`:(int) SCVMM feature usage. This determines the total number of SCVMM configurations in the fabric. 
* `shared_l3_out_count`:(int) SharedL3Out feature usage. This determines the total number of Shared L3 out configured across the fabric. 
* `smart_call_home`:(string) Smart callhome feature usage. This determines if this feature is being enabled or disabled. 
* `snmp`:(string) SNMP feature usage. This determines if this feature is enabled or disabled. 
* `spine_count`:(int) Number of Spines. This determines the total number of spine switches in the fabric. 
* `ssh_over_v6_count`:(int) Ssh over IPv6 feature usage. This determines the total number of IPv6 configurtaions in the fabric. 
* `syslog_over_v6_count`:(int) Syslog over IPv6 feature usage. This determines the total number of IPv6 configurtaions in the fabric. 
* `tenant_count`:(int) Number of tenants. This determines the total number of tenants configured across the fabric. 
* `tier_two_leaf_count`:(int) Number of tier 2 Leafs. This determines the total number of tier 2 Leaf switches in the fabric. 
* `twamp`:(string) TWAMP feature usage. This determines if this feature is enabled or disabled. 
* `useg`:(string) VMM uSegmentation feature usage. This determines if microsegmentation feature is enabled or disabled. 
* `vpod_count`:(int) Virtual pod feature usage. This determines the total number of virtual POD configurations in the fabrics. 
