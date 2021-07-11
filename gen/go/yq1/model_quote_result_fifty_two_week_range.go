/*
 * Yahoo Finance
 *
 * Yahoo Finance API specification
 *
 * API version: 1.0.8
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package yq1

import (
	"encoding/json"
)

// QuoteResultFiftyTwoWeekRange struct for QuoteResultFiftyTwoWeekRange
type QuoteResultFiftyTwoWeekRange struct {
	Raw *string `json:"raw,omitempty"`
	Fmt *string `json:"fmt,omitempty"`
}

// NewQuoteResultFiftyTwoWeekRange instantiates a new QuoteResultFiftyTwoWeekRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuoteResultFiftyTwoWeekRange() *QuoteResultFiftyTwoWeekRange {
	this := QuoteResultFiftyTwoWeekRange{}
	return &this
}

// NewQuoteResultFiftyTwoWeekRangeWithDefaults instantiates a new QuoteResultFiftyTwoWeekRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuoteResultFiftyTwoWeekRangeWithDefaults() *QuoteResultFiftyTwoWeekRange {
	this := QuoteResultFiftyTwoWeekRange{}
	return &this
}

// GetRaw returns the Raw field value if set, zero value otherwise.
func (o *QuoteResultFiftyTwoWeekRange) GetRaw() string {
	if o == nil || o.Raw == nil {
		var ret string
		return ret
	}
	return *o.Raw
}

// GetRawOk returns a tuple with the Raw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteResultFiftyTwoWeekRange) GetRawOk() (*string, bool) {
	if o == nil || o.Raw == nil {
		return nil, false
	}
	return o.Raw, true
}

// HasRaw returns a boolean if a field has been set.
func (o *QuoteResultFiftyTwoWeekRange) HasRaw() bool {
	if o != nil && o.Raw != nil {
		return true
	}

	return false
}

// SetRaw gets a reference to the given string and assigns it to the Raw field.
func (o *QuoteResultFiftyTwoWeekRange) SetRaw(v string) {
	o.Raw = &v
}

// GetFmt returns the Fmt field value if set, zero value otherwise.
func (o *QuoteResultFiftyTwoWeekRange) GetFmt() string {
	if o == nil || o.Fmt == nil {
		var ret string
		return ret
	}
	return *o.Fmt
}

// GetFmtOk returns a tuple with the Fmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteResultFiftyTwoWeekRange) GetFmtOk() (*string, bool) {
	if o == nil || o.Fmt == nil {
		return nil, false
	}
	return o.Fmt, true
}

// HasFmt returns a boolean if a field has been set.
func (o *QuoteResultFiftyTwoWeekRange) HasFmt() bool {
	if o != nil && o.Fmt != nil {
		return true
	}

	return false
}

// SetFmt gets a reference to the given string and assigns it to the Fmt field.
func (o *QuoteResultFiftyTwoWeekRange) SetFmt(v string) {
	o.Fmt = &v
}

func (o QuoteResultFiftyTwoWeekRange) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Raw != nil {
		toSerialize["raw"] = o.Raw
	}
	if o.Fmt != nil {
		toSerialize["fmt"] = o.Fmt
	}
	return json.Marshal(toSerialize)
}

type NullableQuoteResultFiftyTwoWeekRange struct {
	value *QuoteResultFiftyTwoWeekRange
	isSet bool
}

func (v NullableQuoteResultFiftyTwoWeekRange) Get() *QuoteResultFiftyTwoWeekRange {
	return v.value
}

func (v *NullableQuoteResultFiftyTwoWeekRange) Set(val *QuoteResultFiftyTwoWeekRange) {
	v.value = val
	v.isSet = true
}

func (v NullableQuoteResultFiftyTwoWeekRange) IsSet() bool {
	return v.isSet
}

func (v *NullableQuoteResultFiftyTwoWeekRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuoteResultFiftyTwoWeekRange(val *QuoteResultFiftyTwoWeekRange) *NullableQuoteResultFiftyTwoWeekRange {
	return &NullableQuoteResultFiftyTwoWeekRange{value: val, isSet: true}
}

func (v NullableQuoteResultFiftyTwoWeekRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuoteResultFiftyTwoWeekRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


