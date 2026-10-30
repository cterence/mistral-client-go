# UsageLimits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NoMonthlyLimit** | Pointer to **bool** | Whether no monthly usage limit is configured. | [optional] [default to false]
**MonthlyLimitReached** | **bool** | Whether the monthly usage limit has been reached. | 
**Usage** | Pointer to **NullableFloat32** | Current usage counted against the limit. | [optional] 
**VibeUsage** | Pointer to **NullableFloat32** | Current Vibe usage counted against the limit. | [optional] 
**TotalUsage** | Pointer to **NullableFloat32** | Total current usage counted against the limit. | [optional] 
**UsageLimit** | Pointer to **NullableFloat32** | Current usage limit. | [optional] 

## Methods

### NewUsageLimits

`func NewUsageLimits(monthlyLimitReached bool, ) *UsageLimits`

NewUsageLimits instantiates a new UsageLimits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageLimitsWithDefaults

`func NewUsageLimitsWithDefaults() *UsageLimits`

NewUsageLimitsWithDefaults instantiates a new UsageLimits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNoMonthlyLimit

`func (o *UsageLimits) GetNoMonthlyLimit() bool`

GetNoMonthlyLimit returns the NoMonthlyLimit field if non-nil, zero value otherwise.

### GetNoMonthlyLimitOk

`func (o *UsageLimits) GetNoMonthlyLimitOk() (*bool, bool)`

GetNoMonthlyLimitOk returns a tuple with the NoMonthlyLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoMonthlyLimit

`func (o *UsageLimits) SetNoMonthlyLimit(v bool)`

SetNoMonthlyLimit sets NoMonthlyLimit field to given value.

### HasNoMonthlyLimit

`func (o *UsageLimits) HasNoMonthlyLimit() bool`

HasNoMonthlyLimit returns a boolean if a field has been set.

### GetMonthlyLimitReached

`func (o *UsageLimits) GetMonthlyLimitReached() bool`

GetMonthlyLimitReached returns the MonthlyLimitReached field if non-nil, zero value otherwise.

### GetMonthlyLimitReachedOk

`func (o *UsageLimits) GetMonthlyLimitReachedOk() (*bool, bool)`

GetMonthlyLimitReachedOk returns a tuple with the MonthlyLimitReached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyLimitReached

`func (o *UsageLimits) SetMonthlyLimitReached(v bool)`

SetMonthlyLimitReached sets MonthlyLimitReached field to given value.


### GetUsage

`func (o *UsageLimits) GetUsage() float32`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *UsageLimits) GetUsageOk() (*float32, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *UsageLimits) SetUsage(v float32)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *UsageLimits) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### SetUsageNil

`func (o *UsageLimits) SetUsageNil(b bool)`

 SetUsageNil sets the value for Usage to be an explicit nil

### UnsetUsage
`func (o *UsageLimits) UnsetUsage()`

UnsetUsage ensures that no value is present for Usage, not even an explicit nil
### GetVibeUsage

`func (o *UsageLimits) GetVibeUsage() float32`

GetVibeUsage returns the VibeUsage field if non-nil, zero value otherwise.

### GetVibeUsageOk

`func (o *UsageLimits) GetVibeUsageOk() (*float32, bool)`

GetVibeUsageOk returns a tuple with the VibeUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVibeUsage

`func (o *UsageLimits) SetVibeUsage(v float32)`

SetVibeUsage sets VibeUsage field to given value.

### HasVibeUsage

`func (o *UsageLimits) HasVibeUsage() bool`

HasVibeUsage returns a boolean if a field has been set.

### SetVibeUsageNil

`func (o *UsageLimits) SetVibeUsageNil(b bool)`

 SetVibeUsageNil sets the value for VibeUsage to be an explicit nil

### UnsetVibeUsage
`func (o *UsageLimits) UnsetVibeUsage()`

UnsetVibeUsage ensures that no value is present for VibeUsage, not even an explicit nil
### GetTotalUsage

`func (o *UsageLimits) GetTotalUsage() float32`

GetTotalUsage returns the TotalUsage field if non-nil, zero value otherwise.

### GetTotalUsageOk

`func (o *UsageLimits) GetTotalUsageOk() (*float32, bool)`

GetTotalUsageOk returns a tuple with the TotalUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUsage

`func (o *UsageLimits) SetTotalUsage(v float32)`

SetTotalUsage sets TotalUsage field to given value.

### HasTotalUsage

`func (o *UsageLimits) HasTotalUsage() bool`

HasTotalUsage returns a boolean if a field has been set.

### SetTotalUsageNil

`func (o *UsageLimits) SetTotalUsageNil(b bool)`

 SetTotalUsageNil sets the value for TotalUsage to be an explicit nil

### UnsetTotalUsage
`func (o *UsageLimits) UnsetTotalUsage()`

UnsetTotalUsage ensures that no value is present for TotalUsage, not even an explicit nil
### GetUsageLimit

`func (o *UsageLimits) GetUsageLimit() float32`

GetUsageLimit returns the UsageLimit field if non-nil, zero value otherwise.

### GetUsageLimitOk

`func (o *UsageLimits) GetUsageLimitOk() (*float32, bool)`

GetUsageLimitOk returns a tuple with the UsageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageLimit

`func (o *UsageLimits) SetUsageLimit(v float32)`

SetUsageLimit sets UsageLimit field to given value.

### HasUsageLimit

`func (o *UsageLimits) HasUsageLimit() bool`

HasUsageLimit returns a boolean if a field has been set.

### SetUsageLimitNil

`func (o *UsageLimits) SetUsageLimitNil(b bool)`

 SetUsageLimitNil sets the value for UsageLimit to be an explicit nil

### UnsetUsageLimit
`func (o *UsageLimits) UnsetUsageLimit()`

UnsetUsageLimit ensures that no value is present for UsageLimit, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


