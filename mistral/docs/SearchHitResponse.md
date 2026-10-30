# SearchHitResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chunk** | [**SearchChunkResponse**](SearchChunkResponse.md) |  | 
**Score** | **float32** |  | 
**Distance** | Pointer to **NullableFloat32** |  | [optional] 
**IndexName** | **string** |  | 

## Methods

### NewSearchHitResponse

`func NewSearchHitResponse(chunk SearchChunkResponse, score float32, indexName string, ) *SearchHitResponse`

NewSearchHitResponse instantiates a new SearchHitResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchHitResponseWithDefaults

`func NewSearchHitResponseWithDefaults() *SearchHitResponse`

NewSearchHitResponseWithDefaults instantiates a new SearchHitResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChunk

`func (o *SearchHitResponse) GetChunk() SearchChunkResponse`

GetChunk returns the Chunk field if non-nil, zero value otherwise.

### GetChunkOk

`func (o *SearchHitResponse) GetChunkOk() (*SearchChunkResponse, bool)`

GetChunkOk returns a tuple with the Chunk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChunk

`func (o *SearchHitResponse) SetChunk(v SearchChunkResponse)`

SetChunk sets Chunk field to given value.


### GetScore

`func (o *SearchHitResponse) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *SearchHitResponse) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *SearchHitResponse) SetScore(v float32)`

SetScore sets Score field to given value.


### GetDistance

`func (o *SearchHitResponse) GetDistance() float32`

GetDistance returns the Distance field if non-nil, zero value otherwise.

### GetDistanceOk

`func (o *SearchHitResponse) GetDistanceOk() (*float32, bool)`

GetDistanceOk returns a tuple with the Distance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistance

`func (o *SearchHitResponse) SetDistance(v float32)`

SetDistance sets Distance field to given value.

### HasDistance

`func (o *SearchHitResponse) HasDistance() bool`

HasDistance returns a boolean if a field has been set.

### SetDistanceNil

`func (o *SearchHitResponse) SetDistanceNil(b bool)`

 SetDistanceNil sets the value for Distance to be an explicit nil

### UnsetDistance
`func (o *SearchHitResponse) UnsetDistance()`

UnsetDistance ensures that no value is present for Distance, not even an explicit nil
### GetIndexName

`func (o *SearchHitResponse) GetIndexName() string`

GetIndexName returns the IndexName field if non-nil, zero value otherwise.

### GetIndexNameOk

`func (o *SearchHitResponse) GetIndexNameOk() (*string, bool)`

GetIndexNameOk returns a tuple with the IndexName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexName

`func (o *SearchHitResponse) SetIndexName(v string)`

SetIndexName sets IndexName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


