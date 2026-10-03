# BatchExecutionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**map[string]BatchExecutionResult**](BatchExecutionResult.md) | Mapping of execution_id to result with status and optional error message | [optional] 

## Methods

### NewBatchExecutionResponse

`func NewBatchExecutionResponse() *BatchExecutionResponse`

NewBatchExecutionResponse instantiates a new BatchExecutionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchExecutionResponseWithDefaults

`func NewBatchExecutionResponseWithDefaults() *BatchExecutionResponse`

NewBatchExecutionResponseWithDefaults instantiates a new BatchExecutionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *BatchExecutionResponse) GetResults() map[string]BatchExecutionResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *BatchExecutionResponse) GetResultsOk() (*map[string]BatchExecutionResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *BatchExecutionResponse) SetResults(v map[string]BatchExecutionResult)`

SetResults sets Results field to given value.

### HasResults

`func (o *BatchExecutionResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


