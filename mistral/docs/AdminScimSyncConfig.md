# AdminScimSyncConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteMissingGroups** | Pointer to **bool** | Delete Organization groups that are absent from the SCIM provider. | [optional] [default to false]
**DeprovisionUsers** | Pointer to **bool** | Remove Organization users that are inactive or absent from the SCIM provider. | [optional] [default to false]
**SyncUsers** | **bool** | Add users found in the SCIM provider to the Organization. | 
**SyncGroups** | **bool** | Synchronize group metadata from the SCIM provider. | 
**SyncMemberships** | **bool** | Synchronize additions and removals of users in SCIM groups. | 

## Methods

### NewAdminScimSyncConfig

`func NewAdminScimSyncConfig(syncUsers bool, syncGroups bool, syncMemberships bool, ) *AdminScimSyncConfig`

NewAdminScimSyncConfig instantiates a new AdminScimSyncConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminScimSyncConfigWithDefaults

`func NewAdminScimSyncConfigWithDefaults() *AdminScimSyncConfig`

NewAdminScimSyncConfigWithDefaults instantiates a new AdminScimSyncConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteMissingGroups

`func (o *AdminScimSyncConfig) GetDeleteMissingGroups() bool`

GetDeleteMissingGroups returns the DeleteMissingGroups field if non-nil, zero value otherwise.

### GetDeleteMissingGroupsOk

`func (o *AdminScimSyncConfig) GetDeleteMissingGroupsOk() (*bool, bool)`

GetDeleteMissingGroupsOk returns a tuple with the DeleteMissingGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteMissingGroups

`func (o *AdminScimSyncConfig) SetDeleteMissingGroups(v bool)`

SetDeleteMissingGroups sets DeleteMissingGroups field to given value.

### HasDeleteMissingGroups

`func (o *AdminScimSyncConfig) HasDeleteMissingGroups() bool`

HasDeleteMissingGroups returns a boolean if a field has been set.

### GetDeprovisionUsers

`func (o *AdminScimSyncConfig) GetDeprovisionUsers() bool`

GetDeprovisionUsers returns the DeprovisionUsers field if non-nil, zero value otherwise.

### GetDeprovisionUsersOk

`func (o *AdminScimSyncConfig) GetDeprovisionUsersOk() (*bool, bool)`

GetDeprovisionUsersOk returns a tuple with the DeprovisionUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprovisionUsers

`func (o *AdminScimSyncConfig) SetDeprovisionUsers(v bool)`

SetDeprovisionUsers sets DeprovisionUsers field to given value.

### HasDeprovisionUsers

`func (o *AdminScimSyncConfig) HasDeprovisionUsers() bool`

HasDeprovisionUsers returns a boolean if a field has been set.

### GetSyncUsers

`func (o *AdminScimSyncConfig) GetSyncUsers() bool`

GetSyncUsers returns the SyncUsers field if non-nil, zero value otherwise.

### GetSyncUsersOk

`func (o *AdminScimSyncConfig) GetSyncUsersOk() (*bool, bool)`

GetSyncUsersOk returns a tuple with the SyncUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncUsers

`func (o *AdminScimSyncConfig) SetSyncUsers(v bool)`

SetSyncUsers sets SyncUsers field to given value.


### GetSyncGroups

`func (o *AdminScimSyncConfig) GetSyncGroups() bool`

GetSyncGroups returns the SyncGroups field if non-nil, zero value otherwise.

### GetSyncGroupsOk

`func (o *AdminScimSyncConfig) GetSyncGroupsOk() (*bool, bool)`

GetSyncGroupsOk returns a tuple with the SyncGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncGroups

`func (o *AdminScimSyncConfig) SetSyncGroups(v bool)`

SetSyncGroups sets SyncGroups field to given value.


### GetSyncMemberships

`func (o *AdminScimSyncConfig) GetSyncMemberships() bool`

GetSyncMemberships returns the SyncMemberships field if non-nil, zero value otherwise.

### GetSyncMembershipsOk

`func (o *AdminScimSyncConfig) GetSyncMembershipsOk() (*bool, bool)`

GetSyncMembershipsOk returns a tuple with the SyncMemberships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncMemberships

`func (o *AdminScimSyncConfig) SetSyncMemberships(v bool)`

SetSyncMemberships sets SyncMemberships field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


