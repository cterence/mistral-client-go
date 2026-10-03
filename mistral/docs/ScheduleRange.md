# ScheduleRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Start** | **int32** |  | 
**End** | Pointer to **int32** |  | [optional] [default to 0]
**Step** | Pointer to **int32** |  | [optional] [default to 0]

## Methods

### NewScheduleRange

`func NewScheduleRange(start int32, ) *ScheduleRange`

NewScheduleRange instantiates a new ScheduleRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleRangeWithDefaults

`func NewScheduleRangeWithDefaults() *ScheduleRange`

NewScheduleRangeWithDefaults instantiates a new ScheduleRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStart

`func (o *ScheduleRange) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ScheduleRange) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ScheduleRange) SetStart(v int32)`

SetStart sets Start field to given value.


### GetEnd

`func (o *ScheduleRange) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *ScheduleRange) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *ScheduleRange) SetEnd(v int32)`

SetEnd sets End field to given value.

### HasEnd

`func (o *ScheduleRange) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetStep

`func (o *ScheduleRange) GetStep() int32`

GetStep returns the Step field if non-nil, zero value otherwise.

### GetStepOk

`func (o *ScheduleRange) GetStepOk() (*int32, bool)`

GetStepOk returns a tuple with the Step field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStep

`func (o *ScheduleRange) SetStep(v int32)`

SetStep sets Step field to given value.

### HasStep

`func (o *ScheduleRange) HasStep() bool`

HasStep returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


