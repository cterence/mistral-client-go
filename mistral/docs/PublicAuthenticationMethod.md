# PublicAuthenticationMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MethodType** | [**AuthenticationType**](AuthenticationType.md) |  | 
**Headers** | Pointer to [**[]ConnectorAuthenticationHeader**](ConnectorAuthenticationHeader.md) |  | [optional] 
**GlobalHeaders** | Pointer to [**map[string]GlobalHeaderValue**](GlobalHeaderValue.md) |  | [optional] [default to {}]
**HasDefaultCredentials** | **bool** |  | 
**GrantType** | Pointer to [**NullableOAuth2GrantType**](OAuth2GrantType.md) |  | [optional] 
**TokenEndpointAuthMethod** | Pointer to [**NullableOAuth2TokenEndpointAuthMethod**](OAuth2TokenEndpointAuthMethod.md) |  | [optional] 
**Oauth2ServerMetadata** | Pointer to [**NullableExtendedOAuthServerMetadata**](ExtendedOAuthServerMetadata.md) |  | [optional] 

## Methods

### NewPublicAuthenticationMethod

`func NewPublicAuthenticationMethod(methodType AuthenticationType, hasDefaultCredentials bool, ) *PublicAuthenticationMethod`

NewPublicAuthenticationMethod instantiates a new PublicAuthenticationMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicAuthenticationMethodWithDefaults

`func NewPublicAuthenticationMethodWithDefaults() *PublicAuthenticationMethod`

NewPublicAuthenticationMethodWithDefaults instantiates a new PublicAuthenticationMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethodType

`func (o *PublicAuthenticationMethod) GetMethodType() AuthenticationType`

GetMethodType returns the MethodType field if non-nil, zero value otherwise.

### GetMethodTypeOk

`func (o *PublicAuthenticationMethod) GetMethodTypeOk() (*AuthenticationType, bool)`

GetMethodTypeOk returns a tuple with the MethodType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodType

`func (o *PublicAuthenticationMethod) SetMethodType(v AuthenticationType)`

SetMethodType sets MethodType field to given value.


### GetHeaders

`func (o *PublicAuthenticationMethod) GetHeaders() []ConnectorAuthenticationHeader`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *PublicAuthenticationMethod) GetHeadersOk() (*[]ConnectorAuthenticationHeader, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *PublicAuthenticationMethod) SetHeaders(v []ConnectorAuthenticationHeader)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *PublicAuthenticationMethod) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### SetHeadersNil

`func (o *PublicAuthenticationMethod) SetHeadersNil(b bool)`

 SetHeadersNil sets the value for Headers to be an explicit nil

### UnsetHeaders
`func (o *PublicAuthenticationMethod) UnsetHeaders()`

UnsetHeaders ensures that no value is present for Headers, not even an explicit nil
### GetGlobalHeaders

`func (o *PublicAuthenticationMethod) GetGlobalHeaders() map[string]GlobalHeaderValue`

GetGlobalHeaders returns the GlobalHeaders field if non-nil, zero value otherwise.

### GetGlobalHeadersOk

`func (o *PublicAuthenticationMethod) GetGlobalHeadersOk() (*map[string]GlobalHeaderValue, bool)`

GetGlobalHeadersOk returns a tuple with the GlobalHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalHeaders

`func (o *PublicAuthenticationMethod) SetGlobalHeaders(v map[string]GlobalHeaderValue)`

SetGlobalHeaders sets GlobalHeaders field to given value.

### HasGlobalHeaders

`func (o *PublicAuthenticationMethod) HasGlobalHeaders() bool`

HasGlobalHeaders returns a boolean if a field has been set.

### GetHasDefaultCredentials

`func (o *PublicAuthenticationMethod) GetHasDefaultCredentials() bool`

GetHasDefaultCredentials returns the HasDefaultCredentials field if non-nil, zero value otherwise.

### GetHasDefaultCredentialsOk

`func (o *PublicAuthenticationMethod) GetHasDefaultCredentialsOk() (*bool, bool)`

GetHasDefaultCredentialsOk returns a tuple with the HasDefaultCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDefaultCredentials

`func (o *PublicAuthenticationMethod) SetHasDefaultCredentials(v bool)`

SetHasDefaultCredentials sets HasDefaultCredentials field to given value.


### GetGrantType

`func (o *PublicAuthenticationMethod) GetGrantType() OAuth2GrantType`

GetGrantType returns the GrantType field if non-nil, zero value otherwise.

### GetGrantTypeOk

`func (o *PublicAuthenticationMethod) GetGrantTypeOk() (*OAuth2GrantType, bool)`

GetGrantTypeOk returns a tuple with the GrantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantType

`func (o *PublicAuthenticationMethod) SetGrantType(v OAuth2GrantType)`

SetGrantType sets GrantType field to given value.

### HasGrantType

`func (o *PublicAuthenticationMethod) HasGrantType() bool`

HasGrantType returns a boolean if a field has been set.

### SetGrantTypeNil

`func (o *PublicAuthenticationMethod) SetGrantTypeNil(b bool)`

 SetGrantTypeNil sets the value for GrantType to be an explicit nil

### UnsetGrantType
`func (o *PublicAuthenticationMethod) UnsetGrantType()`

UnsetGrantType ensures that no value is present for GrantType, not even an explicit nil
### GetTokenEndpointAuthMethod

`func (o *PublicAuthenticationMethod) GetTokenEndpointAuthMethod() OAuth2TokenEndpointAuthMethod`

GetTokenEndpointAuthMethod returns the TokenEndpointAuthMethod field if non-nil, zero value otherwise.

### GetTokenEndpointAuthMethodOk

`func (o *PublicAuthenticationMethod) GetTokenEndpointAuthMethodOk() (*OAuth2TokenEndpointAuthMethod, bool)`

GetTokenEndpointAuthMethodOk returns a tuple with the TokenEndpointAuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEndpointAuthMethod

`func (o *PublicAuthenticationMethod) SetTokenEndpointAuthMethod(v OAuth2TokenEndpointAuthMethod)`

SetTokenEndpointAuthMethod sets TokenEndpointAuthMethod field to given value.

### HasTokenEndpointAuthMethod

`func (o *PublicAuthenticationMethod) HasTokenEndpointAuthMethod() bool`

HasTokenEndpointAuthMethod returns a boolean if a field has been set.

### SetTokenEndpointAuthMethodNil

`func (o *PublicAuthenticationMethod) SetTokenEndpointAuthMethodNil(b bool)`

 SetTokenEndpointAuthMethodNil sets the value for TokenEndpointAuthMethod to be an explicit nil

### UnsetTokenEndpointAuthMethod
`func (o *PublicAuthenticationMethod) UnsetTokenEndpointAuthMethod()`

UnsetTokenEndpointAuthMethod ensures that no value is present for TokenEndpointAuthMethod, not even an explicit nil
### GetOauth2ServerMetadata

`func (o *PublicAuthenticationMethod) GetOauth2ServerMetadata() ExtendedOAuthServerMetadata`

GetOauth2ServerMetadata returns the Oauth2ServerMetadata field if non-nil, zero value otherwise.

### GetOauth2ServerMetadataOk

`func (o *PublicAuthenticationMethod) GetOauth2ServerMetadataOk() (*ExtendedOAuthServerMetadata, bool)`

GetOauth2ServerMetadataOk returns a tuple with the Oauth2ServerMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2ServerMetadata

`func (o *PublicAuthenticationMethod) SetOauth2ServerMetadata(v ExtendedOAuthServerMetadata)`

SetOauth2ServerMetadata sets Oauth2ServerMetadata field to given value.

### HasOauth2ServerMetadata

`func (o *PublicAuthenticationMethod) HasOauth2ServerMetadata() bool`

HasOauth2ServerMetadata returns a boolean if a field has been set.

### SetOauth2ServerMetadataNil

`func (o *PublicAuthenticationMethod) SetOauth2ServerMetadataNil(b bool)`

 SetOauth2ServerMetadataNil sets the value for Oauth2ServerMetadata to be an explicit nil

### UnsetOauth2ServerMetadata
`func (o *PublicAuthenticationMethod) UnsetOauth2ServerMetadata()`

UnsetOauth2ServerMetadata ensures that no value is present for Oauth2ServerMetadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


