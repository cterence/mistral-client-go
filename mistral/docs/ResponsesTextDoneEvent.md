# ResponsesTextDoneEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentIndex** | **int32** |  | 
**ItemId** | **string** |  | 
**Logprobs** | [**[]ResponsesTextDoneEventLogprob**](ResponsesTextDoneEventLogprob.md) |  | 
**OutputIndex** | **int32** |  | 
**SequenceNumber** | **int32** |  | 
**Text** | **string** |  | 
**Type** | **string** |  | 

## Methods

### NewResponsesTextDoneEvent

`func NewResponsesTextDoneEvent(contentIndex int32, itemId string, logprobs []ResponsesTextDoneEventLogprob, outputIndex int32, sequenceNumber int32, text string, type_ string, ) *ResponsesTextDoneEvent`

NewResponsesTextDoneEvent instantiates a new ResponsesTextDoneEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsesTextDoneEventWithDefaults

`func NewResponsesTextDoneEventWithDefaults() *ResponsesTextDoneEvent`

NewResponsesTextDoneEventWithDefaults instantiates a new ResponsesTextDoneEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentIndex

`func (o *ResponsesTextDoneEvent) GetContentIndex() int32`

GetContentIndex returns the ContentIndex field if non-nil, zero value otherwise.

### GetContentIndexOk

`func (o *ResponsesTextDoneEvent) GetContentIndexOk() (*int32, bool)`

GetContentIndexOk returns a tuple with the ContentIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentIndex

`func (o *ResponsesTextDoneEvent) SetContentIndex(v int32)`

SetContentIndex sets ContentIndex field to given value.


### GetItemId

`func (o *ResponsesTextDoneEvent) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *ResponsesTextDoneEvent) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *ResponsesTextDoneEvent) SetItemId(v string)`

SetItemId sets ItemId field to given value.


### GetLogprobs

`func (o *ResponsesTextDoneEvent) GetLogprobs() []ResponsesTextDoneEventLogprob`

GetLogprobs returns the Logprobs field if non-nil, zero value otherwise.

### GetLogprobsOk

`func (o *ResponsesTextDoneEvent) GetLogprobsOk() (*[]ResponsesTextDoneEventLogprob, bool)`

GetLogprobsOk returns a tuple with the Logprobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogprobs

`func (o *ResponsesTextDoneEvent) SetLogprobs(v []ResponsesTextDoneEventLogprob)`

SetLogprobs sets Logprobs field to given value.


### GetOutputIndex

`func (o *ResponsesTextDoneEvent) GetOutputIndex() int32`

GetOutputIndex returns the OutputIndex field if non-nil, zero value otherwise.

### GetOutputIndexOk

`func (o *ResponsesTextDoneEvent) GetOutputIndexOk() (*int32, bool)`

GetOutputIndexOk returns a tuple with the OutputIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputIndex

`func (o *ResponsesTextDoneEvent) SetOutputIndex(v int32)`

SetOutputIndex sets OutputIndex field to given value.


### GetSequenceNumber

`func (o *ResponsesTextDoneEvent) GetSequenceNumber() int32`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *ResponsesTextDoneEvent) GetSequenceNumberOk() (*int32, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *ResponsesTextDoneEvent) SetSequenceNumber(v int32)`

SetSequenceNumber sets SequenceNumber field to given value.


### GetText

`func (o *ResponsesTextDoneEvent) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ResponsesTextDoneEvent) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ResponsesTextDoneEvent) SetText(v string)`

SetText sets Text field to given value.


### GetType

`func (o *ResponsesTextDoneEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResponsesTextDoneEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResponsesTextDoneEvent) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


