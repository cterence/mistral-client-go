# ResponsesContentPartAddedEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentIndex** | **int32** |  | 
**ItemId** | **string** |  | 
**OutputIndex** | **int32** |  | 
**Part** | [**Part**](Part.md) |  | 
**SequenceNumber** | **int32** |  | 
**Type** | **string** |  | 

## Methods

### NewResponsesContentPartAddedEvent

`func NewResponsesContentPartAddedEvent(contentIndex int32, itemId string, outputIndex int32, part Part, sequenceNumber int32, type_ string, ) *ResponsesContentPartAddedEvent`

NewResponsesContentPartAddedEvent instantiates a new ResponsesContentPartAddedEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsesContentPartAddedEventWithDefaults

`func NewResponsesContentPartAddedEventWithDefaults() *ResponsesContentPartAddedEvent`

NewResponsesContentPartAddedEventWithDefaults instantiates a new ResponsesContentPartAddedEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentIndex

`func (o *ResponsesContentPartAddedEvent) GetContentIndex() int32`

GetContentIndex returns the ContentIndex field if non-nil, zero value otherwise.

### GetContentIndexOk

`func (o *ResponsesContentPartAddedEvent) GetContentIndexOk() (*int32, bool)`

GetContentIndexOk returns a tuple with the ContentIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentIndex

`func (o *ResponsesContentPartAddedEvent) SetContentIndex(v int32)`

SetContentIndex sets ContentIndex field to given value.


### GetItemId

`func (o *ResponsesContentPartAddedEvent) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *ResponsesContentPartAddedEvent) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *ResponsesContentPartAddedEvent) SetItemId(v string)`

SetItemId sets ItemId field to given value.


### GetOutputIndex

`func (o *ResponsesContentPartAddedEvent) GetOutputIndex() int32`

GetOutputIndex returns the OutputIndex field if non-nil, zero value otherwise.

### GetOutputIndexOk

`func (o *ResponsesContentPartAddedEvent) GetOutputIndexOk() (*int32, bool)`

GetOutputIndexOk returns a tuple with the OutputIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputIndex

`func (o *ResponsesContentPartAddedEvent) SetOutputIndex(v int32)`

SetOutputIndex sets OutputIndex field to given value.


### GetPart

`func (o *ResponsesContentPartAddedEvent) GetPart() Part`

GetPart returns the Part field if non-nil, zero value otherwise.

### GetPartOk

`func (o *ResponsesContentPartAddedEvent) GetPartOk() (*Part, bool)`

GetPartOk returns a tuple with the Part field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPart

`func (o *ResponsesContentPartAddedEvent) SetPart(v Part)`

SetPart sets Part field to given value.


### GetSequenceNumber

`func (o *ResponsesContentPartAddedEvent) GetSequenceNumber() int32`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *ResponsesContentPartAddedEvent) GetSequenceNumberOk() (*int32, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *ResponsesContentPartAddedEvent) SetSequenceNumber(v int32)`

SetSequenceNumber sets SequenceNumber field to given value.


### GetType

`func (o *ResponsesContentPartAddedEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResponsesContentPartAddedEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResponsesContentPartAddedEvent) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


