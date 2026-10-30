# ResponsesOutputItemDoneEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Item** | [**Item**](Item.md) |  | 
**OutputIndex** | **int32** |  | 
**SequenceNumber** | **int32** |  | 
**Type** | **string** |  | 

## Methods

### NewResponsesOutputItemDoneEvent

`func NewResponsesOutputItemDoneEvent(item Item, outputIndex int32, sequenceNumber int32, type_ string, ) *ResponsesOutputItemDoneEvent`

NewResponsesOutputItemDoneEvent instantiates a new ResponsesOutputItemDoneEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsesOutputItemDoneEventWithDefaults

`func NewResponsesOutputItemDoneEventWithDefaults() *ResponsesOutputItemDoneEvent`

NewResponsesOutputItemDoneEventWithDefaults instantiates a new ResponsesOutputItemDoneEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItem

`func (o *ResponsesOutputItemDoneEvent) GetItem() Item`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *ResponsesOutputItemDoneEvent) GetItemOk() (*Item, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *ResponsesOutputItemDoneEvent) SetItem(v Item)`

SetItem sets Item field to given value.


### GetOutputIndex

`func (o *ResponsesOutputItemDoneEvent) GetOutputIndex() int32`

GetOutputIndex returns the OutputIndex field if non-nil, zero value otherwise.

### GetOutputIndexOk

`func (o *ResponsesOutputItemDoneEvent) GetOutputIndexOk() (*int32, bool)`

GetOutputIndexOk returns a tuple with the OutputIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputIndex

`func (o *ResponsesOutputItemDoneEvent) SetOutputIndex(v int32)`

SetOutputIndex sets OutputIndex field to given value.


### GetSequenceNumber

`func (o *ResponsesOutputItemDoneEvent) GetSequenceNumber() int32`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *ResponsesOutputItemDoneEvent) GetSequenceNumberOk() (*int32, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *ResponsesOutputItemDoneEvent) SetSequenceNumber(v int32)`

SetSequenceNumber sets SequenceNumber field to given value.


### GetType

`func (o *ResponsesOutputItemDoneEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResponsesOutputItemDoneEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResponsesOutputItemDoneEvent) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


