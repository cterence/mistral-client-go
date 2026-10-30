# SpeechRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | Pointer to **NullableString** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Custom type for metadata with embedded validation. | [optional] 
**Stream** | Pointer to **bool** |  | [optional] [default to false]
**PromptCacheKey** | Pointer to **NullableString** |  | [optional] 
**VoiceId** | Pointer to **NullableString** | The preset or custom voice to use for generating the speech. | [optional] 
**RefAudio** | Pointer to [**NullableRefAudio**](RefAudio.md) |  | [optional] 
**Input** | **string** | Text to generate a speech from | 
**PhonemeHints** | Pointer to [**map[string][]IPAPhoneme**](array.md) | Optional pronunciation hints mapping a (case-sensitive) word to its ordered list of IPA phones. | [optional] 
**ResponseFormat** | Pointer to [**SpeechOutputFormat**](SpeechOutputFormat.md) | Output audio format. Defaults to mp3. | [optional] [default to SPEECHOUTPUTFORMAT_MP3]

## Methods

### NewSpeechRequest

`func NewSpeechRequest(input string, ) *SpeechRequest`

NewSpeechRequest instantiates a new SpeechRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpeechRequestWithDefaults

`func NewSpeechRequestWithDefaults() *SpeechRequest`

NewSpeechRequestWithDefaults instantiates a new SpeechRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *SpeechRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *SpeechRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *SpeechRequest) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *SpeechRequest) HasModel() bool`

HasModel returns a boolean if a field has been set.

### SetModelNil

`func (o *SpeechRequest) SetModelNil(b bool)`

 SetModelNil sets the value for Model to be an explicit nil

### UnsetModel
`func (o *SpeechRequest) UnsetModel()`

UnsetModel ensures that no value is present for Model, not even an explicit nil
### GetMetadata

`func (o *SpeechRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SpeechRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SpeechRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SpeechRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *SpeechRequest) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *SpeechRequest) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetStream

`func (o *SpeechRequest) GetStream() bool`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *SpeechRequest) GetStreamOk() (*bool, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *SpeechRequest) SetStream(v bool)`

SetStream sets Stream field to given value.

### HasStream

`func (o *SpeechRequest) HasStream() bool`

HasStream returns a boolean if a field has been set.

### GetPromptCacheKey

`func (o *SpeechRequest) GetPromptCacheKey() string`

GetPromptCacheKey returns the PromptCacheKey field if non-nil, zero value otherwise.

### GetPromptCacheKeyOk

`func (o *SpeechRequest) GetPromptCacheKeyOk() (*string, bool)`

GetPromptCacheKeyOk returns a tuple with the PromptCacheKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptCacheKey

`func (o *SpeechRequest) SetPromptCacheKey(v string)`

SetPromptCacheKey sets PromptCacheKey field to given value.

### HasPromptCacheKey

`func (o *SpeechRequest) HasPromptCacheKey() bool`

HasPromptCacheKey returns a boolean if a field has been set.

### SetPromptCacheKeyNil

`func (o *SpeechRequest) SetPromptCacheKeyNil(b bool)`

 SetPromptCacheKeyNil sets the value for PromptCacheKey to be an explicit nil

### UnsetPromptCacheKey
`func (o *SpeechRequest) UnsetPromptCacheKey()`

UnsetPromptCacheKey ensures that no value is present for PromptCacheKey, not even an explicit nil
### GetVoiceId

`func (o *SpeechRequest) GetVoiceId() string`

GetVoiceId returns the VoiceId field if non-nil, zero value otherwise.

### GetVoiceIdOk

`func (o *SpeechRequest) GetVoiceIdOk() (*string, bool)`

GetVoiceIdOk returns a tuple with the VoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceId

`func (o *SpeechRequest) SetVoiceId(v string)`

SetVoiceId sets VoiceId field to given value.

### HasVoiceId

`func (o *SpeechRequest) HasVoiceId() bool`

HasVoiceId returns a boolean if a field has been set.

### SetVoiceIdNil

`func (o *SpeechRequest) SetVoiceIdNil(b bool)`

 SetVoiceIdNil sets the value for VoiceId to be an explicit nil

### UnsetVoiceId
`func (o *SpeechRequest) UnsetVoiceId()`

UnsetVoiceId ensures that no value is present for VoiceId, not even an explicit nil
### GetRefAudio

`func (o *SpeechRequest) GetRefAudio() RefAudio`

GetRefAudio returns the RefAudio field if non-nil, zero value otherwise.

### GetRefAudioOk

`func (o *SpeechRequest) GetRefAudioOk() (*RefAudio, bool)`

GetRefAudioOk returns a tuple with the RefAudio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefAudio

`func (o *SpeechRequest) SetRefAudio(v RefAudio)`

SetRefAudio sets RefAudio field to given value.

### HasRefAudio

`func (o *SpeechRequest) HasRefAudio() bool`

HasRefAudio returns a boolean if a field has been set.

### SetRefAudioNil

`func (o *SpeechRequest) SetRefAudioNil(b bool)`

 SetRefAudioNil sets the value for RefAudio to be an explicit nil

### UnsetRefAudio
`func (o *SpeechRequest) UnsetRefAudio()`

UnsetRefAudio ensures that no value is present for RefAudio, not even an explicit nil
### GetInput

`func (o *SpeechRequest) GetInput() string`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *SpeechRequest) GetInputOk() (*string, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *SpeechRequest) SetInput(v string)`

SetInput sets Input field to given value.


### GetPhonemeHints

`func (o *SpeechRequest) GetPhonemeHints() map[string][]IPAPhoneme`

GetPhonemeHints returns the PhonemeHints field if non-nil, zero value otherwise.

### GetPhonemeHintsOk

`func (o *SpeechRequest) GetPhonemeHintsOk() (*map[string][]IPAPhoneme, bool)`

GetPhonemeHintsOk returns a tuple with the PhonemeHints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhonemeHints

`func (o *SpeechRequest) SetPhonemeHints(v map[string][]IPAPhoneme)`

SetPhonemeHints sets PhonemeHints field to given value.

### HasPhonemeHints

`func (o *SpeechRequest) HasPhonemeHints() bool`

HasPhonemeHints returns a boolean if a field has been set.

### SetPhonemeHintsNil

`func (o *SpeechRequest) SetPhonemeHintsNil(b bool)`

 SetPhonemeHintsNil sets the value for PhonemeHints to be an explicit nil

### UnsetPhonemeHints
`func (o *SpeechRequest) UnsetPhonemeHints()`

UnsetPhonemeHints ensures that no value is present for PhonemeHints, not even an explicit nil
### GetResponseFormat

`func (o *SpeechRequest) GetResponseFormat() SpeechOutputFormat`

GetResponseFormat returns the ResponseFormat field if non-nil, zero value otherwise.

### GetResponseFormatOk

`func (o *SpeechRequest) GetResponseFormatOk() (*SpeechOutputFormat, bool)`

GetResponseFormatOk returns a tuple with the ResponseFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseFormat

`func (o *SpeechRequest) SetResponseFormat(v SpeechOutputFormat)`

SetResponseFormat sets ResponseFormat field to given value.

### HasResponseFormat

`func (o *SpeechRequest) HasResponseFormat() bool`

HasResponseFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


