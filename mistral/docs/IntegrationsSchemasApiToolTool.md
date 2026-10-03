# IntegrationsSchemasApiToolTool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Description** | **string** |  | 
**SystemPrompt** | Pointer to **NullableString** |  | [optional] 
**Locale** | Pointer to [**NullableIntegrationsSchemasTurbineToolLocale**](IntegrationsSchemasTurbineToolLocale.md) |  | [optional] 
**Jsonschema** | Pointer to **map[string]interface{}** |  | [optional] 
**ExecutionConfig** | [**NullableExecutionConfig**](ExecutionConfig.md) |  | 
**Visibility** | [**ResourceVisibility**](ResourceVisibility.md) |  | 
**CreatedAt** | **time.Time** |  | 
**ModifiedAt** | **time.Time** |  | 
**Active** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewIntegrationsSchemasApiToolTool

`func NewIntegrationsSchemasApiToolTool(id string, name string, description string, executionConfig NullableExecutionConfig, visibility ResourceVisibility, createdAt time.Time, modifiedAt time.Time, ) *IntegrationsSchemasApiToolTool`

NewIntegrationsSchemasApiToolTool instantiates a new IntegrationsSchemasApiToolTool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationsSchemasApiToolToolWithDefaults

`func NewIntegrationsSchemasApiToolToolWithDefaults() *IntegrationsSchemasApiToolTool`

NewIntegrationsSchemasApiToolToolWithDefaults instantiates a new IntegrationsSchemasApiToolTool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IntegrationsSchemasApiToolTool) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IntegrationsSchemasApiToolTool) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IntegrationsSchemasApiToolTool) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *IntegrationsSchemasApiToolTool) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IntegrationsSchemasApiToolTool) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IntegrationsSchemasApiToolTool) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *IntegrationsSchemasApiToolTool) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IntegrationsSchemasApiToolTool) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IntegrationsSchemasApiToolTool) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetSystemPrompt

`func (o *IntegrationsSchemasApiToolTool) GetSystemPrompt() string`

GetSystemPrompt returns the SystemPrompt field if non-nil, zero value otherwise.

### GetSystemPromptOk

`func (o *IntegrationsSchemasApiToolTool) GetSystemPromptOk() (*string, bool)`

GetSystemPromptOk returns a tuple with the SystemPrompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemPrompt

`func (o *IntegrationsSchemasApiToolTool) SetSystemPrompt(v string)`

SetSystemPrompt sets SystemPrompt field to given value.

### HasSystemPrompt

`func (o *IntegrationsSchemasApiToolTool) HasSystemPrompt() bool`

HasSystemPrompt returns a boolean if a field has been set.

### SetSystemPromptNil

`func (o *IntegrationsSchemasApiToolTool) SetSystemPromptNil(b bool)`

 SetSystemPromptNil sets the value for SystemPrompt to be an explicit nil

### UnsetSystemPrompt
`func (o *IntegrationsSchemasApiToolTool) UnsetSystemPrompt()`

UnsetSystemPrompt ensures that no value is present for SystemPrompt, not even an explicit nil
### GetLocale

`func (o *IntegrationsSchemasApiToolTool) GetLocale() IntegrationsSchemasTurbineToolLocale`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *IntegrationsSchemasApiToolTool) GetLocaleOk() (*IntegrationsSchemasTurbineToolLocale, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *IntegrationsSchemasApiToolTool) SetLocale(v IntegrationsSchemasTurbineToolLocale)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *IntegrationsSchemasApiToolTool) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### SetLocaleNil

`func (o *IntegrationsSchemasApiToolTool) SetLocaleNil(b bool)`

 SetLocaleNil sets the value for Locale to be an explicit nil

### UnsetLocale
`func (o *IntegrationsSchemasApiToolTool) UnsetLocale()`

UnsetLocale ensures that no value is present for Locale, not even an explicit nil
### GetJsonschema

`func (o *IntegrationsSchemasApiToolTool) GetJsonschema() map[string]interface{}`

GetJsonschema returns the Jsonschema field if non-nil, zero value otherwise.

### GetJsonschemaOk

`func (o *IntegrationsSchemasApiToolTool) GetJsonschemaOk() (*map[string]interface{}, bool)`

GetJsonschemaOk returns a tuple with the Jsonschema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonschema

`func (o *IntegrationsSchemasApiToolTool) SetJsonschema(v map[string]interface{})`

SetJsonschema sets Jsonschema field to given value.

### HasJsonschema

`func (o *IntegrationsSchemasApiToolTool) HasJsonschema() bool`

HasJsonschema returns a boolean if a field has been set.

### SetJsonschemaNil

`func (o *IntegrationsSchemasApiToolTool) SetJsonschemaNil(b bool)`

 SetJsonschemaNil sets the value for Jsonschema to be an explicit nil

### UnsetJsonschema
`func (o *IntegrationsSchemasApiToolTool) UnsetJsonschema()`

UnsetJsonschema ensures that no value is present for Jsonschema, not even an explicit nil
### GetExecutionConfig

`func (o *IntegrationsSchemasApiToolTool) GetExecutionConfig() ExecutionConfig`

GetExecutionConfig returns the ExecutionConfig field if non-nil, zero value otherwise.

### GetExecutionConfigOk

`func (o *IntegrationsSchemasApiToolTool) GetExecutionConfigOk() (*ExecutionConfig, bool)`

GetExecutionConfigOk returns a tuple with the ExecutionConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionConfig

`func (o *IntegrationsSchemasApiToolTool) SetExecutionConfig(v ExecutionConfig)`

SetExecutionConfig sets ExecutionConfig field to given value.


### SetExecutionConfigNil

`func (o *IntegrationsSchemasApiToolTool) SetExecutionConfigNil(b bool)`

 SetExecutionConfigNil sets the value for ExecutionConfig to be an explicit nil

### UnsetExecutionConfig
`func (o *IntegrationsSchemasApiToolTool) UnsetExecutionConfig()`

UnsetExecutionConfig ensures that no value is present for ExecutionConfig, not even an explicit nil
### GetVisibility

`func (o *IntegrationsSchemasApiToolTool) GetVisibility() ResourceVisibility`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *IntegrationsSchemasApiToolTool) GetVisibilityOk() (*ResourceVisibility, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *IntegrationsSchemasApiToolTool) SetVisibility(v ResourceVisibility)`

SetVisibility sets Visibility field to given value.


### GetCreatedAt

`func (o *IntegrationsSchemasApiToolTool) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IntegrationsSchemasApiToolTool) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IntegrationsSchemasApiToolTool) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetModifiedAt

`func (o *IntegrationsSchemasApiToolTool) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *IntegrationsSchemasApiToolTool) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *IntegrationsSchemasApiToolTool) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetActive

`func (o *IntegrationsSchemasApiToolTool) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *IntegrationsSchemasApiToolTool) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *IntegrationsSchemasApiToolTool) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *IntegrationsSchemasApiToolTool) HasActive() bool`

HasActive returns a boolean if a field has been set.

### SetActiveNil

`func (o *IntegrationsSchemasApiToolTool) SetActiveNil(b bool)`

 SetActiveNil sets the value for Active to be an explicit nil

### UnsetActive
`func (o *IntegrationsSchemasApiToolTool) UnsetActive()`

UnsetActive ensures that no value is present for Active, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


