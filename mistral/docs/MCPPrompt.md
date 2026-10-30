# MCPPrompt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Title** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Arguments** | Pointer to [**[]PromptArgument**](PromptArgument.md) |  | [optional] 
**Icons** | Pointer to [**[]MCPServerIcon**](MCPServerIcon.md) |  | [optional] 
**Meta** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewMCPPrompt

`func NewMCPPrompt(name string, ) *MCPPrompt`

NewMCPPrompt instantiates a new MCPPrompt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMCPPromptWithDefaults

`func NewMCPPromptWithDefaults() *MCPPrompt`

NewMCPPromptWithDefaults instantiates a new MCPPrompt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MCPPrompt) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MCPPrompt) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MCPPrompt) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *MCPPrompt) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MCPPrompt) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *MCPPrompt) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *MCPPrompt) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *MCPPrompt) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *MCPPrompt) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *MCPPrompt) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MCPPrompt) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MCPPrompt) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MCPPrompt) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MCPPrompt) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MCPPrompt) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetArguments

`func (o *MCPPrompt) GetArguments() []PromptArgument`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *MCPPrompt) GetArgumentsOk() (*[]PromptArgument, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *MCPPrompt) SetArguments(v []PromptArgument)`

SetArguments sets Arguments field to given value.

### HasArguments

`func (o *MCPPrompt) HasArguments() bool`

HasArguments returns a boolean if a field has been set.

### SetArgumentsNil

`func (o *MCPPrompt) SetArgumentsNil(b bool)`

 SetArgumentsNil sets the value for Arguments to be an explicit nil

### UnsetArguments
`func (o *MCPPrompt) UnsetArguments()`

UnsetArguments ensures that no value is present for Arguments, not even an explicit nil
### GetIcons

`func (o *MCPPrompt) GetIcons() []MCPServerIcon`

GetIcons returns the Icons field if non-nil, zero value otherwise.

### GetIconsOk

`func (o *MCPPrompt) GetIconsOk() (*[]MCPServerIcon, bool)`

GetIconsOk returns a tuple with the Icons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcons

`func (o *MCPPrompt) SetIcons(v []MCPServerIcon)`

SetIcons sets Icons field to given value.

### HasIcons

`func (o *MCPPrompt) HasIcons() bool`

HasIcons returns a boolean if a field has been set.

### SetIconsNil

`func (o *MCPPrompt) SetIconsNil(b bool)`

 SetIconsNil sets the value for Icons to be an explicit nil

### UnsetIcons
`func (o *MCPPrompt) UnsetIcons()`

UnsetIcons ensures that no value is present for Icons, not even an explicit nil
### GetMeta

`func (o *MCPPrompt) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MCPPrompt) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MCPPrompt) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *MCPPrompt) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### SetMetaNil

`func (o *MCPPrompt) SetMetaNil(b bool)`

 SetMetaNil sets the value for Meta to be an explicit nil

### UnsetMeta
`func (o *MCPPrompt) UnsetMeta()`

UnsetMeta ensures that no value is present for Meta, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


