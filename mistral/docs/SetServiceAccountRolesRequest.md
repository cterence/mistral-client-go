# SetServiceAccountRolesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoleIds** | **[]string** | Workspace role UUIDs to assign. Required and non-empty; omitted or empty lists return 422. Select member explicitly for no role-granted permissions; workspace-wide sharing may still grant access. | 

## Methods

### NewSetServiceAccountRolesRequest

`func NewSetServiceAccountRolesRequest(roleIds []string, ) *SetServiceAccountRolesRequest`

NewSetServiceAccountRolesRequest instantiates a new SetServiceAccountRolesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetServiceAccountRolesRequestWithDefaults

`func NewSetServiceAccountRolesRequestWithDefaults() *SetServiceAccountRolesRequest`

NewSetServiceAccountRolesRequestWithDefaults instantiates a new SetServiceAccountRolesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoleIds

`func (o *SetServiceAccountRolesRequest) GetRoleIds() []string`

GetRoleIds returns the RoleIds field if non-nil, zero value otherwise.

### GetRoleIdsOk

`func (o *SetServiceAccountRolesRequest) GetRoleIdsOk() (*[]string, bool)`

GetRoleIdsOk returns a tuple with the RoleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleIds

`func (o *SetServiceAccountRolesRequest) SetRoleIds(v []string)`

SetRoleIds sets RoleIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


