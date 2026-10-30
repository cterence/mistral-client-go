# AdminUserGroupsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int32** | Total number of user groups that match the request. | 
**Items** | [**[]AdminUserGroupOut**](AdminUserGroupOut.md) | User groups on this page. | 
**Page** | **int32** | Page number returned. | 
**PageSize** | **int32** | Maximum number of results per page. | 

## Methods

### NewAdminUserGroupsOut

`func NewAdminUserGroupsOut(total int32, items []AdminUserGroupOut, page int32, pageSize int32, ) *AdminUserGroupsOut`

NewAdminUserGroupsOut instantiates a new AdminUserGroupsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminUserGroupsOutWithDefaults

`func NewAdminUserGroupsOutWithDefaults() *AdminUserGroupsOut`

NewAdminUserGroupsOutWithDefaults instantiates a new AdminUserGroupsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *AdminUserGroupsOut) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *AdminUserGroupsOut) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *AdminUserGroupsOut) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetItems

`func (o *AdminUserGroupsOut) GetItems() []AdminUserGroupOut`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *AdminUserGroupsOut) GetItemsOk() (*[]AdminUserGroupOut, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *AdminUserGroupsOut) SetItems(v []AdminUserGroupOut)`

SetItems sets Items field to given value.


### GetPage

`func (o *AdminUserGroupsOut) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *AdminUserGroupsOut) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *AdminUserGroupsOut) SetPage(v int32)`

SetPage sets Page field to given value.


### GetPageSize

`func (o *AdminUserGroupsOut) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *AdminUserGroupsOut) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *AdminUserGroupsOut) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


