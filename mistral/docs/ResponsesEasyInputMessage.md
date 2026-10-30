# ResponsesEasyInputMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | [**Content6**](Content6.md) |  | 
**Role** | **string** |  | 
**Phase** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewResponsesEasyInputMessage

`func NewResponsesEasyInputMessage(content Content6, role string, ) *ResponsesEasyInputMessage`

NewResponsesEasyInputMessage instantiates a new ResponsesEasyInputMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsesEasyInputMessageWithDefaults

`func NewResponsesEasyInputMessageWithDefaults() *ResponsesEasyInputMessage`

NewResponsesEasyInputMessageWithDefaults instantiates a new ResponsesEasyInputMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *ResponsesEasyInputMessage) GetContent() Content6`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ResponsesEasyInputMessage) GetContentOk() (*Content6, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ResponsesEasyInputMessage) SetContent(v Content6)`

SetContent sets Content field to given value.


### GetRole

`func (o *ResponsesEasyInputMessage) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ResponsesEasyInputMessage) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ResponsesEasyInputMessage) SetRole(v string)`

SetRole sets Role field to given value.


### GetPhase

`func (o *ResponsesEasyInputMessage) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *ResponsesEasyInputMessage) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *ResponsesEasyInputMessage) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *ResponsesEasyInputMessage) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### SetPhaseNil

`func (o *ResponsesEasyInputMessage) SetPhaseNil(b bool)`

 SetPhaseNil sets the value for Phase to be an explicit nil

### UnsetPhase
`func (o *ResponsesEasyInputMessage) UnsetPhase()`

UnsetPhase ensures that no value is present for Phase, not even an explicit nil
### GetType

`func (o *ResponsesEasyInputMessage) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResponsesEasyInputMessage) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResponsesEasyInputMessage) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ResponsesEasyInputMessage) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *ResponsesEasyInputMessage) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ResponsesEasyInputMessage) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


