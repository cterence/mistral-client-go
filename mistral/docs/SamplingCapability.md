# SamplingCapability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to **map[string]interface{}** | Capability for context inclusion during sampling.  Indicates support for non-&#39;none&#39; values in the includeContext parameter. SOFT-DEPRECATED: New implementations should use tools parameter instead. | [optional] 
**Tools** | Pointer to **map[string]interface{}** | Capability indicating support for tool calling during sampling.  When present in ClientCapabilities.sampling, indicates that the client supports the tools and toolChoice parameters in sampling requests. | [optional] 

## Methods

### NewSamplingCapability

`func NewSamplingCapability() *SamplingCapability`

NewSamplingCapability instantiates a new SamplingCapability object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSamplingCapabilityWithDefaults

`func NewSamplingCapabilityWithDefaults() *SamplingCapability`

NewSamplingCapabilityWithDefaults instantiates a new SamplingCapability object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *SamplingCapability) GetContext() map[string]interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *SamplingCapability) GetContextOk() (*map[string]interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *SamplingCapability) SetContext(v map[string]interface{})`

SetContext sets Context field to given value.

### HasContext

`func (o *SamplingCapability) HasContext() bool`

HasContext returns a boolean if a field has been set.

### SetContextNil

`func (o *SamplingCapability) SetContextNil(b bool)`

 SetContextNil sets the value for Context to be an explicit nil

### UnsetContext
`func (o *SamplingCapability) UnsetContext()`

UnsetContext ensures that no value is present for Context, not even an explicit nil
### GetTools

`func (o *SamplingCapability) GetTools() map[string]interface{}`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *SamplingCapability) GetToolsOk() (*map[string]interface{}, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *SamplingCapability) SetTools(v map[string]interface{})`

SetTools sets Tools field to given value.

### HasTools

`func (o *SamplingCapability) HasTools() bool`

HasTools returns a boolean if a field has been set.

### SetToolsNil

`func (o *SamplingCapability) SetToolsNil(b bool)`

 SetToolsNil sets the value for Tools to be an explicit nil

### UnsetTools
`func (o *SamplingCapability) UnsetTools()`

UnsetTools ensures that no value is present for Tools, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


