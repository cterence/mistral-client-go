# EvaluatorV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**DeletedAt** | **NullableTime** |  | 
**Name** | **string** |  | 
**Description** | **NullableString** |  | 
**Tags** | **[]string** |  | 
**NumScores** | **int32** |  | 
**Goal** | Pointer to [**NullableGoalSpec**](GoalSpec.md) |  | [optional] 
**AggregateGoal** | Pointer to [**NullableGoalSpec**](GoalSpec.md) |  | [optional] 
**Statistics** | Pointer to [**[]EvaluatorV2StatisticsInner**](EvaluatorV2StatisticsInner.md) |  | [optional] 
**Weight** | Pointer to **float32** |  | [optional] [default to 1.0]
**Direction** | Pointer to **NullableString** |  | [optional] 
**MinValue** | Pointer to **NullableFloat32** |  | [optional] 
**MaxValue** | Pointer to **NullableFloat32** |  | [optional] 

## Methods

### NewEvaluatorV2

`func NewEvaluatorV2(id string, createdAt time.Time, updatedAt time.Time, deletedAt NullableTime, name string, description NullableString, tags []string, numScores int32, ) *EvaluatorV2`

NewEvaluatorV2 instantiates a new EvaluatorV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvaluatorV2WithDefaults

`func NewEvaluatorV2WithDefaults() *EvaluatorV2`

NewEvaluatorV2WithDefaults instantiates a new EvaluatorV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EvaluatorV2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EvaluatorV2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EvaluatorV2) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *EvaluatorV2) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EvaluatorV2) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EvaluatorV2) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *EvaluatorV2) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EvaluatorV2) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EvaluatorV2) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *EvaluatorV2) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *EvaluatorV2) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *EvaluatorV2) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.


### SetDeletedAtNil

`func (o *EvaluatorV2) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *EvaluatorV2) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetName

`func (o *EvaluatorV2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EvaluatorV2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EvaluatorV2) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *EvaluatorV2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EvaluatorV2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EvaluatorV2) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *EvaluatorV2) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *EvaluatorV2) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetTags

`func (o *EvaluatorV2) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *EvaluatorV2) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *EvaluatorV2) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetNumScores

`func (o *EvaluatorV2) GetNumScores() int32`

GetNumScores returns the NumScores field if non-nil, zero value otherwise.

### GetNumScoresOk

`func (o *EvaluatorV2) GetNumScoresOk() (*int32, bool)`

GetNumScoresOk returns a tuple with the NumScores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumScores

`func (o *EvaluatorV2) SetNumScores(v int32)`

SetNumScores sets NumScores field to given value.


### GetGoal

`func (o *EvaluatorV2) GetGoal() GoalSpec`

GetGoal returns the Goal field if non-nil, zero value otherwise.

### GetGoalOk

`func (o *EvaluatorV2) GetGoalOk() (*GoalSpec, bool)`

GetGoalOk returns a tuple with the Goal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoal

`func (o *EvaluatorV2) SetGoal(v GoalSpec)`

SetGoal sets Goal field to given value.

### HasGoal

`func (o *EvaluatorV2) HasGoal() bool`

HasGoal returns a boolean if a field has been set.

### SetGoalNil

`func (o *EvaluatorV2) SetGoalNil(b bool)`

 SetGoalNil sets the value for Goal to be an explicit nil

### UnsetGoal
`func (o *EvaluatorV2) UnsetGoal()`

UnsetGoal ensures that no value is present for Goal, not even an explicit nil
### GetAggregateGoal

`func (o *EvaluatorV2) GetAggregateGoal() GoalSpec`

GetAggregateGoal returns the AggregateGoal field if non-nil, zero value otherwise.

### GetAggregateGoalOk

`func (o *EvaluatorV2) GetAggregateGoalOk() (*GoalSpec, bool)`

GetAggregateGoalOk returns a tuple with the AggregateGoal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregateGoal

`func (o *EvaluatorV2) SetAggregateGoal(v GoalSpec)`

SetAggregateGoal sets AggregateGoal field to given value.

### HasAggregateGoal

`func (o *EvaluatorV2) HasAggregateGoal() bool`

HasAggregateGoal returns a boolean if a field has been set.

### SetAggregateGoalNil

`func (o *EvaluatorV2) SetAggregateGoalNil(b bool)`

 SetAggregateGoalNil sets the value for AggregateGoal to be an explicit nil

### UnsetAggregateGoal
`func (o *EvaluatorV2) UnsetAggregateGoal()`

UnsetAggregateGoal ensures that no value is present for AggregateGoal, not even an explicit nil
### GetStatistics

`func (o *EvaluatorV2) GetStatistics() []EvaluatorV2StatisticsInner`

GetStatistics returns the Statistics field if non-nil, zero value otherwise.

### GetStatisticsOk

`func (o *EvaluatorV2) GetStatisticsOk() (*[]EvaluatorV2StatisticsInner, bool)`

GetStatisticsOk returns a tuple with the Statistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatistics

`func (o *EvaluatorV2) SetStatistics(v []EvaluatorV2StatisticsInner)`

SetStatistics sets Statistics field to given value.

### HasStatistics

`func (o *EvaluatorV2) HasStatistics() bool`

HasStatistics returns a boolean if a field has been set.

### SetStatisticsNil

`func (o *EvaluatorV2) SetStatisticsNil(b bool)`

 SetStatisticsNil sets the value for Statistics to be an explicit nil

### UnsetStatistics
`func (o *EvaluatorV2) UnsetStatistics()`

UnsetStatistics ensures that no value is present for Statistics, not even an explicit nil
### GetWeight

`func (o *EvaluatorV2) GetWeight() float32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *EvaluatorV2) GetWeightOk() (*float32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *EvaluatorV2) SetWeight(v float32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *EvaluatorV2) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetDirection

`func (o *EvaluatorV2) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *EvaluatorV2) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *EvaluatorV2) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *EvaluatorV2) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### SetDirectionNil

`func (o *EvaluatorV2) SetDirectionNil(b bool)`

 SetDirectionNil sets the value for Direction to be an explicit nil

### UnsetDirection
`func (o *EvaluatorV2) UnsetDirection()`

UnsetDirection ensures that no value is present for Direction, not even an explicit nil
### GetMinValue

`func (o *EvaluatorV2) GetMinValue() float32`

GetMinValue returns the MinValue field if non-nil, zero value otherwise.

### GetMinValueOk

`func (o *EvaluatorV2) GetMinValueOk() (*float32, bool)`

GetMinValueOk returns a tuple with the MinValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinValue

`func (o *EvaluatorV2) SetMinValue(v float32)`

SetMinValue sets MinValue field to given value.

### HasMinValue

`func (o *EvaluatorV2) HasMinValue() bool`

HasMinValue returns a boolean if a field has been set.

### SetMinValueNil

`func (o *EvaluatorV2) SetMinValueNil(b bool)`

 SetMinValueNil sets the value for MinValue to be an explicit nil

### UnsetMinValue
`func (o *EvaluatorV2) UnsetMinValue()`

UnsetMinValue ensures that no value is present for MinValue, not even an explicit nil
### GetMaxValue

`func (o *EvaluatorV2) GetMaxValue() float32`

GetMaxValue returns the MaxValue field if non-nil, zero value otherwise.

### GetMaxValueOk

`func (o *EvaluatorV2) GetMaxValueOk() (*float32, bool)`

GetMaxValueOk returns a tuple with the MaxValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxValue

`func (o *EvaluatorV2) SetMaxValue(v float32)`

SetMaxValue sets MaxValue field to given value.

### HasMaxValue

`func (o *EvaluatorV2) HasMaxValue() bool`

HasMaxValue returns a boolean if a field has been set.

### SetMaxValueNil

`func (o *EvaluatorV2) SetMaxValueNil(b bool)`

 SetMaxValueNil sets the value for MaxValue to be an explicit nil

### UnsetMaxValue
`func (o *EvaluatorV2) UnsetMaxValue()`

UnsetMaxValue ensures that no value is present for MaxValue, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


