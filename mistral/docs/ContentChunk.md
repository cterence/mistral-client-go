# ContentChunk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "text"]
**Text** | **string** |  | 
**ImageUrl** | [**ImageUrl**](ImageUrl.md) |  | 
**DocumentUrl** | **string** |  | 
**DocumentName** | Pointer to **string** | The filename of the document | [optional] 
**ReferenceIds** | [**[]ReferenceIdsInner**](ReferenceIdsInner.md) |  | 
**FileId** | **string** |  | 
**Thinking** | [**[]ThinkingInner**](ThinkingInner.md) |  | 
**Signature** | Pointer to **string** | Signature to replay some reasoning blocks across turns. | [optional] 
**Closed** | Pointer to **bool** | Whether the thinking chunk is closed or not. Currently only used for prefixing. | [optional] [default to true]
**InputAudio** | [**InputAudio**](InputAudio.md) |  | 
**Uri** | **string** |  | 
**Content** | [**Content4**](Content4.md) |  | 
**Metadata** | Pointer to [**map[string]ResourceLinkChunkMetadataValue**](ResourceLinkChunkMetadataValue.md) |  | [optional] 

## Methods

### NewContentChunk

`func NewContentChunk(text string, imageUrl ImageUrl, documentUrl string, referenceIds []ReferenceIdsInner, fileId string, thinking []ThinkingInner, inputAudio InputAudio, uri string, content Content4, ) *ContentChunk`

NewContentChunk instantiates a new ContentChunk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentChunkWithDefaults

`func NewContentChunkWithDefaults() *ContentChunk`

NewContentChunkWithDefaults instantiates a new ContentChunk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ContentChunk) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContentChunk) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContentChunk) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ContentChunk) HasType() bool`

HasType returns a boolean if a field has been set.

### GetText

`func (o *ContentChunk) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ContentChunk) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ContentChunk) SetText(v string)`

SetText sets Text field to given value.


### GetImageUrl

`func (o *ContentChunk) GetImageUrl() ImageUrl`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *ContentChunk) GetImageUrlOk() (*ImageUrl, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *ContentChunk) SetImageUrl(v ImageUrl)`

SetImageUrl sets ImageUrl field to given value.


### GetDocumentUrl

`func (o *ContentChunk) GetDocumentUrl() string`

GetDocumentUrl returns the DocumentUrl field if non-nil, zero value otherwise.

### GetDocumentUrlOk

`func (o *ContentChunk) GetDocumentUrlOk() (*string, bool)`

GetDocumentUrlOk returns a tuple with the DocumentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentUrl

`func (o *ContentChunk) SetDocumentUrl(v string)`

SetDocumentUrl sets DocumentUrl field to given value.


### GetDocumentName

`func (o *ContentChunk) GetDocumentName() string`

GetDocumentName returns the DocumentName field if non-nil, zero value otherwise.

### GetDocumentNameOk

`func (o *ContentChunk) GetDocumentNameOk() (*string, bool)`

GetDocumentNameOk returns a tuple with the DocumentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentName

`func (o *ContentChunk) SetDocumentName(v string)`

SetDocumentName sets DocumentName field to given value.

### HasDocumentName

`func (o *ContentChunk) HasDocumentName() bool`

HasDocumentName returns a boolean if a field has been set.

### GetReferenceIds

`func (o *ContentChunk) GetReferenceIds() []ReferenceIdsInner`

GetReferenceIds returns the ReferenceIds field if non-nil, zero value otherwise.

### GetReferenceIdsOk

`func (o *ContentChunk) GetReferenceIdsOk() (*[]ReferenceIdsInner, bool)`

GetReferenceIdsOk returns a tuple with the ReferenceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceIds

`func (o *ContentChunk) SetReferenceIds(v []ReferenceIdsInner)`

SetReferenceIds sets ReferenceIds field to given value.


### GetFileId

`func (o *ContentChunk) GetFileId() string`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *ContentChunk) GetFileIdOk() (*string, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *ContentChunk) SetFileId(v string)`

SetFileId sets FileId field to given value.


### GetThinking

`func (o *ContentChunk) GetThinking() []ThinkingInner`

GetThinking returns the Thinking field if non-nil, zero value otherwise.

### GetThinkingOk

`func (o *ContentChunk) GetThinkingOk() (*[]ThinkingInner, bool)`

GetThinkingOk returns a tuple with the Thinking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThinking

`func (o *ContentChunk) SetThinking(v []ThinkingInner)`

SetThinking sets Thinking field to given value.


### GetSignature

`func (o *ContentChunk) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *ContentChunk) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *ContentChunk) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *ContentChunk) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetClosed

`func (o *ContentChunk) GetClosed() bool`

GetClosed returns the Closed field if non-nil, zero value otherwise.

### GetClosedOk

`func (o *ContentChunk) GetClosedOk() (*bool, bool)`

GetClosedOk returns a tuple with the Closed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosed

`func (o *ContentChunk) SetClosed(v bool)`

SetClosed sets Closed field to given value.

### HasClosed

`func (o *ContentChunk) HasClosed() bool`

HasClosed returns a boolean if a field has been set.

### GetInputAudio

`func (o *ContentChunk) GetInputAudio() InputAudio`

GetInputAudio returns the InputAudio field if non-nil, zero value otherwise.

### GetInputAudioOk

`func (o *ContentChunk) GetInputAudioOk() (*InputAudio, bool)`

GetInputAudioOk returns a tuple with the InputAudio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputAudio

`func (o *ContentChunk) SetInputAudio(v InputAudio)`

SetInputAudio sets InputAudio field to given value.


### GetUri

`func (o *ContentChunk) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ContentChunk) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ContentChunk) SetUri(v string)`

SetUri sets Uri field to given value.


### GetContent

`func (o *ContentChunk) GetContent() Content4`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ContentChunk) GetContentOk() (*Content4, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ContentChunk) SetContent(v Content4)`

SetContent sets Content field to given value.


### GetMetadata

`func (o *ContentChunk) GetMetadata() map[string]ResourceLinkChunkMetadataValue`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ContentChunk) GetMetadataOk() (*map[string]ResourceLinkChunkMetadataValue, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ContentChunk) SetMetadata(v map[string]ResourceLinkChunkMetadataValue)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ContentChunk) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


