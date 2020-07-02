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

// GraphicsCard Graphics Card present in a server.
type GraphicsCard struct {
	EquipmentBase
	// The id of the graphics card.
	CardId *int64 `json:"CardId,omitempty"`
	// The device id of the graphics card.
	DeviceId *int64 `json:"DeviceId,omitempty"`
	// The expander slot information of the card.
	ExpanderSlot *string `json:"ExpanderSlot,omitempty"`
	// The firmware version of the graphics card.
	FirmwareVersion *string `json:"FirmwareVersion,omitempty"`
	// The current mode of the graphics card.
	Mode *string `json:"Mode,omitempty"`
	// The number of controllers under each card.
	NumGpus *string `json:"NumGpus,omitempty"`
	// The current operational state of the graphics card.
	OperState *string `json:"OperState,omitempty"`
	// The PCI address of the graphics card.
	PciAddress *string `json:"PciAddress,omitempty"`
	// This list contains the PCI address of all controllers for corresponding card.
	PciAddressList *string `json:"PciAddressList,omitempty"`
	// The PCI slot name of the graphics card.
	PciSlot         *string                      `json:"PciSlot,omitempty"`
	ComputeBlade    *ComputeBladeRelationship    `json:"ComputeBlade,omitempty"`
	ComputeBoard    *ComputeBoardRelationship    `json:"ComputeBoard,omitempty"`
	ComputeRackUnit *ComputeRackUnitRelationship `json:"ComputeRackUnit,omitempty"`
	// An array of relationships to graphicsController resources.
	GraphicsControllers []GraphicsControllerRelationship     `json:"GraphicsControllers,omitempty"`
	InventoryDeviceInfo *InventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
	RegisteredDevice    *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	// An array of relationships to firmwareRunningFirmware resources.
	RunningFirmware []FirmwareRunningFirmwareRelationship `json:"RunningFirmware,omitempty"`
}

// NewGraphicsCard instantiates a new GraphicsCard object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGraphicsCard() *GraphicsCard {
	this := GraphicsCard{}
	return &this
}

// NewGraphicsCardWithDefaults instantiates a new GraphicsCard object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGraphicsCardWithDefaults() *GraphicsCard {
	this := GraphicsCard{}
	return &this
}

// GetCardId returns the CardId field value if set, zero value otherwise.
func (o *GraphicsCard) GetCardId() int64 {
	if o == nil || o.CardId == nil {
		var ret int64
		return ret
	}
	return *o.CardId
}

// GetCardIdOk returns a tuple with the CardId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphicsCard) GetCardIdOk() (*int64, bool) {
	if o == nil || o.CardId == nil {
		return nil, false
	}
	return o.CardId, true
}

// HasCardId returns a boolean if a field has been set.
func (o *GraphicsCard) HasCardId() bool {
	if o != nil && o.CardId != nil {
		return true
	}

	return false
}

// SetCardId gets a reference to the given int64 and assigns it to the CardId field.
func (o *GraphicsCard) SetCardId(v int64) {
	o.CardId = &v
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *GraphicsCard) GetDeviceId() int64 {
	if o == nil || o.DeviceId == nil {
		var ret int64
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphicsCard) GetDeviceIdOk() (*int64, bool) {
	if o == nil || o.DeviceId == nil {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *GraphicsCard) HasDeviceId() bool {
	if o != nil && o.DeviceId != nil {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given int64 and assigns it to the DeviceId field.
func (o *GraphicsCard) SetDeviceId(v int64) {
	o.DeviceId = &v
}

// GetExpanderSlot returns the ExpanderSlot field value if set, zero value otherwise.
func (o *GraphicsCard) GetExpanderSlot() string {
	if o == nil || o.ExpanderSlot == nil {
		var ret string
		return ret
	}
	return *o.ExpanderSlot
}

// GetExpanderSlotOk returns a tuple with the ExpanderSlot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphicsCard) GetExpanderSlotOk() (*string, bool) {
	if o == nil || o.ExpanderSlot == nil {
		return nil, false
	}
	return o.ExpanderSlot, true
}

// HasExpanderSlot returns a boolean if a field has been set.
func (o *GraphicsCard) HasExpanderSlot() bool {
	if o != nil && o.ExpanderSlot != nil {
		return true
	}

	return false
}

// SetExpanderSlot gets a reference to the given string and assigns it to the ExpanderSlot field.
func (o *GraphicsCard) SetExpanderSlot(v string) {
	o.ExpanderSlot = &v
}

// GetFirmwareVersion returns the FirmwareVersion field value if set, zero value otherwise.
func (o *GraphicsCard) GetFirmwareVersion() string {
	if o == nil || o.FirmwareVersion == nil {
		var ret string
		return ret
	}
	return *o.FirmwareVersion
}

// GetFirmwareVersionOk returns a tuple with the FirmwareVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphicsCard) GetFirmwareVersionOk() (*string, bool) {
	if o == nil || o.FirmwareVersion == nil {
		return nil, false
	}
	return o.FirmwareVersion, true
}

// HasFirmwareVersion returns a boolean if a field has been set.
func (o *GraphicsCard) HasFirmwareVersion() bool {
	if o != nil && o.FirmwareVersion != nil {
		return true
	}

	return false
}

// SetFirmwareVersion gets a reference to the given string and assigns it to the FirmwareVersion field.
func (o *GraphicsCard) SetFirmwareVersion(v string) {
	o.FirmwareVersion = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *GraphicsCard) GetMode() string {
	if o == nil || o.Mode == nil {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphicsCard) GetModeOk() (*string, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *GraphicsCard) HasMode() bool {
	if o != nil && o.Mode != nil {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *GraphicsCard) SetMode(v string) {
	o.Mode = &v
}

// GetNumGpus returns the NumGpus field value if set, zero value otherwise.
func (o *GraphicsCard) GetNumGpus() string {
	if o == nil || o.NumGpus == nil {
		var ret string
		return ret
	}
	return *o.NumGpus
}

// GetNumGpusOk returns a tuple with the NumGpus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphicsCard) GetNumGpusOk() (*string, bool) {
	if o == nil || o.NumGpus == nil {
		return nil, false
	}
	return o.NumGpus, true
}

// HasNumGpus returns a boolean if a field has been set.
func (o *GraphicsCard) HasNumGpus() bool {
	if o != nil && o.NumGpus != nil {
		return true
	}

	return false
}

// SetNumGpus gets a reference to the given string and assigns it to the NumGpus field.
func (o *GraphicsCard) SetNumGpus(v string) {
	o.NumGpus = &v
}

// GetOperState returns the OperState field value if set, zero value otherwise.
func (o *GraphicsCard) GetOperState() string {
	if o == nil || o.OperState == nil {
		var ret string
		return ret
	}
	return *o.OperState
}

// GetOperStateOk returns a tuple with the OperState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphicsCard) GetOperStateOk() (*string, bool) {
	if o == nil || o.OperState == nil {
		return nil, false
	}
	return o.OperState, true
}

// HasOperState returns a boolean if a field has been set.
func (o *GraphicsCard) HasOperState() bool {
	if o != nil && o.OperState != nil {
		return true
	}

	return false
}

// SetOperState gets a reference to the given string and assigns it to the OperState field.
func (o *GraphicsCard) SetOperState(v string) {
	o.OperState = &v
}

// GetPciAddress returns the PciAddress field value if set, zero value otherwise.
func (o *GraphicsCard) GetPciAddress() string {
	if o == nil || o.PciAddress == nil {
		var ret string
		return ret
	}
	return *o.PciAddress
}

// GetPciAddressOk returns a tuple with the PciAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphicsCard) GetPciAddressOk() (*string, bool) {
	if o == nil || o.PciAddress == nil {
		return nil, false
	}
	return o.PciAddress, true
}

// HasPciAddress returns a boolean if a field has been set.
func (o *GraphicsCard) HasPciAddress() bool {
	if o != nil && o.PciAddress != nil {
		return true
	}

	return false
}

// SetPciAddress gets a reference to the given string and assigns it to the PciAddress field.
func (o *GraphicsCard) SetPciAddress(v string) {
	o.PciAddress = &v
}

// GetPciAddressList returns the PciAddressList field value if set, zero value otherwise.
func (o *GraphicsCard) GetPciAddressList() string {
	if o == nil || o.PciAddressList == nil {
		var ret string
		return ret
	}
	return *o.PciAddressList
}

// GetPciAddressListOk returns a tuple with the PciAddressList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphicsCard) GetPciAddressListOk() (*string, bool) {
	if o == nil || o.PciAddressList == nil {
		return nil, false
	}
	return o.PciAddressList, true
}

// HasPciAddressList returns a boolean if a field has been set.
func (o *GraphicsCard) HasPciAddressList() bool {
	if o != nil && o.PciAddressList != nil {
		return true
	}

	return false
}

// SetPciAddressList gets a reference to the given string and assigns it to the PciAddressList field.
func (o *GraphicsCard) SetPciAddressList(v string) {
	o.PciAddressList = &v
}

// GetPciSlot returns the PciSlot field value if set, zero value otherwise.
func (o *GraphicsCard) GetPciSlot() string {
	if o == nil || o.PciSlot == nil {
		var ret string
		return ret
	}
	return *o.PciSlot
}

// GetPciSlotOk returns a tuple with the PciSlot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphicsCard) GetPciSlotOk() (*string, bool) {
	if o == nil || o.PciSlot == nil {
		return nil, false
	}
	return o.PciSlot, true
}

// HasPciSlot returns a boolean if a field has been set.
func (o *GraphicsCard) HasPciSlot() bool {
	if o != nil && o.PciSlot != nil {
		return true
	}

	return false
}

// SetPciSlot gets a reference to the given string and assigns it to the PciSlot field.
func (o *GraphicsCard) SetPciSlot(v string) {
	o.PciSlot = &v
}

// GetComputeBlade returns the ComputeBlade field value if set, zero value otherwise.
func (o *GraphicsCard) GetComputeBlade() ComputeBladeRelationship {
	if o == nil || o.ComputeBlade == nil {
		var ret ComputeBladeRelationship
		return ret
	}
	return *o.ComputeBlade
}

// GetComputeBladeOk returns a tuple with the ComputeBlade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphicsCard) GetComputeBladeOk() (*ComputeBladeRelationship, bool) {
	if o == nil || o.ComputeBlade == nil {
		return nil, false
	}
	return o.ComputeBlade, true
}

// HasComputeBlade returns a boolean if a field has been set.
func (o *GraphicsCard) HasComputeBlade() bool {
	if o != nil && o.ComputeBlade != nil {
		return true
	}

	return false
}

// SetComputeBlade gets a reference to the given ComputeBladeRelationship and assigns it to the ComputeBlade field.
func (o *GraphicsCard) SetComputeBlade(v ComputeBladeRelationship) {
	o.ComputeBlade = &v
}

// GetComputeBoard returns the ComputeBoard field value if set, zero value otherwise.
func (o *GraphicsCard) GetComputeBoard() ComputeBoardRelationship {
	if o == nil || o.ComputeBoard == nil {
		var ret ComputeBoardRelationship
		return ret
	}
	return *o.ComputeBoard
}

// GetComputeBoardOk returns a tuple with the ComputeBoard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphicsCard) GetComputeBoardOk() (*ComputeBoardRelationship, bool) {
	if o == nil || o.ComputeBoard == nil {
		return nil, false
	}
	return o.ComputeBoard, true
}

// HasComputeBoard returns a boolean if a field has been set.
func (o *GraphicsCard) HasComputeBoard() bool {
	if o != nil && o.ComputeBoard != nil {
		return true
	}

	return false
}

// SetComputeBoard gets a reference to the given ComputeBoardRelationship and assigns it to the ComputeBoard field.
func (o *GraphicsCard) SetComputeBoard(v ComputeBoardRelationship) {
	o.ComputeBoard = &v
}

// GetComputeRackUnit returns the ComputeRackUnit field value if set, zero value otherwise.
func (o *GraphicsCard) GetComputeRackUnit() ComputeRackUnitRelationship {
	if o == nil || o.ComputeRackUnit == nil {
		var ret ComputeRackUnitRelationship
		return ret
	}
	return *o.ComputeRackUnit
}

// GetComputeRackUnitOk returns a tuple with the ComputeRackUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphicsCard) GetComputeRackUnitOk() (*ComputeRackUnitRelationship, bool) {
	if o == nil || o.ComputeRackUnit == nil {
		return nil, false
	}
	return o.ComputeRackUnit, true
}

// HasComputeRackUnit returns a boolean if a field has been set.
func (o *GraphicsCard) HasComputeRackUnit() bool {
	if o != nil && o.ComputeRackUnit != nil {
		return true
	}

	return false
}

// SetComputeRackUnit gets a reference to the given ComputeRackUnitRelationship and assigns it to the ComputeRackUnit field.
func (o *GraphicsCard) SetComputeRackUnit(v ComputeRackUnitRelationship) {
	o.ComputeRackUnit = &v
}

// GetGraphicsControllers returns the GraphicsControllers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GraphicsCard) GetGraphicsControllers() []GraphicsControllerRelationship {
	if o == nil {
		var ret []GraphicsControllerRelationship
		return ret
	}
	return o.GraphicsControllers
}

// GetGraphicsControllersOk returns a tuple with the GraphicsControllers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GraphicsCard) GetGraphicsControllersOk() (*[]GraphicsControllerRelationship, bool) {
	if o == nil || o.GraphicsControllers == nil {
		return nil, false
	}
	return &o.GraphicsControllers, true
}

// HasGraphicsControllers returns a boolean if a field has been set.
func (o *GraphicsCard) HasGraphicsControllers() bool {
	if o != nil && o.GraphicsControllers != nil {
		return true
	}

	return false
}

// SetGraphicsControllers gets a reference to the given []GraphicsControllerRelationship and assigns it to the GraphicsControllers field.
func (o *GraphicsCard) SetGraphicsControllers(v []GraphicsControllerRelationship) {
	o.GraphicsControllers = v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *GraphicsCard) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphicsCard) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *GraphicsCard) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *GraphicsCard) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *GraphicsCard) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphicsCard) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *GraphicsCard) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *GraphicsCard) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

// GetRunningFirmware returns the RunningFirmware field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GraphicsCard) GetRunningFirmware() []FirmwareRunningFirmwareRelationship {
	if o == nil {
		var ret []FirmwareRunningFirmwareRelationship
		return ret
	}
	return o.RunningFirmware
}

// GetRunningFirmwareOk returns a tuple with the RunningFirmware field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GraphicsCard) GetRunningFirmwareOk() (*[]FirmwareRunningFirmwareRelationship, bool) {
	if o == nil || o.RunningFirmware == nil {
		return nil, false
	}
	return &o.RunningFirmware, true
}

// HasRunningFirmware returns a boolean if a field has been set.
func (o *GraphicsCard) HasRunningFirmware() bool {
	if o != nil && o.RunningFirmware != nil {
		return true
	}

	return false
}

// SetRunningFirmware gets a reference to the given []FirmwareRunningFirmwareRelationship and assigns it to the RunningFirmware field.
func (o *GraphicsCard) SetRunningFirmware(v []FirmwareRunningFirmwareRelationship) {
	o.RunningFirmware = v
}

func (o GraphicsCard) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedEquipmentBase, errEquipmentBase := json.Marshal(o.EquipmentBase)
	if errEquipmentBase != nil {
		return []byte{}, errEquipmentBase
	}
	errEquipmentBase = json.Unmarshal([]byte(serializedEquipmentBase), &toSerialize)
	if errEquipmentBase != nil {
		return []byte{}, errEquipmentBase
	}
	if o.CardId != nil {
		toSerialize["CardId"] = o.CardId
	}
	if o.DeviceId != nil {
		toSerialize["DeviceId"] = o.DeviceId
	}
	if o.ExpanderSlot != nil {
		toSerialize["ExpanderSlot"] = o.ExpanderSlot
	}
	if o.FirmwareVersion != nil {
		toSerialize["FirmwareVersion"] = o.FirmwareVersion
	}
	if o.Mode != nil {
		toSerialize["Mode"] = o.Mode
	}
	if o.NumGpus != nil {
		toSerialize["NumGpus"] = o.NumGpus
	}
	if o.OperState != nil {
		toSerialize["OperState"] = o.OperState
	}
	if o.PciAddress != nil {
		toSerialize["PciAddress"] = o.PciAddress
	}
	if o.PciAddressList != nil {
		toSerialize["PciAddressList"] = o.PciAddressList
	}
	if o.PciSlot != nil {
		toSerialize["PciSlot"] = o.PciSlot
	}
	if o.ComputeBlade != nil {
		toSerialize["ComputeBlade"] = o.ComputeBlade
	}
	if o.ComputeBoard != nil {
		toSerialize["ComputeBoard"] = o.ComputeBoard
	}
	if o.ComputeRackUnit != nil {
		toSerialize["ComputeRackUnit"] = o.ComputeRackUnit
	}
	if o.GraphicsControllers != nil {
		toSerialize["GraphicsControllers"] = o.GraphicsControllers
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	if o.RunningFirmware != nil {
		toSerialize["RunningFirmware"] = o.RunningFirmware
	}
	return json.Marshal(toSerialize)
}

type NullableGraphicsCard struct {
	value *GraphicsCard
	isSet bool
}

func (v NullableGraphicsCard) Get() *GraphicsCard {
	return v.value
}

func (v *NullableGraphicsCard) Set(val *GraphicsCard) {
	v.value = val
	v.isSet = true
}

func (v NullableGraphicsCard) IsSet() bool {
	return v.isSet
}

func (v *NullableGraphicsCard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGraphicsCard(val *GraphicsCard) *NullableGraphicsCard {
	return &NullableGraphicsCard{value: val, isSet: true}
}

func (v NullableGraphicsCard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGraphicsCard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
