# OCRBlockConfidenceScores

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AverageContentConfidenceScore** | Pointer to **NullableFloat32** | Average confidence over the block&#39;s content (caption) tokens. None when the block has no textual content (e.g. image-only entry). | [optional] 
**MinimumContentConfidenceScore** | Pointer to **NullableFloat32** | Minimum per-word content confidence in the block. None when the block has no textual content. | [optional] 
**BlockTypeConfidenceScore** | Pointer to **NullableFloat32** | Confidence in the block type (e.g. &#39;text&#39;, &#39;title&#39;, &#39;table&#39;). None when the entry had no block type or the block type span could not be located. | [optional] 

## Methods

### NewOCRBlockConfidenceScores

`func NewOCRBlockConfidenceScores() *OCRBlockConfidenceScores`

NewOCRBlockConfidenceScores instantiates a new OCRBlockConfidenceScores object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOCRBlockConfidenceScoresWithDefaults

`func NewOCRBlockConfidenceScoresWithDefaults() *OCRBlockConfidenceScores`

NewOCRBlockConfidenceScoresWithDefaults instantiates a new OCRBlockConfidenceScores object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAverageContentConfidenceScore

`func (o *OCRBlockConfidenceScores) GetAverageContentConfidenceScore() float32`

GetAverageContentConfidenceScore returns the AverageContentConfidenceScore field if non-nil, zero value otherwise.

### GetAverageContentConfidenceScoreOk

`func (o *OCRBlockConfidenceScores) GetAverageContentConfidenceScoreOk() (*float32, bool)`

GetAverageContentConfidenceScoreOk returns a tuple with the AverageContentConfidenceScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageContentConfidenceScore

`func (o *OCRBlockConfidenceScores) SetAverageContentConfidenceScore(v float32)`

SetAverageContentConfidenceScore sets AverageContentConfidenceScore field to given value.

### HasAverageContentConfidenceScore

`func (o *OCRBlockConfidenceScores) HasAverageContentConfidenceScore() bool`

HasAverageContentConfidenceScore returns a boolean if a field has been set.

### SetAverageContentConfidenceScoreNil

`func (o *OCRBlockConfidenceScores) SetAverageContentConfidenceScoreNil(b bool)`

 SetAverageContentConfidenceScoreNil sets the value for AverageContentConfidenceScore to be an explicit nil

### UnsetAverageContentConfidenceScore
`func (o *OCRBlockConfidenceScores) UnsetAverageContentConfidenceScore()`

UnsetAverageContentConfidenceScore ensures that no value is present for AverageContentConfidenceScore, not even an explicit nil
### GetMinimumContentConfidenceScore

`func (o *OCRBlockConfidenceScores) GetMinimumContentConfidenceScore() float32`

GetMinimumContentConfidenceScore returns the MinimumContentConfidenceScore field if non-nil, zero value otherwise.

### GetMinimumContentConfidenceScoreOk

`func (o *OCRBlockConfidenceScores) GetMinimumContentConfidenceScoreOk() (*float32, bool)`

GetMinimumContentConfidenceScoreOk returns a tuple with the MinimumContentConfidenceScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumContentConfidenceScore

`func (o *OCRBlockConfidenceScores) SetMinimumContentConfidenceScore(v float32)`

SetMinimumContentConfidenceScore sets MinimumContentConfidenceScore field to given value.

### HasMinimumContentConfidenceScore

`func (o *OCRBlockConfidenceScores) HasMinimumContentConfidenceScore() bool`

HasMinimumContentConfidenceScore returns a boolean if a field has been set.

### SetMinimumContentConfidenceScoreNil

`func (o *OCRBlockConfidenceScores) SetMinimumContentConfidenceScoreNil(b bool)`

 SetMinimumContentConfidenceScoreNil sets the value for MinimumContentConfidenceScore to be an explicit nil

### UnsetMinimumContentConfidenceScore
`func (o *OCRBlockConfidenceScores) UnsetMinimumContentConfidenceScore()`

UnsetMinimumContentConfidenceScore ensures that no value is present for MinimumContentConfidenceScore, not even an explicit nil
### GetBlockTypeConfidenceScore

`func (o *OCRBlockConfidenceScores) GetBlockTypeConfidenceScore() float32`

GetBlockTypeConfidenceScore returns the BlockTypeConfidenceScore field if non-nil, zero value otherwise.

### GetBlockTypeConfidenceScoreOk

`func (o *OCRBlockConfidenceScores) GetBlockTypeConfidenceScoreOk() (*float32, bool)`

GetBlockTypeConfidenceScoreOk returns a tuple with the BlockTypeConfidenceScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockTypeConfidenceScore

`func (o *OCRBlockConfidenceScores) SetBlockTypeConfidenceScore(v float32)`

SetBlockTypeConfidenceScore sets BlockTypeConfidenceScore field to given value.

### HasBlockTypeConfidenceScore

`func (o *OCRBlockConfidenceScores) HasBlockTypeConfidenceScore() bool`

HasBlockTypeConfidenceScore returns a boolean if a field has been set.

### SetBlockTypeConfidenceScoreNil

`func (o *OCRBlockConfidenceScores) SetBlockTypeConfidenceScoreNil(b bool)`

 SetBlockTypeConfidenceScoreNil sets the value for BlockTypeConfidenceScore to be an explicit nil

### UnsetBlockTypeConfidenceScore
`func (o *OCRBlockConfidenceScores) UnsetBlockTypeConfidenceScore()`

UnsetBlockTypeConfidenceScore ensures that no value is present for BlockTypeConfidenceScore, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


