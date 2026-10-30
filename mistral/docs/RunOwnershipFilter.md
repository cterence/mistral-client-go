# RunOwnershipFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Combinator** | Pointer to [**FilterCombinator**](FilterCombinator.md) |  | [optional] [default to FILTERCOMBINATOR_OR]
**CreatorIds** | Pointer to **[]string** |  | [optional] [default to {}]
**ProjectIds** | Pointer to **[]string** |  | [optional] [default to {}]
**StarredByMe** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewRunOwnershipFilter

`func NewRunOwnershipFilter() *RunOwnershipFilter`

NewRunOwnershipFilter instantiates a new RunOwnershipFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunOwnershipFilterWithDefaults

`func NewRunOwnershipFilterWithDefaults() *RunOwnershipFilter`

NewRunOwnershipFilterWithDefaults instantiates a new RunOwnershipFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCombinator

`func (o *RunOwnershipFilter) GetCombinator() FilterCombinator`

GetCombinator returns the Combinator field if non-nil, zero value otherwise.

### GetCombinatorOk

`func (o *RunOwnershipFilter) GetCombinatorOk() (*FilterCombinator, bool)`

GetCombinatorOk returns a tuple with the Combinator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCombinator

`func (o *RunOwnershipFilter) SetCombinator(v FilterCombinator)`

SetCombinator sets Combinator field to given value.

### HasCombinator

`func (o *RunOwnershipFilter) HasCombinator() bool`

HasCombinator returns a boolean if a field has been set.

### GetCreatorIds

`func (o *RunOwnershipFilter) GetCreatorIds() []string`

GetCreatorIds returns the CreatorIds field if non-nil, zero value otherwise.

### GetCreatorIdsOk

`func (o *RunOwnershipFilter) GetCreatorIdsOk() (*[]string, bool)`

GetCreatorIdsOk returns a tuple with the CreatorIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorIds

`func (o *RunOwnershipFilter) SetCreatorIds(v []string)`

SetCreatorIds sets CreatorIds field to given value.

### HasCreatorIds

`func (o *RunOwnershipFilter) HasCreatorIds() bool`

HasCreatorIds returns a boolean if a field has been set.

### GetProjectIds

`func (o *RunOwnershipFilter) GetProjectIds() []string`

GetProjectIds returns the ProjectIds field if non-nil, zero value otherwise.

### GetProjectIdsOk

`func (o *RunOwnershipFilter) GetProjectIdsOk() (*[]string, bool)`

GetProjectIdsOk returns a tuple with the ProjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIds

`func (o *RunOwnershipFilter) SetProjectIds(v []string)`

SetProjectIds sets ProjectIds field to given value.

### HasProjectIds

`func (o *RunOwnershipFilter) HasProjectIds() bool`

HasProjectIds returns a boolean if a field has been set.

### GetStarredByMe

`func (o *RunOwnershipFilter) GetStarredByMe() bool`

GetStarredByMe returns the StarredByMe field if non-nil, zero value otherwise.

### GetStarredByMeOk

`func (o *RunOwnershipFilter) GetStarredByMeOk() (*bool, bool)`

GetStarredByMeOk returns a tuple with the StarredByMe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarredByMe

`func (o *RunOwnershipFilter) SetStarredByMe(v bool)`

SetStarredByMe sets StarredByMe field to given value.

### HasStarredByMe

`func (o *RunOwnershipFilter) HasStarredByMe() bool`

HasStarredByMe returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


