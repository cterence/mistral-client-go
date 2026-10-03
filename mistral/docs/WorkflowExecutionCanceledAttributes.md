# WorkflowExecutionCanceledAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskId** | **string** | Unique identifier for the task within the workflow execution. | 
**Reason** | Pointer to **NullableString** | Optional reason provided for the cancellation. | [optional] 

## Methods

### NewWorkflowExecutionCanceledAttributes

`func NewWorkflowExecutionCanceledAttributes(taskId string, ) *WorkflowExecutionCanceledAttributes`

NewWorkflowExecutionCanceledAttributes instantiates a new WorkflowExecutionCanceledAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowExecutionCanceledAttributesWithDefaults

`func NewWorkflowExecutionCanceledAttributesWithDefaults() *WorkflowExecutionCanceledAttributes`

NewWorkflowExecutionCanceledAttributesWithDefaults instantiates a new WorkflowExecutionCanceledAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskId

`func (o *WorkflowExecutionCanceledAttributes) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *WorkflowExecutionCanceledAttributes) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *WorkflowExecutionCanceledAttributes) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.


### GetReason

`func (o *WorkflowExecutionCanceledAttributes) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *WorkflowExecutionCanceledAttributes) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *WorkflowExecutionCanceledAttributes) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *WorkflowExecutionCanceledAttributes) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *WorkflowExecutionCanceledAttributes) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *WorkflowExecutionCanceledAttributes) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


