# ScheduleFutureExecution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScheduledAt** | **time.Time** | Time the execution is scheduled to run. | 

## Methods

### NewScheduleFutureExecution

`func NewScheduleFutureExecution(scheduledAt time.Time, ) *ScheduleFutureExecution`

NewScheduleFutureExecution instantiates a new ScheduleFutureExecution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleFutureExecutionWithDefaults

`func NewScheduleFutureExecutionWithDefaults() *ScheduleFutureExecution`

NewScheduleFutureExecutionWithDefaults instantiates a new ScheduleFutureExecution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScheduledAt

`func (o *ScheduleFutureExecution) GetScheduledAt() time.Time`

GetScheduledAt returns the ScheduledAt field if non-nil, zero value otherwise.

### GetScheduledAtOk

`func (o *ScheduleFutureExecution) GetScheduledAtOk() (*time.Time, bool)`

GetScheduledAtOk returns a tuple with the ScheduledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledAt

`func (o *ScheduleFutureExecution) SetScheduledAt(v time.Time)`

SetScheduledAt sets ScheduledAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


