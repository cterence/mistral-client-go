# JudgeDefinitionV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | **string** |  | 
**Prompt** | **string** |  | 
**Temperature** | Pointer to **NullableFloat32** |  | [optional] 
**TopP** | Pointer to **NullableFloat32** |  | [optional] 
**MaxTokens** | Pointer to **NullableInt32** |  | [optional] 
**RandomSeed** | Pointer to **NullableInt32** |  | [optional] 
**ResultSpec** | [**EvaluationResultSpec**](EvaluationResultSpec.md) |  | 
**Mapping** | Pointer to **map[string]string** |  | [optional] 
**UseCache** | Pointer to **bool** |  | [optional] [default to true]

## Methods

### NewJudgeDefinitionV1

`func NewJudgeDefinitionV1(model string, prompt string, resultSpec EvaluationResultSpec, ) *JudgeDefinitionV1`

NewJudgeDefinitionV1 instantiates a new JudgeDefinitionV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJudgeDefinitionV1WithDefaults

`func NewJudgeDefinitionV1WithDefaults() *JudgeDefinitionV1`

NewJudgeDefinitionV1WithDefaults instantiates a new JudgeDefinitionV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *JudgeDefinitionV1) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *JudgeDefinitionV1) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *JudgeDefinitionV1) SetModel(v string)`

SetModel sets Model field to given value.


### GetPrompt

`func (o *JudgeDefinitionV1) GetPrompt() string`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *JudgeDefinitionV1) GetPromptOk() (*string, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *JudgeDefinitionV1) SetPrompt(v string)`

SetPrompt sets Prompt field to given value.


### GetTemperature

`func (o *JudgeDefinitionV1) GetTemperature() float32`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *JudgeDefinitionV1) GetTemperatureOk() (*float32, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *JudgeDefinitionV1) SetTemperature(v float32)`

SetTemperature sets Temperature field to given value.

### HasTemperature

`func (o *JudgeDefinitionV1) HasTemperature() bool`

HasTemperature returns a boolean if a field has been set.

### SetTemperatureNil

`func (o *JudgeDefinitionV1) SetTemperatureNil(b bool)`

 SetTemperatureNil sets the value for Temperature to be an explicit nil

### UnsetTemperature
`func (o *JudgeDefinitionV1) UnsetTemperature()`

UnsetTemperature ensures that no value is present for Temperature, not even an explicit nil
### GetTopP

`func (o *JudgeDefinitionV1) GetTopP() float32`

GetTopP returns the TopP field if non-nil, zero value otherwise.

### GetTopPOk

`func (o *JudgeDefinitionV1) GetTopPOk() (*float32, bool)`

GetTopPOk returns a tuple with the TopP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopP

`func (o *JudgeDefinitionV1) SetTopP(v float32)`

SetTopP sets TopP field to given value.

### HasTopP

`func (o *JudgeDefinitionV1) HasTopP() bool`

HasTopP returns a boolean if a field has been set.

### SetTopPNil

`func (o *JudgeDefinitionV1) SetTopPNil(b bool)`

 SetTopPNil sets the value for TopP to be an explicit nil

### UnsetTopP
`func (o *JudgeDefinitionV1) UnsetTopP()`

UnsetTopP ensures that no value is present for TopP, not even an explicit nil
### GetMaxTokens

`func (o *JudgeDefinitionV1) GetMaxTokens() int32`

GetMaxTokens returns the MaxTokens field if non-nil, zero value otherwise.

### GetMaxTokensOk

`func (o *JudgeDefinitionV1) GetMaxTokensOk() (*int32, bool)`

GetMaxTokensOk returns a tuple with the MaxTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTokens

`func (o *JudgeDefinitionV1) SetMaxTokens(v int32)`

SetMaxTokens sets MaxTokens field to given value.

### HasMaxTokens

`func (o *JudgeDefinitionV1) HasMaxTokens() bool`

HasMaxTokens returns a boolean if a field has been set.

### SetMaxTokensNil

`func (o *JudgeDefinitionV1) SetMaxTokensNil(b bool)`

 SetMaxTokensNil sets the value for MaxTokens to be an explicit nil

### UnsetMaxTokens
`func (o *JudgeDefinitionV1) UnsetMaxTokens()`

UnsetMaxTokens ensures that no value is present for MaxTokens, not even an explicit nil
### GetRandomSeed

`func (o *JudgeDefinitionV1) GetRandomSeed() int32`

GetRandomSeed returns the RandomSeed field if non-nil, zero value otherwise.

### GetRandomSeedOk

`func (o *JudgeDefinitionV1) GetRandomSeedOk() (*int32, bool)`

GetRandomSeedOk returns a tuple with the RandomSeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRandomSeed

`func (o *JudgeDefinitionV1) SetRandomSeed(v int32)`

SetRandomSeed sets RandomSeed field to given value.

### HasRandomSeed

`func (o *JudgeDefinitionV1) HasRandomSeed() bool`

HasRandomSeed returns a boolean if a field has been set.

### SetRandomSeedNil

`func (o *JudgeDefinitionV1) SetRandomSeedNil(b bool)`

 SetRandomSeedNil sets the value for RandomSeed to be an explicit nil

### UnsetRandomSeed
`func (o *JudgeDefinitionV1) UnsetRandomSeed()`

UnsetRandomSeed ensures that no value is present for RandomSeed, not even an explicit nil
### GetResultSpec

`func (o *JudgeDefinitionV1) GetResultSpec() EvaluationResultSpec`

GetResultSpec returns the ResultSpec field if non-nil, zero value otherwise.

### GetResultSpecOk

`func (o *JudgeDefinitionV1) GetResultSpecOk() (*EvaluationResultSpec, bool)`

GetResultSpecOk returns a tuple with the ResultSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultSpec

`func (o *JudgeDefinitionV1) SetResultSpec(v EvaluationResultSpec)`

SetResultSpec sets ResultSpec field to given value.


### GetMapping

`func (o *JudgeDefinitionV1) GetMapping() map[string]string`

GetMapping returns the Mapping field if non-nil, zero value otherwise.

### GetMappingOk

`func (o *JudgeDefinitionV1) GetMappingOk() (*map[string]string, bool)`

GetMappingOk returns a tuple with the Mapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapping

`func (o *JudgeDefinitionV1) SetMapping(v map[string]string)`

SetMapping sets Mapping field to given value.

### HasMapping

`func (o *JudgeDefinitionV1) HasMapping() bool`

HasMapping returns a boolean if a field has been set.

### GetUseCache

`func (o *JudgeDefinitionV1) GetUseCache() bool`

GetUseCache returns the UseCache field if non-nil, zero value otherwise.

### GetUseCacheOk

`func (o *JudgeDefinitionV1) GetUseCacheOk() (*bool, bool)`

GetUseCacheOk returns a tuple with the UseCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCache

`func (o *JudgeDefinitionV1) SetUseCache(v bool)`

SetUseCache sets UseCache field to given value.

### HasUseCache

`func (o *JudgeDefinitionV1) HasUseCache() bool`

HasUseCache returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


