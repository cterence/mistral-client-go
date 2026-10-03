# ScheduleDefinitionOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Input** | **interface{}** | Input to provide to the workflow when starting it. | 
**Calendars** | Pointer to [**[]ScheduleCalendar**](ScheduleCalendar.md) | Calendar-based specification of times. | [optional] 
**Intervals** | Pointer to [**[]ScheduleInterval**](ScheduleInterval.md) | Interval-based specification of times. | [optional] 
**CronExpressions** | Pointer to **[]string** | Cron-based specification of times. | [optional] 
**Skip** | Pointer to [**[]ScheduleCalendar**](ScheduleCalendar.md) | Set of calendar times to skip. | [optional] 
**StartAt** | Pointer to **NullableTime** | Time after which the first action may be run. | [optional] 
**EndAt** | Pointer to **NullableTime** | Time after which no more actions will be run. | [optional] 
**Jitter** | Pointer to **NullableString** | Jitter to apply each action.  An action&#39;s scheduled time will be incremented by a random value between 0 and this value if present (but not past the next schedule).  | [optional] 
**TimeZoneName** | Pointer to **NullableString** | IANA time zone name, for example &#x60;&#x60;US/Central&#x60;&#x60;. | [optional] 
**Policy** | Pointer to [**SchedulePolicy**](SchedulePolicy.md) | Policy for the schedule. | [optional] 
**ScheduleId** | **string** | Unique identifier for the schedule. | 

## Methods

### NewScheduleDefinitionOutput

`func NewScheduleDefinitionOutput(input interface{}, scheduleId string, ) *ScheduleDefinitionOutput`

NewScheduleDefinitionOutput instantiates a new ScheduleDefinitionOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleDefinitionOutputWithDefaults

`func NewScheduleDefinitionOutputWithDefaults() *ScheduleDefinitionOutput`

NewScheduleDefinitionOutputWithDefaults instantiates a new ScheduleDefinitionOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInput

`func (o *ScheduleDefinitionOutput) GetInput() interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *ScheduleDefinitionOutput) GetInputOk() (*interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *ScheduleDefinitionOutput) SetInput(v interface{})`

SetInput sets Input field to given value.


### SetInputNil

`func (o *ScheduleDefinitionOutput) SetInputNil(b bool)`

 SetInputNil sets the value for Input to be an explicit nil

### UnsetInput
`func (o *ScheduleDefinitionOutput) UnsetInput()`

UnsetInput ensures that no value is present for Input, not even an explicit nil
### GetCalendars

`func (o *ScheduleDefinitionOutput) GetCalendars() []ScheduleCalendar`

GetCalendars returns the Calendars field if non-nil, zero value otherwise.

### GetCalendarsOk

`func (o *ScheduleDefinitionOutput) GetCalendarsOk() (*[]ScheduleCalendar, bool)`

GetCalendarsOk returns a tuple with the Calendars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalendars

`func (o *ScheduleDefinitionOutput) SetCalendars(v []ScheduleCalendar)`

SetCalendars sets Calendars field to given value.

### HasCalendars

`func (o *ScheduleDefinitionOutput) HasCalendars() bool`

HasCalendars returns a boolean if a field has been set.

### GetIntervals

`func (o *ScheduleDefinitionOutput) GetIntervals() []ScheduleInterval`

GetIntervals returns the Intervals field if non-nil, zero value otherwise.

### GetIntervalsOk

`func (o *ScheduleDefinitionOutput) GetIntervalsOk() (*[]ScheduleInterval, bool)`

GetIntervalsOk returns a tuple with the Intervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervals

`func (o *ScheduleDefinitionOutput) SetIntervals(v []ScheduleInterval)`

SetIntervals sets Intervals field to given value.

### HasIntervals

`func (o *ScheduleDefinitionOutput) HasIntervals() bool`

HasIntervals returns a boolean if a field has been set.

### GetCronExpressions

`func (o *ScheduleDefinitionOutput) GetCronExpressions() []string`

GetCronExpressions returns the CronExpressions field if non-nil, zero value otherwise.

### GetCronExpressionsOk

`func (o *ScheduleDefinitionOutput) GetCronExpressionsOk() (*[]string, bool)`

GetCronExpressionsOk returns a tuple with the CronExpressions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronExpressions

`func (o *ScheduleDefinitionOutput) SetCronExpressions(v []string)`

SetCronExpressions sets CronExpressions field to given value.

### HasCronExpressions

`func (o *ScheduleDefinitionOutput) HasCronExpressions() bool`

HasCronExpressions returns a boolean if a field has been set.

### GetSkip

`func (o *ScheduleDefinitionOutput) GetSkip() []ScheduleCalendar`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *ScheduleDefinitionOutput) GetSkipOk() (*[]ScheduleCalendar, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *ScheduleDefinitionOutput) SetSkip(v []ScheduleCalendar)`

SetSkip sets Skip field to given value.

### HasSkip

`func (o *ScheduleDefinitionOutput) HasSkip() bool`

HasSkip returns a boolean if a field has been set.

### GetStartAt

`func (o *ScheduleDefinitionOutput) GetStartAt() time.Time`

GetStartAt returns the StartAt field if non-nil, zero value otherwise.

### GetStartAtOk

`func (o *ScheduleDefinitionOutput) GetStartAtOk() (*time.Time, bool)`

GetStartAtOk returns a tuple with the StartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAt

`func (o *ScheduleDefinitionOutput) SetStartAt(v time.Time)`

SetStartAt sets StartAt field to given value.

### HasStartAt

`func (o *ScheduleDefinitionOutput) HasStartAt() bool`

HasStartAt returns a boolean if a field has been set.

### SetStartAtNil

`func (o *ScheduleDefinitionOutput) SetStartAtNil(b bool)`

 SetStartAtNil sets the value for StartAt to be an explicit nil

### UnsetStartAt
`func (o *ScheduleDefinitionOutput) UnsetStartAt()`

UnsetStartAt ensures that no value is present for StartAt, not even an explicit nil
### GetEndAt

`func (o *ScheduleDefinitionOutput) GetEndAt() time.Time`

GetEndAt returns the EndAt field if non-nil, zero value otherwise.

### GetEndAtOk

`func (o *ScheduleDefinitionOutput) GetEndAtOk() (*time.Time, bool)`

GetEndAtOk returns a tuple with the EndAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAt

`func (o *ScheduleDefinitionOutput) SetEndAt(v time.Time)`

SetEndAt sets EndAt field to given value.

### HasEndAt

`func (o *ScheduleDefinitionOutput) HasEndAt() bool`

HasEndAt returns a boolean if a field has been set.

### SetEndAtNil

`func (o *ScheduleDefinitionOutput) SetEndAtNil(b bool)`

 SetEndAtNil sets the value for EndAt to be an explicit nil

### UnsetEndAt
`func (o *ScheduleDefinitionOutput) UnsetEndAt()`

UnsetEndAt ensures that no value is present for EndAt, not even an explicit nil
### GetJitter

`func (o *ScheduleDefinitionOutput) GetJitter() string`

GetJitter returns the Jitter field if non-nil, zero value otherwise.

### GetJitterOk

`func (o *ScheduleDefinitionOutput) GetJitterOk() (*string, bool)`

GetJitterOk returns a tuple with the Jitter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJitter

`func (o *ScheduleDefinitionOutput) SetJitter(v string)`

SetJitter sets Jitter field to given value.

### HasJitter

`func (o *ScheduleDefinitionOutput) HasJitter() bool`

HasJitter returns a boolean if a field has been set.

### SetJitterNil

`func (o *ScheduleDefinitionOutput) SetJitterNil(b bool)`

 SetJitterNil sets the value for Jitter to be an explicit nil

### UnsetJitter
`func (o *ScheduleDefinitionOutput) UnsetJitter()`

UnsetJitter ensures that no value is present for Jitter, not even an explicit nil
### GetTimeZoneName

`func (o *ScheduleDefinitionOutput) GetTimeZoneName() string`

GetTimeZoneName returns the TimeZoneName field if non-nil, zero value otherwise.

### GetTimeZoneNameOk

`func (o *ScheduleDefinitionOutput) GetTimeZoneNameOk() (*string, bool)`

GetTimeZoneNameOk returns a tuple with the TimeZoneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZoneName

`func (o *ScheduleDefinitionOutput) SetTimeZoneName(v string)`

SetTimeZoneName sets TimeZoneName field to given value.

### HasTimeZoneName

`func (o *ScheduleDefinitionOutput) HasTimeZoneName() bool`

HasTimeZoneName returns a boolean if a field has been set.

### SetTimeZoneNameNil

`func (o *ScheduleDefinitionOutput) SetTimeZoneNameNil(b bool)`

 SetTimeZoneNameNil sets the value for TimeZoneName to be an explicit nil

### UnsetTimeZoneName
`func (o *ScheduleDefinitionOutput) UnsetTimeZoneName()`

UnsetTimeZoneName ensures that no value is present for TimeZoneName, not even an explicit nil
### GetPolicy

`func (o *ScheduleDefinitionOutput) GetPolicy() SchedulePolicy`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *ScheduleDefinitionOutput) GetPolicyOk() (*SchedulePolicy, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *ScheduleDefinitionOutput) SetPolicy(v SchedulePolicy)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *ScheduleDefinitionOutput) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetScheduleId

`func (o *ScheduleDefinitionOutput) GetScheduleId() string`

GetScheduleId returns the ScheduleId field if non-nil, zero value otherwise.

### GetScheduleIdOk

`func (o *ScheduleDefinitionOutput) GetScheduleIdOk() (*string, bool)`

GetScheduleIdOk returns a tuple with the ScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleId

`func (o *ScheduleDefinitionOutput) SetScheduleId(v string)`

SetScheduleId sets ScheduleId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


