/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-07-11T05:47:33Z.
 *
 * API version: 1.0.9-1999
 * Contact: intersight@cisco.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package intersight

import (
	"encoding/json"
)

// EquipmentChassisAllOf Definition of the list of properties defined in 'equipment.Chassis', excluding properties defined in parent classes.
type EquipmentChassisAllOf struct {
	AlarmSummary *ComputeAlarmSummary `json:"AlarmSummary,omitempty"`
	// The assigned identifier for a chassis.
	ChassisId *int64 `json:"ChassisId,omitempty"`
	// This field identifies the connectivity path for the chassis enclosure.
	ConnectionPath *string `json:"ConnectionPath,omitempty"`
	// This field identifies the connectivity status for the chassis enclosure.
	ConnectionStatus *string `json:"ConnectionStatus,omitempty"`
	// This field is to provide description for chassis model.
	Description *string `json:"Description,omitempty"`
	// This field summarizes the faults on the chassis enclosure.
	FaultSummary *int64 `json:"FaultSummary,omitempty"`
	// The management mode of the blade server chassis.
	ManagementMode *string `json:"ManagementMode,omitempty"`
	// This field identifies the name for the chassis enclosure.
	Name *string `json:"Name,omitempty"`
	// This field identifies the Chassis Operational State.
	OperState *string `json:"OperState,omitempty"`
	// Part Number identifier for the chassis enclosure.
	PartNumber *string `json:"PartNumber,omitempty"`
	// This field identifies the Product ID for the chassis enclosure.
	Pid *string `json:"Pid,omitempty"`
	// The platform type that the chassis is a part of.
	PlatformType *string `json:"PlatformType,omitempty"`
	// This field identifies the Product Name for the chassis enclosure.
	ProductName *string `json:"ProductName,omitempty"`
	// This field identifies the Stock Keeping Unit for the chassis enclosure.
	Sku *string `json:"Sku,omitempty"`
	// This field identifies the Vendor ID for the chassis enclosure.
	Vid *string `json:"Vid,omitempty"`
	// An array of relationships to computeBlade resources.
	Blades []ComputeBladeRelationship `json:"Blades,omitempty"`
	// An array of relationships to equipmentFanModule resources.
	Fanmodules          []EquipmentFanModuleRelationship `json:"Fanmodules,omitempty"`
	InventoryDeviceInfo *InventoryDeviceInfoRelationship `json:"InventoryDeviceInfo,omitempty"`
	// An array of relationships to equipmentIoCard resources.
	Ioms       []EquipmentIoCardRelationship    `json:"Ioms,omitempty"`
	LocatorLed *EquipmentLocatorLedRelationship `json:"LocatorLed,omitempty"`
	PsuControl *EquipmentPsuControlRelationship `json:"PsuControl,omitempty"`
	// An array of relationships to equipmentPsu resources.
	Psus             []EquipmentPsuRelationship           `json:"Psus,omitempty"`
	RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	// An array of relationships to storageSasExpander resources.
	Sasexpanders []StorageSasExpanderRelationship `json:"Sasexpanders,omitempty"`
	// An array of relationships to equipmentSystemIoController resources.
	Siocs []EquipmentSystemIoControllerRelationship `json:"Siocs,omitempty"`
	// An array of relationships to storageEnclosure resources.
	StorageEnclosures    []StorageEnclosureRelationship `json:"StorageEnclosures,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EquipmentChassisAllOf EquipmentChassisAllOf

// NewEquipmentChassisAllOf instantiates a new EquipmentChassisAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquipmentChassisAllOf() *EquipmentChassisAllOf {
	this := EquipmentChassisAllOf{}
	var managementMode string = "IntersightStandalone"
	this.ManagementMode = &managementMode
	return &this
}

// NewEquipmentChassisAllOfWithDefaults instantiates a new EquipmentChassisAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquipmentChassisAllOfWithDefaults() *EquipmentChassisAllOf {
	this := EquipmentChassisAllOf{}
	var managementMode string = "IntersightStandalone"
	this.ManagementMode = &managementMode
	return &this
}

// GetAlarmSummary returns the AlarmSummary field value if set, zero value otherwise.
func (o *EquipmentChassisAllOf) GetAlarmSummary() ComputeAlarmSummary {
	if o == nil || o.AlarmSummary == nil {
		var ret ComputeAlarmSummary
		return ret
	}
	return *o.AlarmSummary
}

// GetAlarmSummaryOk returns a tuple with the AlarmSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentChassisAllOf) GetAlarmSummaryOk() (*ComputeAlarmSummary, bool) {
	if o == nil || o.AlarmSummary == nil {
		return nil, false
	}
	return o.AlarmSummary, true
}

// HasAlarmSummary returns a boolean if a field has been set.
func (o *EquipmentChassisAllOf) HasAlarmSummary() bool {
	if o != nil && o.AlarmSummary != nil {
		return true
	}

	return false
}

// SetAlarmSummary gets a reference to the given ComputeAlarmSummary and assigns it to the AlarmSummary field.
func (o *EquipmentChassisAllOf) SetAlarmSummary(v ComputeAlarmSummary) {
	o.AlarmSummary = &v
}

// GetChassisId returns the ChassisId field value if set, zero value otherwise.
func (o *EquipmentChassisAllOf) GetChassisId() int64 {
	if o == nil || o.ChassisId == nil {
		var ret int64
		return ret
	}
	return *o.ChassisId
}

// GetChassisIdOk returns a tuple with the ChassisId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentChassisAllOf) GetChassisIdOk() (*int64, bool) {
	if o == nil || o.ChassisId == nil {
		return nil, false
	}
	return o.ChassisId, true
}

// HasChassisId returns a boolean if a field has been set.
func (o *EquipmentChassisAllOf) HasChassisId() bool {
	if o != nil && o.ChassisId != nil {
		return true
	}

	return false
}

// SetChassisId gets a reference to the given int64 and assigns it to the ChassisId field.
func (o *EquipmentChassisAllOf) SetChassisId(v int64) {
	o.ChassisId = &v
}

// GetConnectionPath returns the ConnectionPath field value if set, zero value otherwise.
func (o *EquipmentChassisAllOf) GetConnectionPath() string {
	if o == nil || o.ConnectionPath == nil {
		var ret string
		return ret
	}
	return *o.ConnectionPath
}

// GetConnectionPathOk returns a tuple with the ConnectionPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentChassisAllOf) GetConnectionPathOk() (*string, bool) {
	if o == nil || o.ConnectionPath == nil {
		return nil, false
	}
	return o.ConnectionPath, true
}

// HasConnectionPath returns a boolean if a field has been set.
func (o *EquipmentChassisAllOf) HasConnectionPath() bool {
	if o != nil && o.ConnectionPath != nil {
		return true
	}

	return false
}

// SetConnectionPath gets a reference to the given string and assigns it to the ConnectionPath field.
func (o *EquipmentChassisAllOf) SetConnectionPath(v string) {
	o.ConnectionPath = &v
}

// GetConnectionStatus returns the ConnectionStatus field value if set, zero value otherwise.
func (o *EquipmentChassisAllOf) GetConnectionStatus() string {
	if o == nil || o.ConnectionStatus == nil {
		var ret string
		return ret
	}
	return *o.ConnectionStatus
}

// GetConnectionStatusOk returns a tuple with the ConnectionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentChassisAllOf) GetConnectionStatusOk() (*string, bool) {
	if o == nil || o.ConnectionStatus == nil {
		return nil, false
	}
	return o.ConnectionStatus, true
}

// HasConnectionStatus returns a boolean if a field has been set.
func (o *EquipmentChassisAllOf) HasConnectionStatus() bool {
	if o != nil && o.ConnectionStatus != nil {
		return true
	}

	return false
}

// SetConnectionStatus gets a reference to the given string and assigns it to the ConnectionStatus field.
func (o *EquipmentChassisAllOf) SetConnectionStatus(v string) {
	o.ConnectionStatus = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EquipmentChassisAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentChassisAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EquipmentChassisAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EquipmentChassisAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetFaultSummary returns the FaultSummary field value if set, zero value otherwise.
func (o *EquipmentChassisAllOf) GetFaultSummary() int64 {
	if o == nil || o.FaultSummary == nil {
		var ret int64
		return ret
	}
	return *o.FaultSummary
}

// GetFaultSummaryOk returns a tuple with the FaultSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentChassisAllOf) GetFaultSummaryOk() (*int64, bool) {
	if o == nil || o.FaultSummary == nil {
		return nil, false
	}
	return o.FaultSummary, true
}

// HasFaultSummary returns a boolean if a field has been set.
func (o *EquipmentChassisAllOf) HasFaultSummary() bool {
	if o != nil && o.FaultSummary != nil {
		return true
	}

	return false
}

// SetFaultSummary gets a reference to the given int64 and assigns it to the FaultSummary field.
func (o *EquipmentChassisAllOf) SetFaultSummary(v int64) {
	o.FaultSummary = &v
}

// GetManagementMode returns the ManagementMode field value if set, zero value otherwise.
func (o *EquipmentChassisAllOf) GetManagementMode() string {
	if o == nil || o.ManagementMode == nil {
		var ret string
		return ret
	}
	return *o.ManagementMode
}

// GetManagementModeOk returns a tuple with the ManagementMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentChassisAllOf) GetManagementModeOk() (*string, bool) {
	if o == nil || o.ManagementMode == nil {
		return nil, false
	}
	return o.ManagementMode, true
}

// HasManagementMode returns a boolean if a field has been set.
func (o *EquipmentChassisAllOf) HasManagementMode() bool {
	if o != nil && o.ManagementMode != nil {
		return true
	}

	return false
}

// SetManagementMode gets a reference to the given string and assigns it to the ManagementMode field.
func (o *EquipmentChassisAllOf) SetManagementMode(v string) {
	o.ManagementMode = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EquipmentChassisAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentChassisAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EquipmentChassisAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EquipmentChassisAllOf) SetName(v string) {
	o.Name = &v
}

// GetOperState returns the OperState field value if set, zero value otherwise.
func (o *EquipmentChassisAllOf) GetOperState() string {
	if o == nil || o.OperState == nil {
		var ret string
		return ret
	}
	return *o.OperState
}

// GetOperStateOk returns a tuple with the OperState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentChassisAllOf) GetOperStateOk() (*string, bool) {
	if o == nil || o.OperState == nil {
		return nil, false
	}
	return o.OperState, true
}

// HasOperState returns a boolean if a field has been set.
func (o *EquipmentChassisAllOf) HasOperState() bool {
	if o != nil && o.OperState != nil {
		return true
	}

	return false
}

// SetOperState gets a reference to the given string and assigns it to the OperState field.
func (o *EquipmentChassisAllOf) SetOperState(v string) {
	o.OperState = &v
}

// GetPartNumber returns the PartNumber field value if set, zero value otherwise.
func (o *EquipmentChassisAllOf) GetPartNumber() string {
	if o == nil || o.PartNumber == nil {
		var ret string
		return ret
	}
	return *o.PartNumber
}

// GetPartNumberOk returns a tuple with the PartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentChassisAllOf) GetPartNumberOk() (*string, bool) {
	if o == nil || o.PartNumber == nil {
		return nil, false
	}
	return o.PartNumber, true
}

// HasPartNumber returns a boolean if a field has been set.
func (o *EquipmentChassisAllOf) HasPartNumber() bool {
	if o != nil && o.PartNumber != nil {
		return true
	}

	return false
}

// SetPartNumber gets a reference to the given string and assigns it to the PartNumber field.
func (o *EquipmentChassisAllOf) SetPartNumber(v string) {
	o.PartNumber = &v
}

// GetPid returns the Pid field value if set, zero value otherwise.
func (o *EquipmentChassisAllOf) GetPid() string {
	if o == nil || o.Pid == nil {
		var ret string
		return ret
	}
	return *o.Pid
}

// GetPidOk returns a tuple with the Pid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentChassisAllOf) GetPidOk() (*string, bool) {
	if o == nil || o.Pid == nil {
		return nil, false
	}
	return o.Pid, true
}

// HasPid returns a boolean if a field has been set.
func (o *EquipmentChassisAllOf) HasPid() bool {
	if o != nil && o.Pid != nil {
		return true
	}

	return false
}

// SetPid gets a reference to the given string and assigns it to the Pid field.
func (o *EquipmentChassisAllOf) SetPid(v string) {
	o.Pid = &v
}

// GetPlatformType returns the PlatformType field value if set, zero value otherwise.
func (o *EquipmentChassisAllOf) GetPlatformType() string {
	if o == nil || o.PlatformType == nil {
		var ret string
		return ret
	}
	return *o.PlatformType
}

// GetPlatformTypeOk returns a tuple with the PlatformType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentChassisAllOf) GetPlatformTypeOk() (*string, bool) {
	if o == nil || o.PlatformType == nil {
		return nil, false
	}
	return o.PlatformType, true
}

// HasPlatformType returns a boolean if a field has been set.
func (o *EquipmentChassisAllOf) HasPlatformType() bool {
	if o != nil && o.PlatformType != nil {
		return true
	}

	return false
}

// SetPlatformType gets a reference to the given string and assigns it to the PlatformType field.
func (o *EquipmentChassisAllOf) SetPlatformType(v string) {
	o.PlatformType = &v
}

// GetProductName returns the ProductName field value if set, zero value otherwise.
func (o *EquipmentChassisAllOf) GetProductName() string {
	if o == nil || o.ProductName == nil {
		var ret string
		return ret
	}
	return *o.ProductName
}

// GetProductNameOk returns a tuple with the ProductName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentChassisAllOf) GetProductNameOk() (*string, bool) {
	if o == nil || o.ProductName == nil {
		return nil, false
	}
	return o.ProductName, true
}

// HasProductName returns a boolean if a field has been set.
func (o *EquipmentChassisAllOf) HasProductName() bool {
	if o != nil && o.ProductName != nil {
		return true
	}

	return false
}

// SetProductName gets a reference to the given string and assigns it to the ProductName field.
func (o *EquipmentChassisAllOf) SetProductName(v string) {
	o.ProductName = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *EquipmentChassisAllOf) GetSku() string {
	if o == nil || o.Sku == nil {
		var ret string
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentChassisAllOf) GetSkuOk() (*string, bool) {
	if o == nil || o.Sku == nil {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *EquipmentChassisAllOf) HasSku() bool {
	if o != nil && o.Sku != nil {
		return true
	}

	return false
}

// SetSku gets a reference to the given string and assigns it to the Sku field.
func (o *EquipmentChassisAllOf) SetSku(v string) {
	o.Sku = &v
}

// GetVid returns the Vid field value if set, zero value otherwise.
func (o *EquipmentChassisAllOf) GetVid() string {
	if o == nil || o.Vid == nil {
		var ret string
		return ret
	}
	return *o.Vid
}

// GetVidOk returns a tuple with the Vid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentChassisAllOf) GetVidOk() (*string, bool) {
	if o == nil || o.Vid == nil {
		return nil, false
	}
	return o.Vid, true
}

// HasVid returns a boolean if a field has been set.
func (o *EquipmentChassisAllOf) HasVid() bool {
	if o != nil && o.Vid != nil {
		return true
	}

	return false
}

// SetVid gets a reference to the given string and assigns it to the Vid field.
func (o *EquipmentChassisAllOf) SetVid(v string) {
	o.Vid = &v
}

// GetBlades returns the Blades field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquipmentChassisAllOf) GetBlades() []ComputeBladeRelationship {
	if o == nil {
		var ret []ComputeBladeRelationship
		return ret
	}
	return o.Blades
}

// GetBladesOk returns a tuple with the Blades field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquipmentChassisAllOf) GetBladesOk() (*[]ComputeBladeRelationship, bool) {
	if o == nil || o.Blades == nil {
		return nil, false
	}
	return &o.Blades, true
}

// HasBlades returns a boolean if a field has been set.
func (o *EquipmentChassisAllOf) HasBlades() bool {
	if o != nil && o.Blades != nil {
		return true
	}

	return false
}

// SetBlades gets a reference to the given []ComputeBladeRelationship and assigns it to the Blades field.
func (o *EquipmentChassisAllOf) SetBlades(v []ComputeBladeRelationship) {
	o.Blades = v
}

// GetFanmodules returns the Fanmodules field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquipmentChassisAllOf) GetFanmodules() []EquipmentFanModuleRelationship {
	if o == nil {
		var ret []EquipmentFanModuleRelationship
		return ret
	}
	return o.Fanmodules
}

// GetFanmodulesOk returns a tuple with the Fanmodules field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquipmentChassisAllOf) GetFanmodulesOk() (*[]EquipmentFanModuleRelationship, bool) {
	if o == nil || o.Fanmodules == nil {
		return nil, false
	}
	return &o.Fanmodules, true
}

// HasFanmodules returns a boolean if a field has been set.
func (o *EquipmentChassisAllOf) HasFanmodules() bool {
	if o != nil && o.Fanmodules != nil {
		return true
	}

	return false
}

// SetFanmodules gets a reference to the given []EquipmentFanModuleRelationship and assigns it to the Fanmodules field.
func (o *EquipmentChassisAllOf) SetFanmodules(v []EquipmentFanModuleRelationship) {
	o.Fanmodules = v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *EquipmentChassisAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentChassisAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *EquipmentChassisAllOf) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *EquipmentChassisAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetIoms returns the Ioms field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquipmentChassisAllOf) GetIoms() []EquipmentIoCardRelationship {
	if o == nil {
		var ret []EquipmentIoCardRelationship
		return ret
	}
	return o.Ioms
}

// GetIomsOk returns a tuple with the Ioms field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquipmentChassisAllOf) GetIomsOk() (*[]EquipmentIoCardRelationship, bool) {
	if o == nil || o.Ioms == nil {
		return nil, false
	}
	return &o.Ioms, true
}

// HasIoms returns a boolean if a field has been set.
func (o *EquipmentChassisAllOf) HasIoms() bool {
	if o != nil && o.Ioms != nil {
		return true
	}

	return false
}

// SetIoms gets a reference to the given []EquipmentIoCardRelationship and assigns it to the Ioms field.
func (o *EquipmentChassisAllOf) SetIoms(v []EquipmentIoCardRelationship) {
	o.Ioms = v
}

// GetLocatorLed returns the LocatorLed field value if set, zero value otherwise.
func (o *EquipmentChassisAllOf) GetLocatorLed() EquipmentLocatorLedRelationship {
	if o == nil || o.LocatorLed == nil {
		var ret EquipmentLocatorLedRelationship
		return ret
	}
	return *o.LocatorLed
}

// GetLocatorLedOk returns a tuple with the LocatorLed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentChassisAllOf) GetLocatorLedOk() (*EquipmentLocatorLedRelationship, bool) {
	if o == nil || o.LocatorLed == nil {
		return nil, false
	}
	return o.LocatorLed, true
}

// HasLocatorLed returns a boolean if a field has been set.
func (o *EquipmentChassisAllOf) HasLocatorLed() bool {
	if o != nil && o.LocatorLed != nil {
		return true
	}

	return false
}

// SetLocatorLed gets a reference to the given EquipmentLocatorLedRelationship and assigns it to the LocatorLed field.
func (o *EquipmentChassisAllOf) SetLocatorLed(v EquipmentLocatorLedRelationship) {
	o.LocatorLed = &v
}

// GetPsuControl returns the PsuControl field value if set, zero value otherwise.
func (o *EquipmentChassisAllOf) GetPsuControl() EquipmentPsuControlRelationship {
	if o == nil || o.PsuControl == nil {
		var ret EquipmentPsuControlRelationship
		return ret
	}
	return *o.PsuControl
}

// GetPsuControlOk returns a tuple with the PsuControl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentChassisAllOf) GetPsuControlOk() (*EquipmentPsuControlRelationship, bool) {
	if o == nil || o.PsuControl == nil {
		return nil, false
	}
	return o.PsuControl, true
}

// HasPsuControl returns a boolean if a field has been set.
func (o *EquipmentChassisAllOf) HasPsuControl() bool {
	if o != nil && o.PsuControl != nil {
		return true
	}

	return false
}

// SetPsuControl gets a reference to the given EquipmentPsuControlRelationship and assigns it to the PsuControl field.
func (o *EquipmentChassisAllOf) SetPsuControl(v EquipmentPsuControlRelationship) {
	o.PsuControl = &v
}

// GetPsus returns the Psus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquipmentChassisAllOf) GetPsus() []EquipmentPsuRelationship {
	if o == nil {
		var ret []EquipmentPsuRelationship
		return ret
	}
	return o.Psus
}

// GetPsusOk returns a tuple with the Psus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquipmentChassisAllOf) GetPsusOk() (*[]EquipmentPsuRelationship, bool) {
	if o == nil || o.Psus == nil {
		return nil, false
	}
	return &o.Psus, true
}

// HasPsus returns a boolean if a field has been set.
func (o *EquipmentChassisAllOf) HasPsus() bool {
	if o != nil && o.Psus != nil {
		return true
	}

	return false
}

// SetPsus gets a reference to the given []EquipmentPsuRelationship and assigns it to the Psus field.
func (o *EquipmentChassisAllOf) SetPsus(v []EquipmentPsuRelationship) {
	o.Psus = v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *EquipmentChassisAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentChassisAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *EquipmentChassisAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *EquipmentChassisAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

// GetSasexpanders returns the Sasexpanders field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquipmentChassisAllOf) GetSasexpanders() []StorageSasExpanderRelationship {
	if o == nil {
		var ret []StorageSasExpanderRelationship
		return ret
	}
	return o.Sasexpanders
}

// GetSasexpandersOk returns a tuple with the Sasexpanders field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquipmentChassisAllOf) GetSasexpandersOk() (*[]StorageSasExpanderRelationship, bool) {
	if o == nil || o.Sasexpanders == nil {
		return nil, false
	}
	return &o.Sasexpanders, true
}

// HasSasexpanders returns a boolean if a field has been set.
func (o *EquipmentChassisAllOf) HasSasexpanders() bool {
	if o != nil && o.Sasexpanders != nil {
		return true
	}

	return false
}

// SetSasexpanders gets a reference to the given []StorageSasExpanderRelationship and assigns it to the Sasexpanders field.
func (o *EquipmentChassisAllOf) SetSasexpanders(v []StorageSasExpanderRelationship) {
	o.Sasexpanders = v
}

// GetSiocs returns the Siocs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquipmentChassisAllOf) GetSiocs() []EquipmentSystemIoControllerRelationship {
	if o == nil {
		var ret []EquipmentSystemIoControllerRelationship
		return ret
	}
	return o.Siocs
}

// GetSiocsOk returns a tuple with the Siocs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquipmentChassisAllOf) GetSiocsOk() (*[]EquipmentSystemIoControllerRelationship, bool) {
	if o == nil || o.Siocs == nil {
		return nil, false
	}
	return &o.Siocs, true
}

// HasSiocs returns a boolean if a field has been set.
func (o *EquipmentChassisAllOf) HasSiocs() bool {
	if o != nil && o.Siocs != nil {
		return true
	}

	return false
}

// SetSiocs gets a reference to the given []EquipmentSystemIoControllerRelationship and assigns it to the Siocs field.
func (o *EquipmentChassisAllOf) SetSiocs(v []EquipmentSystemIoControllerRelationship) {
	o.Siocs = v
}

// GetStorageEnclosures returns the StorageEnclosures field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquipmentChassisAllOf) GetStorageEnclosures() []StorageEnclosureRelationship {
	if o == nil {
		var ret []StorageEnclosureRelationship
		return ret
	}
	return o.StorageEnclosures
}

// GetStorageEnclosuresOk returns a tuple with the StorageEnclosures field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquipmentChassisAllOf) GetStorageEnclosuresOk() (*[]StorageEnclosureRelationship, bool) {
	if o == nil || o.StorageEnclosures == nil {
		return nil, false
	}
	return &o.StorageEnclosures, true
}

// HasStorageEnclosures returns a boolean if a field has been set.
func (o *EquipmentChassisAllOf) HasStorageEnclosures() bool {
	if o != nil && o.StorageEnclosures != nil {
		return true
	}

	return false
}

// SetStorageEnclosures gets a reference to the given []StorageEnclosureRelationship and assigns it to the StorageEnclosures field.
func (o *EquipmentChassisAllOf) SetStorageEnclosures(v []StorageEnclosureRelationship) {
	o.StorageEnclosures = v
}

func (o EquipmentChassisAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AlarmSummary != nil {
		toSerialize["AlarmSummary"] = o.AlarmSummary
	}
	if o.ChassisId != nil {
		toSerialize["ChassisId"] = o.ChassisId
	}
	if o.ConnectionPath != nil {
		toSerialize["ConnectionPath"] = o.ConnectionPath
	}
	if o.ConnectionStatus != nil {
		toSerialize["ConnectionStatus"] = o.ConnectionStatus
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.FaultSummary != nil {
		toSerialize["FaultSummary"] = o.FaultSummary
	}
	if o.ManagementMode != nil {
		toSerialize["ManagementMode"] = o.ManagementMode
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.OperState != nil {
		toSerialize["OperState"] = o.OperState
	}
	if o.PartNumber != nil {
		toSerialize["PartNumber"] = o.PartNumber
	}
	if o.Pid != nil {
		toSerialize["Pid"] = o.Pid
	}
	if o.PlatformType != nil {
		toSerialize["PlatformType"] = o.PlatformType
	}
	if o.ProductName != nil {
		toSerialize["ProductName"] = o.ProductName
	}
	if o.Sku != nil {
		toSerialize["Sku"] = o.Sku
	}
	if o.Vid != nil {
		toSerialize["Vid"] = o.Vid
	}
	if o.Blades != nil {
		toSerialize["Blades"] = o.Blades
	}
	if o.Fanmodules != nil {
		toSerialize["Fanmodules"] = o.Fanmodules
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.Ioms != nil {
		toSerialize["Ioms"] = o.Ioms
	}
	if o.LocatorLed != nil {
		toSerialize["LocatorLed"] = o.LocatorLed
	}
	if o.PsuControl != nil {
		toSerialize["PsuControl"] = o.PsuControl
	}
	if o.Psus != nil {
		toSerialize["Psus"] = o.Psus
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	if o.Sasexpanders != nil {
		toSerialize["Sasexpanders"] = o.Sasexpanders
	}
	if o.Siocs != nil {
		toSerialize["Siocs"] = o.Siocs
	}
	if o.StorageEnclosures != nil {
		toSerialize["StorageEnclosures"] = o.StorageEnclosures
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EquipmentChassisAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varEquipmentChassisAllOf := _EquipmentChassisAllOf{}

	if err = json.Unmarshal(bytes, &varEquipmentChassisAllOf); err == nil {
		*o = EquipmentChassisAllOf(varEquipmentChassisAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "AlarmSummary")
		delete(additionalProperties, "ChassisId")
		delete(additionalProperties, "ConnectionPath")
		delete(additionalProperties, "ConnectionStatus")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "FaultSummary")
		delete(additionalProperties, "ManagementMode")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "OperState")
		delete(additionalProperties, "PartNumber")
		delete(additionalProperties, "Pid")
		delete(additionalProperties, "PlatformType")
		delete(additionalProperties, "ProductName")
		delete(additionalProperties, "Sku")
		delete(additionalProperties, "Vid")
		delete(additionalProperties, "Blades")
		delete(additionalProperties, "Fanmodules")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "Ioms")
		delete(additionalProperties, "LocatorLed")
		delete(additionalProperties, "PsuControl")
		delete(additionalProperties, "Psus")
		delete(additionalProperties, "RegisteredDevice")
		delete(additionalProperties, "Sasexpanders")
		delete(additionalProperties, "Siocs")
		delete(additionalProperties, "StorageEnclosures")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEquipmentChassisAllOf struct {
	value *EquipmentChassisAllOf
	isSet bool
}

func (v NullableEquipmentChassisAllOf) Get() *EquipmentChassisAllOf {
	return v.value
}

func (v *NullableEquipmentChassisAllOf) Set(val *EquipmentChassisAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEquipmentChassisAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEquipmentChassisAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquipmentChassisAllOf(val *EquipmentChassisAllOf) *NullableEquipmentChassisAllOf {
	return &NullableEquipmentChassisAllOf{value: val, isSet: true}
}

func (v NullableEquipmentChassisAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquipmentChassisAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
