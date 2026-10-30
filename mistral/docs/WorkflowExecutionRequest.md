# WorkflowExecutionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExecutionId** | Pointer to **NullableString** | Allows you to specify a custom execution ID. If not provided, a random ID will be generated. | [optional] 
**Input** | Pointer to [**NullableInput5**](Input5.md) |  | [optional] 
**WaitForResult** | Pointer to **bool** | If true, wait for the workflow to complete and return the result directly. | [optional] [default to false]
**TimeoutSeconds** | Pointer to **NullableFloat32** | Maximum time to wait for completion when wait_for_result is true. | [optional] 
**CustomTracingAttributes** | Pointer to **map[string]string** |  | [optional] 
**ForceNewTrace** | Pointer to **bool** | If true, ignore the caller&#39;s trace context and start a new, independent trace for this execution instead of joining the caller&#39;s trace. | [optional] [default to false]
**Traceparent** | Pointer to **NullableString** | W3C trace context to join this execution to the caller&#39;s trace. Ignored when force_new_trace is set. | [optional] 
**Extensions** | Pointer to **map[string]interface{}** | Output JSON schema of the query&#39;s model | [optional] 
**TaskQueue** | Pointer to **NullableString** | Deprecated. Use deployment_name instead. | [optional] 
**DeploymentName** | Pointer to **NullableString** | Name of the deployment to route this execution to | [optional] 

## Methods

### NewWorkflowExecutionRequest

`func NewWorkflowExecutionRequest() *WorkflowExecutionRequest`

NewWorkflowExecutionRequest instantiates a new WorkflowExecutionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowExecutionRequestWithDefaults

`func NewWorkflowExecutionRequestWithDefaults() *WorkflowExecutionRequest`

NewWorkflowExecutionRequestWithDefaults instantiates a new WorkflowExecutionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExecutionId

`func (o *WorkflowExecutionRequest) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *WorkflowExecutionRequest) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *WorkflowExecutionRequest) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *WorkflowExecutionRequest) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### SetExecutionIdNil

`func (o *WorkflowExecutionRequest) SetExecutionIdNil(b bool)`

 SetExecutionIdNil sets the value for ExecutionId to be an explicit nil

### UnsetExecutionId
`func (o *WorkflowExecutionRequest) UnsetExecutionId()`

UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
### GetInput

`func (o *WorkflowExecutionRequest) GetInput() Input5`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *WorkflowExecutionRequest) GetInputOk() (*Input5, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *WorkflowExecutionRequest) SetInput(v Input5)`

SetInput sets Input field to given value.

### HasInput

`func (o *WorkflowExecutionRequest) HasInput() bool`

HasInput returns a boolean if a field has been set.

### SetInputNil

`func (o *WorkflowExecutionRequest) SetInputNil(b bool)`

 SetInputNil sets the value for Input to be an explicit nil

### UnsetInput
`func (o *WorkflowExecutionRequest) UnsetInput()`

UnsetInput ensures that no value is present for Input, not even an explicit nil
### GetWaitForResult

`func (o *WorkflowExecutionRequest) GetWaitForResult() bool`

GetWaitForResult returns the WaitForResult field if non-nil, zero value otherwise.

### GetWaitForResultOk

`func (o *WorkflowExecutionRequest) GetWaitForResultOk() (*bool, bool)`

GetWaitForResultOk returns a tuple with the WaitForResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitForResult

`func (o *WorkflowExecutionRequest) SetWaitForResult(v bool)`

SetWaitForResult sets WaitForResult field to given value.

### HasWaitForResult

`func (o *WorkflowExecutionRequest) HasWaitForResult() bool`

HasWaitForResult returns a boolean if a field has been set.

### GetTimeoutSeconds

`func (o *WorkflowExecutionRequest) GetTimeoutSeconds() float32`

GetTimeoutSeconds returns the TimeoutSeconds field if non-nil, zero value otherwise.

### GetTimeoutSecondsOk

`func (o *WorkflowExecutionRequest) GetTimeoutSecondsOk() (*float32, bool)`

GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSeconds

`func (o *WorkflowExecutionRequest) SetTimeoutSeconds(v float32)`

SetTimeoutSeconds sets TimeoutSeconds field to given value.

### HasTimeoutSeconds

`func (o *WorkflowExecutionRequest) HasTimeoutSeconds() bool`

HasTimeoutSeconds returns a boolean if a field has been set.

### SetTimeoutSecondsNil

`func (o *WorkflowExecutionRequest) SetTimeoutSecondsNil(b bool)`

 SetTimeoutSecondsNil sets the value for TimeoutSeconds to be an explicit nil

### UnsetTimeoutSeconds
`func (o *WorkflowExecutionRequest) UnsetTimeoutSeconds()`

UnsetTimeoutSeconds ensures that no value is present for TimeoutSeconds, not even an explicit nil
### GetCustomTracingAttributes

`func (o *WorkflowExecutionRequest) GetCustomTracingAttributes() map[string]string`

GetCustomTracingAttributes returns the CustomTracingAttributes field if non-nil, zero value otherwise.

### GetCustomTracingAttributesOk

`func (o *WorkflowExecutionRequest) GetCustomTracingAttributesOk() (*map[string]string, bool)`

GetCustomTracingAttributesOk returns a tuple with the CustomTracingAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTracingAttributes

`func (o *WorkflowExecutionRequest) SetCustomTracingAttributes(v map[string]string)`

SetCustomTracingAttributes sets CustomTracingAttributes field to given value.

### HasCustomTracingAttributes

`func (o *WorkflowExecutionRequest) HasCustomTracingAttributes() bool`

HasCustomTracingAttributes returns a boolean if a field has been set.

### SetCustomTracingAttributesNil

`func (o *WorkflowExecutionRequest) SetCustomTracingAttributesNil(b bool)`

 SetCustomTracingAttributesNil sets the value for CustomTracingAttributes to be an explicit nil

### UnsetCustomTracingAttributes
`func (o *WorkflowExecutionRequest) UnsetCustomTracingAttributes()`

UnsetCustomTracingAttributes ensures that no value is present for CustomTracingAttributes, not even an explicit nil
### GetForceNewTrace

`func (o *WorkflowExecutionRequest) GetForceNewTrace() bool`

GetForceNewTrace returns the ForceNewTrace field if non-nil, zero value otherwise.

### GetForceNewTraceOk

`func (o *WorkflowExecutionRequest) GetForceNewTraceOk() (*bool, bool)`

GetForceNewTraceOk returns a tuple with the ForceNewTrace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceNewTrace

`func (o *WorkflowExecutionRequest) SetForceNewTrace(v bool)`

SetForceNewTrace sets ForceNewTrace field to given value.

### HasForceNewTrace

`func (o *WorkflowExecutionRequest) HasForceNewTrace() bool`

HasForceNewTrace returns a boolean if a field has been set.

### GetTraceparent

`func (o *WorkflowExecutionRequest) GetTraceparent() string`

GetTraceparent returns the Traceparent field if non-nil, zero value otherwise.

### GetTraceparentOk

`func (o *WorkflowExecutionRequest) GetTraceparentOk() (*string, bool)`

GetTraceparentOk returns a tuple with the Traceparent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceparent

`func (o *WorkflowExecutionRequest) SetTraceparent(v string)`

SetTraceparent sets Traceparent field to given value.

### HasTraceparent

`func (o *WorkflowExecutionRequest) HasTraceparent() bool`

HasTraceparent returns a boolean if a field has been set.

### SetTraceparentNil

`func (o *WorkflowExecutionRequest) SetTraceparentNil(b bool)`

 SetTraceparentNil sets the value for Traceparent to be an explicit nil

### UnsetTraceparent
`func (o *WorkflowExecutionRequest) UnsetTraceparent()`

UnsetTraceparent ensures that no value is present for Traceparent, not even an explicit nil
### GetExtensions

`func (o *WorkflowExecutionRequest) GetExtensions() map[string]interface{}`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *WorkflowExecutionRequest) GetExtensionsOk() (*map[string]interface{}, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *WorkflowExecutionRequest) SetExtensions(v map[string]interface{})`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *WorkflowExecutionRequest) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### SetExtensionsNil

`func (o *WorkflowExecutionRequest) SetExtensionsNil(b bool)`

 SetExtensionsNil sets the value for Extensions to be an explicit nil

### UnsetExtensions
`func (o *WorkflowExecutionRequest) UnsetExtensions()`

UnsetExtensions ensures that no value is present for Extensions, not even an explicit nil
### GetTaskQueue

`func (o *WorkflowExecutionRequest) GetTaskQueue() string`

GetTaskQueue returns the TaskQueue field if non-nil, zero value otherwise.

### GetTaskQueueOk

`func (o *WorkflowExecutionRequest) GetTaskQueueOk() (*string, bool)`

GetTaskQueueOk returns a tuple with the TaskQueue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskQueue

`func (o *WorkflowExecutionRequest) SetTaskQueue(v string)`

SetTaskQueue sets TaskQueue field to given value.

### HasTaskQueue

`func (o *WorkflowExecutionRequest) HasTaskQueue() bool`

HasTaskQueue returns a boolean if a field has been set.

### SetTaskQueueNil

`func (o *WorkflowExecutionRequest) SetTaskQueueNil(b bool)`

 SetTaskQueueNil sets the value for TaskQueue to be an explicit nil

### UnsetTaskQueue
`func (o *WorkflowExecutionRequest) UnsetTaskQueue()`

UnsetTaskQueue ensures that no value is present for TaskQueue, not even an explicit nil
### GetDeploymentName

`func (o *WorkflowExecutionRequest) GetDeploymentName() string`

GetDeploymentName returns the DeploymentName field if non-nil, zero value otherwise.

### GetDeploymentNameOk

`func (o *WorkflowExecutionRequest) GetDeploymentNameOk() (*string, bool)`

GetDeploymentNameOk returns a tuple with the DeploymentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentName

`func (o *WorkflowExecutionRequest) SetDeploymentName(v string)`

SetDeploymentName sets DeploymentName field to given value.

### HasDeploymentName

`func (o *WorkflowExecutionRequest) HasDeploymentName() bool`

HasDeploymentName returns a boolean if a field has been set.

### SetDeploymentNameNil

`func (o *WorkflowExecutionRequest) SetDeploymentNameNil(b bool)`

 SetDeploymentNameNil sets the value for DeploymentName to be an explicit nil

### UnsetDeploymentName
`func (o *WorkflowExecutionRequest) UnsetDeploymentName()`

UnsetDeploymentName ensures that no value is present for DeploymentName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


