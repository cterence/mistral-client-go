# CreateDeploymentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Spec** | [**DeploymentWorkerSpecInput**](DeploymentWorkerSpecInput.md) |  | 
**Resources** | Pointer to [**NullableDeploymentResourceConfig**](DeploymentResourceConfig.md) |  | [optional] 
**Hardened** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewCreateDeploymentRequest

`func NewCreateDeploymentRequest(name string, spec DeploymentWorkerSpecInput, ) *CreateDeploymentRequest`

NewCreateDeploymentRequest instantiates a new CreateDeploymentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDeploymentRequestWithDefaults

`func NewCreateDeploymentRequestWithDefaults() *CreateDeploymentRequest`

NewCreateDeploymentRequestWithDefaults instantiates a new CreateDeploymentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateDeploymentRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateDeploymentRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateDeploymentRequest) SetName(v string)`

SetName sets Name field to given value.


### GetSpec

`func (o *CreateDeploymentRequest) GetSpec() DeploymentWorkerSpecInput`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *CreateDeploymentRequest) GetSpecOk() (*DeploymentWorkerSpecInput, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *CreateDeploymentRequest) SetSpec(v DeploymentWorkerSpecInput)`

SetSpec sets Spec field to given value.


### GetResources

`func (o *CreateDeploymentRequest) GetResources() DeploymentResourceConfig`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *CreateDeploymentRequest) GetResourcesOk() (*DeploymentResourceConfig, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *CreateDeploymentRequest) SetResources(v DeploymentResourceConfig)`

SetResources sets Resources field to given value.

### HasResources

`func (o *CreateDeploymentRequest) HasResources() bool`

HasResources returns a boolean if a field has been set.

### SetResourcesNil

`func (o *CreateDeploymentRequest) SetResourcesNil(b bool)`

 SetResourcesNil sets the value for Resources to be an explicit nil

### UnsetResources
`func (o *CreateDeploymentRequest) UnsetResources()`

UnsetResources ensures that no value is present for Resources, not even an explicit nil
### GetHardened

`func (o *CreateDeploymentRequest) GetHardened() bool`

GetHardened returns the Hardened field if non-nil, zero value otherwise.

### GetHardenedOk

`func (o *CreateDeploymentRequest) GetHardenedOk() (*bool, bool)`

GetHardenedOk returns a tuple with the Hardened field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardened

`func (o *CreateDeploymentRequest) SetHardened(v bool)`

SetHardened sets Hardened field to given value.

### HasHardened

`func (o *CreateDeploymentRequest) HasHardened() bool`

HasHardened returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


