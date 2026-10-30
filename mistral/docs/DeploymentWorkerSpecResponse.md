# DeploymentWorkerSpecResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GithubUrl** | **string** |  | 
**Type** | Pointer to **string** |  | [optional] [default to "workflows_worker"]
**Revision** | Pointer to **NullableString** |  | [optional] 
**BackendSpec** | [**BackendSpec**](BackendSpec.md) |  | 
**RestartedAt** | Pointer to **NullableString** |  | [optional] 
**Commit** | Pointer to [**NullableGitCommitMetadata**](GitCommitMetadata.md) |  | [optional] 

## Methods

### NewDeploymentWorkerSpecResponse

`func NewDeploymentWorkerSpecResponse(githubUrl string, backendSpec BackendSpec, ) *DeploymentWorkerSpecResponse`

NewDeploymentWorkerSpecResponse instantiates a new DeploymentWorkerSpecResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentWorkerSpecResponseWithDefaults

`func NewDeploymentWorkerSpecResponseWithDefaults() *DeploymentWorkerSpecResponse`

NewDeploymentWorkerSpecResponseWithDefaults instantiates a new DeploymentWorkerSpecResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGithubUrl

`func (o *DeploymentWorkerSpecResponse) GetGithubUrl() string`

GetGithubUrl returns the GithubUrl field if non-nil, zero value otherwise.

### GetGithubUrlOk

`func (o *DeploymentWorkerSpecResponse) GetGithubUrlOk() (*string, bool)`

GetGithubUrlOk returns a tuple with the GithubUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubUrl

`func (o *DeploymentWorkerSpecResponse) SetGithubUrl(v string)`

SetGithubUrl sets GithubUrl field to given value.


### GetType

`func (o *DeploymentWorkerSpecResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DeploymentWorkerSpecResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DeploymentWorkerSpecResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DeploymentWorkerSpecResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRevision

`func (o *DeploymentWorkerSpecResponse) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *DeploymentWorkerSpecResponse) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *DeploymentWorkerSpecResponse) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *DeploymentWorkerSpecResponse) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### SetRevisionNil

`func (o *DeploymentWorkerSpecResponse) SetRevisionNil(b bool)`

 SetRevisionNil sets the value for Revision to be an explicit nil

### UnsetRevision
`func (o *DeploymentWorkerSpecResponse) UnsetRevision()`

UnsetRevision ensures that no value is present for Revision, not even an explicit nil
### GetBackendSpec

`func (o *DeploymentWorkerSpecResponse) GetBackendSpec() BackendSpec`

GetBackendSpec returns the BackendSpec field if non-nil, zero value otherwise.

### GetBackendSpecOk

`func (o *DeploymentWorkerSpecResponse) GetBackendSpecOk() (*BackendSpec, bool)`

GetBackendSpecOk returns a tuple with the BackendSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendSpec

`func (o *DeploymentWorkerSpecResponse) SetBackendSpec(v BackendSpec)`

SetBackendSpec sets BackendSpec field to given value.


### GetRestartedAt

`func (o *DeploymentWorkerSpecResponse) GetRestartedAt() string`

GetRestartedAt returns the RestartedAt field if non-nil, zero value otherwise.

### GetRestartedAtOk

`func (o *DeploymentWorkerSpecResponse) GetRestartedAtOk() (*string, bool)`

GetRestartedAtOk returns a tuple with the RestartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartedAt

`func (o *DeploymentWorkerSpecResponse) SetRestartedAt(v string)`

SetRestartedAt sets RestartedAt field to given value.

### HasRestartedAt

`func (o *DeploymentWorkerSpecResponse) HasRestartedAt() bool`

HasRestartedAt returns a boolean if a field has been set.

### SetRestartedAtNil

`func (o *DeploymentWorkerSpecResponse) SetRestartedAtNil(b bool)`

 SetRestartedAtNil sets the value for RestartedAt to be an explicit nil

### UnsetRestartedAt
`func (o *DeploymentWorkerSpecResponse) UnsetRestartedAt()`

UnsetRestartedAt ensures that no value is present for RestartedAt, not even an explicit nil
### GetCommit

`func (o *DeploymentWorkerSpecResponse) GetCommit() GitCommitMetadata`

GetCommit returns the Commit field if non-nil, zero value otherwise.

### GetCommitOk

`func (o *DeploymentWorkerSpecResponse) GetCommitOk() (*GitCommitMetadata, bool)`

GetCommitOk returns a tuple with the Commit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommit

`func (o *DeploymentWorkerSpecResponse) SetCommit(v GitCommitMetadata)`

SetCommit sets Commit field to given value.

### HasCommit

`func (o *DeploymentWorkerSpecResponse) HasCommit() bool`

HasCommit returns a boolean if a field has been set.

### SetCommitNil

`func (o *DeploymentWorkerSpecResponse) SetCommitNil(b bool)`

 SetCommitNil sets the value for Commit to be an explicit nil

### UnsetCommit
`func (o *DeploymentWorkerSpecResponse) UnsetCommit()`

UnsetCommit ensures that no value is present for Commit, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


