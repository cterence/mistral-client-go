# WorkflowExecutionStartedAttributesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskId** | **string** | Unique identifier for the task within the workflow execution. | 
**WorkflowName** | **string** | The registered name of the workflow being executed. | 
**Input** | [**JSONPayloadResponse**](JSONPayloadResponse.md) | The input arguments passed to the workflow. | 

## Methods

### NewWorkflowExecutionStartedAttributesResponse

`func NewWorkflowExecutionStartedAttributesResponse(taskId string, workflowName string, input JSONPayloadResponse, ) *WorkflowExecutionStartedAttributesResponse`

NewWorkflowExecutionStartedAttributesResponse instantiates a new WorkflowExecutionStartedAttributesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowExecutionStartedAttributesResponseWithDefaults

`func NewWorkflowExecutionStartedAttributesResponseWithDefaults() *WorkflowExecutionStartedAttributesResponse`

NewWorkflowExecutionStartedAttributesResponseWithDefaults instantiates a new WorkflowExecutionStartedAttributesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskId

`func (o *WorkflowExecutionStartedAttributesResponse) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *WorkflowExecutionStartedAttributesResponse) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *WorkflowExecutionStartedAttributesResponse) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.


### GetWorkflowName

`func (o *WorkflowExecutionStartedAttributesResponse) GetWorkflowName() string`

GetWorkflowName returns the WorkflowName field if non-nil, zero value otherwise.

### GetWorkflowNameOk

`func (o *WorkflowExecutionStartedAttributesResponse) GetWorkflowNameOk() (*string, bool)`

GetWorkflowNameOk returns a tuple with the WorkflowName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowName

`func (o *WorkflowExecutionStartedAttributesResponse) SetWorkflowName(v string)`

SetWorkflowName sets WorkflowName field to given value.


### GetInput

`func (o *WorkflowExecutionStartedAttributesResponse) GetInput() JSONPayloadResponse`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *WorkflowExecutionStartedAttributesResponse) GetInputOk() (*JSONPayloadResponse, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *WorkflowExecutionStartedAttributesResponse) SetInput(v JSONPayloadResponse)`

SetInput sets Input field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


