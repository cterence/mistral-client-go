# AdminOrganizationMemberOUT

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | **string** | Organization member ID. | 
**OidId** | **NullableString** | Identity provider ID for the member. | 
**Name** | **NullableString** | Name of the Organization member. | 
**Email** | **NullableString** | Email address of the Organization member. | 
**IsSsoOutsider** | Pointer to **NullableBool** | Whether the member is outside the SSO domain. | [optional] 
**Workspaces** | Pointer to [**[]MemberWorkspaceInfo**](MemberWorkspaceInfo.md) | Workspaces the member belongs to. | [optional] 
**Subscriptions** | Pointer to [**[]MemberSubscriptionOUT**](MemberSubscriptionOUT.md) | Subscriptions assigned to the member. | [optional] [default to {}]
**RawRoles** | [**RawRoles**](RawRoles.md) |  | 
**RawRole** | [**RawRole**](RawRole.md) |  | 
**SubscriptionTypes** | Pointer to **[]string** | Product seats assigned to the member. | [optional] 
**CreatedAt** | **time.Time** | Date the member was added to the Organization. | 

## Methods

### NewAdminOrganizationMemberOUT

`func NewAdminOrganizationMemberOUT(uuid string, oidId NullableString, name NullableString, email NullableString, rawRoles RawRoles, rawRole RawRole, createdAt time.Time, ) *AdminOrganizationMemberOUT`

NewAdminOrganizationMemberOUT instantiates a new AdminOrganizationMemberOUT object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminOrganizationMemberOUTWithDefaults

`func NewAdminOrganizationMemberOUTWithDefaults() *AdminOrganizationMemberOUT`

NewAdminOrganizationMemberOUTWithDefaults instantiates a new AdminOrganizationMemberOUT object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *AdminOrganizationMemberOUT) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *AdminOrganizationMemberOUT) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *AdminOrganizationMemberOUT) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetOidId

`func (o *AdminOrganizationMemberOUT) GetOidId() string`

GetOidId returns the OidId field if non-nil, zero value otherwise.

### GetOidIdOk

`func (o *AdminOrganizationMemberOUT) GetOidIdOk() (*string, bool)`

GetOidIdOk returns a tuple with the OidId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidId

`func (o *AdminOrganizationMemberOUT) SetOidId(v string)`

SetOidId sets OidId field to given value.


### SetOidIdNil

`func (o *AdminOrganizationMemberOUT) SetOidIdNil(b bool)`

 SetOidIdNil sets the value for OidId to be an explicit nil

### UnsetOidId
`func (o *AdminOrganizationMemberOUT) UnsetOidId()`

UnsetOidId ensures that no value is present for OidId, not even an explicit nil
### GetName

`func (o *AdminOrganizationMemberOUT) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdminOrganizationMemberOUT) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdminOrganizationMemberOUT) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *AdminOrganizationMemberOUT) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AdminOrganizationMemberOUT) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetEmail

`func (o *AdminOrganizationMemberOUT) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AdminOrganizationMemberOUT) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AdminOrganizationMemberOUT) SetEmail(v string)`

SetEmail sets Email field to given value.


### SetEmailNil

`func (o *AdminOrganizationMemberOUT) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *AdminOrganizationMemberOUT) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetIsSsoOutsider

`func (o *AdminOrganizationMemberOUT) GetIsSsoOutsider() bool`

GetIsSsoOutsider returns the IsSsoOutsider field if non-nil, zero value otherwise.

### GetIsSsoOutsiderOk

`func (o *AdminOrganizationMemberOUT) GetIsSsoOutsiderOk() (*bool, bool)`

GetIsSsoOutsiderOk returns a tuple with the IsSsoOutsider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSsoOutsider

`func (o *AdminOrganizationMemberOUT) SetIsSsoOutsider(v bool)`

SetIsSsoOutsider sets IsSsoOutsider field to given value.

### HasIsSsoOutsider

`func (o *AdminOrganizationMemberOUT) HasIsSsoOutsider() bool`

HasIsSsoOutsider returns a boolean if a field has been set.

### SetIsSsoOutsiderNil

`func (o *AdminOrganizationMemberOUT) SetIsSsoOutsiderNil(b bool)`

 SetIsSsoOutsiderNil sets the value for IsSsoOutsider to be an explicit nil

### UnsetIsSsoOutsider
`func (o *AdminOrganizationMemberOUT) UnsetIsSsoOutsider()`

UnsetIsSsoOutsider ensures that no value is present for IsSsoOutsider, not even an explicit nil
### GetWorkspaces

`func (o *AdminOrganizationMemberOUT) GetWorkspaces() []MemberWorkspaceInfo`

GetWorkspaces returns the Workspaces field if non-nil, zero value otherwise.

### GetWorkspacesOk

`func (o *AdminOrganizationMemberOUT) GetWorkspacesOk() (*[]MemberWorkspaceInfo, bool)`

GetWorkspacesOk returns a tuple with the Workspaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaces

`func (o *AdminOrganizationMemberOUT) SetWorkspaces(v []MemberWorkspaceInfo)`

SetWorkspaces sets Workspaces field to given value.

### HasWorkspaces

`func (o *AdminOrganizationMemberOUT) HasWorkspaces() bool`

HasWorkspaces returns a boolean if a field has been set.

### SetWorkspacesNil

`func (o *AdminOrganizationMemberOUT) SetWorkspacesNil(b bool)`

 SetWorkspacesNil sets the value for Workspaces to be an explicit nil

### UnsetWorkspaces
`func (o *AdminOrganizationMemberOUT) UnsetWorkspaces()`

UnsetWorkspaces ensures that no value is present for Workspaces, not even an explicit nil
### GetSubscriptions

`func (o *AdminOrganizationMemberOUT) GetSubscriptions() []MemberSubscriptionOUT`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *AdminOrganizationMemberOUT) GetSubscriptionsOk() (*[]MemberSubscriptionOUT, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *AdminOrganizationMemberOUT) SetSubscriptions(v []MemberSubscriptionOUT)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *AdminOrganizationMemberOUT) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.

### GetRawRoles

`func (o *AdminOrganizationMemberOUT) GetRawRoles() RawRoles`

GetRawRoles returns the RawRoles field if non-nil, zero value otherwise.

### GetRawRolesOk

`func (o *AdminOrganizationMemberOUT) GetRawRolesOk() (*RawRoles, bool)`

GetRawRolesOk returns a tuple with the RawRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawRoles

`func (o *AdminOrganizationMemberOUT) SetRawRoles(v RawRoles)`

SetRawRoles sets RawRoles field to given value.


### GetRawRole

`func (o *AdminOrganizationMemberOUT) GetRawRole() RawRole`

GetRawRole returns the RawRole field if non-nil, zero value otherwise.

### GetRawRoleOk

`func (o *AdminOrganizationMemberOUT) GetRawRoleOk() (*RawRole, bool)`

GetRawRoleOk returns a tuple with the RawRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawRole

`func (o *AdminOrganizationMemberOUT) SetRawRole(v RawRole)`

SetRawRole sets RawRole field to given value.


### GetSubscriptionTypes

`func (o *AdminOrganizationMemberOUT) GetSubscriptionTypes() []string`

GetSubscriptionTypes returns the SubscriptionTypes field if non-nil, zero value otherwise.

### GetSubscriptionTypesOk

`func (o *AdminOrganizationMemberOUT) GetSubscriptionTypesOk() (*[]string, bool)`

GetSubscriptionTypesOk returns a tuple with the SubscriptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionTypes

`func (o *AdminOrganizationMemberOUT) SetSubscriptionTypes(v []string)`

SetSubscriptionTypes sets SubscriptionTypes field to given value.

### HasSubscriptionTypes

`func (o *AdminOrganizationMemberOUT) HasSubscriptionTypes() bool`

HasSubscriptionTypes returns a boolean if a field has been set.

### SetSubscriptionTypesNil

`func (o *AdminOrganizationMemberOUT) SetSubscriptionTypesNil(b bool)`

 SetSubscriptionTypesNil sets the value for SubscriptionTypes to be an explicit nil

### UnsetSubscriptionTypes
`func (o *AdminOrganizationMemberOUT) UnsetSubscriptionTypes()`

UnsetSubscriptionTypes ensures that no value is present for SubscriptionTypes, not even an explicit nil
### GetCreatedAt

`func (o *AdminOrganizationMemberOUT) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AdminOrganizationMemberOUT) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AdminOrganizationMemberOUT) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


