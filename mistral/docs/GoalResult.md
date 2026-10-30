# GoalResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Passed** | **NullableBool** |  | 
**Results** | [**[]GoalResultItem**](GoalResultItem.md) |  | 

## Methods

### NewGoalResult

`func NewGoalResult(passed NullableBool, results []GoalResultItem, ) *GoalResult`

NewGoalResult instantiates a new GoalResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoalResultWithDefaults

`func NewGoalResultWithDefaults() *GoalResult`

NewGoalResultWithDefaults instantiates a new GoalResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassed

`func (o *GoalResult) GetPassed() bool`

GetPassed returns the Passed field if non-nil, zero value otherwise.

### GetPassedOk

`func (o *GoalResult) GetPassedOk() (*bool, bool)`

GetPassedOk returns a tuple with the Passed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassed

`func (o *GoalResult) SetPassed(v bool)`

SetPassed sets Passed field to given value.


### SetPassedNil

`func (o *GoalResult) SetPassedNil(b bool)`

 SetPassedNil sets the value for Passed to be an explicit nil

### UnsetPassed
`func (o *GoalResult) UnsetPassed()`

UnsetPassed ensures that no value is present for Passed, not even an explicit nil
### GetResults

`func (o *GoalResult) GetResults() []GoalResultItem`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *GoalResult) GetResultsOk() (*[]GoalResultItem, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *GoalResult) SetResults(v []GoalResultItem)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


