# DeploymentK8sBackendSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "kubernetes"]
**Entrypoint** | Pointer to **NullableString** | Either a &#39;module:function&#39; reference (e.g. &#39;worker:main&#39;) or a &#39;.py&#39; script path relative to &#39;working_dir&#39;. | [optional] 
**WorkingDir** | Pointer to **NullableString** | Path in the repo holding the worker&#39;s pyproject.toml, for monorepo layouts. Leave empty for single-package repos. | [optional] 

## Methods

### NewDeploymentK8sBackendSpec

`func NewDeploymentK8sBackendSpec() *DeploymentK8sBackendSpec`

NewDeploymentK8sBackendSpec instantiates a new DeploymentK8sBackendSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentK8sBackendSpecWithDefaults

`func NewDeploymentK8sBackendSpecWithDefaults() *DeploymentK8sBackendSpec`

NewDeploymentK8sBackendSpecWithDefaults instantiates a new DeploymentK8sBackendSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DeploymentK8sBackendSpec) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DeploymentK8sBackendSpec) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DeploymentK8sBackendSpec) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DeploymentK8sBackendSpec) HasType() bool`

HasType returns a boolean if a field has been set.

### GetEntrypoint

`func (o *DeploymentK8sBackendSpec) GetEntrypoint() string`

GetEntrypoint returns the Entrypoint field if non-nil, zero value otherwise.

### GetEntrypointOk

`func (o *DeploymentK8sBackendSpec) GetEntrypointOk() (*string, bool)`

GetEntrypointOk returns a tuple with the Entrypoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntrypoint

`func (o *DeploymentK8sBackendSpec) SetEntrypoint(v string)`

SetEntrypoint sets Entrypoint field to given value.

### HasEntrypoint

`func (o *DeploymentK8sBackendSpec) HasEntrypoint() bool`

HasEntrypoint returns a boolean if a field has been set.

### SetEntrypointNil

`func (o *DeploymentK8sBackendSpec) SetEntrypointNil(b bool)`

 SetEntrypointNil sets the value for Entrypoint to be an explicit nil

### UnsetEntrypoint
`func (o *DeploymentK8sBackendSpec) UnsetEntrypoint()`

UnsetEntrypoint ensures that no value is present for Entrypoint, not even an explicit nil
### GetWorkingDir

`func (o *DeploymentK8sBackendSpec) GetWorkingDir() string`

GetWorkingDir returns the WorkingDir field if non-nil, zero value otherwise.

### GetWorkingDirOk

`func (o *DeploymentK8sBackendSpec) GetWorkingDirOk() (*string, bool)`

GetWorkingDirOk returns a tuple with the WorkingDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingDir

`func (o *DeploymentK8sBackendSpec) SetWorkingDir(v string)`

SetWorkingDir sets WorkingDir field to given value.

### HasWorkingDir

`func (o *DeploymentK8sBackendSpec) HasWorkingDir() bool`

HasWorkingDir returns a boolean if a field has been set.

### SetWorkingDirNil

`func (o *DeploymentK8sBackendSpec) SetWorkingDirNil(b bool)`

 SetWorkingDirNil sets the value for WorkingDir to be an explicit nil

### UnsetWorkingDir
`func (o *DeploymentK8sBackendSpec) UnsetWorkingDir()`

UnsetWorkingDir ensures that no value is present for WorkingDir, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


