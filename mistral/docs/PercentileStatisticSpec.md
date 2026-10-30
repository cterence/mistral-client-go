# PercentileStatisticSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** |  | [optional] [default to "percentile"]
**Percentile** | [**Percentile**](Percentile.md) |  | 
**Goal** | Pointer to [**NullableGoalSpec**](GoalSpec.md) |  | [optional] 

## Methods

### NewPercentileStatisticSpec

`func NewPercentileStatisticSpec(percentile Percentile, ) *PercentileStatisticSpec`

NewPercentileStatisticSpec instantiates a new PercentileStatisticSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPercentileStatisticSpecWithDefaults

`func NewPercentileStatisticSpecWithDefaults() *PercentileStatisticSpec`

NewPercentileStatisticSpecWithDefaults instantiates a new PercentileStatisticSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *PercentileStatisticSpec) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *PercentileStatisticSpec) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *PercentileStatisticSpec) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *PercentileStatisticSpec) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetPercentile

`func (o *PercentileStatisticSpec) GetPercentile() Percentile`

GetPercentile returns the Percentile field if non-nil, zero value otherwise.

### GetPercentileOk

`func (o *PercentileStatisticSpec) GetPercentileOk() (*Percentile, bool)`

GetPercentileOk returns a tuple with the Percentile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentile

`func (o *PercentileStatisticSpec) SetPercentile(v Percentile)`

SetPercentile sets Percentile field to given value.


### GetGoal

`func (o *PercentileStatisticSpec) GetGoal() GoalSpec`

GetGoal returns the Goal field if non-nil, zero value otherwise.

### GetGoalOk

`func (o *PercentileStatisticSpec) GetGoalOk() (*GoalSpec, bool)`

GetGoalOk returns a tuple with the Goal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoal

`func (o *PercentileStatisticSpec) SetGoal(v GoalSpec)`

SetGoal sets Goal field to given value.

### HasGoal

`func (o *PercentileStatisticSpec) HasGoal() bool`

HasGoal returns a boolean if a field has been set.

### SetGoalNil

`func (o *PercentileStatisticSpec) SetGoalNil(b bool)`

 SetGoalNil sets the value for Goal to be an explicit nil

### UnsetGoal
`func (o *PercentileStatisticSpec) UnsetGoal()`

UnsetGoal ensures that no value is present for Goal, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


