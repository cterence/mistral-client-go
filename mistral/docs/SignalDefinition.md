# SignalDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the signal | 
**Description** | Pointer to **NullableString** | Description of the signal | [optional] 
**InputSchema** | **map[string]interface{}** | Input JSON schema of the signal&#39;s model | 

## Methods

### NewSignalDefinition

`func NewSignalDefinition(name string, inputSchema map[string]interface{}, ) *SignalDefinition`

NewSignalDefinition instantiates a new SignalDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignalDefinitionWithDefaults

`func NewSignalDefinitionWithDefaults() *SignalDefinition`

NewSignalDefinitionWithDefaults instantiates a new SignalDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SignalDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SignalDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SignalDefinition) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *SignalDefinition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SignalDefinition) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SignalDefinition) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SignalDefinition) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *SignalDefinition) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SignalDefinition) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetInputSchema

`func (o *SignalDefinition) GetInputSchema() map[string]interface{}`

GetInputSchema returns the InputSchema field if non-nil, zero value otherwise.

### GetInputSchemaOk

`func (o *SignalDefinition) GetInputSchemaOk() (*map[string]interface{}, bool)`

GetInputSchemaOk returns a tuple with the InputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputSchema

`func (o *SignalDefinition) SetInputSchema(v map[string]interface{})`

SetInputSchema sets InputSchema field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


