# CreateSkillVersionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Definition** | [**SkillDefinition**](SkillDefinition.md) |  | 
**Notes** | Pointer to **string** | Notes for this version. | [optional] 
**Aliases** | Pointer to **[]string** | Aliases pointing to this version. | [optional] 

## Methods

### NewCreateSkillVersionRequest

`func NewCreateSkillVersionRequest(definition SkillDefinition, ) *CreateSkillVersionRequest`

NewCreateSkillVersionRequest instantiates a new CreateSkillVersionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSkillVersionRequestWithDefaults

`func NewCreateSkillVersionRequestWithDefaults() *CreateSkillVersionRequest`

NewCreateSkillVersionRequestWithDefaults instantiates a new CreateSkillVersionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefinition

`func (o *CreateSkillVersionRequest) GetDefinition() SkillDefinition`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *CreateSkillVersionRequest) GetDefinitionOk() (*SkillDefinition, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *CreateSkillVersionRequest) SetDefinition(v SkillDefinition)`

SetDefinition sets Definition field to given value.


### GetNotes

`func (o *CreateSkillVersionRequest) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *CreateSkillVersionRequest) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *CreateSkillVersionRequest) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *CreateSkillVersionRequest) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetAliases

`func (o *CreateSkillVersionRequest) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *CreateSkillVersionRequest) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *CreateSkillVersionRequest) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *CreateSkillVersionRequest) HasAliases() bool`

HasAliases returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


