# RealtimeTranscriptionClientMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "session.update"]
**Session** | [**RealtimeTranscriptionSessionUpdatePayload**](RealtimeTranscriptionSessionUpdatePayload.md) |  | 
**Audio** | **string** | Base64-encoded raw PCM bytes matching the current audio_format. Max decoded size: 262144 bytes. | 

## Methods

### NewRealtimeTranscriptionClientMessage

`func NewRealtimeTranscriptionClientMessage(session RealtimeTranscriptionSessionUpdatePayload, audio string, ) *RealtimeTranscriptionClientMessage`

NewRealtimeTranscriptionClientMessage instantiates a new RealtimeTranscriptionClientMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRealtimeTranscriptionClientMessageWithDefaults

`func NewRealtimeTranscriptionClientMessageWithDefaults() *RealtimeTranscriptionClientMessage`

NewRealtimeTranscriptionClientMessageWithDefaults instantiates a new RealtimeTranscriptionClientMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RealtimeTranscriptionClientMessage) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RealtimeTranscriptionClientMessage) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RealtimeTranscriptionClientMessage) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RealtimeTranscriptionClientMessage) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSession

`func (o *RealtimeTranscriptionClientMessage) GetSession() RealtimeTranscriptionSessionUpdatePayload`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *RealtimeTranscriptionClientMessage) GetSessionOk() (*RealtimeTranscriptionSessionUpdatePayload, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *RealtimeTranscriptionClientMessage) SetSession(v RealtimeTranscriptionSessionUpdatePayload)`

SetSession sets Session field to given value.


### GetAudio

`func (o *RealtimeTranscriptionClientMessage) GetAudio() string`

GetAudio returns the Audio field if non-nil, zero value otherwise.

### GetAudioOk

`func (o *RealtimeTranscriptionClientMessage) GetAudioOk() (*string, bool)`

GetAudioOk returns a tuple with the Audio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudio

`func (o *RealtimeTranscriptionClientMessage) SetAudio(v string)`

SetAudio sets Audio field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


