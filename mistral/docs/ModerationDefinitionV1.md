# ModerationDefinitionV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | Pointer to **string** |  | [optional] [default to "mistral-moderation-latest"]
**TargetAttributes** | Pointer to **[]string** |  | [optional] [default to {"gen_ai.input.messages", "gen_ai.output.messages"}]

## Methods

### NewModerationDefinitionV1

`func NewModerationDefinitionV1() *ModerationDefinitionV1`

NewModerationDefinitionV1 instantiates a new ModerationDefinitionV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModerationDefinitionV1WithDefaults

`func NewModerationDefinitionV1WithDefaults() *ModerationDefinitionV1`

NewModerationDefinitionV1WithDefaults instantiates a new ModerationDefinitionV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *ModerationDefinitionV1) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ModerationDefinitionV1) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ModerationDefinitionV1) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ModerationDefinitionV1) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetTargetAttributes

`func (o *ModerationDefinitionV1) GetTargetAttributes() []string`

GetTargetAttributes returns the TargetAttributes field if non-nil, zero value otherwise.

### GetTargetAttributesOk

`func (o *ModerationDefinitionV1) GetTargetAttributesOk() (*[]string, bool)`

GetTargetAttributesOk returns a tuple with the TargetAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAttributes

`func (o *ModerationDefinitionV1) SetTargetAttributes(v []string)`

SetTargetAttributes sets TargetAttributes field to given value.

### HasTargetAttributes

`func (o *ModerationDefinitionV1) HasTargetAttributes() bool`

HasTargetAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


