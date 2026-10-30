# RunStatisticsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RunId** | **string** |  | 
**Statistics** | [**map[string]EvaluatorStatistics**](EvaluatorStatistics.md) |  | 
**GoalResults** | Pointer to [**map[string]GoalResult**](GoalResult.md) |  | [optional] [default to {}]
**Passed** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewRunStatisticsResponse

`func NewRunStatisticsResponse(runId string, statistics map[string]EvaluatorStatistics, ) *RunStatisticsResponse`

NewRunStatisticsResponse instantiates a new RunStatisticsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunStatisticsResponseWithDefaults

`func NewRunStatisticsResponseWithDefaults() *RunStatisticsResponse`

NewRunStatisticsResponseWithDefaults instantiates a new RunStatisticsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRunId

`func (o *RunStatisticsResponse) GetRunId() string`

GetRunId returns the RunId field if non-nil, zero value otherwise.

### GetRunIdOk

`func (o *RunStatisticsResponse) GetRunIdOk() (*string, bool)`

GetRunIdOk returns a tuple with the RunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunId

`func (o *RunStatisticsResponse) SetRunId(v string)`

SetRunId sets RunId field to given value.


### GetStatistics

`func (o *RunStatisticsResponse) GetStatistics() map[string]EvaluatorStatistics`

GetStatistics returns the Statistics field if non-nil, zero value otherwise.

### GetStatisticsOk

`func (o *RunStatisticsResponse) GetStatisticsOk() (*map[string]EvaluatorStatistics, bool)`

GetStatisticsOk returns a tuple with the Statistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatistics

`func (o *RunStatisticsResponse) SetStatistics(v map[string]EvaluatorStatistics)`

SetStatistics sets Statistics field to given value.


### GetGoalResults

`func (o *RunStatisticsResponse) GetGoalResults() map[string]GoalResult`

GetGoalResults returns the GoalResults field if non-nil, zero value otherwise.

### GetGoalResultsOk

`func (o *RunStatisticsResponse) GetGoalResultsOk() (*map[string]GoalResult, bool)`

GetGoalResultsOk returns a tuple with the GoalResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoalResults

`func (o *RunStatisticsResponse) SetGoalResults(v map[string]GoalResult)`

SetGoalResults sets GoalResults field to given value.

### HasGoalResults

`func (o *RunStatisticsResponse) HasGoalResults() bool`

HasGoalResults returns a boolean if a field has been set.

### GetPassed

`func (o *RunStatisticsResponse) GetPassed() bool`

GetPassed returns the Passed field if non-nil, zero value otherwise.

### GetPassedOk

`func (o *RunStatisticsResponse) GetPassedOk() (*bool, bool)`

GetPassedOk returns a tuple with the Passed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassed

`func (o *RunStatisticsResponse) SetPassed(v bool)`

SetPassed sets Passed field to given value.

### HasPassed

`func (o *RunStatisticsResponse) HasPassed() bool`

HasPassed returns a boolean if a field has been set.

### SetPassedNil

`func (o *RunStatisticsResponse) SetPassedNil(b bool)`

 SetPassedNil sets the value for Passed to be an explicit nil

### UnsetPassed
`func (o *RunStatisticsResponse) UnsetPassed()`

UnsetPassed ensures that no value is present for Passed, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


