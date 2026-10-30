# ConnectorTool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Description** | **string** |  | 
**SystemPrompt** | Pointer to **NullableString** |  | [optional] 
**Locale** | Pointer to [**NullableConnectorToolLocale**](ConnectorToolLocale.md) |  | [optional] 
**Jsonschema** | Pointer to **map[string]interface{}** | Output JSON schema of the query&#39;s model | [optional] 
**ExecutionConfig** | [**NullableExecutionConfig**](ExecutionConfig.md) |  | 
**Visibility** | [**ResourceVisibility**](ResourceVisibility.md) |  | 
**CreatedAt** | **time.Time** |  | 
**ModifiedAt** | **time.Time** |  | 
**Active** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewConnectorTool

`func NewConnectorTool(id string, name string, description string, executionConfig NullableExecutionConfig, visibility ResourceVisibility, createdAt time.Time, modifiedAt time.Time, ) *ConnectorTool`

NewConnectorTool instantiates a new ConnectorTool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorToolWithDefaults

`func NewConnectorToolWithDefaults() *ConnectorTool`

NewConnectorToolWithDefaults instantiates a new ConnectorTool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConnectorTool) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConnectorTool) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConnectorTool) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ConnectorTool) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnectorTool) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnectorTool) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ConnectorTool) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConnectorTool) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConnectorTool) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetSystemPrompt

`func (o *ConnectorTool) GetSystemPrompt() string`

GetSystemPrompt returns the SystemPrompt field if non-nil, zero value otherwise.

### GetSystemPromptOk

`func (o *ConnectorTool) GetSystemPromptOk() (*string, bool)`

GetSystemPromptOk returns a tuple with the SystemPrompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemPrompt

`func (o *ConnectorTool) SetSystemPrompt(v string)`

SetSystemPrompt sets SystemPrompt field to given value.

### HasSystemPrompt

`func (o *ConnectorTool) HasSystemPrompt() bool`

HasSystemPrompt returns a boolean if a field has been set.

### SetSystemPromptNil

`func (o *ConnectorTool) SetSystemPromptNil(b bool)`

 SetSystemPromptNil sets the value for SystemPrompt to be an explicit nil

### UnsetSystemPrompt
`func (o *ConnectorTool) UnsetSystemPrompt()`

UnsetSystemPrompt ensures that no value is present for SystemPrompt, not even an explicit nil
### GetLocale

`func (o *ConnectorTool) GetLocale() ConnectorToolLocale`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *ConnectorTool) GetLocaleOk() (*ConnectorToolLocale, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *ConnectorTool) SetLocale(v ConnectorToolLocale)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *ConnectorTool) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### SetLocaleNil

`func (o *ConnectorTool) SetLocaleNil(b bool)`

 SetLocaleNil sets the value for Locale to be an explicit nil

### UnsetLocale
`func (o *ConnectorTool) UnsetLocale()`

UnsetLocale ensures that no value is present for Locale, not even an explicit nil
### GetJsonschema

`func (o *ConnectorTool) GetJsonschema() map[string]interface{}`

GetJsonschema returns the Jsonschema field if non-nil, zero value otherwise.

### GetJsonschemaOk

`func (o *ConnectorTool) GetJsonschemaOk() (*map[string]interface{}, bool)`

GetJsonschemaOk returns a tuple with the Jsonschema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonschema

`func (o *ConnectorTool) SetJsonschema(v map[string]interface{})`

SetJsonschema sets Jsonschema field to given value.

### HasJsonschema

`func (o *ConnectorTool) HasJsonschema() bool`

HasJsonschema returns a boolean if a field has been set.

### SetJsonschemaNil

`func (o *ConnectorTool) SetJsonschemaNil(b bool)`

 SetJsonschemaNil sets the value for Jsonschema to be an explicit nil

### UnsetJsonschema
`func (o *ConnectorTool) UnsetJsonschema()`

UnsetJsonschema ensures that no value is present for Jsonschema, not even an explicit nil
### GetExecutionConfig

`func (o *ConnectorTool) GetExecutionConfig() ExecutionConfig`

GetExecutionConfig returns the ExecutionConfig field if non-nil, zero value otherwise.

### GetExecutionConfigOk

`func (o *ConnectorTool) GetExecutionConfigOk() (*ExecutionConfig, bool)`

GetExecutionConfigOk returns a tuple with the ExecutionConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionConfig

`func (o *ConnectorTool) SetExecutionConfig(v ExecutionConfig)`

SetExecutionConfig sets ExecutionConfig field to given value.


### SetExecutionConfigNil

`func (o *ConnectorTool) SetExecutionConfigNil(b bool)`

 SetExecutionConfigNil sets the value for ExecutionConfig to be an explicit nil

### UnsetExecutionConfig
`func (o *ConnectorTool) UnsetExecutionConfig()`

UnsetExecutionConfig ensures that no value is present for ExecutionConfig, not even an explicit nil
### GetVisibility

`func (o *ConnectorTool) GetVisibility() ResourceVisibility`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ConnectorTool) GetVisibilityOk() (*ResourceVisibility, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ConnectorTool) SetVisibility(v ResourceVisibility)`

SetVisibility sets Visibility field to given value.


### GetCreatedAt

`func (o *ConnectorTool) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ConnectorTool) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ConnectorTool) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetModifiedAt

`func (o *ConnectorTool) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *ConnectorTool) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *ConnectorTool) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetActive

`func (o *ConnectorTool) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ConnectorTool) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ConnectorTool) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ConnectorTool) HasActive() bool`

HasActive returns a boolean if a field has been set.

### SetActiveNil

`func (o *ConnectorTool) SetActiveNil(b bool)`

 SetActiveNil sets the value for Active to be an explicit nil

### UnsetActive
`func (o *ConnectorTool) UnsetActive()`

UnsetActive ensures that no value is present for Active, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


