# EvaluatorStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "numeric"]
**Values** | Pointer to [**[]ScalarStatisticValue**](ScalarStatisticValue.md) |  | [optional] 
**SampleCount** | Pointer to **int32** |  | [optional] [default to 0]
**Avg** | Pointer to **float32** |  | [optional] 
**Min** | Pointer to **float32** |  | [optional] 
**Max** | Pointer to **float32** |  | [optional] 
**Std** | Pointer to **float32** |  | [optional] 
**Count** | **int32** |  | 
**Frequencies** | **map[string]int32** |  | 
**Mode** | **[]string** |  | 
**ExcludedCount** | Pointer to **int32** |  | [optional] [default to 0]
**ConfusionMatrix** | Pointer to [**ConfusionMatrix**](ConfusionMatrix.md) |  | [optional] 

## Methods

### NewEvaluatorStatistics

`func NewEvaluatorStatistics(count int32, frequencies map[string]int32, mode []string, ) *EvaluatorStatistics`

NewEvaluatorStatistics instantiates a new EvaluatorStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvaluatorStatisticsWithDefaults

`func NewEvaluatorStatisticsWithDefaults() *EvaluatorStatistics`

NewEvaluatorStatisticsWithDefaults instantiates a new EvaluatorStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EvaluatorStatistics) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EvaluatorStatistics) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EvaluatorStatistics) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EvaluatorStatistics) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValues

`func (o *EvaluatorStatistics) GetValues() []ScalarStatisticValue`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *EvaluatorStatistics) GetValuesOk() (*[]ScalarStatisticValue, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *EvaluatorStatistics) SetValues(v []ScalarStatisticValue)`

SetValues sets Values field to given value.

### HasValues

`func (o *EvaluatorStatistics) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetSampleCount

`func (o *EvaluatorStatistics) GetSampleCount() int32`

GetSampleCount returns the SampleCount field if non-nil, zero value otherwise.

### GetSampleCountOk

`func (o *EvaluatorStatistics) GetSampleCountOk() (*int32, bool)`

GetSampleCountOk returns a tuple with the SampleCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleCount

`func (o *EvaluatorStatistics) SetSampleCount(v int32)`

SetSampleCount sets SampleCount field to given value.

### HasSampleCount

`func (o *EvaluatorStatistics) HasSampleCount() bool`

HasSampleCount returns a boolean if a field has been set.

### GetAvg

`func (o *EvaluatorStatistics) GetAvg() float32`

GetAvg returns the Avg field if non-nil, zero value otherwise.

### GetAvgOk

`func (o *EvaluatorStatistics) GetAvgOk() (*float32, bool)`

GetAvgOk returns a tuple with the Avg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvg

`func (o *EvaluatorStatistics) SetAvg(v float32)`

SetAvg sets Avg field to given value.

### HasAvg

`func (o *EvaluatorStatistics) HasAvg() bool`

HasAvg returns a boolean if a field has been set.

### GetMin

`func (o *EvaluatorStatistics) GetMin() float32`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *EvaluatorStatistics) GetMinOk() (*float32, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *EvaluatorStatistics) SetMin(v float32)`

SetMin sets Min field to given value.

### HasMin

`func (o *EvaluatorStatistics) HasMin() bool`

HasMin returns a boolean if a field has been set.

### GetMax

`func (o *EvaluatorStatistics) GetMax() float32`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *EvaluatorStatistics) GetMaxOk() (*float32, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *EvaluatorStatistics) SetMax(v float32)`

SetMax sets Max field to given value.

### HasMax

`func (o *EvaluatorStatistics) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetStd

`func (o *EvaluatorStatistics) GetStd() float32`

GetStd returns the Std field if non-nil, zero value otherwise.

### GetStdOk

`func (o *EvaluatorStatistics) GetStdOk() (*float32, bool)`

GetStdOk returns a tuple with the Std field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStd

`func (o *EvaluatorStatistics) SetStd(v float32)`

SetStd sets Std field to given value.

### HasStd

`func (o *EvaluatorStatistics) HasStd() bool`

HasStd returns a boolean if a field has been set.

### GetCount

`func (o *EvaluatorStatistics) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *EvaluatorStatistics) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *EvaluatorStatistics) SetCount(v int32)`

SetCount sets Count field to given value.


### GetFrequencies

`func (o *EvaluatorStatistics) GetFrequencies() map[string]int32`

GetFrequencies returns the Frequencies field if non-nil, zero value otherwise.

### GetFrequenciesOk

`func (o *EvaluatorStatistics) GetFrequenciesOk() (*map[string]int32, bool)`

GetFrequenciesOk returns a tuple with the Frequencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencies

`func (o *EvaluatorStatistics) SetFrequencies(v map[string]int32)`

SetFrequencies sets Frequencies field to given value.


### GetMode

`func (o *EvaluatorStatistics) GetMode() []string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *EvaluatorStatistics) GetModeOk() (*[]string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *EvaluatorStatistics) SetMode(v []string)`

SetMode sets Mode field to given value.


### GetExcludedCount

`func (o *EvaluatorStatistics) GetExcludedCount() int32`

GetExcludedCount returns the ExcludedCount field if non-nil, zero value otherwise.

### GetExcludedCountOk

`func (o *EvaluatorStatistics) GetExcludedCountOk() (*int32, bool)`

GetExcludedCountOk returns a tuple with the ExcludedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedCount

`func (o *EvaluatorStatistics) SetExcludedCount(v int32)`

SetExcludedCount sets ExcludedCount field to given value.

### HasExcludedCount

`func (o *EvaluatorStatistics) HasExcludedCount() bool`

HasExcludedCount returns a boolean if a field has been set.

### GetConfusionMatrix

`func (o *EvaluatorStatistics) GetConfusionMatrix() ConfusionMatrix`

GetConfusionMatrix returns the ConfusionMatrix field if non-nil, zero value otherwise.

### GetConfusionMatrixOk

`func (o *EvaluatorStatistics) GetConfusionMatrixOk() (*ConfusionMatrix, bool)`

GetConfusionMatrixOk returns a tuple with the ConfusionMatrix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfusionMatrix

`func (o *EvaluatorStatistics) SetConfusionMatrix(v ConfusionMatrix)`

SetConfusionMatrix sets ConfusionMatrix field to given value.

### HasConfusionMatrix

`func (o *EvaluatorStatistics) HasConfusionMatrix() bool`

HasConfusionMatrix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


