# WhoamiUserPrincipal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrincipalType** | Pointer to **string** | Discriminator identifying a human user principal. | [optional] [default to "user"]
**Id** | **string** | The user&#39;s unique identifier. | 
**Email** | Pointer to **NullableString** | The user&#39;s email address, when known. | [optional] 
**FirstName** | Pointer to **NullableString** | The user&#39;s first name, when known. | [optional] 
**LastName** | Pointer to **NullableString** | The user&#39;s last name, when known. | [optional] 

## Methods

### NewWhoamiUserPrincipal

`func NewWhoamiUserPrincipal(id string, ) *WhoamiUserPrincipal`

NewWhoamiUserPrincipal instantiates a new WhoamiUserPrincipal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhoamiUserPrincipalWithDefaults

`func NewWhoamiUserPrincipalWithDefaults() *WhoamiUserPrincipal`

NewWhoamiUserPrincipalWithDefaults instantiates a new WhoamiUserPrincipal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrincipalType

`func (o *WhoamiUserPrincipal) GetPrincipalType() string`

GetPrincipalType returns the PrincipalType field if non-nil, zero value otherwise.

### GetPrincipalTypeOk

`func (o *WhoamiUserPrincipal) GetPrincipalTypeOk() (*string, bool)`

GetPrincipalTypeOk returns a tuple with the PrincipalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalType

`func (o *WhoamiUserPrincipal) SetPrincipalType(v string)`

SetPrincipalType sets PrincipalType field to given value.

### HasPrincipalType

`func (o *WhoamiUserPrincipal) HasPrincipalType() bool`

HasPrincipalType returns a boolean if a field has been set.

### GetId

`func (o *WhoamiUserPrincipal) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WhoamiUserPrincipal) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WhoamiUserPrincipal) SetId(v string)`

SetId sets Id field to given value.


### GetEmail

`func (o *WhoamiUserPrincipal) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *WhoamiUserPrincipal) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *WhoamiUserPrincipal) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *WhoamiUserPrincipal) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *WhoamiUserPrincipal) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *WhoamiUserPrincipal) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetFirstName

`func (o *WhoamiUserPrincipal) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *WhoamiUserPrincipal) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *WhoamiUserPrincipal) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *WhoamiUserPrincipal) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *WhoamiUserPrincipal) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *WhoamiUserPrincipal) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *WhoamiUserPrincipal) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *WhoamiUserPrincipal) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *WhoamiUserPrincipal) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *WhoamiUserPrincipal) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *WhoamiUserPrincipal) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *WhoamiUserPrincipal) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


