# ReadRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceId** | **string** |  | 
**StartOffset** | Pointer to **NullableInt32** |  | [optional] 
**EndOffset** | Pointer to **NullableInt32** |  | [optional] 
**TopK** | Pointer to **int32** |  | [optional] [default to 20]
**ContentType** | Pointer to **string** |  | [optional] [default to "content"]

## Methods

### NewReadRequest

`func NewReadRequest(sourceId string, ) *ReadRequest`

NewReadRequest instantiates a new ReadRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadRequestWithDefaults

`func NewReadRequestWithDefaults() *ReadRequest`

NewReadRequestWithDefaults instantiates a new ReadRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceId

`func (o *ReadRequest) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *ReadRequest) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *ReadRequest) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetStartOffset

`func (o *ReadRequest) GetStartOffset() int32`

GetStartOffset returns the StartOffset field if non-nil, zero value otherwise.

### GetStartOffsetOk

`func (o *ReadRequest) GetStartOffsetOk() (*int32, bool)`

GetStartOffsetOk returns a tuple with the StartOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartOffset

`func (o *ReadRequest) SetStartOffset(v int32)`

SetStartOffset sets StartOffset field to given value.

### HasStartOffset

`func (o *ReadRequest) HasStartOffset() bool`

HasStartOffset returns a boolean if a field has been set.

### SetStartOffsetNil

`func (o *ReadRequest) SetStartOffsetNil(b bool)`

 SetStartOffsetNil sets the value for StartOffset to be an explicit nil

### UnsetStartOffset
`func (o *ReadRequest) UnsetStartOffset()`

UnsetStartOffset ensures that no value is present for StartOffset, not even an explicit nil
### GetEndOffset

`func (o *ReadRequest) GetEndOffset() int32`

GetEndOffset returns the EndOffset field if non-nil, zero value otherwise.

### GetEndOffsetOk

`func (o *ReadRequest) GetEndOffsetOk() (*int32, bool)`

GetEndOffsetOk returns a tuple with the EndOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndOffset

`func (o *ReadRequest) SetEndOffset(v int32)`

SetEndOffset sets EndOffset field to given value.

### HasEndOffset

`func (o *ReadRequest) HasEndOffset() bool`

HasEndOffset returns a boolean if a field has been set.

### SetEndOffsetNil

`func (o *ReadRequest) SetEndOffsetNil(b bool)`

 SetEndOffsetNil sets the value for EndOffset to be an explicit nil

### UnsetEndOffset
`func (o *ReadRequest) UnsetEndOffset()`

UnsetEndOffset ensures that no value is present for EndOffset, not even an explicit nil
### GetTopK

`func (o *ReadRequest) GetTopK() int32`

GetTopK returns the TopK field if non-nil, zero value otherwise.

### GetTopKOk

`func (o *ReadRequest) GetTopKOk() (*int32, bool)`

GetTopKOk returns a tuple with the TopK field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopK

`func (o *ReadRequest) SetTopK(v int32)`

SetTopK sets TopK field to given value.

### HasTopK

`func (o *ReadRequest) HasTopK() bool`

HasTopK returns a boolean if a field has been set.

### GetContentType

`func (o *ReadRequest) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *ReadRequest) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *ReadRequest) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *ReadRequest) HasContentType() bool`

HasContentType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


