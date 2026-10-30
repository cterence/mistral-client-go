# Document

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**LibraryId** | **string** |  | 
**Hash** | **NullableString** |  | 
**MimeType** | **NullableString** |  | 
**Extension** | **NullableString** |  | 
**Size** | **NullableInt32** |  | 
**Name** | **string** |  | 
**Summary** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**LastProcessedAt** | Pointer to **NullableTime** |  | [optional] 
**NumberOfPages** | Pointer to **NullableInt32** |  | [optional] 
**ProcessStatus** | [**ProcessStatus**](ProcessStatus.md) | Processing status of the document.  | 
**UploadedById** | **NullableString** |  | 
**UploadedByType** | **NullableString** | Filter by agent name | 
**TokensProcessingMainContent** | Pointer to **NullableInt32** |  | [optional] 
**TokensProcessingSummary** | Pointer to **NullableInt32** |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 
**Attributes** | Pointer to **map[string]interface{}** | Output JSON schema of the query&#39;s model | [optional] 
**ExpiresAt** | Pointer to **NullableTime** | If set, the document will be automatically deleted after this date. | [optional] 
**ProcessingStatus** | **string** |  | [readonly] 
**TokensProcessingTotal** | **int32** |  | [readonly] 

## Methods

### NewDocument

`func NewDocument(id string, libraryId string, hash NullableString, mimeType NullableString, extension NullableString, size NullableInt32, name string, createdAt time.Time, processStatus ProcessStatus, uploadedById NullableString, uploadedByType NullableString, processingStatus string, tokensProcessingTotal int32, ) *Document`

NewDocument instantiates a new Document object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentWithDefaults

`func NewDocumentWithDefaults() *Document`

NewDocumentWithDefaults instantiates a new Document object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Document) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Document) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Document) SetId(v string)`

SetId sets Id field to given value.


### GetLibraryId

`func (o *Document) GetLibraryId() string`

GetLibraryId returns the LibraryId field if non-nil, zero value otherwise.

### GetLibraryIdOk

`func (o *Document) GetLibraryIdOk() (*string, bool)`

GetLibraryIdOk returns a tuple with the LibraryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraryId

`func (o *Document) SetLibraryId(v string)`

SetLibraryId sets LibraryId field to given value.


### GetHash

`func (o *Document) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *Document) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *Document) SetHash(v string)`

SetHash sets Hash field to given value.


### SetHashNil

`func (o *Document) SetHashNil(b bool)`

 SetHashNil sets the value for Hash to be an explicit nil

### UnsetHash
`func (o *Document) UnsetHash()`

UnsetHash ensures that no value is present for Hash, not even an explicit nil
### GetMimeType

`func (o *Document) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *Document) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *Document) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.


### SetMimeTypeNil

`func (o *Document) SetMimeTypeNil(b bool)`

 SetMimeTypeNil sets the value for MimeType to be an explicit nil

### UnsetMimeType
`func (o *Document) UnsetMimeType()`

UnsetMimeType ensures that no value is present for MimeType, not even an explicit nil
### GetExtension

`func (o *Document) GetExtension() string`

GetExtension returns the Extension field if non-nil, zero value otherwise.

### GetExtensionOk

`func (o *Document) GetExtensionOk() (*string, bool)`

GetExtensionOk returns a tuple with the Extension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtension

`func (o *Document) SetExtension(v string)`

SetExtension sets Extension field to given value.


### SetExtensionNil

`func (o *Document) SetExtensionNil(b bool)`

 SetExtensionNil sets the value for Extension to be an explicit nil

### UnsetExtension
`func (o *Document) UnsetExtension()`

UnsetExtension ensures that no value is present for Extension, not even an explicit nil
### GetSize

`func (o *Document) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Document) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Document) SetSize(v int32)`

SetSize sets Size field to given value.


### SetSizeNil

`func (o *Document) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *Document) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetName

`func (o *Document) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Document) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Document) SetName(v string)`

SetName sets Name field to given value.


### GetSummary

`func (o *Document) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *Document) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *Document) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *Document) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### SetSummaryNil

`func (o *Document) SetSummaryNil(b bool)`

 SetSummaryNil sets the value for Summary to be an explicit nil

### UnsetSummary
`func (o *Document) UnsetSummary()`

UnsetSummary ensures that no value is present for Summary, not even an explicit nil
### GetCreatedAt

`func (o *Document) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Document) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Document) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetLastProcessedAt

`func (o *Document) GetLastProcessedAt() time.Time`

GetLastProcessedAt returns the LastProcessedAt field if non-nil, zero value otherwise.

### GetLastProcessedAtOk

`func (o *Document) GetLastProcessedAtOk() (*time.Time, bool)`

GetLastProcessedAtOk returns a tuple with the LastProcessedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastProcessedAt

`func (o *Document) SetLastProcessedAt(v time.Time)`

SetLastProcessedAt sets LastProcessedAt field to given value.

### HasLastProcessedAt

`func (o *Document) HasLastProcessedAt() bool`

HasLastProcessedAt returns a boolean if a field has been set.

### SetLastProcessedAtNil

`func (o *Document) SetLastProcessedAtNil(b bool)`

 SetLastProcessedAtNil sets the value for LastProcessedAt to be an explicit nil

### UnsetLastProcessedAt
`func (o *Document) UnsetLastProcessedAt()`

UnsetLastProcessedAt ensures that no value is present for LastProcessedAt, not even an explicit nil
### GetNumberOfPages

`func (o *Document) GetNumberOfPages() int32`

GetNumberOfPages returns the NumberOfPages field if non-nil, zero value otherwise.

### GetNumberOfPagesOk

`func (o *Document) GetNumberOfPagesOk() (*int32, bool)`

GetNumberOfPagesOk returns a tuple with the NumberOfPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfPages

`func (o *Document) SetNumberOfPages(v int32)`

SetNumberOfPages sets NumberOfPages field to given value.

### HasNumberOfPages

`func (o *Document) HasNumberOfPages() bool`

HasNumberOfPages returns a boolean if a field has been set.

### SetNumberOfPagesNil

`func (o *Document) SetNumberOfPagesNil(b bool)`

 SetNumberOfPagesNil sets the value for NumberOfPages to be an explicit nil

### UnsetNumberOfPages
`func (o *Document) UnsetNumberOfPages()`

UnsetNumberOfPages ensures that no value is present for NumberOfPages, not even an explicit nil
### GetProcessStatus

`func (o *Document) GetProcessStatus() ProcessStatus`

GetProcessStatus returns the ProcessStatus field if non-nil, zero value otherwise.

### GetProcessStatusOk

`func (o *Document) GetProcessStatusOk() (*ProcessStatus, bool)`

GetProcessStatusOk returns a tuple with the ProcessStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessStatus

`func (o *Document) SetProcessStatus(v ProcessStatus)`

SetProcessStatus sets ProcessStatus field to given value.


### GetUploadedById

`func (o *Document) GetUploadedById() string`

GetUploadedById returns the UploadedById field if non-nil, zero value otherwise.

### GetUploadedByIdOk

`func (o *Document) GetUploadedByIdOk() (*string, bool)`

GetUploadedByIdOk returns a tuple with the UploadedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadedById

`func (o *Document) SetUploadedById(v string)`

SetUploadedById sets UploadedById field to given value.


### SetUploadedByIdNil

`func (o *Document) SetUploadedByIdNil(b bool)`

 SetUploadedByIdNil sets the value for UploadedById to be an explicit nil

### UnsetUploadedById
`func (o *Document) UnsetUploadedById()`

UnsetUploadedById ensures that no value is present for UploadedById, not even an explicit nil
### GetUploadedByType

`func (o *Document) GetUploadedByType() string`

GetUploadedByType returns the UploadedByType field if non-nil, zero value otherwise.

### GetUploadedByTypeOk

`func (o *Document) GetUploadedByTypeOk() (*string, bool)`

GetUploadedByTypeOk returns a tuple with the UploadedByType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadedByType

`func (o *Document) SetUploadedByType(v string)`

SetUploadedByType sets UploadedByType field to given value.


### SetUploadedByTypeNil

`func (o *Document) SetUploadedByTypeNil(b bool)`

 SetUploadedByTypeNil sets the value for UploadedByType to be an explicit nil

### UnsetUploadedByType
`func (o *Document) UnsetUploadedByType()`

UnsetUploadedByType ensures that no value is present for UploadedByType, not even an explicit nil
### GetTokensProcessingMainContent

`func (o *Document) GetTokensProcessingMainContent() int32`

GetTokensProcessingMainContent returns the TokensProcessingMainContent field if non-nil, zero value otherwise.

### GetTokensProcessingMainContentOk

`func (o *Document) GetTokensProcessingMainContentOk() (*int32, bool)`

GetTokensProcessingMainContentOk returns a tuple with the TokensProcessingMainContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokensProcessingMainContent

`func (o *Document) SetTokensProcessingMainContent(v int32)`

SetTokensProcessingMainContent sets TokensProcessingMainContent field to given value.

### HasTokensProcessingMainContent

`func (o *Document) HasTokensProcessingMainContent() bool`

HasTokensProcessingMainContent returns a boolean if a field has been set.

### SetTokensProcessingMainContentNil

`func (o *Document) SetTokensProcessingMainContentNil(b bool)`

 SetTokensProcessingMainContentNil sets the value for TokensProcessingMainContent to be an explicit nil

### UnsetTokensProcessingMainContent
`func (o *Document) UnsetTokensProcessingMainContent()`

UnsetTokensProcessingMainContent ensures that no value is present for TokensProcessingMainContent, not even an explicit nil
### GetTokensProcessingSummary

`func (o *Document) GetTokensProcessingSummary() int32`

GetTokensProcessingSummary returns the TokensProcessingSummary field if non-nil, zero value otherwise.

### GetTokensProcessingSummaryOk

`func (o *Document) GetTokensProcessingSummaryOk() (*int32, bool)`

GetTokensProcessingSummaryOk returns a tuple with the TokensProcessingSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokensProcessingSummary

`func (o *Document) SetTokensProcessingSummary(v int32)`

SetTokensProcessingSummary sets TokensProcessingSummary field to given value.

### HasTokensProcessingSummary

`func (o *Document) HasTokensProcessingSummary() bool`

HasTokensProcessingSummary returns a boolean if a field has been set.

### SetTokensProcessingSummaryNil

`func (o *Document) SetTokensProcessingSummaryNil(b bool)`

 SetTokensProcessingSummaryNil sets the value for TokensProcessingSummary to be an explicit nil

### UnsetTokensProcessingSummary
`func (o *Document) UnsetTokensProcessingSummary()`

UnsetTokensProcessingSummary ensures that no value is present for TokensProcessingSummary, not even an explicit nil
### GetUrl

`func (o *Document) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Document) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Document) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Document) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *Document) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *Document) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetAttributes

`func (o *Document) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Document) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Document) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *Document) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *Document) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *Document) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil
### GetExpiresAt

`func (o *Document) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *Document) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *Document) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *Document) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *Document) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *Document) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetProcessingStatus

`func (o *Document) GetProcessingStatus() string`

GetProcessingStatus returns the ProcessingStatus field if non-nil, zero value otherwise.

### GetProcessingStatusOk

`func (o *Document) GetProcessingStatusOk() (*string, bool)`

GetProcessingStatusOk returns a tuple with the ProcessingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingStatus

`func (o *Document) SetProcessingStatus(v string)`

SetProcessingStatus sets ProcessingStatus field to given value.


### GetTokensProcessingTotal

`func (o *Document) GetTokensProcessingTotal() int32`

GetTokensProcessingTotal returns the TokensProcessingTotal field if non-nil, zero value otherwise.

### GetTokensProcessingTotalOk

`func (o *Document) GetTokensProcessingTotalOk() (*int32, bool)`

GetTokensProcessingTotalOk returns a tuple with the TokensProcessingTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokensProcessingTotal

`func (o *Document) SetTokensProcessingTotal(v int32)`

SetTokensProcessingTotal sets TokensProcessingTotal field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


