# UpdateRunInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExecutionTime** | **time.Time** |  | 
**ChunksCount** | **int32** |  | 

## Methods

### NewUpdateRunInfo

`func NewUpdateRunInfo(executionTime time.Time, chunksCount int32, ) *UpdateRunInfo`

NewUpdateRunInfo instantiates a new UpdateRunInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRunInfoWithDefaults

`func NewUpdateRunInfoWithDefaults() *UpdateRunInfo`

NewUpdateRunInfoWithDefaults instantiates a new UpdateRunInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExecutionTime

`func (o *UpdateRunInfo) GetExecutionTime() time.Time`

GetExecutionTime returns the ExecutionTime field if non-nil, zero value otherwise.

### GetExecutionTimeOk

`func (o *UpdateRunInfo) GetExecutionTimeOk() (*time.Time, bool)`

GetExecutionTimeOk returns a tuple with the ExecutionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionTime

`func (o *UpdateRunInfo) SetExecutionTime(v time.Time)`

SetExecutionTime sets ExecutionTime field to given value.


### GetChunksCount

`func (o *UpdateRunInfo) GetChunksCount() int32`

GetChunksCount returns the ChunksCount field if non-nil, zero value otherwise.

### GetChunksCountOk

`func (o *UpdateRunInfo) GetChunksCountOk() (*int32, bool)`

GetChunksCountOk returns a tuple with the ChunksCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChunksCount

`func (o *UpdateRunInfo) SetChunksCount(v int32)`

SetChunksCount sets ChunksCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


