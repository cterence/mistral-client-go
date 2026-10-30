# ChatCompletionResponseTokenLogprobsTopK

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | **string** |  | 
**Logprob** | **float32** |  | 
**Bytes** | **[]int32** |  | 
**TokenId** | **int32** |  | 
**TopLogprobs** | Pointer to [**[]ChatCompletionResponseTokenLogprobs**](ChatCompletionResponseTokenLogprobs.md) |  | [optional] 

## Methods

### NewChatCompletionResponseTokenLogprobsTopK

`func NewChatCompletionResponseTokenLogprobsTopK(token string, logprob float32, bytes []int32, tokenId int32, ) *ChatCompletionResponseTokenLogprobsTopK`

NewChatCompletionResponseTokenLogprobsTopK instantiates a new ChatCompletionResponseTokenLogprobsTopK object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatCompletionResponseTokenLogprobsTopKWithDefaults

`func NewChatCompletionResponseTokenLogprobsTopKWithDefaults() *ChatCompletionResponseTokenLogprobsTopK`

NewChatCompletionResponseTokenLogprobsTopKWithDefaults instantiates a new ChatCompletionResponseTokenLogprobsTopK object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *ChatCompletionResponseTokenLogprobsTopK) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ChatCompletionResponseTokenLogprobsTopK) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ChatCompletionResponseTokenLogprobsTopK) SetToken(v string)`

SetToken sets Token field to given value.


### GetLogprob

`func (o *ChatCompletionResponseTokenLogprobsTopK) GetLogprob() float32`

GetLogprob returns the Logprob field if non-nil, zero value otherwise.

### GetLogprobOk

`func (o *ChatCompletionResponseTokenLogprobsTopK) GetLogprobOk() (*float32, bool)`

GetLogprobOk returns a tuple with the Logprob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogprob

`func (o *ChatCompletionResponseTokenLogprobsTopK) SetLogprob(v float32)`

SetLogprob sets Logprob field to given value.


### GetBytes

`func (o *ChatCompletionResponseTokenLogprobsTopK) GetBytes() []int32`

GetBytes returns the Bytes field if non-nil, zero value otherwise.

### GetBytesOk

`func (o *ChatCompletionResponseTokenLogprobsTopK) GetBytesOk() (*[]int32, bool)`

GetBytesOk returns a tuple with the Bytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytes

`func (o *ChatCompletionResponseTokenLogprobsTopK) SetBytes(v []int32)`

SetBytes sets Bytes field to given value.


### GetTokenId

`func (o *ChatCompletionResponseTokenLogprobsTopK) GetTokenId() int32`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *ChatCompletionResponseTokenLogprobsTopK) GetTokenIdOk() (*int32, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *ChatCompletionResponseTokenLogprobsTopK) SetTokenId(v int32)`

SetTokenId sets TokenId field to given value.


### GetTopLogprobs

`func (o *ChatCompletionResponseTokenLogprobsTopK) GetTopLogprobs() []ChatCompletionResponseTokenLogprobs`

GetTopLogprobs returns the TopLogprobs field if non-nil, zero value otherwise.

### GetTopLogprobsOk

`func (o *ChatCompletionResponseTokenLogprobsTopK) GetTopLogprobsOk() (*[]ChatCompletionResponseTokenLogprobs, bool)`

GetTopLogprobsOk returns a tuple with the TopLogprobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopLogprobs

`func (o *ChatCompletionResponseTokenLogprobsTopK) SetTopLogprobs(v []ChatCompletionResponseTokenLogprobs)`

SetTopLogprobs sets TopLogprobs field to given value.

### HasTopLogprobs

`func (o *ChatCompletionResponseTokenLogprobsTopK) HasTopLogprobs() bool`

HasTopLogprobs returns a boolean if a field has been set.

### SetTopLogprobsNil

`func (o *ChatCompletionResponseTokenLogprobsTopK) SetTopLogprobsNil(b bool)`

 SetTopLogprobsNil sets the value for TopLogprobs to be an explicit nil

### UnsetTopLogprobs
`func (o *ChatCompletionResponseTokenLogprobsTopK) UnsetTopLogprobs()`

UnsetTopLogprobs ensures that no value is present for TopLogprobs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


