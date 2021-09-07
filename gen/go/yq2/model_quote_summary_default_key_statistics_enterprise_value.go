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

// QuoteSummaryDefaultKeyStatisticsEnterpriseValue struct for QuoteSummaryDefaultKeyStatisticsEnterpriseValue
type QuoteSummaryDefaultKeyStatisticsEnterpriseValue struct {
	Raw *int64 `json:"raw,omitempty"`
	Fmt *string `json:"fmt,omitempty"`
	LongFmt *string `json:"longFmt,omitempty"`
}

// NewQuoteSummaryDefaultKeyStatisticsEnterpriseValue instantiates a new QuoteSummaryDefaultKeyStatisticsEnterpriseValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuoteSummaryDefaultKeyStatisticsEnterpriseValue() *QuoteSummaryDefaultKeyStatisticsEnterpriseValue {
	this := QuoteSummaryDefaultKeyStatisticsEnterpriseValue{}
	return &this
}

// NewQuoteSummaryDefaultKeyStatisticsEnterpriseValueWithDefaults instantiates a new QuoteSummaryDefaultKeyStatisticsEnterpriseValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuoteSummaryDefaultKeyStatisticsEnterpriseValueWithDefaults() *QuoteSummaryDefaultKeyStatisticsEnterpriseValue {
	this := QuoteSummaryDefaultKeyStatisticsEnterpriseValue{}
	return &this
}

// GetRaw returns the Raw field value if set, zero value otherwise.
func (o *QuoteSummaryDefaultKeyStatisticsEnterpriseValue) GetRaw() int64 {
	if o == nil || o.Raw == nil {
		var ret int64
		return ret
	}
	return *o.Raw
}

// GetRawOk returns a tuple with the Raw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteSummaryDefaultKeyStatisticsEnterpriseValue) GetRawOk() (*int64, bool) {
	if o == nil || o.Raw == nil {
		return nil, false
	}
	return o.Raw, true
}

// HasRaw returns a boolean if a field has been set.
func (o *QuoteSummaryDefaultKeyStatisticsEnterpriseValue) HasRaw() bool {
	if o != nil && o.Raw != nil {
		return true
	}

	return false
}

// SetRaw gets a reference to the given int64 and assigns it to the Raw field.
func (o *QuoteSummaryDefaultKeyStatisticsEnterpriseValue) SetRaw(v int64) {
	o.Raw = &v
}

// GetFmt returns the Fmt field value if set, zero value otherwise.
func (o *QuoteSummaryDefaultKeyStatisticsEnterpriseValue) GetFmt() string {
	if o == nil || o.Fmt == nil {
		var ret string
		return ret
	}
	return *o.Fmt
}

// GetFmtOk returns a tuple with the Fmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteSummaryDefaultKeyStatisticsEnterpriseValue) GetFmtOk() (*string, bool) {
	if o == nil || o.Fmt == nil {
		return nil, false
	}
	return o.Fmt, true
}

// HasFmt returns a boolean if a field has been set.
func (o *QuoteSummaryDefaultKeyStatisticsEnterpriseValue) HasFmt() bool {
	if o != nil && o.Fmt != nil {
		return true
	}

	return false
}

// SetFmt gets a reference to the given string and assigns it to the Fmt field.
func (o *QuoteSummaryDefaultKeyStatisticsEnterpriseValue) SetFmt(v string) {
	o.Fmt = &v
}

// GetLongFmt returns the LongFmt field value if set, zero value otherwise.
func (o *QuoteSummaryDefaultKeyStatisticsEnterpriseValue) GetLongFmt() string {
	if o == nil || o.LongFmt == nil {
		var ret string
		return ret
	}
	return *o.LongFmt
}

// GetLongFmtOk returns a tuple with the LongFmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteSummaryDefaultKeyStatisticsEnterpriseValue) GetLongFmtOk() (*string, bool) {
	if o == nil || o.LongFmt == nil {
		return nil, false
	}
	return o.LongFmt, true
}

// HasLongFmt returns a boolean if a field has been set.
func (o *QuoteSummaryDefaultKeyStatisticsEnterpriseValue) HasLongFmt() bool {
	if o != nil && o.LongFmt != nil {
		return true
	}

	return false
}

// SetLongFmt gets a reference to the given string and assigns it to the LongFmt field.
func (o *QuoteSummaryDefaultKeyStatisticsEnterpriseValue) SetLongFmt(v string) {
	o.LongFmt = &v
}

func (o QuoteSummaryDefaultKeyStatisticsEnterpriseValue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Raw != nil {
		toSerialize["raw"] = o.Raw
	}
	if o.Fmt != nil {
		toSerialize["fmt"] = o.Fmt
	}
	if o.LongFmt != nil {
		toSerialize["longFmt"] = o.LongFmt
	}
	return json.Marshal(toSerialize)
}

type NullableQuoteSummaryDefaultKeyStatisticsEnterpriseValue struct {
	value *QuoteSummaryDefaultKeyStatisticsEnterpriseValue
	isSet bool
}

func (v NullableQuoteSummaryDefaultKeyStatisticsEnterpriseValue) Get() *QuoteSummaryDefaultKeyStatisticsEnterpriseValue {
	return v.value
}

func (v *NullableQuoteSummaryDefaultKeyStatisticsEnterpriseValue) Set(val *QuoteSummaryDefaultKeyStatisticsEnterpriseValue) {
	v.value = val
	v.isSet = true
}

func (v NullableQuoteSummaryDefaultKeyStatisticsEnterpriseValue) IsSet() bool {
	return v.isSet
}

func (v *NullableQuoteSummaryDefaultKeyStatisticsEnterpriseValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuoteSummaryDefaultKeyStatisticsEnterpriseValue(val *QuoteSummaryDefaultKeyStatisticsEnterpriseValue) *NullableQuoteSummaryDefaultKeyStatisticsEnterpriseValue {
	return &NullableQuoteSummaryDefaultKeyStatisticsEnterpriseValue{value: val, isSet: true}
}

func (v NullableQuoteSummaryDefaultKeyStatisticsEnterpriseValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuoteSummaryDefaultKeyStatisticsEnterpriseValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

