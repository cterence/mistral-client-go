# UpdatePipelineRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Selectors** | Pointer to [**[]PipelineConfigSelector**](PipelineConfigSelector.md) |  | [optional] 
**Definitions** | Pointer to [**[]LlmJudgeReference**](LlmJudgeReference.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdatePipelineRequest

`func NewUpdatePipelineRequest() *UpdatePipelineRequest`

NewUpdatePipelineRequest instantiates a new UpdatePipelineRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePipelineRequestWithDefaults

`func NewUpdatePipelineRequestWithDefaults() *UpdatePipelineRequest`

NewUpdatePipelineRequestWithDefaults instantiates a new UpdatePipelineRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdatePipelineRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdatePipelineRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdatePipelineRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdatePipelineRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdatePipelineRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdatePipelineRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdatePipelineRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdatePipelineRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdatePipelineRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdatePipelineRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSelectors

`func (o *UpdatePipelineRequest) GetSelectors() []PipelineConfigSelector`

GetSelectors returns the Selectors field if non-nil, zero value otherwise.

### GetSelectorsOk

`func (o *UpdatePipelineRequest) GetSelectorsOk() (*[]PipelineConfigSelector, bool)`

GetSelectorsOk returns a tuple with the Selectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectors

`func (o *UpdatePipelineRequest) SetSelectors(v []PipelineConfigSelector)`

SetSelectors sets Selectors field to given value.

### HasSelectors

`func (o *UpdatePipelineRequest) HasSelectors() bool`

HasSelectors returns a boolean if a field has been set.

### GetDefinitions

`func (o *UpdatePipelineRequest) GetDefinitions() []LlmJudgeReference`

GetDefinitions returns the Definitions field if non-nil, zero value otherwise.

### GetDefinitionsOk

`func (o *UpdatePipelineRequest) GetDefinitionsOk() (*[]LlmJudgeReference, bool)`

GetDefinitionsOk returns a tuple with the Definitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinitions

`func (o *UpdatePipelineRequest) SetDefinitions(v []LlmJudgeReference)`

SetDefinitions sets Definitions field to given value.

### HasDefinitions

`func (o *UpdatePipelineRequest) HasDefinitions() bool`

HasDefinitions returns a boolean if a field has been set.

### GetEnabled

`func (o *UpdatePipelineRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdatePipelineRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdatePipelineRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdatePipelineRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


