# ResourceLinkChunk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "resource_link"]
**Uri** | **string** |  | 
**Metadata** | Pointer to [**map[string]ResourceLinkChunkMetadataValue**](ResourceLinkChunkMetadataValue.md) |  | [optional] 

## Methods

### NewResourceLinkChunk

`func NewResourceLinkChunk(uri string, ) *ResourceLinkChunk`

NewResourceLinkChunk instantiates a new ResourceLinkChunk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceLinkChunkWithDefaults

`func NewResourceLinkChunkWithDefaults() *ResourceLinkChunk`

NewResourceLinkChunkWithDefaults instantiates a new ResourceLinkChunk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ResourceLinkChunk) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResourceLinkChunk) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResourceLinkChunk) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ResourceLinkChunk) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUri

`func (o *ResourceLinkChunk) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ResourceLinkChunk) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ResourceLinkChunk) SetUri(v string)`

SetUri sets Uri field to given value.


### GetMetadata

`func (o *ResourceLinkChunk) GetMetadata() map[string]ResourceLinkChunkMetadataValue`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ResourceLinkChunk) GetMetadataOk() (*map[string]ResourceLinkChunkMetadataValue, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ResourceLinkChunk) SetMetadata(v map[string]ResourceLinkChunkMetadataValue)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ResourceLinkChunk) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *ResourceLinkChunk) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *ResourceLinkChunk) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


