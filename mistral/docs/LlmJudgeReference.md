# LlmJudgeReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Slug** | **string** |  | 
**Mapping** | Pointer to **map[string]string** |  | [optional] 
**Sampling** | Pointer to [**NullablePipelineConfigSampling**](PipelineConfigSampling.md) |  | [optional] 

## Methods

### NewLlmJudgeReference

`func NewLlmJudgeReference(slug string, ) *LlmJudgeReference`

NewLlmJudgeReference instantiates a new LlmJudgeReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLlmJudgeReferenceWithDefaults

`func NewLlmJudgeReferenceWithDefaults() *LlmJudgeReference`

NewLlmJudgeReferenceWithDefaults instantiates a new LlmJudgeReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSlug

`func (o *LlmJudgeReference) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *LlmJudgeReference) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *LlmJudgeReference) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetMapping

`func (o *LlmJudgeReference) GetMapping() map[string]string`

GetMapping returns the Mapping field if non-nil, zero value otherwise.

### GetMappingOk

`func (o *LlmJudgeReference) GetMappingOk() (*map[string]string, bool)`

GetMappingOk returns a tuple with the Mapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapping

`func (o *LlmJudgeReference) SetMapping(v map[string]string)`

SetMapping sets Mapping field to given value.

### HasMapping

`func (o *LlmJudgeReference) HasMapping() bool`

HasMapping returns a boolean if a field has been set.

### GetSampling

`func (o *LlmJudgeReference) GetSampling() PipelineConfigSampling`

GetSampling returns the Sampling field if non-nil, zero value otherwise.

### GetSamplingOk

`func (o *LlmJudgeReference) GetSamplingOk() (*PipelineConfigSampling, bool)`

GetSamplingOk returns a tuple with the Sampling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampling

`func (o *LlmJudgeReference) SetSampling(v PipelineConfigSampling)`

SetSampling sets Sampling field to given value.

### HasSampling

`func (o *LlmJudgeReference) HasSampling() bool`

HasSampling returns a boolean if a field has been set.

### SetSamplingNil

`func (o *LlmJudgeReference) SetSamplingNil(b bool)`

 SetSamplingNil sets the value for Sampling to be an explicit nil

### UnsetSampling
`func (o *LlmJudgeReference) UnsetSampling()`

UnsetSampling ensures that no value is present for Sampling, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


