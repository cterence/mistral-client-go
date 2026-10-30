# KeywordRetriever

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TopK** | Pointer to **int32** |  | [optional] [default to 20]
**Type** | Pointer to **string** |  | [optional] [default to "keyword"]
**Query** | **string** |  | 
**Filter** | Pointer to [**NullableKeywordRetrieverFilter**](KeywordRetrieverFilter.md) |  | [optional] 
**Field** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewKeywordRetriever

`func NewKeywordRetriever(query string, ) *KeywordRetriever`

NewKeywordRetriever instantiates a new KeywordRetriever object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeywordRetrieverWithDefaults

`func NewKeywordRetrieverWithDefaults() *KeywordRetriever`

NewKeywordRetrieverWithDefaults instantiates a new KeywordRetriever object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTopK

`func (o *KeywordRetriever) GetTopK() int32`

GetTopK returns the TopK field if non-nil, zero value otherwise.

### GetTopKOk

`func (o *KeywordRetriever) GetTopKOk() (*int32, bool)`

GetTopKOk returns a tuple with the TopK field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopK

`func (o *KeywordRetriever) SetTopK(v int32)`

SetTopK sets TopK field to given value.

### HasTopK

`func (o *KeywordRetriever) HasTopK() bool`

HasTopK returns a boolean if a field has been set.

### GetType

`func (o *KeywordRetriever) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KeywordRetriever) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KeywordRetriever) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *KeywordRetriever) HasType() bool`

HasType returns a boolean if a field has been set.

### GetQuery

`func (o *KeywordRetriever) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *KeywordRetriever) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *KeywordRetriever) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetFilter

`func (o *KeywordRetriever) GetFilter() KeywordRetrieverFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *KeywordRetriever) GetFilterOk() (*KeywordRetrieverFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *KeywordRetriever) SetFilter(v KeywordRetrieverFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *KeywordRetriever) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### SetFilterNil

`func (o *KeywordRetriever) SetFilterNil(b bool)`

 SetFilterNil sets the value for Filter to be an explicit nil

### UnsetFilter
`func (o *KeywordRetriever) UnsetFilter()`

UnsetFilter ensures that no value is present for Filter, not even an explicit nil
### GetField

`func (o *KeywordRetriever) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *KeywordRetriever) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *KeywordRetriever) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *KeywordRetriever) HasField() bool`

HasField returns a boolean if a field has been set.

### SetFieldNil

`func (o *KeywordRetriever) SetFieldNil(b bool)`

 SetFieldNil sets the value for Field to be an explicit nil

### UnsetField
`func (o *KeywordRetriever) UnsetField()`

UnsetField ensures that no value is present for Field, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


