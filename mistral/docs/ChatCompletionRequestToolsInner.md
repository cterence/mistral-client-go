# ChatCompletionRequestToolsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**ToolTypes**](ToolTypes.md) |  | [optional] [default to TOOLTYPES_FUNCTION]
**Function** | [**Function**](Function.md) |  | 
**ToolConfiguration** | Pointer to [**ToolConfiguration**](ToolConfiguration.md) |  | [optional] 
**LibraryIds** | **[]string** | Ids of the library in which to search. | 
**ConnectorId** | **string** |  | 
**Authorization** | Pointer to [**NullableCustomConnectorAuthorization**](CustomConnectorAuthorization.md) |  | [optional] 

## Methods

### NewChatCompletionRequestToolsInner

`func NewChatCompletionRequestToolsInner(function Function, libraryIds []string, connectorId string, ) *ChatCompletionRequestToolsInner`

NewChatCompletionRequestToolsInner instantiates a new ChatCompletionRequestToolsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatCompletionRequestToolsInnerWithDefaults

`func NewChatCompletionRequestToolsInnerWithDefaults() *ChatCompletionRequestToolsInner`

NewChatCompletionRequestToolsInnerWithDefaults instantiates a new ChatCompletionRequestToolsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ChatCompletionRequestToolsInner) GetType() ToolTypes`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ChatCompletionRequestToolsInner) GetTypeOk() (*ToolTypes, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ChatCompletionRequestToolsInner) SetType(v ToolTypes)`

SetType sets Type field to given value.

### HasType

`func (o *ChatCompletionRequestToolsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFunction

`func (o *ChatCompletionRequestToolsInner) GetFunction() Function`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *ChatCompletionRequestToolsInner) GetFunctionOk() (*Function, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *ChatCompletionRequestToolsInner) SetFunction(v Function)`

SetFunction sets Function field to given value.


### GetToolConfiguration

`func (o *ChatCompletionRequestToolsInner) GetToolConfiguration() ToolConfiguration`

GetToolConfiguration returns the ToolConfiguration field if non-nil, zero value otherwise.

### GetToolConfigurationOk

`func (o *ChatCompletionRequestToolsInner) GetToolConfigurationOk() (*ToolConfiguration, bool)`

GetToolConfigurationOk returns a tuple with the ToolConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolConfiguration

`func (o *ChatCompletionRequestToolsInner) SetToolConfiguration(v ToolConfiguration)`

SetToolConfiguration sets ToolConfiguration field to given value.

### HasToolConfiguration

`func (o *ChatCompletionRequestToolsInner) HasToolConfiguration() bool`

HasToolConfiguration returns a boolean if a field has been set.

### GetLibraryIds

`func (o *ChatCompletionRequestToolsInner) GetLibraryIds() []string`

GetLibraryIds returns the LibraryIds field if non-nil, zero value otherwise.

### GetLibraryIdsOk

`func (o *ChatCompletionRequestToolsInner) GetLibraryIdsOk() (*[]string, bool)`

GetLibraryIdsOk returns a tuple with the LibraryIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraryIds

`func (o *ChatCompletionRequestToolsInner) SetLibraryIds(v []string)`

SetLibraryIds sets LibraryIds field to given value.


### GetConnectorId

`func (o *ChatCompletionRequestToolsInner) GetConnectorId() string`

GetConnectorId returns the ConnectorId field if non-nil, zero value otherwise.

### GetConnectorIdOk

`func (o *ChatCompletionRequestToolsInner) GetConnectorIdOk() (*string, bool)`

GetConnectorIdOk returns a tuple with the ConnectorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorId

`func (o *ChatCompletionRequestToolsInner) SetConnectorId(v string)`

SetConnectorId sets ConnectorId field to given value.


### GetAuthorization

`func (o *ChatCompletionRequestToolsInner) GetAuthorization() CustomConnectorAuthorization`

GetAuthorization returns the Authorization field if non-nil, zero value otherwise.

### GetAuthorizationOk

`func (o *ChatCompletionRequestToolsInner) GetAuthorizationOk() (*CustomConnectorAuthorization, bool)`

GetAuthorizationOk returns a tuple with the Authorization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorization

`func (o *ChatCompletionRequestToolsInner) SetAuthorization(v CustomConnectorAuthorization)`

SetAuthorization sets Authorization field to given value.

### HasAuthorization

`func (o *ChatCompletionRequestToolsInner) HasAuthorization() bool`

HasAuthorization returns a boolean if a field has been set.

### SetAuthorizationNil

`func (o *ChatCompletionRequestToolsInner) SetAuthorizationNil(b bool)`

 SetAuthorizationNil sets the value for Authorization to be an explicit nil

### UnsetAuthorization
`func (o *ChatCompletionRequestToolsInner) UnsetAuthorization()`

UnsetAuthorization ensures that no value is present for Authorization, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


