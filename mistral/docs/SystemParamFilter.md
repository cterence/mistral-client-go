# SystemParamFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** |  | 
**Operator** | [**SystemParamOperator**](SystemParamOperator.md) |  | 
**Value** | [**Value**](Value.md) |  | 

## Methods

### NewSystemParamFilter

`func NewSystemParamFilter(key string, operator SystemParamOperator, value Value, ) *SystemParamFilter`

NewSystemParamFilter instantiates a new SystemParamFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemParamFilterWithDefaults

`func NewSystemParamFilterWithDefaults() *SystemParamFilter`

NewSystemParamFilterWithDefaults instantiates a new SystemParamFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *SystemParamFilter) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *SystemParamFilter) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *SystemParamFilter) SetKey(v string)`

SetKey sets Key field to given value.


### GetOperator

`func (o *SystemParamFilter) GetOperator() SystemParamOperator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *SystemParamFilter) GetOperatorOk() (*SystemParamOperator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *SystemParamFilter) SetOperator(v SystemParamOperator)`

SetOperator sets Operator field to given value.


### GetValue

`func (o *SystemParamFilter) GetValue() Value`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SystemParamFilter) GetValueOk() (*Value, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SystemParamFilter) SetValue(v Value)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


