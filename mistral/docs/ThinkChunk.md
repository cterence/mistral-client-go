# ThinkChunk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "thinking"]
**Thinking** | [**[]ThinkingInner**](ThinkingInner.md) |  | 
**Closed** | Pointer to **bool** | Whether the thinking chunk is closed or not. Currently only used for prefixing. | [optional] [default to true]

## Methods

### NewThinkChunk

`func NewThinkChunk(thinking []ThinkingInner, ) *ThinkChunk`

NewThinkChunk instantiates a new ThinkChunk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThinkChunkWithDefaults

`func NewThinkChunkWithDefaults() *ThinkChunk`

NewThinkChunkWithDefaults instantiates a new ThinkChunk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ThinkChunk) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ThinkChunk) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ThinkChunk) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ThinkChunk) HasType() bool`

HasType returns a boolean if a field has been set.

### GetThinking

`func (o *ThinkChunk) GetThinking() []ThinkingInner`

GetThinking returns the Thinking field if non-nil, zero value otherwise.

### GetThinkingOk

`func (o *ThinkChunk) GetThinkingOk() (*[]ThinkingInner, bool)`

GetThinkingOk returns a tuple with the Thinking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThinking

`func (o *ThinkChunk) SetThinking(v []ThinkingInner)`

SetThinking sets Thinking field to given value.


### GetClosed

`func (o *ThinkChunk) GetClosed() bool`

GetClosed returns the Closed field if non-nil, zero value otherwise.

### GetClosedOk

`func (o *ThinkChunk) GetClosedOk() (*bool, bool)`

GetClosedOk returns a tuple with the Closed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosed

`func (o *ThinkChunk) SetClosed(v bool)`

SetClosed sets Closed field to given value.

### HasClosed

`func (o *ThinkChunk) HasClosed() bool`

HasClosed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


