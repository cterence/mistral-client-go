# GoalSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operator** | Pointer to **NullableString** |  | [optional] 
**Value** | Pointer to [**NullableAnyOf**](anyOf&lt;&gt;.md) |  | [optional] 
**Min** | Pointer to **NullableFloat32** |  | [optional] 
**Max** | Pointer to **NullableFloat32** |  | [optional] 
**Values** | Pointer to **[]interface{}** |  | [optional] 
**Metric** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewGoalSpec

`func NewGoalSpec() *GoalSpec`

NewGoalSpec instantiates a new GoalSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoalSpecWithDefaults

`func NewGoalSpecWithDefaults() *GoalSpec`

NewGoalSpecWithDefaults instantiates a new GoalSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperator

`func (o *GoalSpec) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *GoalSpec) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *GoalSpec) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *GoalSpec) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### SetOperatorNil

`func (o *GoalSpec) SetOperatorNil(b bool)`

 SetOperatorNil sets the value for Operator to be an explicit nil

### UnsetOperator
`func (o *GoalSpec) UnsetOperator()`

UnsetOperator ensures that no value is present for Operator, not even an explicit nil
### GetValue

`func (o *GoalSpec) GetValue() AnyOf`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GoalSpec) GetValueOk() (*AnyOf, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GoalSpec) SetValue(v AnyOf)`

SetValue sets Value field to given value.

### HasValue

`func (o *GoalSpec) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *GoalSpec) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *GoalSpec) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetMin

`func (o *GoalSpec) GetMin() float32`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *GoalSpec) GetMinOk() (*float32, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *GoalSpec) SetMin(v float32)`

SetMin sets Min field to given value.

### HasMin

`func (o *GoalSpec) HasMin() bool`

HasMin returns a boolean if a field has been set.

### SetMinNil

`func (o *GoalSpec) SetMinNil(b bool)`

 SetMinNil sets the value for Min to be an explicit nil

### UnsetMin
`func (o *GoalSpec) UnsetMin()`

UnsetMin ensures that no value is present for Min, not even an explicit nil
### GetMax

`func (o *GoalSpec) GetMax() float32`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *GoalSpec) GetMaxOk() (*float32, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *GoalSpec) SetMax(v float32)`

SetMax sets Max field to given value.

### HasMax

`func (o *GoalSpec) HasMax() bool`

HasMax returns a boolean if a field has been set.

### SetMaxNil

`func (o *GoalSpec) SetMaxNil(b bool)`

 SetMaxNil sets the value for Max to be an explicit nil

### UnsetMax
`func (o *GoalSpec) UnsetMax()`

UnsetMax ensures that no value is present for Max, not even an explicit nil
### GetValues

`func (o *GoalSpec) GetValues() []interface{}`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GoalSpec) GetValuesOk() (*[]interface{}, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GoalSpec) SetValues(v []interface{})`

SetValues sets Values field to given value.

### HasValues

`func (o *GoalSpec) HasValues() bool`

HasValues returns a boolean if a field has been set.

### SetValuesNil

`func (o *GoalSpec) SetValuesNil(b bool)`

 SetValuesNil sets the value for Values to be an explicit nil

### UnsetValues
`func (o *GoalSpec) UnsetValues()`

UnsetValues ensures that no value is present for Values, not even an explicit nil
### GetMetric

`func (o *GoalSpec) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *GoalSpec) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *GoalSpec) SetMetric(v string)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *GoalSpec) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### SetMetricNil

`func (o *GoalSpec) SetMetricNil(b bool)`

 SetMetricNil sets the value for Metric to be an explicit nil

### UnsetMetric
`func (o *GoalSpec) UnsetMetric()`

UnsetMetric ensures that no value is present for Metric, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


