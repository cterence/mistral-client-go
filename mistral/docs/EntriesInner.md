# EntriesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | Pointer to **string** |  | [optional] [default to "entry"]
**Type** | Pointer to **string** |  | [optional] [default to "message.input"]
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**CompletedAt** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Role** | **string** |  | [default to "assistant"]
**Content** | [**Content1**](Content1.md) |  | 
**Prefix** | Pointer to **bool** |  | [optional] [default to false]
**AgentId** | Pointer to **string** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**ToolCallId** | **string** |  | 
**Result** | **string** |  | 
**Name** | [**Name**](Name.md) |  | 
**Arguments** | **string** |  | 
**ConfirmationStatus** | Pointer to **string** |  | [optional] 
**Info** | Pointer to **map[string]interface{}** |  | [optional] 
**PreviousAgentId** | **string** |  | 
**PreviousAgentName** | **string** |  | 
**NextAgentId** | **string** |  | 
**NextAgentName** | **string** |  | 

## Methods

### NewEntriesInner

`func NewEntriesInner(role string, content Content1, toolCallId string, result string, name Name, arguments string, previousAgentId string, previousAgentName string, nextAgentId string, nextAgentName string, ) *EntriesInner`

NewEntriesInner instantiates a new EntriesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntriesInnerWithDefaults

`func NewEntriesInnerWithDefaults() *EntriesInner`

NewEntriesInnerWithDefaults instantiates a new EntriesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *EntriesInner) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *EntriesInner) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *EntriesInner) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *EntriesInner) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetType

`func (o *EntriesInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EntriesInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EntriesInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EntriesInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *EntriesInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EntriesInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EntriesInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EntriesInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCompletedAt

`func (o *EntriesInner) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *EntriesInner) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *EntriesInner) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *EntriesInner) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetId

`func (o *EntriesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntriesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntriesInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EntriesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRole

`func (o *EntriesInner) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *EntriesInner) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *EntriesInner) SetRole(v string)`

SetRole sets Role field to given value.


### GetContent

`func (o *EntriesInner) GetContent() Content1`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *EntriesInner) GetContentOk() (*Content1, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *EntriesInner) SetContent(v Content1)`

SetContent sets Content field to given value.


### GetPrefix

`func (o *EntriesInner) GetPrefix() bool`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *EntriesInner) GetPrefixOk() (*bool, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *EntriesInner) SetPrefix(v bool)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *EntriesInner) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetAgentId

`func (o *EntriesInner) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *EntriesInner) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *EntriesInner) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.

### HasAgentId

`func (o *EntriesInner) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.

### GetModel

`func (o *EntriesInner) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *EntriesInner) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *EntriesInner) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *EntriesInner) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetToolCallId

`func (o *EntriesInner) GetToolCallId() string`

GetToolCallId returns the ToolCallId field if non-nil, zero value otherwise.

### GetToolCallIdOk

`func (o *EntriesInner) GetToolCallIdOk() (*string, bool)`

GetToolCallIdOk returns a tuple with the ToolCallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolCallId

`func (o *EntriesInner) SetToolCallId(v string)`

SetToolCallId sets ToolCallId field to given value.


### GetResult

`func (o *EntriesInner) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *EntriesInner) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *EntriesInner) SetResult(v string)`

SetResult sets Result field to given value.


### GetName

`func (o *EntriesInner) GetName() Name`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EntriesInner) GetNameOk() (*Name, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EntriesInner) SetName(v Name)`

SetName sets Name field to given value.


### GetArguments

`func (o *EntriesInner) GetArguments() string`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *EntriesInner) GetArgumentsOk() (*string, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *EntriesInner) SetArguments(v string)`

SetArguments sets Arguments field to given value.


### GetConfirmationStatus

`func (o *EntriesInner) GetConfirmationStatus() string`

GetConfirmationStatus returns the ConfirmationStatus field if non-nil, zero value otherwise.

### GetConfirmationStatusOk

`func (o *EntriesInner) GetConfirmationStatusOk() (*string, bool)`

GetConfirmationStatusOk returns a tuple with the ConfirmationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmationStatus

`func (o *EntriesInner) SetConfirmationStatus(v string)`

SetConfirmationStatus sets ConfirmationStatus field to given value.

### HasConfirmationStatus

`func (o *EntriesInner) HasConfirmationStatus() bool`

HasConfirmationStatus returns a boolean if a field has been set.

### GetInfo

`func (o *EntriesInner) GetInfo() map[string]interface{}`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *EntriesInner) GetInfoOk() (*map[string]interface{}, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *EntriesInner) SetInfo(v map[string]interface{})`

SetInfo sets Info field to given value.

### HasInfo

`func (o *EntriesInner) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetPreviousAgentId

`func (o *EntriesInner) GetPreviousAgentId() string`

GetPreviousAgentId returns the PreviousAgentId field if non-nil, zero value otherwise.

### GetPreviousAgentIdOk

`func (o *EntriesInner) GetPreviousAgentIdOk() (*string, bool)`

GetPreviousAgentIdOk returns a tuple with the PreviousAgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousAgentId

`func (o *EntriesInner) SetPreviousAgentId(v string)`

SetPreviousAgentId sets PreviousAgentId field to given value.


### GetPreviousAgentName

`func (o *EntriesInner) GetPreviousAgentName() string`

GetPreviousAgentName returns the PreviousAgentName field if non-nil, zero value otherwise.

### GetPreviousAgentNameOk

`func (o *EntriesInner) GetPreviousAgentNameOk() (*string, bool)`

GetPreviousAgentNameOk returns a tuple with the PreviousAgentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousAgentName

`func (o *EntriesInner) SetPreviousAgentName(v string)`

SetPreviousAgentName sets PreviousAgentName field to given value.


### GetNextAgentId

`func (o *EntriesInner) GetNextAgentId() string`

GetNextAgentId returns the NextAgentId field if non-nil, zero value otherwise.

### GetNextAgentIdOk

`func (o *EntriesInner) GetNextAgentIdOk() (*string, bool)`

GetNextAgentIdOk returns a tuple with the NextAgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextAgentId

`func (o *EntriesInner) SetNextAgentId(v string)`

SetNextAgentId sets NextAgentId field to given value.


### GetNextAgentName

`func (o *EntriesInner) GetNextAgentName() string`

GetNextAgentName returns the NextAgentName field if non-nil, zero value otherwise.

### GetNextAgentNameOk

`func (o *EntriesInner) GetNextAgentNameOk() (*string, bool)`

GetNextAgentNameOk returns a tuple with the NextAgentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextAgentName

`func (o *EntriesInner) SetNextAgentName(v string)`

SetNextAgentName sets NextAgentName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


