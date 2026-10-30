# OAuth2ClientCredentialsAuthMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MethodType** | Pointer to **string** |  | [optional] [default to "oauth2"]
**Headers** | Pointer to [**[]ConnectorAuthenticationHeader**](ConnectorAuthenticationHeader.md) | Optional headers sent with requests for this auth method. | [optional] 
**GrantType** | **string** |  | 
**TokenEndpointAuthMethod** | [**OAuth2TokenEndpointAuthMethod**](OAuth2TokenEndpointAuthMethod.md) |  | 
**Oauth2ServerMetadata** | [**ExtendedOAuthServerMetadata**](ExtendedOAuthServerMetadata.md) | OAuth2 authorization server metadata (endpoints, scopes_supported, ...). | 

## Methods

### NewOAuth2ClientCredentialsAuthMethod

`func NewOAuth2ClientCredentialsAuthMethod(grantType string, tokenEndpointAuthMethod OAuth2TokenEndpointAuthMethod, oauth2ServerMetadata ExtendedOAuthServerMetadata, ) *OAuth2ClientCredentialsAuthMethod`

NewOAuth2ClientCredentialsAuthMethod instantiates a new OAuth2ClientCredentialsAuthMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuth2ClientCredentialsAuthMethodWithDefaults

`func NewOAuth2ClientCredentialsAuthMethodWithDefaults() *OAuth2ClientCredentialsAuthMethod`

NewOAuth2ClientCredentialsAuthMethodWithDefaults instantiates a new OAuth2ClientCredentialsAuthMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethodType

`func (o *OAuth2ClientCredentialsAuthMethod) GetMethodType() string`

GetMethodType returns the MethodType field if non-nil, zero value otherwise.

### GetMethodTypeOk

`func (o *OAuth2ClientCredentialsAuthMethod) GetMethodTypeOk() (*string, bool)`

GetMethodTypeOk returns a tuple with the MethodType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodType

`func (o *OAuth2ClientCredentialsAuthMethod) SetMethodType(v string)`

SetMethodType sets MethodType field to given value.

### HasMethodType

`func (o *OAuth2ClientCredentialsAuthMethod) HasMethodType() bool`

HasMethodType returns a boolean if a field has been set.

### GetHeaders

`func (o *OAuth2ClientCredentialsAuthMethod) GetHeaders() []ConnectorAuthenticationHeader`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *OAuth2ClientCredentialsAuthMethod) GetHeadersOk() (*[]ConnectorAuthenticationHeader, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *OAuth2ClientCredentialsAuthMethod) SetHeaders(v []ConnectorAuthenticationHeader)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *OAuth2ClientCredentialsAuthMethod) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### SetHeadersNil

`func (o *OAuth2ClientCredentialsAuthMethod) SetHeadersNil(b bool)`

 SetHeadersNil sets the value for Headers to be an explicit nil

### UnsetHeaders
`func (o *OAuth2ClientCredentialsAuthMethod) UnsetHeaders()`

UnsetHeaders ensures that no value is present for Headers, not even an explicit nil
### GetGrantType

`func (o *OAuth2ClientCredentialsAuthMethod) GetGrantType() string`

GetGrantType returns the GrantType field if non-nil, zero value otherwise.

### GetGrantTypeOk

`func (o *OAuth2ClientCredentialsAuthMethod) GetGrantTypeOk() (*string, bool)`

GetGrantTypeOk returns a tuple with the GrantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantType

`func (o *OAuth2ClientCredentialsAuthMethod) SetGrantType(v string)`

SetGrantType sets GrantType field to given value.


### GetTokenEndpointAuthMethod

`func (o *OAuth2ClientCredentialsAuthMethod) GetTokenEndpointAuthMethod() OAuth2TokenEndpointAuthMethod`

GetTokenEndpointAuthMethod returns the TokenEndpointAuthMethod field if non-nil, zero value otherwise.

### GetTokenEndpointAuthMethodOk

`func (o *OAuth2ClientCredentialsAuthMethod) GetTokenEndpointAuthMethodOk() (*OAuth2TokenEndpointAuthMethod, bool)`

GetTokenEndpointAuthMethodOk returns a tuple with the TokenEndpointAuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEndpointAuthMethod

`func (o *OAuth2ClientCredentialsAuthMethod) SetTokenEndpointAuthMethod(v OAuth2TokenEndpointAuthMethod)`

SetTokenEndpointAuthMethod sets TokenEndpointAuthMethod field to given value.


### GetOauth2ServerMetadata

`func (o *OAuth2ClientCredentialsAuthMethod) GetOauth2ServerMetadata() ExtendedOAuthServerMetadata`

GetOauth2ServerMetadata returns the Oauth2ServerMetadata field if non-nil, zero value otherwise.

### GetOauth2ServerMetadataOk

`func (o *OAuth2ClientCredentialsAuthMethod) GetOauth2ServerMetadataOk() (*ExtendedOAuthServerMetadata, bool)`

GetOauth2ServerMetadataOk returns a tuple with the Oauth2ServerMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2ServerMetadata

`func (o *OAuth2ClientCredentialsAuthMethod) SetOauth2ServerMetadata(v ExtendedOAuthServerMetadata)`

SetOauth2ServerMetadata sets Oauth2ServerMetadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


