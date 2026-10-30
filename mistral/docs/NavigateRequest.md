# NavigateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceId** | **string** |  | 
**StartOffset** | **int32** |  | 
**EndOffset** | **int32** |  | 
**Direction** | [**NavigationDirection**](NavigationDirection.md) |  | 
**TopK** | Pointer to **int32** |  | [optional] [default to 1]
**ContentType** | Pointer to **string** |  | [optional] [default to "content"]

## Methods

### NewNavigateRequest

`func NewNavigateRequest(sourceId string, startOffset int32, endOffset int32, direction NavigationDirection, ) *NavigateRequest`

NewNavigateRequest instantiates a new NavigateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNavigateRequestWithDefaults

`func NewNavigateRequestWithDefaults() *NavigateRequest`

NewNavigateRequestWithDefaults instantiates a new NavigateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceId

`func (o *NavigateRequest) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *NavigateRequest) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *NavigateRequest) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetStartOffset

`func (o *NavigateRequest) GetStartOffset() int32`

GetStartOffset returns the StartOffset field if non-nil, zero value otherwise.

### GetStartOffsetOk

`func (o *NavigateRequest) GetStartOffsetOk() (*int32, bool)`

GetStartOffsetOk returns a tuple with the StartOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartOffset

`func (o *NavigateRequest) SetStartOffset(v int32)`

SetStartOffset sets StartOffset field to given value.


### GetEndOffset

`func (o *NavigateRequest) GetEndOffset() int32`

GetEndOffset returns the EndOffset field if non-nil, zero value otherwise.

### GetEndOffsetOk

`func (o *NavigateRequest) GetEndOffsetOk() (*int32, bool)`

GetEndOffsetOk returns a tuple with the EndOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndOffset

`func (o *NavigateRequest) SetEndOffset(v int32)`

SetEndOffset sets EndOffset field to given value.


### GetDirection

`func (o *NavigateRequest) GetDirection() NavigationDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *NavigateRequest) GetDirectionOk() (*NavigationDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *NavigateRequest) SetDirection(v NavigationDirection)`

SetDirection sets Direction field to given value.


### GetTopK

`func (o *NavigateRequest) GetTopK() int32`

GetTopK returns the TopK field if non-nil, zero value otherwise.

### GetTopKOk

`func (o *NavigateRequest) GetTopKOk() (*int32, bool)`

GetTopKOk returns a tuple with the TopK field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopK

`func (o *NavigateRequest) SetTopK(v int32)`

SetTopK sets TopK field to given value.

### HasTopK

`func (o *NavigateRequest) HasTopK() bool`

HasTopK returns a boolean if a field has been set.

### GetContentType

`func (o *NavigateRequest) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *NavigateRequest) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *NavigateRequest) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *NavigateRequest) HasContentType() bool`

HasContentType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


