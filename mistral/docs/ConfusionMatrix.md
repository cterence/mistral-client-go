# ConfusionMatrix

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Labels** | **[]string** |  | 
**Matrix** | **[][]int32** |  | 

## Methods

### NewConfusionMatrix

`func NewConfusionMatrix(labels []string, matrix [][]int32, ) *ConfusionMatrix`

NewConfusionMatrix instantiates a new ConfusionMatrix object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfusionMatrixWithDefaults

`func NewConfusionMatrixWithDefaults() *ConfusionMatrix`

NewConfusionMatrixWithDefaults instantiates a new ConfusionMatrix object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabels

`func (o *ConfusionMatrix) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ConfusionMatrix) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ConfusionMatrix) SetLabels(v []string)`

SetLabels sets Labels field to given value.


### GetMatrix

`func (o *ConfusionMatrix) GetMatrix() [][]int32`

GetMatrix returns the Matrix field if non-nil, zero value otherwise.

### GetMatrixOk

`func (o *ConfusionMatrix) GetMatrixOk() (*[][]int32, bool)`

GetMatrixOk returns a tuple with the Matrix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatrix

`func (o *ConfusionMatrix) SetMatrix(v [][]int32)`

SetMatrix sets Matrix field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


