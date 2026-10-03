# UsageInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PromptTokens** | **int32** |  | [default to 0]
**CompletionTokens** | **int32** |  | [default to 0]
**TotalTokens** | **int32** |  | [default to 0]
**PromptAudioSeconds** | Pointer to **NullableInt32** |  | [optional] 
**NumCachedTokens** | Pointer to **NullableInt32** |  | [optional] 
**PromptTokensDetails** | Pointer to [**NullablePromptTokensDetails**](PromptTokensDetails.md) |  | [optional] 
**PromptTokenDetails** | Pointer to [**NullablePromptTokensDetails**](PromptTokensDetails.md) |  | [optional] 

## Methods

### NewUsageInfo

`func NewUsageInfo(promptTokens int32, completionTokens int32, totalTokens int32, ) *UsageInfo`

NewUsageInfo instantiates a new UsageInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageInfoWithDefaults

`func NewUsageInfoWithDefaults() *UsageInfo`

NewUsageInfoWithDefaults instantiates a new UsageInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPromptTokens

`func (o *UsageInfo) GetPromptTokens() int32`

GetPromptTokens returns the PromptTokens field if non-nil, zero value otherwise.

### GetPromptTokensOk

`func (o *UsageInfo) GetPromptTokensOk() (*int32, bool)`

GetPromptTokensOk returns a tuple with the PromptTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptTokens

`func (o *UsageInfo) SetPromptTokens(v int32)`

SetPromptTokens sets PromptTokens field to given value.


### GetCompletionTokens

`func (o *UsageInfo) GetCompletionTokens() int32`

GetCompletionTokens returns the CompletionTokens field if non-nil, zero value otherwise.

### GetCompletionTokensOk

`func (o *UsageInfo) GetCompletionTokensOk() (*int32, bool)`

GetCompletionTokensOk returns a tuple with the CompletionTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTokens

`func (o *UsageInfo) SetCompletionTokens(v int32)`

SetCompletionTokens sets CompletionTokens field to given value.


### GetTotalTokens

`func (o *UsageInfo) GetTotalTokens() int32`

GetTotalTokens returns the TotalTokens field if non-nil, zero value otherwise.

### GetTotalTokensOk

`func (o *UsageInfo) GetTotalTokensOk() (*int32, bool)`

GetTotalTokensOk returns a tuple with the TotalTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTokens

`func (o *UsageInfo) SetTotalTokens(v int32)`

SetTotalTokens sets TotalTokens field to given value.


### GetPromptAudioSeconds

`func (o *UsageInfo) GetPromptAudioSeconds() int32`

GetPromptAudioSeconds returns the PromptAudioSeconds field if non-nil, zero value otherwise.

### GetPromptAudioSecondsOk

`func (o *UsageInfo) GetPromptAudioSecondsOk() (*int32, bool)`

GetPromptAudioSecondsOk returns a tuple with the PromptAudioSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptAudioSeconds

`func (o *UsageInfo) SetPromptAudioSeconds(v int32)`

SetPromptAudioSeconds sets PromptAudioSeconds field to given value.

### HasPromptAudioSeconds

`func (o *UsageInfo) HasPromptAudioSeconds() bool`

HasPromptAudioSeconds returns a boolean if a field has been set.

### SetPromptAudioSecondsNil

`func (o *UsageInfo) SetPromptAudioSecondsNil(b bool)`

 SetPromptAudioSecondsNil sets the value for PromptAudioSeconds to be an explicit nil

### UnsetPromptAudioSeconds
`func (o *UsageInfo) UnsetPromptAudioSeconds()`

UnsetPromptAudioSeconds ensures that no value is present for PromptAudioSeconds, not even an explicit nil
### GetNumCachedTokens

`func (o *UsageInfo) GetNumCachedTokens() int32`

GetNumCachedTokens returns the NumCachedTokens field if non-nil, zero value otherwise.

### GetNumCachedTokensOk

`func (o *UsageInfo) GetNumCachedTokensOk() (*int32, bool)`

GetNumCachedTokensOk returns a tuple with the NumCachedTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumCachedTokens

`func (o *UsageInfo) SetNumCachedTokens(v int32)`

SetNumCachedTokens sets NumCachedTokens field to given value.

### HasNumCachedTokens

`func (o *UsageInfo) HasNumCachedTokens() bool`

HasNumCachedTokens returns a boolean if a field has been set.

### SetNumCachedTokensNil

`func (o *UsageInfo) SetNumCachedTokensNil(b bool)`

 SetNumCachedTokensNil sets the value for NumCachedTokens to be an explicit nil

### UnsetNumCachedTokens
`func (o *UsageInfo) UnsetNumCachedTokens()`

UnsetNumCachedTokens ensures that no value is present for NumCachedTokens, not even an explicit nil
### GetPromptTokensDetails

`func (o *UsageInfo) GetPromptTokensDetails() PromptTokensDetails`

GetPromptTokensDetails returns the PromptTokensDetails field if non-nil, zero value otherwise.

### GetPromptTokensDetailsOk

`func (o *UsageInfo) GetPromptTokensDetailsOk() (*PromptTokensDetails, bool)`

GetPromptTokensDetailsOk returns a tuple with the PromptTokensDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptTokensDetails

`func (o *UsageInfo) SetPromptTokensDetails(v PromptTokensDetails)`

SetPromptTokensDetails sets PromptTokensDetails field to given value.

### HasPromptTokensDetails

`func (o *UsageInfo) HasPromptTokensDetails() bool`

HasPromptTokensDetails returns a boolean if a field has been set.

### SetPromptTokensDetailsNil

`func (o *UsageInfo) SetPromptTokensDetailsNil(b bool)`

 SetPromptTokensDetailsNil sets the value for PromptTokensDetails to be an explicit nil

### UnsetPromptTokensDetails
`func (o *UsageInfo) UnsetPromptTokensDetails()`

UnsetPromptTokensDetails ensures that no value is present for PromptTokensDetails, not even an explicit nil
### GetPromptTokenDetails

`func (o *UsageInfo) GetPromptTokenDetails() PromptTokensDetails`

GetPromptTokenDetails returns the PromptTokenDetails field if non-nil, zero value otherwise.

### GetPromptTokenDetailsOk

`func (o *UsageInfo) GetPromptTokenDetailsOk() (*PromptTokensDetails, bool)`

GetPromptTokenDetailsOk returns a tuple with the PromptTokenDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptTokenDetails

`func (o *UsageInfo) SetPromptTokenDetails(v PromptTokensDetails)`

SetPromptTokenDetails sets PromptTokenDetails field to given value.

### HasPromptTokenDetails

`func (o *UsageInfo) HasPromptTokenDetails() bool`

HasPromptTokenDetails returns a boolean if a field has been set.

### SetPromptTokenDetailsNil

`func (o *UsageInfo) SetPromptTokenDetailsNil(b bool)`

 SetPromptTokenDetailsNil sets the value for PromptTokenDetails to be an explicit nil

### UnsetPromptTokenDetails
`func (o *UsageInfo) UnsetPromptTokenDetails()`

UnsetPromptTokenDetails ensures that no value is present for PromptTokenDetails, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


