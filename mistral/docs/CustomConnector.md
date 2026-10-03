# CustomConnector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "connector"]
**ConnectorId** | **string** |  | 
**Authorization** | Pointer to [**NullableCustomConnectorAuthorization**](CustomConnectorAuthorization.md) |  | [optional] 
**ToolConfiguration** | Pointer to [**NullableToolConfiguration**](ToolConfiguration.md) |  | [optional] 

## Methods

### NewCustomConnector

`func NewCustomConnector(connectorId string, ) *CustomConnector`

NewCustomConnector instantiates a new CustomConnector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomConnectorWithDefaults

`func NewCustomConnectorWithDefaults() *CustomConnector`

NewCustomConnectorWithDefaults instantiates a new CustomConnector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CustomConnector) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomConnector) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomConnector) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CustomConnector) HasType() bool`

HasType returns a boolean if a field has been set.

### GetConnectorId

`func (o *CustomConnector) GetConnectorId() string`

GetConnectorId returns the ConnectorId field if non-nil, zero value otherwise.

### GetConnectorIdOk

`func (o *CustomConnector) GetConnectorIdOk() (*string, bool)`

GetConnectorIdOk returns a tuple with the ConnectorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorId

`func (o *CustomConnector) SetConnectorId(v string)`

SetConnectorId sets ConnectorId field to given value.


### GetAuthorization

`func (o *CustomConnector) GetAuthorization() CustomConnectorAuthorization`

GetAuthorization returns the Authorization field if non-nil, zero value otherwise.

### GetAuthorizationOk

`func (o *CustomConnector) GetAuthorizationOk() (*CustomConnectorAuthorization, bool)`

GetAuthorizationOk returns a tuple with the Authorization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorization

`func (o *CustomConnector) SetAuthorization(v CustomConnectorAuthorization)`

SetAuthorization sets Authorization field to given value.

### HasAuthorization

`func (o *CustomConnector) HasAuthorization() bool`

HasAuthorization returns a boolean if a field has been set.

### SetAuthorizationNil

`func (o *CustomConnector) SetAuthorizationNil(b bool)`

 SetAuthorizationNil sets the value for Authorization to be an explicit nil

### UnsetAuthorization
`func (o *CustomConnector) UnsetAuthorization()`

UnsetAuthorization ensures that no value is present for Authorization, not even an explicit nil
### GetToolConfiguration

`func (o *CustomConnector) GetToolConfiguration() ToolConfiguration`

GetToolConfiguration returns the ToolConfiguration field if non-nil, zero value otherwise.

### GetToolConfigurationOk

`func (o *CustomConnector) GetToolConfigurationOk() (*ToolConfiguration, bool)`

GetToolConfigurationOk returns a tuple with the ToolConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolConfiguration

`func (o *CustomConnector) SetToolConfiguration(v ToolConfiguration)`

SetToolConfiguration sets ToolConfiguration field to given value.

### HasToolConfiguration

`func (o *CustomConnector) HasToolConfiguration() bool`

HasToolConfiguration returns a boolean if a field has been set.

### SetToolConfigurationNil

`func (o *CustomConnector) SetToolConfigurationNil(b bool)`

 SetToolConfigurationNil sets the value for ToolConfiguration to be an explicit nil

### UnsetToolConfiguration
`func (o *CustomConnector) UnsetToolConfiguration()`

UnsetToolConfiguration ensures that no value is present for ToolConfiguration, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


