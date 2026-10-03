# WorkflowTaskTimedOutResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | **string** | Unique identifier for this event instance. | 
**EventTimestamp** | **int32** | Unix timestamp in nanoseconds when the event was created. | 
**RootWorkflowExecId** | **string** | Execution ID of the root workflow that initiated this execution chain. | 
**ParentWorkflowExecId** | **NullableString** | Execution ID of the parent workflow that initiated this execution. If this is a root workflow, this field is not set. | 
**WorkflowExecId** | **string** | Execution ID of the workflow that emitted this event. | 
**WorkflowRunId** | **string** | Run ID of the workflow execution. Changes on continue-as-new while workflow_exec_id stays the same. | 
**WorkflowName** | **string** | The registered name of the workflow that emitted this event. | 
**EventType** | **string** | Event type discriminator. | [default to "WORKFLOW_TASK_TIMED_OUT"]
**Attributes** | [**WorkflowTaskTimedOutAttributes**](WorkflowTaskTimedOutAttributes.md) | Event-specific attributes. | 

## Methods

### NewWorkflowTaskTimedOutResponse

`func NewWorkflowTaskTimedOutResponse(eventId string, eventTimestamp int32, rootWorkflowExecId string, parentWorkflowExecId NullableString, workflowExecId string, workflowRunId string, workflowName string, eventType string, attributes WorkflowTaskTimedOutAttributes, ) *WorkflowTaskTimedOutResponse`

NewWorkflowTaskTimedOutResponse instantiates a new WorkflowTaskTimedOutResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTaskTimedOutResponseWithDefaults

`func NewWorkflowTaskTimedOutResponseWithDefaults() *WorkflowTaskTimedOutResponse`

NewWorkflowTaskTimedOutResponseWithDefaults instantiates a new WorkflowTaskTimedOutResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventId

`func (o *WorkflowTaskTimedOutResponse) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *WorkflowTaskTimedOutResponse) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *WorkflowTaskTimedOutResponse) SetEventId(v string)`

SetEventId sets EventId field to given value.


### GetEventTimestamp

`func (o *WorkflowTaskTimedOutResponse) GetEventTimestamp() int32`

GetEventTimestamp returns the EventTimestamp field if non-nil, zero value otherwise.

### GetEventTimestampOk

`func (o *WorkflowTaskTimedOutResponse) GetEventTimestampOk() (*int32, bool)`

GetEventTimestampOk returns a tuple with the EventTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTimestamp

`func (o *WorkflowTaskTimedOutResponse) SetEventTimestamp(v int32)`

SetEventTimestamp sets EventTimestamp field to given value.


### GetRootWorkflowExecId

`func (o *WorkflowTaskTimedOutResponse) GetRootWorkflowExecId() string`

GetRootWorkflowExecId returns the RootWorkflowExecId field if non-nil, zero value otherwise.

### GetRootWorkflowExecIdOk

`func (o *WorkflowTaskTimedOutResponse) GetRootWorkflowExecIdOk() (*string, bool)`

GetRootWorkflowExecIdOk returns a tuple with the RootWorkflowExecId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootWorkflowExecId

`func (o *WorkflowTaskTimedOutResponse) SetRootWorkflowExecId(v string)`

SetRootWorkflowExecId sets RootWorkflowExecId field to given value.


### GetParentWorkflowExecId

`func (o *WorkflowTaskTimedOutResponse) GetParentWorkflowExecId() string`

GetParentWorkflowExecId returns the ParentWorkflowExecId field if non-nil, zero value otherwise.

### GetParentWorkflowExecIdOk

`func (o *WorkflowTaskTimedOutResponse) GetParentWorkflowExecIdOk() (*string, bool)`

GetParentWorkflowExecIdOk returns a tuple with the ParentWorkflowExecId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentWorkflowExecId

`func (o *WorkflowTaskTimedOutResponse) SetParentWorkflowExecId(v string)`

SetParentWorkflowExecId sets ParentWorkflowExecId field to given value.


### SetParentWorkflowExecIdNil

`func (o *WorkflowTaskTimedOutResponse) SetParentWorkflowExecIdNil(b bool)`

 SetParentWorkflowExecIdNil sets the value for ParentWorkflowExecId to be an explicit nil

### UnsetParentWorkflowExecId
`func (o *WorkflowTaskTimedOutResponse) UnsetParentWorkflowExecId()`

UnsetParentWorkflowExecId ensures that no value is present for ParentWorkflowExecId, not even an explicit nil
### GetWorkflowExecId

`func (o *WorkflowTaskTimedOutResponse) GetWorkflowExecId() string`

GetWorkflowExecId returns the WorkflowExecId field if non-nil, zero value otherwise.

### GetWorkflowExecIdOk

`func (o *WorkflowTaskTimedOutResponse) GetWorkflowExecIdOk() (*string, bool)`

GetWorkflowExecIdOk returns a tuple with the WorkflowExecId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowExecId

`func (o *WorkflowTaskTimedOutResponse) SetWorkflowExecId(v string)`

SetWorkflowExecId sets WorkflowExecId field to given value.


### GetWorkflowRunId

`func (o *WorkflowTaskTimedOutResponse) GetWorkflowRunId() string`

GetWorkflowRunId returns the WorkflowRunId field if non-nil, zero value otherwise.

### GetWorkflowRunIdOk

`func (o *WorkflowTaskTimedOutResponse) GetWorkflowRunIdOk() (*string, bool)`

GetWorkflowRunIdOk returns a tuple with the WorkflowRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowRunId

`func (o *WorkflowTaskTimedOutResponse) SetWorkflowRunId(v string)`

SetWorkflowRunId sets WorkflowRunId field to given value.


### GetWorkflowName

`func (o *WorkflowTaskTimedOutResponse) GetWorkflowName() string`

GetWorkflowName returns the WorkflowName field if non-nil, zero value otherwise.

### GetWorkflowNameOk

`func (o *WorkflowTaskTimedOutResponse) GetWorkflowNameOk() (*string, bool)`

GetWorkflowNameOk returns a tuple with the WorkflowName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowName

`func (o *WorkflowTaskTimedOutResponse) SetWorkflowName(v string)`

SetWorkflowName sets WorkflowName field to given value.


### GetEventType

`func (o *WorkflowTaskTimedOutResponse) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *WorkflowTaskTimedOutResponse) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *WorkflowTaskTimedOutResponse) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetAttributes

`func (o *WorkflowTaskTimedOutResponse) GetAttributes() WorkflowTaskTimedOutAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *WorkflowTaskTimedOutResponse) GetAttributesOk() (*WorkflowTaskTimedOutAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *WorkflowTaskTimedOutResponse) SetAttributes(v WorkflowTaskTimedOutAttributes)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


