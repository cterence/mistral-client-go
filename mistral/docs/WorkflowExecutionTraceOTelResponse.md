# WorkflowExecutionTraceOTelResponse

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
**DataSource** | **string** | The data source of the trace | 
**OtelTraceId** | Pointer to **NullableString** | The ID of the trace | [optional] 
**OtelTraceData** | Pointer to [**NullableTempoGetTraceResponse**](TempoGetTraceResponse.md) | The raw OpenTelemetry trace data | [optional] 

## Methods

### NewWorkflowExecutionTraceOTelResponse

`func NewWorkflowExecutionTraceOTelResponse(workflowName string, executionId string, rootExecutionId string, status NullableWorkflowExecutionStatus, startTime time.Time, endTime NullableTime, result NullableAnyOf, dataSource string, ) *WorkflowExecutionTraceOTelResponse`

NewWorkflowExecutionTraceOTelResponse instantiates a new WorkflowExecutionTraceOTelResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowExecutionTraceOTelResponseWithDefaults

`func NewWorkflowExecutionTraceOTelResponseWithDefaults() *WorkflowExecutionTraceOTelResponse`

NewWorkflowExecutionTraceOTelResponseWithDefaults instantiates a new WorkflowExecutionTraceOTelResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowName

`func (o *WorkflowExecutionTraceOTelResponse) GetWorkflowName() string`

GetWorkflowName returns the WorkflowName field if non-nil, zero value otherwise.

### GetWorkflowNameOk

`func (o *WorkflowExecutionTraceOTelResponse) GetWorkflowNameOk() (*string, bool)`

GetWorkflowNameOk returns a tuple with the WorkflowName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowName

`func (o *WorkflowExecutionTraceOTelResponse) SetWorkflowName(v string)`

SetWorkflowName sets WorkflowName field to given value.


### GetExecutionId

`func (o *WorkflowExecutionTraceOTelResponse) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *WorkflowExecutionTraceOTelResponse) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *WorkflowExecutionTraceOTelResponse) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.


### GetParentExecutionId

`func (o *WorkflowExecutionTraceOTelResponse) GetParentExecutionId() string`

GetParentExecutionId returns the ParentExecutionId field if non-nil, zero value otherwise.

### GetParentExecutionIdOk

`func (o *WorkflowExecutionTraceOTelResponse) GetParentExecutionIdOk() (*string, bool)`

GetParentExecutionIdOk returns a tuple with the ParentExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentExecutionId

`func (o *WorkflowExecutionTraceOTelResponse) SetParentExecutionId(v string)`

SetParentExecutionId sets ParentExecutionId field to given value.

### HasParentExecutionId

`func (o *WorkflowExecutionTraceOTelResponse) HasParentExecutionId() bool`

HasParentExecutionId returns a boolean if a field has been set.

### SetParentExecutionIdNil

`func (o *WorkflowExecutionTraceOTelResponse) SetParentExecutionIdNil(b bool)`

 SetParentExecutionIdNil sets the value for ParentExecutionId to be an explicit nil

### UnsetParentExecutionId
`func (o *WorkflowExecutionTraceOTelResponse) UnsetParentExecutionId()`

UnsetParentExecutionId ensures that no value is present for ParentExecutionId, not even an explicit nil
### GetRootExecutionId

`func (o *WorkflowExecutionTraceOTelResponse) GetRootExecutionId() string`

GetRootExecutionId returns the RootExecutionId field if non-nil, zero value otherwise.

### GetRootExecutionIdOk

`func (o *WorkflowExecutionTraceOTelResponse) GetRootExecutionIdOk() (*string, bool)`

GetRootExecutionIdOk returns a tuple with the RootExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootExecutionId

`func (o *WorkflowExecutionTraceOTelResponse) SetRootExecutionId(v string)`

SetRootExecutionId sets RootExecutionId field to given value.


### GetStatus

`func (o *WorkflowExecutionTraceOTelResponse) GetStatus() WorkflowExecutionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkflowExecutionTraceOTelResponse) GetStatusOk() (*WorkflowExecutionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkflowExecutionTraceOTelResponse) SetStatus(v WorkflowExecutionStatus)`

SetStatus sets Status field to given value.


### SetStatusNil

`func (o *WorkflowExecutionTraceOTelResponse) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *WorkflowExecutionTraceOTelResponse) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetStartTime

`func (o *WorkflowExecutionTraceOTelResponse) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *WorkflowExecutionTraceOTelResponse) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *WorkflowExecutionTraceOTelResponse) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.


### GetEndTime

`func (o *WorkflowExecutionTraceOTelResponse) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *WorkflowExecutionTraceOTelResponse) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *WorkflowExecutionTraceOTelResponse) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.


### SetEndTimeNil

`func (o *WorkflowExecutionTraceOTelResponse) SetEndTimeNil(b bool)`

 SetEndTimeNil sets the value for EndTime to be an explicit nil

### UnsetEndTime
`func (o *WorkflowExecutionTraceOTelResponse) UnsetEndTime()`

UnsetEndTime ensures that no value is present for EndTime, not even an explicit nil
### GetTotalDurationMs

`func (o *WorkflowExecutionTraceOTelResponse) GetTotalDurationMs() int32`

GetTotalDurationMs returns the TotalDurationMs field if non-nil, zero value otherwise.

### GetTotalDurationMsOk

`func (o *WorkflowExecutionTraceOTelResponse) GetTotalDurationMsOk() (*int32, bool)`

GetTotalDurationMsOk returns a tuple with the TotalDurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDurationMs

`func (o *WorkflowExecutionTraceOTelResponse) SetTotalDurationMs(v int32)`

SetTotalDurationMs sets TotalDurationMs field to given value.

### HasTotalDurationMs

`func (o *WorkflowExecutionTraceOTelResponse) HasTotalDurationMs() bool`

HasTotalDurationMs returns a boolean if a field has been set.

### SetTotalDurationMsNil

`func (o *WorkflowExecutionTraceOTelResponse) SetTotalDurationMsNil(b bool)`

 SetTotalDurationMsNil sets the value for TotalDurationMs to be an explicit nil

### UnsetTotalDurationMs
`func (o *WorkflowExecutionTraceOTelResponse) UnsetTotalDurationMs()`

UnsetTotalDurationMs ensures that no value is present for TotalDurationMs, not even an explicit nil
### GetResult

`func (o *WorkflowExecutionTraceOTelResponse) GetResult() AnyOf`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *WorkflowExecutionTraceOTelResponse) GetResultOk() (*AnyOf, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *WorkflowExecutionTraceOTelResponse) SetResult(v AnyOf)`

SetResult sets Result field to given value.


### SetResultNil

`func (o *WorkflowExecutionTraceOTelResponse) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *WorkflowExecutionTraceOTelResponse) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetDataSource

`func (o *WorkflowExecutionTraceOTelResponse) GetDataSource() string`

GetDataSource returns the DataSource field if non-nil, zero value otherwise.

### GetDataSourceOk

`func (o *WorkflowExecutionTraceOTelResponse) GetDataSourceOk() (*string, bool)`

GetDataSourceOk returns a tuple with the DataSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSource

`func (o *WorkflowExecutionTraceOTelResponse) SetDataSource(v string)`

SetDataSource sets DataSource field to given value.


### GetOtelTraceId

`func (o *WorkflowExecutionTraceOTelResponse) GetOtelTraceId() string`

GetOtelTraceId returns the OtelTraceId field if non-nil, zero value otherwise.

### GetOtelTraceIdOk

`func (o *WorkflowExecutionTraceOTelResponse) GetOtelTraceIdOk() (*string, bool)`

GetOtelTraceIdOk returns a tuple with the OtelTraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtelTraceId

`func (o *WorkflowExecutionTraceOTelResponse) SetOtelTraceId(v string)`

SetOtelTraceId sets OtelTraceId field to given value.

### HasOtelTraceId

`func (o *WorkflowExecutionTraceOTelResponse) HasOtelTraceId() bool`

HasOtelTraceId returns a boolean if a field has been set.

### SetOtelTraceIdNil

`func (o *WorkflowExecutionTraceOTelResponse) SetOtelTraceIdNil(b bool)`

 SetOtelTraceIdNil sets the value for OtelTraceId to be an explicit nil

### UnsetOtelTraceId
`func (o *WorkflowExecutionTraceOTelResponse) UnsetOtelTraceId()`

UnsetOtelTraceId ensures that no value is present for OtelTraceId, not even an explicit nil
### GetOtelTraceData

`func (o *WorkflowExecutionTraceOTelResponse) GetOtelTraceData() TempoGetTraceResponse`

GetOtelTraceData returns the OtelTraceData field if non-nil, zero value otherwise.

### GetOtelTraceDataOk

`func (o *WorkflowExecutionTraceOTelResponse) GetOtelTraceDataOk() (*TempoGetTraceResponse, bool)`

GetOtelTraceDataOk returns a tuple with the OtelTraceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtelTraceData

`func (o *WorkflowExecutionTraceOTelResponse) SetOtelTraceData(v TempoGetTraceResponse)`

SetOtelTraceData sets OtelTraceData field to given value.

### HasOtelTraceData

`func (o *WorkflowExecutionTraceOTelResponse) HasOtelTraceData() bool`

HasOtelTraceData returns a boolean if a field has been set.

### SetOtelTraceDataNil

`func (o *WorkflowExecutionTraceOTelResponse) SetOtelTraceDataNil(b bool)`

 SetOtelTraceDataNil sets the value for OtelTraceData to be an explicit nil

### UnsetOtelTraceData
`func (o *WorkflowExecutionTraceOTelResponse) UnsetOtelTraceData()`

UnsetOtelTraceData ensures that no value is present for OtelTraceData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


