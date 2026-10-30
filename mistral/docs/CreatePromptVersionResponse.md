# CreatePromptVersionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **int32** |  | [optional] 
**Deduplicated** | Pointer to **bool** |  | [optional] 

## Methods

### NewCreatePromptVersionResponse

`func NewCreatePromptVersionResponse() *CreatePromptVersionResponse`

NewCreatePromptVersionResponse instantiates a new CreatePromptVersionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePromptVersionResponseWithDefaults

`func NewCreatePromptVersionResponseWithDefaults() *CreatePromptVersionResponse`

NewCreatePromptVersionResponseWithDefaults instantiates a new CreatePromptVersionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *CreatePromptVersionResponse) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CreatePromptVersionResponse) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CreatePromptVersionResponse) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CreatePromptVersionResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetDeduplicated

`func (o *CreatePromptVersionResponse) GetDeduplicated() bool`

GetDeduplicated returns the Deduplicated field if non-nil, zero value otherwise.

### GetDeduplicatedOk

`func (o *CreatePromptVersionResponse) GetDeduplicatedOk() (*bool, bool)`

GetDeduplicatedOk returns a tuple with the Deduplicated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeduplicated

`func (o *CreatePromptVersionResponse) SetDeduplicated(v bool)`

SetDeduplicated sets Deduplicated field to given value.

### HasDeduplicated

`func (o *CreatePromptVersionResponse) HasDeduplicated() bool`

HasDeduplicated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


