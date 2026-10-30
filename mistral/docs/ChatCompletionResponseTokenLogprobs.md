# ChatCompletionResponseTokenLogprobs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | **string** |  | 
**Logprob** | **float32** |  | 
**Bytes** | **[]int32** |  | 
**TokenId** | **int32** |  | 

## Methods

### NewChatCompletionResponseTokenLogprobs

`func NewChatCompletionResponseTokenLogprobs(token string, logprob float32, bytes []int32, tokenId int32, ) *ChatCompletionResponseTokenLogprobs`

NewChatCompletionResponseTokenLogprobs instantiates a new ChatCompletionResponseTokenLogprobs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatCompletionResponseTokenLogprobsWithDefaults

`func NewChatCompletionResponseTokenLogprobsWithDefaults() *ChatCompletionResponseTokenLogprobs`

NewChatCompletionResponseTokenLogprobsWithDefaults instantiates a new ChatCompletionResponseTokenLogprobs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *ChatCompletionResponseTokenLogprobs) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ChatCompletionResponseTokenLogprobs) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ChatCompletionResponseTokenLogprobs) SetToken(v string)`

SetToken sets Token field to given value.


### GetLogprob

`func (o *ChatCompletionResponseTokenLogprobs) GetLogprob() float32`

GetLogprob returns the Logprob field if non-nil, zero value otherwise.

### GetLogprobOk

`func (o *ChatCompletionResponseTokenLogprobs) GetLogprobOk() (*float32, bool)`

GetLogprobOk returns a tuple with the Logprob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogprob

`func (o *ChatCompletionResponseTokenLogprobs) SetLogprob(v float32)`

SetLogprob sets Logprob field to given value.


### GetBytes

`func (o *ChatCompletionResponseTokenLogprobs) GetBytes() []int32`

GetBytes returns the Bytes field if non-nil, zero value otherwise.

### GetBytesOk

`func (o *ChatCompletionResponseTokenLogprobs) GetBytesOk() (*[]int32, bool)`

GetBytesOk returns a tuple with the Bytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytes

`func (o *ChatCompletionResponseTokenLogprobs) SetBytes(v []int32)`

SetBytes sets Bytes field to given value.


### GetTokenId

`func (o *ChatCompletionResponseTokenLogprobs) GetTokenId() int32`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *ChatCompletionResponseTokenLogprobs) GetTokenIdOk() (*int32, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *ChatCompletionResponseTokenLogprobs) SetTokenId(v int32)`

SetTokenId sets TokenId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


