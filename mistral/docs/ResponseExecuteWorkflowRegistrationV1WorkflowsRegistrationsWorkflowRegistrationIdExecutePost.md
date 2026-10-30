# ResponseExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePost

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

### NewResponseExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePost

`func NewResponseExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePost(workflowName string, executionId string, rootExecutionId string, status WorkflowExecutionStatus, startTime time.Time, endTime time.Time, result interface{}, ) *ResponseExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePost`

NewResponseExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePost instantiates a new ResponseExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePostWithDefaults

`func NewResponseExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePostWithDefaults() *ResponseExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePost`

NewResponseExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePostWithDefaults instantiates a new ResponseExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowName

`func (o *ResponseExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePost) GetWorkflowName() string`

GetWorkflowName returns the WorkflowName field if non-nil, zero value otherwise.

### GetWorkflowNameOk

`func (o *ResponseExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePost) GetWorkflowNameOk() (*string, bool)`

GetWorkflowNameOk returns a tuple with the WorkflowName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowName

`func (o *ResponseExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePost) SetWorkflowName(v string)`

SetWorkflowName sets WorkflowName field to given value.


### GetWorkflowId

`func (o *ResponseExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePost) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *ResponseExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePost) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *ResponseExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePost) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.

### HasWorkflowId

`func (o *ResponseExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePost) HasWorkflowId() bool`

HasWorkflowId returns a boolean if a field has been set.

### GetDeploymentName

`func (o *ResponseExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePost) GetDeploymentName() string`

GetDeploymentName returns the DeploymentName field if non-nil, zero value otherwise.

### GetDeploymentNameOk

`func (o *ResponseExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePost) GetDeploymentNameOk() (*string, bool)`

GetDeploymentNameOk returns a tuple with the DeploymentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentName

`func (o *ResponseExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePost) SetDeploymentName(v string)`

SetDeploymentName sets DeploymentName field to given value.

### HasDeploymentName

`func (o *ResponseExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePost) HasDeploymentName() bool`

HasDeploymentName returns a boolean if a field has been set.

### GetExecutionId

`func (o *ResponseExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePost) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *ResponseExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePost) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *ResponseExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePost) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.


### GetParentExecutionId

`func (o *ResponseExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePost) GetParentExecutionId() string`

GetParentExecutionId returns the ParentExecutionId field if non-nil, zero value otherwise.

### GetParentExecutionIdOk

`func (o *ResponseExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePost) GetParentExecutionIdOk() (*string, bool)`

GetParentExecutionIdOk returns a tuple with the ParentExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentExecutionId

`func (o *ResponseExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePost) SetParentExecutionId(v string)`

SetParentExecutionId sets ParentExecutionId field to given value.

### HasParentExecutionId

`func (o *ResponseExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePost) HasParentExecutionId() bool`

HasParentExecutionId returns a boolean if a field has been set.

### GetRootExecutionId

`func (o *ResponseExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePost) GetRootExecutionId() string`

GetRootExecutionId returns the RootExecutionId field if non-nil, zero value otherwise.

### GetRootExecutionIdOk

`func (o *ResponseExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePost) GetRootExecutionIdOk() (*string, bool)`

GetRootExecutionIdOk returns a tuple with the RootExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootExecutionId

`func (o *ResponseExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePost) SetRootExecutionId(v string)`

SetRootExecutionId sets RootExecutionId field to given value.


### GetRunId

`func (o *ResponseExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePost) GetRunId() string`

GetRunId returns the RunId field if non-nil, zero value otherwise.

### GetRunIdOk

`func (o *ResponseExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePost) GetRunIdOk() (*string, bool)`

GetRunIdOk returns a tuple with the RunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunId

`func (o *ResponseExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePost) SetRunId(v string)`

SetRunId sets RunId field to given value.

### HasRunId

`func (o *ResponseExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePost) HasRunId() bool`

HasRunId returns a boolean if a field has been set.

### GetUserId

`func (o *ResponseExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePost) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ResponseExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePost) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ResponseExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePost) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ResponseExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePost) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetStatus

`func (o *ResponseExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePost) GetStatus() WorkflowExecutionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResponseExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePost) GetStatusOk() (*WorkflowExecutionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResponseExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePost) SetStatus(v WorkflowExecutionStatus)`

SetStatus sets Status field to given value.


### GetStartTime

`func (o *ResponseExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePost) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ResponseExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePost) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ResponseExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePost) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.


### GetEndTime

`func (o *ResponseExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePost) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *ResponseExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePost) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *ResponseExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePost) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.


### GetTotalDurationMs

`func (o *ResponseExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePost) GetTotalDurationMs() int32`

GetTotalDurationMs returns the TotalDurationMs field if non-nil, zero value otherwise.

### GetTotalDurationMsOk

`func (o *ResponseExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePost) GetTotalDurationMsOk() (*int32, bool)`

GetTotalDurationMsOk returns a tuple with the TotalDurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDurationMs

`func (o *ResponseExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePost) SetTotalDurationMs(v int32)`

SetTotalDurationMs sets TotalDurationMs field to given value.

### HasTotalDurationMs

`func (o *ResponseExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePost) HasTotalDurationMs() bool`

HasTotalDurationMs returns a boolean if a field has been set.

### GetResult

`func (o *ResponseExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePost) GetResult() interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ResponseExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePost) GetResultOk() (*interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ResponseExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePost) SetResult(v interface{})`

SetResult sets Result field to given value.


### SetResultNil

`func (o *ResponseExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePost) SetResultNil(b bool)`

 SetResultNil sets the value for Result to be an explicit nil

### UnsetResult
`func (o *ResponseExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePost) UnsetResult()`

UnsetResult ensures that no value is present for Result, not even an explicit nil
### GetSearchKeys

`func (o *ResponseExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePost) GetSearchKeys() map[string]string`

GetSearchKeys returns the SearchKeys field if non-nil, zero value otherwise.

### GetSearchKeysOk

`func (o *ResponseExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePost) GetSearchKeysOk() (*map[string]string, bool)`

GetSearchKeysOk returns a tuple with the SearchKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchKeys

`func (o *ResponseExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePost) SetSearchKeys(v map[string]string)`

SetSearchKeys sets SearchKeys field to given value.

### HasSearchKeys

`func (o *ResponseExecuteWorkflowRegistrationV1WorkflowsRegistrationsWorkflowRegistrationIdExecutePost) HasSearchKeys() bool`

HasSearchKeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


