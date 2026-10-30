# ExecutionTraceInfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OtelTraceId** | Pointer to **NullableString** | The ID of the trace, if available | [optional] 
**HasTraceData** | Pointer to **bool** | Whether trace data is available in the trace backend for this execution | [optional] [default to false]

## Methods

### NewExecutionTraceInfoResponse

`func NewExecutionTraceInfoResponse() *ExecutionTraceInfoResponse`

NewExecutionTraceInfoResponse instantiates a new ExecutionTraceInfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecutionTraceInfoResponseWithDefaults

`func NewExecutionTraceInfoResponseWithDefaults() *ExecutionTraceInfoResponse`

NewExecutionTraceInfoResponseWithDefaults instantiates a new ExecutionTraceInfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOtelTraceId

`func (o *ExecutionTraceInfoResponse) GetOtelTraceId() string`

GetOtelTraceId returns the OtelTraceId field if non-nil, zero value otherwise.

### GetOtelTraceIdOk

`func (o *ExecutionTraceInfoResponse) GetOtelTraceIdOk() (*string, bool)`

GetOtelTraceIdOk returns a tuple with the OtelTraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtelTraceId

`func (o *ExecutionTraceInfoResponse) SetOtelTraceId(v string)`

SetOtelTraceId sets OtelTraceId field to given value.

### HasOtelTraceId

`func (o *ExecutionTraceInfoResponse) HasOtelTraceId() bool`

HasOtelTraceId returns a boolean if a field has been set.

### SetOtelTraceIdNil

`func (o *ExecutionTraceInfoResponse) SetOtelTraceIdNil(b bool)`

 SetOtelTraceIdNil sets the value for OtelTraceId to be an explicit nil

### UnsetOtelTraceId
`func (o *ExecutionTraceInfoResponse) UnsetOtelTraceId()`

UnsetOtelTraceId ensures that no value is present for OtelTraceId, not even an explicit nil
### GetHasTraceData

`func (o *ExecutionTraceInfoResponse) GetHasTraceData() bool`

GetHasTraceData returns the HasTraceData field if non-nil, zero value otherwise.

### GetHasTraceDataOk

`func (o *ExecutionTraceInfoResponse) GetHasTraceDataOk() (*bool, bool)`

GetHasTraceDataOk returns a tuple with the HasTraceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasTraceData

`func (o *ExecutionTraceInfoResponse) SetHasTraceData(v bool)`

SetHasTraceData sets HasTraceData field to given value.

### HasHasTraceData

`func (o *ExecutionTraceInfoResponse) HasHasTraceData() bool`

HasHasTraceData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


