# OAuth2AuthorizationCodeAuthMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MethodType** | Pointer to **string** |  | [optional] [default to "oauth2"]
**Headers** | Pointer to [**[]ConnectorAuthenticationHeader**](ConnectorAuthenticationHeader.md) | Optional headers sent with requests for this auth method. | [optional] 
**GrantType** | Pointer to **string** |  | [optional] [default to "authorization_code"]
**AuthData** | [**AuthData**](AuthData.md) | OAuth2 client credentials (client_id required; client_secret optional for PKCE flows). | 
**Oauth2ServerMetadata** | Pointer to [**NullableExtendedOAuthServerMetadata**](ExtendedOAuthServerMetadata.md) | OAuth2 authorization server metadata (endpoints, scopes_supported, ...). | [optional] 

## Methods

### NewOAuth2AuthorizationCodeAuthMethod

`func NewOAuth2AuthorizationCodeAuthMethod(authData AuthData, ) *OAuth2AuthorizationCodeAuthMethod`

NewOAuth2AuthorizationCodeAuthMethod instantiates a new OAuth2AuthorizationCodeAuthMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuth2AuthorizationCodeAuthMethodWithDefaults

`func NewOAuth2AuthorizationCodeAuthMethodWithDefaults() *OAuth2AuthorizationCodeAuthMethod`

NewOAuth2AuthorizationCodeAuthMethodWithDefaults instantiates a new OAuth2AuthorizationCodeAuthMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethodType

`func (o *OAuth2AuthorizationCodeAuthMethod) GetMethodType() string`

GetMethodType returns the MethodType field if non-nil, zero value otherwise.

### GetMethodTypeOk

`func (o *OAuth2AuthorizationCodeAuthMethod) GetMethodTypeOk() (*string, bool)`

GetMethodTypeOk returns a tuple with the MethodType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodType

`func (o *OAuth2AuthorizationCodeAuthMethod) SetMethodType(v string)`

SetMethodType sets MethodType field to given value.

### HasMethodType

`func (o *OAuth2AuthorizationCodeAuthMethod) HasMethodType() bool`

HasMethodType returns a boolean if a field has been set.

### GetHeaders

`func (o *OAuth2AuthorizationCodeAuthMethod) GetHeaders() []ConnectorAuthenticationHeader`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *OAuth2AuthorizationCodeAuthMethod) GetHeadersOk() (*[]ConnectorAuthenticationHeader, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *OAuth2AuthorizationCodeAuthMethod) SetHeaders(v []ConnectorAuthenticationHeader)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *OAuth2AuthorizationCodeAuthMethod) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### SetHeadersNil

`func (o *OAuth2AuthorizationCodeAuthMethod) SetHeadersNil(b bool)`

 SetHeadersNil sets the value for Headers to be an explicit nil

### UnsetHeaders
`func (o *OAuth2AuthorizationCodeAuthMethod) UnsetHeaders()`

UnsetHeaders ensures that no value is present for Headers, not even an explicit nil
### GetGrantType

`func (o *OAuth2AuthorizationCodeAuthMethod) GetGrantType() string`

GetGrantType returns the GrantType field if non-nil, zero value otherwise.

### GetGrantTypeOk

`func (o *OAuth2AuthorizationCodeAuthMethod) GetGrantTypeOk() (*string, bool)`

GetGrantTypeOk returns a tuple with the GrantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantType

`func (o *OAuth2AuthorizationCodeAuthMethod) SetGrantType(v string)`

SetGrantType sets GrantType field to given value.

### HasGrantType

`func (o *OAuth2AuthorizationCodeAuthMethod) HasGrantType() bool`

HasGrantType returns a boolean if a field has been set.

### GetAuthData

`func (o *OAuth2AuthorizationCodeAuthMethod) GetAuthData() AuthData`

GetAuthData returns the AuthData field if non-nil, zero value otherwise.

### GetAuthDataOk

`func (o *OAuth2AuthorizationCodeAuthMethod) GetAuthDataOk() (*AuthData, bool)`

GetAuthDataOk returns a tuple with the AuthData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthData

`func (o *OAuth2AuthorizationCodeAuthMethod) SetAuthData(v AuthData)`

SetAuthData sets AuthData field to given value.


### GetOauth2ServerMetadata

`func (o *OAuth2AuthorizationCodeAuthMethod) GetOauth2ServerMetadata() ExtendedOAuthServerMetadata`

GetOauth2ServerMetadata returns the Oauth2ServerMetadata field if non-nil, zero value otherwise.

### GetOauth2ServerMetadataOk

`func (o *OAuth2AuthorizationCodeAuthMethod) GetOauth2ServerMetadataOk() (*ExtendedOAuthServerMetadata, bool)`

GetOauth2ServerMetadataOk returns a tuple with the Oauth2ServerMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2ServerMetadata

`func (o *OAuth2AuthorizationCodeAuthMethod) SetOauth2ServerMetadata(v ExtendedOAuthServerMetadata)`

SetOauth2ServerMetadata sets Oauth2ServerMetadata field to given value.

### HasOauth2ServerMetadata

`func (o *OAuth2AuthorizationCodeAuthMethod) HasOauth2ServerMetadata() bool`

HasOauth2ServerMetadata returns a boolean if a field has been set.

### SetOauth2ServerMetadataNil

`func (o *OAuth2AuthorizationCodeAuthMethod) SetOauth2ServerMetadataNil(b bool)`

 SetOauth2ServerMetadataNil sets the value for Oauth2ServerMetadata to be an explicit nil

### UnsetOauth2ServerMetadata
`func (o *OAuth2AuthorizationCodeAuthMethod) UnsetOauth2ServerMetadata()`

UnsetOauth2ServerMetadata ensures that no value is present for Oauth2ServerMetadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


