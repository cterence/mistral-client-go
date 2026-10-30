# EvaluationRunRecordFilterInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InputFilters** | Pointer to [**NullableContentSearchFilterGroup**](ContentSearchFilterGroup.md) |  | [optional] 
**OutputFilters** | Pointer to [**NullableContentSearchFilterGroup**](ContentSearchFilterGroup.md) |  | [optional] 
**ScoreFilters** | Pointer to [**NullableScoreFilterGroup**](ScoreFilterGroup.md) |  | [optional] 
**GoalFilters** | Pointer to [**NullableGoalStatusFilterGroup**](GoalStatusFilterGroup.md) |  | [optional] 
**Sort** | Pointer to [**NullableEvaluationRunRecordFilterInputSort**](EvaluationRunRecordFilterInputSort.md) |  | [optional] 

## Methods

### NewEvaluationRunRecordFilterInput

`func NewEvaluationRunRecordFilterInput() *EvaluationRunRecordFilterInput`

NewEvaluationRunRecordFilterInput instantiates a new EvaluationRunRecordFilterInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvaluationRunRecordFilterInputWithDefaults

`func NewEvaluationRunRecordFilterInputWithDefaults() *EvaluationRunRecordFilterInput`

NewEvaluationRunRecordFilterInputWithDefaults instantiates a new EvaluationRunRecordFilterInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInputFilters

`func (o *EvaluationRunRecordFilterInput) GetInputFilters() ContentSearchFilterGroup`

GetInputFilters returns the InputFilters field if non-nil, zero value otherwise.

### GetInputFiltersOk

`func (o *EvaluationRunRecordFilterInput) GetInputFiltersOk() (*ContentSearchFilterGroup, bool)`

GetInputFiltersOk returns a tuple with the InputFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputFilters

`func (o *EvaluationRunRecordFilterInput) SetInputFilters(v ContentSearchFilterGroup)`

SetInputFilters sets InputFilters field to given value.

### HasInputFilters

`func (o *EvaluationRunRecordFilterInput) HasInputFilters() bool`

HasInputFilters returns a boolean if a field has been set.

### SetInputFiltersNil

`func (o *EvaluationRunRecordFilterInput) SetInputFiltersNil(b bool)`

 SetInputFiltersNil sets the value for InputFilters to be an explicit nil

### UnsetInputFilters
`func (o *EvaluationRunRecordFilterInput) UnsetInputFilters()`

UnsetInputFilters ensures that no value is present for InputFilters, not even an explicit nil
### GetOutputFilters

`func (o *EvaluationRunRecordFilterInput) GetOutputFilters() ContentSearchFilterGroup`

GetOutputFilters returns the OutputFilters field if non-nil, zero value otherwise.

### GetOutputFiltersOk

`func (o *EvaluationRunRecordFilterInput) GetOutputFiltersOk() (*ContentSearchFilterGroup, bool)`

GetOutputFiltersOk returns a tuple with the OutputFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputFilters

`func (o *EvaluationRunRecordFilterInput) SetOutputFilters(v ContentSearchFilterGroup)`

SetOutputFilters sets OutputFilters field to given value.

### HasOutputFilters

`func (o *EvaluationRunRecordFilterInput) HasOutputFilters() bool`

HasOutputFilters returns a boolean if a field has been set.

### SetOutputFiltersNil

`func (o *EvaluationRunRecordFilterInput) SetOutputFiltersNil(b bool)`

 SetOutputFiltersNil sets the value for OutputFilters to be an explicit nil

### UnsetOutputFilters
`func (o *EvaluationRunRecordFilterInput) UnsetOutputFilters()`

UnsetOutputFilters ensures that no value is present for OutputFilters, not even an explicit nil
### GetScoreFilters

`func (o *EvaluationRunRecordFilterInput) GetScoreFilters() ScoreFilterGroup`

GetScoreFilters returns the ScoreFilters field if non-nil, zero value otherwise.

### GetScoreFiltersOk

`func (o *EvaluationRunRecordFilterInput) GetScoreFiltersOk() (*ScoreFilterGroup, bool)`

GetScoreFiltersOk returns a tuple with the ScoreFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreFilters

`func (o *EvaluationRunRecordFilterInput) SetScoreFilters(v ScoreFilterGroup)`

SetScoreFilters sets ScoreFilters field to given value.

### HasScoreFilters

`func (o *EvaluationRunRecordFilterInput) HasScoreFilters() bool`

HasScoreFilters returns a boolean if a field has been set.

### SetScoreFiltersNil

`func (o *EvaluationRunRecordFilterInput) SetScoreFiltersNil(b bool)`

 SetScoreFiltersNil sets the value for ScoreFilters to be an explicit nil

### UnsetScoreFilters
`func (o *EvaluationRunRecordFilterInput) UnsetScoreFilters()`

UnsetScoreFilters ensures that no value is present for ScoreFilters, not even an explicit nil
### GetGoalFilters

`func (o *EvaluationRunRecordFilterInput) GetGoalFilters() GoalStatusFilterGroup`

GetGoalFilters returns the GoalFilters field if non-nil, zero value otherwise.

### GetGoalFiltersOk

`func (o *EvaluationRunRecordFilterInput) GetGoalFiltersOk() (*GoalStatusFilterGroup, bool)`

GetGoalFiltersOk returns a tuple with the GoalFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoalFilters

`func (o *EvaluationRunRecordFilterInput) SetGoalFilters(v GoalStatusFilterGroup)`

SetGoalFilters sets GoalFilters field to given value.

### HasGoalFilters

`func (o *EvaluationRunRecordFilterInput) HasGoalFilters() bool`

HasGoalFilters returns a boolean if a field has been set.

### SetGoalFiltersNil

`func (o *EvaluationRunRecordFilterInput) SetGoalFiltersNil(b bool)`

 SetGoalFiltersNil sets the value for GoalFilters to be an explicit nil

### UnsetGoalFilters
`func (o *EvaluationRunRecordFilterInput) UnsetGoalFilters()`

UnsetGoalFilters ensures that no value is present for GoalFilters, not even an explicit nil
### GetSort

`func (o *EvaluationRunRecordFilterInput) GetSort() EvaluationRunRecordFilterInputSort`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *EvaluationRunRecordFilterInput) GetSortOk() (*EvaluationRunRecordFilterInputSort, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *EvaluationRunRecordFilterInput) SetSort(v EvaluationRunRecordFilterInputSort)`

SetSort sets Sort field to given value.

### HasSort

`func (o *EvaluationRunRecordFilterInput) HasSort() bool`

HasSort returns a boolean if a field has been set.

### SetSortNil

`func (o *EvaluationRunRecordFilterInput) SetSortNil(b bool)`

 SetSortNil sets the value for Sort to be an explicit nil

### UnsetSort
`func (o *EvaluationRunRecordFilterInput) UnsetSort()`

UnsetSort ensures that no value is present for Sort, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


