# VibeWorkspaceStatsOUT

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTime** | **int32** | Start of the queried window, as a Unix timestamp in seconds. | 
**EndTime** | **int32** | End of the queried window, as a Unix timestamp in seconds. | 
**Sessions** | [**[]VibeSessionStat**](VibeSessionStat.md) |  | 
**UserPrompts** | [**[]VibeUserPromptsStat**](VibeUserPromptsStat.md) |  | 
**ActiveUsers** | [**[]VibeActiveUsersStat**](VibeActiveUsersStat.md) |  | 
**ConsumedTokens** | [**[]VibeConsumedTokensStat**](VibeConsumedTokensStat.md) |  | 
**ToolCalls** | [**[]VibeToolCallsStat**](VibeToolCallsStat.md) |  | 
**ToolCallsByName** | [**[]VibeToolCallsByNameStat**](VibeToolCallsByNameStat.md) |  | 
**SessionDurations** | [**[]VibeSessionDurationStat**](VibeSessionDurationStat.md) |  | 

## Methods

### NewVibeWorkspaceStatsOUT

`func NewVibeWorkspaceStatsOUT(startTime int32, endTime int32, sessions []VibeSessionStat, userPrompts []VibeUserPromptsStat, activeUsers []VibeActiveUsersStat, consumedTokens []VibeConsumedTokensStat, toolCalls []VibeToolCallsStat, toolCallsByName []VibeToolCallsByNameStat, sessionDurations []VibeSessionDurationStat, ) *VibeWorkspaceStatsOUT`

NewVibeWorkspaceStatsOUT instantiates a new VibeWorkspaceStatsOUT object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVibeWorkspaceStatsOUTWithDefaults

`func NewVibeWorkspaceStatsOUTWithDefaults() *VibeWorkspaceStatsOUT`

NewVibeWorkspaceStatsOUTWithDefaults instantiates a new VibeWorkspaceStatsOUT object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTime

`func (o *VibeWorkspaceStatsOUT) GetStartTime() int32`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *VibeWorkspaceStatsOUT) GetStartTimeOk() (*int32, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *VibeWorkspaceStatsOUT) SetStartTime(v int32)`

SetStartTime sets StartTime field to given value.


### GetEndTime

`func (o *VibeWorkspaceStatsOUT) GetEndTime() int32`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *VibeWorkspaceStatsOUT) GetEndTimeOk() (*int32, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *VibeWorkspaceStatsOUT) SetEndTime(v int32)`

SetEndTime sets EndTime field to given value.


### GetSessions

`func (o *VibeWorkspaceStatsOUT) GetSessions() []VibeSessionStat`

GetSessions returns the Sessions field if non-nil, zero value otherwise.

### GetSessionsOk

`func (o *VibeWorkspaceStatsOUT) GetSessionsOk() (*[]VibeSessionStat, bool)`

GetSessionsOk returns a tuple with the Sessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessions

`func (o *VibeWorkspaceStatsOUT) SetSessions(v []VibeSessionStat)`

SetSessions sets Sessions field to given value.


### GetUserPrompts

`func (o *VibeWorkspaceStatsOUT) GetUserPrompts() []VibeUserPromptsStat`

GetUserPrompts returns the UserPrompts field if non-nil, zero value otherwise.

### GetUserPromptsOk

`func (o *VibeWorkspaceStatsOUT) GetUserPromptsOk() (*[]VibeUserPromptsStat, bool)`

GetUserPromptsOk returns a tuple with the UserPrompts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPrompts

`func (o *VibeWorkspaceStatsOUT) SetUserPrompts(v []VibeUserPromptsStat)`

SetUserPrompts sets UserPrompts field to given value.


### GetActiveUsers

`func (o *VibeWorkspaceStatsOUT) GetActiveUsers() []VibeActiveUsersStat`

GetActiveUsers returns the ActiveUsers field if non-nil, zero value otherwise.

### GetActiveUsersOk

`func (o *VibeWorkspaceStatsOUT) GetActiveUsersOk() (*[]VibeActiveUsersStat, bool)`

GetActiveUsersOk returns a tuple with the ActiveUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveUsers

`func (o *VibeWorkspaceStatsOUT) SetActiveUsers(v []VibeActiveUsersStat)`

SetActiveUsers sets ActiveUsers field to given value.


### GetConsumedTokens

`func (o *VibeWorkspaceStatsOUT) GetConsumedTokens() []VibeConsumedTokensStat`

GetConsumedTokens returns the ConsumedTokens field if non-nil, zero value otherwise.

### GetConsumedTokensOk

`func (o *VibeWorkspaceStatsOUT) GetConsumedTokensOk() (*[]VibeConsumedTokensStat, bool)`

GetConsumedTokensOk returns a tuple with the ConsumedTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumedTokens

`func (o *VibeWorkspaceStatsOUT) SetConsumedTokens(v []VibeConsumedTokensStat)`

SetConsumedTokens sets ConsumedTokens field to given value.


### GetToolCalls

`func (o *VibeWorkspaceStatsOUT) GetToolCalls() []VibeToolCallsStat`

GetToolCalls returns the ToolCalls field if non-nil, zero value otherwise.

### GetToolCallsOk

`func (o *VibeWorkspaceStatsOUT) GetToolCallsOk() (*[]VibeToolCallsStat, bool)`

GetToolCallsOk returns a tuple with the ToolCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolCalls

`func (o *VibeWorkspaceStatsOUT) SetToolCalls(v []VibeToolCallsStat)`

SetToolCalls sets ToolCalls field to given value.


### GetToolCallsByName

`func (o *VibeWorkspaceStatsOUT) GetToolCallsByName() []VibeToolCallsByNameStat`

GetToolCallsByName returns the ToolCallsByName field if non-nil, zero value otherwise.

### GetToolCallsByNameOk

`func (o *VibeWorkspaceStatsOUT) GetToolCallsByNameOk() (*[]VibeToolCallsByNameStat, bool)`

GetToolCallsByNameOk returns a tuple with the ToolCallsByName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolCallsByName

`func (o *VibeWorkspaceStatsOUT) SetToolCallsByName(v []VibeToolCallsByNameStat)`

SetToolCallsByName sets ToolCallsByName field to given value.


### GetSessionDurations

`func (o *VibeWorkspaceStatsOUT) GetSessionDurations() []VibeSessionDurationStat`

GetSessionDurations returns the SessionDurations field if non-nil, zero value otherwise.

### GetSessionDurationsOk

`func (o *VibeWorkspaceStatsOUT) GetSessionDurationsOk() (*[]VibeSessionDurationStat, bool)`

GetSessionDurationsOk returns a tuple with the SessionDurations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionDurations

`func (o *VibeWorkspaceStatsOUT) SetSessionDurations(v []VibeSessionDurationStat)`

SetSessionDurations sets SessionDurations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


