# WorkflowGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Workflow** | [**WorkflowWithWorkerStatus**](WorkflowWithWorkerStatus.md) | The workflow spec | 

## Methods

### NewWorkflowGetResponse

`func NewWorkflowGetResponse(workflow WorkflowWithWorkerStatus, ) *WorkflowGetResponse`

NewWorkflowGetResponse instantiates a new WorkflowGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowGetResponseWithDefaults

`func NewWorkflowGetResponseWithDefaults() *WorkflowGetResponse`

NewWorkflowGetResponseWithDefaults instantiates a new WorkflowGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflow

`func (o *WorkflowGetResponse) GetWorkflow() WorkflowWithWorkerStatus`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *WorkflowGetResponse) GetWorkflowOk() (*WorkflowWithWorkerStatus, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *WorkflowGetResponse) SetWorkflow(v WorkflowWithWorkerStatus)`

SetWorkflow sets Workflow field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


