# ScoreSortInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** |  | [optional] [default to "score"]
**EvaluatorName** | **string** |  | 
**Metric** | Pointer to [**NumericScoreMetric**](NumericScoreMetric.md) |  | [optional] [default to NUMERICSCOREMETRIC_AVG]
**Direction** | Pointer to [**SortDirection**](SortDirection.md) |  | [optional] [default to SORTDIRECTION_ASC]

## Methods

### NewScoreSortInput

`func NewScoreSortInput(evaluatorName string, ) *ScoreSortInput`

NewScoreSortInput instantiates a new ScoreSortInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScoreSortInputWithDefaults

`func NewScoreSortInputWithDefaults() *ScoreSortInput`

NewScoreSortInputWithDefaults instantiates a new ScoreSortInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *ScoreSortInput) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ScoreSortInput) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ScoreSortInput) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ScoreSortInput) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetEvaluatorName

`func (o *ScoreSortInput) GetEvaluatorName() string`

GetEvaluatorName returns the EvaluatorName field if non-nil, zero value otherwise.

### GetEvaluatorNameOk

`func (o *ScoreSortInput) GetEvaluatorNameOk() (*string, bool)`

GetEvaluatorNameOk returns a tuple with the EvaluatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluatorName

`func (o *ScoreSortInput) SetEvaluatorName(v string)`

SetEvaluatorName sets EvaluatorName field to given value.


### GetMetric

`func (o *ScoreSortInput) GetMetric() NumericScoreMetric`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *ScoreSortInput) GetMetricOk() (*NumericScoreMetric, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *ScoreSortInput) SetMetric(v NumericScoreMetric)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *ScoreSortInput) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetDirection

`func (o *ScoreSortInput) GetDirection() SortDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *ScoreSortInput) GetDirectionOk() (*SortDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *ScoreSortInput) SetDirection(v SortDirection)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *ScoreSortInput) HasDirection() bool`

HasDirection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


