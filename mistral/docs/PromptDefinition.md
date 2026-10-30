# PromptDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | **string** | Prompt template content. | 
**Variables** | Pointer to [**[]PromptVariable**](PromptVariable.md) | Variables used by the prompt. | [optional] 

## Methods

### NewPromptDefinition

`func NewPromptDefinition(content string, ) *PromptDefinition`

NewPromptDefinition instantiates a new PromptDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPromptDefinitionWithDefaults

`func NewPromptDefinitionWithDefaults() *PromptDefinition`

NewPromptDefinitionWithDefaults instantiates a new PromptDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *PromptDefinition) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *PromptDefinition) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *PromptDefinition) SetContent(v string)`

SetContent sets Content field to given value.


### GetVariables

`func (o *PromptDefinition) GetVariables() []PromptVariable`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *PromptDefinition) GetVariablesOk() (*[]PromptVariable, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *PromptDefinition) SetVariables(v []PromptVariable)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *PromptDefinition) HasVariables() bool`

HasVariables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


