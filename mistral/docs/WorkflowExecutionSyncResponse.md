# WorkflowExecutionSyncResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowName** | **string** | Name of the workflow that was executed | 
**ExecutionId** | **string** | ID of the workflow execution | 
**Result** | **interface{}** | The result of the workflow execution | 

## Methods

### NewWorkflowExecutionSyncResponse

`func NewWorkflowExecutionSyncResponse(workflowName string, executionId string, result interface{}, ) *WorkflowExecutionSyncResponse`

NewWorkflowExecutionSyncResponse instantiates a new WorkflowExecutionSyncResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowExecutionSyncResponseWithDefaults

`func NewWorkflowExecutionSyncResponseWithDefaults() *WorkflowExecutionSyncResponse`

NewWorkflowExecutionSyncResponseWithDefaults instantiates a new WorkflowExecutionSyncResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowName

`func (o *WorkflowExecutionSyncResponse) GetWorkflowName() string`

GetWorkflowName returns the WorkflowName field if non-nil, zero value otherwise.

### GetWorkflowNameOk

`func (o *WorkflowExecutionSyncResponse) GetWorkflowNameOk() (*string, bool)`

GetWorkflowNameOk returns a tuple with the WorkflowName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowName

`func (o *WorkflowExecutionSyncResponse) SetWorkflowName(v string)`

SetWorkflowName sets WorkflowName field to given value.


### GetExecutionId

`func (o *WorkflowExecutionSyncResponse) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *WorkflowExecutionSyncResponse) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *WorkflowExecutionSyncResponse) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.


### GetResult

`func (o *WorkflowExecutionSyncResponse) GetResult() interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *WorkflowExecutionSyncResponse) GetResultOk() (*interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *WorkflowExecutionSyncResponse) SetResult(v interface{})`

SetResult sets Result field to given value.


### SetResultNil

`func (o *WorkflowExecutionSyncResponse) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *WorkflowExecutionSyncResponse) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


