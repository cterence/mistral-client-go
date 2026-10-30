# ConnectorMCPPublicUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **NullableString** | Optional human-readable title for the connector. | [optional] 
**Name** | Pointer to **NullableString** | The name of the connector. | [optional] 
**Description** | Pointer to **NullableString** | The description of the connector. | [optional] 
**IconUrl** | Pointer to **NullableString** | The optional url of the icon you want to associate to the connector. | [optional] 
**Server** | Pointer to **NullableString** | New connector server URL. | [optional] 
**AuthMethods** | Pointer to [**[]ConnectorMCPPublicUpdateAuthMethodsInner**](ConnectorMCPPublicUpdateAuthMethodsInner.md) | Authentication methods supported by the connector. | [optional] 
**Protocol** | Pointer to **string** |  | [optional] [default to "mcp"]
**SystemPrompt** | Pointer to **NullableString** | Optional system prompt for the connector. | [optional] 

## Methods

### NewConnectorMCPPublicUpdate

`func NewConnectorMCPPublicUpdate() *ConnectorMCPPublicUpdate`

NewConnectorMCPPublicUpdate instantiates a new ConnectorMCPPublicUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorMCPPublicUpdateWithDefaults

`func NewConnectorMCPPublicUpdateWithDefaults() *ConnectorMCPPublicUpdate`

NewConnectorMCPPublicUpdateWithDefaults instantiates a new ConnectorMCPPublicUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *ConnectorMCPPublicUpdate) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ConnectorMCPPublicUpdate) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ConnectorMCPPublicUpdate) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ConnectorMCPPublicUpdate) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *ConnectorMCPPublicUpdate) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ConnectorMCPPublicUpdate) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetName

`func (o *ConnectorMCPPublicUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnectorMCPPublicUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnectorMCPPublicUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConnectorMCPPublicUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ConnectorMCPPublicUpdate) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ConnectorMCPPublicUpdate) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *ConnectorMCPPublicUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConnectorMCPPublicUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConnectorMCPPublicUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConnectorMCPPublicUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ConnectorMCPPublicUpdate) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ConnectorMCPPublicUpdate) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIconUrl

`func (o *ConnectorMCPPublicUpdate) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *ConnectorMCPPublicUpdate) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *ConnectorMCPPublicUpdate) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *ConnectorMCPPublicUpdate) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### SetIconUrlNil

`func (o *ConnectorMCPPublicUpdate) SetIconUrlNil(b bool)`

 SetIconUrlNil sets the value for IconUrl to be an explicit nil

### UnsetIconUrl
`func (o *ConnectorMCPPublicUpdate) UnsetIconUrl()`

UnsetIconUrl ensures that no value is present for IconUrl, not even an explicit nil
### GetServer

`func (o *ConnectorMCPPublicUpdate) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *ConnectorMCPPublicUpdate) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *ConnectorMCPPublicUpdate) SetServer(v string)`

SetServer sets Server field to given value.

### HasServer

`func (o *ConnectorMCPPublicUpdate) HasServer() bool`

HasServer returns a boolean if a field has been set.

### SetServerNil

`func (o *ConnectorMCPPublicUpdate) SetServerNil(b bool)`

 SetServerNil sets the value for Server to be an explicit nil

### UnsetServer
`func (o *ConnectorMCPPublicUpdate) UnsetServer()`

UnsetServer ensures that no value is present for Server, not even an explicit nil
### GetAuthMethods

`func (o *ConnectorMCPPublicUpdate) GetAuthMethods() []ConnectorMCPPublicUpdateAuthMethodsInner`

GetAuthMethods returns the AuthMethods field if non-nil, zero value otherwise.

### GetAuthMethodsOk

`func (o *ConnectorMCPPublicUpdate) GetAuthMethodsOk() (*[]ConnectorMCPPublicUpdateAuthMethodsInner, bool)`

GetAuthMethodsOk returns a tuple with the AuthMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMethods

`func (o *ConnectorMCPPublicUpdate) SetAuthMethods(v []ConnectorMCPPublicUpdateAuthMethodsInner)`

SetAuthMethods sets AuthMethods field to given value.

### HasAuthMethods

`func (o *ConnectorMCPPublicUpdate) HasAuthMethods() bool`

HasAuthMethods returns a boolean if a field has been set.

### SetAuthMethodsNil

`func (o *ConnectorMCPPublicUpdate) SetAuthMethodsNil(b bool)`

 SetAuthMethodsNil sets the value for AuthMethods to be an explicit nil

### UnsetAuthMethods
`func (o *ConnectorMCPPublicUpdate) UnsetAuthMethods()`

UnsetAuthMethods ensures that no value is present for AuthMethods, not even an explicit nil
### GetProtocol

`func (o *ConnectorMCPPublicUpdate) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ConnectorMCPPublicUpdate) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ConnectorMCPPublicUpdate) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *ConnectorMCPPublicUpdate) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetSystemPrompt

`func (o *ConnectorMCPPublicUpdate) GetSystemPrompt() string`

GetSystemPrompt returns the SystemPrompt field if non-nil, zero value otherwise.

### GetSystemPromptOk

`func (o *ConnectorMCPPublicUpdate) GetSystemPromptOk() (*string, bool)`

GetSystemPromptOk returns a tuple with the SystemPrompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemPrompt

`func (o *ConnectorMCPPublicUpdate) SetSystemPrompt(v string)`

SetSystemPrompt sets SystemPrompt field to given value.

### HasSystemPrompt

`func (o *ConnectorMCPPublicUpdate) HasSystemPrompt() bool`

HasSystemPrompt returns a boolean if a field has been set.

### SetSystemPromptNil

`func (o *ConnectorMCPPublicUpdate) SetSystemPromptNil(b bool)`

 SetSystemPromptNil sets the value for SystemPrompt to be an explicit nil

### UnsetSystemPrompt
`func (o *ConnectorMCPPublicUpdate) UnsetSystemPrompt()`

UnsetSystemPrompt ensures that no value is present for SystemPrompt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


