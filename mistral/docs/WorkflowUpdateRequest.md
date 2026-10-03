# WorkflowUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **NullableString** | New display name value | [optional] 
**Description** | Pointer to **NullableString** | New description value | [optional] 
**AvailableInChatAssistant** | Pointer to **NullableBool** | Whether to make the workflow available in the chat assistant | [optional] 

## Methods

### NewWorkflowUpdateRequest

`func NewWorkflowUpdateRequest() *WorkflowUpdateRequest`

NewWorkflowUpdateRequest instantiates a new WorkflowUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowUpdateRequestWithDefaults

`func NewWorkflowUpdateRequestWithDefaults() *WorkflowUpdateRequest`

NewWorkflowUpdateRequestWithDefaults instantiates a new WorkflowUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *WorkflowUpdateRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *WorkflowUpdateRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *WorkflowUpdateRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *WorkflowUpdateRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *WorkflowUpdateRequest) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *WorkflowUpdateRequest) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetDescription

`func (o *WorkflowUpdateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowUpdateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowUpdateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowUpdateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *WorkflowUpdateRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *WorkflowUpdateRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetAvailableInChatAssistant

`func (o *WorkflowUpdateRequest) GetAvailableInChatAssistant() bool`

GetAvailableInChatAssistant returns the AvailableInChatAssistant field if non-nil, zero value otherwise.

### GetAvailableInChatAssistantOk

`func (o *WorkflowUpdateRequest) GetAvailableInChatAssistantOk() (*bool, bool)`

GetAvailableInChatAssistantOk returns a tuple with the AvailableInChatAssistant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableInChatAssistant

`func (o *WorkflowUpdateRequest) SetAvailableInChatAssistant(v bool)`

SetAvailableInChatAssistant sets AvailableInChatAssistant field to given value.

### HasAvailableInChatAssistant

`func (o *WorkflowUpdateRequest) HasAvailableInChatAssistant() bool`

HasAvailableInChatAssistant returns a boolean if a field has been set.

### SetAvailableInChatAssistantNil

`func (o *WorkflowUpdateRequest) SetAvailableInChatAssistantNil(b bool)`

 SetAvailableInChatAssistantNil sets the value for AvailableInChatAssistant to be an explicit nil

### UnsetAvailableInChatAssistant
`func (o *WorkflowUpdateRequest) UnsetAvailableInChatAssistant()`

UnsetAvailableInChatAssistant ensures that no value is present for AvailableInChatAssistant, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


