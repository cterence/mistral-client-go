# ClassificationStatisticSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**PositiveLabel** | Pointer to **NullableString** |  | [optional] 
**Average** | Pointer to **NullableString** |  | [optional] 
**ExpectedKey** | Pointer to **string** |  | [optional] [default to "expected"]
**Labels** | Pointer to **[]string** |  | [optional] 
**Goal** | Pointer to [**NullableGoalSpec**](GoalSpec.md) |  | [optional] 

## Methods

### NewClassificationStatisticSpec

`func NewClassificationStatisticSpec(kind string, ) *ClassificationStatisticSpec`

NewClassificationStatisticSpec instantiates a new ClassificationStatisticSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClassificationStatisticSpecWithDefaults

`func NewClassificationStatisticSpecWithDefaults() *ClassificationStatisticSpec`

NewClassificationStatisticSpecWithDefaults instantiates a new ClassificationStatisticSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *ClassificationStatisticSpec) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ClassificationStatisticSpec) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ClassificationStatisticSpec) SetKind(v string)`

SetKind sets Kind field to given value.


### GetPositiveLabel

`func (o *ClassificationStatisticSpec) GetPositiveLabel() string`

GetPositiveLabel returns the PositiveLabel field if non-nil, zero value otherwise.

### GetPositiveLabelOk

`func (o *ClassificationStatisticSpec) GetPositiveLabelOk() (*string, bool)`

GetPositiveLabelOk returns a tuple with the PositiveLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositiveLabel

`func (o *ClassificationStatisticSpec) SetPositiveLabel(v string)`

SetPositiveLabel sets PositiveLabel field to given value.

### HasPositiveLabel

`func (o *ClassificationStatisticSpec) HasPositiveLabel() bool`

HasPositiveLabel returns a boolean if a field has been set.

### SetPositiveLabelNil

`func (o *ClassificationStatisticSpec) SetPositiveLabelNil(b bool)`

 SetPositiveLabelNil sets the value for PositiveLabel to be an explicit nil

### UnsetPositiveLabel
`func (o *ClassificationStatisticSpec) UnsetPositiveLabel()`

UnsetPositiveLabel ensures that no value is present for PositiveLabel, not even an explicit nil
### GetAverage

`func (o *ClassificationStatisticSpec) GetAverage() string`

GetAverage returns the Average field if non-nil, zero value otherwise.

### GetAverageOk

`func (o *ClassificationStatisticSpec) GetAverageOk() (*string, bool)`

GetAverageOk returns a tuple with the Average field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverage

`func (o *ClassificationStatisticSpec) SetAverage(v string)`

SetAverage sets Average field to given value.

### HasAverage

`func (o *ClassificationStatisticSpec) HasAverage() bool`

HasAverage returns a boolean if a field has been set.

### SetAverageNil

`func (o *ClassificationStatisticSpec) SetAverageNil(b bool)`

 SetAverageNil sets the value for Average to be an explicit nil

### UnsetAverage
`func (o *ClassificationStatisticSpec) UnsetAverage()`

UnsetAverage ensures that no value is present for Average, not even an explicit nil
### GetExpectedKey

`func (o *ClassificationStatisticSpec) GetExpectedKey() string`

GetExpectedKey returns the ExpectedKey field if non-nil, zero value otherwise.

### GetExpectedKeyOk

`func (o *ClassificationStatisticSpec) GetExpectedKeyOk() (*string, bool)`

GetExpectedKeyOk returns a tuple with the ExpectedKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedKey

`func (o *ClassificationStatisticSpec) SetExpectedKey(v string)`

SetExpectedKey sets ExpectedKey field to given value.

### HasExpectedKey

`func (o *ClassificationStatisticSpec) HasExpectedKey() bool`

HasExpectedKey returns a boolean if a field has been set.

### GetLabels

`func (o *ClassificationStatisticSpec) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ClassificationStatisticSpec) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ClassificationStatisticSpec) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ClassificationStatisticSpec) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *ClassificationStatisticSpec) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *ClassificationStatisticSpec) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetGoal

`func (o *ClassificationStatisticSpec) GetGoal() GoalSpec`

GetGoal returns the Goal field if non-nil, zero value otherwise.

### GetGoalOk

`func (o *ClassificationStatisticSpec) GetGoalOk() (*GoalSpec, bool)`

GetGoalOk returns a tuple with the Goal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoal

`func (o *ClassificationStatisticSpec) SetGoal(v GoalSpec)`

SetGoal sets Goal field to given value.

### HasGoal

`func (o *ClassificationStatisticSpec) HasGoal() bool`

HasGoal returns a boolean if a field has been set.

### SetGoalNil

`func (o *ClassificationStatisticSpec) SetGoalNil(b bool)`

 SetGoalNil sets the value for Goal to be an explicit nil

### UnsetGoal
`func (o *ClassificationStatisticSpec) UnsetGoal()`

UnsetGoal ensures that no value is present for Goal, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


