# QuoteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to [**[]QuoteResult**](QuoteResult.md) |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 

## Methods

### NewQuoteResponse

`func NewQuoteResponse() *QuoteResponse`

NewQuoteResponse instantiates a new QuoteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuoteResponseWithDefaults

`func NewQuoteResponseWithDefaults() *QuoteResponse`

NewQuoteResponseWithDefaults instantiates a new QuoteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *QuoteResponse) GetResult() []QuoteResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *QuoteResponse) GetResultOk() (*[]QuoteResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *QuoteResponse) SetResult(v []QuoteResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *QuoteResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetError

`func (o *QuoteResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *QuoteResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *QuoteResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *QuoteResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


