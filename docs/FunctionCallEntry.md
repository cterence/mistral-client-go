# FunctionCallEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | Pointer to **string** |  | [optional] [default to "entry"]
**Type** | Pointer to **string** |  | [optional] [default to "function.call"]
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**CompletedAt** | Pointer to **NullableTime** |  | [optional] 
**AgentId** | Pointer to **NullableString** |  | [optional] 
**Model** | Pointer to **NullableString** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**ToolCallId** | **string** |  | 
**Name** | **string** |  | 
**Arguments** | [**FunctionCallEntryArguments**](FunctionCallEntryArguments.md) |  | 
**ConfirmationStatus** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewFunctionCallEntry

`func NewFunctionCallEntry(toolCallId string, name string, arguments FunctionCallEntryArguments, ) *FunctionCallEntry`

NewFunctionCallEntry instantiates a new FunctionCallEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionCallEntryWithDefaults

`func NewFunctionCallEntryWithDefaults() *FunctionCallEntry`

NewFunctionCallEntryWithDefaults instantiates a new FunctionCallEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *FunctionCallEntry) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *FunctionCallEntry) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *FunctionCallEntry) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *FunctionCallEntry) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetType

`func (o *FunctionCallEntry) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FunctionCallEntry) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FunctionCallEntry) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FunctionCallEntry) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FunctionCallEntry) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FunctionCallEntry) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FunctionCallEntry) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FunctionCallEntry) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCompletedAt

`func (o *FunctionCallEntry) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *FunctionCallEntry) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *FunctionCallEntry) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *FunctionCallEntry) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### SetCompletedAtNil

`func (o *FunctionCallEntry) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *FunctionCallEntry) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
### GetAgentId

`func (o *FunctionCallEntry) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *FunctionCallEntry) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *FunctionCallEntry) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.

### HasAgentId

`func (o *FunctionCallEntry) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.

### SetAgentIdNil

`func (o *FunctionCallEntry) SetAgentIdNil(b bool)`

 SetAgentIdNil sets the value for AgentId to be an explicit nil

### UnsetAgentId
`func (o *FunctionCallEntry) UnsetAgentId()`

UnsetAgentId ensures that no value is present for AgentId, not even an explicit nil
### GetModel

`func (o *FunctionCallEntry) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *FunctionCallEntry) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *FunctionCallEntry) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *FunctionCallEntry) HasModel() bool`

HasModel returns a boolean if a field has been set.

### SetModelNil

`func (o *FunctionCallEntry) SetModelNil(b bool)`

 SetModelNil sets the value for Model to be an explicit nil

### UnsetModel
`func (o *FunctionCallEntry) UnsetModel()`

UnsetModel ensures that no value is present for Model, not even an explicit nil
### GetId

`func (o *FunctionCallEntry) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FunctionCallEntry) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FunctionCallEntry) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FunctionCallEntry) HasId() bool`

HasId returns a boolean if a field has been set.

### GetToolCallId

`func (o *FunctionCallEntry) GetToolCallId() string`

GetToolCallId returns the ToolCallId field if non-nil, zero value otherwise.

### GetToolCallIdOk

`func (o *FunctionCallEntry) GetToolCallIdOk() (*string, bool)`

GetToolCallIdOk returns a tuple with the ToolCallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolCallId

`func (o *FunctionCallEntry) SetToolCallId(v string)`

SetToolCallId sets ToolCallId field to given value.


### GetName

`func (o *FunctionCallEntry) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FunctionCallEntry) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FunctionCallEntry) SetName(v string)`

SetName sets Name field to given value.


### GetArguments

`func (o *FunctionCallEntry) GetArguments() FunctionCallEntryArguments`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *FunctionCallEntry) GetArgumentsOk() (*FunctionCallEntryArguments, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *FunctionCallEntry) SetArguments(v FunctionCallEntryArguments)`

SetArguments sets Arguments field to given value.


### GetConfirmationStatus

`func (o *FunctionCallEntry) GetConfirmationStatus() string`

GetConfirmationStatus returns the ConfirmationStatus field if non-nil, zero value otherwise.

### GetConfirmationStatusOk

`func (o *FunctionCallEntry) GetConfirmationStatusOk() (*string, bool)`

GetConfirmationStatusOk returns a tuple with the ConfirmationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmationStatus

`func (o *FunctionCallEntry) SetConfirmationStatus(v string)`

SetConfirmationStatus sets ConfirmationStatus field to given value.

### HasConfirmationStatus

`func (o *FunctionCallEntry) HasConfirmationStatus() bool`

HasConfirmationStatus returns a boolean if a field has been set.

### SetConfirmationStatusNil

`func (o *FunctionCallEntry) SetConfirmationStatusNil(b bool)`

 SetConfirmationStatusNil sets the value for ConfirmationStatus to be an explicit nil

### UnsetConfirmationStatus
`func (o *FunctionCallEntry) UnsetConfirmationStatus()`

UnsetConfirmationStatus ensures that no value is present for ConfirmationStatus, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


