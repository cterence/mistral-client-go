# MistralEmbeddingModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "mistral"]
**Name** | **string** |  | 
**Dtype** | Pointer to [**VectorDType**](VectorDType.md) |  | [optional] [default to VECTORDTYPE_FLOAT16]
**Dimensions** | **int32** |  | 
**DistanceMetric** | Pointer to [**DistanceMetric**](DistanceMetric.md) |  | [optional] [default to DISTANCEMETRIC_COSINE]

## Methods

### NewMistralEmbeddingModel

`func NewMistralEmbeddingModel(name string, dimensions int32, ) *MistralEmbeddingModel`

NewMistralEmbeddingModel instantiates a new MistralEmbeddingModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMistralEmbeddingModelWithDefaults

`func NewMistralEmbeddingModelWithDefaults() *MistralEmbeddingModel`

NewMistralEmbeddingModelWithDefaults instantiates a new MistralEmbeddingModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MistralEmbeddingModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MistralEmbeddingModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MistralEmbeddingModel) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MistralEmbeddingModel) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *MistralEmbeddingModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MistralEmbeddingModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MistralEmbeddingModel) SetName(v string)`

SetName sets Name field to given value.


### GetDtype

`func (o *MistralEmbeddingModel) GetDtype() VectorDType`

GetDtype returns the Dtype field if non-nil, zero value otherwise.

### GetDtypeOk

`func (o *MistralEmbeddingModel) GetDtypeOk() (*VectorDType, bool)`

GetDtypeOk returns a tuple with the Dtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtype

`func (o *MistralEmbeddingModel) SetDtype(v VectorDType)`

SetDtype sets Dtype field to given value.

### HasDtype

`func (o *MistralEmbeddingModel) HasDtype() bool`

HasDtype returns a boolean if a field has been set.

### GetDimensions

`func (o *MistralEmbeddingModel) GetDimensions() int32`

GetDimensions returns the Dimensions field if non-nil, zero value otherwise.

### GetDimensionsOk

`func (o *MistralEmbeddingModel) GetDimensionsOk() (*int32, bool)`

GetDimensionsOk returns a tuple with the Dimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensions

`func (o *MistralEmbeddingModel) SetDimensions(v int32)`

SetDimensions sets Dimensions field to given value.


### GetDistanceMetric

`func (o *MistralEmbeddingModel) GetDistanceMetric() DistanceMetric`

GetDistanceMetric returns the DistanceMetric field if non-nil, zero value otherwise.

### GetDistanceMetricOk

`func (o *MistralEmbeddingModel) GetDistanceMetricOk() (*DistanceMetric, bool)`

GetDistanceMetricOk returns a tuple with the DistanceMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistanceMetric

`func (o *MistralEmbeddingModel) SetDistanceMetric(v DistanceMetric)`

SetDistanceMetric sets DistanceMetric field to given value.

### HasDistanceMetric

`func (o *MistralEmbeddingModel) HasDistanceMetric() bool`

HasDistanceMetric returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


