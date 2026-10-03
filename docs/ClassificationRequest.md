# ClassificationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | **string** | ID of the model to use. | 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**Input** | [**Input1**](Input1.md) |  | 

## Methods

### NewClassificationRequest

`func NewClassificationRequest(model string, input Input1, ) *ClassificationRequest`

NewClassificationRequest instantiates a new ClassificationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClassificationRequestWithDefaults

`func NewClassificationRequestWithDefaults() *ClassificationRequest`

NewClassificationRequestWithDefaults instantiates a new ClassificationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *ClassificationRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ClassificationRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ClassificationRequest) SetModel(v string)`

SetModel sets Model field to given value.


### GetMetadata

`func (o *ClassificationRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ClassificationRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ClassificationRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ClassificationRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *ClassificationRequest) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *ClassificationRequest) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetInput

`func (o *ClassificationRequest) GetInput() Input1`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *ClassificationRequest) GetInputOk() (*Input1, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *ClassificationRequest) SetInput(v Input1)`

SetInput sets Input field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


