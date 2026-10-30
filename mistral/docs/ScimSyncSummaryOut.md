# ScimSyncSummaryOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | [**ScimSyncUsersSummaryOut**](ScimSyncUsersSummaryOut.md) |  | 
**Groups** | [**ScimSyncGroupsSummaryOut**](ScimSyncGroupsSummaryOut.md) |  | 
**Memberships** | [**[]ScimSyncMembershipActionOut**](ScimSyncMembershipActionOut.md) |  | 

## Methods

### NewScimSyncSummaryOut

`func NewScimSyncSummaryOut(users ScimSyncUsersSummaryOut, groups ScimSyncGroupsSummaryOut, memberships []ScimSyncMembershipActionOut, ) *ScimSyncSummaryOut`

NewScimSyncSummaryOut instantiates a new ScimSyncSummaryOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScimSyncSummaryOutWithDefaults

`func NewScimSyncSummaryOutWithDefaults() *ScimSyncSummaryOut`

NewScimSyncSummaryOutWithDefaults instantiates a new ScimSyncSummaryOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *ScimSyncSummaryOut) GetUsers() ScimSyncUsersSummaryOut`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *ScimSyncSummaryOut) GetUsersOk() (*ScimSyncUsersSummaryOut, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *ScimSyncSummaryOut) SetUsers(v ScimSyncUsersSummaryOut)`

SetUsers sets Users field to given value.


### GetGroups

`func (o *ScimSyncSummaryOut) GetGroups() ScimSyncGroupsSummaryOut`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *ScimSyncSummaryOut) GetGroupsOk() (*ScimSyncGroupsSummaryOut, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *ScimSyncSummaryOut) SetGroups(v ScimSyncGroupsSummaryOut)`

SetGroups sets Groups field to given value.


### GetMemberships

`func (o *ScimSyncSummaryOut) GetMemberships() []ScimSyncMembershipActionOut`

GetMemberships returns the Memberships field if non-nil, zero value otherwise.

### GetMembershipsOk

`func (o *ScimSyncSummaryOut) GetMembershipsOk() (*[]ScimSyncMembershipActionOut, bool)`

GetMembershipsOk returns a tuple with the Memberships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberships

`func (o *ScimSyncSummaryOut) SetMemberships(v []ScimSyncMembershipActionOut)`

SetMemberships sets Memberships field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


