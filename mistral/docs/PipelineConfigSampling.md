# PipelineConfigSampling

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "probabilistic"]
**Percentage** | **float32** |  | 
**Attribute** | **string** |  | 

## Methods

### NewPipelineConfigSampling

`func NewPipelineConfigSampling(percentage float32, attribute string, ) *PipelineConfigSampling`

NewPipelineConfigSampling instantiates a new PipelineConfigSampling object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineConfigSamplingWithDefaults

`func NewPipelineConfigSamplingWithDefaults() *PipelineConfigSampling`

NewPipelineConfigSamplingWithDefaults instantiates a new PipelineConfigSampling object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PipelineConfigSampling) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PipelineConfigSampling) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PipelineConfigSampling) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PipelineConfigSampling) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPercentage

`func (o *PipelineConfigSampling) GetPercentage() float32`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *PipelineConfigSampling) GetPercentageOk() (*float32, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *PipelineConfigSampling) SetPercentage(v float32)`

SetPercentage sets Percentage field to given value.


### GetAttribute

`func (o *PipelineConfigSampling) GetAttribute() string`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *PipelineConfigSampling) GetAttributeOk() (*string, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *PipelineConfigSampling) SetAttribute(v string)`

SetAttribute sets Attribute field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


