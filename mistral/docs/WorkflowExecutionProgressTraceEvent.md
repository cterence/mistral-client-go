# WorkflowExecutionProgressTraceEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**EventType**](EventType.md) |  | [optional] [default to EVENTTYPE_EVENT_PROGRESS]
**Name** | **string** | Name of the event | 
**Id** | **string** | The ID of the event | 
**TimestampUnixNano** | **int32** | The timestamp of the event in nanoseconds since the Unix epoch | 
**Attributes** | [**map[string]WorkflowExecutionTraceSummaryAttributesValues**](WorkflowExecutionTraceSummaryAttributesValues.md) | The attributes of the event | 
**Internal** | Pointer to **bool** | Whether the event is internal | [optional] [default to false]
**Status** | Pointer to [**EventProgressStatus**](EventProgressStatus.md) | The progress message | [optional] [default to EVENTPROGRESSSTATUS_RUNNING]
**StartTimeUnixMs** | **int32** | The start time of the event in milliseconds since the Unix epoch | 
**EndTimeUnixMs** | Pointer to **NullableInt32** | The end time of the event in milliseconds since the Unix epoch | [optional] 
**Error** | Pointer to **NullableString** | The error message, if any | [optional] 

## Methods

### NewWorkflowExecutionProgressTraceEvent

`func NewWorkflowExecutionProgressTraceEvent(name string, id string, timestampUnixNano int32, attributes map[string]WorkflowExecutionTraceSummaryAttributesValues, startTimeUnixMs int32, ) *WorkflowExecutionProgressTraceEvent`

NewWorkflowExecutionProgressTraceEvent instantiates a new WorkflowExecutionProgressTraceEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowExecutionProgressTraceEventWithDefaults

`func NewWorkflowExecutionProgressTraceEventWithDefaults() *WorkflowExecutionProgressTraceEvent`

NewWorkflowExecutionProgressTraceEventWithDefaults instantiates a new WorkflowExecutionProgressTraceEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WorkflowExecutionProgressTraceEvent) GetType() EventType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WorkflowExecutionProgressTraceEvent) GetTypeOk() (*EventType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WorkflowExecutionProgressTraceEvent) SetType(v EventType)`

SetType sets Type field to given value.

### HasType

`func (o *WorkflowExecutionProgressTraceEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *WorkflowExecutionProgressTraceEvent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowExecutionProgressTraceEvent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowExecutionProgressTraceEvent) SetName(v string)`

SetName sets Name field to given value.


### GetId

`func (o *WorkflowExecutionProgressTraceEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowExecutionProgressTraceEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowExecutionProgressTraceEvent) SetId(v string)`

SetId sets Id field to given value.


### GetTimestampUnixNano

`func (o *WorkflowExecutionProgressTraceEvent) GetTimestampUnixNano() int32`

GetTimestampUnixNano returns the TimestampUnixNano field if non-nil, zero value otherwise.

### GetTimestampUnixNanoOk

`func (o *WorkflowExecutionProgressTraceEvent) GetTimestampUnixNanoOk() (*int32, bool)`

GetTimestampUnixNanoOk returns a tuple with the TimestampUnixNano field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampUnixNano

`func (o *WorkflowExecutionProgressTraceEvent) SetTimestampUnixNano(v int32)`

SetTimestampUnixNano sets TimestampUnixNano field to given value.


### GetAttributes

`func (o *WorkflowExecutionProgressTraceEvent) GetAttributes() map[string]WorkflowExecutionTraceSummaryAttributesValues`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *WorkflowExecutionProgressTraceEvent) GetAttributesOk() (*map[string]WorkflowExecutionTraceSummaryAttributesValues, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *WorkflowExecutionProgressTraceEvent) SetAttributes(v map[string]WorkflowExecutionTraceSummaryAttributesValues)`

SetAttributes sets Attributes field to given value.


### GetInternal

`func (o *WorkflowExecutionProgressTraceEvent) GetInternal() bool`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *WorkflowExecutionProgressTraceEvent) GetInternalOk() (*bool, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *WorkflowExecutionProgressTraceEvent) SetInternal(v bool)`

SetInternal sets Internal field to given value.

### HasInternal

`func (o *WorkflowExecutionProgressTraceEvent) HasInternal() bool`

HasInternal returns a boolean if a field has been set.

### GetStatus

`func (o *WorkflowExecutionProgressTraceEvent) GetStatus() EventProgressStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkflowExecutionProgressTraceEvent) GetStatusOk() (*EventProgressStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkflowExecutionProgressTraceEvent) SetStatus(v EventProgressStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WorkflowExecutionProgressTraceEvent) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStartTimeUnixMs

`func (o *WorkflowExecutionProgressTraceEvent) GetStartTimeUnixMs() int32`

GetStartTimeUnixMs returns the StartTimeUnixMs field if non-nil, zero value otherwise.

### GetStartTimeUnixMsOk

`func (o *WorkflowExecutionProgressTraceEvent) GetStartTimeUnixMsOk() (*int32, bool)`

GetStartTimeUnixMsOk returns a tuple with the StartTimeUnixMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeUnixMs

`func (o *WorkflowExecutionProgressTraceEvent) SetStartTimeUnixMs(v int32)`

SetStartTimeUnixMs sets StartTimeUnixMs field to given value.


### GetEndTimeUnixMs

`func (o *WorkflowExecutionProgressTraceEvent) GetEndTimeUnixMs() int32`

GetEndTimeUnixMs returns the EndTimeUnixMs field if non-nil, zero value otherwise.

### GetEndTimeUnixMsOk

`func (o *WorkflowExecutionProgressTraceEvent) GetEndTimeUnixMsOk() (*int32, bool)`

GetEndTimeUnixMsOk returns a tuple with the EndTimeUnixMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeUnixMs

`func (o *WorkflowExecutionProgressTraceEvent) SetEndTimeUnixMs(v int32)`

SetEndTimeUnixMs sets EndTimeUnixMs field to given value.

### HasEndTimeUnixMs

`func (o *WorkflowExecutionProgressTraceEvent) HasEndTimeUnixMs() bool`

HasEndTimeUnixMs returns a boolean if a field has been set.

### SetEndTimeUnixMsNil

`func (o *WorkflowExecutionProgressTraceEvent) SetEndTimeUnixMsNil(b bool)`

 SetEndTimeUnixMsNil sets the value for EndTimeUnixMs to be an explicit nil

### UnsetEndTimeUnixMs
`func (o *WorkflowExecutionProgressTraceEvent) UnsetEndTimeUnixMs()`

UnsetEndTimeUnixMs ensures that no value is present for EndTimeUnixMs, not even an explicit nil
### GetError

`func (o *WorkflowExecutionProgressTraceEvent) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *WorkflowExecutionProgressTraceEvent) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *WorkflowExecutionProgressTraceEvent) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *WorkflowExecutionProgressTraceEvent) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *WorkflowExecutionProgressTraceEvent) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *WorkflowExecutionProgressTraceEvent) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


