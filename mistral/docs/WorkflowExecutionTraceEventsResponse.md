# WorkflowExecutionTraceEventsResponse

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
**Events** | Pointer to [**[]EventsInner1**](EventsInner1.md) | The events of the workflow execution | [optional] 

## Methods

### NewWorkflowExecutionTraceEventsResponse

`func NewWorkflowExecutionTraceEventsResponse(workflowName string, executionId string, rootExecutionId string, status NullableWorkflowExecutionStatus, startTime time.Time, endTime NullableTime, result NullableAnyOf, ) *WorkflowExecutionTraceEventsResponse`

NewWorkflowExecutionTraceEventsResponse instantiates a new WorkflowExecutionTraceEventsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowExecutionTraceEventsResponseWithDefaults

`func NewWorkflowExecutionTraceEventsResponseWithDefaults() *WorkflowExecutionTraceEventsResponse`

NewWorkflowExecutionTraceEventsResponseWithDefaults instantiates a new WorkflowExecutionTraceEventsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowName

`func (o *WorkflowExecutionTraceEventsResponse) GetWorkflowName() string`

GetWorkflowName returns the WorkflowName field if non-nil, zero value otherwise.

### GetWorkflowNameOk

`func (o *WorkflowExecutionTraceEventsResponse) GetWorkflowNameOk() (*string, bool)`

GetWorkflowNameOk returns a tuple with the WorkflowName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowName

`func (o *WorkflowExecutionTraceEventsResponse) SetWorkflowName(v string)`

SetWorkflowName sets WorkflowName field to given value.


### GetExecutionId

`func (o *WorkflowExecutionTraceEventsResponse) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *WorkflowExecutionTraceEventsResponse) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *WorkflowExecutionTraceEventsResponse) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.


### GetParentExecutionId

`func (o *WorkflowExecutionTraceEventsResponse) GetParentExecutionId() string`

GetParentExecutionId returns the ParentExecutionId field if non-nil, zero value otherwise.

### GetParentExecutionIdOk

`func (o *WorkflowExecutionTraceEventsResponse) GetParentExecutionIdOk() (*string, bool)`

GetParentExecutionIdOk returns a tuple with the ParentExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentExecutionId

`func (o *WorkflowExecutionTraceEventsResponse) SetParentExecutionId(v string)`

SetParentExecutionId sets ParentExecutionId field to given value.

### HasParentExecutionId

`func (o *WorkflowExecutionTraceEventsResponse) HasParentExecutionId() bool`

HasParentExecutionId returns a boolean if a field has been set.

### SetParentExecutionIdNil

`func (o *WorkflowExecutionTraceEventsResponse) SetParentExecutionIdNil(b bool)`

 SetParentExecutionIdNil sets the value for ParentExecutionId to be an explicit nil

### UnsetParentExecutionId
`func (o *WorkflowExecutionTraceEventsResponse) UnsetParentExecutionId()`

UnsetParentExecutionId ensures that no value is present for ParentExecutionId, not even an explicit nil
### GetRootExecutionId

`func (o *WorkflowExecutionTraceEventsResponse) GetRootExecutionId() string`

GetRootExecutionId returns the RootExecutionId field if non-nil, zero value otherwise.

### GetRootExecutionIdOk

`func (o *WorkflowExecutionTraceEventsResponse) GetRootExecutionIdOk() (*string, bool)`

GetRootExecutionIdOk returns a tuple with the RootExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootExecutionId

`func (o *WorkflowExecutionTraceEventsResponse) SetRootExecutionId(v string)`

SetRootExecutionId sets RootExecutionId field to given value.


### GetStatus

`func (o *WorkflowExecutionTraceEventsResponse) GetStatus() WorkflowExecutionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkflowExecutionTraceEventsResponse) GetStatusOk() (*WorkflowExecutionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkflowExecutionTraceEventsResponse) SetStatus(v WorkflowExecutionStatus)`

SetStatus sets Status field to given value.


### SetStatusNil

`func (o *WorkflowExecutionTraceEventsResponse) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *WorkflowExecutionTraceEventsResponse) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetStartTime

`func (o *WorkflowExecutionTraceEventsResponse) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *WorkflowExecutionTraceEventsResponse) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *WorkflowExecutionTraceEventsResponse) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.


### GetEndTime

`func (o *WorkflowExecutionTraceEventsResponse) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *WorkflowExecutionTraceEventsResponse) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *WorkflowExecutionTraceEventsResponse) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.


### SetEndTimeNil

`func (o *WorkflowExecutionTraceEventsResponse) SetEndTimeNil(b bool)`

 SetEndTimeNil sets the value for EndTime to be an explicit nil

### UnsetEndTime
`func (o *WorkflowExecutionTraceEventsResponse) UnsetEndTime()`

UnsetEndTime ensures that no value is present for EndTime, not even an explicit nil
### GetTotalDurationMs

`func (o *WorkflowExecutionTraceEventsResponse) GetTotalDurationMs() int32`

GetTotalDurationMs returns the TotalDurationMs field if non-nil, zero value otherwise.

### GetTotalDurationMsOk

`func (o *WorkflowExecutionTraceEventsResponse) GetTotalDurationMsOk() (*int32, bool)`

GetTotalDurationMsOk returns a tuple with the TotalDurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDurationMs

`func (o *WorkflowExecutionTraceEventsResponse) SetTotalDurationMs(v int32)`

SetTotalDurationMs sets TotalDurationMs field to given value.

### HasTotalDurationMs

`func (o *WorkflowExecutionTraceEventsResponse) HasTotalDurationMs() bool`

HasTotalDurationMs returns a boolean if a field has been set.

### SetTotalDurationMsNil

`func (o *WorkflowExecutionTraceEventsResponse) SetTotalDurationMsNil(b bool)`

 SetTotalDurationMsNil sets the value for TotalDurationMs to be an explicit nil

### UnsetTotalDurationMs
`func (o *WorkflowExecutionTraceEventsResponse) UnsetTotalDurationMs()`

UnsetTotalDurationMs ensures that no value is present for TotalDurationMs, not even an explicit nil
### GetResult

`func (o *WorkflowExecutionTraceEventsResponse) GetResult() AnyOf`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *WorkflowExecutionTraceEventsResponse) GetResultOk() (*AnyOf, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *WorkflowExecutionTraceEventsResponse) SetResult(v AnyOf)`

SetResult sets Result field to given value.


### SetResultNil

`func (o *WorkflowExecutionTraceEventsResponse) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *WorkflowExecutionTraceEventsResponse) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetEvents

`func (o *WorkflowExecutionTraceEventsResponse) GetEvents() []EventsInner1`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *WorkflowExecutionTraceEventsResponse) GetEventsOk() (*[]EventsInner1, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *WorkflowExecutionTraceEventsResponse) SetEvents(v []EventsInner1)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *WorkflowExecutionTraceEventsResponse) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


