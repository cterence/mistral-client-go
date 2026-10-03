# FilterGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AND** | Pointer to [**[]FilterGroupANDInner**](FilterGroupANDInner.md) |  | [optional] 
**OR** | Pointer to [**[]FilterGroupANDInner**](FilterGroupANDInner.md) |  | [optional] 

## Methods

### NewFilterGroup

`func NewFilterGroup() *FilterGroup`

NewFilterGroup instantiates a new FilterGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilterGroupWithDefaults

`func NewFilterGroupWithDefaults() *FilterGroup`

NewFilterGroupWithDefaults instantiates a new FilterGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAND

`func (o *FilterGroup) GetAND() []FilterGroupANDInner`

GetAND returns the AND field if non-nil, zero value otherwise.

### GetANDOk

`func (o *FilterGroup) GetANDOk() (*[]FilterGroupANDInner, bool)`

GetANDOk returns a tuple with the AND field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAND

`func (o *FilterGroup) SetAND(v []FilterGroupANDInner)`

SetAND sets AND field to given value.

### HasAND

`func (o *FilterGroup) HasAND() bool`

HasAND returns a boolean if a field has been set.

### SetANDNil

`func (o *FilterGroup) SetANDNil(b bool)`

 SetANDNil sets the value for AND to be an explicit nil

### UnsetAND
`func (o *FilterGroup) UnsetAND()`

UnsetAND ensures that no value is present for AND, not even an explicit nil
### GetOR

`func (o *FilterGroup) GetOR() []FilterGroupANDInner`

GetOR returns the OR field if non-nil, zero value otherwise.

### GetOROk

`func (o *FilterGroup) GetOROk() (*[]FilterGroupANDInner, bool)`

GetOROk returns a tuple with the OR field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOR

`func (o *FilterGroup) SetOR(v []FilterGroupANDInner)`

SetOR sets OR field to given value.

### HasOR

`func (o *FilterGroup) HasOR() bool`

HasOR returns a boolean if a field has been set.

### SetORNil

`func (o *FilterGroup) SetORNil(b bool)`

 SetORNil sets the value for OR to be an explicit nil

### UnsetOR
`func (o *FilterGroup) UnsetOR()`

UnsetOR ensures that no value is present for OR, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


