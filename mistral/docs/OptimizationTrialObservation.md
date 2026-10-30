# OptimizationTrialObservation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObservationType** | Pointer to **NullableString** |  | [optional] 
**Score** | Pointer to **NullableFloat32** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] [default to {}]
**Run** | [**EvaluationRunV2Response**](EvaluationRunV2Response.md) |  | 

## Methods

### NewOptimizationTrialObservation

`func NewOptimizationTrialObservation(run EvaluationRunV2Response, ) *OptimizationTrialObservation`

NewOptimizationTrialObservation instantiates a new OptimizationTrialObservation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptimizationTrialObservationWithDefaults

`func NewOptimizationTrialObservationWithDefaults() *OptimizationTrialObservation`

NewOptimizationTrialObservationWithDefaults instantiates a new OptimizationTrialObservation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObservationType

`func (o *OptimizationTrialObservation) GetObservationType() string`

GetObservationType returns the ObservationType field if non-nil, zero value otherwise.

### GetObservationTypeOk

`func (o *OptimizationTrialObservation) GetObservationTypeOk() (*string, bool)`

GetObservationTypeOk returns a tuple with the ObservationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservationType

`func (o *OptimizationTrialObservation) SetObservationType(v string)`

SetObservationType sets ObservationType field to given value.

### HasObservationType

`func (o *OptimizationTrialObservation) HasObservationType() bool`

HasObservationType returns a boolean if a field has been set.

### SetObservationTypeNil

`func (o *OptimizationTrialObservation) SetObservationTypeNil(b bool)`

 SetObservationTypeNil sets the value for ObservationType to be an explicit nil

### UnsetObservationType
`func (o *OptimizationTrialObservation) UnsetObservationType()`

UnsetObservationType ensures that no value is present for ObservationType, not even an explicit nil
### GetScore

`func (o *OptimizationTrialObservation) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *OptimizationTrialObservation) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *OptimizationTrialObservation) SetScore(v float32)`

SetScore sets Score field to given value.

### HasScore

`func (o *OptimizationTrialObservation) HasScore() bool`

HasScore returns a boolean if a field has been set.

### SetScoreNil

`func (o *OptimizationTrialObservation) SetScoreNil(b bool)`

 SetScoreNil sets the value for Score to be an explicit nil

### UnsetScore
`func (o *OptimizationTrialObservation) UnsetScore()`

UnsetScore ensures that no value is present for Score, not even an explicit nil
### GetMetadata

`func (o *OptimizationTrialObservation) GetMetadata() map[string]*interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *OptimizationTrialObservation) GetMetadataOk() (*map[string]*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *OptimizationTrialObservation) SetMetadata(v map[string]*interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *OptimizationTrialObservation) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetRun

`func (o *OptimizationTrialObservation) GetRun() EvaluationRunV2Response`

GetRun returns the Run field if non-nil, zero value otherwise.

### GetRunOk

`func (o *OptimizationTrialObservation) GetRunOk() (*EvaluationRunV2Response, bool)`

GetRunOk returns a tuple with the Run field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRun

`func (o *OptimizationTrialObservation) SetRun(v EvaluationRunV2Response)`

SetRun sets Run field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


