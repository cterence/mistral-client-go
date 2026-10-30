# NavigationChunkResponse

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

### NewNavigationChunkResponse

`func NewNavigationChunkResponse(id string, sourceId string, locator string, startOffset NullableInt32, endOffset NullableInt32, chunkType string, content string, ) *NavigationChunkResponse`

NewNavigationChunkResponse instantiates a new NavigationChunkResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNavigationChunkResponseWithDefaults

`func NewNavigationChunkResponseWithDefaults() *NavigationChunkResponse`

NewNavigationChunkResponseWithDefaults instantiates a new NavigationChunkResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NavigationChunkResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NavigationChunkResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NavigationChunkResponse) SetId(v string)`

SetId sets Id field to given value.


### GetSourceId

`func (o *NavigationChunkResponse) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *NavigationChunkResponse) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *NavigationChunkResponse) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetLocator

`func (o *NavigationChunkResponse) GetLocator() string`

GetLocator returns the Locator field if non-nil, zero value otherwise.

### GetLocatorOk

`func (o *NavigationChunkResponse) GetLocatorOk() (*string, bool)`

GetLocatorOk returns a tuple with the Locator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocator

`func (o *NavigationChunkResponse) SetLocator(v string)`

SetLocator sets Locator field to given value.


### GetStartOffset

`func (o *NavigationChunkResponse) GetStartOffset() int32`

GetStartOffset returns the StartOffset field if non-nil, zero value otherwise.

### GetStartOffsetOk

`func (o *NavigationChunkResponse) GetStartOffsetOk() (*int32, bool)`

GetStartOffsetOk returns a tuple with the StartOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartOffset

`func (o *NavigationChunkResponse) SetStartOffset(v int32)`

SetStartOffset sets StartOffset field to given value.


### SetStartOffsetNil

`func (o *NavigationChunkResponse) SetStartOffsetNil(b bool)`

 SetStartOffsetNil sets the value for StartOffset to be an explicit nil

### UnsetStartOffset
`func (o *NavigationChunkResponse) UnsetStartOffset()`

UnsetStartOffset ensures that no value is present for StartOffset, not even an explicit nil
### GetEndOffset

`func (o *NavigationChunkResponse) GetEndOffset() int32`

GetEndOffset returns the EndOffset field if non-nil, zero value otherwise.

### GetEndOffsetOk

`func (o *NavigationChunkResponse) GetEndOffsetOk() (*int32, bool)`

GetEndOffsetOk returns a tuple with the EndOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndOffset

`func (o *NavigationChunkResponse) SetEndOffset(v int32)`

SetEndOffset sets EndOffset field to given value.


### SetEndOffsetNil

`func (o *NavigationChunkResponse) SetEndOffsetNil(b bool)`

 SetEndOffsetNil sets the value for EndOffset to be an explicit nil

### UnsetEndOffset
`func (o *NavigationChunkResponse) UnsetEndOffset()`

UnsetEndOffset ensures that no value is present for EndOffset, not even an explicit nil
### GetChunkType

`func (o *NavigationChunkResponse) GetChunkType() string`

GetChunkType returns the ChunkType field if non-nil, zero value otherwise.

### GetChunkTypeOk

`func (o *NavigationChunkResponse) GetChunkTypeOk() (*string, bool)`

GetChunkTypeOk returns a tuple with the ChunkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChunkType

`func (o *NavigationChunkResponse) SetChunkType(v string)`

SetChunkType sets ChunkType field to given value.


### GetParentRef

`func (o *NavigationChunkResponse) GetParentRef() string`

GetParentRef returns the ParentRef field if non-nil, zero value otherwise.

### GetParentRefOk

`func (o *NavigationChunkResponse) GetParentRefOk() (*string, bool)`

GetParentRefOk returns a tuple with the ParentRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentRef

`func (o *NavigationChunkResponse) SetParentRef(v string)`

SetParentRef sets ParentRef field to given value.

### HasParentRef

`func (o *NavigationChunkResponse) HasParentRef() bool`

HasParentRef returns a boolean if a field has been set.

### SetParentRefNil

`func (o *NavigationChunkResponse) SetParentRefNil(b bool)`

 SetParentRefNil sets the value for ParentRef to be an explicit nil

### UnsetParentRef
`func (o *NavigationChunkResponse) UnsetParentRef()`

UnsetParentRef ensures that no value is present for ParentRef, not even an explicit nil
### GetContent

`func (o *NavigationChunkResponse) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *NavigationChunkResponse) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *NavigationChunkResponse) SetContent(v string)`

SetContent sets Content field to given value.


### GetMetadata

`func (o *NavigationChunkResponse) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NavigationChunkResponse) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NavigationChunkResponse) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *NavigationChunkResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


