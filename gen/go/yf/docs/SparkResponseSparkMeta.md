# SparkResponseSparkMeta

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
**CurrentTradingPeriod** | Pointer to [**SparkResponseSparkMetaCurrentTradingPeriod**](SparkResponseSparkMetaCurrentTradingPeriod.md) |  | [optional] 
**TradingPeriods** | Pointer to **[][]map[string]interface{}** |  | [optional] 
**DataGranularity** | Pointer to **string** |  | [optional] 
**Range** | Pointer to **string** |  | [optional] 
**ValidRanges** | Pointer to **[]string** |  | [optional] 

## Methods

### NewSparkResponseSparkMeta

`func NewSparkResponseSparkMeta() *SparkResponseSparkMeta`

NewSparkResponseSparkMeta instantiates a new SparkResponseSparkMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSparkResponseSparkMetaWithDefaults

`func NewSparkResponseSparkMetaWithDefaults() *SparkResponseSparkMeta`

NewSparkResponseSparkMetaWithDefaults instantiates a new SparkResponseSparkMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *SparkResponseSparkMeta) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *SparkResponseSparkMeta) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *SparkResponseSparkMeta) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *SparkResponseSparkMeta) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetSymbol

`func (o *SparkResponseSparkMeta) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *SparkResponseSparkMeta) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *SparkResponseSparkMeta) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *SparkResponseSparkMeta) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetExchangeName

`func (o *SparkResponseSparkMeta) GetExchangeName() string`

GetExchangeName returns the ExchangeName field if non-nil, zero value otherwise.

### GetExchangeNameOk

`func (o *SparkResponseSparkMeta) GetExchangeNameOk() (*string, bool)`

GetExchangeNameOk returns a tuple with the ExchangeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeName

`func (o *SparkResponseSparkMeta) SetExchangeName(v string)`

SetExchangeName sets ExchangeName field to given value.

### HasExchangeName

`func (o *SparkResponseSparkMeta) HasExchangeName() bool`

HasExchangeName returns a boolean if a field has been set.

### GetInstrumentType

`func (o *SparkResponseSparkMeta) GetInstrumentType() string`

GetInstrumentType returns the InstrumentType field if non-nil, zero value otherwise.

### GetInstrumentTypeOk

`func (o *SparkResponseSparkMeta) GetInstrumentTypeOk() (*string, bool)`

GetInstrumentTypeOk returns a tuple with the InstrumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentType

`func (o *SparkResponseSparkMeta) SetInstrumentType(v string)`

SetInstrumentType sets InstrumentType field to given value.

### HasInstrumentType

`func (o *SparkResponseSparkMeta) HasInstrumentType() bool`

HasInstrumentType returns a boolean if a field has been set.

### GetFirstTradeDate

`func (o *SparkResponseSparkMeta) GetFirstTradeDate() int32`

GetFirstTradeDate returns the FirstTradeDate field if non-nil, zero value otherwise.

### GetFirstTradeDateOk

`func (o *SparkResponseSparkMeta) GetFirstTradeDateOk() (*int32, bool)`

GetFirstTradeDateOk returns a tuple with the FirstTradeDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstTradeDate

`func (o *SparkResponseSparkMeta) SetFirstTradeDate(v int32)`

SetFirstTradeDate sets FirstTradeDate field to given value.

### HasFirstTradeDate

`func (o *SparkResponseSparkMeta) HasFirstTradeDate() bool`

HasFirstTradeDate returns a boolean if a field has been set.

### GetRegularMarketTime

`func (o *SparkResponseSparkMeta) GetRegularMarketTime() int32`

GetRegularMarketTime returns the RegularMarketTime field if non-nil, zero value otherwise.

### GetRegularMarketTimeOk

`func (o *SparkResponseSparkMeta) GetRegularMarketTimeOk() (*int32, bool)`

GetRegularMarketTimeOk returns a tuple with the RegularMarketTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegularMarketTime

`func (o *SparkResponseSparkMeta) SetRegularMarketTime(v int32)`

SetRegularMarketTime sets RegularMarketTime field to given value.

### HasRegularMarketTime

`func (o *SparkResponseSparkMeta) HasRegularMarketTime() bool`

HasRegularMarketTime returns a boolean if a field has been set.

### GetGmtoffset

`func (o *SparkResponseSparkMeta) GetGmtoffset() int32`

GetGmtoffset returns the Gmtoffset field if non-nil, zero value otherwise.

### GetGmtoffsetOk

`func (o *SparkResponseSparkMeta) GetGmtoffsetOk() (*int32, bool)`

GetGmtoffsetOk returns a tuple with the Gmtoffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmtoffset

`func (o *SparkResponseSparkMeta) SetGmtoffset(v int32)`

SetGmtoffset sets Gmtoffset field to given value.

### HasGmtoffset

`func (o *SparkResponseSparkMeta) HasGmtoffset() bool`

HasGmtoffset returns a boolean if a field has been set.

### GetTimezone

`func (o *SparkResponseSparkMeta) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *SparkResponseSparkMeta) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *SparkResponseSparkMeta) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *SparkResponseSparkMeta) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetExchangeTimezoneName

`func (o *SparkResponseSparkMeta) GetExchangeTimezoneName() string`

GetExchangeTimezoneName returns the ExchangeTimezoneName field if non-nil, zero value otherwise.

### GetExchangeTimezoneNameOk

`func (o *SparkResponseSparkMeta) GetExchangeTimezoneNameOk() (*string, bool)`

GetExchangeTimezoneNameOk returns a tuple with the ExchangeTimezoneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeTimezoneName

`func (o *SparkResponseSparkMeta) SetExchangeTimezoneName(v string)`

SetExchangeTimezoneName sets ExchangeTimezoneName field to given value.

### HasExchangeTimezoneName

`func (o *SparkResponseSparkMeta) HasExchangeTimezoneName() bool`

HasExchangeTimezoneName returns a boolean if a field has been set.

### GetRegularMarketPrice

`func (o *SparkResponseSparkMeta) GetRegularMarketPrice() float32`

GetRegularMarketPrice returns the RegularMarketPrice field if non-nil, zero value otherwise.

### GetRegularMarketPriceOk

`func (o *SparkResponseSparkMeta) GetRegularMarketPriceOk() (*float32, bool)`

GetRegularMarketPriceOk returns a tuple with the RegularMarketPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegularMarketPrice

`func (o *SparkResponseSparkMeta) SetRegularMarketPrice(v float32)`

SetRegularMarketPrice sets RegularMarketPrice field to given value.

### HasRegularMarketPrice

`func (o *SparkResponseSparkMeta) HasRegularMarketPrice() bool`

HasRegularMarketPrice returns a boolean if a field has been set.

### GetChartPreviousClose

`func (o *SparkResponseSparkMeta) GetChartPreviousClose() float32`

GetChartPreviousClose returns the ChartPreviousClose field if non-nil, zero value otherwise.

### GetChartPreviousCloseOk

`func (o *SparkResponseSparkMeta) GetChartPreviousCloseOk() (*float32, bool)`

GetChartPreviousCloseOk returns a tuple with the ChartPreviousClose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChartPreviousClose

`func (o *SparkResponseSparkMeta) SetChartPreviousClose(v float32)`

SetChartPreviousClose sets ChartPreviousClose field to given value.

### HasChartPreviousClose

`func (o *SparkResponseSparkMeta) HasChartPreviousClose() bool`

HasChartPreviousClose returns a boolean if a field has been set.

### GetPreviousClose

`func (o *SparkResponseSparkMeta) GetPreviousClose() float32`

GetPreviousClose returns the PreviousClose field if non-nil, zero value otherwise.

### GetPreviousCloseOk

`func (o *SparkResponseSparkMeta) GetPreviousCloseOk() (*float32, bool)`

GetPreviousCloseOk returns a tuple with the PreviousClose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousClose

`func (o *SparkResponseSparkMeta) SetPreviousClose(v float32)`

SetPreviousClose sets PreviousClose field to given value.

### HasPreviousClose

`func (o *SparkResponseSparkMeta) HasPreviousClose() bool`

HasPreviousClose returns a boolean if a field has been set.

### GetScale

`func (o *SparkResponseSparkMeta) GetScale() int32`

GetScale returns the Scale field if non-nil, zero value otherwise.

### GetScaleOk

`func (o *SparkResponseSparkMeta) GetScaleOk() (*int32, bool)`

GetScaleOk returns a tuple with the Scale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScale

`func (o *SparkResponseSparkMeta) SetScale(v int32)`

SetScale sets Scale field to given value.

### HasScale

`func (o *SparkResponseSparkMeta) HasScale() bool`

HasScale returns a boolean if a field has been set.

### GetPriceHint

`func (o *SparkResponseSparkMeta) GetPriceHint() int32`

GetPriceHint returns the PriceHint field if non-nil, zero value otherwise.

### GetPriceHintOk

`func (o *SparkResponseSparkMeta) GetPriceHintOk() (*int32, bool)`

GetPriceHintOk returns a tuple with the PriceHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceHint

`func (o *SparkResponseSparkMeta) SetPriceHint(v int32)`

SetPriceHint sets PriceHint field to given value.

### HasPriceHint

`func (o *SparkResponseSparkMeta) HasPriceHint() bool`

HasPriceHint returns a boolean if a field has been set.

### GetCurrentTradingPeriod

`func (o *SparkResponseSparkMeta) GetCurrentTradingPeriod() SparkResponseSparkMetaCurrentTradingPeriod`

GetCurrentTradingPeriod returns the CurrentTradingPeriod field if non-nil, zero value otherwise.

### GetCurrentTradingPeriodOk

`func (o *SparkResponseSparkMeta) GetCurrentTradingPeriodOk() (*SparkResponseSparkMetaCurrentTradingPeriod, bool)`

GetCurrentTradingPeriodOk returns a tuple with the CurrentTradingPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentTradingPeriod

`func (o *SparkResponseSparkMeta) SetCurrentTradingPeriod(v SparkResponseSparkMetaCurrentTradingPeriod)`

SetCurrentTradingPeriod sets CurrentTradingPeriod field to given value.

### HasCurrentTradingPeriod

`func (o *SparkResponseSparkMeta) HasCurrentTradingPeriod() bool`

HasCurrentTradingPeriod returns a boolean if a field has been set.

### GetTradingPeriods

`func (o *SparkResponseSparkMeta) GetTradingPeriods() [][]map[string]interface{}`

GetTradingPeriods returns the TradingPeriods field if non-nil, zero value otherwise.

### GetTradingPeriodsOk

`func (o *SparkResponseSparkMeta) GetTradingPeriodsOk() (*[][]map[string]interface{}, bool)`

GetTradingPeriodsOk returns a tuple with the TradingPeriods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradingPeriods

`func (o *SparkResponseSparkMeta) SetTradingPeriods(v [][]map[string]interface{})`

SetTradingPeriods sets TradingPeriods field to given value.

### HasTradingPeriods

`func (o *SparkResponseSparkMeta) HasTradingPeriods() bool`

HasTradingPeriods returns a boolean if a field has been set.

### GetDataGranularity

`func (o *SparkResponseSparkMeta) GetDataGranularity() string`

GetDataGranularity returns the DataGranularity field if non-nil, zero value otherwise.

### GetDataGranularityOk

`func (o *SparkResponseSparkMeta) GetDataGranularityOk() (*string, bool)`

GetDataGranularityOk returns a tuple with the DataGranularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataGranularity

`func (o *SparkResponseSparkMeta) SetDataGranularity(v string)`

SetDataGranularity sets DataGranularity field to given value.

### HasDataGranularity

`func (o *SparkResponseSparkMeta) HasDataGranularity() bool`

HasDataGranularity returns a boolean if a field has been set.

### GetRange

`func (o *SparkResponseSparkMeta) GetRange() string`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *SparkResponseSparkMeta) GetRangeOk() (*string, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *SparkResponseSparkMeta) SetRange(v string)`

SetRange sets Range field to given value.

### HasRange

`func (o *SparkResponseSparkMeta) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetValidRanges

`func (o *SparkResponseSparkMeta) GetValidRanges() []string`

GetValidRanges returns the ValidRanges field if non-nil, zero value otherwise.

### GetValidRangesOk

`func (o *SparkResponseSparkMeta) GetValidRangesOk() (*[]string, bool)`

GetValidRangesOk returns a tuple with the ValidRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidRanges

`func (o *SparkResponseSparkMeta) SetValidRanges(v []string)`

SetValidRanges sets ValidRanges field to given value.

### HasValidRanges

`func (o *SparkResponseSparkMeta) HasValidRanges() bool`

HasValidRanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


