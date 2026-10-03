# AgentUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instructions** | Pointer to **NullableString** | Instruction prompt the model will follow during the conversation. | [optional] 
**Tools** | Pointer to [**[]ToolsInner**](ToolsInner.md) | List of tools which are available to the model during the conversation. | [optional] 
**CompletionArgs** | Pointer to [**CompletionArgs**](CompletionArgs.md) | Completion arguments that will be used to generate assistant responses. Can be overridden at each message request. | [optional] 
**Guardrails** | Pointer to [**[]GuardrailConfig**](GuardrailConfig.md) |  | [optional] 
**Model** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Handoffs** | Pointer to **[]string** |  | [optional] 
**DeploymentChat** | Pointer to **NullableBool** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Custom type for metadata with embedded validation. | [optional] 
**VersionMessage** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAgentUpdateRequest

`func NewAgentUpdateRequest() *AgentUpdateRequest`

NewAgentUpdateRequest instantiates a new AgentUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentUpdateRequestWithDefaults

`func NewAgentUpdateRequestWithDefaults() *AgentUpdateRequest`

NewAgentUpdateRequestWithDefaults instantiates a new AgentUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstructions

`func (o *AgentUpdateRequest) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *AgentUpdateRequest) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *AgentUpdateRequest) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *AgentUpdateRequest) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### SetInstructionsNil

`func (o *AgentUpdateRequest) SetInstructionsNil(b bool)`

 SetInstructionsNil sets the value for Instructions to be an explicit nil

### UnsetInstructions
`func (o *AgentUpdateRequest) UnsetInstructions()`

UnsetInstructions ensures that no value is present for Instructions, not even an explicit nil
### GetTools

`func (o *AgentUpdateRequest) GetTools() []ToolsInner`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *AgentUpdateRequest) GetToolsOk() (*[]ToolsInner, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *AgentUpdateRequest) SetTools(v []ToolsInner)`

SetTools sets Tools field to given value.

### HasTools

`func (o *AgentUpdateRequest) HasTools() bool`

HasTools returns a boolean if a field has been set.

### GetCompletionArgs

`func (o *AgentUpdateRequest) GetCompletionArgs() CompletionArgs`

GetCompletionArgs returns the CompletionArgs field if non-nil, zero value otherwise.

### GetCompletionArgsOk

`func (o *AgentUpdateRequest) GetCompletionArgsOk() (*CompletionArgs, bool)`

GetCompletionArgsOk returns a tuple with the CompletionArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionArgs

`func (o *AgentUpdateRequest) SetCompletionArgs(v CompletionArgs)`

SetCompletionArgs sets CompletionArgs field to given value.

### HasCompletionArgs

`func (o *AgentUpdateRequest) HasCompletionArgs() bool`

HasCompletionArgs returns a boolean if a field has been set.

### GetGuardrails

`func (o *AgentUpdateRequest) GetGuardrails() []GuardrailConfig`

GetGuardrails returns the Guardrails field if non-nil, zero value otherwise.

### GetGuardrailsOk

`func (o *AgentUpdateRequest) GetGuardrailsOk() (*[]GuardrailConfig, bool)`

GetGuardrailsOk returns a tuple with the Guardrails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuardrails

`func (o *AgentUpdateRequest) SetGuardrails(v []GuardrailConfig)`

SetGuardrails sets Guardrails field to given value.

### HasGuardrails

`func (o *AgentUpdateRequest) HasGuardrails() bool`

HasGuardrails returns a boolean if a field has been set.

### SetGuardrailsNil

`func (o *AgentUpdateRequest) SetGuardrailsNil(b bool)`

 SetGuardrailsNil sets the value for Guardrails to be an explicit nil

### UnsetGuardrails
`func (o *AgentUpdateRequest) UnsetGuardrails()`

UnsetGuardrails ensures that no value is present for Guardrails, not even an explicit nil
### GetModel

`func (o *AgentUpdateRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AgentUpdateRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AgentUpdateRequest) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *AgentUpdateRequest) HasModel() bool`

HasModel returns a boolean if a field has been set.

### SetModelNil

`func (o *AgentUpdateRequest) SetModelNil(b bool)`

 SetModelNil sets the value for Model to be an explicit nil

### UnsetModel
`func (o *AgentUpdateRequest) UnsetModel()`

UnsetModel ensures that no value is present for Model, not even an explicit nil
### GetName

`func (o *AgentUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AgentUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AgentUpdateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AgentUpdateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AgentUpdateRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AgentUpdateRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *AgentUpdateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AgentUpdateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AgentUpdateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AgentUpdateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AgentUpdateRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AgentUpdateRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetHandoffs

`func (o *AgentUpdateRequest) GetHandoffs() []string`

GetHandoffs returns the Handoffs field if non-nil, zero value otherwise.

### GetHandoffsOk

`func (o *AgentUpdateRequest) GetHandoffsOk() (*[]string, bool)`

GetHandoffsOk returns a tuple with the Handoffs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandoffs

`func (o *AgentUpdateRequest) SetHandoffs(v []string)`

SetHandoffs sets Handoffs field to given value.

### HasHandoffs

`func (o *AgentUpdateRequest) HasHandoffs() bool`

HasHandoffs returns a boolean if a field has been set.

### SetHandoffsNil

`func (o *AgentUpdateRequest) SetHandoffsNil(b bool)`

 SetHandoffsNil sets the value for Handoffs to be an explicit nil

### UnsetHandoffs
`func (o *AgentUpdateRequest) UnsetHandoffs()`

UnsetHandoffs ensures that no value is present for Handoffs, not even an explicit nil
### GetDeploymentChat

`func (o *AgentUpdateRequest) GetDeploymentChat() bool`

GetDeploymentChat returns the DeploymentChat field if non-nil, zero value otherwise.

### GetDeploymentChatOk

`func (o *AgentUpdateRequest) GetDeploymentChatOk() (*bool, bool)`

GetDeploymentChatOk returns a tuple with the DeploymentChat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentChat

`func (o *AgentUpdateRequest) SetDeploymentChat(v bool)`

SetDeploymentChat sets DeploymentChat field to given value.

### HasDeploymentChat

`func (o *AgentUpdateRequest) HasDeploymentChat() bool`

HasDeploymentChat returns a boolean if a field has been set.

### SetDeploymentChatNil

`func (o *AgentUpdateRequest) SetDeploymentChatNil(b bool)`

 SetDeploymentChatNil sets the value for DeploymentChat to be an explicit nil

### UnsetDeploymentChat
`func (o *AgentUpdateRequest) UnsetDeploymentChat()`

UnsetDeploymentChat ensures that no value is present for DeploymentChat, not even an explicit nil
### GetMetadata

`func (o *AgentUpdateRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AgentUpdateRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AgentUpdateRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AgentUpdateRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *AgentUpdateRequest) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *AgentUpdateRequest) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetVersionMessage

`func (o *AgentUpdateRequest) GetVersionMessage() string`

GetVersionMessage returns the VersionMessage field if non-nil, zero value otherwise.

### GetVersionMessageOk

`func (o *AgentUpdateRequest) GetVersionMessageOk() (*string, bool)`

GetVersionMessageOk returns a tuple with the VersionMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionMessage

`func (o *AgentUpdateRequest) SetVersionMessage(v string)`

SetVersionMessage sets VersionMessage field to given value.

### HasVersionMessage

`func (o *AgentUpdateRequest) HasVersionMessage() bool`

HasVersionMessage returns a boolean if a field has been set.

### SetVersionMessageNil

`func (o *AgentUpdateRequest) SetVersionMessageNil(b bool)`

 SetVersionMessageNil sets the value for VersionMessage to be an explicit nil

### UnsetVersionMessage
`func (o *AgentUpdateRequest) UnsetVersionMessage()`

UnsetVersionMessage ensures that no value is present for VersionMessage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


