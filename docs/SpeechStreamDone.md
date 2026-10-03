# SpeechStreamDone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "speech.audio.done"]
**Usage** | [**UsageInfo**](UsageInfo.md) |  | 

## Methods

### NewSpeechStreamDone

`func NewSpeechStreamDone(usage UsageInfo, ) *SpeechStreamDone`

NewSpeechStreamDone instantiates a new SpeechStreamDone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpeechStreamDoneWithDefaults

`func NewSpeechStreamDoneWithDefaults() *SpeechStreamDone`

NewSpeechStreamDoneWithDefaults instantiates a new SpeechStreamDone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SpeechStreamDone) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SpeechStreamDone) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SpeechStreamDone) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SpeechStreamDone) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUsage

`func (o *SpeechStreamDone) GetUsage() UsageInfo`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *SpeechStreamDone) GetUsageOk() (*UsageInfo, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *SpeechStreamDone) SetUsage(v UsageInfo)`

SetUsage sets Usage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


