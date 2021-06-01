# SparkResponseSpark

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to [**[]SparkResponseSparkResult**](SparkResponseSparkResult.md) |  | [optional] 
**Error** | Pointer to [**Error**](Error.md) |  | [optional] 

## Methods

### NewSparkResponseSpark

`func NewSparkResponseSpark() *SparkResponseSpark`

NewSparkResponseSpark instantiates a new SparkResponseSpark object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSparkResponseSparkWithDefaults

`func NewSparkResponseSparkWithDefaults() *SparkResponseSpark`

NewSparkResponseSparkWithDefaults instantiates a new SparkResponseSpark object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *SparkResponseSpark) GetResult() []SparkResponseSparkResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *SparkResponseSpark) GetResultOk() (*[]SparkResponseSparkResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *SparkResponseSpark) SetResult(v []SparkResponseSparkResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *SparkResponseSpark) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetError

`func (o *SparkResponseSpark) GetError() Error`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *SparkResponseSpark) GetErrorOk() (*Error, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *SparkResponseSpark) SetError(v Error)`

SetError sets Error field to given value.

### HasError

`func (o *SparkResponseSpark) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


