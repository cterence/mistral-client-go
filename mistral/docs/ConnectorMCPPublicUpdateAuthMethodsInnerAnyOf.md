# ConnectorMCPPublicUpdateAuthMethodsInnerAnyOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MethodType** | Pointer to **string** |  | [optional] [default to "none"]
**Headers** | Pointer to [**[]ConnectorAuthenticationHeader**](ConnectorAuthenticationHeader.md) | Headers whose values must be provided as credentials. | [optional] 

## Methods

### NewConnectorMCPPublicUpdateAuthMethodsInnerAnyOf

`func NewConnectorMCPPublicUpdateAuthMethodsInnerAnyOf() *ConnectorMCPPublicUpdateAuthMethodsInnerAnyOf`

NewConnectorMCPPublicUpdateAuthMethodsInnerAnyOf instantiates a new ConnectorMCPPublicUpdateAuthMethodsInnerAnyOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorMCPPublicUpdateAuthMethodsInnerAnyOfWithDefaults

`func NewConnectorMCPPublicUpdateAuthMethodsInnerAnyOfWithDefaults() *ConnectorMCPPublicUpdateAuthMethodsInnerAnyOf`

NewConnectorMCPPublicUpdateAuthMethodsInnerAnyOfWithDefaults instantiates a new ConnectorMCPPublicUpdateAuthMethodsInnerAnyOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethodType

`func (o *ConnectorMCPPublicUpdateAuthMethodsInnerAnyOf) GetMethodType() string`

GetMethodType returns the MethodType field if non-nil, zero value otherwise.

### GetMethodTypeOk

`func (o *ConnectorMCPPublicUpdateAuthMethodsInnerAnyOf) GetMethodTypeOk() (*string, bool)`

GetMethodTypeOk returns a tuple with the MethodType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodType

`func (o *ConnectorMCPPublicUpdateAuthMethodsInnerAnyOf) SetMethodType(v string)`

SetMethodType sets MethodType field to given value.

### HasMethodType

`func (o *ConnectorMCPPublicUpdateAuthMethodsInnerAnyOf) HasMethodType() bool`

HasMethodType returns a boolean if a field has been set.

### GetHeaders

`func (o *ConnectorMCPPublicUpdateAuthMethodsInnerAnyOf) GetHeaders() []ConnectorAuthenticationHeader`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *ConnectorMCPPublicUpdateAuthMethodsInnerAnyOf) GetHeadersOk() (*[]ConnectorAuthenticationHeader, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *ConnectorMCPPublicUpdateAuthMethodsInnerAnyOf) SetHeaders(v []ConnectorAuthenticationHeader)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *ConnectorMCPPublicUpdateAuthMethodsInnerAnyOf) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


