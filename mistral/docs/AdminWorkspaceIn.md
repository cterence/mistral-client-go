# AdminWorkspaceIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Workspace name. | 
**Description** | Pointer to **string** | Workspace description. | [optional] [default to ""]
**Icon** | Pointer to **string** | Workspace icon. | [optional] [default to ""]
**AddAllOrgMembers** | Pointer to **bool** | Whether to add all Organization members to the Workspace. | [optional] [default to false]
**AdminUserId** | **string** | User ID to grant the Workspace Admin role to. | 

## Methods

### NewAdminWorkspaceIn

`func NewAdminWorkspaceIn(name string, adminUserId string, ) *AdminWorkspaceIn`

NewAdminWorkspaceIn instantiates a new AdminWorkspaceIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminWorkspaceInWithDefaults

`func NewAdminWorkspaceInWithDefaults() *AdminWorkspaceIn`

NewAdminWorkspaceInWithDefaults instantiates a new AdminWorkspaceIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AdminWorkspaceIn) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdminWorkspaceIn) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdminWorkspaceIn) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AdminWorkspaceIn) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AdminWorkspaceIn) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AdminWorkspaceIn) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AdminWorkspaceIn) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIcon

`func (o *AdminWorkspaceIn) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *AdminWorkspaceIn) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *AdminWorkspaceIn) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *AdminWorkspaceIn) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetAddAllOrgMembers

`func (o *AdminWorkspaceIn) GetAddAllOrgMembers() bool`

GetAddAllOrgMembers returns the AddAllOrgMembers field if non-nil, zero value otherwise.

### GetAddAllOrgMembersOk

`func (o *AdminWorkspaceIn) GetAddAllOrgMembersOk() (*bool, bool)`

GetAddAllOrgMembersOk returns a tuple with the AddAllOrgMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddAllOrgMembers

`func (o *AdminWorkspaceIn) SetAddAllOrgMembers(v bool)`

SetAddAllOrgMembers sets AddAllOrgMembers field to given value.

### HasAddAllOrgMembers

`func (o *AdminWorkspaceIn) HasAddAllOrgMembers() bool`

HasAddAllOrgMembers returns a boolean if a field has been set.

### GetAdminUserId

`func (o *AdminWorkspaceIn) GetAdminUserId() string`

GetAdminUserId returns the AdminUserId field if non-nil, zero value otherwise.

### GetAdminUserIdOk

`func (o *AdminWorkspaceIn) GetAdminUserIdOk() (*string, bool)`

GetAdminUserIdOk returns a tuple with the AdminUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminUserId

`func (o *AdminWorkspaceIn) SetAdminUserId(v string)`

SetAdminUserId sets AdminUserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


