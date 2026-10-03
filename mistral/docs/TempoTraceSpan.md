# TempoTraceSpan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TraceId** | **string** | The trace ID of the scope | 
**SpanId** | **string** | The span ID of the scope | 
**ParentSpanId** | Pointer to **NullableString** | The parent span ID of the scope | [optional] 
**Name** | **string** | The name of the scope | 
**Kind** | [**TempoTraceScopeKind**](TempoTraceScopeKind.md) | The kind of the scope | 
**StartTimeUnixNano** | **string** | The start time of the scope in Unix nano | 
**EndTimeUnixNano** | **string** | The end time of the scope in Unix nano | 
**Attributes** | Pointer to [**[]TempoTraceAttribute**](TempoTraceAttribute.md) | The attributes of the scope | [optional] 
**Events** | Pointer to [**[]TempoTraceEvent**](TempoTraceEvent.md) | The events of the scope | [optional] 

## Methods

### NewTempoTraceSpan

`func NewTempoTraceSpan(traceId string, spanId string, name string, kind TempoTraceScopeKind, startTimeUnixNano string, endTimeUnixNano string, ) *TempoTraceSpan`

NewTempoTraceSpan instantiates a new TempoTraceSpan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTempoTraceSpanWithDefaults

`func NewTempoTraceSpanWithDefaults() *TempoTraceSpan`

NewTempoTraceSpanWithDefaults instantiates a new TempoTraceSpan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTraceId

`func (o *TempoTraceSpan) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *TempoTraceSpan) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *TempoTraceSpan) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.


### GetSpanId

`func (o *TempoTraceSpan) GetSpanId() string`

GetSpanId returns the SpanId field if non-nil, zero value otherwise.

### GetSpanIdOk

`func (o *TempoTraceSpan) GetSpanIdOk() (*string, bool)`

GetSpanIdOk returns a tuple with the SpanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanId

`func (o *TempoTraceSpan) SetSpanId(v string)`

SetSpanId sets SpanId field to given value.


### GetParentSpanId

`func (o *TempoTraceSpan) GetParentSpanId() string`

GetParentSpanId returns the ParentSpanId field if non-nil, zero value otherwise.

### GetParentSpanIdOk

`func (o *TempoTraceSpan) GetParentSpanIdOk() (*string, bool)`

GetParentSpanIdOk returns a tuple with the ParentSpanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSpanId

`func (o *TempoTraceSpan) SetParentSpanId(v string)`

SetParentSpanId sets ParentSpanId field to given value.

### HasParentSpanId

`func (o *TempoTraceSpan) HasParentSpanId() bool`

HasParentSpanId returns a boolean if a field has been set.

### SetParentSpanIdNil

`func (o *TempoTraceSpan) SetParentSpanIdNil(b bool)`

 SetParentSpanIdNil sets the value for ParentSpanId to be an explicit nil

### UnsetParentSpanId
`func (o *TempoTraceSpan) UnsetParentSpanId()`

UnsetParentSpanId ensures that no value is present for ParentSpanId, not even an explicit nil
### GetName

`func (o *TempoTraceSpan) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TempoTraceSpan) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TempoTraceSpan) SetName(v string)`

SetName sets Name field to given value.


### GetKind

`func (o *TempoTraceSpan) GetKind() TempoTraceScopeKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *TempoTraceSpan) GetKindOk() (*TempoTraceScopeKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *TempoTraceSpan) SetKind(v TempoTraceScopeKind)`

SetKind sets Kind field to given value.


### GetStartTimeUnixNano

`func (o *TempoTraceSpan) GetStartTimeUnixNano() string`

GetStartTimeUnixNano returns the StartTimeUnixNano field if non-nil, zero value otherwise.

### GetStartTimeUnixNanoOk

`func (o *TempoTraceSpan) GetStartTimeUnixNanoOk() (*string, bool)`

GetStartTimeUnixNanoOk returns a tuple with the StartTimeUnixNano field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimeUnixNano

`func (o *TempoTraceSpan) SetStartTimeUnixNano(v string)`

SetStartTimeUnixNano sets StartTimeUnixNano field to given value.


### GetEndTimeUnixNano

`func (o *TempoTraceSpan) GetEndTimeUnixNano() string`

GetEndTimeUnixNano returns the EndTimeUnixNano field if non-nil, zero value otherwise.

### GetEndTimeUnixNanoOk

`func (o *TempoTraceSpan) GetEndTimeUnixNanoOk() (*string, bool)`

GetEndTimeUnixNanoOk returns a tuple with the EndTimeUnixNano field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTimeUnixNano

`func (o *TempoTraceSpan) SetEndTimeUnixNano(v string)`

SetEndTimeUnixNano sets EndTimeUnixNano field to given value.


### GetAttributes

`func (o *TempoTraceSpan) GetAttributes() []TempoTraceAttribute`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TempoTraceSpan) GetAttributesOk() (*[]TempoTraceAttribute, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TempoTraceSpan) SetAttributes(v []TempoTraceAttribute)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *TempoTraceSpan) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetEvents

`func (o *TempoTraceSpan) GetEvents() []TempoTraceEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *TempoTraceSpan) GetEventsOk() (*[]TempoTraceEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *TempoTraceSpan) SetEvents(v []TempoTraceEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *TempoTraceSpan) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


