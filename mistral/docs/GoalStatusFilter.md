# GoalStatusFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EvaluatorName** | Pointer to **NullableString** |  | [optional] 
**Values** | [**[]GoalStatus**](GoalStatus.md) |  | 

## Methods

### NewGoalStatusFilter

`func NewGoalStatusFilter(values []GoalStatus, ) *GoalStatusFilter`

NewGoalStatusFilter instantiates a new GoalStatusFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoalStatusFilterWithDefaults

`func NewGoalStatusFilterWithDefaults() *GoalStatusFilter`

NewGoalStatusFilterWithDefaults instantiates a new GoalStatusFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvaluatorName

`func (o *GoalStatusFilter) GetEvaluatorName() string`

GetEvaluatorName returns the EvaluatorName field if non-nil, zero value otherwise.

### GetEvaluatorNameOk

`func (o *GoalStatusFilter) GetEvaluatorNameOk() (*string, bool)`

GetEvaluatorNameOk returns a tuple with the EvaluatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluatorName

`func (o *GoalStatusFilter) SetEvaluatorName(v string)`

SetEvaluatorName sets EvaluatorName field to given value.

### HasEvaluatorName

`func (o *GoalStatusFilter) HasEvaluatorName() bool`

HasEvaluatorName returns a boolean if a field has been set.

### SetEvaluatorNameNil

`func (o *GoalStatusFilter) SetEvaluatorNameNil(b bool)`

 SetEvaluatorNameNil sets the value for EvaluatorName to be an explicit nil

### UnsetEvaluatorName
`func (o *GoalStatusFilter) UnsetEvaluatorName()`

UnsetEvaluatorName ensures that no value is present for EvaluatorName, not even an explicit nil
### GetValues

`func (o *GoalStatusFilter) GetValues() []GoalStatus`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GoalStatusFilter) GetValuesOk() (*[]GoalStatus, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GoalStatusFilter) SetValues(v []GoalStatus)`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


