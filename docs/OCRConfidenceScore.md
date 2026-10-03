# OCRConfidenceScore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | **string** | The word or text segment | 
**Confidence** | **float32** | Confidence score (0-1) | 
**StartIndex** | **int32** | Start index of the text in the page markdown string | 

## Methods

### NewOCRConfidenceScore

`func NewOCRConfidenceScore(text string, confidence float32, startIndex int32, ) *OCRConfidenceScore`

NewOCRConfidenceScore instantiates a new OCRConfidenceScore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOCRConfidenceScoreWithDefaults

`func NewOCRConfidenceScoreWithDefaults() *OCRConfidenceScore`

NewOCRConfidenceScoreWithDefaults instantiates a new OCRConfidenceScore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *OCRConfidenceScore) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *OCRConfidenceScore) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *OCRConfidenceScore) SetText(v string)`

SetText sets Text field to given value.


### GetConfidence

`func (o *OCRConfidenceScore) GetConfidence() float32`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *OCRConfidenceScore) GetConfidenceOk() (*float32, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *OCRConfidenceScore) SetConfidence(v float32)`

SetConfidence sets Confidence field to given value.


### GetStartIndex

`func (o *OCRConfidenceScore) GetStartIndex() int32`

GetStartIndex returns the StartIndex field if non-nil, zero value otherwise.

### GetStartIndexOk

`func (o *OCRConfidenceScore) GetStartIndexOk() (*int32, bool)`

GetStartIndexOk returns a tuple with the StartIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartIndex

`func (o *OCRConfidenceScore) SetStartIndex(v int32)`

SetStartIndex sets StartIndex field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


