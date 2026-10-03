# ScheduleCalendar

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Second** | Pointer to [**[]ScheduleRange**](ScheduleRange.md) |  | [optional] [default to {{start=0, end=0, step=0}}]
**Minute** | Pointer to [**[]ScheduleRange**](ScheduleRange.md) |  | [optional] [default to {{start=0, end=0, step=0}}]
**Hour** | Pointer to [**[]ScheduleRange**](ScheduleRange.md) |  | [optional] [default to {{start=0, end=0, step=0}}]
**DayOfMonth** | Pointer to [**[]ScheduleRange**](ScheduleRange.md) |  | [optional] [default to {{start=1, end=31, step=0}}]
**Month** | Pointer to [**[]ScheduleRange**](ScheduleRange.md) |  | [optional] [default to {{start=1, end=12, step=0}}]
**Year** | Pointer to [**[]ScheduleRange**](ScheduleRange.md) |  | [optional] [default to {}]
**DayOfWeek** | Pointer to [**[]ScheduleRange**](ScheduleRange.md) |  | [optional] [default to {{start=0, end=6, step=0}}]
**Comment** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewScheduleCalendar

`func NewScheduleCalendar() *ScheduleCalendar`

NewScheduleCalendar instantiates a new ScheduleCalendar object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleCalendarWithDefaults

`func NewScheduleCalendarWithDefaults() *ScheduleCalendar`

NewScheduleCalendarWithDefaults instantiates a new ScheduleCalendar object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecond

`func (o *ScheduleCalendar) GetSecond() []ScheduleRange`

GetSecond returns the Second field if non-nil, zero value otherwise.

### GetSecondOk

`func (o *ScheduleCalendar) GetSecondOk() (*[]ScheduleRange, bool)`

GetSecondOk returns a tuple with the Second field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecond

`func (o *ScheduleCalendar) SetSecond(v []ScheduleRange)`

SetSecond sets Second field to given value.

### HasSecond

`func (o *ScheduleCalendar) HasSecond() bool`

HasSecond returns a boolean if a field has been set.

### GetMinute

`func (o *ScheduleCalendar) GetMinute() []ScheduleRange`

GetMinute returns the Minute field if non-nil, zero value otherwise.

### GetMinuteOk

`func (o *ScheduleCalendar) GetMinuteOk() (*[]ScheduleRange, bool)`

GetMinuteOk returns a tuple with the Minute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinute

`func (o *ScheduleCalendar) SetMinute(v []ScheduleRange)`

SetMinute sets Minute field to given value.

### HasMinute

`func (o *ScheduleCalendar) HasMinute() bool`

HasMinute returns a boolean if a field has been set.

### GetHour

`func (o *ScheduleCalendar) GetHour() []ScheduleRange`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *ScheduleCalendar) GetHourOk() (*[]ScheduleRange, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *ScheduleCalendar) SetHour(v []ScheduleRange)`

SetHour sets Hour field to given value.

### HasHour

`func (o *ScheduleCalendar) HasHour() bool`

HasHour returns a boolean if a field has been set.

### GetDayOfMonth

`func (o *ScheduleCalendar) GetDayOfMonth() []ScheduleRange`

GetDayOfMonth returns the DayOfMonth field if non-nil, zero value otherwise.

### GetDayOfMonthOk

`func (o *ScheduleCalendar) GetDayOfMonthOk() (*[]ScheduleRange, bool)`

GetDayOfMonthOk returns a tuple with the DayOfMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfMonth

`func (o *ScheduleCalendar) SetDayOfMonth(v []ScheduleRange)`

SetDayOfMonth sets DayOfMonth field to given value.

### HasDayOfMonth

`func (o *ScheduleCalendar) HasDayOfMonth() bool`

HasDayOfMonth returns a boolean if a field has been set.

### GetMonth

`func (o *ScheduleCalendar) GetMonth() []ScheduleRange`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *ScheduleCalendar) GetMonthOk() (*[]ScheduleRange, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *ScheduleCalendar) SetMonth(v []ScheduleRange)`

SetMonth sets Month field to given value.

### HasMonth

`func (o *ScheduleCalendar) HasMonth() bool`

HasMonth returns a boolean if a field has been set.

### GetYear

`func (o *ScheduleCalendar) GetYear() []ScheduleRange`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *ScheduleCalendar) GetYearOk() (*[]ScheduleRange, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *ScheduleCalendar) SetYear(v []ScheduleRange)`

SetYear sets Year field to given value.

### HasYear

`func (o *ScheduleCalendar) HasYear() bool`

HasYear returns a boolean if a field has been set.

### GetDayOfWeek

`func (o *ScheduleCalendar) GetDayOfWeek() []ScheduleRange`

GetDayOfWeek returns the DayOfWeek field if non-nil, zero value otherwise.

### GetDayOfWeekOk

`func (o *ScheduleCalendar) GetDayOfWeekOk() (*[]ScheduleRange, bool)`

GetDayOfWeekOk returns a tuple with the DayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfWeek

`func (o *ScheduleCalendar) SetDayOfWeek(v []ScheduleRange)`

SetDayOfWeek sets DayOfWeek field to given value.

### HasDayOfWeek

`func (o *ScheduleCalendar) HasDayOfWeek() bool`

HasDayOfWeek returns a boolean if a field has been set.

### GetComment

`func (o *ScheduleCalendar) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ScheduleCalendar) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ScheduleCalendar) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ScheduleCalendar) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *ScheduleCalendar) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *ScheduleCalendar) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


