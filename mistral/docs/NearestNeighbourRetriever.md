# NearestNeighbourRetriever

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TopK** | Pointer to **int32** |  | [optional] [default to 20]
**Type** | Pointer to **string** |  | [optional] [default to "nearest_neighbour"]
**Query** | Pointer to **NullableString** |  | [optional] 
**QueryEmbedding** | Pointer to **[]float32** |  | [optional] 
**Field** | Pointer to **NullableString** |  | [optional] 
**Filter** | Pointer to [**NullableKeywordRetrieverFilter**](KeywordRetrieverFilter.md) |  | [optional] 
**MaxCandidates** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewNearestNeighbourRetriever

`func NewNearestNeighbourRetriever() *NearestNeighbourRetriever`

NewNearestNeighbourRetriever instantiates a new NearestNeighbourRetriever object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNearestNeighbourRetrieverWithDefaults

`func NewNearestNeighbourRetrieverWithDefaults() *NearestNeighbourRetriever`

NewNearestNeighbourRetrieverWithDefaults instantiates a new NearestNeighbourRetriever object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTopK

`func (o *NearestNeighbourRetriever) GetTopK() int32`

GetTopK returns the TopK field if non-nil, zero value otherwise.

### GetTopKOk

`func (o *NearestNeighbourRetriever) GetTopKOk() (*int32, bool)`

GetTopKOk returns a tuple with the TopK field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopK

`func (o *NearestNeighbourRetriever) SetTopK(v int32)`

SetTopK sets TopK field to given value.

### HasTopK

`func (o *NearestNeighbourRetriever) HasTopK() bool`

HasTopK returns a boolean if a field has been set.

### GetType

`func (o *NearestNeighbourRetriever) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NearestNeighbourRetriever) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NearestNeighbourRetriever) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NearestNeighbourRetriever) HasType() bool`

HasType returns a boolean if a field has been set.

### GetQuery

`func (o *NearestNeighbourRetriever) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *NearestNeighbourRetriever) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *NearestNeighbourRetriever) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *NearestNeighbourRetriever) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### SetQueryNil

`func (o *NearestNeighbourRetriever) SetQueryNil(b bool)`

 SetQueryNil sets the value for Query to be an explicit nil

### UnsetQuery
`func (o *NearestNeighbourRetriever) UnsetQuery()`

UnsetQuery ensures that no value is present for Query, not even an explicit nil
### GetQueryEmbedding

`func (o *NearestNeighbourRetriever) GetQueryEmbedding() []float32`

GetQueryEmbedding returns the QueryEmbedding field if non-nil, zero value otherwise.

### GetQueryEmbeddingOk

`func (o *NearestNeighbourRetriever) GetQueryEmbeddingOk() (*[]float32, bool)`

GetQueryEmbeddingOk returns a tuple with the QueryEmbedding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryEmbedding

`func (o *NearestNeighbourRetriever) SetQueryEmbedding(v []float32)`

SetQueryEmbedding sets QueryEmbedding field to given value.

### HasQueryEmbedding

`func (o *NearestNeighbourRetriever) HasQueryEmbedding() bool`

HasQueryEmbedding returns a boolean if a field has been set.

### SetQueryEmbeddingNil

`func (o *NearestNeighbourRetriever) SetQueryEmbeddingNil(b bool)`

 SetQueryEmbeddingNil sets the value for QueryEmbedding to be an explicit nil

### UnsetQueryEmbedding
`func (o *NearestNeighbourRetriever) UnsetQueryEmbedding()`

UnsetQueryEmbedding ensures that no value is present for QueryEmbedding, not even an explicit nil
### GetField

`func (o *NearestNeighbourRetriever) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *NearestNeighbourRetriever) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *NearestNeighbourRetriever) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *NearestNeighbourRetriever) HasField() bool`

HasField returns a boolean if a field has been set.

### SetFieldNil

`func (o *NearestNeighbourRetriever) SetFieldNil(b bool)`

 SetFieldNil sets the value for Field to be an explicit nil

### UnsetField
`func (o *NearestNeighbourRetriever) UnsetField()`

UnsetField ensures that no value is present for Field, not even an explicit nil
### GetFilter

`func (o *NearestNeighbourRetriever) GetFilter() KeywordRetrieverFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *NearestNeighbourRetriever) GetFilterOk() (*KeywordRetrieverFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *NearestNeighbourRetriever) SetFilter(v KeywordRetrieverFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *NearestNeighbourRetriever) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### SetFilterNil

`func (o *NearestNeighbourRetriever) SetFilterNil(b bool)`

 SetFilterNil sets the value for Filter to be an explicit nil

### UnsetFilter
`func (o *NearestNeighbourRetriever) UnsetFilter()`

UnsetFilter ensures that no value is present for Filter, not even an explicit nil
### GetMaxCandidates

`func (o *NearestNeighbourRetriever) GetMaxCandidates() int32`

GetMaxCandidates returns the MaxCandidates field if non-nil, zero value otherwise.

### GetMaxCandidatesOk

`func (o *NearestNeighbourRetriever) GetMaxCandidatesOk() (*int32, bool)`

GetMaxCandidatesOk returns a tuple with the MaxCandidates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCandidates

`func (o *NearestNeighbourRetriever) SetMaxCandidates(v int32)`

SetMaxCandidates sets MaxCandidates field to given value.

### HasMaxCandidates

`func (o *NearestNeighbourRetriever) HasMaxCandidates() bool`

HasMaxCandidates returns a boolean if a field has been set.

### SetMaxCandidatesNil

`func (o *NearestNeighbourRetriever) SetMaxCandidatesNil(b bool)`

 SetMaxCandidatesNil sets the value for MaxCandidates to be an explicit nil

### UnsetMaxCandidates
`func (o *NearestNeighbourRetriever) UnsetMaxCandidates()`

UnsetMaxCandidates ensures that no value is present for MaxCandidates, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


