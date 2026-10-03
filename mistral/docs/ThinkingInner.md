# ThinkingInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "text"]
**Text** | **string** |  | 
**Tool** | [**Tool1**](Tool1.md) |  | 
**Title** | **string** |  | 
**Url** | Pointer to **string** |  | [optional] 
**Favicon** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ReferenceIds** | **[]int32** |  | 

## Methods

### NewThinkingInner

`func NewThinkingInner(text string, tool Tool1, title string, referenceIds []int32, ) *ThinkingInner`

NewThinkingInner instantiates a new ThinkingInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThinkingInnerWithDefaults

`func NewThinkingInnerWithDefaults() *ThinkingInner`

NewThinkingInnerWithDefaults instantiates a new ThinkingInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ThinkingInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ThinkingInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ThinkingInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ThinkingInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetText

`func (o *ThinkingInner) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ThinkingInner) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ThinkingInner) SetText(v string)`

SetText sets Text field to given value.


### GetTool

`func (o *ThinkingInner) GetTool() Tool1`

GetTool returns the Tool field if non-nil, zero value otherwise.

### GetToolOk

`func (o *ThinkingInner) GetToolOk() (*Tool1, bool)`

GetToolOk returns a tuple with the Tool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTool

`func (o *ThinkingInner) SetTool(v Tool1)`

SetTool sets Tool field to given value.


### GetTitle

`func (o *ThinkingInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ThinkingInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ThinkingInner) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetUrl

`func (o *ThinkingInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ThinkingInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ThinkingInner) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ThinkingInner) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetFavicon

`func (o *ThinkingInner) GetFavicon() string`

GetFavicon returns the Favicon field if non-nil, zero value otherwise.

### GetFaviconOk

`func (o *ThinkingInner) GetFaviconOk() (*string, bool)`

GetFaviconOk returns a tuple with the Favicon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavicon

`func (o *ThinkingInner) SetFavicon(v string)`

SetFavicon sets Favicon field to given value.

### HasFavicon

`func (o *ThinkingInner) HasFavicon() bool`

HasFavicon returns a boolean if a field has been set.

### GetDescription

`func (o *ThinkingInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ThinkingInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ThinkingInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ThinkingInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetReferenceIds

`func (o *ThinkingInner) GetReferenceIds() []int32`

GetReferenceIds returns the ReferenceIds field if non-nil, zero value otherwise.

### GetReferenceIdsOk

`func (o *ThinkingInner) GetReferenceIdsOk() (*[]int32, bool)`

GetReferenceIdsOk returns a tuple with the ReferenceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceIds

`func (o *ThinkingInner) SetReferenceIds(v []int32)`

SetReferenceIds sets ReferenceIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


