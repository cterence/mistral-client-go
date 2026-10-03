# StreamEventSsePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Stream** | **string** |  | 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Data** | [**Data3**](Data3.md) |  | 
**WorkflowContext** | [**StreamEventWorkflowContext**](StreamEventWorkflowContext.md) |  | 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**BrokerSequence** | **int32** |  | 

## Methods

### NewStreamEventSsePayload

`func NewStreamEventSsePayload(stream string, data Data3, workflowContext StreamEventWorkflowContext, brokerSequence int32, ) *StreamEventSsePayload`

NewStreamEventSsePayload instantiates a new StreamEventSsePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamEventSsePayloadWithDefaults

`func NewStreamEventSsePayloadWithDefaults() *StreamEventSsePayload`

NewStreamEventSsePayloadWithDefaults instantiates a new StreamEventSsePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStream

`func (o *StreamEventSsePayload) GetStream() string`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *StreamEventSsePayload) GetStreamOk() (*string, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *StreamEventSsePayload) SetStream(v string)`

SetStream sets Stream field to given value.


### GetTimestamp

`func (o *StreamEventSsePayload) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *StreamEventSsePayload) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *StreamEventSsePayload) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *StreamEventSsePayload) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetData

`func (o *StreamEventSsePayload) GetData() Data3`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *StreamEventSsePayload) GetDataOk() (*Data3, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *StreamEventSsePayload) SetData(v Data3)`

SetData sets Data field to given value.


### GetWorkflowContext

`func (o *StreamEventSsePayload) GetWorkflowContext() StreamEventWorkflowContext`

GetWorkflowContext returns the WorkflowContext field if non-nil, zero value otherwise.

### GetWorkflowContextOk

`func (o *StreamEventSsePayload) GetWorkflowContextOk() (*StreamEventWorkflowContext, bool)`

GetWorkflowContextOk returns a tuple with the WorkflowContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowContext

`func (o *StreamEventSsePayload) SetWorkflowContext(v StreamEventWorkflowContext)`

SetWorkflowContext sets WorkflowContext field to given value.


### GetMetadata

`func (o *StreamEventSsePayload) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *StreamEventSsePayload) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *StreamEventSsePayload) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *StreamEventSsePayload) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetBrokerSequence

`func (o *StreamEventSsePayload) GetBrokerSequence() int32`

GetBrokerSequence returns the BrokerSequence field if non-nil, zero value otherwise.

### GetBrokerSequenceOk

`func (o *StreamEventSsePayload) GetBrokerSequenceOk() (*int32, bool)`

GetBrokerSequenceOk returns a tuple with the BrokerSequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerSequence

`func (o *StreamEventSsePayload) SetBrokerSequence(v int32)`

SetBrokerSequence sets BrokerSequence field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


