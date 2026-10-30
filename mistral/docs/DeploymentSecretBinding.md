# DeploymentSecretBinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvVarName** | **string** | Environment variable the worker process reads; the secret&#39;s name lives in &#39;reference&#39;. | 
**Reference** | **string** | Secret reference as &#39;secret:workspace:&lt;NAME&gt;&#39;. Resolved at boot by the runtime-init container; the value never reaches the control plane. | 

## Methods

### NewDeploymentSecretBinding

`func NewDeploymentSecretBinding(envVarName string, reference string, ) *DeploymentSecretBinding`

NewDeploymentSecretBinding instantiates a new DeploymentSecretBinding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentSecretBindingWithDefaults

`func NewDeploymentSecretBindingWithDefaults() *DeploymentSecretBinding`

NewDeploymentSecretBindingWithDefaults instantiates a new DeploymentSecretBinding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvVarName

`func (o *DeploymentSecretBinding) GetEnvVarName() string`

GetEnvVarName returns the EnvVarName field if non-nil, zero value otherwise.

### GetEnvVarNameOk

`func (o *DeploymentSecretBinding) GetEnvVarNameOk() (*string, bool)`

GetEnvVarNameOk returns a tuple with the EnvVarName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvVarName

`func (o *DeploymentSecretBinding) SetEnvVarName(v string)`

SetEnvVarName sets EnvVarName field to given value.


### GetReference

`func (o *DeploymentSecretBinding) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *DeploymentSecretBinding) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *DeploymentSecretBinding) SetReference(v string)`

SetReference sets Reference field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


