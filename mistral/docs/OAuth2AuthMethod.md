# OAuth2AuthMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MethodType** | Pointer to **string** |  | [optional] [default to "oauth2"]
**Headers** | Pointer to [**[]ConnectorAuthenticationHeader**](ConnectorAuthenticationHeader.md) | Optional headers sent with requests for this auth method. | [optional] 
**GrantType** | **string** |  | 
**AuthData** | [**AuthData**](AuthData.md) | OAuth2 client credentials (client_id required; client_secret optional for PKCE flows). | 
**Oauth2ServerMetadata** | [**ExtendedOAuthServerMetadata**](ExtendedOAuthServerMetadata.md) | OAuth2 authorization server metadata (endpoints, scopes_supported, ...). | 
**TokenEndpointAuthMethod** | [**OAuth2TokenEndpointAuthMethod**](OAuth2TokenEndpointAuthMethod.md) |  | 

## Methods

### NewOAuth2AuthMethod

`func NewOAuth2AuthMethod(grantType string, authData AuthData, oauth2ServerMetadata ExtendedOAuthServerMetadata, tokenEndpointAuthMethod OAuth2TokenEndpointAuthMethod, ) *OAuth2AuthMethod`

NewOAuth2AuthMethod instantiates a new OAuth2AuthMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuth2AuthMethodWithDefaults

`func NewOAuth2AuthMethodWithDefaults() *OAuth2AuthMethod`

NewOAuth2AuthMethodWithDefaults instantiates a new OAuth2AuthMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethodType

`func (o *OAuth2AuthMethod) GetMethodType() string`

GetMethodType returns the MethodType field if non-nil, zero value otherwise.

### GetMethodTypeOk

`func (o *OAuth2AuthMethod) GetMethodTypeOk() (*string, bool)`

GetMethodTypeOk returns a tuple with the MethodType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodType

`func (o *OAuth2AuthMethod) SetMethodType(v string)`

SetMethodType sets MethodType field to given value.

### HasMethodType

`func (o *OAuth2AuthMethod) HasMethodType() bool`

HasMethodType returns a boolean if a field has been set.

### GetHeaders

`func (o *OAuth2AuthMethod) GetHeaders() []ConnectorAuthenticationHeader`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *OAuth2AuthMethod) GetHeadersOk() (*[]ConnectorAuthenticationHeader, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *OAuth2AuthMethod) SetHeaders(v []ConnectorAuthenticationHeader)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *OAuth2AuthMethod) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetGrantType

`func (o *OAuth2AuthMethod) GetGrantType() string`

GetGrantType returns the GrantType field if non-nil, zero value otherwise.

### GetGrantTypeOk

`func (o *OAuth2AuthMethod) GetGrantTypeOk() (*string, bool)`

GetGrantTypeOk returns a tuple with the GrantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantType

`func (o *OAuth2AuthMethod) SetGrantType(v string)`

SetGrantType sets GrantType field to given value.


### GetAuthData

`func (o *OAuth2AuthMethod) GetAuthData() AuthData`

GetAuthData returns the AuthData field if non-nil, zero value otherwise.

### GetAuthDataOk

`func (o *OAuth2AuthMethod) GetAuthDataOk() (*AuthData, bool)`

GetAuthDataOk returns a tuple with the AuthData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthData

`func (o *OAuth2AuthMethod) SetAuthData(v AuthData)`

SetAuthData sets AuthData field to given value.


### GetOauth2ServerMetadata

`func (o *OAuth2AuthMethod) GetOauth2ServerMetadata() ExtendedOAuthServerMetadata`

GetOauth2ServerMetadata returns the Oauth2ServerMetadata field if non-nil, zero value otherwise.

### GetOauth2ServerMetadataOk

`func (o *OAuth2AuthMethod) GetOauth2ServerMetadataOk() (*ExtendedOAuthServerMetadata, bool)`

GetOauth2ServerMetadataOk returns a tuple with the Oauth2ServerMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2ServerMetadata

`func (o *OAuth2AuthMethod) SetOauth2ServerMetadata(v ExtendedOAuthServerMetadata)`

SetOauth2ServerMetadata sets Oauth2ServerMetadata field to given value.


### GetTokenEndpointAuthMethod

`func (o *OAuth2AuthMethod) GetTokenEndpointAuthMethod() OAuth2TokenEndpointAuthMethod`

GetTokenEndpointAuthMethod returns the TokenEndpointAuthMethod field if non-nil, zero value otherwise.

### GetTokenEndpointAuthMethodOk

`func (o *OAuth2AuthMethod) GetTokenEndpointAuthMethodOk() (*OAuth2TokenEndpointAuthMethod, bool)`

GetTokenEndpointAuthMethodOk returns a tuple with the TokenEndpointAuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEndpointAuthMethod

`func (o *OAuth2AuthMethod) SetTokenEndpointAuthMethod(v OAuth2TokenEndpointAuthMethod)`

SetTokenEndpointAuthMethod sets TokenEndpointAuthMethod field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


