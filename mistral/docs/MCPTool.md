# MCPTool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Title** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**InputSchema** | **map[string]interface{}** |  | 
**OutputSchema** | Pointer to **map[string]interface{}** |  | [optional] 
**Icons** | Pointer to [**[]MCPServerIcon**](MCPServerIcon.md) |  | [optional] 
**Annotations** | Pointer to [**NullableToolAnnotations**](ToolAnnotations.md) |  | [optional] 
**Meta** | Pointer to [**NullableMCPToolMeta**](MCPToolMeta.md) |  | [optional] 
**Execution** | Pointer to [**NullableToolExecution**](ToolExecution.md) |  | [optional] 

## Methods

### NewMCPTool

`func NewMCPTool(name string, inputSchema map[string]interface{}, ) *MCPTool`

NewMCPTool instantiates a new MCPTool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMCPToolWithDefaults

`func NewMCPToolWithDefaults() *MCPTool`

NewMCPToolWithDefaults instantiates a new MCPTool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MCPTool) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MCPTool) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MCPTool) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *MCPTool) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MCPTool) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *MCPTool) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *MCPTool) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *MCPTool) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *MCPTool) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *MCPTool) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MCPTool) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MCPTool) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MCPTool) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MCPTool) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MCPTool) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetInputSchema

`func (o *MCPTool) GetInputSchema() map[string]interface{}`

GetInputSchema returns the InputSchema field if non-nil, zero value otherwise.

### GetInputSchemaOk

`func (o *MCPTool) GetInputSchemaOk() (*map[string]interface{}, bool)`

GetInputSchemaOk returns a tuple with the InputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputSchema

`func (o *MCPTool) SetInputSchema(v map[string]interface{})`

SetInputSchema sets InputSchema field to given value.


### GetOutputSchema

`func (o *MCPTool) GetOutputSchema() map[string]interface{}`

GetOutputSchema returns the OutputSchema field if non-nil, zero value otherwise.

### GetOutputSchemaOk

`func (o *MCPTool) GetOutputSchemaOk() (*map[string]interface{}, bool)`

GetOutputSchemaOk returns a tuple with the OutputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputSchema

`func (o *MCPTool) SetOutputSchema(v map[string]interface{})`

SetOutputSchema sets OutputSchema field to given value.

### HasOutputSchema

`func (o *MCPTool) HasOutputSchema() bool`

HasOutputSchema returns a boolean if a field has been set.

### SetOutputSchemaNil

`func (o *MCPTool) SetOutputSchemaNil(b bool)`

 SetOutputSchemaNil sets the value for OutputSchema to be an explicit nil

### UnsetOutputSchema
`func (o *MCPTool) UnsetOutputSchema()`

UnsetOutputSchema ensures that no value is present for OutputSchema, not even an explicit nil
### GetIcons

`func (o *MCPTool) GetIcons() []MCPServerIcon`

GetIcons returns the Icons field if non-nil, zero value otherwise.

### GetIconsOk

`func (o *MCPTool) GetIconsOk() (*[]MCPServerIcon, bool)`

GetIconsOk returns a tuple with the Icons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcons

`func (o *MCPTool) SetIcons(v []MCPServerIcon)`

SetIcons sets Icons field to given value.

### HasIcons

`func (o *MCPTool) HasIcons() bool`

HasIcons returns a boolean if a field has been set.

### SetIconsNil

`func (o *MCPTool) SetIconsNil(b bool)`

 SetIconsNil sets the value for Icons to be an explicit nil

### UnsetIcons
`func (o *MCPTool) UnsetIcons()`

UnsetIcons ensures that no value is present for Icons, not even an explicit nil
### GetAnnotations

`func (o *MCPTool) GetAnnotations() ToolAnnotations`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *MCPTool) GetAnnotationsOk() (*ToolAnnotations, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *MCPTool) SetAnnotations(v ToolAnnotations)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *MCPTool) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### SetAnnotationsNil

`func (o *MCPTool) SetAnnotationsNil(b bool)`

 SetAnnotationsNil sets the value for Annotations to be an explicit nil

### UnsetAnnotations
`func (o *MCPTool) UnsetAnnotations()`

UnsetAnnotations ensures that no value is present for Annotations, not even an explicit nil
### GetMeta

`func (o *MCPTool) GetMeta() MCPToolMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MCPTool) GetMetaOk() (*MCPToolMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MCPTool) SetMeta(v MCPToolMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *MCPTool) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### SetMetaNil

`func (o *MCPTool) SetMetaNil(b bool)`

 SetMetaNil sets the value for Meta to be an explicit nil

### UnsetMeta
`func (o *MCPTool) UnsetMeta()`

UnsetMeta ensures that no value is present for Meta, not even an explicit nil
### GetExecution

`func (o *MCPTool) GetExecution() ToolExecution`

GetExecution returns the Execution field if non-nil, zero value otherwise.

### GetExecutionOk

`func (o *MCPTool) GetExecutionOk() (*ToolExecution, bool)`

GetExecutionOk returns a tuple with the Execution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecution

`func (o *MCPTool) SetExecution(v ToolExecution)`

SetExecution sets Execution field to given value.

### HasExecution

`func (o *MCPTool) HasExecution() bool`

HasExecution returns a boolean if a field has been set.

### SetExecutionNil

`func (o *MCPTool) SetExecutionNil(b bool)`

 SetExecutionNil sets the value for Execution to be an explicit nil

### UnsetExecution
`func (o *MCPTool) UnsetExecution()`

UnsetExecution ensures that no value is present for Execution, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


