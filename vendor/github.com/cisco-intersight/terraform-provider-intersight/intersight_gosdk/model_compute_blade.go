/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-06-17T02:04:50-07:00.
 *
 * API version: 1.0.9-1867
 * Contact: intersight@cisco.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// ComputeBlade Server which is housed in a chassis and shares some of the hardware with other servers in the chassis.
type ComputeBlade struct {
	ComputePhysical
	// The id of the chassis that the blade is located in.
	ChassisId *string `json:"ChassisId,omitempty"`
	// The mode of the server that determines it is scaled.
	ScaledMode *string `json:"ScaledMode,omitempty"`
	// The slot number in the chassis that the blade is located in.
	SlotId *int64 `json:"SlotId,omitempty"`
	// An array of relationships to adapterUnit resources.
	Adapters []AdapterUnitRelationship `json:"Adapters,omitempty"`
	// An array of relationships to biosUnit resources.
	BiosUnits        []BiosUnitRelationship            `json:"BiosUnits,omitempty"`
	Bmc              *ManagementControllerRelationship `json:"Bmc,omitempty"`
	Board            *ComputeBoardRelationship         `json:"Board,omitempty"`
	EquipmentChassis *EquipmentChassisRelationship     `json:"EquipmentChassis,omitempty"`
	// An array of relationships to equipmentIoExpander resources.
	EquipmentIoExpanders []EquipmentIoExpanderRelationship `json:"EquipmentIoExpanders,omitempty"`
	// An array of relationships to inventoryGenericInventoryHolder resources.
	GenericInventoryHolders []InventoryGenericInventoryHolderRelationship `json:"GenericInventoryHolders,omitempty"`
	// An array of relationships to graphicsCard resources.
	GraphicsCards       []GraphicsCardRelationship       `json:"GraphicsCards,omitempty"`
	InventoryDeviceInfo *InventoryDeviceInfoRelationship `json:"InventoryDeviceInfo,omitempty"`
	LocatorLed          *EquipmentLocatorLedRelationship `json:"LocatorLed,omitempty"`
	// An array of relationships to memoryArray resources.
	MemoryArrays []MemoryArrayRelationship `json:"MemoryArrays,omitempty"`
	// An array of relationships to pciDevice resources.
	PciDevices []PciDeviceRelationship `json:"PciDevices,omitempty"`
	// An array of relationships to processorUnit resources.
	Processors       []ProcessorUnitRelationship          `json:"Processors,omitempty"`
	RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	// An array of relationships to storageController resources.
	StorageControllers []StorageControllerRelationship `json:"StorageControllers,omitempty"`
	// An array of relationships to storageEnclosure resources.
	StorageEnclosures    []StorageEnclosureRelationship      `json:"StorageEnclosures,omitempty"`
	TopSystem            *TopSystemRelationship              `json:"TopSystem,omitempty"`
	UemConnection        *InventoryUemConnectionRelationship `json:"UemConnection,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ComputeBlade ComputeBlade

// NewComputeBlade instantiates a new ComputeBlade object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComputeBlade() *ComputeBlade {
	this := ComputeBlade{}
	return &this
}

// NewComputeBladeWithDefaults instantiates a new ComputeBlade object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComputeBladeWithDefaults() *ComputeBlade {
	this := ComputeBlade{}
	return &this
}

// GetChassisId returns the ChassisId field value if set, zero value otherwise.
func (o *ComputeBlade) GetChassisId() string {
	if o == nil || o.ChassisId == nil {
		var ret string
		return ret
	}
	return *o.ChassisId
}

// GetChassisIdOk returns a tuple with the ChassisId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeBlade) GetChassisIdOk() (*string, bool) {
	if o == nil || o.ChassisId == nil {
		return nil, false
	}
	return o.ChassisId, true
}

// HasChassisId returns a boolean if a field has been set.
func (o *ComputeBlade) HasChassisId() bool {
	if o != nil && o.ChassisId != nil {
		return true
	}

	return false
}

// SetChassisId gets a reference to the given string and assigns it to the ChassisId field.
func (o *ComputeBlade) SetChassisId(v string) {
	o.ChassisId = &v
}

// GetScaledMode returns the ScaledMode field value if set, zero value otherwise.
func (o *ComputeBlade) GetScaledMode() string {
	if o == nil || o.ScaledMode == nil {
		var ret string
		return ret
	}
	return *o.ScaledMode
}

// GetScaledModeOk returns a tuple with the ScaledMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeBlade) GetScaledModeOk() (*string, bool) {
	if o == nil || o.ScaledMode == nil {
		return nil, false
	}
	return o.ScaledMode, true
}

// HasScaledMode returns a boolean if a field has been set.
func (o *ComputeBlade) HasScaledMode() bool {
	if o != nil && o.ScaledMode != nil {
		return true
	}

	return false
}

// SetScaledMode gets a reference to the given string and assigns it to the ScaledMode field.
func (o *ComputeBlade) SetScaledMode(v string) {
	o.ScaledMode = &v
}

// GetSlotId returns the SlotId field value if set, zero value otherwise.
func (o *ComputeBlade) GetSlotId() int64 {
	if o == nil || o.SlotId == nil {
		var ret int64
		return ret
	}
	return *o.SlotId
}

// GetSlotIdOk returns a tuple with the SlotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeBlade) GetSlotIdOk() (*int64, bool) {
	if o == nil || o.SlotId == nil {
		return nil, false
	}
	return o.SlotId, true
}

// HasSlotId returns a boolean if a field has been set.
func (o *ComputeBlade) HasSlotId() bool {
	if o != nil && o.SlotId != nil {
		return true
	}

	return false
}

// SetSlotId gets a reference to the given int64 and assigns it to the SlotId field.
func (o *ComputeBlade) SetSlotId(v int64) {
	o.SlotId = &v
}

// GetAdapters returns the Adapters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComputeBlade) GetAdapters() []AdapterUnitRelationship {
	if o == nil {
		var ret []AdapterUnitRelationship
		return ret
	}
	return o.Adapters
}

// GetAdaptersOk returns a tuple with the Adapters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComputeBlade) GetAdaptersOk() (*[]AdapterUnitRelationship, bool) {
	if o == nil || o.Adapters == nil {
		return nil, false
	}
	return &o.Adapters, true
}

// HasAdapters returns a boolean if a field has been set.
func (o *ComputeBlade) HasAdapters() bool {
	if o != nil && o.Adapters != nil {
		return true
	}

	return false
}

// SetAdapters gets a reference to the given []AdapterUnitRelationship and assigns it to the Adapters field.
func (o *ComputeBlade) SetAdapters(v []AdapterUnitRelationship) {
	o.Adapters = v
}

// GetBiosUnits returns the BiosUnits field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComputeBlade) GetBiosUnits() []BiosUnitRelationship {
	if o == nil {
		var ret []BiosUnitRelationship
		return ret
	}
	return o.BiosUnits
}

// GetBiosUnitsOk returns a tuple with the BiosUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComputeBlade) GetBiosUnitsOk() (*[]BiosUnitRelationship, bool) {
	if o == nil || o.BiosUnits == nil {
		return nil, false
	}
	return &o.BiosUnits, true
}

// HasBiosUnits returns a boolean if a field has been set.
func (o *ComputeBlade) HasBiosUnits() bool {
	if o != nil && o.BiosUnits != nil {
		return true
	}

	return false
}

// SetBiosUnits gets a reference to the given []BiosUnitRelationship and assigns it to the BiosUnits field.
func (o *ComputeBlade) SetBiosUnits(v []BiosUnitRelationship) {
	o.BiosUnits = v
}

// GetBmc returns the Bmc field value if set, zero value otherwise.
func (o *ComputeBlade) GetBmc() ManagementControllerRelationship {
	if o == nil || o.Bmc == nil {
		var ret ManagementControllerRelationship
		return ret
	}
	return *o.Bmc
}

// GetBmcOk returns a tuple with the Bmc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeBlade) GetBmcOk() (*ManagementControllerRelationship, bool) {
	if o == nil || o.Bmc == nil {
		return nil, false
	}
	return o.Bmc, true
}

// HasBmc returns a boolean if a field has been set.
func (o *ComputeBlade) HasBmc() bool {
	if o != nil && o.Bmc != nil {
		return true
	}

	return false
}

// SetBmc gets a reference to the given ManagementControllerRelationship and assigns it to the Bmc field.
func (o *ComputeBlade) SetBmc(v ManagementControllerRelationship) {
	o.Bmc = &v
}

// GetBoard returns the Board field value if set, zero value otherwise.
func (o *ComputeBlade) GetBoard() ComputeBoardRelationship {
	if o == nil || o.Board == nil {
		var ret ComputeBoardRelationship
		return ret
	}
	return *o.Board
}

// GetBoardOk returns a tuple with the Board field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeBlade) GetBoardOk() (*ComputeBoardRelationship, bool) {
	if o == nil || o.Board == nil {
		return nil, false
	}
	return o.Board, true
}

// HasBoard returns a boolean if a field has been set.
func (o *ComputeBlade) HasBoard() bool {
	if o != nil && o.Board != nil {
		return true
	}

	return false
}

// SetBoard gets a reference to the given ComputeBoardRelationship and assigns it to the Board field.
func (o *ComputeBlade) SetBoard(v ComputeBoardRelationship) {
	o.Board = &v
}

// GetEquipmentChassis returns the EquipmentChassis field value if set, zero value otherwise.
func (o *ComputeBlade) GetEquipmentChassis() EquipmentChassisRelationship {
	if o == nil || o.EquipmentChassis == nil {
		var ret EquipmentChassisRelationship
		return ret
	}
	return *o.EquipmentChassis
}

// GetEquipmentChassisOk returns a tuple with the EquipmentChassis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeBlade) GetEquipmentChassisOk() (*EquipmentChassisRelationship, bool) {
	if o == nil || o.EquipmentChassis == nil {
		return nil, false
	}
	return o.EquipmentChassis, true
}

// HasEquipmentChassis returns a boolean if a field has been set.
func (o *ComputeBlade) HasEquipmentChassis() bool {
	if o != nil && o.EquipmentChassis != nil {
		return true
	}

	return false
}

// SetEquipmentChassis gets a reference to the given EquipmentChassisRelationship and assigns it to the EquipmentChassis field.
func (o *ComputeBlade) SetEquipmentChassis(v EquipmentChassisRelationship) {
	o.EquipmentChassis = &v
}

// GetEquipmentIoExpanders returns the EquipmentIoExpanders field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComputeBlade) GetEquipmentIoExpanders() []EquipmentIoExpanderRelationship {
	if o == nil {
		var ret []EquipmentIoExpanderRelationship
		return ret
	}
	return o.EquipmentIoExpanders
}

// GetEquipmentIoExpandersOk returns a tuple with the EquipmentIoExpanders field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComputeBlade) GetEquipmentIoExpandersOk() (*[]EquipmentIoExpanderRelationship, bool) {
	if o == nil || o.EquipmentIoExpanders == nil {
		return nil, false
	}
	return &o.EquipmentIoExpanders, true
}

// HasEquipmentIoExpanders returns a boolean if a field has been set.
func (o *ComputeBlade) HasEquipmentIoExpanders() bool {
	if o != nil && o.EquipmentIoExpanders != nil {
		return true
	}

	return false
}

// SetEquipmentIoExpanders gets a reference to the given []EquipmentIoExpanderRelationship and assigns it to the EquipmentIoExpanders field.
func (o *ComputeBlade) SetEquipmentIoExpanders(v []EquipmentIoExpanderRelationship) {
	o.EquipmentIoExpanders = v
}

// GetGenericInventoryHolders returns the GenericInventoryHolders field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComputeBlade) GetGenericInventoryHolders() []InventoryGenericInventoryHolderRelationship {
	if o == nil {
		var ret []InventoryGenericInventoryHolderRelationship
		return ret
	}
	return o.GenericInventoryHolders
}

// GetGenericInventoryHoldersOk returns a tuple with the GenericInventoryHolders field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComputeBlade) GetGenericInventoryHoldersOk() (*[]InventoryGenericInventoryHolderRelationship, bool) {
	if o == nil || o.GenericInventoryHolders == nil {
		return nil, false
	}
	return &o.GenericInventoryHolders, true
}

// HasGenericInventoryHolders returns a boolean if a field has been set.
func (o *ComputeBlade) HasGenericInventoryHolders() bool {
	if o != nil && o.GenericInventoryHolders != nil {
		return true
	}

	return false
}

// SetGenericInventoryHolders gets a reference to the given []InventoryGenericInventoryHolderRelationship and assigns it to the GenericInventoryHolders field.
func (o *ComputeBlade) SetGenericInventoryHolders(v []InventoryGenericInventoryHolderRelationship) {
	o.GenericInventoryHolders = v
}

// GetGraphicsCards returns the GraphicsCards field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComputeBlade) GetGraphicsCards() []GraphicsCardRelationship {
	if o == nil {
		var ret []GraphicsCardRelationship
		return ret
	}
	return o.GraphicsCards
}

// GetGraphicsCardsOk returns a tuple with the GraphicsCards field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComputeBlade) GetGraphicsCardsOk() (*[]GraphicsCardRelationship, bool) {
	if o == nil || o.GraphicsCards == nil {
		return nil, false
	}
	return &o.GraphicsCards, true
}

// HasGraphicsCards returns a boolean if a field has been set.
func (o *ComputeBlade) HasGraphicsCards() bool {
	if o != nil && o.GraphicsCards != nil {
		return true
	}

	return false
}

// SetGraphicsCards gets a reference to the given []GraphicsCardRelationship and assigns it to the GraphicsCards field.
func (o *ComputeBlade) SetGraphicsCards(v []GraphicsCardRelationship) {
	o.GraphicsCards = v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *ComputeBlade) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeBlade) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *ComputeBlade) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *ComputeBlade) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetLocatorLed returns the LocatorLed field value if set, zero value otherwise.
func (o *ComputeBlade) GetLocatorLed() EquipmentLocatorLedRelationship {
	if o == nil || o.LocatorLed == nil {
		var ret EquipmentLocatorLedRelationship
		return ret
	}
	return *o.LocatorLed
}

// GetLocatorLedOk returns a tuple with the LocatorLed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeBlade) GetLocatorLedOk() (*EquipmentLocatorLedRelationship, bool) {
	if o == nil || o.LocatorLed == nil {
		return nil, false
	}
	return o.LocatorLed, true
}

// HasLocatorLed returns a boolean if a field has been set.
func (o *ComputeBlade) HasLocatorLed() bool {
	if o != nil && o.LocatorLed != nil {
		return true
	}

	return false
}

// SetLocatorLed gets a reference to the given EquipmentLocatorLedRelationship and assigns it to the LocatorLed field.
func (o *ComputeBlade) SetLocatorLed(v EquipmentLocatorLedRelationship) {
	o.LocatorLed = &v
}

// GetMemoryArrays returns the MemoryArrays field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComputeBlade) GetMemoryArrays() []MemoryArrayRelationship {
	if o == nil {
		var ret []MemoryArrayRelationship
		return ret
	}
	return o.MemoryArrays
}

// GetMemoryArraysOk returns a tuple with the MemoryArrays field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComputeBlade) GetMemoryArraysOk() (*[]MemoryArrayRelationship, bool) {
	if o == nil || o.MemoryArrays == nil {
		return nil, false
	}
	return &o.MemoryArrays, true
}

// HasMemoryArrays returns a boolean if a field has been set.
func (o *ComputeBlade) HasMemoryArrays() bool {
	if o != nil && o.MemoryArrays != nil {
		return true
	}

	return false
}

// SetMemoryArrays gets a reference to the given []MemoryArrayRelationship and assigns it to the MemoryArrays field.
func (o *ComputeBlade) SetMemoryArrays(v []MemoryArrayRelationship) {
	o.MemoryArrays = v
}

// GetPciDevices returns the PciDevices field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComputeBlade) GetPciDevices() []PciDeviceRelationship {
	if o == nil {
		var ret []PciDeviceRelationship
		return ret
	}
	return o.PciDevices
}

// GetPciDevicesOk returns a tuple with the PciDevices field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComputeBlade) GetPciDevicesOk() (*[]PciDeviceRelationship, bool) {
	if o == nil || o.PciDevices == nil {
		return nil, false
	}
	return &o.PciDevices, true
}

// HasPciDevices returns a boolean if a field has been set.
func (o *ComputeBlade) HasPciDevices() bool {
	if o != nil && o.PciDevices != nil {
		return true
	}

	return false
}

// SetPciDevices gets a reference to the given []PciDeviceRelationship and assigns it to the PciDevices field.
func (o *ComputeBlade) SetPciDevices(v []PciDeviceRelationship) {
	o.PciDevices = v
}

// GetProcessors returns the Processors field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComputeBlade) GetProcessors() []ProcessorUnitRelationship {
	if o == nil {
		var ret []ProcessorUnitRelationship
		return ret
	}
	return o.Processors
}

// GetProcessorsOk returns a tuple with the Processors field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComputeBlade) GetProcessorsOk() (*[]ProcessorUnitRelationship, bool) {
	if o == nil || o.Processors == nil {
		return nil, false
	}
	return &o.Processors, true
}

// HasProcessors returns a boolean if a field has been set.
func (o *ComputeBlade) HasProcessors() bool {
	if o != nil && o.Processors != nil {
		return true
	}

	return false
}

// SetProcessors gets a reference to the given []ProcessorUnitRelationship and assigns it to the Processors field.
func (o *ComputeBlade) SetProcessors(v []ProcessorUnitRelationship) {
	o.Processors = v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *ComputeBlade) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeBlade) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *ComputeBlade) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *ComputeBlade) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

// GetStorageControllers returns the StorageControllers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComputeBlade) GetStorageControllers() []StorageControllerRelationship {
	if o == nil {
		var ret []StorageControllerRelationship
		return ret
	}
	return o.StorageControllers
}

// GetStorageControllersOk returns a tuple with the StorageControllers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComputeBlade) GetStorageControllersOk() (*[]StorageControllerRelationship, bool) {
	if o == nil || o.StorageControllers == nil {
		return nil, false
	}
	return &o.StorageControllers, true
}

// HasStorageControllers returns a boolean if a field has been set.
func (o *ComputeBlade) HasStorageControllers() bool {
	if o != nil && o.StorageControllers != nil {
		return true
	}

	return false
}

// SetStorageControllers gets a reference to the given []StorageControllerRelationship and assigns it to the StorageControllers field.
func (o *ComputeBlade) SetStorageControllers(v []StorageControllerRelationship) {
	o.StorageControllers = v
}

// GetStorageEnclosures returns the StorageEnclosures field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ComputeBlade) GetStorageEnclosures() []StorageEnclosureRelationship {
	if o == nil {
		var ret []StorageEnclosureRelationship
		return ret
	}
	return o.StorageEnclosures
}

// GetStorageEnclosuresOk returns a tuple with the StorageEnclosures field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ComputeBlade) GetStorageEnclosuresOk() (*[]StorageEnclosureRelationship, bool) {
	if o == nil || o.StorageEnclosures == nil {
		return nil, false
	}
	return &o.StorageEnclosures, true
}

// HasStorageEnclosures returns a boolean if a field has been set.
func (o *ComputeBlade) HasStorageEnclosures() bool {
	if o != nil && o.StorageEnclosures != nil {
		return true
	}

	return false
}

// SetStorageEnclosures gets a reference to the given []StorageEnclosureRelationship and assigns it to the StorageEnclosures field.
func (o *ComputeBlade) SetStorageEnclosures(v []StorageEnclosureRelationship) {
	o.StorageEnclosures = v
}

// GetTopSystem returns the TopSystem field value if set, zero value otherwise.
func (o *ComputeBlade) GetTopSystem() TopSystemRelationship {
	if o == nil || o.TopSystem == nil {
		var ret TopSystemRelationship
		return ret
	}
	return *o.TopSystem
}

// GetTopSystemOk returns a tuple with the TopSystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeBlade) GetTopSystemOk() (*TopSystemRelationship, bool) {
	if o == nil || o.TopSystem == nil {
		return nil, false
	}
	return o.TopSystem, true
}

// HasTopSystem returns a boolean if a field has been set.
func (o *ComputeBlade) HasTopSystem() bool {
	if o != nil && o.TopSystem != nil {
		return true
	}

	return false
}

// SetTopSystem gets a reference to the given TopSystemRelationship and assigns it to the TopSystem field.
func (o *ComputeBlade) SetTopSystem(v TopSystemRelationship) {
	o.TopSystem = &v
}

// GetUemConnection returns the UemConnection field value if set, zero value otherwise.
func (o *ComputeBlade) GetUemConnection() InventoryUemConnectionRelationship {
	if o == nil || o.UemConnection == nil {
		var ret InventoryUemConnectionRelationship
		return ret
	}
	return *o.UemConnection
}

// GetUemConnectionOk returns a tuple with the UemConnection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeBlade) GetUemConnectionOk() (*InventoryUemConnectionRelationship, bool) {
	if o == nil || o.UemConnection == nil {
		return nil, false
	}
	return o.UemConnection, true
}

// HasUemConnection returns a boolean if a field has been set.
func (o *ComputeBlade) HasUemConnection() bool {
	if o != nil && o.UemConnection != nil {
		return true
	}

	return false
}

// SetUemConnection gets a reference to the given InventoryUemConnectionRelationship and assigns it to the UemConnection field.
func (o *ComputeBlade) SetUemConnection(v InventoryUemConnectionRelationship) {
	o.UemConnection = &v
}

func (o ComputeBlade) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedComputePhysical, errComputePhysical := json.Marshal(o.ComputePhysical)
	if errComputePhysical != nil {
		return []byte{}, errComputePhysical
	}
	errComputePhysical = json.Unmarshal([]byte(serializedComputePhysical), &toSerialize)
	if errComputePhysical != nil {
		return []byte{}, errComputePhysical
	}
	if o.ChassisId != nil {
		toSerialize["ChassisId"] = o.ChassisId
	}
	if o.ScaledMode != nil {
		toSerialize["ScaledMode"] = o.ScaledMode
	}
	if o.SlotId != nil {
		toSerialize["SlotId"] = o.SlotId
	}
	if o.Adapters != nil {
		toSerialize["Adapters"] = o.Adapters
	}
	if o.BiosUnits != nil {
		toSerialize["BiosUnits"] = o.BiosUnits
	}
	if o.Bmc != nil {
		toSerialize["Bmc"] = o.Bmc
	}
	if o.Board != nil {
		toSerialize["Board"] = o.Board
	}
	if o.EquipmentChassis != nil {
		toSerialize["EquipmentChassis"] = o.EquipmentChassis
	}
	if o.EquipmentIoExpanders != nil {
		toSerialize["EquipmentIoExpanders"] = o.EquipmentIoExpanders
	}
	if o.GenericInventoryHolders != nil {
		toSerialize["GenericInventoryHolders"] = o.GenericInventoryHolders
	}
	if o.GraphicsCards != nil {
		toSerialize["GraphicsCards"] = o.GraphicsCards
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.LocatorLed != nil {
		toSerialize["LocatorLed"] = o.LocatorLed
	}
	if o.MemoryArrays != nil {
		toSerialize["MemoryArrays"] = o.MemoryArrays
	}
	if o.PciDevices != nil {
		toSerialize["PciDevices"] = o.PciDevices
	}
	if o.Processors != nil {
		toSerialize["Processors"] = o.Processors
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	if o.StorageControllers != nil {
		toSerialize["StorageControllers"] = o.StorageControllers
	}
	if o.StorageEnclosures != nil {
		toSerialize["StorageEnclosures"] = o.StorageEnclosures
	}
	if o.TopSystem != nil {
		toSerialize["TopSystem"] = o.TopSystem
	}
	if o.UemConnection != nil {
		toSerialize["UemConnection"] = o.UemConnection
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ComputeBlade) UnmarshalJSON(bytes []byte) (err error) {
	type ComputeBladeWithoutEmbeddedStruct struct {
		// The id of the chassis that the blade is located in.
		ChassisId *string `json:"ChassisId,omitempty"`
		// The mode of the server that determines it is scaled.
		ScaledMode *string `json:"ScaledMode,omitempty"`
		// The slot number in the chassis that the blade is located in.
		SlotId *int64 `json:"SlotId,omitempty"`
		// An array of relationships to adapterUnit resources.
		Adapters []AdapterUnitRelationship `json:"Adapters,omitempty"`
		// An array of relationships to biosUnit resources.
		BiosUnits        []BiosUnitRelationship            `json:"BiosUnits,omitempty"`
		Bmc              *ManagementControllerRelationship `json:"Bmc,omitempty"`
		Board            *ComputeBoardRelationship         `json:"Board,omitempty"`
		EquipmentChassis *EquipmentChassisRelationship     `json:"EquipmentChassis,omitempty"`
		// An array of relationships to equipmentIoExpander resources.
		EquipmentIoExpanders []EquipmentIoExpanderRelationship `json:"EquipmentIoExpanders,omitempty"`
		// An array of relationships to inventoryGenericInventoryHolder resources.
		GenericInventoryHolders []InventoryGenericInventoryHolderRelationship `json:"GenericInventoryHolders,omitempty"`
		// An array of relationships to graphicsCard resources.
		GraphicsCards       []GraphicsCardRelationship       `json:"GraphicsCards,omitempty"`
		InventoryDeviceInfo *InventoryDeviceInfoRelationship `json:"InventoryDeviceInfo,omitempty"`
		LocatorLed          *EquipmentLocatorLedRelationship `json:"LocatorLed,omitempty"`
		// An array of relationships to memoryArray resources.
		MemoryArrays []MemoryArrayRelationship `json:"MemoryArrays,omitempty"`
		// An array of relationships to pciDevice resources.
		PciDevices []PciDeviceRelationship `json:"PciDevices,omitempty"`
		// An array of relationships to processorUnit resources.
		Processors       []ProcessorUnitRelationship          `json:"Processors,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
		// An array of relationships to storageController resources.
		StorageControllers []StorageControllerRelationship `json:"StorageControllers,omitempty"`
		// An array of relationships to storageEnclosure resources.
		StorageEnclosures []StorageEnclosureRelationship      `json:"StorageEnclosures,omitempty"`
		TopSystem         *TopSystemRelationship              `json:"TopSystem,omitempty"`
		UemConnection     *InventoryUemConnectionRelationship `json:"UemConnection,omitempty"`
	}

	varComputeBladeWithoutEmbeddedStruct := ComputeBladeWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varComputeBladeWithoutEmbeddedStruct)
	if err == nil {
		varComputeBlade := _ComputeBlade{}
		varComputeBlade.ChassisId = varComputeBladeWithoutEmbeddedStruct.ChassisId
		varComputeBlade.ScaledMode = varComputeBladeWithoutEmbeddedStruct.ScaledMode
		varComputeBlade.SlotId = varComputeBladeWithoutEmbeddedStruct.SlotId
		varComputeBlade.Adapters = varComputeBladeWithoutEmbeddedStruct.Adapters
		varComputeBlade.BiosUnits = varComputeBladeWithoutEmbeddedStruct.BiosUnits
		varComputeBlade.Bmc = varComputeBladeWithoutEmbeddedStruct.Bmc
		varComputeBlade.Board = varComputeBladeWithoutEmbeddedStruct.Board
		varComputeBlade.EquipmentChassis = varComputeBladeWithoutEmbeddedStruct.EquipmentChassis
		varComputeBlade.EquipmentIoExpanders = varComputeBladeWithoutEmbeddedStruct.EquipmentIoExpanders
		varComputeBlade.GenericInventoryHolders = varComputeBladeWithoutEmbeddedStruct.GenericInventoryHolders
		varComputeBlade.GraphicsCards = varComputeBladeWithoutEmbeddedStruct.GraphicsCards
		varComputeBlade.InventoryDeviceInfo = varComputeBladeWithoutEmbeddedStruct.InventoryDeviceInfo
		varComputeBlade.LocatorLed = varComputeBladeWithoutEmbeddedStruct.LocatorLed
		varComputeBlade.MemoryArrays = varComputeBladeWithoutEmbeddedStruct.MemoryArrays
		varComputeBlade.PciDevices = varComputeBladeWithoutEmbeddedStruct.PciDevices
		varComputeBlade.Processors = varComputeBladeWithoutEmbeddedStruct.Processors
		varComputeBlade.RegisteredDevice = varComputeBladeWithoutEmbeddedStruct.RegisteredDevice
		varComputeBlade.StorageControllers = varComputeBladeWithoutEmbeddedStruct.StorageControllers
		varComputeBlade.StorageEnclosures = varComputeBladeWithoutEmbeddedStruct.StorageEnclosures
		varComputeBlade.TopSystem = varComputeBladeWithoutEmbeddedStruct.TopSystem
		varComputeBlade.UemConnection = varComputeBladeWithoutEmbeddedStruct.UemConnection
		*o = ComputeBlade(varComputeBlade)
	} else {
		return err
	}

	varComputeBlade := _ComputeBlade{}

	err = json.Unmarshal(bytes, &varComputeBlade)
	if err == nil {
		o.ComputePhysical = varComputeBlade.ComputePhysical
	} else {
		return err
	}
	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ChassisId")
		delete(additionalProperties, "ScaledMode")
		delete(additionalProperties, "SlotId")
		delete(additionalProperties, "Adapters")
		delete(additionalProperties, "BiosUnits")
		delete(additionalProperties, "Bmc")
		delete(additionalProperties, "Board")
		delete(additionalProperties, "EquipmentChassis")
		delete(additionalProperties, "EquipmentIoExpanders")
		delete(additionalProperties, "GenericInventoryHolders")
		delete(additionalProperties, "GraphicsCards")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "LocatorLed")
		delete(additionalProperties, "MemoryArrays")
		delete(additionalProperties, "PciDevices")
		delete(additionalProperties, "Processors")
		delete(additionalProperties, "RegisteredDevice")
		delete(additionalProperties, "StorageControllers")
		delete(additionalProperties, "StorageEnclosures")
		delete(additionalProperties, "TopSystem")
		delete(additionalProperties, "UemConnection")

		// remove fields from embedded structs
		reflectComputePhysical := reflect.ValueOf(o.ComputePhysical)
		for i := 0; i < reflectComputePhysical.Type().NumField(); i++ {
			t := reflectComputePhysical.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableComputeBlade struct {
	value *ComputeBlade
	isSet bool
}

func (v NullableComputeBlade) Get() *ComputeBlade {
	return v.value
}

func (v *NullableComputeBlade) Set(val *ComputeBlade) {
	v.value = val
	v.isSet = true
}

func (v NullableComputeBlade) IsSet() bool {
	return v.isSet
}

func (v *NullableComputeBlade) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComputeBlade(val *ComputeBlade) *NullableComputeBlade {
	return &NullableComputeBlade{value: val, isSet: true}
}

func (v NullableComputeBlade) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComputeBlade) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
