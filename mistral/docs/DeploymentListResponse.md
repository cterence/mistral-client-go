# DeploymentListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deployments** | [**[]DeploymentResponse**](DeploymentResponse.md) | List of deployments | 
**NextCursor** | **NullableString** | Cursor for the next page of results | 
**WorkspaceId** | **string** | Workspace ID the results are scoped to | 

## Methods

### NewDeploymentListResponse

`func NewDeploymentListResponse(deployments []DeploymentResponse, nextCursor NullableString, workspaceId string, ) *DeploymentListResponse`

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


### GetNextCursor

`func (o *DeploymentListResponse) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *DeploymentListResponse) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *DeploymentListResponse) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.


### SetNextCursorNil

`func (o *DeploymentListResponse) SetNextCursorNil(b bool)`

 SetNextCursorNil sets the value for NextCursor to be an explicit nil

### UnsetNextCursor
`func (o *DeploymentListResponse) UnsetNextCursor()`

UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil
### GetWorkspaceId

`func (o *DeploymentListResponse) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *DeploymentListResponse) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *DeploymentListResponse) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


