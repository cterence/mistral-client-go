# ConversationRequestBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Inputs** | [**ConversationInputs**](ConversationInputs.md) |  | 
**Stream** | Pointer to **NullableBool** |  | [optional] 
**Store** | Pointer to **NullableBool** |  | [optional] 
**HandoffExecution** | Pointer to **NullableString** |  | [optional] 
**Instructions** | Pointer to **NullableString** |  | [optional] 
**Tools** | Pointer to [**[]ToolsInner**](ToolsInner.md) |  | [optional] 
**CompletionArgs** | Pointer to [**NullableCompletionArgs**](CompletionArgs.md) |  | [optional] 
**Guardrails** | Pointer to [**[]GuardrailConfig**](GuardrailConfig.md) |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Custom type for metadata with embedded validation. | [optional] 
**AgentId** | Pointer to **NullableString** |  | [optional] 
**AgentVersion** | Pointer to [**NullableAgentVersion1**](AgentVersion1.md) |  | [optional] 
**Model** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewConversationRequestBase

`func NewConversationRequestBase(inputs ConversationInputs, ) *ConversationRequestBase`

NewConversationRequestBase instantiates a new ConversationRequestBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConversationRequestBaseWithDefaults

`func NewConversationRequestBaseWithDefaults() *ConversationRequestBase`

NewConversationRequestBaseWithDefaults instantiates a new ConversationRequestBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInputs

`func (o *ConversationRequestBase) GetInputs() ConversationInputs`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *ConversationRequestBase) GetInputsOk() (*ConversationInputs, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *ConversationRequestBase) SetInputs(v ConversationInputs)`

SetInputs sets Inputs field to given value.


### GetStream

`func (o *ConversationRequestBase) GetStream() bool`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *ConversationRequestBase) GetStreamOk() (*bool, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *ConversationRequestBase) SetStream(v bool)`

SetStream sets Stream field to given value.

### HasStream

`func (o *ConversationRequestBase) HasStream() bool`

HasStream returns a boolean if a field has been set.

### SetStreamNil

`func (o *ConversationRequestBase) SetStreamNil(b bool)`

 SetStreamNil sets the value for Stream to be an explicit nil

### UnsetStream
`func (o *ConversationRequestBase) UnsetStream()`

UnsetStream ensures that no value is present for Stream, not even an explicit nil
### GetStore

`func (o *ConversationRequestBase) GetStore() bool`

GetStore returns the Store field if non-nil, zero value otherwise.

### GetStoreOk

`func (o *ConversationRequestBase) GetStoreOk() (*bool, bool)`

GetStoreOk returns a tuple with the Store field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStore

`func (o *ConversationRequestBase) SetStore(v bool)`

SetStore sets Store field to given value.

### HasStore

`func (o *ConversationRequestBase) HasStore() bool`

HasStore returns a boolean if a field has been set.

### SetStoreNil

`func (o *ConversationRequestBase) SetStoreNil(b bool)`

 SetStoreNil sets the value for Store to be an explicit nil

### UnsetStore
`func (o *ConversationRequestBase) UnsetStore()`

UnsetStore ensures that no value is present for Store, not even an explicit nil
### GetHandoffExecution

`func (o *ConversationRequestBase) GetHandoffExecution() string`

GetHandoffExecution returns the HandoffExecution field if non-nil, zero value otherwise.

### GetHandoffExecutionOk

`func (o *ConversationRequestBase) GetHandoffExecutionOk() (*string, bool)`

GetHandoffExecutionOk returns a tuple with the HandoffExecution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandoffExecution

`func (o *ConversationRequestBase) SetHandoffExecution(v string)`

SetHandoffExecution sets HandoffExecution field to given value.

### HasHandoffExecution

`func (o *ConversationRequestBase) HasHandoffExecution() bool`

HasHandoffExecution returns a boolean if a field has been set.

### SetHandoffExecutionNil

`func (o *ConversationRequestBase) SetHandoffExecutionNil(b bool)`

 SetHandoffExecutionNil sets the value for HandoffExecution to be an explicit nil

### UnsetHandoffExecution
`func (o *ConversationRequestBase) UnsetHandoffExecution()`

UnsetHandoffExecution ensures that no value is present for HandoffExecution, not even an explicit nil
### GetInstructions

`func (o *ConversationRequestBase) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *ConversationRequestBase) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *ConversationRequestBase) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *ConversationRequestBase) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### SetInstructionsNil

`func (o *ConversationRequestBase) SetInstructionsNil(b bool)`

 SetInstructionsNil sets the value for Instructions to be an explicit nil

### UnsetInstructions
`func (o *ConversationRequestBase) UnsetInstructions()`

UnsetInstructions ensures that no value is present for Instructions, not even an explicit nil
### GetTools

`func (o *ConversationRequestBase) GetTools() []ToolsInner`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *ConversationRequestBase) GetToolsOk() (*[]ToolsInner, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *ConversationRequestBase) SetTools(v []ToolsInner)`

SetTools sets Tools field to given value.

### HasTools

`func (o *ConversationRequestBase) HasTools() bool`

HasTools returns a boolean if a field has been set.

### SetToolsNil

`func (o *ConversationRequestBase) SetToolsNil(b bool)`

 SetToolsNil sets the value for Tools to be an explicit nil

### UnsetTools
`func (o *ConversationRequestBase) UnsetTools()`

UnsetTools ensures that no value is present for Tools, not even an explicit nil
### GetCompletionArgs

`func (o *ConversationRequestBase) GetCompletionArgs() CompletionArgs`

GetCompletionArgs returns the CompletionArgs field if non-nil, zero value otherwise.

### GetCompletionArgsOk

`func (o *ConversationRequestBase) GetCompletionArgsOk() (*CompletionArgs, bool)`

GetCompletionArgsOk returns a tuple with the CompletionArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionArgs

`func (o *ConversationRequestBase) SetCompletionArgs(v CompletionArgs)`

SetCompletionArgs sets CompletionArgs field to given value.

### HasCompletionArgs

`func (o *ConversationRequestBase) HasCompletionArgs() bool`

HasCompletionArgs returns a boolean if a field has been set.

### SetCompletionArgsNil

`func (o *ConversationRequestBase) SetCompletionArgsNil(b bool)`

 SetCompletionArgsNil sets the value for CompletionArgs to be an explicit nil

### UnsetCompletionArgs
`func (o *ConversationRequestBase) UnsetCompletionArgs()`

UnsetCompletionArgs ensures that no value is present for CompletionArgs, not even an explicit nil
### GetGuardrails

`func (o *ConversationRequestBase) GetGuardrails() []GuardrailConfig`

GetGuardrails returns the Guardrails field if non-nil, zero value otherwise.

### GetGuardrailsOk

`func (o *ConversationRequestBase) GetGuardrailsOk() (*[]GuardrailConfig, bool)`

GetGuardrailsOk returns a tuple with the Guardrails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuardrails

`func (o *ConversationRequestBase) SetGuardrails(v []GuardrailConfig)`

SetGuardrails sets Guardrails field to given value.

### HasGuardrails

`func (o *ConversationRequestBase) HasGuardrails() bool`

HasGuardrails returns a boolean if a field has been set.

### SetGuardrailsNil

`func (o *ConversationRequestBase) SetGuardrailsNil(b bool)`

 SetGuardrailsNil sets the value for Guardrails to be an explicit nil

### UnsetGuardrails
`func (o *ConversationRequestBase) UnsetGuardrails()`

UnsetGuardrails ensures that no value is present for Guardrails, not even an explicit nil
### GetName

`func (o *ConversationRequestBase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConversationRequestBase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConversationRequestBase) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConversationRequestBase) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ConversationRequestBase) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ConversationRequestBase) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *ConversationRequestBase) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConversationRequestBase) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConversationRequestBase) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConversationRequestBase) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ConversationRequestBase) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ConversationRequestBase) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetMetadata

`func (o *ConversationRequestBase) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ConversationRequestBase) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ConversationRequestBase) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ConversationRequestBase) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *ConversationRequestBase) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *ConversationRequestBase) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetAgentId

`func (o *ConversationRequestBase) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *ConversationRequestBase) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *ConversationRequestBase) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.

### HasAgentId

`func (o *ConversationRequestBase) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.

### SetAgentIdNil

`func (o *ConversationRequestBase) SetAgentIdNil(b bool)`

 SetAgentIdNil sets the value for AgentId to be an explicit nil

### UnsetAgentId
`func (o *ConversationRequestBase) UnsetAgentId()`

UnsetAgentId ensures that no value is present for AgentId, not even an explicit nil
### GetAgentVersion

`func (o *ConversationRequestBase) GetAgentVersion() AgentVersion1`

GetAgentVersion returns the AgentVersion field if non-nil, zero value otherwise.

### GetAgentVersionOk

`func (o *ConversationRequestBase) GetAgentVersionOk() (*AgentVersion1, bool)`

GetAgentVersionOk returns a tuple with the AgentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentVersion

`func (o *ConversationRequestBase) SetAgentVersion(v AgentVersion1)`

SetAgentVersion sets AgentVersion field to given value.

### HasAgentVersion

`func (o *ConversationRequestBase) HasAgentVersion() bool`

HasAgentVersion returns a boolean if a field has been set.

### SetAgentVersionNil

`func (o *ConversationRequestBase) SetAgentVersionNil(b bool)`

 SetAgentVersionNil sets the value for AgentVersion to be an explicit nil

### UnsetAgentVersion
`func (o *ConversationRequestBase) UnsetAgentVersion()`

UnsetAgentVersion ensures that no value is present for AgentVersion, not even an explicit nil
### GetModel

`func (o *ConversationRequestBase) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ConversationRequestBase) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ConversationRequestBase) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ConversationRequestBase) HasModel() bool`

HasModel returns a boolean if a field has been set.

### SetModelNil

`func (o *ConversationRequestBase) SetModelNil(b bool)`

 SetModelNil sets the value for Model to be an explicit nil

### UnsetModel
`func (o *ConversationRequestBase) UnsetModel()`

UnsetModel ensures that no value is present for Model, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


