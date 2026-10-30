# WorkflowBulkArchiveResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Archived** | [**[]Workflow**](Workflow.md) | Workflows that were successfully archived or were already archived | 
**Errored** | Pointer to [**[]WorkflowBulkError**](WorkflowBulkError.md) | Workflows that could not be archived and the corresponding error messages | [optional] 

## Methods

### NewWorkflowBulkArchiveResponse

`func NewWorkflowBulkArchiveResponse(archived []Workflow, ) *WorkflowBulkArchiveResponse`

NewWorkflowBulkArchiveResponse instantiates a new WorkflowBulkArchiveResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowBulkArchiveResponseWithDefaults

`func NewWorkflowBulkArchiveResponseWithDefaults() *WorkflowBulkArchiveResponse`

NewWorkflowBulkArchiveResponseWithDefaults instantiates a new WorkflowBulkArchiveResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchived

`func (o *WorkflowBulkArchiveResponse) GetArchived() []Workflow`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *WorkflowBulkArchiveResponse) GetArchivedOk() (*[]Workflow, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *WorkflowBulkArchiveResponse) SetArchived(v []Workflow)`

SetArchived sets Archived field to given value.


### GetErrored

`func (o *WorkflowBulkArchiveResponse) GetErrored() []WorkflowBulkError`

GetErrored returns the Errored field if non-nil, zero value otherwise.

### GetErroredOk

`func (o *WorkflowBulkArchiveResponse) GetErroredOk() (*[]WorkflowBulkError, bool)`

GetErroredOk returns a tuple with the Errored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrored

`func (o *WorkflowBulkArchiveResponse) SetErrored(v []WorkflowBulkError)`

SetErrored sets Errored field to given value.

### HasErrored

`func (o *WorkflowBulkArchiveResponse) HasErrored() bool`

HasErrored returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


