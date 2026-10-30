# UserIdentity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Email** | **NullableString** |  | 
**FirstName** | **NullableString** |  | 
**LastName** | **NullableString** |  | 
**Workspace** | Pointer to [**NullableUserIdentityWorkspace**](UserIdentityWorkspace.md) |  | [optional] 
**Organization** | Pointer to [**NullableUserIdentityOrganization**](UserIdentityOrganization.md) |  | [optional] 
**ApiKey** | Pointer to [**NullableUserIdentityApiKey**](UserIdentityApiKey.md) |  | [optional] 

## Methods

### NewUserIdentity

`func NewUserIdentity(id string, email NullableString, firstName NullableString, lastName NullableString, ) *UserIdentity`

NewUserIdentity instantiates a new UserIdentity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserIdentityWithDefaults

`func NewUserIdentityWithDefaults() *UserIdentity`

NewUserIdentityWithDefaults instantiates a new UserIdentity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserIdentity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserIdentity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserIdentity) SetId(v string)`

SetId sets Id field to given value.


### GetEmail

`func (o *UserIdentity) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserIdentity) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserIdentity) SetEmail(v string)`

SetEmail sets Email field to given value.


### SetEmailNil

`func (o *UserIdentity) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *UserIdentity) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetFirstName

`func (o *UserIdentity) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UserIdentity) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UserIdentity) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### SetFirstNameNil

`func (o *UserIdentity) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *UserIdentity) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *UserIdentity) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UserIdentity) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UserIdentity) SetLastName(v string)`

SetLastName sets LastName field to given value.


### SetLastNameNil

`func (o *UserIdentity) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *UserIdentity) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetWorkspace

`func (o *UserIdentity) GetWorkspace() UserIdentityWorkspace`

GetWorkspace returns the Workspace field if non-nil, zero value otherwise.

### GetWorkspaceOk

`func (o *UserIdentity) GetWorkspaceOk() (*UserIdentityWorkspace, bool)`

GetWorkspaceOk returns a tuple with the Workspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspace

`func (o *UserIdentity) SetWorkspace(v UserIdentityWorkspace)`

SetWorkspace sets Workspace field to given value.

### HasWorkspace

`func (o *UserIdentity) HasWorkspace() bool`

HasWorkspace returns a boolean if a field has been set.

### SetWorkspaceNil

`func (o *UserIdentity) SetWorkspaceNil(b bool)`

 SetWorkspaceNil sets the value for Workspace to be an explicit nil

### UnsetWorkspace
`func (o *UserIdentity) UnsetWorkspace()`

UnsetWorkspace ensures that no value is present for Workspace, not even an explicit nil
### GetOrganization

`func (o *UserIdentity) GetOrganization() UserIdentityOrganization`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *UserIdentity) GetOrganizationOk() (*UserIdentityOrganization, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *UserIdentity) SetOrganization(v UserIdentityOrganization)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *UserIdentity) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### SetOrganizationNil

`func (o *UserIdentity) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *UserIdentity) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
### GetApiKey

`func (o *UserIdentity) GetApiKey() UserIdentityApiKey`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *UserIdentity) GetApiKeyOk() (*UserIdentityApiKey, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *UserIdentity) SetApiKey(v UserIdentityApiKey)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *UserIdentity) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### SetApiKeyNil

`func (o *UserIdentity) SetApiKeyNil(b bool)`

 SetApiKeyNil sets the value for ApiKey to be an explicit nil

### UnsetApiKey
`func (o *UserIdentity) UnsetApiKey()`

UnsetApiKey ensures that no value is present for ApiKey, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


