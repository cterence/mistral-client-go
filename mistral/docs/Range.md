# Range

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "range"]
**Field** | **string** |  | 
**Gt** | Pointer to [**NullableGt**](Gt.md) |  | [optional] 
**Gte** | Pointer to [**NullableGte**](Gte.md) |  | [optional] 
**Lt** | Pointer to [**NullableLt**](Lt.md) |  | [optional] 
**Lte** | Pointer to [**NullableLte**](Lte.md) |  | [optional] 

## Methods

### NewRange

`func NewRange(field string, ) *Range`

NewRange instantiates a new Range object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRangeWithDefaults

`func NewRangeWithDefaults() *Range`

NewRangeWithDefaults instantiates a new Range object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Range) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Range) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Range) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Range) HasType() bool`

HasType returns a boolean if a field has been set.

### GetField

`func (o *Range) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *Range) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *Range) SetField(v string)`

SetField sets Field field to given value.


### GetGt

`func (o *Range) GetGt() Gt`

GetGt returns the Gt field if non-nil, zero value otherwise.

### GetGtOk

`func (o *Range) GetGtOk() (*Gt, bool)`

GetGtOk returns a tuple with the Gt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGt

`func (o *Range) SetGt(v Gt)`

SetGt sets Gt field to given value.

### HasGt

`func (o *Range) HasGt() bool`

HasGt returns a boolean if a field has been set.

### SetGtNil

`func (o *Range) SetGtNil(b bool)`

 SetGtNil sets the value for Gt to be an explicit nil

### UnsetGt
`func (o *Range) UnsetGt()`

UnsetGt ensures that no value is present for Gt, not even an explicit nil
### GetGte

`func (o *Range) GetGte() Gte`

GetGte returns the Gte field if non-nil, zero value otherwise.

### GetGteOk

`func (o *Range) GetGteOk() (*Gte, bool)`

GetGteOk returns a tuple with the Gte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGte

`func (o *Range) SetGte(v Gte)`

SetGte sets Gte field to given value.

### HasGte

`func (o *Range) HasGte() bool`

HasGte returns a boolean if a field has been set.

### SetGteNil

`func (o *Range) SetGteNil(b bool)`

 SetGteNil sets the value for Gte to be an explicit nil

### UnsetGte
`func (o *Range) UnsetGte()`

UnsetGte ensures that no value is present for Gte, not even an explicit nil
### GetLt

`func (o *Range) GetLt() Lt`

GetLt returns the Lt field if non-nil, zero value otherwise.

### GetLtOk

`func (o *Range) GetLtOk() (*Lt, bool)`

GetLtOk returns a tuple with the Lt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLt

`func (o *Range) SetLt(v Lt)`

SetLt sets Lt field to given value.

### HasLt

`func (o *Range) HasLt() bool`

HasLt returns a boolean if a field has been set.

### SetLtNil

`func (o *Range) SetLtNil(b bool)`

 SetLtNil sets the value for Lt to be an explicit nil

### UnsetLt
`func (o *Range) UnsetLt()`

UnsetLt ensures that no value is present for Lt, not even an explicit nil
### GetLte

`func (o *Range) GetLte() Lte`

GetLte returns the Lte field if non-nil, zero value otherwise.

### GetLteOk

`func (o *Range) GetLteOk() (*Lte, bool)`

GetLteOk returns a tuple with the Lte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLte

`func (o *Range) SetLte(v Lte)`

SetLte sets Lte field to given value.

### HasLte

`func (o *Range) HasLte() bool`

HasLte returns a boolean if a field has been set.

### SetLteNil

`func (o *Range) SetLteNil(b bool)`

 SetLteNil sets the value for Lte to be an explicit nil

### UnsetLte
`func (o *Range) UnsetLte()`

UnsetLte ensures that no value is present for Lte, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


