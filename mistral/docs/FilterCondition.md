# FilterCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | **string** |  | 
**Op** | **string** |  | 
**Value** | **interface{}** |  | 

## Methods

### NewFilterCondition

`func NewFilterCondition(field string, op string, value interface{}, ) *FilterCondition`

NewFilterCondition instantiates a new FilterCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilterConditionWithDefaults

`func NewFilterConditionWithDefaults() *FilterCondition`

NewFilterConditionWithDefaults instantiates a new FilterCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *FilterCondition) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *FilterCondition) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *FilterCondition) SetField(v string)`

SetField sets Field field to given value.


### GetOp

`func (o *FilterCondition) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *FilterCondition) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *FilterCondition) SetOp(v string)`

SetOp sets Op field to given value.


### GetValue

`func (o *FilterCondition) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *FilterCondition) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *FilterCondition) SetValue(v interface{})`

SetValue sets Value field to given value.


### SetValueNil

`func (o *FilterCondition) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *FilterCondition) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


