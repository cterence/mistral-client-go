# OCRTableObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Table ID for extracted table in a page | 
**Content** | **string** | Content of the table in the given format | 
**Format** | **string** | Format of the table | 
**WordConfidenceScores** | Pointer to [**[]OCRConfidenceScore**](OCRConfidenceScore.md) | Per-word confidence scores for the table content. Returned when confidence_scores_granularity is set to &#39;word&#39;. | [optional] 

## Methods

### NewOCRTableObject

`func NewOCRTableObject(id string, content string, format string, ) *OCRTableObject`

NewOCRTableObject instantiates a new OCRTableObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOCRTableObjectWithDefaults

`func NewOCRTableObjectWithDefaults() *OCRTableObject`

NewOCRTableObjectWithDefaults instantiates a new OCRTableObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OCRTableObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OCRTableObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OCRTableObject) SetId(v string)`

SetId sets Id field to given value.


### GetContent

`func (o *OCRTableObject) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *OCRTableObject) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *OCRTableObject) SetContent(v string)`

SetContent sets Content field to given value.


### GetFormat

`func (o *OCRTableObject) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *OCRTableObject) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *OCRTableObject) SetFormat(v string)`

SetFormat sets Format field to given value.


### GetWordConfidenceScores

`func (o *OCRTableObject) GetWordConfidenceScores() []OCRConfidenceScore`

GetWordConfidenceScores returns the WordConfidenceScores field if non-nil, zero value otherwise.

### GetWordConfidenceScoresOk

`func (o *OCRTableObject) GetWordConfidenceScoresOk() (*[]OCRConfidenceScore, bool)`

GetWordConfidenceScoresOk returns a tuple with the WordConfidenceScores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWordConfidenceScores

`func (o *OCRTableObject) SetWordConfidenceScores(v []OCRConfidenceScore)`

SetWordConfidenceScores sets WordConfidenceScores field to given value.

### HasWordConfidenceScores

`func (o *OCRTableObject) HasWordConfidenceScores() bool`

HasWordConfidenceScores returns a boolean if a field has been set.

### SetWordConfidenceScoresNil

`func (o *OCRTableObject) SetWordConfidenceScoresNil(b bool)`

 SetWordConfidenceScoresNil sets the value for WordConfidenceScores to be an explicit nil

### UnsetWordConfidenceScores
`func (o *OCRTableObject) UnsetWordConfidenceScores()`

UnsetWordConfidenceScores ensures that no value is present for WordConfidenceScores, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


