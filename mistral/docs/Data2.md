# Data2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "transcription.text.delta"]
**Text** | **string** |  | 
**AudioLanguage** | **string** |  | 
**Start** | **float32** |  | 
**End** | **float32** |  | 
**SpeakerId** | Pointer to **string** |  | [optional] 
**Model** | **string** |  | 
**Language** | **string** |  | 
**Segments** | Pointer to [**[]TranscriptionSegmentChunk**](TranscriptionSegmentChunk.md) |  | [optional] 
**Usage** | [**UsageInfo**](UsageInfo.md) |  | 

## Methods

### NewData2

`func NewData2(text string, audioLanguage string, start float32, end float32, model string, language string, usage UsageInfo, ) *Data2`

NewData2 instantiates a new Data2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewData2WithDefaults

`func NewData2WithDefaults() *Data2`

NewData2WithDefaults instantiates a new Data2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Data2) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Data2) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Data2) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Data2) HasType() bool`

HasType returns a boolean if a field has been set.

### GetText

`func (o *Data2) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *Data2) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *Data2) SetText(v string)`

SetText sets Text field to given value.


### GetAudioLanguage

`func (o *Data2) GetAudioLanguage() string`

GetAudioLanguage returns the AudioLanguage field if non-nil, zero value otherwise.

### GetAudioLanguageOk

`func (o *Data2) GetAudioLanguageOk() (*string, bool)`

GetAudioLanguageOk returns a tuple with the AudioLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioLanguage

`func (o *Data2) SetAudioLanguage(v string)`

SetAudioLanguage sets AudioLanguage field to given value.


### GetStart

`func (o *Data2) GetStart() float32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *Data2) GetStartOk() (*float32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *Data2) SetStart(v float32)`

SetStart sets Start field to given value.


### GetEnd

`func (o *Data2) GetEnd() float32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *Data2) GetEndOk() (*float32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *Data2) SetEnd(v float32)`

SetEnd sets End field to given value.


### GetSpeakerId

`func (o *Data2) GetSpeakerId() string`

GetSpeakerId returns the SpeakerId field if non-nil, zero value otherwise.

### GetSpeakerIdOk

`func (o *Data2) GetSpeakerIdOk() (*string, bool)`

GetSpeakerIdOk returns a tuple with the SpeakerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeakerId

`func (o *Data2) SetSpeakerId(v string)`

SetSpeakerId sets SpeakerId field to given value.

### HasSpeakerId

`func (o *Data2) HasSpeakerId() bool`

HasSpeakerId returns a boolean if a field has been set.

### GetModel

`func (o *Data2) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *Data2) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *Data2) SetModel(v string)`

SetModel sets Model field to given value.


### GetLanguage

`func (o *Data2) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *Data2) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *Data2) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### GetSegments

`func (o *Data2) GetSegments() []TranscriptionSegmentChunk`

GetSegments returns the Segments field if non-nil, zero value otherwise.

### GetSegmentsOk

`func (o *Data2) GetSegmentsOk() (*[]TranscriptionSegmentChunk, bool)`

GetSegmentsOk returns a tuple with the Segments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegments

`func (o *Data2) SetSegments(v []TranscriptionSegmentChunk)`

SetSegments sets Segments field to given value.

### HasSegments

`func (o *Data2) HasSegments() bool`

HasSegments returns a boolean if a field has been set.

### GetUsage

`func (o *Data2) GetUsage() UsageInfo`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *Data2) GetUsageOk() (*UsageInfo, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *Data2) SetUsage(v UsageInfo)`

SetUsage sets Usage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


