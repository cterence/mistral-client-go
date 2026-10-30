# EvaluatorV2StatisticsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**Goal** | Pointer to [**GoalSpec**](GoalSpec.md) |  | [optional] 
**Percentile** | [**Percentile**](Percentile.md) |  | 
**PositiveLabel** | Pointer to **string** |  | [optional] 
**Average** | Pointer to **string** |  | [optional] 
**ExpectedKey** | Pointer to **string** |  | [optional] [default to "expected"]
**Labels** | Pointer to **[]string** |  | [optional] 

## Methods

### NewEvaluatorV2StatisticsInner

`func NewEvaluatorV2StatisticsInner(kind string, percentile Percentile, ) *EvaluatorV2StatisticsInner`

NewEvaluatorV2StatisticsInner instantiates a new EvaluatorV2StatisticsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvaluatorV2StatisticsInnerWithDefaults

`func NewEvaluatorV2StatisticsInnerWithDefaults() *EvaluatorV2StatisticsInner`

NewEvaluatorV2StatisticsInnerWithDefaults instantiates a new EvaluatorV2StatisticsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *EvaluatorV2StatisticsInner) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *EvaluatorV2StatisticsInner) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *EvaluatorV2StatisticsInner) SetKind(v string)`

SetKind sets Kind field to given value.


### GetGoal

`func (o *EvaluatorV2StatisticsInner) GetGoal() GoalSpec`

GetGoal returns the Goal field if non-nil, zero value otherwise.

### GetGoalOk

`func (o *EvaluatorV2StatisticsInner) GetGoalOk() (*GoalSpec, bool)`

GetGoalOk returns a tuple with the Goal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoal

`func (o *EvaluatorV2StatisticsInner) SetGoal(v GoalSpec)`

SetGoal sets Goal field to given value.

### HasGoal

`func (o *EvaluatorV2StatisticsInner) HasGoal() bool`

HasGoal returns a boolean if a field has been set.

### GetPercentile

`func (o *EvaluatorV2StatisticsInner) GetPercentile() Percentile`

GetPercentile returns the Percentile field if non-nil, zero value otherwise.

### GetPercentileOk

`func (o *EvaluatorV2StatisticsInner) GetPercentileOk() (*Percentile, bool)`

GetPercentileOk returns a tuple with the Percentile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentile

`func (o *EvaluatorV2StatisticsInner) SetPercentile(v Percentile)`

SetPercentile sets Percentile field to given value.


### GetPositiveLabel

`func (o *EvaluatorV2StatisticsInner) GetPositiveLabel() string`

GetPositiveLabel returns the PositiveLabel field if non-nil, zero value otherwise.

### GetPositiveLabelOk

`func (o *EvaluatorV2StatisticsInner) GetPositiveLabelOk() (*string, bool)`

GetPositiveLabelOk returns a tuple with the PositiveLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositiveLabel

`func (o *EvaluatorV2StatisticsInner) SetPositiveLabel(v string)`

SetPositiveLabel sets PositiveLabel field to given value.

### HasPositiveLabel

`func (o *EvaluatorV2StatisticsInner) HasPositiveLabel() bool`

HasPositiveLabel returns a boolean if a field has been set.

### GetAverage

`func (o *EvaluatorV2StatisticsInner) GetAverage() string`

GetAverage returns the Average field if non-nil, zero value otherwise.

### GetAverageOk

`func (o *EvaluatorV2StatisticsInner) GetAverageOk() (*string, bool)`

GetAverageOk returns a tuple with the Average field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverage

`func (o *EvaluatorV2StatisticsInner) SetAverage(v string)`

SetAverage sets Average field to given value.

### HasAverage

`func (o *EvaluatorV2StatisticsInner) HasAverage() bool`

HasAverage returns a boolean if a field has been set.

### GetExpectedKey

`func (o *EvaluatorV2StatisticsInner) GetExpectedKey() string`

GetExpectedKey returns the ExpectedKey field if non-nil, zero value otherwise.

### GetExpectedKeyOk

`func (o *EvaluatorV2StatisticsInner) GetExpectedKeyOk() (*string, bool)`

GetExpectedKeyOk returns a tuple with the ExpectedKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedKey

`func (o *EvaluatorV2StatisticsInner) SetExpectedKey(v string)`

SetExpectedKey sets ExpectedKey field to given value.

### HasExpectedKey

`func (o *EvaluatorV2StatisticsInner) HasExpectedKey() bool`

HasExpectedKey returns a boolean if a field has been set.

### GetLabels

`func (o *EvaluatorV2StatisticsInner) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *EvaluatorV2StatisticsInner) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *EvaluatorV2StatisticsInner) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *EvaluatorV2StatisticsInner) HasLabels() bool`

HasLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


