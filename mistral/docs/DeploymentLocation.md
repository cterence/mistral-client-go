# DeploymentLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocationType** | [**LocationType**](LocationType.md) | Where the deployment runs: &#39;local&#39;, &#39;k8s&#39;, or &#39;managed&#39; | 
**K8sCluster** | Pointer to **NullableString** | K8s cluster name, if applicable | [optional] 
**K8sNamespace** | Pointer to **NullableString** | K8s namespace, if applicable | [optional] 

## Methods

### NewDeploymentLocation

`func NewDeploymentLocation(locationType LocationType, ) *DeploymentLocation`

NewDeploymentLocation instantiates a new DeploymentLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentLocationWithDefaults

`func NewDeploymentLocationWithDefaults() *DeploymentLocation`

NewDeploymentLocationWithDefaults instantiates a new DeploymentLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocationType

`func (o *DeploymentLocation) GetLocationType() LocationType`

GetLocationType returns the LocationType field if non-nil, zero value otherwise.

### GetLocationTypeOk

`func (o *DeploymentLocation) GetLocationTypeOk() (*LocationType, bool)`

GetLocationTypeOk returns a tuple with the LocationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationType

`func (o *DeploymentLocation) SetLocationType(v LocationType)`

SetLocationType sets LocationType field to given value.


### GetK8sCluster

`func (o *DeploymentLocation) GetK8sCluster() string`

GetK8sCluster returns the K8sCluster field if non-nil, zero value otherwise.

### GetK8sClusterOk

`func (o *DeploymentLocation) GetK8sClusterOk() (*string, bool)`

GetK8sClusterOk returns a tuple with the K8sCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sCluster

`func (o *DeploymentLocation) SetK8sCluster(v string)`

SetK8sCluster sets K8sCluster field to given value.

### HasK8sCluster

`func (o *DeploymentLocation) HasK8sCluster() bool`

HasK8sCluster returns a boolean if a field has been set.

### SetK8sClusterNil

`func (o *DeploymentLocation) SetK8sClusterNil(b bool)`

 SetK8sClusterNil sets the value for K8sCluster to be an explicit nil

### UnsetK8sCluster
`func (o *DeploymentLocation) UnsetK8sCluster()`

UnsetK8sCluster ensures that no value is present for K8sCluster, not even an explicit nil
### GetK8sNamespace

`func (o *DeploymentLocation) GetK8sNamespace() string`

GetK8sNamespace returns the K8sNamespace field if non-nil, zero value otherwise.

### GetK8sNamespaceOk

`func (o *DeploymentLocation) GetK8sNamespaceOk() (*string, bool)`

GetK8sNamespaceOk returns a tuple with the K8sNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sNamespace

`func (o *DeploymentLocation) SetK8sNamespace(v string)`

SetK8sNamespace sets K8sNamespace field to given value.

### HasK8sNamespace

`func (o *DeploymentLocation) HasK8sNamespace() bool`

HasK8sNamespace returns a boolean if a field has been set.

### SetK8sNamespaceNil

`func (o *DeploymentLocation) SetK8sNamespaceNil(b bool)`

 SetK8sNamespaceNil sets the value for K8sNamespace to be an explicit nil

### UnsetK8sNamespace
`func (o *DeploymentLocation) UnsetK8sNamespace()`

UnsetK8sNamespace ensures that no value is present for K8sNamespace, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


