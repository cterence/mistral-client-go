# WorkflowExecutionStartedAttributesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskId** | **string** | Unique identifier for the task within the workflow execution. | 
**WorkflowName** | **string** | The registered name of the workflow being executed. | 
**DisplayName** | Pointer to **NullableString** | The user-friendly display name of the workflow, if available. | [optional] 
**Input** | [**JSONPayloadResponse**](JSONPayloadResponse.md) | The input arguments passed to the workflow. | 
**Attempt** | Pointer to **int32** | Workflow retry attempt number. 1 on first run and CAN; &gt;1 on workflow-level retries. | [optional] [default to 1]

## Methods

### NewWorkflowExecutionStartedAttributesResponse

`func NewWorkflowExecutionStartedAttributesResponse(taskId string, workflowName string, input JSONPayloadResponse, ) *WorkflowExecutionStartedAttributesResponse`

NewWorkflowExecutionStartedAttributesResponse instantiates a new WorkflowExecutionStartedAttributesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowExecutionStartedAttributesResponseWithDefaults

`func NewWorkflowExecutionStartedAttributesResponseWithDefaults() *WorkflowExecutionStartedAttributesResponse`

NewWorkflowExecutionStartedAttributesResponseWithDefaults instantiates a new WorkflowExecutionStartedAttributesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskId

`func (o *WorkflowExecutionStartedAttributesResponse) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *WorkflowExecutionStartedAttributesResponse) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *WorkflowExecutionStartedAttributesResponse) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.


### GetWorkflowName

`func (o *WorkflowExecutionStartedAttributesResponse) GetWorkflowName() string`

GetWorkflowName returns the WorkflowName field if non-nil, zero value otherwise.

### GetWorkflowNameOk

`func (o *WorkflowExecutionStartedAttributesResponse) GetWorkflowNameOk() (*string, bool)`

GetWorkflowNameOk returns a tuple with the WorkflowName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowName

`func (o *WorkflowExecutionStartedAttributesResponse) SetWorkflowName(v string)`

SetWorkflowName sets WorkflowName field to given value.


### GetDisplayName

`func (o *WorkflowExecutionStartedAttributesResponse) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *WorkflowExecutionStartedAttributesResponse) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *WorkflowExecutionStartedAttributesResponse) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *WorkflowExecutionStartedAttributesResponse) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *WorkflowExecutionStartedAttributesResponse) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *WorkflowExecutionStartedAttributesResponse) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetInput

`func (o *WorkflowExecutionStartedAttributesResponse) GetInput() JSONPayloadResponse`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *WorkflowExecutionStartedAttributesResponse) GetInputOk() (*JSONPayloadResponse, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *WorkflowExecutionStartedAttributesResponse) SetInput(v JSONPayloadResponse)`

SetInput sets Input field to given value.


### GetAttempt

`func (o *WorkflowExecutionStartedAttributesResponse) GetAttempt() int32`

GetAttempt returns the Attempt field if non-nil, zero value otherwise.

### GetAttemptOk

`func (o *WorkflowExecutionStartedAttributesResponse) GetAttemptOk() (*int32, bool)`

GetAttemptOk returns a tuple with the Attempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttempt

`func (o *WorkflowExecutionStartedAttributesResponse) SetAttempt(v int32)`

SetAttempt sets Attempt field to given value.

### HasAttempt

`func (o *WorkflowExecutionStartedAttributesResponse) HasAttempt() bool`

HasAttempt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


