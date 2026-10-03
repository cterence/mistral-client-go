# StreamV1WorkflowsExecutionsExecutionIdStreamGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | Pointer to **NullableString** | Filter by agent name | [optional] 
**Data** | Pointer to [**StreamEventSsePayload**](StreamEventSsePayload.md) |  | [optional] 
**Id** | Pointer to **NullableString** | Filter by agent name | [optional] 
**Retry** | Pointer to **NullableInt32** | Set before sending audio. Streaming delay updates are rejected after audio starts. | [optional] 

## Methods

### NewStreamV1WorkflowsExecutionsExecutionIdStreamGet200Response

`func NewStreamV1WorkflowsExecutionsExecutionIdStreamGet200Response() *StreamV1WorkflowsExecutionsExecutionIdStreamGet200Response`

NewStreamV1WorkflowsExecutionsExecutionIdStreamGet200Response instantiates a new StreamV1WorkflowsExecutionsExecutionIdStreamGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamV1WorkflowsExecutionsExecutionIdStreamGet200ResponseWithDefaults

`func NewStreamV1WorkflowsExecutionsExecutionIdStreamGet200ResponseWithDefaults() *StreamV1WorkflowsExecutionsExecutionIdStreamGet200Response`

NewStreamV1WorkflowsExecutionsExecutionIdStreamGet200ResponseWithDefaults instantiates a new StreamV1WorkflowsExecutionsExecutionIdStreamGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *StreamV1WorkflowsExecutionsExecutionIdStreamGet200Response) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *StreamV1WorkflowsExecutionsExecutionIdStreamGet200Response) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *StreamV1WorkflowsExecutionsExecutionIdStreamGet200Response) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *StreamV1WorkflowsExecutionsExecutionIdStreamGet200Response) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### SetEventNil

`func (o *StreamV1WorkflowsExecutionsExecutionIdStreamGet200Response) SetEventNil(b bool)`

 SetEventNil sets the value for Event to be an explicit nil

### UnsetEvent
`func (o *StreamV1WorkflowsExecutionsExecutionIdStreamGet200Response) UnsetEvent()`

UnsetEvent ensures that no value is present for Event, not even an explicit nil
### GetData

`func (o *StreamV1WorkflowsExecutionsExecutionIdStreamGet200Response) GetData() StreamEventSsePayload`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *StreamV1WorkflowsExecutionsExecutionIdStreamGet200Response) GetDataOk() (*StreamEventSsePayload, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *StreamV1WorkflowsExecutionsExecutionIdStreamGet200Response) SetData(v StreamEventSsePayload)`

SetData sets Data field to given value.

### HasData

`func (o *StreamV1WorkflowsExecutionsExecutionIdStreamGet200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetId

`func (o *StreamV1WorkflowsExecutionsExecutionIdStreamGet200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StreamV1WorkflowsExecutionsExecutionIdStreamGet200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StreamV1WorkflowsExecutionsExecutionIdStreamGet200Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *StreamV1WorkflowsExecutionsExecutionIdStreamGet200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *StreamV1WorkflowsExecutionsExecutionIdStreamGet200Response) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *StreamV1WorkflowsExecutionsExecutionIdStreamGet200Response) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetRetry

`func (o *StreamV1WorkflowsExecutionsExecutionIdStreamGet200Response) GetRetry() int32`

GetRetry returns the Retry field if non-nil, zero value otherwise.

### GetRetryOk

`func (o *StreamV1WorkflowsExecutionsExecutionIdStreamGet200Response) GetRetryOk() (*int32, bool)`

GetRetryOk returns a tuple with the Retry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetry

`func (o *StreamV1WorkflowsExecutionsExecutionIdStreamGet200Response) SetRetry(v int32)`

SetRetry sets Retry field to given value.

### HasRetry

`func (o *StreamV1WorkflowsExecutionsExecutionIdStreamGet200Response) HasRetry() bool`

HasRetry returns a boolean if a field has been set.

### SetRetryNil

`func (o *StreamV1WorkflowsExecutionsExecutionIdStreamGet200Response) SetRetryNil(b bool)`

 SetRetryNil sets the value for Retry to be an explicit nil

### UnsetRetry
`func (o *StreamV1WorkflowsExecutionsExecutionIdStreamGet200Response) UnsetRetry()`

UnsetRetry ensures that no value is present for Retry, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


