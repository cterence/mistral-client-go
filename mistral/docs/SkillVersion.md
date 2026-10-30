# SkillVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **int32** |  | [optional] 
**Definition** | Pointer to [**SkillDefinition**](SkillDefinition.md) |  | [optional] 
**Notes** | Pointer to **string** | Notes for this version. | [optional] 
**Aliases** | Pointer to **[]string** | Aliases pointing to this version. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Creation time. | [optional] 

## Methods

### NewSkillVersion

`func NewSkillVersion() *SkillVersion`

NewSkillVersion instantiates a new SkillVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkillVersionWithDefaults

`func NewSkillVersionWithDefaults() *SkillVersion`

NewSkillVersionWithDefaults instantiates a new SkillVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *SkillVersion) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SkillVersion) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SkillVersion) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SkillVersion) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetDefinition

`func (o *SkillVersion) GetDefinition() SkillDefinition`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *SkillVersion) GetDefinitionOk() (*SkillDefinition, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *SkillVersion) SetDefinition(v SkillDefinition)`

SetDefinition sets Definition field to given value.

### HasDefinition

`func (o *SkillVersion) HasDefinition() bool`

HasDefinition returns a boolean if a field has been set.

### GetNotes

`func (o *SkillVersion) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *SkillVersion) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *SkillVersion) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *SkillVersion) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetAliases

`func (o *SkillVersion) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *SkillVersion) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *SkillVersion) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *SkillVersion) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SkillVersion) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SkillVersion) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SkillVersion) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SkillVersion) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


