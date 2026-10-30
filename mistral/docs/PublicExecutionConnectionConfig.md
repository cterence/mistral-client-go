# PublicExecutionConnectionConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ConnectionConfigType**](ConnectionConfigType.md) |  | 
**Server** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Id** | Pointer to **NullableString** |  | [optional] 
**ToolConfiguration** | Pointer to [**NullableToolExecutionConfiguration**](ToolExecutionConfiguration.md) |  | [optional] 
**HostedInternally** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewPublicExecutionConnectionConfig

`func NewPublicExecutionConnectionConfig(type_ ConnectionConfigType, ) *PublicExecutionConnectionConfig`

NewPublicExecutionConnectionConfig instantiates a new PublicExecutionConnectionConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicExecutionConnectionConfigWithDefaults

`func NewPublicExecutionConnectionConfigWithDefaults() *PublicExecutionConnectionConfig`

NewPublicExecutionConnectionConfigWithDefaults instantiates a new PublicExecutionConnectionConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PublicExecutionConnectionConfig) GetType() ConnectionConfigType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PublicExecutionConnectionConfig) GetTypeOk() (*ConnectionConfigType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PublicExecutionConnectionConfig) SetType(v ConnectionConfigType)`

SetType sets Type field to given value.


### GetServer

`func (o *PublicExecutionConnectionConfig) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *PublicExecutionConnectionConfig) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *PublicExecutionConnectionConfig) SetServer(v string)`

SetServer sets Server field to given value.

### HasServer

`func (o *PublicExecutionConnectionConfig) HasServer() bool`

HasServer returns a boolean if a field has been set.

### SetServerNil

`func (o *PublicExecutionConnectionConfig) SetServerNil(b bool)`

 SetServerNil sets the value for Server to be an explicit nil

### UnsetServer
`func (o *PublicExecutionConnectionConfig) UnsetServer()`

UnsetServer ensures that no value is present for Server, not even an explicit nil
### GetName

`func (o *PublicExecutionConnectionConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PublicExecutionConnectionConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PublicExecutionConnectionConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PublicExecutionConnectionConfig) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PublicExecutionConnectionConfig) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PublicExecutionConnectionConfig) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetId

`func (o *PublicExecutionConnectionConfig) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PublicExecutionConnectionConfig) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PublicExecutionConnectionConfig) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PublicExecutionConnectionConfig) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *PublicExecutionConnectionConfig) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *PublicExecutionConnectionConfig) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetToolConfiguration

`func (o *PublicExecutionConnectionConfig) GetToolConfiguration() ToolExecutionConfiguration`

GetToolConfiguration returns the ToolConfiguration field if non-nil, zero value otherwise.

### GetToolConfigurationOk

`func (o *PublicExecutionConnectionConfig) GetToolConfigurationOk() (*ToolExecutionConfiguration, bool)`

GetToolConfigurationOk returns a tuple with the ToolConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolConfiguration

`func (o *PublicExecutionConnectionConfig) SetToolConfiguration(v ToolExecutionConfiguration)`

SetToolConfiguration sets ToolConfiguration field to given value.

### HasToolConfiguration

`func (o *PublicExecutionConnectionConfig) HasToolConfiguration() bool`

HasToolConfiguration returns a boolean if a field has been set.

### SetToolConfigurationNil

`func (o *PublicExecutionConnectionConfig) SetToolConfigurationNil(b bool)`

 SetToolConfigurationNil sets the value for ToolConfiguration to be an explicit nil

### UnsetToolConfiguration
`func (o *PublicExecutionConnectionConfig) UnsetToolConfiguration()`

UnsetToolConfiguration ensures that no value is present for ToolConfiguration, not even an explicit nil
### GetHostedInternally

`func (o *PublicExecutionConnectionConfig) GetHostedInternally() bool`

GetHostedInternally returns the HostedInternally field if non-nil, zero value otherwise.

### GetHostedInternallyOk

`func (o *PublicExecutionConnectionConfig) GetHostedInternallyOk() (*bool, bool)`

GetHostedInternallyOk returns a tuple with the HostedInternally field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostedInternally

`func (o *PublicExecutionConnectionConfig) SetHostedInternally(v bool)`

SetHostedInternally sets HostedInternally field to given value.

### HasHostedInternally

`func (o *PublicExecutionConnectionConfig) HasHostedInternally() bool`

HasHostedInternally returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


