# ResponseExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowName** | **string** | Name of the workflow that was executed | 
**ExecutionId** | **string** | ID of the workflow execution | 
**ParentExecutionId** | Pointer to **string** | The parent execution ID of the workflow execution | [optional] 
**RootExecutionId** | **string** | The root execution ID of the workflow execution | 
**Status** | [**WorkflowExecutionStatus**](WorkflowExecutionStatus.md) | The status of the workflow execution | 
**StartTime** | **time.Time** | The start time of the workflow execution | 
**EndTime** | **time.Time** | The end time of the workflow execution, if available | 
**TotalDurationMs** | Pointer to **int32** | The total duration of the trace in milliseconds | [optional] 
**Result** | **interface{}** | The result of the workflow execution | 

## Methods

### NewResponseExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePost

`func NewResponseExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePost(workflowName string, executionId string, rootExecutionId string, status WorkflowExecutionStatus, startTime time.Time, endTime time.Time, result interface{}, ) *ResponseExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePost`

NewResponseExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePost instantiates a new ResponseExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePostWithDefaults

`func NewResponseExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePostWithDefaults() *ResponseExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePost`

NewResponseExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePostWithDefaults instantiates a new ResponseExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowName

`func (o *ResponseExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePost) GetWorkflowName() string`

GetWorkflowName returns the WorkflowName field if non-nil, zero value otherwise.

### GetWorkflowNameOk

`func (o *ResponseExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePost) GetWorkflowNameOk() (*string, bool)`

GetWorkflowNameOk returns a tuple with the WorkflowName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowName

`func (o *ResponseExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePost) SetWorkflowName(v string)`

SetWorkflowName sets WorkflowName field to given value.


### GetExecutionId

`func (o *ResponseExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePost) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *ResponseExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePost) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *ResponseExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePost) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.


### GetParentExecutionId

`func (o *ResponseExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePost) GetParentExecutionId() string`

GetParentExecutionId returns the ParentExecutionId field if non-nil, zero value otherwise.

### GetParentExecutionIdOk

`func (o *ResponseExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePost) GetParentExecutionIdOk() (*string, bool)`

GetParentExecutionIdOk returns a tuple with the ParentExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentExecutionId

`func (o *ResponseExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePost) SetParentExecutionId(v string)`

SetParentExecutionId sets ParentExecutionId field to given value.

### HasParentExecutionId

`func (o *ResponseExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePost) HasParentExecutionId() bool`

HasParentExecutionId returns a boolean if a field has been set.

### GetRootExecutionId

`func (o *ResponseExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePost) GetRootExecutionId() string`

GetRootExecutionId returns the RootExecutionId field if non-nil, zero value otherwise.

### GetRootExecutionIdOk

`func (o *ResponseExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePost) GetRootExecutionIdOk() (*string, bool)`

GetRootExecutionIdOk returns a tuple with the RootExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootExecutionId

`func (o *ResponseExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePost) SetRootExecutionId(v string)`

SetRootExecutionId sets RootExecutionId field to given value.


### GetStatus

`func (o *ResponseExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePost) GetStatus() WorkflowExecutionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResponseExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePost) GetStatusOk() (*WorkflowExecutionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResponseExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePost) SetStatus(v WorkflowExecutionStatus)`

SetStatus sets Status field to given value.


### GetStartTime

`func (o *ResponseExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePost) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ResponseExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePost) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ResponseExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePost) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.


### GetEndTime

`func (o *ResponseExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePost) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *ResponseExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePost) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *ResponseExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePost) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.


### GetTotalDurationMs

`func (o *ResponseExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePost) GetTotalDurationMs() int32`

GetTotalDurationMs returns the TotalDurationMs field if non-nil, zero value otherwise.

### GetTotalDurationMsOk

`func (o *ResponseExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePost) GetTotalDurationMsOk() (*int32, bool)`

GetTotalDurationMsOk returns a tuple with the TotalDurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDurationMs

`func (o *ResponseExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePost) SetTotalDurationMs(v int32)`

SetTotalDurationMs sets TotalDurationMs field to given value.

### HasTotalDurationMs

`func (o *ResponseExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePost) HasTotalDurationMs() bool`

HasTotalDurationMs returns a boolean if a field has been set.

### GetResult

`func (o *ResponseExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePost) GetResult() interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ResponseExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePost) GetResultOk() (*interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ResponseExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePost) SetResult(v interface{})`

SetResult sets Result field to given value.


### SetResultNil

`func (o *ResponseExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePost) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *ResponseExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePost) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


