# Output1AnyOfInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | **string** |  | 
**Type** | **string** |  | 
**PromptCacheBreakpoint** | Pointer to [**NullableResponsesPromptCacheBreakpoint**](ResponsesPromptCacheBreakpoint.md) |  | [optional] 
**Detail** | Pointer to **NullableString** |  | [optional] 
**FileId** | Pointer to **NullableString** |  | [optional] 
**ImageUrl** | Pointer to **NullableString** |  | [optional] 
**FileData** | Pointer to **NullableString** |  | [optional] 
**FileUrl** | Pointer to **NullableString** |  | [optional] 
**Filename** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewOutput1AnyOfInner

`func NewOutput1AnyOfInner(text string, type_ string, ) *Output1AnyOfInner`

NewOutput1AnyOfInner instantiates a new Output1AnyOfInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutput1AnyOfInnerWithDefaults

`func NewOutput1AnyOfInnerWithDefaults() *Output1AnyOfInner`

NewOutput1AnyOfInnerWithDefaults instantiates a new Output1AnyOfInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *Output1AnyOfInner) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *Output1AnyOfInner) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *Output1AnyOfInner) SetText(v string)`

SetText sets Text field to given value.


### GetType

`func (o *Output1AnyOfInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Output1AnyOfInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Output1AnyOfInner) SetType(v string)`

SetType sets Type field to given value.


### GetPromptCacheBreakpoint

`func (o *Output1AnyOfInner) GetPromptCacheBreakpoint() ResponsesPromptCacheBreakpoint`

GetPromptCacheBreakpoint returns the PromptCacheBreakpoint field if non-nil, zero value otherwise.

### GetPromptCacheBreakpointOk

`func (o *Output1AnyOfInner) GetPromptCacheBreakpointOk() (*ResponsesPromptCacheBreakpoint, bool)`

GetPromptCacheBreakpointOk returns a tuple with the PromptCacheBreakpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptCacheBreakpoint

`func (o *Output1AnyOfInner) SetPromptCacheBreakpoint(v ResponsesPromptCacheBreakpoint)`

SetPromptCacheBreakpoint sets PromptCacheBreakpoint field to given value.

### HasPromptCacheBreakpoint

`func (o *Output1AnyOfInner) HasPromptCacheBreakpoint() bool`

HasPromptCacheBreakpoint returns a boolean if a field has been set.

### SetPromptCacheBreakpointNil

`func (o *Output1AnyOfInner) SetPromptCacheBreakpointNil(b bool)`

 SetPromptCacheBreakpointNil sets the value for PromptCacheBreakpoint to be an explicit nil

### UnsetPromptCacheBreakpoint
`func (o *Output1AnyOfInner) UnsetPromptCacheBreakpoint()`

UnsetPromptCacheBreakpoint ensures that no value is present for PromptCacheBreakpoint, not even an explicit nil
### GetDetail

`func (o *Output1AnyOfInner) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *Output1AnyOfInner) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *Output1AnyOfInner) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *Output1AnyOfInner) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### SetDetailNil

`func (o *Output1AnyOfInner) SetDetailNil(b bool)`

 SetDetailNil sets the value for Detail to be an explicit nil

### UnsetDetail
`func (o *Output1AnyOfInner) UnsetDetail()`

UnsetDetail ensures that no value is present for Detail, not even an explicit nil
### GetFileId

`func (o *Output1AnyOfInner) GetFileId() string`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *Output1AnyOfInner) GetFileIdOk() (*string, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *Output1AnyOfInner) SetFileId(v string)`

SetFileId sets FileId field to given value.

### HasFileId

`func (o *Output1AnyOfInner) HasFileId() bool`

HasFileId returns a boolean if a field has been set.

### SetFileIdNil

`func (o *Output1AnyOfInner) SetFileIdNil(b bool)`

 SetFileIdNil sets the value for FileId to be an explicit nil

### UnsetFileId
`func (o *Output1AnyOfInner) UnsetFileId()`

UnsetFileId ensures that no value is present for FileId, not even an explicit nil
### GetImageUrl

`func (o *Output1AnyOfInner) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *Output1AnyOfInner) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *Output1AnyOfInner) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *Output1AnyOfInner) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### SetImageUrlNil

`func (o *Output1AnyOfInner) SetImageUrlNil(b bool)`

 SetImageUrlNil sets the value for ImageUrl to be an explicit nil

### UnsetImageUrl
`func (o *Output1AnyOfInner) UnsetImageUrl()`

UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
### GetFileData

`func (o *Output1AnyOfInner) GetFileData() string`

GetFileData returns the FileData field if non-nil, zero value otherwise.

### GetFileDataOk

`func (o *Output1AnyOfInner) GetFileDataOk() (*string, bool)`

GetFileDataOk returns a tuple with the FileData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileData

`func (o *Output1AnyOfInner) SetFileData(v string)`

SetFileData sets FileData field to given value.

### HasFileData

`func (o *Output1AnyOfInner) HasFileData() bool`

HasFileData returns a boolean if a field has been set.

### SetFileDataNil

`func (o *Output1AnyOfInner) SetFileDataNil(b bool)`

 SetFileDataNil sets the value for FileData to be an explicit nil

### UnsetFileData
`func (o *Output1AnyOfInner) UnsetFileData()`

UnsetFileData ensures that no value is present for FileData, not even an explicit nil
### GetFileUrl

`func (o *Output1AnyOfInner) GetFileUrl() string`

GetFileUrl returns the FileUrl field if non-nil, zero value otherwise.

### GetFileUrlOk

`func (o *Output1AnyOfInner) GetFileUrlOk() (*string, bool)`

GetFileUrlOk returns a tuple with the FileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUrl

`func (o *Output1AnyOfInner) SetFileUrl(v string)`

SetFileUrl sets FileUrl field to given value.

### HasFileUrl

`func (o *Output1AnyOfInner) HasFileUrl() bool`

HasFileUrl returns a boolean if a field has been set.

### SetFileUrlNil

`func (o *Output1AnyOfInner) SetFileUrlNil(b bool)`

 SetFileUrlNil sets the value for FileUrl to be an explicit nil

### UnsetFileUrl
`func (o *Output1AnyOfInner) UnsetFileUrl()`

UnsetFileUrl ensures that no value is present for FileUrl, not even an explicit nil
### GetFilename

`func (o *Output1AnyOfInner) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *Output1AnyOfInner) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *Output1AnyOfInner) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *Output1AnyOfInner) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### SetFilenameNil

`func (o *Output1AnyOfInner) SetFilenameNil(b bool)`

 SetFilenameNil sets the value for Filename to be an explicit nil

### UnsetFilename
`func (o *Output1AnyOfInner) UnsetFilename()`

UnsetFilename ensures that no value is present for Filename, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


