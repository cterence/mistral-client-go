# RemoveWorkspaceMembersOUT

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeletedMembersCount** | **int32** | Number of Workspace members removed. | 
**NotDeletedMembers** | Pointer to **[]string** | Users that could not be removed. | [optional] 

## Methods

### NewRemoveWorkspaceMembersOUT

`func NewRemoveWorkspaceMembersOUT(deletedMembersCount int32, ) *RemoveWorkspaceMembersOUT`

NewRemoveWorkspaceMembersOUT instantiates a new RemoveWorkspaceMembersOUT object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoveWorkspaceMembersOUTWithDefaults

`func NewRemoveWorkspaceMembersOUTWithDefaults() *RemoveWorkspaceMembersOUT`

NewRemoveWorkspaceMembersOUTWithDefaults instantiates a new RemoveWorkspaceMembersOUT object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeletedMembersCount

`func (o *RemoveWorkspaceMembersOUT) GetDeletedMembersCount() int32`

GetDeletedMembersCount returns the DeletedMembersCount field if non-nil, zero value otherwise.

### GetDeletedMembersCountOk

`func (o *RemoveWorkspaceMembersOUT) GetDeletedMembersCountOk() (*int32, bool)`

GetDeletedMembersCountOk returns a tuple with the DeletedMembersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedMembersCount

`func (o *RemoveWorkspaceMembersOUT) SetDeletedMembersCount(v int32)`

SetDeletedMembersCount sets DeletedMembersCount field to given value.


### GetNotDeletedMembers

`func (o *RemoveWorkspaceMembersOUT) GetNotDeletedMembers() []string`

GetNotDeletedMembers returns the NotDeletedMembers field if non-nil, zero value otherwise.

### GetNotDeletedMembersOk

`func (o *RemoveWorkspaceMembersOUT) GetNotDeletedMembersOk() (*[]string, bool)`

GetNotDeletedMembersOk returns a tuple with the NotDeletedMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotDeletedMembers

`func (o *RemoveWorkspaceMembersOUT) SetNotDeletedMembers(v []string)`

SetNotDeletedMembers sets NotDeletedMembers field to given value.

### HasNotDeletedMembers

`func (o *RemoveWorkspaceMembersOUT) HasNotDeletedMembers() bool`

HasNotDeletedMembers returns a boolean if a field has been set.

### SetNotDeletedMembersNil

`func (o *RemoveWorkspaceMembersOUT) SetNotDeletedMembersNil(b bool)`

 SetNotDeletedMembersNil sets the value for NotDeletedMembers to be an explicit nil

### UnsetNotDeletedMembers
`func (o *RemoveWorkspaceMembersOUT) UnsetNotDeletedMembers()`

UnsetNotDeletedMembers ensures that no value is present for NotDeletedMembers, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


