# ChatCompletionEventPreview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | **string** |  | 
**CorrelationId** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**ExtraFields** | [**map[string]ExtraFieldsValue**](ExtraFieldsValue.md) |  | 
**NbInputTokens** | **int32** |  | 
**NbOutputTokens** | **int32** |  | 

## Methods

### NewChatCompletionEventPreview

`func NewChatCompletionEventPreview(eventId string, correlationId string, createdAt time.Time, extraFields map[string]ExtraFieldsValue, nbInputTokens int32, nbOutputTokens int32, ) *ChatCompletionEventPreview`

NewChatCompletionEventPreview instantiates a new ChatCompletionEventPreview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatCompletionEventPreviewWithDefaults

`func NewChatCompletionEventPreviewWithDefaults() *ChatCompletionEventPreview`

NewChatCompletionEventPreviewWithDefaults instantiates a new ChatCompletionEventPreview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventId

`func (o *ChatCompletionEventPreview) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *ChatCompletionEventPreview) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *ChatCompletionEventPreview) SetEventId(v string)`

SetEventId sets EventId field to given value.


### GetCorrelationId

`func (o *ChatCompletionEventPreview) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *ChatCompletionEventPreview) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *ChatCompletionEventPreview) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.


### GetCreatedAt

`func (o *ChatCompletionEventPreview) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ChatCompletionEventPreview) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ChatCompletionEventPreview) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetExtraFields

`func (o *ChatCompletionEventPreview) GetExtraFields() map[string]ExtraFieldsValue`

GetExtraFields returns the ExtraFields field if non-nil, zero value otherwise.

### GetExtraFieldsOk

`func (o *ChatCompletionEventPreview) GetExtraFieldsOk() (*map[string]ExtraFieldsValue, bool)`

GetExtraFieldsOk returns a tuple with the ExtraFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraFields

`func (o *ChatCompletionEventPreview) SetExtraFields(v map[string]ExtraFieldsValue)`

SetExtraFields sets ExtraFields field to given value.


### GetNbInputTokens

`func (o *ChatCompletionEventPreview) GetNbInputTokens() int32`

GetNbInputTokens returns the NbInputTokens field if non-nil, zero value otherwise.

### GetNbInputTokensOk

`func (o *ChatCompletionEventPreview) GetNbInputTokensOk() (*int32, bool)`

GetNbInputTokensOk returns a tuple with the NbInputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNbInputTokens

`func (o *ChatCompletionEventPreview) SetNbInputTokens(v int32)`

SetNbInputTokens sets NbInputTokens field to given value.


### GetNbOutputTokens

`func (o *ChatCompletionEventPreview) GetNbOutputTokens() int32`

GetNbOutputTokens returns the NbOutputTokens field if non-nil, zero value otherwise.

### GetNbOutputTokensOk

`func (o *ChatCompletionEventPreview) GetNbOutputTokensOk() (*int32, bool)`

GetNbOutputTokensOk returns a tuple with the NbOutputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNbOutputTokens

`func (o *ChatCompletionEventPreview) SetNbOutputTokens(v int32)`

SetNbOutputTokens sets NbOutputTokens field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


