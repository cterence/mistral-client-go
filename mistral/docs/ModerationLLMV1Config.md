# ModerationLLMV1Config

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ModelName** | Pointer to **string** | Override model name. Should be omitted in general. | [optional] [default to "mistral-moderation-2411"]
**CustomCategoryThresholds** | Pointer to [**NullableModerationLLMV1CategoryThresholds**](ModerationLLMV1CategoryThresholds.md) |  | [optional] 
**IgnoreOtherCategories** | Pointer to **bool** | If true, only evaluate categories in custom_category_thresholds; others are ignored. | [optional] [default to false]
**Action** | Pointer to [**ModerationLLMAction**](ModerationLLMAction.md) | Action to take if any score is above the threshold for any category. | [optional] [default to MODERATIONLLMACTION_NONE]

## Methods

### NewModerationLLMV1Config

`func NewModerationLLMV1Config() *ModerationLLMV1Config`

NewModerationLLMV1Config instantiates a new ModerationLLMV1Config object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModerationLLMV1ConfigWithDefaults

`func NewModerationLLMV1ConfigWithDefaults() *ModerationLLMV1Config`

NewModerationLLMV1ConfigWithDefaults instantiates a new ModerationLLMV1Config object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModelName

`func (o *ModerationLLMV1Config) GetModelName() string`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *ModerationLLMV1Config) GetModelNameOk() (*string, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *ModerationLLMV1Config) SetModelName(v string)`

SetModelName sets ModelName field to given value.

### HasModelName

`func (o *ModerationLLMV1Config) HasModelName() bool`

HasModelName returns a boolean if a field has been set.

### GetCustomCategoryThresholds

`func (o *ModerationLLMV1Config) GetCustomCategoryThresholds() ModerationLLMV1CategoryThresholds`

GetCustomCategoryThresholds returns the CustomCategoryThresholds field if non-nil, zero value otherwise.

### GetCustomCategoryThresholdsOk

`func (o *ModerationLLMV1Config) GetCustomCategoryThresholdsOk() (*ModerationLLMV1CategoryThresholds, bool)`

GetCustomCategoryThresholdsOk returns a tuple with the CustomCategoryThresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCategoryThresholds

`func (o *ModerationLLMV1Config) SetCustomCategoryThresholds(v ModerationLLMV1CategoryThresholds)`

SetCustomCategoryThresholds sets CustomCategoryThresholds field to given value.

### HasCustomCategoryThresholds

`func (o *ModerationLLMV1Config) HasCustomCategoryThresholds() bool`

HasCustomCategoryThresholds returns a boolean if a field has been set.

### SetCustomCategoryThresholdsNil

`func (o *ModerationLLMV1Config) SetCustomCategoryThresholdsNil(b bool)`

 SetCustomCategoryThresholdsNil sets the value for CustomCategoryThresholds to be an explicit nil

### UnsetCustomCategoryThresholds
`func (o *ModerationLLMV1Config) UnsetCustomCategoryThresholds()`

UnsetCustomCategoryThresholds ensures that no value is present for CustomCategoryThresholds, not even an explicit nil
### GetIgnoreOtherCategories

`func (o *ModerationLLMV1Config) GetIgnoreOtherCategories() bool`

GetIgnoreOtherCategories returns the IgnoreOtherCategories field if non-nil, zero value otherwise.

### GetIgnoreOtherCategoriesOk

`func (o *ModerationLLMV1Config) GetIgnoreOtherCategoriesOk() (*bool, bool)`

GetIgnoreOtherCategoriesOk returns a tuple with the IgnoreOtherCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreOtherCategories

`func (o *ModerationLLMV1Config) SetIgnoreOtherCategories(v bool)`

SetIgnoreOtherCategories sets IgnoreOtherCategories field to given value.

### HasIgnoreOtherCategories

`func (o *ModerationLLMV1Config) HasIgnoreOtherCategories() bool`

HasIgnoreOtherCategories returns a boolean if a field has been set.

### GetAction

`func (o *ModerationLLMV1Config) GetAction() ModerationLLMAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ModerationLLMV1Config) GetActionOk() (*ModerationLLMAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ModerationLLMV1Config) SetAction(v ModerationLLMAction)`

SetAction sets Action field to given value.

### HasAction

`func (o *ModerationLLMV1Config) HasAction() bool`

HasAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


