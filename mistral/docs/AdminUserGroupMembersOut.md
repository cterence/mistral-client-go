# AdminUserGroupMembersOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int32** | Total number of group members that match the request. | 
**Members** | [**[]AdminUserGroupMemberOut**](AdminUserGroupMemberOut.md) | Group members on this page. | 
**Page** | **int32** | Page number returned. | 
**PageSize** | **int32** | Maximum number of results per page. | 

## Methods

### NewAdminUserGroupMembersOut

`func NewAdminUserGroupMembersOut(total int32, members []AdminUserGroupMemberOut, page int32, pageSize int32, ) *AdminUserGroupMembersOut`

NewAdminUserGroupMembersOut instantiates a new AdminUserGroupMembersOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminUserGroupMembersOutWithDefaults

`func NewAdminUserGroupMembersOutWithDefaults() *AdminUserGroupMembersOut`

NewAdminUserGroupMembersOutWithDefaults instantiates a new AdminUserGroupMembersOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *AdminUserGroupMembersOut) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *AdminUserGroupMembersOut) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *AdminUserGroupMembersOut) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetMembers

`func (o *AdminUserGroupMembersOut) GetMembers() []AdminUserGroupMemberOut`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *AdminUserGroupMembersOut) GetMembersOk() (*[]AdminUserGroupMemberOut, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *AdminUserGroupMembersOut) SetMembers(v []AdminUserGroupMemberOut)`

SetMembers sets Members field to given value.


### GetPage

`func (o *AdminUserGroupMembersOut) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *AdminUserGroupMembersOut) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *AdminUserGroupMembersOut) SetPage(v int32)`

SetPage sets Page field to given value.


### GetPageSize

`func (o *AdminUserGroupMembersOut) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *AdminUserGroupMembersOut) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *AdminUserGroupMembersOut) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


