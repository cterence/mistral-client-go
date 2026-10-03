# WorkflowExecutionContinuedAsNewAttributesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskId** | **string** | Unique identifier for the task within the workflow execution. | 
**NewExecutionRunId** | **string** | The run ID of the new workflow execution that continues this workflow. | 
**WorkflowName** | **string** | The registered name of the continued workflow. | 
**Input** | [**JSONPayloadResponse**](JSONPayloadResponse.md) | The input arguments passed to the new workflow execution. | 

## Methods

### NewWorkflowExecutionContinuedAsNewAttributesResponse

`func NewWorkflowExecutionContinuedAsNewAttributesResponse(taskId string, newExecutionRunId string, workflowName string, input JSONPayloadResponse, ) *WorkflowExecutionContinuedAsNewAttributesResponse`

NewWorkflowExecutionContinuedAsNewAttributesResponse instantiates a new WorkflowExecutionContinuedAsNewAttributesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowExecutionContinuedAsNewAttributesResponseWithDefaults

`func NewWorkflowExecutionContinuedAsNewAttributesResponseWithDefaults() *WorkflowExecutionContinuedAsNewAttributesResponse`

NewWorkflowExecutionContinuedAsNewAttributesResponseWithDefaults instantiates a new WorkflowExecutionContinuedAsNewAttributesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskId

`func (o *WorkflowExecutionContinuedAsNewAttributesResponse) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *WorkflowExecutionContinuedAsNewAttributesResponse) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *WorkflowExecutionContinuedAsNewAttributesResponse) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.


### GetNewExecutionRunId

`func (o *WorkflowExecutionContinuedAsNewAttributesResponse) GetNewExecutionRunId() string`

GetNewExecutionRunId returns the NewExecutionRunId field if non-nil, zero value otherwise.

### GetNewExecutionRunIdOk

`func (o *WorkflowExecutionContinuedAsNewAttributesResponse) GetNewExecutionRunIdOk() (*string, bool)`

GetNewExecutionRunIdOk returns a tuple with the NewExecutionRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewExecutionRunId

`func (o *WorkflowExecutionContinuedAsNewAttributesResponse) SetNewExecutionRunId(v string)`

SetNewExecutionRunId sets NewExecutionRunId field to given value.


### GetWorkflowName

`func (o *WorkflowExecutionContinuedAsNewAttributesResponse) GetWorkflowName() string`

GetWorkflowName returns the WorkflowName field if non-nil, zero value otherwise.

### GetWorkflowNameOk

`func (o *WorkflowExecutionContinuedAsNewAttributesResponse) GetWorkflowNameOk() (*string, bool)`

GetWorkflowNameOk returns a tuple with the WorkflowName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowName

`func (o *WorkflowExecutionContinuedAsNewAttributesResponse) SetWorkflowName(v string)`

SetWorkflowName sets WorkflowName field to given value.


### GetInput

`func (o *WorkflowExecutionContinuedAsNewAttributesResponse) GetInput() JSONPayloadResponse`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *WorkflowExecutionContinuedAsNewAttributesResponse) GetInputOk() (*JSONPayloadResponse, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *WorkflowExecutionContinuedAsNewAttributesResponse) SetInput(v JSONPayloadResponse)`

SetInput sets Input field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


