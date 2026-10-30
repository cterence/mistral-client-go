# StreamWorkflowExecutionLogs200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | **time.Time** |  | 
**TraceId** | **string** |  | 
**SpanId** | **string** |  | 
**SeverityText** | **string** |  | 
**Body** | **string** |  | 
**LogAttributes** | **map[string]string** |  | 
**Error** | **string** |  | 
**Reason** | **string** |  | 

## Methods

### NewStreamWorkflowExecutionLogs200ResponseData

`func NewStreamWorkflowExecutionLogs200ResponseData(timestamp time.Time, traceId string, spanId string, severityText string, body string, logAttributes map[string]string, error_ string, reason string, ) *StreamWorkflowExecutionLogs200ResponseData`

NewStreamWorkflowExecutionLogs200ResponseData instantiates a new StreamWorkflowExecutionLogs200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamWorkflowExecutionLogs200ResponseDataWithDefaults

`func NewStreamWorkflowExecutionLogs200ResponseDataWithDefaults() *StreamWorkflowExecutionLogs200ResponseData`

NewStreamWorkflowExecutionLogs200ResponseDataWithDefaults instantiates a new StreamWorkflowExecutionLogs200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *StreamWorkflowExecutionLogs200ResponseData) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *StreamWorkflowExecutionLogs200ResponseData) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *StreamWorkflowExecutionLogs200ResponseData) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetTraceId

`func (o *StreamWorkflowExecutionLogs200ResponseData) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *StreamWorkflowExecutionLogs200ResponseData) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *StreamWorkflowExecutionLogs200ResponseData) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.


### GetSpanId

`func (o *StreamWorkflowExecutionLogs200ResponseData) GetSpanId() string`

GetSpanId returns the SpanId field if non-nil, zero value otherwise.

### GetSpanIdOk

`func (o *StreamWorkflowExecutionLogs200ResponseData) GetSpanIdOk() (*string, bool)`

GetSpanIdOk returns a tuple with the SpanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanId

`func (o *StreamWorkflowExecutionLogs200ResponseData) SetSpanId(v string)`

SetSpanId sets SpanId field to given value.


### GetSeverityText

`func (o *StreamWorkflowExecutionLogs200ResponseData) GetSeverityText() string`

GetSeverityText returns the SeverityText field if non-nil, zero value otherwise.

### GetSeverityTextOk

`func (o *StreamWorkflowExecutionLogs200ResponseData) GetSeverityTextOk() (*string, bool)`

GetSeverityTextOk returns a tuple with the SeverityText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityText

`func (o *StreamWorkflowExecutionLogs200ResponseData) SetSeverityText(v string)`

SetSeverityText sets SeverityText field to given value.


### GetBody

`func (o *StreamWorkflowExecutionLogs200ResponseData) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *StreamWorkflowExecutionLogs200ResponseData) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *StreamWorkflowExecutionLogs200ResponseData) SetBody(v string)`

SetBody sets Body field to given value.


### GetLogAttributes

`func (o *StreamWorkflowExecutionLogs200ResponseData) GetLogAttributes() map[string]string`

GetLogAttributes returns the LogAttributes field if non-nil, zero value otherwise.

### GetLogAttributesOk

`func (o *StreamWorkflowExecutionLogs200ResponseData) GetLogAttributesOk() (*map[string]string, bool)`

GetLogAttributesOk returns a tuple with the LogAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogAttributes

`func (o *StreamWorkflowExecutionLogs200ResponseData) SetLogAttributes(v map[string]string)`

SetLogAttributes sets LogAttributes field to given value.


### GetError

`func (o *StreamWorkflowExecutionLogs200ResponseData) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *StreamWorkflowExecutionLogs200ResponseData) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *StreamWorkflowExecutionLogs200ResponseData) SetError(v string)`

SetError sets Error field to given value.


### GetReason

`func (o *StreamWorkflowExecutionLogs200ResponseData) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *StreamWorkflowExecutionLogs200ResponseData) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *StreamWorkflowExecutionLogs200ResponseData) SetReason(v string)`

SetReason sets Reason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


