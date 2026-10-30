# NumericStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "numeric"]
**Values** | Pointer to [**[]ScalarStatisticValue**](ScalarStatisticValue.md) |  | [optional] 
**SampleCount** | Pointer to **int32** |  | [optional] [default to 0]
**Avg** | Pointer to **NullableFloat32** |  | [optional] 
**Min** | Pointer to **NullableFloat32** |  | [optional] 
**Max** | Pointer to **NullableFloat32** |  | [optional] 
**Std** | Pointer to **NullableFloat32** |  | [optional] 
**Count** | Pointer to **int32** |  | [optional] [default to 0]

## Methods

### NewNumericStatistics

`func NewNumericStatistics() *NumericStatistics`

NewNumericStatistics instantiates a new NumericStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNumericStatisticsWithDefaults

`func NewNumericStatisticsWithDefaults() *NumericStatistics`

NewNumericStatisticsWithDefaults instantiates a new NumericStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *NumericStatistics) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NumericStatistics) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NumericStatistics) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NumericStatistics) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValues

`func (o *NumericStatistics) GetValues() []ScalarStatisticValue`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *NumericStatistics) GetValuesOk() (*[]ScalarStatisticValue, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *NumericStatistics) SetValues(v []ScalarStatisticValue)`

SetValues sets Values field to given value.

### HasValues

`func (o *NumericStatistics) HasValues() bool`

HasValues returns a boolean if a field has been set.

### SetValuesNil

`func (o *NumericStatistics) SetValuesNil(b bool)`

 SetValuesNil sets the value for Values to be an explicit nil

### UnsetValues
`func (o *NumericStatistics) UnsetValues()`

UnsetValues ensures that no value is present for Values, not even an explicit nil
### GetSampleCount

`func (o *NumericStatistics) GetSampleCount() int32`

GetSampleCount returns the SampleCount field if non-nil, zero value otherwise.

### GetSampleCountOk

`func (o *NumericStatistics) GetSampleCountOk() (*int32, bool)`

GetSampleCountOk returns a tuple with the SampleCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleCount

`func (o *NumericStatistics) SetSampleCount(v int32)`

SetSampleCount sets SampleCount field to given value.

### HasSampleCount

`func (o *NumericStatistics) HasSampleCount() bool`

HasSampleCount returns a boolean if a field has been set.

### GetAvg

`func (o *NumericStatistics) GetAvg() float32`

GetAvg returns the Avg field if non-nil, zero value otherwise.

### GetAvgOk

`func (o *NumericStatistics) GetAvgOk() (*float32, bool)`

GetAvgOk returns a tuple with the Avg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvg

`func (o *NumericStatistics) SetAvg(v float32)`

SetAvg sets Avg field to given value.

### HasAvg

`func (o *NumericStatistics) HasAvg() bool`

HasAvg returns a boolean if a field has been set.

### SetAvgNil

`func (o *NumericStatistics) SetAvgNil(b bool)`

 SetAvgNil sets the value for Avg to be an explicit nil

### UnsetAvg
`func (o *NumericStatistics) UnsetAvg()`

UnsetAvg ensures that no value is present for Avg, not even an explicit nil
### GetMin

`func (o *NumericStatistics) GetMin() float32`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *NumericStatistics) GetMinOk() (*float32, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *NumericStatistics) SetMin(v float32)`

SetMin sets Min field to given value.

### HasMin

`func (o *NumericStatistics) HasMin() bool`

HasMin returns a boolean if a field has been set.

### SetMinNil

`func (o *NumericStatistics) SetMinNil(b bool)`

 SetMinNil sets the value for Min to be an explicit nil

### UnsetMin
`func (o *NumericStatistics) UnsetMin()`

UnsetMin ensures that no value is present for Min, not even an explicit nil
### GetMax

`func (o *NumericStatistics) GetMax() float32`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *NumericStatistics) GetMaxOk() (*float32, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *NumericStatistics) SetMax(v float32)`

SetMax sets Max field to given value.

### HasMax

`func (o *NumericStatistics) HasMax() bool`

HasMax returns a boolean if a field has been set.

### SetMaxNil

`func (o *NumericStatistics) SetMaxNil(b bool)`

 SetMaxNil sets the value for Max to be an explicit nil

### UnsetMax
`func (o *NumericStatistics) UnsetMax()`

UnsetMax ensures that no value is present for Max, not even an explicit nil
### GetStd

`func (o *NumericStatistics) GetStd() float32`

GetStd returns the Std field if non-nil, zero value otherwise.

### GetStdOk

`func (o *NumericStatistics) GetStdOk() (*float32, bool)`

GetStdOk returns a tuple with the Std field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStd

`func (o *NumericStatistics) SetStd(v float32)`

SetStd sets Std field to given value.

### HasStd

`func (o *NumericStatistics) HasStd() bool`

HasStd returns a boolean if a field has been set.

### SetStdNil

`func (o *NumericStatistics) SetStdNil(b bool)`

 SetStdNil sets the value for Std to be an explicit nil

### UnsetStd
`func (o *NumericStatistics) UnsetStd()`

UnsetStd ensures that no value is present for Std, not even an explicit nil
### GetCount

`func (o *NumericStatistics) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *NumericStatistics) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *NumericStatistics) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *NumericStatistics) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


