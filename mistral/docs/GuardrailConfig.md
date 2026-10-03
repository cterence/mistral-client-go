# GuardrailConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockOnError** | Pointer to **bool** | If true, return HTTP 403 and block request in the event of a server-side error | [optional] [default to false]
**ModerationLlmV1** | Pointer to [**NullableModerationLLMV1Config**](ModerationLLMV1Config.md) |  | [optional] 
**ModerationLlmV2** | Pointer to [**NullableModerationLLMV2Config**](ModerationLLMV2Config.md) |  | [optional] 

## Methods

### NewGuardrailConfig

`func NewGuardrailConfig() *GuardrailConfig`

NewGuardrailConfig instantiates a new GuardrailConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuardrailConfigWithDefaults

`func NewGuardrailConfigWithDefaults() *GuardrailConfig`

NewGuardrailConfigWithDefaults instantiates a new GuardrailConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockOnError

`func (o *GuardrailConfig) GetBlockOnError() bool`

GetBlockOnError returns the BlockOnError field if non-nil, zero value otherwise.

### GetBlockOnErrorOk

`func (o *GuardrailConfig) GetBlockOnErrorOk() (*bool, bool)`

GetBlockOnErrorOk returns a tuple with the BlockOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockOnError

`func (o *GuardrailConfig) SetBlockOnError(v bool)`

SetBlockOnError sets BlockOnError field to given value.

### HasBlockOnError

`func (o *GuardrailConfig) HasBlockOnError() bool`

HasBlockOnError returns a boolean if a field has been set.

### GetModerationLlmV1

`func (o *GuardrailConfig) GetModerationLlmV1() ModerationLLMV1Config`

GetModerationLlmV1 returns the ModerationLlmV1 field if non-nil, zero value otherwise.

### GetModerationLlmV1Ok

`func (o *GuardrailConfig) GetModerationLlmV1Ok() (*ModerationLLMV1Config, bool)`

GetModerationLlmV1Ok returns a tuple with the ModerationLlmV1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModerationLlmV1

`func (o *GuardrailConfig) SetModerationLlmV1(v ModerationLLMV1Config)`

SetModerationLlmV1 sets ModerationLlmV1 field to given value.

### HasModerationLlmV1

`func (o *GuardrailConfig) HasModerationLlmV1() bool`

HasModerationLlmV1 returns a boolean if a field has been set.

### SetModerationLlmV1Nil

`func (o *GuardrailConfig) SetModerationLlmV1Nil(b bool)`

 SetModerationLlmV1Nil sets the value for ModerationLlmV1 to be an explicit nil

### UnsetModerationLlmV1
`func (o *GuardrailConfig) UnsetModerationLlmV1()`

UnsetModerationLlmV1 ensures that no value is present for ModerationLlmV1, not even an explicit nil
### GetModerationLlmV2

`func (o *GuardrailConfig) GetModerationLlmV2() ModerationLLMV2Config`

GetModerationLlmV2 returns the ModerationLlmV2 field if non-nil, zero value otherwise.

### GetModerationLlmV2Ok

`func (o *GuardrailConfig) GetModerationLlmV2Ok() (*ModerationLLMV2Config, bool)`

GetModerationLlmV2Ok returns a tuple with the ModerationLlmV2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModerationLlmV2

`func (o *GuardrailConfig) SetModerationLlmV2(v ModerationLLMV2Config)`

SetModerationLlmV2 sets ModerationLlmV2 field to given value.

### HasModerationLlmV2

`func (o *GuardrailConfig) HasModerationLlmV2() bool`

HasModerationLlmV2 returns a boolean if a field has been set.

### SetModerationLlmV2Nil

`func (o *GuardrailConfig) SetModerationLlmV2Nil(b bool)`

 SetModerationLlmV2Nil sets the value for ModerationLlmV2 to be an explicit nil

### UnsetModerationLlmV2
`func (o *GuardrailConfig) UnsetModerationLlmV2()`

UnsetModerationLlmV2 ensures that no value is present for ModerationLlmV2, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


