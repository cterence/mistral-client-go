# OptimizationDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**DeletedAt** | **NullableTime** |  | 
**CreatorId** | **string** |  | 
**ProjectId** | **string** |  | 
**EvaluationId** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Steer** | Pointer to **NullableString** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] [default to {}]
**Algorithm** | **string** |  | 
**Status** | **string** |  | 
**Source** | **string** |  | 
**WorkflowExecutionId** | Pointer to **NullableString** |  | [optional] 
**WorkflowRunId** | Pointer to **NullableString** |  | [optional] 
**ObjectiveSpec** | Pointer to **map[string]interface{}** |  | [optional] 
**Outcome** | Pointer to **map[string]interface{}** |  | [optional] 
**Evaluation** | Pointer to [**NullableOptimizationReference**](OptimizationReference.md) |  | [optional] 
**Project** | Pointer to [**NullableOptimizationReference**](OptimizationReference.md) |  | [optional] 

## Methods

### NewOptimizationDetail

`func NewOptimizationDetail(id string, createdAt time.Time, updatedAt time.Time, deletedAt NullableTime, creatorId string, projectId string, name string, algorithm string, status string, source string, ) *OptimizationDetail`

NewOptimizationDetail instantiates a new OptimizationDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptimizationDetailWithDefaults

`func NewOptimizationDetailWithDefaults() *OptimizationDetail`

NewOptimizationDetailWithDefaults instantiates a new OptimizationDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OptimizationDetail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OptimizationDetail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OptimizationDetail) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *OptimizationDetail) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OptimizationDetail) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OptimizationDetail) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *OptimizationDetail) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OptimizationDetail) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OptimizationDetail) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *OptimizationDetail) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *OptimizationDetail) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *OptimizationDetail) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.


### SetDeletedAtNil

`func (o *OptimizationDetail) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *OptimizationDetail) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetCreatorId

`func (o *OptimizationDetail) GetCreatorId() string`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *OptimizationDetail) GetCreatorIdOk() (*string, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *OptimizationDetail) SetCreatorId(v string)`

SetCreatorId sets CreatorId field to given value.


### GetProjectId

`func (o *OptimizationDetail) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *OptimizationDetail) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *OptimizationDetail) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetEvaluationId

`func (o *OptimizationDetail) GetEvaluationId() string`

GetEvaluationId returns the EvaluationId field if non-nil, zero value otherwise.

### GetEvaluationIdOk

`func (o *OptimizationDetail) GetEvaluationIdOk() (*string, bool)`

GetEvaluationIdOk returns a tuple with the EvaluationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationId

`func (o *OptimizationDetail) SetEvaluationId(v string)`

SetEvaluationId sets EvaluationId field to given value.

### HasEvaluationId

`func (o *OptimizationDetail) HasEvaluationId() bool`

HasEvaluationId returns a boolean if a field has been set.

### SetEvaluationIdNil

`func (o *OptimizationDetail) SetEvaluationIdNil(b bool)`

 SetEvaluationIdNil sets the value for EvaluationId to be an explicit nil

### UnsetEvaluationId
`func (o *OptimizationDetail) UnsetEvaluationId()`

UnsetEvaluationId ensures that no value is present for EvaluationId, not even an explicit nil
### GetName

`func (o *OptimizationDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OptimizationDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OptimizationDetail) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *OptimizationDetail) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OptimizationDetail) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OptimizationDetail) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OptimizationDetail) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *OptimizationDetail) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *OptimizationDetail) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSteer

`func (o *OptimizationDetail) GetSteer() string`

GetSteer returns the Steer field if non-nil, zero value otherwise.

### GetSteerOk

`func (o *OptimizationDetail) GetSteerOk() (*string, bool)`

GetSteerOk returns a tuple with the Steer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteer

`func (o *OptimizationDetail) SetSteer(v string)`

SetSteer sets Steer field to given value.

### HasSteer

`func (o *OptimizationDetail) HasSteer() bool`

HasSteer returns a boolean if a field has been set.

### SetSteerNil

`func (o *OptimizationDetail) SetSteerNil(b bool)`

 SetSteerNil sets the value for Steer to be an explicit nil

### UnsetSteer
`func (o *OptimizationDetail) UnsetSteer()`

UnsetSteer ensures that no value is present for Steer, not even an explicit nil
### GetTags

`func (o *OptimizationDetail) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *OptimizationDetail) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *OptimizationDetail) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *OptimizationDetail) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetAlgorithm

`func (o *OptimizationDetail) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *OptimizationDetail) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *OptimizationDetail) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.


### GetStatus

`func (o *OptimizationDetail) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OptimizationDetail) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OptimizationDetail) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSource

`func (o *OptimizationDetail) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *OptimizationDetail) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *OptimizationDetail) SetSource(v string)`

SetSource sets Source field to given value.


### GetWorkflowExecutionId

`func (o *OptimizationDetail) GetWorkflowExecutionId() string`

GetWorkflowExecutionId returns the WorkflowExecutionId field if non-nil, zero value otherwise.

### GetWorkflowExecutionIdOk

`func (o *OptimizationDetail) GetWorkflowExecutionIdOk() (*string, bool)`

GetWorkflowExecutionIdOk returns a tuple with the WorkflowExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowExecutionId

`func (o *OptimizationDetail) SetWorkflowExecutionId(v string)`

SetWorkflowExecutionId sets WorkflowExecutionId field to given value.

### HasWorkflowExecutionId

`func (o *OptimizationDetail) HasWorkflowExecutionId() bool`

HasWorkflowExecutionId returns a boolean if a field has been set.

### SetWorkflowExecutionIdNil

`func (o *OptimizationDetail) SetWorkflowExecutionIdNil(b bool)`

 SetWorkflowExecutionIdNil sets the value for WorkflowExecutionId to be an explicit nil

### UnsetWorkflowExecutionId
`func (o *OptimizationDetail) UnsetWorkflowExecutionId()`

UnsetWorkflowExecutionId ensures that no value is present for WorkflowExecutionId, not even an explicit nil
### GetWorkflowRunId

`func (o *OptimizationDetail) GetWorkflowRunId() string`

GetWorkflowRunId returns the WorkflowRunId field if non-nil, zero value otherwise.

### GetWorkflowRunIdOk

`func (o *OptimizationDetail) GetWorkflowRunIdOk() (*string, bool)`

GetWorkflowRunIdOk returns a tuple with the WorkflowRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowRunId

`func (o *OptimizationDetail) SetWorkflowRunId(v string)`

SetWorkflowRunId sets WorkflowRunId field to given value.

### HasWorkflowRunId

`func (o *OptimizationDetail) HasWorkflowRunId() bool`

HasWorkflowRunId returns a boolean if a field has been set.

### SetWorkflowRunIdNil

`func (o *OptimizationDetail) SetWorkflowRunIdNil(b bool)`

 SetWorkflowRunIdNil sets the value for WorkflowRunId to be an explicit nil

### UnsetWorkflowRunId
`func (o *OptimizationDetail) UnsetWorkflowRunId()`

UnsetWorkflowRunId ensures that no value is present for WorkflowRunId, not even an explicit nil
### GetObjectiveSpec

`func (o *OptimizationDetail) GetObjectiveSpec() map[string]*interface{}`

GetObjectiveSpec returns the ObjectiveSpec field if non-nil, zero value otherwise.

### GetObjectiveSpecOk

`func (o *OptimizationDetail) GetObjectiveSpecOk() (*map[string]*interface{}, bool)`

GetObjectiveSpecOk returns a tuple with the ObjectiveSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectiveSpec

`func (o *OptimizationDetail) SetObjectiveSpec(v map[string]*interface{})`

SetObjectiveSpec sets ObjectiveSpec field to given value.

### HasObjectiveSpec

`func (o *OptimizationDetail) HasObjectiveSpec() bool`

HasObjectiveSpec returns a boolean if a field has been set.

### SetObjectiveSpecNil

`func (o *OptimizationDetail) SetObjectiveSpecNil(b bool)`

 SetObjectiveSpecNil sets the value for ObjectiveSpec to be an explicit nil

### UnsetObjectiveSpec
`func (o *OptimizationDetail) UnsetObjectiveSpec()`

UnsetObjectiveSpec ensures that no value is present for ObjectiveSpec, not even an explicit nil
### GetOutcome

`func (o *OptimizationDetail) GetOutcome() map[string]*interface{}`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *OptimizationDetail) GetOutcomeOk() (*map[string]*interface{}, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *OptimizationDetail) SetOutcome(v map[string]*interface{})`

SetOutcome sets Outcome field to given value.

### HasOutcome

`func (o *OptimizationDetail) HasOutcome() bool`

HasOutcome returns a boolean if a field has been set.

### SetOutcomeNil

`func (o *OptimizationDetail) SetOutcomeNil(b bool)`

 SetOutcomeNil sets the value for Outcome to be an explicit nil

### UnsetOutcome
`func (o *OptimizationDetail) UnsetOutcome()`

UnsetOutcome ensures that no value is present for Outcome, not even an explicit nil
### GetEvaluation

`func (o *OptimizationDetail) GetEvaluation() OptimizationReference`

GetEvaluation returns the Evaluation field if non-nil, zero value otherwise.

### GetEvaluationOk

`func (o *OptimizationDetail) GetEvaluationOk() (*OptimizationReference, bool)`

GetEvaluationOk returns a tuple with the Evaluation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluation

`func (o *OptimizationDetail) SetEvaluation(v OptimizationReference)`

SetEvaluation sets Evaluation field to given value.

### HasEvaluation

`func (o *OptimizationDetail) HasEvaluation() bool`

HasEvaluation returns a boolean if a field has been set.

### SetEvaluationNil

`func (o *OptimizationDetail) SetEvaluationNil(b bool)`

 SetEvaluationNil sets the value for Evaluation to be an explicit nil

### UnsetEvaluation
`func (o *OptimizationDetail) UnsetEvaluation()`

UnsetEvaluation ensures that no value is present for Evaluation, not even an explicit nil
### GetProject

`func (o *OptimizationDetail) GetProject() OptimizationReference`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *OptimizationDetail) GetProjectOk() (*OptimizationReference, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *OptimizationDetail) SetProject(v OptimizationReference)`

SetProject sets Project field to given value.

### HasProject

`func (o *OptimizationDetail) HasProject() bool`

HasProject returns a boolean if a field has been set.

### SetProjectNil

`func (o *OptimizationDetail) SetProjectNil(b bool)`

 SetProjectNil sets the value for Project to be an explicit nil

### UnsetProject
`func (o *OptimizationDetail) UnsetProject()`

UnsetProject ensures that no value is present for Project, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


