# TempoTraceEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the event | 
**TimeUnixNano** | **string** | The time of the event in Unix nano | 
**Attributes** | Pointer to [**[]TempoTraceAttribute**](TempoTraceAttribute.md) | The attributes of the event | [optional] 

## Methods

### NewTempoTraceEvent

`func NewTempoTraceEvent(name string, timeUnixNano string, ) *TempoTraceEvent`

NewTempoTraceEvent instantiates a new TempoTraceEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTempoTraceEventWithDefaults

`func NewTempoTraceEventWithDefaults() *TempoTraceEvent`

NewTempoTraceEventWithDefaults instantiates a new TempoTraceEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TempoTraceEvent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TempoTraceEvent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TempoTraceEvent) SetName(v string)`

SetName sets Name field to given value.


### GetTimeUnixNano

`func (o *TempoTraceEvent) GetTimeUnixNano() string`

GetTimeUnixNano returns the TimeUnixNano field if non-nil, zero value otherwise.

### GetTimeUnixNanoOk

`func (o *TempoTraceEvent) GetTimeUnixNanoOk() (*string, bool)`

GetTimeUnixNanoOk returns a tuple with the TimeUnixNano field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeUnixNano

`func (o *TempoTraceEvent) SetTimeUnixNano(v string)`

SetTimeUnixNano sets TimeUnixNano field to given value.


### GetAttributes

`func (o *TempoTraceEvent) GetAttributes() []TempoTraceAttribute`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *TempoTraceEvent) GetAttributesOk() (*[]TempoTraceAttribute, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *TempoTraceEvent) SetAttributes(v []TempoTraceAttribute)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *TempoTraceEvent) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


