# TempoTraceScopeSpan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scope** | [**TempoTraceScope**](TempoTraceScope.md) | The scope of the span | 
**Spans** | Pointer to [**[]TempoTraceSpan**](TempoTraceSpan.md) | The spans of the scope | [optional] 

## Methods

### NewTempoTraceScopeSpan

`func NewTempoTraceScopeSpan(scope TempoTraceScope, ) *TempoTraceScopeSpan`

NewTempoTraceScopeSpan instantiates a new TempoTraceScopeSpan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTempoTraceScopeSpanWithDefaults

`func NewTempoTraceScopeSpanWithDefaults() *TempoTraceScopeSpan`

NewTempoTraceScopeSpanWithDefaults instantiates a new TempoTraceScopeSpan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScope

`func (o *TempoTraceScopeSpan) GetScope() TempoTraceScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *TempoTraceScopeSpan) GetScopeOk() (*TempoTraceScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *TempoTraceScopeSpan) SetScope(v TempoTraceScope)`

SetScope sets Scope field to given value.


### GetSpans

`func (o *TempoTraceScopeSpan) GetSpans() []TempoTraceSpan`

GetSpans returns the Spans field if non-nil, zero value otherwise.

### GetSpansOk

`func (o *TempoTraceScopeSpan) GetSpansOk() (*[]TempoTraceSpan, bool)`

GetSpansOk returns a tuple with the Spans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpans

`func (o *TempoTraceScopeSpan) SetSpans(v []TempoTraceSpan)`

SetSpans sets Spans field to given value.

### HasSpans

`func (o *TempoTraceScopeSpan) HasSpans() bool`

HasSpans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


