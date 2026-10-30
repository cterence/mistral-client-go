# EvaluationRunRecordFilterInputSort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** |  | [optional] [default to "score"]
**Field** | [**RecordSortField**](RecordSortField.md) |  | 
**Direction** | Pointer to [**SortDirection**](SortDirection.md) |  | [optional] [default to SORTDIRECTION_ASC]
**EvaluatorName** | **string** |  | 
**Metric** | Pointer to [**NumericScoreMetric**](NumericScoreMetric.md) |  | [optional] [default to NUMERICSCOREMETRIC_AVG]

## Methods

### NewEvaluationRunRecordFilterInputSort

`func NewEvaluationRunRecordFilterInputSort(field RecordSortField, evaluatorName string, ) *EvaluationRunRecordFilterInputSort`

NewEvaluationRunRecordFilterInputSort instantiates a new EvaluationRunRecordFilterInputSort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvaluationRunRecordFilterInputSortWithDefaults

`func NewEvaluationRunRecordFilterInputSortWithDefaults() *EvaluationRunRecordFilterInputSort`

NewEvaluationRunRecordFilterInputSortWithDefaults instantiates a new EvaluationRunRecordFilterInputSort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *EvaluationRunRecordFilterInputSort) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *EvaluationRunRecordFilterInputSort) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *EvaluationRunRecordFilterInputSort) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *EvaluationRunRecordFilterInputSort) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetField

`func (o *EvaluationRunRecordFilterInputSort) GetField() RecordSortField`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *EvaluationRunRecordFilterInputSort) GetFieldOk() (*RecordSortField, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *EvaluationRunRecordFilterInputSort) SetField(v RecordSortField)`

SetField sets Field field to given value.


### GetDirection

`func (o *EvaluationRunRecordFilterInputSort) GetDirection() SortDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *EvaluationRunRecordFilterInputSort) GetDirectionOk() (*SortDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *EvaluationRunRecordFilterInputSort) SetDirection(v SortDirection)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *EvaluationRunRecordFilterInputSort) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetEvaluatorName

`func (o *EvaluationRunRecordFilterInputSort) GetEvaluatorName() string`

GetEvaluatorName returns the EvaluatorName field if non-nil, zero value otherwise.

### GetEvaluatorNameOk

`func (o *EvaluationRunRecordFilterInputSort) GetEvaluatorNameOk() (*string, bool)`

GetEvaluatorNameOk returns a tuple with the EvaluatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluatorName

`func (o *EvaluationRunRecordFilterInputSort) SetEvaluatorName(v string)`

SetEvaluatorName sets EvaluatorName field to given value.


### GetMetric

`func (o *EvaluationRunRecordFilterInputSort) GetMetric() NumericScoreMetric`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *EvaluationRunRecordFilterInputSort) GetMetricOk() (*NumericScoreMetric, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *EvaluationRunRecordFilterInputSort) SetMetric(v NumericScoreMetric)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *EvaluationRunRecordFilterInputSort) HasMetric() bool`

HasMetric returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


