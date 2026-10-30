# UpdateSkillRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SharingScope** | Pointer to [**RegistrySharingScope**](RegistrySharingScope.md) | Registry sharing scope. | [optional] 
**WorkspaceRelation** | Pointer to [**ShareRelation**](ShareRelation.md) | Access every workspace member gets when sharing_scope is &#x60;workspace&#x60;. Defaults to &#x60;writer&#x60;, which is what a workspace publish granted before this field existed, so omitting it preserves today&#39;s behaviour exactly. Rejected unless sharing_scope is &#x60;workspace&#x60;. | [optional] 

## Methods

### NewUpdateSkillRequest

`func NewUpdateSkillRequest() *UpdateSkillRequest`

NewUpdateSkillRequest instantiates a new UpdateSkillRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSkillRequestWithDefaults

`func NewUpdateSkillRequestWithDefaults() *UpdateSkillRequest`

NewUpdateSkillRequestWithDefaults instantiates a new UpdateSkillRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSharingScope

`func (o *UpdateSkillRequest) GetSharingScope() RegistrySharingScope`

GetSharingScope returns the SharingScope field if non-nil, zero value otherwise.

### GetSharingScopeOk

`func (o *UpdateSkillRequest) GetSharingScopeOk() (*RegistrySharingScope, bool)`

GetSharingScopeOk returns a tuple with the SharingScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharingScope

`func (o *UpdateSkillRequest) SetSharingScope(v RegistrySharingScope)`

SetSharingScope sets SharingScope field to given value.

### HasSharingScope

`func (o *UpdateSkillRequest) HasSharingScope() bool`

HasSharingScope returns a boolean if a field has been set.

### GetWorkspaceRelation

`func (o *UpdateSkillRequest) GetWorkspaceRelation() ShareRelation`

GetWorkspaceRelation returns the WorkspaceRelation field if non-nil, zero value otherwise.

### GetWorkspaceRelationOk

`func (o *UpdateSkillRequest) GetWorkspaceRelationOk() (*ShareRelation, bool)`

GetWorkspaceRelationOk returns a tuple with the WorkspaceRelation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceRelation

`func (o *UpdateSkillRequest) SetWorkspaceRelation(v ShareRelation)`

SetWorkspaceRelation sets WorkspaceRelation field to given value.

### HasWorkspaceRelation

`func (o *UpdateSkillRequest) HasWorkspaceRelation() bool`

HasWorkspaceRelation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


