# DeploymentMistralCloudBackendSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "mistral_cloud"]
**BuildDirectory** | Pointer to **NullableString** | Docker build context, as a path in the repo. Defaults to the repo root. | [optional] 
**DockerfilePath** | Pointer to **NullableString** | Path to the Dockerfile, relative to &#39;build_directory&#39;. Defaults to &#39;Dockerfile&#39;. | [optional] 
**Secrets** | Pointer to [**[]DeploymentSecretBinding**](DeploymentSecretBinding.md) | Secrets to bind to the worker, as &#39;secret:workspace:&lt;NAME&gt;&#39; references. Each is available during build and run: as a Dockerfile build arg (declare a matching &#39;ARG&#39;) and as an env var. On update, the field carries the complete set: send the full list to rebind, an empty list to unbind all, omit it to leave the bindings unchanged. | [optional] 

## Methods

### NewDeploymentMistralCloudBackendSpec

`func NewDeploymentMistralCloudBackendSpec() *DeploymentMistralCloudBackendSpec`

NewDeploymentMistralCloudBackendSpec instantiates a new DeploymentMistralCloudBackendSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentMistralCloudBackendSpecWithDefaults

`func NewDeploymentMistralCloudBackendSpecWithDefaults() *DeploymentMistralCloudBackendSpec`

NewDeploymentMistralCloudBackendSpecWithDefaults instantiates a new DeploymentMistralCloudBackendSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DeploymentMistralCloudBackendSpec) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DeploymentMistralCloudBackendSpec) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DeploymentMistralCloudBackendSpec) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DeploymentMistralCloudBackendSpec) HasType() bool`

HasType returns a boolean if a field has been set.

### GetBuildDirectory

`func (o *DeploymentMistralCloudBackendSpec) GetBuildDirectory() string`

GetBuildDirectory returns the BuildDirectory field if non-nil, zero value otherwise.

### GetBuildDirectoryOk

`func (o *DeploymentMistralCloudBackendSpec) GetBuildDirectoryOk() (*string, bool)`

GetBuildDirectoryOk returns a tuple with the BuildDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildDirectory

`func (o *DeploymentMistralCloudBackendSpec) SetBuildDirectory(v string)`

SetBuildDirectory sets BuildDirectory field to given value.

### HasBuildDirectory

`func (o *DeploymentMistralCloudBackendSpec) HasBuildDirectory() bool`

HasBuildDirectory returns a boolean if a field has been set.

### SetBuildDirectoryNil

`func (o *DeploymentMistralCloudBackendSpec) SetBuildDirectoryNil(b bool)`

 SetBuildDirectoryNil sets the value for BuildDirectory to be an explicit nil

### UnsetBuildDirectory
`func (o *DeploymentMistralCloudBackendSpec) UnsetBuildDirectory()`

UnsetBuildDirectory ensures that no value is present for BuildDirectory, not even an explicit nil
### GetDockerfilePath

`func (o *DeploymentMistralCloudBackendSpec) GetDockerfilePath() string`

GetDockerfilePath returns the DockerfilePath field if non-nil, zero value otherwise.

### GetDockerfilePathOk

`func (o *DeploymentMistralCloudBackendSpec) GetDockerfilePathOk() (*string, bool)`

GetDockerfilePathOk returns a tuple with the DockerfilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerfilePath

`func (o *DeploymentMistralCloudBackendSpec) SetDockerfilePath(v string)`

SetDockerfilePath sets DockerfilePath field to given value.

### HasDockerfilePath

`func (o *DeploymentMistralCloudBackendSpec) HasDockerfilePath() bool`

HasDockerfilePath returns a boolean if a field has been set.

### SetDockerfilePathNil

`func (o *DeploymentMistralCloudBackendSpec) SetDockerfilePathNil(b bool)`

 SetDockerfilePathNil sets the value for DockerfilePath to be an explicit nil

### UnsetDockerfilePath
`func (o *DeploymentMistralCloudBackendSpec) UnsetDockerfilePath()`

UnsetDockerfilePath ensures that no value is present for DockerfilePath, not even an explicit nil
### GetSecrets

`func (o *DeploymentMistralCloudBackendSpec) GetSecrets() []DeploymentSecretBinding`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *DeploymentMistralCloudBackendSpec) GetSecretsOk() (*[]DeploymentSecretBinding, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *DeploymentMistralCloudBackendSpec) SetSecrets(v []DeploymentSecretBinding)`

SetSecrets sets Secrets field to given value.

### HasSecrets

`func (o *DeploymentMistralCloudBackendSpec) HasSecrets() bool`

HasSecrets returns a boolean if a field has been set.

### SetSecretsNil

`func (o *DeploymentMistralCloudBackendSpec) SetSecretsNil(b bool)`

 SetSecretsNil sets the value for Secrets to be an explicit nil

### UnsetSecrets
`func (o *DeploymentMistralCloudBackendSpec) UnsetSecrets()`

UnsetSecrets ensures that no value is present for Secrets, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


