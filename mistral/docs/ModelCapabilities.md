# ModelCapabilities

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompletionChat** | Pointer to **bool** |  | [optional] [default to false]
**FunctionCalling** | Pointer to **bool** |  | [optional] [default to false]
**CompletionFim** | Pointer to **bool** |  | [optional] [default to false]
**FineTuning** | Pointer to **bool** |  | [optional] [default to false]
**Vision** | Pointer to **bool** |  | [optional] [default to false]
**Ocr** | Pointer to **bool** |  | [optional] [default to false]
**Classification** | Pointer to **bool** |  | [optional] [default to false]
**Moderation** | Pointer to **bool** |  | [optional] [default to false]
**Audio** | Pointer to **bool** |  | [optional] [default to false]
**AudioTranscription** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewModelCapabilities

`func NewModelCapabilities() *ModelCapabilities`

NewModelCapabilities instantiates a new ModelCapabilities object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelCapabilitiesWithDefaults

`func NewModelCapabilitiesWithDefaults() *ModelCapabilities`

NewModelCapabilitiesWithDefaults instantiates a new ModelCapabilities object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompletionChat

`func (o *ModelCapabilities) GetCompletionChat() bool`

GetCompletionChat returns the CompletionChat field if non-nil, zero value otherwise.

### GetCompletionChatOk

`func (o *ModelCapabilities) GetCompletionChatOk() (*bool, bool)`

GetCompletionChatOk returns a tuple with the CompletionChat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionChat

`func (o *ModelCapabilities) SetCompletionChat(v bool)`

SetCompletionChat sets CompletionChat field to given value.

### HasCompletionChat

`func (o *ModelCapabilities) HasCompletionChat() bool`

HasCompletionChat returns a boolean if a field has been set.

### GetFunctionCalling

`func (o *ModelCapabilities) GetFunctionCalling() bool`

GetFunctionCalling returns the FunctionCalling field if non-nil, zero value otherwise.

### GetFunctionCallingOk

`func (o *ModelCapabilities) GetFunctionCallingOk() (*bool, bool)`

GetFunctionCallingOk returns a tuple with the FunctionCalling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionCalling

`func (o *ModelCapabilities) SetFunctionCalling(v bool)`

SetFunctionCalling sets FunctionCalling field to given value.

### HasFunctionCalling

`func (o *ModelCapabilities) HasFunctionCalling() bool`

HasFunctionCalling returns a boolean if a field has been set.

### GetCompletionFim

`func (o *ModelCapabilities) GetCompletionFim() bool`

GetCompletionFim returns the CompletionFim field if non-nil, zero value otherwise.

### GetCompletionFimOk

`func (o *ModelCapabilities) GetCompletionFimOk() (*bool, bool)`

GetCompletionFimOk returns a tuple with the CompletionFim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionFim

`func (o *ModelCapabilities) SetCompletionFim(v bool)`

SetCompletionFim sets CompletionFim field to given value.

### HasCompletionFim

`func (o *ModelCapabilities) HasCompletionFim() bool`

HasCompletionFim returns a boolean if a field has been set.

### GetFineTuning

`func (o *ModelCapabilities) GetFineTuning() bool`

GetFineTuning returns the FineTuning field if non-nil, zero value otherwise.

### GetFineTuningOk

`func (o *ModelCapabilities) GetFineTuningOk() (*bool, bool)`

GetFineTuningOk returns a tuple with the FineTuning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFineTuning

`func (o *ModelCapabilities) SetFineTuning(v bool)`

SetFineTuning sets FineTuning field to given value.

### HasFineTuning

`func (o *ModelCapabilities) HasFineTuning() bool`

HasFineTuning returns a boolean if a field has been set.

### GetVision

`func (o *ModelCapabilities) GetVision() bool`

GetVision returns the Vision field if non-nil, zero value otherwise.

### GetVisionOk

`func (o *ModelCapabilities) GetVisionOk() (*bool, bool)`

GetVisionOk returns a tuple with the Vision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVision

`func (o *ModelCapabilities) SetVision(v bool)`

SetVision sets Vision field to given value.

### HasVision

`func (o *ModelCapabilities) HasVision() bool`

HasVision returns a boolean if a field has been set.

### GetOcr

`func (o *ModelCapabilities) GetOcr() bool`

GetOcr returns the Ocr field if non-nil, zero value otherwise.

### GetOcrOk

`func (o *ModelCapabilities) GetOcrOk() (*bool, bool)`

GetOcrOk returns a tuple with the Ocr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcr

`func (o *ModelCapabilities) SetOcr(v bool)`

SetOcr sets Ocr field to given value.

### HasOcr

`func (o *ModelCapabilities) HasOcr() bool`

HasOcr returns a boolean if a field has been set.

### GetClassification

`func (o *ModelCapabilities) GetClassification() bool`

GetClassification returns the Classification field if non-nil, zero value otherwise.

### GetClassificationOk

`func (o *ModelCapabilities) GetClassificationOk() (*bool, bool)`

GetClassificationOk returns a tuple with the Classification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassification

`func (o *ModelCapabilities) SetClassification(v bool)`

SetClassification sets Classification field to given value.

### HasClassification

`func (o *ModelCapabilities) HasClassification() bool`

HasClassification returns a boolean if a field has been set.

### GetModeration

`func (o *ModelCapabilities) GetModeration() bool`

GetModeration returns the Moderation field if non-nil, zero value otherwise.

### GetModerationOk

`func (o *ModelCapabilities) GetModerationOk() (*bool, bool)`

GetModerationOk returns a tuple with the Moderation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModeration

`func (o *ModelCapabilities) SetModeration(v bool)`

SetModeration sets Moderation field to given value.

### HasModeration

`func (o *ModelCapabilities) HasModeration() bool`

HasModeration returns a boolean if a field has been set.

### GetAudio

`func (o *ModelCapabilities) GetAudio() bool`

GetAudio returns the Audio field if non-nil, zero value otherwise.

### GetAudioOk

`func (o *ModelCapabilities) GetAudioOk() (*bool, bool)`

GetAudioOk returns a tuple with the Audio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudio

`func (o *ModelCapabilities) SetAudio(v bool)`

SetAudio sets Audio field to given value.

### HasAudio

`func (o *ModelCapabilities) HasAudio() bool`

HasAudio returns a boolean if a field has been set.

### GetAudioTranscription

`func (o *ModelCapabilities) GetAudioTranscription() bool`

GetAudioTranscription returns the AudioTranscription field if non-nil, zero value otherwise.

### GetAudioTranscriptionOk

`func (o *ModelCapabilities) GetAudioTranscriptionOk() (*bool, bool)`

GetAudioTranscriptionOk returns a tuple with the AudioTranscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioTranscription

`func (o *ModelCapabilities) SetAudioTranscription(v bool)`

SetAudioTranscription sets AudioTranscription field to given value.

### HasAudioTranscription

`func (o *ModelCapabilities) HasAudioTranscription() bool`

HasAudioTranscription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


