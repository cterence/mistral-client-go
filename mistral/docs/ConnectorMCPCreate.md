# ConnectorMCPCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the connector. Should be 64 char length maximum, alphanumeric, only underscores/dashes. | 
**Description** | **string** | The description of the connector. | 
**IconUrl** | Pointer to **NullableString** | The optional url of the icon you want to associate to the connector. | [optional] 
**Visibility** | Pointer to [**ResourceVisibility**](ResourceVisibility.md) | Visibility of the connector. Use &#39;shared_workspace&#39; for workspace scoped connectors, or &#39;private&#39; for private connectors. | [optional] [default to RESOURCEVISIBILITY_SHARED_ORG]
**Server** | **string** | The url of the MCP server. | 
**Headers** | Pointer to **map[string]interface{}** | Optional organization-level headers to be sent with the request to the mcp server. | [optional] 
**AuthData** | Pointer to [**NullableAuthData**](AuthData.md) | Optional additional authentication data for the connector. | [optional] 
**SystemPrompt** | Pointer to **NullableString** | Optional system prompt for the connector. | [optional] 

## Methods

### NewConnectorMCPCreate

`func NewConnectorMCPCreate(name string, description string, server string, ) *ConnectorMCPCreate`

NewConnectorMCPCreate instantiates a new ConnectorMCPCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorMCPCreateWithDefaults

`func NewConnectorMCPCreateWithDefaults() *ConnectorMCPCreate`

NewConnectorMCPCreateWithDefaults instantiates a new ConnectorMCPCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ConnectorMCPCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnectorMCPCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnectorMCPCreate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ConnectorMCPCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConnectorMCPCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConnectorMCPCreate) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetIconUrl

`func (o *ConnectorMCPCreate) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *ConnectorMCPCreate) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *ConnectorMCPCreate) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *ConnectorMCPCreate) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### SetIconUrlNil

`func (o *ConnectorMCPCreate) SetIconUrlNil(b bool)`

 SetIconUrlNil sets the value for IconUrl to be an explicit nil

### UnsetIconUrl
`func (o *ConnectorMCPCreate) UnsetIconUrl()`

UnsetIconUrl ensures that no value is present for IconUrl, not even an explicit nil
### GetVisibility

`func (o *ConnectorMCPCreate) GetVisibility() ResourceVisibility`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ConnectorMCPCreate) GetVisibilityOk() (*ResourceVisibility, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ConnectorMCPCreate) SetVisibility(v ResourceVisibility)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ConnectorMCPCreate) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetServer

`func (o *ConnectorMCPCreate) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *ConnectorMCPCreate) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *ConnectorMCPCreate) SetServer(v string)`

SetServer sets Server field to given value.


### GetHeaders

`func (o *ConnectorMCPCreate) GetHeaders() map[string]interface{}`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *ConnectorMCPCreate) GetHeadersOk() (*map[string]interface{}, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *ConnectorMCPCreate) SetHeaders(v map[string]interface{})`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *ConnectorMCPCreate) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### SetHeadersNil

`func (o *ConnectorMCPCreate) SetHeadersNil(b bool)`

 SetHeadersNil sets the value for Headers to be an explicit nil

### UnsetHeaders
`func (o *ConnectorMCPCreate) UnsetHeaders()`

UnsetHeaders ensures that no value is present for Headers, not even an explicit nil
### GetAuthData

`func (o *ConnectorMCPCreate) GetAuthData() AuthData`

GetAuthData returns the AuthData field if non-nil, zero value otherwise.

### GetAuthDataOk

`func (o *ConnectorMCPCreate) GetAuthDataOk() (*AuthData, bool)`

GetAuthDataOk returns a tuple with the AuthData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthData

`func (o *ConnectorMCPCreate) SetAuthData(v AuthData)`

SetAuthData sets AuthData field to given value.

### HasAuthData

`func (o *ConnectorMCPCreate) HasAuthData() bool`

HasAuthData returns a boolean if a field has been set.

### SetAuthDataNil

`func (o *ConnectorMCPCreate) SetAuthDataNil(b bool)`

 SetAuthDataNil sets the value for AuthData to be an explicit nil

### UnsetAuthData
`func (o *ConnectorMCPCreate) UnsetAuthData()`

UnsetAuthData ensures that no value is present for AuthData, not even an explicit nil
### GetSystemPrompt

`func (o *ConnectorMCPCreate) GetSystemPrompt() string`

GetSystemPrompt returns the SystemPrompt field if non-nil, zero value otherwise.

### GetSystemPromptOk

`func (o *ConnectorMCPCreate) GetSystemPromptOk() (*string, bool)`

GetSystemPromptOk returns a tuple with the SystemPrompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemPrompt

`func (o *ConnectorMCPCreate) SetSystemPrompt(v string)`

SetSystemPrompt sets SystemPrompt field to given value.

### HasSystemPrompt

`func (o *ConnectorMCPCreate) HasSystemPrompt() bool`

HasSystemPrompt returns a boolean if a field has been set.

### SetSystemPromptNil

`func (o *ConnectorMCPCreate) SetSystemPromptNil(b bool)`

 SetSystemPromptNil sets the value for SystemPrompt to be an explicit nil

### UnsetSystemPrompt
`func (o *ConnectorMCPCreate) UnsetSystemPrompt()`

UnsetSystemPrompt ensures that no value is present for SystemPrompt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


