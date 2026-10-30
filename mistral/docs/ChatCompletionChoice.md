# ChatCompletionChoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | **int32** |  | 
**Message** | Pointer to [**AssistantMessage**](AssistantMessage.md) |  | [optional] 
**Messages** | Pointer to [**[]DeltaMessage**](DeltaMessage.md) |  | [optional] 
**FinishReason** | **string** |  | 
**Logprobs** | Pointer to [**NullableChatCompletionResponseChoiceLogprobs**](ChatCompletionResponseChoiceLogprobs.md) |  | [optional] 
**PromptLogprobs** | Pointer to [**NullableChatCompletionResponseChoiceLogprobs**](ChatCompletionResponseChoiceLogprobs.md) |  | [optional] 

## Methods

### NewChatCompletionChoice

`func NewChatCompletionChoice(index int32, finishReason string, ) *ChatCompletionChoice`

NewChatCompletionChoice instantiates a new ChatCompletionChoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatCompletionChoiceWithDefaults

`func NewChatCompletionChoiceWithDefaults() *ChatCompletionChoice`

NewChatCompletionChoiceWithDefaults instantiates a new ChatCompletionChoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *ChatCompletionChoice) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *ChatCompletionChoice) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *ChatCompletionChoice) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetMessage

`func (o *ChatCompletionChoice) GetMessage() AssistantMessage`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ChatCompletionChoice) GetMessageOk() (*AssistantMessage, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ChatCompletionChoice) SetMessage(v AssistantMessage)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ChatCompletionChoice) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetMessages

`func (o *ChatCompletionChoice) GetMessages() []DeltaMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *ChatCompletionChoice) GetMessagesOk() (*[]DeltaMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *ChatCompletionChoice) SetMessages(v []DeltaMessage)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *ChatCompletionChoice) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetFinishReason

`func (o *ChatCompletionChoice) GetFinishReason() string`

GetFinishReason returns the FinishReason field if non-nil, zero value otherwise.

### GetFinishReasonOk

`func (o *ChatCompletionChoice) GetFinishReasonOk() (*string, bool)`

GetFinishReasonOk returns a tuple with the FinishReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishReason

`func (o *ChatCompletionChoice) SetFinishReason(v string)`

SetFinishReason sets FinishReason field to given value.


### GetLogprobs

`func (o *ChatCompletionChoice) GetLogprobs() ChatCompletionResponseChoiceLogprobs`

GetLogprobs returns the Logprobs field if non-nil, zero value otherwise.

### GetLogprobsOk

`func (o *ChatCompletionChoice) GetLogprobsOk() (*ChatCompletionResponseChoiceLogprobs, bool)`

GetLogprobsOk returns a tuple with the Logprobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogprobs

`func (o *ChatCompletionChoice) SetLogprobs(v ChatCompletionResponseChoiceLogprobs)`

SetLogprobs sets Logprobs field to given value.

### HasLogprobs

`func (o *ChatCompletionChoice) HasLogprobs() bool`

HasLogprobs returns a boolean if a field has been set.

### SetLogprobsNil

`func (o *ChatCompletionChoice) SetLogprobsNil(b bool)`

 SetLogprobsNil sets the value for Logprobs to be an explicit nil

### UnsetLogprobs
`func (o *ChatCompletionChoice) UnsetLogprobs()`

UnsetLogprobs ensures that no value is present for Logprobs, not even an explicit nil
### GetPromptLogprobs

`func (o *ChatCompletionChoice) GetPromptLogprobs() ChatCompletionResponseChoiceLogprobs`

GetPromptLogprobs returns the PromptLogprobs field if non-nil, zero value otherwise.

### GetPromptLogprobsOk

`func (o *ChatCompletionChoice) GetPromptLogprobsOk() (*ChatCompletionResponseChoiceLogprobs, bool)`

GetPromptLogprobsOk returns a tuple with the PromptLogprobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptLogprobs

`func (o *ChatCompletionChoice) SetPromptLogprobs(v ChatCompletionResponseChoiceLogprobs)`

SetPromptLogprobs sets PromptLogprobs field to given value.

### HasPromptLogprobs

`func (o *ChatCompletionChoice) HasPromptLogprobs() bool`

HasPromptLogprobs returns a boolean if a field has been set.

### SetPromptLogprobsNil

`func (o *ChatCompletionChoice) SetPromptLogprobsNil(b bool)`

 SetPromptLogprobsNil sets the value for PromptLogprobs to be an explicit nil

### UnsetPromptLogprobs
`func (o *ChatCompletionChoice) UnsetPromptLogprobs()`

UnsetPromptLogprobs ensures that no value is present for PromptLogprobs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


