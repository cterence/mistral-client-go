# GoalResultItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metric** | **string** |  | 
**Operator** | **string** |  | 
**Target** | Pointer to **NullableFloat32** |  | [optional] 
**TargetMin** | Pointer to **NullableFloat32** |  | [optional] 
**TargetMax** | Pointer to **NullableFloat32** |  | [optional] 
**Actual** | **float32** |  | 
**Passed** | **bool** |  | 
**Direction** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewGoalResultItem

`func NewGoalResultItem(metric string, operator string, actual float32, passed bool, ) *GoalResultItem`

NewGoalResultItem instantiates a new GoalResultItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoalResultItemWithDefaults

`func NewGoalResultItemWithDefaults() *GoalResultItem`

NewGoalResultItemWithDefaults instantiates a new GoalResultItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetric

`func (o *GoalResultItem) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *GoalResultItem) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *GoalResultItem) SetMetric(v string)`

SetMetric sets Metric field to given value.


### GetOperator

`func (o *GoalResultItem) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *GoalResultItem) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *GoalResultItem) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetTarget

`func (o *GoalResultItem) GetTarget() float32`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *GoalResultItem) GetTargetOk() (*float32, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *GoalResultItem) SetTarget(v float32)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *GoalResultItem) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### SetTargetNil

`func (o *GoalResultItem) SetTargetNil(b bool)`

 SetTargetNil sets the value for Target to be an explicit nil

### UnsetTarget
`func (o *GoalResultItem) UnsetTarget()`

UnsetTarget ensures that no value is present for Target, not even an explicit nil
### GetTargetMin

`func (o *GoalResultItem) GetTargetMin() float32`

GetTargetMin returns the TargetMin field if non-nil, zero value otherwise.

### GetTargetMinOk

`func (o *GoalResultItem) GetTargetMinOk() (*float32, bool)`

GetTargetMinOk returns a tuple with the TargetMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMin

`func (o *GoalResultItem) SetTargetMin(v float32)`

SetTargetMin sets TargetMin field to given value.

### HasTargetMin

`func (o *GoalResultItem) HasTargetMin() bool`

HasTargetMin returns a boolean if a field has been set.

### SetTargetMinNil

`func (o *GoalResultItem) SetTargetMinNil(b bool)`

 SetTargetMinNil sets the value for TargetMin to be an explicit nil

### UnsetTargetMin
`func (o *GoalResultItem) UnsetTargetMin()`

UnsetTargetMin ensures that no value is present for TargetMin, not even an explicit nil
### GetTargetMax

`func (o *GoalResultItem) GetTargetMax() float32`

GetTargetMax returns the TargetMax field if non-nil, zero value otherwise.

### GetTargetMaxOk

`func (o *GoalResultItem) GetTargetMaxOk() (*float32, bool)`

GetTargetMaxOk returns a tuple with the TargetMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMax

`func (o *GoalResultItem) SetTargetMax(v float32)`

SetTargetMax sets TargetMax field to given value.

### HasTargetMax

`func (o *GoalResultItem) HasTargetMax() bool`

HasTargetMax returns a boolean if a field has been set.

### SetTargetMaxNil

`func (o *GoalResultItem) SetTargetMaxNil(b bool)`

 SetTargetMaxNil sets the value for TargetMax to be an explicit nil

### UnsetTargetMax
`func (o *GoalResultItem) UnsetTargetMax()`

UnsetTargetMax ensures that no value is present for TargetMax, not even an explicit nil
### GetActual

`func (o *GoalResultItem) GetActual() float32`

GetActual returns the Actual field if non-nil, zero value otherwise.

### GetActualOk

`func (o *GoalResultItem) GetActualOk() (*float32, bool)`

GetActualOk returns a tuple with the Actual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActual

`func (o *GoalResultItem) SetActual(v float32)`

SetActual sets Actual field to given value.


### GetPassed

`func (o *GoalResultItem) GetPassed() bool`

GetPassed returns the Passed field if non-nil, zero value otherwise.

### GetPassedOk

`func (o *GoalResultItem) GetPassedOk() (*bool, bool)`

GetPassedOk returns a tuple with the Passed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassed

`func (o *GoalResultItem) SetPassed(v bool)`

SetPassed sets Passed field to given value.


### GetDirection

`func (o *GoalResultItem) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *GoalResultItem) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *GoalResultItem) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *GoalResultItem) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### SetDirectionNil

`func (o *GoalResultItem) SetDirectionNil(b bool)`

 SetDirectionNil sets the value for Direction to be an explicit nil

### UnsetDirection
`func (o *GoalResultItem) UnsetDirection()`

UnsetDirection ensures that no value is present for Direction, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


