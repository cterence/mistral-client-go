# ClientCapabilities

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Experimental** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Sampling** | Pointer to [**NullableSamplingCapability**](SamplingCapability.md) |  | [optional] 
**Elicitation** | Pointer to [**NullableElicitationCapability**](ElicitationCapability.md) |  | [optional] 
**Roots** | Pointer to [**NullableRootsCapability**](RootsCapability.md) |  | [optional] 
**Extensions** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Tasks** | Pointer to [**NullableClientTasksCapability**](ClientTasksCapability.md) |  | [optional] 

## Methods

### NewClientCapabilities

`func NewClientCapabilities() *ClientCapabilities`

NewClientCapabilities instantiates a new ClientCapabilities object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientCapabilitiesWithDefaults

`func NewClientCapabilitiesWithDefaults() *ClientCapabilities`

NewClientCapabilitiesWithDefaults instantiates a new ClientCapabilities object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExperimental

`func (o *ClientCapabilities) GetExperimental() map[string]map[string]interface{}`

GetExperimental returns the Experimental field if non-nil, zero value otherwise.

### GetExperimentalOk

`func (o *ClientCapabilities) GetExperimentalOk() (*map[string]map[string]interface{}, bool)`

GetExperimentalOk returns a tuple with the Experimental field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimental

`func (o *ClientCapabilities) SetExperimental(v map[string]map[string]interface{})`

SetExperimental sets Experimental field to given value.

### HasExperimental

`func (o *ClientCapabilities) HasExperimental() bool`

HasExperimental returns a boolean if a field has been set.

### SetExperimentalNil

`func (o *ClientCapabilities) SetExperimentalNil(b bool)`

 SetExperimentalNil sets the value for Experimental to be an explicit nil

### UnsetExperimental
`func (o *ClientCapabilities) UnsetExperimental()`

UnsetExperimental ensures that no value is present for Experimental, not even an explicit nil
### GetSampling

`func (o *ClientCapabilities) GetSampling() SamplingCapability`

GetSampling returns the Sampling field if non-nil, zero value otherwise.

### GetSamplingOk

`func (o *ClientCapabilities) GetSamplingOk() (*SamplingCapability, bool)`

GetSamplingOk returns a tuple with the Sampling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampling

`func (o *ClientCapabilities) SetSampling(v SamplingCapability)`

SetSampling sets Sampling field to given value.

### HasSampling

`func (o *ClientCapabilities) HasSampling() bool`

HasSampling returns a boolean if a field has been set.

### SetSamplingNil

`func (o *ClientCapabilities) SetSamplingNil(b bool)`

 SetSamplingNil sets the value for Sampling to be an explicit nil

### UnsetSampling
`func (o *ClientCapabilities) UnsetSampling()`

UnsetSampling ensures that no value is present for Sampling, not even an explicit nil
### GetElicitation

`func (o *ClientCapabilities) GetElicitation() ElicitationCapability`

GetElicitation returns the Elicitation field if non-nil, zero value otherwise.

### GetElicitationOk

`func (o *ClientCapabilities) GetElicitationOk() (*ElicitationCapability, bool)`

GetElicitationOk returns a tuple with the Elicitation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElicitation

`func (o *ClientCapabilities) SetElicitation(v ElicitationCapability)`

SetElicitation sets Elicitation field to given value.

### HasElicitation

`func (o *ClientCapabilities) HasElicitation() bool`

HasElicitation returns a boolean if a field has been set.

### SetElicitationNil

`func (o *ClientCapabilities) SetElicitationNil(b bool)`

 SetElicitationNil sets the value for Elicitation to be an explicit nil

### UnsetElicitation
`func (o *ClientCapabilities) UnsetElicitation()`

UnsetElicitation ensures that no value is present for Elicitation, not even an explicit nil
### GetRoots

`func (o *ClientCapabilities) GetRoots() RootsCapability`

GetRoots returns the Roots field if non-nil, zero value otherwise.

### GetRootsOk

`func (o *ClientCapabilities) GetRootsOk() (*RootsCapability, bool)`

GetRootsOk returns a tuple with the Roots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoots

`func (o *ClientCapabilities) SetRoots(v RootsCapability)`

SetRoots sets Roots field to given value.

### HasRoots

`func (o *ClientCapabilities) HasRoots() bool`

HasRoots returns a boolean if a field has been set.

### SetRootsNil

`func (o *ClientCapabilities) SetRootsNil(b bool)`

 SetRootsNil sets the value for Roots to be an explicit nil

### UnsetRoots
`func (o *ClientCapabilities) UnsetRoots()`

UnsetRoots ensures that no value is present for Roots, not even an explicit nil
### GetExtensions

`func (o *ClientCapabilities) GetExtensions() map[string]map[string]interface{}`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *ClientCapabilities) GetExtensionsOk() (*map[string]map[string]interface{}, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *ClientCapabilities) SetExtensions(v map[string]map[string]interface{})`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *ClientCapabilities) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### SetExtensionsNil

`func (o *ClientCapabilities) SetExtensionsNil(b bool)`

 SetExtensionsNil sets the value for Extensions to be an explicit nil

### UnsetExtensions
`func (o *ClientCapabilities) UnsetExtensions()`

UnsetExtensions ensures that no value is present for Extensions, not even an explicit nil
### GetTasks

`func (o *ClientCapabilities) GetTasks() ClientTasksCapability`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *ClientCapabilities) GetTasksOk() (*ClientTasksCapability, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *ClientCapabilities) SetTasks(v ClientTasksCapability)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *ClientCapabilities) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### SetTasksNil

`func (o *ClientCapabilities) SetTasksNil(b bool)`

 SetTasksNil sets the value for Tasks to be an explicit nil

### UnsetTasks
`func (o *ClientCapabilities) UnsetTasks()`

UnsetTasks ensures that no value is present for Tasks, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


