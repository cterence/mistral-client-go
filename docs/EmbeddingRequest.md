# EmbeddingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | **string** | ID of the model to use. | 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**Input** | [**Input1**](Input1.md) |  | 
**OutputDimension** | Pointer to **NullableInt32** | The dimension of the output embeddings when feature available. If not provided, a default output dimension will be used. | [optional] 
**OutputDtype** | Pointer to [**EmbeddingDtype**](EmbeddingDtype.md) | The data type of the output embeddings when feature available. If not provided, a default output data type will be used. | [optional] [default to EMBEDDINGDTYPE_FLOAT]
**EncodingFormat** | Pointer to [**EncodingFormat**](EncodingFormat.md) | The format of embeddings in the response. | [optional] [default to ENCODINGFORMAT_FLOAT]

## Methods

### NewEmbeddingRequest

`func NewEmbeddingRequest(model string, input Input1, ) *EmbeddingRequest`

NewEmbeddingRequest instantiates a new EmbeddingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmbeddingRequestWithDefaults

`func NewEmbeddingRequestWithDefaults() *EmbeddingRequest`

NewEmbeddingRequestWithDefaults instantiates a new EmbeddingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *EmbeddingRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *EmbeddingRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *EmbeddingRequest) SetModel(v string)`

SetModel sets Model field to given value.


### GetMetadata

`func (o *EmbeddingRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *EmbeddingRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *EmbeddingRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *EmbeddingRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *EmbeddingRequest) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *EmbeddingRequest) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetInput

`func (o *EmbeddingRequest) GetInput() Input1`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *EmbeddingRequest) GetInputOk() (*Input1, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *EmbeddingRequest) SetInput(v Input1)`

SetInput sets Input field to given value.


### GetOutputDimension

`func (o *EmbeddingRequest) GetOutputDimension() int32`

GetOutputDimension returns the OutputDimension field if non-nil, zero value otherwise.

### GetOutputDimensionOk

`func (o *EmbeddingRequest) GetOutputDimensionOk() (*int32, bool)`

GetOutputDimensionOk returns a tuple with the OutputDimension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputDimension

`func (o *EmbeddingRequest) SetOutputDimension(v int32)`

SetOutputDimension sets OutputDimension field to given value.

### HasOutputDimension

`func (o *EmbeddingRequest) HasOutputDimension() bool`

HasOutputDimension returns a boolean if a field has been set.

### SetOutputDimensionNil

`func (o *EmbeddingRequest) SetOutputDimensionNil(b bool)`

 SetOutputDimensionNil sets the value for OutputDimension to be an explicit nil

### UnsetOutputDimension
`func (o *EmbeddingRequest) UnsetOutputDimension()`

UnsetOutputDimension ensures that no value is present for OutputDimension, not even an explicit nil
### GetOutputDtype

`func (o *EmbeddingRequest) GetOutputDtype() EmbeddingDtype`

GetOutputDtype returns the OutputDtype field if non-nil, zero value otherwise.

### GetOutputDtypeOk

`func (o *EmbeddingRequest) GetOutputDtypeOk() (*EmbeddingDtype, bool)`

GetOutputDtypeOk returns a tuple with the OutputDtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputDtype

`func (o *EmbeddingRequest) SetOutputDtype(v EmbeddingDtype)`

SetOutputDtype sets OutputDtype field to given value.

### HasOutputDtype

`func (o *EmbeddingRequest) HasOutputDtype() bool`

HasOutputDtype returns a boolean if a field has been set.

### GetEncodingFormat

`func (o *EmbeddingRequest) GetEncodingFormat() EncodingFormat`

GetEncodingFormat returns the EncodingFormat field if non-nil, zero value otherwise.

### GetEncodingFormatOk

`func (o *EmbeddingRequest) GetEncodingFormatOk() (*EncodingFormat, bool)`

GetEncodingFormatOk returns a tuple with the EncodingFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncodingFormat

`func (o *EmbeddingRequest) SetEncodingFormat(v EncodingFormat)`

SetEncodingFormat sets EncodingFormat field to given value.

### HasEncodingFormat

`func (o *EmbeddingRequest) HasEncodingFormat() bool`

HasEncodingFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


