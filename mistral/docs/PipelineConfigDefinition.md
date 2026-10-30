# PipelineConfigDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetAttributes** | **[]string** |  | [default to {"gen_ai.input.messages", "gen_ai.output.messages"}]
**Patterns** | [**[]DetectionPattern**](DetectionPattern.md) |  | 
**Model** | **string** |  | 
**Prompt** | **string** |  | 
**Temperature** | Pointer to **float32** |  | [optional] 
**TopP** | Pointer to **float32** |  | [optional] 
**MaxTokens** | Pointer to **int32** |  | [optional] 
**RandomSeed** | Pointer to **int32** |  | [optional] 
**ResultSpec** | [**EvaluationResultSpec**](EvaluationResultSpec.md) |  | 
**Mapping** | Pointer to **map[string]string** |  | [optional] 
**UseCache** | Pointer to **bool** |  | [optional] [default to true]
**Destination** | [**OTLPDestination**](OTLPDestination.md) |  | 

## Methods

### NewPipelineConfigDefinition

`func NewPipelineConfigDefinition(targetAttributes []string, patterns []DetectionPattern, model string, prompt string, resultSpec EvaluationResultSpec, destination OTLPDestination, ) *PipelineConfigDefinition`

NewPipelineConfigDefinition instantiates a new PipelineConfigDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineConfigDefinitionWithDefaults

`func NewPipelineConfigDefinitionWithDefaults() *PipelineConfigDefinition`

NewPipelineConfigDefinitionWithDefaults instantiates a new PipelineConfigDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetAttributes

`func (o *PipelineConfigDefinition) GetTargetAttributes() []string`

GetTargetAttributes returns the TargetAttributes field if non-nil, zero value otherwise.

### GetTargetAttributesOk

`func (o *PipelineConfigDefinition) GetTargetAttributesOk() (*[]string, bool)`

GetTargetAttributesOk returns a tuple with the TargetAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAttributes

`func (o *PipelineConfigDefinition) SetTargetAttributes(v []string)`

SetTargetAttributes sets TargetAttributes field to given value.


### GetPatterns

`func (o *PipelineConfigDefinition) GetPatterns() []DetectionPattern`

GetPatterns returns the Patterns field if non-nil, zero value otherwise.

### GetPatternsOk

`func (o *PipelineConfigDefinition) GetPatternsOk() (*[]DetectionPattern, bool)`

GetPatternsOk returns a tuple with the Patterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatterns

`func (o *PipelineConfigDefinition) SetPatterns(v []DetectionPattern)`

SetPatterns sets Patterns field to given value.


### GetModel

`func (o *PipelineConfigDefinition) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *PipelineConfigDefinition) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *PipelineConfigDefinition) SetModel(v string)`

SetModel sets Model field to given value.


### GetPrompt

`func (o *PipelineConfigDefinition) GetPrompt() string`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *PipelineConfigDefinition) GetPromptOk() (*string, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *PipelineConfigDefinition) SetPrompt(v string)`

SetPrompt sets Prompt field to given value.


### GetTemperature

`func (o *PipelineConfigDefinition) GetTemperature() float32`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *PipelineConfigDefinition) GetTemperatureOk() (*float32, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *PipelineConfigDefinition) SetTemperature(v float32)`

SetTemperature sets Temperature field to given value.

### HasTemperature

`func (o *PipelineConfigDefinition) HasTemperature() bool`

HasTemperature returns a boolean if a field has been set.

### GetTopP

`func (o *PipelineConfigDefinition) GetTopP() float32`

GetTopP returns the TopP field if non-nil, zero value otherwise.

### GetTopPOk

`func (o *PipelineConfigDefinition) GetTopPOk() (*float32, bool)`

GetTopPOk returns a tuple with the TopP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopP

`func (o *PipelineConfigDefinition) SetTopP(v float32)`

SetTopP sets TopP field to given value.

### HasTopP

`func (o *PipelineConfigDefinition) HasTopP() bool`

HasTopP returns a boolean if a field has been set.

### GetMaxTokens

`func (o *PipelineConfigDefinition) GetMaxTokens() int32`

GetMaxTokens returns the MaxTokens field if non-nil, zero value otherwise.

### GetMaxTokensOk

`func (o *PipelineConfigDefinition) GetMaxTokensOk() (*int32, bool)`

GetMaxTokensOk returns a tuple with the MaxTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTokens

`func (o *PipelineConfigDefinition) SetMaxTokens(v int32)`

SetMaxTokens sets MaxTokens field to given value.

### HasMaxTokens

`func (o *PipelineConfigDefinition) HasMaxTokens() bool`

HasMaxTokens returns a boolean if a field has been set.

### GetRandomSeed

`func (o *PipelineConfigDefinition) GetRandomSeed() int32`

GetRandomSeed returns the RandomSeed field if non-nil, zero value otherwise.

### GetRandomSeedOk

`func (o *PipelineConfigDefinition) GetRandomSeedOk() (*int32, bool)`

GetRandomSeedOk returns a tuple with the RandomSeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRandomSeed

`func (o *PipelineConfigDefinition) SetRandomSeed(v int32)`

SetRandomSeed sets RandomSeed field to given value.

### HasRandomSeed

`func (o *PipelineConfigDefinition) HasRandomSeed() bool`

HasRandomSeed returns a boolean if a field has been set.

### GetResultSpec

`func (o *PipelineConfigDefinition) GetResultSpec() EvaluationResultSpec`

GetResultSpec returns the ResultSpec field if non-nil, zero value otherwise.

### GetResultSpecOk

`func (o *PipelineConfigDefinition) GetResultSpecOk() (*EvaluationResultSpec, bool)`

GetResultSpecOk returns a tuple with the ResultSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultSpec

`func (o *PipelineConfigDefinition) SetResultSpec(v EvaluationResultSpec)`

SetResultSpec sets ResultSpec field to given value.


### GetMapping

`func (o *PipelineConfigDefinition) GetMapping() map[string]string`

GetMapping returns the Mapping field if non-nil, zero value otherwise.

### GetMappingOk

`func (o *PipelineConfigDefinition) GetMappingOk() (*map[string]string, bool)`

GetMappingOk returns a tuple with the Mapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapping

`func (o *PipelineConfigDefinition) SetMapping(v map[string]string)`

SetMapping sets Mapping field to given value.

### HasMapping

`func (o *PipelineConfigDefinition) HasMapping() bool`

HasMapping returns a boolean if a field has been set.

### GetUseCache

`func (o *PipelineConfigDefinition) GetUseCache() bool`

GetUseCache returns the UseCache field if non-nil, zero value otherwise.

### GetUseCacheOk

`func (o *PipelineConfigDefinition) GetUseCacheOk() (*bool, bool)`

GetUseCacheOk returns a tuple with the UseCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCache

`func (o *PipelineConfigDefinition) SetUseCache(v bool)`

SetUseCache sets UseCache field to given value.

### HasUseCache

`func (o *PipelineConfigDefinition) HasUseCache() bool`

HasUseCache returns a boolean if a field has been set.

### GetDestination

`func (o *PipelineConfigDefinition) GetDestination() OTLPDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *PipelineConfigDefinition) GetDestinationOk() (*OTLPDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *PipelineConfigDefinition) SetDestination(v OTLPDestination)`

SetDestination sets Destination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


