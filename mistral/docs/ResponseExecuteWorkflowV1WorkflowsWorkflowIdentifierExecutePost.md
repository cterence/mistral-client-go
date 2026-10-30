# ResponseExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowName** | **string** | Name of the workflow that was executed | 
**WorkflowId** | Pointer to **string** | The ID of the workflow | [optional] 
**DeploymentName** | Pointer to **string** | The name of the deployment that ran this execution | [optional] 
**ExecutionId** | **string** | ID of the workflow execution | 
**ParentExecutionId** | Pointer to **string** | The parent execution ID of the workflow execution | [optional] 
**RootExecutionId** | **string** | The root execution ID of the workflow execution | 
**RunId** | Pointer to **string** | The unique run identifier (database UUID) | [optional] 
**UserId** | Pointer to **string** | The ID of the user who triggered the execution | [optional] 
**Status** | [**WorkflowExecutionStatus**](WorkflowExecutionStatus.md) | The status of the workflow execution | 
**StartTime** | **time.Time** | The start time of the workflow execution | 
**EndTime** | **time.Time** | The end time of the workflow execution, if available | 
**TotalDurationMs** | Pointer to **int32** | The total duration of the trace in milliseconds | [optional] 
**Result** | **interface{}** | The result of the workflow execution | 
**SearchKeys** | Pointer to **map[string]string** | The execution&#39;s search keys (metadata), if requested via include_search_keys. | [optional] 

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


### GetWorkflowId

`func (o *ResponseExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePost) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *ResponseExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePost) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *ResponseExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePost) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.

### HasWorkflowId

`func (o *ResponseExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePost) HasWorkflowId() bool`

HasWorkflowId returns a boolean if a field has been set.

### GetDeploymentName

`func (o *ResponseExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePost) GetDeploymentName() string`

GetDeploymentName returns the DeploymentName field if non-nil, zero value otherwise.

### GetDeploymentNameOk

`func (o *ResponseExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePost) GetDeploymentNameOk() (*string, bool)`

GetDeploymentNameOk returns a tuple with the DeploymentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentName

`func (o *ResponseExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePost) SetDeploymentName(v string)`

SetDeploymentName sets DeploymentName field to given value.

### HasDeploymentName

`func (o *ResponseExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePost) HasDeploymentName() bool`

HasDeploymentName returns a boolean if a field has been set.

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


### GetRunId

`func (o *ResponseExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePost) GetRunId() string`

GetRunId returns the RunId field if non-nil, zero value otherwise.

### GetRunIdOk

`func (o *ResponseExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePost) GetRunIdOk() (*string, bool)`

GetRunIdOk returns a tuple with the RunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunId

`func (o *ResponseExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePost) SetRunId(v string)`

SetRunId sets RunId field to given value.

### HasRunId

`func (o *ResponseExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePost) HasRunId() bool`

HasRunId returns a boolean if a field has been set.

### GetUserId

`func (o *ResponseExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePost) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ResponseExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePost) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ResponseExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePost) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ResponseExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePost) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

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
### GetSearchKeys

`func (o *ResponseExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePost) GetSearchKeys() map[string]string`

GetSearchKeys returns the SearchKeys field if non-nil, zero value otherwise.

### GetSearchKeysOk

`func (o *ResponseExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePost) GetSearchKeysOk() (*map[string]string, bool)`

GetSearchKeysOk returns a tuple with the SearchKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchKeys

`func (o *ResponseExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePost) SetSearchKeys(v map[string]string)`

SetSearchKeys sets SearchKeys field to given value.

### HasSearchKeys

`func (o *ResponseExecuteWorkflowV1WorkflowsWorkflowIdentifierExecutePost) HasSearchKeys() bool`

HasSearchKeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


