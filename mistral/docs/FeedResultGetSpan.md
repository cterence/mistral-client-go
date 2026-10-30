# FeedResultGetSpan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]GetSpan**](GetSpan.md) |  | [optional] 
**Next** | Pointer to **NullableString** |  | [optional] 
**Cursor** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewFeedResultGetSpan

`func NewFeedResultGetSpan() *FeedResultGetSpan`

NewFeedResultGetSpan instantiates a new FeedResultGetSpan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeedResultGetSpanWithDefaults

`func NewFeedResultGetSpanWithDefaults() *FeedResultGetSpan`

NewFeedResultGetSpanWithDefaults instantiates a new FeedResultGetSpan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *FeedResultGetSpan) GetResults() []GetSpan`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *FeedResultGetSpan) GetResultsOk() (*[]GetSpan, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *FeedResultGetSpan) SetResults(v []GetSpan)`

SetResults sets Results field to given value.

### HasResults

`func (o *FeedResultGetSpan) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetNext

`func (o *FeedResultGetSpan) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *FeedResultGetSpan) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *FeedResultGetSpan) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *FeedResultGetSpan) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *FeedResultGetSpan) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *FeedResultGetSpan) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetCursor

`func (o *FeedResultGetSpan) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *FeedResultGetSpan) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *FeedResultGetSpan) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *FeedResultGetSpan) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### SetCursorNil

`func (o *FeedResultGetSpan) SetCursorNil(b bool)`

 SetCursorNil sets the value for Cursor to be an explicit nil

### UnsetCursor
`func (o *FeedResultGetSpan) UnsetCursor()`

UnsetCursor ensures that no value is present for Cursor, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


