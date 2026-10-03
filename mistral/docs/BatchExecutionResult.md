# BatchExecutionResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | Status of the operation (success/failure) | 
**Error** | Pointer to **NullableString** | Error message if operation failed | [optional] 

## Methods

### NewBatchExecutionResult

`func NewBatchExecutionResult(status string, ) *BatchExecutionResult`

NewBatchExecutionResult instantiates a new BatchExecutionResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchExecutionResultWithDefaults

`func NewBatchExecutionResultWithDefaults() *BatchExecutionResult`

NewBatchExecutionResultWithDefaults instantiates a new BatchExecutionResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *BatchExecutionResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BatchExecutionResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BatchExecutionResult) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetError

`func (o *BatchExecutionResult) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *BatchExecutionResult) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *BatchExecutionResult) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *BatchExecutionResult) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *BatchExecutionResult) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *BatchExecutionResult) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


