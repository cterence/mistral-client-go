# Principal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrincipalType** | Pointer to **string** | Discriminator identifying a service-account principal. | [optional] [default to "service_account"]
**Id** | **string** | The user&#39;s unique identifier. | 
**Email** | Pointer to **string** | The user&#39;s email address, when known. | [optional] 
**FirstName** | Pointer to **string** | The user&#39;s first name, when known. | [optional] 
**LastName** | Pointer to **string** | The user&#39;s last name, when known. | [optional] 
**ServiceAccountId** | **string** | The service account&#39;s unique identifier. | 
**Name** | Pointer to **string** | The service account&#39;s display name, when known. | [optional] 

## Methods

### NewPrincipal

`func NewPrincipal(id string, serviceAccountId string, ) *Principal`

NewPrincipal instantiates a new Principal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrincipalWithDefaults

`func NewPrincipalWithDefaults() *Principal`

NewPrincipalWithDefaults instantiates a new Principal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrincipalType

`func (o *Principal) GetPrincipalType() string`

GetPrincipalType returns the PrincipalType field if non-nil, zero value otherwise.

### GetPrincipalTypeOk

`func (o *Principal) GetPrincipalTypeOk() (*string, bool)`

GetPrincipalTypeOk returns a tuple with the PrincipalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalType

`func (o *Principal) SetPrincipalType(v string)`

SetPrincipalType sets PrincipalType field to given value.

### HasPrincipalType

`func (o *Principal) HasPrincipalType() bool`

HasPrincipalType returns a boolean if a field has been set.

### GetId

`func (o *Principal) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Principal) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Principal) SetId(v string)`

SetId sets Id field to given value.


### GetEmail

`func (o *Principal) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Principal) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Principal) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Principal) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFirstName

`func (o *Principal) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *Principal) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *Principal) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *Principal) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *Principal) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *Principal) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *Principal) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *Principal) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetServiceAccountId

`func (o *Principal) GetServiceAccountId() string`

GetServiceAccountId returns the ServiceAccountId field if non-nil, zero value otherwise.

### GetServiceAccountIdOk

`func (o *Principal) GetServiceAccountIdOk() (*string, bool)`

GetServiceAccountIdOk returns a tuple with the ServiceAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountId

`func (o *Principal) SetServiceAccountId(v string)`

SetServiceAccountId sets ServiceAccountId field to given value.


### GetName

`func (o *Principal) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Principal) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Principal) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Principal) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


