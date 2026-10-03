# PaginatedResultDatasetImportTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]DatasetImportTask**](DatasetImportTask.md) |  | [optional] 
**Count** | **int32** |  | 
**Next** | Pointer to **NullableString** |  | [optional] 
**Previous** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPaginatedResultDatasetImportTask

`func NewPaginatedResultDatasetImportTask(count int32, ) *PaginatedResultDatasetImportTask`

NewPaginatedResultDatasetImportTask instantiates a new PaginatedResultDatasetImportTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedResultDatasetImportTaskWithDefaults

`func NewPaginatedResultDatasetImportTaskWithDefaults() *PaginatedResultDatasetImportTask`

NewPaginatedResultDatasetImportTaskWithDefaults instantiates a new PaginatedResultDatasetImportTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *PaginatedResultDatasetImportTask) GetResults() []DatasetImportTask`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedResultDatasetImportTask) GetResultsOk() (*[]DatasetImportTask, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedResultDatasetImportTask) SetResults(v []DatasetImportTask)`

SetResults sets Results field to given value.

### HasResults

`func (o *PaginatedResultDatasetImportTask) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetCount

`func (o *PaginatedResultDatasetImportTask) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedResultDatasetImportTask) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedResultDatasetImportTask) SetCount(v int32)`

SetCount sets Count field to given value.


### GetNext

`func (o *PaginatedResultDatasetImportTask) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedResultDatasetImportTask) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedResultDatasetImportTask) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedResultDatasetImportTask) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *PaginatedResultDatasetImportTask) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *PaginatedResultDatasetImportTask) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *PaginatedResultDatasetImportTask) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginatedResultDatasetImportTask) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginatedResultDatasetImportTask) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginatedResultDatasetImportTask) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *PaginatedResultDatasetImportTask) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *PaginatedResultDatasetImportTask) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


