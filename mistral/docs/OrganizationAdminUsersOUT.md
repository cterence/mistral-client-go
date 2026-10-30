# OrganizationAdminUsersOUT

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Members** | [**[]AdminOrganizationMemberOUT**](AdminOrganizationMemberOUT.md) | Organization members on this page. | 
**Invites** | [**[]OrganizationInviteOUT**](OrganizationInviteOUT.md) | Pending Organization invitations on this page. | 
**Total** | **NullableInt32** | Total number of users and invitations that match the request. | 
**Page** | **NullableInt32** | Page number returned. | 
**PageSize** | **NullableInt32** | Maximum number of results per page. | 

## Methods

### NewOrganizationAdminUsersOUT

`func NewOrganizationAdminUsersOUT(members []AdminOrganizationMemberOUT, invites []OrganizationInviteOUT, total NullableInt32, page NullableInt32, pageSize NullableInt32, ) *OrganizationAdminUsersOUT`

NewOrganizationAdminUsersOUT instantiates a new OrganizationAdminUsersOUT object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationAdminUsersOUTWithDefaults

`func NewOrganizationAdminUsersOUTWithDefaults() *OrganizationAdminUsersOUT`

NewOrganizationAdminUsersOUTWithDefaults instantiates a new OrganizationAdminUsersOUT object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMembers

`func (o *OrganizationAdminUsersOUT) GetMembers() []AdminOrganizationMemberOUT`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *OrganizationAdminUsersOUT) GetMembersOk() (*[]AdminOrganizationMemberOUT, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *OrganizationAdminUsersOUT) SetMembers(v []AdminOrganizationMemberOUT)`

SetMembers sets Members field to given value.


### GetInvites

`func (o *OrganizationAdminUsersOUT) GetInvites() []OrganizationInviteOUT`

GetInvites returns the Invites field if non-nil, zero value otherwise.

### GetInvitesOk

`func (o *OrganizationAdminUsersOUT) GetInvitesOk() (*[]OrganizationInviteOUT, bool)`

GetInvitesOk returns a tuple with the Invites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvites

`func (o *OrganizationAdminUsersOUT) SetInvites(v []OrganizationInviteOUT)`

SetInvites sets Invites field to given value.


### GetTotal

`func (o *OrganizationAdminUsersOUT) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *OrganizationAdminUsersOUT) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *OrganizationAdminUsersOUT) SetTotal(v int32)`

SetTotal sets Total field to given value.


### SetTotalNil

`func (o *OrganizationAdminUsersOUT) SetTotalNil(b bool)`

 SetTotalNil sets the value for Total to be an explicit nil

### UnsetTotal
`func (o *OrganizationAdminUsersOUT) UnsetTotal()`

UnsetTotal ensures that no value is present for Total, not even an explicit nil
### GetPage

`func (o *OrganizationAdminUsersOUT) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *OrganizationAdminUsersOUT) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *OrganizationAdminUsersOUT) SetPage(v int32)`

SetPage sets Page field to given value.


### SetPageNil

`func (o *OrganizationAdminUsersOUT) SetPageNil(b bool)`

 SetPageNil sets the value for Page to be an explicit nil

### UnsetPage
`func (o *OrganizationAdminUsersOUT) UnsetPage()`

UnsetPage ensures that no value is present for Page, not even an explicit nil
### GetPageSize

`func (o *OrganizationAdminUsersOUT) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *OrganizationAdminUsersOUT) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *OrganizationAdminUsersOUT) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.


### SetPageSizeNil

`func (o *OrganizationAdminUsersOUT) SetPageSizeNil(b bool)`

 SetPageSizeNil sets the value for PageSize to be an explicit nil

### UnsetPageSize
`func (o *OrganizationAdminUsersOUT) UnsetPageSize()`

UnsetPageSize ensures that no value is present for PageSize, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


