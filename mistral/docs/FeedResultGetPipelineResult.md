# FeedResultGetPipelineResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]PipelineResult**](PipelineResult.md) |  | [optional] 
**Next** | Pointer to **NullableString** |  | [optional] 
**Cursor** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewFeedResultGetPipelineResult

`func NewFeedResultGetPipelineResult() *FeedResultGetPipelineResult`

NewFeedResultGetPipelineResult instantiates a new FeedResultGetPipelineResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeedResultGetPipelineResultWithDefaults

`func NewFeedResultGetPipelineResultWithDefaults() *FeedResultGetPipelineResult`

NewFeedResultGetPipelineResultWithDefaults instantiates a new FeedResultGetPipelineResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *FeedResultGetPipelineResult) GetResults() []PipelineResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *FeedResultGetPipelineResult) GetResultsOk() (*[]PipelineResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *FeedResultGetPipelineResult) SetResults(v []PipelineResult)`

SetResults sets Results field to given value.

### HasResults

`func (o *FeedResultGetPipelineResult) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetNext

`func (o *FeedResultGetPipelineResult) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *FeedResultGetPipelineResult) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *FeedResultGetPipelineResult) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *FeedResultGetPipelineResult) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *FeedResultGetPipelineResult) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *FeedResultGetPipelineResult) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetCursor

`func (o *FeedResultGetPipelineResult) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *FeedResultGetPipelineResult) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *FeedResultGetPipelineResult) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *FeedResultGetPipelineResult) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### SetCursorNil

`func (o *FeedResultGetPipelineResult) SetCursorNil(b bool)`

 SetCursorNil sets the value for Cursor to be an explicit nil

### UnsetCursor
`func (o *FeedResultGetPipelineResult) UnsetCursor()`

UnsetCursor ensures that no value is present for Cursor, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


