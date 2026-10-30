# AdminProvisionGroupToWorkspaceIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserGroupUuid** | **string** | User group ID to provision. | 
**WorkspaceUuid** | **string** | Workspace ID where the group is provisioned. | 
**WorkspaceRole** | Pointer to **NullableString** | Workspace role value to assign to the group. Mutually exclusive with &#39;workspace_role_name&#39;. | [optional] 
**WorkspaceRoleName** | Pointer to **NullableString** | Workspace role name to assign to the group. Mutually exclusive with &#39;workspace_role&#39;. | [optional] 

## Methods

### NewAdminProvisionGroupToWorkspaceIn

`func NewAdminProvisionGroupToWorkspaceIn(userGroupUuid string, workspaceUuid string, ) *AdminProvisionGroupToWorkspaceIn`

NewAdminProvisionGroupToWorkspaceIn instantiates a new AdminProvisionGroupToWorkspaceIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminProvisionGroupToWorkspaceInWithDefaults

`func NewAdminProvisionGroupToWorkspaceInWithDefaults() *AdminProvisionGroupToWorkspaceIn`

NewAdminProvisionGroupToWorkspaceInWithDefaults instantiates a new AdminProvisionGroupToWorkspaceIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserGroupUuid

`func (o *AdminProvisionGroupToWorkspaceIn) GetUserGroupUuid() string`

GetUserGroupUuid returns the UserGroupUuid field if non-nil, zero value otherwise.

### GetUserGroupUuidOk

`func (o *AdminProvisionGroupToWorkspaceIn) GetUserGroupUuidOk() (*string, bool)`

GetUserGroupUuidOk returns a tuple with the UserGroupUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroupUuid

`func (o *AdminProvisionGroupToWorkspaceIn) SetUserGroupUuid(v string)`

SetUserGroupUuid sets UserGroupUuid field to given value.


### GetWorkspaceUuid

`func (o *AdminProvisionGroupToWorkspaceIn) GetWorkspaceUuid() string`

GetWorkspaceUuid returns the WorkspaceUuid field if non-nil, zero value otherwise.

### GetWorkspaceUuidOk

`func (o *AdminProvisionGroupToWorkspaceIn) GetWorkspaceUuidOk() (*string, bool)`

GetWorkspaceUuidOk returns a tuple with the WorkspaceUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceUuid

`func (o *AdminProvisionGroupToWorkspaceIn) SetWorkspaceUuid(v string)`

SetWorkspaceUuid sets WorkspaceUuid field to given value.


### GetWorkspaceRole

`func (o *AdminProvisionGroupToWorkspaceIn) GetWorkspaceRole() string`

GetWorkspaceRole returns the WorkspaceRole field if non-nil, zero value otherwise.

### GetWorkspaceRoleOk

`func (o *AdminProvisionGroupToWorkspaceIn) GetWorkspaceRoleOk() (*string, bool)`

GetWorkspaceRoleOk returns a tuple with the WorkspaceRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceRole

`func (o *AdminProvisionGroupToWorkspaceIn) SetWorkspaceRole(v string)`

SetWorkspaceRole sets WorkspaceRole field to given value.

### HasWorkspaceRole

`func (o *AdminProvisionGroupToWorkspaceIn) HasWorkspaceRole() bool`

HasWorkspaceRole returns a boolean if a field has been set.

### SetWorkspaceRoleNil

`func (o *AdminProvisionGroupToWorkspaceIn) SetWorkspaceRoleNil(b bool)`

 SetWorkspaceRoleNil sets the value for WorkspaceRole to be an explicit nil

### UnsetWorkspaceRole
`func (o *AdminProvisionGroupToWorkspaceIn) UnsetWorkspaceRole()`

UnsetWorkspaceRole ensures that no value is present for WorkspaceRole, not even an explicit nil
### GetWorkspaceRoleName

`func (o *AdminProvisionGroupToWorkspaceIn) GetWorkspaceRoleName() string`

GetWorkspaceRoleName returns the WorkspaceRoleName field if non-nil, zero value otherwise.

### GetWorkspaceRoleNameOk

`func (o *AdminProvisionGroupToWorkspaceIn) GetWorkspaceRoleNameOk() (*string, bool)`

GetWorkspaceRoleNameOk returns a tuple with the WorkspaceRoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceRoleName

`func (o *AdminProvisionGroupToWorkspaceIn) SetWorkspaceRoleName(v string)`

SetWorkspaceRoleName sets WorkspaceRoleName field to given value.

### HasWorkspaceRoleName

`func (o *AdminProvisionGroupToWorkspaceIn) HasWorkspaceRoleName() bool`

HasWorkspaceRoleName returns a boolean if a field has been set.

### SetWorkspaceRoleNameNil

`func (o *AdminProvisionGroupToWorkspaceIn) SetWorkspaceRoleNameNil(b bool)`

 SetWorkspaceRoleNameNil sets the value for WorkspaceRoleName to be an explicit nil

### UnsetWorkspaceRoleName
`func (o *AdminProvisionGroupToWorkspaceIn) UnsetWorkspaceRoleName()`

UnsetWorkspaceRoleName ensures that no value is present for WorkspaceRoleName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


