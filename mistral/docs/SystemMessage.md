# SystemMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | Pointer to **string** |  | [optional] [default to "system"]
**Content** | [**Content5**](Content5.md) |  | 

## Methods

### NewSystemMessage

`func NewSystemMessage(content Content5, ) *SystemMessage`

NewSystemMessage instantiates a new SystemMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemMessageWithDefaults

`func NewSystemMessageWithDefaults() *SystemMessage`

NewSystemMessageWithDefaults instantiates a new SystemMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *SystemMessage) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *SystemMessage) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *SystemMessage) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *SystemMessage) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetContent

`func (o *SystemMessage) GetContent() Content5`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *SystemMessage) GetContentOk() (*Content5, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *SystemMessage) SetContent(v Content5)`

SetContent sets Content field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


