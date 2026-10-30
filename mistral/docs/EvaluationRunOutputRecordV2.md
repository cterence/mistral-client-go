# EvaluationRunOutputRecordV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**DeletedAt** | **NullableTime** |  | 
**Metadata** | **map[string]interface{}** |  | 
**RecordInfo** | Pointer to **map[string]interface{}** |  | [optional] 
**Generations** | Pointer to [**[]EvaluationRunGenerationV2**](EvaluationRunGenerationV2.md) |  | [optional] [default to {}]
**Statistics** | Pointer to [**map[string]EvaluatorStatistics**](EvaluatorStatistics.md) |  | [optional] [default to {}]
**GoalResults** | Pointer to [**map[string]GoalResult**](GoalResult.md) |  | [optional] [default to {}]

## Methods

### NewEvaluationRunOutputRecordV2

`func NewEvaluationRunOutputRecordV2(id string, createdAt time.Time, updatedAt time.Time, deletedAt NullableTime, metadata map[string]*interface{}, ) *EvaluationRunOutputRecordV2`

NewEvaluationRunOutputRecordV2 instantiates a new EvaluationRunOutputRecordV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvaluationRunOutputRecordV2WithDefaults

`func NewEvaluationRunOutputRecordV2WithDefaults() *EvaluationRunOutputRecordV2`

NewEvaluationRunOutputRecordV2WithDefaults instantiates a new EvaluationRunOutputRecordV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EvaluationRunOutputRecordV2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EvaluationRunOutputRecordV2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EvaluationRunOutputRecordV2) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *EvaluationRunOutputRecordV2) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EvaluationRunOutputRecordV2) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EvaluationRunOutputRecordV2) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *EvaluationRunOutputRecordV2) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EvaluationRunOutputRecordV2) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EvaluationRunOutputRecordV2) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *EvaluationRunOutputRecordV2) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *EvaluationRunOutputRecordV2) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *EvaluationRunOutputRecordV2) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.


### SetDeletedAtNil

`func (o *EvaluationRunOutputRecordV2) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *EvaluationRunOutputRecordV2) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetMetadata

`func (o *EvaluationRunOutputRecordV2) GetMetadata() map[string]*interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *EvaluationRunOutputRecordV2) GetMetadataOk() (*map[string]*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *EvaluationRunOutputRecordV2) SetMetadata(v map[string]*interface{})`

SetMetadata sets Metadata field to given value.


### GetRecordInfo

`func (o *EvaluationRunOutputRecordV2) GetRecordInfo() map[string]*interface{}`

GetRecordInfo returns the RecordInfo field if non-nil, zero value otherwise.

### GetRecordInfoOk

`func (o *EvaluationRunOutputRecordV2) GetRecordInfoOk() (*map[string]*interface{}, bool)`

GetRecordInfoOk returns a tuple with the RecordInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordInfo

`func (o *EvaluationRunOutputRecordV2) SetRecordInfo(v map[string]*interface{})`

SetRecordInfo sets RecordInfo field to given value.

### HasRecordInfo

`func (o *EvaluationRunOutputRecordV2) HasRecordInfo() bool`

HasRecordInfo returns a boolean if a field has been set.

### SetRecordInfoNil

`func (o *EvaluationRunOutputRecordV2) SetRecordInfoNil(b bool)`

 SetRecordInfoNil sets the value for RecordInfo to be an explicit nil

### UnsetRecordInfo
`func (o *EvaluationRunOutputRecordV2) UnsetRecordInfo()`

UnsetRecordInfo ensures that no value is present for RecordInfo, not even an explicit nil
### GetGenerations

`func (o *EvaluationRunOutputRecordV2) GetGenerations() []EvaluationRunGenerationV2`

GetGenerations returns the Generations field if non-nil, zero value otherwise.

### GetGenerationsOk

`func (o *EvaluationRunOutputRecordV2) GetGenerationsOk() (*[]EvaluationRunGenerationV2, bool)`

GetGenerationsOk returns a tuple with the Generations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerations

`func (o *EvaluationRunOutputRecordV2) SetGenerations(v []EvaluationRunGenerationV2)`

SetGenerations sets Generations field to given value.

### HasGenerations

`func (o *EvaluationRunOutputRecordV2) HasGenerations() bool`

HasGenerations returns a boolean if a field has been set.

### GetStatistics

`func (o *EvaluationRunOutputRecordV2) GetStatistics() map[string]EvaluatorStatistics`

GetStatistics returns the Statistics field if non-nil, zero value otherwise.

### GetStatisticsOk

`func (o *EvaluationRunOutputRecordV2) GetStatisticsOk() (*map[string]EvaluatorStatistics, bool)`

GetStatisticsOk returns a tuple with the Statistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatistics

`func (o *EvaluationRunOutputRecordV2) SetStatistics(v map[string]EvaluatorStatistics)`

SetStatistics sets Statistics field to given value.

### HasStatistics

`func (o *EvaluationRunOutputRecordV2) HasStatistics() bool`

HasStatistics returns a boolean if a field has been set.

### GetGoalResults

`func (o *EvaluationRunOutputRecordV2) GetGoalResults() map[string]GoalResult`

GetGoalResults returns the GoalResults field if non-nil, zero value otherwise.

### GetGoalResultsOk

`func (o *EvaluationRunOutputRecordV2) GetGoalResultsOk() (*map[string]GoalResult, bool)`

GetGoalResultsOk returns a tuple with the GoalResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoalResults

`func (o *EvaluationRunOutputRecordV2) SetGoalResults(v map[string]GoalResult)`

SetGoalResults sets GoalResults field to given value.

### HasGoalResults

`func (o *EvaluationRunOutputRecordV2) HasGoalResults() bool`

HasGoalResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


