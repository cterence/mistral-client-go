# CreateResponseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CreatedAt** | **float32** |  | 
**Error** | Pointer to [**NullableResponsesError**](ResponsesError.md) |  | [optional] 
**IncompleteDetails** | Pointer to [**NullableResponsesIncompleteDetails**](ResponsesIncompleteDetails.md) |  | [optional] 
**Instructions** | Pointer to **NullableString** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Model** | **string** |  | 
**Object** | **string** |  | 
**Output** | [**[]Output2Inner**](Output2Inner.md) |  | 
**ParallelToolCalls** | **bool** |  | 
**Temperature** | Pointer to **NullableFloat32** |  | [optional] 
**ToolChoice** | [**ToolChoice1**](ToolChoice1.md) |  | 
**Tools** | [**[]Tools1Inner**](Tools1Inner.md) |  | 
**TopP** | Pointer to **NullableFloat32** |  | [optional] 
**Background** | Pointer to **NullableBool** |  | [optional] 
**CompletedAt** | Pointer to **NullableFloat32** |  | [optional] 
**MaxOutputTokens** | Pointer to **NullableInt32** |  | [optional] 
**MaxToolCalls** | Pointer to **NullableInt32** |  | [optional] 
**PreviousResponseId** | Pointer to **NullableString** |  | [optional] 
**PromptCacheKey** | Pointer to **NullableString** |  | [optional] 
**PromptCacheRetention** | Pointer to **NullableString** |  | [optional] 
**Reasoning** | Pointer to [**NullableResponsesSupportedReasoning**](ResponsesSupportedReasoning.md) |  | [optional] 
**SafetyIdentifier** | Pointer to **NullableString** |  | [optional] 
**ServiceTier** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**Text** | Pointer to [**NullableResponsesTextConfig**](ResponsesTextConfig.md) |  | [optional] 
**TopLogprobs** | Pointer to **NullableInt32** |  | [optional] 
**Truncation** | Pointer to **NullableString** |  | [optional] 
**Usage** | Pointer to [**NullableResponsesUsage**](ResponsesUsage.md) |  | [optional] 
**User** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCreateResponseResponse

`func NewCreateResponseResponse(id string, createdAt float32, model string, object string, output []Output2Inner, parallelToolCalls bool, toolChoice ToolChoice1, tools []Tools1Inner, ) *CreateResponseResponse`

NewCreateResponseResponse instantiates a new CreateResponseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateResponseResponseWithDefaults

`func NewCreateResponseResponseWithDefaults() *CreateResponseResponse`

NewCreateResponseResponseWithDefaults instantiates a new CreateResponseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateResponseResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateResponseResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateResponseResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *CreateResponseResponse) GetCreatedAt() float32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CreateResponseResponse) GetCreatedAtOk() (*float32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CreateResponseResponse) SetCreatedAt(v float32)`

SetCreatedAt sets CreatedAt field to given value.


### GetError

`func (o *CreateResponseResponse) GetError() ResponsesError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *CreateResponseResponse) GetErrorOk() (*ResponsesError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *CreateResponseResponse) SetError(v ResponsesError)`

SetError sets Error field to given value.

### HasError

`func (o *CreateResponseResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *CreateResponseResponse) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *CreateResponseResponse) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetIncompleteDetails

`func (o *CreateResponseResponse) GetIncompleteDetails() ResponsesIncompleteDetails`

GetIncompleteDetails returns the IncompleteDetails field if non-nil, zero value otherwise.

### GetIncompleteDetailsOk

`func (o *CreateResponseResponse) GetIncompleteDetailsOk() (*ResponsesIncompleteDetails, bool)`

GetIncompleteDetailsOk returns a tuple with the IncompleteDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncompleteDetails

`func (o *CreateResponseResponse) SetIncompleteDetails(v ResponsesIncompleteDetails)`

SetIncompleteDetails sets IncompleteDetails field to given value.

### HasIncompleteDetails

`func (o *CreateResponseResponse) HasIncompleteDetails() bool`

HasIncompleteDetails returns a boolean if a field has been set.

### SetIncompleteDetailsNil

`func (o *CreateResponseResponse) SetIncompleteDetailsNil(b bool)`

 SetIncompleteDetailsNil sets the value for IncompleteDetails to be an explicit nil

### UnsetIncompleteDetails
`func (o *CreateResponseResponse) UnsetIncompleteDetails()`

UnsetIncompleteDetails ensures that no value is present for IncompleteDetails, not even an explicit nil
### GetInstructions

`func (o *CreateResponseResponse) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *CreateResponseResponse) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *CreateResponseResponse) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *CreateResponseResponse) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### SetInstructionsNil

`func (o *CreateResponseResponse) SetInstructionsNil(b bool)`

 SetInstructionsNil sets the value for Instructions to be an explicit nil

### UnsetInstructions
`func (o *CreateResponseResponse) UnsetInstructions()`

UnsetInstructions ensures that no value is present for Instructions, not even an explicit nil
### GetMetadata

`func (o *CreateResponseResponse) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateResponseResponse) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateResponseResponse) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateResponseResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *CreateResponseResponse) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *CreateResponseResponse) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetModel

`func (o *CreateResponseResponse) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *CreateResponseResponse) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *CreateResponseResponse) SetModel(v string)`

SetModel sets Model field to given value.


### GetObject

`func (o *CreateResponseResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *CreateResponseResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *CreateResponseResponse) SetObject(v string)`

SetObject sets Object field to given value.


### GetOutput

`func (o *CreateResponseResponse) GetOutput() []Output2Inner`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *CreateResponseResponse) GetOutputOk() (*[]Output2Inner, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *CreateResponseResponse) SetOutput(v []Output2Inner)`

SetOutput sets Output field to given value.


### GetParallelToolCalls

`func (o *CreateResponseResponse) GetParallelToolCalls() bool`

GetParallelToolCalls returns the ParallelToolCalls field if non-nil, zero value otherwise.

### GetParallelToolCallsOk

`func (o *CreateResponseResponse) GetParallelToolCallsOk() (*bool, bool)`

GetParallelToolCallsOk returns a tuple with the ParallelToolCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParallelToolCalls

`func (o *CreateResponseResponse) SetParallelToolCalls(v bool)`

SetParallelToolCalls sets ParallelToolCalls field to given value.


### GetTemperature

`func (o *CreateResponseResponse) GetTemperature() float32`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *CreateResponseResponse) GetTemperatureOk() (*float32, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *CreateResponseResponse) SetTemperature(v float32)`

SetTemperature sets Temperature field to given value.

### HasTemperature

`func (o *CreateResponseResponse) HasTemperature() bool`

HasTemperature returns a boolean if a field has been set.

### SetTemperatureNil

`func (o *CreateResponseResponse) SetTemperatureNil(b bool)`

 SetTemperatureNil sets the value for Temperature to be an explicit nil

### UnsetTemperature
`func (o *CreateResponseResponse) UnsetTemperature()`

UnsetTemperature ensures that no value is present for Temperature, not even an explicit nil
### GetToolChoice

`func (o *CreateResponseResponse) GetToolChoice() ToolChoice1`

GetToolChoice returns the ToolChoice field if non-nil, zero value otherwise.

### GetToolChoiceOk

`func (o *CreateResponseResponse) GetToolChoiceOk() (*ToolChoice1, bool)`

GetToolChoiceOk returns a tuple with the ToolChoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolChoice

`func (o *CreateResponseResponse) SetToolChoice(v ToolChoice1)`

SetToolChoice sets ToolChoice field to given value.


### GetTools

`func (o *CreateResponseResponse) GetTools() []Tools1Inner`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *CreateResponseResponse) GetToolsOk() (*[]Tools1Inner, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *CreateResponseResponse) SetTools(v []Tools1Inner)`

SetTools sets Tools field to given value.


### GetTopP

`func (o *CreateResponseResponse) GetTopP() float32`

GetTopP returns the TopP field if non-nil, zero value otherwise.

### GetTopPOk

`func (o *CreateResponseResponse) GetTopPOk() (*float32, bool)`

GetTopPOk returns a tuple with the TopP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopP

`func (o *CreateResponseResponse) SetTopP(v float32)`

SetTopP sets TopP field to given value.

### HasTopP

`func (o *CreateResponseResponse) HasTopP() bool`

HasTopP returns a boolean if a field has been set.

### SetTopPNil

`func (o *CreateResponseResponse) SetTopPNil(b bool)`

 SetTopPNil sets the value for TopP to be an explicit nil

### UnsetTopP
`func (o *CreateResponseResponse) UnsetTopP()`

UnsetTopP ensures that no value is present for TopP, not even an explicit nil
### GetBackground

`func (o *CreateResponseResponse) GetBackground() bool`

GetBackground returns the Background field if non-nil, zero value otherwise.

### GetBackgroundOk

`func (o *CreateResponseResponse) GetBackgroundOk() (*bool, bool)`

GetBackgroundOk returns a tuple with the Background field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackground

`func (o *CreateResponseResponse) SetBackground(v bool)`

SetBackground sets Background field to given value.

### HasBackground

`func (o *CreateResponseResponse) HasBackground() bool`

HasBackground returns a boolean if a field has been set.

### SetBackgroundNil

`func (o *CreateResponseResponse) SetBackgroundNil(b bool)`

 SetBackgroundNil sets the value for Background to be an explicit nil

### UnsetBackground
`func (o *CreateResponseResponse) UnsetBackground()`

UnsetBackground ensures that no value is present for Background, not even an explicit nil
### GetCompletedAt

`func (o *CreateResponseResponse) GetCompletedAt() float32`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *CreateResponseResponse) GetCompletedAtOk() (*float32, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *CreateResponseResponse) SetCompletedAt(v float32)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *CreateResponseResponse) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### SetCompletedAtNil

`func (o *CreateResponseResponse) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *CreateResponseResponse) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
### GetMaxOutputTokens

`func (o *CreateResponseResponse) GetMaxOutputTokens() int32`

GetMaxOutputTokens returns the MaxOutputTokens field if non-nil, zero value otherwise.

### GetMaxOutputTokensOk

`func (o *CreateResponseResponse) GetMaxOutputTokensOk() (*int32, bool)`

GetMaxOutputTokensOk returns a tuple with the MaxOutputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxOutputTokens

`func (o *CreateResponseResponse) SetMaxOutputTokens(v int32)`

SetMaxOutputTokens sets MaxOutputTokens field to given value.

### HasMaxOutputTokens

`func (o *CreateResponseResponse) HasMaxOutputTokens() bool`

HasMaxOutputTokens returns a boolean if a field has been set.

### SetMaxOutputTokensNil

`func (o *CreateResponseResponse) SetMaxOutputTokensNil(b bool)`

 SetMaxOutputTokensNil sets the value for MaxOutputTokens to be an explicit nil

### UnsetMaxOutputTokens
`func (o *CreateResponseResponse) UnsetMaxOutputTokens()`

UnsetMaxOutputTokens ensures that no value is present for MaxOutputTokens, not even an explicit nil
### GetMaxToolCalls

`func (o *CreateResponseResponse) GetMaxToolCalls() int32`

GetMaxToolCalls returns the MaxToolCalls field if non-nil, zero value otherwise.

### GetMaxToolCallsOk

`func (o *CreateResponseResponse) GetMaxToolCallsOk() (*int32, bool)`

GetMaxToolCallsOk returns a tuple with the MaxToolCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxToolCalls

`func (o *CreateResponseResponse) SetMaxToolCalls(v int32)`

SetMaxToolCalls sets MaxToolCalls field to given value.

### HasMaxToolCalls

`func (o *CreateResponseResponse) HasMaxToolCalls() bool`

HasMaxToolCalls returns a boolean if a field has been set.

### SetMaxToolCallsNil

`func (o *CreateResponseResponse) SetMaxToolCallsNil(b bool)`

 SetMaxToolCallsNil sets the value for MaxToolCalls to be an explicit nil

### UnsetMaxToolCalls
`func (o *CreateResponseResponse) UnsetMaxToolCalls()`

UnsetMaxToolCalls ensures that no value is present for MaxToolCalls, not even an explicit nil
### GetPreviousResponseId

`func (o *CreateResponseResponse) GetPreviousResponseId() string`

GetPreviousResponseId returns the PreviousResponseId field if non-nil, zero value otherwise.

### GetPreviousResponseIdOk

`func (o *CreateResponseResponse) GetPreviousResponseIdOk() (*string, bool)`

GetPreviousResponseIdOk returns a tuple with the PreviousResponseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousResponseId

`func (o *CreateResponseResponse) SetPreviousResponseId(v string)`

SetPreviousResponseId sets PreviousResponseId field to given value.

### HasPreviousResponseId

`func (o *CreateResponseResponse) HasPreviousResponseId() bool`

HasPreviousResponseId returns a boolean if a field has been set.

### SetPreviousResponseIdNil

`func (o *CreateResponseResponse) SetPreviousResponseIdNil(b bool)`

 SetPreviousResponseIdNil sets the value for PreviousResponseId to be an explicit nil

### UnsetPreviousResponseId
`func (o *CreateResponseResponse) UnsetPreviousResponseId()`

UnsetPreviousResponseId ensures that no value is present for PreviousResponseId, not even an explicit nil
### GetPromptCacheKey

`func (o *CreateResponseResponse) GetPromptCacheKey() string`

GetPromptCacheKey returns the PromptCacheKey field if non-nil, zero value otherwise.

### GetPromptCacheKeyOk

`func (o *CreateResponseResponse) GetPromptCacheKeyOk() (*string, bool)`

GetPromptCacheKeyOk returns a tuple with the PromptCacheKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptCacheKey

`func (o *CreateResponseResponse) SetPromptCacheKey(v string)`

SetPromptCacheKey sets PromptCacheKey field to given value.

### HasPromptCacheKey

`func (o *CreateResponseResponse) HasPromptCacheKey() bool`

HasPromptCacheKey returns a boolean if a field has been set.

### SetPromptCacheKeyNil

`func (o *CreateResponseResponse) SetPromptCacheKeyNil(b bool)`

 SetPromptCacheKeyNil sets the value for PromptCacheKey to be an explicit nil

### UnsetPromptCacheKey
`func (o *CreateResponseResponse) UnsetPromptCacheKey()`

UnsetPromptCacheKey ensures that no value is present for PromptCacheKey, not even an explicit nil
### GetPromptCacheRetention

`func (o *CreateResponseResponse) GetPromptCacheRetention() string`

GetPromptCacheRetention returns the PromptCacheRetention field if non-nil, zero value otherwise.

### GetPromptCacheRetentionOk

`func (o *CreateResponseResponse) GetPromptCacheRetentionOk() (*string, bool)`

GetPromptCacheRetentionOk returns a tuple with the PromptCacheRetention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptCacheRetention

`func (o *CreateResponseResponse) SetPromptCacheRetention(v string)`

SetPromptCacheRetention sets PromptCacheRetention field to given value.

### HasPromptCacheRetention

`func (o *CreateResponseResponse) HasPromptCacheRetention() bool`

HasPromptCacheRetention returns a boolean if a field has been set.

### SetPromptCacheRetentionNil

`func (o *CreateResponseResponse) SetPromptCacheRetentionNil(b bool)`

 SetPromptCacheRetentionNil sets the value for PromptCacheRetention to be an explicit nil

### UnsetPromptCacheRetention
`func (o *CreateResponseResponse) UnsetPromptCacheRetention()`

UnsetPromptCacheRetention ensures that no value is present for PromptCacheRetention, not even an explicit nil
### GetReasoning

`func (o *CreateResponseResponse) GetReasoning() ResponsesSupportedReasoning`

GetReasoning returns the Reasoning field if non-nil, zero value otherwise.

### GetReasoningOk

`func (o *CreateResponseResponse) GetReasoningOk() (*ResponsesSupportedReasoning, bool)`

GetReasoningOk returns a tuple with the Reasoning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasoning

`func (o *CreateResponseResponse) SetReasoning(v ResponsesSupportedReasoning)`

SetReasoning sets Reasoning field to given value.

### HasReasoning

`func (o *CreateResponseResponse) HasReasoning() bool`

HasReasoning returns a boolean if a field has been set.

### SetReasoningNil

`func (o *CreateResponseResponse) SetReasoningNil(b bool)`

 SetReasoningNil sets the value for Reasoning to be an explicit nil

### UnsetReasoning
`func (o *CreateResponseResponse) UnsetReasoning()`

UnsetReasoning ensures that no value is present for Reasoning, not even an explicit nil
### GetSafetyIdentifier

`func (o *CreateResponseResponse) GetSafetyIdentifier() string`

GetSafetyIdentifier returns the SafetyIdentifier field if non-nil, zero value otherwise.

### GetSafetyIdentifierOk

`func (o *CreateResponseResponse) GetSafetyIdentifierOk() (*string, bool)`

GetSafetyIdentifierOk returns a tuple with the SafetyIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafetyIdentifier

`func (o *CreateResponseResponse) SetSafetyIdentifier(v string)`

SetSafetyIdentifier sets SafetyIdentifier field to given value.

### HasSafetyIdentifier

`func (o *CreateResponseResponse) HasSafetyIdentifier() bool`

HasSafetyIdentifier returns a boolean if a field has been set.

### SetSafetyIdentifierNil

`func (o *CreateResponseResponse) SetSafetyIdentifierNil(b bool)`

 SetSafetyIdentifierNil sets the value for SafetyIdentifier to be an explicit nil

### UnsetSafetyIdentifier
`func (o *CreateResponseResponse) UnsetSafetyIdentifier()`

UnsetSafetyIdentifier ensures that no value is present for SafetyIdentifier, not even an explicit nil
### GetServiceTier

`func (o *CreateResponseResponse) GetServiceTier() string`

GetServiceTier returns the ServiceTier field if non-nil, zero value otherwise.

### GetServiceTierOk

`func (o *CreateResponseResponse) GetServiceTierOk() (*string, bool)`

GetServiceTierOk returns a tuple with the ServiceTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTier

`func (o *CreateResponseResponse) SetServiceTier(v string)`

SetServiceTier sets ServiceTier field to given value.

### HasServiceTier

`func (o *CreateResponseResponse) HasServiceTier() bool`

HasServiceTier returns a boolean if a field has been set.

### SetServiceTierNil

`func (o *CreateResponseResponse) SetServiceTierNil(b bool)`

 SetServiceTierNil sets the value for ServiceTier to be an explicit nil

### UnsetServiceTier
`func (o *CreateResponseResponse) UnsetServiceTier()`

UnsetServiceTier ensures that no value is present for ServiceTier, not even an explicit nil
### GetStatus

`func (o *CreateResponseResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateResponseResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateResponseResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateResponseResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *CreateResponseResponse) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *CreateResponseResponse) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetText

`func (o *CreateResponseResponse) GetText() ResponsesTextConfig`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *CreateResponseResponse) GetTextOk() (*ResponsesTextConfig, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *CreateResponseResponse) SetText(v ResponsesTextConfig)`

SetText sets Text field to given value.

### HasText

`func (o *CreateResponseResponse) HasText() bool`

HasText returns a boolean if a field has been set.

### SetTextNil

`func (o *CreateResponseResponse) SetTextNil(b bool)`

 SetTextNil sets the value for Text to be an explicit nil

### UnsetText
`func (o *CreateResponseResponse) UnsetText()`

UnsetText ensures that no value is present for Text, not even an explicit nil
### GetTopLogprobs

`func (o *CreateResponseResponse) GetTopLogprobs() int32`

GetTopLogprobs returns the TopLogprobs field if non-nil, zero value otherwise.

### GetTopLogprobsOk

`func (o *CreateResponseResponse) GetTopLogprobsOk() (*int32, bool)`

GetTopLogprobsOk returns a tuple with the TopLogprobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopLogprobs

`func (o *CreateResponseResponse) SetTopLogprobs(v int32)`

SetTopLogprobs sets TopLogprobs field to given value.

### HasTopLogprobs

`func (o *CreateResponseResponse) HasTopLogprobs() bool`

HasTopLogprobs returns a boolean if a field has been set.

### SetTopLogprobsNil

`func (o *CreateResponseResponse) SetTopLogprobsNil(b bool)`

 SetTopLogprobsNil sets the value for TopLogprobs to be an explicit nil

### UnsetTopLogprobs
`func (o *CreateResponseResponse) UnsetTopLogprobs()`

UnsetTopLogprobs ensures that no value is present for TopLogprobs, not even an explicit nil
### GetTruncation

`func (o *CreateResponseResponse) GetTruncation() string`

GetTruncation returns the Truncation field if non-nil, zero value otherwise.

### GetTruncationOk

`func (o *CreateResponseResponse) GetTruncationOk() (*string, bool)`

GetTruncationOk returns a tuple with the Truncation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncation

`func (o *CreateResponseResponse) SetTruncation(v string)`

SetTruncation sets Truncation field to given value.

### HasTruncation

`func (o *CreateResponseResponse) HasTruncation() bool`

HasTruncation returns a boolean if a field has been set.

### SetTruncationNil

`func (o *CreateResponseResponse) SetTruncationNil(b bool)`

 SetTruncationNil sets the value for Truncation to be an explicit nil

### UnsetTruncation
`func (o *CreateResponseResponse) UnsetTruncation()`

UnsetTruncation ensures that no value is present for Truncation, not even an explicit nil
### GetUsage

`func (o *CreateResponseResponse) GetUsage() ResponsesUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *CreateResponseResponse) GetUsageOk() (*ResponsesUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *CreateResponseResponse) SetUsage(v ResponsesUsage)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *CreateResponseResponse) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### SetUsageNil

`func (o *CreateResponseResponse) SetUsageNil(b bool)`

 SetUsageNil sets the value for Usage to be an explicit nil

### UnsetUsage
`func (o *CreateResponseResponse) UnsetUsage()`

UnsetUsage ensures that no value is present for Usage, not even an explicit nil
### GetUser

`func (o *CreateResponseResponse) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *CreateResponseResponse) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *CreateResponseResponse) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *CreateResponseResponse) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *CreateResponseResponse) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *CreateResponseResponse) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


