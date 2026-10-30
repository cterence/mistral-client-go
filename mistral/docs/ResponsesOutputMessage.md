# ResponsesOutputMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Content** | [**[]Content7Inner**](Content7Inner.md) |  | 
**Role** | **string** |  | 
**Status** | **string** |  | 
**Type** | **string** |  | 
**Phase** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewResponsesOutputMessage

`func NewResponsesOutputMessage(id string, content []Content7Inner, role string, status string, type_ string, ) *ResponsesOutputMessage`

NewResponsesOutputMessage instantiates a new ResponsesOutputMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsesOutputMessageWithDefaults

`func NewResponsesOutputMessageWithDefaults() *ResponsesOutputMessage`

NewResponsesOutputMessageWithDefaults instantiates a new ResponsesOutputMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResponsesOutputMessage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResponsesOutputMessage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResponsesOutputMessage) SetId(v string)`

SetId sets Id field to given value.


### GetContent

`func (o *ResponsesOutputMessage) GetContent() []Content7Inner`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ResponsesOutputMessage) GetContentOk() (*[]Content7Inner, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ResponsesOutputMessage) SetContent(v []Content7Inner)`

SetContent sets Content field to given value.


### GetRole

`func (o *ResponsesOutputMessage) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ResponsesOutputMessage) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ResponsesOutputMessage) SetRole(v string)`

SetRole sets Role field to given value.


### GetStatus

`func (o *ResponsesOutputMessage) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResponsesOutputMessage) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResponsesOutputMessage) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetType

`func (o *ResponsesOutputMessage) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResponsesOutputMessage) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResponsesOutputMessage) SetType(v string)`

SetType sets Type field to given value.


### GetPhase

`func (o *ResponsesOutputMessage) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *ResponsesOutputMessage) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *ResponsesOutputMessage) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *ResponsesOutputMessage) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### SetPhaseNil

`func (o *ResponsesOutputMessage) SetPhaseNil(b bool)`

 SetPhaseNil sets the value for Phase to be an explicit nil

### UnsetPhase
`func (o *ResponsesOutputMessage) UnsetPhase()`

UnsetPhase ensures that no value is present for Phase, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


