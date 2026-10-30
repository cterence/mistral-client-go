# ScalarStatisticValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "scalar"]
**Id** | **string** |  | 
**Value** | Pointer to **NullableFloat32** |  | [optional] 

## Methods

### NewScalarStatisticValue

`func NewScalarStatisticValue(id string, ) *ScalarStatisticValue`

NewScalarStatisticValue instantiates a new ScalarStatisticValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScalarStatisticValueWithDefaults

`func NewScalarStatisticValueWithDefaults() *ScalarStatisticValue`

NewScalarStatisticValueWithDefaults instantiates a new ScalarStatisticValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ScalarStatisticValue) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ScalarStatisticValue) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ScalarStatisticValue) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ScalarStatisticValue) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *ScalarStatisticValue) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScalarStatisticValue) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScalarStatisticValue) SetId(v string)`

SetId sets Id field to given value.


### GetValue

`func (o *ScalarStatisticValue) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ScalarStatisticValue) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ScalarStatisticValue) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *ScalarStatisticValue) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *ScalarStatisticValue) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *ScalarStatisticValue) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


