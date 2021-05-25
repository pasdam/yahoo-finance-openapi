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

// ChartResponseChartMetaCurrentTradingPeriod struct for ChartResponseChartMetaCurrentTradingPeriod
type ChartResponseChartMetaCurrentTradingPeriod struct {
	Pre *ChartResponseChartMetaCurrentTradingPeriodPre `json:"pre,omitempty"`
	Regular *ChartResponseChartMetaCurrentTradingPeriodRegular `json:"regular,omitempty"`
	Post *ChartResponseChartMetaCurrentTradingPeriodPost `json:"post,omitempty"`
}

// NewChartResponseChartMetaCurrentTradingPeriod instantiates a new ChartResponseChartMetaCurrentTradingPeriod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChartResponseChartMetaCurrentTradingPeriod() *ChartResponseChartMetaCurrentTradingPeriod {
	this := ChartResponseChartMetaCurrentTradingPeriod{}
	return &this
}

// NewChartResponseChartMetaCurrentTradingPeriodWithDefaults instantiates a new ChartResponseChartMetaCurrentTradingPeriod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChartResponseChartMetaCurrentTradingPeriodWithDefaults() *ChartResponseChartMetaCurrentTradingPeriod {
	this := ChartResponseChartMetaCurrentTradingPeriod{}
	return &this
}

// GetPre returns the Pre field value if set, zero value otherwise.
func (o *ChartResponseChartMetaCurrentTradingPeriod) GetPre() ChartResponseChartMetaCurrentTradingPeriodPre {
	if o == nil || o.Pre == nil {
		var ret ChartResponseChartMetaCurrentTradingPeriodPre
		return ret
	}
	return *o.Pre
}

// GetPreOk returns a tuple with the Pre field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChartResponseChartMetaCurrentTradingPeriod) GetPreOk() (*ChartResponseChartMetaCurrentTradingPeriodPre, bool) {
	if o == nil || o.Pre == nil {
		return nil, false
	}
	return o.Pre, true
}

// HasPre returns a boolean if a field has been set.
func (o *ChartResponseChartMetaCurrentTradingPeriod) HasPre() bool {
	if o != nil && o.Pre != nil {
		return true
	}

	return false
}

// SetPre gets a reference to the given ChartResponseChartMetaCurrentTradingPeriodPre and assigns it to the Pre field.
func (o *ChartResponseChartMetaCurrentTradingPeriod) SetPre(v ChartResponseChartMetaCurrentTradingPeriodPre) {
	o.Pre = &v
}

// GetRegular returns the Regular field value if set, zero value otherwise.
func (o *ChartResponseChartMetaCurrentTradingPeriod) GetRegular() ChartResponseChartMetaCurrentTradingPeriodRegular {
	if o == nil || o.Regular == nil {
		var ret ChartResponseChartMetaCurrentTradingPeriodRegular
		return ret
	}
	return *o.Regular
}

// GetRegularOk returns a tuple with the Regular field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChartResponseChartMetaCurrentTradingPeriod) GetRegularOk() (*ChartResponseChartMetaCurrentTradingPeriodRegular, bool) {
	if o == nil || o.Regular == nil {
		return nil, false
	}
	return o.Regular, true
}

// HasRegular returns a boolean if a field has been set.
func (o *ChartResponseChartMetaCurrentTradingPeriod) HasRegular() bool {
	if o != nil && o.Regular != nil {
		return true
	}

	return false
}

// SetRegular gets a reference to the given ChartResponseChartMetaCurrentTradingPeriodRegular and assigns it to the Regular field.
func (o *ChartResponseChartMetaCurrentTradingPeriod) SetRegular(v ChartResponseChartMetaCurrentTradingPeriodRegular) {
	o.Regular = &v
}

// GetPost returns the Post field value if set, zero value otherwise.
func (o *ChartResponseChartMetaCurrentTradingPeriod) GetPost() ChartResponseChartMetaCurrentTradingPeriodPost {
	if o == nil || o.Post == nil {
		var ret ChartResponseChartMetaCurrentTradingPeriodPost
		return ret
	}
	return *o.Post
}

// GetPostOk returns a tuple with the Post field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChartResponseChartMetaCurrentTradingPeriod) GetPostOk() (*ChartResponseChartMetaCurrentTradingPeriodPost, bool) {
	if o == nil || o.Post == nil {
		return nil, false
	}
	return o.Post, true
}

// HasPost returns a boolean if a field has been set.
func (o *ChartResponseChartMetaCurrentTradingPeriod) HasPost() bool {
	if o != nil && o.Post != nil {
		return true
	}

	return false
}

// SetPost gets a reference to the given ChartResponseChartMetaCurrentTradingPeriodPost and assigns it to the Post field.
func (o *ChartResponseChartMetaCurrentTradingPeriod) SetPost(v ChartResponseChartMetaCurrentTradingPeriodPost) {
	o.Post = &v
}

func (o ChartResponseChartMetaCurrentTradingPeriod) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Pre != nil {
		toSerialize["pre"] = o.Pre
	}
	if o.Regular != nil {
		toSerialize["regular"] = o.Regular
	}
	if o.Post != nil {
		toSerialize["post"] = o.Post
	}
	return json.Marshal(toSerialize)
}

type NullableChartResponseChartMetaCurrentTradingPeriod struct {
	value *ChartResponseChartMetaCurrentTradingPeriod
	isSet bool
}

func (v NullableChartResponseChartMetaCurrentTradingPeriod) Get() *ChartResponseChartMetaCurrentTradingPeriod {
	return v.value
}

func (v *NullableChartResponseChartMetaCurrentTradingPeriod) Set(val *ChartResponseChartMetaCurrentTradingPeriod) {
	v.value = val
	v.isSet = true
}

func (v NullableChartResponseChartMetaCurrentTradingPeriod) IsSet() bool {
	return v.isSet
}

func (v *NullableChartResponseChartMetaCurrentTradingPeriod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChartResponseChartMetaCurrentTradingPeriod(val *ChartResponseChartMetaCurrentTradingPeriod) *NullableChartResponseChartMetaCurrentTradingPeriod {
	return &NullableChartResponseChartMetaCurrentTradingPeriod{value: val, isSet: true}
}

func (v NullableChartResponseChartMetaCurrentTradingPeriod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChartResponseChartMetaCurrentTradingPeriod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


