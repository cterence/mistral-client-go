# MemberWorkspaceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | **string** | Workspace ID. | 
**Name** | **string** | Workspace name. | 
**IsDefault** | **bool** | Whether this is the default Workspace for the Organization. | 
**RawRoles** | [**RawRoles1**](RawRoles1.md) |  | 
**RawRole** | [**RawRole1**](RawRole1.md) |  | 

## Methods

### NewMemberWorkspaceInfo

`func NewMemberWorkspaceInfo(uuid string, name string, isDefault bool, rawRoles RawRoles1, rawRole RawRole1, ) *MemberWorkspaceInfo`

NewMemberWorkspaceInfo instantiates a new MemberWorkspaceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberWorkspaceInfoWithDefaults

`func NewMemberWorkspaceInfoWithDefaults() *MemberWorkspaceInfo`

NewMemberWorkspaceInfoWithDefaults instantiates a new MemberWorkspaceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *MemberWorkspaceInfo) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *MemberWorkspaceInfo) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *MemberWorkspaceInfo) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetName

`func (o *MemberWorkspaceInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MemberWorkspaceInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MemberWorkspaceInfo) SetName(v string)`

SetName sets Name field to given value.


### GetIsDefault

`func (o *MemberWorkspaceInfo) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *MemberWorkspaceInfo) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *MemberWorkspaceInfo) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.


### GetRawRoles

`func (o *MemberWorkspaceInfo) GetRawRoles() RawRoles1`

GetRawRoles returns the RawRoles field if non-nil, zero value otherwise.

### GetRawRolesOk

`func (o *MemberWorkspaceInfo) GetRawRolesOk() (*RawRoles1, bool)`

GetRawRolesOk returns a tuple with the RawRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawRoles

`func (o *MemberWorkspaceInfo) SetRawRoles(v RawRoles1)`

SetRawRoles sets RawRoles field to given value.


### GetRawRole

`func (o *MemberWorkspaceInfo) GetRawRole() RawRole1`

GetRawRole returns the RawRole field if non-nil, zero value otherwise.

### GetRawRoleOk

`func (o *MemberWorkspaceInfo) GetRawRoleOk() (*RawRole1, bool)`

GetRawRoleOk returns a tuple with the RawRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawRole

`func (o *MemberWorkspaceInfo) SetRawRole(v RawRole1)`

SetRawRole sets RawRole field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


