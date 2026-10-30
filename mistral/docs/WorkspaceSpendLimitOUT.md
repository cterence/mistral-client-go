# WorkspaceSpendLimitOUT

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **NullableFloat32** | Monthly spending limit for the Workspace. | 
**Currency** | **NullableString** | Currency used for the Workspace spending limit. | 
**NoMonthlyLimit** | **NullableBool** | Whether the Workspace has no monthly spending limit. | 

## Methods

### NewWorkspaceSpendLimitOUT

`func NewWorkspaceSpendLimitOUT(value NullableFloat32, currency NullableString, noMonthlyLimit NullableBool, ) *WorkspaceSpendLimitOUT`

NewWorkspaceSpendLimitOUT instantiates a new WorkspaceSpendLimitOUT object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceSpendLimitOUTWithDefaults

`func NewWorkspaceSpendLimitOUTWithDefaults() *WorkspaceSpendLimitOUT`

NewWorkspaceSpendLimitOUTWithDefaults instantiates a new WorkspaceSpendLimitOUT object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *WorkspaceSpendLimitOUT) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *WorkspaceSpendLimitOUT) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *WorkspaceSpendLimitOUT) SetValue(v float32)`

SetValue sets Value field to given value.


### SetValueNil

`func (o *WorkspaceSpendLimitOUT) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *WorkspaceSpendLimitOUT) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetCurrency

`func (o *WorkspaceSpendLimitOUT) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *WorkspaceSpendLimitOUT) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *WorkspaceSpendLimitOUT) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### SetCurrencyNil

`func (o *WorkspaceSpendLimitOUT) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *WorkspaceSpendLimitOUT) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetNoMonthlyLimit

`func (o *WorkspaceSpendLimitOUT) GetNoMonthlyLimit() bool`

GetNoMonthlyLimit returns the NoMonthlyLimit field if non-nil, zero value otherwise.

### GetNoMonthlyLimitOk

`func (o *WorkspaceSpendLimitOUT) GetNoMonthlyLimitOk() (*bool, bool)`

GetNoMonthlyLimitOk returns a tuple with the NoMonthlyLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoMonthlyLimit

`func (o *WorkspaceSpendLimitOUT) SetNoMonthlyLimit(v bool)`

SetNoMonthlyLimit sets NoMonthlyLimit field to given value.


### SetNoMonthlyLimitNil

`func (o *WorkspaceSpendLimitOUT) SetNoMonthlyLimitNil(b bool)`

 SetNoMonthlyLimitNil sets the value for NoMonthlyLimit to be an explicit nil

### UnsetNoMonthlyLimit
`func (o *WorkspaceSpendLimitOUT) UnsetNoMonthlyLimit()`

UnsetNoMonthlyLimit ensures that no value is present for NoMonthlyLimit, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


