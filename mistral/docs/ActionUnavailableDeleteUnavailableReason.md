# ActionUnavailableDeleteUnavailableReason

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] [default to "unavailable"]
**Reason** | Pointer to [**NullableDeleteUnavailableReason**](DeleteUnavailableReason.md) |  | [optional] 

## Methods

### NewActionUnavailableDeleteUnavailableReason

`func NewActionUnavailableDeleteUnavailableReason() *ActionUnavailableDeleteUnavailableReason`

NewActionUnavailableDeleteUnavailableReason instantiates a new ActionUnavailableDeleteUnavailableReason object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionUnavailableDeleteUnavailableReasonWithDefaults

`func NewActionUnavailableDeleteUnavailableReasonWithDefaults() *ActionUnavailableDeleteUnavailableReason`

NewActionUnavailableDeleteUnavailableReasonWithDefaults instantiates a new ActionUnavailableDeleteUnavailableReason object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ActionUnavailableDeleteUnavailableReason) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ActionUnavailableDeleteUnavailableReason) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ActionUnavailableDeleteUnavailableReason) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ActionUnavailableDeleteUnavailableReason) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetReason

`func (o *ActionUnavailableDeleteUnavailableReason) GetReason() DeleteUnavailableReason`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ActionUnavailableDeleteUnavailableReason) GetReasonOk() (*DeleteUnavailableReason, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ActionUnavailableDeleteUnavailableReason) SetReason(v DeleteUnavailableReason)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ActionUnavailableDeleteUnavailableReason) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *ActionUnavailableDeleteUnavailableReason) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *ActionUnavailableDeleteUnavailableReason) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


