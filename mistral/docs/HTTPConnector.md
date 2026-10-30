# HTTPConnector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Title** | Pointer to **NullableString** |  | [optional] 
**Description** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**ModifiedAt** | **time.Time** |  | 
**Server** | Pointer to **NullableString** |  | [optional] 
**IconUrl** | Pointer to **NullableString** |  | [optional] 
**ServerCard** | Pointer to [**NullableMCPServerCard**](MCPServerCard.md) |  | [optional] 
**OwnerId** | Pointer to **NullableString** |  | [optional] 
**OwnerType** | [**ConsumerType**](ConsumerType.md) |  | 
**Visibility** | [**ResourceVisibility**](ResourceVisibility.md) |  | 
**CreatorId** | Pointer to **NullableString** |  | [optional] 
**Locale** | Pointer to [**NullableConnectorLocale**](ConnectorLocale.md) |  | [optional] 
**SystemPrompt** | Pointer to **NullableString** |  | [optional] 
**SupportedAuthMethods** | Pointer to [**[]PublicAuthenticationMethod**](PublicAuthenticationMethod.md) |  | [optional] 
**ConnectionPreferences** | Pointer to [**[]ConnectionPreference**](ConnectionPreference.md) |  | [optional] 
**ConnectionCredentials** | Pointer to [**[]AuthenticationConfiguration**](AuthenticationConfiguration.md) |  | [optional] 
**BlockedBy** | Pointer to [**[]ConsumerType**](ConsumerType.md) |  | [optional] [default to {}]
**PrivateToolExecution** | **bool** |  | 
**SupportsMcp** | **bool** | Whether this connector exposes an MCP-compatible tool interface. | 
**Mistral** | Pointer to **bool** |  | [optional] [default to false]
**IsAuthenticated** | Pointer to **NullableBool** |  | [optional] 
**Tools** | Pointer to [**[]ConnectorTool**](ConnectorTool.md) |  | [optional] 
**ConnectionConfig** | Pointer to [**NullablePublicConnectionConfig**](PublicConnectionConfig.md) |  | [optional] 
**ExecutionEnv** | Pointer to [**NullablePublicExecutionEnv**](PublicExecutionEnv.md) |  | [optional] 
**Protocol** | **string** |  | 

## Methods

### NewHTTPConnector

`func NewHTTPConnector(id string, name string, description string, createdAt time.Time, modifiedAt time.Time, ownerType ConsumerType, visibility ResourceVisibility, privateToolExecution bool, supportsMcp bool, protocol string, ) *HTTPConnector`

NewHTTPConnector instantiates a new HTTPConnector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHTTPConnectorWithDefaults

`func NewHTTPConnectorWithDefaults() *HTTPConnector`

NewHTTPConnectorWithDefaults instantiates a new HTTPConnector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HTTPConnector) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HTTPConnector) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HTTPConnector) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *HTTPConnector) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HTTPConnector) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HTTPConnector) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *HTTPConnector) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *HTTPConnector) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *HTTPConnector) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *HTTPConnector) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *HTTPConnector) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *HTTPConnector) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *HTTPConnector) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HTTPConnector) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HTTPConnector) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCreatedAt

`func (o *HTTPConnector) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *HTTPConnector) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *HTTPConnector) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetModifiedAt

`func (o *HTTPConnector) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *HTTPConnector) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *HTTPConnector) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetServer

`func (o *HTTPConnector) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *HTTPConnector) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *HTTPConnector) SetServer(v string)`

SetServer sets Server field to given value.

### HasServer

`func (o *HTTPConnector) HasServer() bool`

HasServer returns a boolean if a field has been set.

### SetServerNil

`func (o *HTTPConnector) SetServerNil(b bool)`

 SetServerNil sets the value for Server to be an explicit nil

### UnsetServer
`func (o *HTTPConnector) UnsetServer()`

UnsetServer ensures that no value is present for Server, not even an explicit nil
### GetIconUrl

`func (o *HTTPConnector) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *HTTPConnector) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *HTTPConnector) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *HTTPConnector) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### SetIconUrlNil

`func (o *HTTPConnector) SetIconUrlNil(b bool)`

 SetIconUrlNil sets the value for IconUrl to be an explicit nil

### UnsetIconUrl
`func (o *HTTPConnector) UnsetIconUrl()`

UnsetIconUrl ensures that no value is present for IconUrl, not even an explicit nil
### GetServerCard

`func (o *HTTPConnector) GetServerCard() MCPServerCard`

GetServerCard returns the ServerCard field if non-nil, zero value otherwise.

### GetServerCardOk

`func (o *HTTPConnector) GetServerCardOk() (*MCPServerCard, bool)`

GetServerCardOk returns a tuple with the ServerCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCard

`func (o *HTTPConnector) SetServerCard(v MCPServerCard)`

SetServerCard sets ServerCard field to given value.

### HasServerCard

`func (o *HTTPConnector) HasServerCard() bool`

HasServerCard returns a boolean if a field has been set.

### SetServerCardNil

`func (o *HTTPConnector) SetServerCardNil(b bool)`

 SetServerCardNil sets the value for ServerCard to be an explicit nil

### UnsetServerCard
`func (o *HTTPConnector) UnsetServerCard()`

UnsetServerCard ensures that no value is present for ServerCard, not even an explicit nil
### GetOwnerId

`func (o *HTTPConnector) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *HTTPConnector) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *HTTPConnector) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *HTTPConnector) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### SetOwnerIdNil

`func (o *HTTPConnector) SetOwnerIdNil(b bool)`

 SetOwnerIdNil sets the value for OwnerId to be an explicit nil

### UnsetOwnerId
`func (o *HTTPConnector) UnsetOwnerId()`

UnsetOwnerId ensures that no value is present for OwnerId, not even an explicit nil
### GetOwnerType

`func (o *HTTPConnector) GetOwnerType() ConsumerType`

GetOwnerType returns the OwnerType field if non-nil, zero value otherwise.

### GetOwnerTypeOk

`func (o *HTTPConnector) GetOwnerTypeOk() (*ConsumerType, bool)`

GetOwnerTypeOk returns a tuple with the OwnerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerType

`func (o *HTTPConnector) SetOwnerType(v ConsumerType)`

SetOwnerType sets OwnerType field to given value.


### GetVisibility

`func (o *HTTPConnector) GetVisibility() ResourceVisibility`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *HTTPConnector) GetVisibilityOk() (*ResourceVisibility, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *HTTPConnector) SetVisibility(v ResourceVisibility)`

SetVisibility sets Visibility field to given value.


### GetCreatorId

`func (o *HTTPConnector) GetCreatorId() string`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *HTTPConnector) GetCreatorIdOk() (*string, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *HTTPConnector) SetCreatorId(v string)`

SetCreatorId sets CreatorId field to given value.

### HasCreatorId

`func (o *HTTPConnector) HasCreatorId() bool`

HasCreatorId returns a boolean if a field has been set.

### SetCreatorIdNil

`func (o *HTTPConnector) SetCreatorIdNil(b bool)`

 SetCreatorIdNil sets the value for CreatorId to be an explicit nil

### UnsetCreatorId
`func (o *HTTPConnector) UnsetCreatorId()`

UnsetCreatorId ensures that no value is present for CreatorId, not even an explicit nil
### GetLocale

`func (o *HTTPConnector) GetLocale() ConnectorLocale`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *HTTPConnector) GetLocaleOk() (*ConnectorLocale, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *HTTPConnector) SetLocale(v ConnectorLocale)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *HTTPConnector) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### SetLocaleNil

`func (o *HTTPConnector) SetLocaleNil(b bool)`

 SetLocaleNil sets the value for Locale to be an explicit nil

### UnsetLocale
`func (o *HTTPConnector) UnsetLocale()`

UnsetLocale ensures that no value is present for Locale, not even an explicit nil
### GetSystemPrompt

`func (o *HTTPConnector) GetSystemPrompt() string`

GetSystemPrompt returns the SystemPrompt field if non-nil, zero value otherwise.

### GetSystemPromptOk

`func (o *HTTPConnector) GetSystemPromptOk() (*string, bool)`

GetSystemPromptOk returns a tuple with the SystemPrompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemPrompt

`func (o *HTTPConnector) SetSystemPrompt(v string)`

SetSystemPrompt sets SystemPrompt field to given value.

### HasSystemPrompt

`func (o *HTTPConnector) HasSystemPrompt() bool`

HasSystemPrompt returns a boolean if a field has been set.

### SetSystemPromptNil

`func (o *HTTPConnector) SetSystemPromptNil(b bool)`

 SetSystemPromptNil sets the value for SystemPrompt to be an explicit nil

### UnsetSystemPrompt
`func (o *HTTPConnector) UnsetSystemPrompt()`

UnsetSystemPrompt ensures that no value is present for SystemPrompt, not even an explicit nil
### GetSupportedAuthMethods

`func (o *HTTPConnector) GetSupportedAuthMethods() []PublicAuthenticationMethod`

GetSupportedAuthMethods returns the SupportedAuthMethods field if non-nil, zero value otherwise.

### GetSupportedAuthMethodsOk

`func (o *HTTPConnector) GetSupportedAuthMethodsOk() (*[]PublicAuthenticationMethod, bool)`

GetSupportedAuthMethodsOk returns a tuple with the SupportedAuthMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedAuthMethods

`func (o *HTTPConnector) SetSupportedAuthMethods(v []PublicAuthenticationMethod)`

SetSupportedAuthMethods sets SupportedAuthMethods field to given value.

### HasSupportedAuthMethods

`func (o *HTTPConnector) HasSupportedAuthMethods() bool`

HasSupportedAuthMethods returns a boolean if a field has been set.

### SetSupportedAuthMethodsNil

`func (o *HTTPConnector) SetSupportedAuthMethodsNil(b bool)`

 SetSupportedAuthMethodsNil sets the value for SupportedAuthMethods to be an explicit nil

### UnsetSupportedAuthMethods
`func (o *HTTPConnector) UnsetSupportedAuthMethods()`

UnsetSupportedAuthMethods ensures that no value is present for SupportedAuthMethods, not even an explicit nil
### GetConnectionPreferences

`func (o *HTTPConnector) GetConnectionPreferences() []ConnectionPreference`

GetConnectionPreferences returns the ConnectionPreferences field if non-nil, zero value otherwise.

### GetConnectionPreferencesOk

`func (o *HTTPConnector) GetConnectionPreferencesOk() (*[]ConnectionPreference, bool)`

GetConnectionPreferencesOk returns a tuple with the ConnectionPreferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionPreferences

`func (o *HTTPConnector) SetConnectionPreferences(v []ConnectionPreference)`

SetConnectionPreferences sets ConnectionPreferences field to given value.

### HasConnectionPreferences

`func (o *HTTPConnector) HasConnectionPreferences() bool`

HasConnectionPreferences returns a boolean if a field has been set.

### SetConnectionPreferencesNil

`func (o *HTTPConnector) SetConnectionPreferencesNil(b bool)`

 SetConnectionPreferencesNil sets the value for ConnectionPreferences to be an explicit nil

### UnsetConnectionPreferences
`func (o *HTTPConnector) UnsetConnectionPreferences()`

UnsetConnectionPreferences ensures that no value is present for ConnectionPreferences, not even an explicit nil
### GetConnectionCredentials

`func (o *HTTPConnector) GetConnectionCredentials() []AuthenticationConfiguration`

GetConnectionCredentials returns the ConnectionCredentials field if non-nil, zero value otherwise.

### GetConnectionCredentialsOk

`func (o *HTTPConnector) GetConnectionCredentialsOk() (*[]AuthenticationConfiguration, bool)`

GetConnectionCredentialsOk returns a tuple with the ConnectionCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionCredentials

`func (o *HTTPConnector) SetConnectionCredentials(v []AuthenticationConfiguration)`

SetConnectionCredentials sets ConnectionCredentials field to given value.

### HasConnectionCredentials

`func (o *HTTPConnector) HasConnectionCredentials() bool`

HasConnectionCredentials returns a boolean if a field has been set.

### SetConnectionCredentialsNil

`func (o *HTTPConnector) SetConnectionCredentialsNil(b bool)`

 SetConnectionCredentialsNil sets the value for ConnectionCredentials to be an explicit nil

### UnsetConnectionCredentials
`func (o *HTTPConnector) UnsetConnectionCredentials()`

UnsetConnectionCredentials ensures that no value is present for ConnectionCredentials, not even an explicit nil
### GetBlockedBy

`func (o *HTTPConnector) GetBlockedBy() []ConsumerType`

GetBlockedBy returns the BlockedBy field if non-nil, zero value otherwise.

### GetBlockedByOk

`func (o *HTTPConnector) GetBlockedByOk() (*[]ConsumerType, bool)`

GetBlockedByOk returns a tuple with the BlockedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedBy

`func (o *HTTPConnector) SetBlockedBy(v []ConsumerType)`

SetBlockedBy sets BlockedBy field to given value.

### HasBlockedBy

`func (o *HTTPConnector) HasBlockedBy() bool`

HasBlockedBy returns a boolean if a field has been set.

### GetPrivateToolExecution

`func (o *HTTPConnector) GetPrivateToolExecution() bool`

GetPrivateToolExecution returns the PrivateToolExecution field if non-nil, zero value otherwise.

### GetPrivateToolExecutionOk

`func (o *HTTPConnector) GetPrivateToolExecutionOk() (*bool, bool)`

GetPrivateToolExecutionOk returns a tuple with the PrivateToolExecution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateToolExecution

`func (o *HTTPConnector) SetPrivateToolExecution(v bool)`

SetPrivateToolExecution sets PrivateToolExecution field to given value.


### GetSupportsMcp

`func (o *HTTPConnector) GetSupportsMcp() bool`

GetSupportsMcp returns the SupportsMcp field if non-nil, zero value otherwise.

### GetSupportsMcpOk

`func (o *HTTPConnector) GetSupportsMcpOk() (*bool, bool)`

GetSupportsMcpOk returns a tuple with the SupportsMcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsMcp

`func (o *HTTPConnector) SetSupportsMcp(v bool)`

SetSupportsMcp sets SupportsMcp field to given value.


### GetMistral

`func (o *HTTPConnector) GetMistral() bool`

GetMistral returns the Mistral field if non-nil, zero value otherwise.

### GetMistralOk

`func (o *HTTPConnector) GetMistralOk() (*bool, bool)`

GetMistralOk returns a tuple with the Mistral field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMistral

`func (o *HTTPConnector) SetMistral(v bool)`

SetMistral sets Mistral field to given value.

### HasMistral

`func (o *HTTPConnector) HasMistral() bool`

HasMistral returns a boolean if a field has been set.

### GetIsAuthenticated

`func (o *HTTPConnector) GetIsAuthenticated() bool`

GetIsAuthenticated returns the IsAuthenticated field if non-nil, zero value otherwise.

### GetIsAuthenticatedOk

`func (o *HTTPConnector) GetIsAuthenticatedOk() (*bool, bool)`

GetIsAuthenticatedOk returns a tuple with the IsAuthenticated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAuthenticated

`func (o *HTTPConnector) SetIsAuthenticated(v bool)`

SetIsAuthenticated sets IsAuthenticated field to given value.

### HasIsAuthenticated

`func (o *HTTPConnector) HasIsAuthenticated() bool`

HasIsAuthenticated returns a boolean if a field has been set.

### SetIsAuthenticatedNil

`func (o *HTTPConnector) SetIsAuthenticatedNil(b bool)`

 SetIsAuthenticatedNil sets the value for IsAuthenticated to be an explicit nil

### UnsetIsAuthenticated
`func (o *HTTPConnector) UnsetIsAuthenticated()`

UnsetIsAuthenticated ensures that no value is present for IsAuthenticated, not even an explicit nil
### GetTools

`func (o *HTTPConnector) GetTools() []ConnectorTool`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *HTTPConnector) GetToolsOk() (*[]ConnectorTool, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *HTTPConnector) SetTools(v []ConnectorTool)`

SetTools sets Tools field to given value.

### HasTools

`func (o *HTTPConnector) HasTools() bool`

HasTools returns a boolean if a field has been set.

### SetToolsNil

`func (o *HTTPConnector) SetToolsNil(b bool)`

 SetToolsNil sets the value for Tools to be an explicit nil

### UnsetTools
`func (o *HTTPConnector) UnsetTools()`

UnsetTools ensures that no value is present for Tools, not even an explicit nil
### GetConnectionConfig

`func (o *HTTPConnector) GetConnectionConfig() PublicConnectionConfig`

GetConnectionConfig returns the ConnectionConfig field if non-nil, zero value otherwise.

### GetConnectionConfigOk

`func (o *HTTPConnector) GetConnectionConfigOk() (*PublicConnectionConfig, bool)`

GetConnectionConfigOk returns a tuple with the ConnectionConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionConfig

`func (o *HTTPConnector) SetConnectionConfig(v PublicConnectionConfig)`

SetConnectionConfig sets ConnectionConfig field to given value.

### HasConnectionConfig

`func (o *HTTPConnector) HasConnectionConfig() bool`

HasConnectionConfig returns a boolean if a field has been set.

### SetConnectionConfigNil

`func (o *HTTPConnector) SetConnectionConfigNil(b bool)`

 SetConnectionConfigNil sets the value for ConnectionConfig to be an explicit nil

### UnsetConnectionConfig
`func (o *HTTPConnector) UnsetConnectionConfig()`

UnsetConnectionConfig ensures that no value is present for ConnectionConfig, not even an explicit nil
### GetExecutionEnv

`func (o *HTTPConnector) GetExecutionEnv() PublicExecutionEnv`

GetExecutionEnv returns the ExecutionEnv field if non-nil, zero value otherwise.

### GetExecutionEnvOk

`func (o *HTTPConnector) GetExecutionEnvOk() (*PublicExecutionEnv, bool)`

GetExecutionEnvOk returns a tuple with the ExecutionEnv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionEnv

`func (o *HTTPConnector) SetExecutionEnv(v PublicExecutionEnv)`

SetExecutionEnv sets ExecutionEnv field to given value.

### HasExecutionEnv

`func (o *HTTPConnector) HasExecutionEnv() bool`

HasExecutionEnv returns a boolean if a field has been set.

### SetExecutionEnvNil

`func (o *HTTPConnector) SetExecutionEnvNil(b bool)`

 SetExecutionEnvNil sets the value for ExecutionEnv to be an explicit nil

### UnsetExecutionEnv
`func (o *HTTPConnector) UnsetExecutionEnv()`

UnsetExecutionEnv ensures that no value is present for ExecutionEnv, not even an explicit nil
### GetProtocol

`func (o *HTTPConnector) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *HTTPConnector) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *HTTPConnector) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


