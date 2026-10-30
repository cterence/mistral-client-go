# DeploymentLogRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | **time.Time** |  | 
**TraceId** | **string** |  | 
**SpanId** | **string** |  | 
**SeverityText** | **string** |  | 
**Body** | **string** |  | 
**LogAttributes** | **map[string]string** |  | 

## Methods

### NewDeploymentLogRecord

`func NewDeploymentLogRecord(timestamp time.Time, traceId string, spanId string, severityText string, body string, logAttributes map[string]string, ) *DeploymentLogRecord`

NewDeploymentLogRecord instantiates a new DeploymentLogRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentLogRecordWithDefaults

`func NewDeploymentLogRecordWithDefaults() *DeploymentLogRecord`

NewDeploymentLogRecordWithDefaults instantiates a new DeploymentLogRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *DeploymentLogRecord) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *DeploymentLogRecord) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *DeploymentLogRecord) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetTraceId

`func (o *DeploymentLogRecord) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *DeploymentLogRecord) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *DeploymentLogRecord) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.


### GetSpanId

`func (o *DeploymentLogRecord) GetSpanId() string`

GetSpanId returns the SpanId field if non-nil, zero value otherwise.

### GetSpanIdOk

`func (o *DeploymentLogRecord) GetSpanIdOk() (*string, bool)`

GetSpanIdOk returns a tuple with the SpanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpanId

`func (o *DeploymentLogRecord) SetSpanId(v string)`

SetSpanId sets SpanId field to given value.


### GetSeverityText

`func (o *DeploymentLogRecord) GetSeverityText() string`

GetSeverityText returns the SeverityText field if non-nil, zero value otherwise.

### GetSeverityTextOk

`func (o *DeploymentLogRecord) GetSeverityTextOk() (*string, bool)`

GetSeverityTextOk returns a tuple with the SeverityText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityText

`func (o *DeploymentLogRecord) SetSeverityText(v string)`

SetSeverityText sets SeverityText field to given value.


### GetBody

`func (o *DeploymentLogRecord) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *DeploymentLogRecord) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *DeploymentLogRecord) SetBody(v string)`

SetBody sets Body field to given value.


### GetLogAttributes

`func (o *DeploymentLogRecord) GetLogAttributes() map[string]string`

GetLogAttributes returns the LogAttributes field if non-nil, zero value otherwise.

### GetLogAttributesOk

`func (o *DeploymentLogRecord) GetLogAttributesOk() (*map[string]string, bool)`

GetLogAttributesOk returns a tuple with the LogAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogAttributes

`func (o *DeploymentLogRecord) SetLogAttributes(v map[string]string)`

SetLogAttributes sets LogAttributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


