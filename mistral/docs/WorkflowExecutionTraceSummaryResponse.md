# WorkflowExecutionTraceSummaryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowName** | **string** | The name of the workflow | 
**WorkflowId** | Pointer to **NullableString** | The ID of the workflow | [optional] 
**DeploymentName** | Pointer to **NullableString** | The name of the deployment that ran this execution | [optional] 
**ExecutionId** | **string** | The ID of the workflow execution | 
**ParentExecutionId** | Pointer to **NullableString** | The parent execution ID of the workflow execution | [optional] 
**RootExecutionId** | **string** | The root execution ID of the workflow execution | 
**RunId** | Pointer to **NullableString** | The unique run identifier (database UUID) | [optional] 
**UserId** | Pointer to **NullableString** | The ID of the user who triggered the execution | [optional] 
**Status** | [**NullableWorkflowExecutionStatus**](WorkflowExecutionStatus.md) | The status of the workflow execution | 
**StartTime** | **time.Time** | The start time of the workflow execution | 
**EndTime** | **NullableTime** | The end time of the workflow execution, if available | 
**TotalDurationMs** | Pointer to **NullableInt32** | The total duration of the trace in milliseconds | [optional] 
**Result** | [**NullableAnyOf**](anyOf&lt;&gt;.md) | The result of the workflow execution, if available | 
**SearchKeys** | Pointer to **map[string]string** | The execution&#39;s search keys (metadata), if requested via include_search_keys. | [optional] 
**SpanTree** | Pointer to [**NullableWorkflowExecutionTraceSummarySpan**](WorkflowExecutionTraceSummarySpan.md) | The root span of the trace | [optional] 

## Methods

### NewWorkflowExecutionTraceSummaryResponse

`func NewWorkflowExecutionTraceSummaryResponse(workflowName string, executionId string, rootExecutionId string, status NullableWorkflowExecutionStatus, startTime time.Time, endTime NullableTime, result NullableAnyOf, ) *WorkflowExecutionTraceSummaryResponse`

NewWorkflowExecutionTraceSummaryResponse instantiates a new WorkflowExecutionTraceSummaryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowExecutionTraceSummaryResponseWithDefaults

`func NewWorkflowExecutionTraceSummaryResponseWithDefaults() *WorkflowExecutionTraceSummaryResponse`

NewWorkflowExecutionTraceSummaryResponseWithDefaults instantiates a new WorkflowExecutionTraceSummaryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowName

`func (o *WorkflowExecutionTraceSummaryResponse) GetWorkflowName() string`

GetWorkflowName returns the WorkflowName field if non-nil, zero value otherwise.

### GetWorkflowNameOk

`func (o *WorkflowExecutionTraceSummaryResponse) GetWorkflowNameOk() (*string, bool)`

GetWorkflowNameOk returns a tuple with the WorkflowName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowName

`func (o *WorkflowExecutionTraceSummaryResponse) SetWorkflowName(v string)`

SetWorkflowName sets WorkflowName field to given value.


### GetWorkflowId

`func (o *WorkflowExecutionTraceSummaryResponse) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *WorkflowExecutionTraceSummaryResponse) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *WorkflowExecutionTraceSummaryResponse) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.

### HasWorkflowId

`func (o *WorkflowExecutionTraceSummaryResponse) HasWorkflowId() bool`

HasWorkflowId returns a boolean if a field has been set.

### SetWorkflowIdNil

`func (o *WorkflowExecutionTraceSummaryResponse) SetWorkflowIdNil(b bool)`

 SetWorkflowIdNil sets the value for WorkflowId to be an explicit nil

### UnsetWorkflowId
`func (o *WorkflowExecutionTraceSummaryResponse) UnsetWorkflowId()`

UnsetWorkflowId ensures that no value is present for WorkflowId, not even an explicit nil
### GetDeploymentName

`func (o *WorkflowExecutionTraceSummaryResponse) GetDeploymentName() string`

GetDeploymentName returns the DeploymentName field if non-nil, zero value otherwise.

### GetDeploymentNameOk

`func (o *WorkflowExecutionTraceSummaryResponse) GetDeploymentNameOk() (*string, bool)`

GetDeploymentNameOk returns a tuple with the DeploymentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentName

`func (o *WorkflowExecutionTraceSummaryResponse) SetDeploymentName(v string)`

SetDeploymentName sets DeploymentName field to given value.

### HasDeploymentName

`func (o *WorkflowExecutionTraceSummaryResponse) HasDeploymentName() bool`

HasDeploymentName returns a boolean if a field has been set.

### SetDeploymentNameNil

`func (o *WorkflowExecutionTraceSummaryResponse) SetDeploymentNameNil(b bool)`

 SetDeploymentNameNil sets the value for DeploymentName to be an explicit nil

### UnsetDeploymentName
`func (o *WorkflowExecutionTraceSummaryResponse) UnsetDeploymentName()`

UnsetDeploymentName ensures that no value is present for DeploymentName, not even an explicit nil
### GetExecutionId

`func (o *WorkflowExecutionTraceSummaryResponse) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *WorkflowExecutionTraceSummaryResponse) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *WorkflowExecutionTraceSummaryResponse) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.


### GetParentExecutionId

`func (o *WorkflowExecutionTraceSummaryResponse) GetParentExecutionId() string`

GetParentExecutionId returns the ParentExecutionId field if non-nil, zero value otherwise.

### GetParentExecutionIdOk

`func (o *WorkflowExecutionTraceSummaryResponse) GetParentExecutionIdOk() (*string, bool)`

GetParentExecutionIdOk returns a tuple with the ParentExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentExecutionId

`func (o *WorkflowExecutionTraceSummaryResponse) SetParentExecutionId(v string)`

SetParentExecutionId sets ParentExecutionId field to given value.

### HasParentExecutionId

`func (o *WorkflowExecutionTraceSummaryResponse) HasParentExecutionId() bool`

HasParentExecutionId returns a boolean if a field has been set.

### SetParentExecutionIdNil

`func (o *WorkflowExecutionTraceSummaryResponse) SetParentExecutionIdNil(b bool)`

 SetParentExecutionIdNil sets the value for ParentExecutionId to be an explicit nil

### UnsetParentExecutionId
`func (o *WorkflowExecutionTraceSummaryResponse) UnsetParentExecutionId()`

UnsetParentExecutionId ensures that no value is present for ParentExecutionId, not even an explicit nil
### GetRootExecutionId

`func (o *WorkflowExecutionTraceSummaryResponse) GetRootExecutionId() string`

GetRootExecutionId returns the RootExecutionId field if non-nil, zero value otherwise.

### GetRootExecutionIdOk

`func (o *WorkflowExecutionTraceSummaryResponse) GetRootExecutionIdOk() (*string, bool)`

GetRootExecutionIdOk returns a tuple with the RootExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootExecutionId

`func (o *WorkflowExecutionTraceSummaryResponse) SetRootExecutionId(v string)`

SetRootExecutionId sets RootExecutionId field to given value.


### GetRunId

`func (o *WorkflowExecutionTraceSummaryResponse) GetRunId() string`

GetRunId returns the RunId field if non-nil, zero value otherwise.

### GetRunIdOk

`func (o *WorkflowExecutionTraceSummaryResponse) GetRunIdOk() (*string, bool)`

GetRunIdOk returns a tuple with the RunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunId

`func (o *WorkflowExecutionTraceSummaryResponse) SetRunId(v string)`

SetRunId sets RunId field to given value.

### HasRunId

`func (o *WorkflowExecutionTraceSummaryResponse) HasRunId() bool`

HasRunId returns a boolean if a field has been set.

### SetRunIdNil

`func (o *WorkflowExecutionTraceSummaryResponse) SetRunIdNil(b bool)`

 SetRunIdNil sets the value for RunId to be an explicit nil

### UnsetRunId
`func (o *WorkflowExecutionTraceSummaryResponse) UnsetRunId()`

UnsetRunId ensures that no value is present for RunId, not even an explicit nil
### GetUserId

`func (o *WorkflowExecutionTraceSummaryResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *WorkflowExecutionTraceSummaryResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *WorkflowExecutionTraceSummaryResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *WorkflowExecutionTraceSummaryResponse) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *WorkflowExecutionTraceSummaryResponse) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *WorkflowExecutionTraceSummaryResponse) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetStatus

`func (o *WorkflowExecutionTraceSummaryResponse) GetStatus() WorkflowExecutionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkflowExecutionTraceSummaryResponse) GetStatusOk() (*WorkflowExecutionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkflowExecutionTraceSummaryResponse) SetStatus(v WorkflowExecutionStatus)`

SetStatus sets Status field to given value.


### SetStatusNil

`func (o *WorkflowExecutionTraceSummaryResponse) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *WorkflowExecutionTraceSummaryResponse) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetStartTime

`func (o *WorkflowExecutionTraceSummaryResponse) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *WorkflowExecutionTraceSummaryResponse) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *WorkflowExecutionTraceSummaryResponse) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.


### GetEndTime

`func (o *WorkflowExecutionTraceSummaryResponse) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *WorkflowExecutionTraceSummaryResponse) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *WorkflowExecutionTraceSummaryResponse) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.


### SetEndTimeNil

`func (o *WorkflowExecutionTraceSummaryResponse) SetEndTimeNil(b bool)`

 SetEndTimeNil sets the value for EndTime to be an explicit nil

### UnsetEndTime
`func (o *WorkflowExecutionTraceSummaryResponse) UnsetEndTime()`

UnsetEndTime ensures that no value is present for EndTime, not even an explicit nil
### GetTotalDurationMs

`func (o *WorkflowExecutionTraceSummaryResponse) GetTotalDurationMs() int32`

GetTotalDurationMs returns the TotalDurationMs field if non-nil, zero value otherwise.

### GetTotalDurationMsOk

`func (o *WorkflowExecutionTraceSummaryResponse) GetTotalDurationMsOk() (*int32, bool)`

GetTotalDurationMsOk returns a tuple with the TotalDurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDurationMs

`func (o *WorkflowExecutionTraceSummaryResponse) SetTotalDurationMs(v int32)`

SetTotalDurationMs sets TotalDurationMs field to given value.

### HasTotalDurationMs

`func (o *WorkflowExecutionTraceSummaryResponse) HasTotalDurationMs() bool`

HasTotalDurationMs returns a boolean if a field has been set.

### SetTotalDurationMsNil

`func (o *WorkflowExecutionTraceSummaryResponse) SetTotalDurationMsNil(b bool)`

 SetTotalDurationMsNil sets the value for TotalDurationMs to be an explicit nil

### UnsetTotalDurationMs
`func (o *WorkflowExecutionTraceSummaryResponse) UnsetTotalDurationMs()`

UnsetTotalDurationMs ensures that no value is present for TotalDurationMs, not even an explicit nil
### GetResult

`func (o *WorkflowExecutionTraceSummaryResponse) GetResult() AnyOf`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *WorkflowExecutionTraceSummaryResponse) GetResultOk() (*AnyOf, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *WorkflowExecutionTraceSummaryResponse) SetResult(v AnyOf)`

SetResult sets Result field to given value.


### SetResultNil

`func (o *WorkflowExecutionTraceSummaryResponse) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *WorkflowExecutionTraceSummaryResponse) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetSearchKeys

`func (o *WorkflowExecutionTraceSummaryResponse) GetSearchKeys() map[string]*string`

GetSearchKeys returns the SearchKeys field if non-nil, zero value otherwise.

### GetSearchKeysOk

`func (o *WorkflowExecutionTraceSummaryResponse) GetSearchKeysOk() (*map[string]*string, bool)`

GetSearchKeysOk returns a tuple with the SearchKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchKeys

`func (o *WorkflowExecutionTraceSummaryResponse) SetSearchKeys(v map[string]*string)`

SetSearchKeys sets SearchKeys field to given value.

### HasSearchKeys

`func (o *WorkflowExecutionTraceSummaryResponse) HasSearchKeys() bool`

HasSearchKeys returns a boolean if a field has been set.

### SetSearchKeysNil

`func (o *WorkflowExecutionTraceSummaryResponse) SetSearchKeysNil(b bool)`

 SetSearchKeysNil sets the value for SearchKeys to be an explicit nil

### UnsetSearchKeys
`func (o *WorkflowExecutionTraceSummaryResponse) UnsetSearchKeys()`

UnsetSearchKeys ensures that no value is present for SearchKeys, not even an explicit nil
### GetSpanTree

`func (o *WorkflowExecutionTraceSummaryResponse) GetSpanTree() WorkflowExecutionTraceSummarySpan`

GetSpanTree returns the SpanTree field if non-nil, zero value otherwise.

### GetSpanTreeOk

`func (o *WorkflowExecutionTraceSummaryResponse) GetSpanTreeOk() (*WorkflowExecutionTraceSummarySpan, bool)`

GetSpanTreeOk returns a tuple with the SpanTree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanTree

`func (o *WorkflowExecutionTraceSummaryResponse) SetSpanTree(v WorkflowExecutionTraceSummarySpan)`

SetSpanTree sets SpanTree field to given value.

### HasSpanTree

`func (o *WorkflowExecutionTraceSummaryResponse) HasSpanTree() bool`

HasSpanTree returns a boolean if a field has been set.

### SetSpanTreeNil

`func (o *WorkflowExecutionTraceSummaryResponse) SetSpanTreeNil(b bool)`

 SetSpanTreeNil sets the value for SpanTree to be an explicit nil

### UnsetSpanTree
`func (o *WorkflowExecutionTraceSummaryResponse) UnsetSpanTree()`

UnsetSpanTree ensures that no value is present for SpanTree, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


