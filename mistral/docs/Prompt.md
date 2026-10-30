# Prompt

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** | Stable object name. | [optional] 
**Definition** | Pointer to [**PromptDefinition**](PromptDefinition.md) |  | [optional] 
**Version** | Pointer to **int32** |  | [optional] 
**Notes** | Pointer to **string** | Notes for this version. | [optional] 
**Aliases** | Pointer to **[]string** | Aliases pointing to this version. | [optional] 
**SharingScope** | Pointer to [**RegistrySharingScope**](RegistrySharingScope.md) | Registry sharing scope. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Creation time. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Last update time. | [optional] 
**LatestVersion** | Pointer to **int32** | Latest version number. | [optional] 
**Title** | Pointer to **string** | Display title. | [optional] 
**Description** | Pointer to **string** | Display description. | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**VersionCreatedAt** | Pointer to **time.Time** | RFC 3339 timestamp. | [optional] 

## Methods

### NewPrompt

`func NewPrompt() *Prompt`

NewPrompt instantiates a new Prompt object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPromptWithDefaults

`func NewPromptWithDefaults() *Prompt`

NewPromptWithDefaults instantiates a new Prompt object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Prompt) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Prompt) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Prompt) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Prompt) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Prompt) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Prompt) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Prompt) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Prompt) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDefinition

`func (o *Prompt) GetDefinition() PromptDefinition`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *Prompt) GetDefinitionOk() (*PromptDefinition, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *Prompt) SetDefinition(v PromptDefinition)`

SetDefinition sets Definition field to given value.

### HasDefinition

`func (o *Prompt) HasDefinition() bool`

HasDefinition returns a boolean if a field has been set.

### GetVersion

`func (o *Prompt) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Prompt) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Prompt) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Prompt) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetNotes

`func (o *Prompt) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *Prompt) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *Prompt) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *Prompt) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetAliases

`func (o *Prompt) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *Prompt) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *Prompt) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *Prompt) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetSharingScope

`func (o *Prompt) GetSharingScope() RegistrySharingScope`

GetSharingScope returns the SharingScope field if non-nil, zero value otherwise.

### GetSharingScopeOk

`func (o *Prompt) GetSharingScopeOk() (*RegistrySharingScope, bool)`

GetSharingScopeOk returns a tuple with the SharingScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharingScope

`func (o *Prompt) SetSharingScope(v RegistrySharingScope)`

SetSharingScope sets SharingScope field to given value.

### HasSharingScope

`func (o *Prompt) HasSharingScope() bool`

HasSharingScope returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Prompt) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Prompt) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Prompt) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Prompt) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Prompt) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Prompt) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Prompt) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Prompt) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetLatestVersion

`func (o *Prompt) GetLatestVersion() int32`

GetLatestVersion returns the LatestVersion field if non-nil, zero value otherwise.

### GetLatestVersionOk

`func (o *Prompt) GetLatestVersionOk() (*int32, bool)`

GetLatestVersionOk returns a tuple with the LatestVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestVersion

`func (o *Prompt) SetLatestVersion(v int32)`

SetLatestVersion sets LatestVersion field to given value.

### HasLatestVersion

`func (o *Prompt) HasLatestVersion() bool`

HasLatestVersion returns a boolean if a field has been set.

### GetTitle

`func (o *Prompt) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Prompt) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Prompt) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Prompt) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *Prompt) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Prompt) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Prompt) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Prompt) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreatedBy

`func (o *Prompt) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Prompt) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Prompt) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Prompt) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetVersionCreatedAt

`func (o *Prompt) GetVersionCreatedAt() time.Time`

GetVersionCreatedAt returns the VersionCreatedAt field if non-nil, zero value otherwise.

### GetVersionCreatedAtOk

`func (o *Prompt) GetVersionCreatedAtOk() (*time.Time, bool)`

GetVersionCreatedAtOk returns a tuple with the VersionCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionCreatedAt

`func (o *Prompt) SetVersionCreatedAt(v time.Time)`

SetVersionCreatedAt sets VersionCreatedAt field to given value.

### HasVersionCreatedAt

`func (o *Prompt) HasVersionCreatedAt() bool`

HasVersionCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


