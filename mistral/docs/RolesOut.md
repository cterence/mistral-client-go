# RolesOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkspaceRoles** | [**[]RoleOut**](RoleOut.md) | Workspace roles available to the Organization. | 
**OrganizationRoles** | [**[]RoleOut**](RoleOut.md) | Organization roles available to the Organization. | 

## Methods

### NewRolesOut

`func NewRolesOut(workspaceRoles []RoleOut, organizationRoles []RoleOut, ) *RolesOut`

NewRolesOut instantiates a new RolesOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRolesOutWithDefaults

`func NewRolesOutWithDefaults() *RolesOut`

NewRolesOutWithDefaults instantiates a new RolesOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkspaceRoles

`func (o *RolesOut) GetWorkspaceRoles() []RoleOut`

GetWorkspaceRoles returns the WorkspaceRoles field if non-nil, zero value otherwise.

### GetWorkspaceRolesOk

`func (o *RolesOut) GetWorkspaceRolesOk() (*[]RoleOut, bool)`

GetWorkspaceRolesOk returns a tuple with the WorkspaceRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceRoles

`func (o *RolesOut) SetWorkspaceRoles(v []RoleOut)`

SetWorkspaceRoles sets WorkspaceRoles field to given value.


### GetOrganizationRoles

`func (o *RolesOut) GetOrganizationRoles() []RoleOut`

GetOrganizationRoles returns the OrganizationRoles field if non-nil, zero value otherwise.

### GetOrganizationRolesOk

`func (o *RolesOut) GetOrganizationRolesOk() (*[]RoleOut, bool)`

GetOrganizationRolesOk returns a tuple with the OrganizationRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationRoles

`func (o *RolesOut) SetOrganizationRoles(v []RoleOut)`

SetOrganizationRoles sets OrganizationRoles field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


