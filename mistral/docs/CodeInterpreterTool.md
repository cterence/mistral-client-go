# CodeInterpreterTool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ToolConfiguration** | Pointer to [**NullableToolConfiguration**](ToolConfiguration.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] [default to "code_interpreter"]

## Methods

### NewCodeInterpreterTool

`func NewCodeInterpreterTool() *CodeInterpreterTool`

NewCodeInterpreterTool instantiates a new CodeInterpreterTool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodeInterpreterToolWithDefaults

`func NewCodeInterpreterToolWithDefaults() *CodeInterpreterTool`

NewCodeInterpreterToolWithDefaults instantiates a new CodeInterpreterTool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToolConfiguration

`func (o *CodeInterpreterTool) GetToolConfiguration() ToolConfiguration`

GetToolConfiguration returns the ToolConfiguration field if non-nil, zero value otherwise.

### GetToolConfigurationOk

`func (o *CodeInterpreterTool) GetToolConfigurationOk() (*ToolConfiguration, bool)`

GetToolConfigurationOk returns a tuple with the ToolConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolConfiguration

`func (o *CodeInterpreterTool) SetToolConfiguration(v ToolConfiguration)`

SetToolConfiguration sets ToolConfiguration field to given value.

### HasToolConfiguration

`func (o *CodeInterpreterTool) HasToolConfiguration() bool`

HasToolConfiguration returns a boolean if a field has been set.

### SetToolConfigurationNil

`func (o *CodeInterpreterTool) SetToolConfigurationNil(b bool)`

 SetToolConfigurationNil sets the value for ToolConfiguration to be an explicit nil

### UnsetToolConfiguration
`func (o *CodeInterpreterTool) UnsetToolConfiguration()`

UnsetToolConfiguration ensures that no value is present for ToolConfiguration, not even an explicit nil
### GetType

`func (o *CodeInterpreterTool) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CodeInterpreterTool) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CodeInterpreterTool) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CodeInterpreterTool) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


