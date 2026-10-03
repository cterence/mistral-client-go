# ConnectorAuthenticationHeader

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**IsRequired** | Pointer to **bool** |  | [optional] [default to true]
**IsSecret** | Pointer to **bool** |  | [optional] [default to true]

## Methods

### NewConnectorAuthenticationHeader

`func NewConnectorAuthenticationHeader(name string, ) *ConnectorAuthenticationHeader`

NewConnectorAuthenticationHeader instantiates a new ConnectorAuthenticationHeader object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorAuthenticationHeaderWithDefaults

`func NewConnectorAuthenticationHeaderWithDefaults() *ConnectorAuthenticationHeader`

NewConnectorAuthenticationHeaderWithDefaults instantiates a new ConnectorAuthenticationHeader object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ConnectorAuthenticationHeader) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnectorAuthenticationHeader) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnectorAuthenticationHeader) SetName(v string)`

SetName sets Name field to given value.


### GetIsRequired

`func (o *ConnectorAuthenticationHeader) GetIsRequired() bool`

GetIsRequired returns the IsRequired field if non-nil, zero value otherwise.

### GetIsRequiredOk

`func (o *ConnectorAuthenticationHeader) GetIsRequiredOk() (*bool, bool)`

GetIsRequiredOk returns a tuple with the IsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequired

`func (o *ConnectorAuthenticationHeader) SetIsRequired(v bool)`

SetIsRequired sets IsRequired field to given value.

### HasIsRequired

`func (o *ConnectorAuthenticationHeader) HasIsRequired() bool`

HasIsRequired returns a boolean if a field has been set.

### GetIsSecret

`func (o *ConnectorAuthenticationHeader) GetIsSecret() bool`

GetIsSecret returns the IsSecret field if non-nil, zero value otherwise.

### GetIsSecretOk

`func (o *ConnectorAuthenticationHeader) GetIsSecretOk() (*bool, bool)`

GetIsSecretOk returns a tuple with the IsSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSecret

`func (o *ConnectorAuthenticationHeader) SetIsSecret(v bool)`

SetIsSecret sets IsSecret field to given value.

### HasIsSecret

`func (o *ConnectorAuthenticationHeader) HasIsSecret() bool`

HasIsSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


