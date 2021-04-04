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

// QuoteResponseRegularMarketChange struct for QuoteResponseRegularMarketChange
type QuoteResponseRegularMarketChange struct {
	Raw *float32 `json:"raw,omitempty"`
	Fmt *string `json:"fmt,omitempty"`
}

// NewQuoteResponseRegularMarketChange instantiates a new QuoteResponseRegularMarketChange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuoteResponseRegularMarketChange() *QuoteResponseRegularMarketChange {
	this := QuoteResponseRegularMarketChange{}
	return &this
}

// NewQuoteResponseRegularMarketChangeWithDefaults instantiates a new QuoteResponseRegularMarketChange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuoteResponseRegularMarketChangeWithDefaults() *QuoteResponseRegularMarketChange {
	this := QuoteResponseRegularMarketChange{}
	return &this
}

// GetRaw returns the Raw field value if set, zero value otherwise.
func (o *QuoteResponseRegularMarketChange) GetRaw() float32 {
	if o == nil || o.Raw == nil {
		var ret float32
		return ret
	}
	return *o.Raw
}

// GetRawOk returns a tuple with the Raw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteResponseRegularMarketChange) GetRawOk() (*float32, bool) {
	if o == nil || o.Raw == nil {
		return nil, false
	}
	return o.Raw, true
}

// HasRaw returns a boolean if a field has been set.
func (o *QuoteResponseRegularMarketChange) HasRaw() bool {
	if o != nil && o.Raw != nil {
		return true
	}

	return false
}

// SetRaw gets a reference to the given float32 and assigns it to the Raw field.
func (o *QuoteResponseRegularMarketChange) SetRaw(v float32) {
	o.Raw = &v
}

// GetFmt returns the Fmt field value if set, zero value otherwise.
func (o *QuoteResponseRegularMarketChange) GetFmt() string {
	if o == nil || o.Fmt == nil {
		var ret string
		return ret
	}
	return *o.Fmt
}

// GetFmtOk returns a tuple with the Fmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteResponseRegularMarketChange) GetFmtOk() (*string, bool) {
	if o == nil || o.Fmt == nil {
		return nil, false
	}
	return o.Fmt, true
}

// HasFmt returns a boolean if a field has been set.
func (o *QuoteResponseRegularMarketChange) HasFmt() bool {
	if o != nil && o.Fmt != nil {
		return true
	}

	return false
}

// SetFmt gets a reference to the given string and assigns it to the Fmt field.
func (o *QuoteResponseRegularMarketChange) SetFmt(v string) {
	o.Fmt = &v
}

func (o QuoteResponseRegularMarketChange) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Raw != nil {
		toSerialize["raw"] = o.Raw
	}
	if o.Fmt != nil {
		toSerialize["fmt"] = o.Fmt
	}
	return json.Marshal(toSerialize)
}

type NullableQuoteResponseRegularMarketChange struct {
	value *QuoteResponseRegularMarketChange
	isSet bool
}

func (v NullableQuoteResponseRegularMarketChange) Get() *QuoteResponseRegularMarketChange {
	return v.value
}

func (v *NullableQuoteResponseRegularMarketChange) Set(val *QuoteResponseRegularMarketChange) {
	v.value = val
	v.isSet = true
}

func (v NullableQuoteResponseRegularMarketChange) IsSet() bool {
	return v.isSet
}

func (v *NullableQuoteResponseRegularMarketChange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuoteResponseRegularMarketChange(val *QuoteResponseRegularMarketChange) *NullableQuoteResponseRegularMarketChange {
	return &NullableQuoteResponseRegularMarketChange{value: val, isSet: true}
}

func (v NullableQuoteResponseRegularMarketChange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuoteResponseRegularMarketChange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


