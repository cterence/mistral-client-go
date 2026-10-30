# UpdateDeploymentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spec** | Pointer to [**NullableWorkflowsWorkerSpecUpdate**](WorkflowsWorkerSpecUpdate.md) |  | [optional] 
**Resources** | Pointer to [**NullableDeploymentResourceConfigUpdate**](DeploymentResourceConfigUpdate.md) |  | [optional] 

## Methods

### NewUpdateDeploymentRequest

`func NewUpdateDeploymentRequest() *UpdateDeploymentRequest`

NewUpdateDeploymentRequest instantiates a new UpdateDeploymentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDeploymentRequestWithDefaults

`func NewUpdateDeploymentRequestWithDefaults() *UpdateDeploymentRequest`

NewUpdateDeploymentRequestWithDefaults instantiates a new UpdateDeploymentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpec

`func (o *UpdateDeploymentRequest) GetSpec() WorkflowsWorkerSpecUpdate`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *UpdateDeploymentRequest) GetSpecOk() (*WorkflowsWorkerSpecUpdate, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *UpdateDeploymentRequest) SetSpec(v WorkflowsWorkerSpecUpdate)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *UpdateDeploymentRequest) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### SetSpecNil

`func (o *UpdateDeploymentRequest) SetSpecNil(b bool)`

 SetSpecNil sets the value for Spec to be an explicit nil

### UnsetSpec
`func (o *UpdateDeploymentRequest) UnsetSpec()`

UnsetSpec ensures that no value is present for Spec, not even an explicit nil
### GetResources

`func (o *UpdateDeploymentRequest) GetResources() DeploymentResourceConfigUpdate`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *UpdateDeploymentRequest) GetResourcesOk() (*DeploymentResourceConfigUpdate, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *UpdateDeploymentRequest) SetResources(v DeploymentResourceConfigUpdate)`

SetResources sets Resources field to given value.

### HasResources

`func (o *UpdateDeploymentRequest) HasResources() bool`

HasResources returns a boolean if a field has been set.

### SetResourcesNil

`func (o *UpdateDeploymentRequest) SetResourcesNil(b bool)`

 SetResourcesNil sets the value for Resources to be an explicit nil

### UnsetResources
`func (o *UpdateDeploymentRequest) UnsetResources()`

UnsetResources ensures that no value is present for Resources, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


