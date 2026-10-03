# WorkflowExecutionTraceEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**EventType**](EventType.md) |  | [optional] [default to EVENTTYPE_EVENT]
**Name** | **string** | Name of the event | 
**Id** | **string** | The ID of the event | 
**TimestampUnixNano** | **int32** | The timestamp of the event in nanoseconds since the Unix epoch | 
**Attributes** | [**map[string]WorkflowExecutionTraceSummaryAttributesValues**](WorkflowExecutionTraceSummaryAttributesValues.md) | The attributes of the event | 
**Internal** | Pointer to **bool** | Whether the event is internal | [optional] [default to false]

## Methods

### NewWorkflowExecutionTraceEvent

`func NewWorkflowExecutionTraceEvent(name string, id string, timestampUnixNano int32, attributes map[string]WorkflowExecutionTraceSummaryAttributesValues, ) *WorkflowExecutionTraceEvent`

NewWorkflowExecutionTraceEvent instantiates a new WorkflowExecutionTraceEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowExecutionTraceEventWithDefaults

`func NewWorkflowExecutionTraceEventWithDefaults() *WorkflowExecutionTraceEvent`

NewWorkflowExecutionTraceEventWithDefaults instantiates a new WorkflowExecutionTraceEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WorkflowExecutionTraceEvent) GetType() EventType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WorkflowExecutionTraceEvent) GetTypeOk() (*EventType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WorkflowExecutionTraceEvent) SetType(v EventType)`

SetType sets Type field to given value.

### HasType

`func (o *WorkflowExecutionTraceEvent) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *WorkflowExecutionTraceEvent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowExecutionTraceEvent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowExecutionTraceEvent) SetName(v string)`

SetName sets Name field to given value.


### GetId

`func (o *WorkflowExecutionTraceEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowExecutionTraceEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowExecutionTraceEvent) SetId(v string)`

SetId sets Id field to given value.


### GetTimestampUnixNano

`func (o *WorkflowExecutionTraceEvent) GetTimestampUnixNano() int32`

GetTimestampUnixNano returns the TimestampUnixNano field if non-nil, zero value otherwise.

### GetTimestampUnixNanoOk

`func (o *WorkflowExecutionTraceEvent) GetTimestampUnixNanoOk() (*int32, bool)`

GetTimestampUnixNanoOk returns a tuple with the TimestampUnixNano field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestampUnixNano

`func (o *WorkflowExecutionTraceEvent) SetTimestampUnixNano(v int32)`

SetTimestampUnixNano sets TimestampUnixNano field to given value.


### GetAttributes

`func (o *WorkflowExecutionTraceEvent) GetAttributes() map[string]WorkflowExecutionTraceSummaryAttributesValues`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *WorkflowExecutionTraceEvent) GetAttributesOk() (*map[string]WorkflowExecutionTraceSummaryAttributesValues, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *WorkflowExecutionTraceEvent) SetAttributes(v map[string]WorkflowExecutionTraceSummaryAttributesValues)`

SetAttributes sets Attributes field to given value.


### GetInternal

`func (o *WorkflowExecutionTraceEvent) GetInternal() bool`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *WorkflowExecutionTraceEvent) GetInternalOk() (*bool, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *WorkflowExecutionTraceEvent) SetInternal(v bool)`

SetInternal sets Internal field to given value.

### HasInternal

`func (o *WorkflowExecutionTraceEvent) HasInternal() bool`

HasInternal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


