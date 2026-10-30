# AdminOrganizationMemberUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoleNames** | Pointer to **[]string** | Simplified role names to assign. Mutually exclusive with &#39;roles&#39;. | [optional] 
**Roles** | Pointer to [**NullableRoles1**](Roles1.md) |  | [optional] 
**RoleName** | Pointer to **NullableString** | Deprecated single role name, kept for backward compatibility. Mutually exclusive with &#39;role&#39;. | [optional] 
**Role** | Pointer to **NullableString** | Deprecated legacy single role value, kept for backward compatibility. Mutually exclusive with &#39;role_name&#39;. | [optional] 
**SubscriptionTypes** | Pointer to **[]string** | Product seats to assign to the member. | [optional] 

## Methods

### NewAdminOrganizationMemberUpdate

`func NewAdminOrganizationMemberUpdate() *AdminOrganizationMemberUpdate`

NewAdminOrganizationMemberUpdate instantiates a new AdminOrganizationMemberUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminOrganizationMemberUpdateWithDefaults

`func NewAdminOrganizationMemberUpdateWithDefaults() *AdminOrganizationMemberUpdate`

NewAdminOrganizationMemberUpdateWithDefaults instantiates a new AdminOrganizationMemberUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoleNames

`func (o *AdminOrganizationMemberUpdate) GetRoleNames() []string`

GetRoleNames returns the RoleNames field if non-nil, zero value otherwise.

### GetRoleNamesOk

`func (o *AdminOrganizationMemberUpdate) GetRoleNamesOk() (*[]string, bool)`

GetRoleNamesOk returns a tuple with the RoleNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleNames

`func (o *AdminOrganizationMemberUpdate) SetRoleNames(v []string)`

SetRoleNames sets RoleNames field to given value.

### HasRoleNames

`func (o *AdminOrganizationMemberUpdate) HasRoleNames() bool`

HasRoleNames returns a boolean if a field has been set.

### SetRoleNamesNil

`func (o *AdminOrganizationMemberUpdate) SetRoleNamesNil(b bool)`

 SetRoleNamesNil sets the value for RoleNames to be an explicit nil

### UnsetRoleNames
`func (o *AdminOrganizationMemberUpdate) UnsetRoleNames()`

UnsetRoleNames ensures that no value is present for RoleNames, not even an explicit nil
### GetRoles

`func (o *AdminOrganizationMemberUpdate) GetRoles() Roles1`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *AdminOrganizationMemberUpdate) GetRolesOk() (*Roles1, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *AdminOrganizationMemberUpdate) SetRoles(v Roles1)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *AdminOrganizationMemberUpdate) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### SetRolesNil

`func (o *AdminOrganizationMemberUpdate) SetRolesNil(b bool)`

 SetRolesNil sets the value for Roles to be an explicit nil

### UnsetRoles
`func (o *AdminOrganizationMemberUpdate) UnsetRoles()`

UnsetRoles ensures that no value is present for Roles, not even an explicit nil
### GetRoleName

`func (o *AdminOrganizationMemberUpdate) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *AdminOrganizationMemberUpdate) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *AdminOrganizationMemberUpdate) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.

### HasRoleName

`func (o *AdminOrganizationMemberUpdate) HasRoleName() bool`

HasRoleName returns a boolean if a field has been set.

### SetRoleNameNil

`func (o *AdminOrganizationMemberUpdate) SetRoleNameNil(b bool)`

 SetRoleNameNil sets the value for RoleName to be an explicit nil

### UnsetRoleName
`func (o *AdminOrganizationMemberUpdate) UnsetRoleName()`

UnsetRoleName ensures that no value is present for RoleName, not even an explicit nil
### GetRole

`func (o *AdminOrganizationMemberUpdate) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *AdminOrganizationMemberUpdate) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *AdminOrganizationMemberUpdate) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *AdminOrganizationMemberUpdate) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *AdminOrganizationMemberUpdate) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *AdminOrganizationMemberUpdate) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetSubscriptionTypes

`func (o *AdminOrganizationMemberUpdate) GetSubscriptionTypes() []string`

GetSubscriptionTypes returns the SubscriptionTypes field if non-nil, zero value otherwise.

### GetSubscriptionTypesOk

`func (o *AdminOrganizationMemberUpdate) GetSubscriptionTypesOk() (*[]string, bool)`

GetSubscriptionTypesOk returns a tuple with the SubscriptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionTypes

`func (o *AdminOrganizationMemberUpdate) SetSubscriptionTypes(v []string)`

SetSubscriptionTypes sets SubscriptionTypes field to given value.

### HasSubscriptionTypes

`func (o *AdminOrganizationMemberUpdate) HasSubscriptionTypes() bool`

HasSubscriptionTypes returns a boolean if a field has been set.

### SetSubscriptionTypesNil

`func (o *AdminOrganizationMemberUpdate) SetSubscriptionTypesNil(b bool)`

 SetSubscriptionTypesNil sets the value for SubscriptionTypes to be an explicit nil

### UnsetSubscriptionTypes
`func (o *AdminOrganizationMemberUpdate) UnsetSubscriptionTypes()`

UnsetSubscriptionTypes ensures that no value is present for SubscriptionTypes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


