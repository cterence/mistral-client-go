# ResponsesInputTextContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | **string** |  | 
**Type** | **string** |  | 
**PromptCacheBreakpoint** | Pointer to [**NullableResponsesPromptCacheBreakpoint**](ResponsesPromptCacheBreakpoint.md) |  | [optional] 

## Methods

### NewResponsesInputTextContent

`func NewResponsesInputTextContent(text string, type_ string, ) *ResponsesInputTextContent`

NewResponsesInputTextContent instantiates a new ResponsesInputTextContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsesInputTextContentWithDefaults

`func NewResponsesInputTextContentWithDefaults() *ResponsesInputTextContent`

NewResponsesInputTextContentWithDefaults instantiates a new ResponsesInputTextContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *ResponsesInputTextContent) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ResponsesInputTextContent) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ResponsesInputTextContent) SetText(v string)`

SetText sets Text field to given value.


### GetType

`func (o *ResponsesInputTextContent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResponsesInputTextContent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResponsesInputTextContent) SetType(v string)`

SetType sets Type field to given value.


### GetPromptCacheBreakpoint

`func (o *ResponsesInputTextContent) GetPromptCacheBreakpoint() ResponsesPromptCacheBreakpoint`

GetPromptCacheBreakpoint returns the PromptCacheBreakpoint field if non-nil, zero value otherwise.

### GetPromptCacheBreakpointOk

`func (o *ResponsesInputTextContent) GetPromptCacheBreakpointOk() (*ResponsesPromptCacheBreakpoint, bool)`

GetPromptCacheBreakpointOk returns a tuple with the PromptCacheBreakpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptCacheBreakpoint

`func (o *ResponsesInputTextContent) SetPromptCacheBreakpoint(v ResponsesPromptCacheBreakpoint)`

SetPromptCacheBreakpoint sets PromptCacheBreakpoint field to given value.

### HasPromptCacheBreakpoint

`func (o *ResponsesInputTextContent) HasPromptCacheBreakpoint() bool`

HasPromptCacheBreakpoint returns a boolean if a field has been set.

### SetPromptCacheBreakpointNil

`func (o *ResponsesInputTextContent) SetPromptCacheBreakpointNil(b bool)`

 SetPromptCacheBreakpointNil sets the value for PromptCacheBreakpoint to be an explicit nil

### UnsetPromptCacheBreakpoint
`func (o *ResponsesInputTextContent) UnsetPromptCacheBreakpoint()`

UnsetPromptCacheBreakpoint ensures that no value is present for PromptCacheBreakpoint, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


