# ConversationRestartRequestBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Inputs** | Pointer to [**ConversationInputs**](ConversationInputs.md) |  | [optional] 
**Stream** | Pointer to **bool** | Whether to stream back partial progress. Otherwise, the server will hold the request open until the timeout or until completion, with the response containing the full result as JSON. | [optional] [default to false]
**Store** | Pointer to **bool** | Whether to store the results into our servers or not. | [optional] [default to true]
**HandoffExecution** | Pointer to **string** |  | [optional] [default to "server"]
**CompletionArgs** | Pointer to [**CompletionArgs**](CompletionArgs.md) | Completion arguments that will be used to generate assistant responses. Can be overridden at each message request. | [optional] 
**Guardrails** | Pointer to [**[]GuardrailConfig**](GuardrailConfig.md) |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Custom metadata for the conversation. | [optional] 
**FromEntryId** | **string** |  | 
**AgentVersion** | Pointer to [**NullableAgentVersion1**](AgentVersion1.md) |  | [optional] 

## Methods

### NewConversationRestartRequestBase

`func NewConversationRestartRequestBase(fromEntryId string, ) *ConversationRestartRequestBase`

NewConversationRestartRequestBase instantiates a new ConversationRestartRequestBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConversationRestartRequestBaseWithDefaults

`func NewConversationRestartRequestBaseWithDefaults() *ConversationRestartRequestBase`

NewConversationRestartRequestBaseWithDefaults instantiates a new ConversationRestartRequestBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInputs

`func (o *ConversationRestartRequestBase) GetInputs() ConversationInputs`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *ConversationRestartRequestBase) GetInputsOk() (*ConversationInputs, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *ConversationRestartRequestBase) SetInputs(v ConversationInputs)`

SetInputs sets Inputs field to given value.

### HasInputs

`func (o *ConversationRestartRequestBase) HasInputs() bool`

HasInputs returns a boolean if a field has been set.

### GetStream

`func (o *ConversationRestartRequestBase) GetStream() bool`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *ConversationRestartRequestBase) GetStreamOk() (*bool, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *ConversationRestartRequestBase) SetStream(v bool)`

SetStream sets Stream field to given value.

### HasStream

`func (o *ConversationRestartRequestBase) HasStream() bool`

HasStream returns a boolean if a field has been set.

### GetStore

`func (o *ConversationRestartRequestBase) GetStore() bool`

GetStore returns the Store field if non-nil, zero value otherwise.

### GetStoreOk

`func (o *ConversationRestartRequestBase) GetStoreOk() (*bool, bool)`

GetStoreOk returns a tuple with the Store field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStore

`func (o *ConversationRestartRequestBase) SetStore(v bool)`

SetStore sets Store field to given value.

### HasStore

`func (o *ConversationRestartRequestBase) HasStore() bool`

HasStore returns a boolean if a field has been set.

### GetHandoffExecution

`func (o *ConversationRestartRequestBase) GetHandoffExecution() string`

GetHandoffExecution returns the HandoffExecution field if non-nil, zero value otherwise.

### GetHandoffExecutionOk

`func (o *ConversationRestartRequestBase) GetHandoffExecutionOk() (*string, bool)`

GetHandoffExecutionOk returns a tuple with the HandoffExecution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandoffExecution

`func (o *ConversationRestartRequestBase) SetHandoffExecution(v string)`

SetHandoffExecution sets HandoffExecution field to given value.

### HasHandoffExecution

`func (o *ConversationRestartRequestBase) HasHandoffExecution() bool`

HasHandoffExecution returns a boolean if a field has been set.

### GetCompletionArgs

`func (o *ConversationRestartRequestBase) GetCompletionArgs() CompletionArgs`

GetCompletionArgs returns the CompletionArgs field if non-nil, zero value otherwise.

### GetCompletionArgsOk

`func (o *ConversationRestartRequestBase) GetCompletionArgsOk() (*CompletionArgs, bool)`

GetCompletionArgsOk returns a tuple with the CompletionArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionArgs

`func (o *ConversationRestartRequestBase) SetCompletionArgs(v CompletionArgs)`

SetCompletionArgs sets CompletionArgs field to given value.

### HasCompletionArgs

`func (o *ConversationRestartRequestBase) HasCompletionArgs() bool`

HasCompletionArgs returns a boolean if a field has been set.

### GetGuardrails

`func (o *ConversationRestartRequestBase) GetGuardrails() []GuardrailConfig`

GetGuardrails returns the Guardrails field if non-nil, zero value otherwise.

### GetGuardrailsOk

`func (o *ConversationRestartRequestBase) GetGuardrailsOk() (*[]GuardrailConfig, bool)`

GetGuardrailsOk returns a tuple with the Guardrails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuardrails

`func (o *ConversationRestartRequestBase) SetGuardrails(v []GuardrailConfig)`

SetGuardrails sets Guardrails field to given value.

### HasGuardrails

`func (o *ConversationRestartRequestBase) HasGuardrails() bool`

HasGuardrails returns a boolean if a field has been set.

### SetGuardrailsNil

`func (o *ConversationRestartRequestBase) SetGuardrailsNil(b bool)`

 SetGuardrailsNil sets the value for Guardrails to be an explicit nil

### UnsetGuardrails
`func (o *ConversationRestartRequestBase) UnsetGuardrails()`

UnsetGuardrails ensures that no value is present for Guardrails, not even an explicit nil
### GetMetadata

`func (o *ConversationRestartRequestBase) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ConversationRestartRequestBase) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ConversationRestartRequestBase) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ConversationRestartRequestBase) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *ConversationRestartRequestBase) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *ConversationRestartRequestBase) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetFromEntryId

`func (o *ConversationRestartRequestBase) GetFromEntryId() string`

GetFromEntryId returns the FromEntryId field if non-nil, zero value otherwise.

### GetFromEntryIdOk

`func (o *ConversationRestartRequestBase) GetFromEntryIdOk() (*string, bool)`

GetFromEntryIdOk returns a tuple with the FromEntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromEntryId

`func (o *ConversationRestartRequestBase) SetFromEntryId(v string)`

SetFromEntryId sets FromEntryId field to given value.


### GetAgentVersion

`func (o *ConversationRestartRequestBase) GetAgentVersion() AgentVersion1`

GetAgentVersion returns the AgentVersion field if non-nil, zero value otherwise.

### GetAgentVersionOk

`func (o *ConversationRestartRequestBase) GetAgentVersionOk() (*AgentVersion1, bool)`

GetAgentVersionOk returns a tuple with the AgentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentVersion

`func (o *ConversationRestartRequestBase) SetAgentVersion(v AgentVersion1)`

SetAgentVersion sets AgentVersion field to given value.

### HasAgentVersion

`func (o *ConversationRestartRequestBase) HasAgentVersion() bool`

HasAgentVersion returns a boolean if a field has been set.

### SetAgentVersionNil

`func (o *ConversationRestartRequestBase) SetAgentVersionNil(b bool)`

 SetAgentVersionNil sets the value for AgentVersion to be an explicit nil

### UnsetAgentVersion
`func (o *ConversationRestartRequestBase) UnsetAgentVersion()`

UnsetAgentVersion ensures that no value is present for AgentVersion, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


