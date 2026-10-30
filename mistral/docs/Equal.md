# Equal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "equal"]
**Field** | **string** |  | 
**Value** | [**Value5**](Value5.md) |  | 

## Methods

### NewEqual

`func NewEqual(field string, value Value5, ) *Equal`

NewEqual instantiates a new Equal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEqualWithDefaults

`func NewEqualWithDefaults() *Equal`

NewEqualWithDefaults instantiates a new Equal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Equal) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Equal) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Equal) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Equal) HasType() bool`

HasType returns a boolean if a field has been set.

### GetField

`func (o *Equal) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *Equal) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *Equal) SetField(v string)`

SetField sets Field field to given value.


### GetValue

`func (o *Equal) GetValue() Value5`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Equal) GetValueOk() (*Value5, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Equal) SetValue(v Value5)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


