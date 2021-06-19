# QuoteResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FullExchangeName** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**FiftyTwoWeekLowChangePercent** | Pointer to [**QuoteResultFiftyTwoWeekLowChangePercent**](QuoteResultFiftyTwoWeekLowChangePercent.md) |  | [optional] 
**GmtOffSetMilliseconds** | Pointer to **int32** |  | [optional] 
**RegularMarketOpen** | Pointer to [**QuoteResultRegularMarketOpen**](QuoteResultRegularMarketOpen.md) |  | [optional] 
**Language** | Pointer to **string** |  | [optional] 
**RegularMarketTime** | Pointer to [**QuoteResultRegularMarketTime**](QuoteResultRegularMarketTime.md) |  | [optional] 
**RegularMarketChangePercent** | Pointer to [**QuoteResultRegularMarketChangePercent**](QuoteResultRegularMarketChangePercent.md) |  | [optional] 
**QuoteType** | Pointer to **string** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**RegularMarketDayRange** | Pointer to [**QuoteResultRegularMarketDayRange**](QuoteResultRegularMarketDayRange.md) |  | [optional] 
**FiftyTwoWeekLowChange** | Pointer to [**QuoteResultFiftyTwoWeekLowChange**](QuoteResultFiftyTwoWeekLowChange.md) |  | [optional] 
**FiftyTwoWeekHighChangePercent** | Pointer to [**QuoteResultFiftyTwoWeekHighChangePercent**](QuoteResultFiftyTwoWeekHighChangePercent.md) |  | [optional] 
**RegularMarketDayHigh** | Pointer to [**QuoteResultRegularMarketDayHigh**](QuoteResultRegularMarketDayHigh.md) |  | [optional] 
**Tradeable** | Pointer to **bool** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**FiftyTwoWeekHigh** | Pointer to [**QuoteResultRegularMarketDayHigh**](QuoteResultRegularMarketDayHigh.md) |  | [optional] 
**RegularMarketPreviousClose** | Pointer to [**QuoteResultRegularMarketPreviousClose**](QuoteResultRegularMarketPreviousClose.md) |  | [optional] 
**ExchangeTimezoneName** | Pointer to **string** |  | [optional] 
**FiftyTwoWeekHighChange** | Pointer to [**QuoteResultFiftyTwoWeekHighChange**](QuoteResultFiftyTwoWeekHighChange.md) |  | [optional] 
**RegularMarketChange** | Pointer to [**QuoteResultRegularMarketChange**](QuoteResultRegularMarketChange.md) |  | [optional] 
**FiftyTwoWeekRange** | Pointer to [**QuoteResultFiftyTwoWeekRange**](QuoteResultFiftyTwoWeekRange.md) |  | [optional] 
**ExchangeDataDelayedBy** | Pointer to **int32** |  | [optional] 
**FirstTradeDateMilliseconds** | Pointer to **int64** |  | [optional] 
**ExchangeTimezoneShortName** | Pointer to **string** |  | [optional] 
**MarketState** | Pointer to **string** |  | [optional] 
**FiftyTwoWeekLow** | Pointer to [**QuoteResultFiftyTwoWeekLow**](QuoteResultFiftyTwoWeekLow.md) |  | [optional] 
**RegularMarketPrice** | Pointer to [**QuoteResultRegularMarketPrice**](QuoteResultRegularMarketPrice.md) |  | [optional] 
**Market** | Pointer to **string** |  | [optional] 
**RegularMarketVolume** | Pointer to [**QuoteResultRegularMarketVolume**](QuoteResultRegularMarketVolume.md) |  | [optional] 
**QuoteSourceName** | Pointer to **string** |  | [optional] 
**MessageBoardId** | Pointer to **string** |  | [optional] 
**PriceHint** | Pointer to **int32** |  | [optional] 
**Exchange** | Pointer to **string** |  | [optional] 
**SourceInterval** | Pointer to **int32** |  | [optional] 
**RegularMarketDayLow** | Pointer to [**QuoteResultRegularMarketDayLow**](QuoteResultRegularMarketDayLow.md) |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**ShortName** | Pointer to **string** |  | [optional] 
**Triggerable** | Pointer to **bool** |  | [optional] 

## Methods

### NewQuoteResult

`func NewQuoteResult() *QuoteResult`

NewQuoteResult instantiates a new QuoteResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuoteResultWithDefaults

`func NewQuoteResultWithDefaults() *QuoteResult`

NewQuoteResultWithDefaults instantiates a new QuoteResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFullExchangeName

`func (o *QuoteResult) GetFullExchangeName() string`

GetFullExchangeName returns the FullExchangeName field if non-nil, zero value otherwise.

### GetFullExchangeNameOk

`func (o *QuoteResult) GetFullExchangeNameOk() (*string, bool)`

GetFullExchangeNameOk returns a tuple with the FullExchangeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullExchangeName

`func (o *QuoteResult) SetFullExchangeName(v string)`

SetFullExchangeName sets FullExchangeName field to given value.

### HasFullExchangeName

`func (o *QuoteResult) HasFullExchangeName() bool`

HasFullExchangeName returns a boolean if a field has been set.

### GetSymbol

`func (o *QuoteResult) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *QuoteResult) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *QuoteResult) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *QuoteResult) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetFiftyTwoWeekLowChangePercent

`func (o *QuoteResult) GetFiftyTwoWeekLowChangePercent() QuoteResultFiftyTwoWeekLowChangePercent`

GetFiftyTwoWeekLowChangePercent returns the FiftyTwoWeekLowChangePercent field if non-nil, zero value otherwise.

### GetFiftyTwoWeekLowChangePercentOk

`func (o *QuoteResult) GetFiftyTwoWeekLowChangePercentOk() (*QuoteResultFiftyTwoWeekLowChangePercent, bool)`

GetFiftyTwoWeekLowChangePercentOk returns a tuple with the FiftyTwoWeekLowChangePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiftyTwoWeekLowChangePercent

`func (o *QuoteResult) SetFiftyTwoWeekLowChangePercent(v QuoteResultFiftyTwoWeekLowChangePercent)`

SetFiftyTwoWeekLowChangePercent sets FiftyTwoWeekLowChangePercent field to given value.

### HasFiftyTwoWeekLowChangePercent

`func (o *QuoteResult) HasFiftyTwoWeekLowChangePercent() bool`

HasFiftyTwoWeekLowChangePercent returns a boolean if a field has been set.

### GetGmtOffSetMilliseconds

`func (o *QuoteResult) GetGmtOffSetMilliseconds() int32`

GetGmtOffSetMilliseconds returns the GmtOffSetMilliseconds field if non-nil, zero value otherwise.

### GetGmtOffSetMillisecondsOk

`func (o *QuoteResult) GetGmtOffSetMillisecondsOk() (*int32, bool)`

GetGmtOffSetMillisecondsOk returns a tuple with the GmtOffSetMilliseconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmtOffSetMilliseconds

`func (o *QuoteResult) SetGmtOffSetMilliseconds(v int32)`

SetGmtOffSetMilliseconds sets GmtOffSetMilliseconds field to given value.

### HasGmtOffSetMilliseconds

`func (o *QuoteResult) HasGmtOffSetMilliseconds() bool`

HasGmtOffSetMilliseconds returns a boolean if a field has been set.

### GetRegularMarketOpen

`func (o *QuoteResult) GetRegularMarketOpen() QuoteResultRegularMarketOpen`

GetRegularMarketOpen returns the RegularMarketOpen field if non-nil, zero value otherwise.

### GetRegularMarketOpenOk

`func (o *QuoteResult) GetRegularMarketOpenOk() (*QuoteResultRegularMarketOpen, bool)`

GetRegularMarketOpenOk returns a tuple with the RegularMarketOpen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegularMarketOpen

`func (o *QuoteResult) SetRegularMarketOpen(v QuoteResultRegularMarketOpen)`

SetRegularMarketOpen sets RegularMarketOpen field to given value.

### HasRegularMarketOpen

`func (o *QuoteResult) HasRegularMarketOpen() bool`

HasRegularMarketOpen returns a boolean if a field has been set.

### GetLanguage

`func (o *QuoteResult) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *QuoteResult) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *QuoteResult) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *QuoteResult) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetRegularMarketTime

`func (o *QuoteResult) GetRegularMarketTime() QuoteResultRegularMarketTime`

GetRegularMarketTime returns the RegularMarketTime field if non-nil, zero value otherwise.

### GetRegularMarketTimeOk

`func (o *QuoteResult) GetRegularMarketTimeOk() (*QuoteResultRegularMarketTime, bool)`

GetRegularMarketTimeOk returns a tuple with the RegularMarketTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegularMarketTime

`func (o *QuoteResult) SetRegularMarketTime(v QuoteResultRegularMarketTime)`

SetRegularMarketTime sets RegularMarketTime field to given value.

### HasRegularMarketTime

`func (o *QuoteResult) HasRegularMarketTime() bool`

HasRegularMarketTime returns a boolean if a field has been set.

### GetRegularMarketChangePercent

`func (o *QuoteResult) GetRegularMarketChangePercent() QuoteResultRegularMarketChangePercent`

GetRegularMarketChangePercent returns the RegularMarketChangePercent field if non-nil, zero value otherwise.

### GetRegularMarketChangePercentOk

`func (o *QuoteResult) GetRegularMarketChangePercentOk() (*QuoteResultRegularMarketChangePercent, bool)`

GetRegularMarketChangePercentOk returns a tuple with the RegularMarketChangePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegularMarketChangePercent

`func (o *QuoteResult) SetRegularMarketChangePercent(v QuoteResultRegularMarketChangePercent)`

SetRegularMarketChangePercent sets RegularMarketChangePercent field to given value.

### HasRegularMarketChangePercent

`func (o *QuoteResult) HasRegularMarketChangePercent() bool`

HasRegularMarketChangePercent returns a boolean if a field has been set.

### GetQuoteType

`func (o *QuoteResult) GetQuoteType() string`

GetQuoteType returns the QuoteType field if non-nil, zero value otherwise.

### GetQuoteTypeOk

`func (o *QuoteResult) GetQuoteTypeOk() (*string, bool)`

GetQuoteTypeOk returns a tuple with the QuoteType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteType

`func (o *QuoteResult) SetQuoteType(v string)`

SetQuoteType sets QuoteType field to given value.

### HasQuoteType

`func (o *QuoteResult) HasQuoteType() bool`

HasQuoteType returns a boolean if a field has been set.

### GetUuid

`func (o *QuoteResult) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *QuoteResult) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *QuoteResult) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *QuoteResult) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetRegularMarketDayRange

`func (o *QuoteResult) GetRegularMarketDayRange() QuoteResultRegularMarketDayRange`

GetRegularMarketDayRange returns the RegularMarketDayRange field if non-nil, zero value otherwise.

### GetRegularMarketDayRangeOk

`func (o *QuoteResult) GetRegularMarketDayRangeOk() (*QuoteResultRegularMarketDayRange, bool)`

GetRegularMarketDayRangeOk returns a tuple with the RegularMarketDayRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegularMarketDayRange

`func (o *QuoteResult) SetRegularMarketDayRange(v QuoteResultRegularMarketDayRange)`

SetRegularMarketDayRange sets RegularMarketDayRange field to given value.

### HasRegularMarketDayRange

`func (o *QuoteResult) HasRegularMarketDayRange() bool`

HasRegularMarketDayRange returns a boolean if a field has been set.

### GetFiftyTwoWeekLowChange

`func (o *QuoteResult) GetFiftyTwoWeekLowChange() QuoteResultFiftyTwoWeekLowChange`

GetFiftyTwoWeekLowChange returns the FiftyTwoWeekLowChange field if non-nil, zero value otherwise.

### GetFiftyTwoWeekLowChangeOk

`func (o *QuoteResult) GetFiftyTwoWeekLowChangeOk() (*QuoteResultFiftyTwoWeekLowChange, bool)`

GetFiftyTwoWeekLowChangeOk returns a tuple with the FiftyTwoWeekLowChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiftyTwoWeekLowChange

`func (o *QuoteResult) SetFiftyTwoWeekLowChange(v QuoteResultFiftyTwoWeekLowChange)`

SetFiftyTwoWeekLowChange sets FiftyTwoWeekLowChange field to given value.

### HasFiftyTwoWeekLowChange

`func (o *QuoteResult) HasFiftyTwoWeekLowChange() bool`

HasFiftyTwoWeekLowChange returns a boolean if a field has been set.

### GetFiftyTwoWeekHighChangePercent

`func (o *QuoteResult) GetFiftyTwoWeekHighChangePercent() QuoteResultFiftyTwoWeekHighChangePercent`

GetFiftyTwoWeekHighChangePercent returns the FiftyTwoWeekHighChangePercent field if non-nil, zero value otherwise.

### GetFiftyTwoWeekHighChangePercentOk

`func (o *QuoteResult) GetFiftyTwoWeekHighChangePercentOk() (*QuoteResultFiftyTwoWeekHighChangePercent, bool)`

GetFiftyTwoWeekHighChangePercentOk returns a tuple with the FiftyTwoWeekHighChangePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiftyTwoWeekHighChangePercent

`func (o *QuoteResult) SetFiftyTwoWeekHighChangePercent(v QuoteResultFiftyTwoWeekHighChangePercent)`

SetFiftyTwoWeekHighChangePercent sets FiftyTwoWeekHighChangePercent field to given value.

### HasFiftyTwoWeekHighChangePercent

`func (o *QuoteResult) HasFiftyTwoWeekHighChangePercent() bool`

HasFiftyTwoWeekHighChangePercent returns a boolean if a field has been set.

### GetRegularMarketDayHigh

`func (o *QuoteResult) GetRegularMarketDayHigh() QuoteResultRegularMarketDayHigh`

GetRegularMarketDayHigh returns the RegularMarketDayHigh field if non-nil, zero value otherwise.

### GetRegularMarketDayHighOk

`func (o *QuoteResult) GetRegularMarketDayHighOk() (*QuoteResultRegularMarketDayHigh, bool)`

GetRegularMarketDayHighOk returns a tuple with the RegularMarketDayHigh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegularMarketDayHigh

`func (o *QuoteResult) SetRegularMarketDayHigh(v QuoteResultRegularMarketDayHigh)`

SetRegularMarketDayHigh sets RegularMarketDayHigh field to given value.

### HasRegularMarketDayHigh

`func (o *QuoteResult) HasRegularMarketDayHigh() bool`

HasRegularMarketDayHigh returns a boolean if a field has been set.

### GetTradeable

`func (o *QuoteResult) GetTradeable() bool`

GetTradeable returns the Tradeable field if non-nil, zero value otherwise.

### GetTradeableOk

`func (o *QuoteResult) GetTradeableOk() (*bool, bool)`

GetTradeableOk returns a tuple with the Tradeable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeable

`func (o *QuoteResult) SetTradeable(v bool)`

SetTradeable sets Tradeable field to given value.

### HasTradeable

`func (o *QuoteResult) HasTradeable() bool`

HasTradeable returns a boolean if a field has been set.

### GetCurrency

`func (o *QuoteResult) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *QuoteResult) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *QuoteResult) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *QuoteResult) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetFiftyTwoWeekHigh

`func (o *QuoteResult) GetFiftyTwoWeekHigh() QuoteResultRegularMarketDayHigh`

GetFiftyTwoWeekHigh returns the FiftyTwoWeekHigh field if non-nil, zero value otherwise.

### GetFiftyTwoWeekHighOk

`func (o *QuoteResult) GetFiftyTwoWeekHighOk() (*QuoteResultRegularMarketDayHigh, bool)`

GetFiftyTwoWeekHighOk returns a tuple with the FiftyTwoWeekHigh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiftyTwoWeekHigh

`func (o *QuoteResult) SetFiftyTwoWeekHigh(v QuoteResultRegularMarketDayHigh)`

SetFiftyTwoWeekHigh sets FiftyTwoWeekHigh field to given value.

### HasFiftyTwoWeekHigh

`func (o *QuoteResult) HasFiftyTwoWeekHigh() bool`

HasFiftyTwoWeekHigh returns a boolean if a field has been set.

### GetRegularMarketPreviousClose

`func (o *QuoteResult) GetRegularMarketPreviousClose() QuoteResultRegularMarketPreviousClose`

GetRegularMarketPreviousClose returns the RegularMarketPreviousClose field if non-nil, zero value otherwise.

### GetRegularMarketPreviousCloseOk

`func (o *QuoteResult) GetRegularMarketPreviousCloseOk() (*QuoteResultRegularMarketPreviousClose, bool)`

GetRegularMarketPreviousCloseOk returns a tuple with the RegularMarketPreviousClose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegularMarketPreviousClose

`func (o *QuoteResult) SetRegularMarketPreviousClose(v QuoteResultRegularMarketPreviousClose)`

SetRegularMarketPreviousClose sets RegularMarketPreviousClose field to given value.

### HasRegularMarketPreviousClose

`func (o *QuoteResult) HasRegularMarketPreviousClose() bool`

HasRegularMarketPreviousClose returns a boolean if a field has been set.

### GetExchangeTimezoneName

`func (o *QuoteResult) GetExchangeTimezoneName() string`

GetExchangeTimezoneName returns the ExchangeTimezoneName field if non-nil, zero value otherwise.

### GetExchangeTimezoneNameOk

`func (o *QuoteResult) GetExchangeTimezoneNameOk() (*string, bool)`

GetExchangeTimezoneNameOk returns a tuple with the ExchangeTimezoneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeTimezoneName

`func (o *QuoteResult) SetExchangeTimezoneName(v string)`

SetExchangeTimezoneName sets ExchangeTimezoneName field to given value.

### HasExchangeTimezoneName

`func (o *QuoteResult) HasExchangeTimezoneName() bool`

HasExchangeTimezoneName returns a boolean if a field has been set.

### GetFiftyTwoWeekHighChange

`func (o *QuoteResult) GetFiftyTwoWeekHighChange() QuoteResultFiftyTwoWeekHighChange`

GetFiftyTwoWeekHighChange returns the FiftyTwoWeekHighChange field if non-nil, zero value otherwise.

### GetFiftyTwoWeekHighChangeOk

`func (o *QuoteResult) GetFiftyTwoWeekHighChangeOk() (*QuoteResultFiftyTwoWeekHighChange, bool)`

GetFiftyTwoWeekHighChangeOk returns a tuple with the FiftyTwoWeekHighChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiftyTwoWeekHighChange

`func (o *QuoteResult) SetFiftyTwoWeekHighChange(v QuoteResultFiftyTwoWeekHighChange)`

SetFiftyTwoWeekHighChange sets FiftyTwoWeekHighChange field to given value.

### HasFiftyTwoWeekHighChange

`func (o *QuoteResult) HasFiftyTwoWeekHighChange() bool`

HasFiftyTwoWeekHighChange returns a boolean if a field has been set.

### GetRegularMarketChange

`func (o *QuoteResult) GetRegularMarketChange() QuoteResultRegularMarketChange`

GetRegularMarketChange returns the RegularMarketChange field if non-nil, zero value otherwise.

### GetRegularMarketChangeOk

`func (o *QuoteResult) GetRegularMarketChangeOk() (*QuoteResultRegularMarketChange, bool)`

GetRegularMarketChangeOk returns a tuple with the RegularMarketChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegularMarketChange

`func (o *QuoteResult) SetRegularMarketChange(v QuoteResultRegularMarketChange)`

SetRegularMarketChange sets RegularMarketChange field to given value.

### HasRegularMarketChange

`func (o *QuoteResult) HasRegularMarketChange() bool`

HasRegularMarketChange returns a boolean if a field has been set.

### GetFiftyTwoWeekRange

`func (o *QuoteResult) GetFiftyTwoWeekRange() QuoteResultFiftyTwoWeekRange`

GetFiftyTwoWeekRange returns the FiftyTwoWeekRange field if non-nil, zero value otherwise.

### GetFiftyTwoWeekRangeOk

`func (o *QuoteResult) GetFiftyTwoWeekRangeOk() (*QuoteResultFiftyTwoWeekRange, bool)`

GetFiftyTwoWeekRangeOk returns a tuple with the FiftyTwoWeekRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiftyTwoWeekRange

`func (o *QuoteResult) SetFiftyTwoWeekRange(v QuoteResultFiftyTwoWeekRange)`

SetFiftyTwoWeekRange sets FiftyTwoWeekRange field to given value.

### HasFiftyTwoWeekRange

`func (o *QuoteResult) HasFiftyTwoWeekRange() bool`

HasFiftyTwoWeekRange returns a boolean if a field has been set.

### GetExchangeDataDelayedBy

`func (o *QuoteResult) GetExchangeDataDelayedBy() int32`

GetExchangeDataDelayedBy returns the ExchangeDataDelayedBy field if non-nil, zero value otherwise.

### GetExchangeDataDelayedByOk

`func (o *QuoteResult) GetExchangeDataDelayedByOk() (*int32, bool)`

GetExchangeDataDelayedByOk returns a tuple with the ExchangeDataDelayedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeDataDelayedBy

`func (o *QuoteResult) SetExchangeDataDelayedBy(v int32)`

SetExchangeDataDelayedBy sets ExchangeDataDelayedBy field to given value.

### HasExchangeDataDelayedBy

`func (o *QuoteResult) HasExchangeDataDelayedBy() bool`

HasExchangeDataDelayedBy returns a boolean if a field has been set.

### GetFirstTradeDateMilliseconds

`func (o *QuoteResult) GetFirstTradeDateMilliseconds() int64`

GetFirstTradeDateMilliseconds returns the FirstTradeDateMilliseconds field if non-nil, zero value otherwise.

### GetFirstTradeDateMillisecondsOk

`func (o *QuoteResult) GetFirstTradeDateMillisecondsOk() (*int64, bool)`

GetFirstTradeDateMillisecondsOk returns a tuple with the FirstTradeDateMilliseconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstTradeDateMilliseconds

`func (o *QuoteResult) SetFirstTradeDateMilliseconds(v int64)`

SetFirstTradeDateMilliseconds sets FirstTradeDateMilliseconds field to given value.

### HasFirstTradeDateMilliseconds

`func (o *QuoteResult) HasFirstTradeDateMilliseconds() bool`

HasFirstTradeDateMilliseconds returns a boolean if a field has been set.

### GetExchangeTimezoneShortName

`func (o *QuoteResult) GetExchangeTimezoneShortName() string`

GetExchangeTimezoneShortName returns the ExchangeTimezoneShortName field if non-nil, zero value otherwise.

### GetExchangeTimezoneShortNameOk

`func (o *QuoteResult) GetExchangeTimezoneShortNameOk() (*string, bool)`

GetExchangeTimezoneShortNameOk returns a tuple with the ExchangeTimezoneShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeTimezoneShortName

`func (o *QuoteResult) SetExchangeTimezoneShortName(v string)`

SetExchangeTimezoneShortName sets ExchangeTimezoneShortName field to given value.

### HasExchangeTimezoneShortName

`func (o *QuoteResult) HasExchangeTimezoneShortName() bool`

HasExchangeTimezoneShortName returns a boolean if a field has been set.

### GetMarketState

`func (o *QuoteResult) GetMarketState() string`

GetMarketState returns the MarketState field if non-nil, zero value otherwise.

### GetMarketStateOk

`func (o *QuoteResult) GetMarketStateOk() (*string, bool)`

GetMarketStateOk returns a tuple with the MarketState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketState

`func (o *QuoteResult) SetMarketState(v string)`

SetMarketState sets MarketState field to given value.

### HasMarketState

`func (o *QuoteResult) HasMarketState() bool`

HasMarketState returns a boolean if a field has been set.

### GetFiftyTwoWeekLow

`func (o *QuoteResult) GetFiftyTwoWeekLow() QuoteResultFiftyTwoWeekLow`

GetFiftyTwoWeekLow returns the FiftyTwoWeekLow field if non-nil, zero value otherwise.

### GetFiftyTwoWeekLowOk

`func (o *QuoteResult) GetFiftyTwoWeekLowOk() (*QuoteResultFiftyTwoWeekLow, bool)`

GetFiftyTwoWeekLowOk returns a tuple with the FiftyTwoWeekLow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiftyTwoWeekLow

`func (o *QuoteResult) SetFiftyTwoWeekLow(v QuoteResultFiftyTwoWeekLow)`

SetFiftyTwoWeekLow sets FiftyTwoWeekLow field to given value.

### HasFiftyTwoWeekLow

`func (o *QuoteResult) HasFiftyTwoWeekLow() bool`

HasFiftyTwoWeekLow returns a boolean if a field has been set.

### GetRegularMarketPrice

`func (o *QuoteResult) GetRegularMarketPrice() QuoteResultRegularMarketPrice`

GetRegularMarketPrice returns the RegularMarketPrice field if non-nil, zero value otherwise.

### GetRegularMarketPriceOk

`func (o *QuoteResult) GetRegularMarketPriceOk() (*QuoteResultRegularMarketPrice, bool)`

GetRegularMarketPriceOk returns a tuple with the RegularMarketPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegularMarketPrice

`func (o *QuoteResult) SetRegularMarketPrice(v QuoteResultRegularMarketPrice)`

SetRegularMarketPrice sets RegularMarketPrice field to given value.

### HasRegularMarketPrice

`func (o *QuoteResult) HasRegularMarketPrice() bool`

HasRegularMarketPrice returns a boolean if a field has been set.

### GetMarket

`func (o *QuoteResult) GetMarket() string`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *QuoteResult) GetMarketOk() (*string, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *QuoteResult) SetMarket(v string)`

SetMarket sets Market field to given value.

### HasMarket

`func (o *QuoteResult) HasMarket() bool`

HasMarket returns a boolean if a field has been set.

### GetRegularMarketVolume

`func (o *QuoteResult) GetRegularMarketVolume() QuoteResultRegularMarketVolume`

GetRegularMarketVolume returns the RegularMarketVolume field if non-nil, zero value otherwise.

### GetRegularMarketVolumeOk

`func (o *QuoteResult) GetRegularMarketVolumeOk() (*QuoteResultRegularMarketVolume, bool)`

GetRegularMarketVolumeOk returns a tuple with the RegularMarketVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegularMarketVolume

`func (o *QuoteResult) SetRegularMarketVolume(v QuoteResultRegularMarketVolume)`

SetRegularMarketVolume sets RegularMarketVolume field to given value.

### HasRegularMarketVolume

`func (o *QuoteResult) HasRegularMarketVolume() bool`

HasRegularMarketVolume returns a boolean if a field has been set.

### GetQuoteSourceName

`func (o *QuoteResult) GetQuoteSourceName() string`

GetQuoteSourceName returns the QuoteSourceName field if non-nil, zero value otherwise.

### GetQuoteSourceNameOk

`func (o *QuoteResult) GetQuoteSourceNameOk() (*string, bool)`

GetQuoteSourceNameOk returns a tuple with the QuoteSourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteSourceName

`func (o *QuoteResult) SetQuoteSourceName(v string)`

SetQuoteSourceName sets QuoteSourceName field to given value.

### HasQuoteSourceName

`func (o *QuoteResult) HasQuoteSourceName() bool`

HasQuoteSourceName returns a boolean if a field has been set.

### GetMessageBoardId

`func (o *QuoteResult) GetMessageBoardId() string`

GetMessageBoardId returns the MessageBoardId field if non-nil, zero value otherwise.

### GetMessageBoardIdOk

`func (o *QuoteResult) GetMessageBoardIdOk() (*string, bool)`

GetMessageBoardIdOk returns a tuple with the MessageBoardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageBoardId

`func (o *QuoteResult) SetMessageBoardId(v string)`

SetMessageBoardId sets MessageBoardId field to given value.

### HasMessageBoardId

`func (o *QuoteResult) HasMessageBoardId() bool`

HasMessageBoardId returns a boolean if a field has been set.

### GetPriceHint

`func (o *QuoteResult) GetPriceHint() int32`

GetPriceHint returns the PriceHint field if non-nil, zero value otherwise.

### GetPriceHintOk

`func (o *QuoteResult) GetPriceHintOk() (*int32, bool)`

GetPriceHintOk returns a tuple with the PriceHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceHint

`func (o *QuoteResult) SetPriceHint(v int32)`

SetPriceHint sets PriceHint field to given value.

### HasPriceHint

`func (o *QuoteResult) HasPriceHint() bool`

HasPriceHint returns a boolean if a field has been set.

### GetExchange

`func (o *QuoteResult) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *QuoteResult) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *QuoteResult) SetExchange(v string)`

SetExchange sets Exchange field to given value.

### HasExchange

`func (o *QuoteResult) HasExchange() bool`

HasExchange returns a boolean if a field has been set.

### GetSourceInterval

`func (o *QuoteResult) GetSourceInterval() int32`

GetSourceInterval returns the SourceInterval field if non-nil, zero value otherwise.

### GetSourceIntervalOk

`func (o *QuoteResult) GetSourceIntervalOk() (*int32, bool)`

GetSourceIntervalOk returns a tuple with the SourceInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceInterval

`func (o *QuoteResult) SetSourceInterval(v int32)`

SetSourceInterval sets SourceInterval field to given value.

### HasSourceInterval

`func (o *QuoteResult) HasSourceInterval() bool`

HasSourceInterval returns a boolean if a field has been set.

### GetRegularMarketDayLow

`func (o *QuoteResult) GetRegularMarketDayLow() QuoteResultRegularMarketDayLow`

GetRegularMarketDayLow returns the RegularMarketDayLow field if non-nil, zero value otherwise.

### GetRegularMarketDayLowOk

`func (o *QuoteResult) GetRegularMarketDayLowOk() (*QuoteResultRegularMarketDayLow, bool)`

GetRegularMarketDayLowOk returns a tuple with the RegularMarketDayLow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegularMarketDayLow

`func (o *QuoteResult) SetRegularMarketDayLow(v QuoteResultRegularMarketDayLow)`

SetRegularMarketDayLow sets RegularMarketDayLow field to given value.

### HasRegularMarketDayLow

`func (o *QuoteResult) HasRegularMarketDayLow() bool`

HasRegularMarketDayLow returns a boolean if a field has been set.

### GetRegion

`func (o *QuoteResult) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *QuoteResult) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *QuoteResult) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *QuoteResult) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetShortName

`func (o *QuoteResult) GetShortName() string`

GetShortName returns the ShortName field if non-nil, zero value otherwise.

### GetShortNameOk

`func (o *QuoteResult) GetShortNameOk() (*string, bool)`

GetShortNameOk returns a tuple with the ShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortName

`func (o *QuoteResult) SetShortName(v string)`

SetShortName sets ShortName field to given value.

### HasShortName

`func (o *QuoteResult) HasShortName() bool`

HasShortName returns a boolean if a field has been set.

### GetTriggerable

`func (o *QuoteResult) GetTriggerable() bool`

GetTriggerable returns the Triggerable field if non-nil, zero value otherwise.

### GetTriggerableOk

`func (o *QuoteResult) GetTriggerableOk() (*bool, bool)`

GetTriggerableOk returns a tuple with the Triggerable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerable

`func (o *QuoteResult) SetTriggerable(v bool)`

SetTriggerable sets Triggerable field to given value.

### HasTriggerable

`func (o *QuoteResult) HasTriggerable() bool`

HasTriggerable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


