# AuthenticationMethodCreateOrUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MethodType** | [**AuthenticationType**](AuthenticationType.md) | The type of authentication method (e.g. oauth2, bearer, none). | 
**Headers** | Pointer to [**[]ConnectorAuthenticationHeader**](ConnectorAuthenticationHeader.md) | Set of headers to connect to the connector | [optional] 
**GlobalHeaders** | Pointer to [**map[string]GlobalHeaderValue**](GlobalHeaderValue.md) | Connector-wide headers keyed by header name, applied to every credential. Secret values are encrypted at rest and never returned in clear. | [optional] [default to {}]
**Oauth2MetadataSecrets** | Pointer to [**NullableOauth2MetadataSecrets**](Oauth2MetadataSecrets.md) | New OAuth2 client credentials (client_id and client_secret). | [optional] 
**Oauth2ServerMetadata** | Pointer to [**NullableExtendedOAuthServerMetadata**](ExtendedOAuthServerMetadata.md) | New OAuth2 authorization server metadata. | [optional] 

## Methods

### NewAuthenticationMethodCreateOrUpdateRequest

`func NewAuthenticationMethodCreateOrUpdateRequest(methodType AuthenticationType, ) *AuthenticationMethodCreateOrUpdateRequest`

NewAuthenticationMethodCreateOrUpdateRequest instantiates a new AuthenticationMethodCreateOrUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthenticationMethodCreateOrUpdateRequestWithDefaults

`func NewAuthenticationMethodCreateOrUpdateRequestWithDefaults() *AuthenticationMethodCreateOrUpdateRequest`

NewAuthenticationMethodCreateOrUpdateRequestWithDefaults instantiates a new AuthenticationMethodCreateOrUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethodType

`func (o *AuthenticationMethodCreateOrUpdateRequest) GetMethodType() AuthenticationType`

GetMethodType returns the MethodType field if non-nil, zero value otherwise.

### GetMethodTypeOk

`func (o *AuthenticationMethodCreateOrUpdateRequest) GetMethodTypeOk() (*AuthenticationType, bool)`

GetMethodTypeOk returns a tuple with the MethodType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodType

`func (o *AuthenticationMethodCreateOrUpdateRequest) SetMethodType(v AuthenticationType)`

SetMethodType sets MethodType field to given value.


### GetHeaders

`func (o *AuthenticationMethodCreateOrUpdateRequest) GetHeaders() []ConnectorAuthenticationHeader`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *AuthenticationMethodCreateOrUpdateRequest) GetHeadersOk() (*[]ConnectorAuthenticationHeader, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *AuthenticationMethodCreateOrUpdateRequest) SetHeaders(v []ConnectorAuthenticationHeader)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *AuthenticationMethodCreateOrUpdateRequest) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### SetHeadersNil

`func (o *AuthenticationMethodCreateOrUpdateRequest) SetHeadersNil(b bool)`

 SetHeadersNil sets the value for Headers to be an explicit nil

### UnsetHeaders
`func (o *AuthenticationMethodCreateOrUpdateRequest) UnsetHeaders()`

UnsetHeaders ensures that no value is present for Headers, not even an explicit nil
### GetGlobalHeaders

`func (o *AuthenticationMethodCreateOrUpdateRequest) GetGlobalHeaders() map[string]GlobalHeaderValue`

GetGlobalHeaders returns the GlobalHeaders field if non-nil, zero value otherwise.

### GetGlobalHeadersOk

`func (o *AuthenticationMethodCreateOrUpdateRequest) GetGlobalHeadersOk() (*map[string]GlobalHeaderValue, bool)`

GetGlobalHeadersOk returns a tuple with the GlobalHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalHeaders

`func (o *AuthenticationMethodCreateOrUpdateRequest) SetGlobalHeaders(v map[string]GlobalHeaderValue)`

SetGlobalHeaders sets GlobalHeaders field to given value.

### HasGlobalHeaders

`func (o *AuthenticationMethodCreateOrUpdateRequest) HasGlobalHeaders() bool`

HasGlobalHeaders returns a boolean if a field has been set.

### GetOauth2MetadataSecrets

`func (o *AuthenticationMethodCreateOrUpdateRequest) GetOauth2MetadataSecrets() Oauth2MetadataSecrets`

GetOauth2MetadataSecrets returns the Oauth2MetadataSecrets field if non-nil, zero value otherwise.

### GetOauth2MetadataSecretsOk

`func (o *AuthenticationMethodCreateOrUpdateRequest) GetOauth2MetadataSecretsOk() (*Oauth2MetadataSecrets, bool)`

GetOauth2MetadataSecretsOk returns a tuple with the Oauth2MetadataSecrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2MetadataSecrets

`func (o *AuthenticationMethodCreateOrUpdateRequest) SetOauth2MetadataSecrets(v Oauth2MetadataSecrets)`

SetOauth2MetadataSecrets sets Oauth2MetadataSecrets field to given value.

### HasOauth2MetadataSecrets

`func (o *AuthenticationMethodCreateOrUpdateRequest) HasOauth2MetadataSecrets() bool`

HasOauth2MetadataSecrets returns a boolean if a field has been set.

### SetOauth2MetadataSecretsNil

`func (o *AuthenticationMethodCreateOrUpdateRequest) SetOauth2MetadataSecretsNil(b bool)`

 SetOauth2MetadataSecretsNil sets the value for Oauth2MetadataSecrets to be an explicit nil

### UnsetOauth2MetadataSecrets
`func (o *AuthenticationMethodCreateOrUpdateRequest) UnsetOauth2MetadataSecrets()`

UnsetOauth2MetadataSecrets ensures that no value is present for Oauth2MetadataSecrets, not even an explicit nil
### GetOauth2ServerMetadata

`func (o *AuthenticationMethodCreateOrUpdateRequest) GetOauth2ServerMetadata() ExtendedOAuthServerMetadata`

GetOauth2ServerMetadata returns the Oauth2ServerMetadata field if non-nil, zero value otherwise.

### GetOauth2ServerMetadataOk

`func (o *AuthenticationMethodCreateOrUpdateRequest) GetOauth2ServerMetadataOk() (*ExtendedOAuthServerMetadata, bool)`

GetOauth2ServerMetadataOk returns a tuple with the Oauth2ServerMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2ServerMetadata

`func (o *AuthenticationMethodCreateOrUpdateRequest) SetOauth2ServerMetadata(v ExtendedOAuthServerMetadata)`

SetOauth2ServerMetadata sets Oauth2ServerMetadata field to given value.

### HasOauth2ServerMetadata

`func (o *AuthenticationMethodCreateOrUpdateRequest) HasOauth2ServerMetadata() bool`

HasOauth2ServerMetadata returns a boolean if a field has been set.

### SetOauth2ServerMetadataNil

`func (o *AuthenticationMethodCreateOrUpdateRequest) SetOauth2ServerMetadataNil(b bool)`

 SetOauth2ServerMetadataNil sets the value for Oauth2ServerMetadata to be an explicit nil

### UnsetOauth2ServerMetadata
`func (o *AuthenticationMethodCreateOrUpdateRequest) UnsetOauth2ServerMetadata()`

UnsetOauth2ServerMetadata ensures that no value is present for Oauth2ServerMetadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


