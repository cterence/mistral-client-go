# ResponsesInputImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Detail** | **string** |  | 
**Type** | **string** |  | 
**FileId** | Pointer to **NullableString** |  | [optional] 
**ImageUrl** | Pointer to **NullableString** |  | [optional] 
**PromptCacheBreakpoint** | Pointer to [**NullableResponsesPromptCacheBreakpoint**](ResponsesPromptCacheBreakpoint.md) |  | [optional] 

## Methods

### NewResponsesInputImage

`func NewResponsesInputImage(detail string, type_ string, ) *ResponsesInputImage`

NewResponsesInputImage instantiates a new ResponsesInputImage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsesInputImageWithDefaults

`func NewResponsesInputImageWithDefaults() *ResponsesInputImage`

NewResponsesInputImageWithDefaults instantiates a new ResponsesInputImage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetail

`func (o *ResponsesInputImage) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ResponsesInputImage) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ResponsesInputImage) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetType

`func (o *ResponsesInputImage) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResponsesInputImage) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResponsesInputImage) SetType(v string)`

SetType sets Type field to given value.


### GetFileId

`func (o *ResponsesInputImage) GetFileId() string`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *ResponsesInputImage) GetFileIdOk() (*string, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *ResponsesInputImage) SetFileId(v string)`

SetFileId sets FileId field to given value.

### HasFileId

`func (o *ResponsesInputImage) HasFileId() bool`

HasFileId returns a boolean if a field has been set.

### SetFileIdNil

`func (o *ResponsesInputImage) SetFileIdNil(b bool)`

 SetFileIdNil sets the value for FileId to be an explicit nil

### UnsetFileId
`func (o *ResponsesInputImage) UnsetFileId()`

UnsetFileId ensures that no value is present for FileId, not even an explicit nil
### GetImageUrl

`func (o *ResponsesInputImage) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *ResponsesInputImage) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *ResponsesInputImage) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *ResponsesInputImage) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### SetImageUrlNil

`func (o *ResponsesInputImage) SetImageUrlNil(b bool)`

 SetImageUrlNil sets the value for ImageUrl to be an explicit nil

### UnsetImageUrl
`func (o *ResponsesInputImage) UnsetImageUrl()`

UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
### GetPromptCacheBreakpoint

`func (o *ResponsesInputImage) GetPromptCacheBreakpoint() ResponsesPromptCacheBreakpoint`

GetPromptCacheBreakpoint returns the PromptCacheBreakpoint field if non-nil, zero value otherwise.

### GetPromptCacheBreakpointOk

`func (o *ResponsesInputImage) GetPromptCacheBreakpointOk() (*ResponsesPromptCacheBreakpoint, bool)`

GetPromptCacheBreakpointOk returns a tuple with the PromptCacheBreakpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptCacheBreakpoint

`func (o *ResponsesInputImage) SetPromptCacheBreakpoint(v ResponsesPromptCacheBreakpoint)`

SetPromptCacheBreakpoint sets PromptCacheBreakpoint field to given value.

### HasPromptCacheBreakpoint

`func (o *ResponsesInputImage) HasPromptCacheBreakpoint() bool`

HasPromptCacheBreakpoint returns a boolean if a field has been set.

### SetPromptCacheBreakpointNil

`func (o *ResponsesInputImage) SetPromptCacheBreakpointNil(b bool)`

 SetPromptCacheBreakpointNil sets the value for PromptCacheBreakpoint to be an explicit nil

### UnsetPromptCacheBreakpoint
`func (o *ResponsesInputImage) UnsetPromptCacheBreakpoint()`

UnsetPromptCacheBreakpoint ensures that no value is present for PromptCacheBreakpoint, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


