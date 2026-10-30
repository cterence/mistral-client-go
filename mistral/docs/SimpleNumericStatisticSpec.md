# SimpleNumericStatisticSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**Goal** | Pointer to [**NullableGoalSpec**](GoalSpec.md) |  | [optional] 

## Methods

### NewSimpleNumericStatisticSpec

`func NewSimpleNumericStatisticSpec(kind string, ) *SimpleNumericStatisticSpec`

NewSimpleNumericStatisticSpec instantiates a new SimpleNumericStatisticSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimpleNumericStatisticSpecWithDefaults

`func NewSimpleNumericStatisticSpecWithDefaults() *SimpleNumericStatisticSpec`

NewSimpleNumericStatisticSpecWithDefaults instantiates a new SimpleNumericStatisticSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *SimpleNumericStatisticSpec) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SimpleNumericStatisticSpec) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SimpleNumericStatisticSpec) SetKind(v string)`

SetKind sets Kind field to given value.


### GetGoal

`func (o *SimpleNumericStatisticSpec) GetGoal() GoalSpec`

GetGoal returns the Goal field if non-nil, zero value otherwise.

### GetGoalOk

`func (o *SimpleNumericStatisticSpec) GetGoalOk() (*GoalSpec, bool)`

GetGoalOk returns a tuple with the Goal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoal

`func (o *SimpleNumericStatisticSpec) SetGoal(v GoalSpec)`

SetGoal sets Goal field to given value.

### HasGoal

`func (o *SimpleNumericStatisticSpec) HasGoal() bool`

HasGoal returns a boolean if a field has been set.

### SetGoalNil

`func (o *SimpleNumericStatisticSpec) SetGoalNil(b bool)`

 SetGoalNil sets the value for Goal to be an explicit nil

### UnsetGoal
`func (o *SimpleNumericStatisticSpec) UnsetGoal()`

UnsetGoal ensures that no value is present for Goal, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


