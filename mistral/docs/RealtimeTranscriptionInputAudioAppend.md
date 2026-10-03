# RealtimeTranscriptionInputAudioAppend

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "input_audio.append"]
**Audio** | **string** | Base64-encoded raw PCM bytes matching the current audio_format. Max decoded size: 262144 bytes. | 

## Methods

### NewRealtimeTranscriptionInputAudioAppend

`func NewRealtimeTranscriptionInputAudioAppend(audio string, ) *RealtimeTranscriptionInputAudioAppend`

NewRealtimeTranscriptionInputAudioAppend instantiates a new RealtimeTranscriptionInputAudioAppend object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRealtimeTranscriptionInputAudioAppendWithDefaults

`func NewRealtimeTranscriptionInputAudioAppendWithDefaults() *RealtimeTranscriptionInputAudioAppend`

NewRealtimeTranscriptionInputAudioAppendWithDefaults instantiates a new RealtimeTranscriptionInputAudioAppend object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RealtimeTranscriptionInputAudioAppend) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RealtimeTranscriptionInputAudioAppend) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RealtimeTranscriptionInputAudioAppend) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RealtimeTranscriptionInputAudioAppend) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAudio

`func (o *RealtimeTranscriptionInputAudioAppend) GetAudio() string`

GetAudio returns the Audio field if non-nil, zero value otherwise.

### GetAudioOk

`func (o *RealtimeTranscriptionInputAudioAppend) GetAudioOk() (*string, bool)`

GetAudioOk returns a tuple with the Audio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudio

`func (o *RealtimeTranscriptionInputAudioAppend) SetAudio(v string)`

SetAudio sets Audio field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


