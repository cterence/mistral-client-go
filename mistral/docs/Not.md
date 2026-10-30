# Not

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "not"]
**Match** | [**Match**](Match.md) |  | 

## Methods

### NewNot

`func NewNot(match Match, ) *Not`

NewNot instantiates a new Not object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotWithDefaults

`func NewNotWithDefaults() *Not`

NewNotWithDefaults instantiates a new Not object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Not) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Not) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Not) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Not) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMatch

`func (o *Not) GetMatch() Match`

GetMatch returns the Match field if non-nil, zero value otherwise.

### GetMatchOk

`func (o *Not) GetMatchOk() (*Match, bool)`

GetMatchOk returns a tuple with the Match field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatch

`func (o *Not) SetMatch(v Match)`

SetMatch sets Match field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


