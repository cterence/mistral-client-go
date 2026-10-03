# RealtimeTranscriptionSessionUpdateMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "session.update"]
**Session** | [**RealtimeTranscriptionSessionUpdatePayload**](RealtimeTranscriptionSessionUpdatePayload.md) |  | 

## Methods

### NewRealtimeTranscriptionSessionUpdateMessage

`func NewRealtimeTranscriptionSessionUpdateMessage(session RealtimeTranscriptionSessionUpdatePayload, ) *RealtimeTranscriptionSessionUpdateMessage`

NewRealtimeTranscriptionSessionUpdateMessage instantiates a new RealtimeTranscriptionSessionUpdateMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRealtimeTranscriptionSessionUpdateMessageWithDefaults

`func NewRealtimeTranscriptionSessionUpdateMessageWithDefaults() *RealtimeTranscriptionSessionUpdateMessage`

NewRealtimeTranscriptionSessionUpdateMessageWithDefaults instantiates a new RealtimeTranscriptionSessionUpdateMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RealtimeTranscriptionSessionUpdateMessage) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RealtimeTranscriptionSessionUpdateMessage) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RealtimeTranscriptionSessionUpdateMessage) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RealtimeTranscriptionSessionUpdateMessage) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSession

`func (o *RealtimeTranscriptionSessionUpdateMessage) GetSession() RealtimeTranscriptionSessionUpdatePayload`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *RealtimeTranscriptionSessionUpdateMessage) GetSessionOk() (*RealtimeTranscriptionSessionUpdatePayload, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *RealtimeTranscriptionSessionUpdateMessage) SetSession(v RealtimeTranscriptionSessionUpdatePayload)`

SetSession sets Session field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


