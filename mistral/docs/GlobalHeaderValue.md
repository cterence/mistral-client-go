# GlobalHeaderValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsSecret** | Pointer to **bool** |  | [optional] [default to true]
**Value** | **string** |  | 

## Methods

### NewGlobalHeaderValue

`func NewGlobalHeaderValue(value string, ) *GlobalHeaderValue`

NewGlobalHeaderValue instantiates a new GlobalHeaderValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalHeaderValueWithDefaults

`func NewGlobalHeaderValueWithDefaults() *GlobalHeaderValue`

NewGlobalHeaderValueWithDefaults instantiates a new GlobalHeaderValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsSecret

`func (o *GlobalHeaderValue) GetIsSecret() bool`

GetIsSecret returns the IsSecret field if non-nil, zero value otherwise.

### GetIsSecretOk

`func (o *GlobalHeaderValue) GetIsSecretOk() (*bool, bool)`

GetIsSecretOk returns a tuple with the IsSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSecret

`func (o *GlobalHeaderValue) SetIsSecret(v bool)`

SetIsSecret sets IsSecret field to given value.

### HasIsSecret

`func (o *GlobalHeaderValue) HasIsSecret() bool`

HasIsSecret returns a boolean if a field has been set.

### GetValue

`func (o *GlobalHeaderValue) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GlobalHeaderValue) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GlobalHeaderValue) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


