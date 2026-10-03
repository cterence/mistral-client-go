# TranscriptionStreamDone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | **string** |  | 
**Text** | **string** |  | 
**Language** | **NullableString** |  | 
**Segments** | Pointer to [**[]TranscriptionSegmentChunk**](TranscriptionSegmentChunk.md) |  | [optional] 
**Usage** | [**UsageInfo**](UsageInfo.md) |  | 
**Type** | Pointer to **string** |  | [optional] [default to "transcription.done"]

## Methods

### NewTranscriptionStreamDone

`func NewTranscriptionStreamDone(model string, text string, language NullableString, usage UsageInfo, ) *TranscriptionStreamDone`

NewTranscriptionStreamDone instantiates a new TranscriptionStreamDone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTranscriptionStreamDoneWithDefaults

`func NewTranscriptionStreamDoneWithDefaults() *TranscriptionStreamDone`

NewTranscriptionStreamDoneWithDefaults instantiates a new TranscriptionStreamDone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *TranscriptionStreamDone) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *TranscriptionStreamDone) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *TranscriptionStreamDone) SetModel(v string)`

SetModel sets Model field to given value.


### GetText

`func (o *TranscriptionStreamDone) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *TranscriptionStreamDone) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *TranscriptionStreamDone) SetText(v string)`

SetText sets Text field to given value.


### GetLanguage

`func (o *TranscriptionStreamDone) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *TranscriptionStreamDone) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *TranscriptionStreamDone) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### SetLanguageNil

`func (o *TranscriptionStreamDone) SetLanguageNil(b bool)`

 SetLanguageNil sets the value for Language to be an explicit nil

### UnsetLanguage
`func (o *TranscriptionStreamDone) UnsetLanguage()`

UnsetLanguage ensures that no value is present for Language, not even an explicit nil
### GetSegments

`func (o *TranscriptionStreamDone) GetSegments() []TranscriptionSegmentChunk`

GetSegments returns the Segments field if non-nil, zero value otherwise.

### GetSegmentsOk

`func (o *TranscriptionStreamDone) GetSegmentsOk() (*[]TranscriptionSegmentChunk, bool)`

GetSegmentsOk returns a tuple with the Segments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegments

`func (o *TranscriptionStreamDone) SetSegments(v []TranscriptionSegmentChunk)`

SetSegments sets Segments field to given value.

### HasSegments

`func (o *TranscriptionStreamDone) HasSegments() bool`

HasSegments returns a boolean if a field has been set.

### GetUsage

`func (o *TranscriptionStreamDone) GetUsage() UsageInfo`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *TranscriptionStreamDone) GetUsageOk() (*UsageInfo, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *TranscriptionStreamDone) SetUsage(v UsageInfo)`

SetUsage sets Usage field to given value.


### GetType

`func (o *TranscriptionStreamDone) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TranscriptionStreamDone) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TranscriptionStreamDone) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TranscriptionStreamDone) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


