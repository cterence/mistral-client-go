# MatchesInner

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

### NewMatchesInner

`func NewMatchesInner(matches []MatchesInner, match Match, field string, value Value5, values []ValuesInner, ) *MatchesInner`

NewMatchesInner instantiates a new MatchesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMatchesInnerWithDefaults

`func NewMatchesInnerWithDefaults() *MatchesInner`

NewMatchesInnerWithDefaults instantiates a new MatchesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MatchesInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MatchesInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MatchesInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MatchesInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMatches

`func (o *MatchesInner) GetMatches() []MatchesInner`

GetMatches returns the Matches field if non-nil, zero value otherwise.

### GetMatchesOk

`func (o *MatchesInner) GetMatchesOk() (*[]MatchesInner, bool)`

GetMatchesOk returns a tuple with the Matches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatches

`func (o *MatchesInner) SetMatches(v []MatchesInner)`

SetMatches sets Matches field to given value.


### GetMatch

`func (o *MatchesInner) GetMatch() Match`

GetMatch returns the Match field if non-nil, zero value otherwise.

### GetMatchOk

`func (o *MatchesInner) GetMatchOk() (*Match, bool)`

GetMatchOk returns a tuple with the Match field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatch

`func (o *MatchesInner) SetMatch(v Match)`

SetMatch sets Match field to given value.


### GetField

`func (o *MatchesInner) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *MatchesInner) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *MatchesInner) SetField(v string)`

SetField sets Field field to given value.


### GetValue

`func (o *MatchesInner) GetValue() Value5`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MatchesInner) GetValueOk() (*Value5, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *MatchesInner) SetValue(v Value5)`

SetValue sets Value field to given value.


### GetGt

`func (o *MatchesInner) GetGt() Gt`

GetGt returns the Gt field if non-nil, zero value otherwise.

### GetGtOk

`func (o *MatchesInner) GetGtOk() (*Gt, bool)`

GetGtOk returns a tuple with the Gt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGt

`func (o *MatchesInner) SetGt(v Gt)`

SetGt sets Gt field to given value.

### HasGt

`func (o *MatchesInner) HasGt() bool`

HasGt returns a boolean if a field has been set.

### SetGtNil

`func (o *MatchesInner) SetGtNil(b bool)`

 SetGtNil sets the value for Gt to be an explicit nil

### UnsetGt
`func (o *MatchesInner) UnsetGt()`

UnsetGt ensures that no value is present for Gt, not even an explicit nil
### GetGte

`func (o *MatchesInner) GetGte() Gte`

GetGte returns the Gte field if non-nil, zero value otherwise.

### GetGteOk

`func (o *MatchesInner) GetGteOk() (*Gte, bool)`

GetGteOk returns a tuple with the Gte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGte

`func (o *MatchesInner) SetGte(v Gte)`

SetGte sets Gte field to given value.

### HasGte

`func (o *MatchesInner) HasGte() bool`

HasGte returns a boolean if a field has been set.

### SetGteNil

`func (o *MatchesInner) SetGteNil(b bool)`

 SetGteNil sets the value for Gte to be an explicit nil

### UnsetGte
`func (o *MatchesInner) UnsetGte()`

UnsetGte ensures that no value is present for Gte, not even an explicit nil
### GetLt

`func (o *MatchesInner) GetLt() Lt`

GetLt returns the Lt field if non-nil, zero value otherwise.

### GetLtOk

`func (o *MatchesInner) GetLtOk() (*Lt, bool)`

GetLtOk returns a tuple with the Lt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLt

`func (o *MatchesInner) SetLt(v Lt)`

SetLt sets Lt field to given value.

### HasLt

`func (o *MatchesInner) HasLt() bool`

HasLt returns a boolean if a field has been set.

### SetLtNil

`func (o *MatchesInner) SetLtNil(b bool)`

 SetLtNil sets the value for Lt to be an explicit nil

### UnsetLt
`func (o *MatchesInner) UnsetLt()`

UnsetLt ensures that no value is present for Lt, not even an explicit nil
### GetLte

`func (o *MatchesInner) GetLte() Lte`

GetLte returns the Lte field if non-nil, zero value otherwise.

### GetLteOk

`func (o *MatchesInner) GetLteOk() (*Lte, bool)`

GetLteOk returns a tuple with the Lte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLte

`func (o *MatchesInner) SetLte(v Lte)`

SetLte sets Lte field to given value.

### HasLte

`func (o *MatchesInner) HasLte() bool`

HasLte returns a boolean if a field has been set.

### SetLteNil

`func (o *MatchesInner) SetLteNil(b bool)`

 SetLteNil sets the value for Lte to be an explicit nil

### UnsetLte
`func (o *MatchesInner) UnsetLte()`

UnsetLte ensures that no value is present for Lte, not even an explicit nil
### GetValues

`func (o *MatchesInner) GetValues() []ValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *MatchesInner) GetValuesOk() (*[]ValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *MatchesInner) SetValues(v []ValuesInner)`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


