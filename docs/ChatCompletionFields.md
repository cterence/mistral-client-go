# ChatCompletionFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldDefinitions** | [**[]BaseFieldDefinition**](BaseFieldDefinition.md) |  | 
**FieldGroups** | [**[]FieldGroup**](FieldGroup.md) |  | 

## Methods

### NewChatCompletionFields

`func NewChatCompletionFields(fieldDefinitions []BaseFieldDefinition, fieldGroups []FieldGroup, ) *ChatCompletionFields`

NewChatCompletionFields instantiates a new ChatCompletionFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatCompletionFieldsWithDefaults

`func NewChatCompletionFieldsWithDefaults() *ChatCompletionFields`

NewChatCompletionFieldsWithDefaults instantiates a new ChatCompletionFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldDefinitions

`func (o *ChatCompletionFields) GetFieldDefinitions() []BaseFieldDefinition`

GetFieldDefinitions returns the FieldDefinitions field if non-nil, zero value otherwise.

### GetFieldDefinitionsOk

`func (o *ChatCompletionFields) GetFieldDefinitionsOk() (*[]BaseFieldDefinition, bool)`

GetFieldDefinitionsOk returns a tuple with the FieldDefinitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldDefinitions

`func (o *ChatCompletionFields) SetFieldDefinitions(v []BaseFieldDefinition)`

SetFieldDefinitions sets FieldDefinitions field to given value.


### GetFieldGroups

`func (o *ChatCompletionFields) GetFieldGroups() []FieldGroup`

GetFieldGroups returns the FieldGroups field if non-nil, zero value otherwise.

### GetFieldGroupsOk

`func (o *ChatCompletionFields) GetFieldGroupsOk() (*[]FieldGroup, bool)`

GetFieldGroupsOk returns a tuple with the FieldGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldGroups

`func (o *ChatCompletionFields) SetFieldGroups(v []FieldGroup)`

SetFieldGroups sets FieldGroups field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


