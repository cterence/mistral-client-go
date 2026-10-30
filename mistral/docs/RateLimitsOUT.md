# RateLimitsOUT

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestsPerSecond** | **int32** | Maximum API requests allowed per second. | 
**TokensLimitsByModel** | [**map[string]TokenLimitsByModel**](TokenLimitsByModel.md) | Token limits for each model. | 

## Methods

### NewRateLimitsOUT

`func NewRateLimitsOUT(requestsPerSecond int32, tokensLimitsByModel map[string]TokenLimitsByModel, ) *RateLimitsOUT`

NewRateLimitsOUT instantiates a new RateLimitsOUT object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRateLimitsOUTWithDefaults

`func NewRateLimitsOUTWithDefaults() *RateLimitsOUT`

NewRateLimitsOUTWithDefaults instantiates a new RateLimitsOUT object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestsPerSecond

`func (o *RateLimitsOUT) GetRequestsPerSecond() int32`

GetRequestsPerSecond returns the RequestsPerSecond field if non-nil, zero value otherwise.

### GetRequestsPerSecondOk

`func (o *RateLimitsOUT) GetRequestsPerSecondOk() (*int32, bool)`

GetRequestsPerSecondOk returns a tuple with the RequestsPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestsPerSecond

`func (o *RateLimitsOUT) SetRequestsPerSecond(v int32)`

SetRequestsPerSecond sets RequestsPerSecond field to given value.


### GetTokensLimitsByModel

`func (o *RateLimitsOUT) GetTokensLimitsByModel() map[string]TokenLimitsByModel`

GetTokensLimitsByModel returns the TokensLimitsByModel field if non-nil, zero value otherwise.

### GetTokensLimitsByModelOk

`func (o *RateLimitsOUT) GetTokensLimitsByModelOk() (*map[string]TokenLimitsByModel, bool)`

GetTokensLimitsByModelOk returns a tuple with the TokensLimitsByModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokensLimitsByModel

`func (o *RateLimitsOUT) SetTokensLimitsByModel(v map[string]TokenLimitsByModel)`

SetTokensLimitsByModel sets TokensLimitsByModel field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


