# ActivityTaskCompletedResponse

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
**EventType** | **string** | Event type discriminator. | [default to "ACTIVITY_TASK_COMPLETED"]
**Attributes** | [**ActivityTaskCompletedAttributesResponse**](ActivityTaskCompletedAttributesResponse.md) | Event-specific attributes. | 

## Methods

### NewActivityTaskCompletedResponse

`func NewActivityTaskCompletedResponse(eventId string, eventTimestamp int32, rootWorkflowExecId string, parentWorkflowExecId NullableString, workflowExecId string, workflowRunId string, workflowName string, eventType string, attributes ActivityTaskCompletedAttributesResponse, ) *ActivityTaskCompletedResponse`

NewActivityTaskCompletedResponse instantiates a new ActivityTaskCompletedResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityTaskCompletedResponseWithDefaults

`func NewActivityTaskCompletedResponseWithDefaults() *ActivityTaskCompletedResponse`

NewActivityTaskCompletedResponseWithDefaults instantiates a new ActivityTaskCompletedResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventId

`func (o *ActivityTaskCompletedResponse) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *ActivityTaskCompletedResponse) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *ActivityTaskCompletedResponse) SetEventId(v string)`

SetEventId sets EventId field to given value.


### GetEventTimestamp

`func (o *ActivityTaskCompletedResponse) GetEventTimestamp() int32`

GetEventTimestamp returns the EventTimestamp field if non-nil, zero value otherwise.

### GetEventTimestampOk

`func (o *ActivityTaskCompletedResponse) GetEventTimestampOk() (*int32, bool)`

GetEventTimestampOk returns a tuple with the EventTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTimestamp

`func (o *ActivityTaskCompletedResponse) SetEventTimestamp(v int32)`

SetEventTimestamp sets EventTimestamp field to given value.


### GetRootWorkflowExecId

`func (o *ActivityTaskCompletedResponse) GetRootWorkflowExecId() string`

GetRootWorkflowExecId returns the RootWorkflowExecId field if non-nil, zero value otherwise.

### GetRootWorkflowExecIdOk

`func (o *ActivityTaskCompletedResponse) GetRootWorkflowExecIdOk() (*string, bool)`

GetRootWorkflowExecIdOk returns a tuple with the RootWorkflowExecId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootWorkflowExecId

`func (o *ActivityTaskCompletedResponse) SetRootWorkflowExecId(v string)`

SetRootWorkflowExecId sets RootWorkflowExecId field to given value.


### GetParentWorkflowExecId

`func (o *ActivityTaskCompletedResponse) GetParentWorkflowExecId() string`

GetParentWorkflowExecId returns the ParentWorkflowExecId field if non-nil, zero value otherwise.

### GetParentWorkflowExecIdOk

`func (o *ActivityTaskCompletedResponse) GetParentWorkflowExecIdOk() (*string, bool)`

GetParentWorkflowExecIdOk returns a tuple with the ParentWorkflowExecId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentWorkflowExecId

`func (o *ActivityTaskCompletedResponse) SetParentWorkflowExecId(v string)`

SetParentWorkflowExecId sets ParentWorkflowExecId field to given value.


### SetParentWorkflowExecIdNil

`func (o *ActivityTaskCompletedResponse) SetParentWorkflowExecIdNil(b bool)`

 SetParentWorkflowExecIdNil sets the value for ParentWorkflowExecId to be an explicit nil

### UnsetParentWorkflowExecId
`func (o *ActivityTaskCompletedResponse) UnsetParentWorkflowExecId()`

UnsetParentWorkflowExecId ensures that no value is present for ParentWorkflowExecId, not even an explicit nil
### GetWorkflowExecId

`func (o *ActivityTaskCompletedResponse) GetWorkflowExecId() string`

GetWorkflowExecId returns the WorkflowExecId field if non-nil, zero value otherwise.

### GetWorkflowExecIdOk

`func (o *ActivityTaskCompletedResponse) GetWorkflowExecIdOk() (*string, bool)`

GetWorkflowExecIdOk returns a tuple with the WorkflowExecId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowExecId

`func (o *ActivityTaskCompletedResponse) SetWorkflowExecId(v string)`

SetWorkflowExecId sets WorkflowExecId field to given value.


### GetWorkflowRunId

`func (o *ActivityTaskCompletedResponse) GetWorkflowRunId() string`

GetWorkflowRunId returns the WorkflowRunId field if non-nil, zero value otherwise.

### GetWorkflowRunIdOk

`func (o *ActivityTaskCompletedResponse) GetWorkflowRunIdOk() (*string, bool)`

GetWorkflowRunIdOk returns a tuple with the WorkflowRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowRunId

`func (o *ActivityTaskCompletedResponse) SetWorkflowRunId(v string)`

SetWorkflowRunId sets WorkflowRunId field to given value.


### GetWorkflowName

`func (o *ActivityTaskCompletedResponse) GetWorkflowName() string`

GetWorkflowName returns the WorkflowName field if non-nil, zero value otherwise.

### GetWorkflowNameOk

`func (o *ActivityTaskCompletedResponse) GetWorkflowNameOk() (*string, bool)`

GetWorkflowNameOk returns a tuple with the WorkflowName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowName

`func (o *ActivityTaskCompletedResponse) SetWorkflowName(v string)`

SetWorkflowName sets WorkflowName field to given value.


### GetEventType

`func (o *ActivityTaskCompletedResponse) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *ActivityTaskCompletedResponse) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *ActivityTaskCompletedResponse) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetAttributes

`func (o *ActivityTaskCompletedResponse) GetAttributes() ActivityTaskCompletedAttributesResponse`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ActivityTaskCompletedResponse) GetAttributesOk() (*ActivityTaskCompletedAttributesResponse, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ActivityTaskCompletedResponse) SetAttributes(v ActivityTaskCompletedAttributesResponse)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


