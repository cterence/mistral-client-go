# ToolsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "function"]
**Function** | [**Function**](Function.md) |  | 
**ToolConfiguration** | Pointer to [**ToolConfiguration**](ToolConfiguration.md) |  | [optional] 
**LibraryIds** | **[]string** | Ids of the library in which to search. | 
**ConnectorId** | **string** |  | 
**Authorization** | Pointer to [**NullableCustomConnectorAuthorization**](CustomConnectorAuthorization.md) |  | [optional] 

## Methods

### NewToolsInner

`func NewToolsInner(function Function, libraryIds []string, connectorId string, ) *ToolsInner`

NewToolsInner instantiates a new ToolsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewToolsInnerWithDefaults

`func NewToolsInnerWithDefaults() *ToolsInner`

NewToolsInnerWithDefaults instantiates a new ToolsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ToolsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ToolsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ToolsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ToolsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFunction

`func (o *ToolsInner) GetFunction() Function`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *ToolsInner) GetFunctionOk() (*Function, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *ToolsInner) SetFunction(v Function)`

SetFunction sets Function field to given value.


### GetToolConfiguration

`func (o *ToolsInner) GetToolConfiguration() ToolConfiguration`

GetToolConfiguration returns the ToolConfiguration field if non-nil, zero value otherwise.

### GetToolConfigurationOk

`func (o *ToolsInner) GetToolConfigurationOk() (*ToolConfiguration, bool)`

GetToolConfigurationOk returns a tuple with the ToolConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolConfiguration

`func (o *ToolsInner) SetToolConfiguration(v ToolConfiguration)`

SetToolConfiguration sets ToolConfiguration field to given value.

### HasToolConfiguration

`func (o *ToolsInner) HasToolConfiguration() bool`

HasToolConfiguration returns a boolean if a field has been set.

### GetLibraryIds

`func (o *ToolsInner) GetLibraryIds() []string`

GetLibraryIds returns the LibraryIds field if non-nil, zero value otherwise.

### GetLibraryIdsOk

`func (o *ToolsInner) GetLibraryIdsOk() (*[]string, bool)`

GetLibraryIdsOk returns a tuple with the LibraryIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraryIds

`func (o *ToolsInner) SetLibraryIds(v []string)`

SetLibraryIds sets LibraryIds field to given value.


### GetConnectorId

`func (o *ToolsInner) GetConnectorId() string`

GetConnectorId returns the ConnectorId field if non-nil, zero value otherwise.

### GetConnectorIdOk

`func (o *ToolsInner) GetConnectorIdOk() (*string, bool)`

GetConnectorIdOk returns a tuple with the ConnectorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorId

`func (o *ToolsInner) SetConnectorId(v string)`

SetConnectorId sets ConnectorId field to given value.


### GetAuthorization

`func (o *ToolsInner) GetAuthorization() CustomConnectorAuthorization`

GetAuthorization returns the Authorization field if non-nil, zero value otherwise.

### GetAuthorizationOk

`func (o *ToolsInner) GetAuthorizationOk() (*CustomConnectorAuthorization, bool)`

GetAuthorizationOk returns a tuple with the Authorization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorization

`func (o *ToolsInner) SetAuthorization(v CustomConnectorAuthorization)`

SetAuthorization sets Authorization field to given value.

### HasAuthorization

`func (o *ToolsInner) HasAuthorization() bool`

HasAuthorization returns a boolean if a field has been set.

### SetAuthorizationNil

`func (o *ToolsInner) SetAuthorizationNil(b bool)`

 SetAuthorizationNil sets the value for Authorization to be an explicit nil

### UnsetAuthorization
`func (o *ToolsInner) UnsetAuthorization()`

UnsetAuthorization ensures that no value is present for Authorization, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


