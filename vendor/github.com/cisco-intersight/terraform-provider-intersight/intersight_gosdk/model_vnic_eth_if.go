/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-06-19T15:15:17Z.
 *
 * API version: 1.0.9-1903
 * Contact: intersight@cisco.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package intersight

import (
	"encoding/json"
)

// VnicEthIf Virtual Ethernet Interface.
type VnicEthIf struct {
	MoBaseMo
	Cdn *VnicCdn `json:"Cdn,omitempty"`
	// Setting this to true esnures that the traffic failsover from one uplink to another auotmatically in case of an uplink failure. It is applicable for Cisco VIC adapters only which are connected to Fabric Interconnect cluster. The uplink if specified determines the primary uplink in case of a failover.
	FailoverEnabled *bool `json:"FailoverEnabled,omitempty"`
	// The MAC address that is assigned to the vnic based on the MAC pool that has been assigned to the LAN Connectivity Policy.
	MacAddress *string `json:"MacAddress,omitempty"`
	// Name of the virtual ethernet interface.
	Name *string `json:"Name,omitempty"`
	// The order in which the virtual interface is brought up. The order assigned to an interface should be unique for all the Ethernet and Fibre-Channel interfaces on each PCI link on a VIC adapter. The maximum value of PCI order is limited by the number of virtual interfaces (Ethernet and Fibre-Channel) on each PCI link on a VIC adapter. All VIC adapters have a single PCI link except VIC 1385 which has two.
	Order     *int64                 `json:"Order,omitempty"`
	Placement *VnicPlacementSettings `json:"Placement,omitempty"`
	// The Standby VIF Id is applicable for failover enabled vNICS. It should be the same as the channel number of the standby vethernet created on switch in order to set up the standby data path.
	StandbyVifId  *int64             `json:"StandbyVifId,omitempty"`
	UsnicSettings *VnicUsnicSettings `json:"UsnicSettings,omitempty"`
	// The Vif Id should be same as the channel number of the vethernet created on switch in order to set up the data path. The property is applicable only for FI attached servers where a vethernet is created on the switch for every vNIC.
	VifId                 *int64                                   `json:"VifId,omitempty"`
	VmqSettings           *VnicVmqSettings                         `json:"VmqSettings,omitempty"`
	EthAdapterPolicy      *VnicEthAdapterPolicyRelationship        `json:"EthAdapterPolicy,omitempty"`
	EthNetworkPolicy      *VnicEthNetworkPolicyRelationship        `json:"EthNetworkPolicy,omitempty"`
	EthQosPolicy          *VnicEthQosPolicyRelationship            `json:"EthQosPolicy,omitempty"`
	LanConnectivityPolicy *VnicLanConnectivityPolicyRelationship   `json:"LanConnectivityPolicy,omitempty"`
	LcpVnic               *VnicEthIfRelationship                   `json:"LcpVnic,omitempty"`
	MacLease              *MacpoolLeaseRelationship                `json:"MacLease,omitempty"`
	MacPool               *MacpoolPoolRelationship                 `json:"MacPool,omitempty"`
	Profile               *PolicyAbstractConfigProfileRelationship `json:"Profile,omitempty"`
	// An array of relationships to vnicEthIf resources.
	SpVnics []VnicEthIfRelationship `json:"SpVnics,omitempty"`
}

// NewVnicEthIf instantiates a new VnicEthIf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnicEthIf() *VnicEthIf {
	this := VnicEthIf{}
	return &this
}

// NewVnicEthIfWithDefaults instantiates a new VnicEthIf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnicEthIfWithDefaults() *VnicEthIf {
	this := VnicEthIf{}
	return &this
}

// GetCdn returns the Cdn field value if set, zero value otherwise.
func (o *VnicEthIf) GetCdn() VnicCdn {
	if o == nil || o.Cdn == nil {
		var ret VnicCdn
		return ret
	}
	return *o.Cdn
}

// GetCdnOk returns a tuple with the Cdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthIf) GetCdnOk() (*VnicCdn, bool) {
	if o == nil || o.Cdn == nil {
		return nil, false
	}
	return o.Cdn, true
}

// HasCdn returns a boolean if a field has been set.
func (o *VnicEthIf) HasCdn() bool {
	if o != nil && o.Cdn != nil {
		return true
	}

	return false
}

// SetCdn gets a reference to the given VnicCdn and assigns it to the Cdn field.
func (o *VnicEthIf) SetCdn(v VnicCdn) {
	o.Cdn = &v
}

// GetFailoverEnabled returns the FailoverEnabled field value if set, zero value otherwise.
func (o *VnicEthIf) GetFailoverEnabled() bool {
	if o == nil || o.FailoverEnabled == nil {
		var ret bool
		return ret
	}
	return *o.FailoverEnabled
}

// GetFailoverEnabledOk returns a tuple with the FailoverEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthIf) GetFailoverEnabledOk() (*bool, bool) {
	if o == nil || o.FailoverEnabled == nil {
		return nil, false
	}
	return o.FailoverEnabled, true
}

// HasFailoverEnabled returns a boolean if a field has been set.
func (o *VnicEthIf) HasFailoverEnabled() bool {
	if o != nil && o.FailoverEnabled != nil {
		return true
	}

	return false
}

// SetFailoverEnabled gets a reference to the given bool and assigns it to the FailoverEnabled field.
func (o *VnicEthIf) SetFailoverEnabled(v bool) {
	o.FailoverEnabled = &v
}

// GetMacAddress returns the MacAddress field value if set, zero value otherwise.
func (o *VnicEthIf) GetMacAddress() string {
	if o == nil || o.MacAddress == nil {
		var ret string
		return ret
	}
	return *o.MacAddress
}

// GetMacAddressOk returns a tuple with the MacAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthIf) GetMacAddressOk() (*string, bool) {
	if o == nil || o.MacAddress == nil {
		return nil, false
	}
	return o.MacAddress, true
}

// HasMacAddress returns a boolean if a field has been set.
func (o *VnicEthIf) HasMacAddress() bool {
	if o != nil && o.MacAddress != nil {
		return true
	}

	return false
}

// SetMacAddress gets a reference to the given string and assigns it to the MacAddress field.
func (o *VnicEthIf) SetMacAddress(v string) {
	o.MacAddress = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VnicEthIf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthIf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VnicEthIf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VnicEthIf) SetName(v string) {
	o.Name = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *VnicEthIf) GetOrder() int64 {
	if o == nil || o.Order == nil {
		var ret int64
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthIf) GetOrderOk() (*int64, bool) {
	if o == nil || o.Order == nil {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *VnicEthIf) HasOrder() bool {
	if o != nil && o.Order != nil {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int64 and assigns it to the Order field.
func (o *VnicEthIf) SetOrder(v int64) {
	o.Order = &v
}

// GetPlacement returns the Placement field value if set, zero value otherwise.
func (o *VnicEthIf) GetPlacement() VnicPlacementSettings {
	if o == nil || o.Placement == nil {
		var ret VnicPlacementSettings
		return ret
	}
	return *o.Placement
}

// GetPlacementOk returns a tuple with the Placement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthIf) GetPlacementOk() (*VnicPlacementSettings, bool) {
	if o == nil || o.Placement == nil {
		return nil, false
	}
	return o.Placement, true
}

// HasPlacement returns a boolean if a field has been set.
func (o *VnicEthIf) HasPlacement() bool {
	if o != nil && o.Placement != nil {
		return true
	}

	return false
}

// SetPlacement gets a reference to the given VnicPlacementSettings and assigns it to the Placement field.
func (o *VnicEthIf) SetPlacement(v VnicPlacementSettings) {
	o.Placement = &v
}

// GetStandbyVifId returns the StandbyVifId field value if set, zero value otherwise.
func (o *VnicEthIf) GetStandbyVifId() int64 {
	if o == nil || o.StandbyVifId == nil {
		var ret int64
		return ret
	}
	return *o.StandbyVifId
}

// GetStandbyVifIdOk returns a tuple with the StandbyVifId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthIf) GetStandbyVifIdOk() (*int64, bool) {
	if o == nil || o.StandbyVifId == nil {
		return nil, false
	}
	return o.StandbyVifId, true
}

// HasStandbyVifId returns a boolean if a field has been set.
func (o *VnicEthIf) HasStandbyVifId() bool {
	if o != nil && o.StandbyVifId != nil {
		return true
	}

	return false
}

// SetStandbyVifId gets a reference to the given int64 and assigns it to the StandbyVifId field.
func (o *VnicEthIf) SetStandbyVifId(v int64) {
	o.StandbyVifId = &v
}

// GetUsnicSettings returns the UsnicSettings field value if set, zero value otherwise.
func (o *VnicEthIf) GetUsnicSettings() VnicUsnicSettings {
	if o == nil || o.UsnicSettings == nil {
		var ret VnicUsnicSettings
		return ret
	}
	return *o.UsnicSettings
}

// GetUsnicSettingsOk returns a tuple with the UsnicSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthIf) GetUsnicSettingsOk() (*VnicUsnicSettings, bool) {
	if o == nil || o.UsnicSettings == nil {
		return nil, false
	}
	return o.UsnicSettings, true
}

// HasUsnicSettings returns a boolean if a field has been set.
func (o *VnicEthIf) HasUsnicSettings() bool {
	if o != nil && o.UsnicSettings != nil {
		return true
	}

	return false
}

// SetUsnicSettings gets a reference to the given VnicUsnicSettings and assigns it to the UsnicSettings field.
func (o *VnicEthIf) SetUsnicSettings(v VnicUsnicSettings) {
	o.UsnicSettings = &v
}

// GetVifId returns the VifId field value if set, zero value otherwise.
func (o *VnicEthIf) GetVifId() int64 {
	if o == nil || o.VifId == nil {
		var ret int64
		return ret
	}
	return *o.VifId
}

// GetVifIdOk returns a tuple with the VifId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthIf) GetVifIdOk() (*int64, bool) {
	if o == nil || o.VifId == nil {
		return nil, false
	}
	return o.VifId, true
}

// HasVifId returns a boolean if a field has been set.
func (o *VnicEthIf) HasVifId() bool {
	if o != nil && o.VifId != nil {
		return true
	}

	return false
}

// SetVifId gets a reference to the given int64 and assigns it to the VifId field.
func (o *VnicEthIf) SetVifId(v int64) {
	o.VifId = &v
}

// GetVmqSettings returns the VmqSettings field value if set, zero value otherwise.
func (o *VnicEthIf) GetVmqSettings() VnicVmqSettings {
	if o == nil || o.VmqSettings == nil {
		var ret VnicVmqSettings
		return ret
	}
	return *o.VmqSettings
}

// GetVmqSettingsOk returns a tuple with the VmqSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthIf) GetVmqSettingsOk() (*VnicVmqSettings, bool) {
	if o == nil || o.VmqSettings == nil {
		return nil, false
	}
	return o.VmqSettings, true
}

// HasVmqSettings returns a boolean if a field has been set.
func (o *VnicEthIf) HasVmqSettings() bool {
	if o != nil && o.VmqSettings != nil {
		return true
	}

	return false
}

// SetVmqSettings gets a reference to the given VnicVmqSettings and assigns it to the VmqSettings field.
func (o *VnicEthIf) SetVmqSettings(v VnicVmqSettings) {
	o.VmqSettings = &v
}

// GetEthAdapterPolicy returns the EthAdapterPolicy field value if set, zero value otherwise.
func (o *VnicEthIf) GetEthAdapterPolicy() VnicEthAdapterPolicyRelationship {
	if o == nil || o.EthAdapterPolicy == nil {
		var ret VnicEthAdapterPolicyRelationship
		return ret
	}
	return *o.EthAdapterPolicy
}

// GetEthAdapterPolicyOk returns a tuple with the EthAdapterPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthIf) GetEthAdapterPolicyOk() (*VnicEthAdapterPolicyRelationship, bool) {
	if o == nil || o.EthAdapterPolicy == nil {
		return nil, false
	}
	return o.EthAdapterPolicy, true
}

// HasEthAdapterPolicy returns a boolean if a field has been set.
func (o *VnicEthIf) HasEthAdapterPolicy() bool {
	if o != nil && o.EthAdapterPolicy != nil {
		return true
	}

	return false
}

// SetEthAdapterPolicy gets a reference to the given VnicEthAdapterPolicyRelationship and assigns it to the EthAdapterPolicy field.
func (o *VnicEthIf) SetEthAdapterPolicy(v VnicEthAdapterPolicyRelationship) {
	o.EthAdapterPolicy = &v
}

// GetEthNetworkPolicy returns the EthNetworkPolicy field value if set, zero value otherwise.
func (o *VnicEthIf) GetEthNetworkPolicy() VnicEthNetworkPolicyRelationship {
	if o == nil || o.EthNetworkPolicy == nil {
		var ret VnicEthNetworkPolicyRelationship
		return ret
	}
	return *o.EthNetworkPolicy
}

// GetEthNetworkPolicyOk returns a tuple with the EthNetworkPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthIf) GetEthNetworkPolicyOk() (*VnicEthNetworkPolicyRelationship, bool) {
	if o == nil || o.EthNetworkPolicy == nil {
		return nil, false
	}
	return o.EthNetworkPolicy, true
}

// HasEthNetworkPolicy returns a boolean if a field has been set.
func (o *VnicEthIf) HasEthNetworkPolicy() bool {
	if o != nil && o.EthNetworkPolicy != nil {
		return true
	}

	return false
}

// SetEthNetworkPolicy gets a reference to the given VnicEthNetworkPolicyRelationship and assigns it to the EthNetworkPolicy field.
func (o *VnicEthIf) SetEthNetworkPolicy(v VnicEthNetworkPolicyRelationship) {
	o.EthNetworkPolicy = &v
}

// GetEthQosPolicy returns the EthQosPolicy field value if set, zero value otherwise.
func (o *VnicEthIf) GetEthQosPolicy() VnicEthQosPolicyRelationship {
	if o == nil || o.EthQosPolicy == nil {
		var ret VnicEthQosPolicyRelationship
		return ret
	}
	return *o.EthQosPolicy
}

// GetEthQosPolicyOk returns a tuple with the EthQosPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthIf) GetEthQosPolicyOk() (*VnicEthQosPolicyRelationship, bool) {
	if o == nil || o.EthQosPolicy == nil {
		return nil, false
	}
	return o.EthQosPolicy, true
}

// HasEthQosPolicy returns a boolean if a field has been set.
func (o *VnicEthIf) HasEthQosPolicy() bool {
	if o != nil && o.EthQosPolicy != nil {
		return true
	}

	return false
}

// SetEthQosPolicy gets a reference to the given VnicEthQosPolicyRelationship and assigns it to the EthQosPolicy field.
func (o *VnicEthIf) SetEthQosPolicy(v VnicEthQosPolicyRelationship) {
	o.EthQosPolicy = &v
}

// GetLanConnectivityPolicy returns the LanConnectivityPolicy field value if set, zero value otherwise.
func (o *VnicEthIf) GetLanConnectivityPolicy() VnicLanConnectivityPolicyRelationship {
	if o == nil || o.LanConnectivityPolicy == nil {
		var ret VnicLanConnectivityPolicyRelationship
		return ret
	}
	return *o.LanConnectivityPolicy
}

// GetLanConnectivityPolicyOk returns a tuple with the LanConnectivityPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthIf) GetLanConnectivityPolicyOk() (*VnicLanConnectivityPolicyRelationship, bool) {
	if o == nil || o.LanConnectivityPolicy == nil {
		return nil, false
	}
	return o.LanConnectivityPolicy, true
}

// HasLanConnectivityPolicy returns a boolean if a field has been set.
func (o *VnicEthIf) HasLanConnectivityPolicy() bool {
	if o != nil && o.LanConnectivityPolicy != nil {
		return true
	}

	return false
}

// SetLanConnectivityPolicy gets a reference to the given VnicLanConnectivityPolicyRelationship and assigns it to the LanConnectivityPolicy field.
func (o *VnicEthIf) SetLanConnectivityPolicy(v VnicLanConnectivityPolicyRelationship) {
	o.LanConnectivityPolicy = &v
}

// GetLcpVnic returns the LcpVnic field value if set, zero value otherwise.
func (o *VnicEthIf) GetLcpVnic() VnicEthIfRelationship {
	if o == nil || o.LcpVnic == nil {
		var ret VnicEthIfRelationship
		return ret
	}
	return *o.LcpVnic
}

// GetLcpVnicOk returns a tuple with the LcpVnic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthIf) GetLcpVnicOk() (*VnicEthIfRelationship, bool) {
	if o == nil || o.LcpVnic == nil {
		return nil, false
	}
	return o.LcpVnic, true
}

// HasLcpVnic returns a boolean if a field has been set.
func (o *VnicEthIf) HasLcpVnic() bool {
	if o != nil && o.LcpVnic != nil {
		return true
	}

	return false
}

// SetLcpVnic gets a reference to the given VnicEthIfRelationship and assigns it to the LcpVnic field.
func (o *VnicEthIf) SetLcpVnic(v VnicEthIfRelationship) {
	o.LcpVnic = &v
}

// GetMacLease returns the MacLease field value if set, zero value otherwise.
func (o *VnicEthIf) GetMacLease() MacpoolLeaseRelationship {
	if o == nil || o.MacLease == nil {
		var ret MacpoolLeaseRelationship
		return ret
	}
	return *o.MacLease
}

// GetMacLeaseOk returns a tuple with the MacLease field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthIf) GetMacLeaseOk() (*MacpoolLeaseRelationship, bool) {
	if o == nil || o.MacLease == nil {
		return nil, false
	}
	return o.MacLease, true
}

// HasMacLease returns a boolean if a field has been set.
func (o *VnicEthIf) HasMacLease() bool {
	if o != nil && o.MacLease != nil {
		return true
	}

	return false
}

// SetMacLease gets a reference to the given MacpoolLeaseRelationship and assigns it to the MacLease field.
func (o *VnicEthIf) SetMacLease(v MacpoolLeaseRelationship) {
	o.MacLease = &v
}

// GetMacPool returns the MacPool field value if set, zero value otherwise.
func (o *VnicEthIf) GetMacPool() MacpoolPoolRelationship {
	if o == nil || o.MacPool == nil {
		var ret MacpoolPoolRelationship
		return ret
	}
	return *o.MacPool
}

// GetMacPoolOk returns a tuple with the MacPool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthIf) GetMacPoolOk() (*MacpoolPoolRelationship, bool) {
	if o == nil || o.MacPool == nil {
		return nil, false
	}
	return o.MacPool, true
}

// HasMacPool returns a boolean if a field has been set.
func (o *VnicEthIf) HasMacPool() bool {
	if o != nil && o.MacPool != nil {
		return true
	}

	return false
}

// SetMacPool gets a reference to the given MacpoolPoolRelationship and assigns it to the MacPool field.
func (o *VnicEthIf) SetMacPool(v MacpoolPoolRelationship) {
	o.MacPool = &v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *VnicEthIf) GetProfile() PolicyAbstractConfigProfileRelationship {
	if o == nil || o.Profile == nil {
		var ret PolicyAbstractConfigProfileRelationship
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicEthIf) GetProfileOk() (*PolicyAbstractConfigProfileRelationship, bool) {
	if o == nil || o.Profile == nil {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *VnicEthIf) HasProfile() bool {
	if o != nil && o.Profile != nil {
		return true
	}

	return false
}

// SetProfile gets a reference to the given PolicyAbstractConfigProfileRelationship and assigns it to the Profile field.
func (o *VnicEthIf) SetProfile(v PolicyAbstractConfigProfileRelationship) {
	o.Profile = &v
}

// GetSpVnics returns the SpVnics field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicEthIf) GetSpVnics() []VnicEthIfRelationship {
	if o == nil {
		var ret []VnicEthIfRelationship
		return ret
	}
	return o.SpVnics
}

// GetSpVnicsOk returns a tuple with the SpVnics field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicEthIf) GetSpVnicsOk() (*[]VnicEthIfRelationship, bool) {
	if o == nil || o.SpVnics == nil {
		return nil, false
	}
	return &o.SpVnics, true
}

// HasSpVnics returns a boolean if a field has been set.
func (o *VnicEthIf) HasSpVnics() bool {
	if o != nil && o.SpVnics != nil {
		return true
	}

	return false
}

// SetSpVnics gets a reference to the given []VnicEthIfRelationship and assigns it to the SpVnics field.
func (o *VnicEthIf) SetSpVnics(v []VnicEthIfRelationship) {
	o.SpVnics = v
}

func (o VnicEthIf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if o.Cdn != nil {
		toSerialize["Cdn"] = o.Cdn
	}
	if o.FailoverEnabled != nil {
		toSerialize["FailoverEnabled"] = o.FailoverEnabled
	}
	if o.MacAddress != nil {
		toSerialize["MacAddress"] = o.MacAddress
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Order != nil {
		toSerialize["Order"] = o.Order
	}
	if o.Placement != nil {
		toSerialize["Placement"] = o.Placement
	}
	if o.StandbyVifId != nil {
		toSerialize["StandbyVifId"] = o.StandbyVifId
	}
	if o.UsnicSettings != nil {
		toSerialize["UsnicSettings"] = o.UsnicSettings
	}
	if o.VifId != nil {
		toSerialize["VifId"] = o.VifId
	}
	if o.VmqSettings != nil {
		toSerialize["VmqSettings"] = o.VmqSettings
	}
	if o.EthAdapterPolicy != nil {
		toSerialize["EthAdapterPolicy"] = o.EthAdapterPolicy
	}
	if o.EthNetworkPolicy != nil {
		toSerialize["EthNetworkPolicy"] = o.EthNetworkPolicy
	}
	if o.EthQosPolicy != nil {
		toSerialize["EthQosPolicy"] = o.EthQosPolicy
	}
	if o.LanConnectivityPolicy != nil {
		toSerialize["LanConnectivityPolicy"] = o.LanConnectivityPolicy
	}
	if o.LcpVnic != nil {
		toSerialize["LcpVnic"] = o.LcpVnic
	}
	if o.MacLease != nil {
		toSerialize["MacLease"] = o.MacLease
	}
	if o.MacPool != nil {
		toSerialize["MacPool"] = o.MacPool
	}
	if o.Profile != nil {
		toSerialize["Profile"] = o.Profile
	}
	if o.SpVnics != nil {
		toSerialize["SpVnics"] = o.SpVnics
	}
	return json.Marshal(toSerialize)
}

type NullableVnicEthIf struct {
	value *VnicEthIf
	isSet bool
}

func (v NullableVnicEthIf) Get() *VnicEthIf {
	return v.value
}

func (v *NullableVnicEthIf) Set(val *VnicEthIf) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicEthIf) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicEthIf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicEthIf(val *VnicEthIf) *NullableVnicEthIf {
	return &NullableVnicEthIf{value: val, isSet: true}
}

func (v NullableVnicEthIf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicEthIf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
