# ToolExecutionStartedEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "tool.execution.started"]
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**OutputIndex** | Pointer to **int32** |  | [optional] [default to 0]
**Id** | **string** |  | 
**Model** | Pointer to **NullableString** |  | [optional] 
**AgentId** | Pointer to **NullableString** |  | [optional] 
**Name** | [**Name**](Name.md) |  | 
**Arguments** | **string** |  | 

## Methods

### NewToolExecutionStartedEvent

`func NewToolExecutionStartedEvent(id string, name Name, arguments string, ) *ToolExecutionStartedEvent`

NewToolExecutionStartedEvent instantiates a new ToolExecutionStartedEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewToolExecutionStartedEventWithDefaults

`func NewToolExecutionStartedEventWithDefaults() *ToolExecutionStartedEvent`

NewToolExecutionStartedEventWithDefaults instantiates a new ToolExecutionStartedEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ToolExecutionStartedEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ToolExecutionStartedEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ToolExecutionStartedEvent) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ToolExecutionStartedEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ToolExecutionStartedEvent) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ToolExecutionStartedEvent) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ToolExecutionStartedEvent) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ToolExecutionStartedEvent) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetOutputIndex

`func (o *ToolExecutionStartedEvent) GetOutputIndex() int32`

GetOutputIndex returns the OutputIndex field if non-nil, zero value otherwise.

### GetOutputIndexOk

`func (o *ToolExecutionStartedEvent) GetOutputIndexOk() (*int32, bool)`

GetOutputIndexOk returns a tuple with the OutputIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputIndex

`func (o *ToolExecutionStartedEvent) SetOutputIndex(v int32)`

SetOutputIndex sets OutputIndex field to given value.

### HasOutputIndex

`func (o *ToolExecutionStartedEvent) HasOutputIndex() bool`

HasOutputIndex returns a boolean if a field has been set.

### GetId

`func (o *ToolExecutionStartedEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ToolExecutionStartedEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ToolExecutionStartedEvent) SetId(v string)`

SetId sets Id field to given value.


### GetModel

`func (o *ToolExecutionStartedEvent) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ToolExecutionStartedEvent) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ToolExecutionStartedEvent) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ToolExecutionStartedEvent) HasModel() bool`

HasModel returns a boolean if a field has been set.

### SetModelNil

`func (o *ToolExecutionStartedEvent) SetModelNil(b bool)`

 SetModelNil sets the value for Model to be an explicit nil

### UnsetModel
`func (o *ToolExecutionStartedEvent) UnsetModel()`

UnsetModel ensures that no value is present for Model, not even an explicit nil
### GetAgentId

`func (o *ToolExecutionStartedEvent) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *ToolExecutionStartedEvent) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *ToolExecutionStartedEvent) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.

### HasAgentId

`func (o *ToolExecutionStartedEvent) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.

### SetAgentIdNil

`func (o *ToolExecutionStartedEvent) SetAgentIdNil(b bool)`

 SetAgentIdNil sets the value for AgentId to be an explicit nil

### UnsetAgentId
`func (o *ToolExecutionStartedEvent) UnsetAgentId()`

UnsetAgentId ensures that no value is present for AgentId, not even an explicit nil
### GetName

`func (o *ToolExecutionStartedEvent) GetName() Name`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ToolExecutionStartedEvent) GetNameOk() (*Name, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ToolExecutionStartedEvent) SetName(v Name)`

SetName sets Name field to given value.


### GetArguments

`func (o *ToolExecutionStartedEvent) GetArguments() string`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *ToolExecutionStartedEvent) GetArgumentsOk() (*string, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *ToolExecutionStartedEvent) SetArguments(v string)`

SetArguments sets Arguments field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


