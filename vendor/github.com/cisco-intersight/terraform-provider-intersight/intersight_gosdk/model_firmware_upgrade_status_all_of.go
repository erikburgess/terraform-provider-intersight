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

// FirmwareUpgradeStatusAllOf Definition of the list of properties defined in 'firmware.UpgradeStatus', excluding properties defined in parent classes.
type FirmwareUpgradeStatusAllOf struct {
	// The error message from the endpoint during the download.
	DownloadError *string `json:"DownloadError,omitempty"`
	// The message from the endpoint during the download.} type: string
	DownloadMessage interface{} `json:"DownloadMessage,omitempty"`
	// The percentage of the image downloaded in the endpoint.
	DownloadPercentage *int64 `json:"DownloadPercentage,omitempty"`
	// The download progress of the file represented as a percentage between 0% and 100%. If progress reporting is not possible a value of -1 is sent.
	DownloadProgress *int64 `json:"DownloadProgress,omitempty"`
	// The number of retries the plugin attempted before succeeding or failing the download.
	DownloadRetries *int64 `json:"DownloadRetries,omitempty"`
	// The image download stages. Example:downloading, flashing.
	DownloadStage *string `json:"DownloadStage,omitempty"`
	// The server power status after the upgrade request is submitted in the endpoint.
	EpPowerStatus *string `json:"EpPowerStatus,omitempty"`
	// The reason for the operation failure.
	OverallError *string `json:"OverallError,omitempty"`
	// The overall percentage of the operation.
	OverallPercentage *int64 `json:"OverallPercentage,omitempty"`
	// The overall status of the operation.
	Overallstatus *string `json:"Overallstatus,omitempty"`
	// Pending reason for the upgrade waiting.
	PendingType *string `json:"PendingType,omitempty"`
	// The sha256checksum of the downloaded file as calculated by the download plugin after successfully downloading a file.
	Sha256checksum *string                           `json:"Sha256checksum,omitempty"`
	Upgrade        *FirmwareUpgradeBaseRelationship  `json:"Upgrade,omitempty"`
	Workflow       *WorkflowWorkflowInfoRelationship `json:"Workflow,omitempty"`
}

// NewFirmwareUpgradeStatusAllOf instantiates a new FirmwareUpgradeStatusAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirmwareUpgradeStatusAllOf() *FirmwareUpgradeStatusAllOf {
	this := FirmwareUpgradeStatusAllOf{}
	var epPowerStatus string = "none"
	this.EpPowerStatus = &epPowerStatus
	var overallstatus string = "none"
	this.Overallstatus = &overallstatus
	var pendingType string = "none"
	this.PendingType = &pendingType
	return &this
}

// NewFirmwareUpgradeStatusAllOfWithDefaults instantiates a new FirmwareUpgradeStatusAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirmwareUpgradeStatusAllOfWithDefaults() *FirmwareUpgradeStatusAllOf {
	this := FirmwareUpgradeStatusAllOf{}
	var epPowerStatus string = "none"
	this.EpPowerStatus = &epPowerStatus
	var overallstatus string = "none"
	this.Overallstatus = &overallstatus
	var pendingType string = "none"
	this.PendingType = &pendingType
	return &this
}

// GetDownloadError returns the DownloadError field value if set, zero value otherwise.
func (o *FirmwareUpgradeStatusAllOf) GetDownloadError() string {
	if o == nil || o.DownloadError == nil {
		var ret string
		return ret
	}
	return *o.DownloadError
}

// GetDownloadErrorOk returns a tuple with the DownloadError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareUpgradeStatusAllOf) GetDownloadErrorOk() (*string, bool) {
	if o == nil || o.DownloadError == nil {
		return nil, false
	}
	return o.DownloadError, true
}

// HasDownloadError returns a boolean if a field has been set.
func (o *FirmwareUpgradeStatusAllOf) HasDownloadError() bool {
	if o != nil && o.DownloadError != nil {
		return true
	}

	return false
}

// SetDownloadError gets a reference to the given string and assigns it to the DownloadError field.
func (o *FirmwareUpgradeStatusAllOf) SetDownloadError(v string) {
	o.DownloadError = &v
}

// GetDownloadMessage returns the DownloadMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirmwareUpgradeStatusAllOf) GetDownloadMessage() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.DownloadMessage
}

// GetDownloadMessageOk returns a tuple with the DownloadMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirmwareUpgradeStatusAllOf) GetDownloadMessageOk() (*interface{}, bool) {
	if o == nil || o.DownloadMessage == nil {
		return nil, false
	}
	return &o.DownloadMessage, true
}

// HasDownloadMessage returns a boolean if a field has been set.
func (o *FirmwareUpgradeStatusAllOf) HasDownloadMessage() bool {
	if o != nil && o.DownloadMessage != nil {
		return true
	}

	return false
}

// SetDownloadMessage gets a reference to the given interface{} and assigns it to the DownloadMessage field.
func (o *FirmwareUpgradeStatusAllOf) SetDownloadMessage(v interface{}) {
	o.DownloadMessage = v
}

// GetDownloadPercentage returns the DownloadPercentage field value if set, zero value otherwise.
func (o *FirmwareUpgradeStatusAllOf) GetDownloadPercentage() int64 {
	if o == nil || o.DownloadPercentage == nil {
		var ret int64
		return ret
	}
	return *o.DownloadPercentage
}

// GetDownloadPercentageOk returns a tuple with the DownloadPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareUpgradeStatusAllOf) GetDownloadPercentageOk() (*int64, bool) {
	if o == nil || o.DownloadPercentage == nil {
		return nil, false
	}
	return o.DownloadPercentage, true
}

// HasDownloadPercentage returns a boolean if a field has been set.
func (o *FirmwareUpgradeStatusAllOf) HasDownloadPercentage() bool {
	if o != nil && o.DownloadPercentage != nil {
		return true
	}

	return false
}

// SetDownloadPercentage gets a reference to the given int64 and assigns it to the DownloadPercentage field.
func (o *FirmwareUpgradeStatusAllOf) SetDownloadPercentage(v int64) {
	o.DownloadPercentage = &v
}

// GetDownloadProgress returns the DownloadProgress field value if set, zero value otherwise.
func (o *FirmwareUpgradeStatusAllOf) GetDownloadProgress() int64 {
	if o == nil || o.DownloadProgress == nil {
		var ret int64
		return ret
	}
	return *o.DownloadProgress
}

// GetDownloadProgressOk returns a tuple with the DownloadProgress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareUpgradeStatusAllOf) GetDownloadProgressOk() (*int64, bool) {
	if o == nil || o.DownloadProgress == nil {
		return nil, false
	}
	return o.DownloadProgress, true
}

// HasDownloadProgress returns a boolean if a field has been set.
func (o *FirmwareUpgradeStatusAllOf) HasDownloadProgress() bool {
	if o != nil && o.DownloadProgress != nil {
		return true
	}

	return false
}

// SetDownloadProgress gets a reference to the given int64 and assigns it to the DownloadProgress field.
func (o *FirmwareUpgradeStatusAllOf) SetDownloadProgress(v int64) {
	o.DownloadProgress = &v
}

// GetDownloadRetries returns the DownloadRetries field value if set, zero value otherwise.
func (o *FirmwareUpgradeStatusAllOf) GetDownloadRetries() int64 {
	if o == nil || o.DownloadRetries == nil {
		var ret int64
		return ret
	}
	return *o.DownloadRetries
}

// GetDownloadRetriesOk returns a tuple with the DownloadRetries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareUpgradeStatusAllOf) GetDownloadRetriesOk() (*int64, bool) {
	if o == nil || o.DownloadRetries == nil {
		return nil, false
	}
	return o.DownloadRetries, true
}

// HasDownloadRetries returns a boolean if a field has been set.
func (o *FirmwareUpgradeStatusAllOf) HasDownloadRetries() bool {
	if o != nil && o.DownloadRetries != nil {
		return true
	}

	return false
}

// SetDownloadRetries gets a reference to the given int64 and assigns it to the DownloadRetries field.
func (o *FirmwareUpgradeStatusAllOf) SetDownloadRetries(v int64) {
	o.DownloadRetries = &v
}

// GetDownloadStage returns the DownloadStage field value if set, zero value otherwise.
func (o *FirmwareUpgradeStatusAllOf) GetDownloadStage() string {
	if o == nil || o.DownloadStage == nil {
		var ret string
		return ret
	}
	return *o.DownloadStage
}

// GetDownloadStageOk returns a tuple with the DownloadStage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareUpgradeStatusAllOf) GetDownloadStageOk() (*string, bool) {
	if o == nil || o.DownloadStage == nil {
		return nil, false
	}
	return o.DownloadStage, true
}

// HasDownloadStage returns a boolean if a field has been set.
func (o *FirmwareUpgradeStatusAllOf) HasDownloadStage() bool {
	if o != nil && o.DownloadStage != nil {
		return true
	}

	return false
}

// SetDownloadStage gets a reference to the given string and assigns it to the DownloadStage field.
func (o *FirmwareUpgradeStatusAllOf) SetDownloadStage(v string) {
	o.DownloadStage = &v
}

// GetEpPowerStatus returns the EpPowerStatus field value if set, zero value otherwise.
func (o *FirmwareUpgradeStatusAllOf) GetEpPowerStatus() string {
	if o == nil || o.EpPowerStatus == nil {
		var ret string
		return ret
	}
	return *o.EpPowerStatus
}

// GetEpPowerStatusOk returns a tuple with the EpPowerStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareUpgradeStatusAllOf) GetEpPowerStatusOk() (*string, bool) {
	if o == nil || o.EpPowerStatus == nil {
		return nil, false
	}
	return o.EpPowerStatus, true
}

// HasEpPowerStatus returns a boolean if a field has been set.
func (o *FirmwareUpgradeStatusAllOf) HasEpPowerStatus() bool {
	if o != nil && o.EpPowerStatus != nil {
		return true
	}

	return false
}

// SetEpPowerStatus gets a reference to the given string and assigns it to the EpPowerStatus field.
func (o *FirmwareUpgradeStatusAllOf) SetEpPowerStatus(v string) {
	o.EpPowerStatus = &v
}

// GetOverallError returns the OverallError field value if set, zero value otherwise.
func (o *FirmwareUpgradeStatusAllOf) GetOverallError() string {
	if o == nil || o.OverallError == nil {
		var ret string
		return ret
	}
	return *o.OverallError
}

// GetOverallErrorOk returns a tuple with the OverallError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareUpgradeStatusAllOf) GetOverallErrorOk() (*string, bool) {
	if o == nil || o.OverallError == nil {
		return nil, false
	}
	return o.OverallError, true
}

// HasOverallError returns a boolean if a field has been set.
func (o *FirmwareUpgradeStatusAllOf) HasOverallError() bool {
	if o != nil && o.OverallError != nil {
		return true
	}

	return false
}

// SetOverallError gets a reference to the given string and assigns it to the OverallError field.
func (o *FirmwareUpgradeStatusAllOf) SetOverallError(v string) {
	o.OverallError = &v
}

// GetOverallPercentage returns the OverallPercentage field value if set, zero value otherwise.
func (o *FirmwareUpgradeStatusAllOf) GetOverallPercentage() int64 {
	if o == nil || o.OverallPercentage == nil {
		var ret int64
		return ret
	}
	return *o.OverallPercentage
}

// GetOverallPercentageOk returns a tuple with the OverallPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareUpgradeStatusAllOf) GetOverallPercentageOk() (*int64, bool) {
	if o == nil || o.OverallPercentage == nil {
		return nil, false
	}
	return o.OverallPercentage, true
}

// HasOverallPercentage returns a boolean if a field has been set.
func (o *FirmwareUpgradeStatusAllOf) HasOverallPercentage() bool {
	if o != nil && o.OverallPercentage != nil {
		return true
	}

	return false
}

// SetOverallPercentage gets a reference to the given int64 and assigns it to the OverallPercentage field.
func (o *FirmwareUpgradeStatusAllOf) SetOverallPercentage(v int64) {
	o.OverallPercentage = &v
}

// GetOverallstatus returns the Overallstatus field value if set, zero value otherwise.
func (o *FirmwareUpgradeStatusAllOf) GetOverallstatus() string {
	if o == nil || o.Overallstatus == nil {
		var ret string
		return ret
	}
	return *o.Overallstatus
}

// GetOverallstatusOk returns a tuple with the Overallstatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareUpgradeStatusAllOf) GetOverallstatusOk() (*string, bool) {
	if o == nil || o.Overallstatus == nil {
		return nil, false
	}
	return o.Overallstatus, true
}

// HasOverallstatus returns a boolean if a field has been set.
func (o *FirmwareUpgradeStatusAllOf) HasOverallstatus() bool {
	if o != nil && o.Overallstatus != nil {
		return true
	}

	return false
}

// SetOverallstatus gets a reference to the given string and assigns it to the Overallstatus field.
func (o *FirmwareUpgradeStatusAllOf) SetOverallstatus(v string) {
	o.Overallstatus = &v
}

// GetPendingType returns the PendingType field value if set, zero value otherwise.
func (o *FirmwareUpgradeStatusAllOf) GetPendingType() string {
	if o == nil || o.PendingType == nil {
		var ret string
		return ret
	}
	return *o.PendingType
}

// GetPendingTypeOk returns a tuple with the PendingType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareUpgradeStatusAllOf) GetPendingTypeOk() (*string, bool) {
	if o == nil || o.PendingType == nil {
		return nil, false
	}
	return o.PendingType, true
}

// HasPendingType returns a boolean if a field has been set.
func (o *FirmwareUpgradeStatusAllOf) HasPendingType() bool {
	if o != nil && o.PendingType != nil {
		return true
	}

	return false
}

// SetPendingType gets a reference to the given string and assigns it to the PendingType field.
func (o *FirmwareUpgradeStatusAllOf) SetPendingType(v string) {
	o.PendingType = &v
}

// GetSha256checksum returns the Sha256checksum field value if set, zero value otherwise.
func (o *FirmwareUpgradeStatusAllOf) GetSha256checksum() string {
	if o == nil || o.Sha256checksum == nil {
		var ret string
		return ret
	}
	return *o.Sha256checksum
}

// GetSha256checksumOk returns a tuple with the Sha256checksum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareUpgradeStatusAllOf) GetSha256checksumOk() (*string, bool) {
	if o == nil || o.Sha256checksum == nil {
		return nil, false
	}
	return o.Sha256checksum, true
}

// HasSha256checksum returns a boolean if a field has been set.
func (o *FirmwareUpgradeStatusAllOf) HasSha256checksum() bool {
	if o != nil && o.Sha256checksum != nil {
		return true
	}

	return false
}

// SetSha256checksum gets a reference to the given string and assigns it to the Sha256checksum field.
func (o *FirmwareUpgradeStatusAllOf) SetSha256checksum(v string) {
	o.Sha256checksum = &v
}

// GetUpgrade returns the Upgrade field value if set, zero value otherwise.
func (o *FirmwareUpgradeStatusAllOf) GetUpgrade() FirmwareUpgradeBaseRelationship {
	if o == nil || o.Upgrade == nil {
		var ret FirmwareUpgradeBaseRelationship
		return ret
	}
	return *o.Upgrade
}

// GetUpgradeOk returns a tuple with the Upgrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareUpgradeStatusAllOf) GetUpgradeOk() (*FirmwareUpgradeBaseRelationship, bool) {
	if o == nil || o.Upgrade == nil {
		return nil, false
	}
	return o.Upgrade, true
}

// HasUpgrade returns a boolean if a field has been set.
func (o *FirmwareUpgradeStatusAllOf) HasUpgrade() bool {
	if o != nil && o.Upgrade != nil {
		return true
	}

	return false
}

// SetUpgrade gets a reference to the given FirmwareUpgradeBaseRelationship and assigns it to the Upgrade field.
func (o *FirmwareUpgradeStatusAllOf) SetUpgrade(v FirmwareUpgradeBaseRelationship) {
	o.Upgrade = &v
}

// GetWorkflow returns the Workflow field value if set, zero value otherwise.
func (o *FirmwareUpgradeStatusAllOf) GetWorkflow() WorkflowWorkflowInfoRelationship {
	if o == nil || o.Workflow == nil {
		var ret WorkflowWorkflowInfoRelationship
		return ret
	}
	return *o.Workflow
}

// GetWorkflowOk returns a tuple with the Workflow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareUpgradeStatusAllOf) GetWorkflowOk() (*WorkflowWorkflowInfoRelationship, bool) {
	if o == nil || o.Workflow == nil {
		return nil, false
	}
	return o.Workflow, true
}

// HasWorkflow returns a boolean if a field has been set.
func (o *FirmwareUpgradeStatusAllOf) HasWorkflow() bool {
	if o != nil && o.Workflow != nil {
		return true
	}

	return false
}

// SetWorkflow gets a reference to the given WorkflowWorkflowInfoRelationship and assigns it to the Workflow field.
func (o *FirmwareUpgradeStatusAllOf) SetWorkflow(v WorkflowWorkflowInfoRelationship) {
	o.Workflow = &v
}

func (o FirmwareUpgradeStatusAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DownloadError != nil {
		toSerialize["DownloadError"] = o.DownloadError
	}
	if o.DownloadMessage != nil {
		toSerialize["DownloadMessage"] = o.DownloadMessage
	}
	if o.DownloadPercentage != nil {
		toSerialize["DownloadPercentage"] = o.DownloadPercentage
	}
	if o.DownloadProgress != nil {
		toSerialize["DownloadProgress"] = o.DownloadProgress
	}
	if o.DownloadRetries != nil {
		toSerialize["DownloadRetries"] = o.DownloadRetries
	}
	if o.DownloadStage != nil {
		toSerialize["DownloadStage"] = o.DownloadStage
	}
	if o.EpPowerStatus != nil {
		toSerialize["EpPowerStatus"] = o.EpPowerStatus
	}
	if o.OverallError != nil {
		toSerialize["OverallError"] = o.OverallError
	}
	if o.OverallPercentage != nil {
		toSerialize["OverallPercentage"] = o.OverallPercentage
	}
	if o.Overallstatus != nil {
		toSerialize["Overallstatus"] = o.Overallstatus
	}
	if o.PendingType != nil {
		toSerialize["PendingType"] = o.PendingType
	}
	if o.Sha256checksum != nil {
		toSerialize["Sha256checksum"] = o.Sha256checksum
	}
	if o.Upgrade != nil {
		toSerialize["Upgrade"] = o.Upgrade
	}
	if o.Workflow != nil {
		toSerialize["Workflow"] = o.Workflow
	}
	return json.Marshal(toSerialize)
}

type NullableFirmwareUpgradeStatusAllOf struct {
	value *FirmwareUpgradeStatusAllOf
	isSet bool
}

func (v NullableFirmwareUpgradeStatusAllOf) Get() *FirmwareUpgradeStatusAllOf {
	return v.value
}

func (v *NullableFirmwareUpgradeStatusAllOf) Set(val *FirmwareUpgradeStatusAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFirmwareUpgradeStatusAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFirmwareUpgradeStatusAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirmwareUpgradeStatusAllOf(val *FirmwareUpgradeStatusAllOf) *NullableFirmwareUpgradeStatusAllOf {
	return &NullableFirmwareUpgradeStatusAllOf{value: val, isSet: true}
}

func (v NullableFirmwareUpgradeStatusAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirmwareUpgradeStatusAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
