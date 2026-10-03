# ConversationRestartStreamRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Inputs** | Pointer to [**ConversationInputs**](ConversationInputs.md) |  | [optional] 
**Stream** | Pointer to **bool** |  | [optional] [default to true]
**Store** | Pointer to **bool** | Whether to store the results into our servers or not. | [optional] [default to true]
**HandoffExecution** | Pointer to **string** |  | [optional] [default to "server"]
**CompletionArgs** | Pointer to [**CompletionArgs**](CompletionArgs.md) | Completion arguments that will be used to generate assistant responses. Can be overridden at each message request. | [optional] 
**Guardrails** | Pointer to [**[]GuardrailConfig**](GuardrailConfig.md) |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Custom metadata for the conversation. | [optional] 
**FromEntryId** | **string** |  | 
**AgentVersion** | Pointer to [**NullableAgentVersion1**](AgentVersion1.md) |  | [optional] 

## Methods

### NewConversationRestartStreamRequest

`func NewConversationRestartStreamRequest(fromEntryId string, ) *ConversationRestartStreamRequest`

NewConversationRestartStreamRequest instantiates a new ConversationRestartStreamRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConversationRestartStreamRequestWithDefaults

`func NewConversationRestartStreamRequestWithDefaults() *ConversationRestartStreamRequest`

NewConversationRestartStreamRequestWithDefaults instantiates a new ConversationRestartStreamRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInputs

`func (o *ConversationRestartStreamRequest) GetInputs() ConversationInputs`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *ConversationRestartStreamRequest) GetInputsOk() (*ConversationInputs, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *ConversationRestartStreamRequest) SetInputs(v ConversationInputs)`

SetInputs sets Inputs field to given value.

### HasInputs

`func (o *ConversationRestartStreamRequest) HasInputs() bool`

HasInputs returns a boolean if a field has been set.

### GetStream

`func (o *ConversationRestartStreamRequest) GetStream() bool`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *ConversationRestartStreamRequest) GetStreamOk() (*bool, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *ConversationRestartStreamRequest) SetStream(v bool)`

SetStream sets Stream field to given value.

### HasStream

`func (o *ConversationRestartStreamRequest) HasStream() bool`

HasStream returns a boolean if a field has been set.

### GetStore

`func (o *ConversationRestartStreamRequest) GetStore() bool`

GetStore returns the Store field if non-nil, zero value otherwise.

### GetStoreOk

`func (o *ConversationRestartStreamRequest) GetStoreOk() (*bool, bool)`

GetStoreOk returns a tuple with the Store field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStore

`func (o *ConversationRestartStreamRequest) SetStore(v bool)`

SetStore sets Store field to given value.

### HasStore

`func (o *ConversationRestartStreamRequest) HasStore() bool`

HasStore returns a boolean if a field has been set.

### GetHandoffExecution

`func (o *ConversationRestartStreamRequest) GetHandoffExecution() string`

GetHandoffExecution returns the HandoffExecution field if non-nil, zero value otherwise.

### GetHandoffExecutionOk

`func (o *ConversationRestartStreamRequest) GetHandoffExecutionOk() (*string, bool)`

GetHandoffExecutionOk returns a tuple with the HandoffExecution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandoffExecution

`func (o *ConversationRestartStreamRequest) SetHandoffExecution(v string)`

SetHandoffExecution sets HandoffExecution field to given value.

### HasHandoffExecution

`func (o *ConversationRestartStreamRequest) HasHandoffExecution() bool`

HasHandoffExecution returns a boolean if a field has been set.

### GetCompletionArgs

`func (o *ConversationRestartStreamRequest) GetCompletionArgs() CompletionArgs`

GetCompletionArgs returns the CompletionArgs field if non-nil, zero value otherwise.

### GetCompletionArgsOk

`func (o *ConversationRestartStreamRequest) GetCompletionArgsOk() (*CompletionArgs, bool)`

GetCompletionArgsOk returns a tuple with the CompletionArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionArgs

`func (o *ConversationRestartStreamRequest) SetCompletionArgs(v CompletionArgs)`

SetCompletionArgs sets CompletionArgs field to given value.

### HasCompletionArgs

`func (o *ConversationRestartStreamRequest) HasCompletionArgs() bool`

HasCompletionArgs returns a boolean if a field has been set.

### GetGuardrails

`func (o *ConversationRestartStreamRequest) GetGuardrails() []GuardrailConfig`

GetGuardrails returns the Guardrails field if non-nil, zero value otherwise.

### GetGuardrailsOk

`func (o *ConversationRestartStreamRequest) GetGuardrailsOk() (*[]GuardrailConfig, bool)`

GetGuardrailsOk returns a tuple with the Guardrails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuardrails

`func (o *ConversationRestartStreamRequest) SetGuardrails(v []GuardrailConfig)`

SetGuardrails sets Guardrails field to given value.

### HasGuardrails

`func (o *ConversationRestartStreamRequest) HasGuardrails() bool`

HasGuardrails returns a boolean if a field has been set.

### GetMetadata

`func (o *ConversationRestartStreamRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ConversationRestartStreamRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ConversationRestartStreamRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ConversationRestartStreamRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetFromEntryId

`func (o *ConversationRestartStreamRequest) GetFromEntryId() string`

GetFromEntryId returns the FromEntryId field if non-nil, zero value otherwise.

### GetFromEntryIdOk

`func (o *ConversationRestartStreamRequest) GetFromEntryIdOk() (*string, bool)`

GetFromEntryIdOk returns a tuple with the FromEntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromEntryId

`func (o *ConversationRestartStreamRequest) SetFromEntryId(v string)`

SetFromEntryId sets FromEntryId field to given value.


### GetAgentVersion

`func (o *ConversationRestartStreamRequest) GetAgentVersion() AgentVersion1`

GetAgentVersion returns the AgentVersion field if non-nil, zero value otherwise.

### GetAgentVersionOk

`func (o *ConversationRestartStreamRequest) GetAgentVersionOk() (*AgentVersion1, bool)`

GetAgentVersionOk returns a tuple with the AgentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentVersion

`func (o *ConversationRestartStreamRequest) SetAgentVersion(v AgentVersion1)`

SetAgentVersion sets AgentVersion field to given value.

### HasAgentVersion

`func (o *ConversationRestartStreamRequest) HasAgentVersion() bool`

HasAgentVersion returns a boolean if a field has been set.

### SetAgentVersionNil

`func (o *ConversationRestartStreamRequest) SetAgentVersionNil(b bool)`

 SetAgentVersionNil sets the value for AgentVersion to be an explicit nil

### UnsetAgentVersion
`func (o *ConversationRestartStreamRequest) UnsetAgentVersion()`

UnsetAgentVersion ensures that no value is present for AgentVersion, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


