# ManagedDeploymentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | **string** |  | 
**Name** | **string** |  | 
**Spec** | [**DeploymentWorkerSpecResponse**](DeploymentWorkerSpecResponse.md) |  | 
**Resources** | [**DeploymentResourceConfig**](DeploymentResourceConfig.md) |  | 
**Status** | [**DeploymentObservedState**](DeploymentObservedState.md) |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**Stopped** | Pointer to **bool** |  | [optional] [default to false]
**RolloutStatus** | Pointer to **NullableString** |  | [optional] 
**CreatedBy** | Pointer to **NullableString** |  | [optional] 
**UpdatedBy** | Pointer to **NullableString** |  | [optional] 
**DeployedBy** | Pointer to **NullableString** |  | [optional] 
**DeployedAt** | Pointer to **NullableTime** |  | [optional] 
**IsHardened** | Pointer to **bool** |  | [optional] [default to false]
**RuntimeCredentialId** | Pointer to **NullableString** |  | [optional] 
**RuntimePrincipalType** | Pointer to [**NullablePrincipalType**](PrincipalType.md) |  | [optional] 

## Methods

### NewManagedDeploymentResponse

`func NewManagedDeploymentResponse(serviceId string, name string, spec DeploymentWorkerSpecResponse, resources DeploymentResourceConfig, status DeploymentObservedState, createdAt time.Time, updatedAt time.Time, ) *ManagedDeploymentResponse`

NewManagedDeploymentResponse instantiates a new ManagedDeploymentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedDeploymentResponseWithDefaults

`func NewManagedDeploymentResponseWithDefaults() *ManagedDeploymentResponse`

NewManagedDeploymentResponseWithDefaults instantiates a new ManagedDeploymentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *ManagedDeploymentResponse) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ManagedDeploymentResponse) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ManagedDeploymentResponse) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetName

`func (o *ManagedDeploymentResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ManagedDeploymentResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ManagedDeploymentResponse) SetName(v string)`

SetName sets Name field to given value.


### GetSpec

`func (o *ManagedDeploymentResponse) GetSpec() DeploymentWorkerSpecResponse`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *ManagedDeploymentResponse) GetSpecOk() (*DeploymentWorkerSpecResponse, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *ManagedDeploymentResponse) SetSpec(v DeploymentWorkerSpecResponse)`

SetSpec sets Spec field to given value.


### GetResources

`func (o *ManagedDeploymentResponse) GetResources() DeploymentResourceConfig`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *ManagedDeploymentResponse) GetResourcesOk() (*DeploymentResourceConfig, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *ManagedDeploymentResponse) SetResources(v DeploymentResourceConfig)`

SetResources sets Resources field to given value.


### GetStatus

`func (o *ManagedDeploymentResponse) GetStatus() DeploymentObservedState`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ManagedDeploymentResponse) GetStatusOk() (*DeploymentObservedState, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ManagedDeploymentResponse) SetStatus(v DeploymentObservedState)`

SetStatus sets Status field to given value.


### GetCreatedAt

`func (o *ManagedDeploymentResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ManagedDeploymentResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ManagedDeploymentResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ManagedDeploymentResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ManagedDeploymentResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ManagedDeploymentResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetStopped

`func (o *ManagedDeploymentResponse) GetStopped() bool`

GetStopped returns the Stopped field if non-nil, zero value otherwise.

### GetStoppedOk

`func (o *ManagedDeploymentResponse) GetStoppedOk() (*bool, bool)`

GetStoppedOk returns a tuple with the Stopped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopped

`func (o *ManagedDeploymentResponse) SetStopped(v bool)`

SetStopped sets Stopped field to given value.

### HasStopped

`func (o *ManagedDeploymentResponse) HasStopped() bool`

HasStopped returns a boolean if a field has been set.

### GetRolloutStatus

`func (o *ManagedDeploymentResponse) GetRolloutStatus() string`

GetRolloutStatus returns the RolloutStatus field if non-nil, zero value otherwise.

### GetRolloutStatusOk

`func (o *ManagedDeploymentResponse) GetRolloutStatusOk() (*string, bool)`

GetRolloutStatusOk returns a tuple with the RolloutStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolloutStatus

`func (o *ManagedDeploymentResponse) SetRolloutStatus(v string)`

SetRolloutStatus sets RolloutStatus field to given value.

### HasRolloutStatus

`func (o *ManagedDeploymentResponse) HasRolloutStatus() bool`

HasRolloutStatus returns a boolean if a field has been set.

### SetRolloutStatusNil

`func (o *ManagedDeploymentResponse) SetRolloutStatusNil(b bool)`

 SetRolloutStatusNil sets the value for RolloutStatus to be an explicit nil

### UnsetRolloutStatus
`func (o *ManagedDeploymentResponse) UnsetRolloutStatus()`

UnsetRolloutStatus ensures that no value is present for RolloutStatus, not even an explicit nil
### GetCreatedBy

`func (o *ManagedDeploymentResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ManagedDeploymentResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ManagedDeploymentResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ManagedDeploymentResponse) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *ManagedDeploymentResponse) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *ManagedDeploymentResponse) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetUpdatedBy

`func (o *ManagedDeploymentResponse) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *ManagedDeploymentResponse) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *ManagedDeploymentResponse) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *ManagedDeploymentResponse) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### SetUpdatedByNil

`func (o *ManagedDeploymentResponse) SetUpdatedByNil(b bool)`

 SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil

### UnsetUpdatedBy
`func (o *ManagedDeploymentResponse) UnsetUpdatedBy()`

UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil
### GetDeployedBy

`func (o *ManagedDeploymentResponse) GetDeployedBy() string`

GetDeployedBy returns the DeployedBy field if non-nil, zero value otherwise.

### GetDeployedByOk

`func (o *ManagedDeploymentResponse) GetDeployedByOk() (*string, bool)`

GetDeployedByOk returns a tuple with the DeployedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployedBy

`func (o *ManagedDeploymentResponse) SetDeployedBy(v string)`

SetDeployedBy sets DeployedBy field to given value.

### HasDeployedBy

`func (o *ManagedDeploymentResponse) HasDeployedBy() bool`

HasDeployedBy returns a boolean if a field has been set.

### SetDeployedByNil

`func (o *ManagedDeploymentResponse) SetDeployedByNil(b bool)`

 SetDeployedByNil sets the value for DeployedBy to be an explicit nil

### UnsetDeployedBy
`func (o *ManagedDeploymentResponse) UnsetDeployedBy()`

UnsetDeployedBy ensures that no value is present for DeployedBy, not even an explicit nil
### GetDeployedAt

`func (o *ManagedDeploymentResponse) GetDeployedAt() time.Time`

GetDeployedAt returns the DeployedAt field if non-nil, zero value otherwise.

### GetDeployedAtOk

`func (o *ManagedDeploymentResponse) GetDeployedAtOk() (*time.Time, bool)`

GetDeployedAtOk returns a tuple with the DeployedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployedAt

`func (o *ManagedDeploymentResponse) SetDeployedAt(v time.Time)`

SetDeployedAt sets DeployedAt field to given value.

### HasDeployedAt

`func (o *ManagedDeploymentResponse) HasDeployedAt() bool`

HasDeployedAt returns a boolean if a field has been set.

### SetDeployedAtNil

`func (o *ManagedDeploymentResponse) SetDeployedAtNil(b bool)`

 SetDeployedAtNil sets the value for DeployedAt to be an explicit nil

### UnsetDeployedAt
`func (o *ManagedDeploymentResponse) UnsetDeployedAt()`

UnsetDeployedAt ensures that no value is present for DeployedAt, not even an explicit nil
### GetIsHardened

`func (o *ManagedDeploymentResponse) GetIsHardened() bool`

GetIsHardened returns the IsHardened field if non-nil, zero value otherwise.

### GetIsHardenedOk

`func (o *ManagedDeploymentResponse) GetIsHardenedOk() (*bool, bool)`

GetIsHardenedOk returns a tuple with the IsHardened field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHardened

`func (o *ManagedDeploymentResponse) SetIsHardened(v bool)`

SetIsHardened sets IsHardened field to given value.

### HasIsHardened

`func (o *ManagedDeploymentResponse) HasIsHardened() bool`

HasIsHardened returns a boolean if a field has been set.

### GetRuntimeCredentialId

`func (o *ManagedDeploymentResponse) GetRuntimeCredentialId() string`

GetRuntimeCredentialId returns the RuntimeCredentialId field if non-nil, zero value otherwise.

### GetRuntimeCredentialIdOk

`func (o *ManagedDeploymentResponse) GetRuntimeCredentialIdOk() (*string, bool)`

GetRuntimeCredentialIdOk returns a tuple with the RuntimeCredentialId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeCredentialId

`func (o *ManagedDeploymentResponse) SetRuntimeCredentialId(v string)`

SetRuntimeCredentialId sets RuntimeCredentialId field to given value.

### HasRuntimeCredentialId

`func (o *ManagedDeploymentResponse) HasRuntimeCredentialId() bool`

HasRuntimeCredentialId returns a boolean if a field has been set.

### SetRuntimeCredentialIdNil

`func (o *ManagedDeploymentResponse) SetRuntimeCredentialIdNil(b bool)`

 SetRuntimeCredentialIdNil sets the value for RuntimeCredentialId to be an explicit nil

### UnsetRuntimeCredentialId
`func (o *ManagedDeploymentResponse) UnsetRuntimeCredentialId()`

UnsetRuntimeCredentialId ensures that no value is present for RuntimeCredentialId, not even an explicit nil
### GetRuntimePrincipalType

`func (o *ManagedDeploymentResponse) GetRuntimePrincipalType() PrincipalType`

GetRuntimePrincipalType returns the RuntimePrincipalType field if non-nil, zero value otherwise.

### GetRuntimePrincipalTypeOk

`func (o *ManagedDeploymentResponse) GetRuntimePrincipalTypeOk() (*PrincipalType, bool)`

GetRuntimePrincipalTypeOk returns a tuple with the RuntimePrincipalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimePrincipalType

`func (o *ManagedDeploymentResponse) SetRuntimePrincipalType(v PrincipalType)`

SetRuntimePrincipalType sets RuntimePrincipalType field to given value.

### HasRuntimePrincipalType

`func (o *ManagedDeploymentResponse) HasRuntimePrincipalType() bool`

HasRuntimePrincipalType returns a boolean if a field has been set.

### SetRuntimePrincipalTypeNil

`func (o *ManagedDeploymentResponse) SetRuntimePrincipalTypeNil(b bool)`

 SetRuntimePrincipalTypeNil sets the value for RuntimePrincipalType to be an explicit nil

### UnsetRuntimePrincipalType
`func (o *ManagedDeploymentResponse) UnsetRuntimePrincipalType()`

UnsetRuntimePrincipalType ensures that no value is present for RuntimePrincipalType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


