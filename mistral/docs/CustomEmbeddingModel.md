# CustomEmbeddingModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "custom"]
**Name** | **string** |  | 
**Dtype** | Pointer to [**VectorDType**](VectorDType.md) |  | [optional] [default to VECTORDTYPE_FLOAT16]
**Dimensions** | **int32** |  | 
**DistanceMetric** | Pointer to [**DistanceMetric**](DistanceMetric.md) |  | [optional] [default to DISTANCEMETRIC_COSINE]

## Methods

### NewCustomEmbeddingModel

`func NewCustomEmbeddingModel(name string, dimensions int32, ) *CustomEmbeddingModel`

NewCustomEmbeddingModel instantiates a new CustomEmbeddingModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomEmbeddingModelWithDefaults

`func NewCustomEmbeddingModelWithDefaults() *CustomEmbeddingModel`

NewCustomEmbeddingModelWithDefaults instantiates a new CustomEmbeddingModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CustomEmbeddingModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomEmbeddingModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomEmbeddingModel) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CustomEmbeddingModel) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *CustomEmbeddingModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomEmbeddingModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomEmbeddingModel) SetName(v string)`

SetName sets Name field to given value.


### GetDtype

`func (o *CustomEmbeddingModel) GetDtype() VectorDType`

GetDtype returns the Dtype field if non-nil, zero value otherwise.

### GetDtypeOk

`func (o *CustomEmbeddingModel) GetDtypeOk() (*VectorDType, bool)`

GetDtypeOk returns a tuple with the Dtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtype

`func (o *CustomEmbeddingModel) SetDtype(v VectorDType)`

SetDtype sets Dtype field to given value.

### HasDtype

`func (o *CustomEmbeddingModel) HasDtype() bool`

HasDtype returns a boolean if a field has been set.

### GetDimensions

`func (o *CustomEmbeddingModel) GetDimensions() int32`

GetDimensions returns the Dimensions field if non-nil, zero value otherwise.

### GetDimensionsOk

`func (o *CustomEmbeddingModel) GetDimensionsOk() (*int32, bool)`

GetDimensionsOk returns a tuple with the Dimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensions

`func (o *CustomEmbeddingModel) SetDimensions(v int32)`

SetDimensions sets Dimensions field to given value.


### GetDistanceMetric

`func (o *CustomEmbeddingModel) GetDistanceMetric() DistanceMetric`

GetDistanceMetric returns the DistanceMetric field if non-nil, zero value otherwise.

### GetDistanceMetricOk

`func (o *CustomEmbeddingModel) GetDistanceMetricOk() (*DistanceMetric, bool)`

GetDistanceMetricOk returns a tuple with the DistanceMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistanceMetric

`func (o *CustomEmbeddingModel) SetDistanceMetric(v DistanceMetric)`

SetDistanceMetric sets DistanceMetric field to given value.

### HasDistanceMetric

`func (o *CustomEmbeddingModel) HasDistanceMetric() bool`

HasDistanceMetric returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


