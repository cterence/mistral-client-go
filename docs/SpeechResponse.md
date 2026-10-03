# SpeechResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AudioData** | **string** | Base64 encoded audio data | 

## Methods

### NewSpeechResponse

`func NewSpeechResponse(audioData string, ) *SpeechResponse`

NewSpeechResponse instantiates a new SpeechResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpeechResponseWithDefaults

`func NewSpeechResponseWithDefaults() *SpeechResponse`

NewSpeechResponseWithDefaults instantiates a new SpeechResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudioData

`func (o *SpeechResponse) GetAudioData() string`

GetAudioData returns the AudioData field if non-nil, zero value otherwise.

### GetAudioDataOk

`func (o *SpeechResponse) GetAudioDataOk() (*string, bool)`

GetAudioDataOk returns a tuple with the AudioData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioData

`func (o *SpeechResponse) SetAudioData(v string)`

SetAudioData sets AudioData field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


