# Skill

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** | Stable object name. | [optional] 
**Definition** | Pointer to [**SkillDefinition**](SkillDefinition.md) |  | [optional] 
**Version** | Pointer to **int32** |  | [optional] 
**Notes** | Pointer to **string** | Notes for this version. | [optional] 
**Aliases** | Pointer to **[]string** | Aliases pointing to this version. | [optional] 
**SharingScope** | Pointer to [**RegistrySharingScope**](RegistrySharingScope.md) | Registry sharing scope. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Creation time. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Last update time. | [optional] 
**LatestVersion** | Pointer to **int32** | Latest version number. | [optional] 

## Methods

### NewSkill

`func NewSkill() *Skill`

NewSkill instantiates a new Skill object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkillWithDefaults

`func NewSkillWithDefaults() *Skill`

NewSkillWithDefaults instantiates a new Skill object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Skill) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Skill) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Skill) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Skill) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Skill) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Skill) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Skill) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Skill) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDefinition

`func (o *Skill) GetDefinition() SkillDefinition`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *Skill) GetDefinitionOk() (*SkillDefinition, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *Skill) SetDefinition(v SkillDefinition)`

SetDefinition sets Definition field to given value.

### HasDefinition

`func (o *Skill) HasDefinition() bool`

HasDefinition returns a boolean if a field has been set.

### GetVersion

`func (o *Skill) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Skill) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Skill) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Skill) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetNotes

`func (o *Skill) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *Skill) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *Skill) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *Skill) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetAliases

`func (o *Skill) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *Skill) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *Skill) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *Skill) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetSharingScope

`func (o *Skill) GetSharingScope() RegistrySharingScope`

GetSharingScope returns the SharingScope field if non-nil, zero value otherwise.

### GetSharingScopeOk

`func (o *Skill) GetSharingScopeOk() (*RegistrySharingScope, bool)`

GetSharingScopeOk returns a tuple with the SharingScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharingScope

`func (o *Skill) SetSharingScope(v RegistrySharingScope)`

SetSharingScope sets SharingScope field to given value.

### HasSharingScope

`func (o *Skill) HasSharingScope() bool`

HasSharingScope returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Skill) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Skill) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Skill) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Skill) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Skill) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Skill) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Skill) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Skill) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetLatestVersion

`func (o *Skill) GetLatestVersion() int32`

GetLatestVersion returns the LatestVersion field if non-nil, zero value otherwise.

### GetLatestVersionOk

`func (o *Skill) GetLatestVersionOk() (*int32, bool)`

GetLatestVersionOk returns a tuple with the LatestVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestVersion

`func (o *Skill) SetLatestVersion(v int32)`

SetLatestVersion sets LatestVersion field to given value.

### HasLatestVersion

`func (o *Skill) HasLatestVersion() bool`

HasLatestVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


