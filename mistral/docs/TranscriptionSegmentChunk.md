# TranscriptionSegmentChunk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "transcription_segment"]
**Text** | **string** |  | 
**Start** | **float32** |  | 
**End** | **float32** |  | 
**Score** | Pointer to **NullableFloat32** |  | [optional] 
**SpeakerId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTranscriptionSegmentChunk

`func NewTranscriptionSegmentChunk(text string, start float32, end float32, ) *TranscriptionSegmentChunk`

NewTranscriptionSegmentChunk instantiates a new TranscriptionSegmentChunk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTranscriptionSegmentChunkWithDefaults

`func NewTranscriptionSegmentChunkWithDefaults() *TranscriptionSegmentChunk`

NewTranscriptionSegmentChunkWithDefaults instantiates a new TranscriptionSegmentChunk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TranscriptionSegmentChunk) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TranscriptionSegmentChunk) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TranscriptionSegmentChunk) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TranscriptionSegmentChunk) HasType() bool`

HasType returns a boolean if a field has been set.

### GetText

`func (o *TranscriptionSegmentChunk) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *TranscriptionSegmentChunk) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *TranscriptionSegmentChunk) SetText(v string)`

SetText sets Text field to given value.


### GetStart

`func (o *TranscriptionSegmentChunk) GetStart() float32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *TranscriptionSegmentChunk) GetStartOk() (*float32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *TranscriptionSegmentChunk) SetStart(v float32)`

SetStart sets Start field to given value.


### GetEnd

`func (o *TranscriptionSegmentChunk) GetEnd() float32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *TranscriptionSegmentChunk) GetEndOk() (*float32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *TranscriptionSegmentChunk) SetEnd(v float32)`

SetEnd sets End field to given value.


### GetScore

`func (o *TranscriptionSegmentChunk) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *TranscriptionSegmentChunk) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *TranscriptionSegmentChunk) SetScore(v float32)`

SetScore sets Score field to given value.

### HasScore

`func (o *TranscriptionSegmentChunk) HasScore() bool`

HasScore returns a boolean if a field has been set.

### SetScoreNil

`func (o *TranscriptionSegmentChunk) SetScoreNil(b bool)`

 SetScoreNil sets the value for Score to be an explicit nil

### UnsetScore
`func (o *TranscriptionSegmentChunk) UnsetScore()`

UnsetScore ensures that no value is present for Score, not even an explicit nil
### GetSpeakerId

`func (o *TranscriptionSegmentChunk) GetSpeakerId() string`

GetSpeakerId returns the SpeakerId field if non-nil, zero value otherwise.

### GetSpeakerIdOk

`func (o *TranscriptionSegmentChunk) GetSpeakerIdOk() (*string, bool)`

GetSpeakerIdOk returns a tuple with the SpeakerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeakerId

`func (o *TranscriptionSegmentChunk) SetSpeakerId(v string)`

SetSpeakerId sets SpeakerId field to given value.

### HasSpeakerId

`func (o *TranscriptionSegmentChunk) HasSpeakerId() bool`

HasSpeakerId returns a boolean if a field has been set.

### SetSpeakerIdNil

`func (o *TranscriptionSegmentChunk) SetSpeakerIdNil(b bool)`

 SetSpeakerIdNil sets the value for SpeakerId to be an explicit nil

### UnsetSpeakerId
`func (o *TranscriptionSegmentChunk) UnsetSpeakerId()`

UnsetSpeakerId ensures that no value is present for SpeakerId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


