# ResourceChunk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "resource"]
**Uri** | **string** |  | 
**Content** | [**Content4**](Content4.md) |  | 

## Methods

### NewResourceChunk

`func NewResourceChunk(uri string, content Content4, ) *ResourceChunk`

NewResourceChunk instantiates a new ResourceChunk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceChunkWithDefaults

`func NewResourceChunkWithDefaults() *ResourceChunk`

NewResourceChunkWithDefaults instantiates a new ResourceChunk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ResourceChunk) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResourceChunk) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResourceChunk) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ResourceChunk) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUri

`func (o *ResourceChunk) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ResourceChunk) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ResourceChunk) SetUri(v string)`

SetUri sets Uri field to given value.


### GetContent

`func (o *ResourceChunk) GetContent() Content4`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ResourceChunk) GetContentOk() (*Content4, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ResourceChunk) SetContent(v Content4)`

SetContent sets Content field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


