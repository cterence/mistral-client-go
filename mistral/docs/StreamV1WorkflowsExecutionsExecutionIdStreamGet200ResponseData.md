# StreamV1WorkflowsExecutionsExecutionIdStreamGet200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Stream** | **string** |  | 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Data** | [**Data3**](Data3.md) |  | 
**WorkflowContext** | [**StreamEventWorkflowContext**](StreamEventWorkflowContext.md) |  | 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**BrokerSequence** | **int32** |  | 
**Error** | **string** |  | 
**Reason** | **string** |  | 

## Methods

### NewStreamV1WorkflowsExecutionsExecutionIdStreamGet200ResponseData

`func NewStreamV1WorkflowsExecutionsExecutionIdStreamGet200ResponseData(stream string, data Data3, workflowContext StreamEventWorkflowContext, brokerSequence int32, error_ string, reason string, ) *StreamV1WorkflowsExecutionsExecutionIdStreamGet200ResponseData`

NewStreamV1WorkflowsExecutionsExecutionIdStreamGet200ResponseData instantiates a new StreamV1WorkflowsExecutionsExecutionIdStreamGet200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamV1WorkflowsExecutionsExecutionIdStreamGet200ResponseDataWithDefaults

`func NewStreamV1WorkflowsExecutionsExecutionIdStreamGet200ResponseDataWithDefaults() *StreamV1WorkflowsExecutionsExecutionIdStreamGet200ResponseData`

NewStreamV1WorkflowsExecutionsExecutionIdStreamGet200ResponseDataWithDefaults instantiates a new StreamV1WorkflowsExecutionsExecutionIdStreamGet200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStream

`func (o *StreamV1WorkflowsExecutionsExecutionIdStreamGet200ResponseData) GetStream() string`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *StreamV1WorkflowsExecutionsExecutionIdStreamGet200ResponseData) GetStreamOk() (*string, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *StreamV1WorkflowsExecutionsExecutionIdStreamGet200ResponseData) SetStream(v string)`

SetStream sets Stream field to given value.


### GetTimestamp

`func (o *StreamV1WorkflowsExecutionsExecutionIdStreamGet200ResponseData) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *StreamV1WorkflowsExecutionsExecutionIdStreamGet200ResponseData) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *StreamV1WorkflowsExecutionsExecutionIdStreamGet200ResponseData) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *StreamV1WorkflowsExecutionsExecutionIdStreamGet200ResponseData) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetData

`func (o *StreamV1WorkflowsExecutionsExecutionIdStreamGet200ResponseData) GetData() Data3`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *StreamV1WorkflowsExecutionsExecutionIdStreamGet200ResponseData) GetDataOk() (*Data3, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *StreamV1WorkflowsExecutionsExecutionIdStreamGet200ResponseData) SetData(v Data3)`

SetData sets Data field to given value.


### GetWorkflowContext

`func (o *StreamV1WorkflowsExecutionsExecutionIdStreamGet200ResponseData) GetWorkflowContext() StreamEventWorkflowContext`

GetWorkflowContext returns the WorkflowContext field if non-nil, zero value otherwise.

### GetWorkflowContextOk

`func (o *StreamV1WorkflowsExecutionsExecutionIdStreamGet200ResponseData) GetWorkflowContextOk() (*StreamEventWorkflowContext, bool)`

GetWorkflowContextOk returns a tuple with the WorkflowContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowContext

`func (o *StreamV1WorkflowsExecutionsExecutionIdStreamGet200ResponseData) SetWorkflowContext(v StreamEventWorkflowContext)`

SetWorkflowContext sets WorkflowContext field to given value.


### GetMetadata

`func (o *StreamV1WorkflowsExecutionsExecutionIdStreamGet200ResponseData) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *StreamV1WorkflowsExecutionsExecutionIdStreamGet200ResponseData) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *StreamV1WorkflowsExecutionsExecutionIdStreamGet200ResponseData) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *StreamV1WorkflowsExecutionsExecutionIdStreamGet200ResponseData) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetBrokerSequence

`func (o *StreamV1WorkflowsExecutionsExecutionIdStreamGet200ResponseData) GetBrokerSequence() int32`

GetBrokerSequence returns the BrokerSequence field if non-nil, zero value otherwise.

### GetBrokerSequenceOk

`func (o *StreamV1WorkflowsExecutionsExecutionIdStreamGet200ResponseData) GetBrokerSequenceOk() (*int32, bool)`

GetBrokerSequenceOk returns a tuple with the BrokerSequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerSequence

`func (o *StreamV1WorkflowsExecutionsExecutionIdStreamGet200ResponseData) SetBrokerSequence(v int32)`

SetBrokerSequence sets BrokerSequence field to given value.


### GetError

`func (o *StreamV1WorkflowsExecutionsExecutionIdStreamGet200ResponseData) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *StreamV1WorkflowsExecutionsExecutionIdStreamGet200ResponseData) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *StreamV1WorkflowsExecutionsExecutionIdStreamGet200ResponseData) SetError(v string)`

SetError sets Error field to given value.


### GetReason

`func (o *StreamV1WorkflowsExecutionsExecutionIdStreamGet200ResponseData) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *StreamV1WorkflowsExecutionsExecutionIdStreamGet200ResponseData) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *StreamV1WorkflowsExecutionsExecutionIdStreamGet200ResponseData) SetReason(v string)`

SetReason sets Reason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


