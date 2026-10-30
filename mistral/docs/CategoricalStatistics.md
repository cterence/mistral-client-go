# CategoricalStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "categorical"]
**Frequencies** | **map[string]int32** |  | 
**Mode** | **[]string** |  | 
**Count** | **int32** |  | 
**Values** | Pointer to [**[]ScalarStatisticValue**](ScalarStatisticValue.md) |  | [optional] 
**ExcludedCount** | Pointer to **int32** |  | [optional] [default to 0]
**ConfusionMatrix** | Pointer to [**NullableConfusionMatrix**](ConfusionMatrix.md) |  | [optional] 

## Methods

### NewCategoricalStatistics

`func NewCategoricalStatistics(frequencies map[string]int32, mode []string, count int32, ) *CategoricalStatistics`

NewCategoricalStatistics instantiates a new CategoricalStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoricalStatisticsWithDefaults

`func NewCategoricalStatisticsWithDefaults() *CategoricalStatistics`

NewCategoricalStatisticsWithDefaults instantiates a new CategoricalStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CategoricalStatistics) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CategoricalStatistics) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CategoricalStatistics) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CategoricalStatistics) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFrequencies

`func (o *CategoricalStatistics) GetFrequencies() map[string]int32`

GetFrequencies returns the Frequencies field if non-nil, zero value otherwise.

### GetFrequenciesOk

`func (o *CategoricalStatistics) GetFrequenciesOk() (*map[string]int32, bool)`

GetFrequenciesOk returns a tuple with the Frequencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencies

`func (o *CategoricalStatistics) SetFrequencies(v map[string]int32)`

SetFrequencies sets Frequencies field to given value.


### GetMode

`func (o *CategoricalStatistics) GetMode() []string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *CategoricalStatistics) GetModeOk() (*[]string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *CategoricalStatistics) SetMode(v []string)`

SetMode sets Mode field to given value.


### GetCount

`func (o *CategoricalStatistics) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *CategoricalStatistics) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *CategoricalStatistics) SetCount(v int32)`

SetCount sets Count field to given value.


### GetValues

`func (o *CategoricalStatistics) GetValues() []ScalarStatisticValue`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *CategoricalStatistics) GetValuesOk() (*[]ScalarStatisticValue, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *CategoricalStatistics) SetValues(v []ScalarStatisticValue)`

SetValues sets Values field to given value.

### HasValues

`func (o *CategoricalStatistics) HasValues() bool`

HasValues returns a boolean if a field has been set.

### SetValuesNil

`func (o *CategoricalStatistics) SetValuesNil(b bool)`

 SetValuesNil sets the value for Values to be an explicit nil

### UnsetValues
`func (o *CategoricalStatistics) UnsetValues()`

UnsetValues ensures that no value is present for Values, not even an explicit nil
### GetExcludedCount

`func (o *CategoricalStatistics) GetExcludedCount() int32`

GetExcludedCount returns the ExcludedCount field if non-nil, zero value otherwise.

### GetExcludedCountOk

`func (o *CategoricalStatistics) GetExcludedCountOk() (*int32, bool)`

GetExcludedCountOk returns a tuple with the ExcludedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedCount

`func (o *CategoricalStatistics) SetExcludedCount(v int32)`

SetExcludedCount sets ExcludedCount field to given value.

### HasExcludedCount

`func (o *CategoricalStatistics) HasExcludedCount() bool`

HasExcludedCount returns a boolean if a field has been set.

### GetConfusionMatrix

`func (o *CategoricalStatistics) GetConfusionMatrix() ConfusionMatrix`

GetConfusionMatrix returns the ConfusionMatrix field if non-nil, zero value otherwise.

### GetConfusionMatrixOk

`func (o *CategoricalStatistics) GetConfusionMatrixOk() (*ConfusionMatrix, bool)`

GetConfusionMatrixOk returns a tuple with the ConfusionMatrix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfusionMatrix

`func (o *CategoricalStatistics) SetConfusionMatrix(v ConfusionMatrix)`

SetConfusionMatrix sets ConfusionMatrix field to given value.

### HasConfusionMatrix

`func (o *CategoricalStatistics) HasConfusionMatrix() bool`

HasConfusionMatrix returns a boolean if a field has been set.

### SetConfusionMatrixNil

`func (o *CategoricalStatistics) SetConfusionMatrixNil(b bool)`

 SetConfusionMatrixNil sets the value for ConfusionMatrix to be an explicit nil

### UnsetConfusionMatrix
`func (o *CategoricalStatistics) UnsetConfusionMatrix()`

UnsetConfusionMatrix ensures that no value is present for ConfusionMatrix, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


