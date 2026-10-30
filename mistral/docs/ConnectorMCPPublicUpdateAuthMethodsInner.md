# ConnectorMCPPublicUpdateAuthMethodsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MethodType** | [**AuthenticationType**](AuthenticationType.md) | The type of authentication method (e.g. oauth2, bearer, none). | 
**Headers** | Pointer to [**[]ConnectorAuthenticationHeader**](ConnectorAuthenticationHeader.md) | Set of headers to connect to the connector | [optional] 
**GlobalHeaders** | Pointer to [**map[string]GlobalHeaderValue**](GlobalHeaderValue.md) | Connector-wide headers keyed by header name, applied to every credential. Secret values are encrypted at rest and never returned in clear. | [optional] [default to {}]
**Oauth2MetadataSecrets** | Pointer to [**Oauth2MetadataSecrets**](Oauth2MetadataSecrets.md) | New OAuth2 client credentials (client_id and client_secret). | [optional] 
**Oauth2ServerMetadata** | Pointer to [**ExtendedOAuthServerMetadata**](ExtendedOAuthServerMetadata.md) | New OAuth2 authorization server metadata. | [optional] 

## Methods

### NewConnectorMCPPublicUpdateAuthMethodsInner

`func NewConnectorMCPPublicUpdateAuthMethodsInner(methodType AuthenticationType, ) *ConnectorMCPPublicUpdateAuthMethodsInner`

NewConnectorMCPPublicUpdateAuthMethodsInner instantiates a new ConnectorMCPPublicUpdateAuthMethodsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorMCPPublicUpdateAuthMethodsInnerWithDefaults

`func NewConnectorMCPPublicUpdateAuthMethodsInnerWithDefaults() *ConnectorMCPPublicUpdateAuthMethodsInner`

NewConnectorMCPPublicUpdateAuthMethodsInnerWithDefaults instantiates a new ConnectorMCPPublicUpdateAuthMethodsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethodType

`func (o *ConnectorMCPPublicUpdateAuthMethodsInner) GetMethodType() AuthenticationType`

GetMethodType returns the MethodType field if non-nil, zero value otherwise.

### GetMethodTypeOk

`func (o *ConnectorMCPPublicUpdateAuthMethodsInner) GetMethodTypeOk() (*AuthenticationType, bool)`

GetMethodTypeOk returns a tuple with the MethodType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodType

`func (o *ConnectorMCPPublicUpdateAuthMethodsInner) SetMethodType(v AuthenticationType)`

SetMethodType sets MethodType field to given value.


### GetHeaders

`func (o *ConnectorMCPPublicUpdateAuthMethodsInner) GetHeaders() []ConnectorAuthenticationHeader`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *ConnectorMCPPublicUpdateAuthMethodsInner) GetHeadersOk() (*[]ConnectorAuthenticationHeader, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *ConnectorMCPPublicUpdateAuthMethodsInner) SetHeaders(v []ConnectorAuthenticationHeader)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *ConnectorMCPPublicUpdateAuthMethodsInner) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetGlobalHeaders

`func (o *ConnectorMCPPublicUpdateAuthMethodsInner) GetGlobalHeaders() map[string]GlobalHeaderValue`

GetGlobalHeaders returns the GlobalHeaders field if non-nil, zero value otherwise.

### GetGlobalHeadersOk

`func (o *ConnectorMCPPublicUpdateAuthMethodsInner) GetGlobalHeadersOk() (*map[string]GlobalHeaderValue, bool)`

GetGlobalHeadersOk returns a tuple with the GlobalHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalHeaders

`func (o *ConnectorMCPPublicUpdateAuthMethodsInner) SetGlobalHeaders(v map[string]GlobalHeaderValue)`

SetGlobalHeaders sets GlobalHeaders field to given value.

### HasGlobalHeaders

`func (o *ConnectorMCPPublicUpdateAuthMethodsInner) HasGlobalHeaders() bool`

HasGlobalHeaders returns a boolean if a field has been set.

### GetOauth2MetadataSecrets

`func (o *ConnectorMCPPublicUpdateAuthMethodsInner) GetOauth2MetadataSecrets() Oauth2MetadataSecrets`

GetOauth2MetadataSecrets returns the Oauth2MetadataSecrets field if non-nil, zero value otherwise.

### GetOauth2MetadataSecretsOk

`func (o *ConnectorMCPPublicUpdateAuthMethodsInner) GetOauth2MetadataSecretsOk() (*Oauth2MetadataSecrets, bool)`

GetOauth2MetadataSecretsOk returns a tuple with the Oauth2MetadataSecrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2MetadataSecrets

`func (o *ConnectorMCPPublicUpdateAuthMethodsInner) SetOauth2MetadataSecrets(v Oauth2MetadataSecrets)`

SetOauth2MetadataSecrets sets Oauth2MetadataSecrets field to given value.

### HasOauth2MetadataSecrets

`func (o *ConnectorMCPPublicUpdateAuthMethodsInner) HasOauth2MetadataSecrets() bool`

HasOauth2MetadataSecrets returns a boolean if a field has been set.

### GetOauth2ServerMetadata

`func (o *ConnectorMCPPublicUpdateAuthMethodsInner) GetOauth2ServerMetadata() ExtendedOAuthServerMetadata`

GetOauth2ServerMetadata returns the Oauth2ServerMetadata field if non-nil, zero value otherwise.

### GetOauth2ServerMetadataOk

`func (o *ConnectorMCPPublicUpdateAuthMethodsInner) GetOauth2ServerMetadataOk() (*ExtendedOAuthServerMetadata, bool)`

GetOauth2ServerMetadataOk returns a tuple with the Oauth2ServerMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2ServerMetadata

`func (o *ConnectorMCPPublicUpdateAuthMethodsInner) SetOauth2ServerMetadata(v ExtendedOAuthServerMetadata)`

SetOauth2ServerMetadata sets Oauth2ServerMetadata field to given value.

### HasOauth2ServerMetadata

`func (o *ConnectorMCPPublicUpdateAuthMethodsInner) HasOauth2ServerMetadata() bool`

HasOauth2ServerMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


