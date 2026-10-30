# ResponsesMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | [**[]Content6AnyOfInner**](Content6AnyOfInner.md) |  | 
**Role** | **string** |  | 
**Status** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewResponsesMessage

`func NewResponsesMessage(content []Content6AnyOfInner, role string, ) *ResponsesMessage`

NewResponsesMessage instantiates a new ResponsesMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsesMessageWithDefaults

`func NewResponsesMessageWithDefaults() *ResponsesMessage`

NewResponsesMessageWithDefaults instantiates a new ResponsesMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *ResponsesMessage) GetContent() []Content6AnyOfInner`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ResponsesMessage) GetContentOk() (*[]Content6AnyOfInner, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ResponsesMessage) SetContent(v []Content6AnyOfInner)`

SetContent sets Content field to given value.


### GetRole

`func (o *ResponsesMessage) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ResponsesMessage) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ResponsesMessage) SetRole(v string)`

SetRole sets Role field to given value.


### GetStatus

`func (o *ResponsesMessage) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResponsesMessage) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResponsesMessage) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ResponsesMessage) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ResponsesMessage) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ResponsesMessage) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetType

`func (o *ResponsesMessage) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResponsesMessage) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResponsesMessage) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ResponsesMessage) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *ResponsesMessage) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ResponsesMessage) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


