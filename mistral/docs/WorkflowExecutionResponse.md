# WorkflowExecutionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowName** | **string** | The name of the workflow | 
**ExecutionId** | **string** | The ID of the workflow execution | 
**ParentExecutionId** | Pointer to **NullableString** | The parent execution ID of the workflow execution | [optional] 
**RootExecutionId** | **string** | The root execution ID of the workflow execution | 
**Status** | [**NullableWorkflowExecutionStatus**](WorkflowExecutionStatus.md) | The status of the workflow execution | 
**StartTime** | **time.Time** | The start time of the workflow execution | 
**EndTime** | **NullableTime** | The end time of the workflow execution, if available | 
**TotalDurationMs** | Pointer to **NullableInt32** | The total duration of the trace in milliseconds | [optional] 
**Result** | [**NullableAnyOf**](anyOf&lt;&gt;.md) | The result of the workflow execution, if available | 

## Methods

### NewWorkflowExecutionResponse

`func NewWorkflowExecutionResponse(workflowName string, executionId string, rootExecutionId string, status NullableWorkflowExecutionStatus, startTime time.Time, endTime NullableTime, result NullableAnyOf, ) *WorkflowExecutionResponse`

NewWorkflowExecutionResponse instantiates a new WorkflowExecutionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowExecutionResponseWithDefaults

`func NewWorkflowExecutionResponseWithDefaults() *WorkflowExecutionResponse`

NewWorkflowExecutionResponseWithDefaults instantiates a new WorkflowExecutionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowName

`func (o *WorkflowExecutionResponse) GetWorkflowName() string`

GetWorkflowName returns the WorkflowName field if non-nil, zero value otherwise.

### GetWorkflowNameOk

`func (o *WorkflowExecutionResponse) GetWorkflowNameOk() (*string, bool)`

GetWorkflowNameOk returns a tuple with the WorkflowName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowName

`func (o *WorkflowExecutionResponse) SetWorkflowName(v string)`

SetWorkflowName sets WorkflowName field to given value.


### GetExecutionId

`func (o *WorkflowExecutionResponse) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *WorkflowExecutionResponse) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *WorkflowExecutionResponse) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.


### GetParentExecutionId

`func (o *WorkflowExecutionResponse) GetParentExecutionId() string`

GetParentExecutionId returns the ParentExecutionId field if non-nil, zero value otherwise.

### GetParentExecutionIdOk

`func (o *WorkflowExecutionResponse) GetParentExecutionIdOk() (*string, bool)`

GetParentExecutionIdOk returns a tuple with the ParentExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentExecutionId

`func (o *WorkflowExecutionResponse) SetParentExecutionId(v string)`

SetParentExecutionId sets ParentExecutionId field to given value.

### HasParentExecutionId

`func (o *WorkflowExecutionResponse) HasParentExecutionId() bool`

HasParentExecutionId returns a boolean if a field has been set.

### SetParentExecutionIdNil

`func (o *WorkflowExecutionResponse) SetParentExecutionIdNil(b bool)`

 SetParentExecutionIdNil sets the value for ParentExecutionId to be an explicit nil

### UnsetParentExecutionId
`func (o *WorkflowExecutionResponse) UnsetParentExecutionId()`

UnsetParentExecutionId ensures that no value is present for ParentExecutionId, not even an explicit nil
### GetRootExecutionId

`func (o *WorkflowExecutionResponse) GetRootExecutionId() string`

GetRootExecutionId returns the RootExecutionId field if non-nil, zero value otherwise.

### GetRootExecutionIdOk

`func (o *WorkflowExecutionResponse) GetRootExecutionIdOk() (*string, bool)`

GetRootExecutionIdOk returns a tuple with the RootExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootExecutionId

`func (o *WorkflowExecutionResponse) SetRootExecutionId(v string)`

SetRootExecutionId sets RootExecutionId field to given value.


### GetStatus

`func (o *WorkflowExecutionResponse) GetStatus() WorkflowExecutionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkflowExecutionResponse) GetStatusOk() (*WorkflowExecutionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkflowExecutionResponse) SetStatus(v WorkflowExecutionStatus)`

SetStatus sets Status field to given value.


### SetStatusNil

`func (o *WorkflowExecutionResponse) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *WorkflowExecutionResponse) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetStartTime

`func (o *WorkflowExecutionResponse) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *WorkflowExecutionResponse) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *WorkflowExecutionResponse) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.


### GetEndTime

`func (o *WorkflowExecutionResponse) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *WorkflowExecutionResponse) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *WorkflowExecutionResponse) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.


### SetEndTimeNil

`func (o *WorkflowExecutionResponse) SetEndTimeNil(b bool)`

 SetEndTimeNil sets the value for EndTime to be an explicit nil

### UnsetEndTime
`func (o *WorkflowExecutionResponse) UnsetEndTime()`

UnsetEndTime ensures that no value is present for EndTime, not even an explicit nil
### GetTotalDurationMs

`func (o *WorkflowExecutionResponse) GetTotalDurationMs() int32`

GetTotalDurationMs returns the TotalDurationMs field if non-nil, zero value otherwise.

### GetTotalDurationMsOk

`func (o *WorkflowExecutionResponse) GetTotalDurationMsOk() (*int32, bool)`

GetTotalDurationMsOk returns a tuple with the TotalDurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDurationMs

`func (o *WorkflowExecutionResponse) SetTotalDurationMs(v int32)`

SetTotalDurationMs sets TotalDurationMs field to given value.

### HasTotalDurationMs

`func (o *WorkflowExecutionResponse) HasTotalDurationMs() bool`

HasTotalDurationMs returns a boolean if a field has been set.

### SetTotalDurationMsNil

`func (o *WorkflowExecutionResponse) SetTotalDurationMsNil(b bool)`

 SetTotalDurationMsNil sets the value for TotalDurationMs to be an explicit nil

### UnsetTotalDurationMs
`func (o *WorkflowExecutionResponse) UnsetTotalDurationMs()`

UnsetTotalDurationMs ensures that no value is present for TotalDurationMs, not even an explicit nil
### GetResult

`func (o *WorkflowExecutionResponse) GetResult() AnyOf`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *WorkflowExecutionResponse) GetResultOk() (*AnyOf, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *WorkflowExecutionResponse) SetResult(v AnyOf)`

SetResult sets Result field to given value.


### SetResultNil

`func (o *WorkflowExecutionResponse) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *WorkflowExecutionResponse) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


