# SpeechRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | Pointer to **NullableString** |  | [optional] 
**Stream** | Pointer to **bool** |  | [optional] [default to false]
**VoiceId** | Pointer to **NullableString** | The preset or custom voice to use for generating the speech. | [optional] 
**RefAudio** | Pointer to **NullableString** | The base64-encoded audio reference for zero-shot voice cloning. | [optional] 
**Input** | **string** | Text to generate speech from. | 
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

`func (o *SpeechRequest) GetRefAudio() string`

GetRefAudio returns the RefAudio field if non-nil, zero value otherwise.

### GetRefAudioOk

`func (o *SpeechRequest) GetRefAudioOk() (*string, bool)`

GetRefAudioOk returns a tuple with the RefAudio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefAudio

`func (o *SpeechRequest) SetRefAudio(v string)`

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


