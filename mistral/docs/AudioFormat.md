# AudioFormat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Encoding** | [**AudioEncoding**](AudioEncoding.md) |  | 
**SampleRate** | **int32** |  | 

## Methods

### NewAudioFormat

`func NewAudioFormat(encoding AudioEncoding, sampleRate int32, ) *AudioFormat`

NewAudioFormat instantiates a new AudioFormat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAudioFormatWithDefaults

`func NewAudioFormatWithDefaults() *AudioFormat`

NewAudioFormatWithDefaults instantiates a new AudioFormat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEncoding

`func (o *AudioFormat) GetEncoding() AudioEncoding`

GetEncoding returns the Encoding field if non-nil, zero value otherwise.

### GetEncodingOk

`func (o *AudioFormat) GetEncodingOk() (*AudioEncoding, bool)`

GetEncodingOk returns a tuple with the Encoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoding

`func (o *AudioFormat) SetEncoding(v AudioEncoding)`

SetEncoding sets Encoding field to given value.


### GetSampleRate

`func (o *AudioFormat) GetSampleRate() int32`

GetSampleRate returns the SampleRate field if non-nil, zero value otherwise.

### GetSampleRateOk

`func (o *AudioFormat) GetSampleRateOk() (*int32, bool)`

GetSampleRateOk returns a tuple with the SampleRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleRate

`func (o *AudioFormat) SetSampleRate(v int32)`

SetSampleRate sets SampleRate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


