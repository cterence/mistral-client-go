# WorkflowBulkError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowId** | **string** | The requested workflow ID | 
**Workflow** | Pointer to [**NullableWorkflow**](Workflow.md) | The workflow, if found | [optional] 
**Message** | **string** | Error message describing why the operation failed | 

## Methods

### NewWorkflowBulkError

`func NewWorkflowBulkError(workflowId string, message string, ) *WorkflowBulkError`

NewWorkflowBulkError instantiates a new WorkflowBulkError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowBulkErrorWithDefaults

`func NewWorkflowBulkErrorWithDefaults() *WorkflowBulkError`

NewWorkflowBulkErrorWithDefaults instantiates a new WorkflowBulkError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowId

`func (o *WorkflowBulkError) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *WorkflowBulkError) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *WorkflowBulkError) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.


### GetWorkflow

`func (o *WorkflowBulkError) GetWorkflow() Workflow`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *WorkflowBulkError) GetWorkflowOk() (*Workflow, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *WorkflowBulkError) SetWorkflow(v Workflow)`

SetWorkflow sets Workflow field to given value.

### HasWorkflow

`func (o *WorkflowBulkError) HasWorkflow() bool`

HasWorkflow returns a boolean if a field has been set.

### SetWorkflowNil

`func (o *WorkflowBulkError) SetWorkflowNil(b bool)`

 SetWorkflowNil sets the value for Workflow to be an explicit nil

### UnsetWorkflow
`func (o *WorkflowBulkError) UnsetWorkflow()`

UnsetWorkflow ensures that no value is present for Workflow, not even an explicit nil
### GetMessage

`func (o *WorkflowBulkError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *WorkflowBulkError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *WorkflowBulkError) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


