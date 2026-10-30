# ResponsesOutputItemAddedEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Item** | [**Item**](Item.md) |  | 
**OutputIndex** | **int32** |  | 
**SequenceNumber** | **int32** |  | 
**Type** | **string** |  | 

## Methods

### NewResponsesOutputItemAddedEvent

`func NewResponsesOutputItemAddedEvent(item Item, outputIndex int32, sequenceNumber int32, type_ string, ) *ResponsesOutputItemAddedEvent`

NewResponsesOutputItemAddedEvent instantiates a new ResponsesOutputItemAddedEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsesOutputItemAddedEventWithDefaults

`func NewResponsesOutputItemAddedEventWithDefaults() *ResponsesOutputItemAddedEvent`

NewResponsesOutputItemAddedEventWithDefaults instantiates a new ResponsesOutputItemAddedEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItem

`func (o *ResponsesOutputItemAddedEvent) GetItem() Item`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *ResponsesOutputItemAddedEvent) GetItemOk() (*Item, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *ResponsesOutputItemAddedEvent) SetItem(v Item)`

SetItem sets Item field to given value.


### GetOutputIndex

`func (o *ResponsesOutputItemAddedEvent) GetOutputIndex() int32`

GetOutputIndex returns the OutputIndex field if non-nil, zero value otherwise.

### GetOutputIndexOk

`func (o *ResponsesOutputItemAddedEvent) GetOutputIndexOk() (*int32, bool)`

GetOutputIndexOk returns a tuple with the OutputIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputIndex

`func (o *ResponsesOutputItemAddedEvent) SetOutputIndex(v int32)`

SetOutputIndex sets OutputIndex field to given value.


### GetSequenceNumber

`func (o *ResponsesOutputItemAddedEvent) GetSequenceNumber() int32`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *ResponsesOutputItemAddedEvent) GetSequenceNumberOk() (*int32, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *ResponsesOutputItemAddedEvent) SetSequenceNumber(v int32)`

SetSequenceNumber sets SequenceNumber field to given value.


### GetType

`func (o *ResponsesOutputItemAddedEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResponsesOutputItemAddedEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResponsesOutputItemAddedEvent) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


