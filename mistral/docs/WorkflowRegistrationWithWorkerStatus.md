# WorkflowRegistrationWithWorkerStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the workflow registration | 
**TaskQueue** | **string** | Project name of the workflow | 
**Definition** | [**WorkflowCodeDefinition**](WorkflowCodeDefinition.md) |  | 
**WorkflowId** | **string** | Workflow ID of the workflow | 
**Workflow** | Pointer to [**NullableWorkflow**](Workflow.md) | Workflow of the workflow registration | [optional] 
**CompatibleWithChatAssistant** | Pointer to **bool** | Whether the workflow is compatible with chat assistant | [optional] [default to false]
**Active** | **bool** | Whether the workflow registration is active | 

## Methods

### NewWorkflowRegistrationWithWorkerStatus

`func NewWorkflowRegistrationWithWorkerStatus(id string, taskQueue string, definition WorkflowCodeDefinition, workflowId string, active bool, ) *WorkflowRegistrationWithWorkerStatus`

NewWorkflowRegistrationWithWorkerStatus instantiates a new WorkflowRegistrationWithWorkerStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowRegistrationWithWorkerStatusWithDefaults

`func NewWorkflowRegistrationWithWorkerStatusWithDefaults() *WorkflowRegistrationWithWorkerStatus`

NewWorkflowRegistrationWithWorkerStatusWithDefaults instantiates a new WorkflowRegistrationWithWorkerStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkflowRegistrationWithWorkerStatus) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowRegistrationWithWorkerStatus) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowRegistrationWithWorkerStatus) SetId(v string)`

SetId sets Id field to given value.


### GetTaskQueue

`func (o *WorkflowRegistrationWithWorkerStatus) GetTaskQueue() string`

GetTaskQueue returns the TaskQueue field if non-nil, zero value otherwise.

### GetTaskQueueOk

`func (o *WorkflowRegistrationWithWorkerStatus) GetTaskQueueOk() (*string, bool)`

GetTaskQueueOk returns a tuple with the TaskQueue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskQueue

`func (o *WorkflowRegistrationWithWorkerStatus) SetTaskQueue(v string)`

SetTaskQueue sets TaskQueue field to given value.


### GetDefinition

`func (o *WorkflowRegistrationWithWorkerStatus) GetDefinition() WorkflowCodeDefinition`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *WorkflowRegistrationWithWorkerStatus) GetDefinitionOk() (*WorkflowCodeDefinition, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *WorkflowRegistrationWithWorkerStatus) SetDefinition(v WorkflowCodeDefinition)`

SetDefinition sets Definition field to given value.


### GetWorkflowId

`func (o *WorkflowRegistrationWithWorkerStatus) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *WorkflowRegistrationWithWorkerStatus) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *WorkflowRegistrationWithWorkerStatus) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.


### GetWorkflow

`func (o *WorkflowRegistrationWithWorkerStatus) GetWorkflow() Workflow`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *WorkflowRegistrationWithWorkerStatus) GetWorkflowOk() (*Workflow, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *WorkflowRegistrationWithWorkerStatus) SetWorkflow(v Workflow)`

SetWorkflow sets Workflow field to given value.

### HasWorkflow

`func (o *WorkflowRegistrationWithWorkerStatus) HasWorkflow() bool`

HasWorkflow returns a boolean if a field has been set.

### SetWorkflowNil

`func (o *WorkflowRegistrationWithWorkerStatus) SetWorkflowNil(b bool)`

 SetWorkflowNil sets the value for Workflow to be an explicit nil

### UnsetWorkflow
`func (o *WorkflowRegistrationWithWorkerStatus) UnsetWorkflow()`

UnsetWorkflow ensures that no value is present for Workflow, not even an explicit nil
### GetCompatibleWithChatAssistant

`func (o *WorkflowRegistrationWithWorkerStatus) GetCompatibleWithChatAssistant() bool`

GetCompatibleWithChatAssistant returns the CompatibleWithChatAssistant field if non-nil, zero value otherwise.

### GetCompatibleWithChatAssistantOk

`func (o *WorkflowRegistrationWithWorkerStatus) GetCompatibleWithChatAssistantOk() (*bool, bool)`

GetCompatibleWithChatAssistantOk returns a tuple with the CompatibleWithChatAssistant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatibleWithChatAssistant

`func (o *WorkflowRegistrationWithWorkerStatus) SetCompatibleWithChatAssistant(v bool)`

SetCompatibleWithChatAssistant sets CompatibleWithChatAssistant field to given value.

### HasCompatibleWithChatAssistant

`func (o *WorkflowRegistrationWithWorkerStatus) HasCompatibleWithChatAssistant() bool`

HasCompatibleWithChatAssistant returns a boolean if a field has been set.

### GetActive

`func (o *WorkflowRegistrationWithWorkerStatus) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *WorkflowRegistrationWithWorkerStatus) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *WorkflowRegistrationWithWorkerStatus) SetActive(v bool)`

SetActive sets Active field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


