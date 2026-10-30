# WorkflowsWorkerSpecUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GithubUrl** | Pointer to **NullableString** |  | [optional] 
**Revision** | Pointer to **NullableString** |  | [optional] 
**BackendSpec** | Pointer to [**NullableDeploymentMistralCloudBackendSpec**](DeploymentMistralCloudBackendSpec.md) | Backend-specific configuration for the Mistral Cloud backend. | [optional] 

## Methods

### NewWorkflowsWorkerSpecUpdate

`func NewWorkflowsWorkerSpecUpdate() *WorkflowsWorkerSpecUpdate`

NewWorkflowsWorkerSpecUpdate instantiates a new WorkflowsWorkerSpecUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowsWorkerSpecUpdateWithDefaults

`func NewWorkflowsWorkerSpecUpdateWithDefaults() *WorkflowsWorkerSpecUpdate`

NewWorkflowsWorkerSpecUpdateWithDefaults instantiates a new WorkflowsWorkerSpecUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGithubUrl

`func (o *WorkflowsWorkerSpecUpdate) GetGithubUrl() string`

GetGithubUrl returns the GithubUrl field if non-nil, zero value otherwise.

### GetGithubUrlOk

`func (o *WorkflowsWorkerSpecUpdate) GetGithubUrlOk() (*string, bool)`

GetGithubUrlOk returns a tuple with the GithubUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubUrl

`func (o *WorkflowsWorkerSpecUpdate) SetGithubUrl(v string)`

SetGithubUrl sets GithubUrl field to given value.

### HasGithubUrl

`func (o *WorkflowsWorkerSpecUpdate) HasGithubUrl() bool`

HasGithubUrl returns a boolean if a field has been set.

### SetGithubUrlNil

`func (o *WorkflowsWorkerSpecUpdate) SetGithubUrlNil(b bool)`

 SetGithubUrlNil sets the value for GithubUrl to be an explicit nil

### UnsetGithubUrl
`func (o *WorkflowsWorkerSpecUpdate) UnsetGithubUrl()`

UnsetGithubUrl ensures that no value is present for GithubUrl, not even an explicit nil
### GetRevision

`func (o *WorkflowsWorkerSpecUpdate) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *WorkflowsWorkerSpecUpdate) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *WorkflowsWorkerSpecUpdate) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *WorkflowsWorkerSpecUpdate) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### SetRevisionNil

`func (o *WorkflowsWorkerSpecUpdate) SetRevisionNil(b bool)`

 SetRevisionNil sets the value for Revision to be an explicit nil

### UnsetRevision
`func (o *WorkflowsWorkerSpecUpdate) UnsetRevision()`

UnsetRevision ensures that no value is present for Revision, not even an explicit nil
### GetBackendSpec

`func (o *WorkflowsWorkerSpecUpdate) GetBackendSpec() DeploymentMistralCloudBackendSpec`

GetBackendSpec returns the BackendSpec field if non-nil, zero value otherwise.

### GetBackendSpecOk

`func (o *WorkflowsWorkerSpecUpdate) GetBackendSpecOk() (*DeploymentMistralCloudBackendSpec, bool)`

GetBackendSpecOk returns a tuple with the BackendSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendSpec

`func (o *WorkflowsWorkerSpecUpdate) SetBackendSpec(v DeploymentMistralCloudBackendSpec)`

SetBackendSpec sets BackendSpec field to given value.

### HasBackendSpec

`func (o *WorkflowsWorkerSpecUpdate) HasBackendSpec() bool`

HasBackendSpec returns a boolean if a field has been set.

### SetBackendSpecNil

`func (o *WorkflowsWorkerSpecUpdate) SetBackendSpecNil(b bool)`

 SetBackendSpecNil sets the value for BackendSpec to be an explicit nil

### UnsetBackendSpec
`func (o *WorkflowsWorkerSpecUpdate) UnsetBackendSpec()`

UnsetBackendSpec ensures that no value is present for BackendSpec, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


