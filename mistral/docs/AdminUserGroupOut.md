# AdminUserGroupOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | **string** | User group ID. | 
**Name** | **string** | Name of the user group. | 
**Description** | **NullableString** | Optional description of the user group. | 
**ExternallyManaged** | Pointer to **NullableBool** | Whether the group is managed by an external system. | [optional] 
**OrganizationRole** | Pointer to **NullableString** | Organization role assigned to the group. | [optional] 
**ParentGroupIds** | Pointer to **[]string** | UUIDs of the groups this group is directly nested inside. | [optional] 

## Methods

### NewAdminUserGroupOut

`func NewAdminUserGroupOut(uuid string, name string, description NullableString, ) *AdminUserGroupOut`

NewAdminUserGroupOut instantiates a new AdminUserGroupOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminUserGroupOutWithDefaults

`func NewAdminUserGroupOutWithDefaults() *AdminUserGroupOut`

NewAdminUserGroupOutWithDefaults instantiates a new AdminUserGroupOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *AdminUserGroupOut) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *AdminUserGroupOut) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *AdminUserGroupOut) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetName

`func (o *AdminUserGroupOut) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdminUserGroupOut) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdminUserGroupOut) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AdminUserGroupOut) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AdminUserGroupOut) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AdminUserGroupOut) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *AdminUserGroupOut) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AdminUserGroupOut) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetExternallyManaged

`func (o *AdminUserGroupOut) GetExternallyManaged() bool`

GetExternallyManaged returns the ExternallyManaged field if non-nil, zero value otherwise.

### GetExternallyManagedOk

`func (o *AdminUserGroupOut) GetExternallyManagedOk() (*bool, bool)`

GetExternallyManagedOk returns a tuple with the ExternallyManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternallyManaged

`func (o *AdminUserGroupOut) SetExternallyManaged(v bool)`

SetExternallyManaged sets ExternallyManaged field to given value.

### HasExternallyManaged

`func (o *AdminUserGroupOut) HasExternallyManaged() bool`

HasExternallyManaged returns a boolean if a field has been set.

### SetExternallyManagedNil

`func (o *AdminUserGroupOut) SetExternallyManagedNil(b bool)`

 SetExternallyManagedNil sets the value for ExternallyManaged to be an explicit nil

### UnsetExternallyManaged
`func (o *AdminUserGroupOut) UnsetExternallyManaged()`

UnsetExternallyManaged ensures that no value is present for ExternallyManaged, not even an explicit nil
### GetOrganizationRole

`func (o *AdminUserGroupOut) GetOrganizationRole() string`

GetOrganizationRole returns the OrganizationRole field if non-nil, zero value otherwise.

### GetOrganizationRoleOk

`func (o *AdminUserGroupOut) GetOrganizationRoleOk() (*string, bool)`

GetOrganizationRoleOk returns a tuple with the OrganizationRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationRole

`func (o *AdminUserGroupOut) SetOrganizationRole(v string)`

SetOrganizationRole sets OrganizationRole field to given value.

### HasOrganizationRole

`func (o *AdminUserGroupOut) HasOrganizationRole() bool`

HasOrganizationRole returns a boolean if a field has been set.

### SetOrganizationRoleNil

`func (o *AdminUserGroupOut) SetOrganizationRoleNil(b bool)`

 SetOrganizationRoleNil sets the value for OrganizationRole to be an explicit nil

### UnsetOrganizationRole
`func (o *AdminUserGroupOut) UnsetOrganizationRole()`

UnsetOrganizationRole ensures that no value is present for OrganizationRole, not even an explicit nil
### GetParentGroupIds

`func (o *AdminUserGroupOut) GetParentGroupIds() []string`

GetParentGroupIds returns the ParentGroupIds field if non-nil, zero value otherwise.

### GetParentGroupIdsOk

`func (o *AdminUserGroupOut) GetParentGroupIdsOk() (*[]string, bool)`

GetParentGroupIdsOk returns a tuple with the ParentGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentGroupIds

`func (o *AdminUserGroupOut) SetParentGroupIds(v []string)`

SetParentGroupIds sets ParentGroupIds field to given value.

### HasParentGroupIds

`func (o *AdminUserGroupOut) HasParentGroupIds() bool`

HasParentGroupIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


