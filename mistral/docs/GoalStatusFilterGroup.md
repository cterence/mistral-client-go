# GoalStatusFilterGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Combinator** | Pointer to [**FilterCombinator**](FilterCombinator.md) |  | [optional] [default to FILTERCOMBINATOR_AND]
**Filters** | Pointer to [**[]GoalStatusFilter**](GoalStatusFilter.md) |  | [optional] [default to {}]

## Methods

### NewGoalStatusFilterGroup

`func NewGoalStatusFilterGroup() *GoalStatusFilterGroup`

NewGoalStatusFilterGroup instantiates a new GoalStatusFilterGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoalStatusFilterGroupWithDefaults

`func NewGoalStatusFilterGroupWithDefaults() *GoalStatusFilterGroup`

NewGoalStatusFilterGroupWithDefaults instantiates a new GoalStatusFilterGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCombinator

`func (o *GoalStatusFilterGroup) GetCombinator() FilterCombinator`

GetCombinator returns the Combinator field if non-nil, zero value otherwise.

### GetCombinatorOk

`func (o *GoalStatusFilterGroup) GetCombinatorOk() (*FilterCombinator, bool)`

GetCombinatorOk returns a tuple with the Combinator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCombinator

`func (o *GoalStatusFilterGroup) SetCombinator(v FilterCombinator)`

SetCombinator sets Combinator field to given value.

### HasCombinator

`func (o *GoalStatusFilterGroup) HasCombinator() bool`

HasCombinator returns a boolean if a field has been set.

### GetFilters

`func (o *GoalStatusFilterGroup) GetFilters() []GoalStatusFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *GoalStatusFilterGroup) GetFiltersOk() (*[]GoalStatusFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *GoalStatusFilterGroup) SetFilters(v []GoalStatusFilter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *GoalStatusFilterGroup) HasFilters() bool`

HasFilters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


