# ExecutionLogSearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | [**[]ExecutionLogRecord**](ExecutionLogRecord.md) |  | 
**NextCursor** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewExecutionLogSearchResponse

`func NewExecutionLogSearchResponse(results []ExecutionLogRecord, ) *ExecutionLogSearchResponse`

NewExecutionLogSearchResponse instantiates a new ExecutionLogSearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecutionLogSearchResponseWithDefaults

`func NewExecutionLogSearchResponseWithDefaults() *ExecutionLogSearchResponse`

NewExecutionLogSearchResponseWithDefaults instantiates a new ExecutionLogSearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *ExecutionLogSearchResponse) GetResults() []ExecutionLogRecord`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ExecutionLogSearchResponse) GetResultsOk() (*[]ExecutionLogRecord, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ExecutionLogSearchResponse) SetResults(v []ExecutionLogRecord)`

SetResults sets Results field to given value.


### GetNextCursor

`func (o *ExecutionLogSearchResponse) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *ExecutionLogSearchResponse) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *ExecutionLogSearchResponse) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *ExecutionLogSearchResponse) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.

### SetNextCursorNil

`func (o *ExecutionLogSearchResponse) SetNextCursorNil(b bool)`

 SetNextCursorNil sets the value for NextCursor to be an explicit nil

### UnsetNextCursor
`func (o *ExecutionLogSearchResponse) UnsetNextCursor()`

UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


