# ConnectorMCPUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | The name of the connector. | [optional] 
**Description** | Pointer to **NullableString** | The description of the connector. | [optional] 
**IconUrl** | Pointer to **NullableString** | The optional url of the icon you want to associate to the connector. | [optional] 
**SystemPrompt** | Pointer to **NullableString** | Optional system prompt for the connector. | [optional] 
**ConnectionConfig** | Pointer to **map[string]interface{}** | Optional new connection config. | [optional] 
**ConnectionSecrets** | Pointer to **map[string]interface{}** | Optional new connection secrets | [optional] 
**Server** | Pointer to **NullableString** | New server url for your mcp connector. | [optional] 
**Headers** | Pointer to **map[string]interface{}** | New headers for your mcp connector. | [optional] 
**AuthData** | Pointer to [**NullableAuthData**](AuthData.md) | New authentication data for your mcp connector. | [optional] 

## Methods

### NewConnectorMCPUpdate

`func NewConnectorMCPUpdate() *ConnectorMCPUpdate`

NewConnectorMCPUpdate instantiates a new ConnectorMCPUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorMCPUpdateWithDefaults

`func NewConnectorMCPUpdateWithDefaults() *ConnectorMCPUpdate`

NewConnectorMCPUpdateWithDefaults instantiates a new ConnectorMCPUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ConnectorMCPUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnectorMCPUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnectorMCPUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConnectorMCPUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ConnectorMCPUpdate) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ConnectorMCPUpdate) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *ConnectorMCPUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConnectorMCPUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConnectorMCPUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConnectorMCPUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ConnectorMCPUpdate) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ConnectorMCPUpdate) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIconUrl

`func (o *ConnectorMCPUpdate) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *ConnectorMCPUpdate) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *ConnectorMCPUpdate) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *ConnectorMCPUpdate) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### SetIconUrlNil

`func (o *ConnectorMCPUpdate) SetIconUrlNil(b bool)`

 SetIconUrlNil sets the value for IconUrl to be an explicit nil

### UnsetIconUrl
`func (o *ConnectorMCPUpdate) UnsetIconUrl()`

UnsetIconUrl ensures that no value is present for IconUrl, not even an explicit nil
### GetSystemPrompt

`func (o *ConnectorMCPUpdate) GetSystemPrompt() string`

GetSystemPrompt returns the SystemPrompt field if non-nil, zero value otherwise.

### GetSystemPromptOk

`func (o *ConnectorMCPUpdate) GetSystemPromptOk() (*string, bool)`

GetSystemPromptOk returns a tuple with the SystemPrompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemPrompt

`func (o *ConnectorMCPUpdate) SetSystemPrompt(v string)`

SetSystemPrompt sets SystemPrompt field to given value.

### HasSystemPrompt

`func (o *ConnectorMCPUpdate) HasSystemPrompt() bool`

HasSystemPrompt returns a boolean if a field has been set.

### SetSystemPromptNil

`func (o *ConnectorMCPUpdate) SetSystemPromptNil(b bool)`

 SetSystemPromptNil sets the value for SystemPrompt to be an explicit nil

### UnsetSystemPrompt
`func (o *ConnectorMCPUpdate) UnsetSystemPrompt()`

UnsetSystemPrompt ensures that no value is present for SystemPrompt, not even an explicit nil
### GetConnectionConfig

`func (o *ConnectorMCPUpdate) GetConnectionConfig() map[string]interface{}`

GetConnectionConfig returns the ConnectionConfig field if non-nil, zero value otherwise.

### GetConnectionConfigOk

`func (o *ConnectorMCPUpdate) GetConnectionConfigOk() (*map[string]interface{}, bool)`

GetConnectionConfigOk returns a tuple with the ConnectionConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionConfig

`func (o *ConnectorMCPUpdate) SetConnectionConfig(v map[string]interface{})`

SetConnectionConfig sets ConnectionConfig field to given value.

### HasConnectionConfig

`func (o *ConnectorMCPUpdate) HasConnectionConfig() bool`

HasConnectionConfig returns a boolean if a field has been set.

### SetConnectionConfigNil

`func (o *ConnectorMCPUpdate) SetConnectionConfigNil(b bool)`

 SetConnectionConfigNil sets the value for ConnectionConfig to be an explicit nil

### UnsetConnectionConfig
`func (o *ConnectorMCPUpdate) UnsetConnectionConfig()`

UnsetConnectionConfig ensures that no value is present for ConnectionConfig, not even an explicit nil
### GetConnectionSecrets

`func (o *ConnectorMCPUpdate) GetConnectionSecrets() map[string]interface{}`

GetConnectionSecrets returns the ConnectionSecrets field if non-nil, zero value otherwise.

### GetConnectionSecretsOk

`func (o *ConnectorMCPUpdate) GetConnectionSecretsOk() (*map[string]interface{}, bool)`

GetConnectionSecretsOk returns a tuple with the ConnectionSecrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionSecrets

`func (o *ConnectorMCPUpdate) SetConnectionSecrets(v map[string]interface{})`

SetConnectionSecrets sets ConnectionSecrets field to given value.

### HasConnectionSecrets

`func (o *ConnectorMCPUpdate) HasConnectionSecrets() bool`

HasConnectionSecrets returns a boolean if a field has been set.

### SetConnectionSecretsNil

`func (o *ConnectorMCPUpdate) SetConnectionSecretsNil(b bool)`

 SetConnectionSecretsNil sets the value for ConnectionSecrets to be an explicit nil

### UnsetConnectionSecrets
`func (o *ConnectorMCPUpdate) UnsetConnectionSecrets()`

UnsetConnectionSecrets ensures that no value is present for ConnectionSecrets, not even an explicit nil
### GetServer

`func (o *ConnectorMCPUpdate) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *ConnectorMCPUpdate) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *ConnectorMCPUpdate) SetServer(v string)`

SetServer sets Server field to given value.

### HasServer

`func (o *ConnectorMCPUpdate) HasServer() bool`

HasServer returns a boolean if a field has been set.

### SetServerNil

`func (o *ConnectorMCPUpdate) SetServerNil(b bool)`

 SetServerNil sets the value for Server to be an explicit nil

### UnsetServer
`func (o *ConnectorMCPUpdate) UnsetServer()`

UnsetServer ensures that no value is present for Server, not even an explicit nil
### GetHeaders

`func (o *ConnectorMCPUpdate) GetHeaders() map[string]interface{}`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *ConnectorMCPUpdate) GetHeadersOk() (*map[string]interface{}, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *ConnectorMCPUpdate) SetHeaders(v map[string]interface{})`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *ConnectorMCPUpdate) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### SetHeadersNil

`func (o *ConnectorMCPUpdate) SetHeadersNil(b bool)`

 SetHeadersNil sets the value for Headers to be an explicit nil

### UnsetHeaders
`func (o *ConnectorMCPUpdate) UnsetHeaders()`

UnsetHeaders ensures that no value is present for Headers, not even an explicit nil
### GetAuthData

`func (o *ConnectorMCPUpdate) GetAuthData() AuthData`

GetAuthData returns the AuthData field if non-nil, zero value otherwise.

### GetAuthDataOk

`func (o *ConnectorMCPUpdate) GetAuthDataOk() (*AuthData, bool)`

GetAuthDataOk returns a tuple with the AuthData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthData

`func (o *ConnectorMCPUpdate) SetAuthData(v AuthData)`

SetAuthData sets AuthData field to given value.

### HasAuthData

`func (o *ConnectorMCPUpdate) HasAuthData() bool`

HasAuthData returns a boolean if a field has been set.

### SetAuthDataNil

`func (o *ConnectorMCPUpdate) SetAuthDataNil(b bool)`

 SetAuthDataNil sets the value for AuthData to be an explicit nil

### UnsetAuthData
`func (o *ConnectorMCPUpdate) UnsetAuthData()`

UnsetAuthData ensures that no value is present for AuthData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


