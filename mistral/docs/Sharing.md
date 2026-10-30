# Sharing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LibraryId** | **string** |  | 
**UserId** | Pointer to **NullableString** |  | [optional] 
**OrgId** | **string** |  | 
**Role** | **string** |  | 
**ShareWithType** | **string** |  | 
**ShareWithUuid** | **NullableString** |  | 

## Methods

### NewSharing

`func NewSharing(libraryId string, orgId string, role string, shareWithType string, shareWithUuid NullableString, ) *Sharing`

NewSharing instantiates a new Sharing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharingWithDefaults

`func NewSharingWithDefaults() *Sharing`

NewSharingWithDefaults instantiates a new Sharing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLibraryId

`func (o *Sharing) GetLibraryId() string`

GetLibraryId returns the LibraryId field if non-nil, zero value otherwise.

### GetLibraryIdOk

`func (o *Sharing) GetLibraryIdOk() (*string, bool)`

GetLibraryIdOk returns a tuple with the LibraryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraryId

`func (o *Sharing) SetLibraryId(v string)`

SetLibraryId sets LibraryId field to given value.


### GetUserId

`func (o *Sharing) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Sharing) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Sharing) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *Sharing) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *Sharing) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *Sharing) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetOrgId

`func (o *Sharing) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *Sharing) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *Sharing) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetRole

`func (o *Sharing) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *Sharing) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *Sharing) SetRole(v string)`

SetRole sets Role field to given value.


### GetShareWithType

`func (o *Sharing) GetShareWithType() string`

GetShareWithType returns the ShareWithType field if non-nil, zero value otherwise.

### GetShareWithTypeOk

`func (o *Sharing) GetShareWithTypeOk() (*string, bool)`

GetShareWithTypeOk returns a tuple with the ShareWithType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareWithType

`func (o *Sharing) SetShareWithType(v string)`

SetShareWithType sets ShareWithType field to given value.


### GetShareWithUuid

`func (o *Sharing) GetShareWithUuid() string`

GetShareWithUuid returns the ShareWithUuid field if non-nil, zero value otherwise.

### GetShareWithUuidOk

`func (o *Sharing) GetShareWithUuidOk() (*string, bool)`

GetShareWithUuidOk returns a tuple with the ShareWithUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareWithUuid

`func (o *Sharing) SetShareWithUuid(v string)`

SetShareWithUuid sets ShareWithUuid field to given value.


### SetShareWithUuidNil

`func (o *Sharing) SetShareWithUuidNil(b bool)`

 SetShareWithUuidNil sets the value for ShareWithUuid to be an explicit nil

### UnsetShareWithUuid
`func (o *Sharing) UnsetShareWithUuid()`

UnsetShareWithUuid ensures that no value is present for ShareWithUuid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


