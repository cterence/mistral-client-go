# Optimization

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

## Methods

### NewOptimization

`func NewOptimization(id string, createdAt time.Time, updatedAt time.Time, deletedAt NullableTime, creatorId string, projectId string, name string, algorithm string, status string, source string, ) *Optimization`

NewOptimization instantiates a new Optimization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptimizationWithDefaults

`func NewOptimizationWithDefaults() *Optimization`

NewOptimizationWithDefaults instantiates a new Optimization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Optimization) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Optimization) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Optimization) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *Optimization) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Optimization) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Optimization) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Optimization) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Optimization) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Optimization) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *Optimization) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *Optimization) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *Optimization) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.


### SetDeletedAtNil

`func (o *Optimization) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *Optimization) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetCreatorId

`func (o *Optimization) GetCreatorId() string`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *Optimization) GetCreatorIdOk() (*string, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *Optimization) SetCreatorId(v string)`

SetCreatorId sets CreatorId field to given value.


### GetProjectId

`func (o *Optimization) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *Optimization) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *Optimization) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetEvaluationId

`func (o *Optimization) GetEvaluationId() string`

GetEvaluationId returns the EvaluationId field if non-nil, zero value otherwise.

### GetEvaluationIdOk

`func (o *Optimization) GetEvaluationIdOk() (*string, bool)`

GetEvaluationIdOk returns a tuple with the EvaluationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationId

`func (o *Optimization) SetEvaluationId(v string)`

SetEvaluationId sets EvaluationId field to given value.

### HasEvaluationId

`func (o *Optimization) HasEvaluationId() bool`

HasEvaluationId returns a boolean if a field has been set.

### SetEvaluationIdNil

`func (o *Optimization) SetEvaluationIdNil(b bool)`

 SetEvaluationIdNil sets the value for EvaluationId to be an explicit nil

### UnsetEvaluationId
`func (o *Optimization) UnsetEvaluationId()`

UnsetEvaluationId ensures that no value is present for EvaluationId, not even an explicit nil
### GetName

`func (o *Optimization) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Optimization) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Optimization) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Optimization) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Optimization) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Optimization) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Optimization) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *Optimization) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Optimization) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSteer

`func (o *Optimization) GetSteer() string`

GetSteer returns the Steer field if non-nil, zero value otherwise.

### GetSteerOk

`func (o *Optimization) GetSteerOk() (*string, bool)`

GetSteerOk returns a tuple with the Steer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteer

`func (o *Optimization) SetSteer(v string)`

SetSteer sets Steer field to given value.

### HasSteer

`func (o *Optimization) HasSteer() bool`

HasSteer returns a boolean if a field has been set.

### SetSteerNil

`func (o *Optimization) SetSteerNil(b bool)`

 SetSteerNil sets the value for Steer to be an explicit nil

### UnsetSteer
`func (o *Optimization) UnsetSteer()`

UnsetSteer ensures that no value is present for Steer, not even an explicit nil
### GetTags

`func (o *Optimization) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Optimization) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Optimization) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Optimization) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetAlgorithm

`func (o *Optimization) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *Optimization) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *Optimization) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.


### GetStatus

`func (o *Optimization) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Optimization) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Optimization) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSource

`func (o *Optimization) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Optimization) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Optimization) SetSource(v string)`

SetSource sets Source field to given value.


### GetWorkflowExecutionId

`func (o *Optimization) GetWorkflowExecutionId() string`

GetWorkflowExecutionId returns the WorkflowExecutionId field if non-nil, zero value otherwise.

### GetWorkflowExecutionIdOk

`func (o *Optimization) GetWorkflowExecutionIdOk() (*string, bool)`

GetWorkflowExecutionIdOk returns a tuple with the WorkflowExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowExecutionId

`func (o *Optimization) SetWorkflowExecutionId(v string)`

SetWorkflowExecutionId sets WorkflowExecutionId field to given value.

### HasWorkflowExecutionId

`func (o *Optimization) HasWorkflowExecutionId() bool`

HasWorkflowExecutionId returns a boolean if a field has been set.

### SetWorkflowExecutionIdNil

`func (o *Optimization) SetWorkflowExecutionIdNil(b bool)`

 SetWorkflowExecutionIdNil sets the value for WorkflowExecutionId to be an explicit nil

### UnsetWorkflowExecutionId
`func (o *Optimization) UnsetWorkflowExecutionId()`

UnsetWorkflowExecutionId ensures that no value is present for WorkflowExecutionId, not even an explicit nil
### GetWorkflowRunId

`func (o *Optimization) GetWorkflowRunId() string`

GetWorkflowRunId returns the WorkflowRunId field if non-nil, zero value otherwise.

### GetWorkflowRunIdOk

`func (o *Optimization) GetWorkflowRunIdOk() (*string, bool)`

GetWorkflowRunIdOk returns a tuple with the WorkflowRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowRunId

`func (o *Optimization) SetWorkflowRunId(v string)`

SetWorkflowRunId sets WorkflowRunId field to given value.

### HasWorkflowRunId

`func (o *Optimization) HasWorkflowRunId() bool`

HasWorkflowRunId returns a boolean if a field has been set.

### SetWorkflowRunIdNil

`func (o *Optimization) SetWorkflowRunIdNil(b bool)`

 SetWorkflowRunIdNil sets the value for WorkflowRunId to be an explicit nil

### UnsetWorkflowRunId
`func (o *Optimization) UnsetWorkflowRunId()`

UnsetWorkflowRunId ensures that no value is present for WorkflowRunId, not even an explicit nil
### GetObjectiveSpec

`func (o *Optimization) GetObjectiveSpec() map[string]*interface{}`

GetObjectiveSpec returns the ObjectiveSpec field if non-nil, zero value otherwise.

### GetObjectiveSpecOk

`func (o *Optimization) GetObjectiveSpecOk() (*map[string]*interface{}, bool)`

GetObjectiveSpecOk returns a tuple with the ObjectiveSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectiveSpec

`func (o *Optimization) SetObjectiveSpec(v map[string]*interface{})`

SetObjectiveSpec sets ObjectiveSpec field to given value.

### HasObjectiveSpec

`func (o *Optimization) HasObjectiveSpec() bool`

HasObjectiveSpec returns a boolean if a field has been set.

### SetObjectiveSpecNil

`func (o *Optimization) SetObjectiveSpecNil(b bool)`

 SetObjectiveSpecNil sets the value for ObjectiveSpec to be an explicit nil

### UnsetObjectiveSpec
`func (o *Optimization) UnsetObjectiveSpec()`

UnsetObjectiveSpec ensures that no value is present for ObjectiveSpec, not even an explicit nil
### GetOutcome

`func (o *Optimization) GetOutcome() map[string]*interface{}`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *Optimization) GetOutcomeOk() (*map[string]*interface{}, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *Optimization) SetOutcome(v map[string]*interface{})`

SetOutcome sets Outcome field to given value.

### HasOutcome

`func (o *Optimization) HasOutcome() bool`

HasOutcome returns a boolean if a field has been set.

### SetOutcomeNil

`func (o *Optimization) SetOutcomeNil(b bool)`

 SetOutcomeNil sets the value for Outcome to be an explicit nil

### UnsetOutcome
`func (o *Optimization) UnsetOutcome()`

UnsetOutcome ensures that no value is present for Outcome, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


