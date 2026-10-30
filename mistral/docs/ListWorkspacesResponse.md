# ListWorkspacesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Workspaces** | [**[]UserWorkspace**](UserWorkspace.md) | The workspaces the authenticated user is a member of, each tagged with the organization it belongs to. | 

## Methods

### NewListWorkspacesResponse

`func NewListWorkspacesResponse(workspaces []UserWorkspace, ) *ListWorkspacesResponse`

NewListWorkspacesResponse instantiates a new ListWorkspacesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListWorkspacesResponseWithDefaults

`func NewListWorkspacesResponseWithDefaults() *ListWorkspacesResponse`

NewListWorkspacesResponseWithDefaults instantiates a new ListWorkspacesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkspaces

`func (o *ListWorkspacesResponse) GetWorkspaces() []UserWorkspace`

GetWorkspaces returns the Workspaces field if non-nil, zero value otherwise.

### GetWorkspacesOk

`func (o *ListWorkspacesResponse) GetWorkspacesOk() (*[]UserWorkspace, bool)`

GetWorkspacesOk returns a tuple with the Workspaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaces

`func (o *ListWorkspacesResponse) SetWorkspaces(v []UserWorkspace)`

SetWorkspaces sets Workspaces field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


