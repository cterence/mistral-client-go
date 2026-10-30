# BearerAuthMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MethodType** | Pointer to **string** |  | [optional] [default to "bearer"]
**Headers** | Pointer to [**[]ConnectorAuthenticationHeader**](ConnectorAuthenticationHeader.md) | Optional additional headers sent with requests. | [optional] 

## Methods

### NewBearerAuthMethod

`func NewBearerAuthMethod() *BearerAuthMethod`

NewBearerAuthMethod instantiates a new BearerAuthMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBearerAuthMethodWithDefaults

`func NewBearerAuthMethodWithDefaults() *BearerAuthMethod`

NewBearerAuthMethodWithDefaults instantiates a new BearerAuthMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethodType

`func (o *BearerAuthMethod) GetMethodType() string`

GetMethodType returns the MethodType field if non-nil, zero value otherwise.

### GetMethodTypeOk

`func (o *BearerAuthMethod) GetMethodTypeOk() (*string, bool)`

GetMethodTypeOk returns a tuple with the MethodType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodType

`func (o *BearerAuthMethod) SetMethodType(v string)`

SetMethodType sets MethodType field to given value.

### HasMethodType

`func (o *BearerAuthMethod) HasMethodType() bool`

HasMethodType returns a boolean if a field has been set.

### GetHeaders

`func (o *BearerAuthMethod) GetHeaders() []ConnectorAuthenticationHeader`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *BearerAuthMethod) GetHeadersOk() (*[]ConnectorAuthenticationHeader, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *BearerAuthMethod) SetHeaders(v []ConnectorAuthenticationHeader)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *BearerAuthMethod) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### SetHeadersNil

`func (o *BearerAuthMethod) SetHeadersNil(b bool)`

 SetHeadersNil sets the value for Headers to be an explicit nil

### UnsetHeaders
`func (o *BearerAuthMethod) UnsetHeaders()`

UnsetHeaders ensures that no value is present for Headers, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


