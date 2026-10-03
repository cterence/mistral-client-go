# ConversationUsageInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PromptTokens** | Pointer to **int32** |  | [optional] [default to 0]
**CompletionTokens** | Pointer to **int32** |  | [optional] [default to 0]
**TotalTokens** | Pointer to **int32** |  | [optional] [default to 0]
**ConnectorTokens** | Pointer to **NullableInt32** |  | [optional] 
**Connectors** | Pointer to **map[string]int32** |  | [optional] 

## Methods

### NewConversationUsageInfo

`func NewConversationUsageInfo() *ConversationUsageInfo`

NewConversationUsageInfo instantiates a new ConversationUsageInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConversationUsageInfoWithDefaults

`func NewConversationUsageInfoWithDefaults() *ConversationUsageInfo`

NewConversationUsageInfoWithDefaults instantiates a new ConversationUsageInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPromptTokens

`func (o *ConversationUsageInfo) GetPromptTokens() int32`

GetPromptTokens returns the PromptTokens field if non-nil, zero value otherwise.

### GetPromptTokensOk

`func (o *ConversationUsageInfo) GetPromptTokensOk() (*int32, bool)`

GetPromptTokensOk returns a tuple with the PromptTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptTokens

`func (o *ConversationUsageInfo) SetPromptTokens(v int32)`

SetPromptTokens sets PromptTokens field to given value.

### HasPromptTokens

`func (o *ConversationUsageInfo) HasPromptTokens() bool`

HasPromptTokens returns a boolean if a field has been set.

### GetCompletionTokens

`func (o *ConversationUsageInfo) GetCompletionTokens() int32`

GetCompletionTokens returns the CompletionTokens field if non-nil, zero value otherwise.

### GetCompletionTokensOk

`func (o *ConversationUsageInfo) GetCompletionTokensOk() (*int32, bool)`

GetCompletionTokensOk returns a tuple with the CompletionTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTokens

`func (o *ConversationUsageInfo) SetCompletionTokens(v int32)`

SetCompletionTokens sets CompletionTokens field to given value.

### HasCompletionTokens

`func (o *ConversationUsageInfo) HasCompletionTokens() bool`

HasCompletionTokens returns a boolean if a field has been set.

### GetTotalTokens

`func (o *ConversationUsageInfo) GetTotalTokens() int32`

GetTotalTokens returns the TotalTokens field if non-nil, zero value otherwise.

### GetTotalTokensOk

`func (o *ConversationUsageInfo) GetTotalTokensOk() (*int32, bool)`

GetTotalTokensOk returns a tuple with the TotalTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTokens

`func (o *ConversationUsageInfo) SetTotalTokens(v int32)`

SetTotalTokens sets TotalTokens field to given value.

### HasTotalTokens

`func (o *ConversationUsageInfo) HasTotalTokens() bool`

HasTotalTokens returns a boolean if a field has been set.

### GetConnectorTokens

`func (o *ConversationUsageInfo) GetConnectorTokens() int32`

GetConnectorTokens returns the ConnectorTokens field if non-nil, zero value otherwise.

### GetConnectorTokensOk

`func (o *ConversationUsageInfo) GetConnectorTokensOk() (*int32, bool)`

GetConnectorTokensOk returns a tuple with the ConnectorTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorTokens

`func (o *ConversationUsageInfo) SetConnectorTokens(v int32)`

SetConnectorTokens sets ConnectorTokens field to given value.

### HasConnectorTokens

`func (o *ConversationUsageInfo) HasConnectorTokens() bool`

HasConnectorTokens returns a boolean if a field has been set.

### SetConnectorTokensNil

`func (o *ConversationUsageInfo) SetConnectorTokensNil(b bool)`

 SetConnectorTokensNil sets the value for ConnectorTokens to be an explicit nil

### UnsetConnectorTokens
`func (o *ConversationUsageInfo) UnsetConnectorTokens()`

UnsetConnectorTokens ensures that no value is present for ConnectorTokens, not even an explicit nil
### GetConnectors

`func (o *ConversationUsageInfo) GetConnectors() map[string]int32`

GetConnectors returns the Connectors field if non-nil, zero value otherwise.

### GetConnectorsOk

`func (o *ConversationUsageInfo) GetConnectorsOk() (*map[string]int32, bool)`

GetConnectorsOk returns a tuple with the Connectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectors

`func (o *ConversationUsageInfo) SetConnectors(v map[string]int32)`

SetConnectors sets Connectors field to given value.

### HasConnectors

`func (o *ConversationUsageInfo) HasConnectors() bool`

HasConnectors returns a boolean if a field has been set.

### SetConnectorsNil

`func (o *ConversationUsageInfo) SetConnectorsNil(b bool)`

 SetConnectorsNil sets the value for Connectors to be an explicit nil

### UnsetConnectors
`func (o *ConversationUsageInfo) UnsetConnectors()`

UnsetConnectors ensures that no value is present for Connectors, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


