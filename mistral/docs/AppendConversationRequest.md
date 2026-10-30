# AppendConversationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Inputs** | Pointer to [**ConversationInputs**](ConversationInputs.md) |  | [optional] 
**Stream** | Pointer to **bool** | Whether to stream back partial progress. Otherwise, the server will hold the request open until the timeout or until completion, with the response containing the full result as JSON. | [optional] [default to false]
**Store** | Pointer to **bool** | Whether to store the results into our servers or not. | [optional] [default to true]
**HandoffExecution** | Pointer to **string** |  | [optional] [default to "server"]
**CompletionArgs** | Pointer to [**CompletionArgs**](CompletionArgs.md) | Completion arguments that will be used to generate assistant responses. Can be overridden at each message request. | [optional] 
**ToolConfirmations** | Pointer to [**[]ToolCallConfirmation**](ToolCallConfirmation.md) |  | [optional] 

## Methods

### NewAppendConversationRequest

`func NewAppendConversationRequest() *AppendConversationRequest`

NewAppendConversationRequest instantiates a new AppendConversationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppendConversationRequestWithDefaults

`func NewAppendConversationRequestWithDefaults() *AppendConversationRequest`

NewAppendConversationRequestWithDefaults instantiates a new AppendConversationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInputs

`func (o *AppendConversationRequest) GetInputs() ConversationInputs`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *AppendConversationRequest) GetInputsOk() (*ConversationInputs, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *AppendConversationRequest) SetInputs(v ConversationInputs)`

SetInputs sets Inputs field to given value.

### HasInputs

`func (o *AppendConversationRequest) HasInputs() bool`

HasInputs returns a boolean if a field has been set.

### GetStream

`func (o *AppendConversationRequest) GetStream() bool`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *AppendConversationRequest) GetStreamOk() (*bool, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *AppendConversationRequest) SetStream(v bool)`

SetStream sets Stream field to given value.

### HasStream

`func (o *AppendConversationRequest) HasStream() bool`

HasStream returns a boolean if a field has been set.

### GetStore

`func (o *AppendConversationRequest) GetStore() bool`

GetStore returns the Store field if non-nil, zero value otherwise.

### GetStoreOk

`func (o *AppendConversationRequest) GetStoreOk() (*bool, bool)`

GetStoreOk returns a tuple with the Store field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStore

`func (o *AppendConversationRequest) SetStore(v bool)`

SetStore sets Store field to given value.

### HasStore

`func (o *AppendConversationRequest) HasStore() bool`

HasStore returns a boolean if a field has been set.

### GetHandoffExecution

`func (o *AppendConversationRequest) GetHandoffExecution() string`

GetHandoffExecution returns the HandoffExecution field if non-nil, zero value otherwise.

### GetHandoffExecutionOk

`func (o *AppendConversationRequest) GetHandoffExecutionOk() (*string, bool)`

GetHandoffExecutionOk returns a tuple with the HandoffExecution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandoffExecution

`func (o *AppendConversationRequest) SetHandoffExecution(v string)`

SetHandoffExecution sets HandoffExecution field to given value.

### HasHandoffExecution

`func (o *AppendConversationRequest) HasHandoffExecution() bool`

HasHandoffExecution returns a boolean if a field has been set.

### GetCompletionArgs

`func (o *AppendConversationRequest) GetCompletionArgs() CompletionArgs`

GetCompletionArgs returns the CompletionArgs field if non-nil, zero value otherwise.

### GetCompletionArgsOk

`func (o *AppendConversationRequest) GetCompletionArgsOk() (*CompletionArgs, bool)`

GetCompletionArgsOk returns a tuple with the CompletionArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionArgs

`func (o *AppendConversationRequest) SetCompletionArgs(v CompletionArgs)`

SetCompletionArgs sets CompletionArgs field to given value.

### HasCompletionArgs

`func (o *AppendConversationRequest) HasCompletionArgs() bool`

HasCompletionArgs returns a boolean if a field has been set.

### GetToolConfirmations

`func (o *AppendConversationRequest) GetToolConfirmations() []ToolCallConfirmation`

GetToolConfirmations returns the ToolConfirmations field if non-nil, zero value otherwise.

### GetToolConfirmationsOk

`func (o *AppendConversationRequest) GetToolConfirmationsOk() (*[]ToolCallConfirmation, bool)`

GetToolConfirmationsOk returns a tuple with the ToolConfirmations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolConfirmations

`func (o *AppendConversationRequest) SetToolConfirmations(v []ToolCallConfirmation)`

SetToolConfirmations sets ToolConfirmations field to given value.

### HasToolConfirmations

`func (o *AppendConversationRequest) HasToolConfirmations() bool`

HasToolConfirmations returns a boolean if a field has been set.

### SetToolConfirmationsNil

`func (o *AppendConversationRequest) SetToolConfirmationsNil(b bool)`

 SetToolConfirmationsNil sets the value for ToolConfirmations to be an explicit nil

### UnsetToolConfirmations
`func (o *AppendConversationRequest) UnsetToolConfirmations()`

UnsetToolConfirmations ensures that no value is present for ToolConfirmations, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


