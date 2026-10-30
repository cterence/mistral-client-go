# SystemNameFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **string** |  | 
**Operator** | Pointer to [**ContentSearchOperator**](ContentSearchOperator.md) |  | [optional] [default to CONTENTSEARCHOPERATOR_EQUALS]

## Methods

### NewSystemNameFilter

`func NewSystemNameFilter(value string, ) *SystemNameFilter`

NewSystemNameFilter instantiates a new SystemNameFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemNameFilterWithDefaults

`func NewSystemNameFilterWithDefaults() *SystemNameFilter`

NewSystemNameFilterWithDefaults instantiates a new SystemNameFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *SystemNameFilter) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SystemNameFilter) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SystemNameFilter) SetValue(v string)`

SetValue sets Value field to given value.


### GetOperator

`func (o *SystemNameFilter) GetOperator() ContentSearchOperator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *SystemNameFilter) GetOperatorOk() (*ContentSearchOperator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *SystemNameFilter) SetOperator(v ContentSearchOperator)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *SystemNameFilter) HasOperator() bool`

HasOperator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


