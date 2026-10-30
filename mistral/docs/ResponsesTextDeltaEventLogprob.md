# ResponsesTextDeltaEventLogprob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | **string** |  | 
**Logprob** | **float32** |  | 
**TopLogprobs** | Pointer to [**[]ResponsesTextDeltaEventLogprobTopLogprob**](ResponsesTextDeltaEventLogprobTopLogprob.md) |  | [optional] 

## Methods

### NewResponsesTextDeltaEventLogprob

`func NewResponsesTextDeltaEventLogprob(token string, logprob float32, ) *ResponsesTextDeltaEventLogprob`

NewResponsesTextDeltaEventLogprob instantiates a new ResponsesTextDeltaEventLogprob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsesTextDeltaEventLogprobWithDefaults

`func NewResponsesTextDeltaEventLogprobWithDefaults() *ResponsesTextDeltaEventLogprob`

NewResponsesTextDeltaEventLogprobWithDefaults instantiates a new ResponsesTextDeltaEventLogprob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *ResponsesTextDeltaEventLogprob) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ResponsesTextDeltaEventLogprob) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ResponsesTextDeltaEventLogprob) SetToken(v string)`

SetToken sets Token field to given value.


### GetLogprob

`func (o *ResponsesTextDeltaEventLogprob) GetLogprob() float32`

GetLogprob returns the Logprob field if non-nil, zero value otherwise.

### GetLogprobOk

`func (o *ResponsesTextDeltaEventLogprob) GetLogprobOk() (*float32, bool)`

GetLogprobOk returns a tuple with the Logprob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogprob

`func (o *ResponsesTextDeltaEventLogprob) SetLogprob(v float32)`

SetLogprob sets Logprob field to given value.


### GetTopLogprobs

`func (o *ResponsesTextDeltaEventLogprob) GetTopLogprobs() []ResponsesTextDeltaEventLogprobTopLogprob`

GetTopLogprobs returns the TopLogprobs field if non-nil, zero value otherwise.

### GetTopLogprobsOk

`func (o *ResponsesTextDeltaEventLogprob) GetTopLogprobsOk() (*[]ResponsesTextDeltaEventLogprobTopLogprob, bool)`

GetTopLogprobsOk returns a tuple with the TopLogprobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopLogprobs

`func (o *ResponsesTextDeltaEventLogprob) SetTopLogprobs(v []ResponsesTextDeltaEventLogprobTopLogprob)`

SetTopLogprobs sets TopLogprobs field to given value.

### HasTopLogprobs

`func (o *ResponsesTextDeltaEventLogprob) HasTopLogprobs() bool`

HasTopLogprobs returns a boolean if a field has been set.

### SetTopLogprobsNil

`func (o *ResponsesTextDeltaEventLogprob) SetTopLogprobsNil(b bool)`

 SetTopLogprobsNil sets the value for TopLogprobs to be an explicit nil

### UnsetTopLogprobs
`func (o *ResponsesTextDeltaEventLogprob) UnsetTopLogprobs()`

UnsetTopLogprobs ensures that no value is present for TopLogprobs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


