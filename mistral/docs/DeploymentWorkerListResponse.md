# DeploymentWorkerListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Workers** | [**[]DeploymentWorkerResponse**](DeploymentWorkerResponse.md) | Workers registered for the deployment | 
**NextCursor** | **NullableString** | Cursor for the next page of results | 

## Methods

### NewDeploymentWorkerListResponse

`func NewDeploymentWorkerListResponse(workers []DeploymentWorkerResponse, nextCursor NullableString, ) *DeploymentWorkerListResponse`

NewDeploymentWorkerListResponse instantiates a new DeploymentWorkerListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentWorkerListResponseWithDefaults

`func NewDeploymentWorkerListResponseWithDefaults() *DeploymentWorkerListResponse`

NewDeploymentWorkerListResponseWithDefaults instantiates a new DeploymentWorkerListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkers

`func (o *DeploymentWorkerListResponse) GetWorkers() []DeploymentWorkerResponse`

GetWorkers returns the Workers field if non-nil, zero value otherwise.

### GetWorkersOk

`func (o *DeploymentWorkerListResponse) GetWorkersOk() (*[]DeploymentWorkerResponse, bool)`

GetWorkersOk returns a tuple with the Workers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkers

`func (o *DeploymentWorkerListResponse) SetWorkers(v []DeploymentWorkerResponse)`

SetWorkers sets Workers field to given value.


### GetNextCursor

`func (o *DeploymentWorkerListResponse) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *DeploymentWorkerListResponse) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *DeploymentWorkerListResponse) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.


### SetNextCursorNil

`func (o *DeploymentWorkerListResponse) SetNextCursorNil(b bool)`

 SetNextCursorNil sets the value for NextCursor to be an explicit nil

### UnsetNextCursor
`func (o *DeploymentWorkerListResponse) UnsetNextCursor()`

UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


