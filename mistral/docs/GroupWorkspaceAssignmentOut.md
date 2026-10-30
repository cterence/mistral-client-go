# GroupWorkspaceAssignmentOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Roles** | [**[]WorkspaceRoleRef**](WorkspaceRoleRef.md) |  | 
**RoleUuid** | Pointer to **NullableString** | Deprecated: Workspace assignments now support multiple roles. This field only contains the UUID of the first entry in &#x60;roles&#x60;; use &#x60;roles&#x60; instead. | [optional] 
**RoleName** | Pointer to **NullableString** | Deprecated: Workspace assignments now support multiple roles. This field only contains the name of the first entry in &#x60;roles&#x60;; use &#x60;roles&#x60; instead. | [optional] 
**WorkspaceUuid** | **string** |  | 
**WorkspaceName** | **string** | Name of the Workspace. | 
**Created** | **time.Time** |  | 

## Methods

### NewGroupWorkspaceAssignmentOut

`func NewGroupWorkspaceAssignmentOut(roles []WorkspaceRoleRef, workspaceUuid string, workspaceName string, created time.Time, ) *GroupWorkspaceAssignmentOut`

NewGroupWorkspaceAssignmentOut instantiates a new GroupWorkspaceAssignmentOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupWorkspaceAssignmentOutWithDefaults

`func NewGroupWorkspaceAssignmentOutWithDefaults() *GroupWorkspaceAssignmentOut`

NewGroupWorkspaceAssignmentOutWithDefaults instantiates a new GroupWorkspaceAssignmentOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoles

`func (o *GroupWorkspaceAssignmentOut) GetRoles() []WorkspaceRoleRef`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *GroupWorkspaceAssignmentOut) GetRolesOk() (*[]WorkspaceRoleRef, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *GroupWorkspaceAssignmentOut) SetRoles(v []WorkspaceRoleRef)`

SetRoles sets Roles field to given value.


### GetRoleUuid

`func (o *GroupWorkspaceAssignmentOut) GetRoleUuid() string`

GetRoleUuid returns the RoleUuid field if non-nil, zero value otherwise.

### GetRoleUuidOk

`func (o *GroupWorkspaceAssignmentOut) GetRoleUuidOk() (*string, bool)`

GetRoleUuidOk returns a tuple with the RoleUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleUuid

`func (o *GroupWorkspaceAssignmentOut) SetRoleUuid(v string)`

SetRoleUuid sets RoleUuid field to given value.

### HasRoleUuid

`func (o *GroupWorkspaceAssignmentOut) HasRoleUuid() bool`

HasRoleUuid returns a boolean if a field has been set.

### SetRoleUuidNil

`func (o *GroupWorkspaceAssignmentOut) SetRoleUuidNil(b bool)`

 SetRoleUuidNil sets the value for RoleUuid to be an explicit nil

### UnsetRoleUuid
`func (o *GroupWorkspaceAssignmentOut) UnsetRoleUuid()`

UnsetRoleUuid ensures that no value is present for RoleUuid, not even an explicit nil
### GetRoleName

`func (o *GroupWorkspaceAssignmentOut) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *GroupWorkspaceAssignmentOut) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *GroupWorkspaceAssignmentOut) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.

### HasRoleName

`func (o *GroupWorkspaceAssignmentOut) HasRoleName() bool`

HasRoleName returns a boolean if a field has been set.

### SetRoleNameNil

`func (o *GroupWorkspaceAssignmentOut) SetRoleNameNil(b bool)`

 SetRoleNameNil sets the value for RoleName to be an explicit nil

### UnsetRoleName
`func (o *GroupWorkspaceAssignmentOut) UnsetRoleName()`

UnsetRoleName ensures that no value is present for RoleName, not even an explicit nil
### GetWorkspaceUuid

`func (o *GroupWorkspaceAssignmentOut) GetWorkspaceUuid() string`

GetWorkspaceUuid returns the WorkspaceUuid field if non-nil, zero value otherwise.

### GetWorkspaceUuidOk

`func (o *GroupWorkspaceAssignmentOut) GetWorkspaceUuidOk() (*string, bool)`

GetWorkspaceUuidOk returns a tuple with the WorkspaceUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceUuid

`func (o *GroupWorkspaceAssignmentOut) SetWorkspaceUuid(v string)`

SetWorkspaceUuid sets WorkspaceUuid field to given value.


### GetWorkspaceName

`func (o *GroupWorkspaceAssignmentOut) GetWorkspaceName() string`

GetWorkspaceName returns the WorkspaceName field if non-nil, zero value otherwise.

### GetWorkspaceNameOk

`func (o *GroupWorkspaceAssignmentOut) GetWorkspaceNameOk() (*string, bool)`

GetWorkspaceNameOk returns a tuple with the WorkspaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceName

`func (o *GroupWorkspaceAssignmentOut) SetWorkspaceName(v string)`

SetWorkspaceName sets WorkspaceName field to given value.


### GetCreated

`func (o *GroupWorkspaceAssignmentOut) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *GroupWorkspaceAssignmentOut) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *GroupWorkspaceAssignmentOut) SetCreated(v time.Time)`

SetCreated sets Created field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


