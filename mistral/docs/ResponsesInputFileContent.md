# ResponsesInputFileContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Detail** | Pointer to **NullableString** |  | [optional] 
**FileData** | Pointer to **NullableString** |  | [optional] 
**FileId** | Pointer to **NullableString** |  | [optional] 
**FileUrl** | Pointer to **NullableString** |  | [optional] 
**Filename** | Pointer to **NullableString** |  | [optional] 
**PromptCacheBreakpoint** | Pointer to [**NullableResponsesPromptCacheBreakpoint**](ResponsesPromptCacheBreakpoint.md) |  | [optional] 

## Methods

### NewResponsesInputFileContent

`func NewResponsesInputFileContent(type_ string, ) *ResponsesInputFileContent`

NewResponsesInputFileContent instantiates a new ResponsesInputFileContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsesInputFileContentWithDefaults

`func NewResponsesInputFileContentWithDefaults() *ResponsesInputFileContent`

NewResponsesInputFileContentWithDefaults instantiates a new ResponsesInputFileContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ResponsesInputFileContent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResponsesInputFileContent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResponsesInputFileContent) SetType(v string)`

SetType sets Type field to given value.


### GetDetail

`func (o *ResponsesInputFileContent) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ResponsesInputFileContent) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ResponsesInputFileContent) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *ResponsesInputFileContent) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### SetDetailNil

`func (o *ResponsesInputFileContent) SetDetailNil(b bool)`

 SetDetailNil sets the value for Detail to be an explicit nil

### UnsetDetail
`func (o *ResponsesInputFileContent) UnsetDetail()`

UnsetDetail ensures that no value is present for Detail, not even an explicit nil
### GetFileData

`func (o *ResponsesInputFileContent) GetFileData() string`

GetFileData returns the FileData field if non-nil, zero value otherwise.

### GetFileDataOk

`func (o *ResponsesInputFileContent) GetFileDataOk() (*string, bool)`

GetFileDataOk returns a tuple with the FileData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileData

`func (o *ResponsesInputFileContent) SetFileData(v string)`

SetFileData sets FileData field to given value.

### HasFileData

`func (o *ResponsesInputFileContent) HasFileData() bool`

HasFileData returns a boolean if a field has been set.

### SetFileDataNil

`func (o *ResponsesInputFileContent) SetFileDataNil(b bool)`

 SetFileDataNil sets the value for FileData to be an explicit nil

### UnsetFileData
`func (o *ResponsesInputFileContent) UnsetFileData()`

UnsetFileData ensures that no value is present for FileData, not even an explicit nil
### GetFileId

`func (o *ResponsesInputFileContent) GetFileId() string`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *ResponsesInputFileContent) GetFileIdOk() (*string, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *ResponsesInputFileContent) SetFileId(v string)`

SetFileId sets FileId field to given value.

### HasFileId

`func (o *ResponsesInputFileContent) HasFileId() bool`

HasFileId returns a boolean if a field has been set.

### SetFileIdNil

`func (o *ResponsesInputFileContent) SetFileIdNil(b bool)`

 SetFileIdNil sets the value for FileId to be an explicit nil

### UnsetFileId
`func (o *ResponsesInputFileContent) UnsetFileId()`

UnsetFileId ensures that no value is present for FileId, not even an explicit nil
### GetFileUrl

`func (o *ResponsesInputFileContent) GetFileUrl() string`

GetFileUrl returns the FileUrl field if non-nil, zero value otherwise.

### GetFileUrlOk

`func (o *ResponsesInputFileContent) GetFileUrlOk() (*string, bool)`

GetFileUrlOk returns a tuple with the FileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUrl

`func (o *ResponsesInputFileContent) SetFileUrl(v string)`

SetFileUrl sets FileUrl field to given value.

### HasFileUrl

`func (o *ResponsesInputFileContent) HasFileUrl() bool`

HasFileUrl returns a boolean if a field has been set.

### SetFileUrlNil

`func (o *ResponsesInputFileContent) SetFileUrlNil(b bool)`

 SetFileUrlNil sets the value for FileUrl to be an explicit nil

### UnsetFileUrl
`func (o *ResponsesInputFileContent) UnsetFileUrl()`

UnsetFileUrl ensures that no value is present for FileUrl, not even an explicit nil
### GetFilename

`func (o *ResponsesInputFileContent) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *ResponsesInputFileContent) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *ResponsesInputFileContent) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *ResponsesInputFileContent) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### SetFilenameNil

`func (o *ResponsesInputFileContent) SetFilenameNil(b bool)`

 SetFilenameNil sets the value for Filename to be an explicit nil

### UnsetFilename
`func (o *ResponsesInputFileContent) UnsetFilename()`

UnsetFilename ensures that no value is present for Filename, not even an explicit nil
### GetPromptCacheBreakpoint

`func (o *ResponsesInputFileContent) GetPromptCacheBreakpoint() ResponsesPromptCacheBreakpoint`

GetPromptCacheBreakpoint returns the PromptCacheBreakpoint field if non-nil, zero value otherwise.

### GetPromptCacheBreakpointOk

`func (o *ResponsesInputFileContent) GetPromptCacheBreakpointOk() (*ResponsesPromptCacheBreakpoint, bool)`

GetPromptCacheBreakpointOk returns a tuple with the PromptCacheBreakpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptCacheBreakpoint

`func (o *ResponsesInputFileContent) SetPromptCacheBreakpoint(v ResponsesPromptCacheBreakpoint)`

SetPromptCacheBreakpoint sets PromptCacheBreakpoint field to given value.

### HasPromptCacheBreakpoint

`func (o *ResponsesInputFileContent) HasPromptCacheBreakpoint() bool`

HasPromptCacheBreakpoint returns a boolean if a field has been set.

### SetPromptCacheBreakpointNil

`func (o *ResponsesInputFileContent) SetPromptCacheBreakpointNil(b bool)`

 SetPromptCacheBreakpointNil sets the value for PromptCacheBreakpoint to be an explicit nil

### UnsetPromptCacheBreakpoint
`func (o *ResponsesInputFileContent) UnsetPromptCacheBreakpoint()`

UnsetPromptCacheBreakpoint ensures that no value is present for PromptCacheBreakpoint, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


