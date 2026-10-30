# DeploymentWorkerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Worker name | 
**CreatedAt** | **time.Time** | When the worker first registered | 
**LastHeartbeat** | **time.Time** | When the worker last registered | 
**UpdatedAt** | **time.Time** | Deprecated alias of last_heartbeat | 
**IsActive** | **bool** | Whether this worker&#39;s liveness key is currently alive | 
**Location** | Pointer to [**NullableDeploymentLocation**](DeploymentLocation.md) | Where the worker is running; null if the worker did not report a location | [optional] 

## Methods

### NewDeploymentWorkerResponse

`func NewDeploymentWorkerResponse(name string, createdAt time.Time, lastHeartbeat time.Time, updatedAt time.Time, isActive bool, ) *DeploymentWorkerResponse`

NewDeploymentWorkerResponse instantiates a new DeploymentWorkerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentWorkerResponseWithDefaults

`func NewDeploymentWorkerResponseWithDefaults() *DeploymentWorkerResponse`

NewDeploymentWorkerResponseWithDefaults instantiates a new DeploymentWorkerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DeploymentWorkerResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeploymentWorkerResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeploymentWorkerResponse) SetName(v string)`

SetName sets Name field to given value.


### GetCreatedAt

`func (o *DeploymentWorkerResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DeploymentWorkerResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DeploymentWorkerResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetLastHeartbeat

`func (o *DeploymentWorkerResponse) GetLastHeartbeat() time.Time`

GetLastHeartbeat returns the LastHeartbeat field if non-nil, zero value otherwise.

### GetLastHeartbeatOk

`func (o *DeploymentWorkerResponse) GetLastHeartbeatOk() (*time.Time, bool)`

GetLastHeartbeatOk returns a tuple with the LastHeartbeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastHeartbeat

`func (o *DeploymentWorkerResponse) SetLastHeartbeat(v time.Time)`

SetLastHeartbeat sets LastHeartbeat field to given value.


### GetUpdatedAt

`func (o *DeploymentWorkerResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DeploymentWorkerResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DeploymentWorkerResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetIsActive

`func (o *DeploymentWorkerResponse) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *DeploymentWorkerResponse) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *DeploymentWorkerResponse) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetLocation

`func (o *DeploymentWorkerResponse) GetLocation() DeploymentLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *DeploymentWorkerResponse) GetLocationOk() (*DeploymentLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *DeploymentWorkerResponse) SetLocation(v DeploymentLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *DeploymentWorkerResponse) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *DeploymentWorkerResponse) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *DeploymentWorkerResponse) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


