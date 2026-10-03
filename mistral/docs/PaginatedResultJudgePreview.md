# PaginatedResultJudgePreview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]JudgePreview**](JudgePreview.md) |  | [optional] 
**Count** | **int32** |  | 
**Next** | Pointer to **NullableString** |  | [optional] 
**Previous** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPaginatedResultJudgePreview

`func NewPaginatedResultJudgePreview(count int32, ) *PaginatedResultJudgePreview`

NewPaginatedResultJudgePreview instantiates a new PaginatedResultJudgePreview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedResultJudgePreviewWithDefaults

`func NewPaginatedResultJudgePreviewWithDefaults() *PaginatedResultJudgePreview`

NewPaginatedResultJudgePreviewWithDefaults instantiates a new PaginatedResultJudgePreview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *PaginatedResultJudgePreview) GetResults() []JudgePreview`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedResultJudgePreview) GetResultsOk() (*[]JudgePreview, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedResultJudgePreview) SetResults(v []JudgePreview)`

SetResults sets Results field to given value.

### HasResults

`func (o *PaginatedResultJudgePreview) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetCount

`func (o *PaginatedResultJudgePreview) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedResultJudgePreview) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedResultJudgePreview) SetCount(v int32)`

SetCount sets Count field to given value.


### GetNext

`func (o *PaginatedResultJudgePreview) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedResultJudgePreview) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedResultJudgePreview) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedResultJudgePreview) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *PaginatedResultJudgePreview) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *PaginatedResultJudgePreview) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *PaginatedResultJudgePreview) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginatedResultJudgePreview) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginatedResultJudgePreview) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginatedResultJudgePreview) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *PaginatedResultJudgePreview) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *PaginatedResultJudgePreview) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


