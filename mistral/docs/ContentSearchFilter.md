# ContentSearchFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **NullableString** |  | [optional] 
**Value** | **string** |  | 
**Operator** | Pointer to [**ContentSearchOperator**](ContentSearchOperator.md) |  | [optional] [default to CONTENTSEARCHOPERATOR_CONTAINS]

## Methods

### NewContentSearchFilter

`func NewContentSearchFilter(value string, ) *ContentSearchFilter`

NewContentSearchFilter instantiates a new ContentSearchFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentSearchFilterWithDefaults

`func NewContentSearchFilterWithDefaults() *ContentSearchFilter`

NewContentSearchFilterWithDefaults instantiates a new ContentSearchFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *ContentSearchFilter) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ContentSearchFilter) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ContentSearchFilter) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ContentSearchFilter) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *ContentSearchFilter) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *ContentSearchFilter) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetValue

`func (o *ContentSearchFilter) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ContentSearchFilter) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ContentSearchFilter) SetValue(v string)`

SetValue sets Value field to given value.


### GetOperator

`func (o *ContentSearchFilter) GetOperator() ContentSearchOperator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *ContentSearchFilter) GetOperatorOk() (*ContentSearchOperator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *ContentSearchFilter) SetOperator(v ContentSearchOperator)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *ContentSearchFilter) HasOperator() bool`

HasOperator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


