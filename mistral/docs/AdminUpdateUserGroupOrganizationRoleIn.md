# AdminUpdateUserGroupOrganizationRoleIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrganizationRole** | **string** | Organization role to assign to the group: a role name (&#x60;member&#x60;, &#x60;billing_manager&#x60;, &#x60;organization_admin&#x60;) or a role UUID. | 

## Methods

### NewAdminUpdateUserGroupOrganizationRoleIn

`func NewAdminUpdateUserGroupOrganizationRoleIn(organizationRole string, ) *AdminUpdateUserGroupOrganizationRoleIn`

NewAdminUpdateUserGroupOrganizationRoleIn instantiates a new AdminUpdateUserGroupOrganizationRoleIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminUpdateUserGroupOrganizationRoleInWithDefaults

`func NewAdminUpdateUserGroupOrganizationRoleInWithDefaults() *AdminUpdateUserGroupOrganizationRoleIn`

NewAdminUpdateUserGroupOrganizationRoleInWithDefaults instantiates a new AdminUpdateUserGroupOrganizationRoleIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganizationRole

`func (o *AdminUpdateUserGroupOrganizationRoleIn) GetOrganizationRole() string`

GetOrganizationRole returns the OrganizationRole field if non-nil, zero value otherwise.

### GetOrganizationRoleOk

`func (o *AdminUpdateUserGroupOrganizationRoleIn) GetOrganizationRoleOk() (*string, bool)`

GetOrganizationRoleOk returns a tuple with the OrganizationRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationRole

`func (o *AdminUpdateUserGroupOrganizationRoleIn) SetOrganizationRole(v string)`

SetOrganizationRole sets OrganizationRole field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


