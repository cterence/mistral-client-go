# OCRPageObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | **int32** | The page index in a pdf document starting from 0 | 
**Markdown** | **string** | The markdown string response of the page | 
**Images** | [**[]OCRImageObject**](OCRImageObject.md) | List of all extracted images in the page | 
**Tables** | Pointer to [**[]OCRTableObject**](OCRTableObject.md) | List of all extracted tables in the page | [optional] 
**Hyperlinks** | Pointer to **[]string** | List of all hyperlinks in the page | [optional] 
**Header** | Pointer to **NullableString** | Header of the page | [optional] 
**Footer** | Pointer to **NullableString** | Footer of the page | [optional] 
**Dimensions** | [**NullableOCRPageDimensions**](OCRPageDimensions.md) | The dimensions of the PDF Page&#39;s screenshot image | 
**ConfidenceScores** | Pointer to [**NullableOCRPageConfidenceScores**](OCRPageConfidenceScores.md) | Confidence scores for the OCR page (populated when confidence_scores_granularity is set) | [optional] 
**Blocks** | Pointer to [**[]OCRPageObjectBlocksInner**](OCRPageObjectBlocksInner.md) | Paragraph-level bounding boxes for all content blocks in reading order (populated when include_blocks is True) | [optional] 

## Methods

### NewOCRPageObject

`func NewOCRPageObject(index int32, markdown string, images []OCRImageObject, dimensions NullableOCRPageDimensions, ) *OCRPageObject`

NewOCRPageObject instantiates a new OCRPageObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOCRPageObjectWithDefaults

`func NewOCRPageObjectWithDefaults() *OCRPageObject`

NewOCRPageObjectWithDefaults instantiates a new OCRPageObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *OCRPageObject) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *OCRPageObject) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *OCRPageObject) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetMarkdown

`func (o *OCRPageObject) GetMarkdown() string`

GetMarkdown returns the Markdown field if non-nil, zero value otherwise.

### GetMarkdownOk

`func (o *OCRPageObject) GetMarkdownOk() (*string, bool)`

GetMarkdownOk returns a tuple with the Markdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkdown

`func (o *OCRPageObject) SetMarkdown(v string)`

SetMarkdown sets Markdown field to given value.


### GetImages

`func (o *OCRPageObject) GetImages() []OCRImageObject`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *OCRPageObject) GetImagesOk() (*[]OCRImageObject, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *OCRPageObject) SetImages(v []OCRImageObject)`

SetImages sets Images field to given value.


### GetTables

`func (o *OCRPageObject) GetTables() []OCRTableObject`

GetTables returns the Tables field if non-nil, zero value otherwise.

### GetTablesOk

`func (o *OCRPageObject) GetTablesOk() (*[]OCRTableObject, bool)`

GetTablesOk returns a tuple with the Tables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTables

`func (o *OCRPageObject) SetTables(v []OCRTableObject)`

SetTables sets Tables field to given value.

### HasTables

`func (o *OCRPageObject) HasTables() bool`

HasTables returns a boolean if a field has been set.

### GetHyperlinks

`func (o *OCRPageObject) GetHyperlinks() []string`

GetHyperlinks returns the Hyperlinks field if non-nil, zero value otherwise.

### GetHyperlinksOk

`func (o *OCRPageObject) GetHyperlinksOk() (*[]string, bool)`

GetHyperlinksOk returns a tuple with the Hyperlinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHyperlinks

`func (o *OCRPageObject) SetHyperlinks(v []string)`

SetHyperlinks sets Hyperlinks field to given value.

### HasHyperlinks

`func (o *OCRPageObject) HasHyperlinks() bool`

HasHyperlinks returns a boolean if a field has been set.

### GetHeader

`func (o *OCRPageObject) GetHeader() string`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *OCRPageObject) GetHeaderOk() (*string, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *OCRPageObject) SetHeader(v string)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *OCRPageObject) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### SetHeaderNil

`func (o *OCRPageObject) SetHeaderNil(b bool)`

 SetHeaderNil sets the value for Header to be an explicit nil

### UnsetHeader
`func (o *OCRPageObject) UnsetHeader()`

UnsetHeader ensures that no value is present for Header, not even an explicit nil
### GetFooter

`func (o *OCRPageObject) GetFooter() string`

GetFooter returns the Footer field if non-nil, zero value otherwise.

### GetFooterOk

`func (o *OCRPageObject) GetFooterOk() (*string, bool)`

GetFooterOk returns a tuple with the Footer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFooter

`func (o *OCRPageObject) SetFooter(v string)`

SetFooter sets Footer field to given value.

### HasFooter

`func (o *OCRPageObject) HasFooter() bool`

HasFooter returns a boolean if a field has been set.

### SetFooterNil

`func (o *OCRPageObject) SetFooterNil(b bool)`

 SetFooterNil sets the value for Footer to be an explicit nil

### UnsetFooter
`func (o *OCRPageObject) UnsetFooter()`

UnsetFooter ensures that no value is present for Footer, not even an explicit nil
### GetDimensions

`func (o *OCRPageObject) GetDimensions() OCRPageDimensions`

GetDimensions returns the Dimensions field if non-nil, zero value otherwise.

### GetDimensionsOk

`func (o *OCRPageObject) GetDimensionsOk() (*OCRPageDimensions, bool)`

GetDimensionsOk returns a tuple with the Dimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensions

`func (o *OCRPageObject) SetDimensions(v OCRPageDimensions)`

SetDimensions sets Dimensions field to given value.


### SetDimensionsNil

`func (o *OCRPageObject) SetDimensionsNil(b bool)`

 SetDimensionsNil sets the value for Dimensions to be an explicit nil

### UnsetDimensions
`func (o *OCRPageObject) UnsetDimensions()`

UnsetDimensions ensures that no value is present for Dimensions, not even an explicit nil
### GetConfidenceScores

`func (o *OCRPageObject) GetConfidenceScores() OCRPageConfidenceScores`

GetConfidenceScores returns the ConfidenceScores field if non-nil, zero value otherwise.

### GetConfidenceScoresOk

`func (o *OCRPageObject) GetConfidenceScoresOk() (*OCRPageConfidenceScores, bool)`

GetConfidenceScoresOk returns a tuple with the ConfidenceScores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidenceScores

`func (o *OCRPageObject) SetConfidenceScores(v OCRPageConfidenceScores)`

SetConfidenceScores sets ConfidenceScores field to given value.

### HasConfidenceScores

`func (o *OCRPageObject) HasConfidenceScores() bool`

HasConfidenceScores returns a boolean if a field has been set.

### SetConfidenceScoresNil

`func (o *OCRPageObject) SetConfidenceScoresNil(b bool)`

 SetConfidenceScoresNil sets the value for ConfidenceScores to be an explicit nil

### UnsetConfidenceScores
`func (o *OCRPageObject) UnsetConfidenceScores()`

UnsetConfidenceScores ensures that no value is present for ConfidenceScores, not even an explicit nil
### GetBlocks

`func (o *OCRPageObject) GetBlocks() []OCRPageObjectBlocksInner`

GetBlocks returns the Blocks field if non-nil, zero value otherwise.

### GetBlocksOk

`func (o *OCRPageObject) GetBlocksOk() (*[]OCRPageObjectBlocksInner, bool)`

GetBlocksOk returns a tuple with the Blocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocks

`func (o *OCRPageObject) SetBlocks(v []OCRPageObjectBlocksInner)`

SetBlocks sets Blocks field to given value.

### HasBlocks

`func (o *OCRPageObject) HasBlocks() bool`

HasBlocks returns a boolean if a field has been set.

### SetBlocksNil

`func (o *OCRPageObject) SetBlocksNil(b bool)`

 SetBlocksNil sets the value for Blocks to be an explicit nil

### UnsetBlocks
`func (o *OCRPageObject) UnsetBlocks()`

UnsetBlocks ensures that no value is present for Blocks, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


