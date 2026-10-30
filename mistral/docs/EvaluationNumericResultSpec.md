# EvaluationNumericResultSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "numeric"]
**Min** | **float32** |  | 
**Max** | **float32** |  | 

## Methods

### NewEvaluationNumericResultSpec

`func NewEvaluationNumericResultSpec(min float32, max float32, ) *EvaluationNumericResultSpec`

NewEvaluationNumericResultSpec instantiates a new EvaluationNumericResultSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvaluationNumericResultSpecWithDefaults

`func NewEvaluationNumericResultSpecWithDefaults() *EvaluationNumericResultSpec`

NewEvaluationNumericResultSpecWithDefaults instantiates a new EvaluationNumericResultSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EvaluationNumericResultSpec) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EvaluationNumericResultSpec) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EvaluationNumericResultSpec) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EvaluationNumericResultSpec) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMin

`func (o *EvaluationNumericResultSpec) GetMin() float32`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *EvaluationNumericResultSpec) GetMinOk() (*float32, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *EvaluationNumericResultSpec) SetMin(v float32)`

SetMin sets Min field to given value.


### GetMax

`func (o *EvaluationNumericResultSpec) GetMax() float32`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *EvaluationNumericResultSpec) GetMaxOk() (*float32, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *EvaluationNumericResultSpec) SetMax(v float32)`

SetMax sets Max field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


