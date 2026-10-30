# PaginatedResultEvaluationRunRecordV2Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]EvaluationRunRecordV2Response**](EvaluationRunRecordV2Response.md) |  | [optional] 
**Count** | **int32** |  | 
**Next** | Pointer to **NullableString** |  | [optional] 
**Previous** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPaginatedResultEvaluationRunRecordV2Response

`func NewPaginatedResultEvaluationRunRecordV2Response(count int32, ) *PaginatedResultEvaluationRunRecordV2Response`

NewPaginatedResultEvaluationRunRecordV2Response instantiates a new PaginatedResultEvaluationRunRecordV2Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedResultEvaluationRunRecordV2ResponseWithDefaults

`func NewPaginatedResultEvaluationRunRecordV2ResponseWithDefaults() *PaginatedResultEvaluationRunRecordV2Response`

NewPaginatedResultEvaluationRunRecordV2ResponseWithDefaults instantiates a new PaginatedResultEvaluationRunRecordV2Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *PaginatedResultEvaluationRunRecordV2Response) GetResults() []EvaluationRunRecordV2Response`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedResultEvaluationRunRecordV2Response) GetResultsOk() (*[]EvaluationRunRecordV2Response, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedResultEvaluationRunRecordV2Response) SetResults(v []EvaluationRunRecordV2Response)`

SetResults sets Results field to given value.

### HasResults

`func (o *PaginatedResultEvaluationRunRecordV2Response) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetCount

`func (o *PaginatedResultEvaluationRunRecordV2Response) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaginatedResultEvaluationRunRecordV2Response) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaginatedResultEvaluationRunRecordV2Response) SetCount(v int32)`

SetCount sets Count field to given value.


### GetNext

`func (o *PaginatedResultEvaluationRunRecordV2Response) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedResultEvaluationRunRecordV2Response) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedResultEvaluationRunRecordV2Response) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedResultEvaluationRunRecordV2Response) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *PaginatedResultEvaluationRunRecordV2Response) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *PaginatedResultEvaluationRunRecordV2Response) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetPrevious

`func (o *PaginatedResultEvaluationRunRecordV2Response) GetPrevious() string`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *PaginatedResultEvaluationRunRecordV2Response) GetPreviousOk() (*string, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *PaginatedResultEvaluationRunRecordV2Response) SetPrevious(v string)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *PaginatedResultEvaluationRunRecordV2Response) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.

### SetPreviousNil

`func (o *PaginatedResultEvaluationRunRecordV2Response) SetPreviousNil(b bool)`

 SetPreviousNil sets the value for Previous to be an explicit nil

### UnsetPrevious
`func (o *PaginatedResultEvaluationRunRecordV2Response) UnsetPrevious()`

UnsetPrevious ensures that no value is present for Previous, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


