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

// ChartMeta struct for ChartMeta
type ChartMeta struct {
	Currency *string `json:"currency,omitempty"`
	Symbol *string `json:"symbol,omitempty"`
	ExchangeName *string `json:"exchangeName,omitempty"`
	InstrumentType *string `json:"instrumentType,omitempty"`
	FirstTradeDate *int32 `json:"firstTradeDate,omitempty"`
	RegularMarketTime *int32 `json:"regularMarketTime,omitempty"`
	Gmtoffset *int32 `json:"gmtoffset,omitempty"`
	Timezone *string `json:"timezone,omitempty"`
	ExchangeTimezoneName *string `json:"exchangeTimezoneName,omitempty"`
	RegularMarketPrice *float32 `json:"regularMarketPrice,omitempty"`
	ChartPreviousClose *float32 `json:"chartPreviousClose,omitempty"`
	PreviousClose *float32 `json:"previousClose,omitempty"`
	Scale *int32 `json:"scale,omitempty"`
	PriceHint *int32 `json:"priceHint,omitempty"`
	CurrentTradingPeriod *ChartMetaCurrentTradingPeriod `json:"currentTradingPeriod,omitempty"`
	TradingPeriods *[][]map[string]interface{} `json:"tradingPeriods,omitempty"`
	DataGranularity *string `json:"dataGranularity,omitempty"`
	Range *string `json:"range,omitempty"`
	ValidRanges *[]string `json:"validRanges,omitempty"`
}

// NewChartMeta instantiates a new ChartMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChartMeta() *ChartMeta {
	this := ChartMeta{}
	return &this
}

// NewChartMetaWithDefaults instantiates a new ChartMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChartMetaWithDefaults() *ChartMeta {
	this := ChartMeta{}
	return &this
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *ChartMeta) GetCurrency() string {
	if o == nil || o.Currency == nil {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChartMeta) GetCurrencyOk() (*string, bool) {
	if o == nil || o.Currency == nil {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *ChartMeta) HasCurrency() bool {
	if o != nil && o.Currency != nil {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *ChartMeta) SetCurrency(v string) {
	o.Currency = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *ChartMeta) GetSymbol() string {
	if o == nil || o.Symbol == nil {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChartMeta) GetSymbolOk() (*string, bool) {
	if o == nil || o.Symbol == nil {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *ChartMeta) HasSymbol() bool {
	if o != nil && o.Symbol != nil {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *ChartMeta) SetSymbol(v string) {
	o.Symbol = &v
}

// GetExchangeName returns the ExchangeName field value if set, zero value otherwise.
func (o *ChartMeta) GetExchangeName() string {
	if o == nil || o.ExchangeName == nil {
		var ret string
		return ret
	}
	return *o.ExchangeName
}

// GetExchangeNameOk returns a tuple with the ExchangeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChartMeta) GetExchangeNameOk() (*string, bool) {
	if o == nil || o.ExchangeName == nil {
		return nil, false
	}
	return o.ExchangeName, true
}

// HasExchangeName returns a boolean if a field has been set.
func (o *ChartMeta) HasExchangeName() bool {
	if o != nil && o.ExchangeName != nil {
		return true
	}

	return false
}

// SetExchangeName gets a reference to the given string and assigns it to the ExchangeName field.
func (o *ChartMeta) SetExchangeName(v string) {
	o.ExchangeName = &v
}

// GetInstrumentType returns the InstrumentType field value if set, zero value otherwise.
func (o *ChartMeta) GetInstrumentType() string {
	if o == nil || o.InstrumentType == nil {
		var ret string
		return ret
	}
	return *o.InstrumentType
}

// GetInstrumentTypeOk returns a tuple with the InstrumentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChartMeta) GetInstrumentTypeOk() (*string, bool) {
	if o == nil || o.InstrumentType == nil {
		return nil, false
	}
	return o.InstrumentType, true
}

// HasInstrumentType returns a boolean if a field has been set.
func (o *ChartMeta) HasInstrumentType() bool {
	if o != nil && o.InstrumentType != nil {
		return true
	}

	return false
}

// SetInstrumentType gets a reference to the given string and assigns it to the InstrumentType field.
func (o *ChartMeta) SetInstrumentType(v string) {
	o.InstrumentType = &v
}

// GetFirstTradeDate returns the FirstTradeDate field value if set, zero value otherwise.
func (o *ChartMeta) GetFirstTradeDate() int32 {
	if o == nil || o.FirstTradeDate == nil {
		var ret int32
		return ret
	}
	return *o.FirstTradeDate
}

// GetFirstTradeDateOk returns a tuple with the FirstTradeDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChartMeta) GetFirstTradeDateOk() (*int32, bool) {
	if o == nil || o.FirstTradeDate == nil {
		return nil, false
	}
	return o.FirstTradeDate, true
}

// HasFirstTradeDate returns a boolean if a field has been set.
func (o *ChartMeta) HasFirstTradeDate() bool {
	if o != nil && o.FirstTradeDate != nil {
		return true
	}

	return false
}

// SetFirstTradeDate gets a reference to the given int32 and assigns it to the FirstTradeDate field.
func (o *ChartMeta) SetFirstTradeDate(v int32) {
	o.FirstTradeDate = &v
}

// GetRegularMarketTime returns the RegularMarketTime field value if set, zero value otherwise.
func (o *ChartMeta) GetRegularMarketTime() int32 {
	if o == nil || o.RegularMarketTime == nil {
		var ret int32
		return ret
	}
	return *o.RegularMarketTime
}

// GetRegularMarketTimeOk returns a tuple with the RegularMarketTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChartMeta) GetRegularMarketTimeOk() (*int32, bool) {
	if o == nil || o.RegularMarketTime == nil {
		return nil, false
	}
	return o.RegularMarketTime, true
}

// HasRegularMarketTime returns a boolean if a field has been set.
func (o *ChartMeta) HasRegularMarketTime() bool {
	if o != nil && o.RegularMarketTime != nil {
		return true
	}

	return false
}

// SetRegularMarketTime gets a reference to the given int32 and assigns it to the RegularMarketTime field.
func (o *ChartMeta) SetRegularMarketTime(v int32) {
	o.RegularMarketTime = &v
}

// GetGmtoffset returns the Gmtoffset field value if set, zero value otherwise.
func (o *ChartMeta) GetGmtoffset() int32 {
	if o == nil || o.Gmtoffset == nil {
		var ret int32
		return ret
	}
	return *o.Gmtoffset
}

// GetGmtoffsetOk returns a tuple with the Gmtoffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChartMeta) GetGmtoffsetOk() (*int32, bool) {
	if o == nil || o.Gmtoffset == nil {
		return nil, false
	}
	return o.Gmtoffset, true
}

// HasGmtoffset returns a boolean if a field has been set.
func (o *ChartMeta) HasGmtoffset() bool {
	if o != nil && o.Gmtoffset != nil {
		return true
	}

	return false
}

// SetGmtoffset gets a reference to the given int32 and assigns it to the Gmtoffset field.
func (o *ChartMeta) SetGmtoffset(v int32) {
	o.Gmtoffset = &v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *ChartMeta) GetTimezone() string {
	if o == nil || o.Timezone == nil {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChartMeta) GetTimezoneOk() (*string, bool) {
	if o == nil || o.Timezone == nil {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *ChartMeta) HasTimezone() bool {
	if o != nil && o.Timezone != nil {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *ChartMeta) SetTimezone(v string) {
	o.Timezone = &v
}

// GetExchangeTimezoneName returns the ExchangeTimezoneName field value if set, zero value otherwise.
func (o *ChartMeta) GetExchangeTimezoneName() string {
	if o == nil || o.ExchangeTimezoneName == nil {
		var ret string
		return ret
	}
	return *o.ExchangeTimezoneName
}

// GetExchangeTimezoneNameOk returns a tuple with the ExchangeTimezoneName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChartMeta) GetExchangeTimezoneNameOk() (*string, bool) {
	if o == nil || o.ExchangeTimezoneName == nil {
		return nil, false
	}
	return o.ExchangeTimezoneName, true
}

// HasExchangeTimezoneName returns a boolean if a field has been set.
func (o *ChartMeta) HasExchangeTimezoneName() bool {
	if o != nil && o.ExchangeTimezoneName != nil {
		return true
	}

	return false
}

// SetExchangeTimezoneName gets a reference to the given string and assigns it to the ExchangeTimezoneName field.
func (o *ChartMeta) SetExchangeTimezoneName(v string) {
	o.ExchangeTimezoneName = &v
}

// GetRegularMarketPrice returns the RegularMarketPrice field value if set, zero value otherwise.
func (o *ChartMeta) GetRegularMarketPrice() float32 {
	if o == nil || o.RegularMarketPrice == nil {
		var ret float32
		return ret
	}
	return *o.RegularMarketPrice
}

// GetRegularMarketPriceOk returns a tuple with the RegularMarketPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChartMeta) GetRegularMarketPriceOk() (*float32, bool) {
	if o == nil || o.RegularMarketPrice == nil {
		return nil, false
	}
	return o.RegularMarketPrice, true
}

// HasRegularMarketPrice returns a boolean if a field has been set.
func (o *ChartMeta) HasRegularMarketPrice() bool {
	if o != nil && o.RegularMarketPrice != nil {
		return true
	}

	return false
}

// SetRegularMarketPrice gets a reference to the given float32 and assigns it to the RegularMarketPrice field.
func (o *ChartMeta) SetRegularMarketPrice(v float32) {
	o.RegularMarketPrice = &v
}

// GetChartPreviousClose returns the ChartPreviousClose field value if set, zero value otherwise.
func (o *ChartMeta) GetChartPreviousClose() float32 {
	if o == nil || o.ChartPreviousClose == nil {
		var ret float32
		return ret
	}
	return *o.ChartPreviousClose
}

// GetChartPreviousCloseOk returns a tuple with the ChartPreviousClose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChartMeta) GetChartPreviousCloseOk() (*float32, bool) {
	if o == nil || o.ChartPreviousClose == nil {
		return nil, false
	}
	return o.ChartPreviousClose, true
}

// HasChartPreviousClose returns a boolean if a field has been set.
func (o *ChartMeta) HasChartPreviousClose() bool {
	if o != nil && o.ChartPreviousClose != nil {
		return true
	}

	return false
}

// SetChartPreviousClose gets a reference to the given float32 and assigns it to the ChartPreviousClose field.
func (o *ChartMeta) SetChartPreviousClose(v float32) {
	o.ChartPreviousClose = &v
}

// GetPreviousClose returns the PreviousClose field value if set, zero value otherwise.
func (o *ChartMeta) GetPreviousClose() float32 {
	if o == nil || o.PreviousClose == nil {
		var ret float32
		return ret
	}
	return *o.PreviousClose
}

// GetPreviousCloseOk returns a tuple with the PreviousClose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChartMeta) GetPreviousCloseOk() (*float32, bool) {
	if o == nil || o.PreviousClose == nil {
		return nil, false
	}
	return o.PreviousClose, true
}

// HasPreviousClose returns a boolean if a field has been set.
func (o *ChartMeta) HasPreviousClose() bool {
	if o != nil && o.PreviousClose != nil {
		return true
	}

	return false
}

// SetPreviousClose gets a reference to the given float32 and assigns it to the PreviousClose field.
func (o *ChartMeta) SetPreviousClose(v float32) {
	o.PreviousClose = &v
}

// GetScale returns the Scale field value if set, zero value otherwise.
func (o *ChartMeta) GetScale() int32 {
	if o == nil || o.Scale == nil {
		var ret int32
		return ret
	}
	return *o.Scale
}

// GetScaleOk returns a tuple with the Scale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChartMeta) GetScaleOk() (*int32, bool) {
	if o == nil || o.Scale == nil {
		return nil, false
	}
	return o.Scale, true
}

// HasScale returns a boolean if a field has been set.
func (o *ChartMeta) HasScale() bool {
	if o != nil && o.Scale != nil {
		return true
	}

	return false
}

// SetScale gets a reference to the given int32 and assigns it to the Scale field.
func (o *ChartMeta) SetScale(v int32) {
	o.Scale = &v
}

// GetPriceHint returns the PriceHint field value if set, zero value otherwise.
func (o *ChartMeta) GetPriceHint() int32 {
	if o == nil || o.PriceHint == nil {
		var ret int32
		return ret
	}
	return *o.PriceHint
}

// GetPriceHintOk returns a tuple with the PriceHint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChartMeta) GetPriceHintOk() (*int32, bool) {
	if o == nil || o.PriceHint == nil {
		return nil, false
	}
	return o.PriceHint, true
}

// HasPriceHint returns a boolean if a field has been set.
func (o *ChartMeta) HasPriceHint() bool {
	if o != nil && o.PriceHint != nil {
		return true
	}

	return false
}

// SetPriceHint gets a reference to the given int32 and assigns it to the PriceHint field.
func (o *ChartMeta) SetPriceHint(v int32) {
	o.PriceHint = &v
}

// GetCurrentTradingPeriod returns the CurrentTradingPeriod field value if set, zero value otherwise.
func (o *ChartMeta) GetCurrentTradingPeriod() ChartMetaCurrentTradingPeriod {
	if o == nil || o.CurrentTradingPeriod == nil {
		var ret ChartMetaCurrentTradingPeriod
		return ret
	}
	return *o.CurrentTradingPeriod
}

// GetCurrentTradingPeriodOk returns a tuple with the CurrentTradingPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChartMeta) GetCurrentTradingPeriodOk() (*ChartMetaCurrentTradingPeriod, bool) {
	if o == nil || o.CurrentTradingPeriod == nil {
		return nil, false
	}
	return o.CurrentTradingPeriod, true
}

// HasCurrentTradingPeriod returns a boolean if a field has been set.
func (o *ChartMeta) HasCurrentTradingPeriod() bool {
	if o != nil && o.CurrentTradingPeriod != nil {
		return true
	}

	return false
}

// SetCurrentTradingPeriod gets a reference to the given ChartMetaCurrentTradingPeriod and assigns it to the CurrentTradingPeriod field.
func (o *ChartMeta) SetCurrentTradingPeriod(v ChartMetaCurrentTradingPeriod) {
	o.CurrentTradingPeriod = &v
}

// GetTradingPeriods returns the TradingPeriods field value if set, zero value otherwise.
func (o *ChartMeta) GetTradingPeriods() [][]map[string]interface{} {
	if o == nil || o.TradingPeriods == nil {
		var ret [][]map[string]interface{}
		return ret
	}
	return *o.TradingPeriods
}

// GetTradingPeriodsOk returns a tuple with the TradingPeriods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChartMeta) GetTradingPeriodsOk() (*[][]map[string]interface{}, bool) {
	if o == nil || o.TradingPeriods == nil {
		return nil, false
	}
	return o.TradingPeriods, true
}

// HasTradingPeriods returns a boolean if a field has been set.
func (o *ChartMeta) HasTradingPeriods() bool {
	if o != nil && o.TradingPeriods != nil {
		return true
	}

	return false
}

// SetTradingPeriods gets a reference to the given [][]map[string]interface{} and assigns it to the TradingPeriods field.
func (o *ChartMeta) SetTradingPeriods(v [][]map[string]interface{}) {
	o.TradingPeriods = &v
}

// GetDataGranularity returns the DataGranularity field value if set, zero value otherwise.
func (o *ChartMeta) GetDataGranularity() string {
	if o == nil || o.DataGranularity == nil {
		var ret string
		return ret
	}
	return *o.DataGranularity
}

// GetDataGranularityOk returns a tuple with the DataGranularity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChartMeta) GetDataGranularityOk() (*string, bool) {
	if o == nil || o.DataGranularity == nil {
		return nil, false
	}
	return o.DataGranularity, true
}

// HasDataGranularity returns a boolean if a field has been set.
func (o *ChartMeta) HasDataGranularity() bool {
	if o != nil && o.DataGranularity != nil {
		return true
	}

	return false
}

// SetDataGranularity gets a reference to the given string and assigns it to the DataGranularity field.
func (o *ChartMeta) SetDataGranularity(v string) {
	o.DataGranularity = &v
}

// GetRange returns the Range field value if set, zero value otherwise.
func (o *ChartMeta) GetRange() string {
	if o == nil || o.Range == nil {
		var ret string
		return ret
	}
	return *o.Range
}

// GetRangeOk returns a tuple with the Range field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChartMeta) GetRangeOk() (*string, bool) {
	if o == nil || o.Range == nil {
		return nil, false
	}
	return o.Range, true
}

// HasRange returns a boolean if a field has been set.
func (o *ChartMeta) HasRange() bool {
	if o != nil && o.Range != nil {
		return true
	}

	return false
}

// SetRange gets a reference to the given string and assigns it to the Range field.
func (o *ChartMeta) SetRange(v string) {
	o.Range = &v
}

// GetValidRanges returns the ValidRanges field value if set, zero value otherwise.
func (o *ChartMeta) GetValidRanges() []string {
	if o == nil || o.ValidRanges == nil {
		var ret []string
		return ret
	}
	return *o.ValidRanges
}

// GetValidRangesOk returns a tuple with the ValidRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChartMeta) GetValidRangesOk() (*[]string, bool) {
	if o == nil || o.ValidRanges == nil {
		return nil, false
	}
	return o.ValidRanges, true
}

// HasValidRanges returns a boolean if a field has been set.
func (o *ChartMeta) HasValidRanges() bool {
	if o != nil && o.ValidRanges != nil {
		return true
	}

	return false
}

// SetValidRanges gets a reference to the given []string and assigns it to the ValidRanges field.
func (o *ChartMeta) SetValidRanges(v []string) {
	o.ValidRanges = &v
}

func (o ChartMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Currency != nil {
		toSerialize["currency"] = o.Currency
	}
	if o.Symbol != nil {
		toSerialize["symbol"] = o.Symbol
	}
	if o.ExchangeName != nil {
		toSerialize["exchangeName"] = o.ExchangeName
	}
	if o.InstrumentType != nil {
		toSerialize["instrumentType"] = o.InstrumentType
	}
	if o.FirstTradeDate != nil {
		toSerialize["firstTradeDate"] = o.FirstTradeDate
	}
	if o.RegularMarketTime != nil {
		toSerialize["regularMarketTime"] = o.RegularMarketTime
	}
	if o.Gmtoffset != nil {
		toSerialize["gmtoffset"] = o.Gmtoffset
	}
	if o.Timezone != nil {
		toSerialize["timezone"] = o.Timezone
	}
	if o.ExchangeTimezoneName != nil {
		toSerialize["exchangeTimezoneName"] = o.ExchangeTimezoneName
	}
	if o.RegularMarketPrice != nil {
		toSerialize["regularMarketPrice"] = o.RegularMarketPrice
	}
	if o.ChartPreviousClose != nil {
		toSerialize["chartPreviousClose"] = o.ChartPreviousClose
	}
	if o.PreviousClose != nil {
		toSerialize["previousClose"] = o.PreviousClose
	}
	if o.Scale != nil {
		toSerialize["scale"] = o.Scale
	}
	if o.PriceHint != nil {
		toSerialize["priceHint"] = o.PriceHint
	}
	if o.CurrentTradingPeriod != nil {
		toSerialize["currentTradingPeriod"] = o.CurrentTradingPeriod
	}
	if o.TradingPeriods != nil {
		toSerialize["tradingPeriods"] = o.TradingPeriods
	}
	if o.DataGranularity != nil {
		toSerialize["dataGranularity"] = o.DataGranularity
	}
	if o.Range != nil {
		toSerialize["range"] = o.Range
	}
	if o.ValidRanges != nil {
		toSerialize["validRanges"] = o.ValidRanges
	}
	return json.Marshal(toSerialize)
}

type NullableChartMeta struct {
	value *ChartMeta
	isSet bool
}

func (v NullableChartMeta) Get() *ChartMeta {
	return v.value
}

func (v *NullableChartMeta) Set(val *ChartMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableChartMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableChartMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChartMeta(val *ChartMeta) *NullableChartMeta {
	return &NullableChartMeta{value: val, isSet: true}
}

func (v NullableChartMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChartMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


