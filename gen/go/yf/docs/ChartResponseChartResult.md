# ChartResponseChartResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**ChartResponseChartMeta**](ChartResponseChartMeta.md) |  | [optional] 
**Timestamp** | Pointer to **[]int32** |  | [optional] 
**Indicators** | Pointer to [**ChartResponseChartIndicators**](ChartResponseChartIndicators.md) |  | [optional] 

## Methods

### NewChartResponseChartResult

`func NewChartResponseChartResult() *ChartResponseChartResult`

NewChartResponseChartResult instantiates a new ChartResponseChartResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChartResponseChartResultWithDefaults

`func NewChartResponseChartResultWithDefaults() *ChartResponseChartResult`

NewChartResponseChartResultWithDefaults instantiates a new ChartResponseChartResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *ChartResponseChartResult) GetMeta() ChartResponseChartMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ChartResponseChartResult) GetMetaOk() (*ChartResponseChartMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ChartResponseChartResult) SetMeta(v ChartResponseChartMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ChartResponseChartResult) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetTimestamp

`func (o *ChartResponseChartResult) GetTimestamp() []int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ChartResponseChartResult) GetTimestampOk() (*[]int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ChartResponseChartResult) SetTimestamp(v []int32)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ChartResponseChartResult) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetIndicators

`func (o *ChartResponseChartResult) GetIndicators() ChartResponseChartIndicators`

GetIndicators returns the Indicators field if non-nil, zero value otherwise.

### GetIndicatorsOk

`func (o *ChartResponseChartResult) GetIndicatorsOk() (*ChartResponseChartIndicators, bool)`

GetIndicatorsOk returns a tuple with the Indicators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndicators

`func (o *ChartResponseChartResult) SetIndicators(v ChartResponseChartIndicators)`

SetIndicators sets Indicators field to given value.

### HasIndicators

`func (o *ChartResponseChartResult) HasIndicators() bool`

HasIndicators returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


