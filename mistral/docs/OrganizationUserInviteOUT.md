# OrganizationUserInviteOUT

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InviteUuid** | **string** | Organization invitation ID. | 
**Email** | **NullableString** | Email address invited to the Organization. | 
**Roles** | [**Roles2**](Roles2.md) |  | 
**Role** | **string** | Deprecated single invite role. Use &#39;roles&#39; instead. | 

## Methods

### NewOrganizationUserInviteOUT

`func NewOrganizationUserInviteOUT(inviteUuid string, email NullableString, roles Roles2, role string, ) *OrganizationUserInviteOUT`

NewOrganizationUserInviteOUT instantiates a new OrganizationUserInviteOUT object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationUserInviteOUTWithDefaults

`func NewOrganizationUserInviteOUTWithDefaults() *OrganizationUserInviteOUT`

NewOrganizationUserInviteOUTWithDefaults instantiates a new OrganizationUserInviteOUT object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInviteUuid

`func (o *OrganizationUserInviteOUT) GetInviteUuid() string`

GetInviteUuid returns the InviteUuid field if non-nil, zero value otherwise.

### GetInviteUuidOk

`func (o *OrganizationUserInviteOUT) GetInviteUuidOk() (*string, bool)`

GetInviteUuidOk returns a tuple with the InviteUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviteUuid

`func (o *OrganizationUserInviteOUT) SetInviteUuid(v string)`

SetInviteUuid sets InviteUuid field to given value.


### GetEmail

`func (o *OrganizationUserInviteOUT) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *OrganizationUserInviteOUT) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *OrganizationUserInviteOUT) SetEmail(v string)`

SetEmail sets Email field to given value.


### SetEmailNil

`func (o *OrganizationUserInviteOUT) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *OrganizationUserInviteOUT) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetRoles

`func (o *OrganizationUserInviteOUT) GetRoles() Roles2`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *OrganizationUserInviteOUT) GetRolesOk() (*Roles2, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *OrganizationUserInviteOUT) SetRoles(v Roles2)`

SetRoles sets Roles field to given value.


### GetRole

`func (o *OrganizationUserInviteOUT) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *OrganizationUserInviteOUT) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *OrganizationUserInviteOUT) SetRole(v string)`

SetRole sets Role field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


