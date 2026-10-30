# ChatCompletionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | **string** | ID of the model to use. You can use the [List Available Models](/api/#tag/models/operation/list_models_v1_models_get) API to see all of your available models, or see our [Model overview](/models) for model descriptions. | 
**Temperature** | Pointer to **NullableFloat32** | What sampling temperature to use, we recommend between 0.0 and 0.7. Higher values like 0.7 will make the output more random, while lower values like 0.2 will make it more focused and deterministic. We generally recommend altering this or &#x60;top_p&#x60; but not both. The default value varies depending on the model you are targeting. Call the &#x60;/models&#x60; endpoint to retrieve the appropriate value. | [optional] 
**TopP** | Pointer to **NullableFloat32** | Nucleus sampling, where the model considers the results of the tokens with &#x60;top_p&#x60; probability mass. So 0.1 means only the tokens comprising the top 10% probability mass are considered. We generally recommend altering this or &#x60;temperature&#x60; but not both. | [optional] 
**MaxTokens** | Pointer to **NullableInt32** | The maximum number of tokens to generate in the completion. The token count of your prompt plus &#x60;max_tokens&#x60; cannot exceed the model&#39;s context length. | [optional] 
**Stream** | Pointer to **bool** | Whether to stream back partial progress. If set, tokens will be sent as data-only server-side events as they become available, with the stream terminated by a data: [DONE] message. Otherwise, the server will hold the request open until the timeout or until completion, with the response containing the full result as JSON. | [optional] [default to false]
**Stop** | Pointer to [**NullableStop**](Stop.md) |  | [optional] 
**RandomSeed** | Pointer to **NullableInt32** | The seed to use for random sampling. If set, different calls will generate deterministic results. | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**Messages** | [**[]MessagesInner**](MessagesInner.md) | The prompt(s) to generate completions for, encoded as a list of dict with role and content. | 
**ResponseFormat** | Pointer to [**ResponseFormat**](ResponseFormat.md) |  | [optional] 
**Tools** | Pointer to [**[]ChatCompletionRequestToolsInner**](ChatCompletionRequestToolsInner.md) | A list of tools the model may call. Use this to provide a list of functions the model may generate JSON inputs for. | [optional] 
**ToolChoice** | Pointer to [**ToolChoice**](ToolChoice.md) |  | [optional] [default to auto]
**PresencePenalty** | Pointer to **NullableFloat32** | The &#x60;presence_penalty&#x60; determines how much the model penalizes the repetition of words or phrases. A higher presence penalty encourages the model to use a wider variety of words and phrases, making the output more diverse and creative. | [optional] 
**FrequencyPenalty** | Pointer to **NullableFloat32** | The &#x60;frequency_penalty&#x60; penalizes the repetition of words based on their frequency in the generated text. A higher frequency penalty discourages the model from repeating words that have already appeared frequently in the output, promoting diversity and reducing repetition. | [optional] 
**MinTokens** | Pointer to **NullableInt32** |  | [optional] 
**RepetitionPenalty** | Pointer to **NullableFloat32** |  | [optional] 
**TopK** | Pointer to **NullableInt32** |  | [optional] 
**Logprobs** | Pointer to **NullableBool** |  | [optional] 
**TopLogprobs** | Pointer to **NullableInt32** |  | [optional] 
**PromptLogprobs** | Pointer to **NullableBool** |  | [optional] 
**TopPromptLogprobs** | Pointer to **NullableInt32** |  | [optional] 
**N** | Pointer to **NullableInt32** | Number of completions to return for each request, input tokens are only billed once. | [optional] 
**Prediction** | Pointer to [**Prediction**](Prediction.md) | Enable users to specify expected results, optimizing response times by leveraging known or predictable content. This approach is especially effective for updating text documents or code files with minimal changes, reducing latency while maintaining high-quality results. | [optional] [default to {type=content, content=}]
**ParallelToolCalls** | Pointer to **bool** | Whether to enable parallel function calling during tool use, when enabled the model can call multiple tools in parallel. | [optional] [default to true]
**ReasoningEffort** | Pointer to [**NullableReasoningEffort**](ReasoningEffort.md) |  | [optional] 
**PromptMode** | Pointer to [**NullableMistralPromptMode**](MistralPromptMode.md) | Allows toggling between the reasoning mode and no system prompt. When set to &#x60;reasoning&#x60; the system prompt for reasoning models will be used. | [optional] 
**Guardrails** | Pointer to [**[]GuardrailConfig**](GuardrailConfig.md) |  | [optional] 
**PromptCacheKey** | Pointer to **NullableString** |  | [optional] 
**ServiceTier** | Pointer to [**NullableRequestedServiceTier**](RequestedServiceTier.md) | Determines whether to serve the request using priority or standard capacity. | [optional] 
**SafePrompt** | Pointer to **bool** | Whether to inject a safety prompt before all conversations. | [optional] [default to false]

## Methods

### NewChatCompletionRequest

`func NewChatCompletionRequest(model string, messages []MessagesInner, ) *ChatCompletionRequest`

NewChatCompletionRequest instantiates a new ChatCompletionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatCompletionRequestWithDefaults

`func NewChatCompletionRequestWithDefaults() *ChatCompletionRequest`

NewChatCompletionRequestWithDefaults instantiates a new ChatCompletionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *ChatCompletionRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ChatCompletionRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ChatCompletionRequest) SetModel(v string)`

SetModel sets Model field to given value.


### GetTemperature

`func (o *ChatCompletionRequest) GetTemperature() float32`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *ChatCompletionRequest) GetTemperatureOk() (*float32, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *ChatCompletionRequest) SetTemperature(v float32)`

SetTemperature sets Temperature field to given value.

### HasTemperature

`func (o *ChatCompletionRequest) HasTemperature() bool`

HasTemperature returns a boolean if a field has been set.

### SetTemperatureNil

`func (o *ChatCompletionRequest) SetTemperatureNil(b bool)`

 SetTemperatureNil sets the value for Temperature to be an explicit nil

### UnsetTemperature
`func (o *ChatCompletionRequest) UnsetTemperature()`

UnsetTemperature ensures that no value is present for Temperature, not even an explicit nil
### GetTopP

`func (o *ChatCompletionRequest) GetTopP() float32`

GetTopP returns the TopP field if non-nil, zero value otherwise.

### GetTopPOk

`func (o *ChatCompletionRequest) GetTopPOk() (*float32, bool)`

GetTopPOk returns a tuple with the TopP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopP

`func (o *ChatCompletionRequest) SetTopP(v float32)`

SetTopP sets TopP field to given value.

### HasTopP

`func (o *ChatCompletionRequest) HasTopP() bool`

HasTopP returns a boolean if a field has been set.

### SetTopPNil

`func (o *ChatCompletionRequest) SetTopPNil(b bool)`

 SetTopPNil sets the value for TopP to be an explicit nil

### UnsetTopP
`func (o *ChatCompletionRequest) UnsetTopP()`

UnsetTopP ensures that no value is present for TopP, not even an explicit nil
### GetMaxTokens

`func (o *ChatCompletionRequest) GetMaxTokens() int32`

GetMaxTokens returns the MaxTokens field if non-nil, zero value otherwise.

### GetMaxTokensOk

`func (o *ChatCompletionRequest) GetMaxTokensOk() (*int32, bool)`

GetMaxTokensOk returns a tuple with the MaxTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTokens

`func (o *ChatCompletionRequest) SetMaxTokens(v int32)`

SetMaxTokens sets MaxTokens field to given value.

### HasMaxTokens

`func (o *ChatCompletionRequest) HasMaxTokens() bool`

HasMaxTokens returns a boolean if a field has been set.

### SetMaxTokensNil

`func (o *ChatCompletionRequest) SetMaxTokensNil(b bool)`

 SetMaxTokensNil sets the value for MaxTokens to be an explicit nil

### UnsetMaxTokens
`func (o *ChatCompletionRequest) UnsetMaxTokens()`

UnsetMaxTokens ensures that no value is present for MaxTokens, not even an explicit nil
### GetStream

`func (o *ChatCompletionRequest) GetStream() bool`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *ChatCompletionRequest) GetStreamOk() (*bool, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *ChatCompletionRequest) SetStream(v bool)`

SetStream sets Stream field to given value.

### HasStream

`func (o *ChatCompletionRequest) HasStream() bool`

HasStream returns a boolean if a field has been set.

### GetStop

`func (o *ChatCompletionRequest) GetStop() Stop`

GetStop returns the Stop field if non-nil, zero value otherwise.

### GetStopOk

`func (o *ChatCompletionRequest) GetStopOk() (*Stop, bool)`

GetStopOk returns a tuple with the Stop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStop

`func (o *ChatCompletionRequest) SetStop(v Stop)`

SetStop sets Stop field to given value.

### HasStop

`func (o *ChatCompletionRequest) HasStop() bool`

HasStop returns a boolean if a field has been set.

### SetStopNil

`func (o *ChatCompletionRequest) SetStopNil(b bool)`

 SetStopNil sets the value for Stop to be an explicit nil

### UnsetStop
`func (o *ChatCompletionRequest) UnsetStop()`

UnsetStop ensures that no value is present for Stop, not even an explicit nil
### GetRandomSeed

`func (o *ChatCompletionRequest) GetRandomSeed() int32`

GetRandomSeed returns the RandomSeed field if non-nil, zero value otherwise.

### GetRandomSeedOk

`func (o *ChatCompletionRequest) GetRandomSeedOk() (*int32, bool)`

GetRandomSeedOk returns a tuple with the RandomSeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRandomSeed

`func (o *ChatCompletionRequest) SetRandomSeed(v int32)`

SetRandomSeed sets RandomSeed field to given value.

### HasRandomSeed

`func (o *ChatCompletionRequest) HasRandomSeed() bool`

HasRandomSeed returns a boolean if a field has been set.

### SetRandomSeedNil

`func (o *ChatCompletionRequest) SetRandomSeedNil(b bool)`

 SetRandomSeedNil sets the value for RandomSeed to be an explicit nil

### UnsetRandomSeed
`func (o *ChatCompletionRequest) UnsetRandomSeed()`

UnsetRandomSeed ensures that no value is present for RandomSeed, not even an explicit nil
### GetMetadata

`func (o *ChatCompletionRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ChatCompletionRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ChatCompletionRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ChatCompletionRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *ChatCompletionRequest) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *ChatCompletionRequest) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetMessages

`func (o *ChatCompletionRequest) GetMessages() []MessagesInner`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *ChatCompletionRequest) GetMessagesOk() (*[]MessagesInner, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *ChatCompletionRequest) SetMessages(v []MessagesInner)`

SetMessages sets Messages field to given value.


### GetResponseFormat

`func (o *ChatCompletionRequest) GetResponseFormat() ResponseFormat`

GetResponseFormat returns the ResponseFormat field if non-nil, zero value otherwise.

### GetResponseFormatOk

`func (o *ChatCompletionRequest) GetResponseFormatOk() (*ResponseFormat, bool)`

GetResponseFormatOk returns a tuple with the ResponseFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseFormat

`func (o *ChatCompletionRequest) SetResponseFormat(v ResponseFormat)`

SetResponseFormat sets ResponseFormat field to given value.

### HasResponseFormat

`func (o *ChatCompletionRequest) HasResponseFormat() bool`

HasResponseFormat returns a boolean if a field has been set.

### GetTools

`func (o *ChatCompletionRequest) GetTools() []ChatCompletionRequestToolsInner`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *ChatCompletionRequest) GetToolsOk() (*[]ChatCompletionRequestToolsInner, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *ChatCompletionRequest) SetTools(v []ChatCompletionRequestToolsInner)`

SetTools sets Tools field to given value.

### HasTools

`func (o *ChatCompletionRequest) HasTools() bool`

HasTools returns a boolean if a field has been set.

### SetToolsNil

`func (o *ChatCompletionRequest) SetToolsNil(b bool)`

 SetToolsNil sets the value for Tools to be an explicit nil

### UnsetTools
`func (o *ChatCompletionRequest) UnsetTools()`

UnsetTools ensures that no value is present for Tools, not even an explicit nil
### GetToolChoice

`func (o *ChatCompletionRequest) GetToolChoice() ToolChoice`

GetToolChoice returns the ToolChoice field if non-nil, zero value otherwise.

### GetToolChoiceOk

`func (o *ChatCompletionRequest) GetToolChoiceOk() (*ToolChoice, bool)`

GetToolChoiceOk returns a tuple with the ToolChoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolChoice

`func (o *ChatCompletionRequest) SetToolChoice(v ToolChoice)`

SetToolChoice sets ToolChoice field to given value.

### HasToolChoice

`func (o *ChatCompletionRequest) HasToolChoice() bool`

HasToolChoice returns a boolean if a field has been set.

### GetPresencePenalty

`func (o *ChatCompletionRequest) GetPresencePenalty() float32`

GetPresencePenalty returns the PresencePenalty field if non-nil, zero value otherwise.

### GetPresencePenaltyOk

`func (o *ChatCompletionRequest) GetPresencePenaltyOk() (*float32, bool)`

GetPresencePenaltyOk returns a tuple with the PresencePenalty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresencePenalty

`func (o *ChatCompletionRequest) SetPresencePenalty(v float32)`

SetPresencePenalty sets PresencePenalty field to given value.

### HasPresencePenalty

`func (o *ChatCompletionRequest) HasPresencePenalty() bool`

HasPresencePenalty returns a boolean if a field has been set.

### SetPresencePenaltyNil

`func (o *ChatCompletionRequest) SetPresencePenaltyNil(b bool)`

 SetPresencePenaltyNil sets the value for PresencePenalty to be an explicit nil

### UnsetPresencePenalty
`func (o *ChatCompletionRequest) UnsetPresencePenalty()`

UnsetPresencePenalty ensures that no value is present for PresencePenalty, not even an explicit nil
### GetFrequencyPenalty

`func (o *ChatCompletionRequest) GetFrequencyPenalty() float32`

GetFrequencyPenalty returns the FrequencyPenalty field if non-nil, zero value otherwise.

### GetFrequencyPenaltyOk

`func (o *ChatCompletionRequest) GetFrequencyPenaltyOk() (*float32, bool)`

GetFrequencyPenaltyOk returns a tuple with the FrequencyPenalty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencyPenalty

`func (o *ChatCompletionRequest) SetFrequencyPenalty(v float32)`

SetFrequencyPenalty sets FrequencyPenalty field to given value.

### HasFrequencyPenalty

`func (o *ChatCompletionRequest) HasFrequencyPenalty() bool`

HasFrequencyPenalty returns a boolean if a field has been set.

### SetFrequencyPenaltyNil

`func (o *ChatCompletionRequest) SetFrequencyPenaltyNil(b bool)`

 SetFrequencyPenaltyNil sets the value for FrequencyPenalty to be an explicit nil

### UnsetFrequencyPenalty
`func (o *ChatCompletionRequest) UnsetFrequencyPenalty()`

UnsetFrequencyPenalty ensures that no value is present for FrequencyPenalty, not even an explicit nil
### GetMinTokens

`func (o *ChatCompletionRequest) GetMinTokens() int32`

GetMinTokens returns the MinTokens field if non-nil, zero value otherwise.

### GetMinTokensOk

`func (o *ChatCompletionRequest) GetMinTokensOk() (*int32, bool)`

GetMinTokensOk returns a tuple with the MinTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinTokens

`func (o *ChatCompletionRequest) SetMinTokens(v int32)`

SetMinTokens sets MinTokens field to given value.

### HasMinTokens

`func (o *ChatCompletionRequest) HasMinTokens() bool`

HasMinTokens returns a boolean if a field has been set.

### SetMinTokensNil

`func (o *ChatCompletionRequest) SetMinTokensNil(b bool)`

 SetMinTokensNil sets the value for MinTokens to be an explicit nil

### UnsetMinTokens
`func (o *ChatCompletionRequest) UnsetMinTokens()`

UnsetMinTokens ensures that no value is present for MinTokens, not even an explicit nil
### GetRepetitionPenalty

`func (o *ChatCompletionRequest) GetRepetitionPenalty() float32`

GetRepetitionPenalty returns the RepetitionPenalty field if non-nil, zero value otherwise.

### GetRepetitionPenaltyOk

`func (o *ChatCompletionRequest) GetRepetitionPenaltyOk() (*float32, bool)`

GetRepetitionPenaltyOk returns a tuple with the RepetitionPenalty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepetitionPenalty

`func (o *ChatCompletionRequest) SetRepetitionPenalty(v float32)`

SetRepetitionPenalty sets RepetitionPenalty field to given value.

### HasRepetitionPenalty

`func (o *ChatCompletionRequest) HasRepetitionPenalty() bool`

HasRepetitionPenalty returns a boolean if a field has been set.

### SetRepetitionPenaltyNil

`func (o *ChatCompletionRequest) SetRepetitionPenaltyNil(b bool)`

 SetRepetitionPenaltyNil sets the value for RepetitionPenalty to be an explicit nil

### UnsetRepetitionPenalty
`func (o *ChatCompletionRequest) UnsetRepetitionPenalty()`

UnsetRepetitionPenalty ensures that no value is present for RepetitionPenalty, not even an explicit nil
### GetTopK

`func (o *ChatCompletionRequest) GetTopK() int32`

GetTopK returns the TopK field if non-nil, zero value otherwise.

### GetTopKOk

`func (o *ChatCompletionRequest) GetTopKOk() (*int32, bool)`

GetTopKOk returns a tuple with the TopK field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopK

`func (o *ChatCompletionRequest) SetTopK(v int32)`

SetTopK sets TopK field to given value.

### HasTopK

`func (o *ChatCompletionRequest) HasTopK() bool`

HasTopK returns a boolean if a field has been set.

### SetTopKNil

`func (o *ChatCompletionRequest) SetTopKNil(b bool)`

 SetTopKNil sets the value for TopK to be an explicit nil

### UnsetTopK
`func (o *ChatCompletionRequest) UnsetTopK()`

UnsetTopK ensures that no value is present for TopK, not even an explicit nil
### GetLogprobs

`func (o *ChatCompletionRequest) GetLogprobs() bool`

GetLogprobs returns the Logprobs field if non-nil, zero value otherwise.

### GetLogprobsOk

`func (o *ChatCompletionRequest) GetLogprobsOk() (*bool, bool)`

GetLogprobsOk returns a tuple with the Logprobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogprobs

`func (o *ChatCompletionRequest) SetLogprobs(v bool)`

SetLogprobs sets Logprobs field to given value.

### HasLogprobs

`func (o *ChatCompletionRequest) HasLogprobs() bool`

HasLogprobs returns a boolean if a field has been set.

### SetLogprobsNil

`func (o *ChatCompletionRequest) SetLogprobsNil(b bool)`

 SetLogprobsNil sets the value for Logprobs to be an explicit nil

### UnsetLogprobs
`func (o *ChatCompletionRequest) UnsetLogprobs()`

UnsetLogprobs ensures that no value is present for Logprobs, not even an explicit nil
### GetTopLogprobs

`func (o *ChatCompletionRequest) GetTopLogprobs() int32`

GetTopLogprobs returns the TopLogprobs field if non-nil, zero value otherwise.

### GetTopLogprobsOk

`func (o *ChatCompletionRequest) GetTopLogprobsOk() (*int32, bool)`

GetTopLogprobsOk returns a tuple with the TopLogprobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopLogprobs

`func (o *ChatCompletionRequest) SetTopLogprobs(v int32)`

SetTopLogprobs sets TopLogprobs field to given value.

### HasTopLogprobs

`func (o *ChatCompletionRequest) HasTopLogprobs() bool`

HasTopLogprobs returns a boolean if a field has been set.

### SetTopLogprobsNil

`func (o *ChatCompletionRequest) SetTopLogprobsNil(b bool)`

 SetTopLogprobsNil sets the value for TopLogprobs to be an explicit nil

### UnsetTopLogprobs
`func (o *ChatCompletionRequest) UnsetTopLogprobs()`

UnsetTopLogprobs ensures that no value is present for TopLogprobs, not even an explicit nil
### GetPromptLogprobs

`func (o *ChatCompletionRequest) GetPromptLogprobs() bool`

GetPromptLogprobs returns the PromptLogprobs field if non-nil, zero value otherwise.

### GetPromptLogprobsOk

`func (o *ChatCompletionRequest) GetPromptLogprobsOk() (*bool, bool)`

GetPromptLogprobsOk returns a tuple with the PromptLogprobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptLogprobs

`func (o *ChatCompletionRequest) SetPromptLogprobs(v bool)`

SetPromptLogprobs sets PromptLogprobs field to given value.

### HasPromptLogprobs

`func (o *ChatCompletionRequest) HasPromptLogprobs() bool`

HasPromptLogprobs returns a boolean if a field has been set.

### SetPromptLogprobsNil

`func (o *ChatCompletionRequest) SetPromptLogprobsNil(b bool)`

 SetPromptLogprobsNil sets the value for PromptLogprobs to be an explicit nil

### UnsetPromptLogprobs
`func (o *ChatCompletionRequest) UnsetPromptLogprobs()`

UnsetPromptLogprobs ensures that no value is present for PromptLogprobs, not even an explicit nil
### GetTopPromptLogprobs

`func (o *ChatCompletionRequest) GetTopPromptLogprobs() int32`

GetTopPromptLogprobs returns the TopPromptLogprobs field if non-nil, zero value otherwise.

### GetTopPromptLogprobsOk

`func (o *ChatCompletionRequest) GetTopPromptLogprobsOk() (*int32, bool)`

GetTopPromptLogprobsOk returns a tuple with the TopPromptLogprobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopPromptLogprobs

`func (o *ChatCompletionRequest) SetTopPromptLogprobs(v int32)`

SetTopPromptLogprobs sets TopPromptLogprobs field to given value.

### HasTopPromptLogprobs

`func (o *ChatCompletionRequest) HasTopPromptLogprobs() bool`

HasTopPromptLogprobs returns a boolean if a field has been set.

### SetTopPromptLogprobsNil

`func (o *ChatCompletionRequest) SetTopPromptLogprobsNil(b bool)`

 SetTopPromptLogprobsNil sets the value for TopPromptLogprobs to be an explicit nil

### UnsetTopPromptLogprobs
`func (o *ChatCompletionRequest) UnsetTopPromptLogprobs()`

UnsetTopPromptLogprobs ensures that no value is present for TopPromptLogprobs, not even an explicit nil
### GetN

`func (o *ChatCompletionRequest) GetN() int32`

GetN returns the N field if non-nil, zero value otherwise.

### GetNOk

`func (o *ChatCompletionRequest) GetNOk() (*int32, bool)`

GetNOk returns a tuple with the N field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN

`func (o *ChatCompletionRequest) SetN(v int32)`

SetN sets N field to given value.

### HasN

`func (o *ChatCompletionRequest) HasN() bool`

HasN returns a boolean if a field has been set.

### SetNNil

`func (o *ChatCompletionRequest) SetNNil(b bool)`

 SetNNil sets the value for N to be an explicit nil

### UnsetN
`func (o *ChatCompletionRequest) UnsetN()`

UnsetN ensures that no value is present for N, not even an explicit nil
### GetPrediction

`func (o *ChatCompletionRequest) GetPrediction() Prediction`

GetPrediction returns the Prediction field if non-nil, zero value otherwise.

### GetPredictionOk

`func (o *ChatCompletionRequest) GetPredictionOk() (*Prediction, bool)`

GetPredictionOk returns a tuple with the Prediction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrediction

`func (o *ChatCompletionRequest) SetPrediction(v Prediction)`

SetPrediction sets Prediction field to given value.

### HasPrediction

`func (o *ChatCompletionRequest) HasPrediction() bool`

HasPrediction returns a boolean if a field has been set.

### GetParallelToolCalls

`func (o *ChatCompletionRequest) GetParallelToolCalls() bool`

GetParallelToolCalls returns the ParallelToolCalls field if non-nil, zero value otherwise.

### GetParallelToolCallsOk

`func (o *ChatCompletionRequest) GetParallelToolCallsOk() (*bool, bool)`

GetParallelToolCallsOk returns a tuple with the ParallelToolCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParallelToolCalls

`func (o *ChatCompletionRequest) SetParallelToolCalls(v bool)`

SetParallelToolCalls sets ParallelToolCalls field to given value.

### HasParallelToolCalls

`func (o *ChatCompletionRequest) HasParallelToolCalls() bool`

HasParallelToolCalls returns a boolean if a field has been set.

### GetReasoningEffort

`func (o *ChatCompletionRequest) GetReasoningEffort() ReasoningEffort`

GetReasoningEffort returns the ReasoningEffort field if non-nil, zero value otherwise.

### GetReasoningEffortOk

`func (o *ChatCompletionRequest) GetReasoningEffortOk() (*ReasoningEffort, bool)`

GetReasoningEffortOk returns a tuple with the ReasoningEffort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasoningEffort

`func (o *ChatCompletionRequest) SetReasoningEffort(v ReasoningEffort)`

SetReasoningEffort sets ReasoningEffort field to given value.

### HasReasoningEffort

`func (o *ChatCompletionRequest) HasReasoningEffort() bool`

HasReasoningEffort returns a boolean if a field has been set.

### SetReasoningEffortNil

`func (o *ChatCompletionRequest) SetReasoningEffortNil(b bool)`

 SetReasoningEffortNil sets the value for ReasoningEffort to be an explicit nil

### UnsetReasoningEffort
`func (o *ChatCompletionRequest) UnsetReasoningEffort()`

UnsetReasoningEffort ensures that no value is present for ReasoningEffort, not even an explicit nil
### GetPromptMode

`func (o *ChatCompletionRequest) GetPromptMode() MistralPromptMode`

GetPromptMode returns the PromptMode field if non-nil, zero value otherwise.

### GetPromptModeOk

`func (o *ChatCompletionRequest) GetPromptModeOk() (*MistralPromptMode, bool)`

GetPromptModeOk returns a tuple with the PromptMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptMode

`func (o *ChatCompletionRequest) SetPromptMode(v MistralPromptMode)`

SetPromptMode sets PromptMode field to given value.

### HasPromptMode

`func (o *ChatCompletionRequest) HasPromptMode() bool`

HasPromptMode returns a boolean if a field has been set.

### SetPromptModeNil

`func (o *ChatCompletionRequest) SetPromptModeNil(b bool)`

 SetPromptModeNil sets the value for PromptMode to be an explicit nil

### UnsetPromptMode
`func (o *ChatCompletionRequest) UnsetPromptMode()`

UnsetPromptMode ensures that no value is present for PromptMode, not even an explicit nil
### GetGuardrails

`func (o *ChatCompletionRequest) GetGuardrails() []GuardrailConfig`

GetGuardrails returns the Guardrails field if non-nil, zero value otherwise.

### GetGuardrailsOk

`func (o *ChatCompletionRequest) GetGuardrailsOk() (*[]GuardrailConfig, bool)`

GetGuardrailsOk returns a tuple with the Guardrails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuardrails

`func (o *ChatCompletionRequest) SetGuardrails(v []GuardrailConfig)`

SetGuardrails sets Guardrails field to given value.

### HasGuardrails

`func (o *ChatCompletionRequest) HasGuardrails() bool`

HasGuardrails returns a boolean if a field has been set.

### SetGuardrailsNil

`func (o *ChatCompletionRequest) SetGuardrailsNil(b bool)`

 SetGuardrailsNil sets the value for Guardrails to be an explicit nil

### UnsetGuardrails
`func (o *ChatCompletionRequest) UnsetGuardrails()`

UnsetGuardrails ensures that no value is present for Guardrails, not even an explicit nil
### GetPromptCacheKey

`func (o *ChatCompletionRequest) GetPromptCacheKey() string`

GetPromptCacheKey returns the PromptCacheKey field if non-nil, zero value otherwise.

### GetPromptCacheKeyOk

`func (o *ChatCompletionRequest) GetPromptCacheKeyOk() (*string, bool)`

GetPromptCacheKeyOk returns a tuple with the PromptCacheKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptCacheKey

`func (o *ChatCompletionRequest) SetPromptCacheKey(v string)`

SetPromptCacheKey sets PromptCacheKey field to given value.

### HasPromptCacheKey

`func (o *ChatCompletionRequest) HasPromptCacheKey() bool`

HasPromptCacheKey returns a boolean if a field has been set.

### SetPromptCacheKeyNil

`func (o *ChatCompletionRequest) SetPromptCacheKeyNil(b bool)`

 SetPromptCacheKeyNil sets the value for PromptCacheKey to be an explicit nil

### UnsetPromptCacheKey
`func (o *ChatCompletionRequest) UnsetPromptCacheKey()`

UnsetPromptCacheKey ensures that no value is present for PromptCacheKey, not even an explicit nil
### GetServiceTier

`func (o *ChatCompletionRequest) GetServiceTier() RequestedServiceTier`

GetServiceTier returns the ServiceTier field if non-nil, zero value otherwise.

### GetServiceTierOk

`func (o *ChatCompletionRequest) GetServiceTierOk() (*RequestedServiceTier, bool)`

GetServiceTierOk returns a tuple with the ServiceTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTier

`func (o *ChatCompletionRequest) SetServiceTier(v RequestedServiceTier)`

SetServiceTier sets ServiceTier field to given value.

### HasServiceTier

`func (o *ChatCompletionRequest) HasServiceTier() bool`

HasServiceTier returns a boolean if a field has been set.

### SetServiceTierNil

`func (o *ChatCompletionRequest) SetServiceTierNil(b bool)`

 SetServiceTierNil sets the value for ServiceTier to be an explicit nil

### UnsetServiceTier
`func (o *ChatCompletionRequest) UnsetServiceTier()`

UnsetServiceTier ensures that no value is present for ServiceTier, not even an explicit nil
### GetSafePrompt

`func (o *ChatCompletionRequest) GetSafePrompt() bool`

GetSafePrompt returns the SafePrompt field if non-nil, zero value otherwise.

### GetSafePromptOk

`func (o *ChatCompletionRequest) GetSafePromptOk() (*bool, bool)`

GetSafePromptOk returns a tuple with the SafePrompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafePrompt

`func (o *ChatCompletionRequest) SetSafePrompt(v bool)`

SetSafePrompt sets SafePrompt field to given value.

### HasSafePrompt

`func (o *ChatCompletionRequest) HasSafePrompt() bool`

HasSafePrompt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


