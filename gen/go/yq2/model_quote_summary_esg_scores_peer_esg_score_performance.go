/*
 * Yahoo Finance
 *
 * Yahoo Finance API specification
 *
 * API version: 1.0.8
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package yq2

import (
	"encoding/json"
)

// QuoteSummaryEsgScoresPeerEsgScorePerformance struct for QuoteSummaryEsgScoresPeerEsgScorePerformance
type QuoteSummaryEsgScoresPeerEsgScorePerformance struct {
	Min *float32 `json:"min,omitempty"`
	Avg *float32 `json:"avg,omitempty"`
	Max *float32 `json:"max,omitempty"`
}

// NewQuoteSummaryEsgScoresPeerEsgScorePerformance instantiates a new QuoteSummaryEsgScoresPeerEsgScorePerformance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuoteSummaryEsgScoresPeerEsgScorePerformance() *QuoteSummaryEsgScoresPeerEsgScorePerformance {
	this := QuoteSummaryEsgScoresPeerEsgScorePerformance{}
	return &this
}

// NewQuoteSummaryEsgScoresPeerEsgScorePerformanceWithDefaults instantiates a new QuoteSummaryEsgScoresPeerEsgScorePerformance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuoteSummaryEsgScoresPeerEsgScorePerformanceWithDefaults() *QuoteSummaryEsgScoresPeerEsgScorePerformance {
	this := QuoteSummaryEsgScoresPeerEsgScorePerformance{}
	return &this
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *QuoteSummaryEsgScoresPeerEsgScorePerformance) GetMin() float32 {
	if o == nil || o.Min == nil {
		var ret float32
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteSummaryEsgScoresPeerEsgScorePerformance) GetMinOk() (*float32, bool) {
	if o == nil || o.Min == nil {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *QuoteSummaryEsgScoresPeerEsgScorePerformance) HasMin() bool {
	if o != nil && o.Min != nil {
		return true
	}

	return false
}

// SetMin gets a reference to the given float32 and assigns it to the Min field.
func (o *QuoteSummaryEsgScoresPeerEsgScorePerformance) SetMin(v float32) {
	o.Min = &v
}

// GetAvg returns the Avg field value if set, zero value otherwise.
func (o *QuoteSummaryEsgScoresPeerEsgScorePerformance) GetAvg() float32 {
	if o == nil || o.Avg == nil {
		var ret float32
		return ret
	}
	return *o.Avg
}

// GetAvgOk returns a tuple with the Avg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteSummaryEsgScoresPeerEsgScorePerformance) GetAvgOk() (*float32, bool) {
	if o == nil || o.Avg == nil {
		return nil, false
	}
	return o.Avg, true
}

// HasAvg returns a boolean if a field has been set.
func (o *QuoteSummaryEsgScoresPeerEsgScorePerformance) HasAvg() bool {
	if o != nil && o.Avg != nil {
		return true
	}

	return false
}

// SetAvg gets a reference to the given float32 and assigns it to the Avg field.
func (o *QuoteSummaryEsgScoresPeerEsgScorePerformance) SetAvg(v float32) {
	o.Avg = &v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *QuoteSummaryEsgScoresPeerEsgScorePerformance) GetMax() float32 {
	if o == nil || o.Max == nil {
		var ret float32
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteSummaryEsgScoresPeerEsgScorePerformance) GetMaxOk() (*float32, bool) {
	if o == nil || o.Max == nil {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *QuoteSummaryEsgScoresPeerEsgScorePerformance) HasMax() bool {
	if o != nil && o.Max != nil {
		return true
	}

	return false
}

// SetMax gets a reference to the given float32 and assigns it to the Max field.
func (o *QuoteSummaryEsgScoresPeerEsgScorePerformance) SetMax(v float32) {
	o.Max = &v
}

func (o QuoteSummaryEsgScoresPeerEsgScorePerformance) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Min != nil {
		toSerialize["min"] = o.Min
	}
	if o.Avg != nil {
		toSerialize["avg"] = o.Avg
	}
	if o.Max != nil {
		toSerialize["max"] = o.Max
	}
	return json.Marshal(toSerialize)
}

type NullableQuoteSummaryEsgScoresPeerEsgScorePerformance struct {
	value *QuoteSummaryEsgScoresPeerEsgScorePerformance
	isSet bool
}

func (v NullableQuoteSummaryEsgScoresPeerEsgScorePerformance) Get() *QuoteSummaryEsgScoresPeerEsgScorePerformance {
	return v.value
}

func (v *NullableQuoteSummaryEsgScoresPeerEsgScorePerformance) Set(val *QuoteSummaryEsgScoresPeerEsgScorePerformance) {
	v.value = val
	v.isSet = true
}

func (v NullableQuoteSummaryEsgScoresPeerEsgScorePerformance) IsSet() bool {
	return v.isSet
}

func (v *NullableQuoteSummaryEsgScoresPeerEsgScorePerformance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuoteSummaryEsgScoresPeerEsgScorePerformance(val *QuoteSummaryEsgScoresPeerEsgScorePerformance) *NullableQuoteSummaryEsgScoresPeerEsgScorePerformance {
	return &NullableQuoteSummaryEsgScoresPeerEsgScorePerformance{value: val, isSet: true}
}

func (v NullableQuoteSummaryEsgScoresPeerEsgScorePerformance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuoteSummaryEsgScoresPeerEsgScorePerformance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


