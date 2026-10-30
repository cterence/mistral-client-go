# Connector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Title** | Pointer to **string** |  | [optional] 
**Description** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**ModifiedAt** | **time.Time** |  | 
**Server** | Pointer to **string** |  | [optional] 
**IconUrl** | Pointer to **string** |  | [optional] 
**ServerCard** | Pointer to [**MCPServerCard**](MCPServerCard.md) |  | [optional] 
**OwnerId** | Pointer to **string** |  | [optional] 
**OwnerType** | [**ConsumerType**](ConsumerType.md) |  | 
**Visibility** | [**ResourceVisibility**](ResourceVisibility.md) |  | 
**CreatorId** | Pointer to **string** |  | [optional] 
**Locale** | Pointer to [**ConnectorLocale**](ConnectorLocale.md) |  | [optional] 
**SystemPrompt** | Pointer to **string** |  | [optional] 
**SupportedAuthMethods** | Pointer to [**[]PublicAuthenticationMethod**](PublicAuthenticationMethod.md) |  | [optional] 
**ConnectionPreferences** | Pointer to [**[]ConnectionPreference**](ConnectionPreference.md) |  | [optional] 
**ConnectionCredentials** | Pointer to [**[]AuthenticationConfiguration**](AuthenticationConfiguration.md) |  | [optional] 
**BlockedBy** | Pointer to [**[]ConsumerType**](ConsumerType.md) |  | [optional] [default to {}]
**PrivateToolExecution** | **bool** |  | 
**SupportsMcp** | **bool** | Whether this connector exposes an MCP-compatible tool interface. | 
**Mistral** | Pointer to **bool** |  | [optional] [default to false]
**IsAuthenticated** | Pointer to **bool** |  | [optional] 
**Tools** | Pointer to [**[]ConnectorTool**](ConnectorTool.md) |  | [optional] 
**ConnectionConfig** | Pointer to [**PublicConnectionConfig**](PublicConnectionConfig.md) |  | [optional] 
**ExecutionEnv** | Pointer to [**PublicExecutionEnv**](PublicExecutionEnv.md) |  | [optional] 
**Protocol** | **string** |  | 

## Methods

### NewConnector

`func NewConnector(id string, name string, description string, createdAt time.Time, modifiedAt time.Time, ownerType ConsumerType, visibility ResourceVisibility, privateToolExecution bool, supportsMcp bool, protocol string, ) *Connector`

NewConnector instantiates a new Connector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorWithDefaults

`func NewConnectorWithDefaults() *Connector`

NewConnectorWithDefaults instantiates a new Connector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Connector) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Connector) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Connector) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Connector) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Connector) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Connector) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *Connector) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Connector) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Connector) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Connector) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *Connector) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Connector) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Connector) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCreatedAt

`func (o *Connector) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Connector) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Connector) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetModifiedAt

`func (o *Connector) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *Connector) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *Connector) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetServer

`func (o *Connector) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *Connector) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *Connector) SetServer(v string)`

SetServer sets Server field to given value.

### HasServer

`func (o *Connector) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetIconUrl

`func (o *Connector) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *Connector) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *Connector) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *Connector) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### GetServerCard

`func (o *Connector) GetServerCard() MCPServerCard`

GetServerCard returns the ServerCard field if non-nil, zero value otherwise.

### GetServerCardOk

`func (o *Connector) GetServerCardOk() (*MCPServerCard, bool)`

GetServerCardOk returns a tuple with the ServerCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCard

`func (o *Connector) SetServerCard(v MCPServerCard)`

SetServerCard sets ServerCard field to given value.

### HasServerCard

`func (o *Connector) HasServerCard() bool`

HasServerCard returns a boolean if a field has been set.

### GetOwnerId

`func (o *Connector) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *Connector) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *Connector) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *Connector) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetOwnerType

`func (o *Connector) GetOwnerType() ConsumerType`

GetOwnerType returns the OwnerType field if non-nil, zero value otherwise.

### GetOwnerTypeOk

`func (o *Connector) GetOwnerTypeOk() (*ConsumerType, bool)`

GetOwnerTypeOk returns a tuple with the OwnerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerType

`func (o *Connector) SetOwnerType(v ConsumerType)`

SetOwnerType sets OwnerType field to given value.


### GetVisibility

`func (o *Connector) GetVisibility() ResourceVisibility`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *Connector) GetVisibilityOk() (*ResourceVisibility, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *Connector) SetVisibility(v ResourceVisibility)`

SetVisibility sets Visibility field to given value.


### GetCreatorId

`func (o *Connector) GetCreatorId() string`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *Connector) GetCreatorIdOk() (*string, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *Connector) SetCreatorId(v string)`

SetCreatorId sets CreatorId field to given value.

### HasCreatorId

`func (o *Connector) HasCreatorId() bool`

HasCreatorId returns a boolean if a field has been set.

### GetLocale

`func (o *Connector) GetLocale() ConnectorLocale`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *Connector) GetLocaleOk() (*ConnectorLocale, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *Connector) SetLocale(v ConnectorLocale)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *Connector) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetSystemPrompt

`func (o *Connector) GetSystemPrompt() string`

GetSystemPrompt returns the SystemPrompt field if non-nil, zero value otherwise.

### GetSystemPromptOk

`func (o *Connector) GetSystemPromptOk() (*string, bool)`

GetSystemPromptOk returns a tuple with the SystemPrompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemPrompt

`func (o *Connector) SetSystemPrompt(v string)`

SetSystemPrompt sets SystemPrompt field to given value.

### HasSystemPrompt

`func (o *Connector) HasSystemPrompt() bool`

HasSystemPrompt returns a boolean if a field has been set.

### GetSupportedAuthMethods

`func (o *Connector) GetSupportedAuthMethods() []PublicAuthenticationMethod`

GetSupportedAuthMethods returns the SupportedAuthMethods field if non-nil, zero value otherwise.

### GetSupportedAuthMethodsOk

`func (o *Connector) GetSupportedAuthMethodsOk() (*[]PublicAuthenticationMethod, bool)`

GetSupportedAuthMethodsOk returns a tuple with the SupportedAuthMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedAuthMethods

`func (o *Connector) SetSupportedAuthMethods(v []PublicAuthenticationMethod)`

SetSupportedAuthMethods sets SupportedAuthMethods field to given value.

### HasSupportedAuthMethods

`func (o *Connector) HasSupportedAuthMethods() bool`

HasSupportedAuthMethods returns a boolean if a field has been set.

### GetConnectionPreferences

`func (o *Connector) GetConnectionPreferences() []ConnectionPreference`

GetConnectionPreferences returns the ConnectionPreferences field if non-nil, zero value otherwise.

### GetConnectionPreferencesOk

`func (o *Connector) GetConnectionPreferencesOk() (*[]ConnectionPreference, bool)`

GetConnectionPreferencesOk returns a tuple with the ConnectionPreferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionPreferences

`func (o *Connector) SetConnectionPreferences(v []ConnectionPreference)`

SetConnectionPreferences sets ConnectionPreferences field to given value.

### HasConnectionPreferences

`func (o *Connector) HasConnectionPreferences() bool`

HasConnectionPreferences returns a boolean if a field has been set.

### GetConnectionCredentials

`func (o *Connector) GetConnectionCredentials() []AuthenticationConfiguration`

GetConnectionCredentials returns the ConnectionCredentials field if non-nil, zero value otherwise.

### GetConnectionCredentialsOk

`func (o *Connector) GetConnectionCredentialsOk() (*[]AuthenticationConfiguration, bool)`

GetConnectionCredentialsOk returns a tuple with the ConnectionCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionCredentials

`func (o *Connector) SetConnectionCredentials(v []AuthenticationConfiguration)`

SetConnectionCredentials sets ConnectionCredentials field to given value.

### HasConnectionCredentials

`func (o *Connector) HasConnectionCredentials() bool`

HasConnectionCredentials returns a boolean if a field has been set.

### GetBlockedBy

`func (o *Connector) GetBlockedBy() []ConsumerType`

GetBlockedBy returns the BlockedBy field if non-nil, zero value otherwise.

### GetBlockedByOk

`func (o *Connector) GetBlockedByOk() (*[]ConsumerType, bool)`

GetBlockedByOk returns a tuple with the BlockedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedBy

`func (o *Connector) SetBlockedBy(v []ConsumerType)`

SetBlockedBy sets BlockedBy field to given value.

### HasBlockedBy

`func (o *Connector) HasBlockedBy() bool`

HasBlockedBy returns a boolean if a field has been set.

### GetPrivateToolExecution

`func (o *Connector) GetPrivateToolExecution() bool`

GetPrivateToolExecution returns the PrivateToolExecution field if non-nil, zero value otherwise.

### GetPrivateToolExecutionOk

`func (o *Connector) GetPrivateToolExecutionOk() (*bool, bool)`

GetPrivateToolExecutionOk returns a tuple with the PrivateToolExecution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateToolExecution

`func (o *Connector) SetPrivateToolExecution(v bool)`

SetPrivateToolExecution sets PrivateToolExecution field to given value.


### GetSupportsMcp

`func (o *Connector) GetSupportsMcp() bool`

GetSupportsMcp returns the SupportsMcp field if non-nil, zero value otherwise.

### GetSupportsMcpOk

`func (o *Connector) GetSupportsMcpOk() (*bool, bool)`

GetSupportsMcpOk returns a tuple with the SupportsMcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsMcp

`func (o *Connector) SetSupportsMcp(v bool)`

SetSupportsMcp sets SupportsMcp field to given value.


### GetMistral

`func (o *Connector) GetMistral() bool`

GetMistral returns the Mistral field if non-nil, zero value otherwise.

### GetMistralOk

`func (o *Connector) GetMistralOk() (*bool, bool)`

GetMistralOk returns a tuple with the Mistral field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMistral

`func (o *Connector) SetMistral(v bool)`

SetMistral sets Mistral field to given value.

### HasMistral

`func (o *Connector) HasMistral() bool`

HasMistral returns a boolean if a field has been set.

### GetIsAuthenticated

`func (o *Connector) GetIsAuthenticated() bool`

GetIsAuthenticated returns the IsAuthenticated field if non-nil, zero value otherwise.

### GetIsAuthenticatedOk

`func (o *Connector) GetIsAuthenticatedOk() (*bool, bool)`

GetIsAuthenticatedOk returns a tuple with the IsAuthenticated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAuthenticated

`func (o *Connector) SetIsAuthenticated(v bool)`

SetIsAuthenticated sets IsAuthenticated field to given value.

### HasIsAuthenticated

`func (o *Connector) HasIsAuthenticated() bool`

HasIsAuthenticated returns a boolean if a field has been set.

### GetTools

`func (o *Connector) GetTools() []ConnectorTool`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *Connector) GetToolsOk() (*[]ConnectorTool, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *Connector) SetTools(v []ConnectorTool)`

SetTools sets Tools field to given value.

### HasTools

`func (o *Connector) HasTools() bool`

HasTools returns a boolean if a field has been set.

### GetConnectionConfig

`func (o *Connector) GetConnectionConfig() PublicConnectionConfig`

GetConnectionConfig returns the ConnectionConfig field if non-nil, zero value otherwise.

### GetConnectionConfigOk

`func (o *Connector) GetConnectionConfigOk() (*PublicConnectionConfig, bool)`

GetConnectionConfigOk returns a tuple with the ConnectionConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionConfig

`func (o *Connector) SetConnectionConfig(v PublicConnectionConfig)`

SetConnectionConfig sets ConnectionConfig field to given value.

### HasConnectionConfig

`func (o *Connector) HasConnectionConfig() bool`

HasConnectionConfig returns a boolean if a field has been set.

### GetExecutionEnv

`func (o *Connector) GetExecutionEnv() PublicExecutionEnv`

GetExecutionEnv returns the ExecutionEnv field if non-nil, zero value otherwise.

### GetExecutionEnvOk

`func (o *Connector) GetExecutionEnvOk() (*PublicExecutionEnv, bool)`

GetExecutionEnvOk returns a tuple with the ExecutionEnv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionEnv

`func (o *Connector) SetExecutionEnv(v PublicExecutionEnv)`

SetExecutionEnv sets ExecutionEnv field to given value.

### HasExecutionEnv

`func (o *Connector) HasExecutionEnv() bool`

HasExecutionEnv returns a boolean if a field has been set.

### GetProtocol

`func (o *Connector) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *Connector) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *Connector) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


