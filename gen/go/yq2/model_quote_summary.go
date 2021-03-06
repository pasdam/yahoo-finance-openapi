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

// QuoteSummary struct for QuoteSummary
type QuoteSummary struct {
	Result *[]QuoteSummaryResult `json:"result,omitempty"`
	Error *Error `json:"error,omitempty"`
}

// NewQuoteSummary instantiates a new QuoteSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuoteSummary() *QuoteSummary {
	this := QuoteSummary{}
	return &this
}

// NewQuoteSummaryWithDefaults instantiates a new QuoteSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuoteSummaryWithDefaults() *QuoteSummary {
	this := QuoteSummary{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *QuoteSummary) GetResult() []QuoteSummaryResult {
	if o == nil || o.Result == nil {
		var ret []QuoteSummaryResult
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteSummary) GetResultOk() (*[]QuoteSummaryResult, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *QuoteSummary) HasResult() bool {
	if o != nil && o.Result != nil {
		return true
	}

	return false
}

// SetResult gets a reference to the given []QuoteSummaryResult and assigns it to the Result field.
func (o *QuoteSummary) SetResult(v []QuoteSummaryResult) {
	o.Result = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *QuoteSummary) GetError() Error {
	if o == nil || o.Error == nil {
		var ret Error
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteSummary) GetErrorOk() (*Error, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *QuoteSummary) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given Error and assigns it to the Error field.
func (o *QuoteSummary) SetError(v Error) {
	o.Error = &v
}

func (o QuoteSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableQuoteSummary struct {
	value *QuoteSummary
	isSet bool
}

func (v NullableQuoteSummary) Get() *QuoteSummary {
	return v.value
}

func (v *NullableQuoteSummary) Set(val *QuoteSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableQuoteSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableQuoteSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuoteSummary(val *QuoteSummary) *NullableQuoteSummary {
	return &NullableQuoteSummary{value: val, isSet: true}
}

func (v NullableQuoteSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuoteSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


