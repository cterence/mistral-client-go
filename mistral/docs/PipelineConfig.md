# PipelineConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**DeletedAt** | **NullableTime** |  | 
**Scope** | [**PipelineConfigScope**](PipelineConfigScope.md) |  | 
**WorkspaceId** | **string** |  | 
**Name** | **string** |  | 
**Slug** | **string** |  | 
**PipelineKind** | [**PipelineKind**](PipelineKind.md) |  | 
**Selectors** | [**[]PipelineConfigSelector**](PipelineConfigSelector.md) |  | 
**Sampling** | Pointer to [**NullablePipelineConfigSampling**](PipelineConfigSampling.md) |  | [optional] 
**Enabled** | **bool** |  | 
**Definition** | [**PipelineConfigDefinition**](PipelineConfigDefinition.md) |  | 
**ScorerSlug** | Pointer to **NullableString** |  | [optional] 
**ScorerVersion** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewPipelineConfig

`func NewPipelineConfig(id string, createdAt time.Time, updatedAt time.Time, deletedAt NullableTime, scope PipelineConfigScope, workspaceId string, name string, slug string, pipelineKind PipelineKind, selectors []PipelineConfigSelector, enabled bool, definition PipelineConfigDefinition, ) *PipelineConfig`

NewPipelineConfig instantiates a new PipelineConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineConfigWithDefaults

`func NewPipelineConfigWithDefaults() *PipelineConfig`

NewPipelineConfigWithDefaults instantiates a new PipelineConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PipelineConfig) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PipelineConfig) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PipelineConfig) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *PipelineConfig) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PipelineConfig) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PipelineConfig) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *PipelineConfig) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PipelineConfig) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PipelineConfig) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *PipelineConfig) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *PipelineConfig) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *PipelineConfig) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.


### SetDeletedAtNil

`func (o *PipelineConfig) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *PipelineConfig) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetScope

`func (o *PipelineConfig) GetScope() PipelineConfigScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *PipelineConfig) GetScopeOk() (*PipelineConfigScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *PipelineConfig) SetScope(v PipelineConfigScope)`

SetScope sets Scope field to given value.


### GetWorkspaceId

`func (o *PipelineConfig) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *PipelineConfig) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *PipelineConfig) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetName

`func (o *PipelineConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PipelineConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PipelineConfig) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *PipelineConfig) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *PipelineConfig) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *PipelineConfig) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetPipelineKind

`func (o *PipelineConfig) GetPipelineKind() PipelineKind`

GetPipelineKind returns the PipelineKind field if non-nil, zero value otherwise.

### GetPipelineKindOk

`func (o *PipelineConfig) GetPipelineKindOk() (*PipelineKind, bool)`

GetPipelineKindOk returns a tuple with the PipelineKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineKind

`func (o *PipelineConfig) SetPipelineKind(v PipelineKind)`

SetPipelineKind sets PipelineKind field to given value.


### GetSelectors

`func (o *PipelineConfig) GetSelectors() []PipelineConfigSelector`

GetSelectors returns the Selectors field if non-nil, zero value otherwise.

### GetSelectorsOk

`func (o *PipelineConfig) GetSelectorsOk() (*[]PipelineConfigSelector, bool)`

GetSelectorsOk returns a tuple with the Selectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectors

`func (o *PipelineConfig) SetSelectors(v []PipelineConfigSelector)`

SetSelectors sets Selectors field to given value.


### GetSampling

`func (o *PipelineConfig) GetSampling() PipelineConfigSampling`

GetSampling returns the Sampling field if non-nil, zero value otherwise.

### GetSamplingOk

`func (o *PipelineConfig) GetSamplingOk() (*PipelineConfigSampling, bool)`

GetSamplingOk returns a tuple with the Sampling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampling

`func (o *PipelineConfig) SetSampling(v PipelineConfigSampling)`

SetSampling sets Sampling field to given value.

### HasSampling

`func (o *PipelineConfig) HasSampling() bool`

HasSampling returns a boolean if a field has been set.

### SetSamplingNil

`func (o *PipelineConfig) SetSamplingNil(b bool)`

 SetSamplingNil sets the value for Sampling to be an explicit nil

### UnsetSampling
`func (o *PipelineConfig) UnsetSampling()`

UnsetSampling ensures that no value is present for Sampling, not even an explicit nil
### GetEnabled

`func (o *PipelineConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PipelineConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PipelineConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetDefinition

`func (o *PipelineConfig) GetDefinition() PipelineConfigDefinition`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *PipelineConfig) GetDefinitionOk() (*PipelineConfigDefinition, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *PipelineConfig) SetDefinition(v PipelineConfigDefinition)`

SetDefinition sets Definition field to given value.


### GetScorerSlug

`func (o *PipelineConfig) GetScorerSlug() string`

GetScorerSlug returns the ScorerSlug field if non-nil, zero value otherwise.

### GetScorerSlugOk

`func (o *PipelineConfig) GetScorerSlugOk() (*string, bool)`

GetScorerSlugOk returns a tuple with the ScorerSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScorerSlug

`func (o *PipelineConfig) SetScorerSlug(v string)`

SetScorerSlug sets ScorerSlug field to given value.

### HasScorerSlug

`func (o *PipelineConfig) HasScorerSlug() bool`

HasScorerSlug returns a boolean if a field has been set.

### SetScorerSlugNil

`func (o *PipelineConfig) SetScorerSlugNil(b bool)`

 SetScorerSlugNil sets the value for ScorerSlug to be an explicit nil

### UnsetScorerSlug
`func (o *PipelineConfig) UnsetScorerSlug()`

UnsetScorerSlug ensures that no value is present for ScorerSlug, not even an explicit nil
### GetScorerVersion

`func (o *PipelineConfig) GetScorerVersion() int32`

GetScorerVersion returns the ScorerVersion field if non-nil, zero value otherwise.

### GetScorerVersionOk

`func (o *PipelineConfig) GetScorerVersionOk() (*int32, bool)`

GetScorerVersionOk returns a tuple with the ScorerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScorerVersion

`func (o *PipelineConfig) SetScorerVersion(v int32)`

SetScorerVersion sets ScorerVersion field to given value.

### HasScorerVersion

`func (o *PipelineConfig) HasScorerVersion() bool`

HasScorerVersion returns a boolean if a field has been set.

### SetScorerVersionNil

`func (o *PipelineConfig) SetScorerVersionNil(b bool)`

 SetScorerVersionNil sets the value for ScorerVersion to be an explicit nil

### UnsetScorerVersion
`func (o *PipelineConfig) UnsetScorerVersion()`

UnsetScorerVersion ensures that no value is present for ScorerVersion, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


