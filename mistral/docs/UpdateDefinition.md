# UpdateDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the update | 
**Description** | Pointer to **NullableString** | Description of the update | [optional] 
**InputSchema** | **map[string]interface{}** | Input JSON schema of the update&#39;s model | 
**OutputSchema** | Pointer to **map[string]interface{}** | Output JSON schema of the query&#39;s model | [optional] 

## Methods

### NewUpdateDefinition

`func NewUpdateDefinition(name string, inputSchema map[string]interface{}, ) *UpdateDefinition`

NewUpdateDefinition instantiates a new UpdateDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDefinitionWithDefaults

`func NewUpdateDefinitionWithDefaults() *UpdateDefinition`

NewUpdateDefinitionWithDefaults instantiates a new UpdateDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateDefinition) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *UpdateDefinition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateDefinition) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateDefinition) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateDefinition) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateDefinition) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateDefinition) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetInputSchema

`func (o *UpdateDefinition) GetInputSchema() map[string]interface{}`

GetInputSchema returns the InputSchema field if non-nil, zero value otherwise.

### GetInputSchemaOk

`func (o *UpdateDefinition) GetInputSchemaOk() (*map[string]interface{}, bool)`

GetInputSchemaOk returns a tuple with the InputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputSchema

`func (o *UpdateDefinition) SetInputSchema(v map[string]interface{})`

SetInputSchema sets InputSchema field to given value.


### GetOutputSchema

`func (o *UpdateDefinition) GetOutputSchema() map[string]interface{}`

GetOutputSchema returns the OutputSchema field if non-nil, zero value otherwise.

### GetOutputSchemaOk

`func (o *UpdateDefinition) GetOutputSchemaOk() (*map[string]interface{}, bool)`

GetOutputSchemaOk returns a tuple with the OutputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputSchema

`func (o *UpdateDefinition) SetOutputSchema(v map[string]interface{})`

SetOutputSchema sets OutputSchema field to given value.

### HasOutputSchema

`func (o *UpdateDefinition) HasOutputSchema() bool`

HasOutputSchema returns a boolean if a field has been set.

### SetOutputSchemaNil

`func (o *UpdateDefinition) SetOutputSchemaNil(b bool)`

 SetOutputSchemaNil sets the value for OutputSchema to be an explicit nil

### UnsetOutputSchema
`func (o *UpdateDefinition) UnsetOutputSchema()`

UnsetOutputSchema ensures that no value is present for OutputSchema, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


