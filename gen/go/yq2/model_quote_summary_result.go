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

// QuoteSummaryResult struct for QuoteSummaryResult
type QuoteSummaryResult struct {
	EsgScores *QuoteSummaryEsgScores `json:"esgScores,omitempty"`
	DefaultKeyStatistics *QuoteSummaryDefaultKeyStatistics `json:"defaultKeyStatistics,omitempty"`
	Earnings *QuoteSummaryEarnings `json:"earnings,omitempty"`
	CalendarEvents *QuoteSummaryCalendarEvents `json:"calendarEvents,omitempty"`
	SummaryProfile *QuoteSummarySummaryProfile `json:"summaryProfile,omitempty"`
}

// NewQuoteSummaryResult instantiates a new QuoteSummaryResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuoteSummaryResult() *QuoteSummaryResult {
	this := QuoteSummaryResult{}
	return &this
}

// NewQuoteSummaryResultWithDefaults instantiates a new QuoteSummaryResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuoteSummaryResultWithDefaults() *QuoteSummaryResult {
	this := QuoteSummaryResult{}
	return &this
}

// GetEsgScores returns the EsgScores field value if set, zero value otherwise.
func (o *QuoteSummaryResult) GetEsgScores() QuoteSummaryEsgScores {
	if o == nil || o.EsgScores == nil {
		var ret QuoteSummaryEsgScores
		return ret
	}
	return *o.EsgScores
}

// GetEsgScoresOk returns a tuple with the EsgScores field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteSummaryResult) GetEsgScoresOk() (*QuoteSummaryEsgScores, bool) {
	if o == nil || o.EsgScores == nil {
		return nil, false
	}
	return o.EsgScores, true
}

// HasEsgScores returns a boolean if a field has been set.
func (o *QuoteSummaryResult) HasEsgScores() bool {
	if o != nil && o.EsgScores != nil {
		return true
	}

	return false
}

// SetEsgScores gets a reference to the given QuoteSummaryEsgScores and assigns it to the EsgScores field.
func (o *QuoteSummaryResult) SetEsgScores(v QuoteSummaryEsgScores) {
	o.EsgScores = &v
}

// GetDefaultKeyStatistics returns the DefaultKeyStatistics field value if set, zero value otherwise.
func (o *QuoteSummaryResult) GetDefaultKeyStatistics() QuoteSummaryDefaultKeyStatistics {
	if o == nil || o.DefaultKeyStatistics == nil {
		var ret QuoteSummaryDefaultKeyStatistics
		return ret
	}
	return *o.DefaultKeyStatistics
}

// GetDefaultKeyStatisticsOk returns a tuple with the DefaultKeyStatistics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteSummaryResult) GetDefaultKeyStatisticsOk() (*QuoteSummaryDefaultKeyStatistics, bool) {
	if o == nil || o.DefaultKeyStatistics == nil {
		return nil, false
	}
	return o.DefaultKeyStatistics, true
}

// HasDefaultKeyStatistics returns a boolean if a field has been set.
func (o *QuoteSummaryResult) HasDefaultKeyStatistics() bool {
	if o != nil && o.DefaultKeyStatistics != nil {
		return true
	}

	return false
}

// SetDefaultKeyStatistics gets a reference to the given QuoteSummaryDefaultKeyStatistics and assigns it to the DefaultKeyStatistics field.
func (o *QuoteSummaryResult) SetDefaultKeyStatistics(v QuoteSummaryDefaultKeyStatistics) {
	o.DefaultKeyStatistics = &v
}

// GetEarnings returns the Earnings field value if set, zero value otherwise.
func (o *QuoteSummaryResult) GetEarnings() QuoteSummaryEarnings {
	if o == nil || o.Earnings == nil {
		var ret QuoteSummaryEarnings
		return ret
	}
	return *o.Earnings
}

// GetEarningsOk returns a tuple with the Earnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteSummaryResult) GetEarningsOk() (*QuoteSummaryEarnings, bool) {
	if o == nil || o.Earnings == nil {
		return nil, false
	}
	return o.Earnings, true
}

// HasEarnings returns a boolean if a field has been set.
func (o *QuoteSummaryResult) HasEarnings() bool {
	if o != nil && o.Earnings != nil {
		return true
	}

	return false
}

// SetEarnings gets a reference to the given QuoteSummaryEarnings and assigns it to the Earnings field.
func (o *QuoteSummaryResult) SetEarnings(v QuoteSummaryEarnings) {
	o.Earnings = &v
}

// GetCalendarEvents returns the CalendarEvents field value if set, zero value otherwise.
func (o *QuoteSummaryResult) GetCalendarEvents() QuoteSummaryCalendarEvents {
	if o == nil || o.CalendarEvents == nil {
		var ret QuoteSummaryCalendarEvents
		return ret
	}
	return *o.CalendarEvents
}

// GetCalendarEventsOk returns a tuple with the CalendarEvents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteSummaryResult) GetCalendarEventsOk() (*QuoteSummaryCalendarEvents, bool) {
	if o == nil || o.CalendarEvents == nil {
		return nil, false
	}
	return o.CalendarEvents, true
}

// HasCalendarEvents returns a boolean if a field has been set.
func (o *QuoteSummaryResult) HasCalendarEvents() bool {
	if o != nil && o.CalendarEvents != nil {
		return true
	}

	return false
}

// SetCalendarEvents gets a reference to the given QuoteSummaryCalendarEvents and assigns it to the CalendarEvents field.
func (o *QuoteSummaryResult) SetCalendarEvents(v QuoteSummaryCalendarEvents) {
	o.CalendarEvents = &v
}

// GetSummaryProfile returns the SummaryProfile field value if set, zero value otherwise.
func (o *QuoteSummaryResult) GetSummaryProfile() QuoteSummarySummaryProfile {
	if o == nil || o.SummaryProfile == nil {
		var ret QuoteSummarySummaryProfile
		return ret
	}
	return *o.SummaryProfile
}

// GetSummaryProfileOk returns a tuple with the SummaryProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuoteSummaryResult) GetSummaryProfileOk() (*QuoteSummarySummaryProfile, bool) {
	if o == nil || o.SummaryProfile == nil {
		return nil, false
	}
	return o.SummaryProfile, true
}

// HasSummaryProfile returns a boolean if a field has been set.
func (o *QuoteSummaryResult) HasSummaryProfile() bool {
	if o != nil && o.SummaryProfile != nil {
		return true
	}

	return false
}

// SetSummaryProfile gets a reference to the given QuoteSummarySummaryProfile and assigns it to the SummaryProfile field.
func (o *QuoteSummaryResult) SetSummaryProfile(v QuoteSummarySummaryProfile) {
	o.SummaryProfile = &v
}

func (o QuoteSummaryResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EsgScores != nil {
		toSerialize["esgScores"] = o.EsgScores
	}
	if o.DefaultKeyStatistics != nil {
		toSerialize["defaultKeyStatistics"] = o.DefaultKeyStatistics
	}
	if o.Earnings != nil {
		toSerialize["earnings"] = o.Earnings
	}
	if o.CalendarEvents != nil {
		toSerialize["calendarEvents"] = o.CalendarEvents
	}
	if o.SummaryProfile != nil {
		toSerialize["summaryProfile"] = o.SummaryProfile
	}
	return json.Marshal(toSerialize)
}

type NullableQuoteSummaryResult struct {
	value *QuoteSummaryResult
	isSet bool
}

func (v NullableQuoteSummaryResult) Get() *QuoteSummaryResult {
	return v.value
}

func (v *NullableQuoteSummaryResult) Set(val *QuoteSummaryResult) {
	v.value = val
	v.isSet = true
}

func (v NullableQuoteSummaryResult) IsSet() bool {
	return v.isSet
}

func (v *NullableQuoteSummaryResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuoteSummaryResult(val *QuoteSummaryResult) *NullableQuoteSummaryResult {
	return &NullableQuoteSummaryResult{value: val, isSet: true}
}

func (v NullableQuoteSummaryResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuoteSummaryResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


