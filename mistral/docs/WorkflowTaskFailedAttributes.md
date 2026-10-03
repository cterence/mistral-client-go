# WorkflowTaskFailedAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskId** | **string** | Unique identifier for the task within the workflow execution. | 
**Failure** | [**Failure**](Failure.md) | Details about the failure that caused the task to fail. | 

## Methods

### NewWorkflowTaskFailedAttributes

`func NewWorkflowTaskFailedAttributes(taskId string, failure Failure, ) *WorkflowTaskFailedAttributes`

NewWorkflowTaskFailedAttributes instantiates a new WorkflowTaskFailedAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTaskFailedAttributesWithDefaults

`func NewWorkflowTaskFailedAttributesWithDefaults() *WorkflowTaskFailedAttributes`

NewWorkflowTaskFailedAttributesWithDefaults instantiates a new WorkflowTaskFailedAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskId

`func (o *WorkflowTaskFailedAttributes) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *WorkflowTaskFailedAttributes) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *WorkflowTaskFailedAttributes) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.


### GetFailure

`func (o *WorkflowTaskFailedAttributes) GetFailure() Failure`

GetFailure returns the Failure field if non-nil, zero value otherwise.

### GetFailureOk

`func (o *WorkflowTaskFailedAttributes) GetFailureOk() (*Failure, bool)`

GetFailureOk returns a tuple with the Failure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailure

`func (o *WorkflowTaskFailedAttributes) SetFailure(v Failure)`

SetFailure sets Failure field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


