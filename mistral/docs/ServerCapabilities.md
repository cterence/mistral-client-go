# ServerCapabilities

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Experimental** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Logging** | Pointer to **map[string]interface{}** | Capability for logging operations. | [optional] 
**Prompts** | Pointer to [**NullablePromptsCapability**](PromptsCapability.md) |  | [optional] 
**Resources** | Pointer to [**NullableResourcesCapability**](ResourcesCapability.md) |  | [optional] 
**Tools** | Pointer to [**NullableToolsCapability**](ToolsCapability.md) |  | [optional] 
**Completions** | Pointer to **map[string]interface{}** | Capability for completions operations. | [optional] 
**Extensions** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Tasks** | Pointer to [**NullableServerTasksCapability**](ServerTasksCapability.md) |  | [optional] 

## Methods

### NewServerCapabilities

`func NewServerCapabilities() *ServerCapabilities`

NewServerCapabilities instantiates a new ServerCapabilities object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerCapabilitiesWithDefaults

`func NewServerCapabilitiesWithDefaults() *ServerCapabilities`

NewServerCapabilitiesWithDefaults instantiates a new ServerCapabilities object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExperimental

`func (o *ServerCapabilities) GetExperimental() map[string]map[string]interface{}`

GetExperimental returns the Experimental field if non-nil, zero value otherwise.

### GetExperimentalOk

`func (o *ServerCapabilities) GetExperimentalOk() (*map[string]map[string]interface{}, bool)`

GetExperimentalOk returns a tuple with the Experimental field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimental

`func (o *ServerCapabilities) SetExperimental(v map[string]map[string]interface{})`

SetExperimental sets Experimental field to given value.

### HasExperimental

`func (o *ServerCapabilities) HasExperimental() bool`

HasExperimental returns a boolean if a field has been set.

### SetExperimentalNil

`func (o *ServerCapabilities) SetExperimentalNil(b bool)`

 SetExperimentalNil sets the value for Experimental to be an explicit nil

### UnsetExperimental
`func (o *ServerCapabilities) UnsetExperimental()`

UnsetExperimental ensures that no value is present for Experimental, not even an explicit nil
### GetLogging

`func (o *ServerCapabilities) GetLogging() map[string]interface{}`

GetLogging returns the Logging field if non-nil, zero value otherwise.

### GetLoggingOk

`func (o *ServerCapabilities) GetLoggingOk() (*map[string]interface{}, bool)`

GetLoggingOk returns a tuple with the Logging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogging

`func (o *ServerCapabilities) SetLogging(v map[string]interface{})`

SetLogging sets Logging field to given value.

### HasLogging

`func (o *ServerCapabilities) HasLogging() bool`

HasLogging returns a boolean if a field has been set.

### SetLoggingNil

`func (o *ServerCapabilities) SetLoggingNil(b bool)`

 SetLoggingNil sets the value for Logging to be an explicit nil

### UnsetLogging
`func (o *ServerCapabilities) UnsetLogging()`

UnsetLogging ensures that no value is present for Logging, not even an explicit nil
### GetPrompts

`func (o *ServerCapabilities) GetPrompts() PromptsCapability`

GetPrompts returns the Prompts field if non-nil, zero value otherwise.

### GetPromptsOk

`func (o *ServerCapabilities) GetPromptsOk() (*PromptsCapability, bool)`

GetPromptsOk returns a tuple with the Prompts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompts

`func (o *ServerCapabilities) SetPrompts(v PromptsCapability)`

SetPrompts sets Prompts field to given value.

### HasPrompts

`func (o *ServerCapabilities) HasPrompts() bool`

HasPrompts returns a boolean if a field has been set.

### SetPromptsNil

`func (o *ServerCapabilities) SetPromptsNil(b bool)`

 SetPromptsNil sets the value for Prompts to be an explicit nil

### UnsetPrompts
`func (o *ServerCapabilities) UnsetPrompts()`

UnsetPrompts ensures that no value is present for Prompts, not even an explicit nil
### GetResources

`func (o *ServerCapabilities) GetResources() ResourcesCapability`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *ServerCapabilities) GetResourcesOk() (*ResourcesCapability, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *ServerCapabilities) SetResources(v ResourcesCapability)`

SetResources sets Resources field to given value.

### HasResources

`func (o *ServerCapabilities) HasResources() bool`

HasResources returns a boolean if a field has been set.

### SetResourcesNil

`func (o *ServerCapabilities) SetResourcesNil(b bool)`

 SetResourcesNil sets the value for Resources to be an explicit nil

### UnsetResources
`func (o *ServerCapabilities) UnsetResources()`

UnsetResources ensures that no value is present for Resources, not even an explicit nil
### GetTools

`func (o *ServerCapabilities) GetTools() ToolsCapability`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *ServerCapabilities) GetToolsOk() (*ToolsCapability, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *ServerCapabilities) SetTools(v ToolsCapability)`

SetTools sets Tools field to given value.

### HasTools

`func (o *ServerCapabilities) HasTools() bool`

HasTools returns a boolean if a field has been set.

### SetToolsNil

`func (o *ServerCapabilities) SetToolsNil(b bool)`

 SetToolsNil sets the value for Tools to be an explicit nil

### UnsetTools
`func (o *ServerCapabilities) UnsetTools()`

UnsetTools ensures that no value is present for Tools, not even an explicit nil
### GetCompletions

`func (o *ServerCapabilities) GetCompletions() map[string]interface{}`

GetCompletions returns the Completions field if non-nil, zero value otherwise.

### GetCompletionsOk

`func (o *ServerCapabilities) GetCompletionsOk() (*map[string]interface{}, bool)`

GetCompletionsOk returns a tuple with the Completions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletions

`func (o *ServerCapabilities) SetCompletions(v map[string]interface{})`

SetCompletions sets Completions field to given value.

### HasCompletions

`func (o *ServerCapabilities) HasCompletions() bool`

HasCompletions returns a boolean if a field has been set.

### SetCompletionsNil

`func (o *ServerCapabilities) SetCompletionsNil(b bool)`

 SetCompletionsNil sets the value for Completions to be an explicit nil

### UnsetCompletions
`func (o *ServerCapabilities) UnsetCompletions()`

UnsetCompletions ensures that no value is present for Completions, not even an explicit nil
### GetExtensions

`func (o *ServerCapabilities) GetExtensions() map[string]map[string]interface{}`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *ServerCapabilities) GetExtensionsOk() (*map[string]map[string]interface{}, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *ServerCapabilities) SetExtensions(v map[string]map[string]interface{})`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *ServerCapabilities) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### SetExtensionsNil

`func (o *ServerCapabilities) SetExtensionsNil(b bool)`

 SetExtensionsNil sets the value for Extensions to be an explicit nil

### UnsetExtensions
`func (o *ServerCapabilities) UnsetExtensions()`

UnsetExtensions ensures that no value is present for Extensions, not even an explicit nil
### GetTasks

`func (o *ServerCapabilities) GetTasks() ServerTasksCapability`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *ServerCapabilities) GetTasksOk() (*ServerTasksCapability, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *ServerCapabilities) SetTasks(v ServerTasksCapability)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *ServerCapabilities) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### SetTasksNil

`func (o *ServerCapabilities) SetTasksNil(b bool)`

 SetTasksNil sets the value for Tasks to be an explicit nil

### UnsetTasks
`func (o *ServerCapabilities) UnsetTasks()`

UnsetTasks ensures that no value is present for Tasks, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


