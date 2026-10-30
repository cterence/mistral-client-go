# SkillDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Model-facing trigger and usage description. | [optional] 
**Body** | Pointer to **string** | Skill body content. | [optional] 
**Assets** | Pointer to [**map[string]SkillAssetContent**](SkillAssetContent.md) | Additional files available to the skill. | [optional] 

## Methods

### NewSkillDefinition

`func NewSkillDefinition() *SkillDefinition`

NewSkillDefinition instantiates a new SkillDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSkillDefinitionWithDefaults

`func NewSkillDefinitionWithDefaults() *SkillDefinition`

NewSkillDefinitionWithDefaults instantiates a new SkillDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *SkillDefinition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SkillDefinition) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SkillDefinition) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SkillDefinition) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetBody

`func (o *SkillDefinition) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *SkillDefinition) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *SkillDefinition) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *SkillDefinition) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetAssets

`func (o *SkillDefinition) GetAssets() map[string]SkillAssetContent`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *SkillDefinition) GetAssetsOk() (*map[string]SkillAssetContent, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *SkillDefinition) SetAssets(v map[string]SkillAssetContent)`

SetAssets sets Assets field to given value.

### HasAssets

`func (o *SkillDefinition) HasAssets() bool`

HasAssets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


