# WorkflowBulkUnarchiveResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Unarchived** | [**[]Workflow**](Workflow.md) | Workflows that were successfully unarchived or were already unarchived | 
**Errored** | Pointer to [**[]WorkflowBulkError**](WorkflowBulkError.md) | Workflows that could not be unarchived and the corresponding error messages | [optional] 

## Methods

### NewWorkflowBulkUnarchiveResponse

`func NewWorkflowBulkUnarchiveResponse(unarchived []Workflow, ) *WorkflowBulkUnarchiveResponse`

NewWorkflowBulkUnarchiveResponse instantiates a new WorkflowBulkUnarchiveResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowBulkUnarchiveResponseWithDefaults

`func NewWorkflowBulkUnarchiveResponseWithDefaults() *WorkflowBulkUnarchiveResponse`

NewWorkflowBulkUnarchiveResponseWithDefaults instantiates a new WorkflowBulkUnarchiveResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnarchived

`func (o *WorkflowBulkUnarchiveResponse) GetUnarchived() []Workflow`

GetUnarchived returns the Unarchived field if non-nil, zero value otherwise.

### GetUnarchivedOk

`func (o *WorkflowBulkUnarchiveResponse) GetUnarchivedOk() (*[]Workflow, bool)`

GetUnarchivedOk returns a tuple with the Unarchived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnarchived

`func (o *WorkflowBulkUnarchiveResponse) SetUnarchived(v []Workflow)`

SetUnarchived sets Unarchived field to given value.


### GetErrored

`func (o *WorkflowBulkUnarchiveResponse) GetErrored() []WorkflowBulkError`

GetErrored returns the Errored field if non-nil, zero value otherwise.

### GetErroredOk

`func (o *WorkflowBulkUnarchiveResponse) GetErroredOk() (*[]WorkflowBulkError, bool)`

GetErroredOk returns a tuple with the Errored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrored

`func (o *WorkflowBulkUnarchiveResponse) SetErrored(v []WorkflowBulkError)`

SetErrored sets Errored field to given value.

### HasErrored

`func (o *WorkflowBulkUnarchiveResponse) HasErrored() bool`

HasErrored returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


