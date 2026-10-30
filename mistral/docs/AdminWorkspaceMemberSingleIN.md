# AdminWorkspaceMemberSingleIN

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoleNames** | Pointer to **[]string** | Simplified role names to assign. Mutually exclusive with &#39;roles&#39;. | [optional] 
**Roles** | Pointer to [**NullableRoles3**](Roles3.md) |  | [optional] 
**RoleName** | Pointer to **NullableString** | Deprecated single role name, kept for backward compatibility. Mutually exclusive with &#39;role&#39;. | [optional] 
**Role** | Pointer to **NullableString** | Deprecated legacy single role value, kept for backward compatibility. Mutually exclusive with &#39;role_name&#39;. | [optional] 
**UserUuid** | **string** | User ID of the Workspace member. | 

## Methods

### NewAdminWorkspaceMemberSingleIN

`func NewAdminWorkspaceMemberSingleIN(userUuid string, ) *AdminWorkspaceMemberSingleIN`

NewAdminWorkspaceMemberSingleIN instantiates a new AdminWorkspaceMemberSingleIN object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminWorkspaceMemberSingleINWithDefaults

`func NewAdminWorkspaceMemberSingleINWithDefaults() *AdminWorkspaceMemberSingleIN`

NewAdminWorkspaceMemberSingleINWithDefaults instantiates a new AdminWorkspaceMemberSingleIN object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoleNames

`func (o *AdminWorkspaceMemberSingleIN) GetRoleNames() []string`

GetRoleNames returns the RoleNames field if non-nil, zero value otherwise.

### GetRoleNamesOk

`func (o *AdminWorkspaceMemberSingleIN) GetRoleNamesOk() (*[]string, bool)`

GetRoleNamesOk returns a tuple with the RoleNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleNames

`func (o *AdminWorkspaceMemberSingleIN) SetRoleNames(v []string)`

SetRoleNames sets RoleNames field to given value.

### HasRoleNames

`func (o *AdminWorkspaceMemberSingleIN) HasRoleNames() bool`

HasRoleNames returns a boolean if a field has been set.

### SetRoleNamesNil

`func (o *AdminWorkspaceMemberSingleIN) SetRoleNamesNil(b bool)`

 SetRoleNamesNil sets the value for RoleNames to be an explicit nil

### UnsetRoleNames
`func (o *AdminWorkspaceMemberSingleIN) UnsetRoleNames()`

UnsetRoleNames ensures that no value is present for RoleNames, not even an explicit nil
### GetRoles

`func (o *AdminWorkspaceMemberSingleIN) GetRoles() Roles3`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *AdminWorkspaceMemberSingleIN) GetRolesOk() (*Roles3, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *AdminWorkspaceMemberSingleIN) SetRoles(v Roles3)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *AdminWorkspaceMemberSingleIN) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### SetRolesNil

`func (o *AdminWorkspaceMemberSingleIN) SetRolesNil(b bool)`

 SetRolesNil sets the value for Roles to be an explicit nil

### UnsetRoles
`func (o *AdminWorkspaceMemberSingleIN) UnsetRoles()`

UnsetRoles ensures that no value is present for Roles, not even an explicit nil
### GetRoleName

`func (o *AdminWorkspaceMemberSingleIN) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *AdminWorkspaceMemberSingleIN) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *AdminWorkspaceMemberSingleIN) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.

### HasRoleName

`func (o *AdminWorkspaceMemberSingleIN) HasRoleName() bool`

HasRoleName returns a boolean if a field has been set.

### SetRoleNameNil

`func (o *AdminWorkspaceMemberSingleIN) SetRoleNameNil(b bool)`

 SetRoleNameNil sets the value for RoleName to be an explicit nil

### UnsetRoleName
`func (o *AdminWorkspaceMemberSingleIN) UnsetRoleName()`

UnsetRoleName ensures that no value is present for RoleName, not even an explicit nil
### GetRole

`func (o *AdminWorkspaceMemberSingleIN) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *AdminWorkspaceMemberSingleIN) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *AdminWorkspaceMemberSingleIN) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *AdminWorkspaceMemberSingleIN) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *AdminWorkspaceMemberSingleIN) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *AdminWorkspaceMemberSingleIN) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetUserUuid

`func (o *AdminWorkspaceMemberSingleIN) GetUserUuid() string`

GetUserUuid returns the UserUuid field if non-nil, zero value otherwise.

### GetUserUuidOk

`func (o *AdminWorkspaceMemberSingleIN) GetUserUuidOk() (*string, bool)`

GetUserUuidOk returns a tuple with the UserUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserUuid

`func (o *AdminWorkspaceMemberSingleIN) SetUserUuid(v string)`

SetUserUuid sets UserUuid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


