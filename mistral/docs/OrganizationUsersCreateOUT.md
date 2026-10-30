# OrganizationUsersCreateOUT

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvalidEmails** | **[]string** | Email addresses that could not be created. | 
**EmailToUserId** | **map[string]string** | Mapping from created user email to user identifier. | 

## Methods

### NewOrganizationUsersCreateOUT

`func NewOrganizationUsersCreateOUT(invalidEmails []string, emailToUserId map[string]string, ) *OrganizationUsersCreateOUT`

NewOrganizationUsersCreateOUT instantiates a new OrganizationUsersCreateOUT object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationUsersCreateOUTWithDefaults

`func NewOrganizationUsersCreateOUTWithDefaults() *OrganizationUsersCreateOUT`

NewOrganizationUsersCreateOUTWithDefaults instantiates a new OrganizationUsersCreateOUT object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvalidEmails

`func (o *OrganizationUsersCreateOUT) GetInvalidEmails() []string`

GetInvalidEmails returns the InvalidEmails field if non-nil, zero value otherwise.

### GetInvalidEmailsOk

`func (o *OrganizationUsersCreateOUT) GetInvalidEmailsOk() (*[]string, bool)`

GetInvalidEmailsOk returns a tuple with the InvalidEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidEmails

`func (o *OrganizationUsersCreateOUT) SetInvalidEmails(v []string)`

SetInvalidEmails sets InvalidEmails field to given value.


### GetEmailToUserId

`func (o *OrganizationUsersCreateOUT) GetEmailToUserId() map[string]string`

GetEmailToUserId returns the EmailToUserId field if non-nil, zero value otherwise.

### GetEmailToUserIdOk

`func (o *OrganizationUsersCreateOUT) GetEmailToUserIdOk() (*map[string]string, bool)`

GetEmailToUserIdOk returns a tuple with the EmailToUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailToUserId

`func (o *OrganizationUsersCreateOUT) SetEmailToUserId(v map[string]string)`

SetEmailToUserId sets EmailToUserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


