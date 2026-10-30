# TelemetrySpanReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TraceId** | **string** |  | 
**SpanId** | **string** |  | 

## Methods

### NewTelemetrySpanReference

`func NewTelemetrySpanReference(traceId string, spanId string, ) *TelemetrySpanReference`

NewTelemetrySpanReference instantiates a new TelemetrySpanReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelemetrySpanReferenceWithDefaults

`func NewTelemetrySpanReferenceWithDefaults() *TelemetrySpanReference`

NewTelemetrySpanReferenceWithDefaults instantiates a new TelemetrySpanReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTraceId

`func (o *TelemetrySpanReference) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *TelemetrySpanReference) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *TelemetrySpanReference) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.


### GetSpanId

`func (o *TelemetrySpanReference) GetSpanId() string`

GetSpanId returns the SpanId field if non-nil, zero value otherwise.

### GetSpanIdOk

`func (o *TelemetrySpanReference) GetSpanIdOk() (*string, bool)`

GetSpanIdOk returns a tuple with the SpanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanId

`func (o *TelemetrySpanReference) SetSpanId(v string)`

SetSpanId sets SpanId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


