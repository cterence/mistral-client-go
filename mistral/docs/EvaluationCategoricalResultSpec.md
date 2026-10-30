# EvaluationCategoricalResultSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "categorical"]
**Labels** | **[]string** |  | 

## Methods

### NewEvaluationCategoricalResultSpec

`func NewEvaluationCategoricalResultSpec(labels []string, ) *EvaluationCategoricalResultSpec`

NewEvaluationCategoricalResultSpec instantiates a new EvaluationCategoricalResultSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvaluationCategoricalResultSpecWithDefaults

`func NewEvaluationCategoricalResultSpecWithDefaults() *EvaluationCategoricalResultSpec`

NewEvaluationCategoricalResultSpecWithDefaults instantiates a new EvaluationCategoricalResultSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EvaluationCategoricalResultSpec) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EvaluationCategoricalResultSpec) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EvaluationCategoricalResultSpec) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EvaluationCategoricalResultSpec) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLabels

`func (o *EvaluationCategoricalResultSpec) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *EvaluationCategoricalResultSpec) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *EvaluationCategoricalResultSpec) SetLabels(v []string)`

SetLabels sets Labels field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


