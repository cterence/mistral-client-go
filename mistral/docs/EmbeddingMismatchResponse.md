# EmbeddingMismatchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "embedding_mismatch"]
**Expected** | **int32** |  | 
**Received** | **int32** |  | 

## Methods

### NewEmbeddingMismatchResponse

`func NewEmbeddingMismatchResponse(expected int32, received int32, ) *EmbeddingMismatchResponse`

NewEmbeddingMismatchResponse instantiates a new EmbeddingMismatchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmbeddingMismatchResponseWithDefaults

`func NewEmbeddingMismatchResponseWithDefaults() *EmbeddingMismatchResponse`

NewEmbeddingMismatchResponseWithDefaults instantiates a new EmbeddingMismatchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EmbeddingMismatchResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EmbeddingMismatchResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EmbeddingMismatchResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EmbeddingMismatchResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetExpected

`func (o *EmbeddingMismatchResponse) GetExpected() int32`

GetExpected returns the Expected field if non-nil, zero value otherwise.

### GetExpectedOk

`func (o *EmbeddingMismatchResponse) GetExpectedOk() (*int32, bool)`

GetExpectedOk returns a tuple with the Expected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpected

`func (o *EmbeddingMismatchResponse) SetExpected(v int32)`

SetExpected sets Expected field to given value.


### GetReceived

`func (o *EmbeddingMismatchResponse) GetReceived() int32`

GetReceived returns the Received field if non-nil, zero value otherwise.

### GetReceivedOk

`func (o *EmbeddingMismatchResponse) GetReceivedOk() (*int32, bool)`

GetReceivedOk returns a tuple with the Received field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceived

`func (o *EmbeddingMismatchResponse) SetReceived(v int32)`

SetReceived sets Received field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


