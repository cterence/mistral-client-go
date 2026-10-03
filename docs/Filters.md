# Filters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AND** | Pointer to [**[]FilterGroupANDInner**](FilterGroupANDInner.md) |  | [optional] 
**OR** | Pointer to [**[]FilterGroupANDInner**](FilterGroupANDInner.md) |  | [optional] 
**Field** | **string** |  | 
**Op** | **string** |  | 
**Value** | **interface{}** |  | 

## Methods

### NewFilters

`func NewFilters(field string, op string, value interface{}, ) *Filters`

NewFilters instantiates a new Filters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFiltersWithDefaults

`func NewFiltersWithDefaults() *Filters`

NewFiltersWithDefaults instantiates a new Filters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAND

`func (o *Filters) GetAND() []FilterGroupANDInner`

GetAND returns the AND field if non-nil, zero value otherwise.

### GetANDOk

`func (o *Filters) GetANDOk() (*[]FilterGroupANDInner, bool)`

GetANDOk returns a tuple with the AND field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAND

`func (o *Filters) SetAND(v []FilterGroupANDInner)`

SetAND sets AND field to given value.

### HasAND

`func (o *Filters) HasAND() bool`

HasAND returns a boolean if a field has been set.

### GetOR

`func (o *Filters) GetOR() []FilterGroupANDInner`

GetOR returns the OR field if non-nil, zero value otherwise.

### GetOROk

`func (o *Filters) GetOROk() (*[]FilterGroupANDInner, bool)`

GetOROk returns a tuple with the OR field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOR

`func (o *Filters) SetOR(v []FilterGroupANDInner)`

SetOR sets OR field to given value.

### HasOR

`func (o *Filters) HasOR() bool`

HasOR returns a boolean if a field has been set.

### GetField

`func (o *Filters) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *Filters) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *Filters) SetField(v string)`

SetField sets Field field to given value.


### GetOp

`func (o *Filters) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *Filters) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *Filters) SetOp(v string)`

SetOp sets Op field to given value.


### GetValue

`func (o *Filters) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Filters) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Filters) SetValue(v interface{})`

SetValue sets Value field to given value.


### SetValueNil

`func (o *Filters) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *Filters) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


