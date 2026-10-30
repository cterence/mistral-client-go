# AcceptedDocumentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] [default to "accepted"]
**Position** | **int32** |  | 
**DocumentId** | **string** |  | 
**ChunkCount** | **int32** |  | 

## Methods

### NewAcceptedDocumentResponse

`func NewAcceptedDocumentResponse(position int32, documentId string, chunkCount int32, ) *AcceptedDocumentResponse`

NewAcceptedDocumentResponse instantiates a new AcceptedDocumentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcceptedDocumentResponseWithDefaults

`func NewAcceptedDocumentResponseWithDefaults() *AcceptedDocumentResponse`

NewAcceptedDocumentResponseWithDefaults instantiates a new AcceptedDocumentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AcceptedDocumentResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AcceptedDocumentResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AcceptedDocumentResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AcceptedDocumentResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPosition

`func (o *AcceptedDocumentResponse) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *AcceptedDocumentResponse) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *AcceptedDocumentResponse) SetPosition(v int32)`

SetPosition sets Position field to given value.


### GetDocumentId

`func (o *AcceptedDocumentResponse) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *AcceptedDocumentResponse) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *AcceptedDocumentResponse) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.


### GetChunkCount

`func (o *AcceptedDocumentResponse) GetChunkCount() int32`

GetChunkCount returns the ChunkCount field if non-nil, zero value otherwise.

### GetChunkCountOk

`func (o *AcceptedDocumentResponse) GetChunkCountOk() (*int32, bool)`

GetChunkCountOk returns a tuple with the ChunkCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChunkCount

`func (o *AcceptedDocumentResponse) SetChunkCount(v int32)`

SetChunkCount sets ChunkCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


