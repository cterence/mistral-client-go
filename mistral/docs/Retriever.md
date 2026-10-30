# Retriever

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TopK** | Pointer to **int32** |  | [optional] [default to 20]
**Type** | Pointer to **string** |  | [optional] [default to "keyword"]
**Query** | **string** |  | 
**Filter** | Pointer to [**NullableKeywordRetrieverFilter**](KeywordRetrieverFilter.md) |  | [optional] 
**Field** | Pointer to **string** |  | [optional] 
**QueryEmbedding** | Pointer to **[]float32** |  | [optional] 
**MaxCandidates** | Pointer to **int32** |  | [optional] 
**Retrievers** | [**[]RetrieversInner**](RetrieversInner.md) |  | 
**Weights** | Pointer to **[]float32** |  | [optional] 
**RankConstant** | Pointer to **int32** |  | [optional] [default to 60]

## Methods

### NewRetriever

`func NewRetriever(query string, retrievers []RetrieversInner, ) *Retriever`

NewRetriever instantiates a new Retriever object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRetrieverWithDefaults

`func NewRetrieverWithDefaults() *Retriever`

NewRetrieverWithDefaults instantiates a new Retriever object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTopK

`func (o *Retriever) GetTopK() int32`

GetTopK returns the TopK field if non-nil, zero value otherwise.

### GetTopKOk

`func (o *Retriever) GetTopKOk() (*int32, bool)`

GetTopKOk returns a tuple with the TopK field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopK

`func (o *Retriever) SetTopK(v int32)`

SetTopK sets TopK field to given value.

### HasTopK

`func (o *Retriever) HasTopK() bool`

HasTopK returns a boolean if a field has been set.

### GetType

`func (o *Retriever) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Retriever) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Retriever) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Retriever) HasType() bool`

HasType returns a boolean if a field has been set.

### GetQuery

`func (o *Retriever) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *Retriever) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *Retriever) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetFilter

`func (o *Retriever) GetFilter() KeywordRetrieverFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *Retriever) GetFilterOk() (*KeywordRetrieverFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *Retriever) SetFilter(v KeywordRetrieverFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *Retriever) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### SetFilterNil

`func (o *Retriever) SetFilterNil(b bool)`

 SetFilterNil sets the value for Filter to be an explicit nil

### UnsetFilter
`func (o *Retriever) UnsetFilter()`

UnsetFilter ensures that no value is present for Filter, not even an explicit nil
### GetField

`func (o *Retriever) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *Retriever) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *Retriever) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *Retriever) HasField() bool`

HasField returns a boolean if a field has been set.

### GetQueryEmbedding

`func (o *Retriever) GetQueryEmbedding() []float32`

GetQueryEmbedding returns the QueryEmbedding field if non-nil, zero value otherwise.

### GetQueryEmbeddingOk

`func (o *Retriever) GetQueryEmbeddingOk() (*[]float32, bool)`

GetQueryEmbeddingOk returns a tuple with the QueryEmbedding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryEmbedding

`func (o *Retriever) SetQueryEmbedding(v []float32)`

SetQueryEmbedding sets QueryEmbedding field to given value.

### HasQueryEmbedding

`func (o *Retriever) HasQueryEmbedding() bool`

HasQueryEmbedding returns a boolean if a field has been set.

### GetMaxCandidates

`func (o *Retriever) GetMaxCandidates() int32`

GetMaxCandidates returns the MaxCandidates field if non-nil, zero value otherwise.

### GetMaxCandidatesOk

`func (o *Retriever) GetMaxCandidatesOk() (*int32, bool)`

GetMaxCandidatesOk returns a tuple with the MaxCandidates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCandidates

`func (o *Retriever) SetMaxCandidates(v int32)`

SetMaxCandidates sets MaxCandidates field to given value.

### HasMaxCandidates

`func (o *Retriever) HasMaxCandidates() bool`

HasMaxCandidates returns a boolean if a field has been set.

### GetRetrievers

`func (o *Retriever) GetRetrievers() []RetrieversInner`

GetRetrievers returns the Retrievers field if non-nil, zero value otherwise.

### GetRetrieversOk

`func (o *Retriever) GetRetrieversOk() (*[]RetrieversInner, bool)`

GetRetrieversOk returns a tuple with the Retrievers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetrievers

`func (o *Retriever) SetRetrievers(v []RetrieversInner)`

SetRetrievers sets Retrievers field to given value.


### GetWeights

`func (o *Retriever) GetWeights() []float32`

GetWeights returns the Weights field if non-nil, zero value otherwise.

### GetWeightsOk

`func (o *Retriever) GetWeightsOk() (*[]float32, bool)`

GetWeightsOk returns a tuple with the Weights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeights

`func (o *Retriever) SetWeights(v []float32)`

SetWeights sets Weights field to given value.

### HasWeights

`func (o *Retriever) HasWeights() bool`

HasWeights returns a boolean if a field has been set.

### GetRankConstant

`func (o *Retriever) GetRankConstant() int32`

GetRankConstant returns the RankConstant field if non-nil, zero value otherwise.

### GetRankConstantOk

`func (o *Retriever) GetRankConstantOk() (*int32, bool)`

GetRankConstantOk returns a tuple with the RankConstant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRankConstant

`func (o *Retriever) SetRankConstant(v int32)`

SetRankConstant sets RankConstant field to given value.

### HasRankConstant

`func (o *Retriever) HasRankConstant() bool`

HasRankConstant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


