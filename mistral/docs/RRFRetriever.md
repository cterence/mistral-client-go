# RRFRetriever

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TopK** | Pointer to **int32** |  | [optional] [default to 20]
**Type** | Pointer to **string** |  | [optional] [default to "rrf"]
**Retrievers** | [**[]RetrieversInner**](RetrieversInner.md) |  | 
**Weights** | Pointer to **[]float32** |  | [optional] 
**RankConstant** | Pointer to **int32** |  | [optional] [default to 60]
**Filter** | Pointer to [**NullableKeywordRetrieverFilter**](KeywordRetrieverFilter.md) |  | [optional] 

## Methods

### NewRRFRetriever

`func NewRRFRetriever(retrievers []RetrieversInner, ) *RRFRetriever`

NewRRFRetriever instantiates a new RRFRetriever object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRRFRetrieverWithDefaults

`func NewRRFRetrieverWithDefaults() *RRFRetriever`

NewRRFRetrieverWithDefaults instantiates a new RRFRetriever object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTopK

`func (o *RRFRetriever) GetTopK() int32`

GetTopK returns the TopK field if non-nil, zero value otherwise.

### GetTopKOk

`func (o *RRFRetriever) GetTopKOk() (*int32, bool)`

GetTopKOk returns a tuple with the TopK field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopK

`func (o *RRFRetriever) SetTopK(v int32)`

SetTopK sets TopK field to given value.

### HasTopK

`func (o *RRFRetriever) HasTopK() bool`

HasTopK returns a boolean if a field has been set.

### GetType

`func (o *RRFRetriever) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RRFRetriever) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RRFRetriever) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RRFRetriever) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRetrievers

`func (o *RRFRetriever) GetRetrievers() []RetrieversInner`

GetRetrievers returns the Retrievers field if non-nil, zero value otherwise.

### GetRetrieversOk

`func (o *RRFRetriever) GetRetrieversOk() (*[]RetrieversInner, bool)`

GetRetrieversOk returns a tuple with the Retrievers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetrievers

`func (o *RRFRetriever) SetRetrievers(v []RetrieversInner)`

SetRetrievers sets Retrievers field to given value.


### GetWeights

`func (o *RRFRetriever) GetWeights() []float32`

GetWeights returns the Weights field if non-nil, zero value otherwise.

### GetWeightsOk

`func (o *RRFRetriever) GetWeightsOk() (*[]float32, bool)`

GetWeightsOk returns a tuple with the Weights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeights

`func (o *RRFRetriever) SetWeights(v []float32)`

SetWeights sets Weights field to given value.

### HasWeights

`func (o *RRFRetriever) HasWeights() bool`

HasWeights returns a boolean if a field has been set.

### SetWeightsNil

`func (o *RRFRetriever) SetWeightsNil(b bool)`

 SetWeightsNil sets the value for Weights to be an explicit nil

### UnsetWeights
`func (o *RRFRetriever) UnsetWeights()`

UnsetWeights ensures that no value is present for Weights, not even an explicit nil
### GetRankConstant

`func (o *RRFRetriever) GetRankConstant() int32`

GetRankConstant returns the RankConstant field if non-nil, zero value otherwise.

### GetRankConstantOk

`func (o *RRFRetriever) GetRankConstantOk() (*int32, bool)`

GetRankConstantOk returns a tuple with the RankConstant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRankConstant

`func (o *RRFRetriever) SetRankConstant(v int32)`

SetRankConstant sets RankConstant field to given value.

### HasRankConstant

`func (o *RRFRetriever) HasRankConstant() bool`

HasRankConstant returns a boolean if a field has been set.

### GetFilter

`func (o *RRFRetriever) GetFilter() KeywordRetrieverFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *RRFRetriever) GetFilterOk() (*KeywordRetrieverFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *RRFRetriever) SetFilter(v KeywordRetrieverFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *RRFRetriever) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### SetFilterNil

`func (o *RRFRetriever) SetFilterNil(b bool)`

 SetFilterNil sets the value for Filter to be an explicit nil

### UnsetFilter
`func (o *RRFRetriever) UnsetFilter()`

UnsetFilter ensures that no value is present for Filter, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


