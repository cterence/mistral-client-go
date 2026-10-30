# DeploymentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the deployment | 
**Name** | **string** | Deployment name | 
**IsActive** | **bool** | Whether at least one worker is currently live | 
**IsHardened** | Pointer to **bool** | Whether the deployment only accepts registrations from authorized principals | [optional] [default to false]
**CreatedAt** | **time.Time** | When the deployment was first registered | 
**LastHeartbeat** | **time.Time** | When the deployment last saw a worker registration | 
**UpdatedAt** | **time.Time** | Deprecated alias of last_heartbeat | 
**Owner** | Pointer to **NullableString** | User id that owns the deployment, or null when it is administrator-managed | [optional] 
**Location** | Pointer to [**NullableDeploymentLocation**](DeploymentLocation.md) | Where the deployment is running | [optional] 
**WorkerCount** | Pointer to **int32** | Number of workers registered to the deployment | [optional] [default to 0]
**ActiveWorkerCount** | Pointer to **int32** | Number of workers currently live within the liveness cutoff | [optional] [default to 0]
**Locations** | Pointer to [**[]LocationType**](LocationType.md) | Distinct location types reported by the deployment&#39;s workers | [optional] 
**Managed** | Pointer to [**NullableManagedDeploymentResponse**](ManagedDeploymentResponse.md) | Live managed service state for managed deployments; null for self-hosted deployments or when managed services are unavailable | [optional] 

## Methods

### NewDeploymentResponse

`func NewDeploymentResponse(id string, name string, isActive bool, createdAt time.Time, lastHeartbeat time.Time, updatedAt time.Time, ) *DeploymentResponse`

NewDeploymentResponse instantiates a new DeploymentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentResponseWithDefaults

`func NewDeploymentResponseWithDefaults() *DeploymentResponse`

NewDeploymentResponseWithDefaults instantiates a new DeploymentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeploymentResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeploymentResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeploymentResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *DeploymentResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeploymentResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeploymentResponse) SetName(v string)`

SetName sets Name field to given value.


### GetIsActive

`func (o *DeploymentResponse) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *DeploymentResponse) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *DeploymentResponse) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetIsHardened

`func (o *DeploymentResponse) GetIsHardened() bool`

GetIsHardened returns the IsHardened field if non-nil, zero value otherwise.

### GetIsHardenedOk

`func (o *DeploymentResponse) GetIsHardenedOk() (*bool, bool)`

GetIsHardenedOk returns a tuple with the IsHardened field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHardened

`func (o *DeploymentResponse) SetIsHardened(v bool)`

SetIsHardened sets IsHardened field to given value.

### HasIsHardened

`func (o *DeploymentResponse) HasIsHardened() bool`

HasIsHardened returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DeploymentResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DeploymentResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DeploymentResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetLastHeartbeat

`func (o *DeploymentResponse) GetLastHeartbeat() time.Time`

GetLastHeartbeat returns the LastHeartbeat field if non-nil, zero value otherwise.

### GetLastHeartbeatOk

`func (o *DeploymentResponse) GetLastHeartbeatOk() (*time.Time, bool)`

GetLastHeartbeatOk returns a tuple with the LastHeartbeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastHeartbeat

`func (o *DeploymentResponse) SetLastHeartbeat(v time.Time)`

SetLastHeartbeat sets LastHeartbeat field to given value.


### GetUpdatedAt

`func (o *DeploymentResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DeploymentResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DeploymentResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetOwner

`func (o *DeploymentResponse) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *DeploymentResponse) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *DeploymentResponse) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *DeploymentResponse) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *DeploymentResponse) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *DeploymentResponse) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetLocation

`func (o *DeploymentResponse) GetLocation() DeploymentLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *DeploymentResponse) GetLocationOk() (*DeploymentLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *DeploymentResponse) SetLocation(v DeploymentLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *DeploymentResponse) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *DeploymentResponse) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *DeploymentResponse) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetWorkerCount

`func (o *DeploymentResponse) GetWorkerCount() int32`

GetWorkerCount returns the WorkerCount field if non-nil, zero value otherwise.

### GetWorkerCountOk

`func (o *DeploymentResponse) GetWorkerCountOk() (*int32, bool)`

GetWorkerCountOk returns a tuple with the WorkerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerCount

`func (o *DeploymentResponse) SetWorkerCount(v int32)`

SetWorkerCount sets WorkerCount field to given value.

### HasWorkerCount

`func (o *DeploymentResponse) HasWorkerCount() bool`

HasWorkerCount returns a boolean if a field has been set.

### GetActiveWorkerCount

`func (o *DeploymentResponse) GetActiveWorkerCount() int32`

GetActiveWorkerCount returns the ActiveWorkerCount field if non-nil, zero value otherwise.

### GetActiveWorkerCountOk

`func (o *DeploymentResponse) GetActiveWorkerCountOk() (*int32, bool)`

GetActiveWorkerCountOk returns a tuple with the ActiveWorkerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveWorkerCount

`func (o *DeploymentResponse) SetActiveWorkerCount(v int32)`

SetActiveWorkerCount sets ActiveWorkerCount field to given value.

### HasActiveWorkerCount

`func (o *DeploymentResponse) HasActiveWorkerCount() bool`

HasActiveWorkerCount returns a boolean if a field has been set.

### GetLocations

`func (o *DeploymentResponse) GetLocations() []LocationType`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *DeploymentResponse) GetLocationsOk() (*[]LocationType, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *DeploymentResponse) SetLocations(v []LocationType)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *DeploymentResponse) HasLocations() bool`

HasLocations returns a boolean if a field has been set.

### GetManaged

`func (o *DeploymentResponse) GetManaged() ManagedDeploymentResponse`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *DeploymentResponse) GetManagedOk() (*ManagedDeploymentResponse, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *DeploymentResponse) SetManaged(v ManagedDeploymentResponse)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *DeploymentResponse) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### SetManagedNil

`func (o *DeploymentResponse) SetManagedNil(b bool)`

 SetManagedNil sets the value for Managed to be an explicit nil

### UnsetManaged
`func (o *DeploymentResponse) UnsetManaged()`

UnsetManaged ensures that no value is present for Managed, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


