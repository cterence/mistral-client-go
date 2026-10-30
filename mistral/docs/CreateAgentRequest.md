# CreateAgentRequest

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

### NewCreateAgentRequest

`func NewCreateAgentRequest(model string, name string, ) *CreateAgentRequest`

NewCreateAgentRequest instantiates a new CreateAgentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAgentRequestWithDefaults

`func NewCreateAgentRequestWithDefaults() *CreateAgentRequest`

NewCreateAgentRequestWithDefaults instantiates a new CreateAgentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstructions

`func (o *CreateAgentRequest) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *CreateAgentRequest) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *CreateAgentRequest) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *CreateAgentRequest) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### SetInstructionsNil

`func (o *CreateAgentRequest) SetInstructionsNil(b bool)`

 SetInstructionsNil sets the value for Instructions to be an explicit nil

### UnsetInstructions
`func (o *CreateAgentRequest) UnsetInstructions()`

UnsetInstructions ensures that no value is present for Instructions, not even an explicit nil
### GetTools

`func (o *CreateAgentRequest) GetTools() []ToolsInner`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *CreateAgentRequest) GetToolsOk() (*[]ToolsInner, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *CreateAgentRequest) SetTools(v []ToolsInner)`

SetTools sets Tools field to given value.

### HasTools

`func (o *CreateAgentRequest) HasTools() bool`

HasTools returns a boolean if a field has been set.

### GetCompletionArgs

`func (o *CreateAgentRequest) GetCompletionArgs() CompletionArgs`

GetCompletionArgs returns the CompletionArgs field if non-nil, zero value otherwise.

### GetCompletionArgsOk

`func (o *CreateAgentRequest) GetCompletionArgsOk() (*CompletionArgs, bool)`

GetCompletionArgsOk returns a tuple with the CompletionArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionArgs

`func (o *CreateAgentRequest) SetCompletionArgs(v CompletionArgs)`

SetCompletionArgs sets CompletionArgs field to given value.

### HasCompletionArgs

`func (o *CreateAgentRequest) HasCompletionArgs() bool`

HasCompletionArgs returns a boolean if a field has been set.

### GetGuardrails

`func (o *CreateAgentRequest) GetGuardrails() []GuardrailConfig`

GetGuardrails returns the Guardrails field if non-nil, zero value otherwise.

### GetGuardrailsOk

`func (o *CreateAgentRequest) GetGuardrailsOk() (*[]GuardrailConfig, bool)`

GetGuardrailsOk returns a tuple with the Guardrails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuardrails

`func (o *CreateAgentRequest) SetGuardrails(v []GuardrailConfig)`

SetGuardrails sets Guardrails field to given value.

### HasGuardrails

`func (o *CreateAgentRequest) HasGuardrails() bool`

HasGuardrails returns a boolean if a field has been set.

### SetGuardrailsNil

`func (o *CreateAgentRequest) SetGuardrailsNil(b bool)`

 SetGuardrailsNil sets the value for Guardrails to be an explicit nil

### UnsetGuardrails
`func (o *CreateAgentRequest) UnsetGuardrails()`

UnsetGuardrails ensures that no value is present for Guardrails, not even an explicit nil
### GetModel

`func (o *CreateAgentRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *CreateAgentRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *CreateAgentRequest) SetModel(v string)`

SetModel sets Model field to given value.


### GetName

`func (o *CreateAgentRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAgentRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAgentRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateAgentRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateAgentRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateAgentRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateAgentRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateAgentRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateAgentRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetHandoffs

`func (o *CreateAgentRequest) GetHandoffs() []string`

GetHandoffs returns the Handoffs field if non-nil, zero value otherwise.

### GetHandoffsOk

`func (o *CreateAgentRequest) GetHandoffsOk() (*[]string, bool)`

GetHandoffsOk returns a tuple with the Handoffs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandoffs

`func (o *CreateAgentRequest) SetHandoffs(v []string)`

SetHandoffs sets Handoffs field to given value.

### HasHandoffs

`func (o *CreateAgentRequest) HasHandoffs() bool`

HasHandoffs returns a boolean if a field has been set.

### SetHandoffsNil

`func (o *CreateAgentRequest) SetHandoffsNil(b bool)`

 SetHandoffsNil sets the value for Handoffs to be an explicit nil

### UnsetHandoffs
`func (o *CreateAgentRequest) UnsetHandoffs()`

UnsetHandoffs ensures that no value is present for Handoffs, not even an explicit nil
### GetMetadata

`func (o *CreateAgentRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateAgentRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateAgentRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateAgentRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *CreateAgentRequest) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *CreateAgentRequest) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetVersionMessage

`func (o *CreateAgentRequest) GetVersionMessage() string`

GetVersionMessage returns the VersionMessage field if non-nil, zero value otherwise.

### GetVersionMessageOk

`func (o *CreateAgentRequest) GetVersionMessageOk() (*string, bool)`

GetVersionMessageOk returns a tuple with the VersionMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionMessage

`func (o *CreateAgentRequest) SetVersionMessage(v string)`

SetVersionMessage sets VersionMessage field to given value.

### HasVersionMessage

`func (o *CreateAgentRequest) HasVersionMessage() bool`

HasVersionMessage returns a boolean if a field has been set.

### SetVersionMessageNil

`func (o *CreateAgentRequest) SetVersionMessageNil(b bool)`

 SetVersionMessageNil sets the value for VersionMessage to be an explicit nil

### UnsetVersionMessage
`func (o *CreateAgentRequest) UnsetVersionMessage()`

UnsetVersionMessage ensures that no value is present for VersionMessage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


