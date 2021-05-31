# ChartResponseChart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to [**[]ChartResponseChartResult**](ChartResponseChartResult.md) |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 

## Methods

### NewChartResponseChart

`func NewChartResponseChart() *ChartResponseChart`

NewChartResponseChart instantiates a new ChartResponseChart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChartResponseChartWithDefaults

`func NewChartResponseChartWithDefaults() *ChartResponseChart`

NewChartResponseChartWithDefaults instantiates a new ChartResponseChart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *ChartResponseChart) GetResult() []ChartResponseChartResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ChartResponseChart) GetResultOk() (*[]ChartResponseChartResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ChartResponseChart) SetResult(v []ChartResponseChartResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *ChartResponseChart) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetError

`func (o *ChartResponseChart) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ChartResponseChart) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ChartResponseChart) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ChartResponseChart) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


