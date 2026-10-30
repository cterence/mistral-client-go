# AdminUpdateUserGroupIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Updated name of the user group. | [optional] 
**Description** | Pointer to **NullableString** | Updated description of the user group. Pass an empty string to clear it; omit the field or pass null to leave it unchanged. | [optional] 
**ParentGroupIds** | Pointer to **[]string** | Full replacement list of parent group UUIDs. Pass an empty list to remove all parents; omit the field to leave the parent groups unchanged. | [optional] 

## Methods

### NewAdminUpdateUserGroupIn

`func NewAdminUpdateUserGroupIn() *AdminUpdateUserGroupIn`

NewAdminUpdateUserGroupIn instantiates a new AdminUpdateUserGroupIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminUpdateUserGroupInWithDefaults

`func NewAdminUpdateUserGroupInWithDefaults() *AdminUpdateUserGroupIn`

NewAdminUpdateUserGroupInWithDefaults instantiates a new AdminUpdateUserGroupIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AdminUpdateUserGroupIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdminUpdateUserGroupIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdminUpdateUserGroupIn) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AdminUpdateUserGroupIn) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AdminUpdateUserGroupIn) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AdminUpdateUserGroupIn) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *AdminUpdateUserGroupIn) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AdminUpdateUserGroupIn) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AdminUpdateUserGroupIn) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AdminUpdateUserGroupIn) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AdminUpdateUserGroupIn) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AdminUpdateUserGroupIn) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetParentGroupIds

`func (o *AdminUpdateUserGroupIn) GetParentGroupIds() []string`

GetParentGroupIds returns the ParentGroupIds field if non-nil, zero value otherwise.

### GetParentGroupIdsOk

`func (o *AdminUpdateUserGroupIn) GetParentGroupIdsOk() (*[]string, bool)`

GetParentGroupIdsOk returns a tuple with the ParentGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentGroupIds

`func (o *AdminUpdateUserGroupIn) SetParentGroupIds(v []string)`

SetParentGroupIds sets ParentGroupIds field to given value.

### HasParentGroupIds

`func (o *AdminUpdateUserGroupIn) HasParentGroupIds() bool`

HasParentGroupIds returns a boolean if a field has been set.

### SetParentGroupIdsNil

`func (o *AdminUpdateUserGroupIn) SetParentGroupIdsNil(b bool)`

 SetParentGroupIdsNil sets the value for ParentGroupIds to be an explicit nil

### UnsetParentGroupIds
`func (o *AdminUpdateUserGroupIn) UnsetParentGroupIds()`

UnsetParentGroupIds ensures that no value is present for ParentGroupIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


