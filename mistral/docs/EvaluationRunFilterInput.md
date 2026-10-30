# EvaluationRunFilterInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RunIds** | Pointer to **[]string** |  | [optional] 
**EvaluationIds** | Pointer to **[]string** |  | [optional] 
**CreatorIds** | Pointer to **[]string** |  | [optional] 
**ProjectIds** | Pointer to **[]string** |  | [optional] 
**Ownership** | Pointer to [**NullableRunOwnershipFilter**](RunOwnershipFilter.md) |  | [optional] 
**Tags** | Pointer to [**NullableTagFilter**](TagFilter.md) |  | [optional] 
**ScoreFilters** | Pointer to [**NullableScoreFilterGroup**](ScoreFilterGroup.md) |  | [optional] 
**GoalFilters** | Pointer to [**NullableGoalStatusFilterGroup**](GoalStatusFilterGroup.md) |  | [optional] 
**System** | Pointer to [**NullableSystemFilter**](SystemFilter.md) |  | [optional] 
**Metadata** | Pointer to [**NullableMetadataFilterGroup**](MetadataFilterGroup.md) |  | [optional] 

## Methods

### NewEvaluationRunFilterInput

`func NewEvaluationRunFilterInput() *EvaluationRunFilterInput`

NewEvaluationRunFilterInput instantiates a new EvaluationRunFilterInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvaluationRunFilterInputWithDefaults

`func NewEvaluationRunFilterInputWithDefaults() *EvaluationRunFilterInput`

NewEvaluationRunFilterInputWithDefaults instantiates a new EvaluationRunFilterInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRunIds

`func (o *EvaluationRunFilterInput) GetRunIds() []string`

GetRunIds returns the RunIds field if non-nil, zero value otherwise.

### GetRunIdsOk

`func (o *EvaluationRunFilterInput) GetRunIdsOk() (*[]string, bool)`

GetRunIdsOk returns a tuple with the RunIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunIds

`func (o *EvaluationRunFilterInput) SetRunIds(v []string)`

SetRunIds sets RunIds field to given value.

### HasRunIds

`func (o *EvaluationRunFilterInput) HasRunIds() bool`

HasRunIds returns a boolean if a field has been set.

### SetRunIdsNil

`func (o *EvaluationRunFilterInput) SetRunIdsNil(b bool)`

 SetRunIdsNil sets the value for RunIds to be an explicit nil

### UnsetRunIds
`func (o *EvaluationRunFilterInput) UnsetRunIds()`

UnsetRunIds ensures that no value is present for RunIds, not even an explicit nil
### GetEvaluationIds

`func (o *EvaluationRunFilterInput) GetEvaluationIds() []string`

GetEvaluationIds returns the EvaluationIds field if non-nil, zero value otherwise.

### GetEvaluationIdsOk

`func (o *EvaluationRunFilterInput) GetEvaluationIdsOk() (*[]string, bool)`

GetEvaluationIdsOk returns a tuple with the EvaluationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationIds

`func (o *EvaluationRunFilterInput) SetEvaluationIds(v []string)`

SetEvaluationIds sets EvaluationIds field to given value.

### HasEvaluationIds

`func (o *EvaluationRunFilterInput) HasEvaluationIds() bool`

HasEvaluationIds returns a boolean if a field has been set.

### SetEvaluationIdsNil

`func (o *EvaluationRunFilterInput) SetEvaluationIdsNil(b bool)`

 SetEvaluationIdsNil sets the value for EvaluationIds to be an explicit nil

### UnsetEvaluationIds
`func (o *EvaluationRunFilterInput) UnsetEvaluationIds()`

UnsetEvaluationIds ensures that no value is present for EvaluationIds, not even an explicit nil
### GetCreatorIds

`func (o *EvaluationRunFilterInput) GetCreatorIds() []string`

GetCreatorIds returns the CreatorIds field if non-nil, zero value otherwise.

### GetCreatorIdsOk

`func (o *EvaluationRunFilterInput) GetCreatorIdsOk() (*[]string, bool)`

GetCreatorIdsOk returns a tuple with the CreatorIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorIds

`func (o *EvaluationRunFilterInput) SetCreatorIds(v []string)`

SetCreatorIds sets CreatorIds field to given value.

### HasCreatorIds

`func (o *EvaluationRunFilterInput) HasCreatorIds() bool`

HasCreatorIds returns a boolean if a field has been set.

### SetCreatorIdsNil

`func (o *EvaluationRunFilterInput) SetCreatorIdsNil(b bool)`

 SetCreatorIdsNil sets the value for CreatorIds to be an explicit nil

### UnsetCreatorIds
`func (o *EvaluationRunFilterInput) UnsetCreatorIds()`

UnsetCreatorIds ensures that no value is present for CreatorIds, not even an explicit nil
### GetProjectIds

`func (o *EvaluationRunFilterInput) GetProjectIds() []string`

GetProjectIds returns the ProjectIds field if non-nil, zero value otherwise.

### GetProjectIdsOk

`func (o *EvaluationRunFilterInput) GetProjectIdsOk() (*[]string, bool)`

GetProjectIdsOk returns a tuple with the ProjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIds

`func (o *EvaluationRunFilterInput) SetProjectIds(v []string)`

SetProjectIds sets ProjectIds field to given value.

### HasProjectIds

`func (o *EvaluationRunFilterInput) HasProjectIds() bool`

HasProjectIds returns a boolean if a field has been set.

### SetProjectIdsNil

`func (o *EvaluationRunFilterInput) SetProjectIdsNil(b bool)`

 SetProjectIdsNil sets the value for ProjectIds to be an explicit nil

### UnsetProjectIds
`func (o *EvaluationRunFilterInput) UnsetProjectIds()`

UnsetProjectIds ensures that no value is present for ProjectIds, not even an explicit nil
### GetOwnership

`func (o *EvaluationRunFilterInput) GetOwnership() RunOwnershipFilter`

GetOwnership returns the Ownership field if non-nil, zero value otherwise.

### GetOwnershipOk

`func (o *EvaluationRunFilterInput) GetOwnershipOk() (*RunOwnershipFilter, bool)`

GetOwnershipOk returns a tuple with the Ownership field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnership

`func (o *EvaluationRunFilterInput) SetOwnership(v RunOwnershipFilter)`

SetOwnership sets Ownership field to given value.

### HasOwnership

`func (o *EvaluationRunFilterInput) HasOwnership() bool`

HasOwnership returns a boolean if a field has been set.

### SetOwnershipNil

`func (o *EvaluationRunFilterInput) SetOwnershipNil(b bool)`

 SetOwnershipNil sets the value for Ownership to be an explicit nil

### UnsetOwnership
`func (o *EvaluationRunFilterInput) UnsetOwnership()`

UnsetOwnership ensures that no value is present for Ownership, not even an explicit nil
### GetTags

`func (o *EvaluationRunFilterInput) GetTags() TagFilter`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *EvaluationRunFilterInput) GetTagsOk() (*TagFilter, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *EvaluationRunFilterInput) SetTags(v TagFilter)`

SetTags sets Tags field to given value.

### HasTags

`func (o *EvaluationRunFilterInput) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *EvaluationRunFilterInput) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *EvaluationRunFilterInput) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetScoreFilters

`func (o *EvaluationRunFilterInput) GetScoreFilters() ScoreFilterGroup`

GetScoreFilters returns the ScoreFilters field if non-nil, zero value otherwise.

### GetScoreFiltersOk

`func (o *EvaluationRunFilterInput) GetScoreFiltersOk() (*ScoreFilterGroup, bool)`

GetScoreFiltersOk returns a tuple with the ScoreFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreFilters

`func (o *EvaluationRunFilterInput) SetScoreFilters(v ScoreFilterGroup)`

SetScoreFilters sets ScoreFilters field to given value.

### HasScoreFilters

`func (o *EvaluationRunFilterInput) HasScoreFilters() bool`

HasScoreFilters returns a boolean if a field has been set.

### SetScoreFiltersNil

`func (o *EvaluationRunFilterInput) SetScoreFiltersNil(b bool)`

 SetScoreFiltersNil sets the value for ScoreFilters to be an explicit nil

### UnsetScoreFilters
`func (o *EvaluationRunFilterInput) UnsetScoreFilters()`

UnsetScoreFilters ensures that no value is present for ScoreFilters, not even an explicit nil
### GetGoalFilters

`func (o *EvaluationRunFilterInput) GetGoalFilters() GoalStatusFilterGroup`

GetGoalFilters returns the GoalFilters field if non-nil, zero value otherwise.

### GetGoalFiltersOk

`func (o *EvaluationRunFilterInput) GetGoalFiltersOk() (*GoalStatusFilterGroup, bool)`

GetGoalFiltersOk returns a tuple with the GoalFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoalFilters

`func (o *EvaluationRunFilterInput) SetGoalFilters(v GoalStatusFilterGroup)`

SetGoalFilters sets GoalFilters field to given value.

### HasGoalFilters

`func (o *EvaluationRunFilterInput) HasGoalFilters() bool`

HasGoalFilters returns a boolean if a field has been set.

### SetGoalFiltersNil

`func (o *EvaluationRunFilterInput) SetGoalFiltersNil(b bool)`

 SetGoalFiltersNil sets the value for GoalFilters to be an explicit nil

### UnsetGoalFilters
`func (o *EvaluationRunFilterInput) UnsetGoalFilters()`

UnsetGoalFilters ensures that no value is present for GoalFilters, not even an explicit nil
### GetSystem

`func (o *EvaluationRunFilterInput) GetSystem() SystemFilter`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *EvaluationRunFilterInput) GetSystemOk() (*SystemFilter, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *EvaluationRunFilterInput) SetSystem(v SystemFilter)`

SetSystem sets System field to given value.

### HasSystem

`func (o *EvaluationRunFilterInput) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### SetSystemNil

`func (o *EvaluationRunFilterInput) SetSystemNil(b bool)`

 SetSystemNil sets the value for System to be an explicit nil

### UnsetSystem
`func (o *EvaluationRunFilterInput) UnsetSystem()`

UnsetSystem ensures that no value is present for System, not even an explicit nil
### GetMetadata

`func (o *EvaluationRunFilterInput) GetMetadata() MetadataFilterGroup`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *EvaluationRunFilterInput) GetMetadataOk() (*MetadataFilterGroup, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *EvaluationRunFilterInput) SetMetadata(v MetadataFilterGroup)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *EvaluationRunFilterInput) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *EvaluationRunFilterInput) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *EvaluationRunFilterInput) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


