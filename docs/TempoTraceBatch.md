# TempoTraceBatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resource** | [**TempoTraceResource**](TempoTraceResource.md) | The resource of the batch | 
**ScopeSpans** | Pointer to [**[]TempoTraceScopeSpan**](TempoTraceScopeSpan.md) | The spans of the scope | [optional] 

## Methods

### NewTempoTraceBatch

`func NewTempoTraceBatch(resource TempoTraceResource, ) *TempoTraceBatch`

NewTempoTraceBatch instantiates a new TempoTraceBatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTempoTraceBatchWithDefaults

`func NewTempoTraceBatchWithDefaults() *TempoTraceBatch`

NewTempoTraceBatchWithDefaults instantiates a new TempoTraceBatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResource

`func (o *TempoTraceBatch) GetResource() TempoTraceResource`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *TempoTraceBatch) GetResourceOk() (*TempoTraceResource, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *TempoTraceBatch) SetResource(v TempoTraceResource)`

SetResource sets Resource field to given value.


### GetScopeSpans

`func (o *TempoTraceBatch) GetScopeSpans() []TempoTraceScopeSpan`

GetScopeSpans returns the ScopeSpans field if non-nil, zero value otherwise.

### GetScopeSpansOk

`func (o *TempoTraceBatch) GetScopeSpansOk() (*[]TempoTraceScopeSpan, bool)`

GetScopeSpansOk returns a tuple with the ScopeSpans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeSpans

`func (o *TempoTraceBatch) SetScopeSpans(v []TempoTraceScopeSpan)`

SetScopeSpans sets ScopeSpans field to given value.

### HasScopeSpans

`func (o *TempoTraceBatch) HasScopeSpans() bool`

HasScopeSpans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


