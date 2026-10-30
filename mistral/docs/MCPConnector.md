# MCPConnector

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

### NewMCPConnector

`func NewMCPConnector(id string, name string, description string, createdAt time.Time, modifiedAt time.Time, ownerType ConsumerType, visibility ResourceVisibility, privateToolExecution bool, supportsMcp bool, protocol string, ) *MCPConnector`

NewMCPConnector instantiates a new MCPConnector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMCPConnectorWithDefaults

`func NewMCPConnectorWithDefaults() *MCPConnector`

NewMCPConnectorWithDefaults instantiates a new MCPConnector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MCPConnector) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MCPConnector) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MCPConnector) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *MCPConnector) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MCPConnector) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MCPConnector) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *MCPConnector) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MCPConnector) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *MCPConnector) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *MCPConnector) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *MCPConnector) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *MCPConnector) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *MCPConnector) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MCPConnector) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MCPConnector) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCreatedAt

`func (o *MCPConnector) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MCPConnector) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MCPConnector) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetModifiedAt

`func (o *MCPConnector) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *MCPConnector) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *MCPConnector) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetServer

`func (o *MCPConnector) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *MCPConnector) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *MCPConnector) SetServer(v string)`

SetServer sets Server field to given value.

### HasServer

`func (o *MCPConnector) HasServer() bool`

HasServer returns a boolean if a field has been set.

### SetServerNil

`func (o *MCPConnector) SetServerNil(b bool)`

 SetServerNil sets the value for Server to be an explicit nil

### UnsetServer
`func (o *MCPConnector) UnsetServer()`

UnsetServer ensures that no value is present for Server, not even an explicit nil
### GetIconUrl

`func (o *MCPConnector) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *MCPConnector) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *MCPConnector) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *MCPConnector) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### SetIconUrlNil

`func (o *MCPConnector) SetIconUrlNil(b bool)`

 SetIconUrlNil sets the value for IconUrl to be an explicit nil

### UnsetIconUrl
`func (o *MCPConnector) UnsetIconUrl()`

UnsetIconUrl ensures that no value is present for IconUrl, not even an explicit nil
### GetServerCard

`func (o *MCPConnector) GetServerCard() MCPServerCard`

GetServerCard returns the ServerCard field if non-nil, zero value otherwise.

### GetServerCardOk

`func (o *MCPConnector) GetServerCardOk() (*MCPServerCard, bool)`

GetServerCardOk returns a tuple with the ServerCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCard

`func (o *MCPConnector) SetServerCard(v MCPServerCard)`

SetServerCard sets ServerCard field to given value.

### HasServerCard

`func (o *MCPConnector) HasServerCard() bool`

HasServerCard returns a boolean if a field has been set.

### SetServerCardNil

`func (o *MCPConnector) SetServerCardNil(b bool)`

 SetServerCardNil sets the value for ServerCard to be an explicit nil

### UnsetServerCard
`func (o *MCPConnector) UnsetServerCard()`

UnsetServerCard ensures that no value is present for ServerCard, not even an explicit nil
### GetOwnerId

`func (o *MCPConnector) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *MCPConnector) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *MCPConnector) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *MCPConnector) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### SetOwnerIdNil

`func (o *MCPConnector) SetOwnerIdNil(b bool)`

 SetOwnerIdNil sets the value for OwnerId to be an explicit nil

### UnsetOwnerId
`func (o *MCPConnector) UnsetOwnerId()`

UnsetOwnerId ensures that no value is present for OwnerId, not even an explicit nil
### GetOwnerType

`func (o *MCPConnector) GetOwnerType() ConsumerType`

GetOwnerType returns the OwnerType field if non-nil, zero value otherwise.

### GetOwnerTypeOk

`func (o *MCPConnector) GetOwnerTypeOk() (*ConsumerType, bool)`

GetOwnerTypeOk returns a tuple with the OwnerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerType

`func (o *MCPConnector) SetOwnerType(v ConsumerType)`

SetOwnerType sets OwnerType field to given value.


### GetVisibility

`func (o *MCPConnector) GetVisibility() ResourceVisibility`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *MCPConnector) GetVisibilityOk() (*ResourceVisibility, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *MCPConnector) SetVisibility(v ResourceVisibility)`

SetVisibility sets Visibility field to given value.


### GetCreatorId

`func (o *MCPConnector) GetCreatorId() string`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *MCPConnector) GetCreatorIdOk() (*string, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *MCPConnector) SetCreatorId(v string)`

SetCreatorId sets CreatorId field to given value.

### HasCreatorId

`func (o *MCPConnector) HasCreatorId() bool`

HasCreatorId returns a boolean if a field has been set.

### SetCreatorIdNil

`func (o *MCPConnector) SetCreatorIdNil(b bool)`

 SetCreatorIdNil sets the value for CreatorId to be an explicit nil

### UnsetCreatorId
`func (o *MCPConnector) UnsetCreatorId()`

UnsetCreatorId ensures that no value is present for CreatorId, not even an explicit nil
### GetLocale

`func (o *MCPConnector) GetLocale() ConnectorLocale`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *MCPConnector) GetLocaleOk() (*ConnectorLocale, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *MCPConnector) SetLocale(v ConnectorLocale)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *MCPConnector) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### SetLocaleNil

`func (o *MCPConnector) SetLocaleNil(b bool)`

 SetLocaleNil sets the value for Locale to be an explicit nil

### UnsetLocale
`func (o *MCPConnector) UnsetLocale()`

UnsetLocale ensures that no value is present for Locale, not even an explicit nil
### GetSystemPrompt

`func (o *MCPConnector) GetSystemPrompt() string`

GetSystemPrompt returns the SystemPrompt field if non-nil, zero value otherwise.

### GetSystemPromptOk

`func (o *MCPConnector) GetSystemPromptOk() (*string, bool)`

GetSystemPromptOk returns a tuple with the SystemPrompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemPrompt

`func (o *MCPConnector) SetSystemPrompt(v string)`

SetSystemPrompt sets SystemPrompt field to given value.

### HasSystemPrompt

`func (o *MCPConnector) HasSystemPrompt() bool`

HasSystemPrompt returns a boolean if a field has been set.

### SetSystemPromptNil

`func (o *MCPConnector) SetSystemPromptNil(b bool)`

 SetSystemPromptNil sets the value for SystemPrompt to be an explicit nil

### UnsetSystemPrompt
`func (o *MCPConnector) UnsetSystemPrompt()`

UnsetSystemPrompt ensures that no value is present for SystemPrompt, not even an explicit nil
### GetSupportedAuthMethods

`func (o *MCPConnector) GetSupportedAuthMethods() []PublicAuthenticationMethod`

GetSupportedAuthMethods returns the SupportedAuthMethods field if non-nil, zero value otherwise.

### GetSupportedAuthMethodsOk

`func (o *MCPConnector) GetSupportedAuthMethodsOk() (*[]PublicAuthenticationMethod, bool)`

GetSupportedAuthMethodsOk returns a tuple with the SupportedAuthMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedAuthMethods

`func (o *MCPConnector) SetSupportedAuthMethods(v []PublicAuthenticationMethod)`

SetSupportedAuthMethods sets SupportedAuthMethods field to given value.

### HasSupportedAuthMethods

`func (o *MCPConnector) HasSupportedAuthMethods() bool`

HasSupportedAuthMethods returns a boolean if a field has been set.

### SetSupportedAuthMethodsNil

`func (o *MCPConnector) SetSupportedAuthMethodsNil(b bool)`

 SetSupportedAuthMethodsNil sets the value for SupportedAuthMethods to be an explicit nil

### UnsetSupportedAuthMethods
`func (o *MCPConnector) UnsetSupportedAuthMethods()`

UnsetSupportedAuthMethods ensures that no value is present for SupportedAuthMethods, not even an explicit nil
### GetConnectionPreferences

`func (o *MCPConnector) GetConnectionPreferences() []ConnectionPreference`

GetConnectionPreferences returns the ConnectionPreferences field if non-nil, zero value otherwise.

### GetConnectionPreferencesOk

`func (o *MCPConnector) GetConnectionPreferencesOk() (*[]ConnectionPreference, bool)`

GetConnectionPreferencesOk returns a tuple with the ConnectionPreferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionPreferences

`func (o *MCPConnector) SetConnectionPreferences(v []ConnectionPreference)`

SetConnectionPreferences sets ConnectionPreferences field to given value.

### HasConnectionPreferences

`func (o *MCPConnector) HasConnectionPreferences() bool`

HasConnectionPreferences returns a boolean if a field has been set.

### SetConnectionPreferencesNil

`func (o *MCPConnector) SetConnectionPreferencesNil(b bool)`

 SetConnectionPreferencesNil sets the value for ConnectionPreferences to be an explicit nil

### UnsetConnectionPreferences
`func (o *MCPConnector) UnsetConnectionPreferences()`

UnsetConnectionPreferences ensures that no value is present for ConnectionPreferences, not even an explicit nil
### GetConnectionCredentials

`func (o *MCPConnector) GetConnectionCredentials() []AuthenticationConfiguration`

GetConnectionCredentials returns the ConnectionCredentials field if non-nil, zero value otherwise.

### GetConnectionCredentialsOk

`func (o *MCPConnector) GetConnectionCredentialsOk() (*[]AuthenticationConfiguration, bool)`

GetConnectionCredentialsOk returns a tuple with the ConnectionCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionCredentials

`func (o *MCPConnector) SetConnectionCredentials(v []AuthenticationConfiguration)`

SetConnectionCredentials sets ConnectionCredentials field to given value.

### HasConnectionCredentials

`func (o *MCPConnector) HasConnectionCredentials() bool`

HasConnectionCredentials returns a boolean if a field has been set.

### SetConnectionCredentialsNil

`func (o *MCPConnector) SetConnectionCredentialsNil(b bool)`

 SetConnectionCredentialsNil sets the value for ConnectionCredentials to be an explicit nil

### UnsetConnectionCredentials
`func (o *MCPConnector) UnsetConnectionCredentials()`

UnsetConnectionCredentials ensures that no value is present for ConnectionCredentials, not even an explicit nil
### GetBlockedBy

`func (o *MCPConnector) GetBlockedBy() []ConsumerType`

GetBlockedBy returns the BlockedBy field if non-nil, zero value otherwise.

### GetBlockedByOk

`func (o *MCPConnector) GetBlockedByOk() (*[]ConsumerType, bool)`

GetBlockedByOk returns a tuple with the BlockedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedBy

`func (o *MCPConnector) SetBlockedBy(v []ConsumerType)`

SetBlockedBy sets BlockedBy field to given value.

### HasBlockedBy

`func (o *MCPConnector) HasBlockedBy() bool`

HasBlockedBy returns a boolean if a field has been set.

### GetPrivateToolExecution

`func (o *MCPConnector) GetPrivateToolExecution() bool`

GetPrivateToolExecution returns the PrivateToolExecution field if non-nil, zero value otherwise.

### GetPrivateToolExecutionOk

`func (o *MCPConnector) GetPrivateToolExecutionOk() (*bool, bool)`

GetPrivateToolExecutionOk returns a tuple with the PrivateToolExecution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateToolExecution

`func (o *MCPConnector) SetPrivateToolExecution(v bool)`

SetPrivateToolExecution sets PrivateToolExecution field to given value.


### GetSupportsMcp

`func (o *MCPConnector) GetSupportsMcp() bool`

GetSupportsMcp returns the SupportsMcp field if non-nil, zero value otherwise.

### GetSupportsMcpOk

`func (o *MCPConnector) GetSupportsMcpOk() (*bool, bool)`

GetSupportsMcpOk returns a tuple with the SupportsMcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsMcp

`func (o *MCPConnector) SetSupportsMcp(v bool)`

SetSupportsMcp sets SupportsMcp field to given value.


### GetMistral

`func (o *MCPConnector) GetMistral() bool`

GetMistral returns the Mistral field if non-nil, zero value otherwise.

### GetMistralOk

`func (o *MCPConnector) GetMistralOk() (*bool, bool)`

GetMistralOk returns a tuple with the Mistral field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMistral

`func (o *MCPConnector) SetMistral(v bool)`

SetMistral sets Mistral field to given value.

### HasMistral

`func (o *MCPConnector) HasMistral() bool`

HasMistral returns a boolean if a field has been set.

### GetIsAuthenticated

`func (o *MCPConnector) GetIsAuthenticated() bool`

GetIsAuthenticated returns the IsAuthenticated field if non-nil, zero value otherwise.

### GetIsAuthenticatedOk

`func (o *MCPConnector) GetIsAuthenticatedOk() (*bool, bool)`

GetIsAuthenticatedOk returns a tuple with the IsAuthenticated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAuthenticated

`func (o *MCPConnector) SetIsAuthenticated(v bool)`

SetIsAuthenticated sets IsAuthenticated field to given value.

### HasIsAuthenticated

`func (o *MCPConnector) HasIsAuthenticated() bool`

HasIsAuthenticated returns a boolean if a field has been set.

### SetIsAuthenticatedNil

`func (o *MCPConnector) SetIsAuthenticatedNil(b bool)`

 SetIsAuthenticatedNil sets the value for IsAuthenticated to be an explicit nil

### UnsetIsAuthenticated
`func (o *MCPConnector) UnsetIsAuthenticated()`

UnsetIsAuthenticated ensures that no value is present for IsAuthenticated, not even an explicit nil
### GetTools

`func (o *MCPConnector) GetTools() []ConnectorTool`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *MCPConnector) GetToolsOk() (*[]ConnectorTool, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *MCPConnector) SetTools(v []ConnectorTool)`

SetTools sets Tools field to given value.

### HasTools

`func (o *MCPConnector) HasTools() bool`

HasTools returns a boolean if a field has been set.

### SetToolsNil

`func (o *MCPConnector) SetToolsNil(b bool)`

 SetToolsNil sets the value for Tools to be an explicit nil

### UnsetTools
`func (o *MCPConnector) UnsetTools()`

UnsetTools ensures that no value is present for Tools, not even an explicit nil
### GetConnectionConfig

`func (o *MCPConnector) GetConnectionConfig() PublicConnectionConfig`

GetConnectionConfig returns the ConnectionConfig field if non-nil, zero value otherwise.

### GetConnectionConfigOk

`func (o *MCPConnector) GetConnectionConfigOk() (*PublicConnectionConfig, bool)`

GetConnectionConfigOk returns a tuple with the ConnectionConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionConfig

`func (o *MCPConnector) SetConnectionConfig(v PublicConnectionConfig)`

SetConnectionConfig sets ConnectionConfig field to given value.

### HasConnectionConfig

`func (o *MCPConnector) HasConnectionConfig() bool`

HasConnectionConfig returns a boolean if a field has been set.

### SetConnectionConfigNil

`func (o *MCPConnector) SetConnectionConfigNil(b bool)`

 SetConnectionConfigNil sets the value for ConnectionConfig to be an explicit nil

### UnsetConnectionConfig
`func (o *MCPConnector) UnsetConnectionConfig()`

UnsetConnectionConfig ensures that no value is present for ConnectionConfig, not even an explicit nil
### GetExecutionEnv

`func (o *MCPConnector) GetExecutionEnv() PublicExecutionEnv`

GetExecutionEnv returns the ExecutionEnv field if non-nil, zero value otherwise.

### GetExecutionEnvOk

`func (o *MCPConnector) GetExecutionEnvOk() (*PublicExecutionEnv, bool)`

GetExecutionEnvOk returns a tuple with the ExecutionEnv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionEnv

`func (o *MCPConnector) SetExecutionEnv(v PublicExecutionEnv)`

SetExecutionEnv sets ExecutionEnv field to given value.

### HasExecutionEnv

`func (o *MCPConnector) HasExecutionEnv() bool`

HasExecutionEnv returns a boolean if a field has been set.

### SetExecutionEnvNil

`func (o *MCPConnector) SetExecutionEnvNil(b bool)`

 SetExecutionEnvNil sets the value for ExecutionEnv to be an explicit nil

### UnsetExecutionEnv
`func (o *MCPConnector) UnsetExecutionEnv()`

UnsetExecutionEnv ensures that no value is present for ExecutionEnv, not even an explicit nil
### GetProtocol

`func (o *MCPConnector) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *MCPConnector) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *MCPConnector) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


