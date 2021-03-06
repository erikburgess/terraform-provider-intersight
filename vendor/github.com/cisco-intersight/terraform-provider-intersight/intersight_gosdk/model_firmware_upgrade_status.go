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

// FirmwareUpgradeStatus The status for the upgrade operation to include the status for the download and upgrade stages.
type FirmwareUpgradeStatus struct {
	MoBaseMo
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
	Sha256checksum       *string                           `json:"Sha256checksum,omitempty"`
	Upgrade              *FirmwareUpgradeBaseRelationship  `json:"Upgrade,omitempty"`
	Workflow             *WorkflowWorkflowInfoRelationship `json:"Workflow,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FirmwareUpgradeStatus FirmwareUpgradeStatus

// NewFirmwareUpgradeStatus instantiates a new FirmwareUpgradeStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirmwareUpgradeStatus() *FirmwareUpgradeStatus {
	this := FirmwareUpgradeStatus{}
	var epPowerStatus string = "none"
	this.EpPowerStatus = &epPowerStatus
	var overallstatus string = "none"
	this.Overallstatus = &overallstatus
	var pendingType string = "none"
	this.PendingType = &pendingType
	return &this
}

// NewFirmwareUpgradeStatusWithDefaults instantiates a new FirmwareUpgradeStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirmwareUpgradeStatusWithDefaults() *FirmwareUpgradeStatus {
	this := FirmwareUpgradeStatus{}
	var epPowerStatus string = "none"
	this.EpPowerStatus = &epPowerStatus
	var overallstatus string = "none"
	this.Overallstatus = &overallstatus
	var pendingType string = "none"
	this.PendingType = &pendingType
	return &this
}

// GetDownloadError returns the DownloadError field value if set, zero value otherwise.
func (o *FirmwareUpgradeStatus) GetDownloadError() string {
	if o == nil || o.DownloadError == nil {
		var ret string
		return ret
	}
	return *o.DownloadError
}

// GetDownloadErrorOk returns a tuple with the DownloadError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareUpgradeStatus) GetDownloadErrorOk() (*string, bool) {
	if o == nil || o.DownloadError == nil {
		return nil, false
	}
	return o.DownloadError, true
}

// HasDownloadError returns a boolean if a field has been set.
func (o *FirmwareUpgradeStatus) HasDownloadError() bool {
	if o != nil && o.DownloadError != nil {
		return true
	}

	return false
}

// SetDownloadError gets a reference to the given string and assigns it to the DownloadError field.
func (o *FirmwareUpgradeStatus) SetDownloadError(v string) {
	o.DownloadError = &v
}

// GetDownloadMessage returns the DownloadMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirmwareUpgradeStatus) GetDownloadMessage() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.DownloadMessage
}

// GetDownloadMessageOk returns a tuple with the DownloadMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirmwareUpgradeStatus) GetDownloadMessageOk() (*interface{}, bool) {
	if o == nil || o.DownloadMessage == nil {
		return nil, false
	}
	return &o.DownloadMessage, true
}

// HasDownloadMessage returns a boolean if a field has been set.
func (o *FirmwareUpgradeStatus) HasDownloadMessage() bool {
	if o != nil && o.DownloadMessage != nil {
		return true
	}

	return false
}

// SetDownloadMessage gets a reference to the given interface{} and assigns it to the DownloadMessage field.
func (o *FirmwareUpgradeStatus) SetDownloadMessage(v interface{}) {
	o.DownloadMessage = v
}

// GetDownloadPercentage returns the DownloadPercentage field value if set, zero value otherwise.
func (o *FirmwareUpgradeStatus) GetDownloadPercentage() int64 {
	if o == nil || o.DownloadPercentage == nil {
		var ret int64
		return ret
	}
	return *o.DownloadPercentage
}

// GetDownloadPercentageOk returns a tuple with the DownloadPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareUpgradeStatus) GetDownloadPercentageOk() (*int64, bool) {
	if o == nil || o.DownloadPercentage == nil {
		return nil, false
	}
	return o.DownloadPercentage, true
}

// HasDownloadPercentage returns a boolean if a field has been set.
func (o *FirmwareUpgradeStatus) HasDownloadPercentage() bool {
	if o != nil && o.DownloadPercentage != nil {
		return true
	}

	return false
}

// SetDownloadPercentage gets a reference to the given int64 and assigns it to the DownloadPercentage field.
func (o *FirmwareUpgradeStatus) SetDownloadPercentage(v int64) {
	o.DownloadPercentage = &v
}

// GetDownloadProgress returns the DownloadProgress field value if set, zero value otherwise.
func (o *FirmwareUpgradeStatus) GetDownloadProgress() int64 {
	if o == nil || o.DownloadProgress == nil {
		var ret int64
		return ret
	}
	return *o.DownloadProgress
}

// GetDownloadProgressOk returns a tuple with the DownloadProgress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareUpgradeStatus) GetDownloadProgressOk() (*int64, bool) {
	if o == nil || o.DownloadProgress == nil {
		return nil, false
	}
	return o.DownloadProgress, true
}

// HasDownloadProgress returns a boolean if a field has been set.
func (o *FirmwareUpgradeStatus) HasDownloadProgress() bool {
	if o != nil && o.DownloadProgress != nil {
		return true
	}

	return false
}

// SetDownloadProgress gets a reference to the given int64 and assigns it to the DownloadProgress field.
func (o *FirmwareUpgradeStatus) SetDownloadProgress(v int64) {
	o.DownloadProgress = &v
}

// GetDownloadRetries returns the DownloadRetries field value if set, zero value otherwise.
func (o *FirmwareUpgradeStatus) GetDownloadRetries() int64 {
	if o == nil || o.DownloadRetries == nil {
		var ret int64
		return ret
	}
	return *o.DownloadRetries
}

// GetDownloadRetriesOk returns a tuple with the DownloadRetries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareUpgradeStatus) GetDownloadRetriesOk() (*int64, bool) {
	if o == nil || o.DownloadRetries == nil {
		return nil, false
	}
	return o.DownloadRetries, true
}

// HasDownloadRetries returns a boolean if a field has been set.
func (o *FirmwareUpgradeStatus) HasDownloadRetries() bool {
	if o != nil && o.DownloadRetries != nil {
		return true
	}

	return false
}

// SetDownloadRetries gets a reference to the given int64 and assigns it to the DownloadRetries field.
func (o *FirmwareUpgradeStatus) SetDownloadRetries(v int64) {
	o.DownloadRetries = &v
}

// GetDownloadStage returns the DownloadStage field value if set, zero value otherwise.
func (o *FirmwareUpgradeStatus) GetDownloadStage() string {
	if o == nil || o.DownloadStage == nil {
		var ret string
		return ret
	}
	return *o.DownloadStage
}

// GetDownloadStageOk returns a tuple with the DownloadStage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareUpgradeStatus) GetDownloadStageOk() (*string, bool) {
	if o == nil || o.DownloadStage == nil {
		return nil, false
	}
	return o.DownloadStage, true
}

// HasDownloadStage returns a boolean if a field has been set.
func (o *FirmwareUpgradeStatus) HasDownloadStage() bool {
	if o != nil && o.DownloadStage != nil {
		return true
	}

	return false
}

// SetDownloadStage gets a reference to the given string and assigns it to the DownloadStage field.
func (o *FirmwareUpgradeStatus) SetDownloadStage(v string) {
	o.DownloadStage = &v
}

// GetEpPowerStatus returns the EpPowerStatus field value if set, zero value otherwise.
func (o *FirmwareUpgradeStatus) GetEpPowerStatus() string {
	if o == nil || o.EpPowerStatus == nil {
		var ret string
		return ret
	}
	return *o.EpPowerStatus
}

// GetEpPowerStatusOk returns a tuple with the EpPowerStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareUpgradeStatus) GetEpPowerStatusOk() (*string, bool) {
	if o == nil || o.EpPowerStatus == nil {
		return nil, false
	}
	return o.EpPowerStatus, true
}

// HasEpPowerStatus returns a boolean if a field has been set.
func (o *FirmwareUpgradeStatus) HasEpPowerStatus() bool {
	if o != nil && o.EpPowerStatus != nil {
		return true
	}

	return false
}

// SetEpPowerStatus gets a reference to the given string and assigns it to the EpPowerStatus field.
func (o *FirmwareUpgradeStatus) SetEpPowerStatus(v string) {
	o.EpPowerStatus = &v
}

// GetOverallError returns the OverallError field value if set, zero value otherwise.
func (o *FirmwareUpgradeStatus) GetOverallError() string {
	if o == nil || o.OverallError == nil {
		var ret string
		return ret
	}
	return *o.OverallError
}

// GetOverallErrorOk returns a tuple with the OverallError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareUpgradeStatus) GetOverallErrorOk() (*string, bool) {
	if o == nil || o.OverallError == nil {
		return nil, false
	}
	return o.OverallError, true
}

// HasOverallError returns a boolean if a field has been set.
func (o *FirmwareUpgradeStatus) HasOverallError() bool {
	if o != nil && o.OverallError != nil {
		return true
	}

	return false
}

// SetOverallError gets a reference to the given string and assigns it to the OverallError field.
func (o *FirmwareUpgradeStatus) SetOverallError(v string) {
	o.OverallError = &v
}

// GetOverallPercentage returns the OverallPercentage field value if set, zero value otherwise.
func (o *FirmwareUpgradeStatus) GetOverallPercentage() int64 {
	if o == nil || o.OverallPercentage == nil {
		var ret int64
		return ret
	}
	return *o.OverallPercentage
}

// GetOverallPercentageOk returns a tuple with the OverallPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareUpgradeStatus) GetOverallPercentageOk() (*int64, bool) {
	if o == nil || o.OverallPercentage == nil {
		return nil, false
	}
	return o.OverallPercentage, true
}

// HasOverallPercentage returns a boolean if a field has been set.
func (o *FirmwareUpgradeStatus) HasOverallPercentage() bool {
	if o != nil && o.OverallPercentage != nil {
		return true
	}

	return false
}

// SetOverallPercentage gets a reference to the given int64 and assigns it to the OverallPercentage field.
func (o *FirmwareUpgradeStatus) SetOverallPercentage(v int64) {
	o.OverallPercentage = &v
}

// GetOverallstatus returns the Overallstatus field value if set, zero value otherwise.
func (o *FirmwareUpgradeStatus) GetOverallstatus() string {
	if o == nil || o.Overallstatus == nil {
		var ret string
		return ret
	}
	return *o.Overallstatus
}

// GetOverallstatusOk returns a tuple with the Overallstatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareUpgradeStatus) GetOverallstatusOk() (*string, bool) {
	if o == nil || o.Overallstatus == nil {
		return nil, false
	}
	return o.Overallstatus, true
}

// HasOverallstatus returns a boolean if a field has been set.
func (o *FirmwareUpgradeStatus) HasOverallstatus() bool {
	if o != nil && o.Overallstatus != nil {
		return true
	}

	return false
}

// SetOverallstatus gets a reference to the given string and assigns it to the Overallstatus field.
func (o *FirmwareUpgradeStatus) SetOverallstatus(v string) {
	o.Overallstatus = &v
}

// GetPendingType returns the PendingType field value if set, zero value otherwise.
func (o *FirmwareUpgradeStatus) GetPendingType() string {
	if o == nil || o.PendingType == nil {
		var ret string
		return ret
	}
	return *o.PendingType
}

// GetPendingTypeOk returns a tuple with the PendingType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareUpgradeStatus) GetPendingTypeOk() (*string, bool) {
	if o == nil || o.PendingType == nil {
		return nil, false
	}
	return o.PendingType, true
}

// HasPendingType returns a boolean if a field has been set.
func (o *FirmwareUpgradeStatus) HasPendingType() bool {
	if o != nil && o.PendingType != nil {
		return true
	}

	return false
}

// SetPendingType gets a reference to the given string and assigns it to the PendingType field.
func (o *FirmwareUpgradeStatus) SetPendingType(v string) {
	o.PendingType = &v
}

// GetSha256checksum returns the Sha256checksum field value if set, zero value otherwise.
func (o *FirmwareUpgradeStatus) GetSha256checksum() string {
	if o == nil || o.Sha256checksum == nil {
		var ret string
		return ret
	}
	return *o.Sha256checksum
}

// GetSha256checksumOk returns a tuple with the Sha256checksum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareUpgradeStatus) GetSha256checksumOk() (*string, bool) {
	if o == nil || o.Sha256checksum == nil {
		return nil, false
	}
	return o.Sha256checksum, true
}

// HasSha256checksum returns a boolean if a field has been set.
func (o *FirmwareUpgradeStatus) HasSha256checksum() bool {
	if o != nil && o.Sha256checksum != nil {
		return true
	}

	return false
}

// SetSha256checksum gets a reference to the given string and assigns it to the Sha256checksum field.
func (o *FirmwareUpgradeStatus) SetSha256checksum(v string) {
	o.Sha256checksum = &v
}

// GetUpgrade returns the Upgrade field value if set, zero value otherwise.
func (o *FirmwareUpgradeStatus) GetUpgrade() FirmwareUpgradeBaseRelationship {
	if o == nil || o.Upgrade == nil {
		var ret FirmwareUpgradeBaseRelationship
		return ret
	}
	return *o.Upgrade
}

// GetUpgradeOk returns a tuple with the Upgrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareUpgradeStatus) GetUpgradeOk() (*FirmwareUpgradeBaseRelationship, bool) {
	if o == nil || o.Upgrade == nil {
		return nil, false
	}
	return o.Upgrade, true
}

// HasUpgrade returns a boolean if a field has been set.
func (o *FirmwareUpgradeStatus) HasUpgrade() bool {
	if o != nil && o.Upgrade != nil {
		return true
	}

	return false
}

// SetUpgrade gets a reference to the given FirmwareUpgradeBaseRelationship and assigns it to the Upgrade field.
func (o *FirmwareUpgradeStatus) SetUpgrade(v FirmwareUpgradeBaseRelationship) {
	o.Upgrade = &v
}

// GetWorkflow returns the Workflow field value if set, zero value otherwise.
func (o *FirmwareUpgradeStatus) GetWorkflow() WorkflowWorkflowInfoRelationship {
	if o == nil || o.Workflow == nil {
		var ret WorkflowWorkflowInfoRelationship
		return ret
	}
	return *o.Workflow
}

// GetWorkflowOk returns a tuple with the Workflow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareUpgradeStatus) GetWorkflowOk() (*WorkflowWorkflowInfoRelationship, bool) {
	if o == nil || o.Workflow == nil {
		return nil, false
	}
	return o.Workflow, true
}

// HasWorkflow returns a boolean if a field has been set.
func (o *FirmwareUpgradeStatus) HasWorkflow() bool {
	if o != nil && o.Workflow != nil {
		return true
	}

	return false
}

// SetWorkflow gets a reference to the given WorkflowWorkflowInfoRelationship and assigns it to the Workflow field.
func (o *FirmwareUpgradeStatus) SetWorkflow(v WorkflowWorkflowInfoRelationship) {
	o.Workflow = &v
}

func (o FirmwareUpgradeStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FirmwareUpgradeStatus) UnmarshalJSON(bytes []byte) (err error) {
	type FirmwareUpgradeStatusWithoutEmbeddedStruct struct {
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

	varFirmwareUpgradeStatusWithoutEmbeddedStruct := FirmwareUpgradeStatusWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varFirmwareUpgradeStatusWithoutEmbeddedStruct)
	if err == nil {
		varFirmwareUpgradeStatus := _FirmwareUpgradeStatus{}
		varFirmwareUpgradeStatus.DownloadError = varFirmwareUpgradeStatusWithoutEmbeddedStruct.DownloadError
		varFirmwareUpgradeStatus.DownloadMessage = varFirmwareUpgradeStatusWithoutEmbeddedStruct.DownloadMessage
		varFirmwareUpgradeStatus.DownloadPercentage = varFirmwareUpgradeStatusWithoutEmbeddedStruct.DownloadPercentage
		varFirmwareUpgradeStatus.DownloadProgress = varFirmwareUpgradeStatusWithoutEmbeddedStruct.DownloadProgress
		varFirmwareUpgradeStatus.DownloadRetries = varFirmwareUpgradeStatusWithoutEmbeddedStruct.DownloadRetries
		varFirmwareUpgradeStatus.DownloadStage = varFirmwareUpgradeStatusWithoutEmbeddedStruct.DownloadStage
		varFirmwareUpgradeStatus.EpPowerStatus = varFirmwareUpgradeStatusWithoutEmbeddedStruct.EpPowerStatus
		varFirmwareUpgradeStatus.OverallError = varFirmwareUpgradeStatusWithoutEmbeddedStruct.OverallError
		varFirmwareUpgradeStatus.OverallPercentage = varFirmwareUpgradeStatusWithoutEmbeddedStruct.OverallPercentage
		varFirmwareUpgradeStatus.Overallstatus = varFirmwareUpgradeStatusWithoutEmbeddedStruct.Overallstatus
		varFirmwareUpgradeStatus.PendingType = varFirmwareUpgradeStatusWithoutEmbeddedStruct.PendingType
		varFirmwareUpgradeStatus.Sha256checksum = varFirmwareUpgradeStatusWithoutEmbeddedStruct.Sha256checksum
		varFirmwareUpgradeStatus.Upgrade = varFirmwareUpgradeStatusWithoutEmbeddedStruct.Upgrade
		varFirmwareUpgradeStatus.Workflow = varFirmwareUpgradeStatusWithoutEmbeddedStruct.Workflow
		*o = FirmwareUpgradeStatus(varFirmwareUpgradeStatus)
	} else {
		return err
	}

	varFirmwareUpgradeStatus := _FirmwareUpgradeStatus{}

	err = json.Unmarshal(bytes, &varFirmwareUpgradeStatus)
	if err == nil {
		o.MoBaseMo = varFirmwareUpgradeStatus.MoBaseMo
	} else {
		return err
	}
	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "DownloadError")
		delete(additionalProperties, "DownloadMessage")
		delete(additionalProperties, "DownloadPercentage")
		delete(additionalProperties, "DownloadProgress")
		delete(additionalProperties, "DownloadRetries")
		delete(additionalProperties, "DownloadStage")
		delete(additionalProperties, "EpPowerStatus")
		delete(additionalProperties, "OverallError")
		delete(additionalProperties, "OverallPercentage")
		delete(additionalProperties, "Overallstatus")
		delete(additionalProperties, "PendingType")
		delete(additionalProperties, "Sha256checksum")
		delete(additionalProperties, "Upgrade")
		delete(additionalProperties, "Workflow")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

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

type NullableFirmwareUpgradeStatus struct {
	value *FirmwareUpgradeStatus
	isSet bool
}

func (v NullableFirmwareUpgradeStatus) Get() *FirmwareUpgradeStatus {
	return v.value
}

func (v *NullableFirmwareUpgradeStatus) Set(val *FirmwareUpgradeStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableFirmwareUpgradeStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableFirmwareUpgradeStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirmwareUpgradeStatus(val *FirmwareUpgradeStatus) *NullableFirmwareUpgradeStatus {
	return &NullableFirmwareUpgradeStatus{value: val, isSet: true}
}

func (v NullableFirmwareUpgradeStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirmwareUpgradeStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
