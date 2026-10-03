# ScheduleDefinition

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
**ScheduleId** | Pointer to **NullableString** | Unique identifier for the schedule. | [optional] 

## Methods

### NewScheduleDefinition

`func NewScheduleDefinition(input interface{}, ) *ScheduleDefinition`

NewScheduleDefinition instantiates a new ScheduleDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleDefinitionWithDefaults

`func NewScheduleDefinitionWithDefaults() *ScheduleDefinition`

NewScheduleDefinitionWithDefaults instantiates a new ScheduleDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInput

`func (o *ScheduleDefinition) GetInput() interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *ScheduleDefinition) GetInputOk() (*interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *ScheduleDefinition) SetInput(v interface{})`

SetInput sets Input field to given value.


### SetInputNil

`func (o *ScheduleDefinition) SetInputNil(b bool)`

 SetInputNil sets the value for Input to be an explicit nil

### UnsetInput
`func (o *ScheduleDefinition) UnsetInput()`

UnsetInput ensures that no value is present for Input, not even an explicit nil
### GetCalendars

`func (o *ScheduleDefinition) GetCalendars() []ScheduleCalendar`

GetCalendars returns the Calendars field if non-nil, zero value otherwise.

### GetCalendarsOk

`func (o *ScheduleDefinition) GetCalendarsOk() (*[]ScheduleCalendar, bool)`

GetCalendarsOk returns a tuple with the Calendars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalendars

`func (o *ScheduleDefinition) SetCalendars(v []ScheduleCalendar)`

SetCalendars sets Calendars field to given value.

### HasCalendars

`func (o *ScheduleDefinition) HasCalendars() bool`

HasCalendars returns a boolean if a field has been set.

### GetIntervals

`func (o *ScheduleDefinition) GetIntervals() []ScheduleInterval`

GetIntervals returns the Intervals field if non-nil, zero value otherwise.

### GetIntervalsOk

`func (o *ScheduleDefinition) GetIntervalsOk() (*[]ScheduleInterval, bool)`

GetIntervalsOk returns a tuple with the Intervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervals

`func (o *ScheduleDefinition) SetIntervals(v []ScheduleInterval)`

SetIntervals sets Intervals field to given value.

### HasIntervals

`func (o *ScheduleDefinition) HasIntervals() bool`

HasIntervals returns a boolean if a field has been set.

### GetCronExpressions

`func (o *ScheduleDefinition) GetCronExpressions() []string`

GetCronExpressions returns the CronExpressions field if non-nil, zero value otherwise.

### GetCronExpressionsOk

`func (o *ScheduleDefinition) GetCronExpressionsOk() (*[]string, bool)`

GetCronExpressionsOk returns a tuple with the CronExpressions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronExpressions

`func (o *ScheduleDefinition) SetCronExpressions(v []string)`

SetCronExpressions sets CronExpressions field to given value.

### HasCronExpressions

`func (o *ScheduleDefinition) HasCronExpressions() bool`

HasCronExpressions returns a boolean if a field has been set.

### GetSkip

`func (o *ScheduleDefinition) GetSkip() []ScheduleCalendar`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *ScheduleDefinition) GetSkipOk() (*[]ScheduleCalendar, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *ScheduleDefinition) SetSkip(v []ScheduleCalendar)`

SetSkip sets Skip field to given value.

### HasSkip

`func (o *ScheduleDefinition) HasSkip() bool`

HasSkip returns a boolean if a field has been set.

### GetStartAt

`func (o *ScheduleDefinition) GetStartAt() time.Time`

GetStartAt returns the StartAt field if non-nil, zero value otherwise.

### GetStartAtOk

`func (o *ScheduleDefinition) GetStartAtOk() (*time.Time, bool)`

GetStartAtOk returns a tuple with the StartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAt

`func (o *ScheduleDefinition) SetStartAt(v time.Time)`

SetStartAt sets StartAt field to given value.

### HasStartAt

`func (o *ScheduleDefinition) HasStartAt() bool`

HasStartAt returns a boolean if a field has been set.

### SetStartAtNil

`func (o *ScheduleDefinition) SetStartAtNil(b bool)`

 SetStartAtNil sets the value for StartAt to be an explicit nil

### UnsetStartAt
`func (o *ScheduleDefinition) UnsetStartAt()`

UnsetStartAt ensures that no value is present for StartAt, not even an explicit nil
### GetEndAt

`func (o *ScheduleDefinition) GetEndAt() time.Time`

GetEndAt returns the EndAt field if non-nil, zero value otherwise.

### GetEndAtOk

`func (o *ScheduleDefinition) GetEndAtOk() (*time.Time, bool)`

GetEndAtOk returns a tuple with the EndAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAt

`func (o *ScheduleDefinition) SetEndAt(v time.Time)`

SetEndAt sets EndAt field to given value.

### HasEndAt

`func (o *ScheduleDefinition) HasEndAt() bool`

HasEndAt returns a boolean if a field has been set.

### SetEndAtNil

`func (o *ScheduleDefinition) SetEndAtNil(b bool)`

 SetEndAtNil sets the value for EndAt to be an explicit nil

### UnsetEndAt
`func (o *ScheduleDefinition) UnsetEndAt()`

UnsetEndAt ensures that no value is present for EndAt, not even an explicit nil
### GetJitter

`func (o *ScheduleDefinition) GetJitter() string`

GetJitter returns the Jitter field if non-nil, zero value otherwise.

### GetJitterOk

`func (o *ScheduleDefinition) GetJitterOk() (*string, bool)`

GetJitterOk returns a tuple with the Jitter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJitter

`func (o *ScheduleDefinition) SetJitter(v string)`

SetJitter sets Jitter field to given value.

### HasJitter

`func (o *ScheduleDefinition) HasJitter() bool`

HasJitter returns a boolean if a field has been set.

### SetJitterNil

`func (o *ScheduleDefinition) SetJitterNil(b bool)`

 SetJitterNil sets the value for Jitter to be an explicit nil

### UnsetJitter
`func (o *ScheduleDefinition) UnsetJitter()`

UnsetJitter ensures that no value is present for Jitter, not even an explicit nil
### GetTimeZoneName

`func (o *ScheduleDefinition) GetTimeZoneName() string`

GetTimeZoneName returns the TimeZoneName field if non-nil, zero value otherwise.

### GetTimeZoneNameOk

`func (o *ScheduleDefinition) GetTimeZoneNameOk() (*string, bool)`

GetTimeZoneNameOk returns a tuple with the TimeZoneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZoneName

`func (o *ScheduleDefinition) SetTimeZoneName(v string)`

SetTimeZoneName sets TimeZoneName field to given value.

### HasTimeZoneName

`func (o *ScheduleDefinition) HasTimeZoneName() bool`

HasTimeZoneName returns a boolean if a field has been set.

### SetTimeZoneNameNil

`func (o *ScheduleDefinition) SetTimeZoneNameNil(b bool)`

 SetTimeZoneNameNil sets the value for TimeZoneName to be an explicit nil

### UnsetTimeZoneName
`func (o *ScheduleDefinition) UnsetTimeZoneName()`

UnsetTimeZoneName ensures that no value is present for TimeZoneName, not even an explicit nil
### GetPolicy

`func (o *ScheduleDefinition) GetPolicy() SchedulePolicy`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *ScheduleDefinition) GetPolicyOk() (*SchedulePolicy, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *ScheduleDefinition) SetPolicy(v SchedulePolicy)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *ScheduleDefinition) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetScheduleId

`func (o *ScheduleDefinition) GetScheduleId() string`

GetScheduleId returns the ScheduleId field if non-nil, zero value otherwise.

### GetScheduleIdOk

`func (o *ScheduleDefinition) GetScheduleIdOk() (*string, bool)`

GetScheduleIdOk returns a tuple with the ScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleId

`func (o *ScheduleDefinition) SetScheduleId(v string)`

SetScheduleId sets ScheduleId field to given value.

### HasScheduleId

`func (o *ScheduleDefinition) HasScheduleId() bool`

HasScheduleId returns a boolean if a field has been set.

### SetScheduleIdNil

`func (o *ScheduleDefinition) SetScheduleIdNil(b bool)`

 SetScheduleIdNil sets the value for ScheduleId to be an explicit nil

### UnsetScheduleId
`func (o *ScheduleDefinition) UnsetScheduleId()`

UnsetScheduleId ensures that no value is present for ScheduleId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


