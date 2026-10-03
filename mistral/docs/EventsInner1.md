# EventsInner1

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
**EndTimeUnixMs** | Pointer to **int32** | The end time of the event in milliseconds since the Unix epoch | [optional] 
**Error** | Pointer to **string** | The error message, if any | [optional] 

## Methods

### NewEventsInner1

`func NewEventsInner1(name string, id string, timestampUnixNano int32, attributes map[string]WorkflowExecutionTraceSummaryAttributesValues, startTimeUnixMs int32, ) *EventsInner1`

NewEventsInner1 instantiates a new EventsInner1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventsInner1WithDefaults

`func NewEventsInner1WithDefaults() *EventsInner1`

NewEventsInner1WithDefaults instantiates a new EventsInner1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EventsInner1) GetType() EventType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventsInner1) GetTypeOk() (*EventType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EventsInner1) SetType(v EventType)`

SetType sets Type field to given value.

### HasType

`func (o *EventsInner1) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *EventsInner1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EventsInner1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EventsInner1) SetName(v string)`

SetName sets Name field to given value.


### GetId

`func (o *EventsInner1) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EventsInner1) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EventsInner1) SetId(v string)`

SetId sets Id field to given value.


### GetTimestampUnixNano

`func (o *EventsInner1) GetTimestampUnixNano() int32`

GetTimestampUnixNano returns the TimestampUnixNano field if non-nil, zero value otherwise.

### GetTimestampUnixNanoOk

`func (o *EventsInner1) GetTimestampUnixNanoOk() (*int32, bool)`

GetTimestampUnixNanoOk returns a tuple with the TimestampUnixNano field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampUnixNano

`func (o *EventsInner1) SetTimestampUnixNano(v int32)`

SetTimestampUnixNano sets TimestampUnixNano field to given value.


### GetAttributes

`func (o *EventsInner1) GetAttributes() map[string]WorkflowExecutionTraceSummaryAttributesValues`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *EventsInner1) GetAttributesOk() (*map[string]WorkflowExecutionTraceSummaryAttributesValues, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *EventsInner1) SetAttributes(v map[string]WorkflowExecutionTraceSummaryAttributesValues)`

SetAttributes sets Attributes field to given value.


### GetInternal

`func (o *EventsInner1) GetInternal() bool`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *EventsInner1) GetInternalOk() (*bool, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *EventsInner1) SetInternal(v bool)`

SetInternal sets Internal field to given value.

### HasInternal

`func (o *EventsInner1) HasInternal() bool`

HasInternal returns a boolean if a field has been set.

### GetStatus

`func (o *EventsInner1) GetStatus() EventProgressStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EventsInner1) GetStatusOk() (*EventProgressStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EventsInner1) SetStatus(v EventProgressStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EventsInner1) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStartTimeUnixMs

`func (o *EventsInner1) GetStartTimeUnixMs() int32`

GetStartTimeUnixMs returns the StartTimeUnixMs field if non-nil, zero value otherwise.

### GetStartTimeUnixMsOk

`func (o *EventsInner1) GetStartTimeUnixMsOk() (*int32, bool)`

GetStartTimeUnixMsOk returns a tuple with the StartTimeUnixMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeUnixMs

`func (o *EventsInner1) SetStartTimeUnixMs(v int32)`

SetStartTimeUnixMs sets StartTimeUnixMs field to given value.


### GetEndTimeUnixMs

`func (o *EventsInner1) GetEndTimeUnixMs() int32`

GetEndTimeUnixMs returns the EndTimeUnixMs field if non-nil, zero value otherwise.

### GetEndTimeUnixMsOk

`func (o *EventsInner1) GetEndTimeUnixMsOk() (*int32, bool)`

GetEndTimeUnixMsOk returns a tuple with the EndTimeUnixMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeUnixMs

`func (o *EventsInner1) SetEndTimeUnixMs(v int32)`

SetEndTimeUnixMs sets EndTimeUnixMs field to given value.

### HasEndTimeUnixMs

`func (o *EventsInner1) HasEndTimeUnixMs() bool`

HasEndTimeUnixMs returns a boolean if a field has been set.

### GetError

`func (o *EventsInner1) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *EventsInner1) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *EventsInner1) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *EventsInner1) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


