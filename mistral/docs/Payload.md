# Payload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Headers** | Pointer to **map[string]interface{}** | Optional scoped credentials to be sent with connector requests. | [optional] 
**GlobalHeaders** | Pointer to [**map[string]GlobalHeaderValue**](GlobalHeaderValue.md) | Optional connector-wide headers, keyed by header name, set at creation and applied to every credential. Secret values are encrypted at rest and never returned in clear. | [optional] [default to {}]
**Oauth2ServerMetadata** | Pointer to [**ExtendedOAuthServerMetadata**](ExtendedOAuthServerMetadata.md) | Optional OAuth2 authorization server metadata (authorization_endpoint, token_endpoint, etc.). When provided, skips .well-known discovery and uses these endpoints directly. | [optional] 
**Oauth2ServerMetadataUrl** | Pointer to **string** | Optional URL to fetch OAuth2 authorization server metadata from (RFC 8414). When provided, the metadata is fetched from this URL and used instead of .well-known discovery. Mutually exclusive with oauth2_server_metadata. | [optional] 
**Name** | **string** | The name of the connector. Should be 64 char length maximum, alphanumeric, only underscores/dashes. | 
**Title** | Pointer to **string** | Optional human-readable title for the connector. | [optional] 
**Description** | **string** | The description of the connector. | 
**IconUrl** | Pointer to **string** | The optional url of the icon you want to associate to the connector. | [optional] 
**Server** | **string** | The URL of the connector server. | 
**Visibility** | Pointer to [**PublicResourceVisibility**](PublicResourceVisibility.md) | Visibility of the connector. Use &#39;shared_workspace&#39; for workspace scoped connectors, or &#39;private&#39; for private connectors. | [optional] [default to PUBLICRESOURCEVISIBILITY_PRIVATE]
**AuthMethods** | Pointer to [**[]ConnectorMCPPublicUpdateAuthMethodsInnerAnyOf**](ConnectorMCPPublicUpdateAuthMethodsInnerAnyOf.md) | Authentication methods supported by the connector. | [optional] 
**SystemPrompt** | Pointer to **string** | Optional system prompt for the connector. | [optional] 
**Protocol** | **string** |  | 

## Methods

### NewPayload

`func NewPayload(name string, description string, server string, protocol string, ) *Payload`

NewPayload instantiates a new Payload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayloadWithDefaults

`func NewPayloadWithDefaults() *Payload`

NewPayloadWithDefaults instantiates a new Payload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeaders

`func (o *Payload) GetHeaders() map[string]interface{}`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *Payload) GetHeadersOk() (*map[string]interface{}, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *Payload) SetHeaders(v map[string]interface{})`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *Payload) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetGlobalHeaders

`func (o *Payload) GetGlobalHeaders() map[string]GlobalHeaderValue`

GetGlobalHeaders returns the GlobalHeaders field if non-nil, zero value otherwise.

### GetGlobalHeadersOk

`func (o *Payload) GetGlobalHeadersOk() (*map[string]GlobalHeaderValue, bool)`

GetGlobalHeadersOk returns a tuple with the GlobalHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalHeaders

`func (o *Payload) SetGlobalHeaders(v map[string]GlobalHeaderValue)`

SetGlobalHeaders sets GlobalHeaders field to given value.

### HasGlobalHeaders

`func (o *Payload) HasGlobalHeaders() bool`

HasGlobalHeaders returns a boolean if a field has been set.

### GetOauth2ServerMetadata

`func (o *Payload) GetOauth2ServerMetadata() ExtendedOAuthServerMetadata`

GetOauth2ServerMetadata returns the Oauth2ServerMetadata field if non-nil, zero value otherwise.

### GetOauth2ServerMetadataOk

`func (o *Payload) GetOauth2ServerMetadataOk() (*ExtendedOAuthServerMetadata, bool)`

GetOauth2ServerMetadataOk returns a tuple with the Oauth2ServerMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2ServerMetadata

`func (o *Payload) SetOauth2ServerMetadata(v ExtendedOAuthServerMetadata)`

SetOauth2ServerMetadata sets Oauth2ServerMetadata field to given value.

### HasOauth2ServerMetadata

`func (o *Payload) HasOauth2ServerMetadata() bool`

HasOauth2ServerMetadata returns a boolean if a field has been set.

### GetOauth2ServerMetadataUrl

`func (o *Payload) GetOauth2ServerMetadataUrl() string`

GetOauth2ServerMetadataUrl returns the Oauth2ServerMetadataUrl field if non-nil, zero value otherwise.

### GetOauth2ServerMetadataUrlOk

`func (o *Payload) GetOauth2ServerMetadataUrlOk() (*string, bool)`

GetOauth2ServerMetadataUrlOk returns a tuple with the Oauth2ServerMetadataUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2ServerMetadataUrl

`func (o *Payload) SetOauth2ServerMetadataUrl(v string)`

SetOauth2ServerMetadataUrl sets Oauth2ServerMetadataUrl field to given value.

### HasOauth2ServerMetadataUrl

`func (o *Payload) HasOauth2ServerMetadataUrl() bool`

HasOauth2ServerMetadataUrl returns a boolean if a field has been set.

### GetName

`func (o *Payload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Payload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Payload) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *Payload) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Payload) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Payload) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Payload) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *Payload) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Payload) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Payload) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetIconUrl

`func (o *Payload) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *Payload) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *Payload) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *Payload) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### GetServer

`func (o *Payload) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *Payload) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *Payload) SetServer(v string)`

SetServer sets Server field to given value.


### GetVisibility

`func (o *Payload) GetVisibility() PublicResourceVisibility`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *Payload) GetVisibilityOk() (*PublicResourceVisibility, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *Payload) SetVisibility(v PublicResourceVisibility)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *Payload) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetAuthMethods

`func (o *Payload) GetAuthMethods() []ConnectorMCPPublicUpdateAuthMethodsInnerAnyOf`

GetAuthMethods returns the AuthMethods field if non-nil, zero value otherwise.

### GetAuthMethodsOk

`func (o *Payload) GetAuthMethodsOk() (*[]ConnectorMCPPublicUpdateAuthMethodsInnerAnyOf, bool)`

GetAuthMethodsOk returns a tuple with the AuthMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMethods

`func (o *Payload) SetAuthMethods(v []ConnectorMCPPublicUpdateAuthMethodsInnerAnyOf)`

SetAuthMethods sets AuthMethods field to given value.

### HasAuthMethods

`func (o *Payload) HasAuthMethods() bool`

HasAuthMethods returns a boolean if a field has been set.

### GetSystemPrompt

`func (o *Payload) GetSystemPrompt() string`

GetSystemPrompt returns the SystemPrompt field if non-nil, zero value otherwise.

### GetSystemPromptOk

`func (o *Payload) GetSystemPromptOk() (*string, bool)`

GetSystemPromptOk returns a tuple with the SystemPrompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemPrompt

`func (o *Payload) SetSystemPrompt(v string)`

SetSystemPrompt sets SystemPrompt field to given value.

### HasSystemPrompt

`func (o *Payload) HasSystemPrompt() bool`

HasSystemPrompt returns a boolean if a field has been set.

### GetProtocol

`func (o *Payload) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *Payload) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *Payload) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


