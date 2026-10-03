# ConversationStreamRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Inputs** | [**ConversationInputs**](ConversationInputs.md) |  | 
**Stream** | Pointer to **bool** |  | [optional] [default to true]
**Store** | Pointer to **bool** |  | [optional] 
**HandoffExecution** | Pointer to **string** |  | [optional] 
**Instructions** | Pointer to **string** |  | [optional] 
**Tools** | Pointer to [**[]ToolsInner**](ToolsInner.md) |  | [optional] 
**CompletionArgs** | Pointer to [**CompletionArgs**](CompletionArgs.md) |  | [optional] 
**Guardrails** | Pointer to [**[]GuardrailConfig**](GuardrailConfig.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Custom type for metadata with embedded validation. | [optional] 
**AgentId** | Pointer to **string** |  | [optional] 
**AgentVersion** | Pointer to [**NullableAgentVersion1**](AgentVersion1.md) |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 

## Methods

### NewConversationStreamRequest

`func NewConversationStreamRequest(inputs ConversationInputs, ) *ConversationStreamRequest`

NewConversationStreamRequest instantiates a new ConversationStreamRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConversationStreamRequestWithDefaults

`func NewConversationStreamRequestWithDefaults() *ConversationStreamRequest`

NewConversationStreamRequestWithDefaults instantiates a new ConversationStreamRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInputs

`func (o *ConversationStreamRequest) GetInputs() ConversationInputs`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *ConversationStreamRequest) GetInputsOk() (*ConversationInputs, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *ConversationStreamRequest) SetInputs(v ConversationInputs)`

SetInputs sets Inputs field to given value.


### GetStream

`func (o *ConversationStreamRequest) GetStream() bool`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *ConversationStreamRequest) GetStreamOk() (*bool, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *ConversationStreamRequest) SetStream(v bool)`

SetStream sets Stream field to given value.

### HasStream

`func (o *ConversationStreamRequest) HasStream() bool`

HasStream returns a boolean if a field has been set.

### GetStore

`func (o *ConversationStreamRequest) GetStore() bool`

GetStore returns the Store field if non-nil, zero value otherwise.

### GetStoreOk

`func (o *ConversationStreamRequest) GetStoreOk() (*bool, bool)`

GetStoreOk returns a tuple with the Store field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStore

`func (o *ConversationStreamRequest) SetStore(v bool)`

SetStore sets Store field to given value.

### HasStore

`func (o *ConversationStreamRequest) HasStore() bool`

HasStore returns a boolean if a field has been set.

### GetHandoffExecution

`func (o *ConversationStreamRequest) GetHandoffExecution() string`

GetHandoffExecution returns the HandoffExecution field if non-nil, zero value otherwise.

### GetHandoffExecutionOk

`func (o *ConversationStreamRequest) GetHandoffExecutionOk() (*string, bool)`

GetHandoffExecutionOk returns a tuple with the HandoffExecution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandoffExecution

`func (o *ConversationStreamRequest) SetHandoffExecution(v string)`

SetHandoffExecution sets HandoffExecution field to given value.

### HasHandoffExecution

`func (o *ConversationStreamRequest) HasHandoffExecution() bool`

HasHandoffExecution returns a boolean if a field has been set.

### GetInstructions

`func (o *ConversationStreamRequest) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *ConversationStreamRequest) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *ConversationStreamRequest) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *ConversationStreamRequest) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### GetTools

`func (o *ConversationStreamRequest) GetTools() []ToolsInner`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *ConversationStreamRequest) GetToolsOk() (*[]ToolsInner, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *ConversationStreamRequest) SetTools(v []ToolsInner)`

SetTools sets Tools field to given value.

### HasTools

`func (o *ConversationStreamRequest) HasTools() bool`

HasTools returns a boolean if a field has been set.

### GetCompletionArgs

`func (o *ConversationStreamRequest) GetCompletionArgs() CompletionArgs`

GetCompletionArgs returns the CompletionArgs field if non-nil, zero value otherwise.

### GetCompletionArgsOk

`func (o *ConversationStreamRequest) GetCompletionArgsOk() (*CompletionArgs, bool)`

GetCompletionArgsOk returns a tuple with the CompletionArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionArgs

`func (o *ConversationStreamRequest) SetCompletionArgs(v CompletionArgs)`

SetCompletionArgs sets CompletionArgs field to given value.

### HasCompletionArgs

`func (o *ConversationStreamRequest) HasCompletionArgs() bool`

HasCompletionArgs returns a boolean if a field has been set.

### GetGuardrails

`func (o *ConversationStreamRequest) GetGuardrails() []GuardrailConfig`

GetGuardrails returns the Guardrails field if non-nil, zero value otherwise.

### GetGuardrailsOk

`func (o *ConversationStreamRequest) GetGuardrailsOk() (*[]GuardrailConfig, bool)`

GetGuardrailsOk returns a tuple with the Guardrails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuardrails

`func (o *ConversationStreamRequest) SetGuardrails(v []GuardrailConfig)`

SetGuardrails sets Guardrails field to given value.

### HasGuardrails

`func (o *ConversationStreamRequest) HasGuardrails() bool`

HasGuardrails returns a boolean if a field has been set.

### GetName

`func (o *ConversationStreamRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConversationStreamRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConversationStreamRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConversationStreamRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ConversationStreamRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConversationStreamRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConversationStreamRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConversationStreamRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMetadata

`func (o *ConversationStreamRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ConversationStreamRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ConversationStreamRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ConversationStreamRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetAgentId

`func (o *ConversationStreamRequest) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *ConversationStreamRequest) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *ConversationStreamRequest) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.

### HasAgentId

`func (o *ConversationStreamRequest) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.

### GetAgentVersion

`func (o *ConversationStreamRequest) GetAgentVersion() AgentVersion1`

GetAgentVersion returns the AgentVersion field if non-nil, zero value otherwise.

### GetAgentVersionOk

`func (o *ConversationStreamRequest) GetAgentVersionOk() (*AgentVersion1, bool)`

GetAgentVersionOk returns a tuple with the AgentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentVersion

`func (o *ConversationStreamRequest) SetAgentVersion(v AgentVersion1)`

SetAgentVersion sets AgentVersion field to given value.

### HasAgentVersion

`func (o *ConversationStreamRequest) HasAgentVersion() bool`

HasAgentVersion returns a boolean if a field has been set.

### SetAgentVersionNil

`func (o *ConversationStreamRequest) SetAgentVersionNil(b bool)`

 SetAgentVersionNil sets the value for AgentVersion to be an explicit nil

### UnsetAgentVersion
`func (o *ConversationStreamRequest) UnsetAgentVersion()`

UnsetAgentVersion ensures that no value is present for AgentVersion, not even an explicit nil
### GetModel

`func (o *ConversationStreamRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ConversationStreamRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ConversationStreamRequest) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ConversationStreamRequest) HasModel() bool`

HasModel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


