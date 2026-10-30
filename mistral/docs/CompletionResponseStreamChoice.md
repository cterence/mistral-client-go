# CompletionResponseStreamChoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | **int32** |  | 
**Delta** | [**DeltaMessage**](DeltaMessage.md) |  | 
**FinishReason** | **NullableString** |  | 
**Logprobs** | Pointer to [**NullableChatCompletionResponseChoiceLogprobs**](ChatCompletionResponseChoiceLogprobs.md) |  | [optional] 
**PromptLogprobs** | Pointer to [**NullableChatCompletionResponseChoiceLogprobs**](ChatCompletionResponseChoiceLogprobs.md) |  | [optional] 

## Methods

### NewCompletionResponseStreamChoice

`func NewCompletionResponseStreamChoice(index int32, delta DeltaMessage, finishReason NullableString, ) *CompletionResponseStreamChoice`

NewCompletionResponseStreamChoice instantiates a new CompletionResponseStreamChoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompletionResponseStreamChoiceWithDefaults

`func NewCompletionResponseStreamChoiceWithDefaults() *CompletionResponseStreamChoice`

NewCompletionResponseStreamChoiceWithDefaults instantiates a new CompletionResponseStreamChoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *CompletionResponseStreamChoice) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *CompletionResponseStreamChoice) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *CompletionResponseStreamChoice) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetDelta

`func (o *CompletionResponseStreamChoice) GetDelta() DeltaMessage`

GetDelta returns the Delta field if non-nil, zero value otherwise.

### GetDeltaOk

`func (o *CompletionResponseStreamChoice) GetDeltaOk() (*DeltaMessage, bool)`

GetDeltaOk returns a tuple with the Delta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelta

`func (o *CompletionResponseStreamChoice) SetDelta(v DeltaMessage)`

SetDelta sets Delta field to given value.


### GetFinishReason

`func (o *CompletionResponseStreamChoice) GetFinishReason() string`

GetFinishReason returns the FinishReason field if non-nil, zero value otherwise.

### GetFinishReasonOk

`func (o *CompletionResponseStreamChoice) GetFinishReasonOk() (*string, bool)`

GetFinishReasonOk returns a tuple with the FinishReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishReason

`func (o *CompletionResponseStreamChoice) SetFinishReason(v string)`

SetFinishReason sets FinishReason field to given value.


### SetFinishReasonNil

`func (o *CompletionResponseStreamChoice) SetFinishReasonNil(b bool)`

 SetFinishReasonNil sets the value for FinishReason to be an explicit nil

### UnsetFinishReason
`func (o *CompletionResponseStreamChoice) UnsetFinishReason()`

UnsetFinishReason ensures that no value is present for FinishReason, not even an explicit nil
### GetLogprobs

`func (o *CompletionResponseStreamChoice) GetLogprobs() ChatCompletionResponseChoiceLogprobs`

GetLogprobs returns the Logprobs field if non-nil, zero value otherwise.

### GetLogprobsOk

`func (o *CompletionResponseStreamChoice) GetLogprobsOk() (*ChatCompletionResponseChoiceLogprobs, bool)`

GetLogprobsOk returns a tuple with the Logprobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogprobs

`func (o *CompletionResponseStreamChoice) SetLogprobs(v ChatCompletionResponseChoiceLogprobs)`

SetLogprobs sets Logprobs field to given value.

### HasLogprobs

`func (o *CompletionResponseStreamChoice) HasLogprobs() bool`

HasLogprobs returns a boolean if a field has been set.

### SetLogprobsNil

`func (o *CompletionResponseStreamChoice) SetLogprobsNil(b bool)`

 SetLogprobsNil sets the value for Logprobs to be an explicit nil

### UnsetLogprobs
`func (o *CompletionResponseStreamChoice) UnsetLogprobs()`

UnsetLogprobs ensures that no value is present for Logprobs, not even an explicit nil
### GetPromptLogprobs

`func (o *CompletionResponseStreamChoice) GetPromptLogprobs() ChatCompletionResponseChoiceLogprobs`

GetPromptLogprobs returns the PromptLogprobs field if non-nil, zero value otherwise.

### GetPromptLogprobsOk

`func (o *CompletionResponseStreamChoice) GetPromptLogprobsOk() (*ChatCompletionResponseChoiceLogprobs, bool)`

GetPromptLogprobsOk returns a tuple with the PromptLogprobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptLogprobs

`func (o *CompletionResponseStreamChoice) SetPromptLogprobs(v ChatCompletionResponseChoiceLogprobs)`

SetPromptLogprobs sets PromptLogprobs field to given value.

### HasPromptLogprobs

`func (o *CompletionResponseStreamChoice) HasPromptLogprobs() bool`

HasPromptLogprobs returns a boolean if a field has been set.

### SetPromptLogprobsNil

`func (o *CompletionResponseStreamChoice) SetPromptLogprobsNil(b bool)`

 SetPromptLogprobsNil sets the value for PromptLogprobs to be an explicit nil

### UnsetPromptLogprobs
`func (o *CompletionResponseStreamChoice) UnsetPromptLogprobs()`

UnsetPromptLogprobs ensures that no value is present for PromptLogprobs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


