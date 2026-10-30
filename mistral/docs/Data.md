# Data

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "speech.audio.delta"]
**AudioData** | **string** |  | 
**Usage** | [**UsageInfo1**](UsageInfo1.md) |  | 

## Methods

### NewData

`func NewData(audioData string, usage UsageInfo1, ) *Data`

NewData instantiates a new Data object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataWithDefaults

`func NewDataWithDefaults() *Data`

NewDataWithDefaults instantiates a new Data object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Data) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Data) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Data) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Data) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAudioData

`func (o *Data) GetAudioData() string`

GetAudioData returns the AudioData field if non-nil, zero value otherwise.

### GetAudioDataOk

`func (o *Data) GetAudioDataOk() (*string, bool)`

GetAudioDataOk returns a tuple with the AudioData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioData

`func (o *Data) SetAudioData(v string)`

SetAudioData sets AudioData field to given value.


### GetUsage

`func (o *Data) GetUsage() UsageInfo1`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *Data) GetUsageOk() (*UsageInfo1, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *Data) SetUsage(v UsageInfo1)`

SetUsage sets Usage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


