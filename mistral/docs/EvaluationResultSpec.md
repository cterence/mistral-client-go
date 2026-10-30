# EvaluationResultSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "boolean"]
**Labels** | **[]string** |  | 
**Min** | **float32** |  | 
**Max** | **float32** |  | 

## Methods

### NewEvaluationResultSpec

`func NewEvaluationResultSpec(labels []string, min float32, max float32, ) *EvaluationResultSpec`

NewEvaluationResultSpec instantiates a new EvaluationResultSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvaluationResultSpecWithDefaults

`func NewEvaluationResultSpecWithDefaults() *EvaluationResultSpec`

NewEvaluationResultSpecWithDefaults instantiates a new EvaluationResultSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EvaluationResultSpec) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EvaluationResultSpec) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EvaluationResultSpec) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EvaluationResultSpec) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLabels

`func (o *EvaluationResultSpec) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *EvaluationResultSpec) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *EvaluationResultSpec) SetLabels(v []string)`

SetLabels sets Labels field to given value.


### GetMin

`func (o *EvaluationResultSpec) GetMin() float32`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *EvaluationResultSpec) GetMinOk() (*float32, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *EvaluationResultSpec) SetMin(v float32)`

SetMin sets Min field to given value.


### GetMax

`func (o *EvaluationResultSpec) GetMax() float32`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *EvaluationResultSpec) GetMaxOk() (*float32, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *EvaluationResultSpec) SetMax(v float32)`

SetMax sets Max field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


