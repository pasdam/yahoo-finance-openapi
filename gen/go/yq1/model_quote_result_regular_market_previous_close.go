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

// QuoteResultRegularMarketPreviousClose struct for QuoteResultRegularMarketPreviousClose
type QuoteResultRegularMarketPreviousClose struct {
	Raw *float32 `json:"raw,omitempty"`
	Fmt *string `json:"fmt,omitempty"`
}

// NewQuoteResultRegularMarketPreviousClose instantiates a new QuoteResultRegularMarketPreviousClose object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuoteResultRegularMarketPreviousClose() *QuoteResultRegularMarketPreviousClose {
	this := QuoteResultRegularMarketPreviousClose{}
	return &this
}

// NewQuoteResultRegularMarketPreviousCloseWithDefaults instantiates a new QuoteResultRegularMarketPreviousClose object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuoteResultRegularMarketPreviousCloseWithDefaults() *QuoteResultRegularMarketPreviousClose {
	this := QuoteResultRegularMarketPreviousClose{}
	return &this
}

// GetRaw returns the Raw field value if set, zero value otherwise.
func (o *QuoteResultRegularMarketPreviousClose) GetRaw() float32 {
	if o == nil || o.Raw == nil {
		var ret float32
		return ret
	}
	return *o.Raw
}

// GetRawOk returns a tuple with the Raw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteResultRegularMarketPreviousClose) GetRawOk() (*float32, bool) {
	if o == nil || o.Raw == nil {
		return nil, false
	}
	return o.Raw, true
}

// HasRaw returns a boolean if a field has been set.
func (o *QuoteResultRegularMarketPreviousClose) HasRaw() bool {
	if o != nil && o.Raw != nil {
		return true
	}

	return false
}

// SetRaw gets a reference to the given float32 and assigns it to the Raw field.
func (o *QuoteResultRegularMarketPreviousClose) SetRaw(v float32) {
	o.Raw = &v
}

// GetFmt returns the Fmt field value if set, zero value otherwise.
func (o *QuoteResultRegularMarketPreviousClose) GetFmt() string {
	if o == nil || o.Fmt == nil {
		var ret string
		return ret
	}
	return *o.Fmt
}

// GetFmtOk returns a tuple with the Fmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteResultRegularMarketPreviousClose) GetFmtOk() (*string, bool) {
	if o == nil || o.Fmt == nil {
		return nil, false
	}
	return o.Fmt, true
}

// HasFmt returns a boolean if a field has been set.
func (o *QuoteResultRegularMarketPreviousClose) HasFmt() bool {
	if o != nil && o.Fmt != nil {
		return true
	}

	return false
}

// SetFmt gets a reference to the given string and assigns it to the Fmt field.
func (o *QuoteResultRegularMarketPreviousClose) SetFmt(v string) {
	o.Fmt = &v
}

func (o QuoteResultRegularMarketPreviousClose) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Raw != nil {
		toSerialize["raw"] = o.Raw
	}
	if o.Fmt != nil {
		toSerialize["fmt"] = o.Fmt
	}
	return json.Marshal(toSerialize)
}

type NullableQuoteResultRegularMarketPreviousClose struct {
	value *QuoteResultRegularMarketPreviousClose
	isSet bool
}

func (v NullableQuoteResultRegularMarketPreviousClose) Get() *QuoteResultRegularMarketPreviousClose {
	return v.value
}

func (v *NullableQuoteResultRegularMarketPreviousClose) Set(val *QuoteResultRegularMarketPreviousClose) {
	v.value = val
	v.isSet = true
}

func (v NullableQuoteResultRegularMarketPreviousClose) IsSet() bool {
	return v.isSet
}

func (v *NullableQuoteResultRegularMarketPreviousClose) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuoteResultRegularMarketPreviousClose(val *QuoteResultRegularMarketPreviousClose) *NullableQuoteResultRegularMarketPreviousClose {
	return &NullableQuoteResultRegularMarketPreviousClose{value: val, isSet: true}
}

func (v NullableQuoteResultRegularMarketPreviousClose) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuoteResultRegularMarketPreviousClose) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


