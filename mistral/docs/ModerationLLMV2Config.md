# ModerationLLMV2Config

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ModelName** | Pointer to **string** | Override model name. Should be omitted in general. | [optional] [default to "mistral-moderation-2603"]
**CustomCategoryThresholds** | Pointer to [**NullableModerationLLMV2CategoryThresholds**](ModerationLLMV2CategoryThresholds.md) |  | [optional] 
**IgnoreOtherCategories** | Pointer to **bool** | If true, only evaluate categories in custom_category_thresholds; others are ignored. | [optional] [default to false]
**Action** | Pointer to [**ModerationLLMAction**](ModerationLLMAction.md) | Action to take if any score is above the threshold for any category. | [optional] [default to MODERATIONLLMACTION_NONE]

## Methods

### NewModerationLLMV2Config

`func NewModerationLLMV2Config() *ModerationLLMV2Config`

NewModerationLLMV2Config instantiates a new ModerationLLMV2Config object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModerationLLMV2ConfigWithDefaults

`func NewModerationLLMV2ConfigWithDefaults() *ModerationLLMV2Config`

NewModerationLLMV2ConfigWithDefaults instantiates a new ModerationLLMV2Config object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModelName

`func (o *ModerationLLMV2Config) GetModelName() string`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *ModerationLLMV2Config) GetModelNameOk() (*string, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *ModerationLLMV2Config) SetModelName(v string)`

SetModelName sets ModelName field to given value.

### HasModelName

`func (o *ModerationLLMV2Config) HasModelName() bool`

HasModelName returns a boolean if a field has been set.

### GetCustomCategoryThresholds

`func (o *ModerationLLMV2Config) GetCustomCategoryThresholds() ModerationLLMV2CategoryThresholds`

GetCustomCategoryThresholds returns the CustomCategoryThresholds field if non-nil, zero value otherwise.

### GetCustomCategoryThresholdsOk

`func (o *ModerationLLMV2Config) GetCustomCategoryThresholdsOk() (*ModerationLLMV2CategoryThresholds, bool)`

GetCustomCategoryThresholdsOk returns a tuple with the CustomCategoryThresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCategoryThresholds

`func (o *ModerationLLMV2Config) SetCustomCategoryThresholds(v ModerationLLMV2CategoryThresholds)`

SetCustomCategoryThresholds sets CustomCategoryThresholds field to given value.

### HasCustomCategoryThresholds

`func (o *ModerationLLMV2Config) HasCustomCategoryThresholds() bool`

HasCustomCategoryThresholds returns a boolean if a field has been set.

### SetCustomCategoryThresholdsNil

`func (o *ModerationLLMV2Config) SetCustomCategoryThresholdsNil(b bool)`

 SetCustomCategoryThresholdsNil sets the value for CustomCategoryThresholds to be an explicit nil

### UnsetCustomCategoryThresholds
`func (o *ModerationLLMV2Config) UnsetCustomCategoryThresholds()`

UnsetCustomCategoryThresholds ensures that no value is present for CustomCategoryThresholds, not even an explicit nil
### GetIgnoreOtherCategories

`func (o *ModerationLLMV2Config) GetIgnoreOtherCategories() bool`

GetIgnoreOtherCategories returns the IgnoreOtherCategories field if non-nil, zero value otherwise.

### GetIgnoreOtherCategoriesOk

`func (o *ModerationLLMV2Config) GetIgnoreOtherCategoriesOk() (*bool, bool)`

GetIgnoreOtherCategoriesOk returns a tuple with the IgnoreOtherCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreOtherCategories

`func (o *ModerationLLMV2Config) SetIgnoreOtherCategories(v bool)`

SetIgnoreOtherCategories sets IgnoreOtherCategories field to given value.

### HasIgnoreOtherCategories

`func (o *ModerationLLMV2Config) HasIgnoreOtherCategories() bool`

HasIgnoreOtherCategories returns a boolean if a field has been set.

### GetAction

`func (o *ModerationLLMV2Config) GetAction() ModerationLLMAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ModerationLLMV2Config) GetActionOk() (*ModerationLLMAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ModerationLLMV2Config) SetAction(v ModerationLLMAction)`

SetAction sets Action field to given value.

### HasAction

`func (o *ModerationLLMV2Config) HasAction() bool`

HasAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


