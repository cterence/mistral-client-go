# ResponsesOutputTextLogprob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | **string** |  | 
**Bytes** | **[]int32** |  | 
**Logprob** | **float32** |  | 
**TopLogprobs** | [**[]ResponsesOutputTextLogprobTopLogprob**](ResponsesOutputTextLogprobTopLogprob.md) |  | 

## Methods

### NewResponsesOutputTextLogprob

`func NewResponsesOutputTextLogprob(token string, bytes []int32, logprob float32, topLogprobs []ResponsesOutputTextLogprobTopLogprob, ) *ResponsesOutputTextLogprob`

NewResponsesOutputTextLogprob instantiates a new ResponsesOutputTextLogprob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsesOutputTextLogprobWithDefaults

`func NewResponsesOutputTextLogprobWithDefaults() *ResponsesOutputTextLogprob`

NewResponsesOutputTextLogprobWithDefaults instantiates a new ResponsesOutputTextLogprob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *ResponsesOutputTextLogprob) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ResponsesOutputTextLogprob) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ResponsesOutputTextLogprob) SetToken(v string)`

SetToken sets Token field to given value.


### GetBytes

`func (o *ResponsesOutputTextLogprob) GetBytes() []int32`

GetBytes returns the Bytes field if non-nil, zero value otherwise.

### GetBytesOk

`func (o *ResponsesOutputTextLogprob) GetBytesOk() (*[]int32, bool)`

GetBytesOk returns a tuple with the Bytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytes

`func (o *ResponsesOutputTextLogprob) SetBytes(v []int32)`

SetBytes sets Bytes field to given value.


### GetLogprob

`func (o *ResponsesOutputTextLogprob) GetLogprob() float32`

GetLogprob returns the Logprob field if non-nil, zero value otherwise.

### GetLogprobOk

`func (o *ResponsesOutputTextLogprob) GetLogprobOk() (*float32, bool)`

GetLogprobOk returns a tuple with the Logprob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogprob

`func (o *ResponsesOutputTextLogprob) SetLogprob(v float32)`

SetLogprob sets Logprob field to given value.


### GetTopLogprobs

`func (o *ResponsesOutputTextLogprob) GetTopLogprobs() []ResponsesOutputTextLogprobTopLogprob`

GetTopLogprobs returns the TopLogprobs field if non-nil, zero value otherwise.

### GetTopLogprobsOk

`func (o *ResponsesOutputTextLogprob) GetTopLogprobsOk() (*[]ResponsesOutputTextLogprobTopLogprob, bool)`

GetTopLogprobsOk returns a tuple with the TopLogprobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopLogprobs

`func (o *ResponsesOutputTextLogprob) SetTopLogprobs(v []ResponsesOutputTextLogprobTopLogprob)`

SetTopLogprobs sets TopLogprobs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


