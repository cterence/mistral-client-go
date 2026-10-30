# PromptVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **int32** |  | [optional] 
**Definition** | Pointer to [**PromptDefinition**](PromptDefinition.md) |  | [optional] 
**Notes** | Pointer to **string** | Notes for this version. | [optional] 
**Aliases** | Pointer to **[]string** | Aliases pointing to this version. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Creation time. | [optional] 

## Methods

### NewPromptVersion

`func NewPromptVersion() *PromptVersion`

NewPromptVersion instantiates a new PromptVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPromptVersionWithDefaults

`func NewPromptVersionWithDefaults() *PromptVersion`

NewPromptVersionWithDefaults instantiates a new PromptVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *PromptVersion) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PromptVersion) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PromptVersion) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PromptVersion) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetDefinition

`func (o *PromptVersion) GetDefinition() PromptDefinition`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *PromptVersion) GetDefinitionOk() (*PromptDefinition, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *PromptVersion) SetDefinition(v PromptDefinition)`

SetDefinition sets Definition field to given value.

### HasDefinition

`func (o *PromptVersion) HasDefinition() bool`

HasDefinition returns a boolean if a field has been set.

### GetNotes

`func (o *PromptVersion) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *PromptVersion) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *PromptVersion) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *PromptVersion) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetAliases

`func (o *PromptVersion) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *PromptVersion) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *PromptVersion) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *PromptVersion) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PromptVersion) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PromptVersion) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PromptVersion) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PromptVersion) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


