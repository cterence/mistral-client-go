# WorkflowTaskTimedOutAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskId** | **string** | Unique identifier for the task within the workflow execution. | 
**TimeoutType** | Pointer to **NullableString** | The type of timeout that occurred (e.g., &#39;START_TO_CLOSE&#39;, &#39;SCHEDULE_TO_START&#39;). | [optional] 

## Methods

### NewWorkflowTaskTimedOutAttributes

`func NewWorkflowTaskTimedOutAttributes(taskId string, ) *WorkflowTaskTimedOutAttributes`

NewWorkflowTaskTimedOutAttributes instantiates a new WorkflowTaskTimedOutAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTaskTimedOutAttributesWithDefaults

`func NewWorkflowTaskTimedOutAttributesWithDefaults() *WorkflowTaskTimedOutAttributes`

NewWorkflowTaskTimedOutAttributesWithDefaults instantiates a new WorkflowTaskTimedOutAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskId

`func (o *WorkflowTaskTimedOutAttributes) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *WorkflowTaskTimedOutAttributes) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *WorkflowTaskTimedOutAttributes) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.


### GetTimeoutType

`func (o *WorkflowTaskTimedOutAttributes) GetTimeoutType() string`

GetTimeoutType returns the TimeoutType field if non-nil, zero value otherwise.

### GetTimeoutTypeOk

`func (o *WorkflowTaskTimedOutAttributes) GetTimeoutTypeOk() (*string, bool)`

GetTimeoutTypeOk returns a tuple with the TimeoutType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutType

`func (o *WorkflowTaskTimedOutAttributes) SetTimeoutType(v string)`

SetTimeoutType sets TimeoutType field to given value.

### HasTimeoutType

`func (o *WorkflowTaskTimedOutAttributes) HasTimeoutType() bool`

HasTimeoutType returns a boolean if a field has been set.

### SetTimeoutTypeNil

`func (o *WorkflowTaskTimedOutAttributes) SetTimeoutTypeNil(b bool)`

 SetTimeoutTypeNil sets the value for TimeoutType to be an explicit nil

### UnsetTimeoutType
`func (o *WorkflowTaskTimedOutAttributes) UnsetTimeoutType()`

UnsetTimeoutType ensures that no value is present for TimeoutType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


