# ResponsesTextDoneEventLogprob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | **string** |  | 
**Logprob** | **float32** |  | 
**TopLogprobs** | Pointer to [**[]ResponsesTextDoneEventLogprobTopLogprob**](ResponsesTextDoneEventLogprobTopLogprob.md) |  | [optional] 

## Methods

### NewResponsesTextDoneEventLogprob

`func NewResponsesTextDoneEventLogprob(token string, logprob float32, ) *ResponsesTextDoneEventLogprob`

NewResponsesTextDoneEventLogprob instantiates a new ResponsesTextDoneEventLogprob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsesTextDoneEventLogprobWithDefaults

`func NewResponsesTextDoneEventLogprobWithDefaults() *ResponsesTextDoneEventLogprob`

NewResponsesTextDoneEventLogprobWithDefaults instantiates a new ResponsesTextDoneEventLogprob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *ResponsesTextDoneEventLogprob) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ResponsesTextDoneEventLogprob) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ResponsesTextDoneEventLogprob) SetToken(v string)`

SetToken sets Token field to given value.


### GetLogprob

`func (o *ResponsesTextDoneEventLogprob) GetLogprob() float32`

GetLogprob returns the Logprob field if non-nil, zero value otherwise.

### GetLogprobOk

`func (o *ResponsesTextDoneEventLogprob) GetLogprobOk() (*float32, bool)`

GetLogprobOk returns a tuple with the Logprob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogprob

`func (o *ResponsesTextDoneEventLogprob) SetLogprob(v float32)`

SetLogprob sets Logprob field to given value.


### GetTopLogprobs

`func (o *ResponsesTextDoneEventLogprob) GetTopLogprobs() []ResponsesTextDoneEventLogprobTopLogprob`

GetTopLogprobs returns the TopLogprobs field if non-nil, zero value otherwise.

### GetTopLogprobsOk

`func (o *ResponsesTextDoneEventLogprob) GetTopLogprobsOk() (*[]ResponsesTextDoneEventLogprobTopLogprob, bool)`

GetTopLogprobsOk returns a tuple with the TopLogprobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopLogprobs

`func (o *ResponsesTextDoneEventLogprob) SetTopLogprobs(v []ResponsesTextDoneEventLogprobTopLogprob)`

SetTopLogprobs sets TopLogprobs field to given value.

### HasTopLogprobs

`func (o *ResponsesTextDoneEventLogprob) HasTopLogprobs() bool`

HasTopLogprobs returns a boolean if a field has been set.

### SetTopLogprobsNil

`func (o *ResponsesTextDoneEventLogprob) SetTopLogprobsNil(b bool)`

 SetTopLogprobsNil sets the value for TopLogprobs to be an explicit nil

### UnsetTopLogprobs
`func (o *ResponsesTextDoneEventLogprob) UnsetTopLogprobs()`

UnsetTopLogprobs ensures that no value is present for TopLogprobs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


