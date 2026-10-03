# TranscriptionStreamSegmentDelta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "transcription.segment"]
**Text** | **string** |  | 
**Start** | **float32** |  | 
**End** | **float32** |  | 
**SpeakerId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTranscriptionStreamSegmentDelta

`func NewTranscriptionStreamSegmentDelta(text string, start float32, end float32, ) *TranscriptionStreamSegmentDelta`

NewTranscriptionStreamSegmentDelta instantiates a new TranscriptionStreamSegmentDelta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTranscriptionStreamSegmentDeltaWithDefaults

`func NewTranscriptionStreamSegmentDeltaWithDefaults() *TranscriptionStreamSegmentDelta`

NewTranscriptionStreamSegmentDeltaWithDefaults instantiates a new TranscriptionStreamSegmentDelta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TranscriptionStreamSegmentDelta) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TranscriptionStreamSegmentDelta) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TranscriptionStreamSegmentDelta) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TranscriptionStreamSegmentDelta) HasType() bool`

HasType returns a boolean if a field has been set.

### GetText

`func (o *TranscriptionStreamSegmentDelta) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *TranscriptionStreamSegmentDelta) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *TranscriptionStreamSegmentDelta) SetText(v string)`

SetText sets Text field to given value.


### GetStart

`func (o *TranscriptionStreamSegmentDelta) GetStart() float32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *TranscriptionStreamSegmentDelta) GetStartOk() (*float32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *TranscriptionStreamSegmentDelta) SetStart(v float32)`

SetStart sets Start field to given value.


### GetEnd

`func (o *TranscriptionStreamSegmentDelta) GetEnd() float32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *TranscriptionStreamSegmentDelta) GetEndOk() (*float32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *TranscriptionStreamSegmentDelta) SetEnd(v float32)`

SetEnd sets End field to given value.


### GetSpeakerId

`func (o *TranscriptionStreamSegmentDelta) GetSpeakerId() string`

GetSpeakerId returns the SpeakerId field if non-nil, zero value otherwise.

### GetSpeakerIdOk

`func (o *TranscriptionStreamSegmentDelta) GetSpeakerIdOk() (*string, bool)`

GetSpeakerIdOk returns a tuple with the SpeakerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeakerId

`func (o *TranscriptionStreamSegmentDelta) SetSpeakerId(v string)`

SetSpeakerId sets SpeakerId field to given value.

### HasSpeakerId

`func (o *TranscriptionStreamSegmentDelta) HasSpeakerId() bool`

HasSpeakerId returns a boolean if a field has been set.

### SetSpeakerIdNil

`func (o *TranscriptionStreamSegmentDelta) SetSpeakerIdNil(b bool)`

 SetSpeakerIdNil sets the value for SpeakerId to be an explicit nil

### UnsetSpeakerId
`func (o *TranscriptionStreamSegmentDelta) UnsetSpeakerId()`

UnsetSpeakerId ensures that no value is present for SpeakerId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


