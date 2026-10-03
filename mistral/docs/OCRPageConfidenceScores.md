# OCRPageConfidenceScores

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WordConfidenceScores** | Pointer to [**[]OCRConfidenceScore**](OCRConfidenceScore.md) | Word-level confidence scores (populated only for &#39;word&#39; granularity) | [optional] 
**AveragePageConfidenceScore** | **float32** | Average confidence score for the page | 
**MinimumPageConfidenceScore** | **float32** | Minimum confidence score for the page | 

## Methods

### NewOCRPageConfidenceScores

`func NewOCRPageConfidenceScores(averagePageConfidenceScore float32, minimumPageConfidenceScore float32, ) *OCRPageConfidenceScores`

NewOCRPageConfidenceScores instantiates a new OCRPageConfidenceScores object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOCRPageConfidenceScoresWithDefaults

`func NewOCRPageConfidenceScoresWithDefaults() *OCRPageConfidenceScores`

NewOCRPageConfidenceScoresWithDefaults instantiates a new OCRPageConfidenceScores object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWordConfidenceScores

`func (o *OCRPageConfidenceScores) GetWordConfidenceScores() []OCRConfidenceScore`

GetWordConfidenceScores returns the WordConfidenceScores field if non-nil, zero value otherwise.

### GetWordConfidenceScoresOk

`func (o *OCRPageConfidenceScores) GetWordConfidenceScoresOk() (*[]OCRConfidenceScore, bool)`

GetWordConfidenceScoresOk returns a tuple with the WordConfidenceScores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWordConfidenceScores

`func (o *OCRPageConfidenceScores) SetWordConfidenceScores(v []OCRConfidenceScore)`

SetWordConfidenceScores sets WordConfidenceScores field to given value.

### HasWordConfidenceScores

`func (o *OCRPageConfidenceScores) HasWordConfidenceScores() bool`

HasWordConfidenceScores returns a boolean if a field has been set.

### GetAveragePageConfidenceScore

`func (o *OCRPageConfidenceScores) GetAveragePageConfidenceScore() float32`

GetAveragePageConfidenceScore returns the AveragePageConfidenceScore field if non-nil, zero value otherwise.

### GetAveragePageConfidenceScoreOk

`func (o *OCRPageConfidenceScores) GetAveragePageConfidenceScoreOk() (*float32, bool)`

GetAveragePageConfidenceScoreOk returns a tuple with the AveragePageConfidenceScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAveragePageConfidenceScore

`func (o *OCRPageConfidenceScores) SetAveragePageConfidenceScore(v float32)`

SetAveragePageConfidenceScore sets AveragePageConfidenceScore field to given value.


### GetMinimumPageConfidenceScore

`func (o *OCRPageConfidenceScores) GetMinimumPageConfidenceScore() float32`

GetMinimumPageConfidenceScore returns the MinimumPageConfidenceScore field if non-nil, zero value otherwise.

### GetMinimumPageConfidenceScoreOk

`func (o *OCRPageConfidenceScores) GetMinimumPageConfidenceScoreOk() (*float32, bool)`

GetMinimumPageConfidenceScoreOk returns a tuple with the MinimumPageConfidenceScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumPageConfidenceScore

`func (o *OCRPageConfidenceScores) SetMinimumPageConfidenceScore(v float32)`

SetMinimumPageConfidenceScore sets MinimumPageConfidenceScore field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


