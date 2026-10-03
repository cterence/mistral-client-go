# WorkflowExecutionListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Executions** | [**[]WorkflowExecutionWithoutResultResponse**](WorkflowExecutionWithoutResultResponse.md) | A list of workflow executions | 
**NextPageToken** | Pointer to **NullableString** | Token to use for fetching the next page of results. Null if this is the last page. | [optional] 

## Methods

### NewWorkflowExecutionListResponse

`func NewWorkflowExecutionListResponse(executions []WorkflowExecutionWithoutResultResponse, ) *WorkflowExecutionListResponse`

NewWorkflowExecutionListResponse instantiates a new WorkflowExecutionListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowExecutionListResponseWithDefaults

`func NewWorkflowExecutionListResponseWithDefaults() *WorkflowExecutionListResponse`

NewWorkflowExecutionListResponseWithDefaults instantiates a new WorkflowExecutionListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExecutions

`func (o *WorkflowExecutionListResponse) GetExecutions() []WorkflowExecutionWithoutResultResponse`

GetExecutions returns the Executions field if non-nil, zero value otherwise.

### GetExecutionsOk

`func (o *WorkflowExecutionListResponse) GetExecutionsOk() (*[]WorkflowExecutionWithoutResultResponse, bool)`

GetExecutionsOk returns a tuple with the Executions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutions

`func (o *WorkflowExecutionListResponse) SetExecutions(v []WorkflowExecutionWithoutResultResponse)`

SetExecutions sets Executions field to given value.


### GetNextPageToken

`func (o *WorkflowExecutionListResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *WorkflowExecutionListResponse) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *WorkflowExecutionListResponse) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *WorkflowExecutionListResponse) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### SetNextPageTokenNil

`func (o *WorkflowExecutionListResponse) SetNextPageTokenNil(b bool)`

 SetNextPageTokenNil sets the value for NextPageToken to be an explicit nil

### UnsetNextPageToken
`func (o *WorkflowExecutionListResponse) UnsetNextPageToken()`

UnsetNextPageToken ensures that no value is present for NextPageToken, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


