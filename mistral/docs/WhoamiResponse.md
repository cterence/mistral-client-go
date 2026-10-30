# WhoamiResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Principal** | [**Principal**](Principal.md) |  | 
**Organization** | Pointer to [**NullableWhoamiOrganization**](WhoamiOrganization.md) | The organization the caller&#39;s session is scoped to, when set. | [optional] 
**Workspace** | Pointer to [**NullableUserIdentityWorkspace**](UserIdentityWorkspace.md) | The workspace the caller&#39;s session is scoped to, when set. | [optional] 

## Methods

### NewWhoamiResponse

`func NewWhoamiResponse(principal Principal, ) *WhoamiResponse`

NewWhoamiResponse instantiates a new WhoamiResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhoamiResponseWithDefaults

`func NewWhoamiResponseWithDefaults() *WhoamiResponse`

NewWhoamiResponseWithDefaults instantiates a new WhoamiResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrincipal

`func (o *WhoamiResponse) GetPrincipal() Principal`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *WhoamiResponse) GetPrincipalOk() (*Principal, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *WhoamiResponse) SetPrincipal(v Principal)`

SetPrincipal sets Principal field to given value.


### GetOrganization

`func (o *WhoamiResponse) GetOrganization() WhoamiOrganization`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *WhoamiResponse) GetOrganizationOk() (*WhoamiOrganization, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *WhoamiResponse) SetOrganization(v WhoamiOrganization)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *WhoamiResponse) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### SetOrganizationNil

`func (o *WhoamiResponse) SetOrganizationNil(b bool)`

 SetOrganizationNil sets the value for Organization to be an explicit nil

### UnsetOrganization
`func (o *WhoamiResponse) UnsetOrganization()`

UnsetOrganization ensures that no value is present for Organization, not even an explicit nil
### GetWorkspace

`func (o *WhoamiResponse) GetWorkspace() UserIdentityWorkspace`

GetWorkspace returns the Workspace field if non-nil, zero value otherwise.

### GetWorkspaceOk

`func (o *WhoamiResponse) GetWorkspaceOk() (*UserIdentityWorkspace, bool)`

GetWorkspaceOk returns a tuple with the Workspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspace

`func (o *WhoamiResponse) SetWorkspace(v UserIdentityWorkspace)`

SetWorkspace sets Workspace field to given value.

### HasWorkspace

`func (o *WhoamiResponse) HasWorkspace() bool`

HasWorkspace returns a boolean if a field has been set.

### SetWorkspaceNil

`func (o *WhoamiResponse) SetWorkspaceNil(b bool)`

 SetWorkspaceNil sets the value for Workspace to be an explicit nil

### UnsetWorkspace
`func (o *WhoamiResponse) UnsetWorkspace()`

UnsetWorkspace ensures that no value is present for Workspace, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


