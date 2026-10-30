# AdminAssignGroupToWorkspaceIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoleNames** | Pointer to **[]string** | Simplified role names to assign. Mutually exclusive with &#39;roles&#39;. | [optional] 
**Roles** | Pointer to **[]string** | Role values to assign. Mutually exclusive with &#39;role_names&#39;. | [optional] 
**Role** | Pointer to **NullableString** | Deprecated single role value. Use &#39;role_names&#39; instead. | [optional] 
**WorkspaceUuid** | **string** |  | 

## Methods

### NewAdminAssignGroupToWorkspaceIn

`func NewAdminAssignGroupToWorkspaceIn(workspaceUuid string, ) *AdminAssignGroupToWorkspaceIn`

NewAdminAssignGroupToWorkspaceIn instantiates a new AdminAssignGroupToWorkspaceIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminAssignGroupToWorkspaceInWithDefaults

`func NewAdminAssignGroupToWorkspaceInWithDefaults() *AdminAssignGroupToWorkspaceIn`

NewAdminAssignGroupToWorkspaceInWithDefaults instantiates a new AdminAssignGroupToWorkspaceIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoleNames

`func (o *AdminAssignGroupToWorkspaceIn) GetRoleNames() []string`

GetRoleNames returns the RoleNames field if non-nil, zero value otherwise.

### GetRoleNamesOk

`func (o *AdminAssignGroupToWorkspaceIn) GetRoleNamesOk() (*[]string, bool)`

GetRoleNamesOk returns a tuple with the RoleNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleNames

`func (o *AdminAssignGroupToWorkspaceIn) SetRoleNames(v []string)`

SetRoleNames sets RoleNames field to given value.

### HasRoleNames

`func (o *AdminAssignGroupToWorkspaceIn) HasRoleNames() bool`

HasRoleNames returns a boolean if a field has been set.

### SetRoleNamesNil

`func (o *AdminAssignGroupToWorkspaceIn) SetRoleNamesNil(b bool)`

 SetRoleNamesNil sets the value for RoleNames to be an explicit nil

### UnsetRoleNames
`func (o *AdminAssignGroupToWorkspaceIn) UnsetRoleNames()`

UnsetRoleNames ensures that no value is present for RoleNames, not even an explicit nil
### GetRoles

`func (o *AdminAssignGroupToWorkspaceIn) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *AdminAssignGroupToWorkspaceIn) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *AdminAssignGroupToWorkspaceIn) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *AdminAssignGroupToWorkspaceIn) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### SetRolesNil

`func (o *AdminAssignGroupToWorkspaceIn) SetRolesNil(b bool)`

 SetRolesNil sets the value for Roles to be an explicit nil

### UnsetRoles
`func (o *AdminAssignGroupToWorkspaceIn) UnsetRoles()`

UnsetRoles ensures that no value is present for Roles, not even an explicit nil
### GetRole

`func (o *AdminAssignGroupToWorkspaceIn) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *AdminAssignGroupToWorkspaceIn) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *AdminAssignGroupToWorkspaceIn) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *AdminAssignGroupToWorkspaceIn) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *AdminAssignGroupToWorkspaceIn) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *AdminAssignGroupToWorkspaceIn) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetWorkspaceUuid

`func (o *AdminAssignGroupToWorkspaceIn) GetWorkspaceUuid() string`

GetWorkspaceUuid returns the WorkspaceUuid field if non-nil, zero value otherwise.

### GetWorkspaceUuidOk

`func (o *AdminAssignGroupToWorkspaceIn) GetWorkspaceUuidOk() (*string, bool)`

GetWorkspaceUuidOk returns a tuple with the WorkspaceUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceUuid

`func (o *AdminAssignGroupToWorkspaceIn) SetWorkspaceUuid(v string)`

SetWorkspaceUuid sets WorkspaceUuid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


