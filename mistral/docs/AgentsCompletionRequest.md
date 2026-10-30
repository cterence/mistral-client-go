# AgentsCompletionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxTokens** | Pointer to **NullableInt32** | The maximum number of tokens to generate in the completion. The token count of your prompt plus &#x60;max_tokens&#x60; cannot exceed the model&#39;s context length. | [optional] 
**Stream** | Pointer to **bool** | Whether to stream back partial progress. If set, tokens will be sent as data-only server-side events as they become available, with the stream terminated by a data: [DONE] message. Otherwise, the server will hold the request open until the timeout or until completion, with the response containing the full result as JSON. | [optional] [default to false]
**Stop** | Pointer to [**NullableStop**](Stop.md) |  | [optional] 
**RandomSeed** | Pointer to **NullableInt32** | The seed to use for random sampling. If set, different calls will generate deterministic results. | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**Messages** | [**[]MessagesInner**](MessagesInner.md) | The prompt(s) to generate completions for, encoded as a list of dict with role and content. | 
**ResponseFormat** | Pointer to [**ResponseFormat**](ResponseFormat.md) |  | [optional] 
**Tools** | Pointer to [**[]ChatCompletionRequestToolsInner**](ChatCompletionRequestToolsInner.md) |  | [optional] 
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
**ParallelToolCalls** | Pointer to **bool** |  | [optional] [default to true]
**ReasoningEffort** | Pointer to [**NullableReasoningEffort**](ReasoningEffort.md) |  | [optional] 
**PromptMode** | Pointer to [**NullableMistralPromptMode**](MistralPromptMode.md) | Allows toggling between the reasoning mode and no system prompt. When set to &#x60;reasoning&#x60; the system prompt for reasoning models will be used. | [optional] 
**Guardrails** | Pointer to [**[]GuardrailConfig**](GuardrailConfig.md) |  | [optional] 
**PromptCacheKey** | Pointer to **NullableString** |  | [optional] 
**ServiceTier** | Pointer to [**NullableRequestedServiceTier**](RequestedServiceTier.md) | Determines whether to serve the request using priority or standard capacity. | [optional] 
**AgentId** | **string** | The ID of the agent to use for this completion. | 

## Methods

### NewAgentsCompletionRequest

`func NewAgentsCompletionRequest(messages []MessagesInner, agentId string, ) *AgentsCompletionRequest`

NewAgentsCompletionRequest instantiates a new AgentsCompletionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentsCompletionRequestWithDefaults

`func NewAgentsCompletionRequestWithDefaults() *AgentsCompletionRequest`

NewAgentsCompletionRequestWithDefaults instantiates a new AgentsCompletionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxTokens

`func (o *AgentsCompletionRequest) GetMaxTokens() int32`

GetMaxTokens returns the MaxTokens field if non-nil, zero value otherwise.

### GetMaxTokensOk

`func (o *AgentsCompletionRequest) GetMaxTokensOk() (*int32, bool)`

GetMaxTokensOk returns a tuple with the MaxTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTokens

`func (o *AgentsCompletionRequest) SetMaxTokens(v int32)`

SetMaxTokens sets MaxTokens field to given value.

### HasMaxTokens

`func (o *AgentsCompletionRequest) HasMaxTokens() bool`

HasMaxTokens returns a boolean if a field has been set.

### SetMaxTokensNil

`func (o *AgentsCompletionRequest) SetMaxTokensNil(b bool)`

 SetMaxTokensNil sets the value for MaxTokens to be an explicit nil

### UnsetMaxTokens
`func (o *AgentsCompletionRequest) UnsetMaxTokens()`

UnsetMaxTokens ensures that no value is present for MaxTokens, not even an explicit nil
### GetStream

`func (o *AgentsCompletionRequest) GetStream() bool`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *AgentsCompletionRequest) GetStreamOk() (*bool, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *AgentsCompletionRequest) SetStream(v bool)`

SetStream sets Stream field to given value.

### HasStream

`func (o *AgentsCompletionRequest) HasStream() bool`

HasStream returns a boolean if a field has been set.

### GetStop

`func (o *AgentsCompletionRequest) GetStop() Stop`

GetStop returns the Stop field if non-nil, zero value otherwise.

### GetStopOk

`func (o *AgentsCompletionRequest) GetStopOk() (*Stop, bool)`

GetStopOk returns a tuple with the Stop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStop

`func (o *AgentsCompletionRequest) SetStop(v Stop)`

SetStop sets Stop field to given value.

### HasStop

`func (o *AgentsCompletionRequest) HasStop() bool`

HasStop returns a boolean if a field has been set.

### SetStopNil

`func (o *AgentsCompletionRequest) SetStopNil(b bool)`

 SetStopNil sets the value for Stop to be an explicit nil

### UnsetStop
`func (o *AgentsCompletionRequest) UnsetStop()`

UnsetStop ensures that no value is present for Stop, not even an explicit nil
### GetRandomSeed

`func (o *AgentsCompletionRequest) GetRandomSeed() int32`

GetRandomSeed returns the RandomSeed field if non-nil, zero value otherwise.

### GetRandomSeedOk

`func (o *AgentsCompletionRequest) GetRandomSeedOk() (*int32, bool)`

GetRandomSeedOk returns a tuple with the RandomSeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRandomSeed

`func (o *AgentsCompletionRequest) SetRandomSeed(v int32)`

SetRandomSeed sets RandomSeed field to given value.

### HasRandomSeed

`func (o *AgentsCompletionRequest) HasRandomSeed() bool`

HasRandomSeed returns a boolean if a field has been set.

### SetRandomSeedNil

`func (o *AgentsCompletionRequest) SetRandomSeedNil(b bool)`

 SetRandomSeedNil sets the value for RandomSeed to be an explicit nil

### UnsetRandomSeed
`func (o *AgentsCompletionRequest) UnsetRandomSeed()`

UnsetRandomSeed ensures that no value is present for RandomSeed, not even an explicit nil
### GetMetadata

`func (o *AgentsCompletionRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AgentsCompletionRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AgentsCompletionRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AgentsCompletionRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *AgentsCompletionRequest) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *AgentsCompletionRequest) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetMessages

`func (o *AgentsCompletionRequest) GetMessages() []MessagesInner`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *AgentsCompletionRequest) GetMessagesOk() (*[]MessagesInner, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *AgentsCompletionRequest) SetMessages(v []MessagesInner)`

SetMessages sets Messages field to given value.


### GetResponseFormat

`func (o *AgentsCompletionRequest) GetResponseFormat() ResponseFormat`

GetResponseFormat returns the ResponseFormat field if non-nil, zero value otherwise.

### GetResponseFormatOk

`func (o *AgentsCompletionRequest) GetResponseFormatOk() (*ResponseFormat, bool)`

GetResponseFormatOk returns a tuple with the ResponseFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseFormat

`func (o *AgentsCompletionRequest) SetResponseFormat(v ResponseFormat)`

SetResponseFormat sets ResponseFormat field to given value.

### HasResponseFormat

`func (o *AgentsCompletionRequest) HasResponseFormat() bool`

HasResponseFormat returns a boolean if a field has been set.

### GetTools

`func (o *AgentsCompletionRequest) GetTools() []ChatCompletionRequestToolsInner`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *AgentsCompletionRequest) GetToolsOk() (*[]ChatCompletionRequestToolsInner, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *AgentsCompletionRequest) SetTools(v []ChatCompletionRequestToolsInner)`

SetTools sets Tools field to given value.

### HasTools

`func (o *AgentsCompletionRequest) HasTools() bool`

HasTools returns a boolean if a field has been set.

### SetToolsNil

`func (o *AgentsCompletionRequest) SetToolsNil(b bool)`

 SetToolsNil sets the value for Tools to be an explicit nil

### UnsetTools
`func (o *AgentsCompletionRequest) UnsetTools()`

UnsetTools ensures that no value is present for Tools, not even an explicit nil
### GetToolChoice

`func (o *AgentsCompletionRequest) GetToolChoice() ToolChoice`

GetToolChoice returns the ToolChoice field if non-nil, zero value otherwise.

### GetToolChoiceOk

`func (o *AgentsCompletionRequest) GetToolChoiceOk() (*ToolChoice, bool)`

GetToolChoiceOk returns a tuple with the ToolChoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolChoice

`func (o *AgentsCompletionRequest) SetToolChoice(v ToolChoice)`

SetToolChoice sets ToolChoice field to given value.

### HasToolChoice

`func (o *AgentsCompletionRequest) HasToolChoice() bool`

HasToolChoice returns a boolean if a field has been set.

### GetPresencePenalty

`func (o *AgentsCompletionRequest) GetPresencePenalty() float32`

GetPresencePenalty returns the PresencePenalty field if non-nil, zero value otherwise.

### GetPresencePenaltyOk

`func (o *AgentsCompletionRequest) GetPresencePenaltyOk() (*float32, bool)`

GetPresencePenaltyOk returns a tuple with the PresencePenalty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresencePenalty

`func (o *AgentsCompletionRequest) SetPresencePenalty(v float32)`

SetPresencePenalty sets PresencePenalty field to given value.

### HasPresencePenalty

`func (o *AgentsCompletionRequest) HasPresencePenalty() bool`

HasPresencePenalty returns a boolean if a field has been set.

### SetPresencePenaltyNil

`func (o *AgentsCompletionRequest) SetPresencePenaltyNil(b bool)`

 SetPresencePenaltyNil sets the value for PresencePenalty to be an explicit nil

### UnsetPresencePenalty
`func (o *AgentsCompletionRequest) UnsetPresencePenalty()`

UnsetPresencePenalty ensures that no value is present for PresencePenalty, not even an explicit nil
### GetFrequencyPenalty

`func (o *AgentsCompletionRequest) GetFrequencyPenalty() float32`

GetFrequencyPenalty returns the FrequencyPenalty field if non-nil, zero value otherwise.

### GetFrequencyPenaltyOk

`func (o *AgentsCompletionRequest) GetFrequencyPenaltyOk() (*float32, bool)`

GetFrequencyPenaltyOk returns a tuple with the FrequencyPenalty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencyPenalty

`func (o *AgentsCompletionRequest) SetFrequencyPenalty(v float32)`

SetFrequencyPenalty sets FrequencyPenalty field to given value.

### HasFrequencyPenalty

`func (o *AgentsCompletionRequest) HasFrequencyPenalty() bool`

HasFrequencyPenalty returns a boolean if a field has been set.

### SetFrequencyPenaltyNil

`func (o *AgentsCompletionRequest) SetFrequencyPenaltyNil(b bool)`

 SetFrequencyPenaltyNil sets the value for FrequencyPenalty to be an explicit nil

### UnsetFrequencyPenalty
`func (o *AgentsCompletionRequest) UnsetFrequencyPenalty()`

UnsetFrequencyPenalty ensures that no value is present for FrequencyPenalty, not even an explicit nil
### GetMinTokens

`func (o *AgentsCompletionRequest) GetMinTokens() int32`

GetMinTokens returns the MinTokens field if non-nil, zero value otherwise.

### GetMinTokensOk

`func (o *AgentsCompletionRequest) GetMinTokensOk() (*int32, bool)`

GetMinTokensOk returns a tuple with the MinTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinTokens

`func (o *AgentsCompletionRequest) SetMinTokens(v int32)`

SetMinTokens sets MinTokens field to given value.

### HasMinTokens

`func (o *AgentsCompletionRequest) HasMinTokens() bool`

HasMinTokens returns a boolean if a field has been set.

### SetMinTokensNil

`func (o *AgentsCompletionRequest) SetMinTokensNil(b bool)`

 SetMinTokensNil sets the value for MinTokens to be an explicit nil

### UnsetMinTokens
`func (o *AgentsCompletionRequest) UnsetMinTokens()`

UnsetMinTokens ensures that no value is present for MinTokens, not even an explicit nil
### GetRepetitionPenalty

`func (o *AgentsCompletionRequest) GetRepetitionPenalty() float32`

GetRepetitionPenalty returns the RepetitionPenalty field if non-nil, zero value otherwise.

### GetRepetitionPenaltyOk

`func (o *AgentsCompletionRequest) GetRepetitionPenaltyOk() (*float32, bool)`

GetRepetitionPenaltyOk returns a tuple with the RepetitionPenalty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepetitionPenalty

`func (o *AgentsCompletionRequest) SetRepetitionPenalty(v float32)`

SetRepetitionPenalty sets RepetitionPenalty field to given value.

### HasRepetitionPenalty

`func (o *AgentsCompletionRequest) HasRepetitionPenalty() bool`

HasRepetitionPenalty returns a boolean if a field has been set.

### SetRepetitionPenaltyNil

`func (o *AgentsCompletionRequest) SetRepetitionPenaltyNil(b bool)`

 SetRepetitionPenaltyNil sets the value for RepetitionPenalty to be an explicit nil

### UnsetRepetitionPenalty
`func (o *AgentsCompletionRequest) UnsetRepetitionPenalty()`

UnsetRepetitionPenalty ensures that no value is present for RepetitionPenalty, not even an explicit nil
### GetTopK

`func (o *AgentsCompletionRequest) GetTopK() int32`

GetTopK returns the TopK field if non-nil, zero value otherwise.

### GetTopKOk

`func (o *AgentsCompletionRequest) GetTopKOk() (*int32, bool)`

GetTopKOk returns a tuple with the TopK field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopK

`func (o *AgentsCompletionRequest) SetTopK(v int32)`

SetTopK sets TopK field to given value.

### HasTopK

`func (o *AgentsCompletionRequest) HasTopK() bool`

HasTopK returns a boolean if a field has been set.

### SetTopKNil

`func (o *AgentsCompletionRequest) SetTopKNil(b bool)`

 SetTopKNil sets the value for TopK to be an explicit nil

### UnsetTopK
`func (o *AgentsCompletionRequest) UnsetTopK()`

UnsetTopK ensures that no value is present for TopK, not even an explicit nil
### GetLogprobs

`func (o *AgentsCompletionRequest) GetLogprobs() bool`

GetLogprobs returns the Logprobs field if non-nil, zero value otherwise.

### GetLogprobsOk

`func (o *AgentsCompletionRequest) GetLogprobsOk() (*bool, bool)`

GetLogprobsOk returns a tuple with the Logprobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogprobs

`func (o *AgentsCompletionRequest) SetLogprobs(v bool)`

SetLogprobs sets Logprobs field to given value.

### HasLogprobs

`func (o *AgentsCompletionRequest) HasLogprobs() bool`

HasLogprobs returns a boolean if a field has been set.

### SetLogprobsNil

`func (o *AgentsCompletionRequest) SetLogprobsNil(b bool)`

 SetLogprobsNil sets the value for Logprobs to be an explicit nil

### UnsetLogprobs
`func (o *AgentsCompletionRequest) UnsetLogprobs()`

UnsetLogprobs ensures that no value is present for Logprobs, not even an explicit nil
### GetTopLogprobs

`func (o *AgentsCompletionRequest) GetTopLogprobs() int32`

GetTopLogprobs returns the TopLogprobs field if non-nil, zero value otherwise.

### GetTopLogprobsOk

`func (o *AgentsCompletionRequest) GetTopLogprobsOk() (*int32, bool)`

GetTopLogprobsOk returns a tuple with the TopLogprobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopLogprobs

`func (o *AgentsCompletionRequest) SetTopLogprobs(v int32)`

SetTopLogprobs sets TopLogprobs field to given value.

### HasTopLogprobs

`func (o *AgentsCompletionRequest) HasTopLogprobs() bool`

HasTopLogprobs returns a boolean if a field has been set.

### SetTopLogprobsNil

`func (o *AgentsCompletionRequest) SetTopLogprobsNil(b bool)`

 SetTopLogprobsNil sets the value for TopLogprobs to be an explicit nil

### UnsetTopLogprobs
`func (o *AgentsCompletionRequest) UnsetTopLogprobs()`

UnsetTopLogprobs ensures that no value is present for TopLogprobs, not even an explicit nil
### GetPromptLogprobs

`func (o *AgentsCompletionRequest) GetPromptLogprobs() bool`

GetPromptLogprobs returns the PromptLogprobs field if non-nil, zero value otherwise.

### GetPromptLogprobsOk

`func (o *AgentsCompletionRequest) GetPromptLogprobsOk() (*bool, bool)`

GetPromptLogprobsOk returns a tuple with the PromptLogprobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptLogprobs

`func (o *AgentsCompletionRequest) SetPromptLogprobs(v bool)`

SetPromptLogprobs sets PromptLogprobs field to given value.

### HasPromptLogprobs

`func (o *AgentsCompletionRequest) HasPromptLogprobs() bool`

HasPromptLogprobs returns a boolean if a field has been set.

### SetPromptLogprobsNil

`func (o *AgentsCompletionRequest) SetPromptLogprobsNil(b bool)`

 SetPromptLogprobsNil sets the value for PromptLogprobs to be an explicit nil

### UnsetPromptLogprobs
`func (o *AgentsCompletionRequest) UnsetPromptLogprobs()`

UnsetPromptLogprobs ensures that no value is present for PromptLogprobs, not even an explicit nil
### GetTopPromptLogprobs

`func (o *AgentsCompletionRequest) GetTopPromptLogprobs() int32`

GetTopPromptLogprobs returns the TopPromptLogprobs field if non-nil, zero value otherwise.

### GetTopPromptLogprobsOk

`func (o *AgentsCompletionRequest) GetTopPromptLogprobsOk() (*int32, bool)`

GetTopPromptLogprobsOk returns a tuple with the TopPromptLogprobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopPromptLogprobs

`func (o *AgentsCompletionRequest) SetTopPromptLogprobs(v int32)`

SetTopPromptLogprobs sets TopPromptLogprobs field to given value.

### HasTopPromptLogprobs

`func (o *AgentsCompletionRequest) HasTopPromptLogprobs() bool`

HasTopPromptLogprobs returns a boolean if a field has been set.

### SetTopPromptLogprobsNil

`func (o *AgentsCompletionRequest) SetTopPromptLogprobsNil(b bool)`

 SetTopPromptLogprobsNil sets the value for TopPromptLogprobs to be an explicit nil

### UnsetTopPromptLogprobs
`func (o *AgentsCompletionRequest) UnsetTopPromptLogprobs()`

UnsetTopPromptLogprobs ensures that no value is present for TopPromptLogprobs, not even an explicit nil
### GetN

`func (o *AgentsCompletionRequest) GetN() int32`

GetN returns the N field if non-nil, zero value otherwise.

### GetNOk

`func (o *AgentsCompletionRequest) GetNOk() (*int32, bool)`

GetNOk returns a tuple with the N field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN

`func (o *AgentsCompletionRequest) SetN(v int32)`

SetN sets N field to given value.

### HasN

`func (o *AgentsCompletionRequest) HasN() bool`

HasN returns a boolean if a field has been set.

### SetNNil

`func (o *AgentsCompletionRequest) SetNNil(b bool)`

 SetNNil sets the value for N to be an explicit nil

### UnsetN
`func (o *AgentsCompletionRequest) UnsetN()`

UnsetN ensures that no value is present for N, not even an explicit nil
### GetPrediction

`func (o *AgentsCompletionRequest) GetPrediction() Prediction`

GetPrediction returns the Prediction field if non-nil, zero value otherwise.

### GetPredictionOk

`func (o *AgentsCompletionRequest) GetPredictionOk() (*Prediction, bool)`

GetPredictionOk returns a tuple with the Prediction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrediction

`func (o *AgentsCompletionRequest) SetPrediction(v Prediction)`

SetPrediction sets Prediction field to given value.

### HasPrediction

`func (o *AgentsCompletionRequest) HasPrediction() bool`

HasPrediction returns a boolean if a field has been set.

### GetParallelToolCalls

`func (o *AgentsCompletionRequest) GetParallelToolCalls() bool`

GetParallelToolCalls returns the ParallelToolCalls field if non-nil, zero value otherwise.

### GetParallelToolCallsOk

`func (o *AgentsCompletionRequest) GetParallelToolCallsOk() (*bool, bool)`

GetParallelToolCallsOk returns a tuple with the ParallelToolCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParallelToolCalls

`func (o *AgentsCompletionRequest) SetParallelToolCalls(v bool)`

SetParallelToolCalls sets ParallelToolCalls field to given value.

### HasParallelToolCalls

`func (o *AgentsCompletionRequest) HasParallelToolCalls() bool`

HasParallelToolCalls returns a boolean if a field has been set.

### GetReasoningEffort

`func (o *AgentsCompletionRequest) GetReasoningEffort() ReasoningEffort`

GetReasoningEffort returns the ReasoningEffort field if non-nil, zero value otherwise.

### GetReasoningEffortOk

`func (o *AgentsCompletionRequest) GetReasoningEffortOk() (*ReasoningEffort, bool)`

GetReasoningEffortOk returns a tuple with the ReasoningEffort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasoningEffort

`func (o *AgentsCompletionRequest) SetReasoningEffort(v ReasoningEffort)`

SetReasoningEffort sets ReasoningEffort field to given value.

### HasReasoningEffort

`func (o *AgentsCompletionRequest) HasReasoningEffort() bool`

HasReasoningEffort returns a boolean if a field has been set.

### SetReasoningEffortNil

`func (o *AgentsCompletionRequest) SetReasoningEffortNil(b bool)`

 SetReasoningEffortNil sets the value for ReasoningEffort to be an explicit nil

### UnsetReasoningEffort
`func (o *AgentsCompletionRequest) UnsetReasoningEffort()`

UnsetReasoningEffort ensures that no value is present for ReasoningEffort, not even an explicit nil
### GetPromptMode

`func (o *AgentsCompletionRequest) GetPromptMode() MistralPromptMode`

GetPromptMode returns the PromptMode field if non-nil, zero value otherwise.

### GetPromptModeOk

`func (o *AgentsCompletionRequest) GetPromptModeOk() (*MistralPromptMode, bool)`

GetPromptModeOk returns a tuple with the PromptMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptMode

`func (o *AgentsCompletionRequest) SetPromptMode(v MistralPromptMode)`

SetPromptMode sets PromptMode field to given value.

### HasPromptMode

`func (o *AgentsCompletionRequest) HasPromptMode() bool`

HasPromptMode returns a boolean if a field has been set.

### SetPromptModeNil

`func (o *AgentsCompletionRequest) SetPromptModeNil(b bool)`

 SetPromptModeNil sets the value for PromptMode to be an explicit nil

### UnsetPromptMode
`func (o *AgentsCompletionRequest) UnsetPromptMode()`

UnsetPromptMode ensures that no value is present for PromptMode, not even an explicit nil
### GetGuardrails

`func (o *AgentsCompletionRequest) GetGuardrails() []GuardrailConfig`

GetGuardrails returns the Guardrails field if non-nil, zero value otherwise.

### GetGuardrailsOk

`func (o *AgentsCompletionRequest) GetGuardrailsOk() (*[]GuardrailConfig, bool)`

GetGuardrailsOk returns a tuple with the Guardrails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuardrails

`func (o *AgentsCompletionRequest) SetGuardrails(v []GuardrailConfig)`

SetGuardrails sets Guardrails field to given value.

### HasGuardrails

`func (o *AgentsCompletionRequest) HasGuardrails() bool`

HasGuardrails returns a boolean if a field has been set.

### SetGuardrailsNil

`func (o *AgentsCompletionRequest) SetGuardrailsNil(b bool)`

 SetGuardrailsNil sets the value for Guardrails to be an explicit nil

### UnsetGuardrails
`func (o *AgentsCompletionRequest) UnsetGuardrails()`

UnsetGuardrails ensures that no value is present for Guardrails, not even an explicit nil
### GetPromptCacheKey

`func (o *AgentsCompletionRequest) GetPromptCacheKey() string`

GetPromptCacheKey returns the PromptCacheKey field if non-nil, zero value otherwise.

### GetPromptCacheKeyOk

`func (o *AgentsCompletionRequest) GetPromptCacheKeyOk() (*string, bool)`

GetPromptCacheKeyOk returns a tuple with the PromptCacheKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptCacheKey

`func (o *AgentsCompletionRequest) SetPromptCacheKey(v string)`

SetPromptCacheKey sets PromptCacheKey field to given value.

### HasPromptCacheKey

`func (o *AgentsCompletionRequest) HasPromptCacheKey() bool`

HasPromptCacheKey returns a boolean if a field has been set.

### SetPromptCacheKeyNil

`func (o *AgentsCompletionRequest) SetPromptCacheKeyNil(b bool)`

 SetPromptCacheKeyNil sets the value for PromptCacheKey to be an explicit nil

### UnsetPromptCacheKey
`func (o *AgentsCompletionRequest) UnsetPromptCacheKey()`

UnsetPromptCacheKey ensures that no value is present for PromptCacheKey, not even an explicit nil
### GetServiceTier

`func (o *AgentsCompletionRequest) GetServiceTier() RequestedServiceTier`

GetServiceTier returns the ServiceTier field if non-nil, zero value otherwise.

### GetServiceTierOk

`func (o *AgentsCompletionRequest) GetServiceTierOk() (*RequestedServiceTier, bool)`

GetServiceTierOk returns a tuple with the ServiceTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTier

`func (o *AgentsCompletionRequest) SetServiceTier(v RequestedServiceTier)`

SetServiceTier sets ServiceTier field to given value.

### HasServiceTier

`func (o *AgentsCompletionRequest) HasServiceTier() bool`

HasServiceTier returns a boolean if a field has been set.

### SetServiceTierNil

`func (o *AgentsCompletionRequest) SetServiceTierNil(b bool)`

 SetServiceTierNil sets the value for ServiceTier to be an explicit nil

### UnsetServiceTier
`func (o *AgentsCompletionRequest) UnsetServiceTier()`

UnsetServiceTier ensures that no value is present for ServiceTier, not even an explicit nil
### GetAgentId

`func (o *AgentsCompletionRequest) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *AgentsCompletionRequest) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *AgentsCompletionRequest) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


