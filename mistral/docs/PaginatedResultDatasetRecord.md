# PaginatedResultDatasetRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]DatasetRecord**](DatasetRecord.md) |  | [optional] 
**Count** | **int32** |  | 
**Next** | Pointer to **NullableString** |  | [optional] 
**Previous** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPaginatedResultDatasetRecord

`func NewPaginatedResultDatasetRecord(count int32, ) *PaginatedResultDatasetRecord`

NewPaginatedResultDatasetRecord instantiates a new PaginatedResultDatasetRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedResultDatasetRecordWithDefaults

`func NewPaginatedResultDatasetRecordWithDefaults() *PaginatedResultDatasetRecord`

NewPaginatedResultDatasetRecordWithDefaults instantiates a new PaginatedResultDatasetRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *PaginatedResultDatasetRecord) GetResults() []DatasetRecord`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedResultDatasetRecord) GetResultsOk() (*[]DatasetRecord, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedResultDatasetRecord) SetResults(v []DatasetRecord)`

SetResults sets Results field to given value.

### HasResults

`func (o *PaginatedResultDatasetRecord) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetCount

`func (o *PaginatedResultDatasetRecord) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedResultDatasetRecord) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedResultDatasetRecord) SetCount(v int32)`

SetCount sets Count field to given value.


### GetNext

`func (o *PaginatedResultDatasetRecord) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedResultDatasetRecord) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedResultDatasetRecord) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedResultDatasetRecord) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *PaginatedResultDatasetRecord) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *PaginatedResultDatasetRecord) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *PaginatedResultDatasetRecord) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginatedResultDatasetRecord) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginatedResultDatasetRecord) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginatedResultDatasetRecord) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *PaginatedResultDatasetRecord) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *PaginatedResultDatasetRecord) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


