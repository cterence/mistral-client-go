# QueryDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the query | 
**Description** | Pointer to **NullableString** | Description of the query | [optional] 
**InputSchema** | **map[string]interface{}** | Input JSON schema of the query&#39;s model | 
**OutputSchema** | Pointer to **map[string]interface{}** | Output JSON schema of the query&#39;s model | [optional] 

## Methods

### NewQueryDefinition

`func NewQueryDefinition(name string, inputSchema map[string]interface{}, ) *QueryDefinition`

NewQueryDefinition instantiates a new QueryDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryDefinitionWithDefaults

`func NewQueryDefinitionWithDefaults() *QueryDefinition`

NewQueryDefinitionWithDefaults instantiates a new QueryDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *QueryDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QueryDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QueryDefinition) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *QueryDefinition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *QueryDefinition) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *QueryDefinition) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *QueryDefinition) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *QueryDefinition) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *QueryDefinition) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetInputSchema

`func (o *QueryDefinition) GetInputSchema() map[string]interface{}`

GetInputSchema returns the InputSchema field if non-nil, zero value otherwise.

### GetInputSchemaOk

`func (o *QueryDefinition) GetInputSchemaOk() (*map[string]interface{}, bool)`

GetInputSchemaOk returns a tuple with the InputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputSchema

`func (o *QueryDefinition) SetInputSchema(v map[string]interface{})`

SetInputSchema sets InputSchema field to given value.


### GetOutputSchema

`func (o *QueryDefinition) GetOutputSchema() map[string]interface{}`

GetOutputSchema returns the OutputSchema field if non-nil, zero value otherwise.

### GetOutputSchemaOk

`func (o *QueryDefinition) GetOutputSchemaOk() (*map[string]interface{}, bool)`

GetOutputSchemaOk returns a tuple with the OutputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputSchema

`func (o *QueryDefinition) SetOutputSchema(v map[string]interface{})`

SetOutputSchema sets OutputSchema field to given value.

### HasOutputSchema

`func (o *QueryDefinition) HasOutputSchema() bool`

HasOutputSchema returns a boolean if a field has been set.

### SetOutputSchemaNil

`func (o *QueryDefinition) SetOutputSchemaNil(b bool)`

 SetOutputSchemaNil sets the value for OutputSchema to be an explicit nil

### UnsetOutputSchema
`func (o *QueryDefinition) UnsetOutputSchema()`

UnsetOutputSchema ensures that no value is present for OutputSchema, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


