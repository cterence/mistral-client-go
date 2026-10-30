# WorkspaceOUT

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

## Methods

### NewWorkspaceOUT

`func NewWorkspaceOUT(uuid string, name string, description NullableString, icon NullableString, isDefault bool, ) *WorkspaceOUT`

NewWorkspaceOUT instantiates a new WorkspaceOUT object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceOUTWithDefaults

`func NewWorkspaceOUTWithDefaults() *WorkspaceOUT`

NewWorkspaceOUTWithDefaults instantiates a new WorkspaceOUT object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *WorkspaceOUT) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *WorkspaceOUT) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *WorkspaceOUT) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetName

`func (o *WorkspaceOUT) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkspaceOUT) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkspaceOUT) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *WorkspaceOUT) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkspaceOUT) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkspaceOUT) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *WorkspaceOUT) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *WorkspaceOUT) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIcon

`func (o *WorkspaceOUT) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *WorkspaceOUT) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *WorkspaceOUT) SetIcon(v string)`

SetIcon sets Icon field to given value.


### SetIconNil

`func (o *WorkspaceOUT) SetIconNil(b bool)`

 SetIconNil sets the value for Icon to be an explicit nil

### UnsetIcon
`func (o *WorkspaceOUT) UnsetIcon()`

UnsetIcon ensures that no value is present for Icon, not even an explicit nil
### GetMembersCount

`func (o *WorkspaceOUT) GetMembersCount() int32`

GetMembersCount returns the MembersCount field if non-nil, zero value otherwise.

### GetMembersCountOk

`func (o *WorkspaceOUT) GetMembersCountOk() (*int32, bool)`

GetMembersCountOk returns a tuple with the MembersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersCount

`func (o *WorkspaceOUT) SetMembersCount(v int32)`

SetMembersCount sets MembersCount field to given value.

### HasMembersCount

`func (o *WorkspaceOUT) HasMembersCount() bool`

HasMembersCount returns a boolean if a field has been set.

### SetMembersCountNil

`func (o *WorkspaceOUT) SetMembersCountNil(b bool)`

 SetMembersCountNil sets the value for MembersCount to be an explicit nil

### UnsetMembersCount
`func (o *WorkspaceOUT) UnsetMembersCount()`

UnsetMembersCount ensures that no value is present for MembersCount, not even an explicit nil
### GetGroupsCount

`func (o *WorkspaceOUT) GetGroupsCount() int32`

GetGroupsCount returns the GroupsCount field if non-nil, zero value otherwise.

### GetGroupsCountOk

`func (o *WorkspaceOUT) GetGroupsCountOk() (*int32, bool)`

GetGroupsCountOk returns a tuple with the GroupsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupsCount

`func (o *WorkspaceOUT) SetGroupsCount(v int32)`

SetGroupsCount sets GroupsCount field to given value.

### HasGroupsCount

`func (o *WorkspaceOUT) HasGroupsCount() bool`

HasGroupsCount returns a boolean if a field has been set.

### SetGroupsCountNil

`func (o *WorkspaceOUT) SetGroupsCountNil(b bool)`

 SetGroupsCountNil sets the value for GroupsCount to be an explicit nil

### UnsetGroupsCount
`func (o *WorkspaceOUT) UnsetGroupsCount()`

UnsetGroupsCount ensures that no value is present for GroupsCount, not even an explicit nil
### GetSpendLimit

`func (o *WorkspaceOUT) GetSpendLimit() WorkspaceSpendLimitOUT`

GetSpendLimit returns the SpendLimit field if non-nil, zero value otherwise.

### GetSpendLimitOk

`func (o *WorkspaceOUT) GetSpendLimitOk() (*WorkspaceSpendLimitOUT, bool)`

GetSpendLimitOk returns a tuple with the SpendLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendLimit

`func (o *WorkspaceOUT) SetSpendLimit(v WorkspaceSpendLimitOUT)`

SetSpendLimit sets SpendLimit field to given value.

### HasSpendLimit

`func (o *WorkspaceOUT) HasSpendLimit() bool`

HasSpendLimit returns a boolean if a field has been set.

### SetSpendLimitNil

`func (o *WorkspaceOUT) SetSpendLimitNil(b bool)`

 SetSpendLimitNil sets the value for SpendLimit to be an explicit nil

### UnsetSpendLimit
`func (o *WorkspaceOUT) UnsetSpendLimit()`

UnsetSpendLimit ensures that no value is present for SpendLimit, not even an explicit nil
### GetIsDefault

`func (o *WorkspaceOUT) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *WorkspaceOUT) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *WorkspaceOUT) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


