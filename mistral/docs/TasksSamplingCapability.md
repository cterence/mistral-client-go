# TasksSamplingCapability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateMessage** | Pointer to **map[string]interface{}** | Capability for task-augmented sampling/createMessage requests (2025-11-25 only). | [optional] 

## Methods

### NewTasksSamplingCapability

`func NewTasksSamplingCapability() *TasksSamplingCapability`

NewTasksSamplingCapability instantiates a new TasksSamplingCapability object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTasksSamplingCapabilityWithDefaults

`func NewTasksSamplingCapabilityWithDefaults() *TasksSamplingCapability`

NewTasksSamplingCapabilityWithDefaults instantiates a new TasksSamplingCapability object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateMessage

`func (o *TasksSamplingCapability) GetCreateMessage() map[string]interface{}`

GetCreateMessage returns the CreateMessage field if non-nil, zero value otherwise.

### GetCreateMessageOk

`func (o *TasksSamplingCapability) GetCreateMessageOk() (*map[string]interface{}, bool)`

GetCreateMessageOk returns a tuple with the CreateMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateMessage

`func (o *TasksSamplingCapability) SetCreateMessage(v map[string]interface{})`

SetCreateMessage sets CreateMessage field to given value.

### HasCreateMessage

`func (o *TasksSamplingCapability) HasCreateMessage() bool`

HasCreateMessage returns a boolean if a field has been set.

### SetCreateMessageNil

`func (o *TasksSamplingCapability) SetCreateMessageNil(b bool)`

 SetCreateMessageNil sets the value for CreateMessage to be an explicit nil

### UnsetCreateMessage
`func (o *TasksSamplingCapability) UnsetCreateMessage()`

UnsetCreateMessage ensures that no value is present for CreateMessage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


