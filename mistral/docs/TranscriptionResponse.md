# TranscriptionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | **string** |  | 
**Text** | **string** |  | 
**Language** | **NullableString** |  | 
**Segments** | Pointer to [**[]TranscriptionSegmentChunk**](TranscriptionSegmentChunk.md) |  | [optional] 
**Usage** | [**UsageInfo**](UsageInfo.md) |  | 

## Methods

### NewTranscriptionResponse

`func NewTranscriptionResponse(model string, text string, language NullableString, usage UsageInfo, ) *TranscriptionResponse`

NewTranscriptionResponse instantiates a new TranscriptionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTranscriptionResponseWithDefaults

`func NewTranscriptionResponseWithDefaults() *TranscriptionResponse`

NewTranscriptionResponseWithDefaults instantiates a new TranscriptionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *TranscriptionResponse) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *TranscriptionResponse) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *TranscriptionResponse) SetModel(v string)`

SetModel sets Model field to given value.


### GetText

`func (o *TranscriptionResponse) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *TranscriptionResponse) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *TranscriptionResponse) SetText(v string)`

SetText sets Text field to given value.


### GetLanguage

`func (o *TranscriptionResponse) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *TranscriptionResponse) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *TranscriptionResponse) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### SetLanguageNil

`func (o *TranscriptionResponse) SetLanguageNil(b bool)`

 SetLanguageNil sets the value for Language to be an explicit nil

### UnsetLanguage
`func (o *TranscriptionResponse) UnsetLanguage()`

UnsetLanguage ensures that no value is present for Language, not even an explicit nil
### GetSegments

`func (o *TranscriptionResponse) GetSegments() []TranscriptionSegmentChunk`

GetSegments returns the Segments field if non-nil, zero value otherwise.

### GetSegmentsOk

`func (o *TranscriptionResponse) GetSegmentsOk() (*[]TranscriptionSegmentChunk, bool)`

GetSegmentsOk returns a tuple with the Segments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegments

`func (o *TranscriptionResponse) SetSegments(v []TranscriptionSegmentChunk)`

SetSegments sets Segments field to given value.

### HasSegments

`func (o *TranscriptionResponse) HasSegments() bool`

HasSegments returns a boolean if a field has been set.

### GetUsage

`func (o *TranscriptionResponse) GetUsage() UsageInfo`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *TranscriptionResponse) GetUsageOk() (*UsageInfo, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *TranscriptionResponse) SetUsage(v UsageInfo)`

SetUsage sets Usage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


