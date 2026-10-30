# PaginatedResultOptimizationTrialWithObservations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]OptimizationTrialWithObservations**](OptimizationTrialWithObservations.md) |  | [optional] 
**Count** | **int32** |  | 
**Next** | Pointer to **NullableString** |  | [optional] 
**Previous** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPaginatedResultOptimizationTrialWithObservations

`func NewPaginatedResultOptimizationTrialWithObservations(count int32, ) *PaginatedResultOptimizationTrialWithObservations`

NewPaginatedResultOptimizationTrialWithObservations instantiates a new PaginatedResultOptimizationTrialWithObservations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedResultOptimizationTrialWithObservationsWithDefaults

`func NewPaginatedResultOptimizationTrialWithObservationsWithDefaults() *PaginatedResultOptimizationTrialWithObservations`

NewPaginatedResultOptimizationTrialWithObservationsWithDefaults instantiates a new PaginatedResultOptimizationTrialWithObservations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *PaginatedResultOptimizationTrialWithObservations) GetResults() []OptimizationTrialWithObservations`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedResultOptimizationTrialWithObservations) GetResultsOk() (*[]OptimizationTrialWithObservations, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedResultOptimizationTrialWithObservations) SetResults(v []OptimizationTrialWithObservations)`

SetResults sets Results field to given value.

### HasResults

`func (o *PaginatedResultOptimizationTrialWithObservations) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetCount

`func (o *PaginatedResultOptimizationTrialWithObservations) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedResultOptimizationTrialWithObservations) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedResultOptimizationTrialWithObservations) SetCount(v int32)`

SetCount sets Count field to given value.


### GetNext

`func (o *PaginatedResultOptimizationTrialWithObservations) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedResultOptimizationTrialWithObservations) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedResultOptimizationTrialWithObservations) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedResultOptimizationTrialWithObservations) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *PaginatedResultOptimizationTrialWithObservations) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *PaginatedResultOptimizationTrialWithObservations) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *PaginatedResultOptimizationTrialWithObservations) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginatedResultOptimizationTrialWithObservations) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginatedResultOptimizationTrialWithObservations) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginatedResultOptimizationTrialWithObservations) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *PaginatedResultOptimizationTrialWithObservations) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *PaginatedResultOptimizationTrialWithObservations) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


