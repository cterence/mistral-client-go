# LimitsContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Completion** | [**UsageLimits**](UsageLimits.md) | Completion usage and rate limits. | 
**LastPaymentFailure** | **bool** | Whether the latest payment attempt failed. | 
**LastPaymentFailureProtection** | **NullableBool** | Whether payment failure protection is enabled. | 
**Currency** | **string** | Currency used for usage and limit amounts. | 

## Methods

### NewLimitsContext

`func NewLimitsContext(completion UsageLimits, lastPaymentFailure bool, lastPaymentFailureProtection NullableBool, currency string, ) *LimitsContext`

NewLimitsContext instantiates a new LimitsContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLimitsContextWithDefaults

`func NewLimitsContextWithDefaults() *LimitsContext`

NewLimitsContextWithDefaults instantiates a new LimitsContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompletion

`func (o *LimitsContext) GetCompletion() UsageLimits`

GetCompletion returns the Completion field if non-nil, zero value otherwise.

### GetCompletionOk

`func (o *LimitsContext) GetCompletionOk() (*UsageLimits, bool)`

GetCompletionOk returns a tuple with the Completion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletion

`func (o *LimitsContext) SetCompletion(v UsageLimits)`

SetCompletion sets Completion field to given value.


### GetLastPaymentFailure

`func (o *LimitsContext) GetLastPaymentFailure() bool`

GetLastPaymentFailure returns the LastPaymentFailure field if non-nil, zero value otherwise.

### GetLastPaymentFailureOk

`func (o *LimitsContext) GetLastPaymentFailureOk() (*bool, bool)`

GetLastPaymentFailureOk returns a tuple with the LastPaymentFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPaymentFailure

`func (o *LimitsContext) SetLastPaymentFailure(v bool)`

SetLastPaymentFailure sets LastPaymentFailure field to given value.


### GetLastPaymentFailureProtection

`func (o *LimitsContext) GetLastPaymentFailureProtection() bool`

GetLastPaymentFailureProtection returns the LastPaymentFailureProtection field if non-nil, zero value otherwise.

### GetLastPaymentFailureProtectionOk

`func (o *LimitsContext) GetLastPaymentFailureProtectionOk() (*bool, bool)`

GetLastPaymentFailureProtectionOk returns a tuple with the LastPaymentFailureProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPaymentFailureProtection

`func (o *LimitsContext) SetLastPaymentFailureProtection(v bool)`

SetLastPaymentFailureProtection sets LastPaymentFailureProtection field to given value.


### SetLastPaymentFailureProtectionNil

`func (o *LimitsContext) SetLastPaymentFailureProtectionNil(b bool)`

 SetLastPaymentFailureProtectionNil sets the value for LastPaymentFailureProtection to be an explicit nil

### UnsetLastPaymentFailureProtection
`func (o *LimitsContext) UnsetLastPaymentFailureProtection()`

UnsetLastPaymentFailureProtection ensures that no value is present for LastPaymentFailureProtection, not even an explicit nil
### GetCurrency

`func (o *LimitsContext) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *LimitsContext) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *LimitsContext) SetCurrency(v string)`

SetCurrency sets Currency field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


