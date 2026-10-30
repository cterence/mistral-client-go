# ScoreFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EvaluatorName** | **string** |  | 
**Metric** | **string** |  | 
**Operator** | [**ScoreOperator**](ScoreOperator.md) |  | 
**Value** | [**Value1**](Value1.md) |  | 

## Methods

### NewScoreFilter

`func NewScoreFilter(evaluatorName string, metric string, operator ScoreOperator, value Value1, ) *ScoreFilter`

NewScoreFilter instantiates a new ScoreFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScoreFilterWithDefaults

`func NewScoreFilterWithDefaults() *ScoreFilter`

NewScoreFilterWithDefaults instantiates a new ScoreFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvaluatorName

`func (o *ScoreFilter) GetEvaluatorName() string`

GetEvaluatorName returns the EvaluatorName field if non-nil, zero value otherwise.

### GetEvaluatorNameOk

`func (o *ScoreFilter) GetEvaluatorNameOk() (*string, bool)`

GetEvaluatorNameOk returns a tuple with the EvaluatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluatorName

`func (o *ScoreFilter) SetEvaluatorName(v string)`

SetEvaluatorName sets EvaluatorName field to given value.


### GetMetric

`func (o *ScoreFilter) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *ScoreFilter) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *ScoreFilter) SetMetric(v string)`

SetMetric sets Metric field to given value.


### GetOperator

`func (o *ScoreFilter) GetOperator() ScoreOperator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *ScoreFilter) GetOperatorOk() (*ScoreOperator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *ScoreFilter) SetOperator(v ScoreOperator)`

SetOperator sets Operator field to given value.


### GetValue

`func (o *ScoreFilter) GetValue() Value1`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ScoreFilter) GetValueOk() (*Value1, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ScoreFilter) SetValue(v Value1)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


