# DeploymentWorkerSpecInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GithubUrl** | **string** |  | 
**Revision** | Pointer to **NullableString** |  | [optional] 
**BackendSpec** | Pointer to [**NullableDeploymentMistralCloudBackendSpec**](DeploymentMistralCloudBackendSpec.md) | Backend-specific configuration for the Mistral Cloud backend. | [optional] 

## Methods

### NewDeploymentWorkerSpecInput

`func NewDeploymentWorkerSpecInput(githubUrl string, ) *DeploymentWorkerSpecInput`

NewDeploymentWorkerSpecInput instantiates a new DeploymentWorkerSpecInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentWorkerSpecInputWithDefaults

`func NewDeploymentWorkerSpecInputWithDefaults() *DeploymentWorkerSpecInput`

NewDeploymentWorkerSpecInputWithDefaults instantiates a new DeploymentWorkerSpecInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGithubUrl

`func (o *DeploymentWorkerSpecInput) GetGithubUrl() string`

GetGithubUrl returns the GithubUrl field if non-nil, zero value otherwise.

### GetGithubUrlOk

`func (o *DeploymentWorkerSpecInput) GetGithubUrlOk() (*string, bool)`

GetGithubUrlOk returns a tuple with the GithubUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubUrl

`func (o *DeploymentWorkerSpecInput) SetGithubUrl(v string)`

SetGithubUrl sets GithubUrl field to given value.


### GetRevision

`func (o *DeploymentWorkerSpecInput) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *DeploymentWorkerSpecInput) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *DeploymentWorkerSpecInput) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *DeploymentWorkerSpecInput) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### SetRevisionNil

`func (o *DeploymentWorkerSpecInput) SetRevisionNil(b bool)`

 SetRevisionNil sets the value for Revision to be an explicit nil

### UnsetRevision
`func (o *DeploymentWorkerSpecInput) UnsetRevision()`

UnsetRevision ensures that no value is present for Revision, not even an explicit nil
### GetBackendSpec

`func (o *DeploymentWorkerSpecInput) GetBackendSpec() DeploymentMistralCloudBackendSpec`

GetBackendSpec returns the BackendSpec field if non-nil, zero value otherwise.

### GetBackendSpecOk

`func (o *DeploymentWorkerSpecInput) GetBackendSpecOk() (*DeploymentMistralCloudBackendSpec, bool)`

GetBackendSpecOk returns a tuple with the BackendSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendSpec

`func (o *DeploymentWorkerSpecInput) SetBackendSpec(v DeploymentMistralCloudBackendSpec)`

SetBackendSpec sets BackendSpec field to given value.

### HasBackendSpec

`func (o *DeploymentWorkerSpecInput) HasBackendSpec() bool`

HasBackendSpec returns a boolean if a field has been set.

### SetBackendSpecNil

`func (o *DeploymentWorkerSpecInput) SetBackendSpecNil(b bool)`

 SetBackendSpecNil sets the value for BackendSpec to be an explicit nil

### UnsetBackendSpec
`func (o *DeploymentWorkerSpecInput) UnsetBackendSpec()`

UnsetBackendSpec ensures that no value is present for BackendSpec, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


