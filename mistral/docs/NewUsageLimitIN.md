# NewUsageLimitIN

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int32** | New monthly usage limit amount. | 
**NoMonthlyLimit** | Pointer to **bool** | Whether to remove the monthly usage limit. | [optional] [default to false]

## Methods

### NewNewUsageLimitIN

`func NewNewUsageLimitIN(amount int32, ) *NewUsageLimitIN`

NewNewUsageLimitIN instantiates a new NewUsageLimitIN object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewUsageLimitINWithDefaults

`func NewNewUsageLimitINWithDefaults() *NewUsageLimitIN`

NewNewUsageLimitINWithDefaults instantiates a new NewUsageLimitIN object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *NewUsageLimitIN) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *NewUsageLimitIN) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *NewUsageLimitIN) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetNoMonthlyLimit

`func (o *NewUsageLimitIN) GetNoMonthlyLimit() bool`

GetNoMonthlyLimit returns the NoMonthlyLimit field if non-nil, zero value otherwise.

### GetNoMonthlyLimitOk

`func (o *NewUsageLimitIN) GetNoMonthlyLimitOk() (*bool, bool)`

GetNoMonthlyLimitOk returns a tuple with the NoMonthlyLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoMonthlyLimit

`func (o *NewUsageLimitIN) SetNoMonthlyLimit(v bool)`

SetNoMonthlyLimit sets NoMonthlyLimit field to given value.

### HasNoMonthlyLimit

`func (o *NewUsageLimitIN) HasNoMonthlyLimit() bool`

HasNoMonthlyLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


