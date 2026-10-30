# RetrieversInner

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

## Methods

### NewRetrieversInner

`func NewRetrieversInner(query string, ) *RetrieversInner`

NewRetrieversInner instantiates a new RetrieversInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRetrieversInnerWithDefaults

`func NewRetrieversInnerWithDefaults() *RetrieversInner`

NewRetrieversInnerWithDefaults instantiates a new RetrieversInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTopK

`func (o *RetrieversInner) GetTopK() int32`

GetTopK returns the TopK field if non-nil, zero value otherwise.

### GetTopKOk

`func (o *RetrieversInner) GetTopKOk() (*int32, bool)`

GetTopKOk returns a tuple with the TopK field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopK

`func (o *RetrieversInner) SetTopK(v int32)`

SetTopK sets TopK field to given value.

### HasTopK

`func (o *RetrieversInner) HasTopK() bool`

HasTopK returns a boolean if a field has been set.

### GetType

`func (o *RetrieversInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RetrieversInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RetrieversInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RetrieversInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetQuery

`func (o *RetrieversInner) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *RetrieversInner) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *RetrieversInner) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetFilter

`func (o *RetrieversInner) GetFilter() KeywordRetrieverFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *RetrieversInner) GetFilterOk() (*KeywordRetrieverFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *RetrieversInner) SetFilter(v KeywordRetrieverFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *RetrieversInner) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### SetFilterNil

`func (o *RetrieversInner) SetFilterNil(b bool)`

 SetFilterNil sets the value for Filter to be an explicit nil

### UnsetFilter
`func (o *RetrieversInner) UnsetFilter()`

UnsetFilter ensures that no value is present for Filter, not even an explicit nil
### GetField

`func (o *RetrieversInner) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *RetrieversInner) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *RetrieversInner) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *RetrieversInner) HasField() bool`

HasField returns a boolean if a field has been set.

### GetQueryEmbedding

`func (o *RetrieversInner) GetQueryEmbedding() []float32`

GetQueryEmbedding returns the QueryEmbedding field if non-nil, zero value otherwise.

### GetQueryEmbeddingOk

`func (o *RetrieversInner) GetQueryEmbeddingOk() (*[]float32, bool)`

GetQueryEmbeddingOk returns a tuple with the QueryEmbedding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryEmbedding

`func (o *RetrieversInner) SetQueryEmbedding(v []float32)`

SetQueryEmbedding sets QueryEmbedding field to given value.

### HasQueryEmbedding

`func (o *RetrieversInner) HasQueryEmbedding() bool`

HasQueryEmbedding returns a boolean if a field has been set.

### GetMaxCandidates

`func (o *RetrieversInner) GetMaxCandidates() int32`

GetMaxCandidates returns the MaxCandidates field if non-nil, zero value otherwise.

### GetMaxCandidatesOk

`func (o *RetrieversInner) GetMaxCandidatesOk() (*int32, bool)`

GetMaxCandidatesOk returns a tuple with the MaxCandidates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCandidates

`func (o *RetrieversInner) SetMaxCandidates(v int32)`

SetMaxCandidates sets MaxCandidates field to given value.

### HasMaxCandidates

`func (o *RetrieversInner) HasMaxCandidates() bool`

HasMaxCandidates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


