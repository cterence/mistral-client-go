# EvaluationRunV2WithRefsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**DeletedAt** | **NullableTime** |  | 
**CreatorId** | **string** |  | 
**EvaluationId** | **string** |  | 
**Name** | **string** |  | 
**Description** | **NullableString** |  | 
**Tags** | **[]string** |  | 
**Metadata** | **map[string]interface{}** |  | 
**NumGenerations** | **int32** |  | 
**RunInfo** | Pointer to **map[string]interface{}** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] [default to "completed"]
**DatasetId** | Pointer to **NullableString** |  | [optional] 
**Evaluators** | [**[]EvaluatorV2**](EvaluatorV2.md) |  | 
**RunEvaluators** | Pointer to [**[]EvaluatorV2**](EvaluatorV2.md) |  | [optional] [default to {}]
**System** | Pointer to [**NullableSystemV2**](SystemV2.md) |  | [optional] 
**Statistics** | Pointer to [**map[string]EvaluatorStatistics**](EvaluatorStatistics.md) |  | [optional] [default to {}]
**GoalResults** | Pointer to [**map[string]GoalResult**](GoalResult.md) |  | [optional] [default to {}]
**Passed** | Pointer to **NullableBool** |  | [optional] 
**RunScores** | Pointer to [**map[string]EvaluationRunScoreV2**](EvaluationRunScoreV2.md) |  | [optional] [default to {}]
**Optimization** | Pointer to [**NullableRunOptimizationRef**](RunOptimizationRef.md) |  | [optional] 
**Evaluation** | Pointer to [**NullableRunEvaluationSummary**](RunEvaluationSummary.md) |  | [optional] 
**Project** | Pointer to [**NullableRunProjectSummary**](RunProjectSummary.md) |  | [optional] 

## Methods

### NewEvaluationRunV2WithRefsResponse

`func NewEvaluationRunV2WithRefsResponse(id string, createdAt time.Time, updatedAt time.Time, deletedAt NullableTime, creatorId string, evaluationId string, name string, description NullableString, tags []string, metadata map[string]*interface{}, numGenerations int32, evaluators []EvaluatorV2, ) *EvaluationRunV2WithRefsResponse`

NewEvaluationRunV2WithRefsResponse instantiates a new EvaluationRunV2WithRefsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvaluationRunV2WithRefsResponseWithDefaults

`func NewEvaluationRunV2WithRefsResponseWithDefaults() *EvaluationRunV2WithRefsResponse`

NewEvaluationRunV2WithRefsResponseWithDefaults instantiates a new EvaluationRunV2WithRefsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EvaluationRunV2WithRefsResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EvaluationRunV2WithRefsResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EvaluationRunV2WithRefsResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *EvaluationRunV2WithRefsResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EvaluationRunV2WithRefsResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EvaluationRunV2WithRefsResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *EvaluationRunV2WithRefsResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EvaluationRunV2WithRefsResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EvaluationRunV2WithRefsResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *EvaluationRunV2WithRefsResponse) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *EvaluationRunV2WithRefsResponse) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *EvaluationRunV2WithRefsResponse) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.


### SetDeletedAtNil

`func (o *EvaluationRunV2WithRefsResponse) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *EvaluationRunV2WithRefsResponse) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetCreatorId

`func (o *EvaluationRunV2WithRefsResponse) GetCreatorId() string`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *EvaluationRunV2WithRefsResponse) GetCreatorIdOk() (*string, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *EvaluationRunV2WithRefsResponse) SetCreatorId(v string)`

SetCreatorId sets CreatorId field to given value.


### GetEvaluationId

`func (o *EvaluationRunV2WithRefsResponse) GetEvaluationId() string`

GetEvaluationId returns the EvaluationId field if non-nil, zero value otherwise.

### GetEvaluationIdOk

`func (o *EvaluationRunV2WithRefsResponse) GetEvaluationIdOk() (*string, bool)`

GetEvaluationIdOk returns a tuple with the EvaluationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationId

`func (o *EvaluationRunV2WithRefsResponse) SetEvaluationId(v string)`

SetEvaluationId sets EvaluationId field to given value.


### GetName

`func (o *EvaluationRunV2WithRefsResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EvaluationRunV2WithRefsResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EvaluationRunV2WithRefsResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *EvaluationRunV2WithRefsResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EvaluationRunV2WithRefsResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EvaluationRunV2WithRefsResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *EvaluationRunV2WithRefsResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *EvaluationRunV2WithRefsResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetTags

`func (o *EvaluationRunV2WithRefsResponse) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *EvaluationRunV2WithRefsResponse) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *EvaluationRunV2WithRefsResponse) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetMetadata

`func (o *EvaluationRunV2WithRefsResponse) GetMetadata() map[string]*interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *EvaluationRunV2WithRefsResponse) GetMetadataOk() (*map[string]*interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *EvaluationRunV2WithRefsResponse) SetMetadata(v map[string]*interface{})`

SetMetadata sets Metadata field to given value.


### GetNumGenerations

`func (o *EvaluationRunV2WithRefsResponse) GetNumGenerations() int32`

GetNumGenerations returns the NumGenerations field if non-nil, zero value otherwise.

### GetNumGenerationsOk

`func (o *EvaluationRunV2WithRefsResponse) GetNumGenerationsOk() (*int32, bool)`

GetNumGenerationsOk returns a tuple with the NumGenerations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumGenerations

`func (o *EvaluationRunV2WithRefsResponse) SetNumGenerations(v int32)`

SetNumGenerations sets NumGenerations field to given value.


### GetRunInfo

`func (o *EvaluationRunV2WithRefsResponse) GetRunInfo() map[string]*interface{}`

GetRunInfo returns the RunInfo field if non-nil, zero value otherwise.

### GetRunInfoOk

`func (o *EvaluationRunV2WithRefsResponse) GetRunInfoOk() (*map[string]*interface{}, bool)`

GetRunInfoOk returns a tuple with the RunInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunInfo

`func (o *EvaluationRunV2WithRefsResponse) SetRunInfo(v map[string]*interface{})`

SetRunInfo sets RunInfo field to given value.

### HasRunInfo

`func (o *EvaluationRunV2WithRefsResponse) HasRunInfo() bool`

HasRunInfo returns a boolean if a field has been set.

### SetRunInfoNil

`func (o *EvaluationRunV2WithRefsResponse) SetRunInfoNil(b bool)`

 SetRunInfoNil sets the value for RunInfo to be an explicit nil

### UnsetRunInfo
`func (o *EvaluationRunV2WithRefsResponse) UnsetRunInfo()`

UnsetRunInfo ensures that no value is present for RunInfo, not even an explicit nil
### GetStatus

`func (o *EvaluationRunV2WithRefsResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EvaluationRunV2WithRefsResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EvaluationRunV2WithRefsResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EvaluationRunV2WithRefsResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDatasetId

`func (o *EvaluationRunV2WithRefsResponse) GetDatasetId() string`

GetDatasetId returns the DatasetId field if non-nil, zero value otherwise.

### GetDatasetIdOk

`func (o *EvaluationRunV2WithRefsResponse) GetDatasetIdOk() (*string, bool)`

GetDatasetIdOk returns a tuple with the DatasetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasetId

`func (o *EvaluationRunV2WithRefsResponse) SetDatasetId(v string)`

SetDatasetId sets DatasetId field to given value.

### HasDatasetId

`func (o *EvaluationRunV2WithRefsResponse) HasDatasetId() bool`

HasDatasetId returns a boolean if a field has been set.

### SetDatasetIdNil

`func (o *EvaluationRunV2WithRefsResponse) SetDatasetIdNil(b bool)`

 SetDatasetIdNil sets the value for DatasetId to be an explicit nil

### UnsetDatasetId
`func (o *EvaluationRunV2WithRefsResponse) UnsetDatasetId()`

UnsetDatasetId ensures that no value is present for DatasetId, not even an explicit nil
### GetEvaluators

`func (o *EvaluationRunV2WithRefsResponse) GetEvaluators() []EvaluatorV2`

GetEvaluators returns the Evaluators field if non-nil, zero value otherwise.

### GetEvaluatorsOk

`func (o *EvaluationRunV2WithRefsResponse) GetEvaluatorsOk() (*[]EvaluatorV2, bool)`

GetEvaluatorsOk returns a tuple with the Evaluators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluators

`func (o *EvaluationRunV2WithRefsResponse) SetEvaluators(v []EvaluatorV2)`

SetEvaluators sets Evaluators field to given value.


### GetRunEvaluators

`func (o *EvaluationRunV2WithRefsResponse) GetRunEvaluators() []EvaluatorV2`

GetRunEvaluators returns the RunEvaluators field if non-nil, zero value otherwise.

### GetRunEvaluatorsOk

`func (o *EvaluationRunV2WithRefsResponse) GetRunEvaluatorsOk() (*[]EvaluatorV2, bool)`

GetRunEvaluatorsOk returns a tuple with the RunEvaluators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunEvaluators

`func (o *EvaluationRunV2WithRefsResponse) SetRunEvaluators(v []EvaluatorV2)`

SetRunEvaluators sets RunEvaluators field to given value.

### HasRunEvaluators

`func (o *EvaluationRunV2WithRefsResponse) HasRunEvaluators() bool`

HasRunEvaluators returns a boolean if a field has been set.

### GetSystem

`func (o *EvaluationRunV2WithRefsResponse) GetSystem() SystemV2`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *EvaluationRunV2WithRefsResponse) GetSystemOk() (*SystemV2, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *EvaluationRunV2WithRefsResponse) SetSystem(v SystemV2)`

SetSystem sets System field to given value.

### HasSystem

`func (o *EvaluationRunV2WithRefsResponse) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### SetSystemNil

`func (o *EvaluationRunV2WithRefsResponse) SetSystemNil(b bool)`

 SetSystemNil sets the value for System to be an explicit nil

### UnsetSystem
`func (o *EvaluationRunV2WithRefsResponse) UnsetSystem()`

UnsetSystem ensures that no value is present for System, not even an explicit nil
### GetStatistics

`func (o *EvaluationRunV2WithRefsResponse) GetStatistics() map[string]EvaluatorStatistics`

GetStatistics returns the Statistics field if non-nil, zero value otherwise.

### GetStatisticsOk

`func (o *EvaluationRunV2WithRefsResponse) GetStatisticsOk() (*map[string]EvaluatorStatistics, bool)`

GetStatisticsOk returns a tuple with the Statistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatistics

`func (o *EvaluationRunV2WithRefsResponse) SetStatistics(v map[string]EvaluatorStatistics)`

SetStatistics sets Statistics field to given value.

### HasStatistics

`func (o *EvaluationRunV2WithRefsResponse) HasStatistics() bool`

HasStatistics returns a boolean if a field has been set.

### GetGoalResults

`func (o *EvaluationRunV2WithRefsResponse) GetGoalResults() map[string]GoalResult`

GetGoalResults returns the GoalResults field if non-nil, zero value otherwise.

### GetGoalResultsOk

`func (o *EvaluationRunV2WithRefsResponse) GetGoalResultsOk() (*map[string]GoalResult, bool)`

GetGoalResultsOk returns a tuple with the GoalResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoalResults

`func (o *EvaluationRunV2WithRefsResponse) SetGoalResults(v map[string]GoalResult)`

SetGoalResults sets GoalResults field to given value.

### HasGoalResults

`func (o *EvaluationRunV2WithRefsResponse) HasGoalResults() bool`

HasGoalResults returns a boolean if a field has been set.

### GetPassed

`func (o *EvaluationRunV2WithRefsResponse) GetPassed() bool`

GetPassed returns the Passed field if non-nil, zero value otherwise.

### GetPassedOk

`func (o *EvaluationRunV2WithRefsResponse) GetPassedOk() (*bool, bool)`

GetPassedOk returns a tuple with the Passed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassed

`func (o *EvaluationRunV2WithRefsResponse) SetPassed(v bool)`

SetPassed sets Passed field to given value.

### HasPassed

`func (o *EvaluationRunV2WithRefsResponse) HasPassed() bool`

HasPassed returns a boolean if a field has been set.

### SetPassedNil

`func (o *EvaluationRunV2WithRefsResponse) SetPassedNil(b bool)`

 SetPassedNil sets the value for Passed to be an explicit nil

### UnsetPassed
`func (o *EvaluationRunV2WithRefsResponse) UnsetPassed()`

UnsetPassed ensures that no value is present for Passed, not even an explicit nil
### GetRunScores

`func (o *EvaluationRunV2WithRefsResponse) GetRunScores() map[string]EvaluationRunScoreV2`

GetRunScores returns the RunScores field if non-nil, zero value otherwise.

### GetRunScoresOk

`func (o *EvaluationRunV2WithRefsResponse) GetRunScoresOk() (*map[string]EvaluationRunScoreV2, bool)`

GetRunScoresOk returns a tuple with the RunScores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunScores

`func (o *EvaluationRunV2WithRefsResponse) SetRunScores(v map[string]EvaluationRunScoreV2)`

SetRunScores sets RunScores field to given value.

### HasRunScores

`func (o *EvaluationRunV2WithRefsResponse) HasRunScores() bool`

HasRunScores returns a boolean if a field has been set.

### GetOptimization

`func (o *EvaluationRunV2WithRefsResponse) GetOptimization() RunOptimizationRef`

GetOptimization returns the Optimization field if non-nil, zero value otherwise.

### GetOptimizationOk

`func (o *EvaluationRunV2WithRefsResponse) GetOptimizationOk() (*RunOptimizationRef, bool)`

GetOptimizationOk returns a tuple with the Optimization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptimization

`func (o *EvaluationRunV2WithRefsResponse) SetOptimization(v RunOptimizationRef)`

SetOptimization sets Optimization field to given value.

### HasOptimization

`func (o *EvaluationRunV2WithRefsResponse) HasOptimization() bool`

HasOptimization returns a boolean if a field has been set.

### SetOptimizationNil

`func (o *EvaluationRunV2WithRefsResponse) SetOptimizationNil(b bool)`

 SetOptimizationNil sets the value for Optimization to be an explicit nil

### UnsetOptimization
`func (o *EvaluationRunV2WithRefsResponse) UnsetOptimization()`

UnsetOptimization ensures that no value is present for Optimization, not even an explicit nil
### GetEvaluation

`func (o *EvaluationRunV2WithRefsResponse) GetEvaluation() RunEvaluationSummary`

GetEvaluation returns the Evaluation field if non-nil, zero value otherwise.

### GetEvaluationOk

`func (o *EvaluationRunV2WithRefsResponse) GetEvaluationOk() (*RunEvaluationSummary, bool)`

GetEvaluationOk returns a tuple with the Evaluation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluation

`func (o *EvaluationRunV2WithRefsResponse) SetEvaluation(v RunEvaluationSummary)`

SetEvaluation sets Evaluation field to given value.

### HasEvaluation

`func (o *EvaluationRunV2WithRefsResponse) HasEvaluation() bool`

HasEvaluation returns a boolean if a field has been set.

### SetEvaluationNil

`func (o *EvaluationRunV2WithRefsResponse) SetEvaluationNil(b bool)`

 SetEvaluationNil sets the value for Evaluation to be an explicit nil

### UnsetEvaluation
`func (o *EvaluationRunV2WithRefsResponse) UnsetEvaluation()`

UnsetEvaluation ensures that no value is present for Evaluation, not even an explicit nil
### GetProject

`func (o *EvaluationRunV2WithRefsResponse) GetProject() RunProjectSummary`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *EvaluationRunV2WithRefsResponse) GetProjectOk() (*RunProjectSummary, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *EvaluationRunV2WithRefsResponse) SetProject(v RunProjectSummary)`

SetProject sets Project field to given value.

### HasProject

`func (o *EvaluationRunV2WithRefsResponse) HasProject() bool`

HasProject returns a boolean if a field has been set.

### SetProjectNil

`func (o *EvaluationRunV2WithRefsResponse) SetProjectNil(b bool)`

 SetProjectNil sets the value for Project to be an explicit nil

### UnsetProject
`func (o *EvaluationRunV2WithRefsResponse) UnsetProject()`

UnsetProject ensures that no value is present for Project, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


