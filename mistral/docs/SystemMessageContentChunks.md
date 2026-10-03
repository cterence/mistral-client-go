# SystemMessageContentChunks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "text"]
**Text** | **string** |  | 
**Thinking** | [**[]ThinkingInner**](ThinkingInner.md) |  | 
**Closed** | Pointer to **bool** | Whether the thinking chunk is closed or not. Currently only used for prefixing. | [optional] [default to true]

## Methods

### NewSystemMessageContentChunks

`func NewSystemMessageContentChunks(text string, thinking []ThinkingInner, ) *SystemMessageContentChunks`

NewSystemMessageContentChunks instantiates a new SystemMessageContentChunks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemMessageContentChunksWithDefaults

`func NewSystemMessageContentChunksWithDefaults() *SystemMessageContentChunks`

NewSystemMessageContentChunksWithDefaults instantiates a new SystemMessageContentChunks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SystemMessageContentChunks) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SystemMessageContentChunks) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SystemMessageContentChunks) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SystemMessageContentChunks) HasType() bool`

HasType returns a boolean if a field has been set.

### GetText

`func (o *SystemMessageContentChunks) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *SystemMessageContentChunks) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *SystemMessageContentChunks) SetText(v string)`

SetText sets Text field to given value.


### GetThinking

`func (o *SystemMessageContentChunks) GetThinking() []ThinkingInner`

GetThinking returns the Thinking field if non-nil, zero value otherwise.

### GetThinkingOk

`func (o *SystemMessageContentChunks) GetThinkingOk() (*[]ThinkingInner, bool)`

GetThinkingOk returns a tuple with the Thinking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThinking

`func (o *SystemMessageContentChunks) SetThinking(v []ThinkingInner)`

SetThinking sets Thinking field to given value.


### GetClosed

`func (o *SystemMessageContentChunks) GetClosed() bool`

GetClosed returns the Closed field if non-nil, zero value otherwise.

### GetClosedOk

`func (o *SystemMessageContentChunks) GetClosedOk() (*bool, bool)`

GetClosedOk returns a tuple with the Closed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosed

`func (o *SystemMessageContentChunks) SetClosed(v bool)`

SetClosed sets Closed field to given value.

### HasClosed

`func (o *SystemMessageContentChunks) HasClosed() bool`

HasClosed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


