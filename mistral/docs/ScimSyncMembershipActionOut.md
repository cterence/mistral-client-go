# ScimSyncMembershipActionOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | [**ScimSyncGroupRefOut**](ScimSyncGroupRefOut.md) |  | 
**UsersAdded** | [**[]ScimSyncUserRefOut**](ScimSyncUserRefOut.md) |  | 
**UsersRemoved** | [**[]ScimSyncUserRefOut**](ScimSyncUserRefOut.md) |  | 

## Methods

### NewScimSyncMembershipActionOut

`func NewScimSyncMembershipActionOut(group ScimSyncGroupRefOut, usersAdded []ScimSyncUserRefOut, usersRemoved []ScimSyncUserRefOut, ) *ScimSyncMembershipActionOut`

NewScimSyncMembershipActionOut instantiates a new ScimSyncMembershipActionOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScimSyncMembershipActionOutWithDefaults

`func NewScimSyncMembershipActionOutWithDefaults() *ScimSyncMembershipActionOut`

NewScimSyncMembershipActionOutWithDefaults instantiates a new ScimSyncMembershipActionOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *ScimSyncMembershipActionOut) GetGroup() ScimSyncGroupRefOut`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *ScimSyncMembershipActionOut) GetGroupOk() (*ScimSyncGroupRefOut, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *ScimSyncMembershipActionOut) SetGroup(v ScimSyncGroupRefOut)`

SetGroup sets Group field to given value.


### GetUsersAdded

`func (o *ScimSyncMembershipActionOut) GetUsersAdded() []ScimSyncUserRefOut`

GetUsersAdded returns the UsersAdded field if non-nil, zero value otherwise.

### GetUsersAddedOk

`func (o *ScimSyncMembershipActionOut) GetUsersAddedOk() (*[]ScimSyncUserRefOut, bool)`

GetUsersAddedOk returns a tuple with the UsersAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsersAdded

`func (o *ScimSyncMembershipActionOut) SetUsersAdded(v []ScimSyncUserRefOut)`

SetUsersAdded sets UsersAdded field to given value.


### GetUsersRemoved

`func (o *ScimSyncMembershipActionOut) GetUsersRemoved() []ScimSyncUserRefOut`

GetUsersRemoved returns the UsersRemoved field if non-nil, zero value otherwise.

### GetUsersRemovedOk

`func (o *ScimSyncMembershipActionOut) GetUsersRemovedOk() (*[]ScimSyncUserRefOut, bool)`

GetUsersRemovedOk returns a tuple with the UsersRemoved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsersRemoved

`func (o *ScimSyncMembershipActionOut) SetUsersRemoved(v []ScimSyncUserRefOut)`

SetUsersRemoved sets UsersRemoved field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


