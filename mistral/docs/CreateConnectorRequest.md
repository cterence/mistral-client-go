# CreateConnectorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Headers** | Pointer to **map[string]interface{}** | Output JSON schema of the query&#39;s model | [optional] 
**GlobalHeaders** | Pointer to [**map[string]GlobalHeaderValue**](GlobalHeaderValue.md) | Optional connector-wide headers, keyed by header name, set at creation and applied to every credential. Secret values are encrypted at rest and never returned in clear. | [optional] [default to {}]
**Oauth2ServerMetadata** | Pointer to [**NullableExtendedOAuthServerMetadata**](ExtendedOAuthServerMetadata.md) | Optional OAuth2 authorization server metadata (authorization_endpoint, token_endpoint, etc.). When provided, skips .well-known discovery and uses these endpoints directly. | [optional] 
**Oauth2ServerMetadataUrl** | Pointer to **NullableString** | Optional URL to fetch OAuth2 authorization server metadata from (RFC 8414). When provided, the metadata is fetched from this URL and used instead of .well-known discovery. Mutually exclusive with oauth2_server_metadata. | [optional] 
**Name** | **string** | The name of the connector. Should be 64 char length maximum, alphanumeric, only underscores/dashes. | 
**Title** | Pointer to **NullableString** | Optional human-readable title for the connector. | [optional] 
**Description** | **string** | The description of the connector. | 
**IconUrl** | Pointer to **NullableString** | The optional url of the icon you want to associate to the connector. | [optional] 
**Server** | **string** | The URL of the connector server. | 
**Visibility** | Pointer to [**PublicResourceVisibility**](PublicResourceVisibility.md) | Visibility of the connector. Use &#39;shared_workspace&#39; for workspace scoped connectors, or &#39;private&#39; for private connectors. | [optional] [default to PUBLICRESOURCEVISIBILITY_PRIVATE]
**AuthMethods** | Pointer to [**[]ConnectorMCPPublicUpdateAuthMethodsInnerAnyOf**](ConnectorMCPPublicUpdateAuthMethodsInnerAnyOf.md) | Authentication methods supported by the connector. | [optional] 
**SystemPrompt** | Pointer to **NullableString** | Optional system prompt for the connector. | [optional] 
**Protocol** | Pointer to **string** | Protocol of the connector. Use &#39;mcp&#39; for MCP connectors. | [optional] [default to "mcp"]

## Methods

### NewCreateConnectorRequest

`func NewCreateConnectorRequest(name string, description string, server string, ) *CreateConnectorRequest`

NewCreateConnectorRequest instantiates a new CreateConnectorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateConnectorRequestWithDefaults

`func NewCreateConnectorRequestWithDefaults() *CreateConnectorRequest`

NewCreateConnectorRequestWithDefaults instantiates a new CreateConnectorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeaders

`func (o *CreateConnectorRequest) GetHeaders() map[string]interface{}`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *CreateConnectorRequest) GetHeadersOk() (*map[string]interface{}, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *CreateConnectorRequest) SetHeaders(v map[string]interface{})`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *CreateConnectorRequest) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### SetHeadersNil

`func (o *CreateConnectorRequest) SetHeadersNil(b bool)`

 SetHeadersNil sets the value for Headers to be an explicit nil

### UnsetHeaders
`func (o *CreateConnectorRequest) UnsetHeaders()`

UnsetHeaders ensures that no value is present for Headers, not even an explicit nil
### GetGlobalHeaders

`func (o *CreateConnectorRequest) GetGlobalHeaders() map[string]GlobalHeaderValue`

GetGlobalHeaders returns the GlobalHeaders field if non-nil, zero value otherwise.

### GetGlobalHeadersOk

`func (o *CreateConnectorRequest) GetGlobalHeadersOk() (*map[string]GlobalHeaderValue, bool)`

GetGlobalHeadersOk returns a tuple with the GlobalHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalHeaders

`func (o *CreateConnectorRequest) SetGlobalHeaders(v map[string]GlobalHeaderValue)`

SetGlobalHeaders sets GlobalHeaders field to given value.

### HasGlobalHeaders

`func (o *CreateConnectorRequest) HasGlobalHeaders() bool`

HasGlobalHeaders returns a boolean if a field has been set.

### GetOauth2ServerMetadata

`func (o *CreateConnectorRequest) GetOauth2ServerMetadata() ExtendedOAuthServerMetadata`

GetOauth2ServerMetadata returns the Oauth2ServerMetadata field if non-nil, zero value otherwise.

### GetOauth2ServerMetadataOk

`func (o *CreateConnectorRequest) GetOauth2ServerMetadataOk() (*ExtendedOAuthServerMetadata, bool)`

GetOauth2ServerMetadataOk returns a tuple with the Oauth2ServerMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2ServerMetadata

`func (o *CreateConnectorRequest) SetOauth2ServerMetadata(v ExtendedOAuthServerMetadata)`

SetOauth2ServerMetadata sets Oauth2ServerMetadata field to given value.

### HasOauth2ServerMetadata

`func (o *CreateConnectorRequest) HasOauth2ServerMetadata() bool`

HasOauth2ServerMetadata returns a boolean if a field has been set.

### SetOauth2ServerMetadataNil

`func (o *CreateConnectorRequest) SetOauth2ServerMetadataNil(b bool)`

 SetOauth2ServerMetadataNil sets the value for Oauth2ServerMetadata to be an explicit nil

### UnsetOauth2ServerMetadata
`func (o *CreateConnectorRequest) UnsetOauth2ServerMetadata()`

UnsetOauth2ServerMetadata ensures that no value is present for Oauth2ServerMetadata, not even an explicit nil
### GetOauth2ServerMetadataUrl

`func (o *CreateConnectorRequest) GetOauth2ServerMetadataUrl() string`

GetOauth2ServerMetadataUrl returns the Oauth2ServerMetadataUrl field if non-nil, zero value otherwise.

### GetOauth2ServerMetadataUrlOk

`func (o *CreateConnectorRequest) GetOauth2ServerMetadataUrlOk() (*string, bool)`

GetOauth2ServerMetadataUrlOk returns a tuple with the Oauth2ServerMetadataUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2ServerMetadataUrl

`func (o *CreateConnectorRequest) SetOauth2ServerMetadataUrl(v string)`

SetOauth2ServerMetadataUrl sets Oauth2ServerMetadataUrl field to given value.

### HasOauth2ServerMetadataUrl

`func (o *CreateConnectorRequest) HasOauth2ServerMetadataUrl() bool`

HasOauth2ServerMetadataUrl returns a boolean if a field has been set.

### SetOauth2ServerMetadataUrlNil

`func (o *CreateConnectorRequest) SetOauth2ServerMetadataUrlNil(b bool)`

 SetOauth2ServerMetadataUrlNil sets the value for Oauth2ServerMetadataUrl to be an explicit nil

### UnsetOauth2ServerMetadataUrl
`func (o *CreateConnectorRequest) UnsetOauth2ServerMetadataUrl()`

UnsetOauth2ServerMetadataUrl ensures that no value is present for Oauth2ServerMetadataUrl, not even an explicit nil
### GetName

`func (o *CreateConnectorRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateConnectorRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateConnectorRequest) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *CreateConnectorRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateConnectorRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateConnectorRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CreateConnectorRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *CreateConnectorRequest) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *CreateConnectorRequest) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *CreateConnectorRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateConnectorRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateConnectorRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetIconUrl

`func (o *CreateConnectorRequest) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *CreateConnectorRequest) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *CreateConnectorRequest) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *CreateConnectorRequest) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### SetIconUrlNil

`func (o *CreateConnectorRequest) SetIconUrlNil(b bool)`

 SetIconUrlNil sets the value for IconUrl to be an explicit nil

### UnsetIconUrl
`func (o *CreateConnectorRequest) UnsetIconUrl()`

UnsetIconUrl ensures that no value is present for IconUrl, not even an explicit nil
### GetServer

`func (o *CreateConnectorRequest) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *CreateConnectorRequest) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *CreateConnectorRequest) SetServer(v string)`

SetServer sets Server field to given value.


### GetVisibility

`func (o *CreateConnectorRequest) GetVisibility() PublicResourceVisibility`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *CreateConnectorRequest) GetVisibilityOk() (*PublicResourceVisibility, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *CreateConnectorRequest) SetVisibility(v PublicResourceVisibility)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *CreateConnectorRequest) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetAuthMethods

`func (o *CreateConnectorRequest) GetAuthMethods() []ConnectorMCPPublicUpdateAuthMethodsInnerAnyOf`

GetAuthMethods returns the AuthMethods field if non-nil, zero value otherwise.

### GetAuthMethodsOk

`func (o *CreateConnectorRequest) GetAuthMethodsOk() (*[]ConnectorMCPPublicUpdateAuthMethodsInnerAnyOf, bool)`

GetAuthMethodsOk returns a tuple with the AuthMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMethods

`func (o *CreateConnectorRequest) SetAuthMethods(v []ConnectorMCPPublicUpdateAuthMethodsInnerAnyOf)`

SetAuthMethods sets AuthMethods field to given value.

### HasAuthMethods

`func (o *CreateConnectorRequest) HasAuthMethods() bool`

HasAuthMethods returns a boolean if a field has been set.

### SetAuthMethodsNil

`func (o *CreateConnectorRequest) SetAuthMethodsNil(b bool)`

 SetAuthMethodsNil sets the value for AuthMethods to be an explicit nil

### UnsetAuthMethods
`func (o *CreateConnectorRequest) UnsetAuthMethods()`

UnsetAuthMethods ensures that no value is present for AuthMethods, not even an explicit nil
### GetSystemPrompt

`func (o *CreateConnectorRequest) GetSystemPrompt() string`

GetSystemPrompt returns the SystemPrompt field if non-nil, zero value otherwise.

### GetSystemPromptOk

`func (o *CreateConnectorRequest) GetSystemPromptOk() (*string, bool)`

GetSystemPromptOk returns a tuple with the SystemPrompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemPrompt

`func (o *CreateConnectorRequest) SetSystemPrompt(v string)`

SetSystemPrompt sets SystemPrompt field to given value.

### HasSystemPrompt

`func (o *CreateConnectorRequest) HasSystemPrompt() bool`

HasSystemPrompt returns a boolean if a field has been set.

### SetSystemPromptNil

`func (o *CreateConnectorRequest) SetSystemPromptNil(b bool)`

 SetSystemPromptNil sets the value for SystemPrompt to be an explicit nil

### UnsetSystemPrompt
`func (o *CreateConnectorRequest) UnsetSystemPrompt()`

UnsetSystemPrompt ensures that no value is present for SystemPrompt, not even an explicit nil
### GetProtocol

`func (o *CreateConnectorRequest) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *CreateConnectorRequest) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *CreateConnectorRequest) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *CreateConnectorRequest) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


