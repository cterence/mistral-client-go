# Content5Inner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Text** | **string** |  | 
**Annotations** | Pointer to [**Annotations**](Annotations.md) |  | [optional] 
**Meta** | Pointer to **map[string]interface{}** |  | [optional] 
**Data** | **string** |  | 
**MimeType** | **string** |  | 
**Name** | **string** |  | 
**Title** | Pointer to **string** |  | [optional] 
**Uri** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **int32** |  | [optional] 
**Icons** | Pointer to [**[]MCPServerIcon**](MCPServerIcon.md) |  | [optional] 
**Resource** | [**Resource**](Resource.md) |  | 

## Methods

### NewContent5Inner

`func NewContent5Inner(type_ string, text string, data string, mimeType string, name string, uri string, resource Resource, ) *Content5Inner`

NewContent5Inner instantiates a new Content5Inner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContent5InnerWithDefaults

`func NewContent5InnerWithDefaults() *Content5Inner`

NewContent5InnerWithDefaults instantiates a new Content5Inner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Content5Inner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Content5Inner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Content5Inner) SetType(v string)`

SetType sets Type field to given value.


### GetText

`func (o *Content5Inner) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *Content5Inner) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *Content5Inner) SetText(v string)`

SetText sets Text field to given value.


### GetAnnotations

`func (o *Content5Inner) GetAnnotations() Annotations`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *Content5Inner) GetAnnotationsOk() (*Annotations, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *Content5Inner) SetAnnotations(v Annotations)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *Content5Inner) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### GetMeta

`func (o *Content5Inner) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *Content5Inner) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *Content5Inner) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *Content5Inner) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetData

`func (o *Content5Inner) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Content5Inner) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Content5Inner) SetData(v string)`

SetData sets Data field to given value.


### GetMimeType

`func (o *Content5Inner) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *Content5Inner) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *Content5Inner) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.


### GetName

`func (o *Content5Inner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Content5Inner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Content5Inner) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *Content5Inner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Content5Inner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Content5Inner) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Content5Inner) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUri

`func (o *Content5Inner) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *Content5Inner) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *Content5Inner) SetUri(v string)`

SetUri sets Uri field to given value.


### GetDescription

`func (o *Content5Inner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Content5Inner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Content5Inner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Content5Inner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSize

`func (o *Content5Inner) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *Content5Inner) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *Content5Inner) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *Content5Inner) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetIcons

`func (o *Content5Inner) GetIcons() []MCPServerIcon`

GetIcons returns the Icons field if non-nil, zero value otherwise.

### GetIconsOk

`func (o *Content5Inner) GetIconsOk() (*[]MCPServerIcon, bool)`

GetIconsOk returns a tuple with the Icons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcons

`func (o *Content5Inner) SetIcons(v []MCPServerIcon)`

SetIcons sets Icons field to given value.

### HasIcons

`func (o *Content5Inner) HasIcons() bool`

HasIcons returns a boolean if a field has been set.

### GetResource

`func (o *Content5Inner) GetResource() Resource`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *Content5Inner) GetResourceOk() (*Resource, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *Content5Inner) SetResource(v Resource)`

SetResource sets Resource field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


