# OrganizationMemberCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoleNames** | Pointer to **[]string** | Simplified role names to assign. Mutually exclusive with &#39;roles&#39;. | [optional] 
**Roles** | Pointer to [**NullableRoles1**](Roles1.md) |  | [optional] 
**RoleName** | Pointer to **NullableString** | Deprecated single role name, kept for backward compatibility. Mutually exclusive with &#39;role&#39;. | [optional] 
**Role** | Pointer to **NullableString** | Deprecated legacy single role value, kept for backward compatibility. Mutually exclusive with &#39;role_name&#39;. | [optional] 
**Email** | **string** | Email address of the user to create. | 
**FirstName** | **string** | First name of the user to create. | 
**LastName** | **string** | Last name of the user to create. | 
**SubscriptionTypes** | Pointer to [**[]PlanType**](PlanType.md) | Product seats to assign to the user. | [optional] 

## Methods

### NewOrganizationMemberCreate

`func NewOrganizationMemberCreate(email string, firstName string, lastName string, ) *OrganizationMemberCreate`

NewOrganizationMemberCreate instantiates a new OrganizationMemberCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationMemberCreateWithDefaults

`func NewOrganizationMemberCreateWithDefaults() *OrganizationMemberCreate`

NewOrganizationMemberCreateWithDefaults instantiates a new OrganizationMemberCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoleNames

`func (o *OrganizationMemberCreate) GetRoleNames() []string`

GetRoleNames returns the RoleNames field if non-nil, zero value otherwise.

### GetRoleNamesOk

`func (o *OrganizationMemberCreate) GetRoleNamesOk() (*[]string, bool)`

GetRoleNamesOk returns a tuple with the RoleNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleNames

`func (o *OrganizationMemberCreate) SetRoleNames(v []string)`

SetRoleNames sets RoleNames field to given value.

### HasRoleNames

`func (o *OrganizationMemberCreate) HasRoleNames() bool`

HasRoleNames returns a boolean if a field has been set.

### SetRoleNamesNil

`func (o *OrganizationMemberCreate) SetRoleNamesNil(b bool)`

 SetRoleNamesNil sets the value for RoleNames to be an explicit nil

### UnsetRoleNames
`func (o *OrganizationMemberCreate) UnsetRoleNames()`

UnsetRoleNames ensures that no value is present for RoleNames, not even an explicit nil
### GetRoles

`func (o *OrganizationMemberCreate) GetRoles() Roles1`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *OrganizationMemberCreate) GetRolesOk() (*Roles1, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *OrganizationMemberCreate) SetRoles(v Roles1)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *OrganizationMemberCreate) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### SetRolesNil

`func (o *OrganizationMemberCreate) SetRolesNil(b bool)`

 SetRolesNil sets the value for Roles to be an explicit nil

### UnsetRoles
`func (o *OrganizationMemberCreate) UnsetRoles()`

UnsetRoles ensures that no value is present for Roles, not even an explicit nil
### GetRoleName

`func (o *OrganizationMemberCreate) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *OrganizationMemberCreate) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *OrganizationMemberCreate) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.

### HasRoleName

`func (o *OrganizationMemberCreate) HasRoleName() bool`

HasRoleName returns a boolean if a field has been set.

### SetRoleNameNil

`func (o *OrganizationMemberCreate) SetRoleNameNil(b bool)`

 SetRoleNameNil sets the value for RoleName to be an explicit nil

### UnsetRoleName
`func (o *OrganizationMemberCreate) UnsetRoleName()`

UnsetRoleName ensures that no value is present for RoleName, not even an explicit nil
### GetRole

`func (o *OrganizationMemberCreate) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *OrganizationMemberCreate) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *OrganizationMemberCreate) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *OrganizationMemberCreate) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *OrganizationMemberCreate) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *OrganizationMemberCreate) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetEmail

`func (o *OrganizationMemberCreate) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *OrganizationMemberCreate) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *OrganizationMemberCreate) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetFirstName

`func (o *OrganizationMemberCreate) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *OrganizationMemberCreate) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *OrganizationMemberCreate) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *OrganizationMemberCreate) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *OrganizationMemberCreate) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *OrganizationMemberCreate) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetSubscriptionTypes

`func (o *OrganizationMemberCreate) GetSubscriptionTypes() []PlanType`

GetSubscriptionTypes returns the SubscriptionTypes field if non-nil, zero value otherwise.

### GetSubscriptionTypesOk

`func (o *OrganizationMemberCreate) GetSubscriptionTypesOk() (*[]PlanType, bool)`

GetSubscriptionTypesOk returns a tuple with the SubscriptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionTypes

`func (o *OrganizationMemberCreate) SetSubscriptionTypes(v []PlanType)`

SetSubscriptionTypes sets SubscriptionTypes field to given value.

### HasSubscriptionTypes

`func (o *OrganizationMemberCreate) HasSubscriptionTypes() bool`

HasSubscriptionTypes returns a boolean if a field has been set.

### SetSubscriptionTypesNil

`func (o *OrganizationMemberCreate) SetSubscriptionTypesNil(b bool)`

 SetSubscriptionTypesNil sets the value for SubscriptionTypes to be an explicit nil

### UnsetSubscriptionTypes
`func (o *OrganizationMemberCreate) UnsetSubscriptionTypes()`

UnsetSubscriptionTypes ensures that no value is present for SubscriptionTypes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


