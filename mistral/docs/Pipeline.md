# Pipeline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**WorkspaceId** | **string** |  | 
**Name** | **string** |  | 
**Slug** | **string** |  | 
**Description** | **NullableString** |  | 
**Selectors** | [**[]PipelineConfigSelector**](PipelineConfigSelector.md) |  | 
**Enabled** | **bool** |  | 
**CreatorId** | **NullableString** |  | 
**PipelineConfigs** | [**[]PipelineConfig**](PipelineConfig.md) |  | 

## Methods

### NewPipeline

`func NewPipeline(id string, createdAt time.Time, updatedAt time.Time, workspaceId string, name string, slug string, description NullableString, selectors []PipelineConfigSelector, enabled bool, creatorId NullableString, pipelineConfigs []PipelineConfig, ) *Pipeline`

NewPipeline instantiates a new Pipeline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineWithDefaults

`func NewPipelineWithDefaults() *Pipeline`

NewPipelineWithDefaults instantiates a new Pipeline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Pipeline) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Pipeline) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Pipeline) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *Pipeline) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Pipeline) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Pipeline) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Pipeline) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Pipeline) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Pipeline) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetWorkspaceId

`func (o *Pipeline) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *Pipeline) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *Pipeline) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetName

`func (o *Pipeline) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Pipeline) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Pipeline) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *Pipeline) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *Pipeline) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *Pipeline) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetDescription

`func (o *Pipeline) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Pipeline) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Pipeline) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *Pipeline) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Pipeline) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSelectors

`func (o *Pipeline) GetSelectors() []PipelineConfigSelector`

GetSelectors returns the Selectors field if non-nil, zero value otherwise.

### GetSelectorsOk

`func (o *Pipeline) GetSelectorsOk() (*[]PipelineConfigSelector, bool)`

GetSelectorsOk returns a tuple with the Selectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectors

`func (o *Pipeline) SetSelectors(v []PipelineConfigSelector)`

SetSelectors sets Selectors field to given value.


### GetEnabled

`func (o *Pipeline) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Pipeline) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Pipeline) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetCreatorId

`func (o *Pipeline) GetCreatorId() string`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *Pipeline) GetCreatorIdOk() (*string, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *Pipeline) SetCreatorId(v string)`

SetCreatorId sets CreatorId field to given value.


### SetCreatorIdNil

`func (o *Pipeline) SetCreatorIdNil(b bool)`

 SetCreatorIdNil sets the value for CreatorId to be an explicit nil

### UnsetCreatorId
`func (o *Pipeline) UnsetCreatorId()`

UnsetCreatorId ensures that no value is present for CreatorId, not even an explicit nil
### GetPipelineConfigs

`func (o *Pipeline) GetPipelineConfigs() []PipelineConfig`

GetPipelineConfigs returns the PipelineConfigs field if non-nil, zero value otherwise.

### GetPipelineConfigsOk

`func (o *Pipeline) GetPipelineConfigsOk() (*[]PipelineConfig, bool)`

GetPipelineConfigsOk returns a tuple with the PipelineConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineConfigs

`func (o *Pipeline) SetPipelineConfigs(v []PipelineConfig)`

SetPipelineConfigs sets PipelineConfigs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


