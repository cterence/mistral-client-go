# ChatTranscriptionEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AudioUrl** | **string** |  | 
**Model** | **string** |  | 
**ResponseMessage** | **map[string]interface{}** |  | 

## Methods

### NewChatTranscriptionEvent

`func NewChatTranscriptionEvent(audioUrl string, model string, responseMessage map[string]interface{}, ) *ChatTranscriptionEvent`

NewChatTranscriptionEvent instantiates a new ChatTranscriptionEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatTranscriptionEventWithDefaults

`func NewChatTranscriptionEventWithDefaults() *ChatTranscriptionEvent`

NewChatTranscriptionEventWithDefaults instantiates a new ChatTranscriptionEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudioUrl

`func (o *ChatTranscriptionEvent) GetAudioUrl() string`

GetAudioUrl returns the AudioUrl field if non-nil, zero value otherwise.

### GetAudioUrlOk

`func (o *ChatTranscriptionEvent) GetAudioUrlOk() (*string, bool)`

GetAudioUrlOk returns a tuple with the AudioUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioUrl

`func (o *ChatTranscriptionEvent) SetAudioUrl(v string)`

SetAudioUrl sets AudioUrl field to given value.


### GetModel

`func (o *ChatTranscriptionEvent) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ChatTranscriptionEvent) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ChatTranscriptionEvent) SetModel(v string)`

SetModel sets Model field to given value.


### GetResponseMessage

`func (o *ChatTranscriptionEvent) GetResponseMessage() map[string]interface{}`

GetResponseMessage returns the ResponseMessage field if non-nil, zero value otherwise.

### GetResponseMessageOk

`func (o *ChatTranscriptionEvent) GetResponseMessageOk() (*map[string]interface{}, bool)`

GetResponseMessageOk returns a tuple with the ResponseMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseMessage

`func (o *ChatTranscriptionEvent) SetResponseMessage(v map[string]interface{})`

SetResponseMessage sets ResponseMessage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


