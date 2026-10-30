# UpdateAgentRequest

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

### NewUpdateAgentRequest

`func NewUpdateAgentRequest() *UpdateAgentRequest`

NewUpdateAgentRequest instantiates a new UpdateAgentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAgentRequestWithDefaults

`func NewUpdateAgentRequestWithDefaults() *UpdateAgentRequest`

NewUpdateAgentRequestWithDefaults instantiates a new UpdateAgentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstructions

`func (o *UpdateAgentRequest) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *UpdateAgentRequest) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *UpdateAgentRequest) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *UpdateAgentRequest) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### SetInstructionsNil

`func (o *UpdateAgentRequest) SetInstructionsNil(b bool)`

 SetInstructionsNil sets the value for Instructions to be an explicit nil

### UnsetInstructions
`func (o *UpdateAgentRequest) UnsetInstructions()`

UnsetInstructions ensures that no value is present for Instructions, not even an explicit nil
### GetTools

`func (o *UpdateAgentRequest) GetTools() []ToolsInner`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *UpdateAgentRequest) GetToolsOk() (*[]ToolsInner, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *UpdateAgentRequest) SetTools(v []ToolsInner)`

SetTools sets Tools field to given value.

### HasTools

`func (o *UpdateAgentRequest) HasTools() bool`

HasTools returns a boolean if a field has been set.

### GetCompletionArgs

`func (o *UpdateAgentRequest) GetCompletionArgs() CompletionArgs`

GetCompletionArgs returns the CompletionArgs field if non-nil, zero value otherwise.

### GetCompletionArgsOk

`func (o *UpdateAgentRequest) GetCompletionArgsOk() (*CompletionArgs, bool)`

GetCompletionArgsOk returns a tuple with the CompletionArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionArgs

`func (o *UpdateAgentRequest) SetCompletionArgs(v CompletionArgs)`

SetCompletionArgs sets CompletionArgs field to given value.

### HasCompletionArgs

`func (o *UpdateAgentRequest) HasCompletionArgs() bool`

HasCompletionArgs returns a boolean if a field has been set.

### GetGuardrails

`func (o *UpdateAgentRequest) GetGuardrails() []GuardrailConfig`

GetGuardrails returns the Guardrails field if non-nil, zero value otherwise.

### GetGuardrailsOk

`func (o *UpdateAgentRequest) GetGuardrailsOk() (*[]GuardrailConfig, bool)`

GetGuardrailsOk returns a tuple with the Guardrails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuardrails

`func (o *UpdateAgentRequest) SetGuardrails(v []GuardrailConfig)`

SetGuardrails sets Guardrails field to given value.

### HasGuardrails

`func (o *UpdateAgentRequest) HasGuardrails() bool`

HasGuardrails returns a boolean if a field has been set.

### SetGuardrailsNil

`func (o *UpdateAgentRequest) SetGuardrailsNil(b bool)`

 SetGuardrailsNil sets the value for Guardrails to be an explicit nil

### UnsetGuardrails
`func (o *UpdateAgentRequest) UnsetGuardrails()`

UnsetGuardrails ensures that no value is present for Guardrails, not even an explicit nil
### GetModel

`func (o *UpdateAgentRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *UpdateAgentRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *UpdateAgentRequest) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *UpdateAgentRequest) HasModel() bool`

HasModel returns a boolean if a field has been set.

### SetModelNil

`func (o *UpdateAgentRequest) SetModelNil(b bool)`

 SetModelNil sets the value for Model to be an explicit nil

### UnsetModel
`func (o *UpdateAgentRequest) UnsetModel()`

UnsetModel ensures that no value is present for Model, not even an explicit nil
### GetName

`func (o *UpdateAgentRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateAgentRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateAgentRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateAgentRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpdateAgentRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateAgentRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *UpdateAgentRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateAgentRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateAgentRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateAgentRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateAgentRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateAgentRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetHandoffs

`func (o *UpdateAgentRequest) GetHandoffs() []string`

GetHandoffs returns the Handoffs field if non-nil, zero value otherwise.

### GetHandoffsOk

`func (o *UpdateAgentRequest) GetHandoffsOk() (*[]string, bool)`

GetHandoffsOk returns a tuple with the Handoffs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandoffs

`func (o *UpdateAgentRequest) SetHandoffs(v []string)`

SetHandoffs sets Handoffs field to given value.

### HasHandoffs

`func (o *UpdateAgentRequest) HasHandoffs() bool`

HasHandoffs returns a boolean if a field has been set.

### SetHandoffsNil

`func (o *UpdateAgentRequest) SetHandoffsNil(b bool)`

 SetHandoffsNil sets the value for Handoffs to be an explicit nil

### UnsetHandoffs
`func (o *UpdateAgentRequest) UnsetHandoffs()`

UnsetHandoffs ensures that no value is present for Handoffs, not even an explicit nil
### GetDeploymentChat

`func (o *UpdateAgentRequest) GetDeploymentChat() bool`

GetDeploymentChat returns the DeploymentChat field if non-nil, zero value otherwise.

### GetDeploymentChatOk

`func (o *UpdateAgentRequest) GetDeploymentChatOk() (*bool, bool)`

GetDeploymentChatOk returns a tuple with the DeploymentChat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentChat

`func (o *UpdateAgentRequest) SetDeploymentChat(v bool)`

SetDeploymentChat sets DeploymentChat field to given value.

### HasDeploymentChat

`func (o *UpdateAgentRequest) HasDeploymentChat() bool`

HasDeploymentChat returns a boolean if a field has been set.

### SetDeploymentChatNil

`func (o *UpdateAgentRequest) SetDeploymentChatNil(b bool)`

 SetDeploymentChatNil sets the value for DeploymentChat to be an explicit nil

### UnsetDeploymentChat
`func (o *UpdateAgentRequest) UnsetDeploymentChat()`

UnsetDeploymentChat ensures that no value is present for DeploymentChat, not even an explicit nil
### GetMetadata

`func (o *UpdateAgentRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UpdateAgentRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UpdateAgentRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UpdateAgentRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *UpdateAgentRequest) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *UpdateAgentRequest) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetVersionMessage

`func (o *UpdateAgentRequest) GetVersionMessage() string`

GetVersionMessage returns the VersionMessage field if non-nil, zero value otherwise.

### GetVersionMessageOk

`func (o *UpdateAgentRequest) GetVersionMessageOk() (*string, bool)`

GetVersionMessageOk returns a tuple with the VersionMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionMessage

`func (o *UpdateAgentRequest) SetVersionMessage(v string)`

SetVersionMessage sets VersionMessage field to given value.

### HasVersionMessage

`func (o *UpdateAgentRequest) HasVersionMessage() bool`

HasVersionMessage returns a boolean if a field has been set.

### SetVersionMessageNil

`func (o *UpdateAgentRequest) SetVersionMessageNil(b bool)`

 SetVersionMessageNil sets the value for VersionMessage to be an explicit nil

### UnsetVersionMessage
`func (o *UpdateAgentRequest) UnsetVersionMessage()`

UnsetVersionMessage ensures that no value is present for VersionMessage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


