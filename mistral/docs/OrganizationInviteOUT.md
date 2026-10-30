# OrganizationInviteOUT

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | **string** | Organization invitation ID. | 
**RawRoles** | [**RawRoles**](RawRoles.md) |  | 
**RawRole** | [**RawRole**](RawRole.md) |  | 
**Email** | **string** | Email address invited to the Organization. | 
**Expired** | **bool** | Whether the invitation has expired. | 
**CreatedAt** | **time.Time** | Time when the invitation was created. | 
**WorkspaceUuids** | Pointer to **[]string** | Workspace IDs the invited user will join. | [optional] [default to {}]

## Methods

### NewOrganizationInviteOUT

`func NewOrganizationInviteOUT(uuid string, rawRoles RawRoles, rawRole RawRole, email string, expired bool, createdAt time.Time, ) *OrganizationInviteOUT`

NewOrganizationInviteOUT instantiates a new OrganizationInviteOUT object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationInviteOUTWithDefaults

`func NewOrganizationInviteOUTWithDefaults() *OrganizationInviteOUT`

NewOrganizationInviteOUTWithDefaults instantiates a new OrganizationInviteOUT object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *OrganizationInviteOUT) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *OrganizationInviteOUT) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *OrganizationInviteOUT) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetRawRoles

`func (o *OrganizationInviteOUT) GetRawRoles() RawRoles`

GetRawRoles returns the RawRoles field if non-nil, zero value otherwise.

### GetRawRolesOk

`func (o *OrganizationInviteOUT) GetRawRolesOk() (*RawRoles, bool)`

GetRawRolesOk returns a tuple with the RawRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawRoles

`func (o *OrganizationInviteOUT) SetRawRoles(v RawRoles)`

SetRawRoles sets RawRoles field to given value.


### GetRawRole

`func (o *OrganizationInviteOUT) GetRawRole() RawRole`

GetRawRole returns the RawRole field if non-nil, zero value otherwise.

### GetRawRoleOk

`func (o *OrganizationInviteOUT) GetRawRoleOk() (*RawRole, bool)`

GetRawRoleOk returns a tuple with the RawRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawRole

`func (o *OrganizationInviteOUT) SetRawRole(v RawRole)`

SetRawRole sets RawRole field to given value.


### GetEmail

`func (o *OrganizationInviteOUT) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *OrganizationInviteOUT) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *OrganizationInviteOUT) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetExpired

`func (o *OrganizationInviteOUT) GetExpired() bool`

GetExpired returns the Expired field if non-nil, zero value otherwise.

### GetExpiredOk

`func (o *OrganizationInviteOUT) GetExpiredOk() (*bool, bool)`

GetExpiredOk returns a tuple with the Expired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpired

`func (o *OrganizationInviteOUT) SetExpired(v bool)`

SetExpired sets Expired field to given value.


### GetCreatedAt

`func (o *OrganizationInviteOUT) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrganizationInviteOUT) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrganizationInviteOUT) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetWorkspaceUuids

`func (o *OrganizationInviteOUT) GetWorkspaceUuids() []string`

GetWorkspaceUuids returns the WorkspaceUuids field if non-nil, zero value otherwise.

### GetWorkspaceUuidsOk

`func (o *OrganizationInviteOUT) GetWorkspaceUuidsOk() (*[]string, bool)`

GetWorkspaceUuidsOk returns a tuple with the WorkspaceUuids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceUuids

`func (o *OrganizationInviteOUT) SetWorkspaceUuids(v []string)`

SetWorkspaceUuids sets WorkspaceUuids field to given value.

### HasWorkspaceUuids

`func (o *OrganizationInviteOUT) HasWorkspaceUuids() bool`

HasWorkspaceUuids returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


