# PublicAuthenticationMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MethodType** | [**OutboundAuthenticationType**](OutboundAuthenticationType.md) |  | 
**Headers** | Pointer to [**[]ConnectorAuthenticationHeader**](ConnectorAuthenticationHeader.md) |  | [optional] 

## Methods

### NewPublicAuthenticationMethod

`func NewPublicAuthenticationMethod(methodType OutboundAuthenticationType, ) *PublicAuthenticationMethod`

NewPublicAuthenticationMethod instantiates a new PublicAuthenticationMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicAuthenticationMethodWithDefaults

`func NewPublicAuthenticationMethodWithDefaults() *PublicAuthenticationMethod`

NewPublicAuthenticationMethodWithDefaults instantiates a new PublicAuthenticationMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethodType

`func (o *PublicAuthenticationMethod) GetMethodType() OutboundAuthenticationType`

GetMethodType returns the MethodType field if non-nil, zero value otherwise.

### GetMethodTypeOk

`func (o *PublicAuthenticationMethod) GetMethodTypeOk() (*OutboundAuthenticationType, bool)`

GetMethodTypeOk returns a tuple with the MethodType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodType

`func (o *PublicAuthenticationMethod) SetMethodType(v OutboundAuthenticationType)`

SetMethodType sets MethodType field to given value.


### GetHeaders

`func (o *PublicAuthenticationMethod) GetHeaders() []ConnectorAuthenticationHeader`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *PublicAuthenticationMethod) GetHeadersOk() (*[]ConnectorAuthenticationHeader, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *PublicAuthenticationMethod) SetHeaders(v []ConnectorAuthenticationHeader)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *PublicAuthenticationMethod) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### SetHeadersNil

`func (o *PublicAuthenticationMethod) SetHeadersNil(b bool)`

 SetHeadersNil sets the value for Headers to be an explicit nil

### UnsetHeaders
`func (o *PublicAuthenticationMethod) UnsetHeaders()`

UnsetHeaders ensures that no value is present for Headers, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


