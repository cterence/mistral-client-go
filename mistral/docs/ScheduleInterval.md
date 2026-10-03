# ScheduleInterval

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Every** | **string** |  | 
**Offset** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewScheduleInterval

`func NewScheduleInterval(every string, ) *ScheduleInterval`

NewScheduleInterval instantiates a new ScheduleInterval object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleIntervalWithDefaults

`func NewScheduleIntervalWithDefaults() *ScheduleInterval`

NewScheduleIntervalWithDefaults instantiates a new ScheduleInterval object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvery

`func (o *ScheduleInterval) GetEvery() string`

GetEvery returns the Every field if non-nil, zero value otherwise.

### GetEveryOk

`func (o *ScheduleInterval) GetEveryOk() (*string, bool)`

GetEveryOk returns a tuple with the Every field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvery

`func (o *ScheduleInterval) SetEvery(v string)`

SetEvery sets Every field to given value.


### GetOffset

`func (o *ScheduleInterval) GetOffset() string`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ScheduleInterval) GetOffsetOk() (*string, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ScheduleInterval) SetOffset(v string)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ScheduleInterval) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### SetOffsetNil

`func (o *ScheduleInterval) SetOffsetNil(b bool)`

 SetOffsetNil sets the value for Offset to be an explicit nil

### UnsetOffset
`func (o *ScheduleInterval) UnsetOffset()`

UnsetOffset ensures that no value is present for Offset, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


