# WorkflowExecutionTraceSummarySpan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SpanId** | **string** | The ID of the span | 
**Name** | **string** | The name of the span | 
**StartTimeUnixNano** | **int32** | The start time of the span in nanoseconds since the Unix epoch | 
**EndTimeUnixNano** | **NullableInt32** | The end time of the span in nanoseconds since the Unix epoch | 
**Attributes** | [**map[string]WorkflowExecutionTraceSummaryAttributesValues**](WorkflowExecutionTraceSummaryAttributesValues.md) | The attributes of the span | 
**Events** | [**[]WorkflowExecutionTraceEvent**](WorkflowExecutionTraceEvent.md) | The events of the span | 
**Children** | Pointer to [**[]WorkflowExecutionTraceSummarySpan**](WorkflowExecutionTraceSummarySpan.md) | The child spans of the span | [optional] 

## Methods

### NewWorkflowExecutionTraceSummarySpan

`func NewWorkflowExecutionTraceSummarySpan(spanId string, name string, startTimeUnixNano int32, endTimeUnixNano NullableInt32, attributes map[string]WorkflowExecutionTraceSummaryAttributesValues, events []WorkflowExecutionTraceEvent, ) *WorkflowExecutionTraceSummarySpan`

NewWorkflowExecutionTraceSummarySpan instantiates a new WorkflowExecutionTraceSummarySpan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowExecutionTraceSummarySpanWithDefaults

`func NewWorkflowExecutionTraceSummarySpanWithDefaults() *WorkflowExecutionTraceSummarySpan`

NewWorkflowExecutionTraceSummarySpanWithDefaults instantiates a new WorkflowExecutionTraceSummarySpan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpanId

`func (o *WorkflowExecutionTraceSummarySpan) GetSpanId() string`

GetSpanId returns the SpanId field if non-nil, zero value otherwise.

### GetSpanIdOk

`func (o *WorkflowExecutionTraceSummarySpan) GetSpanIdOk() (*string, bool)`

GetSpanIdOk returns a tuple with the SpanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanId

`func (o *WorkflowExecutionTraceSummarySpan) SetSpanId(v string)`

SetSpanId sets SpanId field to given value.


### GetName

`func (o *WorkflowExecutionTraceSummarySpan) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowExecutionTraceSummarySpan) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowExecutionTraceSummarySpan) SetName(v string)`

SetName sets Name field to given value.


### GetStartTimeUnixNano

`func (o *WorkflowExecutionTraceSummarySpan) GetStartTimeUnixNano() int32`

GetStartTimeUnixNano returns the StartTimeUnixNano field if non-nil, zero value otherwise.

### GetStartTimeUnixNanoOk

`func (o *WorkflowExecutionTraceSummarySpan) GetStartTimeUnixNanoOk() (*int32, bool)`

GetStartTimeUnixNanoOk returns a tuple with the StartTimeUnixNano field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeUnixNano

`func (o *WorkflowExecutionTraceSummarySpan) SetStartTimeUnixNano(v int32)`

SetStartTimeUnixNano sets StartTimeUnixNano field to given value.


### GetEndTimeUnixNano

`func (o *WorkflowExecutionTraceSummarySpan) GetEndTimeUnixNano() int32`

GetEndTimeUnixNano returns the EndTimeUnixNano field if non-nil, zero value otherwise.

### GetEndTimeUnixNanoOk

`func (o *WorkflowExecutionTraceSummarySpan) GetEndTimeUnixNanoOk() (*int32, bool)`

GetEndTimeUnixNanoOk returns a tuple with the EndTimeUnixNano field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeUnixNano

`func (o *WorkflowExecutionTraceSummarySpan) SetEndTimeUnixNano(v int32)`

SetEndTimeUnixNano sets EndTimeUnixNano field to given value.


### SetEndTimeUnixNanoNil

`func (o *WorkflowExecutionTraceSummarySpan) SetEndTimeUnixNanoNil(b bool)`

 SetEndTimeUnixNanoNil sets the value for EndTimeUnixNano to be an explicit nil

### UnsetEndTimeUnixNano
`func (o *WorkflowExecutionTraceSummarySpan) UnsetEndTimeUnixNano()`

UnsetEndTimeUnixNano ensures that no value is present for EndTimeUnixNano, not even an explicit nil
### GetAttributes

`func (o *WorkflowExecutionTraceSummarySpan) GetAttributes() map[string]WorkflowExecutionTraceSummaryAttributesValues`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *WorkflowExecutionTraceSummarySpan) GetAttributesOk() (*map[string]WorkflowExecutionTraceSummaryAttributesValues, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *WorkflowExecutionTraceSummarySpan) SetAttributes(v map[string]WorkflowExecutionTraceSummaryAttributesValues)`

SetAttributes sets Attributes field to given value.


### GetEvents

`func (o *WorkflowExecutionTraceSummarySpan) GetEvents() []WorkflowExecutionTraceEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *WorkflowExecutionTraceSummarySpan) GetEventsOk() (*[]WorkflowExecutionTraceEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *WorkflowExecutionTraceSummarySpan) SetEvents(v []WorkflowExecutionTraceEvent)`

SetEvents sets Events field to given value.


### GetChildren

`func (o *WorkflowExecutionTraceSummarySpan) GetChildren() []WorkflowExecutionTraceSummarySpan`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *WorkflowExecutionTraceSummarySpan) GetChildrenOk() (*[]WorkflowExecutionTraceSummarySpan, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *WorkflowExecutionTraceSummarySpan) SetChildren(v []WorkflowExecutionTraceSummarySpan)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *WorkflowExecutionTraceSummarySpan) HasChildren() bool`

HasChildren returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


