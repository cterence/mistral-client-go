# WorkflowExecutionCompletedAttributesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskId** | **string** | Unique identifier for the task within the workflow execution. | 
**Result** | [**JSONPayloadResponse**](JSONPayloadResponse.md) | The final result returned by the workflow. | 
**Attempt** | Pointer to **int32** | Workflow retry attempt number. 1 on first run and CAN; &gt;1 on workflow-level retries. | [optional] [default to 1]

## Methods

### NewWorkflowExecutionCompletedAttributesResponse

`func NewWorkflowExecutionCompletedAttributesResponse(taskId string, result JSONPayloadResponse, ) *WorkflowExecutionCompletedAttributesResponse`

NewWorkflowExecutionCompletedAttributesResponse instantiates a new WorkflowExecutionCompletedAttributesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowExecutionCompletedAttributesResponseWithDefaults

`func NewWorkflowExecutionCompletedAttributesResponseWithDefaults() *WorkflowExecutionCompletedAttributesResponse`

NewWorkflowExecutionCompletedAttributesResponseWithDefaults instantiates a new WorkflowExecutionCompletedAttributesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskId

`func (o *WorkflowExecutionCompletedAttributesResponse) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *WorkflowExecutionCompletedAttributesResponse) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *WorkflowExecutionCompletedAttributesResponse) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.


### GetResult

`func (o *WorkflowExecutionCompletedAttributesResponse) GetResult() JSONPayloadResponse`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *WorkflowExecutionCompletedAttributesResponse) GetResultOk() (*JSONPayloadResponse, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *WorkflowExecutionCompletedAttributesResponse) SetResult(v JSONPayloadResponse)`

SetResult sets Result field to given value.


### GetAttempt

`func (o *WorkflowExecutionCompletedAttributesResponse) GetAttempt() int32`

GetAttempt returns the Attempt field if non-nil, zero value otherwise.

### GetAttemptOk

`func (o *WorkflowExecutionCompletedAttributesResponse) GetAttemptOk() (*int32, bool)`

GetAttemptOk returns a tuple with the Attempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttempt

`func (o *WorkflowExecutionCompletedAttributesResponse) SetAttempt(v int32)`

SetAttempt sets Attempt field to given value.

### HasAttempt

`func (o *WorkflowExecutionCompletedAttributesResponse) HasAttempt() bool`

HasAttempt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


