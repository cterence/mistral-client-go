# WorkflowRegistrationGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowRegistration** | [**WorkflowRegistrationWithWorkerStatus**](WorkflowRegistrationWithWorkerStatus.md) | The workflow registration | 
**WorkflowVersion** | [**WorkflowRegistrationWithWorkerStatus**](WorkflowRegistrationWithWorkerStatus.md) | Deprecated: use workflow_registration | [readonly] 

## Methods

### NewWorkflowRegistrationGetResponse

`func NewWorkflowRegistrationGetResponse(workflowRegistration WorkflowRegistrationWithWorkerStatus, workflowVersion WorkflowRegistrationWithWorkerStatus, ) *WorkflowRegistrationGetResponse`

NewWorkflowRegistrationGetResponse instantiates a new WorkflowRegistrationGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowRegistrationGetResponseWithDefaults

`func NewWorkflowRegistrationGetResponseWithDefaults() *WorkflowRegistrationGetResponse`

NewWorkflowRegistrationGetResponseWithDefaults instantiates a new WorkflowRegistrationGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowRegistration

`func (o *WorkflowRegistrationGetResponse) GetWorkflowRegistration() WorkflowRegistrationWithWorkerStatus`

GetWorkflowRegistration returns the WorkflowRegistration field if non-nil, zero value otherwise.

### GetWorkflowRegistrationOk

`func (o *WorkflowRegistrationGetResponse) GetWorkflowRegistrationOk() (*WorkflowRegistrationWithWorkerStatus, bool)`

GetWorkflowRegistrationOk returns a tuple with the WorkflowRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowRegistration

`func (o *WorkflowRegistrationGetResponse) SetWorkflowRegistration(v WorkflowRegistrationWithWorkerStatus)`

SetWorkflowRegistration sets WorkflowRegistration field to given value.


### GetWorkflowVersion

`func (o *WorkflowRegistrationGetResponse) GetWorkflowVersion() WorkflowRegistrationWithWorkerStatus`

GetWorkflowVersion returns the WorkflowVersion field if non-nil, zero value otherwise.

### GetWorkflowVersionOk

`func (o *WorkflowRegistrationGetResponse) GetWorkflowVersionOk() (*WorkflowRegistrationWithWorkerStatus, bool)`

GetWorkflowVersionOk returns a tuple with the WorkflowVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowVersion

`func (o *WorkflowRegistrationGetResponse) SetWorkflowVersion(v WorkflowRegistrationWithWorkerStatus)`

SetWorkflowVersion sets WorkflowVersion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


