# MessageTokens

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | [**Roles**](Roles.md) |  | 
**TotalTokens** | Pointer to **NullableInt32** |  | [optional] 
**SettingsTokens** | Pointer to **NullableInt32** |  | [optional] 
**Truncated** | Pointer to **bool** |  | [optional] [default to false]
**UsageCount** | Pointer to **int32** |  | [optional] [default to 1]

## Methods

### NewMessageTokens

`func NewMessageTokens(role Roles, ) *MessageTokens`

NewMessageTokens instantiates a new MessageTokens object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageTokensWithDefaults

`func NewMessageTokensWithDefaults() *MessageTokens`

NewMessageTokensWithDefaults instantiates a new MessageTokens object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *MessageTokens) GetRole() Roles`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *MessageTokens) GetRoleOk() (*Roles, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *MessageTokens) SetRole(v Roles)`

SetRole sets Role field to given value.


### GetTotalTokens

`func (o *MessageTokens) GetTotalTokens() int32`

GetTotalTokens returns the TotalTokens field if non-nil, zero value otherwise.

### GetTotalTokensOk

`func (o *MessageTokens) GetTotalTokensOk() (*int32, bool)`

GetTotalTokensOk returns a tuple with the TotalTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTokens

`func (o *MessageTokens) SetTotalTokens(v int32)`

SetTotalTokens sets TotalTokens field to given value.

### HasTotalTokens

`func (o *MessageTokens) HasTotalTokens() bool`

HasTotalTokens returns a boolean if a field has been set.

### SetTotalTokensNil

`func (o *MessageTokens) SetTotalTokensNil(b bool)`

 SetTotalTokensNil sets the value for TotalTokens to be an explicit nil

### UnsetTotalTokens
`func (o *MessageTokens) UnsetTotalTokens()`

UnsetTotalTokens ensures that no value is present for TotalTokens, not even an explicit nil
### GetSettingsTokens

`func (o *MessageTokens) GetSettingsTokens() int32`

GetSettingsTokens returns the SettingsTokens field if non-nil, zero value otherwise.

### GetSettingsTokensOk

`func (o *MessageTokens) GetSettingsTokensOk() (*int32, bool)`

GetSettingsTokensOk returns a tuple with the SettingsTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingsTokens

`func (o *MessageTokens) SetSettingsTokens(v int32)`

SetSettingsTokens sets SettingsTokens field to given value.

### HasSettingsTokens

`func (o *MessageTokens) HasSettingsTokens() bool`

HasSettingsTokens returns a boolean if a field has been set.

### SetSettingsTokensNil

`func (o *MessageTokens) SetSettingsTokensNil(b bool)`

 SetSettingsTokensNil sets the value for SettingsTokens to be an explicit nil

### UnsetSettingsTokens
`func (o *MessageTokens) UnsetSettingsTokens()`

UnsetSettingsTokens ensures that no value is present for SettingsTokens, not even an explicit nil
### GetTruncated

`func (o *MessageTokens) GetTruncated() bool`

GetTruncated returns the Truncated field if non-nil, zero value otherwise.

### GetTruncatedOk

`func (o *MessageTokens) GetTruncatedOk() (*bool, bool)`

GetTruncatedOk returns a tuple with the Truncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncated

`func (o *MessageTokens) SetTruncated(v bool)`

SetTruncated sets Truncated field to given value.

### HasTruncated

`func (o *MessageTokens) HasTruncated() bool`

HasTruncated returns a boolean if a field has been set.

### GetUsageCount

`func (o *MessageTokens) GetUsageCount() int32`

GetUsageCount returns the UsageCount field if non-nil, zero value otherwise.

### GetUsageCountOk

`func (o *MessageTokens) GetUsageCountOk() (*int32, bool)`

GetUsageCountOk returns a tuple with the UsageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageCount

`func (o *MessageTokens) SetUsageCount(v int32)`

SetUsageCount sets UsageCount field to given value.

### HasUsageCount

`func (o *MessageTokens) HasUsageCount() bool`

HasUsageCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


