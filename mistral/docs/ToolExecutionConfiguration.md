# ToolExecutionConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequiresConfirmation** | Pointer to [**NullableRequiresConfirmation**](RequiresConfirmation.md) |  | [optional] 
**SkipConfirmation** | Pointer to [**NullableSkipConfirmation**](SkipConfirmation.md) |  | [optional] 
**Include** | Pointer to **[]string** |  | [optional] 
**Exclude** | Pointer to **[]string** |  | [optional] 

## Methods

### NewToolExecutionConfiguration

`func NewToolExecutionConfiguration() *ToolExecutionConfiguration`

NewToolExecutionConfiguration instantiates a new ToolExecutionConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewToolExecutionConfigurationWithDefaults

`func NewToolExecutionConfigurationWithDefaults() *ToolExecutionConfiguration`

NewToolExecutionConfigurationWithDefaults instantiates a new ToolExecutionConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequiresConfirmation

`func (o *ToolExecutionConfiguration) GetRequiresConfirmation() RequiresConfirmation`

GetRequiresConfirmation returns the RequiresConfirmation field if non-nil, zero value otherwise.

### GetRequiresConfirmationOk

`func (o *ToolExecutionConfiguration) GetRequiresConfirmationOk() (*RequiresConfirmation, bool)`

GetRequiresConfirmationOk returns a tuple with the RequiresConfirmation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresConfirmation

`func (o *ToolExecutionConfiguration) SetRequiresConfirmation(v RequiresConfirmation)`

SetRequiresConfirmation sets RequiresConfirmation field to given value.

### HasRequiresConfirmation

`func (o *ToolExecutionConfiguration) HasRequiresConfirmation() bool`

HasRequiresConfirmation returns a boolean if a field has been set.

### SetRequiresConfirmationNil

`func (o *ToolExecutionConfiguration) SetRequiresConfirmationNil(b bool)`

 SetRequiresConfirmationNil sets the value for RequiresConfirmation to be an explicit nil

### UnsetRequiresConfirmation
`func (o *ToolExecutionConfiguration) UnsetRequiresConfirmation()`

UnsetRequiresConfirmation ensures that no value is present for RequiresConfirmation, not even an explicit nil
### GetSkipConfirmation

`func (o *ToolExecutionConfiguration) GetSkipConfirmation() SkipConfirmation`

GetSkipConfirmation returns the SkipConfirmation field if non-nil, zero value otherwise.

### GetSkipConfirmationOk

`func (o *ToolExecutionConfiguration) GetSkipConfirmationOk() (*SkipConfirmation, bool)`

GetSkipConfirmationOk returns a tuple with the SkipConfirmation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipConfirmation

`func (o *ToolExecutionConfiguration) SetSkipConfirmation(v SkipConfirmation)`

SetSkipConfirmation sets SkipConfirmation field to given value.

### HasSkipConfirmation

`func (o *ToolExecutionConfiguration) HasSkipConfirmation() bool`

HasSkipConfirmation returns a boolean if a field has been set.

### SetSkipConfirmationNil

`func (o *ToolExecutionConfiguration) SetSkipConfirmationNil(b bool)`

 SetSkipConfirmationNil sets the value for SkipConfirmation to be an explicit nil

### UnsetSkipConfirmation
`func (o *ToolExecutionConfiguration) UnsetSkipConfirmation()`

UnsetSkipConfirmation ensures that no value is present for SkipConfirmation, not even an explicit nil
### GetInclude

`func (o *ToolExecutionConfiguration) GetInclude() []string`

GetInclude returns the Include field if non-nil, zero value otherwise.

### GetIncludeOk

`func (o *ToolExecutionConfiguration) GetIncludeOk() (*[]string, bool)`

GetIncludeOk returns a tuple with the Include field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInclude

`func (o *ToolExecutionConfiguration) SetInclude(v []string)`

SetInclude sets Include field to given value.

### HasInclude

`func (o *ToolExecutionConfiguration) HasInclude() bool`

HasInclude returns a boolean if a field has been set.

### SetIncludeNil

`func (o *ToolExecutionConfiguration) SetIncludeNil(b bool)`

 SetIncludeNil sets the value for Include to be an explicit nil

### UnsetInclude
`func (o *ToolExecutionConfiguration) UnsetInclude()`

UnsetInclude ensures that no value is present for Include, not even an explicit nil
### GetExclude

`func (o *ToolExecutionConfiguration) GetExclude() []string`

GetExclude returns the Exclude field if non-nil, zero value otherwise.

### GetExcludeOk

`func (o *ToolExecutionConfiguration) GetExcludeOk() (*[]string, bool)`

GetExcludeOk returns a tuple with the Exclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclude

`func (o *ToolExecutionConfiguration) SetExclude(v []string)`

SetExclude sets Exclude field to given value.

### HasExclude

`func (o *ToolExecutionConfiguration) HasExclude() bool`

HasExclude returns a boolean if a field has been set.

### SetExcludeNil

`func (o *ToolExecutionConfiguration) SetExcludeNil(b bool)`

 SetExcludeNil sets the value for Exclude to be an explicit nil

### UnsetExclude
`func (o *ToolExecutionConfiguration) UnsetExclude()`

UnsetExclude ensures that no value is present for Exclude, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


