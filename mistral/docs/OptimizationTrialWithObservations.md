# OptimizationTrialWithObservations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**DeletedAt** | **NullableTime** |  | 
**OptimizationId** | **string** |  | 
**AlgorithmTrialKey** | **string** |  | 
**Status** | **string** |  | 
**ParentTrialId** | Pointer to **NullableString** |  | [optional] 
**Hypothesis** | Pointer to **NullableString** |  | [optional] 
**Resolution** | Pointer to **NullableString** |  | [optional] 
**Metadata** | **map[string]interface{}** |  | 
**Observations** | Pointer to [**[]OptimizationTrialObservation**](OptimizationTrialObservation.md) |  | [optional] [default to {}]

## Methods

### NewOptimizationTrialWithObservations

`func NewOptimizationTrialWithObservations(id string, createdAt time.Time, updatedAt time.Time, deletedAt NullableTime, optimizationId string, algorithmTrialKey string, status string, metadata map[string]*interface{}, ) *OptimizationTrialWithObservations`

NewOptimizationTrialWithObservations instantiates a new OptimizationTrialWithObservations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptimizationTrialWithObservationsWithDefaults

`func NewOptimizationTrialWithObservationsWithDefaults() *OptimizationTrialWithObservations`

NewOptimizationTrialWithObservationsWithDefaults instantiates a new OptimizationTrialWithObservations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OptimizationTrialWithObservations) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OptimizationTrialWithObservations) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OptimizationTrialWithObservations) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *OptimizationTrialWithObservations) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OptimizationTrialWithObservations) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OptimizationTrialWithObservations) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *OptimizationTrialWithObservations) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OptimizationTrialWithObservations) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OptimizationTrialWithObservations) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *OptimizationTrialWithObservations) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *OptimizationTrialWithObservations) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *OptimizationTrialWithObservations) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.


### SetDeletedAtNil

`func (o *OptimizationTrialWithObservations) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *OptimizationTrialWithObservations) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetOptimizationId

`func (o *OptimizationTrialWithObservations) GetOptimizationId() string`

GetOptimizationId returns the OptimizationId field if non-nil, zero value otherwise.

### GetOptimizationIdOk

`func (o *OptimizationTrialWithObservations) GetOptimizationIdOk() (*string, bool)`

GetOptimizationIdOk returns a tuple with the OptimizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptimizationId

`func (o *OptimizationTrialWithObservations) SetOptimizationId(v string)`

SetOptimizationId sets OptimizationId field to given value.


### GetAlgorithmTrialKey

`func (o *OptimizationTrialWithObservations) GetAlgorithmTrialKey() string`

GetAlgorithmTrialKey returns the AlgorithmTrialKey field if non-nil, zero value otherwise.

### GetAlgorithmTrialKeyOk

`func (o *OptimizationTrialWithObservations) GetAlgorithmTrialKeyOk() (*string, bool)`

GetAlgorithmTrialKeyOk returns a tuple with the AlgorithmTrialKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithmTrialKey

`func (o *OptimizationTrialWithObservations) SetAlgorithmTrialKey(v string)`

SetAlgorithmTrialKey sets AlgorithmTrialKey field to given value.


### GetStatus

`func (o *OptimizationTrialWithObservations) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OptimizationTrialWithObservations) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OptimizationTrialWithObservations) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetParentTrialId

`func (o *OptimizationTrialWithObservations) GetParentTrialId() string`

GetParentTrialId returns the ParentTrialId field if non-nil, zero value otherwise.

### GetParentTrialIdOk

`func (o *OptimizationTrialWithObservations) GetParentTrialIdOk() (*string, bool)`

GetParentTrialIdOk returns a tuple with the ParentTrialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTrialId

`func (o *OptimizationTrialWithObservations) SetParentTrialId(v string)`

SetParentTrialId sets ParentTrialId field to given value.

### HasParentTrialId

`func (o *OptimizationTrialWithObservations) HasParentTrialId() bool`

HasParentTrialId returns a boolean if a field has been set.

### SetParentTrialIdNil

`func (o *OptimizationTrialWithObservations) SetParentTrialIdNil(b bool)`

 SetParentTrialIdNil sets the value for ParentTrialId to be an explicit nil

### UnsetParentTrialId
`func (o *OptimizationTrialWithObservations) UnsetParentTrialId()`

UnsetParentTrialId ensures that no value is present for ParentTrialId, not even an explicit nil
### GetHypothesis

`func (o *OptimizationTrialWithObservations) GetHypothesis() string`

GetHypothesis returns the Hypothesis field if non-nil, zero value otherwise.

### GetHypothesisOk

`func (o *OptimizationTrialWithObservations) GetHypothesisOk() (*string, bool)`

GetHypothesisOk returns a tuple with the Hypothesis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypothesis

`func (o *OptimizationTrialWithObservations) SetHypothesis(v string)`

SetHypothesis sets Hypothesis field to given value.

### HasHypothesis

`func (o *OptimizationTrialWithObservations) HasHypothesis() bool`

HasHypothesis returns a boolean if a field has been set.

### SetHypothesisNil

`func (o *OptimizationTrialWithObservations) SetHypothesisNil(b bool)`

 SetHypothesisNil sets the value for Hypothesis to be an explicit nil

### UnsetHypothesis
`func (o *OptimizationTrialWithObservations) UnsetHypothesis()`

UnsetHypothesis ensures that no value is present for Hypothesis, not even an explicit nil
### GetResolution

`func (o *OptimizationTrialWithObservations) GetResolution() string`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *OptimizationTrialWithObservations) GetResolutionOk() (*string, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *OptimizationTrialWithObservations) SetResolution(v string)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *OptimizationTrialWithObservations) HasResolution() bool`

HasResolution returns a boolean if a field has been set.

### SetResolutionNil

`func (o *OptimizationTrialWithObservations) SetResolutionNil(b bool)`

 SetResolutionNil sets the value for Resolution to be an explicit nil

### UnsetResolution
`func (o *OptimizationTrialWithObservations) UnsetResolution()`

UnsetResolution ensures that no value is present for Resolution, not even an explicit nil
### GetMetadata

`func (o *OptimizationTrialWithObservations) GetMetadata() map[string]*interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *OptimizationTrialWithObservations) GetMetadataOk() (*map[string]*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *OptimizationTrialWithObservations) SetMetadata(v map[string]*interface{})`

SetMetadata sets Metadata field to given value.


### GetObservations

`func (o *OptimizationTrialWithObservations) GetObservations() []OptimizationTrialObservation`

GetObservations returns the Observations field if non-nil, zero value otherwise.

### GetObservationsOk

`func (o *OptimizationTrialWithObservations) GetObservationsOk() (*[]OptimizationTrialObservation, bool)`

GetObservationsOk returns a tuple with the Observations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservations

`func (o *OptimizationTrialWithObservations) SetObservations(v []OptimizationTrialObservation)`

SetObservations sets Observations field to given value.

### HasObservations

`func (o *OptimizationTrialWithObservations) HasObservations() bool`

HasObservations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


