# TokenLimitsByModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokensPerMinute** | **int32** | Maximum tokens allowed per minute. | 
**TokensPerMonth** | **int32** | Maximum tokens allowed per month. | 

## Methods

### NewTokenLimitsByModel

`func NewTokenLimitsByModel(tokensPerMinute int32, tokensPerMonth int32, ) *TokenLimitsByModel`

NewTokenLimitsByModel instantiates a new TokenLimitsByModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenLimitsByModelWithDefaults

`func NewTokenLimitsByModelWithDefaults() *TokenLimitsByModel`

NewTokenLimitsByModelWithDefaults instantiates a new TokenLimitsByModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokensPerMinute

`func (o *TokenLimitsByModel) GetTokensPerMinute() int32`

GetTokensPerMinute returns the TokensPerMinute field if non-nil, zero value otherwise.

### GetTokensPerMinuteOk

`func (o *TokenLimitsByModel) GetTokensPerMinuteOk() (*int32, bool)`

GetTokensPerMinuteOk returns a tuple with the TokensPerMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokensPerMinute

`func (o *TokenLimitsByModel) SetTokensPerMinute(v int32)`

SetTokensPerMinute sets TokensPerMinute field to given value.


### GetTokensPerMonth

`func (o *TokenLimitsByModel) GetTokensPerMonth() int32`

GetTokensPerMonth returns the TokensPerMonth field if non-nil, zero value otherwise.

### GetTokensPerMonthOk

`func (o *TokenLimitsByModel) GetTokensPerMonthOk() (*int32, bool)`

GetTokensPerMonthOk returns a tuple with the TokensPerMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokensPerMonth

`func (o *TokenLimitsByModel) SetTokensPerMonth(v int32)`

SetTokensPerMonth sets TokensPerMonth field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


