# AdminUserGroupMemberOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserUuid** | **string** | User ID of the group member. | 
**Name** | **NullableString** | Name of the group member. | 
**Email** | **NullableString** | Email address of the group member. | 

## Methods

### NewAdminUserGroupMemberOut

`func NewAdminUserGroupMemberOut(userUuid string, name NullableString, email NullableString, ) *AdminUserGroupMemberOut`

NewAdminUserGroupMemberOut instantiates a new AdminUserGroupMemberOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminUserGroupMemberOutWithDefaults

`func NewAdminUserGroupMemberOutWithDefaults() *AdminUserGroupMemberOut`

NewAdminUserGroupMemberOutWithDefaults instantiates a new AdminUserGroupMemberOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserUuid

`func (o *AdminUserGroupMemberOut) GetUserUuid() string`

GetUserUuid returns the UserUuid field if non-nil, zero value otherwise.

### GetUserUuidOk

`func (o *AdminUserGroupMemberOut) GetUserUuidOk() (*string, bool)`

GetUserUuidOk returns a tuple with the UserUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserUuid

`func (o *AdminUserGroupMemberOut) SetUserUuid(v string)`

SetUserUuid sets UserUuid field to given value.


### GetName

`func (o *AdminUserGroupMemberOut) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdminUserGroupMemberOut) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdminUserGroupMemberOut) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *AdminUserGroupMemberOut) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AdminUserGroupMemberOut) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetEmail

`func (o *AdminUserGroupMemberOut) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AdminUserGroupMemberOut) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AdminUserGroupMemberOut) SetEmail(v string)`

SetEmail sets Email field to given value.


### SetEmailNil

`func (o *AdminUserGroupMemberOut) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *AdminUserGroupMemberOut) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


