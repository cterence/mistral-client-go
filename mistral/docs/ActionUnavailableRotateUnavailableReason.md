# ActionUnavailableRotateUnavailableReason

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] [default to "unavailable"]
**Reason** | Pointer to [**NullableRotateUnavailableReason**](RotateUnavailableReason.md) |  | [optional] 

## Methods

### NewActionUnavailableRotateUnavailableReason

`func NewActionUnavailableRotateUnavailableReason() *ActionUnavailableRotateUnavailableReason`

NewActionUnavailableRotateUnavailableReason instantiates a new ActionUnavailableRotateUnavailableReason object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionUnavailableRotateUnavailableReasonWithDefaults

`func NewActionUnavailableRotateUnavailableReasonWithDefaults() *ActionUnavailableRotateUnavailableReason`

NewActionUnavailableRotateUnavailableReasonWithDefaults instantiates a new ActionUnavailableRotateUnavailableReason object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ActionUnavailableRotateUnavailableReason) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ActionUnavailableRotateUnavailableReason) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ActionUnavailableRotateUnavailableReason) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ActionUnavailableRotateUnavailableReason) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetReason

`func (o *ActionUnavailableRotateUnavailableReason) GetReason() RotateUnavailableReason`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ActionUnavailableRotateUnavailableReason) GetReasonOk() (*RotateUnavailableReason, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ActionUnavailableRotateUnavailableReason) SetReason(v RotateUnavailableReason)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ActionUnavailableRotateUnavailableReason) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *ActionUnavailableRotateUnavailableReason) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *ActionUnavailableRotateUnavailableReason) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


