# FineTuningDataJSON

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Training** | [**map[string]map[string][]map[string]interface{}**](map.md) | Fine-tuning training usage. | 
**Storage** | **map[string]int32** | Fine-tuning storage usage. | 

## Methods

### NewFineTuningDataJSON

`func NewFineTuningDataJSON(training map[string]map[string][]map[string]interface{}, storage map[string]int32, ) *FineTuningDataJSON`

NewFineTuningDataJSON instantiates a new FineTuningDataJSON object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFineTuningDataJSONWithDefaults

`func NewFineTuningDataJSONWithDefaults() *FineTuningDataJSON`

NewFineTuningDataJSONWithDefaults instantiates a new FineTuningDataJSON object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTraining

`func (o *FineTuningDataJSON) GetTraining() map[string]map[string][]map[string]interface{}`

GetTraining returns the Training field if non-nil, zero value otherwise.

### GetTrainingOk

`func (o *FineTuningDataJSON) GetTrainingOk() (*map[string]map[string][]map[string]interface{}, bool)`

GetTrainingOk returns a tuple with the Training field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraining

`func (o *FineTuningDataJSON) SetTraining(v map[string]map[string][]map[string]interface{})`

SetTraining sets Training field to given value.


### GetStorage

`func (o *FineTuningDataJSON) GetStorage() map[string]int32`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *FineTuningDataJSON) GetStorageOk() (*map[string]int32, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *FineTuningDataJSON) SetStorage(v map[string]int32)`

SetStorage sets Storage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


