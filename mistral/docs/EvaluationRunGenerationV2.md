# EvaluationRunGenerationV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**DeletedAt** | **NullableTime** |  | 
**Output** | [**NullableAnyOf**](anyOf&lt;&gt;.md) |  | 
**Metadata** | **map[string]interface{}** |  | 
**Status** | **string** |  | 
**Error** | Pointer to **NullableString** |  | [optional] 
**Scores** | Pointer to [**map[string][]EvaluationRunScoreV2**](array.md) |  | [optional] [default to {}]
**Statistics** | Pointer to [**map[string]EvaluatorStatistics**](EvaluatorStatistics.md) |  | [optional] [default to {}]
**GoalResults** | Pointer to [**map[string]GoalResult**](GoalResult.md) |  | [optional] [default to {}]

## Methods

### NewEvaluationRunGenerationV2

`func NewEvaluationRunGenerationV2(id string, createdAt time.Time, updatedAt time.Time, deletedAt NullableTime, output NullableAnyOf, metadata map[string]*interface{}, status string, ) *EvaluationRunGenerationV2`

NewEvaluationRunGenerationV2 instantiates a new EvaluationRunGenerationV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvaluationRunGenerationV2WithDefaults

`func NewEvaluationRunGenerationV2WithDefaults() *EvaluationRunGenerationV2`

NewEvaluationRunGenerationV2WithDefaults instantiates a new EvaluationRunGenerationV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EvaluationRunGenerationV2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EvaluationRunGenerationV2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EvaluationRunGenerationV2) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *EvaluationRunGenerationV2) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EvaluationRunGenerationV2) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EvaluationRunGenerationV2) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *EvaluationRunGenerationV2) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EvaluationRunGenerationV2) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EvaluationRunGenerationV2) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *EvaluationRunGenerationV2) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *EvaluationRunGenerationV2) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *EvaluationRunGenerationV2) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.


### SetDeletedAtNil

`func (o *EvaluationRunGenerationV2) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *EvaluationRunGenerationV2) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetOutput

`func (o *EvaluationRunGenerationV2) GetOutput() AnyOf`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *EvaluationRunGenerationV2) GetOutputOk() (*AnyOf, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *EvaluationRunGenerationV2) SetOutput(v AnyOf)`

SetOutput sets Output field to given value.


### SetOutputNil

`func (o *EvaluationRunGenerationV2) SetOutputNil(b bool)`

 SetOutputNil sets the value for Output to be an explicit nil

### UnsetOutput
`func (o *EvaluationRunGenerationV2) UnsetOutput()`

UnsetOutput ensures that no value is present for Output, not even an explicit nil
### GetMetadata

`func (o *EvaluationRunGenerationV2) GetMetadata() map[string]*interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *EvaluationRunGenerationV2) GetMetadataOk() (*map[string]*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *EvaluationRunGenerationV2) SetMetadata(v map[string]*interface{})`

SetMetadata sets Metadata field to given value.


### GetStatus

`func (o *EvaluationRunGenerationV2) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EvaluationRunGenerationV2) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EvaluationRunGenerationV2) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetError

`func (o *EvaluationRunGenerationV2) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *EvaluationRunGenerationV2) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *EvaluationRunGenerationV2) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *EvaluationRunGenerationV2) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *EvaluationRunGenerationV2) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *EvaluationRunGenerationV2) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetScores

`func (o *EvaluationRunGenerationV2) GetScores() map[string][]EvaluationRunScoreV2`

GetScores returns the Scores field if non-nil, zero value otherwise.

### GetScoresOk

`func (o *EvaluationRunGenerationV2) GetScoresOk() (*map[string][]EvaluationRunScoreV2, bool)`

GetScoresOk returns a tuple with the Scores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScores

`func (o *EvaluationRunGenerationV2) SetScores(v map[string][]EvaluationRunScoreV2)`

SetScores sets Scores field to given value.

### HasScores

`func (o *EvaluationRunGenerationV2) HasScores() bool`

HasScores returns a boolean if a field has been set.

### GetStatistics

`func (o *EvaluationRunGenerationV2) GetStatistics() map[string]EvaluatorStatistics`

GetStatistics returns the Statistics field if non-nil, zero value otherwise.

### GetStatisticsOk

`func (o *EvaluationRunGenerationV2) GetStatisticsOk() (*map[string]EvaluatorStatistics, bool)`

GetStatisticsOk returns a tuple with the Statistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatistics

`func (o *EvaluationRunGenerationV2) SetStatistics(v map[string]EvaluatorStatistics)`

SetStatistics sets Statistics field to given value.

### HasStatistics

`func (o *EvaluationRunGenerationV2) HasStatistics() bool`

HasStatistics returns a boolean if a field has been set.

### GetGoalResults

`func (o *EvaluationRunGenerationV2) GetGoalResults() map[string]GoalResult`

GetGoalResults returns the GoalResults field if non-nil, zero value otherwise.

### GetGoalResultsOk

`func (o *EvaluationRunGenerationV2) GetGoalResultsOk() (*map[string]GoalResult, bool)`

GetGoalResultsOk returns a tuple with the GoalResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoalResults

`func (o *EvaluationRunGenerationV2) SetGoalResults(v map[string]GoalResult)`

SetGoalResults sets GoalResults field to given value.

### HasGoalResults

`func (o *EvaluationRunGenerationV2) HasGoalResults() bool`

HasGoalResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


