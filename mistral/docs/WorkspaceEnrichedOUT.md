# WorkspaceEnrichedOUT

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | **string** | Workspace ID. | 
**Name** | **string** | Workspace name. | 
**Description** | **NullableString** | Workspace description. | 
**Icon** | **NullableString** | Workspace icon. | 
**MembersCount** | Pointer to **NullableInt32** | Number of members in the Workspace. | [optional] 
**GroupsCount** | Pointer to **NullableInt32** | Number of groups assigned to the Workspace. | [optional] 
**SpendLimit** | Pointer to [**NullableWorkspaceSpendLimitOUT**](WorkspaceSpendLimitOUT.md) | Workspace spending limit. | [optional] 
**IsDefault** | **bool** | Whether this is the default Workspace for the Organization. | 
**RawRoles** | **[]string** | Roles granted for the Workspace. | 
**RawRole** | **string** | Deprecated single role for the Workspace. Use &#39;raw_roles&#39; instead. | 

## Methods

### NewWorkspaceEnrichedOUT

`func NewWorkspaceEnrichedOUT(uuid string, name string, description NullableString, icon NullableString, isDefault bool, rawRoles []string, rawRole string, ) *WorkspaceEnrichedOUT`

NewWorkspaceEnrichedOUT instantiates a new WorkspaceEnrichedOUT object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceEnrichedOUTWithDefaults

`func NewWorkspaceEnrichedOUTWithDefaults() *WorkspaceEnrichedOUT`

NewWorkspaceEnrichedOUTWithDefaults instantiates a new WorkspaceEnrichedOUT object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *WorkspaceEnrichedOUT) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *WorkspaceEnrichedOUT) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *WorkspaceEnrichedOUT) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetName

`func (o *WorkspaceEnrichedOUT) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkspaceEnrichedOUT) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkspaceEnrichedOUT) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *WorkspaceEnrichedOUT) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkspaceEnrichedOUT) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkspaceEnrichedOUT) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *WorkspaceEnrichedOUT) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *WorkspaceEnrichedOUT) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIcon

`func (o *WorkspaceEnrichedOUT) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *WorkspaceEnrichedOUT) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *WorkspaceEnrichedOUT) SetIcon(v string)`

SetIcon sets Icon field to given value.


### SetIconNil

`func (o *WorkspaceEnrichedOUT) SetIconNil(b bool)`

 SetIconNil sets the value for Icon to be an explicit nil

### UnsetIcon
`func (o *WorkspaceEnrichedOUT) UnsetIcon()`

UnsetIcon ensures that no value is present for Icon, not even an explicit nil
### GetMembersCount

`func (o *WorkspaceEnrichedOUT) GetMembersCount() int32`

GetMembersCount returns the MembersCount field if non-nil, zero value otherwise.

### GetMembersCountOk

`func (o *WorkspaceEnrichedOUT) GetMembersCountOk() (*int32, bool)`

GetMembersCountOk returns a tuple with the MembersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersCount

`func (o *WorkspaceEnrichedOUT) SetMembersCount(v int32)`

SetMembersCount sets MembersCount field to given value.

### HasMembersCount

`func (o *WorkspaceEnrichedOUT) HasMembersCount() bool`

HasMembersCount returns a boolean if a field has been set.

### SetMembersCountNil

`func (o *WorkspaceEnrichedOUT) SetMembersCountNil(b bool)`

 SetMembersCountNil sets the value for MembersCount to be an explicit nil

### UnsetMembersCount
`func (o *WorkspaceEnrichedOUT) UnsetMembersCount()`

UnsetMembersCount ensures that no value is present for MembersCount, not even an explicit nil
### GetGroupsCount

`func (o *WorkspaceEnrichedOUT) GetGroupsCount() int32`

GetGroupsCount returns the GroupsCount field if non-nil, zero value otherwise.

### GetGroupsCountOk

`func (o *WorkspaceEnrichedOUT) GetGroupsCountOk() (*int32, bool)`

GetGroupsCountOk returns a tuple with the GroupsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupsCount

`func (o *WorkspaceEnrichedOUT) SetGroupsCount(v int32)`

SetGroupsCount sets GroupsCount field to given value.

### HasGroupsCount

`func (o *WorkspaceEnrichedOUT) HasGroupsCount() bool`

HasGroupsCount returns a boolean if a field has been set.

### SetGroupsCountNil

`func (o *WorkspaceEnrichedOUT) SetGroupsCountNil(b bool)`

 SetGroupsCountNil sets the value for GroupsCount to be an explicit nil

### UnsetGroupsCount
`func (o *WorkspaceEnrichedOUT) UnsetGroupsCount()`

UnsetGroupsCount ensures that no value is present for GroupsCount, not even an explicit nil
### GetSpendLimit

`func (o *WorkspaceEnrichedOUT) GetSpendLimit() WorkspaceSpendLimitOUT`

GetSpendLimit returns the SpendLimit field if non-nil, zero value otherwise.

### GetSpendLimitOk

`func (o *WorkspaceEnrichedOUT) GetSpendLimitOk() (*WorkspaceSpendLimitOUT, bool)`

GetSpendLimitOk returns a tuple with the SpendLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendLimit

`func (o *WorkspaceEnrichedOUT) SetSpendLimit(v WorkspaceSpendLimitOUT)`

SetSpendLimit sets SpendLimit field to given value.

### HasSpendLimit

`func (o *WorkspaceEnrichedOUT) HasSpendLimit() bool`

HasSpendLimit returns a boolean if a field has been set.

### SetSpendLimitNil

`func (o *WorkspaceEnrichedOUT) SetSpendLimitNil(b bool)`

 SetSpendLimitNil sets the value for SpendLimit to be an explicit nil

### UnsetSpendLimit
`func (o *WorkspaceEnrichedOUT) UnsetSpendLimit()`

UnsetSpendLimit ensures that no value is present for SpendLimit, not even an explicit nil
### GetIsDefault

`func (o *WorkspaceEnrichedOUT) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *WorkspaceEnrichedOUT) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *WorkspaceEnrichedOUT) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.


### GetRawRoles

`func (o *WorkspaceEnrichedOUT) GetRawRoles() []string`

GetRawRoles returns the RawRoles field if non-nil, zero value otherwise.

### GetRawRolesOk

`func (o *WorkspaceEnrichedOUT) GetRawRolesOk() (*[]string, bool)`

GetRawRolesOk returns a tuple with the RawRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawRoles

`func (o *WorkspaceEnrichedOUT) SetRawRoles(v []string)`

SetRawRoles sets RawRoles field to given value.


### GetRawRole

`func (o *WorkspaceEnrichedOUT) GetRawRole() string`

GetRawRole returns the RawRole field if non-nil, zero value otherwise.

### GetRawRoleOk

`func (o *WorkspaceEnrichedOUT) GetRawRoleOk() (*string, bool)`

GetRawRoleOk returns a tuple with the RawRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawRole

`func (o *WorkspaceEnrichedOUT) SetRawRole(v string)`

SetRawRole sets RawRole field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


