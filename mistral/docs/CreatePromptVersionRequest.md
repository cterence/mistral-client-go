# CreatePromptVersionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Definition** | [**PromptDefinition**](PromptDefinition.md) |  | 
**Notes** | Pointer to **string** | Notes for this version. | [optional] 
**Aliases** | Pointer to **[]string** | Aliases pointing to this version. | [optional] 

## Methods

### NewCreatePromptVersionRequest

`func NewCreatePromptVersionRequest(definition PromptDefinition, ) *CreatePromptVersionRequest`

NewCreatePromptVersionRequest instantiates a new CreatePromptVersionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePromptVersionRequestWithDefaults

`func NewCreatePromptVersionRequestWithDefaults() *CreatePromptVersionRequest`

NewCreatePromptVersionRequestWithDefaults instantiates a new CreatePromptVersionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefinition

`func (o *CreatePromptVersionRequest) GetDefinition() PromptDefinition`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *CreatePromptVersionRequest) GetDefinitionOk() (*PromptDefinition, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *CreatePromptVersionRequest) SetDefinition(v PromptDefinition)`

SetDefinition sets Definition field to given value.


### GetNotes

`func (o *CreatePromptVersionRequest) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *CreatePromptVersionRequest) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *CreatePromptVersionRequest) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *CreatePromptVersionRequest) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetAliases

`func (o *CreatePromptVersionRequest) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *CreatePromptVersionRequest) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *CreatePromptVersionRequest) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *CreatePromptVersionRequest) HasAliases() bool`

HasAliases returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


