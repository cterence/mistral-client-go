# ScheduleRecentExecution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScheduledAt** | **time.Time** | Time the execution was scheduled to run. | 
**StartedAt** | **time.Time** | Actual time the execution started. | 
**ExecutionId** | **string** | ID of the workflow execution that was started. | 

## Methods

### NewScheduleRecentExecution

`func NewScheduleRecentExecution(scheduledAt time.Time, startedAt time.Time, executionId string, ) *ScheduleRecentExecution`

NewScheduleRecentExecution instantiates a new ScheduleRecentExecution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleRecentExecutionWithDefaults

`func NewScheduleRecentExecutionWithDefaults() *ScheduleRecentExecution`

NewScheduleRecentExecutionWithDefaults instantiates a new ScheduleRecentExecution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScheduledAt

`func (o *ScheduleRecentExecution) GetScheduledAt() time.Time`

GetScheduledAt returns the ScheduledAt field if non-nil, zero value otherwise.

### GetScheduledAtOk

`func (o *ScheduleRecentExecution) GetScheduledAtOk() (*time.Time, bool)`

GetScheduledAtOk returns a tuple with the ScheduledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledAt

`func (o *ScheduleRecentExecution) SetScheduledAt(v time.Time)`

SetScheduledAt sets ScheduledAt field to given value.


### GetStartedAt

`func (o *ScheduleRecentExecution) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *ScheduleRecentExecution) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *ScheduleRecentExecution) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### GetExecutionId

`func (o *ScheduleRecentExecution) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *ScheduleRecentExecution) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *ScheduleRecentExecution) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


