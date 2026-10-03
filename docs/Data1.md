# Data1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "speech.audio.delta"]
**AudioData** | **string** |  | 
**Usage** | [**UsageInfo**](UsageInfo.md) |  | 

## Methods

### NewData1

`func NewData1(audioData string, usage UsageInfo, ) *Data1`

NewData1 instantiates a new Data1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewData1WithDefaults

`func NewData1WithDefaults() *Data1`

NewData1WithDefaults instantiates a new Data1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Data1) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Data1) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Data1) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Data1) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAudioData

`func (o *Data1) GetAudioData() string`

GetAudioData returns the AudioData field if non-nil, zero value otherwise.

### GetAudioDataOk

`func (o *Data1) GetAudioDataOk() (*string, bool)`

GetAudioDataOk returns a tuple with the AudioData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioData

`func (o *Data1) SetAudioData(v string)`

SetAudioData sets AudioData field to given value.


### GetUsage

`func (o *Data1) GetUsage() UsageInfo`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *Data1) GetUsageOk() (*UsageInfo, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *Data1) SetUsage(v UsageInfo)`

SetUsage sets Usage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


