# CreateHTTPConnectorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Headers** | Pointer to **map[string]interface{}** | Optional scoped credentials to be sent with connector requests. | [optional] 
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
**Protocol** | **string** |  | 

## Methods

### NewCreateHTTPConnectorRequest

`func NewCreateHTTPConnectorRequest(name string, description string, server string, protocol string, ) *CreateHTTPConnectorRequest`

NewCreateHTTPConnectorRequest instantiates a new CreateHTTPConnectorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateHTTPConnectorRequestWithDefaults

`func NewCreateHTTPConnectorRequestWithDefaults() *CreateHTTPConnectorRequest`

NewCreateHTTPConnectorRequestWithDefaults instantiates a new CreateHTTPConnectorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeaders

`func (o *CreateHTTPConnectorRequest) GetHeaders() map[string]interface{}`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *CreateHTTPConnectorRequest) GetHeadersOk() (*map[string]interface{}, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *CreateHTTPConnectorRequest) SetHeaders(v map[string]interface{})`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *CreateHTTPConnectorRequest) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### SetHeadersNil

`func (o *CreateHTTPConnectorRequest) SetHeadersNil(b bool)`

 SetHeadersNil sets the value for Headers to be an explicit nil

### UnsetHeaders
`func (o *CreateHTTPConnectorRequest) UnsetHeaders()`

UnsetHeaders ensures that no value is present for Headers, not even an explicit nil
### GetGlobalHeaders

`func (o *CreateHTTPConnectorRequest) GetGlobalHeaders() map[string]GlobalHeaderValue`

GetGlobalHeaders returns the GlobalHeaders field if non-nil, zero value otherwise.

### GetGlobalHeadersOk

`func (o *CreateHTTPConnectorRequest) GetGlobalHeadersOk() (*map[string]GlobalHeaderValue, bool)`

GetGlobalHeadersOk returns a tuple with the GlobalHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalHeaders

`func (o *CreateHTTPConnectorRequest) SetGlobalHeaders(v map[string]GlobalHeaderValue)`

SetGlobalHeaders sets GlobalHeaders field to given value.

### HasGlobalHeaders

`func (o *CreateHTTPConnectorRequest) HasGlobalHeaders() bool`

HasGlobalHeaders returns a boolean if a field has been set.

### GetOauth2ServerMetadata

`func (o *CreateHTTPConnectorRequest) GetOauth2ServerMetadata() ExtendedOAuthServerMetadata`

GetOauth2ServerMetadata returns the Oauth2ServerMetadata field if non-nil, zero value otherwise.

### GetOauth2ServerMetadataOk

`func (o *CreateHTTPConnectorRequest) GetOauth2ServerMetadataOk() (*ExtendedOAuthServerMetadata, bool)`

GetOauth2ServerMetadataOk returns a tuple with the Oauth2ServerMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2ServerMetadata

`func (o *CreateHTTPConnectorRequest) SetOauth2ServerMetadata(v ExtendedOAuthServerMetadata)`

SetOauth2ServerMetadata sets Oauth2ServerMetadata field to given value.

### HasOauth2ServerMetadata

`func (o *CreateHTTPConnectorRequest) HasOauth2ServerMetadata() bool`

HasOauth2ServerMetadata returns a boolean if a field has been set.

### SetOauth2ServerMetadataNil

`func (o *CreateHTTPConnectorRequest) SetOauth2ServerMetadataNil(b bool)`

 SetOauth2ServerMetadataNil sets the value for Oauth2ServerMetadata to be an explicit nil

### UnsetOauth2ServerMetadata
`func (o *CreateHTTPConnectorRequest) UnsetOauth2ServerMetadata()`

UnsetOauth2ServerMetadata ensures that no value is present for Oauth2ServerMetadata, not even an explicit nil
### GetOauth2ServerMetadataUrl

`func (o *CreateHTTPConnectorRequest) GetOauth2ServerMetadataUrl() string`

GetOauth2ServerMetadataUrl returns the Oauth2ServerMetadataUrl field if non-nil, zero value otherwise.

### GetOauth2ServerMetadataUrlOk

`func (o *CreateHTTPConnectorRequest) GetOauth2ServerMetadataUrlOk() (*string, bool)`

GetOauth2ServerMetadataUrlOk returns a tuple with the Oauth2ServerMetadataUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2ServerMetadataUrl

`func (o *CreateHTTPConnectorRequest) SetOauth2ServerMetadataUrl(v string)`

SetOauth2ServerMetadataUrl sets Oauth2ServerMetadataUrl field to given value.

### HasOauth2ServerMetadataUrl

`func (o *CreateHTTPConnectorRequest) HasOauth2ServerMetadataUrl() bool`

HasOauth2ServerMetadataUrl returns a boolean if a field has been set.

### SetOauth2ServerMetadataUrlNil

`func (o *CreateHTTPConnectorRequest) SetOauth2ServerMetadataUrlNil(b bool)`

 SetOauth2ServerMetadataUrlNil sets the value for Oauth2ServerMetadataUrl to be an explicit nil

### UnsetOauth2ServerMetadataUrl
`func (o *CreateHTTPConnectorRequest) UnsetOauth2ServerMetadataUrl()`

UnsetOauth2ServerMetadataUrl ensures that no value is present for Oauth2ServerMetadataUrl, not even an explicit nil
### GetName

`func (o *CreateHTTPConnectorRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateHTTPConnectorRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateHTTPConnectorRequest) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *CreateHTTPConnectorRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateHTTPConnectorRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateHTTPConnectorRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CreateHTTPConnectorRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *CreateHTTPConnectorRequest) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *CreateHTTPConnectorRequest) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *CreateHTTPConnectorRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateHTTPConnectorRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateHTTPConnectorRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetIconUrl

`func (o *CreateHTTPConnectorRequest) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *CreateHTTPConnectorRequest) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *CreateHTTPConnectorRequest) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *CreateHTTPConnectorRequest) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### SetIconUrlNil

`func (o *CreateHTTPConnectorRequest) SetIconUrlNil(b bool)`

 SetIconUrlNil sets the value for IconUrl to be an explicit nil

### UnsetIconUrl
`func (o *CreateHTTPConnectorRequest) UnsetIconUrl()`

UnsetIconUrl ensures that no value is present for IconUrl, not even an explicit nil
### GetServer

`func (o *CreateHTTPConnectorRequest) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *CreateHTTPConnectorRequest) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *CreateHTTPConnectorRequest) SetServer(v string)`

SetServer sets Server field to given value.


### GetVisibility

`func (o *CreateHTTPConnectorRequest) GetVisibility() PublicResourceVisibility`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *CreateHTTPConnectorRequest) GetVisibilityOk() (*PublicResourceVisibility, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *CreateHTTPConnectorRequest) SetVisibility(v PublicResourceVisibility)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *CreateHTTPConnectorRequest) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetAuthMethods

`func (o *CreateHTTPConnectorRequest) GetAuthMethods() []ConnectorMCPPublicUpdateAuthMethodsInnerAnyOf`

GetAuthMethods returns the AuthMethods field if non-nil, zero value otherwise.

### GetAuthMethodsOk

`func (o *CreateHTTPConnectorRequest) GetAuthMethodsOk() (*[]ConnectorMCPPublicUpdateAuthMethodsInnerAnyOf, bool)`

GetAuthMethodsOk returns a tuple with the AuthMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMethods

`func (o *CreateHTTPConnectorRequest) SetAuthMethods(v []ConnectorMCPPublicUpdateAuthMethodsInnerAnyOf)`

SetAuthMethods sets AuthMethods field to given value.

### HasAuthMethods

`func (o *CreateHTTPConnectorRequest) HasAuthMethods() bool`

HasAuthMethods returns a boolean if a field has been set.

### SetAuthMethodsNil

`func (o *CreateHTTPConnectorRequest) SetAuthMethodsNil(b bool)`

 SetAuthMethodsNil sets the value for AuthMethods to be an explicit nil

### UnsetAuthMethods
`func (o *CreateHTTPConnectorRequest) UnsetAuthMethods()`

UnsetAuthMethods ensures that no value is present for AuthMethods, not even an explicit nil
### GetSystemPrompt

`func (o *CreateHTTPConnectorRequest) GetSystemPrompt() string`

GetSystemPrompt returns the SystemPrompt field if non-nil, zero value otherwise.

### GetSystemPromptOk

`func (o *CreateHTTPConnectorRequest) GetSystemPromptOk() (*string, bool)`

GetSystemPromptOk returns a tuple with the SystemPrompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemPrompt

`func (o *CreateHTTPConnectorRequest) SetSystemPrompt(v string)`

SetSystemPrompt sets SystemPrompt field to given value.

### HasSystemPrompt

`func (o *CreateHTTPConnectorRequest) HasSystemPrompt() bool`

HasSystemPrompt returns a boolean if a field has been set.

### SetSystemPromptNil

`func (o *CreateHTTPConnectorRequest) SetSystemPromptNil(b bool)`

 SetSystemPromptNil sets the value for SystemPrompt to be an explicit nil

### UnsetSystemPrompt
`func (o *CreateHTTPConnectorRequest) UnsetSystemPrompt()`

UnsetSystemPrompt ensures that no value is present for SystemPrompt, not even an explicit nil
### GetProtocol

`func (o *CreateHTTPConnectorRequest) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *CreateHTTPConnectorRequest) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *CreateHTTPConnectorRequest) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


