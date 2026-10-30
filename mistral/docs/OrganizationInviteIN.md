# OrganizationInviteIN

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoleNames** | Pointer to **[]string** | Simplified role names to assign. Mutually exclusive with &#39;roles&#39;. | [optional] 
**Roles** | Pointer to [**NullableRoles1**](Roles1.md) |  | [optional] 
**RoleName** | Pointer to **NullableString** | Deprecated single role name, kept for backward compatibility. Mutually exclusive with &#39;role&#39;. | [optional] 
**Role** | Pointer to **NullableString** | Deprecated legacy single role value, kept for backward compatibility. Mutually exclusive with &#39;role_name&#39;. | [optional] 
**Email** | **string** | Email address, comma-separated emails, or newline-separated emails to invite. | 
**SubscriptionType** | Pointer to [**NullablePlanType**](PlanType.md) | Deprecated single product seat to assign. | [optional] 
**SubscriptionTypes** | Pointer to [**[]PlanType**](PlanType.md) | Product seats to assign to invited users. | [optional] 
**SubscriptionSeatAutomaticGranting** | Pointer to **bool** | Whether to grant subscription seats automatically. | [optional] [default to false]
**EmailLanguage** | Pointer to **NullableString** | Language used for invitation emails. | [optional] 
**WorkspaceUuids** | Pointer to **[]string** | Workspace IDs the invited users should join. | [optional] 

## Methods

### NewOrganizationInviteIN

`func NewOrganizationInviteIN(email string, ) *OrganizationInviteIN`

NewOrganizationInviteIN instantiates a new OrganizationInviteIN object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationInviteINWithDefaults

`func NewOrganizationInviteINWithDefaults() *OrganizationInviteIN`

NewOrganizationInviteINWithDefaults instantiates a new OrganizationInviteIN object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoleNames

`func (o *OrganizationInviteIN) GetRoleNames() []string`

GetRoleNames returns the RoleNames field if non-nil, zero value otherwise.

### GetRoleNamesOk

`func (o *OrganizationInviteIN) GetRoleNamesOk() (*[]string, bool)`

GetRoleNamesOk returns a tuple with the RoleNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleNames

`func (o *OrganizationInviteIN) SetRoleNames(v []string)`

SetRoleNames sets RoleNames field to given value.

### HasRoleNames

`func (o *OrganizationInviteIN) HasRoleNames() bool`

HasRoleNames returns a boolean if a field has been set.

### SetRoleNamesNil

`func (o *OrganizationInviteIN) SetRoleNamesNil(b bool)`

 SetRoleNamesNil sets the value for RoleNames to be an explicit nil

### UnsetRoleNames
`func (o *OrganizationInviteIN) UnsetRoleNames()`

UnsetRoleNames ensures that no value is present for RoleNames, not even an explicit nil
### GetRoles

`func (o *OrganizationInviteIN) GetRoles() Roles1`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *OrganizationInviteIN) GetRolesOk() (*Roles1, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *OrganizationInviteIN) SetRoles(v Roles1)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *OrganizationInviteIN) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### SetRolesNil

`func (o *OrganizationInviteIN) SetRolesNil(b bool)`

 SetRolesNil sets the value for Roles to be an explicit nil

### UnsetRoles
`func (o *OrganizationInviteIN) UnsetRoles()`

UnsetRoles ensures that no value is present for Roles, not even an explicit nil
### GetRoleName

`func (o *OrganizationInviteIN) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *OrganizationInviteIN) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *OrganizationInviteIN) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.

### HasRoleName

`func (o *OrganizationInviteIN) HasRoleName() bool`

HasRoleName returns a boolean if a field has been set.

### SetRoleNameNil

`func (o *OrganizationInviteIN) SetRoleNameNil(b bool)`

 SetRoleNameNil sets the value for RoleName to be an explicit nil

### UnsetRoleName
`func (o *OrganizationInviteIN) UnsetRoleName()`

UnsetRoleName ensures that no value is present for RoleName, not even an explicit nil
### GetRole

`func (o *OrganizationInviteIN) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *OrganizationInviteIN) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *OrganizationInviteIN) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *OrganizationInviteIN) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *OrganizationInviteIN) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *OrganizationInviteIN) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetEmail

`func (o *OrganizationInviteIN) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *OrganizationInviteIN) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *OrganizationInviteIN) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetSubscriptionType

`func (o *OrganizationInviteIN) GetSubscriptionType() PlanType`

GetSubscriptionType returns the SubscriptionType field if non-nil, zero value otherwise.

### GetSubscriptionTypeOk

`func (o *OrganizationInviteIN) GetSubscriptionTypeOk() (*PlanType, bool)`

GetSubscriptionTypeOk returns a tuple with the SubscriptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionType

`func (o *OrganizationInviteIN) SetSubscriptionType(v PlanType)`

SetSubscriptionType sets SubscriptionType field to given value.

### HasSubscriptionType

`func (o *OrganizationInviteIN) HasSubscriptionType() bool`

HasSubscriptionType returns a boolean if a field has been set.

### SetSubscriptionTypeNil

`func (o *OrganizationInviteIN) SetSubscriptionTypeNil(b bool)`

 SetSubscriptionTypeNil sets the value for SubscriptionType to be an explicit nil

### UnsetSubscriptionType
`func (o *OrganizationInviteIN) UnsetSubscriptionType()`

UnsetSubscriptionType ensures that no value is present for SubscriptionType, not even an explicit nil
### GetSubscriptionTypes

`func (o *OrganizationInviteIN) GetSubscriptionTypes() []PlanType`

GetSubscriptionTypes returns the SubscriptionTypes field if non-nil, zero value otherwise.

### GetSubscriptionTypesOk

`func (o *OrganizationInviteIN) GetSubscriptionTypesOk() (*[]PlanType, bool)`

GetSubscriptionTypesOk returns a tuple with the SubscriptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionTypes

`func (o *OrganizationInviteIN) SetSubscriptionTypes(v []PlanType)`

SetSubscriptionTypes sets SubscriptionTypes field to given value.

### HasSubscriptionTypes

`func (o *OrganizationInviteIN) HasSubscriptionTypes() bool`

HasSubscriptionTypes returns a boolean if a field has been set.

### SetSubscriptionTypesNil

`func (o *OrganizationInviteIN) SetSubscriptionTypesNil(b bool)`

 SetSubscriptionTypesNil sets the value for SubscriptionTypes to be an explicit nil

### UnsetSubscriptionTypes
`func (o *OrganizationInviteIN) UnsetSubscriptionTypes()`

UnsetSubscriptionTypes ensures that no value is present for SubscriptionTypes, not even an explicit nil
### GetSubscriptionSeatAutomaticGranting

`func (o *OrganizationInviteIN) GetSubscriptionSeatAutomaticGranting() bool`

GetSubscriptionSeatAutomaticGranting returns the SubscriptionSeatAutomaticGranting field if non-nil, zero value otherwise.

### GetSubscriptionSeatAutomaticGrantingOk

`func (o *OrganizationInviteIN) GetSubscriptionSeatAutomaticGrantingOk() (*bool, bool)`

GetSubscriptionSeatAutomaticGrantingOk returns a tuple with the SubscriptionSeatAutomaticGranting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionSeatAutomaticGranting

`func (o *OrganizationInviteIN) SetSubscriptionSeatAutomaticGranting(v bool)`

SetSubscriptionSeatAutomaticGranting sets SubscriptionSeatAutomaticGranting field to given value.

### HasSubscriptionSeatAutomaticGranting

`func (o *OrganizationInviteIN) HasSubscriptionSeatAutomaticGranting() bool`

HasSubscriptionSeatAutomaticGranting returns a boolean if a field has been set.

### GetEmailLanguage

`func (o *OrganizationInviteIN) GetEmailLanguage() string`

GetEmailLanguage returns the EmailLanguage field if non-nil, zero value otherwise.

### GetEmailLanguageOk

`func (o *OrganizationInviteIN) GetEmailLanguageOk() (*string, bool)`

GetEmailLanguageOk returns a tuple with the EmailLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailLanguage

`func (o *OrganizationInviteIN) SetEmailLanguage(v string)`

SetEmailLanguage sets EmailLanguage field to given value.

### HasEmailLanguage

`func (o *OrganizationInviteIN) HasEmailLanguage() bool`

HasEmailLanguage returns a boolean if a field has been set.

### SetEmailLanguageNil

`func (o *OrganizationInviteIN) SetEmailLanguageNil(b bool)`

 SetEmailLanguageNil sets the value for EmailLanguage to be an explicit nil

### UnsetEmailLanguage
`func (o *OrganizationInviteIN) UnsetEmailLanguage()`

UnsetEmailLanguage ensures that no value is present for EmailLanguage, not even an explicit nil
### GetWorkspaceUuids

`func (o *OrganizationInviteIN) GetWorkspaceUuids() []string`

GetWorkspaceUuids returns the WorkspaceUuids field if non-nil, zero value otherwise.

### GetWorkspaceUuidsOk

`func (o *OrganizationInviteIN) GetWorkspaceUuidsOk() (*[]string, bool)`

GetWorkspaceUuidsOk returns a tuple with the WorkspaceUuids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceUuids

`func (o *OrganizationInviteIN) SetWorkspaceUuids(v []string)`

SetWorkspaceUuids sets WorkspaceUuids field to given value.

### HasWorkspaceUuids

`func (o *OrganizationInviteIN) HasWorkspaceUuids() bool`

HasWorkspaceUuids returns a boolean if a field has been set.

### SetWorkspaceUuidsNil

`func (o *OrganizationInviteIN) SetWorkspaceUuidsNil(b bool)`

 SetWorkspaceUuidsNil sets the value for WorkspaceUuids to be an explicit nil

### UnsetWorkspaceUuids
`func (o *OrganizationInviteIN) UnsetWorkspaceUuids()`

UnsetWorkspaceUuids ensures that no value is present for WorkspaceUuids, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


