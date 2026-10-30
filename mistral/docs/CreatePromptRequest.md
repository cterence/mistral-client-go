# CreatePromptRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Stable object name. | 
**Definition** | [**PromptDefinition**](PromptDefinition.md) |  | 
**Title** | Pointer to **string** | Display title. | [optional] 
**Description** | Pointer to **string** | Display description. | [optional] 
**Notes** | Pointer to **string** | Notes for this version. | [optional] 
**SharingScope** | Pointer to [**RegistrySharingScope**](RegistrySharingScope.md) | Registry sharing scope. | [optional] 
**Aliases** | Pointer to **[]string** | Aliases pointing to this version. | [optional] 
**WorkspaceRelation** | Pointer to [**ShareRelation**](ShareRelation.md) | Access every workspace member gets when sharing_scope is &#x60;workspace&#x60;. Defaults to &#x60;writer&#x60;, which is what a workspace publish granted before this field existed, so omitting it preserves today&#39;s behaviour exactly. Rejected unless sharing_scope is &#x60;workspace&#x60;. | [optional] 

## Methods

### NewCreatePromptRequest

`func NewCreatePromptRequest(name string, definition PromptDefinition, ) *CreatePromptRequest`

NewCreatePromptRequest instantiates a new CreatePromptRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePromptRequestWithDefaults

`func NewCreatePromptRequestWithDefaults() *CreatePromptRequest`

NewCreatePromptRequestWithDefaults instantiates a new CreatePromptRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreatePromptRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreatePromptRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreatePromptRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDefinition

`func (o *CreatePromptRequest) GetDefinition() PromptDefinition`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *CreatePromptRequest) GetDefinitionOk() (*PromptDefinition, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *CreatePromptRequest) SetDefinition(v PromptDefinition)`

SetDefinition sets Definition field to given value.


### GetTitle

`func (o *CreatePromptRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreatePromptRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreatePromptRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CreatePromptRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *CreatePromptRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreatePromptRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreatePromptRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreatePromptRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetNotes

`func (o *CreatePromptRequest) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *CreatePromptRequest) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *CreatePromptRequest) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *CreatePromptRequest) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetSharingScope

`func (o *CreatePromptRequest) GetSharingScope() RegistrySharingScope`

GetSharingScope returns the SharingScope field if non-nil, zero value otherwise.

### GetSharingScopeOk

`func (o *CreatePromptRequest) GetSharingScopeOk() (*RegistrySharingScope, bool)`

GetSharingScopeOk returns a tuple with the SharingScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharingScope

`func (o *CreatePromptRequest) SetSharingScope(v RegistrySharingScope)`

SetSharingScope sets SharingScope field to given value.

### HasSharingScope

`func (o *CreatePromptRequest) HasSharingScope() bool`

HasSharingScope returns a boolean if a field has been set.

### GetAliases

`func (o *CreatePromptRequest) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *CreatePromptRequest) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *CreatePromptRequest) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *CreatePromptRequest) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetWorkspaceRelation

`func (o *CreatePromptRequest) GetWorkspaceRelation() ShareRelation`

GetWorkspaceRelation returns the WorkspaceRelation field if non-nil, zero value otherwise.

### GetWorkspaceRelationOk

`func (o *CreatePromptRequest) GetWorkspaceRelationOk() (*ShareRelation, bool)`

GetWorkspaceRelationOk returns a tuple with the WorkspaceRelation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceRelation

`func (o *CreatePromptRequest) SetWorkspaceRelation(v ShareRelation)`

SetWorkspaceRelation sets WorkspaceRelation field to given value.

### HasWorkspaceRelation

`func (o *CreatePromptRequest) HasWorkspaceRelation() bool`

HasWorkspaceRelation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


