# Agent

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
**Object** | Pointer to **string** |  | [optional] [default to "agent"]
**Id** | **string** |  | 
**Version** | **int32** |  | 
**Versions** | **[]int32** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**DeploymentChat** | **bool** |  | 
**Source** | **string** |  | 
**VersionMessage** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAgent

`func NewAgent(model string, name string, id string, version int32, versions []int32, createdAt time.Time, updatedAt time.Time, deploymentChat bool, source string, ) *Agent`

NewAgent instantiates a new Agent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentWithDefaults

`func NewAgentWithDefaults() *Agent`

NewAgentWithDefaults instantiates a new Agent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstructions

`func (o *Agent) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *Agent) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *Agent) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *Agent) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### SetInstructionsNil

`func (o *Agent) SetInstructionsNil(b bool)`

 SetInstructionsNil sets the value for Instructions to be an explicit nil

### UnsetInstructions
`func (o *Agent) UnsetInstructions()`

UnsetInstructions ensures that no value is present for Instructions, not even an explicit nil
### GetTools

`func (o *Agent) GetTools() []ToolsInner`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *Agent) GetToolsOk() (*[]ToolsInner, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *Agent) SetTools(v []ToolsInner)`

SetTools sets Tools field to given value.

### HasTools

`func (o *Agent) HasTools() bool`

HasTools returns a boolean if a field has been set.

### GetCompletionArgs

`func (o *Agent) GetCompletionArgs() CompletionArgs`

GetCompletionArgs returns the CompletionArgs field if non-nil, zero value otherwise.

### GetCompletionArgsOk

`func (o *Agent) GetCompletionArgsOk() (*CompletionArgs, bool)`

GetCompletionArgsOk returns a tuple with the CompletionArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionArgs

`func (o *Agent) SetCompletionArgs(v CompletionArgs)`

SetCompletionArgs sets CompletionArgs field to given value.

### HasCompletionArgs

`func (o *Agent) HasCompletionArgs() bool`

HasCompletionArgs returns a boolean if a field has been set.

### GetGuardrails

`func (o *Agent) GetGuardrails() []GuardrailConfig`

GetGuardrails returns the Guardrails field if non-nil, zero value otherwise.

### GetGuardrailsOk

`func (o *Agent) GetGuardrailsOk() (*[]GuardrailConfig, bool)`

GetGuardrailsOk returns a tuple with the Guardrails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuardrails

`func (o *Agent) SetGuardrails(v []GuardrailConfig)`

SetGuardrails sets Guardrails field to given value.

### HasGuardrails

`func (o *Agent) HasGuardrails() bool`

HasGuardrails returns a boolean if a field has been set.

### SetGuardrailsNil

`func (o *Agent) SetGuardrailsNil(b bool)`

 SetGuardrailsNil sets the value for Guardrails to be an explicit nil

### UnsetGuardrails
`func (o *Agent) UnsetGuardrails()`

UnsetGuardrails ensures that no value is present for Guardrails, not even an explicit nil
### GetModel

`func (o *Agent) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *Agent) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *Agent) SetModel(v string)`

SetModel sets Model field to given value.


### GetName

`func (o *Agent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Agent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Agent) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Agent) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Agent) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Agent) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Agent) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Agent) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Agent) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetHandoffs

`func (o *Agent) GetHandoffs() []string`

GetHandoffs returns the Handoffs field if non-nil, zero value otherwise.

### GetHandoffsOk

`func (o *Agent) GetHandoffsOk() (*[]string, bool)`

GetHandoffsOk returns a tuple with the Handoffs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandoffs

`func (o *Agent) SetHandoffs(v []string)`

SetHandoffs sets Handoffs field to given value.

### HasHandoffs

`func (o *Agent) HasHandoffs() bool`

HasHandoffs returns a boolean if a field has been set.

### SetHandoffsNil

`func (o *Agent) SetHandoffsNil(b bool)`

 SetHandoffsNil sets the value for Handoffs to be an explicit nil

### UnsetHandoffs
`func (o *Agent) UnsetHandoffs()`

UnsetHandoffs ensures that no value is present for Handoffs, not even an explicit nil
### GetMetadata

`func (o *Agent) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Agent) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Agent) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Agent) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *Agent) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *Agent) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetObject

`func (o *Agent) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *Agent) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *Agent) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *Agent) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetId

`func (o *Agent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Agent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Agent) SetId(v string)`

SetId sets Id field to given value.


### GetVersion

`func (o *Agent) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Agent) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Agent) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetVersions

`func (o *Agent) GetVersions() []int32`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *Agent) GetVersionsOk() (*[]int32, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *Agent) SetVersions(v []int32)`

SetVersions sets Versions field to given value.


### GetCreatedAt

`func (o *Agent) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Agent) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Agent) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Agent) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Agent) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Agent) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeploymentChat

`func (o *Agent) GetDeploymentChat() bool`

GetDeploymentChat returns the DeploymentChat field if non-nil, zero value otherwise.

### GetDeploymentChatOk

`func (o *Agent) GetDeploymentChatOk() (*bool, bool)`

GetDeploymentChatOk returns a tuple with the DeploymentChat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentChat

`func (o *Agent) SetDeploymentChat(v bool)`

SetDeploymentChat sets DeploymentChat field to given value.


### GetSource

`func (o *Agent) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Agent) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Agent) SetSource(v string)`

SetSource sets Source field to given value.


### GetVersionMessage

`func (o *Agent) GetVersionMessage() string`

GetVersionMessage returns the VersionMessage field if non-nil, zero value otherwise.

### GetVersionMessageOk

`func (o *Agent) GetVersionMessageOk() (*string, bool)`

GetVersionMessageOk returns a tuple with the VersionMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionMessage

`func (o *Agent) SetVersionMessage(v string)`

SetVersionMessage sets VersionMessage field to given value.

### HasVersionMessage

`func (o *Agent) HasVersionMessage() bool`

HasVersionMessage returns a boolean if a field has been set.

### SetVersionMessageNil

`func (o *Agent) SetVersionMessageNil(b bool)`

 SetVersionMessageNil sets the value for VersionMessage to be an explicit nil

### UnsetVersionMessage
`func (o *Agent) UnsetVersionMessage()`

UnsetVersionMessage ensures that no value is present for VersionMessage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


