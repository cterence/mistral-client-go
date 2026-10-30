# PaginatedResultEvaluationRunV2WithRefsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]EvaluationRunV2WithRefsResponse**](EvaluationRunV2WithRefsResponse.md) |  | [optional] 
**Count** | **int32** |  | 
**Next** | Pointer to **NullableString** |  | [optional] 
**Previous** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPaginatedResultEvaluationRunV2WithRefsResponse

`func NewPaginatedResultEvaluationRunV2WithRefsResponse(count int32, ) *PaginatedResultEvaluationRunV2WithRefsResponse`

NewPaginatedResultEvaluationRunV2WithRefsResponse instantiates a new PaginatedResultEvaluationRunV2WithRefsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedResultEvaluationRunV2WithRefsResponseWithDefaults

`func NewPaginatedResultEvaluationRunV2WithRefsResponseWithDefaults() *PaginatedResultEvaluationRunV2WithRefsResponse`

NewPaginatedResultEvaluationRunV2WithRefsResponseWithDefaults instantiates a new PaginatedResultEvaluationRunV2WithRefsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *PaginatedResultEvaluationRunV2WithRefsResponse) GetResults() []EvaluationRunV2WithRefsResponse`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedResultEvaluationRunV2WithRefsResponse) GetResultsOk() (*[]EvaluationRunV2WithRefsResponse, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedResultEvaluationRunV2WithRefsResponse) SetResults(v []EvaluationRunV2WithRefsResponse)`

SetResults sets Results field to given value.

### HasResults

`func (o *PaginatedResultEvaluationRunV2WithRefsResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetCount

`func (o *PaginatedResultEvaluationRunV2WithRefsResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedResultEvaluationRunV2WithRefsResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedResultEvaluationRunV2WithRefsResponse) SetCount(v int32)`

SetCount sets Count field to given value.


### GetNext

`func (o *PaginatedResultEvaluationRunV2WithRefsResponse) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedResultEvaluationRunV2WithRefsResponse) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedResultEvaluationRunV2WithRefsResponse) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedResultEvaluationRunV2WithRefsResponse) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *PaginatedResultEvaluationRunV2WithRefsResponse) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *PaginatedResultEvaluationRunV2WithRefsResponse) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *PaginatedResultEvaluationRunV2WithRefsResponse) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginatedResultEvaluationRunV2WithRefsResponse) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginatedResultEvaluationRunV2WithRefsResponse) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginatedResultEvaluationRunV2WithRefsResponse) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *PaginatedResultEvaluationRunV2WithRefsResponse) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *PaginatedResultEvaluationRunV2WithRefsResponse) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


