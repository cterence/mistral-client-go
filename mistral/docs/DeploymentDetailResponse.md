# DeploymentDetailResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the deployment | 
**Name** | **string** | Deployment name | 
**IsActive** | **bool** | Whether at least one worker is currently live | 
**CreatedAt** | **time.Time** | When the deployment was first registered | 
**UpdatedAt** | **time.Time** | When the deployment was last updated | 
**Workers** | [**[]DeploymentWorkerResponse**](DeploymentWorkerResponse.md) | Workers registered for the deployment | 

## Methods

### NewDeploymentDetailResponse

`func NewDeploymentDetailResponse(id string, name string, isActive bool, createdAt time.Time, updatedAt time.Time, workers []DeploymentWorkerResponse, ) *DeploymentDetailResponse`

NewDeploymentDetailResponse instantiates a new DeploymentDetailResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentDetailResponseWithDefaults

`func NewDeploymentDetailResponseWithDefaults() *DeploymentDetailResponse`

NewDeploymentDetailResponseWithDefaults instantiates a new DeploymentDetailResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeploymentDetailResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeploymentDetailResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeploymentDetailResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *DeploymentDetailResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeploymentDetailResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeploymentDetailResponse) SetName(v string)`

SetName sets Name field to given value.


### GetIsActive

`func (o *DeploymentDetailResponse) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *DeploymentDetailResponse) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *DeploymentDetailResponse) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetCreatedAt

`func (o *DeploymentDetailResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DeploymentDetailResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DeploymentDetailResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *DeploymentDetailResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DeploymentDetailResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DeploymentDetailResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetWorkers

`func (o *DeploymentDetailResponse) GetWorkers() []DeploymentWorkerResponse`

GetWorkers returns the Workers field if non-nil, zero value otherwise.

### GetWorkersOk

`func (o *DeploymentDetailResponse) GetWorkersOk() (*[]DeploymentWorkerResponse, bool)`

GetWorkersOk returns a tuple with the Workers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkers

`func (o *DeploymentDetailResponse) SetWorkers(v []DeploymentWorkerResponse)`

SetWorkers sets Workers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


