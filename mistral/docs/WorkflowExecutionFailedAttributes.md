# WorkflowExecutionFailedAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskId** | **string** | Unique identifier for the task within the workflow execution. | 
**Failure** | [**Failure**](Failure.md) | Details about the failure that caused the workflow to fail. | 

## Methods

### NewWorkflowExecutionFailedAttributes

`func NewWorkflowExecutionFailedAttributes(taskId string, failure Failure, ) *WorkflowExecutionFailedAttributes`

NewWorkflowExecutionFailedAttributes instantiates a new WorkflowExecutionFailedAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowExecutionFailedAttributesWithDefaults

`func NewWorkflowExecutionFailedAttributesWithDefaults() *WorkflowExecutionFailedAttributes`

NewWorkflowExecutionFailedAttributesWithDefaults instantiates a new WorkflowExecutionFailedAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskId

`func (o *WorkflowExecutionFailedAttributes) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *WorkflowExecutionFailedAttributes) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *WorkflowExecutionFailedAttributes) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.


### GetFailure

`func (o *WorkflowExecutionFailedAttributes) GetFailure() Failure`

GetFailure returns the Failure field if non-nil, zero value otherwise.

### GetFailureOk

`func (o *WorkflowExecutionFailedAttributes) GetFailureOk() (*Failure, bool)`

GetFailureOk returns a tuple with the Failure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailure

`func (o *WorkflowExecutionFailedAttributes) SetFailure(v Failure)`

SetFailure sets Failure field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


