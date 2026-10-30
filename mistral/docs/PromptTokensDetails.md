# PromptTokensDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Messages** | Pointer to [**[]MessageTokens**](MessageTokens.md) |  | [optional] 
**CachedTokens** | Pointer to **int32** |  | [optional] [default to 0]
**AudioTokens** | Pointer to **int32** |  | [optional] [default to 0]

## Methods

### NewPromptTokensDetails

`func NewPromptTokensDetails() *PromptTokensDetails`

NewPromptTokensDetails instantiates a new PromptTokensDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPromptTokensDetailsWithDefaults

`func NewPromptTokensDetailsWithDefaults() *PromptTokensDetails`

NewPromptTokensDetailsWithDefaults instantiates a new PromptTokensDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessages

`func (o *PromptTokensDetails) GetMessages() []MessageTokens`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *PromptTokensDetails) GetMessagesOk() (*[]MessageTokens, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *PromptTokensDetails) SetMessages(v []MessageTokens)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *PromptTokensDetails) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetCachedTokens

`func (o *PromptTokensDetails) GetCachedTokens() int32`

GetCachedTokens returns the CachedTokens field if non-nil, zero value otherwise.

### GetCachedTokensOk

`func (o *PromptTokensDetails) GetCachedTokensOk() (*int32, bool)`

GetCachedTokensOk returns a tuple with the CachedTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCachedTokens

`func (o *PromptTokensDetails) SetCachedTokens(v int32)`

SetCachedTokens sets CachedTokens field to given value.

### HasCachedTokens

`func (o *PromptTokensDetails) HasCachedTokens() bool`

HasCachedTokens returns a boolean if a field has been set.

### GetAudioTokens

`func (o *PromptTokensDetails) GetAudioTokens() int32`

GetAudioTokens returns the AudioTokens field if non-nil, zero value otherwise.

### GetAudioTokensOk

`func (o *PromptTokensDetails) GetAudioTokensOk() (*int32, bool)`

GetAudioTokensOk returns a tuple with the AudioTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioTokens

`func (o *PromptTokensDetails) SetAudioTokens(v int32)`

SetAudioTokens sets AudioTokens field to given value.

### HasAudioTokens

`func (o *PromptTokensDetails) HasAudioTokens() bool`

HasAudioTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


