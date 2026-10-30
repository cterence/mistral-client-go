# CreateResponseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | **string** |  | 
**Input** | [**Input6**](Input6.md) |  | 
**Instructions** | Pointer to **NullableString** |  | [optional] 
**ParallelToolCalls** | Pointer to **bool** |  | [optional] [default to true]
**PreviousResponseId** | Pointer to **interface{}** |  | [optional] 
**Truncation** | Pointer to **string** |  | [optional] [default to "disabled"]
**Tools** | Pointer to [**[]Tools1Inner**](Tools1Inner.md) |  | [optional] 
**ToolChoice** | Pointer to [**ToolChoice2**](ToolChoice2.md) |  | [optional] [default to auto]
**Reasoning** | Pointer to [**NullableResponsesSupportedReasoning**](ResponsesSupportedReasoning.md) |  | [optional] 
**Conversation** | Pointer to **interface{}** |  | [optional] 
**Prompt** | Pointer to **interface{}** |  | [optional] 
**Text** | Pointer to [**NullableResponsesTextConfig**](ResponsesTextConfig.md) |  | [optional] 
**Background** | Pointer to **bool** |  | [optional] [default to false]
**Stream** | Pointer to **bool** |  | [optional] [default to false]
**Store** | Pointer to **NullableBool** |  | [optional] 
**MaxOutputTokens** | Pointer to **NullableInt32** |  | [optional] 
**Temperature** | Pointer to **NullableFloat32** |  | [optional] 
**TopP** | Pointer to **NullableFloat32** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**PromptCacheKey** | Pointer to **NullableString** |  | [optional] 
**MaxToolCalls** | Pointer to **NullableInt32** |  | [optional] 
**Moderation** | Pointer to **interface{}** |  | [optional] 
**TopLogprobs** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewCreateResponseRequest

`func NewCreateResponseRequest(model string, input Input6, ) *CreateResponseRequest`

NewCreateResponseRequest instantiates a new CreateResponseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateResponseRequestWithDefaults

`func NewCreateResponseRequestWithDefaults() *CreateResponseRequest`

NewCreateResponseRequestWithDefaults instantiates a new CreateResponseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *CreateResponseRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *CreateResponseRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *CreateResponseRequest) SetModel(v string)`

SetModel sets Model field to given value.


### GetInput

`func (o *CreateResponseRequest) GetInput() Input6`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *CreateResponseRequest) GetInputOk() (*Input6, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *CreateResponseRequest) SetInput(v Input6)`

SetInput sets Input field to given value.


### GetInstructions

`func (o *CreateResponseRequest) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *CreateResponseRequest) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *CreateResponseRequest) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *CreateResponseRequest) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### SetInstructionsNil

`func (o *CreateResponseRequest) SetInstructionsNil(b bool)`

 SetInstructionsNil sets the value for Instructions to be an explicit nil

### UnsetInstructions
`func (o *CreateResponseRequest) UnsetInstructions()`

UnsetInstructions ensures that no value is present for Instructions, not even an explicit nil
### GetParallelToolCalls

`func (o *CreateResponseRequest) GetParallelToolCalls() bool`

GetParallelToolCalls returns the ParallelToolCalls field if non-nil, zero value otherwise.

### GetParallelToolCallsOk

`func (o *CreateResponseRequest) GetParallelToolCallsOk() (*bool, bool)`

GetParallelToolCallsOk returns a tuple with the ParallelToolCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParallelToolCalls

`func (o *CreateResponseRequest) SetParallelToolCalls(v bool)`

SetParallelToolCalls sets ParallelToolCalls field to given value.

### HasParallelToolCalls

`func (o *CreateResponseRequest) HasParallelToolCalls() bool`

HasParallelToolCalls returns a boolean if a field has been set.

### GetPreviousResponseId

`func (o *CreateResponseRequest) GetPreviousResponseId() interface{}`

GetPreviousResponseId returns the PreviousResponseId field if non-nil, zero value otherwise.

### GetPreviousResponseIdOk

`func (o *CreateResponseRequest) GetPreviousResponseIdOk() (*interface{}, bool)`

GetPreviousResponseIdOk returns a tuple with the PreviousResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousResponseId

`func (o *CreateResponseRequest) SetPreviousResponseId(v interface{})`

SetPreviousResponseId sets PreviousResponseId field to given value.

### HasPreviousResponseId

`func (o *CreateResponseRequest) HasPreviousResponseId() bool`

HasPreviousResponseId returns a boolean if a field has been set.

### SetPreviousResponseIdNil

`func (o *CreateResponseRequest) SetPreviousResponseIdNil(b bool)`

 SetPreviousResponseIdNil sets the value for PreviousResponseId to be an explicit nil

### UnsetPreviousResponseId
`func (o *CreateResponseRequest) UnsetPreviousResponseId()`

UnsetPreviousResponseId ensures that no value is present for PreviousResponseId, not even an explicit nil
### GetTruncation

`func (o *CreateResponseRequest) GetTruncation() string`

GetTruncation returns the Truncation field if non-nil, zero value otherwise.

### GetTruncationOk

`func (o *CreateResponseRequest) GetTruncationOk() (*string, bool)`

GetTruncationOk returns a tuple with the Truncation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncation

`func (o *CreateResponseRequest) SetTruncation(v string)`

SetTruncation sets Truncation field to given value.

### HasTruncation

`func (o *CreateResponseRequest) HasTruncation() bool`

HasTruncation returns a boolean if a field has been set.

### GetTools

`func (o *CreateResponseRequest) GetTools() []Tools1Inner`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *CreateResponseRequest) GetToolsOk() (*[]Tools1Inner, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *CreateResponseRequest) SetTools(v []Tools1Inner)`

SetTools sets Tools field to given value.

### HasTools

`func (o *CreateResponseRequest) HasTools() bool`

HasTools returns a boolean if a field has been set.

### GetToolChoice

`func (o *CreateResponseRequest) GetToolChoice() ToolChoice2`

GetToolChoice returns the ToolChoice field if non-nil, zero value otherwise.

### GetToolChoiceOk

`func (o *CreateResponseRequest) GetToolChoiceOk() (*ToolChoice2, bool)`

GetToolChoiceOk returns a tuple with the ToolChoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolChoice

`func (o *CreateResponseRequest) SetToolChoice(v ToolChoice2)`

SetToolChoice sets ToolChoice field to given value.

### HasToolChoice

`func (o *CreateResponseRequest) HasToolChoice() bool`

HasToolChoice returns a boolean if a field has been set.

### GetReasoning

`func (o *CreateResponseRequest) GetReasoning() ResponsesSupportedReasoning`

GetReasoning returns the Reasoning field if non-nil, zero value otherwise.

### GetReasoningOk

`func (o *CreateResponseRequest) GetReasoningOk() (*ResponsesSupportedReasoning, bool)`

GetReasoningOk returns a tuple with the Reasoning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasoning

`func (o *CreateResponseRequest) SetReasoning(v ResponsesSupportedReasoning)`

SetReasoning sets Reasoning field to given value.

### HasReasoning

`func (o *CreateResponseRequest) HasReasoning() bool`

HasReasoning returns a boolean if a field has been set.

### SetReasoningNil

`func (o *CreateResponseRequest) SetReasoningNil(b bool)`

 SetReasoningNil sets the value for Reasoning to be an explicit nil

### UnsetReasoning
`func (o *CreateResponseRequest) UnsetReasoning()`

UnsetReasoning ensures that no value is present for Reasoning, not even an explicit nil
### GetConversation

`func (o *CreateResponseRequest) GetConversation() interface{}`

GetConversation returns the Conversation field if non-nil, zero value otherwise.

### GetConversationOk

`func (o *CreateResponseRequest) GetConversationOk() (*interface{}, bool)`

GetConversationOk returns a tuple with the Conversation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversation

`func (o *CreateResponseRequest) SetConversation(v interface{})`

SetConversation sets Conversation field to given value.

### HasConversation

`func (o *CreateResponseRequest) HasConversation() bool`

HasConversation returns a boolean if a field has been set.

### SetConversationNil

`func (o *CreateResponseRequest) SetConversationNil(b bool)`

 SetConversationNil sets the value for Conversation to be an explicit nil

### UnsetConversation
`func (o *CreateResponseRequest) UnsetConversation()`

UnsetConversation ensures that no value is present for Conversation, not even an explicit nil
### GetPrompt

`func (o *CreateResponseRequest) GetPrompt() interface{}`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *CreateResponseRequest) GetPromptOk() (*interface{}, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *CreateResponseRequest) SetPrompt(v interface{})`

SetPrompt sets Prompt field to given value.

### HasPrompt

`func (o *CreateResponseRequest) HasPrompt() bool`

HasPrompt returns a boolean if a field has been set.

### SetPromptNil

`func (o *CreateResponseRequest) SetPromptNil(b bool)`

 SetPromptNil sets the value for Prompt to be an explicit nil

### UnsetPrompt
`func (o *CreateResponseRequest) UnsetPrompt()`

UnsetPrompt ensures that no value is present for Prompt, not even an explicit nil
### GetText

`func (o *CreateResponseRequest) GetText() ResponsesTextConfig`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *CreateResponseRequest) GetTextOk() (*ResponsesTextConfig, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *CreateResponseRequest) SetText(v ResponsesTextConfig)`

SetText sets Text field to given value.

### HasText

`func (o *CreateResponseRequest) HasText() bool`

HasText returns a boolean if a field has been set.

### SetTextNil

`func (o *CreateResponseRequest) SetTextNil(b bool)`

 SetTextNil sets the value for Text to be an explicit nil

### UnsetText
`func (o *CreateResponseRequest) UnsetText()`

UnsetText ensures that no value is present for Text, not even an explicit nil
### GetBackground

`func (o *CreateResponseRequest) GetBackground() bool`

GetBackground returns the Background field if non-nil, zero value otherwise.

### GetBackgroundOk

`func (o *CreateResponseRequest) GetBackgroundOk() (*bool, bool)`

GetBackgroundOk returns a tuple with the Background field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackground

`func (o *CreateResponseRequest) SetBackground(v bool)`

SetBackground sets Background field to given value.

### HasBackground

`func (o *CreateResponseRequest) HasBackground() bool`

HasBackground returns a boolean if a field has been set.

### GetStream

`func (o *CreateResponseRequest) GetStream() bool`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *CreateResponseRequest) GetStreamOk() (*bool, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *CreateResponseRequest) SetStream(v bool)`

SetStream sets Stream field to given value.

### HasStream

`func (o *CreateResponseRequest) HasStream() bool`

HasStream returns a boolean if a field has been set.

### GetStore

`func (o *CreateResponseRequest) GetStore() bool`

GetStore returns the Store field if non-nil, zero value otherwise.

### GetStoreOk

`func (o *CreateResponseRequest) GetStoreOk() (*bool, bool)`

GetStoreOk returns a tuple with the Store field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStore

`func (o *CreateResponseRequest) SetStore(v bool)`

SetStore sets Store field to given value.

### HasStore

`func (o *CreateResponseRequest) HasStore() bool`

HasStore returns a boolean if a field has been set.

### SetStoreNil

`func (o *CreateResponseRequest) SetStoreNil(b bool)`

 SetStoreNil sets the value for Store to be an explicit nil

### UnsetStore
`func (o *CreateResponseRequest) UnsetStore()`

UnsetStore ensures that no value is present for Store, not even an explicit nil
### GetMaxOutputTokens

`func (o *CreateResponseRequest) GetMaxOutputTokens() int32`

GetMaxOutputTokens returns the MaxOutputTokens field if non-nil, zero value otherwise.

### GetMaxOutputTokensOk

`func (o *CreateResponseRequest) GetMaxOutputTokensOk() (*int32, bool)`

GetMaxOutputTokensOk returns a tuple with the MaxOutputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxOutputTokens

`func (o *CreateResponseRequest) SetMaxOutputTokens(v int32)`

SetMaxOutputTokens sets MaxOutputTokens field to given value.

### HasMaxOutputTokens

`func (o *CreateResponseRequest) HasMaxOutputTokens() bool`

HasMaxOutputTokens returns a boolean if a field has been set.

### SetMaxOutputTokensNil

`func (o *CreateResponseRequest) SetMaxOutputTokensNil(b bool)`

 SetMaxOutputTokensNil sets the value for MaxOutputTokens to be an explicit nil

### UnsetMaxOutputTokens
`func (o *CreateResponseRequest) UnsetMaxOutputTokens()`

UnsetMaxOutputTokens ensures that no value is present for MaxOutputTokens, not even an explicit nil
### GetTemperature

`func (o *CreateResponseRequest) GetTemperature() float32`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *CreateResponseRequest) GetTemperatureOk() (*float32, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *CreateResponseRequest) SetTemperature(v float32)`

SetTemperature sets Temperature field to given value.

### HasTemperature

`func (o *CreateResponseRequest) HasTemperature() bool`

HasTemperature returns a boolean if a field has been set.

### SetTemperatureNil

`func (o *CreateResponseRequest) SetTemperatureNil(b bool)`

 SetTemperatureNil sets the value for Temperature to be an explicit nil

### UnsetTemperature
`func (o *CreateResponseRequest) UnsetTemperature()`

UnsetTemperature ensures that no value is present for Temperature, not even an explicit nil
### GetTopP

`func (o *CreateResponseRequest) GetTopP() float32`

GetTopP returns the TopP field if non-nil, zero value otherwise.

### GetTopPOk

`func (o *CreateResponseRequest) GetTopPOk() (*float32, bool)`

GetTopPOk returns a tuple with the TopP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopP

`func (o *CreateResponseRequest) SetTopP(v float32)`

SetTopP sets TopP field to given value.

### HasTopP

`func (o *CreateResponseRequest) HasTopP() bool`

HasTopP returns a boolean if a field has been set.

### SetTopPNil

`func (o *CreateResponseRequest) SetTopPNil(b bool)`

 SetTopPNil sets the value for TopP to be an explicit nil

### UnsetTopP
`func (o *CreateResponseRequest) UnsetTopP()`

UnsetTopP ensures that no value is present for TopP, not even an explicit nil
### GetMetadata

`func (o *CreateResponseRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateResponseRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateResponseRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateResponseRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *CreateResponseRequest) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *CreateResponseRequest) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetPromptCacheKey

`func (o *CreateResponseRequest) GetPromptCacheKey() string`

GetPromptCacheKey returns the PromptCacheKey field if non-nil, zero value otherwise.

### GetPromptCacheKeyOk

`func (o *CreateResponseRequest) GetPromptCacheKeyOk() (*string, bool)`

GetPromptCacheKeyOk returns a tuple with the PromptCacheKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptCacheKey

`func (o *CreateResponseRequest) SetPromptCacheKey(v string)`

SetPromptCacheKey sets PromptCacheKey field to given value.

### HasPromptCacheKey

`func (o *CreateResponseRequest) HasPromptCacheKey() bool`

HasPromptCacheKey returns a boolean if a field has been set.

### SetPromptCacheKeyNil

`func (o *CreateResponseRequest) SetPromptCacheKeyNil(b bool)`

 SetPromptCacheKeyNil sets the value for PromptCacheKey to be an explicit nil

### UnsetPromptCacheKey
`func (o *CreateResponseRequest) UnsetPromptCacheKey()`

UnsetPromptCacheKey ensures that no value is present for PromptCacheKey, not even an explicit nil
### GetMaxToolCalls

`func (o *CreateResponseRequest) GetMaxToolCalls() int32`

GetMaxToolCalls returns the MaxToolCalls field if non-nil, zero value otherwise.

### GetMaxToolCallsOk

`func (o *CreateResponseRequest) GetMaxToolCallsOk() (*int32, bool)`

GetMaxToolCallsOk returns a tuple with the MaxToolCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxToolCalls

`func (o *CreateResponseRequest) SetMaxToolCalls(v int32)`

SetMaxToolCalls sets MaxToolCalls field to given value.

### HasMaxToolCalls

`func (o *CreateResponseRequest) HasMaxToolCalls() bool`

HasMaxToolCalls returns a boolean if a field has been set.

### SetMaxToolCallsNil

`func (o *CreateResponseRequest) SetMaxToolCallsNil(b bool)`

 SetMaxToolCallsNil sets the value for MaxToolCalls to be an explicit nil

### UnsetMaxToolCalls
`func (o *CreateResponseRequest) UnsetMaxToolCalls()`

UnsetMaxToolCalls ensures that no value is present for MaxToolCalls, not even an explicit nil
### GetModeration

`func (o *CreateResponseRequest) GetModeration() interface{}`

GetModeration returns the Moderation field if non-nil, zero value otherwise.

### GetModerationOk

`func (o *CreateResponseRequest) GetModerationOk() (*interface{}, bool)`

GetModerationOk returns a tuple with the Moderation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModeration

`func (o *CreateResponseRequest) SetModeration(v interface{})`

SetModeration sets Moderation field to given value.

### HasModeration

`func (o *CreateResponseRequest) HasModeration() bool`

HasModeration returns a boolean if a field has been set.

### SetModerationNil

`func (o *CreateResponseRequest) SetModerationNil(b bool)`

 SetModerationNil sets the value for Moderation to be an explicit nil

### UnsetModeration
`func (o *CreateResponseRequest) UnsetModeration()`

UnsetModeration ensures that no value is present for Moderation, not even an explicit nil
### GetTopLogprobs

`func (o *CreateResponseRequest) GetTopLogprobs() int32`

GetTopLogprobs returns the TopLogprobs field if non-nil, zero value otherwise.

### GetTopLogprobsOk

`func (o *CreateResponseRequest) GetTopLogprobsOk() (*int32, bool)`

GetTopLogprobsOk returns a tuple with the TopLogprobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopLogprobs

`func (o *CreateResponseRequest) SetTopLogprobs(v int32)`

SetTopLogprobs sets TopLogprobs field to given value.

### HasTopLogprobs

`func (o *CreateResponseRequest) HasTopLogprobs() bool`

HasTopLogprobs returns a boolean if a field has been set.

### SetTopLogprobsNil

`func (o *CreateResponseRequest) SetTopLogprobsNil(b bool)`

 SetTopLogprobsNil sets the value for TopLogprobs to be an explicit nil

### UnsetTopLogprobs
`func (o *CreateResponseRequest) UnsetTopLogprobs()`

UnsetTopLogprobs ensures that no value is present for TopLogprobs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


