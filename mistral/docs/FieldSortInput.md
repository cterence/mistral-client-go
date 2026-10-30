# FieldSortInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** |  | [optional] [default to "field"]
**Field** | [**RecordSortField**](RecordSortField.md) |  | 
**Direction** | Pointer to [**SortDirection**](SortDirection.md) |  | [optional] [default to SORTDIRECTION_ASC]

## Methods

### NewFieldSortInput

`func NewFieldSortInput(field RecordSortField, ) *FieldSortInput`

NewFieldSortInput instantiates a new FieldSortInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldSortInputWithDefaults

`func NewFieldSortInputWithDefaults() *FieldSortInput`

NewFieldSortInputWithDefaults instantiates a new FieldSortInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *FieldSortInput) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *FieldSortInput) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *FieldSortInput) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *FieldSortInput) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetField

`func (o *FieldSortInput) GetField() RecordSortField`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *FieldSortInput) GetFieldOk() (*RecordSortField, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *FieldSortInput) SetField(v RecordSortField)`

SetField sets Field field to given value.


### GetDirection

`func (o *FieldSortInput) GetDirection() SortDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *FieldSortInput) GetDirectionOk() (*SortDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *FieldSortInput) SetDirection(v SortDirection)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *FieldSortInput) HasDirection() bool`

HasDirection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


