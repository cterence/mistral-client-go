# FeedResultChatCompletionEventPreview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]ChatCompletionEventPreview**](ChatCompletionEventPreview.md) |  | [optional] 
**Next** | Pointer to **NullableString** |  | [optional] 
**Cursor** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewFeedResultChatCompletionEventPreview

`func NewFeedResultChatCompletionEventPreview() *FeedResultChatCompletionEventPreview`

NewFeedResultChatCompletionEventPreview instantiates a new FeedResultChatCompletionEventPreview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeedResultChatCompletionEventPreviewWithDefaults

`func NewFeedResultChatCompletionEventPreviewWithDefaults() *FeedResultChatCompletionEventPreview`

NewFeedResultChatCompletionEventPreviewWithDefaults instantiates a new FeedResultChatCompletionEventPreview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *FeedResultChatCompletionEventPreview) GetResults() []ChatCompletionEventPreview`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *FeedResultChatCompletionEventPreview) GetResultsOk() (*[]ChatCompletionEventPreview, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *FeedResultChatCompletionEventPreview) SetResults(v []ChatCompletionEventPreview)`

SetResults sets Results field to given value.

### HasResults

`func (o *FeedResultChatCompletionEventPreview) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetNext

`func (o *FeedResultChatCompletionEventPreview) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *FeedResultChatCompletionEventPreview) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *FeedResultChatCompletionEventPreview) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *FeedResultChatCompletionEventPreview) HasNext() bool`

HasNext returns a boolean if a field has been set.

### SetNextNil

`func (o *FeedResultChatCompletionEventPreview) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *FeedResultChatCompletionEventPreview) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil
### GetCursor

`func (o *FeedResultChatCompletionEventPreview) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *FeedResultChatCompletionEventPreview) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *FeedResultChatCompletionEventPreview) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *FeedResultChatCompletionEventPreview) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### SetCursorNil

`func (o *FeedResultChatCompletionEventPreview) SetCursorNil(b bool)`

 SetCursorNil sets the value for Cursor to be an explicit nil

### UnsetCursor
`func (o *FeedResultChatCompletionEventPreview) UnsetCursor()`

UnsetCursor ensures that no value is present for Cursor, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


