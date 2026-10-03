# WorkflowListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Workflows** | Pointer to [**[]WorkflowBasicDefinition**](WorkflowBasicDefinition.md) | A list of workflows | [optional] 
**NextCursor** | **NullableString** |  | 

## Methods

### NewWorkflowListResponse

`func NewWorkflowListResponse(nextCursor NullableString, ) *WorkflowListResponse`

NewWorkflowListResponse instantiates a new WorkflowListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowListResponseWithDefaults

`func NewWorkflowListResponseWithDefaults() *WorkflowListResponse`

NewWorkflowListResponseWithDefaults instantiates a new WorkflowListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflows

`func (o *WorkflowListResponse) GetWorkflows() []WorkflowBasicDefinition`

GetWorkflows returns the Workflows field if non-nil, zero value otherwise.

### GetWorkflowsOk

`func (o *WorkflowListResponse) GetWorkflowsOk() (*[]WorkflowBasicDefinition, bool)`

GetWorkflowsOk returns a tuple with the Workflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflows

`func (o *WorkflowListResponse) SetWorkflows(v []WorkflowBasicDefinition)`

SetWorkflows sets Workflows field to given value.

### HasWorkflows

`func (o *WorkflowListResponse) HasWorkflows() bool`

HasWorkflows returns a boolean if a field has been set.

### GetNextCursor

`func (o *WorkflowListResponse) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *WorkflowListResponse) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *WorkflowListResponse) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.


### SetNextCursorNil

`func (o *WorkflowListResponse) SetNextCursorNil(b bool)`

 SetNextCursorNil sets the value for NextCursor to be an explicit nil

### UnsetNextCursor
`func (o *WorkflowListResponse) UnsetNextCursor()`

UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


