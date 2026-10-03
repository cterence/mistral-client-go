# CustomConnectorAuthorization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "oauth2-token"]
**Value** | **string** |  | 

## Methods

### NewCustomConnectorAuthorization

`func NewCustomConnectorAuthorization(value string, ) *CustomConnectorAuthorization`

NewCustomConnectorAuthorization instantiates a new CustomConnectorAuthorization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomConnectorAuthorizationWithDefaults

`func NewCustomConnectorAuthorizationWithDefaults() *CustomConnectorAuthorization`

NewCustomConnectorAuthorizationWithDefaults instantiates a new CustomConnectorAuthorization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CustomConnectorAuthorization) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomConnectorAuthorization) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomConnectorAuthorization) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CustomConnectorAuthorization) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *CustomConnectorAuthorization) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CustomConnectorAuthorization) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CustomConnectorAuthorization) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


