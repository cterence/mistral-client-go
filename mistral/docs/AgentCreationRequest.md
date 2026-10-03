# AgentCreationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instructions** | Pointer to **NullableString** | Instruction prompt the model will follow during the conversation. | [optional] 
**Tools** | Pointer to [**[]ToolsInner**](ToolsInner.md) | List of tools which are available to the model during the conversation. | [optional] 
**CompletionArgs** | Pointer to [**CompletionArgs**](CompletionArgs.md) | Completion arguments that will be used to generate assistant responses. Can be overridden at each message request. | [optional] 
**Guardrails** | Pointer to [**[]GuardrailConfig**](GuardrailConfig.md) |  | [optional] 
**Model** | **string** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Handoffs** | Pointer to **[]string** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Custom type for metadata with embedded validation. | [optional] 
**VersionMessage** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAgentCreationRequest

`func NewAgentCreationRequest(model string, name string, ) *AgentCreationRequest`

NewAgentCreationRequest instantiates a new AgentCreationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentCreationRequestWithDefaults

`func NewAgentCreationRequestWithDefaults() *AgentCreationRequest`

NewAgentCreationRequestWithDefaults instantiates a new AgentCreationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstructions

`func (o *AgentCreationRequest) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *AgentCreationRequest) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *AgentCreationRequest) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *AgentCreationRequest) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### SetInstructionsNil

`func (o *AgentCreationRequest) SetInstructionsNil(b bool)`

 SetInstructionsNil sets the value for Instructions to be an explicit nil

### UnsetInstructions
`func (o *AgentCreationRequest) UnsetInstructions()`

UnsetInstructions ensures that no value is present for Instructions, not even an explicit nil
### GetTools

`func (o *AgentCreationRequest) GetTools() []ToolsInner`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *AgentCreationRequest) GetToolsOk() (*[]ToolsInner, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *AgentCreationRequest) SetTools(v []ToolsInner)`

SetTools sets Tools field to given value.

### HasTools

`func (o *AgentCreationRequest) HasTools() bool`

HasTools returns a boolean if a field has been set.

### GetCompletionArgs

`func (o *AgentCreationRequest) GetCompletionArgs() CompletionArgs`

GetCompletionArgs returns the CompletionArgs field if non-nil, zero value otherwise.

### GetCompletionArgsOk

`func (o *AgentCreationRequest) GetCompletionArgsOk() (*CompletionArgs, bool)`

GetCompletionArgsOk returns a tuple with the CompletionArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionArgs

`func (o *AgentCreationRequest) SetCompletionArgs(v CompletionArgs)`

SetCompletionArgs sets CompletionArgs field to given value.

### HasCompletionArgs

`func (o *AgentCreationRequest) HasCompletionArgs() bool`

HasCompletionArgs returns a boolean if a field has been set.

### GetGuardrails

`func (o *AgentCreationRequest) GetGuardrails() []GuardrailConfig`

GetGuardrails returns the Guardrails field if non-nil, zero value otherwise.

### GetGuardrailsOk

`func (o *AgentCreationRequest) GetGuardrailsOk() (*[]GuardrailConfig, bool)`

GetGuardrailsOk returns a tuple with the Guardrails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuardrails

`func (o *AgentCreationRequest) SetGuardrails(v []GuardrailConfig)`

SetGuardrails sets Guardrails field to given value.

### HasGuardrails

`func (o *AgentCreationRequest) HasGuardrails() bool`

HasGuardrails returns a boolean if a field has been set.

### SetGuardrailsNil

`func (o *AgentCreationRequest) SetGuardrailsNil(b bool)`

 SetGuardrailsNil sets the value for Guardrails to be an explicit nil

### UnsetGuardrails
`func (o *AgentCreationRequest) UnsetGuardrails()`

UnsetGuardrails ensures that no value is present for Guardrails, not even an explicit nil
### GetModel

`func (o *AgentCreationRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AgentCreationRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AgentCreationRequest) SetModel(v string)`

SetModel sets Model field to given value.


### GetName

`func (o *AgentCreationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AgentCreationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AgentCreationRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AgentCreationRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AgentCreationRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AgentCreationRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AgentCreationRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AgentCreationRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AgentCreationRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetHandoffs

`func (o *AgentCreationRequest) GetHandoffs() []string`

GetHandoffs returns the Handoffs field if non-nil, zero value otherwise.

### GetHandoffsOk

`func (o *AgentCreationRequest) GetHandoffsOk() (*[]string, bool)`

GetHandoffsOk returns a tuple with the Handoffs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandoffs

`func (o *AgentCreationRequest) SetHandoffs(v []string)`

SetHandoffs sets Handoffs field to given value.

### HasHandoffs

`func (o *AgentCreationRequest) HasHandoffs() bool`

HasHandoffs returns a boolean if a field has been set.

### SetHandoffsNil

`func (o *AgentCreationRequest) SetHandoffsNil(b bool)`

 SetHandoffsNil sets the value for Handoffs to be an explicit nil

### UnsetHandoffs
`func (o *AgentCreationRequest) UnsetHandoffs()`

UnsetHandoffs ensures that no value is present for Handoffs, not even an explicit nil
### GetMetadata

`func (o *AgentCreationRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AgentCreationRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AgentCreationRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AgentCreationRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *AgentCreationRequest) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *AgentCreationRequest) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetVersionMessage

`func (o *AgentCreationRequest) GetVersionMessage() string`

GetVersionMessage returns the VersionMessage field if non-nil, zero value otherwise.

### GetVersionMessageOk

`func (o *AgentCreationRequest) GetVersionMessageOk() (*string, bool)`

GetVersionMessageOk returns a tuple with the VersionMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionMessage

`func (o *AgentCreationRequest) SetVersionMessage(v string)`

SetVersionMessage sets VersionMessage field to given value.

### HasVersionMessage

`func (o *AgentCreationRequest) HasVersionMessage() bool`

HasVersionMessage returns a boolean if a field has been set.

### SetVersionMessageNil

`func (o *AgentCreationRequest) SetVersionMessageNil(b bool)`

 SetVersionMessageNil sets the value for VersionMessage to be an explicit nil

### UnsetVersionMessage
`func (o *AgentCreationRequest) UnsetVersionMessage()`

UnsetVersionMessage ensures that no value is present for VersionMessage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


