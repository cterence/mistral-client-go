# DeploymentListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deployments** | [**[]DeploymentResponse**](DeploymentResponse.md) | List of deployments | 

## Methods

### NewDeploymentListResponse

`func NewDeploymentListResponse(deployments []DeploymentResponse, ) *DeploymentListResponse`

NewDeploymentListResponse instantiates a new DeploymentListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentListResponseWithDefaults

`func NewDeploymentListResponseWithDefaults() *DeploymentListResponse`

NewDeploymentListResponseWithDefaults instantiates a new DeploymentListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeployments

`func (o *DeploymentListResponse) GetDeployments() []DeploymentResponse`

GetDeployments returns the Deployments field if non-nil, zero value otherwise.

### GetDeploymentsOk

`func (o *DeploymentListResponse) GetDeploymentsOk() (*[]DeploymentResponse, bool)`

GetDeploymentsOk returns a tuple with the Deployments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployments

`func (o *DeploymentListResponse) SetDeployments(v []DeploymentResponse)`

SetDeployments sets Deployments field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


