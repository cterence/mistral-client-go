# In

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "in"]
**Field** | **string** |  | 
**Values** | [**[]ValuesInner**](ValuesInner.md) |  | 

## Methods

### NewIn

`func NewIn(field string, values []ValuesInner, ) *In`

NewIn instantiates a new In object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInWithDefaults

`func NewInWithDefaults() *In`

NewInWithDefaults instantiates a new In object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *In) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *In) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *In) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *In) HasType() bool`

HasType returns a boolean if a field has been set.

### GetField

`func (o *In) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *In) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *In) SetField(v string)`

SetField sets Field field to given value.


### GetValues

`func (o *In) GetValues() []ValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *In) GetValuesOk() (*[]ValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *In) SetValues(v []ValuesInner)`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


