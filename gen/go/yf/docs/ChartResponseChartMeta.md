# ChartResponseChartMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**ExchangeName** | Pointer to **string** |  | [optional] 
**InstrumentType** | Pointer to **string** |  | [optional] 
**FirstTradeDate** | Pointer to **int32** |  | [optional] 
**RegularMarketTime** | Pointer to **int32** |  | [optional] 
**Gmtoffset** | Pointer to **int32** |  | [optional] 
**Timezone** | Pointer to **string** |  | [optional] 
**ExchangeTimezoneName** | Pointer to **string** |  | [optional] 
**RegularMarketPrice** | Pointer to **float32** |  | [optional] 
**ChartPreviousClose** | Pointer to **float32** |  | [optional] 
**PreviousClose** | Pointer to **float32** |  | [optional] 
**Scale** | Pointer to **int32** |  | [optional] 
**PriceHint** | Pointer to **int32** |  | [optional] 
**CurrentTradingPeriod** | Pointer to [**ChartResponseChartMetaCurrentTradingPeriod**](ChartResponseChartMetaCurrentTradingPeriod.md) |  | [optional] 
**TradingPeriods** | Pointer to **[][]map[string]interface{}** |  | [optional] 
**DataGranularity** | Pointer to **string** |  | [optional] 
**Range** | Pointer to **string** |  | [optional] 
**ValidRanges** | Pointer to **[]string** |  | [optional] 

## Methods

### NewChartResponseChartMeta

`func NewChartResponseChartMeta() *ChartResponseChartMeta`

NewChartResponseChartMeta instantiates a new ChartResponseChartMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChartResponseChartMetaWithDefaults

`func NewChartResponseChartMetaWithDefaults() *ChartResponseChartMeta`

NewChartResponseChartMetaWithDefaults instantiates a new ChartResponseChartMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *ChartResponseChartMeta) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ChartResponseChartMeta) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ChartResponseChartMeta) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ChartResponseChartMeta) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetSymbol

`func (o *ChartResponseChartMeta) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *ChartResponseChartMeta) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *ChartResponseChartMeta) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *ChartResponseChartMeta) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetExchangeName

`func (o *ChartResponseChartMeta) GetExchangeName() string`

GetExchangeName returns the ExchangeName field if non-nil, zero value otherwise.

### GetExchangeNameOk

`func (o *ChartResponseChartMeta) GetExchangeNameOk() (*string, bool)`

GetExchangeNameOk returns a tuple with the ExchangeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeName

`func (o *ChartResponseChartMeta) SetExchangeName(v string)`

SetExchangeName sets ExchangeName field to given value.

### HasExchangeName

`func (o *ChartResponseChartMeta) HasExchangeName() bool`

HasExchangeName returns a boolean if a field has been set.

### GetInstrumentType

`func (o *ChartResponseChartMeta) GetInstrumentType() string`

GetInstrumentType returns the InstrumentType field if non-nil, zero value otherwise.

### GetInstrumentTypeOk

`func (o *ChartResponseChartMeta) GetInstrumentTypeOk() (*string, bool)`

GetInstrumentTypeOk returns a tuple with the InstrumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentType

`func (o *ChartResponseChartMeta) SetInstrumentType(v string)`

SetInstrumentType sets InstrumentType field to given value.

### HasInstrumentType

`func (o *ChartResponseChartMeta) HasInstrumentType() bool`

HasInstrumentType returns a boolean if a field has been set.

### GetFirstTradeDate

`func (o *ChartResponseChartMeta) GetFirstTradeDate() int32`

GetFirstTradeDate returns the FirstTradeDate field if non-nil, zero value otherwise.

### GetFirstTradeDateOk

`func (o *ChartResponseChartMeta) GetFirstTradeDateOk() (*int32, bool)`

GetFirstTradeDateOk returns a tuple with the FirstTradeDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstTradeDate

`func (o *ChartResponseChartMeta) SetFirstTradeDate(v int32)`

SetFirstTradeDate sets FirstTradeDate field to given value.

### HasFirstTradeDate

`func (o *ChartResponseChartMeta) HasFirstTradeDate() bool`

HasFirstTradeDate returns a boolean if a field has been set.

### GetRegularMarketTime

`func (o *ChartResponseChartMeta) GetRegularMarketTime() int32`

GetRegularMarketTime returns the RegularMarketTime field if non-nil, zero value otherwise.

### GetRegularMarketTimeOk

`func (o *ChartResponseChartMeta) GetRegularMarketTimeOk() (*int32, bool)`

GetRegularMarketTimeOk returns a tuple with the RegularMarketTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegularMarketTime

`func (o *ChartResponseChartMeta) SetRegularMarketTime(v int32)`

SetRegularMarketTime sets RegularMarketTime field to given value.

### HasRegularMarketTime

`func (o *ChartResponseChartMeta) HasRegularMarketTime() bool`

HasRegularMarketTime returns a boolean if a field has been set.

### GetGmtoffset

`func (o *ChartResponseChartMeta) GetGmtoffset() int32`

GetGmtoffset returns the Gmtoffset field if non-nil, zero value otherwise.

### GetGmtoffsetOk

`func (o *ChartResponseChartMeta) GetGmtoffsetOk() (*int32, bool)`

GetGmtoffsetOk returns a tuple with the Gmtoffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmtoffset

`func (o *ChartResponseChartMeta) SetGmtoffset(v int32)`

SetGmtoffset sets Gmtoffset field to given value.

### HasGmtoffset

`func (o *ChartResponseChartMeta) HasGmtoffset() bool`

HasGmtoffset returns a boolean if a field has been set.

### GetTimezone

`func (o *ChartResponseChartMeta) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *ChartResponseChartMeta) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *ChartResponseChartMeta) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *ChartResponseChartMeta) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetExchangeTimezoneName

`func (o *ChartResponseChartMeta) GetExchangeTimezoneName() string`

GetExchangeTimezoneName returns the ExchangeTimezoneName field if non-nil, zero value otherwise.

### GetExchangeTimezoneNameOk

`func (o *ChartResponseChartMeta) GetExchangeTimezoneNameOk() (*string, bool)`

GetExchangeTimezoneNameOk returns a tuple with the ExchangeTimezoneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeTimezoneName

`func (o *ChartResponseChartMeta) SetExchangeTimezoneName(v string)`

SetExchangeTimezoneName sets ExchangeTimezoneName field to given value.

### HasExchangeTimezoneName

`func (o *ChartResponseChartMeta) HasExchangeTimezoneName() bool`

HasExchangeTimezoneName returns a boolean if a field has been set.

### GetRegularMarketPrice

`func (o *ChartResponseChartMeta) GetRegularMarketPrice() float32`

GetRegularMarketPrice returns the RegularMarketPrice field if non-nil, zero value otherwise.

### GetRegularMarketPriceOk

`func (o *ChartResponseChartMeta) GetRegularMarketPriceOk() (*float32, bool)`

GetRegularMarketPriceOk returns a tuple with the RegularMarketPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegularMarketPrice

`func (o *ChartResponseChartMeta) SetRegularMarketPrice(v float32)`

SetRegularMarketPrice sets RegularMarketPrice field to given value.

### HasRegularMarketPrice

`func (o *ChartResponseChartMeta) HasRegularMarketPrice() bool`

HasRegularMarketPrice returns a boolean if a field has been set.

### GetChartPreviousClose

`func (o *ChartResponseChartMeta) GetChartPreviousClose() float32`

GetChartPreviousClose returns the ChartPreviousClose field if non-nil, zero value otherwise.

### GetChartPreviousCloseOk

`func (o *ChartResponseChartMeta) GetChartPreviousCloseOk() (*float32, bool)`

GetChartPreviousCloseOk returns a tuple with the ChartPreviousClose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartPreviousClose

`func (o *ChartResponseChartMeta) SetChartPreviousClose(v float32)`

SetChartPreviousClose sets ChartPreviousClose field to given value.

### HasChartPreviousClose

`func (o *ChartResponseChartMeta) HasChartPreviousClose() bool`

HasChartPreviousClose returns a boolean if a field has been set.

### GetPreviousClose

`func (o *ChartResponseChartMeta) GetPreviousClose() float32`

GetPreviousClose returns the PreviousClose field if non-nil, zero value otherwise.

### GetPreviousCloseOk

`func (o *ChartResponseChartMeta) GetPreviousCloseOk() (*float32, bool)`

GetPreviousCloseOk returns a tuple with the PreviousClose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousClose

`func (o *ChartResponseChartMeta) SetPreviousClose(v float32)`

SetPreviousClose sets PreviousClose field to given value.

### HasPreviousClose

`func (o *ChartResponseChartMeta) HasPreviousClose() bool`

HasPreviousClose returns a boolean if a field has been set.

### GetScale

`func (o *ChartResponseChartMeta) GetScale() int32`

GetScale returns the Scale field if non-nil, zero value otherwise.

### GetScaleOk

`func (o *ChartResponseChartMeta) GetScaleOk() (*int32, bool)`

GetScaleOk returns a tuple with the Scale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScale

`func (o *ChartResponseChartMeta) SetScale(v int32)`

SetScale sets Scale field to given value.

### HasScale

`func (o *ChartResponseChartMeta) HasScale() bool`

HasScale returns a boolean if a field has been set.

### GetPriceHint

`func (o *ChartResponseChartMeta) GetPriceHint() int32`

GetPriceHint returns the PriceHint field if non-nil, zero value otherwise.

### GetPriceHintOk

`func (o *ChartResponseChartMeta) GetPriceHintOk() (*int32, bool)`

GetPriceHintOk returns a tuple with the PriceHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceHint

`func (o *ChartResponseChartMeta) SetPriceHint(v int32)`

SetPriceHint sets PriceHint field to given value.

### HasPriceHint

`func (o *ChartResponseChartMeta) HasPriceHint() bool`

HasPriceHint returns a boolean if a field has been set.

### GetCurrentTradingPeriod

`func (o *ChartResponseChartMeta) GetCurrentTradingPeriod() ChartResponseChartMetaCurrentTradingPeriod`

GetCurrentTradingPeriod returns the CurrentTradingPeriod field if non-nil, zero value otherwise.

### GetCurrentTradingPeriodOk

`func (o *ChartResponseChartMeta) GetCurrentTradingPeriodOk() (*ChartResponseChartMetaCurrentTradingPeriod, bool)`

GetCurrentTradingPeriodOk returns a tuple with the CurrentTradingPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentTradingPeriod

`func (o *ChartResponseChartMeta) SetCurrentTradingPeriod(v ChartResponseChartMetaCurrentTradingPeriod)`

SetCurrentTradingPeriod sets CurrentTradingPeriod field to given value.

### HasCurrentTradingPeriod

`func (o *ChartResponseChartMeta) HasCurrentTradingPeriod() bool`

HasCurrentTradingPeriod returns a boolean if a field has been set.

### GetTradingPeriods

`func (o *ChartResponseChartMeta) GetTradingPeriods() [][]map[string]interface{}`

GetTradingPeriods returns the TradingPeriods field if non-nil, zero value otherwise.

### GetTradingPeriodsOk

`func (o *ChartResponseChartMeta) GetTradingPeriodsOk() (*[][]map[string]interface{}, bool)`

GetTradingPeriodsOk returns a tuple with the TradingPeriods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradingPeriods

`func (o *ChartResponseChartMeta) SetTradingPeriods(v [][]map[string]interface{})`

SetTradingPeriods sets TradingPeriods field to given value.

### HasTradingPeriods

`func (o *ChartResponseChartMeta) HasTradingPeriods() bool`

HasTradingPeriods returns a boolean if a field has been set.

### GetDataGranularity

`func (o *ChartResponseChartMeta) GetDataGranularity() string`

GetDataGranularity returns the DataGranularity field if non-nil, zero value otherwise.

### GetDataGranularityOk

`func (o *ChartResponseChartMeta) GetDataGranularityOk() (*string, bool)`

GetDataGranularityOk returns a tuple with the DataGranularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataGranularity

`func (o *ChartResponseChartMeta) SetDataGranularity(v string)`

SetDataGranularity sets DataGranularity field to given value.

### HasDataGranularity

`func (o *ChartResponseChartMeta) HasDataGranularity() bool`

HasDataGranularity returns a boolean if a field has been set.

### GetRange

`func (o *ChartResponseChartMeta) GetRange() string`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *ChartResponseChartMeta) GetRangeOk() (*string, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *ChartResponseChartMeta) SetRange(v string)`

SetRange sets Range field to given value.

### HasRange

`func (o *ChartResponseChartMeta) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetValidRanges

`func (o *ChartResponseChartMeta) GetValidRanges() []string`

GetValidRanges returns the ValidRanges field if non-nil, zero value otherwise.

### GetValidRangesOk

`func (o *ChartResponseChartMeta) GetValidRangesOk() (*[]string, bool)`

GetValidRangesOk returns a tuple with the ValidRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidRanges

`func (o *ChartResponseChartMeta) SetValidRanges(v []string)`

SetValidRanges sets ValidRanges field to given value.

### HasValidRanges

`func (o *ChartResponseChartMeta) HasValidRanges() bool`

HasValidRanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


