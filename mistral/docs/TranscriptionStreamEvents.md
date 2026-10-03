# TranscriptionStreamEvents

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | [**TranscriptionStreamEventTypes**](TranscriptionStreamEventTypes.md) |  | 
**Data** | [**Data2**](Data2.md) |  | 

## Methods

### NewTranscriptionStreamEvents

`func NewTranscriptionStreamEvents(event TranscriptionStreamEventTypes, data Data2, ) *TranscriptionStreamEvents`

NewTranscriptionStreamEvents instantiates a new TranscriptionStreamEvents object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTranscriptionStreamEventsWithDefaults

`func NewTranscriptionStreamEventsWithDefaults() *TranscriptionStreamEvents`

NewTranscriptionStreamEventsWithDefaults instantiates a new TranscriptionStreamEvents object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *TranscriptionStreamEvents) GetEvent() TranscriptionStreamEventTypes`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *TranscriptionStreamEvents) GetEventOk() (*TranscriptionStreamEventTypes, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *TranscriptionStreamEvents) SetEvent(v TranscriptionStreamEventTypes)`

SetEvent sets Event field to given value.


### GetData

`func (o *TranscriptionStreamEvents) GetData() Data2`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TranscriptionStreamEvents) GetDataOk() (*Data2, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TranscriptionStreamEvents) SetData(v Data2)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


