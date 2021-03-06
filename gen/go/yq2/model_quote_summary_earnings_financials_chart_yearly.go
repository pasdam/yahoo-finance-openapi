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

// QuoteSummaryEarningsFinancialsChartYearly struct for QuoteSummaryEarningsFinancialsChartYearly
type QuoteSummaryEarningsFinancialsChartYearly struct {
	Date *int32 `json:"date,omitempty"`
	Revenue *QuoteSummaryDefaultKeyStatisticsEnterpriseValue `json:"revenue,omitempty"`
	Earnings *QuoteSummaryDefaultKeyStatisticsEnterpriseValue `json:"earnings,omitempty"`
}

// NewQuoteSummaryEarningsFinancialsChartYearly instantiates a new QuoteSummaryEarningsFinancialsChartYearly object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuoteSummaryEarningsFinancialsChartYearly() *QuoteSummaryEarningsFinancialsChartYearly {
	this := QuoteSummaryEarningsFinancialsChartYearly{}
	return &this
}

// NewQuoteSummaryEarningsFinancialsChartYearlyWithDefaults instantiates a new QuoteSummaryEarningsFinancialsChartYearly object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuoteSummaryEarningsFinancialsChartYearlyWithDefaults() *QuoteSummaryEarningsFinancialsChartYearly {
	this := QuoteSummaryEarningsFinancialsChartYearly{}
	return &this
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *QuoteSummaryEarningsFinancialsChartYearly) GetDate() int32 {
	if o == nil || o.Date == nil {
		var ret int32
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteSummaryEarningsFinancialsChartYearly) GetDateOk() (*int32, bool) {
	if o == nil || o.Date == nil {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *QuoteSummaryEarningsFinancialsChartYearly) HasDate() bool {
	if o != nil && o.Date != nil {
		return true
	}

	return false
}

// SetDate gets a reference to the given int32 and assigns it to the Date field.
func (o *QuoteSummaryEarningsFinancialsChartYearly) SetDate(v int32) {
	o.Date = &v
}

// GetRevenue returns the Revenue field value if set, zero value otherwise.
func (o *QuoteSummaryEarningsFinancialsChartYearly) GetRevenue() QuoteSummaryDefaultKeyStatisticsEnterpriseValue {
	if o == nil || o.Revenue == nil {
		var ret QuoteSummaryDefaultKeyStatisticsEnterpriseValue
		return ret
	}
	return *o.Revenue
}

// GetRevenueOk returns a tuple with the Revenue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteSummaryEarningsFinancialsChartYearly) GetRevenueOk() (*QuoteSummaryDefaultKeyStatisticsEnterpriseValue, bool) {
	if o == nil || o.Revenue == nil {
		return nil, false
	}
	return o.Revenue, true
}

// HasRevenue returns a boolean if a field has been set.
func (o *QuoteSummaryEarningsFinancialsChartYearly) HasRevenue() bool {
	if o != nil && o.Revenue != nil {
		return true
	}

	return false
}

// SetRevenue gets a reference to the given QuoteSummaryDefaultKeyStatisticsEnterpriseValue and assigns it to the Revenue field.
func (o *QuoteSummaryEarningsFinancialsChartYearly) SetRevenue(v QuoteSummaryDefaultKeyStatisticsEnterpriseValue) {
	o.Revenue = &v
}

// GetEarnings returns the Earnings field value if set, zero value otherwise.
func (o *QuoteSummaryEarningsFinancialsChartYearly) GetEarnings() QuoteSummaryDefaultKeyStatisticsEnterpriseValue {
	if o == nil || o.Earnings == nil {
		var ret QuoteSummaryDefaultKeyStatisticsEnterpriseValue
		return ret
	}
	return *o.Earnings
}

// GetEarningsOk returns a tuple with the Earnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteSummaryEarningsFinancialsChartYearly) GetEarningsOk() (*QuoteSummaryDefaultKeyStatisticsEnterpriseValue, bool) {
	if o == nil || o.Earnings == nil {
		return nil, false
	}
	return o.Earnings, true
}

// HasEarnings returns a boolean if a field has been set.
func (o *QuoteSummaryEarningsFinancialsChartYearly) HasEarnings() bool {
	if o != nil && o.Earnings != nil {
		return true
	}

	return false
}

// SetEarnings gets a reference to the given QuoteSummaryDefaultKeyStatisticsEnterpriseValue and assigns it to the Earnings field.
func (o *QuoteSummaryEarningsFinancialsChartYearly) SetEarnings(v QuoteSummaryDefaultKeyStatisticsEnterpriseValue) {
	o.Earnings = &v
}

func (o QuoteSummaryEarningsFinancialsChartYearly) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Date != nil {
		toSerialize["date"] = o.Date
	}
	if o.Revenue != nil {
		toSerialize["revenue"] = o.Revenue
	}
	if o.Earnings != nil {
		toSerialize["earnings"] = o.Earnings
	}
	return json.Marshal(toSerialize)
}

type NullableQuoteSummaryEarningsFinancialsChartYearly struct {
	value *QuoteSummaryEarningsFinancialsChartYearly
	isSet bool
}

func (v NullableQuoteSummaryEarningsFinancialsChartYearly) Get() *QuoteSummaryEarningsFinancialsChartYearly {
	return v.value
}

func (v *NullableQuoteSummaryEarningsFinancialsChartYearly) Set(val *QuoteSummaryEarningsFinancialsChartYearly) {
	v.value = val
	v.isSet = true
}

func (v NullableQuoteSummaryEarningsFinancialsChartYearly) IsSet() bool {
	return v.isSet
}

func (v *NullableQuoteSummaryEarningsFinancialsChartYearly) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuoteSummaryEarningsFinancialsChartYearly(val *QuoteSummaryEarningsFinancialsChartYearly) *NullableQuoteSummaryEarningsFinancialsChartYearly {
	return &NullableQuoteSummaryEarningsFinancialsChartYearly{value: val, isSet: true}
}

func (v NullableQuoteSummaryEarningsFinancialsChartYearly) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuoteSummaryEarningsFinancialsChartYearly) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


