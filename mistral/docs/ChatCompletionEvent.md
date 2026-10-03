# ChatCompletionEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | **string** |  | 
**CorrelationId** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**ExtraFields** | [**map[string]ExtraFieldsValue**](ExtraFieldsValue.md) |  | 
**NbInputTokens** | **int32** |  | 
**NbOutputTokens** | **int32** |  | 
**EnabledTools** | **[]map[string]interface{}** |  | 
**RequestMessages** | **[]map[string]interface{}** |  | 
**ResponseMessages** | **[]map[string]interface{}** |  | 
**NbMessages** | **int32** |  | 
**ChatTranscriptionEvents** | [**[]ChatTranscriptionEvent**](ChatTranscriptionEvent.md) |  | 

## Methods

### NewChatCompletionEvent

`func NewChatCompletionEvent(eventId string, correlationId string, createdAt time.Time, extraFields map[string]ExtraFieldsValue, nbInputTokens int32, nbOutputTokens int32, enabledTools []map[string]interface{}, requestMessages []map[string]interface{}, responseMessages []map[string]interface{}, nbMessages int32, chatTranscriptionEvents []ChatTranscriptionEvent, ) *ChatCompletionEvent`

NewChatCompletionEvent instantiates a new ChatCompletionEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatCompletionEventWithDefaults

`func NewChatCompletionEventWithDefaults() *ChatCompletionEvent`

NewChatCompletionEventWithDefaults instantiates a new ChatCompletionEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventId

`func (o *ChatCompletionEvent) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *ChatCompletionEvent) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *ChatCompletionEvent) SetEventId(v string)`

SetEventId sets EventId field to given value.


### GetCorrelationId

`func (o *ChatCompletionEvent) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *ChatCompletionEvent) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *ChatCompletionEvent) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.


### GetCreatedAt

`func (o *ChatCompletionEvent) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ChatCompletionEvent) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ChatCompletionEvent) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetExtraFields

`func (o *ChatCompletionEvent) GetExtraFields() map[string]ExtraFieldsValue`

GetExtraFields returns the ExtraFields field if non-nil, zero value otherwise.

### GetExtraFieldsOk

`func (o *ChatCompletionEvent) GetExtraFieldsOk() (*map[string]ExtraFieldsValue, bool)`

GetExtraFieldsOk returns a tuple with the ExtraFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraFields

`func (o *ChatCompletionEvent) SetExtraFields(v map[string]ExtraFieldsValue)`

SetExtraFields sets ExtraFields field to given value.


### GetNbInputTokens

`func (o *ChatCompletionEvent) GetNbInputTokens() int32`

GetNbInputTokens returns the NbInputTokens field if non-nil, zero value otherwise.

### GetNbInputTokensOk

`func (o *ChatCompletionEvent) GetNbInputTokensOk() (*int32, bool)`

GetNbInputTokensOk returns a tuple with the NbInputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNbInputTokens

`func (o *ChatCompletionEvent) SetNbInputTokens(v int32)`

SetNbInputTokens sets NbInputTokens field to given value.


### GetNbOutputTokens

`func (o *ChatCompletionEvent) GetNbOutputTokens() int32`

GetNbOutputTokens returns the NbOutputTokens field if non-nil, zero value otherwise.

### GetNbOutputTokensOk

`func (o *ChatCompletionEvent) GetNbOutputTokensOk() (*int32, bool)`

GetNbOutputTokensOk returns a tuple with the NbOutputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNbOutputTokens

`func (o *ChatCompletionEvent) SetNbOutputTokens(v int32)`

SetNbOutputTokens sets NbOutputTokens field to given value.


### GetEnabledTools

`func (o *ChatCompletionEvent) GetEnabledTools() []map[string]interface{}`

GetEnabledTools returns the EnabledTools field if non-nil, zero value otherwise.

### GetEnabledToolsOk

`func (o *ChatCompletionEvent) GetEnabledToolsOk() (*[]map[string]interface{}, bool)`

GetEnabledToolsOk returns a tuple with the EnabledTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledTools

`func (o *ChatCompletionEvent) SetEnabledTools(v []map[string]interface{})`

SetEnabledTools sets EnabledTools field to given value.


### GetRequestMessages

`func (o *ChatCompletionEvent) GetRequestMessages() []map[string]interface{}`

GetRequestMessages returns the RequestMessages field if non-nil, zero value otherwise.

### GetRequestMessagesOk

`func (o *ChatCompletionEvent) GetRequestMessagesOk() (*[]map[string]interface{}, bool)`

GetRequestMessagesOk returns a tuple with the RequestMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestMessages

`func (o *ChatCompletionEvent) SetRequestMessages(v []map[string]interface{})`

SetRequestMessages sets RequestMessages field to given value.


### GetResponseMessages

`func (o *ChatCompletionEvent) GetResponseMessages() []map[string]interface{}`

GetResponseMessages returns the ResponseMessages field if non-nil, zero value otherwise.

### GetResponseMessagesOk

`func (o *ChatCompletionEvent) GetResponseMessagesOk() (*[]map[string]interface{}, bool)`

GetResponseMessagesOk returns a tuple with the ResponseMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseMessages

`func (o *ChatCompletionEvent) SetResponseMessages(v []map[string]interface{})`

SetResponseMessages sets ResponseMessages field to given value.


### GetNbMessages

`func (o *ChatCompletionEvent) GetNbMessages() int32`

GetNbMessages returns the NbMessages field if non-nil, zero value otherwise.

### GetNbMessagesOk

`func (o *ChatCompletionEvent) GetNbMessagesOk() (*int32, bool)`

GetNbMessagesOk returns a tuple with the NbMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNbMessages

`func (o *ChatCompletionEvent) SetNbMessages(v int32)`

SetNbMessages sets NbMessages field to given value.


### GetChatTranscriptionEvents

`func (o *ChatCompletionEvent) GetChatTranscriptionEvents() []ChatTranscriptionEvent`

GetChatTranscriptionEvents returns the ChatTranscriptionEvents field if non-nil, zero value otherwise.

### GetChatTranscriptionEventsOk

`func (o *ChatCompletionEvent) GetChatTranscriptionEventsOk() (*[]ChatTranscriptionEvent, bool)`

GetChatTranscriptionEventsOk returns a tuple with the ChatTranscriptionEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatTranscriptionEvents

`func (o *ChatCompletionEvent) SetChatTranscriptionEvents(v []ChatTranscriptionEvent)`

SetChatTranscriptionEvents sets ChatTranscriptionEvents field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


