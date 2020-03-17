package intersight

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/cisco-intersight/terraform-provider-intersight/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceFaultInstance() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceFaultInstanceRead,
		Schema: map[string]*schema.Schema{
			"acknowledged": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"affected_dn": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"affected_mo_id": {
				Description: "Managed object Id which was affected.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"affected_mo_type": {
				Description: "Managed object type which was affected.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"ancestor_mo_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ancestor_mo_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"code": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"creation_time": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": {
				Description: "Short summary of the fault found.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"device_mo_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dn": {
				Description: "The Distinguished Name unambiguously identifies an object in the system.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"last_transition_time": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"moid": {
				Description: "The unique identifier of this Managed Object instance.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"num_occurrences": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"object_type": {
				Description: "The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"original_severity": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"permission_resources": {
				Description: "A slice of all permission resources (organizations) associated with this object. Permission ties resources and its associated roles/privileges.\nThese resources which can be specified in a permission is PermissionResource. Currently only organizations can be specified in permission.\nAll logical and physical resources part of an organization will have organization in PermissionResources field.\nIf DeviceRegistration contains another DeviceRegistration and if parent is in org1 and child is part of org2, then child objects will\nhave PermissionResources as org1 and org2. Parent Objects will have PermissionResources as org1.\nAll profiles/policies created with in an organization will have the organization as PermissionResources.",
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"moid": {
							Description: "The Moid of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The Object Type of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'\nis set and 'moid' is empty/absent from the request, Intersight will determine the Moid of the\nresource matching the filter expression and populate it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request. An error is returned if the filter\nmatches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"previous_severity": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"registered_device": {
				Description: "The Device to which this Managed Object is associated.",
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"moid": {
							Description: "The Moid of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"object_type": {
							Description: "The Object Type of the referenced REST resource.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"selector": {
							Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients. If 'moid' is set this field is ignored. If 'selector'\nis set and 'moid' is empty/absent from the request, Intersight will determine the Moid of the\nresource matching the filter expression and populate it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request. An error is returned if the filter\nmatches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
					},
				},
			},
			"rn": {
				Description: "The Relative Name uniquely identifies an object within a given context.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"rule": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"severity": {
				Description: "Severity of the fault found.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"tags": {
				Description: "The array of tags, which allow to add key, value meta-data to managed objects.",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"additional_properties": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: SuppressDiffAdditionProps,
						},
						"key": {
							Description: "The string representation of a tag key.",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"object_type": {
							Description: "The concrete type of this complex type.\nThe ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the \nObjectType is optional. \nThe type is ambiguous when a managed object contains an array of nested documents, and the documents in the array\nare heterogeneous, i.e. the array can contain nested documents of different types.",
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"value": {
							Description: "The string representation of a tag value.",
							Type:        schema.TypeString,
							Optional:    true,
						},
					},
				},
				Computed: true,
			},
		},
	}
}
func dataSourceFaultInstanceRead(d *schema.ResourceData, meta interface{}) error {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)

	url := "fault/Instances"
	var o models.FaultInstance
	if v, ok := d.GetOk("acknowledged"); ok {
		x := (v.(string))
		o.Acknowledged = x
	}
	if v, ok := d.GetOk("affected_dn"); ok {
		x := (v.(string))
		o.AffectedDn = x
	}
	if v, ok := d.GetOk("affected_mo_id"); ok {
		x := (v.(string))
		o.AffectedMoID = x
	}
	if v, ok := d.GetOk("affected_mo_type"); ok {
		x := (v.(string))
		o.AffectedMoType = x
	}
	if v, ok := d.GetOk("ancestor_mo_id"); ok {
		x := (v.(string))
		o.AncestorMoID = x
	}
	if v, ok := d.GetOk("ancestor_mo_type"); ok {
		x := (v.(string))
		o.AncestorMoType = x
	}
	if v, ok := d.GetOk("code"); ok {
		x := (v.(string))
		o.Code = x
	}
	if v, ok := d.GetOk("creation_time"); ok {
		x := (v.(string))
		o.CreationTime = x
	}
	if v, ok := d.GetOk("description"); ok {
		x := (v.(string))
		o.Description = x
	}
	if v, ok := d.GetOk("device_mo_id"); ok {
		x := (v.(string))
		o.DeviceMoID = x
	}
	if v, ok := d.GetOk("dn"); ok {
		x := (v.(string))
		o.Dn = x
	}
	if v, ok := d.GetOk("last_transition_time"); ok {
		x := (v.(string))
		o.LastTransitionTime = x
	}
	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.Moid = x
	}
	if v, ok := d.GetOk("num_occurrences"); ok {
		x := int64(v.(int))
		o.NumOccurrences = x
	}
	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.ObjectType = x
	}
	if v, ok := d.GetOk("original_severity"); ok {
		x := (v.(string))
		o.OriginalSeverity = x
	}
	if v, ok := d.GetOk("previous_severity"); ok {
		x := (v.(string))
		o.PreviousSeverity = x
	}
	if v, ok := d.GetOk("rn"); ok {
		x := (v.(string))
		o.Rn = x
	}
	if v, ok := d.GetOk("rule"); ok {
		x := (v.(string))
		o.Rule = x
	}
	if v, ok := d.GetOk("severity"); ok {
		x := (v.(string))
		o.Severity = x
	}

	data, err := o.MarshalJSON()
	body, err := conn.SendGetRequest(url, data)
	if err != nil {
		return err
	}
	var x = make(map[string]interface{})
	if err = json.Unmarshal(body, &x); err != nil {
		return err
	}
	result := x["Results"]
	if result == nil {
		return fmt.Errorf("your query returned no results. Please change your search criteria and try again")
	}
	switch reflect.TypeOf(result).Kind() {
	case reflect.Slice:
		r := reflect.ValueOf(result)
		for i := 0; i < r.Len(); i++ {
			var s models.FaultInstance
			oo, _ := json.Marshal(r.Index(i).Interface())
			if err = s.UnmarshalJSON(oo); err != nil {
				return err
			}
			if err := d.Set("acknowledged", (s.Acknowledged)); err != nil {
				return err
			}
			if err := d.Set("affected_dn", (s.AffectedDn)); err != nil {
				return err
			}
			if err := d.Set("affected_mo_id", (s.AffectedMoID)); err != nil {
				return err
			}
			if err := d.Set("affected_mo_type", (s.AffectedMoType)); err != nil {
				return err
			}
			if err := d.Set("ancestor_mo_id", (s.AncestorMoID)); err != nil {
				return err
			}
			if err := d.Set("ancestor_mo_type", (s.AncestorMoType)); err != nil {
				return err
			}
			if err := d.Set("code", (s.Code)); err != nil {
				return err
			}
			if err := d.Set("creation_time", (s.CreationTime)); err != nil {
				return err
			}
			if err := d.Set("description", (s.Description)); err != nil {
				return err
			}
			if err := d.Set("device_mo_id", (s.DeviceMoID)); err != nil {
				return err
			}
			if err := d.Set("dn", (s.Dn)); err != nil {
				return err
			}
			if err := d.Set("last_transition_time", (s.LastTransitionTime)); err != nil {
				return err
			}
			if err := d.Set("moid", (s.Moid)); err != nil {
				return err
			}
			if err := d.Set("num_occurrences", (s.NumOccurrences)); err != nil {
				return err
			}
			if err := d.Set("object_type", (s.ObjectType)); err != nil {
				return err
			}
			if err := d.Set("original_severity", (s.OriginalSeverity)); err != nil {
				return err
			}

			if err := d.Set("permission_resources", flattenListMoBaseMoRef(s.PermissionResources, d)); err != nil {
				return err
			}
			if err := d.Set("previous_severity", (s.PreviousSeverity)); err != nil {
				return err
			}

			if err := d.Set("registered_device", flattenMapAssetDeviceRegistrationRef(s.RegisteredDevice, d)); err != nil {
				return err
			}
			if err := d.Set("rn", (s.Rn)); err != nil {
				return err
			}
			if err := d.Set("rule", (s.Rule)); err != nil {
				return err
			}
			if err := d.Set("severity", (s.Severity)); err != nil {
				return err
			}

			if err := d.Set("tags", flattenListMoTag(s.Tags, d)); err != nil {
				return err
			}
			d.SetId(s.Moid)
		}
	}
	return nil
}
