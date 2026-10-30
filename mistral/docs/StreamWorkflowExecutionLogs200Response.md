# StreamWorkflowExecutionLogs200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **NullableString** | Filter by agent name | [optional] 
**Data** | Pointer to [**StreamWorkflowExecutionLogs200ResponseData**](StreamWorkflowExecutionLogs200ResponseData.md) |  | [optional] 

## Methods

### NewStreamWorkflowExecutionLogs200Response

`func NewStreamWorkflowExecutionLogs200Response() *StreamWorkflowExecutionLogs200Response`

NewStreamWorkflowExecutionLogs200Response instantiates a new StreamWorkflowExecutionLogs200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamWorkflowExecutionLogs200ResponseWithDefaults

`func NewStreamWorkflowExecutionLogs200ResponseWithDefaults() *StreamWorkflowExecutionLogs200Response`

NewStreamWorkflowExecutionLogs200ResponseWithDefaults instantiates a new StreamWorkflowExecutionLogs200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *StreamWorkflowExecutionLogs200Response) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *StreamWorkflowExecutionLogs200Response) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *StreamWorkflowExecutionLogs200Response) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *StreamWorkflowExecutionLogs200Response) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetId

`func (o *StreamWorkflowExecutionLogs200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StreamWorkflowExecutionLogs200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StreamWorkflowExecutionLogs200Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *StreamWorkflowExecutionLogs200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *StreamWorkflowExecutionLogs200Response) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *StreamWorkflowExecutionLogs200Response) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetData

`func (o *StreamWorkflowExecutionLogs200Response) GetData() StreamWorkflowExecutionLogs200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *StreamWorkflowExecutionLogs200Response) GetDataOk() (*StreamWorkflowExecutionLogs200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *StreamWorkflowExecutionLogs200Response) SetData(v StreamWorkflowExecutionLogs200ResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *StreamWorkflowExecutionLogs200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


