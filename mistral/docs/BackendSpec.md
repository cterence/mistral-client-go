# BackendSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "mistral_cloud"]
**BuildDirectory** | Pointer to **string** | Docker build context, as a path in the repo. Defaults to the repo root. | [optional] 
**DockerfilePath** | Pointer to **string** | Path to the Dockerfile, relative to &#39;build_directory&#39;. Defaults to &#39;Dockerfile&#39;. | [optional] 
**Secrets** | Pointer to [**[]DeploymentSecretBinding**](DeploymentSecretBinding.md) | Secrets to bind to the worker, as &#39;secret:workspace:&lt;NAME&gt;&#39; references. Each is available during build and run: as a Dockerfile build arg (declare a matching &#39;ARG&#39;) and as an env var. On update, the field carries the complete set: send the full list to rebind, an empty list to unbind all, omit it to leave the bindings unchanged. | [optional] 
**Entrypoint** | Pointer to **string** | Either a &#39;module:function&#39; reference (e.g. &#39;worker:main&#39;) or a &#39;.py&#39; script path relative to &#39;working_dir&#39;. | [optional] 
**WorkingDir** | Pointer to **string** | Path in the repo holding the worker&#39;s pyproject.toml, for monorepo layouts. Leave empty for single-package repos. | [optional] 

## Methods

### NewBackendSpec

`func NewBackendSpec() *BackendSpec`

NewBackendSpec instantiates a new BackendSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackendSpecWithDefaults

`func NewBackendSpecWithDefaults() *BackendSpec`

NewBackendSpecWithDefaults instantiates a new BackendSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BackendSpec) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BackendSpec) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BackendSpec) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BackendSpec) HasType() bool`

HasType returns a boolean if a field has been set.

### GetBuildDirectory

`func (o *BackendSpec) GetBuildDirectory() string`

GetBuildDirectory returns the BuildDirectory field if non-nil, zero value otherwise.

### GetBuildDirectoryOk

`func (o *BackendSpec) GetBuildDirectoryOk() (*string, bool)`

GetBuildDirectoryOk returns a tuple with the BuildDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildDirectory

`func (o *BackendSpec) SetBuildDirectory(v string)`

SetBuildDirectory sets BuildDirectory field to given value.

### HasBuildDirectory

`func (o *BackendSpec) HasBuildDirectory() bool`

HasBuildDirectory returns a boolean if a field has been set.

### GetDockerfilePath

`func (o *BackendSpec) GetDockerfilePath() string`

GetDockerfilePath returns the DockerfilePath field if non-nil, zero value otherwise.

### GetDockerfilePathOk

`func (o *BackendSpec) GetDockerfilePathOk() (*string, bool)`

GetDockerfilePathOk returns a tuple with the DockerfilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerfilePath

`func (o *BackendSpec) SetDockerfilePath(v string)`

SetDockerfilePath sets DockerfilePath field to given value.

### HasDockerfilePath

`func (o *BackendSpec) HasDockerfilePath() bool`

HasDockerfilePath returns a boolean if a field has been set.

### GetSecrets

`func (o *BackendSpec) GetSecrets() []DeploymentSecretBinding`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *BackendSpec) GetSecretsOk() (*[]DeploymentSecretBinding, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *BackendSpec) SetSecrets(v []DeploymentSecretBinding)`

SetSecrets sets Secrets field to given value.

### HasSecrets

`func (o *BackendSpec) HasSecrets() bool`

HasSecrets returns a boolean if a field has been set.

### GetEntrypoint

`func (o *BackendSpec) GetEntrypoint() string`

GetEntrypoint returns the Entrypoint field if non-nil, zero value otherwise.

### GetEntrypointOk

`func (o *BackendSpec) GetEntrypointOk() (*string, bool)`

GetEntrypointOk returns a tuple with the Entrypoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntrypoint

`func (o *BackendSpec) SetEntrypoint(v string)`

SetEntrypoint sets Entrypoint field to given value.

### HasEntrypoint

`func (o *BackendSpec) HasEntrypoint() bool`

HasEntrypoint returns a boolean if a field has been set.

### GetWorkingDir

`func (o *BackendSpec) GetWorkingDir() string`

GetWorkingDir returns the WorkingDir field if non-nil, zero value otherwise.

### GetWorkingDirOk

`func (o *BackendSpec) GetWorkingDirOk() (*string, bool)`

GetWorkingDirOk returns a tuple with the WorkingDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingDir

`func (o *BackendSpec) SetWorkingDir(v string)`

SetWorkingDir sets WorkingDir field to given value.

### HasWorkingDir

`func (o *BackendSpec) HasWorkingDir() bool`

HasWorkingDir returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


