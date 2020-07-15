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

// MemoryAbstractUnit Abstract class for all memory units.
type MemoryAbstractUnit struct {
	EquipmentBase
	// This represents the administrative state of the memory unit on a server.
	AdminState *string `json:"AdminState,omitempty"`
	// This represents the memory array to which the memory unit belongs to.
	ArrayId *int64 `json:"ArrayId,omitempty"`
	// This represents the memory bank of the memory unit on a server.
	Bank *int64 `json:"Bank,omitempty"`
	// This represents the memory capacity in MiB of the memory unit on a server.
	Capacity *string `json:"Capacity,omitempty"`
	// This represents the clock of the memory unit on a server.
	Clock *string `json:"Clock,omitempty"`
	// This represents the form factor of the memory unit on a server.
	FormFactor *string `json:"FormFactor,omitempty"`
	// This represents the latency of the memory unit on a server.
	Latency *string `json:"Latency,omitempty"`
	// This represents the location of the memory unit on a server.
	Location *string `json:"Location,omitempty"`
	// This represents the operational power state of the memory unit on a server.
	OperPowerState *string `json:"OperPowerState,omitempty"`
	// This represents the operational state of the memory unit on a server.
	OperState *string `json:"OperState,omitempty"`
	// This represents the operability of the memory unit on a server.
	Operability *string `json:"Operability,omitempty"`
	// This represents the presence state of the memory unit on a server.
	Presence *string `json:"Presence,omitempty"`
	// This represents the set of the memory unit on a server.
	Set *int64 `json:"Set,omitempty"`
	// This represents the speed of the memory unit on a server.
	Speed *string `json:"Speed,omitempty"`
	// This represents the thremal state of the memory unit on a server.
	Thermal *string `json:"Thermal,omitempty"`
	// This represents the memory type of the memory unit on a server.
	Type *string `json:"Type,omitempty"`
	// This represents the visibility of the memory unit on a server.
	Visibility *string `json:"Visibility,omitempty"`
	// This represents the width of the memory unit on a server.
	Width                *string `json:"Width,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MemoryAbstractUnit MemoryAbstractUnit

// NewMemoryAbstractUnit instantiates a new MemoryAbstractUnit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMemoryAbstractUnit() *MemoryAbstractUnit {
	this := MemoryAbstractUnit{}
	return &this
}

// NewMemoryAbstractUnitWithDefaults instantiates a new MemoryAbstractUnit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMemoryAbstractUnitWithDefaults() *MemoryAbstractUnit {
	this := MemoryAbstractUnit{}
	return &this
}

// GetAdminState returns the AdminState field value if set, zero value otherwise.
func (o *MemoryAbstractUnit) GetAdminState() string {
	if o == nil || o.AdminState == nil {
		var ret string
		return ret
	}
	return *o.AdminState
}

// GetAdminStateOk returns a tuple with the AdminState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryAbstractUnit) GetAdminStateOk() (*string, bool) {
	if o == nil || o.AdminState == nil {
		return nil, false
	}
	return o.AdminState, true
}

// HasAdminState returns a boolean if a field has been set.
func (o *MemoryAbstractUnit) HasAdminState() bool {
	if o != nil && o.AdminState != nil {
		return true
	}

	return false
}

// SetAdminState gets a reference to the given string and assigns it to the AdminState field.
func (o *MemoryAbstractUnit) SetAdminState(v string) {
	o.AdminState = &v
}

// GetArrayId returns the ArrayId field value if set, zero value otherwise.
func (o *MemoryAbstractUnit) GetArrayId() int64 {
	if o == nil || o.ArrayId == nil {
		var ret int64
		return ret
	}
	return *o.ArrayId
}

// GetArrayIdOk returns a tuple with the ArrayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryAbstractUnit) GetArrayIdOk() (*int64, bool) {
	if o == nil || o.ArrayId == nil {
		return nil, false
	}
	return o.ArrayId, true
}

// HasArrayId returns a boolean if a field has been set.
func (o *MemoryAbstractUnit) HasArrayId() bool {
	if o != nil && o.ArrayId != nil {
		return true
	}

	return false
}

// SetArrayId gets a reference to the given int64 and assigns it to the ArrayId field.
func (o *MemoryAbstractUnit) SetArrayId(v int64) {
	o.ArrayId = &v
}

// GetBank returns the Bank field value if set, zero value otherwise.
func (o *MemoryAbstractUnit) GetBank() int64 {
	if o == nil || o.Bank == nil {
		var ret int64
		return ret
	}
	return *o.Bank
}

// GetBankOk returns a tuple with the Bank field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryAbstractUnit) GetBankOk() (*int64, bool) {
	if o == nil || o.Bank == nil {
		return nil, false
	}
	return o.Bank, true
}

// HasBank returns a boolean if a field has been set.
func (o *MemoryAbstractUnit) HasBank() bool {
	if o != nil && o.Bank != nil {
		return true
	}

	return false
}

// SetBank gets a reference to the given int64 and assigns it to the Bank field.
func (o *MemoryAbstractUnit) SetBank(v int64) {
	o.Bank = &v
}

// GetCapacity returns the Capacity field value if set, zero value otherwise.
func (o *MemoryAbstractUnit) GetCapacity() string {
	if o == nil || o.Capacity == nil {
		var ret string
		return ret
	}
	return *o.Capacity
}

// GetCapacityOk returns a tuple with the Capacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryAbstractUnit) GetCapacityOk() (*string, bool) {
	if o == nil || o.Capacity == nil {
		return nil, false
	}
	return o.Capacity, true
}

// HasCapacity returns a boolean if a field has been set.
func (o *MemoryAbstractUnit) HasCapacity() bool {
	if o != nil && o.Capacity != nil {
		return true
	}

	return false
}

// SetCapacity gets a reference to the given string and assigns it to the Capacity field.
func (o *MemoryAbstractUnit) SetCapacity(v string) {
	o.Capacity = &v
}

// GetClock returns the Clock field value if set, zero value otherwise.
func (o *MemoryAbstractUnit) GetClock() string {
	if o == nil || o.Clock == nil {
		var ret string
		return ret
	}
	return *o.Clock
}

// GetClockOk returns a tuple with the Clock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryAbstractUnit) GetClockOk() (*string, bool) {
	if o == nil || o.Clock == nil {
		return nil, false
	}
	return o.Clock, true
}

// HasClock returns a boolean if a field has been set.
func (o *MemoryAbstractUnit) HasClock() bool {
	if o != nil && o.Clock != nil {
		return true
	}

	return false
}

// SetClock gets a reference to the given string and assigns it to the Clock field.
func (o *MemoryAbstractUnit) SetClock(v string) {
	o.Clock = &v
}

// GetFormFactor returns the FormFactor field value if set, zero value otherwise.
func (o *MemoryAbstractUnit) GetFormFactor() string {
	if o == nil || o.FormFactor == nil {
		var ret string
		return ret
	}
	return *o.FormFactor
}

// GetFormFactorOk returns a tuple with the FormFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryAbstractUnit) GetFormFactorOk() (*string, bool) {
	if o == nil || o.FormFactor == nil {
		return nil, false
	}
	return o.FormFactor, true
}

// HasFormFactor returns a boolean if a field has been set.
func (o *MemoryAbstractUnit) HasFormFactor() bool {
	if o != nil && o.FormFactor != nil {
		return true
	}

	return false
}

// SetFormFactor gets a reference to the given string and assigns it to the FormFactor field.
func (o *MemoryAbstractUnit) SetFormFactor(v string) {
	o.FormFactor = &v
}

// GetLatency returns the Latency field value if set, zero value otherwise.
func (o *MemoryAbstractUnit) GetLatency() string {
	if o == nil || o.Latency == nil {
		var ret string
		return ret
	}
	return *o.Latency
}

// GetLatencyOk returns a tuple with the Latency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryAbstractUnit) GetLatencyOk() (*string, bool) {
	if o == nil || o.Latency == nil {
		return nil, false
	}
	return o.Latency, true
}

// HasLatency returns a boolean if a field has been set.
func (o *MemoryAbstractUnit) HasLatency() bool {
	if o != nil && o.Latency != nil {
		return true
	}

	return false
}

// SetLatency gets a reference to the given string and assigns it to the Latency field.
func (o *MemoryAbstractUnit) SetLatency(v string) {
	o.Latency = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *MemoryAbstractUnit) GetLocation() string {
	if o == nil || o.Location == nil {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryAbstractUnit) GetLocationOk() (*string, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *MemoryAbstractUnit) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *MemoryAbstractUnit) SetLocation(v string) {
	o.Location = &v
}

// GetOperPowerState returns the OperPowerState field value if set, zero value otherwise.
func (o *MemoryAbstractUnit) GetOperPowerState() string {
	if o == nil || o.OperPowerState == nil {
		var ret string
		return ret
	}
	return *o.OperPowerState
}

// GetOperPowerStateOk returns a tuple with the OperPowerState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryAbstractUnit) GetOperPowerStateOk() (*string, bool) {
	if o == nil || o.OperPowerState == nil {
		return nil, false
	}
	return o.OperPowerState, true
}

// HasOperPowerState returns a boolean if a field has been set.
func (o *MemoryAbstractUnit) HasOperPowerState() bool {
	if o != nil && o.OperPowerState != nil {
		return true
	}

	return false
}

// SetOperPowerState gets a reference to the given string and assigns it to the OperPowerState field.
func (o *MemoryAbstractUnit) SetOperPowerState(v string) {
	o.OperPowerState = &v
}

// GetOperState returns the OperState field value if set, zero value otherwise.
func (o *MemoryAbstractUnit) GetOperState() string {
	if o == nil || o.OperState == nil {
		var ret string
		return ret
	}
	return *o.OperState
}

// GetOperStateOk returns a tuple with the OperState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryAbstractUnit) GetOperStateOk() (*string, bool) {
	if o == nil || o.OperState == nil {
		return nil, false
	}
	return o.OperState, true
}

// HasOperState returns a boolean if a field has been set.
func (o *MemoryAbstractUnit) HasOperState() bool {
	if o != nil && o.OperState != nil {
		return true
	}

	return false
}

// SetOperState gets a reference to the given string and assigns it to the OperState field.
func (o *MemoryAbstractUnit) SetOperState(v string) {
	o.OperState = &v
}

// GetOperability returns the Operability field value if set, zero value otherwise.
func (o *MemoryAbstractUnit) GetOperability() string {
	if o == nil || o.Operability == nil {
		var ret string
		return ret
	}
	return *o.Operability
}

// GetOperabilityOk returns a tuple with the Operability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryAbstractUnit) GetOperabilityOk() (*string, bool) {
	if o == nil || o.Operability == nil {
		return nil, false
	}
	return o.Operability, true
}

// HasOperability returns a boolean if a field has been set.
func (o *MemoryAbstractUnit) HasOperability() bool {
	if o != nil && o.Operability != nil {
		return true
	}

	return false
}

// SetOperability gets a reference to the given string and assigns it to the Operability field.
func (o *MemoryAbstractUnit) SetOperability(v string) {
	o.Operability = &v
}

// GetPresence returns the Presence field value if set, zero value otherwise.
func (o *MemoryAbstractUnit) GetPresence() string {
	if o == nil || o.Presence == nil {
		var ret string
		return ret
	}
	return *o.Presence
}

// GetPresenceOk returns a tuple with the Presence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryAbstractUnit) GetPresenceOk() (*string, bool) {
	if o == nil || o.Presence == nil {
		return nil, false
	}
	return o.Presence, true
}

// HasPresence returns a boolean if a field has been set.
func (o *MemoryAbstractUnit) HasPresence() bool {
	if o != nil && o.Presence != nil {
		return true
	}

	return false
}

// SetPresence gets a reference to the given string and assigns it to the Presence field.
func (o *MemoryAbstractUnit) SetPresence(v string) {
	o.Presence = &v
}

// GetSet returns the Set field value if set, zero value otherwise.
func (o *MemoryAbstractUnit) GetSet() int64 {
	if o == nil || o.Set == nil {
		var ret int64
		return ret
	}
	return *o.Set
}

// GetSetOk returns a tuple with the Set field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryAbstractUnit) GetSetOk() (*int64, bool) {
	if o == nil || o.Set == nil {
		return nil, false
	}
	return o.Set, true
}

// HasSet returns a boolean if a field has been set.
func (o *MemoryAbstractUnit) HasSet() bool {
	if o != nil && o.Set != nil {
		return true
	}

	return false
}

// SetSet gets a reference to the given int64 and assigns it to the Set field.
func (o *MemoryAbstractUnit) SetSet(v int64) {
	o.Set = &v
}

// GetSpeed returns the Speed field value if set, zero value otherwise.
func (o *MemoryAbstractUnit) GetSpeed() string {
	if o == nil || o.Speed == nil {
		var ret string
		return ret
	}
	return *o.Speed
}

// GetSpeedOk returns a tuple with the Speed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryAbstractUnit) GetSpeedOk() (*string, bool) {
	if o == nil || o.Speed == nil {
		return nil, false
	}
	return o.Speed, true
}

// HasSpeed returns a boolean if a field has been set.
func (o *MemoryAbstractUnit) HasSpeed() bool {
	if o != nil && o.Speed != nil {
		return true
	}

	return false
}

// SetSpeed gets a reference to the given string and assigns it to the Speed field.
func (o *MemoryAbstractUnit) SetSpeed(v string) {
	o.Speed = &v
}

// GetThermal returns the Thermal field value if set, zero value otherwise.
func (o *MemoryAbstractUnit) GetThermal() string {
	if o == nil || o.Thermal == nil {
		var ret string
		return ret
	}
	return *o.Thermal
}

// GetThermalOk returns a tuple with the Thermal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryAbstractUnit) GetThermalOk() (*string, bool) {
	if o == nil || o.Thermal == nil {
		return nil, false
	}
	return o.Thermal, true
}

// HasThermal returns a boolean if a field has been set.
func (o *MemoryAbstractUnit) HasThermal() bool {
	if o != nil && o.Thermal != nil {
		return true
	}

	return false
}

// SetThermal gets a reference to the given string and assigns it to the Thermal field.
func (o *MemoryAbstractUnit) SetThermal(v string) {
	o.Thermal = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MemoryAbstractUnit) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryAbstractUnit) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MemoryAbstractUnit) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *MemoryAbstractUnit) SetType(v string) {
	o.Type = &v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *MemoryAbstractUnit) GetVisibility() string {
	if o == nil || o.Visibility == nil {
		var ret string
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryAbstractUnit) GetVisibilityOk() (*string, bool) {
	if o == nil || o.Visibility == nil {
		return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *MemoryAbstractUnit) HasVisibility() bool {
	if o != nil && o.Visibility != nil {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given string and assigns it to the Visibility field.
func (o *MemoryAbstractUnit) SetVisibility(v string) {
	o.Visibility = &v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *MemoryAbstractUnit) GetWidth() string {
	if o == nil || o.Width == nil {
		var ret string
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemoryAbstractUnit) GetWidthOk() (*string, bool) {
	if o == nil || o.Width == nil {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *MemoryAbstractUnit) HasWidth() bool {
	if o != nil && o.Width != nil {
		return true
	}

	return false
}

// SetWidth gets a reference to the given string and assigns it to the Width field.
func (o *MemoryAbstractUnit) SetWidth(v string) {
	o.Width = &v
}

func (o MemoryAbstractUnit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedEquipmentBase, errEquipmentBase := json.Marshal(o.EquipmentBase)
	if errEquipmentBase != nil {
		return []byte{}, errEquipmentBase
	}
	errEquipmentBase = json.Unmarshal([]byte(serializedEquipmentBase), &toSerialize)
	if errEquipmentBase != nil {
		return []byte{}, errEquipmentBase
	}
	if o.AdminState != nil {
		toSerialize["AdminState"] = o.AdminState
	}
	if o.ArrayId != nil {
		toSerialize["ArrayId"] = o.ArrayId
	}
	if o.Bank != nil {
		toSerialize["Bank"] = o.Bank
	}
	if o.Capacity != nil {
		toSerialize["Capacity"] = o.Capacity
	}
	if o.Clock != nil {
		toSerialize["Clock"] = o.Clock
	}
	if o.FormFactor != nil {
		toSerialize["FormFactor"] = o.FormFactor
	}
	if o.Latency != nil {
		toSerialize["Latency"] = o.Latency
	}
	if o.Location != nil {
		toSerialize["Location"] = o.Location
	}
	if o.OperPowerState != nil {
		toSerialize["OperPowerState"] = o.OperPowerState
	}
	if o.OperState != nil {
		toSerialize["OperState"] = o.OperState
	}
	if o.Operability != nil {
		toSerialize["Operability"] = o.Operability
	}
	if o.Presence != nil {
		toSerialize["Presence"] = o.Presence
	}
	if o.Set != nil {
		toSerialize["Set"] = o.Set
	}
	if o.Speed != nil {
		toSerialize["Speed"] = o.Speed
	}
	if o.Thermal != nil {
		toSerialize["Thermal"] = o.Thermal
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}
	if o.Visibility != nil {
		toSerialize["Visibility"] = o.Visibility
	}
	if o.Width != nil {
		toSerialize["Width"] = o.Width
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MemoryAbstractUnit) UnmarshalJSON(bytes []byte) (err error) {
	type MemoryAbstractUnitWithoutEmbeddedStruct struct {
		// This represents the administrative state of the memory unit on a server.
		AdminState *string `json:"AdminState,omitempty"`
		// This represents the memory array to which the memory unit belongs to.
		ArrayId *int64 `json:"ArrayId,omitempty"`
		// This represents the memory bank of the memory unit on a server.
		Bank *int64 `json:"Bank,omitempty"`
		// This represents the memory capacity in MiB of the memory unit on a server.
		Capacity *string `json:"Capacity,omitempty"`
		// This represents the clock of the memory unit on a server.
		Clock *string `json:"Clock,omitempty"`
		// This represents the form factor of the memory unit on a server.
		FormFactor *string `json:"FormFactor,omitempty"`
		// This represents the latency of the memory unit on a server.
		Latency *string `json:"Latency,omitempty"`
		// This represents the location of the memory unit on a server.
		Location *string `json:"Location,omitempty"`
		// This represents the operational power state of the memory unit on a server.
		OperPowerState *string `json:"OperPowerState,omitempty"`
		// This represents the operational state of the memory unit on a server.
		OperState *string `json:"OperState,omitempty"`
		// This represents the operability of the memory unit on a server.
		Operability *string `json:"Operability,omitempty"`
		// This represents the presence state of the memory unit on a server.
		Presence *string `json:"Presence,omitempty"`
		// This represents the set of the memory unit on a server.
		Set *int64 `json:"Set,omitempty"`
		// This represents the speed of the memory unit on a server.
		Speed *string `json:"Speed,omitempty"`
		// This represents the thremal state of the memory unit on a server.
		Thermal *string `json:"Thermal,omitempty"`
		// This represents the memory type of the memory unit on a server.
		Type *string `json:"Type,omitempty"`
		// This represents the visibility of the memory unit on a server.
		Visibility *string `json:"Visibility,omitempty"`
		// This represents the width of the memory unit on a server.
		Width *string `json:"Width,omitempty"`
	}

	varMemoryAbstractUnitWithoutEmbeddedStruct := MemoryAbstractUnitWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varMemoryAbstractUnitWithoutEmbeddedStruct)
	if err == nil {
		varMemoryAbstractUnit := _MemoryAbstractUnit{}
		varMemoryAbstractUnit.AdminState = varMemoryAbstractUnitWithoutEmbeddedStruct.AdminState
		varMemoryAbstractUnit.ArrayId = varMemoryAbstractUnitWithoutEmbeddedStruct.ArrayId
		varMemoryAbstractUnit.Bank = varMemoryAbstractUnitWithoutEmbeddedStruct.Bank
		varMemoryAbstractUnit.Capacity = varMemoryAbstractUnitWithoutEmbeddedStruct.Capacity
		varMemoryAbstractUnit.Clock = varMemoryAbstractUnitWithoutEmbeddedStruct.Clock
		varMemoryAbstractUnit.FormFactor = varMemoryAbstractUnitWithoutEmbeddedStruct.FormFactor
		varMemoryAbstractUnit.Latency = varMemoryAbstractUnitWithoutEmbeddedStruct.Latency
		varMemoryAbstractUnit.Location = varMemoryAbstractUnitWithoutEmbeddedStruct.Location
		varMemoryAbstractUnit.OperPowerState = varMemoryAbstractUnitWithoutEmbeddedStruct.OperPowerState
		varMemoryAbstractUnit.OperState = varMemoryAbstractUnitWithoutEmbeddedStruct.OperState
		varMemoryAbstractUnit.Operability = varMemoryAbstractUnitWithoutEmbeddedStruct.Operability
		varMemoryAbstractUnit.Presence = varMemoryAbstractUnitWithoutEmbeddedStruct.Presence
		varMemoryAbstractUnit.Set = varMemoryAbstractUnitWithoutEmbeddedStruct.Set
		varMemoryAbstractUnit.Speed = varMemoryAbstractUnitWithoutEmbeddedStruct.Speed
		varMemoryAbstractUnit.Thermal = varMemoryAbstractUnitWithoutEmbeddedStruct.Thermal
		varMemoryAbstractUnit.Type = varMemoryAbstractUnitWithoutEmbeddedStruct.Type
		varMemoryAbstractUnit.Visibility = varMemoryAbstractUnitWithoutEmbeddedStruct.Visibility
		varMemoryAbstractUnit.Width = varMemoryAbstractUnitWithoutEmbeddedStruct.Width
		*o = MemoryAbstractUnit(varMemoryAbstractUnit)
	} else {
		return err
	}

	varMemoryAbstractUnit := _MemoryAbstractUnit{}

	err = json.Unmarshal(bytes, &varMemoryAbstractUnit)
	if err == nil {
		o.EquipmentBase = varMemoryAbstractUnit.EquipmentBase
	} else {
		return err
	}
	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "AdminState")
		delete(additionalProperties, "ArrayId")
		delete(additionalProperties, "Bank")
		delete(additionalProperties, "Capacity")
		delete(additionalProperties, "Clock")
		delete(additionalProperties, "FormFactor")
		delete(additionalProperties, "Latency")
		delete(additionalProperties, "Location")
		delete(additionalProperties, "OperPowerState")
		delete(additionalProperties, "OperState")
		delete(additionalProperties, "Operability")
		delete(additionalProperties, "Presence")
		delete(additionalProperties, "Set")
		delete(additionalProperties, "Speed")
		delete(additionalProperties, "Thermal")
		delete(additionalProperties, "Type")
		delete(additionalProperties, "Visibility")
		delete(additionalProperties, "Width")

		// remove fields from embedded structs
		reflectEquipmentBase := reflect.ValueOf(o.EquipmentBase)
		for i := 0; i < reflectEquipmentBase.Type().NumField(); i++ {
			t := reflectEquipmentBase.Type().Field(i)

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

type NullableMemoryAbstractUnit struct {
	value *MemoryAbstractUnit
	isSet bool
}

func (v NullableMemoryAbstractUnit) Get() *MemoryAbstractUnit {
	return v.value
}

func (v *NullableMemoryAbstractUnit) Set(val *MemoryAbstractUnit) {
	v.value = val
	v.isSet = true
}

func (v NullableMemoryAbstractUnit) IsSet() bool {
	return v.isSet
}

func (v *NullableMemoryAbstractUnit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMemoryAbstractUnit(val *MemoryAbstractUnit) *NullableMemoryAbstractUnit {
	return &NullableMemoryAbstractUnit{value: val, isSet: true}
}

func (v NullableMemoryAbstractUnit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMemoryAbstractUnit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
