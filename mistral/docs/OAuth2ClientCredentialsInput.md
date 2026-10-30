# OAuth2ClientCredentialsInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GrantType** | **string** |  | 
**ClientId** | **string** |  | 
**ClientSecret** | Pointer to **NullableString** |  | [optional] 
**ClientCertificate** | Pointer to **NullableString** | PEM X.509 certificate registered on the OAuth2 client, for private_key_jwt. | [optional] 
**ClientPrivateKey** | Pointer to **NullableString** | Unencrypted PEM RSA private key matching client_certificate, for private_key_jwt. | [optional] 

## Methods

### NewOAuth2ClientCredentialsInput

`func NewOAuth2ClientCredentialsInput(grantType string, clientId string, ) *OAuth2ClientCredentialsInput`

NewOAuth2ClientCredentialsInput instantiates a new OAuth2ClientCredentialsInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuth2ClientCredentialsInputWithDefaults

`func NewOAuth2ClientCredentialsInputWithDefaults() *OAuth2ClientCredentialsInput`

NewOAuth2ClientCredentialsInputWithDefaults instantiates a new OAuth2ClientCredentialsInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrantType

`func (o *OAuth2ClientCredentialsInput) GetGrantType() string`

GetGrantType returns the GrantType field if non-nil, zero value otherwise.

### GetGrantTypeOk

`func (o *OAuth2ClientCredentialsInput) GetGrantTypeOk() (*string, bool)`

GetGrantTypeOk returns a tuple with the GrantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantType

`func (o *OAuth2ClientCredentialsInput) SetGrantType(v string)`

SetGrantType sets GrantType field to given value.


### GetClientId

`func (o *OAuth2ClientCredentialsInput) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *OAuth2ClientCredentialsInput) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *OAuth2ClientCredentialsInput) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *OAuth2ClientCredentialsInput) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *OAuth2ClientCredentialsInput) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *OAuth2ClientCredentialsInput) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *OAuth2ClientCredentialsInput) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### SetClientSecretNil

`func (o *OAuth2ClientCredentialsInput) SetClientSecretNil(b bool)`

 SetClientSecretNil sets the value for ClientSecret to be an explicit nil

### UnsetClientSecret
`func (o *OAuth2ClientCredentialsInput) UnsetClientSecret()`

UnsetClientSecret ensures that no value is present for ClientSecret, not even an explicit nil
### GetClientCertificate

`func (o *OAuth2ClientCredentialsInput) GetClientCertificate() string`

GetClientCertificate returns the ClientCertificate field if non-nil, zero value otherwise.

### GetClientCertificateOk

`func (o *OAuth2ClientCredentialsInput) GetClientCertificateOk() (*string, bool)`

GetClientCertificateOk returns a tuple with the ClientCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCertificate

`func (o *OAuth2ClientCredentialsInput) SetClientCertificate(v string)`

SetClientCertificate sets ClientCertificate field to given value.

### HasClientCertificate

`func (o *OAuth2ClientCredentialsInput) HasClientCertificate() bool`

HasClientCertificate returns a boolean if a field has been set.

### SetClientCertificateNil

`func (o *OAuth2ClientCredentialsInput) SetClientCertificateNil(b bool)`

 SetClientCertificateNil sets the value for ClientCertificate to be an explicit nil

### UnsetClientCertificate
`func (o *OAuth2ClientCredentialsInput) UnsetClientCertificate()`

UnsetClientCertificate ensures that no value is present for ClientCertificate, not even an explicit nil
### GetClientPrivateKey

`func (o *OAuth2ClientCredentialsInput) GetClientPrivateKey() string`

GetClientPrivateKey returns the ClientPrivateKey field if non-nil, zero value otherwise.

### GetClientPrivateKeyOk

`func (o *OAuth2ClientCredentialsInput) GetClientPrivateKeyOk() (*string, bool)`

GetClientPrivateKeyOk returns a tuple with the ClientPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientPrivateKey

`func (o *OAuth2ClientCredentialsInput) SetClientPrivateKey(v string)`

SetClientPrivateKey sets ClientPrivateKey field to given value.

### HasClientPrivateKey

`func (o *OAuth2ClientCredentialsInput) HasClientPrivateKey() bool`

HasClientPrivateKey returns a boolean if a field has been set.

### SetClientPrivateKeyNil

`func (o *OAuth2ClientCredentialsInput) SetClientPrivateKeyNil(b bool)`

 SetClientPrivateKeyNil sets the value for ClientPrivateKey to be an explicit nil

### UnsetClientPrivateKey
`func (o *OAuth2ClientCredentialsInput) UnsetClientPrivateKey()`

UnsetClientPrivateKey ensures that no value is present for ClientPrivateKey, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


