# UpdatePromptVersionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Notes** | Pointer to **string** | Notes for this version. | [optional] 
**Aliases** | Pointer to [**AliasList**](AliasList.md) | Aliases pointing to this version. | [optional] 

## Methods

### NewUpdatePromptVersionRequest

`func NewUpdatePromptVersionRequest() *UpdatePromptVersionRequest`

NewUpdatePromptVersionRequest instantiates a new UpdatePromptVersionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePromptVersionRequestWithDefaults

`func NewUpdatePromptVersionRequestWithDefaults() *UpdatePromptVersionRequest`

NewUpdatePromptVersionRequestWithDefaults instantiates a new UpdatePromptVersionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotes

`func (o *UpdatePromptVersionRequest) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *UpdatePromptVersionRequest) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *UpdatePromptVersionRequest) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *UpdatePromptVersionRequest) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetAliases

`func (o *UpdatePromptVersionRequest) GetAliases() AliasList`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *UpdatePromptVersionRequest) GetAliasesOk() (*AliasList, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *UpdatePromptVersionRequest) SetAliases(v AliasList)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *UpdatePromptVersionRequest) HasAliases() bool`

HasAliases returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


