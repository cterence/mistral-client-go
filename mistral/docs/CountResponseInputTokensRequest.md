# CountResponseInputTokensRequest

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

## Methods

### NewCountResponseInputTokensRequest

`func NewCountResponseInputTokensRequest(model string, input Input6, ) *CountResponseInputTokensRequest`

NewCountResponseInputTokensRequest instantiates a new CountResponseInputTokensRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCountResponseInputTokensRequestWithDefaults

`func NewCountResponseInputTokensRequestWithDefaults() *CountResponseInputTokensRequest`

NewCountResponseInputTokensRequestWithDefaults instantiates a new CountResponseInputTokensRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *CountResponseInputTokensRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *CountResponseInputTokensRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *CountResponseInputTokensRequest) SetModel(v string)`

SetModel sets Model field to given value.


### GetInput

`func (o *CountResponseInputTokensRequest) GetInput() Input6`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *CountResponseInputTokensRequest) GetInputOk() (*Input6, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *CountResponseInputTokensRequest) SetInput(v Input6)`

SetInput sets Input field to given value.


### GetInstructions

`func (o *CountResponseInputTokensRequest) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *CountResponseInputTokensRequest) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *CountResponseInputTokensRequest) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *CountResponseInputTokensRequest) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### SetInstructionsNil

`func (o *CountResponseInputTokensRequest) SetInstructionsNil(b bool)`

 SetInstructionsNil sets the value for Instructions to be an explicit nil

### UnsetInstructions
`func (o *CountResponseInputTokensRequest) UnsetInstructions()`

UnsetInstructions ensures that no value is present for Instructions, not even an explicit nil
### GetParallelToolCalls

`func (o *CountResponseInputTokensRequest) GetParallelToolCalls() bool`

GetParallelToolCalls returns the ParallelToolCalls field if non-nil, zero value otherwise.

### GetParallelToolCallsOk

`func (o *CountResponseInputTokensRequest) GetParallelToolCallsOk() (*bool, bool)`

GetParallelToolCallsOk returns a tuple with the ParallelToolCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParallelToolCalls

`func (o *CountResponseInputTokensRequest) SetParallelToolCalls(v bool)`

SetParallelToolCalls sets ParallelToolCalls field to given value.

### HasParallelToolCalls

`func (o *CountResponseInputTokensRequest) HasParallelToolCalls() bool`

HasParallelToolCalls returns a boolean if a field has been set.

### GetPreviousResponseId

`func (o *CountResponseInputTokensRequest) GetPreviousResponseId() interface{}`

GetPreviousResponseId returns the PreviousResponseId field if non-nil, zero value otherwise.

### GetPreviousResponseIdOk

`func (o *CountResponseInputTokensRequest) GetPreviousResponseIdOk() (*interface{}, bool)`

GetPreviousResponseIdOk returns a tuple with the PreviousResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousResponseId

`func (o *CountResponseInputTokensRequest) SetPreviousResponseId(v interface{})`

SetPreviousResponseId sets PreviousResponseId field to given value.

### HasPreviousResponseId

`func (o *CountResponseInputTokensRequest) HasPreviousResponseId() bool`

HasPreviousResponseId returns a boolean if a field has been set.

### SetPreviousResponseIdNil

`func (o *CountResponseInputTokensRequest) SetPreviousResponseIdNil(b bool)`

 SetPreviousResponseIdNil sets the value for PreviousResponseId to be an explicit nil

### UnsetPreviousResponseId
`func (o *CountResponseInputTokensRequest) UnsetPreviousResponseId()`

UnsetPreviousResponseId ensures that no value is present for PreviousResponseId, not even an explicit nil
### GetTruncation

`func (o *CountResponseInputTokensRequest) GetTruncation() string`

GetTruncation returns the Truncation field if non-nil, zero value otherwise.

### GetTruncationOk

`func (o *CountResponseInputTokensRequest) GetTruncationOk() (*string, bool)`

GetTruncationOk returns a tuple with the Truncation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncation

`func (o *CountResponseInputTokensRequest) SetTruncation(v string)`

SetTruncation sets Truncation field to given value.

### HasTruncation

`func (o *CountResponseInputTokensRequest) HasTruncation() bool`

HasTruncation returns a boolean if a field has been set.

### GetTools

`func (o *CountResponseInputTokensRequest) GetTools() []Tools1Inner`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *CountResponseInputTokensRequest) GetToolsOk() (*[]Tools1Inner, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *CountResponseInputTokensRequest) SetTools(v []Tools1Inner)`

SetTools sets Tools field to given value.

### HasTools

`func (o *CountResponseInputTokensRequest) HasTools() bool`

HasTools returns a boolean if a field has been set.

### GetToolChoice

`func (o *CountResponseInputTokensRequest) GetToolChoice() ToolChoice2`

GetToolChoice returns the ToolChoice field if non-nil, zero value otherwise.

### GetToolChoiceOk

`func (o *CountResponseInputTokensRequest) GetToolChoiceOk() (*ToolChoice2, bool)`

GetToolChoiceOk returns a tuple with the ToolChoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolChoice

`func (o *CountResponseInputTokensRequest) SetToolChoice(v ToolChoice2)`

SetToolChoice sets ToolChoice field to given value.

### HasToolChoice

`func (o *CountResponseInputTokensRequest) HasToolChoice() bool`

HasToolChoice returns a boolean if a field has been set.

### GetReasoning

`func (o *CountResponseInputTokensRequest) GetReasoning() ResponsesSupportedReasoning`

GetReasoning returns the Reasoning field if non-nil, zero value otherwise.

### GetReasoningOk

`func (o *CountResponseInputTokensRequest) GetReasoningOk() (*ResponsesSupportedReasoning, bool)`

GetReasoningOk returns a tuple with the Reasoning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasoning

`func (o *CountResponseInputTokensRequest) SetReasoning(v ResponsesSupportedReasoning)`

SetReasoning sets Reasoning field to given value.

### HasReasoning

`func (o *CountResponseInputTokensRequest) HasReasoning() bool`

HasReasoning returns a boolean if a field has been set.

### SetReasoningNil

`func (o *CountResponseInputTokensRequest) SetReasoningNil(b bool)`

 SetReasoningNil sets the value for Reasoning to be an explicit nil

### UnsetReasoning
`func (o *CountResponseInputTokensRequest) UnsetReasoning()`

UnsetReasoning ensures that no value is present for Reasoning, not even an explicit nil
### GetConversation

`func (o *CountResponseInputTokensRequest) GetConversation() interface{}`

GetConversation returns the Conversation field if non-nil, zero value otherwise.

### GetConversationOk

`func (o *CountResponseInputTokensRequest) GetConversationOk() (*interface{}, bool)`

GetConversationOk returns a tuple with the Conversation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversation

`func (o *CountResponseInputTokensRequest) SetConversation(v interface{})`

SetConversation sets Conversation field to given value.

### HasConversation

`func (o *CountResponseInputTokensRequest) HasConversation() bool`

HasConversation returns a boolean if a field has been set.

### SetConversationNil

`func (o *CountResponseInputTokensRequest) SetConversationNil(b bool)`

 SetConversationNil sets the value for Conversation to be an explicit nil

### UnsetConversation
`func (o *CountResponseInputTokensRequest) UnsetConversation()`

UnsetConversation ensures that no value is present for Conversation, not even an explicit nil
### GetPrompt

`func (o *CountResponseInputTokensRequest) GetPrompt() interface{}`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *CountResponseInputTokensRequest) GetPromptOk() (*interface{}, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *CountResponseInputTokensRequest) SetPrompt(v interface{})`

SetPrompt sets Prompt field to given value.

### HasPrompt

`func (o *CountResponseInputTokensRequest) HasPrompt() bool`

HasPrompt returns a boolean if a field has been set.

### SetPromptNil

`func (o *CountResponseInputTokensRequest) SetPromptNil(b bool)`

 SetPromptNil sets the value for Prompt to be an explicit nil

### UnsetPrompt
`func (o *CountResponseInputTokensRequest) UnsetPrompt()`

UnsetPrompt ensures that no value is present for Prompt, not even an explicit nil
### GetText

`func (o *CountResponseInputTokensRequest) GetText() ResponsesTextConfig`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *CountResponseInputTokensRequest) GetTextOk() (*ResponsesTextConfig, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *CountResponseInputTokensRequest) SetText(v ResponsesTextConfig)`

SetText sets Text field to given value.

### HasText

`func (o *CountResponseInputTokensRequest) HasText() bool`

HasText returns a boolean if a field has been set.

### SetTextNil

`func (o *CountResponseInputTokensRequest) SetTextNil(b bool)`

 SetTextNil sets the value for Text to be an explicit nil

### UnsetText
`func (o *CountResponseInputTokensRequest) UnsetText()`

UnsetText ensures that no value is present for Text, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


