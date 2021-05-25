/*
 * Yahoo Finance
 *
 * Yahoo Finance API specification
 *
 * API version: 1.0.2
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package yf

import (
	"encoding/json"
)

// ChartResponseChartIndicators struct for ChartResponseChartIndicators
type ChartResponseChartIndicators struct {
	Quote *[]ChartResponseChartIndicatorsQuote `json:"quote,omitempty"`
}

// NewChartResponseChartIndicators instantiates a new ChartResponseChartIndicators object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChartResponseChartIndicators() *ChartResponseChartIndicators {
	this := ChartResponseChartIndicators{}
	return &this
}

// NewChartResponseChartIndicatorsWithDefaults instantiates a new ChartResponseChartIndicators object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChartResponseChartIndicatorsWithDefaults() *ChartResponseChartIndicators {
	this := ChartResponseChartIndicators{}
	return &this
}

// GetQuote returns the Quote field value if set, zero value otherwise.
func (o *ChartResponseChartIndicators) GetQuote() []ChartResponseChartIndicatorsQuote {
	if o == nil || o.Quote == nil {
		var ret []ChartResponseChartIndicatorsQuote
		return ret
	}
	return *o.Quote
}

// GetQuoteOk returns a tuple with the Quote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChartResponseChartIndicators) GetQuoteOk() (*[]ChartResponseChartIndicatorsQuote, bool) {
	if o == nil || o.Quote == nil {
		return nil, false
	}
	return o.Quote, true
}

// HasQuote returns a boolean if a field has been set.
func (o *ChartResponseChartIndicators) HasQuote() bool {
	if o != nil && o.Quote != nil {
		return true
	}

	return false
}

// SetQuote gets a reference to the given []ChartResponseChartIndicatorsQuote and assigns it to the Quote field.
func (o *ChartResponseChartIndicators) SetQuote(v []ChartResponseChartIndicatorsQuote) {
	o.Quote = &v
}

func (o ChartResponseChartIndicators) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Quote != nil {
		toSerialize["quote"] = o.Quote
	}
	return json.Marshal(toSerialize)
}

type NullableChartResponseChartIndicators struct {
	value *ChartResponseChartIndicators
	isSet bool
}

func (v NullableChartResponseChartIndicators) Get() *ChartResponseChartIndicators {
	return v.value
}

func (v *NullableChartResponseChartIndicators) Set(val *ChartResponseChartIndicators) {
	v.value = val
	v.isSet = true
}

func (v NullableChartResponseChartIndicators) IsSet() bool {
	return v.isSet
}

func (v *NullableChartResponseChartIndicators) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChartResponseChartIndicators(val *ChartResponseChartIndicators) *NullableChartResponseChartIndicators {
	return &NullableChartResponseChartIndicators{value: val, isSet: true}
}

func (v NullableChartResponseChartIndicators) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChartResponseChartIndicators) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


