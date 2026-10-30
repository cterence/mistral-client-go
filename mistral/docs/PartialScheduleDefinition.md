# PartialScheduleDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Input** | Pointer to **interface{}** | Input to provide to the workflow when starting it. | [optional] 
**Calendars** | Pointer to [**[]ScheduleCalendar**](ScheduleCalendar.md) | Calendar-based specification of times. | [optional] 
**Intervals** | Pointer to [**[]ScheduleInterval**](ScheduleInterval.md) | Interval-based specification of times. | [optional] 
**CronExpressions** | Pointer to **[]string** | Cron-based specification of times. | [optional] 
**Skip** | Pointer to [**[]ScheduleCalendar**](ScheduleCalendar.md) | Set of calendar times to skip. | [optional] 
**StartAt** | Pointer to **NullableTime** | Time after which the first action may be run. | [optional] 
**EndAt** | Pointer to **NullableTime** | Time after which no more actions will be run. | [optional] 
**Jitter** | Pointer to **NullableString** | Jitter to apply each action.  An action&#39;s scheduled time will be incremented by a random value between 0 and this value if present (but not past the next schedule).  | [optional] 
**TimeZoneName** | Pointer to **NullableString** | IANA time zone name, for example &#x60;&#x60;US/Central&#x60;&#x60;. | [optional] 
**Policy** | Pointer to [**SchedulePolicy**](SchedulePolicy.md) | Policy for the schedule. | [optional] 
**MaxExecutions** | Pointer to **NullableInt32** | Maximum number of times this schedule will trigger a workflow execution. Once this limit is reached, no further executions are triggered automatically. null means unlimited. | [optional] 

## Methods

### NewPartialScheduleDefinition

`func NewPartialScheduleDefinition() *PartialScheduleDefinition`

NewPartialScheduleDefinition instantiates a new PartialScheduleDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartialScheduleDefinitionWithDefaults

`func NewPartialScheduleDefinitionWithDefaults() *PartialScheduleDefinition`

NewPartialScheduleDefinitionWithDefaults instantiates a new PartialScheduleDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInput

`func (o *PartialScheduleDefinition) GetInput() interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *PartialScheduleDefinition) GetInputOk() (*interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *PartialScheduleDefinition) SetInput(v interface{})`

SetInput sets Input field to given value.

### HasInput

`func (o *PartialScheduleDefinition) HasInput() bool`

HasInput returns a boolean if a field has been set.

### SetInputNil

`func (o *PartialScheduleDefinition) SetInputNil(b bool)`

 SetInputNil sets the value for Input to be an explicit nil

### UnsetInput
`func (o *PartialScheduleDefinition) UnsetInput()`

UnsetInput ensures that no value is present for Input, not even an explicit nil
### GetCalendars

`func (o *PartialScheduleDefinition) GetCalendars() []ScheduleCalendar`

GetCalendars returns the Calendars field if non-nil, zero value otherwise.

### GetCalendarsOk

`func (o *PartialScheduleDefinition) GetCalendarsOk() (*[]ScheduleCalendar, bool)`

GetCalendarsOk returns a tuple with the Calendars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalendars

`func (o *PartialScheduleDefinition) SetCalendars(v []ScheduleCalendar)`

SetCalendars sets Calendars field to given value.

### HasCalendars

`func (o *PartialScheduleDefinition) HasCalendars() bool`

HasCalendars returns a boolean if a field has been set.

### GetIntervals

`func (o *PartialScheduleDefinition) GetIntervals() []ScheduleInterval`

GetIntervals returns the Intervals field if non-nil, zero value otherwise.

### GetIntervalsOk

`func (o *PartialScheduleDefinition) GetIntervalsOk() (*[]ScheduleInterval, bool)`

GetIntervalsOk returns a tuple with the Intervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervals

`func (o *PartialScheduleDefinition) SetIntervals(v []ScheduleInterval)`

SetIntervals sets Intervals field to given value.

### HasIntervals

`func (o *PartialScheduleDefinition) HasIntervals() bool`

HasIntervals returns a boolean if a field has been set.

### GetCronExpressions

`func (o *PartialScheduleDefinition) GetCronExpressions() []string`

GetCronExpressions returns the CronExpressions field if non-nil, zero value otherwise.

### GetCronExpressionsOk

`func (o *PartialScheduleDefinition) GetCronExpressionsOk() (*[]string, bool)`

GetCronExpressionsOk returns a tuple with the CronExpressions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronExpressions

`func (o *PartialScheduleDefinition) SetCronExpressions(v []string)`

SetCronExpressions sets CronExpressions field to given value.

### HasCronExpressions

`func (o *PartialScheduleDefinition) HasCronExpressions() bool`

HasCronExpressions returns a boolean if a field has been set.

### GetSkip

`func (o *PartialScheduleDefinition) GetSkip() []ScheduleCalendar`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *PartialScheduleDefinition) GetSkipOk() (*[]ScheduleCalendar, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *PartialScheduleDefinition) SetSkip(v []ScheduleCalendar)`

SetSkip sets Skip field to given value.

### HasSkip

`func (o *PartialScheduleDefinition) HasSkip() bool`

HasSkip returns a boolean if a field has been set.

### GetStartAt

`func (o *PartialScheduleDefinition) GetStartAt() time.Time`

GetStartAt returns the StartAt field if non-nil, zero value otherwise.

### GetStartAtOk

`func (o *PartialScheduleDefinition) GetStartAtOk() (*time.Time, bool)`

GetStartAtOk returns a tuple with the StartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAt

`func (o *PartialScheduleDefinition) SetStartAt(v time.Time)`

SetStartAt sets StartAt field to given value.

### HasStartAt

`func (o *PartialScheduleDefinition) HasStartAt() bool`

HasStartAt returns a boolean if a field has been set.

### SetStartAtNil

`func (o *PartialScheduleDefinition) SetStartAtNil(b bool)`

 SetStartAtNil sets the value for StartAt to be an explicit nil

### UnsetStartAt
`func (o *PartialScheduleDefinition) UnsetStartAt()`

UnsetStartAt ensures that no value is present for StartAt, not even an explicit nil
### GetEndAt

`func (o *PartialScheduleDefinition) GetEndAt() time.Time`

GetEndAt returns the EndAt field if non-nil, zero value otherwise.

### GetEndAtOk

`func (o *PartialScheduleDefinition) GetEndAtOk() (*time.Time, bool)`

GetEndAtOk returns a tuple with the EndAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAt

`func (o *PartialScheduleDefinition) SetEndAt(v time.Time)`

SetEndAt sets EndAt field to given value.

### HasEndAt

`func (o *PartialScheduleDefinition) HasEndAt() bool`

HasEndAt returns a boolean if a field has been set.

### SetEndAtNil

`func (o *PartialScheduleDefinition) SetEndAtNil(b bool)`

 SetEndAtNil sets the value for EndAt to be an explicit nil

### UnsetEndAt
`func (o *PartialScheduleDefinition) UnsetEndAt()`

UnsetEndAt ensures that no value is present for EndAt, not even an explicit nil
### GetJitter

`func (o *PartialScheduleDefinition) GetJitter() string`

GetJitter returns the Jitter field if non-nil, zero value otherwise.

### GetJitterOk

`func (o *PartialScheduleDefinition) GetJitterOk() (*string, bool)`

GetJitterOk returns a tuple with the Jitter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJitter

`func (o *PartialScheduleDefinition) SetJitter(v string)`

SetJitter sets Jitter field to given value.

### HasJitter

`func (o *PartialScheduleDefinition) HasJitter() bool`

HasJitter returns a boolean if a field has been set.

### SetJitterNil

`func (o *PartialScheduleDefinition) SetJitterNil(b bool)`

 SetJitterNil sets the value for Jitter to be an explicit nil

### UnsetJitter
`func (o *PartialScheduleDefinition) UnsetJitter()`

UnsetJitter ensures that no value is present for Jitter, not even an explicit nil
### GetTimeZoneName

`func (o *PartialScheduleDefinition) GetTimeZoneName() string`

GetTimeZoneName returns the TimeZoneName field if non-nil, zero value otherwise.

### GetTimeZoneNameOk

`func (o *PartialScheduleDefinition) GetTimeZoneNameOk() (*string, bool)`

GetTimeZoneNameOk returns a tuple with the TimeZoneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZoneName

`func (o *PartialScheduleDefinition) SetTimeZoneName(v string)`

SetTimeZoneName sets TimeZoneName field to given value.

### HasTimeZoneName

`func (o *PartialScheduleDefinition) HasTimeZoneName() bool`

HasTimeZoneName returns a boolean if a field has been set.

### SetTimeZoneNameNil

`func (o *PartialScheduleDefinition) SetTimeZoneNameNil(b bool)`

 SetTimeZoneNameNil sets the value for TimeZoneName to be an explicit nil

### UnsetTimeZoneName
`func (o *PartialScheduleDefinition) UnsetTimeZoneName()`

UnsetTimeZoneName ensures that no value is present for TimeZoneName, not even an explicit nil
### GetPolicy

`func (o *PartialScheduleDefinition) GetPolicy() SchedulePolicy`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *PartialScheduleDefinition) GetPolicyOk() (*SchedulePolicy, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *PartialScheduleDefinition) SetPolicy(v SchedulePolicy)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *PartialScheduleDefinition) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetMaxExecutions

`func (o *PartialScheduleDefinition) GetMaxExecutions() int32`

GetMaxExecutions returns the MaxExecutions field if non-nil, zero value otherwise.

### GetMaxExecutionsOk

`func (o *PartialScheduleDefinition) GetMaxExecutionsOk() (*int32, bool)`

GetMaxExecutionsOk returns a tuple with the MaxExecutions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxExecutions

`func (o *PartialScheduleDefinition) SetMaxExecutions(v int32)`

SetMaxExecutions sets MaxExecutions field to given value.

### HasMaxExecutions

`func (o *PartialScheduleDefinition) HasMaxExecutions() bool`

HasMaxExecutions returns a boolean if a field has been set.

### SetMaxExecutionsNil

`func (o *PartialScheduleDefinition) SetMaxExecutionsNil(b bool)`

 SetMaxExecutionsNil sets the value for MaxExecutions to be an explicit nil

### UnsetMaxExecutions
`func (o *PartialScheduleDefinition) UnsetMaxExecutions()`

UnsetMaxExecutions ensures that no value is present for MaxExecutions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


