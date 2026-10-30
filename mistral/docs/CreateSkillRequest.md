# CreateSkillRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Stable object name. | 
**Definition** | [**SkillDefinition**](SkillDefinition.md) |  | 
**Notes** | Pointer to **string** | Notes for this version. | [optional] 
**SharingScope** | Pointer to [**RegistrySharingScope**](RegistrySharingScope.md) | Registry sharing scope. | [optional] 
**Aliases** | Pointer to **[]string** | Aliases pointing to this version. | [optional] 
**WorkspaceRelation** | Pointer to [**ShareRelation**](ShareRelation.md) | Access every workspace member gets when sharing_scope is &#x60;workspace&#x60;. Defaults to &#x60;writer&#x60;, which is what a workspace publish granted before this field existed, so omitting it preserves today&#39;s behaviour exactly. Rejected unless sharing_scope is &#x60;workspace&#x60;. | [optional] 

## Methods

### NewCreateSkillRequest

`func NewCreateSkillRequest(name string, definition SkillDefinition, ) *CreateSkillRequest`

NewCreateSkillRequest instantiates a new CreateSkillRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSkillRequestWithDefaults

`func NewCreateSkillRequestWithDefaults() *CreateSkillRequest`

NewCreateSkillRequestWithDefaults instantiates a new CreateSkillRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateSkillRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateSkillRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateSkillRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDefinition

`func (o *CreateSkillRequest) GetDefinition() SkillDefinition`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *CreateSkillRequest) GetDefinitionOk() (*SkillDefinition, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *CreateSkillRequest) SetDefinition(v SkillDefinition)`

SetDefinition sets Definition field to given value.


### GetNotes

`func (o *CreateSkillRequest) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *CreateSkillRequest) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *CreateSkillRequest) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *CreateSkillRequest) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetSharingScope

`func (o *CreateSkillRequest) GetSharingScope() RegistrySharingScope`

GetSharingScope returns the SharingScope field if non-nil, zero value otherwise.

### GetSharingScopeOk

`func (o *CreateSkillRequest) GetSharingScopeOk() (*RegistrySharingScope, bool)`

GetSharingScopeOk returns a tuple with the SharingScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharingScope

`func (o *CreateSkillRequest) SetSharingScope(v RegistrySharingScope)`

SetSharingScope sets SharingScope field to given value.

### HasSharingScope

`func (o *CreateSkillRequest) HasSharingScope() bool`

HasSharingScope returns a boolean if a field has been set.

### GetAliases

`func (o *CreateSkillRequest) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *CreateSkillRequest) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *CreateSkillRequest) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *CreateSkillRequest) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetWorkspaceRelation

`func (o *CreateSkillRequest) GetWorkspaceRelation() ShareRelation`

GetWorkspaceRelation returns the WorkspaceRelation field if non-nil, zero value otherwise.

### GetWorkspaceRelationOk

`func (o *CreateSkillRequest) GetWorkspaceRelationOk() (*ShareRelation, bool)`

GetWorkspaceRelationOk returns a tuple with the WorkspaceRelation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceRelation

`func (o *CreateSkillRequest) SetWorkspaceRelation(v ShareRelation)`

SetWorkspaceRelation sets WorkspaceRelation field to given value.

### HasWorkspaceRelation

`func (o *CreateSkillRequest) HasWorkspaceRelation() bool`

HasWorkspaceRelation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


