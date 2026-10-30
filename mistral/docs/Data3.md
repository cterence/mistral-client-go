# Data3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | **string** | Unique identifier for this event instance. | 
**EventTimestamp** | **int32** | Unix timestamp in nanoseconds when the event was created. | 
**RootWorkflowExecId** | **string** | Execution ID of the root workflow that initiated this execution chain. | 
**ParentWorkflowExecId** | **string** | Execution ID of the parent workflow that initiated this execution. If this is a root workflow, this field is not set. | 
**ContinuedRunId** | Pointer to **string** | Run ID of the execution this run continued from. Non-null for continue-as-new runs. | [optional] 
**FirstExecutionRunId** | Pointer to **string** | Run ID of the first execution in this workflow chain. Equals workflow_run_id on a fresh start. Continue-as-new, retry and reset runs keep the original value. | [optional] 
**ChainRunId** | Pointer to **string** | Identifies one start of an execution. Shared by its continue-as-new, retry and reset runs, and by every sub-workflow below it. Starting the same execution ID again gives a new value. | [optional] 
**ScheduleId** | Pointer to **string** | Temporal schedule ID that triggered this execution, if any. | [optional] 
**WorkflowExecId** | **string** | Execution ID of the workflow that emitted this event. | 
**WorkflowRunId** | **string** | Run ID of the workflow execution. Changes on continue-as-new while workflow_exec_id stays the same. | 
**WorkflowName** | **string** | The registered name of the workflow that emitted this event. | 
**EventType** | **string** | Event type discriminator. | [default to "ACTIVITY_TASK_FAILED"]
**Attributes** | [**ActivityTaskFailedAttributes**](ActivityTaskFailedAttributes.md) | Event-specific attributes. | 

## Methods

### NewData3

`func NewData3(eventId string, eventTimestamp int32, rootWorkflowExecId string, parentWorkflowExecId string, workflowExecId string, workflowRunId string, workflowName string, eventType string, attributes ActivityTaskFailedAttributes, ) *Data3`

NewData3 instantiates a new Data3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewData3WithDefaults

`func NewData3WithDefaults() *Data3`

NewData3WithDefaults instantiates a new Data3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventId

`func (o *Data3) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *Data3) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *Data3) SetEventId(v string)`

SetEventId sets EventId field to given value.


### GetEventTimestamp

`func (o *Data3) GetEventTimestamp() int32`

GetEventTimestamp returns the EventTimestamp field if non-nil, zero value otherwise.

### GetEventTimestampOk

`func (o *Data3) GetEventTimestampOk() (*int32, bool)`

GetEventTimestampOk returns a tuple with the EventTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTimestamp

`func (o *Data3) SetEventTimestamp(v int32)`

SetEventTimestamp sets EventTimestamp field to given value.


### GetRootWorkflowExecId

`func (o *Data3) GetRootWorkflowExecId() string`

GetRootWorkflowExecId returns the RootWorkflowExecId field if non-nil, zero value otherwise.

### GetRootWorkflowExecIdOk

`func (o *Data3) GetRootWorkflowExecIdOk() (*string, bool)`

GetRootWorkflowExecIdOk returns a tuple with the RootWorkflowExecId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootWorkflowExecId

`func (o *Data3) SetRootWorkflowExecId(v string)`

SetRootWorkflowExecId sets RootWorkflowExecId field to given value.


### GetParentWorkflowExecId

`func (o *Data3) GetParentWorkflowExecId() string`

GetParentWorkflowExecId returns the ParentWorkflowExecId field if non-nil, zero value otherwise.

### GetParentWorkflowExecIdOk

`func (o *Data3) GetParentWorkflowExecIdOk() (*string, bool)`

GetParentWorkflowExecIdOk returns a tuple with the ParentWorkflowExecId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentWorkflowExecId

`func (o *Data3) SetParentWorkflowExecId(v string)`

SetParentWorkflowExecId sets ParentWorkflowExecId field to given value.


### GetContinuedRunId

`func (o *Data3) GetContinuedRunId() string`

GetContinuedRunId returns the ContinuedRunId field if non-nil, zero value otherwise.

### GetContinuedRunIdOk

`func (o *Data3) GetContinuedRunIdOk() (*string, bool)`

GetContinuedRunIdOk returns a tuple with the ContinuedRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinuedRunId

`func (o *Data3) SetContinuedRunId(v string)`

SetContinuedRunId sets ContinuedRunId field to given value.

### HasContinuedRunId

`func (o *Data3) HasContinuedRunId() bool`

HasContinuedRunId returns a boolean if a field has been set.

### GetFirstExecutionRunId

`func (o *Data3) GetFirstExecutionRunId() string`

GetFirstExecutionRunId returns the FirstExecutionRunId field if non-nil, zero value otherwise.

### GetFirstExecutionRunIdOk

`func (o *Data3) GetFirstExecutionRunIdOk() (*string, bool)`

GetFirstExecutionRunIdOk returns a tuple with the FirstExecutionRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstExecutionRunId

`func (o *Data3) SetFirstExecutionRunId(v string)`

SetFirstExecutionRunId sets FirstExecutionRunId field to given value.

### HasFirstExecutionRunId

`func (o *Data3) HasFirstExecutionRunId() bool`

HasFirstExecutionRunId returns a boolean if a field has been set.

### GetChainRunId

`func (o *Data3) GetChainRunId() string`

GetChainRunId returns the ChainRunId field if non-nil, zero value otherwise.

### GetChainRunIdOk

`func (o *Data3) GetChainRunIdOk() (*string, bool)`

GetChainRunIdOk returns a tuple with the ChainRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainRunId

`func (o *Data3) SetChainRunId(v string)`

SetChainRunId sets ChainRunId field to given value.

### HasChainRunId

`func (o *Data3) HasChainRunId() bool`

HasChainRunId returns a boolean if a field has been set.

### GetScheduleId

`func (o *Data3) GetScheduleId() string`

GetScheduleId returns the ScheduleId field if non-nil, zero value otherwise.

### GetScheduleIdOk

`func (o *Data3) GetScheduleIdOk() (*string, bool)`

GetScheduleIdOk returns a tuple with the ScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleId

`func (o *Data3) SetScheduleId(v string)`

SetScheduleId sets ScheduleId field to given value.

### HasScheduleId

`func (o *Data3) HasScheduleId() bool`

HasScheduleId returns a boolean if a field has been set.

### GetWorkflowExecId

`func (o *Data3) GetWorkflowExecId() string`

GetWorkflowExecId returns the WorkflowExecId field if non-nil, zero value otherwise.

### GetWorkflowExecIdOk

`func (o *Data3) GetWorkflowExecIdOk() (*string, bool)`

GetWorkflowExecIdOk returns a tuple with the WorkflowExecId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowExecId

`func (o *Data3) SetWorkflowExecId(v string)`

SetWorkflowExecId sets WorkflowExecId field to given value.


### GetWorkflowRunId

`func (o *Data3) GetWorkflowRunId() string`

GetWorkflowRunId returns the WorkflowRunId field if non-nil, zero value otherwise.

### GetWorkflowRunIdOk

`func (o *Data3) GetWorkflowRunIdOk() (*string, bool)`

GetWorkflowRunIdOk returns a tuple with the WorkflowRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowRunId

`func (o *Data3) SetWorkflowRunId(v string)`

SetWorkflowRunId sets WorkflowRunId field to given value.


### GetWorkflowName

`func (o *Data3) GetWorkflowName() string`

GetWorkflowName returns the WorkflowName field if non-nil, zero value otherwise.

### GetWorkflowNameOk

`func (o *Data3) GetWorkflowNameOk() (*string, bool)`

GetWorkflowNameOk returns a tuple with the WorkflowName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowName

`func (o *Data3) SetWorkflowName(v string)`

SetWorkflowName sets WorkflowName field to given value.


### GetEventType

`func (o *Data3) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *Data3) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *Data3) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetAttributes

`func (o *Data3) GetAttributes() ActivityTaskFailedAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Data3) GetAttributesOk() (*ActivityTaskFailedAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Data3) SetAttributes(v ActivityTaskFailedAttributes)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


