# StreamDeploymentLogs200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **NullableString** | Filter by agent name | [optional] 
**Data** | Pointer to [**StreamDeploymentLogs200ResponseData**](StreamDeploymentLogs200ResponseData.md) |  | [optional] 

## Methods

### NewStreamDeploymentLogs200Response

`func NewStreamDeploymentLogs200Response() *StreamDeploymentLogs200Response`

NewStreamDeploymentLogs200Response instantiates a new StreamDeploymentLogs200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamDeploymentLogs200ResponseWithDefaults

`func NewStreamDeploymentLogs200ResponseWithDefaults() *StreamDeploymentLogs200Response`

NewStreamDeploymentLogs200ResponseWithDefaults instantiates a new StreamDeploymentLogs200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *StreamDeploymentLogs200Response) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *StreamDeploymentLogs200Response) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *StreamDeploymentLogs200Response) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *StreamDeploymentLogs200Response) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetId

`func (o *StreamDeploymentLogs200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StreamDeploymentLogs200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StreamDeploymentLogs200Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *StreamDeploymentLogs200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *StreamDeploymentLogs200Response) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *StreamDeploymentLogs200Response) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetData

`func (o *StreamDeploymentLogs200Response) GetData() StreamDeploymentLogs200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *StreamDeploymentLogs200Response) GetDataOk() (*StreamDeploymentLogs200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *StreamDeploymentLogs200Response) SetData(v StreamDeploymentLogs200ResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *StreamDeploymentLogs200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


