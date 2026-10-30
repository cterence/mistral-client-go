# MCPServerCardTool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Title** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**InputSchema** | **map[string]interface{}** |  | 
**Execution** | Pointer to [**NullableToolExecution**](ToolExecution.md) |  | [optional] 
**OutputSchema** | Pointer to **map[string]interface{}** |  | [optional] 
**Icons** | Pointer to [**[]MCPServerIcon**](MCPServerIcon.md) |  | [optional] 
**Annotations** | Pointer to [**NullableToolAnnotations**](ToolAnnotations.md) |  | [optional] 
**Meta** | Pointer to [**NullableMCPToolMeta**](MCPToolMeta.md) |  | [optional] 

## Methods

### NewMCPServerCardTool

`func NewMCPServerCardTool(name string, inputSchema map[string]interface{}, ) *MCPServerCardTool`

NewMCPServerCardTool instantiates a new MCPServerCardTool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMCPServerCardToolWithDefaults

`func NewMCPServerCardToolWithDefaults() *MCPServerCardTool`

NewMCPServerCardToolWithDefaults instantiates a new MCPServerCardTool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MCPServerCardTool) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MCPServerCardTool) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MCPServerCardTool) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *MCPServerCardTool) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MCPServerCardTool) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *MCPServerCardTool) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *MCPServerCardTool) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *MCPServerCardTool) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *MCPServerCardTool) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *MCPServerCardTool) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MCPServerCardTool) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MCPServerCardTool) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MCPServerCardTool) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MCPServerCardTool) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MCPServerCardTool) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetInputSchema

`func (o *MCPServerCardTool) GetInputSchema() map[string]interface{}`

GetInputSchema returns the InputSchema field if non-nil, zero value otherwise.

### GetInputSchemaOk

`func (o *MCPServerCardTool) GetInputSchemaOk() (*map[string]interface{}, bool)`

GetInputSchemaOk returns a tuple with the InputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputSchema

`func (o *MCPServerCardTool) SetInputSchema(v map[string]interface{})`

SetInputSchema sets InputSchema field to given value.


### GetExecution

`func (o *MCPServerCardTool) GetExecution() ToolExecution`

GetExecution returns the Execution field if non-nil, zero value otherwise.

### GetExecutionOk

`func (o *MCPServerCardTool) GetExecutionOk() (*ToolExecution, bool)`

GetExecutionOk returns a tuple with the Execution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecution

`func (o *MCPServerCardTool) SetExecution(v ToolExecution)`

SetExecution sets Execution field to given value.

### HasExecution

`func (o *MCPServerCardTool) HasExecution() bool`

HasExecution returns a boolean if a field has been set.

### SetExecutionNil

`func (o *MCPServerCardTool) SetExecutionNil(b bool)`

 SetExecutionNil sets the value for Execution to be an explicit nil

### UnsetExecution
`func (o *MCPServerCardTool) UnsetExecution()`

UnsetExecution ensures that no value is present for Execution, not even an explicit nil
### GetOutputSchema

`func (o *MCPServerCardTool) GetOutputSchema() map[string]interface{}`

GetOutputSchema returns the OutputSchema field if non-nil, zero value otherwise.

### GetOutputSchemaOk

`func (o *MCPServerCardTool) GetOutputSchemaOk() (*map[string]interface{}, bool)`

GetOutputSchemaOk returns a tuple with the OutputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputSchema

`func (o *MCPServerCardTool) SetOutputSchema(v map[string]interface{})`

SetOutputSchema sets OutputSchema field to given value.

### HasOutputSchema

`func (o *MCPServerCardTool) HasOutputSchema() bool`

HasOutputSchema returns a boolean if a field has been set.

### SetOutputSchemaNil

`func (o *MCPServerCardTool) SetOutputSchemaNil(b bool)`

 SetOutputSchemaNil sets the value for OutputSchema to be an explicit nil

### UnsetOutputSchema
`func (o *MCPServerCardTool) UnsetOutputSchema()`

UnsetOutputSchema ensures that no value is present for OutputSchema, not even an explicit nil
### GetIcons

`func (o *MCPServerCardTool) GetIcons() []MCPServerIcon`

GetIcons returns the Icons field if non-nil, zero value otherwise.

### GetIconsOk

`func (o *MCPServerCardTool) GetIconsOk() (*[]MCPServerIcon, bool)`

GetIconsOk returns a tuple with the Icons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcons

`func (o *MCPServerCardTool) SetIcons(v []MCPServerIcon)`

SetIcons sets Icons field to given value.

### HasIcons

`func (o *MCPServerCardTool) HasIcons() bool`

HasIcons returns a boolean if a field has been set.

### SetIconsNil

`func (o *MCPServerCardTool) SetIconsNil(b bool)`

 SetIconsNil sets the value for Icons to be an explicit nil

### UnsetIcons
`func (o *MCPServerCardTool) UnsetIcons()`

UnsetIcons ensures that no value is present for Icons, not even an explicit nil
### GetAnnotations

`func (o *MCPServerCardTool) GetAnnotations() ToolAnnotations`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *MCPServerCardTool) GetAnnotationsOk() (*ToolAnnotations, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *MCPServerCardTool) SetAnnotations(v ToolAnnotations)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *MCPServerCardTool) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### SetAnnotationsNil

`func (o *MCPServerCardTool) SetAnnotationsNil(b bool)`

 SetAnnotationsNil sets the value for Annotations to be an explicit nil

### UnsetAnnotations
`func (o *MCPServerCardTool) UnsetAnnotations()`

UnsetAnnotations ensures that no value is present for Annotations, not even an explicit nil
### GetMeta

`func (o *MCPServerCardTool) GetMeta() MCPToolMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MCPServerCardTool) GetMetaOk() (*MCPToolMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MCPServerCardTool) SetMeta(v MCPToolMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *MCPServerCardTool) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### SetMetaNil

`func (o *MCPServerCardTool) SetMetaNil(b bool)`

 SetMetaNil sets the value for Meta to be an explicit nil

### UnsetMeta
`func (o *MCPServerCardTool) UnsetMeta()`

UnsetMeta ensures that no value is present for Meta, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


