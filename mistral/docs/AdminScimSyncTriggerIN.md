# AdminScimSyncTriggerIN

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DryRun** | Pointer to **bool** | Preview all synchronization changes without applying them. | [optional] [default to true]
**SyncConfig** | Pointer to [**NullableAdminScimSyncConfig**](AdminScimSyncConfig.md) | Categories to synchronize. Required when dry_run is false; ignored for dry runs. | [optional] 

## Methods

### NewAdminScimSyncTriggerIN

`func NewAdminScimSyncTriggerIN() *AdminScimSyncTriggerIN`

NewAdminScimSyncTriggerIN instantiates a new AdminScimSyncTriggerIN object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminScimSyncTriggerINWithDefaults

`func NewAdminScimSyncTriggerINWithDefaults() *AdminScimSyncTriggerIN`

NewAdminScimSyncTriggerINWithDefaults instantiates a new AdminScimSyncTriggerIN object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDryRun

`func (o *AdminScimSyncTriggerIN) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *AdminScimSyncTriggerIN) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *AdminScimSyncTriggerIN) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *AdminScimSyncTriggerIN) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetSyncConfig

`func (o *AdminScimSyncTriggerIN) GetSyncConfig() AdminScimSyncConfig`

GetSyncConfig returns the SyncConfig field if non-nil, zero value otherwise.

### GetSyncConfigOk

`func (o *AdminScimSyncTriggerIN) GetSyncConfigOk() (*AdminScimSyncConfig, bool)`

GetSyncConfigOk returns a tuple with the SyncConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncConfig

`func (o *AdminScimSyncTriggerIN) SetSyncConfig(v AdminScimSyncConfig)`

SetSyncConfig sets SyncConfig field to given value.

### HasSyncConfig

`func (o *AdminScimSyncTriggerIN) HasSyncConfig() bool`

HasSyncConfig returns a boolean if a field has been set.

### SetSyncConfigNil

`func (o *AdminScimSyncTriggerIN) SetSyncConfigNil(b bool)`

 SetSyncConfigNil sets the value for SyncConfig to be an explicit nil

### UnsetSyncConfig
`func (o *AdminScimSyncTriggerIN) UnsetSyncConfig()`

UnsetSyncConfig ensures that no value is present for SyncConfig, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


