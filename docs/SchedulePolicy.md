# SchedulePolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CatchupWindowSeconds** | Pointer to **int32** | After a Temporal server is unavailable, amount of time in seconds in the past to execute missed actions. | [optional] [default to 31536000]
**Overlap** | Pointer to [**ScheduleOverlapPolicy**](ScheduleOverlapPolicy.md) | Policy controlling what to do when a workflow is already running. | [optional] [default to SCHEDULEOVERLAPPOLICY__1]
**PauseOnFailure** | Pointer to **bool** | Whether to pause the schedule after a workflow failure. | [optional] [default to false]

## Methods

### NewSchedulePolicy

`func NewSchedulePolicy() *SchedulePolicy`

NewSchedulePolicy instantiates a new SchedulePolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedulePolicyWithDefaults

`func NewSchedulePolicyWithDefaults() *SchedulePolicy`

NewSchedulePolicyWithDefaults instantiates a new SchedulePolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCatchupWindowSeconds

`func (o *SchedulePolicy) GetCatchupWindowSeconds() int32`

GetCatchupWindowSeconds returns the CatchupWindowSeconds field if non-nil, zero value otherwise.

### GetCatchupWindowSecondsOk

`func (o *SchedulePolicy) GetCatchupWindowSecondsOk() (*int32, bool)`

GetCatchupWindowSecondsOk returns a tuple with the CatchupWindowSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatchupWindowSeconds

`func (o *SchedulePolicy) SetCatchupWindowSeconds(v int32)`

SetCatchupWindowSeconds sets CatchupWindowSeconds field to given value.

### HasCatchupWindowSeconds

`func (o *SchedulePolicy) HasCatchupWindowSeconds() bool`

HasCatchupWindowSeconds returns a boolean if a field has been set.

### GetOverlap

`func (o *SchedulePolicy) GetOverlap() ScheduleOverlapPolicy`

GetOverlap returns the Overlap field if non-nil, zero value otherwise.

### GetOverlapOk

`func (o *SchedulePolicy) GetOverlapOk() (*ScheduleOverlapPolicy, bool)`

GetOverlapOk returns a tuple with the Overlap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverlap

`func (o *SchedulePolicy) SetOverlap(v ScheduleOverlapPolicy)`

SetOverlap sets Overlap field to given value.

### HasOverlap

`func (o *SchedulePolicy) HasOverlap() bool`

HasOverlap returns a boolean if a field has been set.

### GetPauseOnFailure

`func (o *SchedulePolicy) GetPauseOnFailure() bool`

GetPauseOnFailure returns the PauseOnFailure field if non-nil, zero value otherwise.

### GetPauseOnFailureOk

`func (o *SchedulePolicy) GetPauseOnFailureOk() (*bool, bool)`

GetPauseOnFailureOk returns a tuple with the PauseOnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPauseOnFailure

`func (o *SchedulePolicy) SetPauseOnFailure(v bool)`

SetPauseOnFailure sets PauseOnFailure field to given value.

### HasPauseOnFailure

`func (o *SchedulePolicy) HasPauseOnFailure() bool`

HasPauseOnFailure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


