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

// QuoteResultRegularMarketTime struct for QuoteResultRegularMarketTime
type QuoteResultRegularMarketTime struct {
	Raw *int32 `json:"raw,omitempty"`
	Fmt *string `json:"fmt,omitempty"`
}

// NewQuoteResultRegularMarketTime instantiates a new QuoteResultRegularMarketTime object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuoteResultRegularMarketTime() *QuoteResultRegularMarketTime {
	this := QuoteResultRegularMarketTime{}
	return &this
}

// NewQuoteResultRegularMarketTimeWithDefaults instantiates a new QuoteResultRegularMarketTime object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuoteResultRegularMarketTimeWithDefaults() *QuoteResultRegularMarketTime {
	this := QuoteResultRegularMarketTime{}
	return &this
}

// GetRaw returns the Raw field value if set, zero value otherwise.
func (o *QuoteResultRegularMarketTime) GetRaw() int32 {
	if o == nil || o.Raw == nil {
		var ret int32
		return ret
	}
	return *o.Raw
}

// GetRawOk returns a tuple with the Raw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteResultRegularMarketTime) GetRawOk() (*int32, bool) {
	if o == nil || o.Raw == nil {
		return nil, false
	}
	return o.Raw, true
}

// HasRaw returns a boolean if a field has been set.
func (o *QuoteResultRegularMarketTime) HasRaw() bool {
	if o != nil && o.Raw != nil {
		return true
	}

	return false
}

// SetRaw gets a reference to the given int32 and assigns it to the Raw field.
func (o *QuoteResultRegularMarketTime) SetRaw(v int32) {
	o.Raw = &v
}

// GetFmt returns the Fmt field value if set, zero value otherwise.
func (o *QuoteResultRegularMarketTime) GetFmt() string {
	if o == nil || o.Fmt == nil {
		var ret string
		return ret
	}
	return *o.Fmt
}

// GetFmtOk returns a tuple with the Fmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteResultRegularMarketTime) GetFmtOk() (*string, bool) {
	if o == nil || o.Fmt == nil {
		return nil, false
	}
	return o.Fmt, true
}

// HasFmt returns a boolean if a field has been set.
func (o *QuoteResultRegularMarketTime) HasFmt() bool {
	if o != nil && o.Fmt != nil {
		return true
	}

	return false
}

// SetFmt gets a reference to the given string and assigns it to the Fmt field.
func (o *QuoteResultRegularMarketTime) SetFmt(v string) {
	o.Fmt = &v
}

func (o QuoteResultRegularMarketTime) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Raw != nil {
		toSerialize["raw"] = o.Raw
	}
	if o.Fmt != nil {
		toSerialize["fmt"] = o.Fmt
	}
	return json.Marshal(toSerialize)
}

type NullableQuoteResultRegularMarketTime struct {
	value *QuoteResultRegularMarketTime
	isSet bool
}

func (v NullableQuoteResultRegularMarketTime) Get() *QuoteResultRegularMarketTime {
	return v.value
}

func (v *NullableQuoteResultRegularMarketTime) Set(val *QuoteResultRegularMarketTime) {
	v.value = val
	v.isSet = true
}

func (v NullableQuoteResultRegularMarketTime) IsSet() bool {
	return v.isSet
}

func (v *NullableQuoteResultRegularMarketTime) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuoteResultRegularMarketTime(val *QuoteResultRegularMarketTime) *NullableQuoteResultRegularMarketTime {
	return &NullableQuoteResultRegularMarketTime{value: val, isSet: true}
}

func (v NullableQuoteResultRegularMarketTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuoteResultRegularMarketTime) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


