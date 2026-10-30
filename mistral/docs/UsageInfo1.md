# UsageInfo1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PromptAudioSeconds** | Pointer to **NullableInt32** |  | [optional] 
**PromptTokens** | Pointer to **int32** |  | [optional] [default to 0]
**TotalTokens** | Pointer to **int32** |  | [optional] [default to 0]
**CompletionTokens** | Pointer to **NullableInt32** |  | [optional] 
**RequestCount** | Pointer to **NullableInt32** |  | [optional] 
**PromptTokensDetails** | Pointer to [**NullablePromptTokensDetails**](PromptTokensDetails.md) |  | [optional] 
**CompletionTokensDetails** | Pointer to [**NullableCompletionTokensDetails**](CompletionTokensDetails.md) |  | [optional] 
**ServiceTier** | Pointer to **NullableString** | The service tier at which the request was processed: standard or priority. | [optional] 
**PromptTokenDetails** | Pointer to [**NullablePromptTokensDetails**](PromptTokensDetails.md) |  | [optional] 
**NumCachedTokens** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewUsageInfo1

`func NewUsageInfo1() *UsageInfo1`

NewUsageInfo1 instantiates a new UsageInfo1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageInfo1WithDefaults

`func NewUsageInfo1WithDefaults() *UsageInfo1`

NewUsageInfo1WithDefaults instantiates a new UsageInfo1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPromptAudioSeconds

`func (o *UsageInfo1) GetPromptAudioSeconds() int32`

GetPromptAudioSeconds returns the PromptAudioSeconds field if non-nil, zero value otherwise.

### GetPromptAudioSecondsOk

`func (o *UsageInfo1) GetPromptAudioSecondsOk() (*int32, bool)`

GetPromptAudioSecondsOk returns a tuple with the PromptAudioSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptAudioSeconds

`func (o *UsageInfo1) SetPromptAudioSeconds(v int32)`

SetPromptAudioSeconds sets PromptAudioSeconds field to given value.

### HasPromptAudioSeconds

`func (o *UsageInfo1) HasPromptAudioSeconds() bool`

HasPromptAudioSeconds returns a boolean if a field has been set.

### SetPromptAudioSecondsNil

`func (o *UsageInfo1) SetPromptAudioSecondsNil(b bool)`

 SetPromptAudioSecondsNil sets the value for PromptAudioSeconds to be an explicit nil

### UnsetPromptAudioSeconds
`func (o *UsageInfo1) UnsetPromptAudioSeconds()`

UnsetPromptAudioSeconds ensures that no value is present for PromptAudioSeconds, not even an explicit nil
### GetPromptTokens

`func (o *UsageInfo1) GetPromptTokens() int32`

GetPromptTokens returns the PromptTokens field if non-nil, zero value otherwise.

### GetPromptTokensOk

`func (o *UsageInfo1) GetPromptTokensOk() (*int32, bool)`

GetPromptTokensOk returns a tuple with the PromptTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptTokens

`func (o *UsageInfo1) SetPromptTokens(v int32)`

SetPromptTokens sets PromptTokens field to given value.

### HasPromptTokens

`func (o *UsageInfo1) HasPromptTokens() bool`

HasPromptTokens returns a boolean if a field has been set.

### GetTotalTokens

`func (o *UsageInfo1) GetTotalTokens() int32`

GetTotalTokens returns the TotalTokens field if non-nil, zero value otherwise.

### GetTotalTokensOk

`func (o *UsageInfo1) GetTotalTokensOk() (*int32, bool)`

GetTotalTokensOk returns a tuple with the TotalTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTokens

`func (o *UsageInfo1) SetTotalTokens(v int32)`

SetTotalTokens sets TotalTokens field to given value.

### HasTotalTokens

`func (o *UsageInfo1) HasTotalTokens() bool`

HasTotalTokens returns a boolean if a field has been set.

### GetCompletionTokens

`func (o *UsageInfo1) GetCompletionTokens() int32`

GetCompletionTokens returns the CompletionTokens field if non-nil, zero value otherwise.

### GetCompletionTokensOk

`func (o *UsageInfo1) GetCompletionTokensOk() (*int32, bool)`

GetCompletionTokensOk returns a tuple with the CompletionTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTokens

`func (o *UsageInfo1) SetCompletionTokens(v int32)`

SetCompletionTokens sets CompletionTokens field to given value.

### HasCompletionTokens

`func (o *UsageInfo1) HasCompletionTokens() bool`

HasCompletionTokens returns a boolean if a field has been set.

### SetCompletionTokensNil

`func (o *UsageInfo1) SetCompletionTokensNil(b bool)`

 SetCompletionTokensNil sets the value for CompletionTokens to be an explicit nil

### UnsetCompletionTokens
`func (o *UsageInfo1) UnsetCompletionTokens()`

UnsetCompletionTokens ensures that no value is present for CompletionTokens, not even an explicit nil
### GetRequestCount

`func (o *UsageInfo1) GetRequestCount() int32`

GetRequestCount returns the RequestCount field if non-nil, zero value otherwise.

### GetRequestCountOk

`func (o *UsageInfo1) GetRequestCountOk() (*int32, bool)`

GetRequestCountOk returns a tuple with the RequestCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestCount

`func (o *UsageInfo1) SetRequestCount(v int32)`

SetRequestCount sets RequestCount field to given value.

### HasRequestCount

`func (o *UsageInfo1) HasRequestCount() bool`

HasRequestCount returns a boolean if a field has been set.

### SetRequestCountNil

`func (o *UsageInfo1) SetRequestCountNil(b bool)`

 SetRequestCountNil sets the value for RequestCount to be an explicit nil

### UnsetRequestCount
`func (o *UsageInfo1) UnsetRequestCount()`

UnsetRequestCount ensures that no value is present for RequestCount, not even an explicit nil
### GetPromptTokensDetails

`func (o *UsageInfo1) GetPromptTokensDetails() PromptTokensDetails`

GetPromptTokensDetails returns the PromptTokensDetails field if non-nil, zero value otherwise.

### GetPromptTokensDetailsOk

`func (o *UsageInfo1) GetPromptTokensDetailsOk() (*PromptTokensDetails, bool)`

GetPromptTokensDetailsOk returns a tuple with the PromptTokensDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptTokensDetails

`func (o *UsageInfo1) SetPromptTokensDetails(v PromptTokensDetails)`

SetPromptTokensDetails sets PromptTokensDetails field to given value.

### HasPromptTokensDetails

`func (o *UsageInfo1) HasPromptTokensDetails() bool`

HasPromptTokensDetails returns a boolean if a field has been set.

### SetPromptTokensDetailsNil

`func (o *UsageInfo1) SetPromptTokensDetailsNil(b bool)`

 SetPromptTokensDetailsNil sets the value for PromptTokensDetails to be an explicit nil

### UnsetPromptTokensDetails
`func (o *UsageInfo1) UnsetPromptTokensDetails()`

UnsetPromptTokensDetails ensures that no value is present for PromptTokensDetails, not even an explicit nil
### GetCompletionTokensDetails

`func (o *UsageInfo1) GetCompletionTokensDetails() CompletionTokensDetails`

GetCompletionTokensDetails returns the CompletionTokensDetails field if non-nil, zero value otherwise.

### GetCompletionTokensDetailsOk

`func (o *UsageInfo1) GetCompletionTokensDetailsOk() (*CompletionTokensDetails, bool)`

GetCompletionTokensDetailsOk returns a tuple with the CompletionTokensDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTokensDetails

`func (o *UsageInfo1) SetCompletionTokensDetails(v CompletionTokensDetails)`

SetCompletionTokensDetails sets CompletionTokensDetails field to given value.

### HasCompletionTokensDetails

`func (o *UsageInfo1) HasCompletionTokensDetails() bool`

HasCompletionTokensDetails returns a boolean if a field has been set.

### SetCompletionTokensDetailsNil

`func (o *UsageInfo1) SetCompletionTokensDetailsNil(b bool)`

 SetCompletionTokensDetailsNil sets the value for CompletionTokensDetails to be an explicit nil

### UnsetCompletionTokensDetails
`func (o *UsageInfo1) UnsetCompletionTokensDetails()`

UnsetCompletionTokensDetails ensures that no value is present for CompletionTokensDetails, not even an explicit nil
### GetServiceTier

`func (o *UsageInfo1) GetServiceTier() string`

GetServiceTier returns the ServiceTier field if non-nil, zero value otherwise.

### GetServiceTierOk

`func (o *UsageInfo1) GetServiceTierOk() (*string, bool)`

GetServiceTierOk returns a tuple with the ServiceTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTier

`func (o *UsageInfo1) SetServiceTier(v string)`

SetServiceTier sets ServiceTier field to given value.

### HasServiceTier

`func (o *UsageInfo1) HasServiceTier() bool`

HasServiceTier returns a boolean if a field has been set.

### SetServiceTierNil

`func (o *UsageInfo1) SetServiceTierNil(b bool)`

 SetServiceTierNil sets the value for ServiceTier to be an explicit nil

### UnsetServiceTier
`func (o *UsageInfo1) UnsetServiceTier()`

UnsetServiceTier ensures that no value is present for ServiceTier, not even an explicit nil
### GetPromptTokenDetails

`func (o *UsageInfo1) GetPromptTokenDetails() PromptTokensDetails`

GetPromptTokenDetails returns the PromptTokenDetails field if non-nil, zero value otherwise.

### GetPromptTokenDetailsOk

`func (o *UsageInfo1) GetPromptTokenDetailsOk() (*PromptTokensDetails, bool)`

GetPromptTokenDetailsOk returns a tuple with the PromptTokenDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptTokenDetails

`func (o *UsageInfo1) SetPromptTokenDetails(v PromptTokensDetails)`

SetPromptTokenDetails sets PromptTokenDetails field to given value.

### HasPromptTokenDetails

`func (o *UsageInfo1) HasPromptTokenDetails() bool`

HasPromptTokenDetails returns a boolean if a field has been set.

### SetPromptTokenDetailsNil

`func (o *UsageInfo1) SetPromptTokenDetailsNil(b bool)`

 SetPromptTokenDetailsNil sets the value for PromptTokenDetails to be an explicit nil

### UnsetPromptTokenDetails
`func (o *UsageInfo1) UnsetPromptTokenDetails()`

UnsetPromptTokenDetails ensures that no value is present for PromptTokenDetails, not even an explicit nil
### GetNumCachedTokens

`func (o *UsageInfo1) GetNumCachedTokens() int32`

GetNumCachedTokens returns the NumCachedTokens field if non-nil, zero value otherwise.

### GetNumCachedTokensOk

`func (o *UsageInfo1) GetNumCachedTokensOk() (*int32, bool)`

GetNumCachedTokensOk returns a tuple with the NumCachedTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumCachedTokens

`func (o *UsageInfo1) SetNumCachedTokens(v int32)`

SetNumCachedTokens sets NumCachedTokens field to given value.

### HasNumCachedTokens

`func (o *UsageInfo1) HasNumCachedTokens() bool`

HasNumCachedTokens returns a boolean if a field has been set.

### SetNumCachedTokensNil

`func (o *UsageInfo1) SetNumCachedTokensNil(b bool)`

 SetNumCachedTokensNil sets the value for NumCachedTokens to be an explicit nil

### UnsetNumCachedTokens
`func (o *UsageInfo1) UnsetNumCachedTokens()`

UnsetNumCachedTokens ensures that no value is present for NumCachedTokens, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


