/*
 * Yahoo Finance
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package yf

import (
	"encoding/json"
)

// QuoteResponseFiftyTwoWeekLowChangePercent struct for QuoteResponseFiftyTwoWeekLowChangePercent
type QuoteResponseFiftyTwoWeekLowChangePercent struct {
	Raw *float32 `json:"raw,omitempty"`
	Fmt *string `json:"fmt,omitempty"`
}

// NewQuoteResponseFiftyTwoWeekLowChangePercent instantiates a new QuoteResponseFiftyTwoWeekLowChangePercent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuoteResponseFiftyTwoWeekLowChangePercent() *QuoteResponseFiftyTwoWeekLowChangePercent {
	this := QuoteResponseFiftyTwoWeekLowChangePercent{}
	return &this
}

// NewQuoteResponseFiftyTwoWeekLowChangePercentWithDefaults instantiates a new QuoteResponseFiftyTwoWeekLowChangePercent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuoteResponseFiftyTwoWeekLowChangePercentWithDefaults() *QuoteResponseFiftyTwoWeekLowChangePercent {
	this := QuoteResponseFiftyTwoWeekLowChangePercent{}
	return &this
}

// GetRaw returns the Raw field value if set, zero value otherwise.
func (o *QuoteResponseFiftyTwoWeekLowChangePercent) GetRaw() float32 {
	if o == nil || o.Raw == nil {
		var ret float32
		return ret
	}
	return *o.Raw
}

// GetRawOk returns a tuple with the Raw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteResponseFiftyTwoWeekLowChangePercent) GetRawOk() (*float32, bool) {
	if o == nil || o.Raw == nil {
		return nil, false
	}
	return o.Raw, true
}

// HasRaw returns a boolean if a field has been set.
func (o *QuoteResponseFiftyTwoWeekLowChangePercent) HasRaw() bool {
	if o != nil && o.Raw != nil {
		return true
	}

	return false
}

// SetRaw gets a reference to the given float32 and assigns it to the Raw field.
func (o *QuoteResponseFiftyTwoWeekLowChangePercent) SetRaw(v float32) {
	o.Raw = &v
}

// GetFmt returns the Fmt field value if set, zero value otherwise.
func (o *QuoteResponseFiftyTwoWeekLowChangePercent) GetFmt() string {
	if o == nil || o.Fmt == nil {
		var ret string
		return ret
	}
	return *o.Fmt
}

// GetFmtOk returns a tuple with the Fmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteResponseFiftyTwoWeekLowChangePercent) GetFmtOk() (*string, bool) {
	if o == nil || o.Fmt == nil {
		return nil, false
	}
	return o.Fmt, true
}

// HasFmt returns a boolean if a field has been set.
func (o *QuoteResponseFiftyTwoWeekLowChangePercent) HasFmt() bool {
	if o != nil && o.Fmt != nil {
		return true
	}

	return false
}

// SetFmt gets a reference to the given string and assigns it to the Fmt field.
func (o *QuoteResponseFiftyTwoWeekLowChangePercent) SetFmt(v string) {
	o.Fmt = &v
}

func (o QuoteResponseFiftyTwoWeekLowChangePercent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Raw != nil {
		toSerialize["raw"] = o.Raw
	}
	if o.Fmt != nil {
		toSerialize["fmt"] = o.Fmt
	}
	return json.Marshal(toSerialize)
}

type NullableQuoteResponseFiftyTwoWeekLowChangePercent struct {
	value *QuoteResponseFiftyTwoWeekLowChangePercent
	isSet bool
}

func (v NullableQuoteResponseFiftyTwoWeekLowChangePercent) Get() *QuoteResponseFiftyTwoWeekLowChangePercent {
	return v.value
}

func (v *NullableQuoteResponseFiftyTwoWeekLowChangePercent) Set(val *QuoteResponseFiftyTwoWeekLowChangePercent) {
	v.value = val
	v.isSet = true
}

func (v NullableQuoteResponseFiftyTwoWeekLowChangePercent) IsSet() bool {
	return v.isSet
}

func (v *NullableQuoteResponseFiftyTwoWeekLowChangePercent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuoteResponseFiftyTwoWeekLowChangePercent(val *QuoteResponseFiftyTwoWeekLowChangePercent) *NullableQuoteResponseFiftyTwoWeekLowChangePercent {
	return &NullableQuoteResponseFiftyTwoWeekLowChangePercent{value: val, isSet: true}
}

func (v NullableQuoteResponseFiftyTwoWeekLowChangePercent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuoteResponseFiftyTwoWeekLowChangePercent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


