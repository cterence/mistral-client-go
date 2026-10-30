# ActivityTaskStartedResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | **string** | Unique identifier for this event instance. | 
**EventTimestamp** | **int32** | Unix timestamp in nanoseconds when the event was created. | 
**RootWorkflowExecId** | **string** | Execution ID of the root workflow that initiated this execution chain. | 
**ParentWorkflowExecId** | **NullableString** | Execution ID of the parent workflow that initiated this execution. If this is a root workflow, this field is not set. | 
**ContinuedRunId** | Pointer to **NullableString** | Run ID of the execution this run continued from. Non-null for continue-as-new runs. | [optional] 
**FirstExecutionRunId** | Pointer to **NullableString** | Run ID of the first execution in this workflow chain. Equals workflow_run_id on a fresh start. Continue-as-new, retry and reset runs keep the original value. | [optional] 
**ChainRunId** | Pointer to **NullableString** | Identifies one start of an execution. Shared by its continue-as-new, retry and reset runs, and by every sub-workflow below it. Starting the same execution ID again gives a new value. | [optional] 
**ScheduleId** | Pointer to **NullableString** | Temporal schedule ID that triggered this execution, if any. | [optional] 
**WorkflowExecId** | **string** | Execution ID of the workflow that emitted this event. | 
**WorkflowRunId** | **string** | Run ID of the workflow execution. Changes on continue-as-new while workflow_exec_id stays the same. | 
**WorkflowName** | **string** | The registered name of the workflow that emitted this event. | 
**EventType** | **string** | Event type discriminator. | [default to "ACTIVITY_TASK_STARTED"]
**Attributes** | [**ActivityTaskStartedAttributesResponse**](ActivityTaskStartedAttributesResponse.md) | Event-specific attributes. | 

## Methods

### NewActivityTaskStartedResponse

`func NewActivityTaskStartedResponse(eventId string, eventTimestamp int32, rootWorkflowExecId string, parentWorkflowExecId NullableString, workflowExecId string, workflowRunId string, workflowName string, eventType string, attributes ActivityTaskStartedAttributesResponse, ) *ActivityTaskStartedResponse`

NewActivityTaskStartedResponse instantiates a new ActivityTaskStartedResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityTaskStartedResponseWithDefaults

`func NewActivityTaskStartedResponseWithDefaults() *ActivityTaskStartedResponse`

NewActivityTaskStartedResponseWithDefaults instantiates a new ActivityTaskStartedResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventId

`func (o *ActivityTaskStartedResponse) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *ActivityTaskStartedResponse) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *ActivityTaskStartedResponse) SetEventId(v string)`

SetEventId sets EventId field to given value.


### GetEventTimestamp

`func (o *ActivityTaskStartedResponse) GetEventTimestamp() int32`

GetEventTimestamp returns the EventTimestamp field if non-nil, zero value otherwise.

### GetEventTimestampOk

`func (o *ActivityTaskStartedResponse) GetEventTimestampOk() (*int32, bool)`

GetEventTimestampOk returns a tuple with the EventTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTimestamp

`func (o *ActivityTaskStartedResponse) SetEventTimestamp(v int32)`

SetEventTimestamp sets EventTimestamp field to given value.


### GetRootWorkflowExecId

`func (o *ActivityTaskStartedResponse) GetRootWorkflowExecId() string`

GetRootWorkflowExecId returns the RootWorkflowExecId field if non-nil, zero value otherwise.

### GetRootWorkflowExecIdOk

`func (o *ActivityTaskStartedResponse) GetRootWorkflowExecIdOk() (*string, bool)`

GetRootWorkflowExecIdOk returns a tuple with the RootWorkflowExecId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootWorkflowExecId

`func (o *ActivityTaskStartedResponse) SetRootWorkflowExecId(v string)`

SetRootWorkflowExecId sets RootWorkflowExecId field to given value.


### GetParentWorkflowExecId

`func (o *ActivityTaskStartedResponse) GetParentWorkflowExecId() string`

GetParentWorkflowExecId returns the ParentWorkflowExecId field if non-nil, zero value otherwise.

### GetParentWorkflowExecIdOk

`func (o *ActivityTaskStartedResponse) GetParentWorkflowExecIdOk() (*string, bool)`

GetParentWorkflowExecIdOk returns a tuple with the ParentWorkflowExecId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentWorkflowExecId

`func (o *ActivityTaskStartedResponse) SetParentWorkflowExecId(v string)`

SetParentWorkflowExecId sets ParentWorkflowExecId field to given value.


### SetParentWorkflowExecIdNil

`func (o *ActivityTaskStartedResponse) SetParentWorkflowExecIdNil(b bool)`

 SetParentWorkflowExecIdNil sets the value for ParentWorkflowExecId to be an explicit nil

### UnsetParentWorkflowExecId
`func (o *ActivityTaskStartedResponse) UnsetParentWorkflowExecId()`

UnsetParentWorkflowExecId ensures that no value is present for ParentWorkflowExecId, not even an explicit nil
### GetContinuedRunId

`func (o *ActivityTaskStartedResponse) GetContinuedRunId() string`

GetContinuedRunId returns the ContinuedRunId field if non-nil, zero value otherwise.

### GetContinuedRunIdOk

`func (o *ActivityTaskStartedResponse) GetContinuedRunIdOk() (*string, bool)`

GetContinuedRunIdOk returns a tuple with the ContinuedRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinuedRunId

`func (o *ActivityTaskStartedResponse) SetContinuedRunId(v string)`

SetContinuedRunId sets ContinuedRunId field to given value.

### HasContinuedRunId

`func (o *ActivityTaskStartedResponse) HasContinuedRunId() bool`

HasContinuedRunId returns a boolean if a field has been set.

### SetContinuedRunIdNil

`func (o *ActivityTaskStartedResponse) SetContinuedRunIdNil(b bool)`

 SetContinuedRunIdNil sets the value for ContinuedRunId to be an explicit nil

### UnsetContinuedRunId
`func (o *ActivityTaskStartedResponse) UnsetContinuedRunId()`

UnsetContinuedRunId ensures that no value is present for ContinuedRunId, not even an explicit nil
### GetFirstExecutionRunId

`func (o *ActivityTaskStartedResponse) GetFirstExecutionRunId() string`

GetFirstExecutionRunId returns the FirstExecutionRunId field if non-nil, zero value otherwise.

### GetFirstExecutionRunIdOk

`func (o *ActivityTaskStartedResponse) GetFirstExecutionRunIdOk() (*string, bool)`

GetFirstExecutionRunIdOk returns a tuple with the FirstExecutionRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstExecutionRunId

`func (o *ActivityTaskStartedResponse) SetFirstExecutionRunId(v string)`

SetFirstExecutionRunId sets FirstExecutionRunId field to given value.

### HasFirstExecutionRunId

`func (o *ActivityTaskStartedResponse) HasFirstExecutionRunId() bool`

HasFirstExecutionRunId returns a boolean if a field has been set.

### SetFirstExecutionRunIdNil

`func (o *ActivityTaskStartedResponse) SetFirstExecutionRunIdNil(b bool)`

 SetFirstExecutionRunIdNil sets the value for FirstExecutionRunId to be an explicit nil

### UnsetFirstExecutionRunId
`func (o *ActivityTaskStartedResponse) UnsetFirstExecutionRunId()`

UnsetFirstExecutionRunId ensures that no value is present for FirstExecutionRunId, not even an explicit nil
### GetChainRunId

`func (o *ActivityTaskStartedResponse) GetChainRunId() string`

GetChainRunId returns the ChainRunId field if non-nil, zero value otherwise.

### GetChainRunIdOk

`func (o *ActivityTaskStartedResponse) GetChainRunIdOk() (*string, bool)`

GetChainRunIdOk returns a tuple with the ChainRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainRunId

`func (o *ActivityTaskStartedResponse) SetChainRunId(v string)`

SetChainRunId sets ChainRunId field to given value.

### HasChainRunId

`func (o *ActivityTaskStartedResponse) HasChainRunId() bool`

HasChainRunId returns a boolean if a field has been set.

### SetChainRunIdNil

`func (o *ActivityTaskStartedResponse) SetChainRunIdNil(b bool)`

 SetChainRunIdNil sets the value for ChainRunId to be an explicit nil

### UnsetChainRunId
`func (o *ActivityTaskStartedResponse) UnsetChainRunId()`

UnsetChainRunId ensures that no value is present for ChainRunId, not even an explicit nil
### GetScheduleId

`func (o *ActivityTaskStartedResponse) GetScheduleId() string`

GetScheduleId returns the ScheduleId field if non-nil, zero value otherwise.

### GetScheduleIdOk

`func (o *ActivityTaskStartedResponse) GetScheduleIdOk() (*string, bool)`

GetScheduleIdOk returns a tuple with the ScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleId

`func (o *ActivityTaskStartedResponse) SetScheduleId(v string)`

SetScheduleId sets ScheduleId field to given value.

### HasScheduleId

`func (o *ActivityTaskStartedResponse) HasScheduleId() bool`

HasScheduleId returns a boolean if a field has been set.

### SetScheduleIdNil

`func (o *ActivityTaskStartedResponse) SetScheduleIdNil(b bool)`

 SetScheduleIdNil sets the value for ScheduleId to be an explicit nil

### UnsetScheduleId
`func (o *ActivityTaskStartedResponse) UnsetScheduleId()`

UnsetScheduleId ensures that no value is present for ScheduleId, not even an explicit nil
### GetWorkflowExecId

`func (o *ActivityTaskStartedResponse) GetWorkflowExecId() string`

GetWorkflowExecId returns the WorkflowExecId field if non-nil, zero value otherwise.

### GetWorkflowExecIdOk

`func (o *ActivityTaskStartedResponse) GetWorkflowExecIdOk() (*string, bool)`

GetWorkflowExecIdOk returns a tuple with the WorkflowExecId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowExecId

`func (o *ActivityTaskStartedResponse) SetWorkflowExecId(v string)`

SetWorkflowExecId sets WorkflowExecId field to given value.


### GetWorkflowRunId

`func (o *ActivityTaskStartedResponse) GetWorkflowRunId() string`

GetWorkflowRunId returns the WorkflowRunId field if non-nil, zero value otherwise.

### GetWorkflowRunIdOk

`func (o *ActivityTaskStartedResponse) GetWorkflowRunIdOk() (*string, bool)`

GetWorkflowRunIdOk returns a tuple with the WorkflowRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowRunId

`func (o *ActivityTaskStartedResponse) SetWorkflowRunId(v string)`

SetWorkflowRunId sets WorkflowRunId field to given value.


### GetWorkflowName

`func (o *ActivityTaskStartedResponse) GetWorkflowName() string`

GetWorkflowName returns the WorkflowName field if non-nil, zero value otherwise.

### GetWorkflowNameOk

`func (o *ActivityTaskStartedResponse) GetWorkflowNameOk() (*string, bool)`

GetWorkflowNameOk returns a tuple with the WorkflowName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowName

`func (o *ActivityTaskStartedResponse) SetWorkflowName(v string)`

SetWorkflowName sets WorkflowName field to given value.


### GetEventType

`func (o *ActivityTaskStartedResponse) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *ActivityTaskStartedResponse) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *ActivityTaskStartedResponse) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetAttributes

`func (o *ActivityTaskStartedResponse) GetAttributes() ActivityTaskStartedAttributesResponse`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ActivityTaskStartedResponse) GetAttributesOk() (*ActivityTaskStartedAttributesResponse, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ActivityTaskStartedResponse) SetAttributes(v ActivityTaskStartedAttributesResponse)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


