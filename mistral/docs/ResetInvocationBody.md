# ResetInvocationBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | **int32** | The event ID to reset the workflow execution to | 
**Reason** | Pointer to **NullableString** | Reason for resetting the workflow execution | [optional] 
**ExcludeSignals** | Pointer to **bool** | Whether to exclude signals that happened after the reset point | [optional] [default to false]
**ExcludeUpdates** | Pointer to **bool** | Whether to exclude updates that happened after the reset point | [optional] [default to false]

## Methods

### NewResetInvocationBody

`func NewResetInvocationBody(eventId int32, ) *ResetInvocationBody`

NewResetInvocationBody instantiates a new ResetInvocationBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResetInvocationBodyWithDefaults

`func NewResetInvocationBodyWithDefaults() *ResetInvocationBody`

NewResetInvocationBodyWithDefaults instantiates a new ResetInvocationBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventId

`func (o *ResetInvocationBody) GetEventId() int32`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *ResetInvocationBody) GetEventIdOk() (*int32, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *ResetInvocationBody) SetEventId(v int32)`

SetEventId sets EventId field to given value.


### GetReason

`func (o *ResetInvocationBody) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ResetInvocationBody) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ResetInvocationBody) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ResetInvocationBody) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *ResetInvocationBody) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *ResetInvocationBody) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetExcludeSignals

`func (o *ResetInvocationBody) GetExcludeSignals() bool`

GetExcludeSignals returns the ExcludeSignals field if non-nil, zero value otherwise.

### GetExcludeSignalsOk

`func (o *ResetInvocationBody) GetExcludeSignalsOk() (*bool, bool)`

GetExcludeSignalsOk returns a tuple with the ExcludeSignals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeSignals

`func (o *ResetInvocationBody) SetExcludeSignals(v bool)`

SetExcludeSignals sets ExcludeSignals field to given value.

### HasExcludeSignals

`func (o *ResetInvocationBody) HasExcludeSignals() bool`

HasExcludeSignals returns a boolean if a field has been set.

### GetExcludeUpdates

`func (o *ResetInvocationBody) GetExcludeUpdates() bool`

GetExcludeUpdates returns the ExcludeUpdates field if non-nil, zero value otherwise.

### GetExcludeUpdatesOk

`func (o *ResetInvocationBody) GetExcludeUpdatesOk() (*bool, bool)`

GetExcludeUpdatesOk returns a tuple with the ExcludeUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeUpdates

`func (o *ResetInvocationBody) SetExcludeUpdates(v bool)`

SetExcludeUpdates sets ExcludeUpdates field to given value.

### HasExcludeUpdates

`func (o *ResetInvocationBody) HasExcludeUpdates() bool`

HasExcludeUpdates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


