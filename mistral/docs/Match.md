# Match

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "and"]
**Matches** | [**[]MatchesInner**](MatchesInner.md) |  | 
**Match** | [**Match**](Match.md) |  | 
**Field** | **string** |  | 
**Value** | [**Value5**](Value5.md) |  | 
**Gt** | Pointer to [**NullableGt**](Gt.md) |  | [optional] 
**Gte** | Pointer to [**NullableGte**](Gte.md) |  | [optional] 
**Lt** | Pointer to [**NullableLt**](Lt.md) |  | [optional] 
**Lte** | Pointer to [**NullableLte**](Lte.md) |  | [optional] 
**Values** | [**[]ValuesInner**](ValuesInner.md) |  | 

## Methods

### NewMatch

`func NewMatch(matches []MatchesInner, match Match, field string, value Value5, values []ValuesInner, ) *Match`

NewMatch instantiates a new Match object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMatchWithDefaults

`func NewMatchWithDefaults() *Match`

NewMatchWithDefaults instantiates a new Match object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Match) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Match) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Match) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Match) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMatches

`func (o *Match) GetMatches() []MatchesInner`

GetMatches returns the Matches field if non-nil, zero value otherwise.

### GetMatchesOk

`func (o *Match) GetMatchesOk() (*[]MatchesInner, bool)`

GetMatchesOk returns a tuple with the Matches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatches

`func (o *Match) SetMatches(v []MatchesInner)`

SetMatches sets Matches field to given value.


### GetMatch

`func (o *Match) GetMatch() Match`

GetMatch returns the Match field if non-nil, zero value otherwise.

### GetMatchOk

`func (o *Match) GetMatchOk() (*Match, bool)`

GetMatchOk returns a tuple with the Match field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatch

`func (o *Match) SetMatch(v Match)`

SetMatch sets Match field to given value.


### GetField

`func (o *Match) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *Match) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *Match) SetField(v string)`

SetField sets Field field to given value.


### GetValue

`func (o *Match) GetValue() Value5`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Match) GetValueOk() (*Value5, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Match) SetValue(v Value5)`

SetValue sets Value field to given value.


### GetGt

`func (o *Match) GetGt() Gt`

GetGt returns the Gt field if non-nil, zero value otherwise.

### GetGtOk

`func (o *Match) GetGtOk() (*Gt, bool)`

GetGtOk returns a tuple with the Gt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGt

`func (o *Match) SetGt(v Gt)`

SetGt sets Gt field to given value.

### HasGt

`func (o *Match) HasGt() bool`

HasGt returns a boolean if a field has been set.

### SetGtNil

`func (o *Match) SetGtNil(b bool)`

 SetGtNil sets the value for Gt to be an explicit nil

### UnsetGt
`func (o *Match) UnsetGt()`

UnsetGt ensures that no value is present for Gt, not even an explicit nil
### GetGte

`func (o *Match) GetGte() Gte`

GetGte returns the Gte field if non-nil, zero value otherwise.

### GetGteOk

`func (o *Match) GetGteOk() (*Gte, bool)`

GetGteOk returns a tuple with the Gte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGte

`func (o *Match) SetGte(v Gte)`

SetGte sets Gte field to given value.

### HasGte

`func (o *Match) HasGte() bool`

HasGte returns a boolean if a field has been set.

### SetGteNil

`func (o *Match) SetGteNil(b bool)`

 SetGteNil sets the value for Gte to be an explicit nil

### UnsetGte
`func (o *Match) UnsetGte()`

UnsetGte ensures that no value is present for Gte, not even an explicit nil
### GetLt

`func (o *Match) GetLt() Lt`

GetLt returns the Lt field if non-nil, zero value otherwise.

### GetLtOk

`func (o *Match) GetLtOk() (*Lt, bool)`

GetLtOk returns a tuple with the Lt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLt

`func (o *Match) SetLt(v Lt)`

SetLt sets Lt field to given value.

### HasLt

`func (o *Match) HasLt() bool`

HasLt returns a boolean if a field has been set.

### SetLtNil

`func (o *Match) SetLtNil(b bool)`

 SetLtNil sets the value for Lt to be an explicit nil

### UnsetLt
`func (o *Match) UnsetLt()`

UnsetLt ensures that no value is present for Lt, not even an explicit nil
### GetLte

`func (o *Match) GetLte() Lte`

GetLte returns the Lte field if non-nil, zero value otherwise.

### GetLteOk

`func (o *Match) GetLteOk() (*Lte, bool)`

GetLteOk returns a tuple with the Lte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLte

`func (o *Match) SetLte(v Lte)`

SetLte sets Lte field to given value.

### HasLte

`func (o *Match) HasLte() bool`

HasLte returns a boolean if a field has been set.

### SetLteNil

`func (o *Match) SetLteNil(b bool)`

 SetLteNil sets the value for Lte to be an explicit nil

### UnsetLte
`func (o *Match) UnsetLte()`

UnsetLte ensures that no value is present for Lte, not even an explicit nil
### GetValues

`func (o *Match) GetValues() []ValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *Match) GetValuesOk() (*[]ValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *Match) SetValues(v []ValuesInner)`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


