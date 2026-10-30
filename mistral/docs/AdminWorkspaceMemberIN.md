# AdminWorkspaceMemberIN

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Members** | Pointer to [**[]AdminWorkspaceMemberSingleIN**](AdminWorkspaceMemberSingleIN.md) | Workspace members to add or update. | [optional] 

## Methods

### NewAdminWorkspaceMemberIN

`func NewAdminWorkspaceMemberIN() *AdminWorkspaceMemberIN`

NewAdminWorkspaceMemberIN instantiates a new AdminWorkspaceMemberIN object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminWorkspaceMemberINWithDefaults

`func NewAdminWorkspaceMemberINWithDefaults() *AdminWorkspaceMemberIN`

NewAdminWorkspaceMemberINWithDefaults instantiates a new AdminWorkspaceMemberIN object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMembers

`func (o *AdminWorkspaceMemberIN) GetMembers() []AdminWorkspaceMemberSingleIN`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *AdminWorkspaceMemberIN) GetMembersOk() (*[]AdminWorkspaceMemberSingleIN, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *AdminWorkspaceMemberIN) SetMembers(v []AdminWorkspaceMemberSingleIN)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *AdminWorkspaceMemberIN) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### SetMembersNil

`func (o *AdminWorkspaceMemberIN) SetMembersNil(b bool)`

 SetMembersNil sets the value for Members to be an explicit nil

### UnsetMembers
`func (o *AdminWorkspaceMemberIN) UnsetMembers()`

UnsetMembers ensures that no value is present for Members, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


