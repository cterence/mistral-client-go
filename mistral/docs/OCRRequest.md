# OCRRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | **NullableString** |  | 
**Id** | Pointer to **string** |  | [optional] 
**Document** | [**Document**](Document.md) |  | 
**Pages** | Pointer to [**NullablePages**](Pages.md) |  | [optional] 
**IncludeImageBase64** | Pointer to **NullableBool** | Include image URLs in response | [optional] 
**ImageLimit** | Pointer to **NullableInt32** | Max images to extract | [optional] 
**ImageMinSize** | Pointer to **NullableInt32** | Minimum height and width of image to extract | [optional] 
**BboxAnnotationFormat** | Pointer to [**NullableResponseFormat**](ResponseFormat.md) | Structured output class for extracting useful information from each extracted bounding box / image from document. Only json_schema is valid for this field | [optional] 
**DocumentAnnotationFormat** | Pointer to [**NullableResponseFormat**](ResponseFormat.md) | Structured output class for extracting useful information from the entire document. Only json_schema is valid for this field | [optional] 
**DocumentAnnotationPrompt** | Pointer to **NullableString** | Optional prompt to guide the model in extracting structured output from the entire document. A document_annotation_format must be provided. | [optional] 
**TableFormat** | Pointer to **NullableString** |  | [optional] 
**ExtractHeader** | Pointer to **bool** |  | [optional] [default to false]
**ExtractFooter** | Pointer to **bool** |  | [optional] [default to false]
**IncludeBlocks** | Pointer to **bool** | Return paragraph-level bounding boxes for all content blocks in the response | [optional] [default to false]
**ConfidenceScoresGranularity** | Pointer to **NullableString** | Granularity for confidence scores: &#39;word&#39; (per-word scores) or &#39;page&#39; (aggregate only). Defaults to None (no confidence scores) to keep response payload small. | [optional] 

## Methods

### NewOCRRequest

`func NewOCRRequest(model NullableString, document Document, ) *OCRRequest`

NewOCRRequest instantiates a new OCRRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOCRRequestWithDefaults

`func NewOCRRequestWithDefaults() *OCRRequest`

NewOCRRequestWithDefaults instantiates a new OCRRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *OCRRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *OCRRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *OCRRequest) SetModel(v string)`

SetModel sets Model field to given value.


### SetModelNil

`func (o *OCRRequest) SetModelNil(b bool)`

 SetModelNil sets the value for Model to be an explicit nil

### UnsetModel
`func (o *OCRRequest) UnsetModel()`

UnsetModel ensures that no value is present for Model, not even an explicit nil
### GetId

`func (o *OCRRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OCRRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OCRRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OCRRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDocument

`func (o *OCRRequest) GetDocument() Document`

GetDocument returns the Document field if non-nil, zero value otherwise.

### GetDocumentOk

`func (o *OCRRequest) GetDocumentOk() (*Document, bool)`

GetDocumentOk returns a tuple with the Document field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocument

`func (o *OCRRequest) SetDocument(v Document)`

SetDocument sets Document field to given value.


### GetPages

`func (o *OCRRequest) GetPages() Pages`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *OCRRequest) GetPagesOk() (*Pages, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *OCRRequest) SetPages(v Pages)`

SetPages sets Pages field to given value.

### HasPages

`func (o *OCRRequest) HasPages() bool`

HasPages returns a boolean if a field has been set.

### SetPagesNil

`func (o *OCRRequest) SetPagesNil(b bool)`

 SetPagesNil sets the value for Pages to be an explicit nil

### UnsetPages
`func (o *OCRRequest) UnsetPages()`

UnsetPages ensures that no value is present for Pages, not even an explicit nil
### GetIncludeImageBase64

`func (o *OCRRequest) GetIncludeImageBase64() bool`

GetIncludeImageBase64 returns the IncludeImageBase64 field if non-nil, zero value otherwise.

### GetIncludeImageBase64Ok

`func (o *OCRRequest) GetIncludeImageBase64Ok() (*bool, bool)`

GetIncludeImageBase64Ok returns a tuple with the IncludeImageBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeImageBase64

`func (o *OCRRequest) SetIncludeImageBase64(v bool)`

SetIncludeImageBase64 sets IncludeImageBase64 field to given value.

### HasIncludeImageBase64

`func (o *OCRRequest) HasIncludeImageBase64() bool`

HasIncludeImageBase64 returns a boolean if a field has been set.

### SetIncludeImageBase64Nil

`func (o *OCRRequest) SetIncludeImageBase64Nil(b bool)`

 SetIncludeImageBase64Nil sets the value for IncludeImageBase64 to be an explicit nil

### UnsetIncludeImageBase64
`func (o *OCRRequest) UnsetIncludeImageBase64()`

UnsetIncludeImageBase64 ensures that no value is present for IncludeImageBase64, not even an explicit nil
### GetImageLimit

`func (o *OCRRequest) GetImageLimit() int32`

GetImageLimit returns the ImageLimit field if non-nil, zero value otherwise.

### GetImageLimitOk

`func (o *OCRRequest) GetImageLimitOk() (*int32, bool)`

GetImageLimitOk returns a tuple with the ImageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageLimit

`func (o *OCRRequest) SetImageLimit(v int32)`

SetImageLimit sets ImageLimit field to given value.

### HasImageLimit

`func (o *OCRRequest) HasImageLimit() bool`

HasImageLimit returns a boolean if a field has been set.

### SetImageLimitNil

`func (o *OCRRequest) SetImageLimitNil(b bool)`

 SetImageLimitNil sets the value for ImageLimit to be an explicit nil

### UnsetImageLimit
`func (o *OCRRequest) UnsetImageLimit()`

UnsetImageLimit ensures that no value is present for ImageLimit, not even an explicit nil
### GetImageMinSize

`func (o *OCRRequest) GetImageMinSize() int32`

GetImageMinSize returns the ImageMinSize field if non-nil, zero value otherwise.

### GetImageMinSizeOk

`func (o *OCRRequest) GetImageMinSizeOk() (*int32, bool)`

GetImageMinSizeOk returns a tuple with the ImageMinSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageMinSize

`func (o *OCRRequest) SetImageMinSize(v int32)`

SetImageMinSize sets ImageMinSize field to given value.

### HasImageMinSize

`func (o *OCRRequest) HasImageMinSize() bool`

HasImageMinSize returns a boolean if a field has been set.

### SetImageMinSizeNil

`func (o *OCRRequest) SetImageMinSizeNil(b bool)`

 SetImageMinSizeNil sets the value for ImageMinSize to be an explicit nil

### UnsetImageMinSize
`func (o *OCRRequest) UnsetImageMinSize()`

UnsetImageMinSize ensures that no value is present for ImageMinSize, not even an explicit nil
### GetBboxAnnotationFormat

`func (o *OCRRequest) GetBboxAnnotationFormat() ResponseFormat`

GetBboxAnnotationFormat returns the BboxAnnotationFormat field if non-nil, zero value otherwise.

### GetBboxAnnotationFormatOk

`func (o *OCRRequest) GetBboxAnnotationFormatOk() (*ResponseFormat, bool)`

GetBboxAnnotationFormatOk returns a tuple with the BboxAnnotationFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBboxAnnotationFormat

`func (o *OCRRequest) SetBboxAnnotationFormat(v ResponseFormat)`

SetBboxAnnotationFormat sets BboxAnnotationFormat field to given value.

### HasBboxAnnotationFormat

`func (o *OCRRequest) HasBboxAnnotationFormat() bool`

HasBboxAnnotationFormat returns a boolean if a field has been set.

### SetBboxAnnotationFormatNil

`func (o *OCRRequest) SetBboxAnnotationFormatNil(b bool)`

 SetBboxAnnotationFormatNil sets the value for BboxAnnotationFormat to be an explicit nil

### UnsetBboxAnnotationFormat
`func (o *OCRRequest) UnsetBboxAnnotationFormat()`

UnsetBboxAnnotationFormat ensures that no value is present for BboxAnnotationFormat, not even an explicit nil
### GetDocumentAnnotationFormat

`func (o *OCRRequest) GetDocumentAnnotationFormat() ResponseFormat`

GetDocumentAnnotationFormat returns the DocumentAnnotationFormat field if non-nil, zero value otherwise.

### GetDocumentAnnotationFormatOk

`func (o *OCRRequest) GetDocumentAnnotationFormatOk() (*ResponseFormat, bool)`

GetDocumentAnnotationFormatOk returns a tuple with the DocumentAnnotationFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentAnnotationFormat

`func (o *OCRRequest) SetDocumentAnnotationFormat(v ResponseFormat)`

SetDocumentAnnotationFormat sets DocumentAnnotationFormat field to given value.

### HasDocumentAnnotationFormat

`func (o *OCRRequest) HasDocumentAnnotationFormat() bool`

HasDocumentAnnotationFormat returns a boolean if a field has been set.

### SetDocumentAnnotationFormatNil

`func (o *OCRRequest) SetDocumentAnnotationFormatNil(b bool)`

 SetDocumentAnnotationFormatNil sets the value for DocumentAnnotationFormat to be an explicit nil

### UnsetDocumentAnnotationFormat
`func (o *OCRRequest) UnsetDocumentAnnotationFormat()`

UnsetDocumentAnnotationFormat ensures that no value is present for DocumentAnnotationFormat, not even an explicit nil
### GetDocumentAnnotationPrompt

`func (o *OCRRequest) GetDocumentAnnotationPrompt() string`

GetDocumentAnnotationPrompt returns the DocumentAnnotationPrompt field if non-nil, zero value otherwise.

### GetDocumentAnnotationPromptOk

`func (o *OCRRequest) GetDocumentAnnotationPromptOk() (*string, bool)`

GetDocumentAnnotationPromptOk returns a tuple with the DocumentAnnotationPrompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentAnnotationPrompt

`func (o *OCRRequest) SetDocumentAnnotationPrompt(v string)`

SetDocumentAnnotationPrompt sets DocumentAnnotationPrompt field to given value.

### HasDocumentAnnotationPrompt

`func (o *OCRRequest) HasDocumentAnnotationPrompt() bool`

HasDocumentAnnotationPrompt returns a boolean if a field has been set.

### SetDocumentAnnotationPromptNil

`func (o *OCRRequest) SetDocumentAnnotationPromptNil(b bool)`

 SetDocumentAnnotationPromptNil sets the value for DocumentAnnotationPrompt to be an explicit nil

### UnsetDocumentAnnotationPrompt
`func (o *OCRRequest) UnsetDocumentAnnotationPrompt()`

UnsetDocumentAnnotationPrompt ensures that no value is present for DocumentAnnotationPrompt, not even an explicit nil
### GetTableFormat

`func (o *OCRRequest) GetTableFormat() string`

GetTableFormat returns the TableFormat field if non-nil, zero value otherwise.

### GetTableFormatOk

`func (o *OCRRequest) GetTableFormatOk() (*string, bool)`

GetTableFormatOk returns a tuple with the TableFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableFormat

`func (o *OCRRequest) SetTableFormat(v string)`

SetTableFormat sets TableFormat field to given value.

### HasTableFormat

`func (o *OCRRequest) HasTableFormat() bool`

HasTableFormat returns a boolean if a field has been set.

### SetTableFormatNil

`func (o *OCRRequest) SetTableFormatNil(b bool)`

 SetTableFormatNil sets the value for TableFormat to be an explicit nil

### UnsetTableFormat
`func (o *OCRRequest) UnsetTableFormat()`

UnsetTableFormat ensures that no value is present for TableFormat, not even an explicit nil
### GetExtractHeader

`func (o *OCRRequest) GetExtractHeader() bool`

GetExtractHeader returns the ExtractHeader field if non-nil, zero value otherwise.

### GetExtractHeaderOk

`func (o *OCRRequest) GetExtractHeaderOk() (*bool, bool)`

GetExtractHeaderOk returns a tuple with the ExtractHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractHeader

`func (o *OCRRequest) SetExtractHeader(v bool)`

SetExtractHeader sets ExtractHeader field to given value.

### HasExtractHeader

`func (o *OCRRequest) HasExtractHeader() bool`

HasExtractHeader returns a boolean if a field has been set.

### GetExtractFooter

`func (o *OCRRequest) GetExtractFooter() bool`

GetExtractFooter returns the ExtractFooter field if non-nil, zero value otherwise.

### GetExtractFooterOk

`func (o *OCRRequest) GetExtractFooterOk() (*bool, bool)`

GetExtractFooterOk returns a tuple with the ExtractFooter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractFooter

`func (o *OCRRequest) SetExtractFooter(v bool)`

SetExtractFooter sets ExtractFooter field to given value.

### HasExtractFooter

`func (o *OCRRequest) HasExtractFooter() bool`

HasExtractFooter returns a boolean if a field has been set.

### GetIncludeBlocks

`func (o *OCRRequest) GetIncludeBlocks() bool`

GetIncludeBlocks returns the IncludeBlocks field if non-nil, zero value otherwise.

### GetIncludeBlocksOk

`func (o *OCRRequest) GetIncludeBlocksOk() (*bool, bool)`

GetIncludeBlocksOk returns a tuple with the IncludeBlocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeBlocks

`func (o *OCRRequest) SetIncludeBlocks(v bool)`

SetIncludeBlocks sets IncludeBlocks field to given value.

### HasIncludeBlocks

`func (o *OCRRequest) HasIncludeBlocks() bool`

HasIncludeBlocks returns a boolean if a field has been set.

### GetConfidenceScoresGranularity

`func (o *OCRRequest) GetConfidenceScoresGranularity() string`

GetConfidenceScoresGranularity returns the ConfidenceScoresGranularity field if non-nil, zero value otherwise.

### GetConfidenceScoresGranularityOk

`func (o *OCRRequest) GetConfidenceScoresGranularityOk() (*string, bool)`

GetConfidenceScoresGranularityOk returns a tuple with the ConfidenceScoresGranularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidenceScoresGranularity

`func (o *OCRRequest) SetConfidenceScoresGranularity(v string)`

SetConfidenceScoresGranularity sets ConfidenceScoresGranularity field to given value.

### HasConfidenceScoresGranularity

`func (o *OCRRequest) HasConfidenceScoresGranularity() bool`

HasConfidenceScoresGranularity returns a boolean if a field has been set.

### SetConfidenceScoresGranularityNil

`func (o *OCRRequest) SetConfidenceScoresGranularityNil(b bool)`

 SetConfidenceScoresGranularityNil sets the value for ConfidenceScoresGranularity to be an explicit nil

### UnsetConfidenceScoresGranularity
`func (o *OCRRequest) UnsetConfidenceScoresGranularity()`

UnsetConfidenceScoresGranularity ensures that no value is present for ConfidenceScoresGranularity, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


