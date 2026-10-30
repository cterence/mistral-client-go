# AuthData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** |  | 
**ClientSecret** | Pointer to **NullableString** |  | [optional] 
**ClientScopes** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAuthData

`func NewAuthData(clientId string, ) *AuthData`

NewAuthData instantiates a new AuthData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthDataWithDefaults

`func NewAuthDataWithDefaults() *AuthData`

NewAuthDataWithDefaults instantiates a new AuthData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *AuthData) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *AuthData) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *AuthData) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *AuthData) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *AuthData) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *AuthData) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *AuthData) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### SetClientSecretNil

`func (o *AuthData) SetClientSecretNil(b bool)`

 SetClientSecretNil sets the value for ClientSecret to be an explicit nil

### UnsetClientSecret
`func (o *AuthData) UnsetClientSecret()`

UnsetClientSecret ensures that no value is present for ClientSecret, not even an explicit nil
### GetClientScopes

`func (o *AuthData) GetClientScopes() []string`

GetClientScopes returns the ClientScopes field if non-nil, zero value otherwise.

### GetClientScopesOk

`func (o *AuthData) GetClientScopesOk() (*[]string, bool)`

GetClientScopesOk returns a tuple with the ClientScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientScopes

`func (o *AuthData) SetClientScopes(v []string)`

SetClientScopes sets ClientScopes field to given value.

### HasClientScopes

`func (o *AuthData) HasClientScopes() bool`

HasClientScopes returns a boolean if a field has been set.

### SetClientScopesNil

`func (o *AuthData) SetClientScopesNil(b bool)`

 SetClientScopesNil sets the value for ClientScopes to be an explicit nil

### UnsetClientScopes
`func (o *AuthData) UnsetClientScopes()`

UnsetClientScopes ensures that no value is present for ClientScopes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


