# And

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "and"]
**Matches** | [**[]MatchesInner**](MatchesInner.md) |  | 

## Methods

### NewAnd

`func NewAnd(matches []MatchesInner, ) *And`

NewAnd instantiates a new And object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAndWithDefaults

`func NewAndWithDefaults() *And`

NewAndWithDefaults instantiates a new And object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *And) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *And) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *And) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *And) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMatches

`func (o *And) GetMatches() []MatchesInner`

GetMatches returns the Matches field if non-nil, zero value otherwise.

### GetMatchesOk

`func (o *And) GetMatchesOk() (*[]MatchesInner, bool)`

GetMatchesOk returns a tuple with the Matches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatches

`func (o *And) SetMatches(v []MatchesInner)`

SetMatches sets Matches field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


