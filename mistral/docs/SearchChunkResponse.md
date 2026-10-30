# SearchChunkResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**SourceId** | **string** |  | 
**Locator** | **string** |  | 
**StartOffset** | **NullableInt32** |  | 
**EndOffset** | **NullableInt32** |  | 
**ChunkType** | **string** |  | 
**ParentRef** | Pointer to **NullableString** |  | [optional] 
**Content** | **string** |  | 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewSearchChunkResponse

`func NewSearchChunkResponse(id string, sourceId string, locator string, startOffset NullableInt32, endOffset NullableInt32, chunkType string, content string, ) *SearchChunkResponse`

NewSearchChunkResponse instantiates a new SearchChunkResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchChunkResponseWithDefaults

`func NewSearchChunkResponseWithDefaults() *SearchChunkResponse`

NewSearchChunkResponseWithDefaults instantiates a new SearchChunkResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SearchChunkResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SearchChunkResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SearchChunkResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSourceId

`func (o *SearchChunkResponse) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *SearchChunkResponse) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *SearchChunkResponse) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetLocator

`func (o *SearchChunkResponse) GetLocator() string`

GetLocator returns the Locator field if non-nil, zero value otherwise.

### GetLocatorOk

`func (o *SearchChunkResponse) GetLocatorOk() (*string, bool)`

GetLocatorOk returns a tuple with the Locator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocator

`func (o *SearchChunkResponse) SetLocator(v string)`

SetLocator sets Locator field to given value.


### GetStartOffset

`func (o *SearchChunkResponse) GetStartOffset() int32`

GetStartOffset returns the StartOffset field if non-nil, zero value otherwise.

### GetStartOffsetOk

`func (o *SearchChunkResponse) GetStartOffsetOk() (*int32, bool)`

GetStartOffsetOk returns a tuple with the StartOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartOffset

`func (o *SearchChunkResponse) SetStartOffset(v int32)`

SetStartOffset sets StartOffset field to given value.


### SetStartOffsetNil

`func (o *SearchChunkResponse) SetStartOffsetNil(b bool)`

 SetStartOffsetNil sets the value for StartOffset to be an explicit nil

### UnsetStartOffset
`func (o *SearchChunkResponse) UnsetStartOffset()`

UnsetStartOffset ensures that no value is present for StartOffset, not even an explicit nil
### GetEndOffset

`func (o *SearchChunkResponse) GetEndOffset() int32`

GetEndOffset returns the EndOffset field if non-nil, zero value otherwise.

### GetEndOffsetOk

`func (o *SearchChunkResponse) GetEndOffsetOk() (*int32, bool)`

GetEndOffsetOk returns a tuple with the EndOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndOffset

`func (o *SearchChunkResponse) SetEndOffset(v int32)`

SetEndOffset sets EndOffset field to given value.


### SetEndOffsetNil

`func (o *SearchChunkResponse) SetEndOffsetNil(b bool)`

 SetEndOffsetNil sets the value for EndOffset to be an explicit nil

### UnsetEndOffset
`func (o *SearchChunkResponse) UnsetEndOffset()`

UnsetEndOffset ensures that no value is present for EndOffset, not even an explicit nil
### GetChunkType

`func (o *SearchChunkResponse) GetChunkType() string`

GetChunkType returns the ChunkType field if non-nil, zero value otherwise.

### GetChunkTypeOk

`func (o *SearchChunkResponse) GetChunkTypeOk() (*string, bool)`

GetChunkTypeOk returns a tuple with the ChunkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChunkType

`func (o *SearchChunkResponse) SetChunkType(v string)`

SetChunkType sets ChunkType field to given value.


### GetParentRef

`func (o *SearchChunkResponse) GetParentRef() string`

GetParentRef returns the ParentRef field if non-nil, zero value otherwise.

### GetParentRefOk

`func (o *SearchChunkResponse) GetParentRefOk() (*string, bool)`

GetParentRefOk returns a tuple with the ParentRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentRef

`func (o *SearchChunkResponse) SetParentRef(v string)`

SetParentRef sets ParentRef field to given value.

### HasParentRef

`func (o *SearchChunkResponse) HasParentRef() bool`

HasParentRef returns a boolean if a field has been set.

### SetParentRefNil

`func (o *SearchChunkResponse) SetParentRefNil(b bool)`

 SetParentRefNil sets the value for ParentRef to be an explicit nil

### UnsetParentRef
`func (o *SearchChunkResponse) UnsetParentRef()`

UnsetParentRef ensures that no value is present for ParentRef, not even an explicit nil
### GetContent

`func (o *SearchChunkResponse) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *SearchChunkResponse) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *SearchChunkResponse) SetContent(v string)`

SetContent sets Content field to given value.


### GetMetadata

`func (o *SearchChunkResponse) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SearchChunkResponse) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SearchChunkResponse) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SearchChunkResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


