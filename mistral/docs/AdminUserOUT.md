# AdminUserOUT

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
**RawRole** | **string** | Deprecated single organization role. Use &#39;raw_roles&#39; instead. | 
**SubscriptionTypes** | Pointer to **[]string** | Product seats assigned to the member. | [optional] 
**CreatedAt** | **time.Time** | Date the member was added to the Organization. | 
**FirstName** | **NullableString** | User first name. | 
**LastName** | **NullableString** | User last name. | 

## Methods

### NewAdminUserOUT

`func NewAdminUserOUT(uuid string, oidId NullableString, name NullableString, email NullableString, rawRoles RawRoles, rawRole string, createdAt time.Time, firstName NullableString, lastName NullableString, ) *AdminUserOUT`

NewAdminUserOUT instantiates a new AdminUserOUT object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminUserOUTWithDefaults

`func NewAdminUserOUTWithDefaults() *AdminUserOUT`

NewAdminUserOUTWithDefaults instantiates a new AdminUserOUT object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *AdminUserOUT) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *AdminUserOUT) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *AdminUserOUT) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetOidId

`func (o *AdminUserOUT) GetOidId() string`

GetOidId returns the OidId field if non-nil, zero value otherwise.

### GetOidIdOk

`func (o *AdminUserOUT) GetOidIdOk() (*string, bool)`

GetOidIdOk returns a tuple with the OidId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidId

`func (o *AdminUserOUT) SetOidId(v string)`

SetOidId sets OidId field to given value.


### SetOidIdNil

`func (o *AdminUserOUT) SetOidIdNil(b bool)`

 SetOidIdNil sets the value for OidId to be an explicit nil

### UnsetOidId
`func (o *AdminUserOUT) UnsetOidId()`

UnsetOidId ensures that no value is present for OidId, not even an explicit nil
### GetName

`func (o *AdminUserOUT) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdminUserOUT) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdminUserOUT) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *AdminUserOUT) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AdminUserOUT) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetEmail

`func (o *AdminUserOUT) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AdminUserOUT) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AdminUserOUT) SetEmail(v string)`

SetEmail sets Email field to given value.


### SetEmailNil

`func (o *AdminUserOUT) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *AdminUserOUT) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetIsSsoOutsider

`func (o *AdminUserOUT) GetIsSsoOutsider() bool`

GetIsSsoOutsider returns the IsSsoOutsider field if non-nil, zero value otherwise.

### GetIsSsoOutsiderOk

`func (o *AdminUserOUT) GetIsSsoOutsiderOk() (*bool, bool)`

GetIsSsoOutsiderOk returns a tuple with the IsSsoOutsider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSsoOutsider

`func (o *AdminUserOUT) SetIsSsoOutsider(v bool)`

SetIsSsoOutsider sets IsSsoOutsider field to given value.

### HasIsSsoOutsider

`func (o *AdminUserOUT) HasIsSsoOutsider() bool`

HasIsSsoOutsider returns a boolean if a field has been set.

### SetIsSsoOutsiderNil

`func (o *AdminUserOUT) SetIsSsoOutsiderNil(b bool)`

 SetIsSsoOutsiderNil sets the value for IsSsoOutsider to be an explicit nil

### UnsetIsSsoOutsider
`func (o *AdminUserOUT) UnsetIsSsoOutsider()`

UnsetIsSsoOutsider ensures that no value is present for IsSsoOutsider, not even an explicit nil
### GetWorkspaces

`func (o *AdminUserOUT) GetWorkspaces() []MemberWorkspaceInfo`

GetWorkspaces returns the Workspaces field if non-nil, zero value otherwise.

### GetWorkspacesOk

`func (o *AdminUserOUT) GetWorkspacesOk() (*[]MemberWorkspaceInfo, bool)`

GetWorkspacesOk returns a tuple with the Workspaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaces

`func (o *AdminUserOUT) SetWorkspaces(v []MemberWorkspaceInfo)`

SetWorkspaces sets Workspaces field to given value.

### HasWorkspaces

`func (o *AdminUserOUT) HasWorkspaces() bool`

HasWorkspaces returns a boolean if a field has been set.

### SetWorkspacesNil

`func (o *AdminUserOUT) SetWorkspacesNil(b bool)`

 SetWorkspacesNil sets the value for Workspaces to be an explicit nil

### UnsetWorkspaces
`func (o *AdminUserOUT) UnsetWorkspaces()`

UnsetWorkspaces ensures that no value is present for Workspaces, not even an explicit nil
### GetSubscriptions

`func (o *AdminUserOUT) GetSubscriptions() []MemberSubscriptionOUT`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *AdminUserOUT) GetSubscriptionsOk() (*[]MemberSubscriptionOUT, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *AdminUserOUT) SetSubscriptions(v []MemberSubscriptionOUT)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *AdminUserOUT) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.

### GetRawRoles

`func (o *AdminUserOUT) GetRawRoles() RawRoles`

GetRawRoles returns the RawRoles field if non-nil, zero value otherwise.

### GetRawRolesOk

`func (o *AdminUserOUT) GetRawRolesOk() (*RawRoles, bool)`

GetRawRolesOk returns a tuple with the RawRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawRoles

`func (o *AdminUserOUT) SetRawRoles(v RawRoles)`

SetRawRoles sets RawRoles field to given value.


### GetRawRole

`func (o *AdminUserOUT) GetRawRole() string`

GetRawRole returns the RawRole field if non-nil, zero value otherwise.

### GetRawRoleOk

`func (o *AdminUserOUT) GetRawRoleOk() (*string, bool)`

GetRawRoleOk returns a tuple with the RawRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawRole

`func (o *AdminUserOUT) SetRawRole(v string)`

SetRawRole sets RawRole field to given value.


### GetSubscriptionTypes

`func (o *AdminUserOUT) GetSubscriptionTypes() []string`

GetSubscriptionTypes returns the SubscriptionTypes field if non-nil, zero value otherwise.

### GetSubscriptionTypesOk

`func (o *AdminUserOUT) GetSubscriptionTypesOk() (*[]string, bool)`

GetSubscriptionTypesOk returns a tuple with the SubscriptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionTypes

`func (o *AdminUserOUT) SetSubscriptionTypes(v []string)`

SetSubscriptionTypes sets SubscriptionTypes field to given value.

### HasSubscriptionTypes

`func (o *AdminUserOUT) HasSubscriptionTypes() bool`

HasSubscriptionTypes returns a boolean if a field has been set.

### SetSubscriptionTypesNil

`func (o *AdminUserOUT) SetSubscriptionTypesNil(b bool)`

 SetSubscriptionTypesNil sets the value for SubscriptionTypes to be an explicit nil

### UnsetSubscriptionTypes
`func (o *AdminUserOUT) UnsetSubscriptionTypes()`

UnsetSubscriptionTypes ensures that no value is present for SubscriptionTypes, not even an explicit nil
### GetCreatedAt

`func (o *AdminUserOUT) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AdminUserOUT) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AdminUserOUT) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetFirstName

`func (o *AdminUserOUT) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *AdminUserOUT) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *AdminUserOUT) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### SetFirstNameNil

`func (o *AdminUserOUT) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *AdminUserOUT) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *AdminUserOUT) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *AdminUserOUT) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *AdminUserOUT) SetLastName(v string)`

SetLastName sets LastName field to given value.


### SetLastNameNil

`func (o *AdminUserOUT) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *AdminUserOUT) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


