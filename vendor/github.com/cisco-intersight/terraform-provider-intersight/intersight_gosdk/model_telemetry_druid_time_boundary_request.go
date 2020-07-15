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
)

// TelemetryDruidTimeBoundaryRequest Time boundary queries return the earliest and latest data points of a data set.
type TelemetryDruidTimeBoundaryRequest struct {
	// null
	QueryType  string                   `json:"queryType"`
	DataSource TelemetryDruidDataSource `json:"dataSource"`
	// Optional, set to maxTime or minTime to return only the latest or earliest timestamp. Default to returning both if not set.
	Bound                *string                     `json:"bound,omitempty"`
	Filter               *TelemetryDruidFilter       `json:"filter,omitempty"`
	Context              *TelemetryDruidQueryContext `json:"context,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidTimeBoundaryRequest TelemetryDruidTimeBoundaryRequest

// NewTelemetryDruidTimeBoundaryRequest instantiates a new TelemetryDruidTimeBoundaryRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidTimeBoundaryRequest(queryType string, dataSource TelemetryDruidDataSource) *TelemetryDruidTimeBoundaryRequest {
	this := TelemetryDruidTimeBoundaryRequest{}
	this.QueryType = queryType
	this.DataSource = dataSource
	return &this
}

// NewTelemetryDruidTimeBoundaryRequestWithDefaults instantiates a new TelemetryDruidTimeBoundaryRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidTimeBoundaryRequestWithDefaults() *TelemetryDruidTimeBoundaryRequest {
	this := TelemetryDruidTimeBoundaryRequest{}
	return &this
}

// GetQueryType returns the QueryType field value
func (o *TelemetryDruidTimeBoundaryRequest) GetQueryType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QueryType
}

// GetQueryTypeOk returns a tuple with the QueryType field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidTimeBoundaryRequest) GetQueryTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryType, true
}

// SetQueryType sets field value
func (o *TelemetryDruidTimeBoundaryRequest) SetQueryType(v string) {
	o.QueryType = v
}

// GetDataSource returns the DataSource field value
func (o *TelemetryDruidTimeBoundaryRequest) GetDataSource() TelemetryDruidDataSource {
	if o == nil {
		var ret TelemetryDruidDataSource
		return ret
	}

	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidTimeBoundaryRequest) GetDataSourceOk() (*TelemetryDruidDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value
func (o *TelemetryDruidTimeBoundaryRequest) SetDataSource(v TelemetryDruidDataSource) {
	o.DataSource = v
}

// GetBound returns the Bound field value if set, zero value otherwise.
func (o *TelemetryDruidTimeBoundaryRequest) GetBound() string {
	if o == nil || o.Bound == nil {
		var ret string
		return ret
	}
	return *o.Bound
}

// GetBoundOk returns a tuple with the Bound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidTimeBoundaryRequest) GetBoundOk() (*string, bool) {
	if o == nil || o.Bound == nil {
		return nil, false
	}
	return o.Bound, true
}

// HasBound returns a boolean if a field has been set.
func (o *TelemetryDruidTimeBoundaryRequest) HasBound() bool {
	if o != nil && o.Bound != nil {
		return true
	}

	return false
}

// SetBound gets a reference to the given string and assigns it to the Bound field.
func (o *TelemetryDruidTimeBoundaryRequest) SetBound(v string) {
	o.Bound = &v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *TelemetryDruidTimeBoundaryRequest) GetFilter() TelemetryDruidFilter {
	if o == nil || o.Filter == nil {
		var ret TelemetryDruidFilter
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidTimeBoundaryRequest) GetFilterOk() (*TelemetryDruidFilter, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *TelemetryDruidTimeBoundaryRequest) HasFilter() bool {
	if o != nil && o.Filter != nil {
		return true
	}

	return false
}

// SetFilter gets a reference to the given TelemetryDruidFilter and assigns it to the Filter field.
func (o *TelemetryDruidTimeBoundaryRequest) SetFilter(v TelemetryDruidFilter) {
	o.Filter = &v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *TelemetryDruidTimeBoundaryRequest) GetContext() TelemetryDruidQueryContext {
	if o == nil || o.Context == nil {
		var ret TelemetryDruidQueryContext
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidTimeBoundaryRequest) GetContextOk() (*TelemetryDruidQueryContext, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *TelemetryDruidTimeBoundaryRequest) HasContext() bool {
	if o != nil && o.Context != nil {
		return true
	}

	return false
}

// SetContext gets a reference to the given TelemetryDruidQueryContext and assigns it to the Context field.
func (o *TelemetryDruidTimeBoundaryRequest) SetContext(v TelemetryDruidQueryContext) {
	o.Context = &v
}

func (o TelemetryDruidTimeBoundaryRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["queryType"] = o.QueryType
	}
	if true {
		toSerialize["dataSource"] = o.DataSource
	}
	if o.Bound != nil {
		toSerialize["bound"] = o.Bound
	}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.Context != nil {
		toSerialize["context"] = o.Context
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TelemetryDruidTimeBoundaryRequest) UnmarshalJSON(bytes []byte) (err error) {
	varTelemetryDruidTimeBoundaryRequest := _TelemetryDruidTimeBoundaryRequest{}

	if err = json.Unmarshal(bytes, &varTelemetryDruidTimeBoundaryRequest); err == nil {
		*o = TelemetryDruidTimeBoundaryRequest(varTelemetryDruidTimeBoundaryRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "queryType")
		delete(additionalProperties, "dataSource")
		delete(additionalProperties, "bound")
		delete(additionalProperties, "filter")
		delete(additionalProperties, "context")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidTimeBoundaryRequest struct {
	value *TelemetryDruidTimeBoundaryRequest
	isSet bool
}

func (v NullableTelemetryDruidTimeBoundaryRequest) Get() *TelemetryDruidTimeBoundaryRequest {
	return v.value
}

func (v *NullableTelemetryDruidTimeBoundaryRequest) Set(val *TelemetryDruidTimeBoundaryRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidTimeBoundaryRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidTimeBoundaryRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidTimeBoundaryRequest(val *TelemetryDruidTimeBoundaryRequest) *NullableTelemetryDruidTimeBoundaryRequest {
	return &NullableTelemetryDruidTimeBoundaryRequest{value: val, isSet: true}
}

func (v NullableTelemetryDruidTimeBoundaryRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidTimeBoundaryRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
