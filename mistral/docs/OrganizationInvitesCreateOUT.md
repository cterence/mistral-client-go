# OrganizationInvitesCreateOUT

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvalidEmails** | **[]string** | Email addresses that could not be invited. | 
**AlreadyMembers** | **[]string** | Email addresses that already belong to the Organization. | 
**InvitedMembersCount** | Pointer to **NullableInt32** | Number of invitations successfully created. | [optional] 

## Methods

### NewOrganizationInvitesCreateOUT

`func NewOrganizationInvitesCreateOUT(invalidEmails []string, alreadyMembers []string, ) *OrganizationInvitesCreateOUT`

NewOrganizationInvitesCreateOUT instantiates a new OrganizationInvitesCreateOUT object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationInvitesCreateOUTWithDefaults

`func NewOrganizationInvitesCreateOUTWithDefaults() *OrganizationInvitesCreateOUT`

NewOrganizationInvitesCreateOUTWithDefaults instantiates a new OrganizationInvitesCreateOUT object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvalidEmails

`func (o *OrganizationInvitesCreateOUT) GetInvalidEmails() []string`

GetInvalidEmails returns the InvalidEmails field if non-nil, zero value otherwise.

### GetInvalidEmailsOk

`func (o *OrganizationInvitesCreateOUT) GetInvalidEmailsOk() (*[]string, bool)`

GetInvalidEmailsOk returns a tuple with the InvalidEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidEmails

`func (o *OrganizationInvitesCreateOUT) SetInvalidEmails(v []string)`

SetInvalidEmails sets InvalidEmails field to given value.


### GetAlreadyMembers

`func (o *OrganizationInvitesCreateOUT) GetAlreadyMembers() []string`

GetAlreadyMembers returns the AlreadyMembers field if non-nil, zero value otherwise.

### GetAlreadyMembersOk

`func (o *OrganizationInvitesCreateOUT) GetAlreadyMembersOk() (*[]string, bool)`

GetAlreadyMembersOk returns a tuple with the AlreadyMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlreadyMembers

`func (o *OrganizationInvitesCreateOUT) SetAlreadyMembers(v []string)`

SetAlreadyMembers sets AlreadyMembers field to given value.


### GetInvitedMembersCount

`func (o *OrganizationInvitesCreateOUT) GetInvitedMembersCount() int32`

GetInvitedMembersCount returns the InvitedMembersCount field if non-nil, zero value otherwise.

### GetInvitedMembersCountOk

`func (o *OrganizationInvitesCreateOUT) GetInvitedMembersCountOk() (*int32, bool)`

GetInvitedMembersCountOk returns a tuple with the InvitedMembersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitedMembersCount

`func (o *OrganizationInvitesCreateOUT) SetInvitedMembersCount(v int32)`

SetInvitedMembersCount sets InvitedMembersCount field to given value.

### HasInvitedMembersCount

`func (o *OrganizationInvitesCreateOUT) HasInvitedMembersCount() bool`

HasInvitedMembersCount returns a boolean if a field has been set.

### SetInvitedMembersCountNil

`func (o *OrganizationInvitesCreateOUT) SetInvitedMembersCountNil(b bool)`

 SetInvitedMembersCountNil sets the value for InvitedMembersCount to be an explicit nil

### UnsetInvitedMembersCount
`func (o *OrganizationInvitesCreateOUT) UnsetInvitedMembersCount()`

UnsetInvitedMembersCount ensures that no value is present for InvitedMembersCount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


