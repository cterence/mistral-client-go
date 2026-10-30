# APIKeyActionsDelete

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] [default to "unavailable"]
**Reason** | Pointer to [**DeleteUnavailableReason**](DeleteUnavailableReason.md) |  | [optional] 

## Methods

### NewAPIKeyActionsDelete

`func NewAPIKeyActionsDelete() *APIKeyActionsDelete`

NewAPIKeyActionsDelete instantiates a new APIKeyActionsDelete object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIKeyActionsDeleteWithDefaults

`func NewAPIKeyActionsDeleteWithDefaults() *APIKeyActionsDelete`

NewAPIKeyActionsDeleteWithDefaults instantiates a new APIKeyActionsDelete object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *APIKeyActionsDelete) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *APIKeyActionsDelete) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *APIKeyActionsDelete) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *APIKeyActionsDelete) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetReason

`func (o *APIKeyActionsDelete) GetReason() DeleteUnavailableReason`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *APIKeyActionsDelete) GetReasonOk() (*DeleteUnavailableReason, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *APIKeyActionsDelete) SetReason(v DeleteUnavailableReason)`

SetReason sets Reason field to given value.

### HasReason

`func (o *APIKeyActionsDelete) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


