# Data4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Response** | [**CreateResponseResponse**](CreateResponseResponse.md) |  | 
**SequenceNumber** | **int32** |  | 
**Type** | **string** |  | 
**Code** | Pointer to **NullableString** |  | [optional] 
**Message** | **string** |  | 
**Param** | Pointer to **NullableString** |  | [optional] 
**Item** | [**Item**](Item.md) |  | 
**OutputIndex** | **int32** |  | 
**ContentIndex** | **int32** |  | 
**ItemId** | **string** |  | 
**Part** | [**Part**](Part.md) |  | 
**Delta** | **string** |  | 
**Logprobs** | [**[]ResponsesTextDoneEventLogprob**](ResponsesTextDoneEventLogprob.md) |  | 
**Text** | **string** |  | 
**Arguments** | **string** |  | 
**Name** | **string** |  | 

## Methods

### NewData4

`func NewData4(response CreateResponseResponse, sequenceNumber int32, type_ string, message string, item Item, outputIndex int32, contentIndex int32, itemId string, part Part, delta string, logprobs []ResponsesTextDoneEventLogprob, text string, arguments string, name string, ) *Data4`

NewData4 instantiates a new Data4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewData4WithDefaults

`func NewData4WithDefaults() *Data4`

NewData4WithDefaults instantiates a new Data4 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponse

`func (o *Data4) GetResponse() CreateResponseResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *Data4) GetResponseOk() (*CreateResponseResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *Data4) SetResponse(v CreateResponseResponse)`

SetResponse sets Response field to given value.


### GetSequenceNumber

`func (o *Data4) GetSequenceNumber() int32`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *Data4) GetSequenceNumberOk() (*int32, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *Data4) SetSequenceNumber(v int32)`

SetSequenceNumber sets SequenceNumber field to given value.


### GetType

`func (o *Data4) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Data4) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Data4) SetType(v string)`

SetType sets Type field to given value.


### GetCode

`func (o *Data4) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Data4) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Data4) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *Data4) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *Data4) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *Data4) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetMessage

`func (o *Data4) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Data4) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Data4) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetParam

`func (o *Data4) GetParam() string`

GetParam returns the Param field if non-nil, zero value otherwise.

### GetParamOk

`func (o *Data4) GetParamOk() (*string, bool)`

GetParamOk returns a tuple with the Param field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParam

`func (o *Data4) SetParam(v string)`

SetParam sets Param field to given value.

### HasParam

`func (o *Data4) HasParam() bool`

HasParam returns a boolean if a field has been set.

### SetParamNil

`func (o *Data4) SetParamNil(b bool)`

 SetParamNil sets the value for Param to be an explicit nil

### UnsetParam
`func (o *Data4) UnsetParam()`

UnsetParam ensures that no value is present for Param, not even an explicit nil
### GetItem

`func (o *Data4) GetItem() Item`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *Data4) GetItemOk() (*Item, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *Data4) SetItem(v Item)`

SetItem sets Item field to given value.


### GetOutputIndex

`func (o *Data4) GetOutputIndex() int32`

GetOutputIndex returns the OutputIndex field if non-nil, zero value otherwise.

### GetOutputIndexOk

`func (o *Data4) GetOutputIndexOk() (*int32, bool)`

GetOutputIndexOk returns a tuple with the OutputIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputIndex

`func (o *Data4) SetOutputIndex(v int32)`

SetOutputIndex sets OutputIndex field to given value.


### GetContentIndex

`func (o *Data4) GetContentIndex() int32`

GetContentIndex returns the ContentIndex field if non-nil, zero value otherwise.

### GetContentIndexOk

`func (o *Data4) GetContentIndexOk() (*int32, bool)`

GetContentIndexOk returns a tuple with the ContentIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentIndex

`func (o *Data4) SetContentIndex(v int32)`

SetContentIndex sets ContentIndex field to given value.


### GetItemId

`func (o *Data4) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *Data4) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *Data4) SetItemId(v string)`

SetItemId sets ItemId field to given value.


### GetPart

`func (o *Data4) GetPart() Part`

GetPart returns the Part field if non-nil, zero value otherwise.

### GetPartOk

`func (o *Data4) GetPartOk() (*Part, bool)`

GetPartOk returns a tuple with the Part field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPart

`func (o *Data4) SetPart(v Part)`

SetPart sets Part field to given value.


### GetDelta

`func (o *Data4) GetDelta() string`

GetDelta returns the Delta field if non-nil, zero value otherwise.

### GetDeltaOk

`func (o *Data4) GetDeltaOk() (*string, bool)`

GetDeltaOk returns a tuple with the Delta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelta

`func (o *Data4) SetDelta(v string)`

SetDelta sets Delta field to given value.


### GetLogprobs

`func (o *Data4) GetLogprobs() []ResponsesTextDoneEventLogprob`

GetLogprobs returns the Logprobs field if non-nil, zero value otherwise.

### GetLogprobsOk

`func (o *Data4) GetLogprobsOk() (*[]ResponsesTextDoneEventLogprob, bool)`

GetLogprobsOk returns a tuple with the Logprobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogprobs

`func (o *Data4) SetLogprobs(v []ResponsesTextDoneEventLogprob)`

SetLogprobs sets Logprobs field to given value.


### GetText

`func (o *Data4) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *Data4) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *Data4) SetText(v string)`

SetText sets Text field to given value.


### GetArguments

`func (o *Data4) GetArguments() string`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *Data4) GetArgumentsOk() (*string, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *Data4) SetArguments(v string)`

SetArguments sets Arguments field to given value.


### GetName

`func (o *Data4) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Data4) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Data4) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


