# RegisterDeploymentRequestVespaIndex

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Fields** | [**[]RegisterDeploymentRequestVespaField**](RegisterDeploymentRequestVespaField.md) |  | 
**Sd** | **string** |  | 
**EmbeddingDimensions** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewRegisterDeploymentRequestVespaIndex

`func NewRegisterDeploymentRequestVespaIndex(name string, fields []RegisterDeploymentRequestVespaField, sd string, ) *RegisterDeploymentRequestVespaIndex`

NewRegisterDeploymentRequestVespaIndex instantiates a new RegisterDeploymentRequestVespaIndex object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterDeploymentRequestVespaIndexWithDefaults

`func NewRegisterDeploymentRequestVespaIndexWithDefaults() *RegisterDeploymentRequestVespaIndex`

NewRegisterDeploymentRequestVespaIndexWithDefaults instantiates a new RegisterDeploymentRequestVespaIndex object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RegisterDeploymentRequestVespaIndex) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RegisterDeploymentRequestVespaIndex) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RegisterDeploymentRequestVespaIndex) SetName(v string)`

SetName sets Name field to given value.


### GetFields

`func (o *RegisterDeploymentRequestVespaIndex) GetFields() []RegisterDeploymentRequestVespaField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *RegisterDeploymentRequestVespaIndex) GetFieldsOk() (*[]RegisterDeploymentRequestVespaField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *RegisterDeploymentRequestVespaIndex) SetFields(v []RegisterDeploymentRequestVespaField)`

SetFields sets Fields field to given value.


### GetSd

`func (o *RegisterDeploymentRequestVespaIndex) GetSd() string`

GetSd returns the Sd field if non-nil, zero value otherwise.

### GetSdOk

`func (o *RegisterDeploymentRequestVespaIndex) GetSdOk() (*string, bool)`

GetSdOk returns a tuple with the Sd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSd

`func (o *RegisterDeploymentRequestVespaIndex) SetSd(v string)`

SetSd sets Sd field to given value.


### GetEmbeddingDimensions

`func (o *RegisterDeploymentRequestVespaIndex) GetEmbeddingDimensions() int32`

GetEmbeddingDimensions returns the EmbeddingDimensions field if non-nil, zero value otherwise.

### GetEmbeddingDimensionsOk

`func (o *RegisterDeploymentRequestVespaIndex) GetEmbeddingDimensionsOk() (*int32, bool)`

GetEmbeddingDimensionsOk returns a tuple with the EmbeddingDimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbeddingDimensions

`func (o *RegisterDeploymentRequestVespaIndex) SetEmbeddingDimensions(v int32)`

SetEmbeddingDimensions sets EmbeddingDimensions field to given value.

### HasEmbeddingDimensions

`func (o *RegisterDeploymentRequestVespaIndex) HasEmbeddingDimensions() bool`

HasEmbeddingDimensions returns a boolean if a field has been set.

### SetEmbeddingDimensionsNil

`func (o *RegisterDeploymentRequestVespaIndex) SetEmbeddingDimensionsNil(b bool)`

 SetEmbeddingDimensionsNil sets the value for EmbeddingDimensions to be an explicit nil

### UnsetEmbeddingDimensions
`func (o *RegisterDeploymentRequestVespaIndex) UnsetEmbeddingDimensions()`

UnsetEmbeddingDimensions ensures that no value is present for EmbeddingDimensions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


