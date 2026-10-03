# RealtimeTranscriptionSessionUpdatePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AudioFormat** | Pointer to [**NullableAudioFormat**](AudioFormat.md) | Set before sending audio. Audio format updates are rejected after audio starts. | [optional] 
**TargetStreamingDelayMs** | Pointer to **NullableInt32** | Set before sending audio. Streaming delay updates are rejected after audio starts. | [optional] 

## Methods

### NewRealtimeTranscriptionSessionUpdatePayload

`func NewRealtimeTranscriptionSessionUpdatePayload() *RealtimeTranscriptionSessionUpdatePayload`

NewRealtimeTranscriptionSessionUpdatePayload instantiates a new RealtimeTranscriptionSessionUpdatePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRealtimeTranscriptionSessionUpdatePayloadWithDefaults

`func NewRealtimeTranscriptionSessionUpdatePayloadWithDefaults() *RealtimeTranscriptionSessionUpdatePayload`

NewRealtimeTranscriptionSessionUpdatePayloadWithDefaults instantiates a new RealtimeTranscriptionSessionUpdatePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudioFormat

`func (o *RealtimeTranscriptionSessionUpdatePayload) GetAudioFormat() AudioFormat`

GetAudioFormat returns the AudioFormat field if non-nil, zero value otherwise.

### GetAudioFormatOk

`func (o *RealtimeTranscriptionSessionUpdatePayload) GetAudioFormatOk() (*AudioFormat, bool)`

GetAudioFormatOk returns a tuple with the AudioFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioFormat

`func (o *RealtimeTranscriptionSessionUpdatePayload) SetAudioFormat(v AudioFormat)`

SetAudioFormat sets AudioFormat field to given value.

### HasAudioFormat

`func (o *RealtimeTranscriptionSessionUpdatePayload) HasAudioFormat() bool`

HasAudioFormat returns a boolean if a field has been set.

### SetAudioFormatNil

`func (o *RealtimeTranscriptionSessionUpdatePayload) SetAudioFormatNil(b bool)`

 SetAudioFormatNil sets the value for AudioFormat to be an explicit nil

### UnsetAudioFormat
`func (o *RealtimeTranscriptionSessionUpdatePayload) UnsetAudioFormat()`

UnsetAudioFormat ensures that no value is present for AudioFormat, not even an explicit nil
### GetTargetStreamingDelayMs

`func (o *RealtimeTranscriptionSessionUpdatePayload) GetTargetStreamingDelayMs() int32`

GetTargetStreamingDelayMs returns the TargetStreamingDelayMs field if non-nil, zero value otherwise.

### GetTargetStreamingDelayMsOk

`func (o *RealtimeTranscriptionSessionUpdatePayload) GetTargetStreamingDelayMsOk() (*int32, bool)`

GetTargetStreamingDelayMsOk returns a tuple with the TargetStreamingDelayMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetStreamingDelayMs

`func (o *RealtimeTranscriptionSessionUpdatePayload) SetTargetStreamingDelayMs(v int32)`

SetTargetStreamingDelayMs sets TargetStreamingDelayMs field to given value.

### HasTargetStreamingDelayMs

`func (o *RealtimeTranscriptionSessionUpdatePayload) HasTargetStreamingDelayMs() bool`

HasTargetStreamingDelayMs returns a boolean if a field has been set.

### SetTargetStreamingDelayMsNil

`func (o *RealtimeTranscriptionSessionUpdatePayload) SetTargetStreamingDelayMsNil(b bool)`

 SetTargetStreamingDelayMsNil sets the value for TargetStreamingDelayMs to be an explicit nil

### UnsetTargetStreamingDelayMs
`func (o *RealtimeTranscriptionSessionUpdatePayload) UnsetTargetStreamingDelayMs()`

UnsetTargetStreamingDelayMs ensures that no value is present for TargetStreamingDelayMs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


