# OutputsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | Pointer to **string** |  | [optional] [default to "entry"]
**Type** | Pointer to **string** |  | [optional] [default to "message.output"]
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**CompletedAt** | Pointer to **time.Time** |  | [optional] 
**AgentId** | Pointer to **string** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Role** | Pointer to **string** |  | [optional] [default to "assistant"]
**Content** | [**Content1**](Content1.md) |  | 
**Name** | **string** |  | 
**Arguments** | [**FunctionCallEntryArguments**](FunctionCallEntryArguments.md) |  | 
**Info** | Pointer to **map[string]interface{}** |  | [optional] 
**ToolCallId** | **string** |  | 
**ConfirmationStatus** | Pointer to **string** |  | [optional] 
**PreviousAgentId** | **string** |  | 
**PreviousAgentName** | **string** |  | 
**NextAgentId** | **string** |  | 
**NextAgentName** | **string** |  | 

## Methods

### NewOutputsInner

`func NewOutputsInner(content Content1, name string, arguments FunctionCallEntryArguments, toolCallId string, previousAgentId string, previousAgentName string, nextAgentId string, nextAgentName string, ) *OutputsInner`

NewOutputsInner instantiates a new OutputsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputsInnerWithDefaults

`func NewOutputsInnerWithDefaults() *OutputsInner`

NewOutputsInnerWithDefaults instantiates a new OutputsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *OutputsInner) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *OutputsInner) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *OutputsInner) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *OutputsInner) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetType

`func (o *OutputsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OutputsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OutputsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OutputsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *OutputsInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OutputsInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OutputsInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *OutputsInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCompletedAt

`func (o *OutputsInner) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *OutputsInner) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *OutputsInner) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *OutputsInner) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetAgentId

`func (o *OutputsInner) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *OutputsInner) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *OutputsInner) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.

### HasAgentId

`func (o *OutputsInner) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.

### GetModel

`func (o *OutputsInner) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *OutputsInner) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *OutputsInner) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *OutputsInner) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetId

`func (o *OutputsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OutputsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OutputsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OutputsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRole

`func (o *OutputsInner) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *OutputsInner) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *OutputsInner) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *OutputsInner) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetContent

`func (o *OutputsInner) GetContent() Content1`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *OutputsInner) GetContentOk() (*Content1, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *OutputsInner) SetContent(v Content1)`

SetContent sets Content field to given value.


### GetName

`func (o *OutputsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OutputsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OutputsInner) SetName(v string)`

SetName sets Name field to given value.


### GetArguments

`func (o *OutputsInner) GetArguments() FunctionCallEntryArguments`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *OutputsInner) GetArgumentsOk() (*FunctionCallEntryArguments, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *OutputsInner) SetArguments(v FunctionCallEntryArguments)`

SetArguments sets Arguments field to given value.


### GetInfo

`func (o *OutputsInner) GetInfo() map[string]interface{}`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *OutputsInner) GetInfoOk() (*map[string]interface{}, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *OutputsInner) SetInfo(v map[string]interface{})`

SetInfo sets Info field to given value.

### HasInfo

`func (o *OutputsInner) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetToolCallId

`func (o *OutputsInner) GetToolCallId() string`

GetToolCallId returns the ToolCallId field if non-nil, zero value otherwise.

### GetToolCallIdOk

`func (o *OutputsInner) GetToolCallIdOk() (*string, bool)`

GetToolCallIdOk returns a tuple with the ToolCallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolCallId

`func (o *OutputsInner) SetToolCallId(v string)`

SetToolCallId sets ToolCallId field to given value.


### GetConfirmationStatus

`func (o *OutputsInner) GetConfirmationStatus() string`

GetConfirmationStatus returns the ConfirmationStatus field if non-nil, zero value otherwise.

### GetConfirmationStatusOk

`func (o *OutputsInner) GetConfirmationStatusOk() (*string, bool)`

GetConfirmationStatusOk returns a tuple with the ConfirmationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmationStatus

`func (o *OutputsInner) SetConfirmationStatus(v string)`

SetConfirmationStatus sets ConfirmationStatus field to given value.

### HasConfirmationStatus

`func (o *OutputsInner) HasConfirmationStatus() bool`

HasConfirmationStatus returns a boolean if a field has been set.

### GetPreviousAgentId

`func (o *OutputsInner) GetPreviousAgentId() string`

GetPreviousAgentId returns the PreviousAgentId field if non-nil, zero value otherwise.

### GetPreviousAgentIdOk

`func (o *OutputsInner) GetPreviousAgentIdOk() (*string, bool)`

GetPreviousAgentIdOk returns a tuple with the PreviousAgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousAgentId

`func (o *OutputsInner) SetPreviousAgentId(v string)`

SetPreviousAgentId sets PreviousAgentId field to given value.


### GetPreviousAgentName

`func (o *OutputsInner) GetPreviousAgentName() string`

GetPreviousAgentName returns the PreviousAgentName field if non-nil, zero value otherwise.

### GetPreviousAgentNameOk

`func (o *OutputsInner) GetPreviousAgentNameOk() (*string, bool)`

GetPreviousAgentNameOk returns a tuple with the PreviousAgentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousAgentName

`func (o *OutputsInner) SetPreviousAgentName(v string)`

SetPreviousAgentName sets PreviousAgentName field to given value.


### GetNextAgentId

`func (o *OutputsInner) GetNextAgentId() string`

GetNextAgentId returns the NextAgentId field if non-nil, zero value otherwise.

### GetNextAgentIdOk

`func (o *OutputsInner) GetNextAgentIdOk() (*string, bool)`

GetNextAgentIdOk returns a tuple with the NextAgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextAgentId

`func (o *OutputsInner) SetNextAgentId(v string)`

SetNextAgentId sets NextAgentId field to given value.


### GetNextAgentName

`func (o *OutputsInner) GetNextAgentName() string`

GetNextAgentName returns the NextAgentName field if non-nil, zero value otherwise.

### GetNextAgentNameOk

`func (o *OutputsInner) GetNextAgentNameOk() (*string, bool)`

GetNextAgentNameOk returns a tuple with the NextAgentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextAgentName

`func (o *OutputsInner) SetNextAgentName(v string)`

SetNextAgentName sets NextAgentName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


