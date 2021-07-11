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

// QuoteSummaryEarningsFinancialsChartQuarterly struct for QuoteSummaryEarningsFinancialsChartQuarterly
type QuoteSummaryEarningsFinancialsChartQuarterly struct {
	Date *string `json:"date,omitempty"`
	Revenue *QuoteSummaryDefaultKeyStatisticsEnterpriseValue `json:"revenue,omitempty"`
	Earnings *QuoteSummaryDefaultKeyStatisticsEnterpriseValue `json:"earnings,omitempty"`
}

// NewQuoteSummaryEarningsFinancialsChartQuarterly instantiates a new QuoteSummaryEarningsFinancialsChartQuarterly object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuoteSummaryEarningsFinancialsChartQuarterly() *QuoteSummaryEarningsFinancialsChartQuarterly {
	this := QuoteSummaryEarningsFinancialsChartQuarterly{}
	return &this
}

// NewQuoteSummaryEarningsFinancialsChartQuarterlyWithDefaults instantiates a new QuoteSummaryEarningsFinancialsChartQuarterly object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuoteSummaryEarningsFinancialsChartQuarterlyWithDefaults() *QuoteSummaryEarningsFinancialsChartQuarterly {
	this := QuoteSummaryEarningsFinancialsChartQuarterly{}
	return &this
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *QuoteSummaryEarningsFinancialsChartQuarterly) GetDate() string {
	if o == nil || o.Date == nil {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteSummaryEarningsFinancialsChartQuarterly) GetDateOk() (*string, bool) {
	if o == nil || o.Date == nil {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *QuoteSummaryEarningsFinancialsChartQuarterly) HasDate() bool {
	if o != nil && o.Date != nil {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *QuoteSummaryEarningsFinancialsChartQuarterly) SetDate(v string) {
	o.Date = &v
}

// GetRevenue returns the Revenue field value if set, zero value otherwise.
func (o *QuoteSummaryEarningsFinancialsChartQuarterly) GetRevenue() QuoteSummaryDefaultKeyStatisticsEnterpriseValue {
	if o == nil || o.Revenue == nil {
		var ret QuoteSummaryDefaultKeyStatisticsEnterpriseValue
		return ret
	}
	return *o.Revenue
}

// GetRevenueOk returns a tuple with the Revenue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteSummaryEarningsFinancialsChartQuarterly) GetRevenueOk() (*QuoteSummaryDefaultKeyStatisticsEnterpriseValue, bool) {
	if o == nil || o.Revenue == nil {
		return nil, false
	}
	return o.Revenue, true
}

// HasRevenue returns a boolean if a field has been set.
func (o *QuoteSummaryEarningsFinancialsChartQuarterly) HasRevenue() bool {
	if o != nil && o.Revenue != nil {
		return true
	}

	return false
}

// SetRevenue gets a reference to the given QuoteSummaryDefaultKeyStatisticsEnterpriseValue and assigns it to the Revenue field.
func (o *QuoteSummaryEarningsFinancialsChartQuarterly) SetRevenue(v QuoteSummaryDefaultKeyStatisticsEnterpriseValue) {
	o.Revenue = &v
}

// GetEarnings returns the Earnings field value if set, zero value otherwise.
func (o *QuoteSummaryEarningsFinancialsChartQuarterly) GetEarnings() QuoteSummaryDefaultKeyStatisticsEnterpriseValue {
	if o == nil || o.Earnings == nil {
		var ret QuoteSummaryDefaultKeyStatisticsEnterpriseValue
		return ret
	}
	return *o.Earnings
}

// GetEarningsOk returns a tuple with the Earnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteSummaryEarningsFinancialsChartQuarterly) GetEarningsOk() (*QuoteSummaryDefaultKeyStatisticsEnterpriseValue, bool) {
	if o == nil || o.Earnings == nil {
		return nil, false
	}
	return o.Earnings, true
}

// HasEarnings returns a boolean if a field has been set.
func (o *QuoteSummaryEarningsFinancialsChartQuarterly) HasEarnings() bool {
	if o != nil && o.Earnings != nil {
		return true
	}

	return false
}

// SetEarnings gets a reference to the given QuoteSummaryDefaultKeyStatisticsEnterpriseValue and assigns it to the Earnings field.
func (o *QuoteSummaryEarningsFinancialsChartQuarterly) SetEarnings(v QuoteSummaryDefaultKeyStatisticsEnterpriseValue) {
	o.Earnings = &v
}

func (o QuoteSummaryEarningsFinancialsChartQuarterly) MarshalJSON() ([]byte, error) {
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

type NullableQuoteSummaryEarningsFinancialsChartQuarterly struct {
	value *QuoteSummaryEarningsFinancialsChartQuarterly
	isSet bool
}

func (v NullableQuoteSummaryEarningsFinancialsChartQuarterly) Get() *QuoteSummaryEarningsFinancialsChartQuarterly {
	return v.value
}

func (v *NullableQuoteSummaryEarningsFinancialsChartQuarterly) Set(val *QuoteSummaryEarningsFinancialsChartQuarterly) {
	v.value = val
	v.isSet = true
}

func (v NullableQuoteSummaryEarningsFinancialsChartQuarterly) IsSet() bool {
	return v.isSet
}

func (v *NullableQuoteSummaryEarningsFinancialsChartQuarterly) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuoteSummaryEarningsFinancialsChartQuarterly(val *QuoteSummaryEarningsFinancialsChartQuarterly) *NullableQuoteSummaryEarningsFinancialsChartQuarterly {
	return &NullableQuoteSummaryEarningsFinancialsChartQuarterly{value: val, isSet: true}
}

func (v NullableQuoteSummaryEarningsFinancialsChartQuarterly) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuoteSummaryEarningsFinancialsChartQuarterly) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


